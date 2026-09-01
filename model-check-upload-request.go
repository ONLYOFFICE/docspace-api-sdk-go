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

// checks if the CheckUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckUploadRequest{}

// CheckUploadRequest The request parameters for checking file uploads.
type CheckUploadRequest struct {
	// The list of file titles.
	FilesTitle []string `json:"filesTitle,omitempty"`
}

// NewCheckUploadRequest instantiates a new CheckUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckUploadRequest() *CheckUploadRequest {
	this := CheckUploadRequest{}
	return &this
}

// NewCheckUploadRequestWithDefaults instantiates a new CheckUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckUploadRequestWithDefaults() *CheckUploadRequest {
	this := CheckUploadRequest{}
	return &this
}

// GetFilesTitle returns the FilesTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckUploadRequest) GetFilesTitle() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FilesTitle
}

// GetFilesTitleOk returns a tuple with the FilesTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckUploadRequest) GetFilesTitleOk() ([]string, bool) {
	if o == nil || IsNil(o.FilesTitle) {
		return nil, false
	}
	return o.FilesTitle, true
}

// HasFilesTitle returns a boolean if a field has been set.
func (o *CheckUploadRequest) IsFilesTitleSet() bool {
	if o != nil && !IsNil(o.FilesTitle) {
		return true
	}

	return false
}

// SetFilesTitle gets a reference to the given []string and assigns it to the FilesTitle field.
func (o *CheckUploadRequest) SetFilesTitle(v []string) {
	o.FilesTitle = v
}

func (o CheckUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FilesTitle != nil {
		toSerialize["filesTitle"] = o.FilesTitle
	}
	return toSerialize, nil
}

type NullableCheckUploadRequest struct {
	value *CheckUploadRequest
	isSet bool
}

func (v NullableCheckUploadRequest) Get() *CheckUploadRequest {
	return v.value
}

func (v *NullableCheckUploadRequest) Set(val *CheckUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckUploadRequest(val *CheckUploadRequest) *NullableCheckUploadRequest {
	return &NullableCheckUploadRequest{value: val, isSet: true}
}

func (v NullableCheckUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

