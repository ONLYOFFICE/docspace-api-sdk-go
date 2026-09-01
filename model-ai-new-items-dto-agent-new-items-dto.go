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

// checks if the AiNewItemsDtoAgentNewItemsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiNewItemsDtoAgentNewItemsDto{}

// AiNewItemsDtoAgentNewItemsDto The new item parameters.
type AiNewItemsDtoAgentNewItemsDto struct {
	// The date and time when the new item was created.
	Date NullableTime `json:"date"`
	// The list of items.
	Items []AiAgentNewItemsDto `json:"items"`
}

type _AiNewItemsDtoAgentNewItemsDto AiNewItemsDtoAgentNewItemsDto

// NewAiNewItemsDtoAgentNewItemsDto instantiates a new AiNewItemsDtoAgentNewItemsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiNewItemsDtoAgentNewItemsDto(date NullableTime, items []AiAgentNewItemsDto) *AiNewItemsDtoAgentNewItemsDto {
	this := AiNewItemsDtoAgentNewItemsDto{}
	this.Date = date
	this.Items = items
	return &this
}

// NewAiNewItemsDtoAgentNewItemsDtoWithDefaults instantiates a new AiNewItemsDtoAgentNewItemsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiNewItemsDtoAgentNewItemsDtoWithDefaults() *AiNewItemsDtoAgentNewItemsDto {
	this := AiNewItemsDtoAgentNewItemsDto{}
	return &this
}

// GetDate returns the Date field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AiNewItemsDtoAgentNewItemsDto) GetDate() time.Time {
	if o == nil || o.Date.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiNewItemsDtoAgentNewItemsDto) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// SetDate sets field value
func (o *AiNewItemsDtoAgentNewItemsDto) SetDate(v time.Time) {
	o.Date.Set(&v)
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []AiAgentNewItemsDto will be returned
func (o *AiNewItemsDtoAgentNewItemsDto) GetItems() []AiAgentNewItemsDto {
	if o == nil {
		var ret []AiAgentNewItemsDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiNewItemsDtoAgentNewItemsDto) GetItemsOk() ([]AiAgentNewItemsDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *AiNewItemsDtoAgentNewItemsDto) SetItems(v []AiAgentNewItemsDto) {
	o.Items = v
}

func (o AiNewItemsDtoAgentNewItemsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiNewItemsDtoAgentNewItemsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date.Get()
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *AiNewItemsDtoAgentNewItemsDto) UnmarshalJSON(data []byte) (err error) {
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

	varAiNewItemsDtoAgentNewItemsDto := _AiNewItemsDtoAgentNewItemsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiNewItemsDtoAgentNewItemsDto)

	if err != nil {
		return err
	}

	*o = AiNewItemsDtoAgentNewItemsDto(varAiNewItemsDtoAgentNewItemsDto)

	return err
}

type NullableAiNewItemsDtoAgentNewItemsDto struct {
	value *AiNewItemsDtoAgentNewItemsDto
	isSet bool
}

func (v NullableAiNewItemsDtoAgentNewItemsDto) Get() *AiNewItemsDtoAgentNewItemsDto {
	return v.value
}

func (v *NullableAiNewItemsDtoAgentNewItemsDto) Set(val *AiNewItemsDtoAgentNewItemsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiNewItemsDtoAgentNewItemsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiNewItemsDtoAgentNewItemsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiNewItemsDtoAgentNewItemsDto(val *AiNewItemsDtoAgentNewItemsDto) *NullableAiNewItemsDtoAgentNewItemsDto {
	return &NullableAiNewItemsDtoAgentNewItemsDto{value: val, isSet: true}
}

func (v NullableAiNewItemsDtoAgentNewItemsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiNewItemsDtoAgentNewItemsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

