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

// checks if the TfaSetupCodeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TfaSetupCodeDto{}

// TfaSetupCodeDto The secret to enrol in an authenticator application, in both of the forms an application can take it.
type TfaSetupCodeDto struct {
	// The label the authenticator application will list the credential under, which is the caller's own email  address. It identifies the entry to a person, and no application checks it.
	Account NullableString `json:"account,omitempty"`
	// The secret in the base32 form that is typed into an application by hand. It describes the very same  credential as `qrCodeSetupImageUrl`, and repeating the call hands back the same value for the account until  the credential is reset.
	ManualEntryKey NullableString `json:"manualEntryKey,omitempty"`
	// The same secret as a scannable image, given as a `data:image/png;base64,` URL that can be rendered  directly - it is not a link to fetch.
	QrCodeSetupImageUrl NullableString `json:"qrCodeSetupImageUrl,omitempty"`
}

// NewTfaSetupCodeDto instantiates a new TfaSetupCodeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaSetupCodeDto() *TfaSetupCodeDto {
	this := TfaSetupCodeDto{}
	return &this
}

// NewTfaSetupCodeDtoWithDefaults instantiates a new TfaSetupCodeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaSetupCodeDtoWithDefaults() *TfaSetupCodeDto {
	this := TfaSetupCodeDto{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaSetupCodeDto) GetAccount() string {
	if o == nil || IsNil(o.Account.Get()) {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaSetupCodeDto) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *TfaSetupCodeDto) IsAccountSet() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *TfaSetupCodeDto) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *TfaSetupCodeDto) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *TfaSetupCodeDto) UnsetAccount() {
	o.Account.Unset()
}

// GetManualEntryKey returns the ManualEntryKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaSetupCodeDto) GetManualEntryKey() string {
	if o == nil || IsNil(o.ManualEntryKey.Get()) {
		var ret string
		return ret
	}
	return *o.ManualEntryKey.Get()
}

// GetManualEntryKeyOk returns a tuple with the ManualEntryKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaSetupCodeDto) GetManualEntryKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManualEntryKey.Get(), o.ManualEntryKey.IsSet()
}

// HasManualEntryKey returns a boolean if a field has been set.
func (o *TfaSetupCodeDto) IsManualEntryKeySet() bool {
	if o != nil && o.ManualEntryKey.IsSet() {
		return true
	}

	return false
}

// SetManualEntryKey gets a reference to the given NullableString and assigns it to the ManualEntryKey field.
func (o *TfaSetupCodeDto) SetManualEntryKey(v string) {
	o.ManualEntryKey.Set(&v)
}
// SetManualEntryKeyNil sets the value for ManualEntryKey to be an explicit nil
func (o *TfaSetupCodeDto) SetManualEntryKeyNil() {
	o.ManualEntryKey.Set(nil)
}

// UnsetManualEntryKey ensures that no value is present for ManualEntryKey, not even an explicit nil
func (o *TfaSetupCodeDto) UnsetManualEntryKey() {
	o.ManualEntryKey.Unset()
}

// GetQrCodeSetupImageUrl returns the QrCodeSetupImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaSetupCodeDto) GetQrCodeSetupImageUrl() string {
	if o == nil || IsNil(o.QrCodeSetupImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.QrCodeSetupImageUrl.Get()
}

// GetQrCodeSetupImageUrlOk returns a tuple with the QrCodeSetupImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaSetupCodeDto) GetQrCodeSetupImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QrCodeSetupImageUrl.Get(), o.QrCodeSetupImageUrl.IsSet()
}

// HasQrCodeSetupImageUrl returns a boolean if a field has been set.
func (o *TfaSetupCodeDto) IsQrCodeSetupImageUrlSet() bool {
	if o != nil && o.QrCodeSetupImageUrl.IsSet() {
		return true
	}

	return false
}

// SetQrCodeSetupImageUrl gets a reference to the given NullableString and assigns it to the QrCodeSetupImageUrl field.
func (o *TfaSetupCodeDto) SetQrCodeSetupImageUrl(v string) {
	o.QrCodeSetupImageUrl.Set(&v)
}
// SetQrCodeSetupImageUrlNil sets the value for QrCodeSetupImageUrl to be an explicit nil
func (o *TfaSetupCodeDto) SetQrCodeSetupImageUrlNil() {
	o.QrCodeSetupImageUrl.Set(nil)
}

// UnsetQrCodeSetupImageUrl ensures that no value is present for QrCodeSetupImageUrl, not even an explicit nil
func (o *TfaSetupCodeDto) UnsetQrCodeSetupImageUrl() {
	o.QrCodeSetupImageUrl.Unset()
}

func (o TfaSetupCodeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TfaSetupCodeDto) ToMap() (map[string]interface{}, error) {
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

type NullableTfaSetupCodeDto struct {
	value *TfaSetupCodeDto
	isSet bool
}

func (v NullableTfaSetupCodeDto) Get() *TfaSetupCodeDto {
	return v.value
}

func (v *NullableTfaSetupCodeDto) Set(val *TfaSetupCodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaSetupCodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaSetupCodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaSetupCodeDto(val *TfaSetupCodeDto) *NullableTfaSetupCodeDto {
	return &NullableTfaSetupCodeDto{value: val, isSet: true}
}

func (v NullableTfaSetupCodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaSetupCodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

