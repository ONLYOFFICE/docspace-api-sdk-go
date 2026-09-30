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

// checks if the UpdatePhotoMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePhotoMemberRequest{}

// UpdatePhotoMemberRequest The request parameters for updating a photo.
type UpdatePhotoMemberRequest struct {
	// The address the portal downloads the new avatar from. It has to be absolute or relative to the portal, and it  has to use HTTPS unless the request itself came over HTTP; an address the portal refuses to fetch is rejected.  It is required - an empty value is answered with 400 rather than clearing the avatar.
	Files NullableString `json:"files,omitempty"`
}

// NewUpdatePhotoMemberRequest instantiates a new UpdatePhotoMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePhotoMemberRequest() *UpdatePhotoMemberRequest {
	this := UpdatePhotoMemberRequest{}
	return &this
}

// NewUpdatePhotoMemberRequestWithDefaults instantiates a new UpdatePhotoMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePhotoMemberRequestWithDefaults() *UpdatePhotoMemberRequest {
	this := UpdatePhotoMemberRequest{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePhotoMemberRequest) GetFiles() string {
	if o == nil || IsNil(o.Files.Get()) {
		var ret string
		return ret
	}
	return *o.Files.Get()
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePhotoMemberRequest) GetFilesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files.Get(), o.Files.IsSet()
}

// HasFiles returns a boolean if a field has been set.
func (o *UpdatePhotoMemberRequest) IsFilesSet() bool {
	if o != nil && o.Files.IsSet() {
		return true
	}

	return false
}

// SetFiles gets a reference to the given NullableString and assigns it to the Files field.
func (o *UpdatePhotoMemberRequest) SetFiles(v string) {
	o.Files.Set(&v)
}
// SetFilesNil sets the value for Files to be an explicit nil
func (o *UpdatePhotoMemberRequest) SetFilesNil() {
	o.Files.Set(nil)
}

// UnsetFiles ensures that no value is present for Files, not even an explicit nil
func (o *UpdatePhotoMemberRequest) UnsetFiles() {
	o.Files.Unset()
}

func (o UpdatePhotoMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePhotoMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Files.IsSet() {
		toSerialize["files"] = o.Files.Get()
	}
	return toSerialize, nil
}

type NullableUpdatePhotoMemberRequest struct {
	value *UpdatePhotoMemberRequest
	isSet bool
}

func (v NullableUpdatePhotoMemberRequest) Get() *UpdatePhotoMemberRequest {
	return v.value
}

func (v *NullableUpdatePhotoMemberRequest) Set(val *UpdatePhotoMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePhotoMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePhotoMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePhotoMemberRequest(val *UpdatePhotoMemberRequest) *NullableUpdatePhotoMemberRequest {
	return &NullableUpdatePhotoMemberRequest{value: val, isSet: true}
}

func (v NullableUpdatePhotoMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePhotoMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

