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

// checks if the RoomSecurityDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomSecurityDto{}

// RoomSecurityDto The room security parameters.
type RoomSecurityDto struct {
	// The list of room members.
	Members []FileShareDto `json:"members,omitempty"`
	// The warning message.
	Warning NullableString `json:"warning,omitempty"`
	Error *RoomSecurityError `json:"error,omitempty"`
}

// NewRoomSecurityDto instantiates a new RoomSecurityDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomSecurityDto() *RoomSecurityDto {
	this := RoomSecurityDto{}
	return &this
}

// NewRoomSecurityDtoWithDefaults instantiates a new RoomSecurityDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomSecurityDtoWithDefaults() *RoomSecurityDto {
	this := RoomSecurityDto{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomSecurityDto) GetMembers() []FileShareDto {
	if o == nil {
		var ret []FileShareDto
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomSecurityDto) GetMembersOk() ([]FileShareDto, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *RoomSecurityDto) IsMembersSet() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []FileShareDto and assigns it to the Members field.
func (o *RoomSecurityDto) SetMembers(v []FileShareDto) {
	o.Members = v
}

// GetWarning returns the Warning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomSecurityDto) GetWarning() string {
	if o == nil || IsNil(o.Warning.Get()) {
		var ret string
		return ret
	}
	return *o.Warning.Get()
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomSecurityDto) GetWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warning.Get(), o.Warning.IsSet()
}

// HasWarning returns a boolean if a field has been set.
func (o *RoomSecurityDto) IsWarningSet() bool {
	if o != nil && o.Warning.IsSet() {
		return true
	}

	return false
}

// SetWarning gets a reference to the given NullableString and assigns it to the Warning field.
func (o *RoomSecurityDto) SetWarning(v string) {
	o.Warning.Set(&v)
}
// SetWarningNil sets the value for Warning to be an explicit nil
func (o *RoomSecurityDto) SetWarningNil() {
	o.Warning.Set(nil)
}

// UnsetWarning ensures that no value is present for Warning, not even an explicit nil
func (o *RoomSecurityDto) UnsetWarning() {
	o.Warning.Unset()
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RoomSecurityDto) GetError() RoomSecurityError {
	if o == nil || IsNil(o.Error) {
		var ret RoomSecurityError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomSecurityDto) GetErrorOk() (*RoomSecurityError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RoomSecurityDto) IsErrorSet() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given RoomSecurityError and assigns it to the Error field.
func (o *RoomSecurityDto) SetError(v RoomSecurityError) {
	o.Error = &v
}

func (o RoomSecurityDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomSecurityDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Warning.IsSet() {
		toSerialize["warning"] = o.Warning.Get()
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableRoomSecurityDto struct {
	value *RoomSecurityDto
	isSet bool
}

func (v NullableRoomSecurityDto) Get() *RoomSecurityDto {
	return v.value
}

func (v *NullableRoomSecurityDto) Set(val *RoomSecurityDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomSecurityDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomSecurityDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomSecurityDto(val *RoomSecurityDto) *NullableRoomSecurityDto {
	return &NullableRoomSecurityDto{value: val, isSet: true}
}

func (v NullableRoomSecurityDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomSecurityDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

