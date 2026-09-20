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

// checks if the ChangeHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeHistory{}

// ChangeHistory The change to make to a revision group of a file.
type ChangeHistory struct {
	// The version the change applies to; 0 means the current version of the file.
	Version int32 `json:"version"`
	// What to do with the revision group: `false` completes the named version, storing its content again as a fresh  version that opens a new group, while `true` folds the last group back into the group before it, so the next  save continues that revision.
	ContinueVersion *bool `json:"continueVersion,omitempty"`
}

type _ChangeHistory ChangeHistory

// NewChangeHistory instantiates a new ChangeHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeHistory(version int32) *ChangeHistory {
	this := ChangeHistory{}
	this.Version = version
	return &this
}

// NewChangeHistoryWithDefaults instantiates a new ChangeHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeHistoryWithDefaults() *ChangeHistory {
	this := ChangeHistory{}
	return &this
}

// GetVersion returns the Version field value
func (o *ChangeHistory) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ChangeHistory) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ChangeHistory) SetVersion(v int32) {
	o.Version = v
}

// GetContinueVersion returns the ContinueVersion field value if set, zero value otherwise.
func (o *ChangeHistory) GetContinueVersion() bool {
	if o == nil || IsNil(o.ContinueVersion) {
		var ret bool
		return ret
	}
	return *o.ContinueVersion
}

// GetContinueVersionOk returns a tuple with the ContinueVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeHistory) GetContinueVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.ContinueVersion) {
		return nil, false
	}
	return o.ContinueVersion, true
}

// HasContinueVersion returns a boolean if a field has been set.
func (o *ChangeHistory) IsContinueVersionSet() bool {
	if o != nil && !IsNil(o.ContinueVersion) {
		return true
	}

	return false
}

// SetContinueVersion gets a reference to the given bool and assigns it to the ContinueVersion field.
func (o *ChangeHistory) SetContinueVersion(v bool) {
	o.ContinueVersion = &v
}

func (o ChangeHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	if !IsNil(o.ContinueVersion) {
		toSerialize["continueVersion"] = o.ContinueVersion
	}
	return toSerialize, nil
}

func (o *ChangeHistory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
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

	varChangeHistory := _ChangeHistory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varChangeHistory)

	if err != nil {
		return err
	}

	*o = ChangeHistory(varChangeHistory)

	return err
}

type NullableChangeHistory struct {
	value *ChangeHistory
	isSet bool
}

func (v NullableChangeHistory) Get() *ChangeHistory {
	return v.value
}

func (v *NullableChangeHistory) Set(val *ChangeHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeHistory(val *ChangeHistory) *NullableChangeHistory {
	return &NullableChangeHistory{value: val, isSet: true}
}

func (v NullableChangeHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

