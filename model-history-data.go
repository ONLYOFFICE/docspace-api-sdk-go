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

// checks if the HistoryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryData{}

// HistoryData The history data.
type HistoryData struct {
	// The name of the action initiator.
	InitiatorName NullableString `json:"initiatorName,omitempty"`
}

// NewHistoryData instantiates a new HistoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryData() *HistoryData {
	this := HistoryData{}
	return &this
}

// NewHistoryDataWithDefaults instantiates a new HistoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryDataWithDefaults() *HistoryData {
	this := HistoryData{}
	return &this
}

// GetInitiatorName returns the InitiatorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryData) GetInitiatorName() string {
	if o == nil || IsNil(o.InitiatorName.Get()) {
		var ret string
		return ret
	}
	return *o.InitiatorName.Get()
}

// GetInitiatorNameOk returns a tuple with the InitiatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryData) GetInitiatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitiatorName.Get(), o.InitiatorName.IsSet()
}

// HasInitiatorName returns a boolean if a field has been set.
func (o *HistoryData) IsInitiatorNameSet() bool {
	if o != nil && o.InitiatorName.IsSet() {
		return true
	}

	return false
}

// SetInitiatorName gets a reference to the given NullableString and assigns it to the InitiatorName field.
func (o *HistoryData) SetInitiatorName(v string) {
	o.InitiatorName.Set(&v)
}
// SetInitiatorNameNil sets the value for InitiatorName to be an explicit nil
func (o *HistoryData) SetInitiatorNameNil() {
	o.InitiatorName.Set(nil)
}

// UnsetInitiatorName ensures that no value is present for InitiatorName, not even an explicit nil
func (o *HistoryData) UnsetInitiatorName() {
	o.InitiatorName.Unset()
}

func (o HistoryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InitiatorName.IsSet() {
		toSerialize["initiatorName"] = o.InitiatorName.Get()
	}
	return toSerialize, nil
}

type NullableHistoryData struct {
	value *HistoryData
	isSet bool
}

func (v NullableHistoryData) Get() *HistoryData {
	return v.value
}

func (v *NullableHistoryData) Set(val *HistoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryData(val *HistoryData) *NullableHistoryData {
	return &NullableHistoryData{value: val, isSet: true}
}

func (v NullableHistoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

