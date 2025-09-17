package service

/*
 Copyright 2019 - 2025 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/cql"
	"github.com/CrunchyData/pg_featureserv/internal/data"
)

func parseRequestParams(r *http.Request) (api.RequestParam, error) {
	queryValues := r.URL.Query()
	paramValues := extractSingleArgs(queryValues)

	param := api.RequestParam{
		Crs:       data.SRID_4326,
		Limit:     conf.Configuration.Paging.LimitDefault,
		Offset:    0,
		Precision: -1,
		BboxCrs:   data.SRID_4326,
		Filter:    "",
		Values:    paramValues,
	}

	// --- crs parameter
	crs, err := parseInt(paramValues, api.ParamCrs, 0, 99999999, data.SRID_4326)
	if err != nil {
		return param, err
	}
	param.Crs = crs

	// --- limit parameter
	limit, err := parseLimit(paramValues)
	if err != nil {
		return param, err
	}
	param.Limit = limit

	// --- offset parameter
	offset, err := parseInt(paramValues, api.ParamOffset, 0, -1, 0)
	if err != nil {
		return param, err
	}
	param.Offset = offset

	// --- bbox parameter
	bbox, err := parseBbox(paramValues)
	if err != nil {
		return param, err
	}
	param.Bbox = bbox

	// --- bbox-crs parameter
	bboxcrs, err := parseInt(paramValues, api.ParamBboxCrs, 0, 99999999, data.SRID_4326)
	if err != nil {
		return param, err
	}
	param.BboxCrs = bboxcrs

	// --- filter parameter
	param.Filter = parseString(paramValues, api.ParamFilter)

	// --- datetime parameter
	param.DateTime = parseString(paramValues, api.ParamDateTime)

	// --- filter-crs parameter
	filterCrs, err := parseInt(paramValues, api.ParamFilterCrs, 0, 99999999, data.SRID_4326)
	if err != nil {
		return param, err
	}
	param.FilterCrs = filterCrs

	// --- properties parameter
	props, err := parseProperties(paramValues)
	if err != nil {
		return param, err
	}
	param.Properties = props

	// --- orderBy parameter
	groupBy, err := parseGroupBy(paramValues)
	if err != nil {
		return param, err
	}
	param.GroupBy = groupBy

	// --- orderBy parameter (DEPRECATED)
	orderBy, err := parseOrderBy(paramValues)
	if err != nil {
		return param, err
	}
	param.SortBy = orderBy

	// --- sortBy parameter
	sortBy, err := parseSortBy(paramValues)
	if err != nil {
		return param, err
	}
	param.SortBy = sortBy

	// --- precision parameter
	precision, err := parseInt(paramValues, api.ParamPrecision, 0, 20, -1)
	if err != nil {
		return param, err
	}
	param.Precision = precision

	// --- transform parameter
	param.TransformFuns, err = parseTransform(paramValues)
	if err != nil {
		return param, err
	}

	return param, nil
}

func extractSingleArgs(queryArgs url.Values) api.NameValMap {
	vals := make(map[string]string)
	for keyRaw := range queryArgs {
		queryval := queryArgs.Get(keyRaw)
		key := strings.ToLower(keyRaw)
		vals[key] = queryval
	}
	return vals
}

func parseString(values api.NameValMap, key string) string {
	return strings.TrimSpace(values[key])
}

func parseInt(values api.NameValMap, key string, minVal int, maxVal int, defaultVal int) (int, error) {
	valStr := values[key]
	// key not present or missing value
	if len(valStr) < 1 {
		return defaultVal, nil
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, fmt.Errorf(api.ErrMsgInvalidParameterValue, key, valStr)
	}
	if val < minVal {
		val = minVal
	}
	if maxVal >= 0 && val > maxVal {
		val = maxVal
	}
	return val, nil
}

func parseLimit(values api.NameValMap) (int, error) {
	val := values[api.ParamLimit]
	if len(val) < 1 {
		return conf.Configuration.Paging.LimitDefault, nil
	}
	limit, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamLimit, val)
	}
	if limit < 0 || limit > conf.Configuration.Paging.LimitMax {
		limit = conf.Configuration.Paging.LimitMax
	}
	return limit, nil
}

/*
parseBbox parses the bbox query parameter, if present, or nll if not
This has the format bbox=minLon,minLat,maxLon,maxLat.
*/
func parseBbox(values api.NameValMap) (*data.Extent, error) {
	val := values[api.ParamBbox]
	if len(val) < 1 {
		return nil, nil
	}
	nums := strings.Split(val, ",")
	var isErr = false
	if len(nums) != 4 {
		return nil, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamBbox, val)
	}
	minLon, err := strconv.ParseFloat(nums[0], 64)
	if err != nil {
		isErr = true
	}
	minLat, err := strconv.ParseFloat(nums[1], 64)
	if err != nil {
		isErr = true
	}
	maxLon, err := strconv.ParseFloat(nums[2], 64)
	if err != nil {
		isErr = true
	}
	maxLat, err := strconv.ParseFloat(nums[3], 64)
	if err != nil {
		isErr = true
	}
	if isErr {
		return nil, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamBbox, val)
	}
	var bbox = data.Extent{Minx: minLon, Miny: minLat, Maxx: maxLon, Maxy: maxLat}
	return &bbox, nil
}

