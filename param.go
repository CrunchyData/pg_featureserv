package main

/*
 Copyright 2019 Crunchy Data Solutions, Inc.
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

	"github.com/CrunchyData/pg_featureserv/api"
	"github.com/CrunchyData/pg_featureserv/conf"
	"github.com/CrunchyData/pg_featureserv/data"
)

func parseRequestParams(r *http.Request) (api.RequestParam, error) {
	queryValues := r.URL.Query()
	paramValues := extractSingleArgs(queryValues)

	param := api.RequestParam{
		Limit:     conf.Configuration.Paging.LimitDefault,
		Offset:    0,
		Precision: -1,
		Values:    paramValues,
	}

	// --- limit parameter
	limit, err := parseLimit(paramValues)
	if err != nil {
		return param, err
	}
	param.Limit = limit

	// --- offset parameter
	offset, err := parseInt(paramValues, api.ParamOffset, 0, conf.Configuration.Paging.LimitMax, 0)
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

	// --- properties parameter
	props, err := parseProperties(paramValues)
	if err != nil {
		return param, err
	}
	param.Properties = props

	// --- orderBy parameter
	orderBy, err := parseOrderBy(paramValues)
	if err != nil {
		return param, err
	}
	param.OrderBy = orderBy

	// --- precision parameter
	precision, err := parseInt(paramValues, api.ParamPrecision, 0, 20, -1)
	if err != nil {
		return param, err
	}
	param.Precision = precision

	// --- transform parameter
	param.TransformFuns = parseTransform(paramValues)

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
	if val > maxVal {
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

// parseProperties computes a lower-case, unique list
// of property names to be returned
func parseProperties(values api.NameValMap) ([]string, error) {
	val := values[api.ParamProperties]
	if len(val) < 1 {
		return nil, nil
	}
	namesRaw := strings.Split(val, ",")
	var names []string
	nameMap := make(map[string]bool)
	for _, name := range namesRaw {
		nameLow := strings.ToLower(name)
		// if a new name add to list
		if _, ok := nameMap[nameLow]; !ok {
			names = append(names, nameLow)
			nameMap[nameLow] = true
		}
	}
	return names, nil
}

const OrderByDirSep = ":"
const OrderByDirD = "d"
const OrderByDirA = "a"

// parseOrderBy determines an order by array
func parseOrderBy(values api.NameValMap) ([]data.Ordering, error) {
	var orderBy []data.Ordering
	val := values[api.ParamOrderBy]
	if len(val) < 1 {
		return orderBy, nil
	}
	valLow := strings.ToLower(val)
	nameDir := strings.Split(valLow, OrderByDirSep)
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
	orderBy = append(orderBy, data.Ordering{Name: name, IsDesc: isDesc})
	return orderBy, nil
}

func parseOrderByDir(dir string) (bool, error) {
	if dir == OrderByDirD {
		return true, nil
	}
	if dir == OrderByDirA {
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
	// no props given => use all properties
	if len(requestNames) == 0 {
		return colNames
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
		sLow := strings.ToLower(s)
		set[sLow] = true
	}
	return set
}

const transformFunSep = "|"
const transformParamSep = ","

func parseTransform(values api.NameValMap) []data.TransformFunction {
	val := values[api.ParamTransform]
	if len(val) < 1 {
		return nil
	}
	funDefs := strings.Split(val, transformFunSep)

	funList := make([]data.TransformFunction, 0)
	for _, fun := range funDefs {
		tf := parseTransformFun(fun)
		if tf.Name != "" {
			funList = append(funList, tf)
		}
	}
	return funList
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
