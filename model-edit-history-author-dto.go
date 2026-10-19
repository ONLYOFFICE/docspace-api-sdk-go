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

// checks if the EditHistoryAuthorDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHistoryAuthorDto{}

// EditHistoryAuthorDto The person a saved revision of a file, or one single change in it, is attributed to.
type EditHistoryAuthorDto struct {
	// The account the revision or the change is attributed to, as the editing service stored it. It is normally the  identifier of a portal account; the empty identifier stands for a change nobody could be named for.
	Id NullableString `json:"id"`
	// The display name of that account as the portal spells it now, which need not be the name that was stored with  the revision. An account that cannot be resolved - one removed from the portal, or a change made through an  anonymous link - is reported as a guest.
	Name NullableString `json:"name,omitempty"`
}

type _EditHistoryAuthorDto EditHistoryAuthorDto

// NewEditHistoryAuthorDto instantiates a new EditHistoryAuthorDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHistoryAuthorDto(id NullableString) *EditHistoryAuthorDto {
	this := EditHistoryAuthorDto{}
	this.Id = id
	return &this
}

// NewEditHistoryAuthorDtoWithDefaults instantiates a new EditHistoryAuthorDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHistoryAuthorDtoWithDefaults() *EditHistoryAuthorDto {
	this := EditHistoryAuthorDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EditHistoryAuthorDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryAuthorDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *EditHistoryAuthorDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryAuthorDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryAuthorDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EditHistoryAuthorDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EditHistoryAuthorDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EditHistoryAuthorDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EditHistoryAuthorDto) UnsetName() {
	o.Name.Unset()
}

func (o EditHistoryAuthorDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHistoryAuthorDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

func (o *EditHistoryAuthorDto) UnmarshalJSON(data []byte) (err error) {
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

	varEditHistoryAuthorDto := _EditHistoryAuthorDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varEditHistoryAuthorDto)

	if err != nil {
		return err
	}

	*o = EditHistoryAuthorDto(varEditHistoryAuthorDto)

	return err
}

type NullableEditHistoryAuthorDto struct {
	value *EditHistoryAuthorDto
	isSet bool
}

func (v NullableEditHistoryAuthorDto) Get() *EditHistoryAuthorDto {
	return v.value
}

func (v *NullableEditHistoryAuthorDto) Set(val *EditHistoryAuthorDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHistoryAuthorDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHistoryAuthorDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHistoryAuthorDto(val *EditHistoryAuthorDto) *NullableEditHistoryAuthorDto {
	return &NullableEditHistoryAuthorDto{value: val, isSet: true}
}

func (v NullableEditHistoryAuthorDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHistoryAuthorDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

