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

// checks if the EmailValidationKeyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailValidationKeyModel{}

// EmailValidationKeyModel The confirmation email parameters.
type EmailValidationKeyModel struct {
	// The email validation key.
	Key NullableString `json:"key,omitempty"`
	// The user type.
	EmplType *EmployeeType `json:"emplType,omitempty"`
	// The email address.
	Email NullableString `json:"email,omitempty"`
	// The encrypted email address.
	EncEmail NullableString `json:"encEmail,omitempty"`
	// The user ID.
	UiD NullableString `json:"uiD,omitempty"`
	// The confirmation email type.
	Type *ConfirmType `json:"type,omitempty"`
	// Specifies whether it is the first time account access or not.
	First NullableString `json:"first,omitempty"`
	// The room ID.
	RoomId NullableString `json:"roomId,omitempty"`
}

// NewEmailValidationKeyModel instantiates a new EmailValidationKeyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailValidationKeyModel() *EmailValidationKeyModel {
	this := EmailValidationKeyModel{}
	return &this
}

// NewEmailValidationKeyModelWithDefaults instantiates a new EmailValidationKeyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailValidationKeyModelWithDefaults() *EmailValidationKeyModel {
	this := EmailValidationKeyModel{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailValidationKeyModel) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailValidationKeyModel) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *EmailValidationKeyModel) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *EmailValidationKeyModel) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *EmailValidationKeyModel) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *EmailValidationKeyModel) UnsetKey() {
	o.Key.Unset()
}

// GetEmplType returns the EmplType field value if set, zero value otherwise.
func (o *EmailValidationKeyModel) GetEmplType() EmployeeType {
	if o == nil || IsNil(o.EmplType) {
		var ret EmployeeType
		return ret
	}
	return *o.EmplType
}

// GetEmplTypeOk returns a tuple with the EmplType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationKeyModel) GetEmplTypeOk() (*EmployeeType, bool) {
	if o == nil || IsNil(o.EmplType) {
		return nil, false
	}
	return o.EmplType, true
}

// HasEmplType returns a boolean if a field has been set.
func (o *EmailValidationKeyModel) IsEmplTypeSet() bool {
	if o != nil && !IsNil(o.EmplType) {
		return true
	}

	return false
}

// SetEmplType gets a reference to the given EmployeeType and assigns it to the EmplType field.
func (o *EmailValidationKeyModel) SetEmplType(v EmployeeType) {
	o.EmplType = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailValidationKeyModel) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailValidationKeyModel) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailValidationKeyModel) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *EmailValidationKeyModel) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *EmailValidationKeyModel) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *EmailValidationKeyModel) UnsetEmail() {
	o.Email.Unset()
}

// GetEncEmail returns the EncEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailValidationKeyModel) GetEncEmail() string {
	if o == nil || IsNil(o.EncEmail.Get()) {
		var ret string
		return ret
	}
	return *o.EncEmail.Get()
}

// GetEncEmailOk returns a tuple with the EncEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailValidationKeyModel) GetEncEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncEmail.Get(), o.EncEmail.IsSet()
}

// HasEncEmail returns a boolean if a field has been set.
func (o *EmailValidationKeyModel) IsEncEmailSet() bool {
	if o != nil && o.EncEmail.IsSet() {
		return true
	}

	return false
}

// SetEncEmail gets a reference to the given NullableString and assigns it to the EncEmail field.
func (o *EmailValidationKeyModel) SetEncEmail(v string) {
	o.EncEmail.Set(&v)
}
// SetEncEmailNil sets the value for EncEmail to be an explicit nil
func (o *EmailValidationKeyModel) SetEncEmailNil() {
	o.EncEmail.Set(nil)
}

// UnsetEncEmail ensures that no value is present for EncEmail, not even an explicit nil
func (o *EmailValidationKeyModel) UnsetEncEmail() {
	o.EncEmail.Unset()
}

// GetUiD returns the UiD field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailValidationKeyModel) GetUiD() string {
	if o == nil || IsNil(o.UiD.Get()) {
		var ret string
		return ret
	}
	return *o.UiD.Get()
}

