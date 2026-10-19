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
)

// checks if the HistoryActionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionDto{}

// HistoryActionDto The action performed on the file.
type HistoryActionDto struct {
	// The action performed on the file.
	Id *MessageAction `json:"id,omitempty"`
	// The same action as a camel-case key, for a client to look up its own wording by.
	Key NullableString `json:"key,omitempty"`
}

// NewHistoryActionDto instantiates a new HistoryActionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionDto() *HistoryActionDto {
	this := HistoryActionDto{}
	return &this
}

// NewHistoryActionDtoWithDefaults instantiates a new HistoryActionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionDtoWithDefaults() *HistoryActionDto {
	this := HistoryActionDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HistoryActionDto) GetId() MessageAction {
	if o == nil || IsNil(o.Id) {
		var ret MessageAction
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryActionDto) GetIdOk() (*MessageAction, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HistoryActionDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given MessageAction and assigns it to the Id field.
func (o *HistoryActionDto) SetId(v MessageAction) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryActionDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryActionDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *HistoryActionDto) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *HistoryActionDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *HistoryActionDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *HistoryActionDto) UnsetKey() {
	o.Key.Unset()
}

func (o HistoryActionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	return toSerialize, nil
}

type NullableHistoryActionDto struct {
	value *HistoryActionDto
	isSet bool
}

func (v NullableHistoryActionDto) Get() *HistoryActionDto {
	return v.value
}

func (v *NullableHistoryActionDto) Set(val *HistoryActionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionDto(val *HistoryActionDto) *NullableHistoryActionDto {
	return &NullableHistoryActionDto{value: val, isSet: true}
}

func (v NullableHistoryActionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

