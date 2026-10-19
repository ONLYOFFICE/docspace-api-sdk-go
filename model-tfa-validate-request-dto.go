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

// checks if the TfaValidateRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TfaValidateRequestDto{}

// TfaValidateRequestDto The one-time code that completes a pending two-factor step, and how long the resulting sign-in lasts.
type TfaValidateRequestDto struct {
	// The code to check - either one from the authenticator application or one of the account's unused backup  codes, which is spent by the check. A wrong code is refused with 400 and counts against the portal login  attempt limit.
	Code NullableString `json:"code"`
	// Whether the sign-in that follows is tied to the browser session. When it is, the session ends with the  browser rather than lasting for the portal session lifetime.
	Session *bool `json:"session,omitempty"`
}

type _TfaValidateRequestDto TfaValidateRequestDto

// NewTfaValidateRequestDto instantiates a new TfaValidateRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaValidateRequestDto(code NullableString) *TfaValidateRequestDto {
	this := TfaValidateRequestDto{}
	this.Code = code
	return &this
}

// NewTfaValidateRequestDtoWithDefaults instantiates a new TfaValidateRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaValidateRequestDtoWithDefaults() *TfaValidateRequestDto {
	this := TfaValidateRequestDto{}
	return &this
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TfaValidateRequestDto) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}

	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaValidateRequestDto) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// SetCode sets field value
func (o *TfaValidateRequestDto) SetCode(v string) {
	o.Code.Set(&v)
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *TfaValidateRequestDto) GetSession() bool {
	if o == nil || IsNil(o.Session) {
		var ret bool
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaValidateRequestDto) GetSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *TfaValidateRequestDto) IsSessionSet() bool {
	if o != nil && !IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given bool and assigns it to the Session field.
func (o *TfaValidateRequestDto) SetSession(v bool) {
	o.Session = &v
}

func (o TfaValidateRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TfaValidateRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code.Get()
	if !IsNil(o.Session) {
		toSerialize["session"] = o.Session
	}
	return toSerialize, nil
}

func (o *TfaValidateRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varTfaValidateRequestDto := _TfaValidateRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTfaValidateRequestDto)

	if err != nil {
		return err
	}

	*o = TfaValidateRequestDto(varTfaValidateRequestDto)

	return err
}

type NullableTfaValidateRequestDto struct {
	value *TfaValidateRequestDto
	isSet bool
}

func (v NullableTfaValidateRequestDto) Get() *TfaValidateRequestDto {
	return v.value
}

func (v *NullableTfaValidateRequestDto) Set(val *TfaValidateRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaValidateRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaValidateRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaValidateRequestDto(val *TfaValidateRequestDto) *NullableTfaValidateRequestDto {
	return &NullableTfaValidateRequestDto{value: val, isSet: true}
}

func (v NullableTfaValidateRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaValidateRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

