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

// checks if the UpdateMemberCultureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMemberCultureRequest{}

// UpdateMemberCultureRequest The culture name parameters.
type UpdateMemberCultureRequest struct {
	// The user culture name (en-US, de, fr, es, ...).
	CultureName string `json:"cultureName"`
}

type _UpdateMemberCultureRequest UpdateMemberCultureRequest

// NewUpdateMemberCultureRequest instantiates a new UpdateMemberCultureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMemberCultureRequest(cultureName string) *UpdateMemberCultureRequest {
	this := UpdateMemberCultureRequest{}
	this.CultureName = cultureName
	return &this
}

// NewUpdateMemberCultureRequestWithDefaults instantiates a new UpdateMemberCultureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMemberCultureRequestWithDefaults() *UpdateMemberCultureRequest {
	this := UpdateMemberCultureRequest{}
	return &this
}

// GetCultureName returns the CultureName field value
func (o *UpdateMemberCultureRequest) GetCultureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CultureName
}

// GetCultureNameOk returns a tuple with the CultureName field value
// and a boolean to check if the value has been set.
func (o *UpdateMemberCultureRequest) GetCultureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CultureName, true
}

// SetCultureName sets field value
func (o *UpdateMemberCultureRequest) SetCultureName(v string) {
	o.CultureName = v
}

func (o UpdateMemberCultureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMemberCultureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cultureName"] = o.CultureName
	return toSerialize, nil
}

func (o *UpdateMemberCultureRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cultureName",
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

	varUpdateMemberCultureRequest := _UpdateMemberCultureRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varUpdateMemberCultureRequest)

	if err != nil {
		return err
	}

	*o = UpdateMemberCultureRequest(varUpdateMemberCultureRequest)

	return err
}

type NullableUpdateMemberCultureRequest struct {
	value *UpdateMemberCultureRequest
	isSet bool
}

func (v NullableUpdateMemberCultureRequest) Get() *UpdateMemberCultureRequest {
	return v.value
}

func (v *NullableUpdateMemberCultureRequest) Set(val *UpdateMemberCultureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMemberCultureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMemberCultureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMemberCultureRequest(val *UpdateMemberCultureRequest) *NullableUpdateMemberCultureRequest {
	return &NullableUpdateMemberCultureRequest{value: val, isSet: true}
}

func (v NullableUpdateMemberCultureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMemberCultureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

