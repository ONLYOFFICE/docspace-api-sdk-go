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
	"bytes"
	"fmt"
)

// checks if the ChangeOwnerRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeOwnerRequestDto{}

// ChangeOwnerRequestDto The rooms and files to hand over, together with the account that takes them.
type ChangeOwnerRequestDto struct {
	// The rooms to hand over, identified as `GET api/2.0/files/rooms` returns them - a number for a room stored on  the portal and a string for one that lives on a connected third-party account. Only rooms belong here; a  folder inside a room is refused.
	FolderIds []BatchRequestDtoAllOfFileIds `json:"folderIds,omitempty"`
	// The files to hand over, identified as a listing operation returns them - a number for a file stored on the  portal and a string for one on a connected third-party account. Only a file kept in the portal's common  section is accepted.
	FileIds []BatchRequestDtoAllOfFileIds `json:"fileIds,omitempty"`
	// The account that becomes the owner of every listed entry. It has to be an active member allowed to manage  rooms, so a deactivated account, a guest or a plain member is rejected, and for a private room the account  must have set up its encryption keys beforehand.
	UserId string `json:"userId"`
}

type _ChangeOwnerRequestDto ChangeOwnerRequestDto

// NewChangeOwnerRequestDto instantiates a new ChangeOwnerRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeOwnerRequestDto(userId string) *ChangeOwnerRequestDto {
	this := ChangeOwnerRequestDto{}
	this.UserId = userId
	return &this
}

// NewChangeOwnerRequestDtoWithDefaults instantiates a new ChangeOwnerRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeOwnerRequestDtoWithDefaults() *ChangeOwnerRequestDto {
	this := ChangeOwnerRequestDto{}
	return &this
}

// GetFolderIds returns the FolderIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeOwnerRequestDto) GetFolderIds() []BatchRequestDtoAllOfFileIds {
	if o == nil {
		var ret []BatchRequestDtoAllOfFileIds
		return ret
	}
	return o.FolderIds
}

// GetFolderIdsOk returns a tuple with the FolderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeOwnerRequestDto) GetFolderIdsOk() ([]BatchRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.FolderIds) {
		return nil, false
	}
	return o.FolderIds, true
}

// HasFolderIds returns a boolean if a field has been set.
func (o *ChangeOwnerRequestDto) IsFolderIdsSet() bool {
	if o != nil && !IsNil(o.FolderIds) {
		return true
	}

	return false
}

// SetFolderIds gets a reference to the given []BatchRequestDtoAllOfFileIds and assigns it to the FolderIds field.
func (o *ChangeOwnerRequestDto) SetFolderIds(v []BatchRequestDtoAllOfFileIds) {
	o.FolderIds = v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeOwnerRequestDto) GetFileIds() []BatchRequestDtoAllOfFileIds {
	if o == nil {
		var ret []BatchRequestDtoAllOfFileIds
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeOwnerRequestDto) GetFileIdsOk() ([]BatchRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *ChangeOwnerRequestDto) IsFileIdsSet() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []BatchRequestDtoAllOfFileIds and assigns it to the FileIds field.
func (o *ChangeOwnerRequestDto) SetFileIds(v []BatchRequestDtoAllOfFileIds) {
	o.FileIds = v
}

// GetUserId returns the UserId field value
func (o *ChangeOwnerRequestDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ChangeOwnerRequestDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ChangeOwnerRequestDto) SetUserId(v string) {
	o.UserId = v
}

func (o ChangeOwnerRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeOwnerRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FolderIds != nil {
		toSerialize["folderIds"] = o.FolderIds
	}
	if o.FileIds != nil {
		toSerialize["fileIds"] = o.FileIds
	}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *ChangeOwnerRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varChangeOwnerRequestDto := _ChangeOwnerRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varChangeOwnerRequestDto)

	if err != nil {
		return err
	}

	*o = ChangeOwnerRequestDto(varChangeOwnerRequestDto)

	return err
}

type NullableChangeOwnerRequestDto struct {
	value *ChangeOwnerRequestDto
	isSet bool
}

func (v NullableChangeOwnerRequestDto) Get() *ChangeOwnerRequestDto {
	return v.value
}

func (v *NullableChangeOwnerRequestDto) Set(val *ChangeOwnerRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeOwnerRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeOwnerRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeOwnerRequestDto(val *ChangeOwnerRequestDto) *NullableChangeOwnerRequestDto {
	return &NullableChangeOwnerRequestDto{value: val, isSet: true}
}

func (v NullableChangeOwnerRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeOwnerRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

