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

// checks if the GroupSummaryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupSummaryDto{}

// GroupSummaryDto The group summary parameters.
type GroupSummaryDto struct {
	// The group ID.
	Id string `json:"id"`
	// The group name.
	Name NullableString `json:"name"`
	// The group manager.
	Manager NullableString `json:"manager,omitempty"`
	// Indicates whether the group is a system group.
	IsSystem NullableBool `json:"isSystem,omitempty"`
}

type _GroupSummaryDto GroupSummaryDto

// NewGroupSummaryDto instantiates a new GroupSummaryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSummaryDto(id string, name NullableString) *GroupSummaryDto {
	this := GroupSummaryDto{}
	this.Id = id
	this.Name = name
	return &this
}

// NewGroupSummaryDtoWithDefaults instantiates a new GroupSummaryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSummaryDtoWithDefaults() *GroupSummaryDto {
	this := GroupSummaryDto{}
	return &this
}

// GetId returns the Id field value
func (o *GroupSummaryDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupSummaryDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupSummaryDto) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupSummaryDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSummaryDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *GroupSummaryDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSummaryDto) GetManager() string {
	if o == nil || IsNil(o.Manager.Get()) {
		var ret string
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSummaryDto) GetManagerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *GroupSummaryDto) IsManagerSet() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableString and assigns it to the Manager field.
func (o *GroupSummaryDto) SetManager(v string) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *GroupSummaryDto) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *GroupSummaryDto) UnsetManager() {
	o.Manager.Unset()
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSummaryDto) GetIsSystem() bool {
	if o == nil || IsNil(o.IsSystem.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSystem.Get()
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSummaryDto) GetIsSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSystem.Get(), o.IsSystem.IsSet()
}

// HasIsSystem returns a boolean if a field has been set.
func (o *GroupSummaryDto) IsIsSystemSet() bool {
	if o != nil && o.IsSystem.IsSet() {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given NullableBool and assigns it to the IsSystem field.
func (o *GroupSummaryDto) SetIsSystem(v bool) {
	o.IsSystem.Set(&v)
}
// SetIsSystemNil sets the value for IsSystem to be an explicit nil
func (o *GroupSummaryDto) SetIsSystemNil() {
	o.IsSystem.Set(nil)
}

// UnsetIsSystem ensures that no value is present for IsSystem, not even an explicit nil
func (o *GroupSummaryDto) UnsetIsSystem() {
	o.IsSystem.Unset()
}

func (o GroupSummaryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupSummaryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	if o.IsSystem.IsSet() {
		toSerialize["isSystem"] = o.IsSystem.Get()
	}
	return toSerialize, nil
}

func (o *GroupSummaryDto) UnmarshalJSON(data []byte) (err error) {
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

	varGroupSummaryDto := _GroupSummaryDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGroupSummaryDto)

	if err != nil {
		return err
	}

	*o = GroupSummaryDto(varGroupSummaryDto)

	return err
}

type NullableGroupSummaryDto struct {
	value *GroupSummaryDto
	isSet bool
}

func (v NullableGroupSummaryDto) Get() *GroupSummaryDto {
	return v.value
}

func (v *NullableGroupSummaryDto) Set(val *GroupSummaryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSummaryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSummaryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSummaryDto(val *GroupSummaryDto) *NullableGroupSummaryDto {
	return &NullableGroupSummaryDto{value: val, isSet: true}
}

func (v NullableGroupSummaryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSummaryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

