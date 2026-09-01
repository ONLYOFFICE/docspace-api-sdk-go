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

// checks if the DocsCloudQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudQuota{}

// DocsCloudQuota Represents the current user quota of a DocsCloud tenant.
type DocsCloudQuota struct {
	// The editor users.
	Users []DocsCloudQuotaUser `json:"users,omitempty"`
	// The viewer users.
	UsersView []DocsCloudQuotaUser `json:"usersView,omitempty"`
}

// NewDocsCloudQuota instantiates a new DocsCloudQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudQuota() *DocsCloudQuota {
	this := DocsCloudQuota{}
	return &this
}

// NewDocsCloudQuotaWithDefaults instantiates a new DocsCloudQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudQuotaWithDefaults() *DocsCloudQuota {
	this := DocsCloudQuota{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudQuota) GetUsers() []DocsCloudQuotaUser {
	if o == nil {
		var ret []DocsCloudQuotaUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudQuota) GetUsersOk() ([]DocsCloudQuotaUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *DocsCloudQuota) IsUsersSet() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []DocsCloudQuotaUser and assigns it to the Users field.
func (o *DocsCloudQuota) SetUsers(v []DocsCloudQuotaUser) {
	o.Users = v
}

// GetUsersView returns the UsersView field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudQuota) GetUsersView() []DocsCloudQuotaUser {
	if o == nil {
		var ret []DocsCloudQuotaUser
		return ret
	}
	return o.UsersView
}

// GetUsersViewOk returns a tuple with the UsersView field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudQuota) GetUsersViewOk() ([]DocsCloudQuotaUser, bool) {
	if o == nil || IsNil(o.UsersView) {
		return nil, false
	}
	return o.UsersView, true
}

// HasUsersView returns a boolean if a field has been set.
func (o *DocsCloudQuota) IsUsersViewSet() bool {
	if o != nil && !IsNil(o.UsersView) {
		return true
	}

	return false
}

// SetUsersView gets a reference to the given []DocsCloudQuotaUser and assigns it to the UsersView field.
func (o *DocsCloudQuota) SetUsersView(v []DocsCloudQuotaUser) {
	o.UsersView = v
}

func (o DocsCloudQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.UsersView != nil {
		toSerialize["usersView"] = o.UsersView
	}
	return toSerialize, nil
}

type NullableDocsCloudQuota struct {
	value *DocsCloudQuota
	isSet bool
}

func (v NullableDocsCloudQuota) Get() *DocsCloudQuota {
	return v.value
}

func (v *NullableDocsCloudQuota) Set(val *DocsCloudQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudQuota(val *DocsCloudQuota) *NullableDocsCloudQuota {
	return &NullableDocsCloudQuota{value: val, isSet: true}
}

func (v NullableDocsCloudQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

