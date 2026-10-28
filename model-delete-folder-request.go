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

// checks if the DeleteFolderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteFolderRequest{}

// DeleteFolderRequest How a folder is to be deleted.
type DeleteFolderRequest struct {
	// Whether the deletion waits for the editing sessions on the contents to end: with true a folder somebody is  working in is removed once they are done, with false the deletion starts at once.
	DeleteAfter *bool `json:"deleteAfter,omitempty"`
	// Whether the folder is discarded for good instead of being moved to the Trash section: with false it can be  restored from Trash, with true it cannot be recovered. Inside a room there is no Trash and the deletion is  final either way.
	Immediately *bool `json:"immediately,omitempty"`
}

// NewDeleteFolderRequest instantiates a new DeleteFolderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFolderRequest() *DeleteFolderRequest {
	this := DeleteFolderRequest{}
	return &this
}

// NewDeleteFolderRequestWithDefaults instantiates a new DeleteFolderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFolderRequestWithDefaults() *DeleteFolderRequest {
	this := DeleteFolderRequest{}
	return &this
}

// GetDeleteAfter returns the DeleteAfter field value if set, zero value otherwise.
func (o *DeleteFolderRequest) GetDeleteAfter() bool {
	if o == nil || IsNil(o.DeleteAfter) {
		var ret bool
		return ret
	}
	return *o.DeleteAfter
}

// GetDeleteAfterOk returns a tuple with the DeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFolderRequest) GetDeleteAfterOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteAfter) {
		return nil, false
	}
	return o.DeleteAfter, true
}

// HasDeleteAfter returns a boolean if a field has been set.
func (o *DeleteFolderRequest) IsDeleteAfterSet() bool {
	if o != nil && !IsNil(o.DeleteAfter) {
		return true
	}

	return false
}

// SetDeleteAfter gets a reference to the given bool and assigns it to the DeleteAfter field.
func (o *DeleteFolderRequest) SetDeleteAfter(v bool) {
	o.DeleteAfter = &v
}

// GetImmediately returns the Immediately field value if set, zero value otherwise.
func (o *DeleteFolderRequest) GetImmediately() bool {
	if o == nil || IsNil(o.Immediately) {
		var ret bool
		return ret
	}
	return *o.Immediately
}

// GetImmediatelyOk returns a tuple with the Immediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFolderRequest) GetImmediatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.Immediately) {
		return nil, false
	}
	return o.Immediately, true
}

// HasImmediately returns a boolean if a field has been set.
func (o *DeleteFolderRequest) IsImmediatelySet() bool {
	if o != nil && !IsNil(o.Immediately) {
		return true
	}

	return false
}

// SetImmediately gets a reference to the given bool and assigns it to the Immediately field.
func (o *DeleteFolderRequest) SetImmediately(v bool) {
	o.Immediately = &v
}

func (o DeleteFolderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteFolderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteAfter) {
		toSerialize["deleteAfter"] = o.DeleteAfter
	}
	if !IsNil(o.Immediately) {
		toSerialize["immediately"] = o.Immediately
	}
	return toSerialize, nil
}

type NullableDeleteFolderRequest struct {
	value *DeleteFolderRequest
	isSet bool
}

func (v NullableDeleteFolderRequest) Get() *DeleteFolderRequest {
	return v.value
}

func (v *NullableDeleteFolderRequest) Set(val *DeleteFolderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFolderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFolderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFolderRequest(val *DeleteFolderRequest) *NullableDeleteFolderRequest {
	return &NullableDeleteFolderRequest{value: val, isSet: true}
}

func (v NullableDeleteFolderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFolderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

