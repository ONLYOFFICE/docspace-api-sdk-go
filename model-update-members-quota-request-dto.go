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

// checks if the UpdateMembersQuotaRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMembersQuotaRequestDto{}

// UpdateMembersQuotaRequestDto The request parameters for updating a user quota.
type UpdateMembersQuotaRequestDto struct {
	// The list of user IDs.
	UserIds []string `json:"userIds,omitempty"`
	Quota *UpdateMembersQuotaRequestDtoQuota `json:"quota,omitempty"`
}

// NewUpdateMembersQuotaRequestDto instantiates a new UpdateMembersQuotaRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMembersQuotaRequestDto() *UpdateMembersQuotaRequestDto {
	this := UpdateMembersQuotaRequestDto{}
	return &this
}

// NewUpdateMembersQuotaRequestDtoWithDefaults instantiates a new UpdateMembersQuotaRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMembersQuotaRequestDtoWithDefaults() *UpdateMembersQuotaRequestDto {
	this := UpdateMembersQuotaRequestDto{}
	return &this
}

// GetUserIds returns the UserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMembersQuotaRequestDto) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMembersQuotaRequestDto) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *UpdateMembersQuotaRequestDto) IsUserIdsSet() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *UpdateMembersQuotaRequestDto) SetUserIds(v []string) {
	o.UserIds = v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *UpdateMembersQuotaRequestDto) GetQuota() UpdateMembersQuotaRequestDtoQuota {
	if o == nil || IsNil(o.Quota) {
		var ret UpdateMembersQuotaRequestDtoQuota
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMembersQuotaRequestDto) GetQuotaOk() (*UpdateMembersQuotaRequestDtoQuota, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *UpdateMembersQuotaRequestDto) IsQuotaSet() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given UpdateMembersQuotaRequestDtoQuota and assigns it to the Quota field.
func (o *UpdateMembersQuotaRequestDto) SetQuota(v UpdateMembersQuotaRequestDtoQuota) {
	o.Quota = &v
}

func (o UpdateMembersQuotaRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMembersQuotaRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	return toSerialize, nil
}

type NullableUpdateMembersQuotaRequestDto struct {
	value *UpdateMembersQuotaRequestDto
	isSet bool
}

func (v NullableUpdateMembersQuotaRequestDto) Get() *UpdateMembersQuotaRequestDto {
	return v.value
}

func (v *NullableUpdateMembersQuotaRequestDto) Set(val *UpdateMembersQuotaRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMembersQuotaRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMembersQuotaRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMembersQuotaRequestDto(val *UpdateMembersQuotaRequestDto) *NullableUpdateMembersQuotaRequestDto {
	return &NullableUpdateMembersQuotaRequestDto{value: val, isSet: true}
}

func (v NullableUpdateMembersQuotaRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMembersQuotaRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

