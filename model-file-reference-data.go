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

// checks if the FileReferenceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileReferenceData{}

// FileReferenceData The pair of values that names a document across portals, as it is written into a spreadsheet formula.
type FileReferenceData struct {
	// The id of the document inside the portal named below.
	FileKey NullableString `json:"fileKey,omitempty"`
	// The portal the document lives in. A reference whose value is not this portal cannot be resolved by the file  key and falls back to the path or the link.
	InstanceId NullableString `json:"instanceId,omitempty"`
	// The room the document lies in. It is filled in only for a document opened in a virtual data room, and stays  empty everywhere else.
	RoomId NullableString `json:"roomId,omitempty"`
	// Whether the caller may manage the room named above; it is only meaningful together with it.
	CanEditRoom *bool `json:"canEditRoom,omitempty"`
}

// NewFileReferenceData instantiates a new FileReferenceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileReferenceData() *FileReferenceData {
	this := FileReferenceData{}
	return &this
}

// NewFileReferenceDataWithDefaults instantiates a new FileReferenceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileReferenceDataWithDefaults() *FileReferenceData {
	this := FileReferenceData{}
	return &this
}

// GetFileKey returns the FileKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReferenceData) GetFileKey() string {
	if o == nil || IsNil(o.FileKey.Get()) {
		var ret string
		return ret
	}
	return *o.FileKey.Get()
}

// GetFileKeyOk returns a tuple with the FileKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReferenceData) GetFileKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileKey.Get(), o.FileKey.IsSet()
}

// HasFileKey returns a boolean if a field has been set.
func (o *FileReferenceData) IsFileKeySet() bool {
	if o != nil && o.FileKey.IsSet() {
		return true
	}

	return false
}

// SetFileKey gets a reference to the given NullableString and assigns it to the FileKey field.
func (o *FileReferenceData) SetFileKey(v string) {
	o.FileKey.Set(&v)
}
// SetFileKeyNil sets the value for FileKey to be an explicit nil
func (o *FileReferenceData) SetFileKeyNil() {
	o.FileKey.Set(nil)
}

// UnsetFileKey ensures that no value is present for FileKey, not even an explicit nil
func (o *FileReferenceData) UnsetFileKey() {
	o.FileKey.Unset()
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReferenceData) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReferenceData) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// HasInstanceId returns a boolean if a field has been set.
func (o *FileReferenceData) IsInstanceIdSet() bool {
	if o != nil && o.InstanceId.IsSet() {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given NullableString and assigns it to the InstanceId field.
func (o *FileReferenceData) SetInstanceId(v string) {
	o.InstanceId.Set(&v)
}
// SetInstanceIdNil sets the value for InstanceId to be an explicit nil
func (o *FileReferenceData) SetInstanceIdNil() {
	o.InstanceId.Set(nil)
}

// UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
func (o *FileReferenceData) UnsetInstanceId() {
	o.InstanceId.Unset()
}

// GetRoomId returns the RoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReferenceData) GetRoomId() string {
	if o == nil || IsNil(o.RoomId.Get()) {
		var ret string
		return ret
	}
	return *o.RoomId.Get()
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReferenceData) GetRoomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomId.Get(), o.RoomId.IsSet()
}

// HasRoomId returns a boolean if a field has been set.
func (o *FileReferenceData) IsRoomIdSet() bool {
	if o != nil && o.RoomId.IsSet() {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given NullableString and assigns it to the RoomId field.
func (o *FileReferenceData) SetRoomId(v string) {
	o.RoomId.Set(&v)
}
// SetRoomIdNil sets the value for RoomId to be an explicit nil
func (o *FileReferenceData) SetRoomIdNil() {
	o.RoomId.Set(nil)
}

// UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
func (o *FileReferenceData) UnsetRoomId() {
	o.RoomId.Unset()
}

// GetCanEditRoom returns the CanEditRoom field value if set, zero value otherwise.
func (o *FileReferenceData) GetCanEditRoom() bool {
	if o == nil || IsNil(o.CanEditRoom) {
		var ret bool
		return ret
	}
	return *o.CanEditRoom
}

// GetCanEditRoomOk returns a tuple with the CanEditRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileReferenceData) GetCanEditRoomOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEditRoom) {
		return nil, false
	}
	return o.CanEditRoom, true
}

// HasCanEditRoom returns a boolean if a field has been set.
func (o *FileReferenceData) IsCanEditRoomSet() bool {
	if o != nil && !IsNil(o.CanEditRoom) {
		return true
	}

	return false
}

// SetCanEditRoom gets a reference to the given bool and assigns it to the CanEditRoom field.
func (o *FileReferenceData) SetCanEditRoom(v bool) {
	o.CanEditRoom = &v
}

func (o FileReferenceData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileReferenceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FileKey.IsSet() {
		toSerialize["fileKey"] = o.FileKey.Get()
	}
	if o.InstanceId.IsSet() {
		toSerialize["instanceId"] = o.InstanceId.Get()
	}
	if o.RoomId.IsSet() {
		toSerialize["roomId"] = o.RoomId.Get()
	}
	if !IsNil(o.CanEditRoom) {
		toSerialize["canEditRoom"] = o.CanEditRoom
	}
	return toSerialize, nil
}

type NullableFileReferenceData struct {
	value *FileReferenceData
	isSet bool
}

func (v NullableFileReferenceData) Get() *FileReferenceData {
	return v.value
}

func (v *NullableFileReferenceData) Set(val *FileReferenceData) {
	v.value = val
	v.isSet = true
}

func (v NullableFileReferenceData) IsSet() bool {
	return v.isSet
}

func (v *NullableFileReferenceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileReferenceData(val *FileReferenceData) *NullableFileReferenceData {
	return &NullableFileReferenceData{value: val, isSet: true}
}

func (v NullableFileReferenceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileReferenceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

