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

// checks if the MembersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembersRequest{}

// MembersRequest The accounts a member operation applies to.
type MembersRequest struct {
	// The accounts the operation applies to. When adding or replacing members, an account that is a guest, is  disabled or does not exist is skipped without an error; when removing them, an ID that is not a member is  skipped as well.
	Members []string `json:"members,omitempty"`
}

// NewMembersRequest instantiates a new MembersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembersRequest() *MembersRequest {
	this := MembersRequest{}
	return &this
}

// NewMembersRequestWithDefaults instantiates a new MembersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembersRequestWithDefaults() *MembersRequest {
	this := MembersRequest{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MembersRequest) GetMembers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MembersRequest) GetMembersOk() ([]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *MembersRequest) IsMembersSet() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *MembersRequest) SetMembers(v []string) {
	o.Members = v
}

func (o MembersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return toSerialize, nil
}

type NullableMembersRequest struct {
	value *MembersRequest
	isSet bool
}

func (v NullableMembersRequest) Get() *MembersRequest {
	return v.value
}

func (v *NullableMembersRequest) Set(val *MembersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMembersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMembersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembersRequest(val *MembersRequest) *NullableMembersRequest {
	return &NullableMembersRequest{value: val, isSet: true}
}

func (v NullableMembersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

