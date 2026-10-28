// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docspace_api_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BatchTagsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchTagsRequestDto{}

// BatchTagsRequestDto The tag names a request attaches to a room or detaches from it.
type BatchTagsRequestDto struct {
	// The tags, by name: a tag has no identifier of its own, and the name is what links a room to it.  `GET api/2.0/files/tags` lists the names already in the portal catalogue. An empty list is accepted and does  nothing, while a blank or overlong entry makes the whole request invalid.
	Names []string `json:"names"`
}

type _BatchTagsRequestDto BatchTagsRequestDto

// NewBatchTagsRequestDto instantiates a new BatchTagsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchTagsRequestDto(names []string) *BatchTagsRequestDto {
	this := BatchTagsRequestDto{}
	this.Names = names
	return &this
}

// NewBatchTagsRequestDtoWithDefaults instantiates a new BatchTagsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchTagsRequestDtoWithDefaults() *BatchTagsRequestDto {
	this := BatchTagsRequestDto{}
	return &this
}

// GetNames returns the Names field value
func (o *BatchTagsRequestDto) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *BatchTagsRequestDto) GetNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Names, true
}

// SetNames sets field value
func (o *BatchTagsRequestDto) SetNames(v []string) {
	o.Names = v
}

func (o BatchTagsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchTagsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["names"] = o.Names
	return toSerialize, nil
}

func (o *BatchTagsRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"names",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBatchTagsRequestDto := _BatchTagsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varBatchTagsRequestDto)

	if err != nil {
		return err
	}

	*o = BatchTagsRequestDto(varBatchTagsRequestDto)

	return err
}

type NullableBatchTagsRequestDto struct {
	value *BatchTagsRequestDto
	isSet bool
}

func (v NullableBatchTagsRequestDto) Get() *BatchTagsRequestDto {
	return v.value
}

func (v *NullableBatchTagsRequestDto) Set(val *BatchTagsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchTagsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchTagsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchTagsRequestDto(val *BatchTagsRequestDto) *NullableBatchTagsRequestDto {
	return &NullableBatchTagsRequestDto{value: val, isSet: true}
}

func (v NullableBatchTagsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchTagsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

