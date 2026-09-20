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

// checks if the StorageEncryptionRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageEncryptionRequestsDto{}

// StorageEncryptionRequestsDto Whether the users are warned before the portals go down for the storage encryption pass.
type StorageEncryptionRequestsDto struct {
	// Whether every user of every portal on the server is mailed before the encryption or decryption pass starts.  The pass runs either way; the flag only decides whether people are told that their portal is about to become  unavailable.
	NotifyUsers *bool `json:"notifyUsers,omitempty"`
}

// NewStorageEncryptionRequestsDto instantiates a new StorageEncryptionRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageEncryptionRequestsDto() *StorageEncryptionRequestsDto {
	this := StorageEncryptionRequestsDto{}
	return &this
}

// NewStorageEncryptionRequestsDtoWithDefaults instantiates a new StorageEncryptionRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageEncryptionRequestsDtoWithDefaults() *StorageEncryptionRequestsDto {
	this := StorageEncryptionRequestsDto{}
	return &this
}

// GetNotifyUsers returns the NotifyUsers field value if set, zero value otherwise.
func (o *StorageEncryptionRequestsDto) GetNotifyUsers() bool {
	if o == nil || IsNil(o.NotifyUsers) {
		var ret bool
		return ret
	}
	return *o.NotifyUsers
}

// GetNotifyUsersOk returns a tuple with the NotifyUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEncryptionRequestsDto) GetNotifyUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyUsers) {
		return nil, false
	}
	return o.NotifyUsers, true
}

// HasNotifyUsers returns a boolean if a field has been set.
func (o *StorageEncryptionRequestsDto) IsNotifyUsersSet() bool {
	if o != nil && !IsNil(o.NotifyUsers) {
		return true
	}

	return false
}

// SetNotifyUsers gets a reference to the given bool and assigns it to the NotifyUsers field.
func (o *StorageEncryptionRequestsDto) SetNotifyUsers(v bool) {
	o.NotifyUsers = &v
}

func (o StorageEncryptionRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageEncryptionRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotifyUsers) {
		toSerialize["notifyUsers"] = o.NotifyUsers
	}
	return toSerialize, nil
}

type NullableStorageEncryptionRequestsDto struct {
	value *StorageEncryptionRequestsDto
	isSet bool
}

func (v NullableStorageEncryptionRequestsDto) Get() *StorageEncryptionRequestsDto {
	return v.value
}

func (v *NullableStorageEncryptionRequestsDto) Set(val *StorageEncryptionRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageEncryptionRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageEncryptionRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageEncryptionRequestsDto(val *StorageEncryptionRequestsDto) *NullableStorageEncryptionRequestsDto {
	return &NullableStorageEncryptionRequestsDto{value: val, isSet: true}
}

func (v NullableStorageEncryptionRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageEncryptionRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

