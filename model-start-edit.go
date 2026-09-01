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

// checks if the StartEdit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartEdit{}

// StartEdit The parameters for starting file editing.
type StartEdit struct {
	// Specifies whether to share the file with other users for editing or not.
	EditingAlone *bool `json:"editingAlone,omitempty"`
}

// NewStartEdit instantiates a new StartEdit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartEdit() *StartEdit {
	this := StartEdit{}
	return &this
}

// NewStartEditWithDefaults instantiates a new StartEdit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartEditWithDefaults() *StartEdit {
	this := StartEdit{}
	return &this
}

// GetEditingAlone returns the EditingAlone field value if set, zero value otherwise.
func (o *StartEdit) GetEditingAlone() bool {
	if o == nil || IsNil(o.EditingAlone) {
		var ret bool
		return ret
	}
	return *o.EditingAlone
}

// GetEditingAloneOk returns a tuple with the EditingAlone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartEdit) GetEditingAloneOk() (*bool, bool) {
	if o == nil || IsNil(o.EditingAlone) {
		return nil, false
	}
	return o.EditingAlone, true
}

// HasEditingAlone returns a boolean if a field has been set.
func (o *StartEdit) IsEditingAloneSet() bool {
	if o != nil && !IsNil(o.EditingAlone) {
		return true
	}

	return false
}

// SetEditingAlone gets a reference to the given bool and assigns it to the EditingAlone field.
func (o *StartEdit) SetEditingAlone(v bool) {
	o.EditingAlone = &v
}

func (o StartEdit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartEdit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EditingAlone) {
		toSerialize["editingAlone"] = o.EditingAlone
	}
	return toSerialize, nil
}

type NullableStartEdit struct {
	value *StartEdit
	isSet bool
}

func (v NullableStartEdit) Get() *StartEdit {
	return v.value
}

func (v *NullableStartEdit) Set(val *StartEdit) {
	v.value = val
	v.isSet = true
}

func (v NullableStartEdit) IsSet() bool {
	return v.isSet
}

func (v *NullableStartEdit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartEdit(val *StartEdit) *NullableStartEdit {
	return &NullableStartEdit{value: val, isSet: true}
}

func (v NullableStartEdit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartEdit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

