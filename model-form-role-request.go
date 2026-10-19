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

// checks if the FormRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormRoleRequest{}

// FormRoleRequest One role of a form and the account that fills it.
type FormRoleRequest struct {
	// The ID of the room the form is in. It is stored with the role as sent, so pass the room the form lives in.
	RoomId *int32 `json:"roomId,omitempty"`
	// The name of a role the form defines, such as the one the form author gave a group of fields.
	RoleName NullableString `json:"roleName,omitempty"`
	// The color the editor marks the fields of this role with, as a hex code.
	RoleColor NullableString `json:"roleColor,omitempty"`
	// The account that fills this role. It is notified once filling starts, unless it is the caller.
	UserId *string `json:"userId,omitempty"`
	// Accepted for compatibility and ignored: the position of the role in the list sets the filling order.
	Sequence *int32 `json:"sequence,omitempty"`
	// Whether this role counts as already submitted. It is stored as sent; send false when filling starts.
	Submitted *bool `json:"submitted,omitempty"`
	// Accepted for compatibility and ignored: the portal records when the role is opened.
	OpenedAt *time.Time `json:"openedAt,omitempty"`
	// Accepted for compatibility and ignored: the portal records when the role is submitted.
	SubmissionDate *time.Time `json:"submissionDate,omitempty"`
}

// NewFormRoleRequest instantiates a new FormRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormRoleRequest() *FormRoleRequest {
	this := FormRoleRequest{}
	return &this
}

// NewFormRoleRequestWithDefaults instantiates a new FormRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormRoleRequestWithDefaults() *FormRoleRequest {
	this := FormRoleRequest{}
	return &this
}

// GetRoomId returns the RoomId field value if set, zero value otherwise.
func (o *FormRoleRequest) GetRoomId() int32 {
	if o == nil || IsNil(o.RoomId) {
		var ret int32
		return ret
	}
	return *o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleRequest) GetRoomIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RoomId) {
		return nil, false
	}
	return o.RoomId, true
}

// HasRoomId returns a boolean if a field has been set.
func (o *FormRoleRequest) IsRoomIdSet() bool {
	if o != nil && !IsNil(o.RoomId) {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given int32 and assigns it to the RoomId field.
func (o *FormRoleRequest) SetRoomId(v int32) {
	o.RoomId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormRoleRequest) GetRoleName() string {
	if o == nil || IsNil(o.RoleName.Get()) {
		var ret string
		return ret
	}
	return *o.RoleName.Get()
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormRoleRequest) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleName.Get(), o.RoleName.IsSet()
}

// HasRoleName returns a boolean if a field has been set.
func (o *FormRoleRequest) IsRoleNameSet() bool {
	if o != nil && o.RoleName.IsSet() {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given NullableString and assigns it to the RoleName field.
func (o *FormRoleRequest) SetRoleName(v string) {
	o.RoleName.Set(&v)
}
// SetRoleNameNil sets the value for RoleName to be an explicit nil
func (o *FormRoleRequest) SetRoleNameNil() {
	o.RoleName.Set(nil)
}

// UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
func (o *FormRoleRequest) UnsetRoleName() {
	o.RoleName.Unset()
}

// GetRoleColor returns the RoleColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormRoleRequest) GetRoleColor() string {
	if o == nil || IsNil(o.RoleColor.Get()) {
		var ret string
		return ret
	}
	return *o.RoleColor.Get()
}

// GetRoleColorOk returns a tuple with the RoleColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormRoleRequest) GetRoleColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleColor.Get(), o.RoleColor.IsSet()
}

// HasRoleColor returns a boolean if a field has been set.
func (o *FormRoleRequest) IsRoleColorSet() bool {
	if o != nil && o.RoleColor.IsSet() {
		return true
	}

	return false
}

// SetRoleColor gets a reference to the given NullableString and assigns it to the RoleColor field.
func (o *FormRoleRequest) SetRoleColor(v string) {
	o.RoleColor.Set(&v)
}
// SetRoleColorNil sets the value for RoleColor to be an explicit nil
func (o *FormRoleRequest) SetRoleColorNil() {
	o.RoleColor.Set(nil)
}

// UnsetRoleColor ensures that no value is present for RoleColor, not even an explicit nil
func (o *FormRoleRequest) UnsetRoleColor() {
	o.RoleColor.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FormRoleRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FormRoleRequest) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FormRoleRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *FormRoleRequest) GetSequence() int32 {
	if o == nil || IsNil(o.Sequence) {
		var ret int32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleRequest) GetSequenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *FormRoleRequest) IsSequenceSet() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int32 and assigns it to the Sequence field.