// parseProperties extracts an array of rawo property names to be included
// returns nil if no properties parameter was specified
// returns[] if properties is present but with no args
func parseProperties(values api.NameValMap) ([]string, error) {
	val, ok := values[api.ParamProperties]
	// no properties param => nil
	if !ok {
		return nil, nil
	}
	// empty properties list  => []
	if len(val) < 1 {
		return []string{}, nil
	}
	// return array of raw property names
	namesRaw := strings.Split(val, ",")
	return namesRaw, nil
}

func parseGroupBy(values api.NameValMap) ([]string, error) {
	val, ok := values[api.ParamGroupBy]
	// no properties param => nil
	if !ok {
		return nil, nil
	}
	// empty list  => []
	if !ok || len(val) < 1 {
		return []string{}, nil
	}
	// TODO: normalize col names
	// return array of raw col names
	namesRaw := strings.Split(val, ",")
	return namesRaw, nil
}

// parseSortBy determines an Sorting array
func parseSortBy(values api.NameValMap) ([]data.Sorting, error) {
	var sorting []data.Sorting
	val := values[api.ParamSortBy]
	if len(val) < 1 {
		return sorting, nil
	}
	valLow := strings.ToLower(val)
	sortCols := strings.Split(valLow, ",")
	sortCol := sortCols[0]
	isDesc := false
	name := strings.TrimSpace(sortCol)
	if strings.HasPrefix(name, "+") {
		name = strings.TrimSpace(name[1:])
		isDesc = false
	} else if strings.HasPrefix(name, "-") {
		name = strings.TrimSpace(name[1:])
		isDesc = true
	}
	sorting = append(sorting, data.Sorting{Name: name, IsDesc: isDesc})
	return sorting, nil
}

// parseOrderBy determines an order by array (DEPRECATED)
func parseOrderBy(values api.NameValMap) ([]data.Sorting, error) {
	var orderBy []data.Sorting
	val := values[api.ParamOrderBy]
	if len(val) < 1 {
		return orderBy, nil
	}
	valLow := strings.ToLower(val)
	nameDir := strings.Split(valLow, api.OrderByDirSep)
	name := nameDir[0]
	isDesc := false
	var err error
	if len(nameDir) >= 2 {
		dirSpec := nameDir[1]
		isDesc, err = parseOrderByDir(dirSpec)
		if err != nil {
			return nil, err
		}
	}
	orderBy = append(orderBy, data.Sorting{Name: name, IsDesc: isDesc})
	return orderBy, nil
}

func parseOrderByDir(dir string) (bool, error) {
	if dir == api.OrderByDirD {
		return true, nil
	}
	if dir == api.OrderByDirA {
		return false, nil
	}
	err := fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamOrderBy, dir)
	return false, err
}

