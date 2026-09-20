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

// checks if the NewItemsDtoFileEntryBaseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewItemsDtoFileEntryBaseDto{}

// NewItemsDtoFileEntryBaseDto One day of the entries the caller has not opened yet, the groups running from the most recent day backwards.
type NewItemsDtoFileEntryBaseDto struct {
	// The day the grouped entries were last changed, written with the offset of the portal time zone. The time part  is the moment of the newest entry of the group.
	Date ApiDateTime `json:"date"`
	// What changed on that day, the most recent first. Folders are left out of it, so an entry here is always a file  or a room that holds them.
	Items []FileEntryBaseDto `json:"items"`
}

type _NewItemsDtoFileEntryBaseDto NewItemsDtoFileEntryBaseDto

// NewNewItemsDtoFileEntryBaseDto instantiates a new NewItemsDtoFileEntryBaseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewItemsDtoFileEntryBaseDto(date ApiDateTime, items []FileEntryBaseDto) *NewItemsDtoFileEntryBaseDto {
	this := NewItemsDtoFileEntryBaseDto{}
	this.Date = date
	this.Items = items
	return &this
}

// NewNewItemsDtoFileEntryBaseDtoWithDefaults instantiates a new NewItemsDtoFileEntryBaseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewItemsDtoFileEntryBaseDtoWithDefaults() *NewItemsDtoFileEntryBaseDto {
	this := NewItemsDtoFileEntryBaseDto{}
	return &this
}

// GetDate returns the Date field value
func (o *NewItemsDtoFileEntryBaseDto) GetDate() ApiDateTime {
	if o == nil {
		var ret ApiDateTime
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *NewItemsDtoFileEntryBaseDto) GetDateOk() (*ApiDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *NewItemsDtoFileEntryBaseDto) SetDate(v ApiDateTime) {
	o.Date = v
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []FileEntryBaseDto will be returned
func (o *NewItemsDtoFileEntryBaseDto) GetItems() []FileEntryBaseDto {
	if o == nil {
		var ret []FileEntryBaseDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewItemsDtoFileEntryBaseDto) GetItemsOk() ([]FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NewItemsDtoFileEntryBaseDto) SetItems(v []FileEntryBaseDto) {
	o.Items = v
}

func (o NewItemsDtoFileEntryBaseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewItemsDtoFileEntryBaseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *NewItemsDtoFileEntryBaseDto) UnmarshalJSON(data []byte) (err error) {
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

	varNewItemsDtoFileEntryBaseDto := _NewItemsDtoFileEntryBaseDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varNewItemsDtoFileEntryBaseDto)

	if err != nil {
		return err
	}

	*o = NewItemsDtoFileEntryBaseDto(varNewItemsDtoFileEntryBaseDto)

	return err
}

type NullableNewItemsDtoFileEntryBaseDto struct {
	value *NewItemsDtoFileEntryBaseDto
	isSet bool
}

func (v NullableNewItemsDtoFileEntryBaseDto) Get() *NewItemsDtoFileEntryBaseDto {
	return v.value
}

func (v *NullableNewItemsDtoFileEntryBaseDto) Set(val *NewItemsDtoFileEntryBaseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNewItemsDtoFileEntryBaseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNewItemsDtoFileEntryBaseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewItemsDtoFileEntryBaseDto(val *NewItemsDtoFileEntryBaseDto) *NullableNewItemsDtoFileEntryBaseDto {
	return &NullableNewItemsDtoFileEntryBaseDto{value: val, isSet: true}
}

func (v NullableNewItemsDtoFileEntryBaseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewItemsDtoFileEntryBaseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

