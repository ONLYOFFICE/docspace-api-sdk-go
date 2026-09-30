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

// checks if the SetPublicDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetPublicDto{}

// SetPublicDto The public access to set on a room template.
type SetPublicDto struct {
	// The identifier of the room template. Take it from `templateId` of `GET api/2.0/files/roomtemplate/status`, or  from the folder list of `GET api/2.0/files/rooms` called with `searchArea` set to 4; an identifier of an  ordinary room is not accepted.
	Id int32 `json:"id"`
	// Whether the Everyone group keeps read access to the template. True shares it with every member allowed to  create rooms; false leaves it reachable only for its owner.
	Public *bool `json:"public,omitempty"`
}

type _SetPublicDto SetPublicDto

// NewSetPublicDto instantiates a new SetPublicDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPublicDto(id int32) *SetPublicDto {
	this := SetPublicDto{}
	this.Id = id
	return &this
}

// NewSetPublicDtoWithDefaults instantiates a new SetPublicDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPublicDtoWithDefaults() *SetPublicDto {
	this := SetPublicDto{}
	return &this
}

// GetId returns the Id field value
func (o *SetPublicDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SetPublicDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SetPublicDto) SetId(v int32) {
	o.Id = v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *SetPublicDto) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetPublicDto) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *SetPublicDto) IsPublicSet() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *SetPublicDto) SetPublic(v bool) {
	o.Public = &v
}

func (o SetPublicDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetPublicDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	return toSerialize, nil
}

func (o *SetPublicDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varSetPublicDto := _SetPublicDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSetPublicDto)

	if err != nil {
		return err
	}

	*o = SetPublicDto(varSetPublicDto)

	return err
}

type NullableSetPublicDto struct {
	value *SetPublicDto
	isSet bool
}

func (v NullableSetPublicDto) Get() *SetPublicDto {
	return v.value
}

func (v *NullableSetPublicDto) Set(val *SetPublicDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPublicDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPublicDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPublicDto(val *SetPublicDto) *NullableSetPublicDto {
	return &NullableSetPublicDto{value: val, isSet: true}
}

func (v NullableSetPublicDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetPublicDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

