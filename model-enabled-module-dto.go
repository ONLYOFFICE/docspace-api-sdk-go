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

// checks if the EnabledModuleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnabledModuleDto{}

// EnabledModuleDto One portal module the calling user may open.
type EnabledModuleDto struct {
	// The module's product class name, HTML-escaped. It is a display-oriented identifier and not the GUID the  access-settings operations work with, so it must not be passed to `GET api/2.0/settings/security/{id}`.
	Id NullableString `json:"id,omitempty"`
	// The module name in the portal language, HTML-escaped and ready to be rendered as text.
	Title NullableString `json:"title,omitempty"`
}

// NewEnabledModuleDto instantiates a new EnabledModuleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnabledModuleDto() *EnabledModuleDto {
	this := EnabledModuleDto{}
	return &this
}

// NewEnabledModuleDtoWithDefaults instantiates a new EnabledModuleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnabledModuleDtoWithDefaults() *EnabledModuleDto {
	this := EnabledModuleDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnabledModuleDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnabledModuleDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *EnabledModuleDto) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *EnabledModuleDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *EnabledModuleDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *EnabledModuleDto) UnsetId() {
	o.Id.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnabledModuleDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnabledModuleDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *EnabledModuleDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *EnabledModuleDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *EnabledModuleDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *EnabledModuleDto) UnsetTitle() {
	o.Title.Unset()
}

func (o EnabledModuleDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnabledModuleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

type NullableEnabledModuleDto struct {
	value *EnabledModuleDto
	isSet bool
}

func (v NullableEnabledModuleDto) Get() *EnabledModuleDto {
	return v.value
}

func (v *NullableEnabledModuleDto) Set(val *EnabledModuleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEnabledModuleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEnabledModuleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnabledModuleDto(val *EnabledModuleDto) *NullableEnabledModuleDto {
	return &NullableEnabledModuleDto{value: val, isSet: true}
}

func (v NullableEnabledModuleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnabledModuleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

