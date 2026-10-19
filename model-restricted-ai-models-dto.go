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

// checks if the RestrictedAiModelsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestrictedAiModelsDto{}

// RestrictedAiModelsDto The AI models the portal is not allowed to use.
type RestrictedAiModelsDto struct {
	// The identifiers of the models the portal is not allowed to use.
	Models []string `json:"models"`
}

type _RestrictedAiModelsDto RestrictedAiModelsDto

// NewRestrictedAiModelsDto instantiates a new RestrictedAiModelsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestrictedAiModelsDto(models []string) *RestrictedAiModelsDto {
	this := RestrictedAiModelsDto{}
	this.Models = models
	return &this
}

// NewRestrictedAiModelsDtoWithDefaults instantiates a new RestrictedAiModelsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictedAiModelsDtoWithDefaults() *RestrictedAiModelsDto {
	this := RestrictedAiModelsDto{}
	return &this
}

// GetModels returns the Models field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *RestrictedAiModelsDto) GetModels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestrictedAiModelsDto) GetModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *RestrictedAiModelsDto) SetModels(v []string) {
	o.Models = v
}

func (o RestrictedAiModelsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestrictedAiModelsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Models != nil {
		toSerialize["models"] = o.Models
	}
	return toSerialize, nil
}

func (o *RestrictedAiModelsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"models",
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

	varRestrictedAiModelsDto := _RestrictedAiModelsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRestrictedAiModelsDto)

	if err != nil {
		return err
	}

	*o = RestrictedAiModelsDto(varRestrictedAiModelsDto)

	return err
}

type NullableRestrictedAiModelsDto struct {
	value *RestrictedAiModelsDto
	isSet bool
}

func (v NullableRestrictedAiModelsDto) Get() *RestrictedAiModelsDto {
	return v.value
}

func (v *NullableRestrictedAiModelsDto) Set(val *RestrictedAiModelsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictedAiModelsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictedAiModelsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictedAiModelsDto(val *RestrictedAiModelsDto) *NullableRestrictedAiModelsDto {
	return &NullableRestrictedAiModelsDto{value: val, isSet: true}
}

func (v NullableRestrictedAiModelsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictedAiModelsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

