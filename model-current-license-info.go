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
	"bytes"
	"fmt"
)

// checks if the CurrentLicenseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentLicenseInfo{}

// CurrentLicenseInfo The two facts about the subscription in force that a payment page needs.
type CurrentLicenseInfo struct {
	// Whether the portal is on a trial rather than a paid subscription. A trial expires at `dueDate` and is not  extended by paying - a plan has to be bought instead.
	Trial bool `json:"trial"`
	// The day the subscription runs out, with the time of day cut off. The largest value a date can hold means  it never runs out, which is how a free or unlimited plan is expressed.
	DueDate time.Time `json:"dueDate"`
}

type _CurrentLicenseInfo CurrentLicenseInfo

// NewCurrentLicenseInfo instantiates a new CurrentLicenseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentLicenseInfo(trial bool, dueDate time.Time) *CurrentLicenseInfo {
	this := CurrentLicenseInfo{}
	this.Trial = trial
	this.DueDate = dueDate
	return &this
}

// NewCurrentLicenseInfoWithDefaults instantiates a new CurrentLicenseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentLicenseInfoWithDefaults() *CurrentLicenseInfo {
	this := CurrentLicenseInfo{}
	return &this
}

// GetTrial returns the Trial field value
func (o *CurrentLicenseInfo) GetTrial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trial
}

// GetTrialOk returns a tuple with the Trial field value
// and a boolean to check if the value has been set.
func (o *CurrentLicenseInfo) GetTrialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trial, true
}

// SetTrial sets field value
func (o *CurrentLicenseInfo) SetTrial(v bool) {
	o.Trial = v
}

// GetDueDate returns the DueDate field value
func (o *CurrentLicenseInfo) GetDueDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value
// and a boolean to check if the value has been set.
func (o *CurrentLicenseInfo) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDate, true
}

// SetDueDate sets field value
func (o *CurrentLicenseInfo) SetDueDate(v time.Time) {
	o.DueDate = v
}

func (o CurrentLicenseInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentLicenseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trial"] = o.Trial
	toSerialize["dueDate"] = o.DueDate
	return toSerialize, nil
}

func (o *CurrentLicenseInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"trial",
		"dueDate",
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

	varCurrentLicenseInfo := _CurrentLicenseInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCurrentLicenseInfo)

	if err != nil {
		return err
	}

	*o = CurrentLicenseInfo(varCurrentLicenseInfo)

	return err
}

type NullableCurrentLicenseInfo struct {
	value *CurrentLicenseInfo
	isSet bool
}

func (v NullableCurrentLicenseInfo) Get() *CurrentLicenseInfo {
	return v.value
}

func (v *NullableCurrentLicenseInfo) Set(val *CurrentLicenseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentLicenseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentLicenseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentLicenseInfo(val *CurrentLicenseInfo) *NullableCurrentLicenseInfo {
	return &NullableCurrentLicenseInfo{value: val, isSet: true}
}

func (v NullableCurrentLicenseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentLicenseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

