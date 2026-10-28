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

// checks if the UpdateTagRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTagRequestDto{}

// UpdateTagRequestDto The parameters for renaming a custom room tag in the portal catalog.
type UpdateTagRequestDto struct {
	// The name of the tag to rename, matched against the catalog exactly as it is stored rather than searched for.  Read the stored spelling from `GET api/2.0/files/tags`.
	OldName NullableString `json:"oldName"`
	// The name to store instead. It has to be free: names are unique across the portal, so a name another tag  already carries is refused, and merging two tags this way is not possible.
	NewName NullableString `json:"newName"`
}

type _UpdateTagRequestDto UpdateTagRequestDto

// NewUpdateTagRequestDto instantiates a new UpdateTagRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTagRequestDto(oldName NullableString, newName NullableString) *UpdateTagRequestDto {
	this := UpdateTagRequestDto{}
	this.OldName = oldName
	this.NewName = newName
	return &this
}

// NewUpdateTagRequestDtoWithDefaults instantiates a new UpdateTagRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTagRequestDtoWithDefaults() *UpdateTagRequestDto {
	this := UpdateTagRequestDto{}
	return &this
}

// GetOldName returns the OldName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateTagRequestDto) GetOldName() string {
	if o == nil || o.OldName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OldName.Get()
}

// GetOldNameOk returns a tuple with the OldName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTagRequestDto) GetOldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldName.Get(), o.OldName.IsSet()
}

// SetOldName sets field value
func (o *UpdateTagRequestDto) SetOldName(v string) {
	o.OldName.Set(&v)
}

// GetNewName returns the NewName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateTagRequestDto) GetNewName() string {
	if o == nil || o.NewName.Get() == nil {
		var ret string
		return ret
	}

	return *o.NewName.Get()
}

// GetNewNameOk returns a tuple with the NewName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTagRequestDto) GetNewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewName.Get(), o.NewName.IsSet()
}

// SetNewName sets field value
func (o *UpdateTagRequestDto) SetNewName(v string) {
	o.NewName.Set(&v)
}

func (o UpdateTagRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTagRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oldName"] = o.OldName.Get()
	toSerialize["newName"] = o.NewName.Get()
	return toSerialize, nil
}

func (o *UpdateTagRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"oldName",
		"newName",
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

	varUpdateTagRequestDto := _UpdateTagRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varUpdateTagRequestDto)

	if err != nil {
		return err
	}

	*o = UpdateTagRequestDto(varUpdateTagRequestDto)

	return err
}

type NullableUpdateTagRequestDto struct {
	value *UpdateTagRequestDto
	isSet bool
}

func (v NullableUpdateTagRequestDto) Get() *UpdateTagRequestDto {
	return v.value
}

func (v *NullableUpdateTagRequestDto) Set(val *UpdateTagRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTagRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTagRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTagRequestDto(val *UpdateTagRequestDto) *NullableUpdateTagRequestDto {
	return &NullableUpdateTagRequestDto{value: val, isSet: true}
}

func (v NullableUpdateTagRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTagRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

