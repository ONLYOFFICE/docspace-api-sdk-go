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

// checks if the UserExistsResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserExistsResponseDto{}

// UserExistsResponseDto The user existence check response parameters.
type UserExistsResponseDto struct {
	// Specifies whether the user exists or not.
	Exists bool `json:"exists"`
	Status *EmployeeStatus `json:"status,omitempty"`
}

type _UserExistsResponseDto UserExistsResponseDto

// NewUserExistsResponseDto instantiates a new UserExistsResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserExistsResponseDto(exists bool) *UserExistsResponseDto {
	this := UserExistsResponseDto{}
	this.Exists = exists
	return &this
}

// NewUserExistsResponseDtoWithDefaults instantiates a new UserExistsResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserExistsResponseDtoWithDefaults() *UserExistsResponseDto {
	this := UserExistsResponseDto{}
	return &this
}

// GetExists returns the Exists field value
func (o *UserExistsResponseDto) GetExists() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exists
}

// GetExistsOk returns a tuple with the Exists field value
// and a boolean to check if the value has been set.
func (o *UserExistsResponseDto) GetExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exists, true
}

// SetExists sets field value
func (o *UserExistsResponseDto) SetExists(v bool) {
	o.Exists = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserExistsResponseDto) GetStatus() EmployeeStatus {
	if o == nil || IsNil(o.Status) {
		var ret EmployeeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExistsResponseDto) GetStatusOk() (*EmployeeStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserExistsResponseDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EmployeeStatus and assigns it to the Status field.
func (o *UserExistsResponseDto) SetStatus(v EmployeeStatus) {
	o.Status = &v
}

func (o UserExistsResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserExistsResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exists"] = o.Exists
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *UserExistsResponseDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exists",
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

	varUserExistsResponseDto := _UserExistsResponseDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varUserExistsResponseDto)

	if err != nil {
		return err
	}

	*o = UserExistsResponseDto(varUserExistsResponseDto)

	return err
}

type NullableUserExistsResponseDto struct {
	value *UserExistsResponseDto
	isSet bool
}

func (v NullableUserExistsResponseDto) Get() *UserExistsResponseDto {
	return v.value
}

func (v *NullableUserExistsResponseDto) Set(val *UserExistsResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserExistsResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserExistsResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserExistsResponseDto(val *UserExistsResponseDto) *NullableUserExistsResponseDto {
	return &NullableUserExistsResponseDto{value: val, isSet: true}
}

func (v NullableUserExistsResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserExistsResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

