package api

/*
 Copyright 2022 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

 Date     : September 2022
 Authors  : Benoit De Mezzo (benoit dot de dot mezzo at oslandia dot com)
*/

// ==================================================
// ================== FunctionInfo ==================

type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Default     string `json:"default,omitempty"`
}

// FunctionInfo is the API metadata for a function
type FunctionInfo struct {
	Name        string `json:"id"`
	Description string `json:"description,omitempty"`

	// these properties are always present but may be empty arrays
	Parameters []*Parameter `json:"parameters"`
	Properties []*Property  `json:"properties"`

	Links []*Link `json:"links"`

	//--- additional data used during processing
	Function *Function `json:"-"`
}

// ===================================================
// ================== FunctionsInfo ==================

// FunctionsInfo is the API metadata for all functions
type FunctionsInfo struct {
	Links     []*Link            `json:"links"`
	Functions []*FunctionSummary `json:"functions"`
}

// FunctionSummary contains a restricted set of function metadata for use in list display and JSON
// This allows not including parameters and properties in list metadata,
// but ensuring those keys are always present in full metadata JSON.
// Note: Collections do not follow same pattern because their list JSON metadata
// is supposed to contain all properties, and they are always expected to have attribute properties
type FunctionSummary struct {
	Name        string  `json:"id"`
	Description string  `json:"description,omitempty"`
	Links       []*Link `json:"links"`

	//--- additional data used during processing
	Function *Function `json:"-"`
	// used for HTML response only
	URLMetadataHTML string `json:"-"`
	URLMetadataJSON string `json:"-"`
	URLItemsHTML    string `json:"-"`
	URLItemsJSON    string `json:"-"`
}

func NewFunctionsInfo(fns []*Function) *FunctionsInfo {
	fnsDoc := FunctionsInfo{Links: []*Link{}, Functions: []*FunctionSummary{}}
	for _, fn := range fns {
		fnDoc := fn.NewFunctionSummary()
		fnsDoc.Functions = append(fnsDoc.Functions, fnDoc)
	}
	return &fnsDoc
}

// ==============================================
// ================== Function ==================

// Function tbd
type Function struct {
	ID             string
	Schema         string
	Name           string
	Description    string
	InNames        []string
	InDbTypes      []string
	InTypeMap      map[string]PGType
	InDefaults     []string
	NumNoDefault   int
	OutNames       []string
	OutDbTypes     []string
	OutJSONTypes   []JSONType
	Types          map[string]PGType
	GeometryColumn string
	IDColumn       string
}

func (fun *Function) IsGeometryFunction() bool {
	for _, typ := range fun.OutDbTypes {
		if typ == "geometry" {
			return true
		}
	}
	return false
}

func (fn *Function) NewFunctionSummary() *FunctionSummary {
	info := FunctionSummary{
		Name:        fn.ID,
		Description: fn.Description,
		Function:    fn,
	}
	return &info
}

func (fn *Function) NewFunctionInfo() *FunctionInfo {
	info := FunctionInfo{
		Name:        fn.ID,
		Description: fn.Description,
		Function:    fn,
	}
	return &info
}

func (fn *Function) FunctionParameters() []*Parameter {
	params := make([]*Parameter, len(fn.InNames))
	for i, name := range fn.InNames {
		params[i] = &Parameter{
			Name: name,
			Type: fn.InDbTypes[i],
			// no description available from db catalog
			Default: fn.InDefaults[i],
		}
	}
	return params
}

func (fn *Function) FunctionProperties() []*Property {
	props := make([]*Property, len(fn.OutNames))
	for i, name := range fn.OutNames {
		props[i] = &Property{
			Name: name,
			Type: string(fn.OutJSONTypes[i]),
			// no description available from db catalog
		}
	}
	return props
}
