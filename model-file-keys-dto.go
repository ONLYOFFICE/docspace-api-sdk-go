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
	"time"
)

// checks if the FileKeysDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileKeysDto{}

// FileKeysDto The encrypted file key issued to one user.
type FileKeysDto struct {
	// The identifier of the user the file key was issued to.
	UserId *string `json:"userId,omitempty"`
	// The identifier of the key pair the file key is encrypted for.
	PublicKeyId *string `json:"publicKeyId,omitempty"`
	// The file key, encrypted with the public key of the pair.
	PrivateKeyEnc NullableString `json:"privateKeyEnc,omitempty"`
	// The identifier of the portal the file belongs to.
	TenantId *int32 `json:"tenantId,omitempty"`
	// The identifier of the file the key unlocks.
	FileId *int32 `json:"fileId,omitempty"`
	// The date and time when the file key was issued.
	CreateOn *time.Time `json:"createOn,omitempty"`
}

// NewFileKeysDto instantiates a new FileKeysDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileKeysDto() *FileKeysDto {
	this := FileKeysDto{}
	return &this
}

// NewFileKeysDtoWithDefaults instantiates a new FileKeysDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileKeysDtoWithDefaults() *FileKeysDto {
	this := FileKeysDto{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FileKeysDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileKeysDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FileKeysDto) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FileKeysDto) SetUserId(v string) {
	o.UserId = &v
}

// GetPublicKeyId returns the PublicKeyId field value if set, zero value otherwise.
func (o *FileKeysDto) GetPublicKeyId() string {
	if o == nil || IsNil(o.PublicKeyId) {
		var ret string
		return ret
	}
	return *o.PublicKeyId
}

// GetPublicKeyIdOk returns a tuple with the PublicKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileKeysDto) GetPublicKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKeyId) {
		return nil, false
	}
	return o.PublicKeyId, true
}

// HasPublicKeyId returns a boolean if a field has been set.
func (o *FileKeysDto) IsPublicKeyIdSet() bool {
	if o != nil && !IsNil(o.PublicKeyId) {
		return true
	}

	return false
}

// SetPublicKeyId gets a reference to the given string and assigns it to the PublicKeyId field.
func (o *FileKeysDto) SetPublicKeyId(v string) {
	o.PublicKeyId = &v
}

// GetPrivateKeyEnc returns the PrivateKeyEnc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileKeysDto) GetPrivateKeyEnc() string {
	if o == nil || IsNil(o.PrivateKeyEnc.Get()) {
		var ret string
		return ret
	}
	return *o.PrivateKeyEnc.Get()
}

// GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileKeysDto) GetPrivateKeyEncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKeyEnc.Get(), o.PrivateKeyEnc.IsSet()
}

// HasPrivateKeyEnc returns a boolean if a field has been set.
func (o *FileKeysDto) IsPrivateKeyEncSet() bool {
	if o != nil && o.PrivateKeyEnc.IsSet() {
		return true
	}

	return false
}

// SetPrivateKeyEnc gets a reference to the given NullableString and assigns it to the PrivateKeyEnc field.
func (o *FileKeysDto) SetPrivateKeyEnc(v string) {
	o.PrivateKeyEnc.Set(&v)
}
// SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil
func (o *FileKeysDto) SetPrivateKeyEncNil() {
	o.PrivateKeyEnc.Set(nil)
}

// UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil
func (o *FileKeysDto) UnsetPrivateKeyEnc() {
	o.PrivateKeyEnc.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *FileKeysDto) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileKeysDto) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *FileKeysDto) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *FileKeysDto) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *FileKeysDto) GetFileId() int32 {
	if o == nil || IsNil(o.FileId) {
		var ret int32
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileKeysDto) GetFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *FileKeysDto) IsFileIdSet() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given int32 and assigns it to the FileId field.
func (o *FileKeysDto) SetFileId(v int32) {
	o.FileId = &v
}

// GetCreateOn returns the CreateOn field value if set, zero value otherwise.
func (o *FileKeysDto) GetCreateOn() time.Time {
	if o == nil || IsNil(o.CreateOn) {
		var ret time.Time
		return ret
	}
	return *o.CreateOn
}

// GetCreateOnOk returns a tuple with the CreateOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileKeysDto) GetCreateOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateOn) {
		return nil, false
	}
	return o.CreateOn, true
}

// HasCreateOn returns a boolean if a field has been set.
func (o *FileKeysDto) IsCreateOnSet() bool {
	if o != nil && !IsNil(o.CreateOn) {
		return true
	}

	return false
}

// SetCreateOn gets a reference to the given time.Time and assigns it to the CreateOn field.
func (o *FileKeysDto) SetCreateOn(v time.Time) {
	o.CreateOn = &v
}

func (o FileKeysDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileKeysDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.PublicKeyId) {
		toSerialize["publicKeyId"] = o.PublicKeyId
	}
	if o.PrivateKeyEnc.IsSet() {
		toSerialize["privateKeyEnc"] = o.PrivateKeyEnc.Get()
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.FileId) {
		toSerialize["fileId"] = o.FileId
	}
	if !IsNil(o.CreateOn) {
		toSerialize["createOn"] = o.CreateOn
	}
	return toSerialize, nil
}

type NullableFileKeysDto struct {
	value *FileKeysDto
	isSet bool
}

func (v NullableFileKeysDto) Get() *FileKeysDto {
	return v.value
}

func (v *NullableFileKeysDto) Set(val *FileKeysDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFileKeysDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFileKeysDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileKeysDto(val *FileKeysDto) *NullableFileKeysDto {
	return &NullableFileKeysDto{value: val, isSet: true}
}

func (v NullableFileKeysDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileKeysDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

