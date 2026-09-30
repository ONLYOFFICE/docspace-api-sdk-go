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
)

// checks if the TfaAppCodeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TfaAppCodeDto{}

// TfaAppCodeDto One backup code of the caller's authenticator credential.
type TfaAppCodeDto struct {
	// Whether the code has already been spent. A spent code is kept in the list but is no longer accepted, so  count the entries where this is `false` to know how many fallbacks remain.
	IsUsed *bool `json:"isUsed,omitempty"`
	// The code itself, in the form it is typed at sign-in - six characters with the default configuration. It is  stored encrypted and decrypted for this answer, so this is the one place a caller can read it.
	Code NullableString `json:"code,omitempty"`
}

// NewTfaAppCodeDto instantiates a new TfaAppCodeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaAppCodeDto() *TfaAppCodeDto {
	this := TfaAppCodeDto{}
	return &this
}

// NewTfaAppCodeDtoWithDefaults instantiates a new TfaAppCodeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaAppCodeDtoWithDefaults() *TfaAppCodeDto {
	this := TfaAppCodeDto{}
	return &this
}

// GetIsUsed returns the IsUsed field value if set, zero value otherwise.
func (o *TfaAppCodeDto) GetIsUsed() bool {
	if o == nil || IsNil(o.IsUsed) {
		var ret bool
		return ret
	}
	return *o.IsUsed
}

// GetIsUsedOk returns a tuple with the IsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaAppCodeDto) GetIsUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUsed) {
		return nil, false
	}
	return o.IsUsed, true
}

// HasIsUsed returns a boolean if a field has been set.
func (o *TfaAppCodeDto) IsIsUsedSet() bool {
	if o != nil && !IsNil(o.IsUsed) {
		return true
	}

	return false
}

// SetIsUsed gets a reference to the given bool and assigns it to the IsUsed field.
func (o *TfaAppCodeDto) SetIsUsed(v bool) {
	o.IsUsed = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaAppCodeDto) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaAppCodeDto) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *TfaAppCodeDto) IsCodeSet() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *TfaAppCodeDto) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *TfaAppCodeDto) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *TfaAppCodeDto) UnsetCode() {
	o.Code.Unset()
}

func (o TfaAppCodeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TfaAppCodeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsUsed) {
		toSerialize["isUsed"] = o.IsUsed
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	return toSerialize, nil
}

type NullableTfaAppCodeDto struct {
	value *TfaAppCodeDto
	isSet bool
}

func (v NullableTfaAppCodeDto) Get() *TfaAppCodeDto {
	return v.value
}

func (v *NullableTfaAppCodeDto) Set(val *TfaAppCodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaAppCodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaAppCodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaAppCodeDto(val *TfaAppCodeDto) *NullableTfaAppCodeDto {
	return &NullableTfaAppCodeDto{value: val, isSet: true}
}

func (v NullableTfaAppCodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaAppCodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

