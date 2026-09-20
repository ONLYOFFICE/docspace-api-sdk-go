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

// checks if the PasswordSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordSettingsDto{}

// PasswordSettingsDto The password policy of the portal, with the expressions a client can check a password against.
type PasswordSettingsDto struct {
	// The shortest password the portal accepts, 8 characters on a portal nobody has configured. Whatever the  policy says, a password longer than 30 characters is refused as well, and that ceiling is not reported  here.
	MinLength int32 `json:"minLength"`
	// Whether at least one uppercase letter is demanded. While it is `false` an uppercase letter is still  allowed - the flag adds a requirement rather than permission.
	UpperCase bool `json:"upperCase"`
	// Whether at least one digit is demanded, read the same way as `upperCase`.
	Digits bool `json:"digits"`
	// Whether at least one special symbol is demanded, read the same way as `upperCase`. Which symbols count is  spelled out by `specSymbolsRegexStr`.
	SpecSymbols bool `json:"specSymbols"`
	// The expression the whole password has to match, which is what defines the alphabet the portal accepts at  all. It comes from the installation's configuration rather than from the portal policy, so it is the same  for every portal of an installation and unaffected by the flags above.
	AllowedCharactersRegexStr NullableString `json:"allowedCharactersRegexStr"`
	// The look-ahead expression that tests the digit requirement, meant to be applied only while `digits` is  `true`. It is always filled in, so its presence is not itself a requirement.
	DigitsRegexStr NullableString `json:"digitsRegexStr"`
	// The look-ahead expression that tests the uppercase requirement, to be applied while `upperCase` is `true`.
	UpperCaseRegexStr NullableString `json:"upperCaseRegexStr"`
	// The look-ahead expression that tests the special-symbol requirement, to be applied while `specSymbols` is  `true`. It also enumerates the symbols the portal treats as special.
	SpecSymbolsRegexStr NullableString `json:"specSymbolsRegexStr"`
}

type _PasswordSettingsDto PasswordSettingsDto

// NewPasswordSettingsDto instantiates a new PasswordSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordSettingsDto(minLength int32, upperCase bool, digits bool, specSymbols bool, allowedCharactersRegexStr NullableString, digitsRegexStr NullableString, upperCaseRegexStr NullableString, specSymbolsRegexStr NullableString) *PasswordSettingsDto {
	this := PasswordSettingsDto{}
	this.MinLength = minLength
	this.UpperCase = upperCase
	this.Digits = digits
	this.SpecSymbols = specSymbols
	this.AllowedCharactersRegexStr = allowedCharactersRegexStr
	this.DigitsRegexStr = digitsRegexStr
	this.UpperCaseRegexStr = upperCaseRegexStr
	this.SpecSymbolsRegexStr = specSymbolsRegexStr
	return &this
}

// NewPasswordSettingsDtoWithDefaults instantiates a new PasswordSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordSettingsDtoWithDefaults() *PasswordSettingsDto {
	this := PasswordSettingsDto{}
	return &this
}

// GetMinLength returns the MinLength field value
func (o *PasswordSettingsDto) GetMinLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value
// and a boolean to check if the value has been set.
func (o *PasswordSettingsDto) GetMinLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinLength, true
}

// SetMinLength sets field value
func (o *PasswordSettingsDto) SetMinLength(v int32) {
	o.MinLength = v
}

// GetUpperCase returns the UpperCase field value
func (o *PasswordSettingsDto) GetUpperCase() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UpperCase
}

// GetUpperCaseOk returns a tuple with the UpperCase field value
// and a boolean to check if the value has been set.
func (o *PasswordSettingsDto) GetUpperCaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpperCase, true
}

// SetUpperCase sets field value
func (o *PasswordSettingsDto) SetUpperCase(v bool) {
	o.UpperCase = v
}

// GetDigits returns the Digits field value
func (o *PasswordSettingsDto) GetDigits() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Digits
}

// GetDigitsOk returns a tuple with the Digits field value
// and a boolean to check if the value has been set.
func (o *PasswordSettingsDto) GetDigitsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digits, true
}

// SetDigits sets field value
func (o *PasswordSettingsDto) SetDigits(v bool) {
	o.Digits = v
}

// GetSpecSymbols returns the SpecSymbols field value
func (o *PasswordSettingsDto) GetSpecSymbols() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SpecSymbols
}

// GetSpecSymbolsOk returns a tuple with the SpecSymbols field value
// and a boolean to check if the value has been set.
func (o *PasswordSettingsDto) GetSpecSymbolsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecSymbols, true
}

// SetSpecSymbols sets field value
func (o *PasswordSettingsDto) SetSpecSymbols(v bool) {
	o.SpecSymbols = v
}

