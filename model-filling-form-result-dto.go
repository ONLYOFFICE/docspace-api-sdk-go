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

// checks if the FillingFormResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FillingFormResultDto{}

// FillingFormResultDto The outcome of one completed form-filling session, as the person who has just filled the form sees it.
type FillingFormResultDto struct {
	// The number this copy was given among the copies made of the same form, counting up from 1. It is the number  the results of the form are ordered by and the one the title of the copy carries.
	FormNumber int32 `json:"formNumber"`
	// The filled copy that the session produced, as an ordinary file: it can be read and downloaded with the file  operations of this API.
	CompletedForm *FileDto `json:"completedForm,omitempty"`
	// The form the copy was made from, so that a client can offer filling it once more.
	OriginalForm *FileDto `json:"originalForm,omitempty"`
	// The account that owns the original form, reported with its email address, so that the person who has just  filled the form knows who receives it and whom to ask about it.
	Manager *EmployeeFullDto `json:"manager,omitempty"`
	// The room the form was filled in. It comes back as 0 when the session was reached through a link shared for  that single form rather than for its room, in which case there is no room the caller could be sent to.
	RoomId int32 `json:"roomId"`
	// Tells whether the calling account may open that room: true for a member of the room and for a portal  administrator, in which case a client can offer going to the room; false for the anonymous caller who filled  the form through a link and can only be shown the copy itself.
	IsRoomMember *bool `json:"isRoomMember,omitempty"`
}

type _FillingFormResultDto FillingFormResultDto

// NewFillingFormResultDto instantiates a new FillingFormResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFillingFormResultDto(formNumber int32, roomId int32) *FillingFormResultDto {
	this := FillingFormResultDto{}
	this.FormNumber = formNumber
	this.RoomId = roomId
	return &this
}

// NewFillingFormResultDtoWithDefaults instantiates a new FillingFormResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFillingFormResultDtoWithDefaults() *FillingFormResultDto {
	this := FillingFormResultDto{}
	return &this
}

// GetFormNumber returns the FormNumber field value
func (o *FillingFormResultDto) GetFormNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FormNumber
}

// GetFormNumberOk returns a tuple with the FormNumber field value
// and a boolean to check if the value has been set.
func (o *FillingFormResultDto) GetFormNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormNumber, true
}

// SetFormNumber sets field value
func (o *FillingFormResultDto) SetFormNumber(v int32) {
	o.FormNumber = v
}

// GetCompletedForm returns the CompletedForm field value if set, zero value otherwise.
func (o *FillingFormResultDto) GetCompletedForm() FileDto {
	if o == nil || IsNil(o.CompletedForm) {
		var ret FileDto
		return ret
	}
	return *o.CompletedForm
}

// GetCompletedFormOk returns a tuple with the CompletedForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillingFormResultDto) GetCompletedFormOk() (*FileDto, bool) {
	if o == nil || IsNil(o.CompletedForm) {
		return nil, false
	}
	return o.CompletedForm, true
}

// HasCompletedForm returns a boolean if a field has been set.
func (o *FillingFormResultDto) IsCompletedFormSet() bool {
	if o != nil && !IsNil(o.CompletedForm) {
		return true
	}

	return false
}

// SetCompletedForm gets a reference to the given FileDto and assigns it to the CompletedForm field.
func (o *FillingFormResultDto) SetCompletedForm(v FileDto) {
	o.CompletedForm = &v
}

// GetOriginalForm returns the OriginalForm field value if set, zero value otherwise.
func (o *FillingFormResultDto) GetOriginalForm() FileDto {
	if o == nil || IsNil(o.OriginalForm) {
		var ret FileDto
		return ret
	}
	return *o.OriginalForm
}

// GetOriginalFormOk returns a tuple with the OriginalForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillingFormResultDto) GetOriginalFormOk() (*FileDto, bool) {
	if o == nil || IsNil(o.OriginalForm) {
		return nil, false
	}
	return o.OriginalForm, true
}

// HasOriginalForm returns a boolean if a field has been set.
func (o *FillingFormResultDto) IsOriginalFormSet() bool {
	if o != nil && !IsNil(o.OriginalForm) {
		return true
	}

	return false
}

// SetOriginalForm gets a reference to the given FileDto and assigns it to the OriginalForm field.
func (o *FillingFormResultDto) SetOriginalForm(v FileDto) {
	o.OriginalForm = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *FillingFormResultDto) GetManager() EmployeeFullDto {
	if o == nil || IsNil(o.Manager) {
		var ret EmployeeFullDto
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillingFormResultDto) GetManagerOk() (*EmployeeFullDto, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *FillingFormResultDto) IsManagerSet() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given EmployeeFullDto and assigns it to the Manager field.
func (o *FillingFormResultDto) SetManager(v EmployeeFullDto) {
	o.Manager = &v
}

// GetRoomId returns the RoomId field value
func (o *FillingFormResultDto) GetRoomId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value
// and a boolean to check if the value has been set.
func (o *FillingFormResultDto) GetRoomIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomId, true
}

// SetRoomId sets field value
func (o *FillingFormResultDto) SetRoomId(v int32) {
	o.RoomId = v
}

// GetIsRoomMember returns the IsRoomMember field value if set, zero value otherwise.
func (o *FillingFormResultDto) GetIsRoomMember() bool {
	if o == nil || IsNil(o.IsRoomMember) {
		var ret bool
		return ret
	}
	return *o.IsRoomMember
}

// GetIsRoomMemberOk returns a tuple with the IsRoomMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillingFormResultDto) GetIsRoomMemberOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoomMember) {
		return nil, false
	}
	return o.IsRoomMember, true
}

// HasIsRoomMember returns a boolean if a field has been set.
func (o *FillingFormResultDto) IsIsRoomMemberSet() bool {
	if o != nil && !IsNil(o.IsRoomMember) {
		return true
	}

	return false
}

// SetIsRoomMember gets a reference to the given bool and assigns it to the IsRoomMember field.
func (o *FillingFormResultDto) SetIsRoomMember(v bool) {
	o.IsRoomMember = &v
}

func (o FillingFormResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FillingFormResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["formNumber"] = o.FormNumber
	if !IsNil(o.CompletedForm) {
		toSerialize["completedForm"] = o.CompletedForm
	}
	if !IsNil(o.OriginalForm) {
		toSerialize["originalForm"] = o.OriginalForm
	}
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	toSerialize["roomId"] = o.RoomId
	if !IsNil(o.IsRoomMember) {
		toSerialize["isRoomMember"] = o.IsRoomMember
	}
	return toSerialize, nil
}

func (o *FillingFormResultDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"formNumber",
		"roomId",
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

	varFillingFormResultDto := _FillingFormResultDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFillingFormResultDto)

	if err != nil {
		return err
	}

	*o = FillingFormResultDto(varFillingFormResultDto)

	return err
}

type NullableFillingFormResultDto struct {
	value *FillingFormResultDto
	isSet bool
}

func (v NullableFillingFormResultDto) Get() *FillingFormResultDto {
	return v.value
}

func (v *NullableFillingFormResultDto) Set(val *FillingFormResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFillingFormResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFillingFormResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFillingFormResultDto(val *FillingFormResultDto) *NullableFillingFormResultDto {
	return &NullableFillingFormResultDto{value: val, isSet: true}
}

func (v NullableFillingFormResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFillingFormResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

