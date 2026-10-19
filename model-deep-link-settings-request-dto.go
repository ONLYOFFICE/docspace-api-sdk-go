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
	"time"
)

// checks if the DeepLinkSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeepLinkSettingsRequestDto{}

// DeepLinkSettingsRequestDto The deep link handling the portal applies on mobile devices.
type DeepLinkSettingsRequestDto struct {
	// Whether a link always opens in the browser, always in the native application, or asks the user each time.
	HandlingMode *DeepLinkHandlingMode `json:"handlingMode,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewDeepLinkSettingsRequestDto instantiates a new DeepLinkSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeepLinkSettingsRequestDto() *DeepLinkSettingsRequestDto {
	this := DeepLinkSettingsRequestDto{}
	return &this
}

// NewDeepLinkSettingsRequestDtoWithDefaults instantiates a new DeepLinkSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeepLinkSettingsRequestDtoWithDefaults() *DeepLinkSettingsRequestDto {
	this := DeepLinkSettingsRequestDto{}
	return &this
}

// GetHandlingMode returns the HandlingMode field value if set, zero value otherwise.
func (o *DeepLinkSettingsRequestDto) GetHandlingMode() DeepLinkHandlingMode {
	if o == nil || IsNil(o.HandlingMode) {
		var ret DeepLinkHandlingMode
		return ret
	}
	return *o.HandlingMode
}

// GetHandlingModeOk returns a tuple with the HandlingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeepLinkSettingsRequestDto) GetHandlingModeOk() (*DeepLinkHandlingMode, bool) {
	if o == nil || IsNil(o.HandlingMode) {
		return nil, false
	}
	return o.HandlingMode, true
}

// HasHandlingMode returns a boolean if a field has been set.
func (o *DeepLinkSettingsRequestDto) IsHandlingModeSet() bool {
	if o != nil && !IsNil(o.HandlingMode) {
		return true
	}

	return false
}

// SetHandlingMode gets a reference to the given DeepLinkHandlingMode and assigns it to the HandlingMode field.
func (o *DeepLinkSettingsRequestDto) SetHandlingMode(v DeepLinkHandlingMode) {
	o.HandlingMode = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *DeepLinkSettingsRequestDto) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeepLinkSettingsRequestDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *DeepLinkSettingsRequestDto) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *DeepLinkSettingsRequestDto) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o DeepLinkSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeepLinkSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HandlingMode) {
		toSerialize["handlingMode"] = o.HandlingMode
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableDeepLinkSettingsRequestDto struct {
	value *DeepLinkSettingsRequestDto
	isSet bool
}

func (v NullableDeepLinkSettingsRequestDto) Get() *DeepLinkSettingsRequestDto {
	return v.value
}

func (v *NullableDeepLinkSettingsRequestDto) Set(val *DeepLinkSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeepLinkSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeepLinkSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeepLinkSettingsRequestDto(val *DeepLinkSettingsRequestDto) *NullableDeepLinkSettingsRequestDto {
	return &NullableDeepLinkSettingsRequestDto{value: val, isSet: true}
}

func (v NullableDeepLinkSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeepLinkSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

