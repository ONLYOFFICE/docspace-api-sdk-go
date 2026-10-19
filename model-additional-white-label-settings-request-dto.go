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

// checks if the AdditionalWhiteLabelSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalWhiteLabelSettingsRequestDto{}

// AdditionalWhiteLabelSettingsRequestDto The body of a change to the help and community resources the interface offers.
type AdditionalWhiteLabelSettingsRequestDto struct {
	// The resource flags to store.
	Settings *SaveAdditionalResourcesRequest `json:"settings,omitempty"`
}

// NewAdditionalWhiteLabelSettingsRequestDto instantiates a new AdditionalWhiteLabelSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalWhiteLabelSettingsRequestDto() *AdditionalWhiteLabelSettingsRequestDto {
	this := AdditionalWhiteLabelSettingsRequestDto{}
	return &this
}

// NewAdditionalWhiteLabelSettingsRequestDtoWithDefaults instantiates a new AdditionalWhiteLabelSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalWhiteLabelSettingsRequestDtoWithDefaults() *AdditionalWhiteLabelSettingsRequestDto {
	this := AdditionalWhiteLabelSettingsRequestDto{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AdditionalWhiteLabelSettingsRequestDto) GetSettings() SaveAdditionalResourcesRequest {
	if o == nil || IsNil(o.Settings) {
		var ret SaveAdditionalResourcesRequest
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettingsRequestDto) GetSettingsOk() (*SaveAdditionalResourcesRequest, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AdditionalWhiteLabelSettingsRequestDto) IsSettingsSet() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SaveAdditionalResourcesRequest and assigns it to the Settings field.
func (o *AdditionalWhiteLabelSettingsRequestDto) SetSettings(v SaveAdditionalResourcesRequest) {
	o.Settings = &v
}

func (o AdditionalWhiteLabelSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalWhiteLabelSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableAdditionalWhiteLabelSettingsRequestDto struct {
	value *AdditionalWhiteLabelSettingsRequestDto
	isSet bool
}

func (v NullableAdditionalWhiteLabelSettingsRequestDto) Get() *AdditionalWhiteLabelSettingsRequestDto {
	return v.value
}

func (v *NullableAdditionalWhiteLabelSettingsRequestDto) Set(val *AdditionalWhiteLabelSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalWhiteLabelSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalWhiteLabelSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalWhiteLabelSettingsRequestDto(val *AdditionalWhiteLabelSettingsRequestDto) *NullableAdditionalWhiteLabelSettingsRequestDto {
	return &NullableAdditionalWhiteLabelSettingsRequestDto{value: val, isSet: true}
}

func (v NullableAdditionalWhiteLabelSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalWhiteLabelSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

