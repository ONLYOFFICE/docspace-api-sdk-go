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

// checks if the WebItemsSecurityRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebItemsSecurityRequestDto{}

// WebItemsSecurityRequestDto The modules switched on or off together, one entry per module.
type WebItemsSecurityRequestDto struct {
	// The modules to switch, each entry pairing a module GUID as its `key` with the new enabled flag as its  `value`. A key that is not a GUID fails the whole request as invalid, and a module listed twice is applied  once, from its first entry. No allow-list travels here: switching a product module on restores the users and  groups it was last restricted to, and everything else is stored as a plain allow or deny for everyone.
	Items []ItemKeyValuePairStringBoolean `json:"items,omitempty"`
}

// NewWebItemsSecurityRequestDto instantiates a new WebItemsSecurityRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebItemsSecurityRequestDto() *WebItemsSecurityRequestDto {
	this := WebItemsSecurityRequestDto{}
	return &this
}

// NewWebItemsSecurityRequestDtoWithDefaults instantiates a new WebItemsSecurityRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebItemsSecurityRequestDtoWithDefaults() *WebItemsSecurityRequestDto {
	this := WebItemsSecurityRequestDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebItemsSecurityRequestDto) GetItems() []ItemKeyValuePairStringBoolean {
	if o == nil {
		var ret []ItemKeyValuePairStringBoolean
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebItemsSecurityRequestDto) GetItemsOk() ([]ItemKeyValuePairStringBoolean, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *WebItemsSecurityRequestDto) IsItemsSet() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ItemKeyValuePairStringBoolean and assigns it to the Items field.
func (o *WebItemsSecurityRequestDto) SetItems(v []ItemKeyValuePairStringBoolean) {
	o.Items = v
}

func (o WebItemsSecurityRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebItemsSecurityRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableWebItemsSecurityRequestDto struct {
	value *WebItemsSecurityRequestDto
	isSet bool
}

func (v NullableWebItemsSecurityRequestDto) Get() *WebItemsSecurityRequestDto {
	return v.value
}

func (v *NullableWebItemsSecurityRequestDto) Set(val *WebItemsSecurityRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebItemsSecurityRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebItemsSecurityRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebItemsSecurityRequestDto(val *WebItemsSecurityRequestDto) *NullableWebItemsSecurityRequestDto {
	return &NullableWebItemsSecurityRequestDto{value: val, isSet: true}
}

func (v NullableWebItemsSecurityRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebItemsSecurityRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

