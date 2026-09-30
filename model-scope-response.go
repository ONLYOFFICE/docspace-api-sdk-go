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

// checks if the ScopeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopeResponse{}

// ScopeResponse One scope from the tenant scope catalogue, as it may be requested by a client.
type ScopeResponse struct {
	// The scope exactly as it is written in an authorization request, for example files:read or openid.
	Name *string `json:"name,omitempty"`
	// The area of the portal the scope belongs to, which is what groups the scopes on the consent screen: files, rooms, contacts, profiles or openid.
	Group *string `json:"group,omitempty"`
	// What the scope allows inside its group: read for read-only access, write for changes, and openid for the identity scope itself.
	Type *string `json:"type,omitempty"`
}

// NewScopeResponse instantiates a new ScopeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeResponse() *ScopeResponse {
	this := ScopeResponse{}
	return &this
}

// NewScopeResponseWithDefaults instantiates a new ScopeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeResponseWithDefaults() *ScopeResponse {
	this := ScopeResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScopeResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScopeResponse) IsNameSet() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScopeResponse) SetName(v string) {
	o.Name = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ScopeResponse) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeResponse) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ScopeResponse) IsGroupSet() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ScopeResponse) SetGroup(v string) {
	o.Group = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScopeResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScopeResponse) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScopeResponse) SetType(v string) {
	o.Type = &v
}

func (o ScopeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableScopeResponse struct {
	value *ScopeResponse
	isSet bool
}

func (v NullableScopeResponse) Get() *ScopeResponse {
	return v.value
}

func (v *NullableScopeResponse) Set(val *ScopeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeResponse(val *ScopeResponse) *NullableScopeResponse {
	return &NullableScopeResponse{value: val, isSet: true}
}

func (v NullableScopeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

