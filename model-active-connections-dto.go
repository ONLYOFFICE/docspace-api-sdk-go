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

// checks if the ActiveConnectionsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveConnectionsDto{}

// ActiveConnectionsDto The active connections parameters.
type ActiveConnectionsDto struct {
	// The login event.
	LoginEvent int32 `json:"loginEvent"`
	// The list of active connection items.
	Items []ActiveConnectionsItemDto `json:"items,omitempty"`
}

type _ActiveConnectionsDto ActiveConnectionsDto

// NewActiveConnectionsDto instantiates a new ActiveConnectionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveConnectionsDto(loginEvent int32) *ActiveConnectionsDto {
	this := ActiveConnectionsDto{}
	this.LoginEvent = loginEvent
	return &this
}

// NewActiveConnectionsDtoWithDefaults instantiates a new ActiveConnectionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveConnectionsDtoWithDefaults() *ActiveConnectionsDto {
	this := ActiveConnectionsDto{}
	return &this
}

// GetLoginEvent returns the LoginEvent field value
func (o *ActiveConnectionsDto) GetLoginEvent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoginEvent
}

// GetLoginEventOk returns a tuple with the LoginEvent field value
// and a boolean to check if the value has been set.
func (o *ActiveConnectionsDto) GetLoginEventOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginEvent, true
}

// SetLoginEvent sets field value
func (o *ActiveConnectionsDto) SetLoginEvent(v int32) {
	o.LoginEvent = v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveConnectionsDto) GetItems() []ActiveConnectionsItemDto {
	if o == nil {
		var ret []ActiveConnectionsItemDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveConnectionsDto) GetItemsOk() ([]ActiveConnectionsItemDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ActiveConnectionsDto) IsItemsSet() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ActiveConnectionsItemDto and assigns it to the Items field.
func (o *ActiveConnectionsDto) SetItems(v []ActiveConnectionsItemDto) {
	o.Items = v
}

func (o ActiveConnectionsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveConnectionsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["loginEvent"] = o.LoginEvent
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *ActiveConnectionsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"loginEvent",
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

	varActiveConnectionsDto := _ActiveConnectionsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varActiveConnectionsDto)

	if err != nil {
		return err
	}

	*o = ActiveConnectionsDto(varActiveConnectionsDto)

	return err
}

type NullableActiveConnectionsDto struct {
	value *ActiveConnectionsDto
	isSet bool
}

func (v NullableActiveConnectionsDto) Get() *ActiveConnectionsDto {
	return v.value
}

func (v *NullableActiveConnectionsDto) Set(val *ActiveConnectionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveConnectionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveConnectionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveConnectionsDto(val *ActiveConnectionsDto) *NullableActiveConnectionsDto {
	return &NullableActiveConnectionsDto{value: val, isSet: true}
}

func (v NullableActiveConnectionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveConnectionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