// normalizePropNames converts the request property name list (if any)
// into a clean list of valid, unique column names
// If the request properties list is empty,
// the full column list is returned
func normalizePropNames(requestNames []string, colNames []string) []string {
	// no properties parameter => use all columns
	if requestNames == nil {
		return colNames
	}
	// empty properties parameter => use NO columns
	if len(requestNames) == 0 {
		return requestNames
	}
	nameSet := toNameSet(requestNames)
	// select cols which appear in set
	var propNames []string
	for _, colName := range colNames {
		if _, ok := nameSet[colName]; ok {
			propNames = append(propNames, colName)
		}
	}
	return propNames
}

func toNameSet(strs []string) map[string]bool {
	set := make(map[string]bool)
	for _, s := range strs {
		// keep case of property column names
		set[s] = true
	}
	return set
}

const (
	transformFunSep   = "|"
	transformParamSep = ","
	functionPrefixST  = "st_"
)

var transformFunctionWhitelist map[string]string

func initTransforms(funNames []string) {
	transformFunctionWhitelist = make(map[string]string)
	for _, name := range funNames {
		nameLow := strings.ToLower(name)
		transformFunctionWhitelist[nameLow] = name
	}
}

// actualFunctionName converts an input function name
// to an actual function name from the whitelist
func actualFunctionName(name string) string {
	nameLow := strings.ToLower(name)
	if actual, ok := transformFunctionWhitelist[nameLow]; ok {
		return actual
	}
	if !strings.HasPrefix(nameLow, functionPrefixST) {
		// supply ST_ prefix if not there and try again
		stName := functionPrefixST + nameLow
		if actual, ok := transformFunctionWhitelist[stName]; ok {
			return actual
		}
	}
	return ""
}

func parseTransform(values api.NameValMap) ([]data.TransformFunction, error) {
	val := values[api.ParamTransform]
	if len(val) < 1 {
		return nil, nil
	}
	funDefs := strings.Split(val, transformFunSep)

	funList := make([]data.TransformFunction, 0)
	for _, fun := range funDefs {
		tf := parseTransformFun(fun)
		actualName := actualFunctionName(tf.Name)
		if len(actualName) <= 0 {
			err := fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamTransform, tf.Name)
			return nil, err
		}
		tf.Name = actualName
		if tf.Name != "" {
			funList = append(funList, tf)
		}
	}
	return funList, nil
}

func parseTransformFun(def string) data.TransformFunction {
	// check for function parameter
	atoms := strings.Split(def, transformParamSep)
	name := atoms[0]
	args := atoms[1:]
	// TODO: harden this by checking arg is a valid number
	// TODO: have whitelist for function names?
	return data.TransformFunction{Name: name, Arg: args}
}

// parseFilter creates a filter list from applicable query parameters
func parseFilter(paramMap map[string]string, colNameMap map[string]string) []*data.PropertyFilter {
	var conds []*data.PropertyFilter
	for name, val := range paramMap {
		//log.Debugf("testing request param %v", name)
		if api.IsParameterReservedName(name) {
			continue
		}
		if _, ok := colNameMap[name]; ok {
			cond := &data.PropertyFilter{Name: name, Value: val}
			conds = append(conds, cond)
			//log.Debugf("Adding filter %v = %v ", name, val)
		}
	}
	return conds
}

