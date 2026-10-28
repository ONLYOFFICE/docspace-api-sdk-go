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

// checks if the DocsCloudLicenseInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudLicenseInfoDto{}

// DocsCloudLicenseInfoDto Represents the license information of a Docs Connect tenant.
type DocsCloudLicenseInfoDto struct {
	// The date and time until which the license is valid.
	Valid *time.Time `json:"valid,omitempty"`
	// Whether the license is a trial.
	Trial *bool `json:"trial,omitempty"`
	// The license build date.
	BuildDate *time.Time `json:"buildDate,omitempty"`
}

// NewDocsCloudLicenseInfoDto instantiates a new DocsCloudLicenseInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudLicenseInfoDto() *DocsCloudLicenseInfoDto {
	this := DocsCloudLicenseInfoDto{}
	return &this
}

// NewDocsCloudLicenseInfoDtoWithDefaults instantiates a new DocsCloudLicenseInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudLicenseInfoDtoWithDefaults() *DocsCloudLicenseInfoDto {
	this := DocsCloudLicenseInfoDto{}
	return &this
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *DocsCloudLicenseInfoDto) GetValid() time.Time {
	if o == nil || IsNil(o.Valid) {
		var ret time.Time
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudLicenseInfoDto) GetValidOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *DocsCloudLicenseInfoDto) IsValidSet() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given time.Time and assigns it to the Valid field.
func (o *DocsCloudLicenseInfoDto) SetValid(v time.Time) {
	o.Valid = &v
}

// GetTrial returns the Trial field value if set, zero value otherwise.
func (o *DocsCloudLicenseInfoDto) GetTrial() bool {
	if o == nil || IsNil(o.Trial) {
		var ret bool
		return ret
	}
	return *o.Trial
}

// GetTrialOk returns a tuple with the Trial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudLicenseInfoDto) GetTrialOk() (*bool, bool) {
	if o == nil || IsNil(o.Trial) {
		return nil, false
	}
	return o.Trial, true
}

// HasTrial returns a boolean if a field has been set.
func (o *DocsCloudLicenseInfoDto) IsTrialSet() bool {
	if o != nil && !IsNil(o.Trial) {
		return true
	}

	return false
}

// SetTrial gets a reference to the given bool and assigns it to the Trial field.
func (o *DocsCloudLicenseInfoDto) SetTrial(v bool) {
	o.Trial = &v
}

// GetBuildDate returns the BuildDate field value if set, zero value otherwise.
func (o *DocsCloudLicenseInfoDto) GetBuildDate() time.Time {
	if o == nil || IsNil(o.BuildDate) {
		var ret time.Time
		return ret
	}
	return *o.BuildDate
}

// GetBuildDateOk returns a tuple with the BuildDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudLicenseInfoDto) GetBuildDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BuildDate) {
		return nil, false
	}
	return o.BuildDate, true
}

// HasBuildDate returns a boolean if a field has been set.
func (o *DocsCloudLicenseInfoDto) IsBuildDateSet() bool {
	if o != nil && !IsNil(o.BuildDate) {
		return true
	}

	return false
}

// SetBuildDate gets a reference to the given time.Time and assigns it to the BuildDate field.
func (o *DocsCloudLicenseInfoDto) SetBuildDate(v time.Time) {
	o.BuildDate = &v
}

func (o DocsCloudLicenseInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudLicenseInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	if !IsNil(o.Trial) {
		toSerialize["trial"] = o.Trial
	}
	if !IsNil(o.BuildDate) {
		toSerialize["buildDate"] = o.BuildDate
	}
	return toSerialize, nil
}

type NullableDocsCloudLicenseInfoDto struct {
	value *DocsCloudLicenseInfoDto
	isSet bool
}

func (v NullableDocsCloudLicenseInfoDto) Get() *DocsCloudLicenseInfoDto {
	return v.value
}

func (v *NullableDocsCloudLicenseInfoDto) Set(val *DocsCloudLicenseInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudLicenseInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudLicenseInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudLicenseInfoDto(val *DocsCloudLicenseInfoDto) *NullableDocsCloudLicenseInfoDto {
	return &NullableDocsCloudLicenseInfoDto{value: val, isSet: true}
}

func (v NullableDocsCloudLicenseInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudLicenseInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

