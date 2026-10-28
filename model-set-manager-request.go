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

// checks if the SetManagerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetManagerRequest{}

// SetManagerRequest The request for setting a group manager.
type SetManagerRequest struct {
	// The account to make the manager. It has to exist, otherwise the operation answers 404, and it is added to the  group at the same time, so it does not have to be a member beforehand.
	UserId string `json:"userId"`
}

type _SetManagerRequest SetManagerRequest

// NewSetManagerRequest instantiates a new SetManagerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetManagerRequest(userId string) *SetManagerRequest {
	this := SetManagerRequest{}
	this.UserId = userId
	return &this
}

// NewSetManagerRequestWithDefaults instantiates a new SetManagerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetManagerRequestWithDefaults() *SetManagerRequest {
	this := SetManagerRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *SetManagerRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SetManagerRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SetManagerRequest) SetUserId(v string) {
	o.UserId = v
}

func (o SetManagerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetManagerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *SetManagerRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSetManagerRequest := _SetManagerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSetManagerRequest)

	if err != nil {
		return err
	}

	*o = SetManagerRequest(varSetManagerRequest)

	return err
}

type NullableSetManagerRequest struct {
	value *SetManagerRequest
	isSet bool
}

func (v NullableSetManagerRequest) Get() *SetManagerRequest {
	return v.value
}

func (v *NullableSetManagerRequest) Set(val *SetManagerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetManagerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetManagerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetManagerRequest(val *SetManagerRequest) *NullableSetManagerRequest {
	return &NullableSetManagerRequest{value: val, isSet: true}
}

func (v NullableSetManagerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetManagerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

