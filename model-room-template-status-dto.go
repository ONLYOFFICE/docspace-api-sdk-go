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

// checks if the RoomTemplateStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomTemplateStatusDto{}

// RoomTemplateStatusDto The progress of the job that builds a room template out of an existing room.
type RoomTemplateStatusDto struct {
	// The template the job is building. It is meaningful once the job has created the template folder, and the  template can be opened with the room operations only after `isCompleted` turns true.
	TemplateId int32 `json:"templateId"`
	// How far the job has got. The value climbs while the contents of the room are being copied and reaches its  maximum at the very end, so it is an indication of life rather than a reliable estimate of the time left.
	Progress float64 `json:"progress"`
	// Why the job stopped. It is empty while the job runs and after a successful one; when it is filled the  half-built template has already been removed, so nothing has to be cleaned up by the caller.
	Error NullableString `json:"error,omitempty"`
	// Whether the job has ended. It is set both after a successful build and after a failure, so `error` is what  tells the two apart, and the record keeps answering with the same values until another job is started.
	IsCompleted bool `json:"isCompleted"`
}

type _RoomTemplateStatusDto RoomTemplateStatusDto

// NewRoomTemplateStatusDto instantiates a new RoomTemplateStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomTemplateStatusDto(templateId int32, progress float64, isCompleted bool) *RoomTemplateStatusDto {
	this := RoomTemplateStatusDto{}
	this.TemplateId = templateId
	this.Progress = progress
	this.IsCompleted = isCompleted
	return &this
}

// NewRoomTemplateStatusDtoWithDefaults instantiates a new RoomTemplateStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomTemplateStatusDtoWithDefaults() *RoomTemplateStatusDto {
	this := RoomTemplateStatusDto{}
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *RoomTemplateStatusDto) GetTemplateId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *RoomTemplateStatusDto) GetTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *RoomTemplateStatusDto) SetTemplateId(v int32) {
	o.TemplateId = v
}

// GetProgress returns the Progress field value
func (o *RoomTemplateStatusDto) GetProgress() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *RoomTemplateStatusDto) GetProgressOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *RoomTemplateStatusDto) SetProgress(v float64) {
	o.Progress = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomTemplateStatusDto) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomTemplateStatusDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *RoomTemplateStatusDto) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *RoomTemplateStatusDto) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *RoomTemplateStatusDto) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *RoomTemplateStatusDto) UnsetError() {
	o.Error.Unset()
}

// GetIsCompleted returns the IsCompleted field value
func (o *RoomTemplateStatusDto) GetIsCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value
// and a boolean to check if the value has been set.
func (o *RoomTemplateStatusDto) GetIsCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCompleted, true
}

// SetIsCompleted sets field value
func (o *RoomTemplateStatusDto) SetIsCompleted(v bool) {
	o.IsCompleted = v
}

func (o RoomTemplateStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomTemplateStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateId"] = o.TemplateId
	toSerialize["progress"] = o.Progress
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	toSerialize["isCompleted"] = o.IsCompleted
	return toSerialize, nil
}

func (o *RoomTemplateStatusDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateId",
		"progress",
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

	varRoomTemplateStatusDto := _RoomTemplateStatusDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRoomTemplateStatusDto)

	if err != nil {
		return err
	}

	*o = RoomTemplateStatusDto(varRoomTemplateStatusDto)

	return err
}

type NullableRoomTemplateStatusDto struct {
	value *RoomTemplateStatusDto
	isSet bool
}

func (v NullableRoomTemplateStatusDto) Get() *RoomTemplateStatusDto {
	return v.value
}

func (v *NullableRoomTemplateStatusDto) Set(val *RoomTemplateStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomTemplateStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomTemplateStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomTemplateStatusDto(val *RoomTemplateStatusDto) *NullableRoomTemplateStatusDto {
	return &NullableRoomTemplateStatusDto{value: val, isSet: true}
}

func (v NullableRoomTemplateStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomTemplateStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

