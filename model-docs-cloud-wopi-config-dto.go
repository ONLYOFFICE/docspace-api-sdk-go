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

// checks if the DocsCloudWopiConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudWopiConfigDto{}

// DocsCloudWopiConfigDto Represents the WOPI configuration of a Docs Connect tenant.
type DocsCloudWopiConfigDto struct {
	// Whether WOPI is enabled.
	Enable *bool `json:"enable,omitempty"`
}

// NewDocsCloudWopiConfigDto instantiates a new DocsCloudWopiConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudWopiConfigDto() *DocsCloudWopiConfigDto {
	this := DocsCloudWopiConfigDto{}
	return &this
}

// NewDocsCloudWopiConfigDtoWithDefaults instantiates a new DocsCloudWopiConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudWopiConfigDtoWithDefaults() *DocsCloudWopiConfigDto {
	this := DocsCloudWopiConfigDto{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *DocsCloudWopiConfigDto) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudWopiConfigDto) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *DocsCloudWopiConfigDto) IsEnableSet() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *DocsCloudWopiConfigDto) SetEnable(v bool) {
	o.Enable = &v
}

func (o DocsCloudWopiConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudWopiConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	return toSerialize, nil
}

type NullableDocsCloudWopiConfigDto struct {
	value *DocsCloudWopiConfigDto
	isSet bool
}

func (v NullableDocsCloudWopiConfigDto) Get() *DocsCloudWopiConfigDto {
	return v.value
}

func (v *NullableDocsCloudWopiConfigDto) Set(val *DocsCloudWopiConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudWopiConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudWopiConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudWopiConfigDto(val *DocsCloudWopiConfigDto) *NullableDocsCloudWopiConfigDto {
	return &NullableDocsCloudWopiConfigDto{value: val, isSet: true}
}

func (v NullableDocsCloudWopiConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudWopiConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

