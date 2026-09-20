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
	"bytes"
	"fmt"
)

// checks if the CookieSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CookieSettingsDto{}

// CookieSettingsDto How long an authentication session of the portal stays valid, and whether that limit is applied.
type CookieSettingsDto struct {
	// How long, in minutes, a session issued from now on remains valid. It is `1440` on a portal that has never  stored a limit, and that stored number is reported whether or not `enabled` puts it to use.
	LifeTime int32 `json:"lifeTime"`
	// Whether the stored lifetime is applied at all. While it is `false` the number above is ignored and an  issued session is honoured for a year.
	Enabled bool `json:"enabled"`
}

type _CookieSettingsDto CookieSettingsDto

// NewCookieSettingsDto instantiates a new CookieSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCookieSettingsDto(lifeTime int32, enabled bool) *CookieSettingsDto {
	this := CookieSettingsDto{}
	this.LifeTime = lifeTime
	this.Enabled = enabled
	return &this
}

// NewCookieSettingsDtoWithDefaults instantiates a new CookieSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCookieSettingsDtoWithDefaults() *CookieSettingsDto {
	this := CookieSettingsDto{}
	return &this
}

// GetLifeTime returns the LifeTime field value
func (o *CookieSettingsDto) GetLifeTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LifeTime
}

// GetLifeTimeOk returns a tuple with the LifeTime field value
// and a boolean to check if the value has been set.
func (o *CookieSettingsDto) GetLifeTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LifeTime, true
}

// SetLifeTime sets field value
func (o *CookieSettingsDto) SetLifeTime(v int32) {
	o.LifeTime = v
}

// GetEnabled returns the Enabled field value
func (o *CookieSettingsDto) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CookieSettingsDto) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CookieSettingsDto) SetEnabled(v bool) {
	o.Enabled = v
}

func (o CookieSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CookieSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lifeTime"] = o.LifeTime
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *CookieSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lifeTime",
		"enabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCookieSettingsDto := _CookieSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCookieSettingsDto)

	if err != nil {
		return err
	}

	*o = CookieSettingsDto(varCookieSettingsDto)

	return err
}

type NullableCookieSettingsDto struct {
	value *CookieSettingsDto
	isSet bool
}

func (v NullableCookieSettingsDto) Get() *CookieSettingsDto {
	return v.value
}

func (v *NullableCookieSettingsDto) Set(val *CookieSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCookieSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCookieSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCookieSettingsDto(val *CookieSettingsDto) *NullableCookieSettingsDto {
	return &NullableCookieSettingsDto{value: val, isSet: true}
}

func (v NullableCookieSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCookieSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

