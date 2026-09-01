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

// checks if the EncryptionSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptionSettings{}

// EncryptionSettings The encryption settings.
type EncryptionSettings struct {
	// The encryption password.
	Password NullableString `json:"password,omitempty"`
	// The encryption status.
	Status *EncryprtionStatus `json:"status,omitempty"`
	// Specifies if the users will be notified about the encryption operation or not.
	NotifyUsers *bool `json:"notifyUsers,omitempty"`
}

// NewEncryptionSettings instantiates a new EncryptionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionSettings() *EncryptionSettings {
	this := EncryptionSettings{}
	return &this
}

// NewEncryptionSettingsWithDefaults instantiates a new EncryptionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionSettingsWithDefaults() *EncryptionSettings {
	this := EncryptionSettings{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionSettings) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionSettings) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *EncryptionSettings) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *EncryptionSettings) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *EncryptionSettings) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *EncryptionSettings) UnsetPassword() {
	o.Password.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EncryptionSettings) GetStatus() EncryprtionStatus {
	if o == nil || IsNil(o.Status) {
		var ret EncryprtionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionSettings) GetStatusOk() (*EncryprtionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EncryptionSettings) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EncryprtionStatus and assigns it to the Status field.
func (o *EncryptionSettings) SetStatus(v EncryprtionStatus) {
	o.Status = &v
}

// GetNotifyUsers returns the NotifyUsers field value if set, zero value otherwise.
func (o *EncryptionSettings) GetNotifyUsers() bool {
	if o == nil || IsNil(o.NotifyUsers) {
		var ret bool
		return ret
	}
	return *o.NotifyUsers
}

// GetNotifyUsersOk returns a tuple with the NotifyUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionSettings) GetNotifyUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyUsers) {
		return nil, false
	}
	return o.NotifyUsers, true
}

// HasNotifyUsers returns a boolean if a field has been set.
func (o *EncryptionSettings) IsNotifyUsersSet() bool {
	if o != nil && !IsNil(o.NotifyUsers) {
		return true
	}

	return false
}

// SetNotifyUsers gets a reference to the given bool and assigns it to the NotifyUsers field.
func (o *EncryptionSettings) SetNotifyUsers(v bool) {
	o.NotifyUsers = &v
}

func (o EncryptionSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptionSettings) ToMap() (map[string]interface{}, error) {
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

type NullableEncryptionSettings struct {
	value *EncryptionSettings
	isSet bool
}

func (v NullableEncryptionSettings) Get() *EncryptionSettings {
	return v.value
}

func (v *NullableEncryptionSettings) Set(val *EncryptionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionSettings(val *EncryptionSettings) *NullableEncryptionSettings {
	return &NullableEncryptionSettings{value: val, isSet: true}
}

func (v NullableEncryptionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

