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
	"bytes"
	"fmt"
)

// checks if the FormRoleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormRoleDto{}

// FormRoleDto The form role parameters.
type FormRoleDto struct {
	// The role name.
	RoleName NullableString `json:"roleName"`
	// The role color.
	RoleColor NullableString `json:"roleColor,omitempty"`
	// The user of the role.
	User *EmployeeFullDto `json:"user,omitempty"`
	// The role sequence.
	Sequence int32 `json:"sequence"`
	// Specifies if the role is submitted.
	Submitted bool `json:"submitted"`
	// The user who stopped the role.
	StopedBy *EmployeeFullDto `json:"stopedBy,omitempty"`
	// The role history.
	History map[string]time.Time `json:"history,omitempty"`
	// The role status.
	RoleStatus *FormFillingStatus `json:"roleStatus,omitempty"`
}

type _FormRoleDto FormRoleDto

// NewFormRoleDto instantiates a new FormRoleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormRoleDto(roleName NullableString, sequence int32, submitted bool) *FormRoleDto {
	this := FormRoleDto{}
	this.RoleName = roleName
	this.Sequence = sequence
	this.Submitted = submitted
	return &this
}

// NewFormRoleDtoWithDefaults instantiates a new FormRoleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormRoleDtoWithDefaults() *FormRoleDto {
	this := FormRoleDto{}
	return &this
}

// GetRoleName returns the RoleName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormRoleDto) GetRoleName() string {
	if o == nil || o.RoleName.Get() == nil {
		var ret string
		return ret
	}

	return *o.RoleName.Get()
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormRoleDto) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleName.Get(), o.RoleName.IsSet()
}

// SetRoleName sets field value
func (o *FormRoleDto) SetRoleName(v string) {
	o.RoleName.Set(&v)
}

// GetRoleColor returns the RoleColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormRoleDto) GetRoleColor() string {
	if o == nil || IsNil(o.RoleColor.Get()) {
		var ret string
		return ret
	}
	return *o.RoleColor.Get()
}

// GetRoleColorOk returns a tuple with the RoleColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormRoleDto) GetRoleColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleColor.Get(), o.RoleColor.IsSet()
}

// HasRoleColor returns a boolean if a field has been set.
func (o *FormRoleDto) IsRoleColorSet() bool {
	if o != nil && o.RoleColor.IsSet() {
		return true
	}

	return false
}

// SetRoleColor gets a reference to the given NullableString and assigns it to the RoleColor field.
func (o *FormRoleDto) SetRoleColor(v string) {
	o.RoleColor.Set(&v)
}
// SetRoleColorNil sets the value for RoleColor to be an explicit nil
func (o *FormRoleDto) SetRoleColorNil() {
	o.RoleColor.Set(nil)
}

// UnsetRoleColor ensures that no value is present for RoleColor, not even an explicit nil
func (o *FormRoleDto) UnsetRoleColor() {
	o.RoleColor.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *FormRoleDto) GetUser() EmployeeFullDto {
	if o == nil || IsNil(o.User) {
		var ret EmployeeFullDto
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleDto) GetUserOk() (*EmployeeFullDto, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *FormRoleDto) IsUserSet() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given EmployeeFullDto and assigns it to the User field.
func (o *FormRoleDto) SetUser(v EmployeeFullDto) {
	o.User = &v
}

// GetSequence returns the Sequence field value
func (o *FormRoleDto) GetSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *FormRoleDto) GetSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *FormRoleDto) SetSequence(v int32) {
	o.Sequence = v
}

// GetSubmitted returns the Submitted field value
func (o *FormRoleDto) GetSubmitted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value
// and a boolean to check if the value has been set.
func (o *FormRoleDto) GetSubmittedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Submitted, true
}

// SetSubmitted sets field value
func (o *FormRoleDto) SetSubmitted(v bool) {
	o.Submitted = v
}

// GetStopedBy returns the StopedBy field value if set, zero value otherwise.
func (o *FormRoleDto) GetStopedBy() EmployeeFullDto {
	if o == nil || IsNil(o.StopedBy) {
		var ret EmployeeFullDto
		return ret
	}
	return *o.StopedBy
}

// GetStopedByOk returns a tuple with the StopedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleDto) GetStopedByOk() (*EmployeeFullDto, bool) {
	if o == nil || IsNil(o.StopedBy) {
		return nil, false
	}
	return o.StopedBy, true
}

// HasStopedBy returns a boolean if a field has been set.
func (o *FormRoleDto) IsStopedBySet() bool {
	if o != nil && !IsNil(o.StopedBy) {
		return true
	}

	return false
}

// SetStopedBy gets a reference to the given EmployeeFullDto and assigns it to the StopedBy field.
func (o *FormRoleDto) SetStopedBy(v EmployeeFullDto) {
	o.StopedBy = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *FormRoleDto) GetHistory() map[string]time.Time {
	if o == nil || IsNil(o.History) {
		var ret map[string]time.Time
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleDto) GetHistoryOk() (map[string]time.Time, bool) {
	if o == nil || IsNil(o.History) {
		return map[string]time.Time{}, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *FormRoleDto) IsHistorySet() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given map[string]time.Time and assigns it to the History field.
func (o *FormRoleDto) SetHistory(v map[string]time.Time) {
	o.History = v
}

// GetRoleStatus returns the RoleStatus field value if set, zero value otherwise.
func (o *FormRoleDto) GetRoleStatus() FormFillingStatus {
	if o == nil || IsNil(o.RoleStatus) {
		var ret FormFillingStatus
		return ret
	}
	return *o.RoleStatus
}

// GetRoleStatusOk returns a tuple with the RoleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormRoleDto) GetRoleStatusOk() (*FormFillingStatus, bool) {
	if o == nil || IsNil(o.RoleStatus) {
		return nil, false
	}
	return o.RoleStatus, true
}

// HasRoleStatus returns a boolean if a field has been set.
func (o *FormRoleDto) IsRoleStatusSet() bool {
	if o != nil && !IsNil(o.RoleStatus) {
		return true
	}

	return false
}

// SetRoleStatus gets a reference to the given FormFillingStatus and assigns it to the RoleStatus field.
func (o *FormRoleDto) SetRoleStatus(v FormFillingStatus) {
	o.RoleStatus = &v
}

func (o FormRoleDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormRoleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleName"] = o.RoleName.Get()
	if o.RoleColor.IsSet() {
		toSerialize["roleColor"] = o.RoleColor.Get()
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	toSerialize["sequence"] = o.Sequence
	toSerialize["submitted"] = o.Submitted
	if !IsNil(o.StopedBy) {
		toSerialize["stopedBy"] = o.StopedBy
	}
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	if !IsNil(o.RoleStatus) {
		toSerialize["roleStatus"] = o.RoleStatus
	}
	return toSerialize, nil
}

func (o *FormRoleDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roleName",
		"sequence",
		"submitted",
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

	varFormRoleDto := _FormRoleDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFormRoleDto)

	if err != nil {
		return err
	}

	*o = FormRoleDto(varFormRoleDto)

	return err
}

type NullableFormRoleDto struct {
	value *FormRoleDto
	isSet bool
}

func (v NullableFormRoleDto) Get() *FormRoleDto {
	return v.value
}

func (v *NullableFormRoleDto) Set(val *FormRoleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFormRoleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFormRoleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormRoleDto(val *FormRoleDto) *NullableFormRoleDto {
	return &NullableFormRoleDto{value: val, isSet: true}
}

func (v NullableFormRoleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormRoleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

