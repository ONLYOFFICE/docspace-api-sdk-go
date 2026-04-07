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

// checks if the FillingFormResultDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FillingFormResultDtoInteger{}

// FillingFormResultDtoInteger The parameters of the form filling result.
type FillingFormResultDtoInteger struct {
	// The filling form number.
	FormNumber int32 `json:"formNumber"`
	CompletedForm *FileDtoInteger `json:"completedForm,omitempty"`
	OriginalForm *FileDtoInteger `json:"originalForm,omitempty"`
	Manager *EmployeeFullDto `json:"manager,omitempty"`
	// The room ID where filling the form.
	RoomId int32 `json:"roomId"`
	// Specifies if the manager who fills the form is a room member or not.
	IsRoomMember *bool `json:"isRoomMember,omitempty"`
}

type _FillingFormResultDtoInteger FillingFormResultDtoInteger

// NewFillingFormResultDtoInteger instantiates a new FillingFormResultDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFillingFormResultDtoInteger(formNumber int32, roomId int32) *FillingFormResultDtoInteger {
	this := FillingFormResultDtoInteger{}
	this.FormNumber = formNumber
	this.RoomId = roomId
	return &this
}

// NewFillingFormResultDtoIntegerWithDefaults instantiates a new FillingFormResultDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFillingFormResultDtoIntegerWithDefaults() *FillingFormResultDtoInteger {
	this := FillingFormResultDtoInteger{}
	return &this
}

// GetFormNumber returns the FormNumber field value
func (o *FillingFormResultDtoInteger) GetFormNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FormNumber
}

// GetFormNumberOk returns a tuple with the FormNumber field value
// and a boolean to check if the value has been set.
func (o *FillingFormResultDtoInteger) GetFormNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormNumber, true
}

// SetFormNumber sets field value
func (o *FillingFormResultDtoInteger) SetFormNumber(v int32) {
	o.FormNumber = v
}

// GetCompletedForm returns the CompletedForm field value if set, zero value otherwise.
func (o *FillingFormResultDtoInteger) GetCompletedForm() FileDtoInteger {
	if o == nil || IsNil(o.CompletedForm) {
		var ret FileDtoInteger
		return ret
	}
	return *o.CompletedForm
}

// GetCompletedFormOk returns a tuple with the CompletedForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillingFormResultDtoInteger) GetCompletedFormOk() (*FileDtoInteger, bool) {
	if o == nil || IsNil(o.CompletedForm) {
		return nil, false
	}
	return o.CompletedForm, true
}

// HasCompletedForm returns a boolean if a field has been set.
func (o *FillingFormResultDtoInteger) IsCompletedFormSet() bool {
	if o != nil && !IsNil(o.CompletedForm) {
		return true
	}

	return false
}

// SetCompletedForm gets a reference to the given FileDtoInteger and assigns it to the CompletedForm field.
func (o *FillingFormResultDtoInteger) SetCompletedForm(v FileDtoInteger) {
	o.CompletedForm = &v
}

// GetOriginalForm returns the OriginalForm field value if set, zero value otherwise.
func (o *FillingFormResultDtoInteger) GetOriginalForm() FileDtoInteger {
	if o == nil || IsNil(o.OriginalForm) {
		var ret FileDtoInteger
		return ret
	}
	return *o.OriginalForm
}

// GetOriginalFormOk returns a tuple with the OriginalForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillingFormResultDtoInteger) GetOriginalFormOk() (*FileDtoInteger, bool) {
	if o == nil || IsNil(o.OriginalForm) {
		return nil, false
	}
	return o.OriginalForm, true
}

// HasOriginalForm returns a boolean if a field has been set.
func (o *FillingFormResultDtoInteger) IsOriginalFormSet() bool {
	if o != nil && !IsNil(o.OriginalForm) {
		return true
	}

	return false
}

// SetOriginalForm gets a reference to the given FileDtoInteger and assigns it to the OriginalForm field.
func (o *FillingFormResultDtoInteger) SetOriginalForm(v FileDtoInteger) {
	o.OriginalForm = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *FillingFormResultDtoInteger) GetManager() EmployeeFullDto {
	if o == nil || IsNil(o.Manager) {
		var ret EmployeeFullDto
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillingFormResultDtoInteger) GetManagerOk() (*EmployeeFullDto, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *FillingFormResultDtoInteger) IsManagerSet() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given EmployeeFullDto and assigns it to the Manager field.
func (o *FillingFormResultDtoInteger) SetManager(v EmployeeFullDto) {
	o.Manager = &v
}

// GetRoomId returns the RoomId field value
func (o *FillingFormResultDtoInteger) GetRoomId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value
// and a boolean to check if the value has been set.
func (o *FillingFormResultDtoInteger) GetRoomIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomId, true
}

// SetRoomId sets field value
func (o *FillingFormResultDtoInteger) SetRoomId(v int32) {
	o.RoomId = v
}

// GetIsRoomMember returns the IsRoomMember field value if set, zero value otherwise.
func (o *FillingFormResultDtoInteger) GetIsRoomMember() bool {
	if o == nil || IsNil(o.IsRoomMember) {
		var ret bool
		return ret
	}
	return *o.IsRoomMember
}

// GetIsRoomMemberOk returns a tuple with the IsRoomMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillingFormResultDtoInteger) GetIsRoomMemberOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoomMember) {
		return nil, false
	}
	return o.IsRoomMember, true
}

// HasIsRoomMember returns a boolean if a field has been set.
func (o *FillingFormResultDtoInteger) IsIsRoomMemberSet() bool {
	if o != nil && !IsNil(o.IsRoomMember) {
		return true
	}

	return false
}

// SetIsRoomMember gets a reference to the given bool and assigns it to the IsRoomMember field.
func (o *FillingFormResultDtoInteger) SetIsRoomMember(v bool) {
	o.IsRoomMember = &v
}

func (o FillingFormResultDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FillingFormResultDtoInteger) ToMap() (map[string]interface{}, error) {
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

func (o *FillingFormResultDtoInteger) UnmarshalJSON(data []byte) (err error) {
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

	varFillingFormResultDtoInteger := _FillingFormResultDtoInteger{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFillingFormResultDtoInteger)

	if err != nil {
		return err
	}

	*o = FillingFormResultDtoInteger(varFillingFormResultDtoInteger)

	return err
}

type NullableFillingFormResultDtoInteger struct {
	value *FillingFormResultDtoInteger
	isSet bool
}

func (v NullableFillingFormResultDtoInteger) Get() *FillingFormResultDtoInteger {
	return v.value
}

func (v *NullableFillingFormResultDtoInteger) Set(val *FillingFormResultDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableFillingFormResultDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableFillingFormResultDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFillingFormResultDtoInteger(val *FillingFormResultDtoInteger) *NullableFillingFormResultDtoInteger {
	return &NullableFillingFormResultDtoInteger{value: val, isSet: true}
}

func (v NullableFillingFormResultDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFillingFormResultDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

