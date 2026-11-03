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

// checks if the AiContextRoom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiContextRoom{}

// AiContextRoom A room the caller can connect to a chat as context.
type AiContextRoom struct {
	// The room ID.
	Id string `json:"id"`
	// The room title.
	Name string `json:"name"`
}

type _AiContextRoom AiContextRoom

// NewAiContextRoom instantiates a new AiContextRoom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiContextRoom(id string, name string) *AiContextRoom {
	this := AiContextRoom{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAiContextRoomWithDefaults instantiates a new AiContextRoom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiContextRoomWithDefaults() *AiContextRoom {
	this := AiContextRoom{}
	return &this
}

// GetId returns the Id field value
func (o *AiContextRoom) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiContextRoom) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiContextRoom) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiContextRoom) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiContextRoom) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiContextRoom) SetName(v string) {
	o.Name = v
}

func (o AiContextRoom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiContextRoom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *AiContextRoom) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varAiContextRoom := _AiContextRoom{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiContextRoom)

	if err != nil {
		return err
	}

	*o = AiContextRoom(varAiContextRoom)

	return err
}

type NullableAiContextRoom struct {
	value *AiContextRoom
	isSet bool
}

func (v NullableAiContextRoom) Get() *AiContextRoom {
	return v.value
}

func (v *NullableAiContextRoom) Set(val *AiContextRoom) {
	v.value = val
	v.isSet = true
}

func (v NullableAiContextRoom) IsSet() bool {
	return v.isSet
}

func (v *NullableAiContextRoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiContextRoom(val *AiContextRoom) *NullableAiContextRoom {
	return &NullableAiContextRoom{value: val, isSet: true}
}

func (v NullableAiContextRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiContextRoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

