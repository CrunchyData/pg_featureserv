package api

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
)

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

 Date     : October 2022
 Authors  : Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

// Generic representation of a strong etag
type StrongEtagData struct {
	Collection string `json:"collection"`
	Srid       int    `json:"srid"`
	Format     string `json:"format"`
	WeakEtag   string `json:"weaketag"`
}

func MakeStrongEtag(collection string, srid int, format string, weakEtag string) *StrongEtagData {
	strongEtagData := StrongEtagData{
		Collection: collection,
		Srid:       srid,
		Format:     format,
		WeakEtag:   weakEtag,
	}
	return &strongEtagData
}

// Returns a decoded structure of the raw Base64 strong etag provided as an argument
func DecodeStrongEtag(encodedStrongEtag string) (*StrongEtagData, error) {
	encodedStrongEtag = strings.ReplaceAll(encodedStrongEtag, "\"", "")

	decodedStrongEtag, err := base64.StdEncoding.DecodeString(encodedStrongEtag)
	if err != nil {
		return nil, errors.New("string has to be Base64 encoded")
	}
	decodedString := string(decodedStrongEtag)
	decodedString = strings.Replace(decodedString, "\"", "", -1)
	elements := strings.Split(decodedString, "-")
	if len(elements) != 4 {
		return nil, errors.New("strong etag contains a wrong number of elements")
	}
	collectionName := elements[0]
	sridValue, err := strconv.Atoi(elements[1])
	if err != nil {
		return nil, errors.New("the provided srid value is not an int")
	}
	format, weakEtag := elements[2], elements[3]
	return MakeStrongEtag(collectionName, sridValue, format, weakEtag), nil
}

func EtagToWeakEtag(strongEtag string) (string, error) {

	weakEtagValue := ""

	// Weak Etag form
	if strings.HasPrefix(strongEtag, "W/") {
		weakEtagValue = strings.Split(strongEtag, "W/")[1]
	} else {
		// Strong Etag form
		strongValue, err := DecodeStrongEtag(strongEtag)
		if err != nil {
			return "", errors.New("wrong strong etag format")
		}
		weakEtagValue = strongValue.WeakEtag
	}

	weakEtagValue = strings.ReplaceAll(weakEtagValue, "\"", "")
	return weakEtagValue, nil
}
