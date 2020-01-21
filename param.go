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
	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
)

func parseRequestParams(r *http.Request) (data.QueryParam, error) {
	queryValues := r.URL.Query()
	param := data.QueryParam{
		Limit:     config.Configuration.Paging.LimitDefault,
		Offset:    0,
		Precision: 4,
		Values:    extractSingleArg(queryValues),
	}

	// --- limit parameter
	limit, err := parseLimit(queryValues)
	if err != nil {
		return param, err
	}
	param.Limit = limit

	// --- limit parameter
	offset, err := parseOffset(queryValues)
	if err != nil {
		return param, err
	}
	param.Offset = offset

	// --- bbox parameter
	bbox, err := parseBbox(queryValues)
	if err != nil {
		return param, err
	}
	param.Bbox = bbox

	// --- precision parameter
	precision, err := parsePrecision(queryValues)
	if err != nil {
		return param, err
	}
	param.Precision = precision

	// --- transform parameter
	param.TransformFuns = parseTransform(queryValues, 0)

	return param, nil
}

func extractSingleArg(queryArgs url.Values) map[string]string {
	vals := make(map[string]string)
	for key := range queryArgs {
		queryval := queryArgs.Get(key)
		vals[key] = queryval
	}
	return vals
}

func parseLimit(values url.Values) (int, error) {
	val := values.Get(api.ParamLimit)
	if len(val) < 1 {
		return config.Configuration.Paging.LimitDefault, nil
	}
	limit, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamLimit, val)
	}
	if limit < 0 || limit > config.Configuration.Paging.LimitMax {
		limit = config.Configuration.Paging.LimitMax
	}
	return limit, nil
}

func parseOffset(values url.Values) (int, error) {
	val := values.Get(api.ParamOffset)
	if len(val) < 1 {
		return 0, nil
	}
	offset, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamLimit, val)
	}
	if offset < 0 {
		offset = 0
	}
	if offset > config.Configuration.Paging.LimitMax {
		offset = config.Configuration.Paging.LimitMax
	}
	return offset, nil
}

func parsePrecision(values url.Values) (int, error) {
	val := values.Get(api.ParamPrecision)
	// no parameter present
	if len(val) < 1 {
		return -1, nil
	}
	prec, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf(api.ErrMsgInvalidParameterValue, api.ParamPrecision, val)
	}
	return prec, nil
}

/*
parseBbox parses the bbox query parameter, if present, or nll if not
This has the format bbox=minLon,minLat,maxLon,maxLat.
*/
func parseBbox(values url.Values) (*data.Extent, error) {
	val := values.Get(api.ParamBbox)
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

const transformFunSep = "|"
const transformParamSep = ","

func parseTransform(values url.Values, index int) []data.TransformFunction {
	val := values.Get(api.ParamTransform)
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
