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

// checks if the CookieSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CookieSettingsRequestsDto{}

// CookieSettingsRequestsDto The request parameters for managing cookie settings.
type CookieSettingsRequestsDto struct {
	// The cookie lifetime in minutes.
	LifeTime *int32 `json:"lifeTime,omitempty"`
	// Specifies whether the cookie settings are enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCookieSettingsRequestsDto instantiates a new CookieSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCookieSettingsRequestsDto() *CookieSettingsRequestsDto {
	this := CookieSettingsRequestsDto{}
	return &this
}

// NewCookieSettingsRequestsDtoWithDefaults instantiates a new CookieSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCookieSettingsRequestsDtoWithDefaults() *CookieSettingsRequestsDto {
	this := CookieSettingsRequestsDto{}
	return &this
}

// GetLifeTime returns the LifeTime field value if set, zero value otherwise.
func (o *CookieSettingsRequestsDto) GetLifeTime() int32 {
	if o == nil || IsNil(o.LifeTime) {
		var ret int32
		return ret
	}
	return *o.LifeTime
}

// GetLifeTimeOk returns a tuple with the LifeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSettingsRequestsDto) GetLifeTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.LifeTime) {
		return nil, false
	}
	return o.LifeTime, true
}

// HasLifeTime returns a boolean if a field has been set.
func (o *CookieSettingsRequestsDto) IsLifeTimeSet() bool {
	if o != nil && !IsNil(o.LifeTime) {
		return true
	}

	return false
}

// SetLifeTime gets a reference to the given int32 and assigns it to the LifeTime field.
func (o *CookieSettingsRequestsDto) SetLifeTime(v int32) {
	o.LifeTime = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CookieSettingsRequestsDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSettingsRequestsDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CookieSettingsRequestsDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CookieSettingsRequestsDto) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CookieSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CookieSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LifeTime) {
		toSerialize["lifeTime"] = o.LifeTime
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableCookieSettingsRequestsDto struct {
	value *CookieSettingsRequestsDto
	isSet bool
}

func (v NullableCookieSettingsRequestsDto) Get() *CookieSettingsRequestsDto {
	return v.value
}

func (v *NullableCookieSettingsRequestsDto) Set(val *CookieSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCookieSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCookieSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCookieSettingsRequestsDto(val *CookieSettingsRequestsDto) *NullableCookieSettingsRequestsDto {
	return &NullableCookieSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableCookieSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCookieSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

