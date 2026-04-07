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

// checks if the LogoCover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogoCover{}

// LogoCover The logo cover information.
type LogoCover struct {
	// The logo cover ID.
	Id NullableString `json:"id"`
	// The logo cover data.
	Data NullableString `json:"data"`
}

type _LogoCover LogoCover

// NewLogoCover instantiates a new LogoCover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogoCover(id NullableString, data NullableString) *LogoCover {
	this := LogoCover{}
	this.Id = id
	this.Data = data
	return &this
}

// NewLogoCoverWithDefaults instantiates a new LogoCover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoCoverWithDefaults() *LogoCover {
	this := LogoCover{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogoCover) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoCover) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *LogoCover) SetId(v string) {
	o.Id.Set(&v)
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LogoCover) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}

	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoCover) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value
func (o *LogoCover) SetData(v string) {
	o.Data.Set(&v)
}

func (o LogoCover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogoCover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["data"] = o.Data.Get()
	return toSerialize, nil
}

func (o *LogoCover) UnmarshalJSON(data []byte) (err error) {
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

	varLogoCover := _LogoCover{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varLogoCover)

	if err != nil {
		return err
	}

	*o = LogoCover(varLogoCover)

	return err
}

type NullableLogoCover struct {
	value *LogoCover
	isSet bool
}

func (v NullableLogoCover) Get() *LogoCover {
	return v.value
}

func (v *NullableLogoCover) Set(val *LogoCover) {
	v.value = val
	v.isSet = true
}

func (v NullableLogoCover) IsSet() bool {
	return v.isSet
}

func (v *NullableLogoCover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogoCover(val *LogoCover) *NullableLogoCover {
	return &NullableLogoCover{value: val, isSet: true}
}

func (v NullableLogoCover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogoCover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