// GetUiDOk returns a tuple with the UiD field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailValidationKeyModel) GetUiDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UiD.Get(), o.UiD.IsSet()
}

// HasUiD returns a boolean if a field has been set.
func (o *EmailValidationKeyModel) IsUiDSet() bool {
	if o != nil && o.UiD.IsSet() {
		return true
	}

	return false
}

// SetUiD gets a reference to the given NullableString and assigns it to the UiD field.
func (o *EmailValidationKeyModel) SetUiD(v string) {
	o.UiD.Set(&v)
}
// SetUiDNil sets the value for UiD to be an explicit nil
func (o *EmailValidationKeyModel) SetUiDNil() {
	o.UiD.Set(nil)
}

// UnsetUiD ensures that no value is present for UiD, not even an explicit nil
func (o *EmailValidationKeyModel) UnsetUiD() {
	o.UiD.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EmailValidationKeyModel) GetType() ConfirmType {
	if o == nil || IsNil(o.Type) {
		var ret ConfirmType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationKeyModel) GetTypeOk() (*ConfirmType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EmailValidationKeyModel) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConfirmType and assigns it to the Type field.
func (o *EmailValidationKeyModel) SetType(v ConfirmType) {
	o.Type = &v
}

// GetFirst returns the First field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailValidationKeyModel) GetFirst() string {
	if o == nil || IsNil(o.First.Get()) {
		var ret string
		return ret
	}
	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailValidationKeyModel) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// HasFirst returns a boolean if a field has been set.
func (o *EmailValidationKeyModel) IsFirstSet() bool {
	if o != nil && o.First.IsSet() {
		return true
	}

	return false
}

// SetFirst gets a reference to the given NullableString and assigns it to the First field.
func (o *EmailValidationKeyModel) SetFirst(v string) {
	o.First.Set(&v)
}
// SetFirstNil sets the value for First to be an explicit nil
func (o *EmailValidationKeyModel) SetFirstNil() {
	o.First.Set(nil)
}

// UnsetFirst ensures that no value is present for First, not even an explicit nil
func (o *EmailValidationKeyModel) UnsetFirst() {
	o.First.Unset()
}

// GetRoomId returns the RoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailValidationKeyModel) GetRoomId() string {
	if o == nil || IsNil(o.RoomId.Get()) {
		var ret string
		return ret
	}
	return *o.RoomId.Get()
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailValidationKeyModel) GetRoomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomId.Get(), o.RoomId.IsSet()
}

// HasRoomId returns a boolean if a field has been set.
func (o *EmailValidationKeyModel) IsRoomIdSet() bool {
	if o != nil && o.RoomId.IsSet() {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given NullableString and assigns it to the RoomId field.
func (o *EmailValidationKeyModel) SetRoomId(v string) {
	o.RoomId.Set(&v)
}
// SetRoomIdNil sets the value for RoomId to be an explicit nil
func (o *EmailValidationKeyModel) SetRoomIdNil() {
	o.RoomId.Set(nil)
}

// UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
func (o *EmailValidationKeyModel) UnsetRoomId() {
	o.RoomId.Unset()
}

func (o EmailValidationKeyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailValidationKeyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !IsNil(o.EmplType) {
		toSerialize["emplType"] = o.EmplType
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.EncEmail.IsSet() {
		toSerialize["encEmail"] = o.EncEmail.Get()
	}
	if o.UiD.IsSet() {
		toSerialize["uiD"] = o.UiD.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.First.IsSet() {
		toSerialize["first"] = o.First.Get()
	}
	if o.RoomId.IsSet() {
		toSerialize["roomId"] = o.RoomId.Get()
	}
	return toSerialize, nil
}

type NullableEmailValidationKeyModel struct {
	value *EmailValidationKeyModel
	isSet bool
}

func (v NullableEmailValidationKeyModel) Get() *EmailValidationKeyModel {
	return v.value
}

func (v *NullableEmailValidationKeyModel) Set(val *EmailValidationKeyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailValidationKeyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailValidationKeyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailValidationKeyModel(val *EmailValidationKeyModel) *NullableEmailValidationKeyModel {
	return &NullableEmailValidationKeyModel{value: val, isSet: true}
}

func (v NullableEmailValidationKeyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailValidationKeyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