// GetAllowedCharactersRegexStr returns the AllowedCharactersRegexStr field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PasswordSettingsDto) GetAllowedCharactersRegexStr() string {
	if o == nil || o.AllowedCharactersRegexStr.Get() == nil {
		var ret string
		return ret
	}

	return *o.AllowedCharactersRegexStr.Get()
}

// GetAllowedCharactersRegexStrOk returns a tuple with the AllowedCharactersRegexStr field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordSettingsDto) GetAllowedCharactersRegexStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedCharactersRegexStr.Get(), o.AllowedCharactersRegexStr.IsSet()
}

// SetAllowedCharactersRegexStr sets field value
func (o *PasswordSettingsDto) SetAllowedCharactersRegexStr(v string) {
	o.AllowedCharactersRegexStr.Set(&v)
}

// GetDigitsRegexStr returns the DigitsRegexStr field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PasswordSettingsDto) GetDigitsRegexStr() string {
	if o == nil || o.DigitsRegexStr.Get() == nil {
		var ret string
		return ret
	}

	return *o.DigitsRegexStr.Get()
}

// GetDigitsRegexStrOk returns a tuple with the DigitsRegexStr field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordSettingsDto) GetDigitsRegexStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DigitsRegexStr.Get(), o.DigitsRegexStr.IsSet()
}

// SetDigitsRegexStr sets field value
func (o *PasswordSettingsDto) SetDigitsRegexStr(v string) {
	o.DigitsRegexStr.Set(&v)
}

// GetUpperCaseRegexStr returns the UpperCaseRegexStr field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PasswordSettingsDto) GetUpperCaseRegexStr() string {
	if o == nil || o.UpperCaseRegexStr.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpperCaseRegexStr.Get()
}

// GetUpperCaseRegexStrOk returns a tuple with the UpperCaseRegexStr field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordSettingsDto) GetUpperCaseRegexStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpperCaseRegexStr.Get(), o.UpperCaseRegexStr.IsSet()
}

// SetUpperCaseRegexStr sets field value
func (o *PasswordSettingsDto) SetUpperCaseRegexStr(v string) {
	o.UpperCaseRegexStr.Set(&v)
}

// GetSpecSymbolsRegexStr returns the SpecSymbolsRegexStr field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PasswordSettingsDto) GetSpecSymbolsRegexStr() string {
	if o == nil || o.SpecSymbolsRegexStr.Get() == nil {
		var ret string
		return ret
	}

	return *o.SpecSymbolsRegexStr.Get()
}

// GetSpecSymbolsRegexStrOk returns a tuple with the SpecSymbolsRegexStr field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordSettingsDto) GetSpecSymbolsRegexStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpecSymbolsRegexStr.Get(), o.SpecSymbolsRegexStr.IsSet()
}

// SetSpecSymbolsRegexStr sets field value
func (o *PasswordSettingsDto) SetSpecSymbolsRegexStr(v string) {
	o.SpecSymbolsRegexStr.Set(&v)
}

func (o PasswordSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minLength"] = o.MinLength
	toSerialize["upperCase"] = o.UpperCase
	toSerialize["digits"] = o.Digits
	toSerialize["specSymbols"] = o.SpecSymbols
	toSerialize["allowedCharactersRegexStr"] = o.AllowedCharactersRegexStr.Get()
	toSerialize["digitsRegexStr"] = o.DigitsRegexStr.Get()
	toSerialize["upperCaseRegexStr"] = o.UpperCaseRegexStr.Get()
	toSerialize["specSymbolsRegexStr"] = o.SpecSymbolsRegexStr.Get()
	return toSerialize, nil
}

func (o *PasswordSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"minLength",
		"upperCase",
		"digits",
		"specSymbols",
		"allowedCharactersRegexStr",
		"digitsRegexStr",
		"upperCaseRegexStr",
		"specSymbolsRegexStr",
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

	varPasswordSettingsDto := _PasswordSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPasswordSettingsDto)

	if err != nil {
		return err
	}

	*o = PasswordSettingsDto(varPasswordSettingsDto)

	return err
}

type NullablePasswordSettingsDto struct {
	value *PasswordSettingsDto
	isSet bool
}

func (v NullablePasswordSettingsDto) Get() *PasswordSettingsDto {
	return v.value
}

func (v *NullablePasswordSettingsDto) Set(val *PasswordSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordSettingsDto(val *PasswordSettingsDto) *NullablePasswordSettingsDto {
	return &NullablePasswordSettingsDto{value: val, isSet: true}
}

func (v NullablePasswordSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

