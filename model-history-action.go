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

// checks if the HistoryAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryAction{}

// HistoryAction The action performed on the file.
type HistoryAction struct {
	// The action performed on the file.
	Id *MessageAction `json:"id,omitempty"`
	// The action performed on the file.
	Key NullableString `json:"key,omitempty"`
}

// NewHistoryAction instantiates a new HistoryAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryAction() *HistoryAction {
	this := HistoryAction{}
	return &this
}

// NewHistoryActionWithDefaults instantiates a new HistoryAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionWithDefaults() *HistoryAction {
	this := HistoryAction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HistoryAction) GetId() MessageAction {
	if o == nil || IsNil(o.Id) {
		var ret MessageAction
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryAction) GetIdOk() (*MessageAction, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HistoryAction) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given MessageAction and assigns it to the Id field.
func (o *HistoryAction) SetId(v MessageAction) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryAction) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryAction) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *HistoryAction) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *HistoryAction) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *HistoryAction) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *HistoryAction) UnsetKey() {
	o.Key.Unset()
}

func (o HistoryAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	return toSerialize, nil
}

type NullableHistoryAction struct {
	value *HistoryAction
	isSet bool
}

func (v NullableHistoryAction) Get() *HistoryAction {
	return v.value
}

func (v *NullableHistoryAction) Set(val *HistoryAction) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryAction) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryAction(val *HistoryAction) *NullableHistoryAction {
	return &NullableHistoryAction{value: val, isSet: true}
}

func (v NullableHistoryAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

