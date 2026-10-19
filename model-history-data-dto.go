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

// checks if the HistoryDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryDataDto{}

// HistoryDataDto The history data.
type HistoryDataDto struct {
	// The name of the action initiator.
	InitiatorName NullableString `json:"initiatorName,omitempty"`
}

// NewHistoryDataDto instantiates a new HistoryDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryDataDto() *HistoryDataDto {
	this := HistoryDataDto{}
	return &this
}

// NewHistoryDataDtoWithDefaults instantiates a new HistoryDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryDataDtoWithDefaults() *HistoryDataDto {
	this := HistoryDataDto{}
	return &this
}

// GetInitiatorName returns the InitiatorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryDataDto) GetInitiatorName() string {
	if o == nil || IsNil(o.InitiatorName.Get()) {
		var ret string
		return ret
	}
	return *o.InitiatorName.Get()
}

// GetInitiatorNameOk returns a tuple with the InitiatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryDataDto) GetInitiatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitiatorName.Get(), o.InitiatorName.IsSet()
}

// HasInitiatorName returns a boolean if a field has been set.
func (o *HistoryDataDto) IsInitiatorNameSet() bool {
	if o != nil && o.InitiatorName.IsSet() {
		return true
	}

	return false
}

// SetInitiatorName gets a reference to the given NullableString and assigns it to the InitiatorName field.
func (o *HistoryDataDto) SetInitiatorName(v string) {
	o.InitiatorName.Set(&v)
}
// SetInitiatorNameNil sets the value for InitiatorName to be an explicit nil
func (o *HistoryDataDto) SetInitiatorNameNil() {
	o.InitiatorName.Set(nil)
}

// UnsetInitiatorName ensures that no value is present for InitiatorName, not even an explicit nil
func (o *HistoryDataDto) UnsetInitiatorName() {
	o.InitiatorName.Unset()
}

func (o HistoryDataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InitiatorName.IsSet() {
		toSerialize["initiatorName"] = o.InitiatorName.Get()
	}
	return toSerialize, nil
}

type NullableHistoryDataDto struct {
	value *HistoryDataDto
	isSet bool
}

func (v NullableHistoryDataDto) Get() *HistoryDataDto {
	return v.value
}

func (v *NullableHistoryDataDto) Set(val *HistoryDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryDataDto(val *HistoryDataDto) *NullableHistoryDataDto {
	return &NullableHistoryDataDto{value: val, isSet: true}
}

func (v NullableHistoryDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

