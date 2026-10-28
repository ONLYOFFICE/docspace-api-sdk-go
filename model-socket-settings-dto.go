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

// checks if the SocketSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SocketSettingsDto{}

// SocketSettingsDto Where a client connects for the portal's live updates.
type SocketSettingsDto struct {
	// The base address of the Socket.IO hub that pushes file changes, presence and quota alerts, always with a  trailing slash. It is empty when the installation runs no hub, and a client must then fall back to  polling rather than guessing an address. The value comes from the installation's configuration and cannot  be changed through this API.
	Url NullableString `json:"url,omitempty"`
}

// NewSocketSettingsDto instantiates a new SocketSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocketSettingsDto() *SocketSettingsDto {
	this := SocketSettingsDto{}
	return &this
}

// NewSocketSettingsDtoWithDefaults instantiates a new SocketSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocketSettingsDtoWithDefaults() *SocketSettingsDto {
	this := SocketSettingsDto{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SocketSettingsDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SocketSettingsDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *SocketSettingsDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *SocketSettingsDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *SocketSettingsDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *SocketSettingsDto) UnsetUrl() {
	o.Url.Unset()
}

func (o SocketSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SocketSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableSocketSettingsDto struct {
	value *SocketSettingsDto
	isSet bool
}

func (v NullableSocketSettingsDto) Get() *SocketSettingsDto {
	return v.value
}

func (v *NullableSocketSettingsDto) Set(val *SocketSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSocketSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSocketSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocketSettingsDto(val *SocketSettingsDto) *NullableSocketSettingsDto {
	return &NullableSocketSettingsDto{value: val, isSet: true}
}

func (v NullableSocketSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocketSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

