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

// checks if the RemoveProviderRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveProviderRequestDto{}

// RemoveProviderRequestDto Request parameters for deleting one or more AI providers.
type RemoveProviderRequestDto struct {
	// The set of AI provider identifiers to delete.
	Ids []int32 `json:"ids"`
}

type _RemoveProviderRequestDto RemoveProviderRequestDto

// NewRemoveProviderRequestDto instantiates a new RemoveProviderRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveProviderRequestDto(ids []int32) *RemoveProviderRequestDto {
	this := RemoveProviderRequestDto{}
	this.Ids = ids
	return &this
}

// NewRemoveProviderRequestDtoWithDefaults instantiates a new RemoveProviderRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveProviderRequestDtoWithDefaults() *RemoveProviderRequestDto {
	this := RemoveProviderRequestDto{}
	return &this
}

// GetIds returns the Ids field value
// If the value is explicit nil, the zero value for []int32 will be returned
func (o *RemoveProviderRequestDto) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoveProviderRequestDto) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *RemoveProviderRequestDto) SetIds(v []int32) {
	o.Ids = v
}

func (o RemoveProviderRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveProviderRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

func (o *RemoveProviderRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
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

	varRemoveProviderRequestDto := _RemoveProviderRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRemoveProviderRequestDto)

	if err != nil {
		return err
	}

	*o = RemoveProviderRequestDto(varRemoveProviderRequestDto)

	return err
}

type NullableRemoveProviderRequestDto struct {
	value *RemoveProviderRequestDto
	isSet bool
}

func (v NullableRemoveProviderRequestDto) Get() *RemoveProviderRequestDto {
	return v.value
}

func (v *NullableRemoveProviderRequestDto) Set(val *RemoveProviderRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveProviderRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveProviderRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveProviderRequestDto(val *RemoveProviderRequestDto) *NullableRemoveProviderRequestDto {
	return &NullableRemoveProviderRequestDto{value: val, isSet: true}
}

func (v NullableRemoveProviderRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveProviderRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

