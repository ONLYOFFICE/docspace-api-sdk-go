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

// checks if the SetupCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetupCode{}

// SetupCode struct for SetupCode
type SetupCode struct {
	Account NullableString `json:"account,omitempty"`
	ManualEntryKey NullableString `json:"manualEntryKey,omitempty"`
	QrCodeSetupImageUrl NullableString `json:"qrCodeSetupImageUrl,omitempty"`
}

// NewSetupCode instantiates a new SetupCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetupCode() *SetupCode {
	this := SetupCode{}
	return &this
}

// NewSetupCodeWithDefaults instantiates a new SetupCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetupCodeWithDefaults() *SetupCode {
	this := SetupCode{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetupCode) GetAccount() string {
	if o == nil || IsNil(o.Account.Get()) {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupCode) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *SetupCode) IsAccountSet() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *SetupCode) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *SetupCode) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *SetupCode) UnsetAccount() {
	o.Account.Unset()
}

// GetManualEntryKey returns the ManualEntryKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetupCode) GetManualEntryKey() string {
	if o == nil || IsNil(o.ManualEntryKey.Get()) {
		var ret string
		return ret
	}
	return *o.ManualEntryKey.Get()
}

// GetManualEntryKeyOk returns a tuple with the ManualEntryKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupCode) GetManualEntryKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManualEntryKey.Get(), o.ManualEntryKey.IsSet()
}

// HasManualEntryKey returns a boolean if a field has been set.
func (o *SetupCode) IsManualEntryKeySet() bool {
	if o != nil && o.ManualEntryKey.IsSet() {
		return true
	}

	return false
}

// SetManualEntryKey gets a reference to the given NullableString and assigns it to the ManualEntryKey field.
func (o *SetupCode) SetManualEntryKey(v string) {
	o.ManualEntryKey.Set(&v)
}
// SetManualEntryKeyNil sets the value for ManualEntryKey to be an explicit nil
func (o *SetupCode) SetManualEntryKeyNil() {
	o.ManualEntryKey.Set(nil)
}

// UnsetManualEntryKey ensures that no value is present for ManualEntryKey, not even an explicit nil
func (o *SetupCode) UnsetManualEntryKey() {
	o.ManualEntryKey.Unset()
}

// GetQrCodeSetupImageUrl returns the QrCodeSetupImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetupCode) GetQrCodeSetupImageUrl() string {
	if o == nil || IsNil(o.QrCodeSetupImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.QrCodeSetupImageUrl.Get()
}

// GetQrCodeSetupImageUrlOk returns a tuple with the QrCodeSetupImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupCode) GetQrCodeSetupImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QrCodeSetupImageUrl.Get(), o.QrCodeSetupImageUrl.IsSet()
}

// HasQrCodeSetupImageUrl returns a boolean if a field has been set.
func (o *SetupCode) IsQrCodeSetupImageUrlSet() bool {
	if o != nil && o.QrCodeSetupImageUrl.IsSet() {
		return true
	}

	return false
}

// SetQrCodeSetupImageUrl gets a reference to the given NullableString and assigns it to the QrCodeSetupImageUrl field.
func (o *SetupCode) SetQrCodeSetupImageUrl(v string) {
	o.QrCodeSetupImageUrl.Set(&v)
}
// SetQrCodeSetupImageUrlNil sets the value for QrCodeSetupImageUrl to be an explicit nil
func (o *SetupCode) SetQrCodeSetupImageUrlNil() {
	o.QrCodeSetupImageUrl.Set(nil)
}

// UnsetQrCodeSetupImageUrl ensures that no value is present for QrCodeSetupImageUrl, not even an explicit nil
func (o *SetupCode) UnsetQrCodeSetupImageUrl() {
	o.QrCodeSetupImageUrl.Unset()
}

func (o SetupCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetupCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.ManualEntryKey.IsSet() {
		toSerialize["manualEntryKey"] = o.ManualEntryKey.Get()
	}
	if o.QrCodeSetupImageUrl.IsSet() {
		toSerialize["qrCodeSetupImageUrl"] = o.QrCodeSetupImageUrl.Get()
	}
	return toSerialize, nil
}

type NullableSetupCode struct {
	value *SetupCode
	isSet bool
}

func (v NullableSetupCode) Get() *SetupCode {
	return v.value
}

func (v *NullableSetupCode) Set(val *SetupCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSetupCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSetupCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetupCode(val *SetupCode) *NullableSetupCode {
	return &NullableSetupCode{value: val, isSet: true}
}

func (v NullableSetupCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetupCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

