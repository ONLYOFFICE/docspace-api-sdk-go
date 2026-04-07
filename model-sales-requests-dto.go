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
	UserName NullableString `json:"userName,omitempty"`
	// The contact email address for the sales inquiry.
	Email NullableString `json:"email"`
	// The details of the sales inquiry or payment request.
	Message NullableString `json:"message"`
}

type _SalesRequestsDto SalesRequestsDto

// NewSalesRequestsDto instantiates a new SalesRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesRequestsDto(email NullableString, message NullableString) *SalesRequestsDto {
	this := SalesRequestsDto{}
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

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalesRequestsDto) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesRequestsDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *SalesRequestsDto) IsUserNameSet() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *SalesRequestsDto) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *SalesRequestsDto) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *SalesRequestsDto) UnsetUserName() {
	o.UserName.Unset()
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SalesRequestsDto) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesRequestsDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *SalesRequestsDto) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SalesRequestsDto) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesRequestsDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *SalesRequestsDto) SetMessage(v string) {
	o.Message.Set(&v)
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
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	toSerialize["email"] = o.Email.Get()
	toSerialize["message"] = o.Message.Get()
	return toSerialize, nil
}

func (o *SalesRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

