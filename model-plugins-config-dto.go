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

// checks if the PluginsConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginsConfigDto{}

// PluginsConfigDto Which editor add-ons the portal connects. It currently connects none.
type PluginsConfigDto struct {
	// The array of absolute URLs to the plugin configuration files.
	PluginsData []string `json:"pluginsData,omitempty"`
}

// NewPluginsConfigDto instantiates a new PluginsConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginsConfigDto() *PluginsConfigDto {
	this := PluginsConfigDto{}
	return &this
}

// NewPluginsConfigDtoWithDefaults instantiates a new PluginsConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsConfigDtoWithDefaults() *PluginsConfigDto {
	this := PluginsConfigDto{}
	return &this
}

// GetPluginsData returns the PluginsData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginsConfigDto) GetPluginsData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PluginsData
}

// GetPluginsDataOk returns a tuple with the PluginsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginsConfigDto) GetPluginsDataOk() ([]string, bool) {
	if o == nil || IsNil(o.PluginsData) {
		return nil, false
	}
	return o.PluginsData, true
}

// HasPluginsData returns a boolean if a field has been set.
func (o *PluginsConfigDto) IsPluginsDataSet() bool {
	if o != nil && !IsNil(o.PluginsData) {
		return true
	}

	return false
}

// SetPluginsData gets a reference to the given []string and assigns it to the PluginsData field.
func (o *PluginsConfigDto) SetPluginsData(v []string) {
	o.PluginsData = v
}

func (o PluginsConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginsConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PluginsData != nil {
		toSerialize["pluginsData"] = o.PluginsData
	}
	return toSerialize, nil
}

type NullablePluginsConfigDto struct {
	value *PluginsConfigDto
	isSet bool
}

func (v NullablePluginsConfigDto) Get() *PluginsConfigDto {
	return v.value
}

func (v *NullablePluginsConfigDto) Set(val *PluginsConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginsConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginsConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginsConfigDto(val *PluginsConfigDto) *NullablePluginsConfigDto {
	return &NullablePluginsConfigDto{value: val, isSet: true}
}

func (v NullablePluginsConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginsConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

