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

// checks if the UpdateMembersRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMembersRequestDto{}

// UpdateMembersRequestDto The request parameters for updating the user information.
type UpdateMembersRequestDto struct {
	// The list of user IDs.
	UserIds []string `json:"userIds,omitempty"`
	// Specifies whether to resend invitation letters to all the users or not.
	ResendAll *bool `json:"resendAll,omitempty"`
}

// NewUpdateMembersRequestDto instantiates a new UpdateMembersRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMembersRequestDto() *UpdateMembersRequestDto {
	this := UpdateMembersRequestDto{}
	return &this
}

// NewUpdateMembersRequestDtoWithDefaults instantiates a new UpdateMembersRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMembersRequestDtoWithDefaults() *UpdateMembersRequestDto {
	this := UpdateMembersRequestDto{}
	return &this
}

// GetUserIds returns the UserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMembersRequestDto) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMembersRequestDto) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *UpdateMembersRequestDto) IsUserIdsSet() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *UpdateMembersRequestDto) SetUserIds(v []string) {
	o.UserIds = v
}

// GetResendAll returns the ResendAll field value if set, zero value otherwise.
func (o *UpdateMembersRequestDto) GetResendAll() bool {
	if o == nil || IsNil(o.ResendAll) {
		var ret bool
		return ret
	}
	return *o.ResendAll
}

// GetResendAllOk returns a tuple with the ResendAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMembersRequestDto) GetResendAllOk() (*bool, bool) {
	if o == nil || IsNil(o.ResendAll) {
		return nil, false
	}
	return o.ResendAll, true
}

// HasResendAll returns a boolean if a field has been set.
func (o *UpdateMembersRequestDto) IsResendAllSet() bool {
	if o != nil && !IsNil(o.ResendAll) {
		return true
	}

	return false
}

// SetResendAll gets a reference to the given bool and assigns it to the ResendAll field.
func (o *UpdateMembersRequestDto) SetResendAll(v bool) {
	o.ResendAll = &v
}

func (o UpdateMembersRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMembersRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	if !IsNil(o.ResendAll) {
		toSerialize["resendAll"] = o.ResendAll
	}
	return toSerialize, nil
}

type NullableUpdateMembersRequestDto struct {
	value *UpdateMembersRequestDto
	isSet bool
}

func (v NullableUpdateMembersRequestDto) Get() *UpdateMembersRequestDto {
	return v.value
}

func (v *NullableUpdateMembersRequestDto) Set(val *UpdateMembersRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMembersRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMembersRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMembersRequestDto(val *UpdateMembersRequestDto) *NullableUpdateMembersRequestDto {
	return &NullableUpdateMembersRequestDto{value: val, isSet: true}
}

func (v NullableUpdateMembersRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMembersRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

