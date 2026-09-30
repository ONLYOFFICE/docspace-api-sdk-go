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

// checks if the AdminMessageSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminMessageSettingsRequestsDto{}

// AdminMessageSettingsRequestsDto The message sent to the portal administrators, with the CAPTCHA proof that a person wrote it.
type AdminMessageSettingsRequestsDto struct {
	// What the sender wants to tell the portal administrators. Markup is stripped before the letter is written, so  a body that carries nothing but markup counts as empty and is refused with 400.
	Message NullableString `json:"message"`
	// The address the sender can be answered at, which the letter is signed with. It has to be a well-formed email  address.
	Email NullableString `json:"email"`
	// The language the letter is written in, as a culture name such as `en-US`. A culture the installation does not  have falls back to the portal language rather than failing the call.
	Culture NullableString `json:"culture,omitempty"`
	// Which CAPTCHA service the proof in `recaptchaResponse` came from. It has to match the service the  installation is configured with, which `GET api/2.0/capabilities` reports; the default value means the  installation is left to decide.
	RecaptchaType *RecaptchaType `json:"recaptchaType,omitempty"`
	// The token the CAPTCHA widget produced in the browser, passed on unchanged for the portal to verify with the  CAPTCHA service. It is single-use and short-lived, so it cannot be reused for a second message.
	RecaptchaResponse NullableString `json:"recaptchaResponse,omitempty"`
}

type _AdminMessageSettingsRequestsDto AdminMessageSettingsRequestsDto

// NewAdminMessageSettingsRequestsDto instantiates a new AdminMessageSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminMessageSettingsRequestsDto(message NullableString, email NullableString) *AdminMessageSettingsRequestsDto {
	this := AdminMessageSettingsRequestsDto{}
	this.Message = message
	this.Email = email
	return &this
}

// NewAdminMessageSettingsRequestsDtoWithDefaults instantiates a new AdminMessageSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminMessageSettingsRequestsDtoWithDefaults() *AdminMessageSettingsRequestsDto {
	this := AdminMessageSettingsRequestsDto{}
	return &this
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AdminMessageSettingsRequestsDto) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminMessageSettingsRequestsDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *AdminMessageSettingsRequestsDto) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AdminMessageSettingsRequestsDto) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminMessageSettingsRequestsDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *AdminMessageSettingsRequestsDto) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetCulture returns the Culture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminMessageSettingsRequestsDto) GetCulture() string {
	if o == nil || IsNil(o.Culture.Get()) {
		var ret string
		return ret
	}
	return *o.Culture.Get()
}

// GetCultureOk returns a tuple with the Culture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminMessageSettingsRequestsDto) GetCultureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Culture.Get(), o.Culture.IsSet()
}

// HasCulture returns a boolean if a field has been set.
func (o *AdminMessageSettingsRequestsDto) IsCultureSet() bool {
	if o != nil && o.Culture.IsSet() {
		return true
	}

	return false
}

// SetCulture gets a reference to the given NullableString and assigns it to the Culture field.
func (o *AdminMessageSettingsRequestsDto) SetCulture(v string) {
	o.Culture.Set(&v)
}
// SetCultureNil sets the value for Culture to be an explicit nil
func (o *AdminMessageSettingsRequestsDto) SetCultureNil() {
	o.Culture.Set(nil)
}

// UnsetCulture ensures that no value is present for Culture, not even an explicit nil
func (o *AdminMessageSettingsRequestsDto) UnsetCulture() {
	o.Culture.Unset()
}

// GetRecaptchaType returns the RecaptchaType field value if set, zero value otherwise.
func (o *AdminMessageSettingsRequestsDto) GetRecaptchaType() RecaptchaType {
	if o == nil || IsNil(o.RecaptchaType) {
		var ret RecaptchaType
		return ret
	}
	return *o.RecaptchaType
}

// GetRecaptchaTypeOk returns a tuple with the RecaptchaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessageSettingsRequestsDto) GetRecaptchaTypeOk() (*RecaptchaType, bool) {
	if o == nil || IsNil(o.RecaptchaType) {
		return nil, false
	}
	return o.RecaptchaType, true
}

// HasRecaptchaType returns a boolean if a field has been set.
func (o *AdminMessageSettingsRequestsDto) IsRecaptchaTypeSet() bool {
	if o != nil && !IsNil(o.RecaptchaType) {
		return true
	}

	return false
}

// SetRecaptchaType gets a reference to the given RecaptchaType and assigns it to the RecaptchaType field.
func (o *AdminMessageSettingsRequestsDto) SetRecaptchaType(v RecaptchaType) {
	o.RecaptchaType = &v
}

// GetRecaptchaResponse returns the RecaptchaResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminMessageSettingsRequestsDto) GetRecaptchaResponse() string {
	if o == nil || IsNil(o.RecaptchaResponse.Get()) {
		var ret string
		return ret
	}
	return *o.RecaptchaResponse.Get()
}

// GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminMessageSettingsRequestsDto) GetRecaptchaResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecaptchaResponse.Get(), o.RecaptchaResponse.IsSet()
}

// HasRecaptchaResponse returns a boolean if a field has been set.
func (o *AdminMessageSettingsRequestsDto) IsRecaptchaResponseSet() bool {
	if o != nil && o.RecaptchaResponse.IsSet() {
		return true
	}

	return false
}

// SetRecaptchaResponse gets a reference to the given NullableString and assigns it to the RecaptchaResponse field.
func (o *AdminMessageSettingsRequestsDto) SetRecaptchaResponse(v string) {
	o.RecaptchaResponse.Set(&v)
}
// SetRecaptchaResponseNil sets the value for RecaptchaResponse to be an explicit nil
func (o *AdminMessageSettingsRequestsDto) SetRecaptchaResponseNil() {
	o.RecaptchaResponse.Set(nil)
}

// UnsetRecaptchaResponse ensures that no value is present for RecaptchaResponse, not even an explicit nil
func (o *AdminMessageSettingsRequestsDto) UnsetRecaptchaResponse() {
	o.RecaptchaResponse.Unset()
}

func (o AdminMessageSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminMessageSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message.Get()
	toSerialize["email"] = o.Email.Get()
	if o.Culture.IsSet() {
		toSerialize["culture"] = o.Culture.Get()
	}
	if !IsNil(o.RecaptchaType) {
		toSerialize["recaptchaType"] = o.RecaptchaType
	}
	if o.RecaptchaResponse.IsSet() {
		toSerialize["recaptchaResponse"] = o.RecaptchaResponse.Get()
	}
	return toSerialize, nil
}

func (o *AdminMessageSettingsRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varAdminMessageSettingsRequestsDto := _AdminMessageSettingsRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAdminMessageSettingsRequestsDto)

	if err != nil {
		return err
	}

	*o = AdminMessageSettingsRequestsDto(varAdminMessageSettingsRequestsDto)

	return err
}

type NullableAdminMessageSettingsRequestsDto struct {
	value *AdminMessageSettingsRequestsDto
	isSet bool
}

func (v NullableAdminMessageSettingsRequestsDto) Get() *AdminMessageSettingsRequestsDto {
	return v.value
}

func (v *NullableAdminMessageSettingsRequestsDto) Set(val *AdminMessageSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminMessageSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminMessageSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminMessageSettingsRequestsDto(val *AdminMessageSettingsRequestsDto) *NullableAdminMessageSettingsRequestsDto {
	return &NullableAdminMessageSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableAdminMessageSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminMessageSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

