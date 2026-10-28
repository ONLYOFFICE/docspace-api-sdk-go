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

// checks if the UpdateFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFileRequest{}

// UpdateFileRequest The changes to make to a file: a new title, an earlier version to restore, or both.
type UpdateFileRequest struct {
	// The new title of the file, without an extension - the stored extension is kept whatever the title says, so a  rename cannot change the format. Left empty, the file keeps its name.
	Title NullableString `json:"title,omitempty"`
	// The version to restore on top of the history, as reported by `GET api/2.0/files/file/{fileId}/history`; 0 or  less leaves the versions untouched.
	LastVersion *int32 `json:"lastVersion,omitempty"`
}

// NewUpdateFileRequest instantiates a new UpdateFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFileRequest() *UpdateFileRequest {
	this := UpdateFileRequest{}
	return &this
}

// NewUpdateFileRequestWithDefaults instantiates a new UpdateFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFileRequestWithDefaults() *UpdateFileRequest {
	this := UpdateFileRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFileRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFileRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateFileRequest) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UpdateFileRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UpdateFileRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UpdateFileRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetLastVersion returns the LastVersion field value if set, zero value otherwise.
func (o *UpdateFileRequest) GetLastVersion() int32 {
	if o == nil || IsNil(o.LastVersion) {
		var ret int32
		return ret
	}
	return *o.LastVersion
}

// GetLastVersionOk returns a tuple with the LastVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileRequest) GetLastVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.LastVersion) {
		return nil, false
	}
	return o.LastVersion, true
}

// HasLastVersion returns a boolean if a field has been set.
func (o *UpdateFileRequest) IsLastVersionSet() bool {
	if o != nil && !IsNil(o.LastVersion) {
		return true
	}

	return false
}

// SetLastVersion gets a reference to the given int32 and assigns it to the LastVersion field.
func (o *UpdateFileRequest) SetLastVersion(v int32) {
	o.LastVersion = &v
}

func (o UpdateFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.LastVersion) {
		toSerialize["lastVersion"] = o.LastVersion
	}
	return toSerialize, nil
}

type NullableUpdateFileRequest struct {
	value *UpdateFileRequest
	isSet bool
}

func (v NullableUpdateFileRequest) Get() *UpdateFileRequest {
	return v.value
}

func (v *NullableUpdateFileRequest) Set(val *UpdateFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFileRequest(val *UpdateFileRequest) *NullableUpdateFileRequest {
	return &NullableUpdateFileRequest{value: val, isSet: true}
}

func (v NullableUpdateFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

