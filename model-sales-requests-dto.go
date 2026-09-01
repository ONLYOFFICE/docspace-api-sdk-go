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

// checks if the SalesRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesRequestsDto{}

// SalesRequestsDto The request parameters for handling sales and payment inquiries in the portal.
type SalesRequestsDto struct {
	// The name of the user submitting the sales request.
	UserName string `json:"userName"`
	// The contact email address for the sales inquiry.
	Email string `json:"email"`
	// The details of the sales inquiry or payment request.
	Message string `json:"message"`
}

type _SalesRequestsDto SalesRequestsDto

// NewSalesRequestsDto instantiates a new SalesRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesRequestsDto(userName string, email string, message string) *SalesRequestsDto {
	this := SalesRequestsDto{}
	this.UserName = userName
	this.Email = email
	this.Message = message
	return &this
}

// NewSalesRequestsDtoWithDefaults instantiates a new SalesRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesRequestsDtoWithDefaults() *SalesRequestsDto {
	this := SalesRequestsDto{}
	return &this
}

// GetUserName returns the UserName field value
func (o *SalesRequestsDto) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *SalesRequestsDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *SalesRequestsDto) SetUserName(v string) {
	o.UserName = v
}

// GetEmail returns the Email field value
func (o *SalesRequestsDto) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SalesRequestsDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SalesRequestsDto) SetEmail(v string) {
	o.Email = v
}

// GetMessage returns the Message field value
func (o *SalesRequestsDto) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SalesRequestsDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SalesRequestsDto) SetMessage(v string) {
	o.Message = v
}

func (o SalesRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userName"] = o.UserName
	toSerialize["email"] = o.Email
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *SalesRequestsDto) UnmarshalJSON(data []byte) (err error) {
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

	varSalesRequestsDto := _SalesRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSalesRequestsDto)

	if err != nil {
		return err
	}

	*o = SalesRequestsDto(varSalesRequestsDto)

	return err
}

type NullableSalesRequestsDto struct {
	value *SalesRequestsDto
	isSet bool
}

func (v NullableSalesRequestsDto) Get() *SalesRequestsDto {
	return v.value
}

func (v *NullableSalesRequestsDto) Set(val *SalesRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesRequestsDto(val *SalesRequestsDto) *NullableSalesRequestsDto {
	return &NullableSalesRequestsDto{value: val, isSet: true}
}

func (v NullableSalesRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

