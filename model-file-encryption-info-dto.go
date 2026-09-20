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

// checks if the FileEncryptionInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileEncryptionInfoDto{}

// FileEncryptionInfoDto The keys the calling account needs in order to open one file of an end-to-end encrypted private room.
type FileEncryptionInfoDto struct {
	// The key pairs of the calling account, never those of the other people in the room. The private half of each  pair is stored encrypted with that person's own password and has to be decrypted on the client. An empty list  means the account has generated no key pair yet, and until it does no file key can be issued to it.
	UserKeys []EncryptionKeyDto `json:"userKeys,omitempty"`
	// The keys of this file that were issued to the calling account, each naming the public key it was encrypted for  so that the client can pick the matching private half. An empty list means the file has not been shared with  this account rather than that the file is unencrypted.
	FileKeys []FileKeys `json:"fileKeys,omitempty"`
}

// NewFileEncryptionInfoDto instantiates a new FileEncryptionInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileEncryptionInfoDto() *FileEncryptionInfoDto {
	this := FileEncryptionInfoDto{}
	return &this
}

// NewFileEncryptionInfoDtoWithDefaults instantiates a new FileEncryptionInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileEncryptionInfoDtoWithDefaults() *FileEncryptionInfoDto {
	this := FileEncryptionInfoDto{}
	return &this
}

// GetUserKeys returns the UserKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEncryptionInfoDto) GetUserKeys() []EncryptionKeyDto {
	if o == nil {
		var ret []EncryptionKeyDto
		return ret
	}
	return o.UserKeys
}

// GetUserKeysOk returns a tuple with the UserKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEncryptionInfoDto) GetUserKeysOk() ([]EncryptionKeyDto, bool) {
	if o == nil || IsNil(o.UserKeys) {
		return nil, false
	}
	return o.UserKeys, true
}

// HasUserKeys returns a boolean if a field has been set.
func (o *FileEncryptionInfoDto) IsUserKeysSet() bool {
	if o != nil && !IsNil(o.UserKeys) {
		return true
	}

	return false
}

// SetUserKeys gets a reference to the given []EncryptionKeyDto and assigns it to the UserKeys field.
func (o *FileEncryptionInfoDto) SetUserKeys(v []EncryptionKeyDto) {
	o.UserKeys = v
}

// GetFileKeys returns the FileKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEncryptionInfoDto) GetFileKeys() []FileKeys {
	if o == nil {
		var ret []FileKeys
		return ret
	}
	return o.FileKeys
}

// GetFileKeysOk returns a tuple with the FileKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEncryptionInfoDto) GetFileKeysOk() ([]FileKeys, bool) {
	if o == nil || IsNil(o.FileKeys) {
		return nil, false
	}
	return o.FileKeys, true
}

// HasFileKeys returns a boolean if a field has been set.
func (o *FileEncryptionInfoDto) IsFileKeysSet() bool {
	if o != nil && !IsNil(o.FileKeys) {
		return true
	}

	return false
}

// SetFileKeys gets a reference to the given []FileKeys and assigns it to the FileKeys field.
func (o *FileEncryptionInfoDto) SetFileKeys(v []FileKeys) {
	o.FileKeys = v
}

func (o FileEncryptionInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileEncryptionInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserKeys != nil {
		toSerialize["userKeys"] = o.UserKeys
	}
	if o.FileKeys != nil {
		toSerialize["fileKeys"] = o.FileKeys
	}
	return toSerialize, nil
}

type NullableFileEncryptionInfoDto struct {
	value *FileEncryptionInfoDto
	isSet bool
}

func (v NullableFileEncryptionInfoDto) Get() *FileEncryptionInfoDto {
	return v.value
}

func (v *NullableFileEncryptionInfoDto) Set(val *FileEncryptionInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFileEncryptionInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFileEncryptionInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileEncryptionInfoDto(val *FileEncryptionInfoDto) *NullableFileEncryptionInfoDto {
	return &NullableFileEncryptionInfoDto{value: val, isSet: true}
}

func (v NullableFileEncryptionInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileEncryptionInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

