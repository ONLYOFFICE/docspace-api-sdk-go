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

// checks if the StartReassignRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartReassignRequestDto{}

// StartReassignRequestDto The request parameters for starting the reassignment process.
type StartReassignRequestDto struct {
	// The user ID whose data will be reassigned to another user.
	FromUserId string `json:"fromUserId"`
	// The user ID to whom all the data will be reassigned.
	ToUserId string `json:"toUserId"`
	// Specifies whether to delete a profile when the data reassignment will be finished or not.
	DeleteProfile *bool `json:"deleteProfile,omitempty"`
}

type _StartReassignRequestDto StartReassignRequestDto

// NewStartReassignRequestDto instantiates a new StartReassignRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartReassignRequestDto(fromUserId string, toUserId string) *StartReassignRequestDto {
	this := StartReassignRequestDto{}
	this.FromUserId = fromUserId
	this.ToUserId = toUserId
	return &this
}

// NewStartReassignRequestDtoWithDefaults instantiates a new StartReassignRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartReassignRequestDtoWithDefaults() *StartReassignRequestDto {
	this := StartReassignRequestDto{}
	return &this
}

// GetFromUserId returns the FromUserId field value
func (o *StartReassignRequestDto) GetFromUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromUserId
}

// GetFromUserIdOk returns a tuple with the FromUserId field value
// and a boolean to check if the value has been set.
func (o *StartReassignRequestDto) GetFromUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromUserId, true
}

// SetFromUserId sets field value
func (o *StartReassignRequestDto) SetFromUserId(v string) {
	o.FromUserId = v
}

// GetToUserId returns the ToUserId field value
func (o *StartReassignRequestDto) GetToUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToUserId
}

// GetToUserIdOk returns a tuple with the ToUserId field value
// and a boolean to check if the value has been set.
func (o *StartReassignRequestDto) GetToUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToUserId, true
}

// SetToUserId sets field value
func (o *StartReassignRequestDto) SetToUserId(v string) {
	o.ToUserId = v
}

// GetDeleteProfile returns the DeleteProfile field value if set, zero value otherwise.
func (o *StartReassignRequestDto) GetDeleteProfile() bool {
	if o == nil || IsNil(o.DeleteProfile) {
		var ret bool
		return ret
	}
	return *o.DeleteProfile
}

// GetDeleteProfileOk returns a tuple with the DeleteProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartReassignRequestDto) GetDeleteProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteProfile) {
		return nil, false
	}
	return o.DeleteProfile, true
}

// HasDeleteProfile returns a boolean if a field has been set.
func (o *StartReassignRequestDto) IsDeleteProfileSet() bool {
	if o != nil && !IsNil(o.DeleteProfile) {
		return true
	}

	return false
}

// SetDeleteProfile gets a reference to the given bool and assigns it to the DeleteProfile field.
func (o *StartReassignRequestDto) SetDeleteProfile(v bool) {
	o.DeleteProfile = &v
}

func (o StartReassignRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartReassignRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fromUserId"] = o.FromUserId
	toSerialize["toUserId"] = o.ToUserId
	if !IsNil(o.DeleteProfile) {
		toSerialize["deleteProfile"] = o.DeleteProfile
	}
	return toSerialize, nil
}

func (o *StartReassignRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromUserId",
		"toUserId",
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

	varStartReassignRequestDto := _StartReassignRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varStartReassignRequestDto)

	if err != nil {
		return err
	}

	*o = StartReassignRequestDto(varStartReassignRequestDto)

	return err
}

type NullableStartReassignRequestDto struct {
	value *StartReassignRequestDto
	isSet bool
}

func (v NullableStartReassignRequestDto) Get() *StartReassignRequestDto {
	return v.value
}

func (v *NullableStartReassignRequestDto) Set(val *StartReassignRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStartReassignRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStartReassignRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartReassignRequestDto(val *StartReassignRequestDto) *NullableStartReassignRequestDto {
	return &NullableStartReassignRequestDto{value: val, isSet: true}
}

func (v NullableStartReassignRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartReassignRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

