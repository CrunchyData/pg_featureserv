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

type CacheDisabled struct {
	// no-ope
}

func (cache CacheDisabled) ContainsWeakEtag(strongEtag string) (bool, error) {
	return false, nil
}

func (cache CacheDisabled) AddWeakEtag(weakEtag string, etag interface{}) (bool, error) {
	return false, nil
}

func (cache CacheDisabled) RemoveWeakEtag(weakEtag string) (bool, error) {
	return false, nil
}

func (cache CacheDisabled) String() string {
	return ""
}

func (cache CacheDisabled) Type() string {
	return "CacheDisabled"
}

func (cache CacheDisabled) Size() int {
	return 0
}
