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

// checks if the SetRestrictedAiModelsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetRestrictedAiModelsRequestDto{}

// SetRestrictedAiModelsRequestDto The complete set of AI chat models that are to be barred on the portal.
type SetRestrictedAiModelsRequestDto struct {
	// The identifiers of the models no user of the portal may pick, taken from  `GET api/2.0/portal/payment/ai-prices`. This is the whole set that is to hold afterwards and not a list of  additions: send the models already barred together with the new one to add a restriction, leave one out to  lift it, and send an empty set to lift them all.
	Models []string `json:"models"`
}

type _SetRestrictedAiModelsRequestDto SetRestrictedAiModelsRequestDto

// NewSetRestrictedAiModelsRequestDto instantiates a new SetRestrictedAiModelsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetRestrictedAiModelsRequestDto(models []string) *SetRestrictedAiModelsRequestDto {
	this := SetRestrictedAiModelsRequestDto{}
	this.Models = models
	return &this
}

// NewSetRestrictedAiModelsRequestDtoWithDefaults instantiates a new SetRestrictedAiModelsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetRestrictedAiModelsRequestDtoWithDefaults() *SetRestrictedAiModelsRequestDto {
	this := SetRestrictedAiModelsRequestDto{}
	return &this
}

// GetModels returns the Models field value
func (o *SetRestrictedAiModelsRequestDto) GetModels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *SetRestrictedAiModelsRequestDto) GetModelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *SetRestrictedAiModelsRequestDto) SetModels(v []string) {
	o.Models = v
}

func (o SetRestrictedAiModelsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetRestrictedAiModelsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["models"] = o.Models
	return toSerialize, nil
}

func (o *SetRestrictedAiModelsRequestDto) UnmarshalJSON(data []byte) (err error) {
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

	varSetRestrictedAiModelsRequestDto := _SetRestrictedAiModelsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSetRestrictedAiModelsRequestDto)

	if err != nil {
		return err
	}

	*o = SetRestrictedAiModelsRequestDto(varSetRestrictedAiModelsRequestDto)

	return err
}

type NullableSetRestrictedAiModelsRequestDto struct {
	value *SetRestrictedAiModelsRequestDto
	isSet bool
}

func (v NullableSetRestrictedAiModelsRequestDto) Get() *SetRestrictedAiModelsRequestDto {
	return v.value
}

func (v *NullableSetRestrictedAiModelsRequestDto) Set(val *SetRestrictedAiModelsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetRestrictedAiModelsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetRestrictedAiModelsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetRestrictedAiModelsRequestDto(val *SetRestrictedAiModelsRequestDto) *NullableSetRestrictedAiModelsRequestDto {
	return &NullableSetRestrictedAiModelsRequestDto{value: val, isSet: true}
}

func (v NullableSetRestrictedAiModelsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetRestrictedAiModelsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

