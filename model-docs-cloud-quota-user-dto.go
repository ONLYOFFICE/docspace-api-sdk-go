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

// checks if the DocsCloudQuotaUserDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudQuotaUserDto{}

// DocsCloudQuotaUserDto Represents a single user entry of a Docs Connect quota.
type DocsCloudQuotaUserDto struct {
	// The user ID.
	UserId NullableString `json:"userId,omitempty"`
	// The expiration date of the user.
	Expire NullableString `json:"expire,omitempty"`
}

// NewDocsCloudQuotaUserDto instantiates a new DocsCloudQuotaUserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudQuotaUserDto() *DocsCloudQuotaUserDto {
	this := DocsCloudQuotaUserDto{}
	return &this
}

// NewDocsCloudQuotaUserDtoWithDefaults instantiates a new DocsCloudQuotaUserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudQuotaUserDtoWithDefaults() *DocsCloudQuotaUserDto {
	this := DocsCloudQuotaUserDto{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudQuotaUserDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudQuotaUserDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *DocsCloudQuotaUserDto) IsUserIdSet() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *DocsCloudQuotaUserDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *DocsCloudQuotaUserDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *DocsCloudQuotaUserDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetExpire returns the Expire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudQuotaUserDto) GetExpire() string {
	if o == nil || IsNil(o.Expire.Get()) {
		var ret string
		return ret
	}
	return *o.Expire.Get()
}

// GetExpireOk returns a tuple with the Expire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudQuotaUserDto) GetExpireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expire.Get(), o.Expire.IsSet()
}

// HasExpire returns a boolean if a field has been set.
func (o *DocsCloudQuotaUserDto) IsExpireSet() bool {
	if o != nil && o.Expire.IsSet() {
		return true
	}

	return false
}

// SetExpire gets a reference to the given NullableString and assigns it to the Expire field.
func (o *DocsCloudQuotaUserDto) SetExpire(v string) {
	o.Expire.Set(&v)
}
// SetExpireNil sets the value for Expire to be an explicit nil
func (o *DocsCloudQuotaUserDto) SetExpireNil() {
	o.Expire.Set(nil)
}

// UnsetExpire ensures that no value is present for Expire, not even an explicit nil
func (o *DocsCloudQuotaUserDto) UnsetExpire() {
	o.Expire.Unset()
}

func (o DocsCloudQuotaUserDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudQuotaUserDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.Expire.IsSet() {
		toSerialize["expire"] = o.Expire.Get()
	}
	return toSerialize, nil
}

type NullableDocsCloudQuotaUserDto struct {
	value *DocsCloudQuotaUserDto
	isSet bool
}

func (v NullableDocsCloudQuotaUserDto) Get() *DocsCloudQuotaUserDto {
	return v.value
}

func (v *NullableDocsCloudQuotaUserDto) Set(val *DocsCloudQuotaUserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudQuotaUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudQuotaUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudQuotaUserDto(val *DocsCloudQuotaUserDto) *NullableDocsCloudQuotaUserDto {
	return &NullableDocsCloudQuotaUserDto{value: val, isSet: true}
}

func (v NullableDocsCloudQuotaUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudQuotaUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

