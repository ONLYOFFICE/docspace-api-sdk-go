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

// checks if the DeleteFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteFileRequest{}

// DeleteFileRequest The parameters of a single file deletion.
type DeleteFileRequest struct {
	// When to delete: `true` waits until the editing session on the file has ended, `false` deletes at once, pulling  the file away from whoever is working on it.
	DeleteAfter *bool `json:"deleteAfter,omitempty"`
	// Where the file goes: `false` moves it to Trash, from where it can be restored, `true` deletes it for good.  Inside a room, where there is no Trash, deletion is always final.
	Immediately *bool `json:"immediately,omitempty"`
}

// NewDeleteFileRequest instantiates a new DeleteFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFileRequest() *DeleteFileRequest {
	this := DeleteFileRequest{}
	return &this
}

// NewDeleteFileRequestWithDefaults instantiates a new DeleteFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFileRequestWithDefaults() *DeleteFileRequest {
	this := DeleteFileRequest{}
	return &this
}

// GetDeleteAfter returns the DeleteAfter field value if set, zero value otherwise.
func (o *DeleteFileRequest) GetDeleteAfter() bool {
	if o == nil || IsNil(o.DeleteAfter) {
		var ret bool
		return ret
	}
	return *o.DeleteAfter
}

// GetDeleteAfterOk returns a tuple with the DeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFileRequest) GetDeleteAfterOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteAfter) {
		return nil, false
	}
	return o.DeleteAfter, true
}

// HasDeleteAfter returns a boolean if a field has been set.
func (o *DeleteFileRequest) IsDeleteAfterSet() bool {
	if o != nil && !IsNil(o.DeleteAfter) {
		return true
	}

	return false
}

// SetDeleteAfter gets a reference to the given bool and assigns it to the DeleteAfter field.
func (o *DeleteFileRequest) SetDeleteAfter(v bool) {
	o.DeleteAfter = &v
}

// GetImmediately returns the Immediately field value if set, zero value otherwise.
func (o *DeleteFileRequest) GetImmediately() bool {
	if o == nil || IsNil(o.Immediately) {
		var ret bool
		return ret
	}
	return *o.Immediately
}

// GetImmediatelyOk returns a tuple with the Immediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFileRequest) GetImmediatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.Immediately) {
		return nil, false
	}
	return o.Immediately, true
}

// HasImmediately returns a boolean if a field has been set.
func (o *DeleteFileRequest) IsImmediatelySet() bool {
	if o != nil && !IsNil(o.Immediately) {
		return true
	}

	return false
}

// SetImmediately gets a reference to the given bool and assigns it to the Immediately field.
func (o *DeleteFileRequest) SetImmediately(v bool) {
	o.Immediately = &v
}

func (o DeleteFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteAfter) {
		toSerialize["deleteAfter"] = o.DeleteAfter
	}
	if !IsNil(o.Immediately) {
		toSerialize["immediately"] = o.Immediately
	}
	return toSerialize, nil
}

type NullableDeleteFileRequest struct {
	value *DeleteFileRequest
	isSet bool
}

func (v NullableDeleteFileRequest) Get() *DeleteFileRequest {
	return v.value
}

func (v *NullableDeleteFileRequest) Set(val *DeleteFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFileRequest(val *DeleteFileRequest) *NullableDeleteFileRequest {
	return &NullableDeleteFileRequest{value: val, isSet: true}
}

func (v NullableDeleteFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

