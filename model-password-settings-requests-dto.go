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

// checks if the PasswordSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordSettingsRequestsDto{}

// PasswordSettingsRequestsDto The request parameters for configuring the password complexity requirements.
type PasswordSettingsRequestsDto struct {
	// The minimum number of characters required for valid passwords.
	MinLength int32 `json:"minLength"`
	// Specifies whether the password should contain the uppercase letters or not.
	UpperCase *bool `json:"upperCase,omitempty"`
	// Specifies whether the password should contain the digits or not.
	Digits *bool `json:"digits,omitempty"`
	// Specifies whether the password should contain the special symbols or not.
	SpecSymbols *bool `json:"specSymbols,omitempty"`
}

type _PasswordSettingsRequestsDto PasswordSettingsRequestsDto

// NewPasswordSettingsRequestsDto instantiates a new PasswordSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordSettingsRequestsDto(minLength int32) *PasswordSettingsRequestsDto {
	this := PasswordSettingsRequestsDto{}
	this.MinLength = minLength
	return &this
}

// NewPasswordSettingsRequestsDtoWithDefaults instantiates a new PasswordSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordSettingsRequestsDtoWithDefaults() *PasswordSettingsRequestsDto {
	this := PasswordSettingsRequestsDto{}
	return &this
}

// GetMinLength returns the MinLength field value
func (o *PasswordSettingsRequestsDto) GetMinLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value
// and a boolean to check if the value has been set.
func (o *PasswordSettingsRequestsDto) GetMinLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinLength, true
}

// SetMinLength sets field value
func (o *PasswordSettingsRequestsDto) SetMinLength(v int32) {
	o.MinLength = v
}

// GetUpperCase returns the UpperCase field value if set, zero value otherwise.
func (o *PasswordSettingsRequestsDto) GetUpperCase() bool {
	if o == nil || IsNil(o.UpperCase) {
		var ret bool
		return ret
	}
	return *o.UpperCase
}

// GetUpperCaseOk returns a tuple with the UpperCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSettingsRequestsDto) GetUpperCaseOk() (*bool, bool) {
	if o == nil || IsNil(o.UpperCase) {
		return nil, false
	}
	return o.UpperCase, true
}

// HasUpperCase returns a boolean if a field has been set.
func (o *PasswordSettingsRequestsDto) IsUpperCaseSet() bool {
	if o != nil && !IsNil(o.UpperCase) {
		return true
	}

	return false
}

// SetUpperCase gets a reference to the given bool and assigns it to the UpperCase field.
func (o *PasswordSettingsRequestsDto) SetUpperCase(v bool) {
	o.UpperCase = &v
}

// GetDigits returns the Digits field value if set, zero value otherwise.
func (o *PasswordSettingsRequestsDto) GetDigits() bool {
	if o == nil || IsNil(o.Digits) {
		var ret bool
		return ret
	}
	return *o.Digits
}

// GetDigitsOk returns a tuple with the Digits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSettingsRequestsDto) GetDigitsOk() (*bool, bool) {
	if o == nil || IsNil(o.Digits) {
		return nil, false
	}
	return o.Digits, true
}

// HasDigits returns a boolean if a field has been set.
func (o *PasswordSettingsRequestsDto) IsDigitsSet() bool {
	if o != nil && !IsNil(o.Digits) {
		return true
	}

	return false
}

// SetDigits gets a reference to the given bool and assigns it to the Digits field.
func (o *PasswordSettingsRequestsDto) SetDigits(v bool) {
	o.Digits = &v
}

// GetSpecSymbols returns the SpecSymbols field value if set, zero value otherwise.
func (o *PasswordSettingsRequestsDto) GetSpecSymbols() bool {
	if o == nil || IsNil(o.SpecSymbols) {
		var ret bool
		return ret
	}
	return *o.SpecSymbols
}

// GetSpecSymbolsOk returns a tuple with the SpecSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSettingsRequestsDto) GetSpecSymbolsOk() (*bool, bool) {
	if o == nil || IsNil(o.SpecSymbols) {
		return nil, false
	}
	return o.SpecSymbols, true
}

// HasSpecSymbols returns a boolean if a field has been set.
func (o *PasswordSettingsRequestsDto) IsSpecSymbolsSet() bool {
	if o != nil && !IsNil(o.SpecSymbols) {
		return true
	}

	return false
}

// SetSpecSymbols gets a reference to the given bool and assigns it to the SpecSymbols field.
func (o *PasswordSettingsRequestsDto) SetSpecSymbols(v bool) {
	o.SpecSymbols = &v
}

func (o PasswordSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minLength"] = o.MinLength
	if !IsNil(o.UpperCase) {
		toSerialize["upperCase"] = o.UpperCase
	}
	if !IsNil(o.Digits) {
		toSerialize["digits"] = o.Digits
	}
	if !IsNil(o.SpecSymbols) {
		toSerialize["specSymbols"] = o.SpecSymbols
	}
	return toSerialize, nil
}

func (o *PasswordSettingsRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"minLength",
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

	varPasswordSettingsRequestsDto := _PasswordSettingsRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPasswordSettingsRequestsDto)

	if err != nil {
		return err
	}

	*o = PasswordSettingsRequestsDto(varPasswordSettingsRequestsDto)

	return err
}

type NullablePasswordSettingsRequestsDto struct {
	value *PasswordSettingsRequestsDto
	isSet bool
}

func (v NullablePasswordSettingsRequestsDto) Get() *PasswordSettingsRequestsDto {
	return v.value
}

func (v *NullablePasswordSettingsRequestsDto) Set(val *PasswordSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordSettingsRequestsDto(val *PasswordSettingsRequestsDto) *NullablePasswordSettingsRequestsDto {
	return &NullablePasswordSettingsRequestsDto{value: val, isSet: true}
}

func (v NullablePasswordSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

