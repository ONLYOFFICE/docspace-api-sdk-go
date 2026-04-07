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

// checks if the SignupAccountRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignupAccountRequestDto{}

// SignupAccountRequestDto The request parameters for creating a third-party account.
type SignupAccountRequestDto struct {
	EmployeeType *EmployeeType `json:"employeeType,omitempty"`
	// The user link key.
	Key NullableString `json:"key"`
	// The user culture code.
	Culture NullableString `json:"culture,omitempty"`
	// The third-party profile in the serialized format
	SerializedProfile NullableString `json:"serializedProfile"`
}

type _SignupAccountRequestDto SignupAccountRequestDto

// NewSignupAccountRequestDto instantiates a new SignupAccountRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignupAccountRequestDto(key NullableString, serializedProfile NullableString) *SignupAccountRequestDto {
	this := SignupAccountRequestDto{}
	this.Key = key
	this.SerializedProfile = serializedProfile
	return &this
}

// NewSignupAccountRequestDtoWithDefaults instantiates a new SignupAccountRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignupAccountRequestDtoWithDefaults() *SignupAccountRequestDto {
	this := SignupAccountRequestDto{}
	return &this
}

// GetEmployeeType returns the EmployeeType field value if set, zero value otherwise.
func (o *SignupAccountRequestDto) GetEmployeeType() EmployeeType {
	if o == nil || IsNil(o.EmployeeType) {
		var ret EmployeeType
		return ret
	}
	return *o.EmployeeType
}

// GetEmployeeTypeOk returns a tuple with the EmployeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupAccountRequestDto) GetEmployeeTypeOk() (*EmployeeType, bool) {
	if o == nil || IsNil(o.EmployeeType) {
		return nil, false
	}
	return o.EmployeeType, true
}

// HasEmployeeType returns a boolean if a field has been set.
func (o *SignupAccountRequestDto) IsEmployeeTypeSet() bool {
	if o != nil && !IsNil(o.EmployeeType) {
		return true
	}

	return false
}

// SetEmployeeType gets a reference to the given EmployeeType and assigns it to the EmployeeType field.
func (o *SignupAccountRequestDto) SetEmployeeType(v EmployeeType) {
	o.EmployeeType = &v
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SignupAccountRequestDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignupAccountRequestDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *SignupAccountRequestDto) SetKey(v string) {
	o.Key.Set(&v)
}

// GetCulture returns the Culture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignupAccountRequestDto) GetCulture() string {
	if o == nil || IsNil(o.Culture.Get()) {
		var ret string
		return ret
	}
	return *o.Culture.Get()
}

// GetCultureOk returns a tuple with the Culture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignupAccountRequestDto) GetCultureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Culture.Get(), o.Culture.IsSet()
}

// HasCulture returns a boolean if a field has been set.
func (o *SignupAccountRequestDto) IsCultureSet() bool {
	if o != nil && o.Culture.IsSet() {
		return true
	}

	return false
}

// SetCulture gets a reference to the given NullableString and assigns it to the Culture field.
func (o *SignupAccountRequestDto) SetCulture(v string) {
	o.Culture.Set(&v)
}
// SetCultureNil sets the value for Culture to be an explicit nil
func (o *SignupAccountRequestDto) SetCultureNil() {
	o.Culture.Set(nil)
}

// UnsetCulture ensures that no value is present for Culture, not even an explicit nil
func (o *SignupAccountRequestDto) UnsetCulture() {
	o.Culture.Unset()
}

// GetSerializedProfile returns the SerializedProfile field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SignupAccountRequestDto) GetSerializedProfile() string {
	if o == nil || o.SerializedProfile.Get() == nil {
		var ret string
		return ret
	}

	return *o.SerializedProfile.Get()
}

// GetSerializedProfileOk returns a tuple with the SerializedProfile field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignupAccountRequestDto) GetSerializedProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SerializedProfile.Get(), o.SerializedProfile.IsSet()
}

// SetSerializedProfile sets field value
func (o *SignupAccountRequestDto) SetSerializedProfile(v string) {
	o.SerializedProfile.Set(&v)
}

func (o SignupAccountRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignupAccountRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmployeeType) {
		toSerialize["employeeType"] = o.EmployeeType
	}
	toSerialize["key"] = o.Key.Get()
	if o.Culture.IsSet() {
		toSerialize["culture"] = o.Culture.Get()
	}
	toSerialize["serializedProfile"] = o.SerializedProfile.Get()
	return toSerialize, nil
}

func (o *SignupAccountRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"serializedProfile",
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

	varSignupAccountRequestDto := _SignupAccountRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSignupAccountRequestDto)

	if err != nil {
		return err
	}

	*o = SignupAccountRequestDto(varSignupAccountRequestDto)

	return err
}

type NullableSignupAccountRequestDto struct {
	value *SignupAccountRequestDto
	isSet bool
}

func (v NullableSignupAccountRequestDto) Get() *SignupAccountRequestDto {
	return v.value
}

func (v *NullableSignupAccountRequestDto) Set(val *SignupAccountRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSignupAccountRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSignupAccountRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignupAccountRequestDto(val *SignupAccountRequestDto) *NullableSignupAccountRequestDto {
	return &NullableSignupAccountRequestDto{value: val, isSet: true}
}

func (v NullableSignupAccountRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignupAccountRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

