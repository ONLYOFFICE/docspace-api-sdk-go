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

// checks if the WebItemsSecurityRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebItemsSecurityRequestsDto{}

// WebItemsSecurityRequestsDto The request parameters for configuring security settings across multiple web modules.
type WebItemsSecurityRequestsDto struct {
	// The list of module security configurations.
	Items []ItemKeyValuePairStringBoolean `json:"items,omitempty"`
}

// NewWebItemsSecurityRequestsDto instantiates a new WebItemsSecurityRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebItemsSecurityRequestsDto() *WebItemsSecurityRequestsDto {
	this := WebItemsSecurityRequestsDto{}
	return &this
}

// NewWebItemsSecurityRequestsDtoWithDefaults instantiates a new WebItemsSecurityRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebItemsSecurityRequestsDtoWithDefaults() *WebItemsSecurityRequestsDto {
	this := WebItemsSecurityRequestsDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebItemsSecurityRequestsDto) GetItems() []ItemKeyValuePairStringBoolean {
	if o == nil {
		var ret []ItemKeyValuePairStringBoolean
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebItemsSecurityRequestsDto) GetItemsOk() ([]ItemKeyValuePairStringBoolean, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *WebItemsSecurityRequestsDto) IsItemsSet() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ItemKeyValuePairStringBoolean and assigns it to the Items field.
func (o *WebItemsSecurityRequestsDto) SetItems(v []ItemKeyValuePairStringBoolean) {
	o.Items = v
}

func (o WebItemsSecurityRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebItemsSecurityRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableWebItemsSecurityRequestsDto struct {
	value *WebItemsSecurityRequestsDto
	isSet bool
}

func (v NullableWebItemsSecurityRequestsDto) Get() *WebItemsSecurityRequestsDto {
	return v.value
}

func (v *NullableWebItemsSecurityRequestsDto) Set(val *WebItemsSecurityRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebItemsSecurityRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebItemsSecurityRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebItemsSecurityRequestsDto(val *WebItemsSecurityRequestsDto) *NullableWebItemsSecurityRequestsDto {
	return &NullableWebItemsSecurityRequestsDto{value: val, isSet: true}
}

func (v NullableWebItemsSecurityRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebItemsSecurityRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