// createQueryParams applies any cross-parameter logic
func createQueryParams(param *api.RequestParam, colNames []string, sourceSRID int) (*data.QueryParam, error) {
	query := data.QueryParam{
		Crs:           param.Crs,
		Limit:         param.Limit,
		Offset:        param.Offset,
		Bbox:          param.Bbox,
		BboxCrs:       param.BboxCrs,
		GroupBy:       param.GroupBy,
		SortBy:        param.SortBy,
		Precision:     param.Precision,
		TransformFuns: param.TransformFuns,
	}
	cols := param.Properties
	// --- if groupby is present it replaces properties (it may be empty)
	if param.GroupBy != nil {
		cols = param.GroupBy
		// ensure a aggregating transform is set to avoid error
		if len(param.TransformFuns) == 0 {
			query.TransformFuns = []data.TransformFunction{
				{
					Name: "st_collect",
				},
				{
					Name: "st_envelope",
				},
			}
		}
	}
	query.Columns = normalizePropNames(cols, colNames)
	//-- convert filter CQL
	sql, err := cql.TranspileToSQL(param.Filter, param.FilterCrs, sourceSRID)
	if err != nil {
		return &query, err
	}
	query.FilterSql = sql

	dtRange, err := parseDateTimeRange(param.DateTime)
	if err != nil {
		return &query, err
	}
	query.DateTime = dtRange

	return &query, nil
}

func parseDateTimeRange(value string) (*data.TimeRange, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return nil, nil
	}
	if !strings.Contains(trimmed, "/") {
		inst, err := parseDateTimeInstant(trimmed)
		if err != nil {
			return nil, err
		}
		if inst == nil {
			return nil, nil
		}
		rng := &data.TimeRange{StartInclusive: true, EndInclusive: true}
		start := copyTime(inst.Time)
		rng.Start = &start
		if inst.DateOnly {
			end := start.Add(24 * time.Hour)
			rng.End = &end
			rng.EndInclusive = false
		} else {
			end := copyTime(inst.Time)
			rng.End = &end
		}
		return rng, nil
	}
	parts := strings.SplitN(trimmed, "/", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamDateTime, value)
	}
	startInst, err := parseDateTimeInstant(parts[0])
	if err != nil {
		return nil, err
	}
	endInst, err := parseDateTimeInstant(parts[1])
	if err != nil {
		return nil, err
	}
	if startInst == nil && endInst == nil {
		return nil, nil
	}
	rng := &data.TimeRange{StartInclusive: true, EndInclusive: true}
	if startInst != nil {
		start := copyTime(startInst.Time)
		rng.Start = &start
	}
	if endInst != nil {
		end := copyTime(endInst.Time)
		rng.End = &end
	}
	if startInst != nil && startInst.DateOnly {
		// already normalized to midnight
		rng.StartInclusive = true
	}
	if endInst != nil && endInst.DateOnly {
		endAdj := rng.End.Add(24 * time.Hour)
		rng.End = &endAdj
		rng.EndInclusive = false
	}
	if rng.Start != nil && rng.End != nil {
		if rng.EndInclusive {
			if rng.Start.After(*rng.End) {
				return nil, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamDateTime, value)
			}
		} else {
			if !rng.Start.Before(*rng.End) {
				return nil, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamDateTime, value)
			}
		}
	}
	return rng, nil
}

type parsedInstant struct {
	Time     time.Time
	DateOnly bool
}

func parseDateTimeInstant(value string) (*parsedInstant, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" || trimmed == ".." {
		return nil, nil
	}
	tm, isDateOnly, err := parseDateTimeLiteral(trimmed)
	if err != nil {
		return nil, err
	}
	return &parsedInstant{Time: tm.UTC(), DateOnly: isDateOnly}, nil
}

func parseDateTimeLiteral(value string) (time.Time, bool, error) {
	layouts := []string{time.RFC3339Nano, time.RFC3339}
	for _, layout := range layouts {
		tm, err := time.Parse(layout, value)
		if err == nil {
			return tm.UTC(), false, nil
		}
	}
	nonZoneLayouts := []string{"2006-01-02T15:04:05", "2006-01-02T15:04"}
	for _, layout := range nonZoneLayouts {
		tm, err := time.ParseInLocation(layout, value, time.UTC)
		if err == nil {
			return tm.UTC(), false, nil
		}
	}
	tm, err := time.Parse("2006-01-02", value)
	if err == nil {
		return tm.UTC(), true, nil
	}
	return time.Time{}, false, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamDateTime, value)
}

func copyTime(t time.Time) time.Time {
	return t.UTC()
}
