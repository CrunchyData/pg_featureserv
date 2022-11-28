package api

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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

// Generic representation of a weak etag
type WeakEtagData struct {
	Collection   string                 `json:"collection"`
	FeatureId    string                 `json:"fid"`
	Etag         string                 `json:"etag"` // for example: xmin
	LastModified string                 `json:"last-modified"`
	Data         map[string]interface{} `json:"-"`
}

// Generic representation of a strong etag
type StrongEtagData struct {
	*WeakEtagData

	Srid   int    `json:"srid"`
	Format string `json:"format"`
}

// weak etag stringer function
func (etag WeakEtagData) String() string {
	return "W/\"" + etag.Etag + "\""
}

// return a key for a cache according to etag.
// If the etag field is not set, return the alternate key
func (etag WeakEtagData) CacheKey() string {
	if etag.Etag != "" {
		return fmt.Sprintf("ET-%s", etag.Etag)
	}
	return etag.AlternateCacheKey()
}

// return a key for a cache according to collection name and feature id.
func (etag WeakEtagData) AlternateCacheKey() string {
	return fmt.Sprintf("CF-%s-%s", etag.Collection, etag.FeatureId)
}

// Marshall etag to byte
func (etag WeakEtagData) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(etag)
	return bytes, err
}

// Unmarshall etag to byte
func (etag *WeakEtagData) UnmarshalBinary(data []byte) error {
	var tmp WeakEtagData
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	etag.Collection = tmp.Collection
	etag.FeatureId = tmp.FeatureId
	etag.Etag = tmp.Etag
	etag.LastModified = tmp.LastModified
	etag.Data = tmp.Data

	return nil
}

// return a new pointer onto weak etag
func MakeWeakEtag(collection string, fid string, weakEtag string, lastModified string) *WeakEtagData {
	weakEtagData := WeakEtagData{
		Collection:   collection,
		FeatureId:    fid,
		Etag:         weakEtag,
		LastModified: lastModified,
	}
	return &weakEtagData
}

// strong etag stringer function
func (sEtag StrongEtagData) String() string {
	return fmt.Sprintf("%s-%s-%d-%s-%s", sEtag.WeakEtagData.Collection, sEtag.WeakEtagData.FeatureId,
		sEtag.Srid, sEtag.Format, sEtag.WeakEtagData.Etag)
}

// base64 encoded version of String() function
func (sEtag StrongEtagData) ToEncodedString() string {
	return base64.StdEncoding.EncodeToString([]byte(sEtag.String()))
}

// returns a key for a cache
func (sEtag StrongEtagData) CacheKey() string {
	return sEtag.WeakEtagData.CacheKey()
}

// Marshall etag to byte
func (etag StrongEtagData) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(etag)
	return bytes, err
}

// return a new pointer onto strong etag
func MakeStrongEtag(collection string, fid string, weakEtag string, lastModified string, srid int, format string) *StrongEtagData {
	weakEtagData := MakeWeakEtag(collection, fid, weakEtag, lastModified)
	strongEtagData := StrongEtagData{
		WeakEtagData: weakEtagData,
		Srid:         srid,
		Format:       format,
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
	if len(elements) != 5 {
		return nil, errors.New("strong etag contains a wrong number of elements")
	}
	collectionName := elements[0]
	sridValue, err := strconv.Atoi(elements[2])
	if err != nil {
		return nil, errors.New("the provided srid value is not an int")
	}
	fid, format, etag := elements[1], elements[3], elements[4]

	return MakeStrongEtag(collectionName, fid, etag, "", sridValue, format), nil
}

// Returns a WeakEtag string from a etag. If etag starts with W/ value is split to get WeakEtag otherwise we decode etag to strong format to get WeakEtag
func EtagStrToObject(etagStr string) (*WeakEtagData, error) {

	if etagStr == "" {
		return nil, nil
	}

	// Weak Etag form
	if strings.HasPrefix(etagStr, "W/") {
		etagStr := strings.Split(etagStr, "W/")[1]
		etagStr = strings.Replace(etagStr, "\"", "", -1)
		weakEtag := MakeWeakEtag("", "", etagStr, "")
		return weakEtag, nil

	} else {
		// Strong Etag form
		strongEtag, err := DecodeStrongEtag(etagStr)
		if err != nil {
			return nil, errors.New("wrong strong etag format")
		}

		return strongEtag.WeakEtagData, nil
	}

}
