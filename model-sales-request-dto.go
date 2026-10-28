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

// checks if the SalesRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesRequestDto{}

// SalesRequestDto Who is writing to the ONLYOFFICE sales team, and what about.
type SalesRequestDto struct {
	// The name the sales team should address the reply to. It is sent as written and is not matched against any  portal account; an empty value fails the request with 400.
	UserName string `json:"userName"`
	// The address the answer is sent to. It has to be a well-formed email address and need not be the caller portal  address; an empty or malformed value fails the request with 400.
	Email string `json:"email"`
	// What is being asked of the sales team - a quote, an invoice, or a plan that cannot be bought online. An empty  value fails the request with 400.
	Message string `json:"message"`
}

type _SalesRequestDto SalesRequestDto

// NewSalesRequestDto instantiates a new SalesRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesRequestDto(userName string, email string, message string) *SalesRequestDto {
	this := SalesRequestDto{}
	this.UserName = userName
	this.Email = email
	this.Message = message
	return &this
}

// NewSalesRequestDtoWithDefaults instantiates a new SalesRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesRequestDtoWithDefaults() *SalesRequestDto {
	this := SalesRequestDto{}
	return &this
}

// GetUserName returns the UserName field value
func (o *SalesRequestDto) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *SalesRequestDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *SalesRequestDto) SetUserName(v string) {
	o.UserName = v
}

// GetEmail returns the Email field value
func (o *SalesRequestDto) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SalesRequestDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SalesRequestDto) SetEmail(v string) {
	o.Email = v
}

// GetMessage returns the Message field value
func (o *SalesRequestDto) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SalesRequestDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SalesRequestDto) SetMessage(v string) {
	o.Message = v
}

func (o SalesRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userName"] = o.UserName
	toSerialize["email"] = o.Email
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *SalesRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userName",
		"email",
		"message",
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

	varSalesRequestDto := _SalesRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSalesRequestDto)

	if err != nil {
		return err
	}

	*o = SalesRequestDto(varSalesRequestDto)

	return err
}

type NullableSalesRequestDto struct {
	value *SalesRequestDto
	isSet bool
}

func (v NullableSalesRequestDto) Get() *SalesRequestDto {
	return v.value
}

func (v *NullableSalesRequestDto) Set(val *SalesRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesRequestDto(val *SalesRequestDto) *NullableSalesRequestDto {
	return &NullableSalesRequestDto{value: val, isSet: true}
}

func (v NullableSalesRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

