package data

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

type Cacher interface {

	// returns true if the weak etag contained into the strong etag in argument is referenced into the cache
	// returns false otherwise
	// an error will be returned if a malformed etag is detected
	ContainsWeakEtag(strongEtag string) (bool, error)

	// adds the weak etag string into the cache and returns true if successful
	// returns false if error occurs during the operation
	AddWeakEtag(weakEtag string, etag interface{}) bool

	// returns a string representation of the cache for dev purpose
	ToString() string

	// returns true if the cache is activated, false otherwise
	IsCacheActive() bool
}
