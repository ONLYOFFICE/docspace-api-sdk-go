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

// checks if the StartUpdateUserTypeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartUpdateUserTypeDto{}

// StartUpdateUserTypeDto The parameters for updating the type of the user or guest when reassigning rooms and shared files.
type StartUpdateUserTypeDto struct {
	// The type to convert the account to. Only `Guest` and `User` are accepted, because they are the types that  cannot own rooms; `RoomAdmin`, `DocSpaceAdmin` and `All` are rejected here and belong to  `PUT api/2.0/people/type/{type}`.
	Type *EmployeeType `json:"type,omitempty"`
	// The ID of the account being converted. It has to be an active account other than the caller, and only the  portal owner may pass the ID of a DocSpace administrator.
	UserId *string `json:"userId,omitempty"`
	// The ID of the administrator who receives the rooms and the shared files of the converted account. It has to be  an active room admin or DocSpace admin other than the converted account, and when it is omitted the data goes  to the caller.
	ReassignUserId NullableString `json:"reassignUserId,omitempty"`
}

// NewStartUpdateUserTypeDto instantiates a new StartUpdateUserTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartUpdateUserTypeDto() *StartUpdateUserTypeDto {
	this := StartUpdateUserTypeDto{}
	return &this
}

// NewStartUpdateUserTypeDtoWithDefaults instantiates a new StartUpdateUserTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartUpdateUserTypeDtoWithDefaults() *StartUpdateUserTypeDto {
	this := StartUpdateUserTypeDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StartUpdateUserTypeDto) GetType() EmployeeType {
	if o == nil || IsNil(o.Type) {
		var ret EmployeeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUpdateUserTypeDto) GetTypeOk() (*EmployeeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StartUpdateUserTypeDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EmployeeType and assigns it to the Type field.
func (o *StartUpdateUserTypeDto) SetType(v EmployeeType) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *StartUpdateUserTypeDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUpdateUserTypeDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *StartUpdateUserTypeDto) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *StartUpdateUserTypeDto) SetUserId(v string) {
	o.UserId = &v
}

// GetReassignUserId returns the ReassignUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartUpdateUserTypeDto) GetReassignUserId() string {
	if o == nil || IsNil(o.ReassignUserId.Get()) {
		var ret string
		return ret
	}
	return *o.ReassignUserId.Get()
}

// GetReassignUserIdOk returns a tuple with the ReassignUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartUpdateUserTypeDto) GetReassignUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReassignUserId.Get(), o.ReassignUserId.IsSet()
}

// HasReassignUserId returns a boolean if a field has been set.
func (o *StartUpdateUserTypeDto) IsReassignUserIdSet() bool {
	if o != nil && o.ReassignUserId.IsSet() {
		return true
	}

	return false
}

// SetReassignUserId gets a reference to the given NullableString and assigns it to the ReassignUserId field.
func (o *StartUpdateUserTypeDto) SetReassignUserId(v string) {
	o.ReassignUserId.Set(&v)
}
// SetReassignUserIdNil sets the value for ReassignUserId to be an explicit nil
func (o *StartUpdateUserTypeDto) SetReassignUserIdNil() {
	o.ReassignUserId.Set(nil)
}

// UnsetReassignUserId ensures that no value is present for ReassignUserId, not even an explicit nil
func (o *StartUpdateUserTypeDto) UnsetReassignUserId() {
	o.ReassignUserId.Unset()
}

func (o StartUpdateUserTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartUpdateUserTypeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if o.ReassignUserId.IsSet() {
		toSerialize["reassignUserId"] = o.ReassignUserId.Get()
	}
	return toSerialize, nil
}

type NullableStartUpdateUserTypeDto struct {
	value *StartUpdateUserTypeDto
	isSet bool
}

func (v NullableStartUpdateUserTypeDto) Get() *StartUpdateUserTypeDto {
	return v.value
}

func (v *NullableStartUpdateUserTypeDto) Set(val *StartUpdateUserTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStartUpdateUserTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStartUpdateUserTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartUpdateUserTypeDto(val *StartUpdateUserTypeDto) *NullableStartUpdateUserTypeDto {
	return &NullableStartUpdateUserTypeDto{value: val, isSet: true}
}

func (v NullableStartUpdateUserTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartUpdateUserTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

