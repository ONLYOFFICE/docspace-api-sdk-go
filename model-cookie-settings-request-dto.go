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

// checks if the CookieSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CookieSettingsRequestDto{}

// CookieSettingsRequestDto How long an authentication session of the portal stays valid, and whether that limit is applied.
type CookieSettingsRequestDto struct {
	// How long, in minutes, a session issued from now on remains valid. A value above 9999 is clamped to 9999  rather than refused, and 0 or less clears the number, which together with `enabled` leaves sessions that  never expire on their own. Any positive value invalidates every session issued before this call, the  caller's included, so the client has to keep the fresh cookie the response carries.
	LifeTime *int32 `json:"lifeTime,omitempty"`
	// Whether the stored lifetime is applied at all. While it is false the number is ignored and an issued session  is honoured for a year; while it is true the connections behind expired sessions are dropped as well.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCookieSettingsRequestDto instantiates a new CookieSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCookieSettingsRequestDto() *CookieSettingsRequestDto {
	this := CookieSettingsRequestDto{}
	return &this
}

// NewCookieSettingsRequestDtoWithDefaults instantiates a new CookieSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCookieSettingsRequestDtoWithDefaults() *CookieSettingsRequestDto {
	this := CookieSettingsRequestDto{}
	return &this
}

// GetLifeTime returns the LifeTime field value if set, zero value otherwise.
func (o *CookieSettingsRequestDto) GetLifeTime() int32 {
	if o == nil || IsNil(o.LifeTime) {
		var ret int32
		return ret
	}
	return *o.LifeTime
}

// GetLifeTimeOk returns a tuple with the LifeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSettingsRequestDto) GetLifeTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.LifeTime) {
		return nil, false
	}
	return o.LifeTime, true
}

// HasLifeTime returns a boolean if a field has been set.
func (o *CookieSettingsRequestDto) IsLifeTimeSet() bool {
	if o != nil && !IsNil(o.LifeTime) {
		return true
	}

	return false
}

// SetLifeTime gets a reference to the given int32 and assigns it to the LifeTime field.
func (o *CookieSettingsRequestDto) SetLifeTime(v int32) {
	o.LifeTime = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CookieSettingsRequestDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSettingsRequestDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CookieSettingsRequestDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CookieSettingsRequestDto) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CookieSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CookieSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LifeTime) {
		toSerialize["lifeTime"] = o.LifeTime
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableCookieSettingsRequestDto struct {
	value *CookieSettingsRequestDto
	isSet bool
}

func (v NullableCookieSettingsRequestDto) Get() *CookieSettingsRequestDto {
	return v.value
}

func (v *NullableCookieSettingsRequestDto) Set(val *CookieSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCookieSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCookieSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCookieSettingsRequestDto(val *CookieSettingsRequestDto) *NullableCookieSettingsRequestDto {
	return &NullableCookieSettingsRequestDto{value: val, isSet: true}
}

func (v NullableCookieSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCookieSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

