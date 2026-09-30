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

// checks if the PluginsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginsDto{}

// PluginsDto What the installation allows to be done with web plugins.
type PluginsDto struct {
	// Whether web plugins run on this portal at all. While it is `false` the operations under  `api/2.0/settings/webplugins` are of no use, whatever the other two flags say. All three are `false`  unless the installation switched plugins on in its configuration.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether an administrator may add a plugin of their own through  `POST api/2.0/settings/webplugins`. While it is `false` only the plugins that ship with the installation  are available.
	Upload *bool `json:"upload,omitempty"`
	// Whether an added plugin may be removed again through `DELETE api/2.0/settings/webplugins/{name}`. The  plugins that ship with the installation cannot be removed regardless of this flag.
	Delete *bool `json:"delete,omitempty"`
}

// NewPluginsDto instantiates a new PluginsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginsDto() *PluginsDto {
	this := PluginsDto{}
	return &this
}

// NewPluginsDtoWithDefaults instantiates a new PluginsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsDtoWithDefaults() *PluginsDto {
	this := PluginsDto{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PluginsDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginsDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PluginsDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PluginsDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUpload returns the Upload field value if set, zero value otherwise.
func (o *PluginsDto) GetUpload() bool {
	if o == nil || IsNil(o.Upload) {
		var ret bool
		return ret
	}
	return *o.Upload
}

// GetUploadOk returns a tuple with the Upload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginsDto) GetUploadOk() (*bool, bool) {
	if o == nil || IsNil(o.Upload) {
		return nil, false
	}
	return o.Upload, true
}

// HasUpload returns a boolean if a field has been set.
func (o *PluginsDto) IsUploadSet() bool {
	if o != nil && !IsNil(o.Upload) {
		return true
	}

	return false
}

// SetUpload gets a reference to the given bool and assigns it to the Upload field.
func (o *PluginsDto) SetUpload(v bool) {
	o.Upload = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *PluginsDto) GetDelete() bool {
	if o == nil || IsNil(o.Delete) {
		var ret bool
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginsDto) GetDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *PluginsDto) IsDeleteSet() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given bool and assigns it to the Delete field.
func (o *PluginsDto) SetDelete(v bool) {
	o.Delete = &v
}

func (o PluginsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Upload) {
		toSerialize["upload"] = o.Upload
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	return toSerialize, nil
}

type NullablePluginsDto struct {
	value *PluginsDto
	isSet bool
}

func (v NullablePluginsDto) Get() *PluginsDto {
	return v.value
}

func (v *NullablePluginsDto) Set(val *PluginsDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginsDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginsDto(val *PluginsDto) *NullablePluginsDto {
	return &NullablePluginsDto{value: val, isSet: true}
}

func (v NullablePluginsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