func (o *FormRoleRequest) SetSequence(v int32) {
	o.Sequence = &v
}

// GetSubmitted returns the Submitted field value if set, zero value otherwise.
func (o *FormRoleRequest) GetSubmitted() bool {
	if o == nil || IsNil(o.Submitted) {
		var ret bool
		return ret
	}
	return *o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleRequest) GetSubmittedOk() (*bool, bool) {
	if o == nil || IsNil(o.Submitted) {
		return nil, false
	}
	return o.Submitted, true
}

// HasSubmitted returns a boolean if a field has been set.
func (o *FormRoleRequest) IsSubmittedSet() bool {
	if o != nil && !IsNil(o.Submitted) {
		return true
	}

	return false
}

// SetSubmitted gets a reference to the given bool and assigns it to the Submitted field.
func (o *FormRoleRequest) SetSubmitted(v bool) {
	o.Submitted = &v
}

// GetOpenedAt returns the OpenedAt field value if set, zero value otherwise.
func (o *FormRoleRequest) GetOpenedAt() time.Time {
	if o == nil || IsNil(o.OpenedAt) {
		var ret time.Time
		return ret
	}
	return *o.OpenedAt
}

// GetOpenedAtOk returns a tuple with the OpenedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleRequest) GetOpenedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OpenedAt) {
		return nil, false
	}
	return o.OpenedAt, true
}

// HasOpenedAt returns a boolean if a field has been set.
func (o *FormRoleRequest) IsOpenedAtSet() bool {
	if o != nil && !IsNil(o.OpenedAt) {
		return true
	}

	return false
}

// SetOpenedAt gets a reference to the given time.Time and assigns it to the OpenedAt field.
func (o *FormRoleRequest) SetOpenedAt(v time.Time) {
	o.OpenedAt = &v
}

// GetSubmissionDate returns the SubmissionDate field value if set, zero value otherwise.
func (o *FormRoleRequest) GetSubmissionDate() time.Time {
	if o == nil || IsNil(o.SubmissionDate) {
		var ret time.Time
		return ret
	}
	return *o.SubmissionDate
}

// GetSubmissionDateOk returns a tuple with the SubmissionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleRequest) GetSubmissionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmissionDate) {
		return nil, false
	}
	return o.SubmissionDate, true
}

// HasSubmissionDate returns a boolean if a field has been set.
func (o *FormRoleRequest) IsSubmissionDateSet() bool {
	if o != nil && !IsNil(o.SubmissionDate) {
		return true
	}

	return false
}

// SetSubmissionDate gets a reference to the given time.Time and assigns it to the SubmissionDate field.
func (o *FormRoleRequest) SetSubmissionDate(v time.Time) {
	o.SubmissionDate = &v
}

func (o FormRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormRoleRequest) ToMap() (map[string]interface{}, error) {
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

type NullableFormRoleRequest struct {
	value *FormRoleRequest
	isSet bool
}

func (v NullableFormRoleRequest) Get() *FormRoleRequest {
	return v.value
}

func (v *NullableFormRoleRequest) Set(val *FormRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFormRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFormRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormRoleRequest(val *FormRoleRequest) *NullableFormRoleRequest {
	return &NullableFormRoleRequest{value: val, isSet: true}
}

func (v NullableFormRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

