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

// checks if the CreateFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFolder{}

// CreateFolder The parameters for creating a folder.
type CreateFolder struct {
	// The folder title to create.
	Title NullableString `json:"title"`
}

type _CreateFolder CreateFolder

// NewCreateFolder instantiates a new CreateFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFolder(title NullableString) *CreateFolder {
	this := CreateFolder{}
	this.Title = title
	return &this
}

// NewCreateFolderWithDefaults instantiates a new CreateFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFolderWithDefaults() *CreateFolder {
	this := CreateFolder{}
	return &this
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateFolder) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFolder) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateFolder) SetTitle(v string) {
	o.Title.Set(&v)
}

func (o CreateFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title.Get()
	return toSerialize, nil
}

func (o *CreateFolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varCreateFolder := _CreateFolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateFolder)

	if err != nil {
		return err
	}

	*o = CreateFolder(varCreateFolder)

	return err
}

type NullableCreateFolder struct {
	value *CreateFolder
	isSet bool
}

func (v NullableCreateFolder) Get() *CreateFolder {
	return v.value
}

func (v *NullableCreateFolder) Set(val *CreateFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFolder(val *CreateFolder) *NullableCreateFolder {
	return &NullableCreateFolder{value: val, isSet: true}
}

func (v NullableCreateFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

