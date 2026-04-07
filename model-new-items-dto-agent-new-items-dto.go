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

// checks if the NewItemsDtoAgentNewItemsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewItemsDtoAgentNewItemsDto{}

// NewItemsDtoAgentNewItemsDto The new item parameters.
type NewItemsDtoAgentNewItemsDto struct {
	Date ApiDateTime `json:"date"`
	// The list of items.
	Items []AgentNewItemsDto `json:"items"`
}

type _NewItemsDtoAgentNewItemsDto NewItemsDtoAgentNewItemsDto

// NewNewItemsDtoAgentNewItemsDto instantiates a new NewItemsDtoAgentNewItemsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewItemsDtoAgentNewItemsDto(date ApiDateTime, items []AgentNewItemsDto) *NewItemsDtoAgentNewItemsDto {
	this := NewItemsDtoAgentNewItemsDto{}
	this.Date = date
	this.Items = items
	return &this
}

// NewNewItemsDtoAgentNewItemsDtoWithDefaults instantiates a new NewItemsDtoAgentNewItemsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewItemsDtoAgentNewItemsDtoWithDefaults() *NewItemsDtoAgentNewItemsDto {
	this := NewItemsDtoAgentNewItemsDto{}
	return &this
}

// GetDate returns the Date field value
func (o *NewItemsDtoAgentNewItemsDto) GetDate() ApiDateTime {
	if o == nil {
		var ret ApiDateTime
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *NewItemsDtoAgentNewItemsDto) GetDateOk() (*ApiDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *NewItemsDtoAgentNewItemsDto) SetDate(v ApiDateTime) {
	o.Date = v
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []AgentNewItemsDto will be returned
func (o *NewItemsDtoAgentNewItemsDto) GetItems() []AgentNewItemsDto {
	if o == nil {
		var ret []AgentNewItemsDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewItemsDtoAgentNewItemsDto) GetItemsOk() ([]AgentNewItemsDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NewItemsDtoAgentNewItemsDto) SetItems(v []AgentNewItemsDto) {
	o.Items = v
}

func (o NewItemsDtoAgentNewItemsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewItemsDtoAgentNewItemsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *NewItemsDtoAgentNewItemsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"items",
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

	varNewItemsDtoAgentNewItemsDto := _NewItemsDtoAgentNewItemsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varNewItemsDtoAgentNewItemsDto)

	if err != nil {
		return err
	}

	*o = NewItemsDtoAgentNewItemsDto(varNewItemsDtoAgentNewItemsDto)

	return err
}

type NullableNewItemsDtoAgentNewItemsDto struct {
	value *NewItemsDtoAgentNewItemsDto
	isSet bool
}

func (v NullableNewItemsDtoAgentNewItemsDto) Get() *NewItemsDtoAgentNewItemsDto {
	return v.value
}

func (v *NullableNewItemsDtoAgentNewItemsDto) Set(val *NewItemsDtoAgentNewItemsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNewItemsDtoAgentNewItemsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNewItemsDtoAgentNewItemsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewItemsDtoAgentNewItemsDto(val *NewItemsDtoAgentNewItemsDto) *NullableNewItemsDtoAgentNewItemsDto {
	return &NullableNewItemsDtoAgentNewItemsDto{value: val, isSet: true}
}

func (v NullableNewItemsDtoAgentNewItemsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewItemsDtoAgentNewItemsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

