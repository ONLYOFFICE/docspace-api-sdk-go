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

// checks if the EmailMemberRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailMemberRequestDto{}

// EmailMemberRequestDto The request parameters for the user email.
type EmailMemberRequestDto struct {
	// The user email address.
	Email string `json:"email"`
	RecaptchaType *RecaptchaType `json:"recaptchaType,omitempty"`
	// The user's response to the CAPTCHA challenge.
	RecaptchaResponse NullableString `json:"recaptchaResponse,omitempty"`
}

type _EmailMemberRequestDto EmailMemberRequestDto

// NewEmailMemberRequestDto instantiates a new EmailMemberRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailMemberRequestDto(email string) *EmailMemberRequestDto {
	this := EmailMemberRequestDto{}
	this.Email = email
	return &this
}

// NewEmailMemberRequestDtoWithDefaults instantiates a new EmailMemberRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailMemberRequestDtoWithDefaults() *EmailMemberRequestDto {
	this := EmailMemberRequestDto{}
	return &this
}

// GetEmail returns the Email field value
func (o *EmailMemberRequestDto) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmailMemberRequestDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmailMemberRequestDto) SetEmail(v string) {
	o.Email = v
}

// GetRecaptchaType returns the RecaptchaType field value if set, zero value otherwise.
func (o *EmailMemberRequestDto) GetRecaptchaType() RecaptchaType {
	if o == nil || IsNil(o.RecaptchaType) {
		var ret RecaptchaType
		return ret
	}
	return *o.RecaptchaType
}

// GetRecaptchaTypeOk returns a tuple with the RecaptchaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMemberRequestDto) GetRecaptchaTypeOk() (*RecaptchaType, bool) {
	if o == nil || IsNil(o.RecaptchaType) {
		return nil, false
	}
	return o.RecaptchaType, true
}

// HasRecaptchaType returns a boolean if a field has been set.
func (o *EmailMemberRequestDto) IsRecaptchaTypeSet() bool {
	if o != nil && !IsNil(o.RecaptchaType) {
		return true
	}

	return false
}

// SetRecaptchaType gets a reference to the given RecaptchaType and assigns it to the RecaptchaType field.
func (o *EmailMemberRequestDto) SetRecaptchaType(v RecaptchaType) {
	o.RecaptchaType = &v
}

// GetRecaptchaResponse returns the RecaptchaResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailMemberRequestDto) GetRecaptchaResponse() string {
	if o == nil || IsNil(o.RecaptchaResponse.Get()) {
		var ret string
		return ret
	}
	return *o.RecaptchaResponse.Get()
}

// GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailMemberRequestDto) GetRecaptchaResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecaptchaResponse.Get(), o.RecaptchaResponse.IsSet()
}

// HasRecaptchaResponse returns a boolean if a field has been set.
func (o *EmailMemberRequestDto) IsRecaptchaResponseSet() bool {
	if o != nil && o.RecaptchaResponse.IsSet() {
		return true
	}

	return false
}

// SetRecaptchaResponse gets a reference to the given NullableString and assigns it to the RecaptchaResponse field.
func (o *EmailMemberRequestDto) SetRecaptchaResponse(v string) {
	o.RecaptchaResponse.Set(&v)
}
// SetRecaptchaResponseNil sets the value for RecaptchaResponse to be an explicit nil
func (o *EmailMemberRequestDto) SetRecaptchaResponseNil() {
	o.RecaptchaResponse.Set(nil)
}

// UnsetRecaptchaResponse ensures that no value is present for RecaptchaResponse, not even an explicit nil
func (o *EmailMemberRequestDto) UnsetRecaptchaResponse() {
	o.RecaptchaResponse.Unset()
}

func (o EmailMemberRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailMemberRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.RecaptchaType) {
		toSerialize["recaptchaType"] = o.RecaptchaType
	}
	if o.RecaptchaResponse.IsSet() {
		toSerialize["recaptchaResponse"] = o.RecaptchaResponse.Get()
	}
	return toSerialize, nil
}

func (o *EmailMemberRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varEmailMemberRequestDto := _EmailMemberRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varEmailMemberRequestDto)

	if err != nil {
		return err
	}

	*o = EmailMemberRequestDto(varEmailMemberRequestDto)

	return err
}

type NullableEmailMemberRequestDto struct {
	value *EmailMemberRequestDto
	isSet bool
}

func (v NullableEmailMemberRequestDto) Get() *EmailMemberRequestDto {
	return v.value
}

func (v *NullableEmailMemberRequestDto) Set(val *EmailMemberRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailMemberRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailMemberRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailMemberRequestDto(val *EmailMemberRequestDto) *NullableEmailMemberRequestDto {
	return &NullableEmailMemberRequestDto{value: val, isSet: true}
}

func (v NullableEmailMemberRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailMemberRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

