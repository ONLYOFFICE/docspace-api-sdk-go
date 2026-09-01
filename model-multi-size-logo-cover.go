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

// checks if the MultiSizeLogoCover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiSizeLogoCover{}

// MultiSizeLogoCover The logo cover information, with the cover data in every available size.
type MultiSizeLogoCover struct {
	// The logo cover ID.
	Id NullableString `json:"id"`
	// The logo cover data.
	Data map[string]*string `json:"data"`
}

type _MultiSizeLogoCover MultiSizeLogoCover

// NewMultiSizeLogoCover instantiates a new MultiSizeLogoCover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiSizeLogoCover(id NullableString, data map[string]*string) *MultiSizeLogoCover {
	this := MultiSizeLogoCover{}
	this.Id = id
	this.Data = data
	return &this
}

// NewMultiSizeLogoCoverWithDefaults instantiates a new MultiSizeLogoCover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiSizeLogoCoverWithDefaults() *MultiSizeLogoCover {
	this := MultiSizeLogoCover{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MultiSizeLogoCover) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultiSizeLogoCover) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *MultiSizeLogoCover) SetId(v string) {
	o.Id.Set(&v)
}

// GetData returns the Data field value
func (o *MultiSizeLogoCover) GetData() map[string]*string {
	if o == nil {
		var ret map[string]*string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MultiSizeLogoCover) GetDataOk() (map[string]*string, bool) {
	if o == nil {
		return map[string]*string{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *MultiSizeLogoCover) SetData(v map[string]*string) {
	o.Data = v
}

func (o MultiSizeLogoCover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiSizeLogoCover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *MultiSizeLogoCover) UnmarshalJSON(data []byte) (err error) {
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

	varMultiSizeLogoCover := _MultiSizeLogoCover{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varMultiSizeLogoCover)

	if err != nil {
		return err
	}

	*o = MultiSizeLogoCover(varMultiSizeLogoCover)

	return err
}

type NullableMultiSizeLogoCover struct {
	value *MultiSizeLogoCover
	isSet bool
}

func (v NullableMultiSizeLogoCover) Get() *MultiSizeLogoCover {
	return v.value
}

func (v *NullableMultiSizeLogoCover) Set(val *MultiSizeLogoCover) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiSizeLogoCover) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiSizeLogoCover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiSizeLogoCover(val *MultiSizeLogoCover) *NullableMultiSizeLogoCover {
	return &NullableMultiSizeLogoCover{value: val, isSet: true}
}

func (v NullableMultiSizeLogoCover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiSizeLogoCover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

