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

// checks if the EncryptionSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptionSettingsDto{}

// EncryptionSettingsDto The state of the portal's storage encryption.
type EncryptionSettingsDto struct {
	// Always an empty string: the encryption password is never returned.
	Password NullableString `json:"password,omitempty"`
	// Whether the storage is encrypted, decrypted, or on its way to either.
	Status *EncryptionStatus `json:"status,omitempty"`
	// Whether the users are notified when the operation starts and ends.
	NotifyUsers *bool `json:"notifyUsers,omitempty"`
}

// NewEncryptionSettingsDto instantiates a new EncryptionSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionSettingsDto() *EncryptionSettingsDto {
	this := EncryptionSettingsDto{}
	return &this
}

// NewEncryptionSettingsDtoWithDefaults instantiates a new EncryptionSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionSettingsDtoWithDefaults() *EncryptionSettingsDto {
	this := EncryptionSettingsDto{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionSettingsDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionSettingsDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *EncryptionSettingsDto) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *EncryptionSettingsDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *EncryptionSettingsDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *EncryptionSettingsDto) UnsetPassword() {
	o.Password.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EncryptionSettingsDto) GetStatus() EncryptionStatus {
	if o == nil || IsNil(o.Status) {
		var ret EncryptionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsDto) GetStatusOk() (*EncryptionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EncryptionSettingsDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EncryptionStatus and assigns it to the Status field.
func (o *EncryptionSettingsDto) SetStatus(v EncryptionStatus) {
	o.Status = &v
}

// GetNotifyUsers returns the NotifyUsers field value if set, zero value otherwise.
func (o *EncryptionSettingsDto) GetNotifyUsers() bool {
	if o == nil || IsNil(o.NotifyUsers) {
		var ret bool
		return ret
	}
	return *o.NotifyUsers
}

// GetNotifyUsersOk returns a tuple with the NotifyUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsDto) GetNotifyUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyUsers) {
		return nil, false
	}
	return o.NotifyUsers, true
}

// HasNotifyUsers returns a boolean if a field has been set.
func (o *EncryptionSettingsDto) IsNotifyUsersSet() bool {
	if o != nil && !IsNil(o.NotifyUsers) {
		return true
	}

	return false
}

// SetNotifyUsers gets a reference to the given bool and assigns it to the NotifyUsers field.
func (o *EncryptionSettingsDto) SetNotifyUsers(v bool) {
	o.NotifyUsers = &v
}

func (o EncryptionSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptionSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.NotifyUsers) {
		toSerialize["notifyUsers"] = o.NotifyUsers
	}
	return toSerialize, nil
}

type NullableEncryptionSettingsDto struct {
	value *EncryptionSettingsDto
	isSet bool
}

func (v NullableEncryptionSettingsDto) Get() *EncryptionSettingsDto {
	return v.value
}

func (v *NullableEncryptionSettingsDto) Set(val *EncryptionSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionSettingsDto(val *EncryptionSettingsDto) *NullableEncryptionSettingsDto {
	return &NullableEncryptionSettingsDto{value: val, isSet: true}
}

func (v NullableEncryptionSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

