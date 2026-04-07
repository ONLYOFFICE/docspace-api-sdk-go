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

// checks if the RenameChatBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenameChatBody{}

// RenameChatBody Parameters for renaming an AI chat session.
type RenameChatBody struct {
	// The new display name for the chat session (maximum 255 characters).
	Name NullableString `json:"name"`
}

type _RenameChatBody RenameChatBody

// NewRenameChatBody instantiates a new RenameChatBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenameChatBody(name NullableString) *RenameChatBody {
	this := RenameChatBody{}
	this.Name = name
	return &this
}

// NewRenameChatBodyWithDefaults instantiates a new RenameChatBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenameChatBodyWithDefaults() *RenameChatBody {
	this := RenameChatBody{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RenameChatBody) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RenameChatBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *RenameChatBody) SetName(v string) {
	o.Name.Set(&v)
}

func (o RenameChatBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenameChatBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	return toSerialize, nil
}

func (o *RenameChatBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varRenameChatBody := _RenameChatBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRenameChatBody)

	if err != nil {
		return err
	}

	*o = RenameChatBody(varRenameChatBody)

	return err
}

type NullableRenameChatBody struct {
	value *RenameChatBody
	isSet bool
}

func (v NullableRenameChatBody) Get() *RenameChatBody {
	return v.value
}

func (v *NullableRenameChatBody) Set(val *RenameChatBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRenameChatBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRenameChatBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenameChatBody(val *RenameChatBody) *NullableRenameChatBody {
	return &NullableRenameChatBody{value: val, isSet: true}
}

func (v NullableRenameChatBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenameChatBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

