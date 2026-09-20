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

// checks if the ChangePasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangePasswordRequest{}

// ChangePasswordRequest The request parameters for updating a user password.
type ChangePasswordRequest struct {
	// The new password in plain text. It is checked against the portal password policy and rejected with 400 when  it is too weak, then hashed by the portal. Send it only over a secure connection, and prefer `passwordHash`  when the client can compute it.
	Password NullableString `json:"password,omitempty"`
	// The new password already hashed by the client, which is what the portal stores. It is a PBKDF2-HMACSHA256  hash of the plain password, computed with the salt, the iteration count and the key size the portal settings  publish, and written as lowercase hexadecimal. When it is sent, `password` is ignored and the password policy  is not applied.
	PasswordHash NullableString `json:"passwordHash,omitempty"`
}

// NewChangePasswordRequest instantiates a new ChangePasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangePasswordRequest() *ChangePasswordRequest {
	this := ChangePasswordRequest{}
	return &this
}

// NewChangePasswordRequestWithDefaults instantiates a new ChangePasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangePasswordRequestWithDefaults() *ChangePasswordRequest {
	this := ChangePasswordRequest{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangePasswordRequest) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangePasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *ChangePasswordRequest) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *ChangePasswordRequest) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *ChangePasswordRequest) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *ChangePasswordRequest) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangePasswordRequest) GetPasswordHash() string {
	if o == nil || IsNil(o.PasswordHash.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangePasswordRequest) GetPasswordHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *ChangePasswordRequest) IsPasswordHashSet() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *ChangePasswordRequest) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *ChangePasswordRequest) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *ChangePasswordRequest) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

func (o ChangePasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangePasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.PasswordHash.IsSet() {
		toSerialize["passwordHash"] = o.PasswordHash.Get()
	}
	return toSerialize, nil
}

type NullableChangePasswordRequest struct {
	value *ChangePasswordRequest
	isSet bool
}

func (v NullableChangePasswordRequest) Get() *ChangePasswordRequest {
	return v.value
}

func (v *NullableChangePasswordRequest) Set(val *ChangePasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangePasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangePasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangePasswordRequest(val *ChangePasswordRequest) *NullableChangePasswordRequest {
	return &NullableChangePasswordRequest{value: val, isSet: true}
}

func (v NullableChangePasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangePasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

