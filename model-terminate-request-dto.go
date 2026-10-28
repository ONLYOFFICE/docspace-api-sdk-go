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

// checks if the TerminateRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminateRequestDto{}

// TerminateRequestDto The request parameters that address the queued job of a single user - a data reassignment, a data deletion or a  user type change.
type TerminateRequestDto struct {
	// The ID of the user whose job is addressed. For a terminate operation it has to be the same ID that was passed  when the job was started.
	UserId string `json:"userId"`
}

type _TerminateRequestDto TerminateRequestDto

// NewTerminateRequestDto instantiates a new TerminateRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminateRequestDto(userId string) *TerminateRequestDto {
	this := TerminateRequestDto{}
	this.UserId = userId
	return &this
}

// NewTerminateRequestDtoWithDefaults instantiates a new TerminateRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminateRequestDtoWithDefaults() *TerminateRequestDto {
	this := TerminateRequestDto{}
	return &this
}

// GetUserId returns the UserId field value
func (o *TerminateRequestDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TerminateRequestDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TerminateRequestDto) SetUserId(v string) {
	o.UserId = v
}

func (o TerminateRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminateRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *TerminateRequestDto) UnmarshalJSON(data []byte) (err error) {
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

	varTerminateRequestDto := _TerminateRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTerminateRequestDto)

	if err != nil {
		return err
	}

	*o = TerminateRequestDto(varTerminateRequestDto)

	return err
}

type NullableTerminateRequestDto struct {
	value *TerminateRequestDto
	isSet bool
}

func (v NullableTerminateRequestDto) Get() *TerminateRequestDto {
	return v.value
}

func (v *NullableTerminateRequestDto) Set(val *TerminateRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminateRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminateRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminateRequestDto(val *TerminateRequestDto) *NullableTerminateRequestDto {
	return &NullableTerminateRequestDto{value: val, isSet: true}
}

func (v NullableTerminateRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminateRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

