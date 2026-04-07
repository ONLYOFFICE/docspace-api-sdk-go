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

// checks if the FileShareParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileShareParams{}

// FileShareParams The collection of file sharing parameters.
type FileShareParams struct {
	// The email address.
	Email NullableString `json:"email,omitempty"`
	// The ID of the user to whom the file will be shared.
	ShareTo *string `json:"shareTo,omitempty"`
	Access *FileShare `json:"access,omitempty"`
}

// NewFileShareParams instantiates a new FileShareParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileShareParams() *FileShareParams {
	this := FileShareParams{}
	return &this
}

// NewFileShareParamsWithDefaults instantiates a new FileShareParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileShareParamsWithDefaults() *FileShareParams {
	this := FileShareParams{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareParams) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareParams) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *FileShareParams) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *FileShareParams) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *FileShareParams) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *FileShareParams) UnsetEmail() {
	o.Email.Unset()
}

// GetShareTo returns the ShareTo field value if set, zero value otherwise.
func (o *FileShareParams) GetShareTo() string {
	if o == nil || IsNil(o.ShareTo) {
		var ret string
		return ret
	}
	return *o.ShareTo
}

// GetShareToOk returns a tuple with the ShareTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareParams) GetShareToOk() (*string, bool) {
	if o == nil || IsNil(o.ShareTo) {
		return nil, false
	}
	return o.ShareTo, true
}

// HasShareTo returns a boolean if a field has been set.
func (o *FileShareParams) IsShareToSet() bool {
	if o != nil && !IsNil(o.ShareTo) {
		return true
	}

	return false
}

// SetShareTo gets a reference to the given string and assigns it to the ShareTo field.
func (o *FileShareParams) SetShareTo(v string) {
	o.ShareTo = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *FileShareParams) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareParams) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *FileShareParams) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *FileShareParams) SetAccess(v FileShare) {
	o.Access = &v
}

func (o FileShareParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileShareParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.ShareTo) {
		toSerialize["shareTo"] = o.ShareTo
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

type NullableFileShareParams struct {
	value *FileShareParams
	isSet bool
}

func (v NullableFileShareParams) Get() *FileShareParams {
	return v.value
}

func (v *NullableFileShareParams) Set(val *FileShareParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFileShareParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFileShareParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileShareParams(val *FileShareParams) *NullableFileShareParams {
	return &NullableFileShareParams{value: val, isSet: true}
}

func (v NullableFileShareParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileShareParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

