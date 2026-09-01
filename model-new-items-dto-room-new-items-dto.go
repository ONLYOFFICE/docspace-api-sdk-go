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

// checks if the NewItemsDtoRoomNewItemsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewItemsDtoRoomNewItemsDto{}

// NewItemsDtoRoomNewItemsDto The new item parameters.
type NewItemsDtoRoomNewItemsDto struct {
	// The date and time when the new item was created.
	Date NullableTime `json:"date"`
	// The list of items.
	Items []RoomNewItemsDto `json:"items"`
}

type _NewItemsDtoRoomNewItemsDto NewItemsDtoRoomNewItemsDto

// NewNewItemsDtoRoomNewItemsDto instantiates a new NewItemsDtoRoomNewItemsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewItemsDtoRoomNewItemsDto(date NullableTime, items []RoomNewItemsDto) *NewItemsDtoRoomNewItemsDto {
	this := NewItemsDtoRoomNewItemsDto{}
	this.Date = date
	this.Items = items
	return &this
}

// NewNewItemsDtoRoomNewItemsDtoWithDefaults instantiates a new NewItemsDtoRoomNewItemsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewItemsDtoRoomNewItemsDtoWithDefaults() *NewItemsDtoRoomNewItemsDto {
	this := NewItemsDtoRoomNewItemsDto{}
	return &this
}

// GetDate returns the Date field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *NewItemsDtoRoomNewItemsDto) GetDate() time.Time {
	if o == nil || o.Date.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewItemsDtoRoomNewItemsDto) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// SetDate sets field value
func (o *NewItemsDtoRoomNewItemsDto) SetDate(v time.Time) {
	o.Date.Set(&v)
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []RoomNewItemsDto will be returned
func (o *NewItemsDtoRoomNewItemsDto) GetItems() []RoomNewItemsDto {
	if o == nil {
		var ret []RoomNewItemsDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewItemsDtoRoomNewItemsDto) GetItemsOk() ([]RoomNewItemsDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NewItemsDtoRoomNewItemsDto) SetItems(v []RoomNewItemsDto) {
	o.Items = v
}

func (o NewItemsDtoRoomNewItemsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewItemsDtoRoomNewItemsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date.Get()
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *NewItemsDtoRoomNewItemsDto) UnmarshalJSON(data []byte) (err error) {
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

	varNewItemsDtoRoomNewItemsDto := _NewItemsDtoRoomNewItemsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varNewItemsDtoRoomNewItemsDto)

	if err != nil {
		return err
	}

	*o = NewItemsDtoRoomNewItemsDto(varNewItemsDtoRoomNewItemsDto)

	return err
}

type NullableNewItemsDtoRoomNewItemsDto struct {
	value *NewItemsDtoRoomNewItemsDto
	isSet bool
}

func (v NullableNewItemsDtoRoomNewItemsDto) Get() *NewItemsDtoRoomNewItemsDto {
	return v.value
}

func (v *NullableNewItemsDtoRoomNewItemsDto) Set(val *NewItemsDtoRoomNewItemsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNewItemsDtoRoomNewItemsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNewItemsDtoRoomNewItemsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewItemsDtoRoomNewItemsDto(val *NewItemsDtoRoomNewItemsDto) *NullableNewItemsDtoRoomNewItemsDto {
	return &NullableNewItemsDtoRoomNewItemsDto{value: val, isSet: true}
}

func (v NullableNewItemsDtoRoomNewItemsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewItemsDtoRoomNewItemsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

