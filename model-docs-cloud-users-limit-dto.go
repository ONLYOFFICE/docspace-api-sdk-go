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

// checks if the DocsCloudUsersLimitDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudUsersLimitDto{}

// DocsCloudUsersLimitDto Represents the user limits of a Docs Connect license.
type DocsCloudUsersLimitDto struct {
	// The maximum number of users who can edit documents.
	Edit *int32 `json:"edit,omitempty"`
	// The maximum number of users who can view documents.
	View *int32 `json:"view,omitempty"`
}

// NewDocsCloudUsersLimitDto instantiates a new DocsCloudUsersLimitDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudUsersLimitDto() *DocsCloudUsersLimitDto {
	this := DocsCloudUsersLimitDto{}
	return &this
}

// NewDocsCloudUsersLimitDtoWithDefaults instantiates a new DocsCloudUsersLimitDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudUsersLimitDtoWithDefaults() *DocsCloudUsersLimitDto {
	this := DocsCloudUsersLimitDto{}
	return &this
}

// GetEdit returns the Edit field value if set, zero value otherwise.
func (o *DocsCloudUsersLimitDto) GetEdit() int32 {
	if o == nil || IsNil(o.Edit) {
		var ret int32
		return ret
	}
	return *o.Edit
}

// GetEditOk returns a tuple with the Edit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUsersLimitDto) GetEditOk() (*int32, bool) {
	if o == nil || IsNil(o.Edit) {
		return nil, false
	}
	return o.Edit, true
}

// HasEdit returns a boolean if a field has been set.
func (o *DocsCloudUsersLimitDto) IsEditSet() bool {
	if o != nil && !IsNil(o.Edit) {
		return true
	}

	return false
}

// SetEdit gets a reference to the given int32 and assigns it to the Edit field.
func (o *DocsCloudUsersLimitDto) SetEdit(v int32) {
	o.Edit = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *DocsCloudUsersLimitDto) GetView() int32 {
	if o == nil || IsNil(o.View) {
		var ret int32
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUsersLimitDto) GetViewOk() (*int32, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *DocsCloudUsersLimitDto) IsViewSet() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given int32 and assigns it to the View field.
func (o *DocsCloudUsersLimitDto) SetView(v int32) {
	o.View = &v
}

func (o DocsCloudUsersLimitDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudUsersLimitDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Edit) {
		toSerialize["edit"] = o.Edit
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	return toSerialize, nil
}

type NullableDocsCloudUsersLimitDto struct {
	value *DocsCloudUsersLimitDto
	isSet bool
}

func (v NullableDocsCloudUsersLimitDto) Get() *DocsCloudUsersLimitDto {
	return v.value
}

func (v *NullableDocsCloudUsersLimitDto) Set(val *DocsCloudUsersLimitDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudUsersLimitDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudUsersLimitDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudUsersLimitDto(val *DocsCloudUsersLimitDto) *NullableDocsCloudUsersLimitDto {
	return &NullableDocsCloudUsersLimitDto{value: val, isSet: true}
}

func (v NullableDocsCloudUsersLimitDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudUsersLimitDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

