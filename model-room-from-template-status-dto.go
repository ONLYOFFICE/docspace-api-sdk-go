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

// checks if the RoomFromTemplateStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomFromTemplateStatusDto{}

// RoomFromTemplateStatusDto The progress parameters of creating a room from the template.
type RoomFromTemplateStatusDto struct {
	// The room ID.
	RoomId int32 `json:"roomId"`
	// The progress of creating a room from the template.
	Progress float64 `json:"progress"`
	// The error message that is sent when a room is not created successfully from the template.
	Error NullableString `json:"error"`
	// Specifies whether the process of creating a room from the template is completed.
	IsCompleted bool `json:"isCompleted"`
}

type _RoomFromTemplateStatusDto RoomFromTemplateStatusDto

// NewRoomFromTemplateStatusDto instantiates a new RoomFromTemplateStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomFromTemplateStatusDto(roomId int32, progress float64, error_ NullableString, isCompleted bool) *RoomFromTemplateStatusDto {
	this := RoomFromTemplateStatusDto{}
	this.RoomId = roomId
	this.Progress = progress
	this.Error = error_
	this.IsCompleted = isCompleted
	return &this
}

// NewRoomFromTemplateStatusDtoWithDefaults instantiates a new RoomFromTemplateStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomFromTemplateStatusDtoWithDefaults() *RoomFromTemplateStatusDto {
	this := RoomFromTemplateStatusDto{}
	return &this
}

// GetRoomId returns the RoomId field value
func (o *RoomFromTemplateStatusDto) GetRoomId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value
// and a boolean to check if the value has been set.
func (o *RoomFromTemplateStatusDto) GetRoomIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomId, true
}

// SetRoomId sets field value
func (o *RoomFromTemplateStatusDto) SetRoomId(v int32) {
	o.RoomId = v
}

// GetProgress returns the Progress field value
func (o *RoomFromTemplateStatusDto) GetProgress() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *RoomFromTemplateStatusDto) GetProgressOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *RoomFromTemplateStatusDto) SetProgress(v float64) {
	o.Progress = v
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RoomFromTemplateStatusDto) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomFromTemplateStatusDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *RoomFromTemplateStatusDto) SetError(v string) {
	o.Error.Set(&v)
}

// GetIsCompleted returns the IsCompleted field value
func (o *RoomFromTemplateStatusDto) GetIsCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value
// and a boolean to check if the value has been set.
func (o *RoomFromTemplateStatusDto) GetIsCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCompleted, true
}

// SetIsCompleted sets field value
func (o *RoomFromTemplateStatusDto) SetIsCompleted(v bool) {
	o.IsCompleted = v
}

func (o RoomFromTemplateStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomFromTemplateStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roomId"] = o.RoomId
	toSerialize["progress"] = o.Progress
	toSerialize["error"] = o.Error.Get()
	toSerialize["isCompleted"] = o.IsCompleted
	return toSerialize, nil
}

func (o *RoomFromTemplateStatusDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roomId",
		"progress",
		"error",
		"isCompleted",
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

	varRoomFromTemplateStatusDto := _RoomFromTemplateStatusDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRoomFromTemplateStatusDto)

	if err != nil {
		return err
	}

	*o = RoomFromTemplateStatusDto(varRoomFromTemplateStatusDto)

	return err
}

type NullableRoomFromTemplateStatusDto struct {
	value *RoomFromTemplateStatusDto
	isSet bool
}

func (v NullableRoomFromTemplateStatusDto) Get() *RoomFromTemplateStatusDto {
	return v.value
}

func (v *NullableRoomFromTemplateStatusDto) Set(val *RoomFromTemplateStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomFromTemplateStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomFromTemplateStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomFromTemplateStatusDto(val *RoomFromTemplateStatusDto) *NullableRoomFromTemplateStatusDto {
	return &NullableRoomFromTemplateStatusDto{value: val, isSet: true}
}

func (v NullableRoomFromTemplateStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomFromTemplateStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

