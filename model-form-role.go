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
	"time"
)

// checks if the FormRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormRole{}

// FormRole The form role.
type FormRole struct {
	// The room ID.
	RoomId *int32 `json:"roomId,omitempty"`
	// The role name.
	RoleName NullableString `json:"roleName,omitempty"`
	// The role color.
	RoleColor NullableString `json:"roleColor,omitempty"`
	// The user ID.
	UserId *string `json:"userId,omitempty"`
	// The role sequence.
	Sequence *int32 `json:"sequence,omitempty"`
	// Specifies if the role was submitted or not.
	Submitted *bool `json:"submitted,omitempty"`
	// The date and time when the role was opened.
	OpenedAt *time.Time `json:"openedAt,omitempty"`
	// The date and time when the role was submitted.
	SubmissionDate *time.Time `json:"submissionDate,omitempty"`
}

// NewFormRole instantiates a new FormRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormRole() *FormRole {
	this := FormRole{}
	return &this
}

// NewFormRoleWithDefaults instantiates a new FormRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormRoleWithDefaults() *FormRole {
	this := FormRole{}
	return &this
}

// GetRoomId returns the RoomId field value if set, zero value otherwise.
func (o *FormRole) GetRoomId() int32 {
	if o == nil || IsNil(o.RoomId) {
		var ret int32
		return ret
	}
	return *o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRole) GetRoomIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RoomId) {
		return nil, false
	}
	return o.RoomId, true
}

// HasRoomId returns a boolean if a field has been set.
func (o *FormRole) IsRoomIdSet() bool {
	if o != nil && !IsNil(o.RoomId) {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given int32 and assigns it to the RoomId field.
func (o *FormRole) SetRoomId(v int32) {
	o.RoomId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormRole) GetRoleName() string {
	if o == nil || IsNil(o.RoleName.Get()) {
		var ret string
		return ret
	}
	return *o.RoleName.Get()
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormRole) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleName.Get(), o.RoleName.IsSet()
}

// HasRoleName returns a boolean if a field has been set.
func (o *FormRole) IsRoleNameSet() bool {
	if o != nil && o.RoleName.IsSet() {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given NullableString and assigns it to the RoleName field.
func (o *FormRole) SetRoleName(v string) {
	o.RoleName.Set(&v)
}
// SetRoleNameNil sets the value for RoleName to be an explicit nil
func (o *FormRole) SetRoleNameNil() {
	o.RoleName.Set(nil)
}

// UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
func (o *FormRole) UnsetRoleName() {
	o.RoleName.Unset()
}

// GetRoleColor returns the RoleColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormRole) GetRoleColor() string {
	if o == nil || IsNil(o.RoleColor.Get()) {
		var ret string
		return ret
	}
	return *o.RoleColor.Get()
}

// GetRoleColorOk returns a tuple with the RoleColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormRole) GetRoleColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleColor.Get(), o.RoleColor.IsSet()
}

// HasRoleColor returns a boolean if a field has been set.
func (o *FormRole) IsRoleColorSet() bool {
	if o != nil && o.RoleColor.IsSet() {
		return true
	}

	return false
}

// SetRoleColor gets a reference to the given NullableString and assigns it to the RoleColor field.
func (o *FormRole) SetRoleColor(v string) {
	o.RoleColor.Set(&v)
}
// SetRoleColorNil sets the value for RoleColor to be an explicit nil
func (o *FormRole) SetRoleColorNil() {
	o.RoleColor.Set(nil)
}

// UnsetRoleColor ensures that no value is present for RoleColor, not even an explicit nil
func (o *FormRole) UnsetRoleColor() {
	o.RoleColor.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FormRole) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRole) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FormRole) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FormRole) SetUserId(v string) {
	o.UserId = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *FormRole) GetSequence() int32 {
	if o == nil || IsNil(o.Sequence) {
		var ret int32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRole) GetSequenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *FormRole) IsSequenceSet() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int32 and assigns it to the Sequence field.
func (o *FormRole) SetSequence(v int32) {
	o.Sequence = &v
}

// GetSubmitted returns the Submitted field value if set, zero value otherwise.
func (o *FormRole) GetSubmitted() bool {
	if o == nil || IsNil(o.Submitted) {
		var ret bool
		return ret
	}
	return *o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRole) GetSubmittedOk() (*bool, bool) {
	if o == nil || IsNil(o.Submitted) {
		return nil, false
	}
	return o.Submitted, true
}

// HasSubmitted returns a boolean if a field has been set.
func (o *FormRole) IsSubmittedSet() bool {
	if o != nil && !IsNil(o.Submitted) {
		return true
	}

	return false
}

// SetSubmitted gets a reference to the given bool and assigns it to the Submitted field.
func (o *FormRole) SetSubmitted(v bool) {
	o.Submitted = &v
}

// GetOpenedAt returns the OpenedAt field value if set, zero value otherwise.
func (o *FormRole) GetOpenedAt() time.Time {
	if o == nil || IsNil(o.OpenedAt) {
		var ret time.Time
		return ret
	}
	return *o.OpenedAt
}

// GetOpenedAtOk returns a tuple with the OpenedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRole) GetOpenedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OpenedAt) {
		return nil, false
	}
	return o.OpenedAt, true
}

// HasOpenedAt returns a boolean if a field has been set.
func (o *FormRole) IsOpenedAtSet() bool {
	if o != nil && !IsNil(o.OpenedAt) {
		return true
	}

	return false
}

// SetOpenedAt gets a reference to the given time.Time and assigns it to the OpenedAt field.
func (o *FormRole) SetOpenedAt(v time.Time) {
	o.OpenedAt = &v
}

// GetSubmissionDate returns the SubmissionDate field value if set, zero value otherwise.
func (o *FormRole) GetSubmissionDate() time.Time {
	if o == nil || IsNil(o.SubmissionDate) {
		var ret time.Time
		return ret
	}
	return *o.SubmissionDate
}

// GetSubmissionDateOk returns a tuple with the SubmissionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRole) GetSubmissionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmissionDate) {
		return nil, false
	}
	return o.SubmissionDate, true
}

// HasSubmissionDate returns a boolean if a field has been set.
func (o *FormRole) IsSubmissionDateSet() bool {
	if o != nil && !IsNil(o.SubmissionDate) {
		return true
	}

	return false
}

// SetSubmissionDate gets a reference to the given time.Time and assigns it to the SubmissionDate field.
func (o *FormRole) SetSubmissionDate(v time.Time) {
	o.SubmissionDate = &v
}

func (o FormRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomId) {
		toSerialize["roomId"] = o.RoomId
	}
	if o.RoleName.IsSet() {
		toSerialize["roleName"] = o.RoleName.Get()
	}
	if o.RoleColor.IsSet() {
		toSerialize["roleColor"] = o.RoleColor.Get()
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !IsNil(o.Submitted) {
		toSerialize["submitted"] = o.Submitted
	}
	if !IsNil(o.OpenedAt) {
		toSerialize["openedAt"] = o.OpenedAt
	}
	if !IsNil(o.SubmissionDate) {
		toSerialize["submissionDate"] = o.SubmissionDate
	}
	return toSerialize, nil
}

type NullableFormRole struct {
	value *FormRole
	isSet bool
}

func (v NullableFormRole) Get() *FormRole {
	return v.value
}

func (v *NullableFormRole) Set(val *FormRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFormRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFormRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormRole(val *FormRole) *NullableFormRole {
	return &NullableFormRole{value: val, isSet: true}
}

func (v NullableFormRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

