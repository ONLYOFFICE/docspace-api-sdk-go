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

// checks if the FormResultsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormResultsDto{}

// FormResultsDto struct for FormResultsDto
type FormResultsDto struct {
	// The date and time when the form was created.
	CreateOn *time.Time `json:"createOn,omitempty"`
	// The list of forms data.
	FormsData []FormsItemData `json:"formsData,omitempty"`
}

// NewFormResultsDto instantiates a new FormResultsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormResultsDto() *FormResultsDto {
	this := FormResultsDto{}
	return &this
}

// NewFormResultsDtoWithDefaults instantiates a new FormResultsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormResultsDtoWithDefaults() *FormResultsDto {
	this := FormResultsDto{}
	return &this
}

// GetCreateOn returns the CreateOn field value if set, zero value otherwise.
func (o *FormResultsDto) GetCreateOn() time.Time {
	if o == nil || IsNil(o.CreateOn) {
		var ret time.Time
		return ret
	}
	return *o.CreateOn
}

// GetCreateOnOk returns a tuple with the CreateOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormResultsDto) GetCreateOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateOn) {
		return nil, false
	}
	return o.CreateOn, true
}

// HasCreateOn returns a boolean if a field has been set.
func (o *FormResultsDto) IsCreateOnSet() bool {
	if o != nil && !IsNil(o.CreateOn) {
		return true
	}

	return false
}

// SetCreateOn gets a reference to the given time.Time and assigns it to the CreateOn field.
func (o *FormResultsDto) SetCreateOn(v time.Time) {
	o.CreateOn = &v
}

// GetFormsData returns the FormsData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormResultsDto) GetFormsData() []FormsItemData {
	if o == nil {
		var ret []FormsItemData
		return ret
	}
	return o.FormsData
}

// GetFormsDataOk returns a tuple with the FormsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormResultsDto) GetFormsDataOk() ([]FormsItemData, bool) {
	if o == nil || IsNil(o.FormsData) {
		return nil, false
	}
	return o.FormsData, true
}

// HasFormsData returns a boolean if a field has been set.
func (o *FormResultsDto) IsFormsDataSet() bool {
	if o != nil && !IsNil(o.FormsData) {
		return true
	}

	return false
}

// SetFormsData gets a reference to the given []FormsItemData and assigns it to the FormsData field.
func (o *FormResultsDto) SetFormsData(v []FormsItemData) {
	o.FormsData = v
}

func (o FormResultsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormResultsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateOn) {
		toSerialize["createOn"] = o.CreateOn
	}
	if o.FormsData != nil {
		toSerialize["formsData"] = o.FormsData
	}
	return toSerialize, nil
}

type NullableFormResultsDto struct {
	value *FormResultsDto
	isSet bool
}

func (v NullableFormResultsDto) Get() *FormResultsDto {
	return v.value
}

func (v *NullableFormResultsDto) Set(val *FormResultsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFormResultsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFormResultsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormResultsDto(val *FormResultsDto) *NullableFormResultsDto {
	return &NullableFormResultsDto{value: val, isSet: true}
}

func (v NullableFormResultsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormResultsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

