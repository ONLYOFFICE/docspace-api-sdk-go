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

// checks if the EditHistoryAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHistoryAuthor{}

// EditHistoryAuthor The information about the file editing history author.
type EditHistoryAuthor struct {
	// The author ID.
	Id NullableString `json:"id"`
	// The author name.
	Name NullableString `json:"name,omitempty"`
}

type _EditHistoryAuthor EditHistoryAuthor

// NewEditHistoryAuthor instantiates a new EditHistoryAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHistoryAuthor(id NullableString) *EditHistoryAuthor {
	this := EditHistoryAuthor{}
	this.Id = id
	return &this
}

// NewEditHistoryAuthorWithDefaults instantiates a new EditHistoryAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHistoryAuthorWithDefaults() *EditHistoryAuthor {
	this := EditHistoryAuthor{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EditHistoryAuthor) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryAuthor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *EditHistoryAuthor) SetId(v string) {
	o.Id.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryAuthor) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryAuthor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EditHistoryAuthor) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EditHistoryAuthor) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EditHistoryAuthor) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EditHistoryAuthor) UnsetName() {
	o.Name.Unset()
}

func (o EditHistoryAuthor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHistoryAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

func (o *EditHistoryAuthor) UnmarshalJSON(data []byte) (err error) {
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

	varEditHistoryAuthor := _EditHistoryAuthor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varEditHistoryAuthor)

	if err != nil {
		return err
	}

	*o = EditHistoryAuthor(varEditHistoryAuthor)

	return err
}

type NullableEditHistoryAuthor struct {
	value *EditHistoryAuthor
	isSet bool
}

func (v NullableEditHistoryAuthor) Get() *EditHistoryAuthor {
	return v.value
}

func (v *NullableEditHistoryAuthor) Set(val *EditHistoryAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHistoryAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHistoryAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHistoryAuthor(val *EditHistoryAuthor) *NullableEditHistoryAuthor {
	return &NullableEditHistoryAuthor{value: val, isSet: true}
}

func (v NullableEditHistoryAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHistoryAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

