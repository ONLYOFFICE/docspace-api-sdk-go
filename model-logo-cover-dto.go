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

// checks if the LogoCoverDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogoCoverDto{}

// LogoCoverDto The logo cover information.
type LogoCoverDto struct {
	// The logo cover ID.
	Id NullableString `json:"id"`
	// The logo cover data.
	Data NullableString `json:"data"`
}

type _LogoCoverDto LogoCoverDto

// NewLogoCoverDto instantiates a new LogoCoverDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogoCoverDto(id NullableString, data NullableString) *LogoCoverDto {
	this := LogoCoverDto{}
	this.Id = id
	this.Data = data
	return &this
}

// NewLogoCoverDtoWithDefaults instantiates a new LogoCoverDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoCoverDtoWithDefaults() *LogoCoverDto {
	this := LogoCoverDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogoCoverDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoCoverDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *LogoCoverDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogoCoverDto) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}

	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoCoverDto) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value
func (o *LogoCoverDto) SetData(v string) {
	o.Data.Set(&v)
}

func (o LogoCoverDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogoCoverDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["data"] = o.Data.Get()
	return toSerialize, nil
}

func (o *LogoCoverDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"data",
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

	varLogoCoverDto := _LogoCoverDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varLogoCoverDto)

	if err != nil {
		return err
	}

	*o = LogoCoverDto(varLogoCoverDto)

	return err
}

type NullableLogoCoverDto struct {
	value *LogoCoverDto
	isSet bool
}

func (v NullableLogoCoverDto) Get() *LogoCoverDto {
	return v.value
}

func (v *NullableLogoCoverDto) Set(val *LogoCoverDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLogoCoverDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLogoCoverDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogoCoverDto(val *LogoCoverDto) *NullableLogoCoverDto {
	return &NullableLogoCoverDto{value: val, isSet: true}
}

func (v NullableLogoCoverDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogoCoverDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

