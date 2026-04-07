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

// checks if the TfaValidateRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TfaValidateRequestsDto{}

// TfaValidateRequestsDto The request parameters for validating the two-factor authentication codes.
type TfaValidateRequestsDto struct {
	// The verification code provided by the user.
	Code NullableString `json:"code"`
	// Specifies whether the authentication is session-based.
	Session *bool `json:"session,omitempty"`
}

type _TfaValidateRequestsDto TfaValidateRequestsDto

// NewTfaValidateRequestsDto instantiates a new TfaValidateRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaValidateRequestsDto(code NullableString) *TfaValidateRequestsDto {
	this := TfaValidateRequestsDto{}
	this.Code = code
	return &this
}

// NewTfaValidateRequestsDtoWithDefaults instantiates a new TfaValidateRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaValidateRequestsDtoWithDefaults() *TfaValidateRequestsDto {
	this := TfaValidateRequestsDto{}
	return &this
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TfaValidateRequestsDto) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}

	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaValidateRequestsDto) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// SetCode sets field value
func (o *TfaValidateRequestsDto) SetCode(v string) {
	o.Code.Set(&v)
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *TfaValidateRequestsDto) GetSession() bool {
	if o == nil || IsNil(o.Session) {
		var ret bool
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaValidateRequestsDto) GetSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *TfaValidateRequestsDto) IsSessionSet() bool {
	if o != nil && !IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given bool and assigns it to the Session field.
func (o *TfaValidateRequestsDto) SetSession(v bool) {
	o.Session = &v
}

func (o TfaValidateRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TfaValidateRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code.Get()
	if !IsNil(o.Session) {
		toSerialize["session"] = o.Session
	}
	return toSerialize, nil
}

func (o *TfaValidateRequestsDto) UnmarshalJSON(data []byte) (err error) {
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

	varTfaValidateRequestsDto := _TfaValidateRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTfaValidateRequestsDto)

	if err != nil {
		return err
	}

	*o = TfaValidateRequestsDto(varTfaValidateRequestsDto)

	return err
}

type NullableTfaValidateRequestsDto struct {
	value *TfaValidateRequestsDto
	isSet bool
}

func (v NullableTfaValidateRequestsDto) Get() *TfaValidateRequestsDto {
	return v.value
}

func (v *NullableTfaValidateRequestsDto) Set(val *TfaValidateRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaValidateRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaValidateRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaValidateRequestsDto(val *TfaValidateRequestsDto) *NullableTfaValidateRequestsDto {
	return &NullableTfaValidateRequestsDto{value: val, isSet: true}
}

func (v NullableTfaValidateRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaValidateRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

