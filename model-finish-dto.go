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

// checks if the FinishDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinishDto{}

// FinishDto The parameters for terminating a process or operation.
type FinishDto struct {
	// Specifies whether to send a welcome email or not.
	IsSendWelcomeEmail bool `json:"isSendWelcomeEmail"`
}

type _FinishDto FinishDto

// NewFinishDto instantiates a new FinishDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinishDto(isSendWelcomeEmail bool) *FinishDto {
	this := FinishDto{}
	this.IsSendWelcomeEmail = isSendWelcomeEmail
	return &this
}

// NewFinishDtoWithDefaults instantiates a new FinishDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinishDtoWithDefaults() *FinishDto {
	this := FinishDto{}
	return &this
}

// GetIsSendWelcomeEmail returns the IsSendWelcomeEmail field value
func (o *FinishDto) GetIsSendWelcomeEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSendWelcomeEmail
}

// GetIsSendWelcomeEmailOk returns a tuple with the IsSendWelcomeEmail field value
// and a boolean to check if the value has been set.
func (o *FinishDto) GetIsSendWelcomeEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSendWelcomeEmail, true
}

// SetIsSendWelcomeEmail sets field value
func (o *FinishDto) SetIsSendWelcomeEmail(v bool) {
	o.IsSendWelcomeEmail = v
}

func (o FinishDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinishDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isSendWelcomeEmail"] = o.IsSendWelcomeEmail
	return toSerialize, nil
}

func (o *FinishDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isSendWelcomeEmail",
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

	varFinishDto := _FinishDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFinishDto)

	if err != nil {
		return err
	}

	*o = FinishDto(varFinishDto)

	return err
}

type NullableFinishDto struct {
	value *FinishDto
	isSet bool
}

func (v NullableFinishDto) Get() *FinishDto {
	return v.value
}

func (v *NullableFinishDto) Set(val *FinishDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFinishDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFinishDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinishDto(val *FinishDto) *NullableFinishDto {
	return &NullableFinishDto{value: val, isSet: true}
}

func (v NullableFinishDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinishDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

