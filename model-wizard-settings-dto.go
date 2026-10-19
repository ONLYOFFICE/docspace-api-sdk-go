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

// checks if the WizardSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WizardSettingsDto{}

// WizardSettingsDto The state of the portal's first-run setup wizard.
type WizardSettingsDto struct {
	// Specifies if the Wizard settings are completed or not
	Completed *bool `json:"completed,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewWizardSettingsDto instantiates a new WizardSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWizardSettingsDto() *WizardSettingsDto {
	this := WizardSettingsDto{}
	return &this
}

// NewWizardSettingsDtoWithDefaults instantiates a new WizardSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWizardSettingsDtoWithDefaults() *WizardSettingsDto {
	this := WizardSettingsDto{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *WizardSettingsDto) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSettingsDto) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *WizardSettingsDto) IsCompletedSet() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *WizardSettingsDto) SetCompleted(v bool) {
	o.Completed = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *WizardSettingsDto) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSettingsDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *WizardSettingsDto) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *WizardSettingsDto) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o WizardSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WizardSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableWizardSettingsDto struct {
	value *WizardSettingsDto
	isSet bool
}

func (v NullableWizardSettingsDto) Get() *WizardSettingsDto {
	return v.value
}

func (v *NullableWizardSettingsDto) Set(val *WizardSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWizardSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWizardSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWizardSettingsDto(val *WizardSettingsDto) *NullableWizardSettingsDto {
	return &NullableWizardSettingsDto{value: val, isSet: true}
}

func (v NullableWizardSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWizardSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

