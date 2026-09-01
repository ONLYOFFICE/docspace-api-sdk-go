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
)

// checks if the ChangeEmailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeEmailRequest{}

// ChangeEmailRequest The request parameters for updating a user email.
type ChangeEmailRequest struct {
	// The user email address.
	Email NullableString `json:"email,omitempty"`
	// The user encrypted email address.
	EncEmail NullableString `json:"encEmail,omitempty"`
}

// NewChangeEmailRequest instantiates a new ChangeEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeEmailRequest() *ChangeEmailRequest {
	this := ChangeEmailRequest{}
	return &this
}

// NewChangeEmailRequestWithDefaults instantiates a new ChangeEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeEmailRequestWithDefaults() *ChangeEmailRequest {
	this := ChangeEmailRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeEmailRequest) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeEmailRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ChangeEmailRequest) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *ChangeEmailRequest) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *ChangeEmailRequest) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *ChangeEmailRequest) UnsetEmail() {
	o.Email.Unset()
}

// GetEncEmail returns the EncEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeEmailRequest) GetEncEmail() string {
	if o == nil || IsNil(o.EncEmail.Get()) {
		var ret string
		return ret
	}
	return *o.EncEmail.Get()
}

// GetEncEmailOk returns a tuple with the EncEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeEmailRequest) GetEncEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncEmail.Get(), o.EncEmail.IsSet()
}

// HasEncEmail returns a boolean if a field has been set.
func (o *ChangeEmailRequest) IsEncEmailSet() bool {
	if o != nil && o.EncEmail.IsSet() {
		return true
	}

	return false
}

// SetEncEmail gets a reference to the given NullableString and assigns it to the EncEmail field.
func (o *ChangeEmailRequest) SetEncEmail(v string) {
	o.EncEmail.Set(&v)
}
// SetEncEmailNil sets the value for EncEmail to be an explicit nil
func (o *ChangeEmailRequest) SetEncEmailNil() {
	o.EncEmail.Set(nil)
}

// UnsetEncEmail ensures that no value is present for EncEmail, not even an explicit nil
func (o *ChangeEmailRequest) UnsetEncEmail() {
	o.EncEmail.Unset()
}

func (o ChangeEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeEmailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.EncEmail.IsSet() {
		toSerialize["encEmail"] = o.EncEmail.Get()
	}
	return toSerialize, nil
}

type NullableChangeEmailRequest struct {
	value *ChangeEmailRequest
	isSet bool
}

func (v NullableChangeEmailRequest) Get() *ChangeEmailRequest {
	return v.value
}

func (v *NullableChangeEmailRequest) Set(val *ChangeEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeEmailRequest(val *ChangeEmailRequest) *NullableChangeEmailRequest {
	return &NullableChangeEmailRequest{value: val, isSet: true}
}

func (v NullableChangeEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

