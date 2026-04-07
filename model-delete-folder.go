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

// checks if the DeleteFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteFolder{}

// DeleteFolder The parameters for deleting a folder.
type DeleteFolder struct {
	// Specifies whether to delete a folder after the editing session is finished or not.
	DeleteAfter *bool `json:"deleteAfter,omitempty"`
	// Specifies whether to move a folder to the \\Trash\\ folder or delete it immediately.
	Immediately *bool `json:"immediately,omitempty"`
}

// NewDeleteFolder instantiates a new DeleteFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFolder() *DeleteFolder {
	this := DeleteFolder{}
	return &this
}

// NewDeleteFolderWithDefaults instantiates a new DeleteFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFolderWithDefaults() *DeleteFolder {
	this := DeleteFolder{}
	return &this
}

// GetDeleteAfter returns the DeleteAfter field value if set, zero value otherwise.
func (o *DeleteFolder) GetDeleteAfter() bool {
	if o == nil || IsNil(o.DeleteAfter) {
		var ret bool
		return ret
	}
	return *o.DeleteAfter
}

// GetDeleteAfterOk returns a tuple with the DeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFolder) GetDeleteAfterOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteAfter) {
		return nil, false
	}
	return o.DeleteAfter, true
}

// HasDeleteAfter returns a boolean if a field has been set.
func (o *DeleteFolder) IsDeleteAfterSet() bool {
	if o != nil && !IsNil(o.DeleteAfter) {
		return true
	}

	return false
}

// SetDeleteAfter gets a reference to the given bool and assigns it to the DeleteAfter field.
func (o *DeleteFolder) SetDeleteAfter(v bool) {
	o.DeleteAfter = &v
}

// GetImmediately returns the Immediately field value if set, zero value otherwise.
func (o *DeleteFolder) GetImmediately() bool {
	if o == nil || IsNil(o.Immediately) {
		var ret bool
		return ret
	}
	return *o.Immediately
}

// GetImmediatelyOk returns a tuple with the Immediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFolder) GetImmediatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.Immediately) {
		return nil, false
	}
	return o.Immediately, true
}

// HasImmediately returns a boolean if a field has been set.
func (o *DeleteFolder) IsImmediatelySet() bool {
	if o != nil && !IsNil(o.Immediately) {
		return true
	}

	return false
}

// SetImmediately gets a reference to the given bool and assigns it to the Immediately field.
func (o *DeleteFolder) SetImmediately(v bool) {
	o.Immediately = &v
}

func (o DeleteFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteAfter) {
		toSerialize["deleteAfter"] = o.DeleteAfter
	}
	if !IsNil(o.Immediately) {
		toSerialize["immediately"] = o.Immediately
	}
	return toSerialize, nil
}

type NullableDeleteFolder struct {
	value *DeleteFolder
	isSet bool
}

func (v NullableDeleteFolder) Get() *DeleteFolder {
	return v.value
}

func (v *NullableDeleteFolder) Set(val *DeleteFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFolder(val *DeleteFolder) *NullableDeleteFolder {
	return &NullableDeleteFolder{value: val, isSet: true}
}

func (v NullableDeleteFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

