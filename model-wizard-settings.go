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

// checks if the WizardSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WizardSettings{}

// WizardSettings The Wizard settings.
type WizardSettings struct {
	// Specifies if the Wizard settings are completed or not
	Completed *bool `json:"completed,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewWizardSettings instantiates a new WizardSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWizardSettings() *WizardSettings {
	this := WizardSettings{}
	return &this
}

// NewWizardSettingsWithDefaults instantiates a new WizardSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWizardSettingsWithDefaults() *WizardSettings {
	this := WizardSettings{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *WizardSettings) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSettings) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *WizardSettings) IsCompletedSet() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *WizardSettings) SetCompleted(v bool) {
	o.Completed = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *WizardSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *WizardSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *WizardSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o WizardSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WizardSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableWizardSettings struct {
	value *WizardSettings
	isSet bool
}

func (v NullableWizardSettings) Get() *WizardSettings {
	return v.value
}

func (v *NullableWizardSettings) Set(val *WizardSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWizardSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWizardSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWizardSettings(val *WizardSettings) *NullableWizardSettings {
	return &NullableWizardSettings{value: val, isSet: true}
}

func (v NullableWizardSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWizardSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

