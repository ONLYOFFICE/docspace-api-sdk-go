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

// checks if the DeepLinkDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeepLinkDto{}

// DeepLinkDto The deep link parameters.
type DeepLinkDto struct {
	// The Android package name.
	AndroidPackageName NullableString `json:"androidPackageName"`
	// The deep link URL.
	Url NullableString `json:"url"`
	// The deep link IOS package ID.
	IosPackageId NullableString `json:"iosPackageId"`
}

type _DeepLinkDto DeepLinkDto

// NewDeepLinkDto instantiates a new DeepLinkDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeepLinkDto(androidPackageName NullableString, url NullableString, iosPackageId NullableString) *DeepLinkDto {
	this := DeepLinkDto{}
	this.AndroidPackageName = androidPackageName
	this.Url = url
	this.IosPackageId = iosPackageId
	return &this
}

// NewDeepLinkDtoWithDefaults instantiates a new DeepLinkDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeepLinkDtoWithDefaults() *DeepLinkDto {
	this := DeepLinkDto{}
	return &this
}

// GetAndroidPackageName returns the AndroidPackageName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeepLinkDto) GetAndroidPackageName() string {
	if o == nil || o.AndroidPackageName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AndroidPackageName.Get()
}

// GetAndroidPackageNameOk returns a tuple with the AndroidPackageName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeepLinkDto) GetAndroidPackageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AndroidPackageName.Get(), o.AndroidPackageName.IsSet()
}

// SetAndroidPackageName sets field value
func (o *DeepLinkDto) SetAndroidPackageName(v string) {
	o.AndroidPackageName.Set(&v)
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeepLinkDto) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeepLinkDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *DeepLinkDto) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetIosPackageId returns the IosPackageId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeepLinkDto) GetIosPackageId() string {
	if o == nil || o.IosPackageId.Get() == nil {
		var ret string
		return ret
	}

	return *o.IosPackageId.Get()
}

// GetIosPackageIdOk returns a tuple with the IosPackageId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeepLinkDto) GetIosPackageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IosPackageId.Get(), o.IosPackageId.IsSet()
}

// SetIosPackageId sets field value
func (o *DeepLinkDto) SetIosPackageId(v string) {
	o.IosPackageId.Set(&v)
}

func (o DeepLinkDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeepLinkDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["androidPackageName"] = o.AndroidPackageName.Get()
	toSerialize["url"] = o.Url.Get()
	toSerialize["iosPackageId"] = o.IosPackageId.Get()
	return toSerialize, nil
}

func (o *DeepLinkDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"androidPackageName",
		"url",
		"iosPackageId",
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

	varDeepLinkDto := _DeepLinkDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDeepLinkDto)

	if err != nil {
		return err
	}

	*o = DeepLinkDto(varDeepLinkDto)

	return err
}

type NullableDeepLinkDto struct {
	value *DeepLinkDto
	isSet bool
}

func (v NullableDeepLinkDto) Get() *DeepLinkDto {
	return v.value
}

func (v *NullableDeepLinkDto) Set(val *DeepLinkDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeepLinkDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeepLinkDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeepLinkDto(val *DeepLinkDto) *NullableDeepLinkDto {
	return &NullableDeepLinkDto{value: val, isSet: true}
}

func (v NullableDeepLinkDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeepLinkDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

