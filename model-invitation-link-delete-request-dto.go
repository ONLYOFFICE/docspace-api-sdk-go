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

// checks if the InvitationLinkDeleteRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitationLinkDeleteRequestDto{}

// InvitationLinkDeleteRequestDto The request parameters for deleting an invitation link.
type InvitationLinkDeleteRequestDto struct {
	// The ID of the invitation link.
	Id string `json:"id"`
}

type _InvitationLinkDeleteRequestDto InvitationLinkDeleteRequestDto

// NewInvitationLinkDeleteRequestDto instantiates a new InvitationLinkDeleteRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationLinkDeleteRequestDto(id string) *InvitationLinkDeleteRequestDto {
	this := InvitationLinkDeleteRequestDto{}
	this.Id = id
	return &this
}

// NewInvitationLinkDeleteRequestDtoWithDefaults instantiates a new InvitationLinkDeleteRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationLinkDeleteRequestDtoWithDefaults() *InvitationLinkDeleteRequestDto {
	this := InvitationLinkDeleteRequestDto{}
	return &this
}

// GetId returns the Id field value
func (o *InvitationLinkDeleteRequestDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvitationLinkDeleteRequestDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvitationLinkDeleteRequestDto) SetId(v string) {
	o.Id = v
}

func (o InvitationLinkDeleteRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitationLinkDeleteRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *InvitationLinkDeleteRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varInvitationLinkDeleteRequestDto := _InvitationLinkDeleteRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInvitationLinkDeleteRequestDto)

	if err != nil {
		return err
	}

	*o = InvitationLinkDeleteRequestDto(varInvitationLinkDeleteRequestDto)

	return err
}

type NullableInvitationLinkDeleteRequestDto struct {
	value *InvitationLinkDeleteRequestDto
	isSet bool
}

func (v NullableInvitationLinkDeleteRequestDto) Get() *InvitationLinkDeleteRequestDto {
	return v.value
}

func (v *NullableInvitationLinkDeleteRequestDto) Set(val *InvitationLinkDeleteRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationLinkDeleteRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationLinkDeleteRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationLinkDeleteRequestDto(val *InvitationLinkDeleteRequestDto) *NullableInvitationLinkDeleteRequestDto {
	return &NullableInvitationLinkDeleteRequestDto{value: val, isSet: true}
}

func (v NullableInvitationLinkDeleteRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationLinkDeleteRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

