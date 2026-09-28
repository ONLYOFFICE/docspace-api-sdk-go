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

// checks if the DocsCloudLicenseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudLicenseInfo{}

// DocsCloudLicenseInfo Represents the license information of a Docs Connect tenant.
type DocsCloudLicenseInfo struct {
	// The date and time until which the license is valid.
	Valid *time.Time `json:"valid,omitempty"`
	// Whether the license is a trial.
	Trial *bool `json:"trial,omitempty"`
	// The license build date.
	BuildDate *time.Time `json:"buildDate,omitempty"`
}

// NewDocsCloudLicenseInfo instantiates a new DocsCloudLicenseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudLicenseInfo() *DocsCloudLicenseInfo {
	this := DocsCloudLicenseInfo{}
	return &this
}

// NewDocsCloudLicenseInfoWithDefaults instantiates a new DocsCloudLicenseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudLicenseInfoWithDefaults() *DocsCloudLicenseInfo {
	this := DocsCloudLicenseInfo{}
	return &this
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *DocsCloudLicenseInfo) GetValid() time.Time {
	if o == nil || IsNil(o.Valid) {
		var ret time.Time
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudLicenseInfo) GetValidOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *DocsCloudLicenseInfo) IsValidSet() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given time.Time and assigns it to the Valid field.
func (o *DocsCloudLicenseInfo) SetValid(v time.Time) {
	o.Valid = &v
}

// GetTrial returns the Trial field value if set, zero value otherwise.
func (o *DocsCloudLicenseInfo) GetTrial() bool {
	if o == nil || IsNil(o.Trial) {
		var ret bool
		return ret
	}
	return *o.Trial
}

// GetTrialOk returns a tuple with the Trial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudLicenseInfo) GetTrialOk() (*bool, bool) {
	if o == nil || IsNil(o.Trial) {
		return nil, false
	}
	return o.Trial, true
}

// HasTrial returns a boolean if a field has been set.
func (o *DocsCloudLicenseInfo) IsTrialSet() bool {
	if o != nil && !IsNil(o.Trial) {
		return true
	}

	return false
}

// SetTrial gets a reference to the given bool and assigns it to the Trial field.
func (o *DocsCloudLicenseInfo) SetTrial(v bool) {
	o.Trial = &v
}

// GetBuildDate returns the BuildDate field value if set, zero value otherwise.
func (o *DocsCloudLicenseInfo) GetBuildDate() time.Time {
	if o == nil || IsNil(o.BuildDate) {
		var ret time.Time
		return ret
	}
	return *o.BuildDate
}

// GetBuildDateOk returns a tuple with the BuildDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudLicenseInfo) GetBuildDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BuildDate) {
		return nil, false
	}
	return o.BuildDate, true
}

// HasBuildDate returns a boolean if a field has been set.
func (o *DocsCloudLicenseInfo) IsBuildDateSet() bool {
	if o != nil && !IsNil(o.BuildDate) {
		return true
	}

	return false
}

// SetBuildDate gets a reference to the given time.Time and assigns it to the BuildDate field.
func (o *DocsCloudLicenseInfo) SetBuildDate(v time.Time) {
	o.BuildDate = &v
}

func (o DocsCloudLicenseInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudLicenseInfo) ToMap() (map[string]interface{}, error) {
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

type NullableDocsCloudLicenseInfo struct {
	value *DocsCloudLicenseInfo
	isSet bool
}

func (v NullableDocsCloudLicenseInfo) Get() *DocsCloudLicenseInfo {
	return v.value
}

func (v *NullableDocsCloudLicenseInfo) Set(val *DocsCloudLicenseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudLicenseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudLicenseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudLicenseInfo(val *DocsCloudLicenseInfo) *NullableDocsCloudLicenseInfo {
	return &NullableDocsCloudLicenseInfo{value: val, isSet: true}
}

func (v NullableDocsCloudLicenseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudLicenseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

