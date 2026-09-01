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

// checks if the AccessRequestKeyDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestKeyDto{}

// AccessRequestKeyDto The encryption key granting one user access to a file.
type AccessRequestKeyDto struct {
	// User ID
	UserId *string `json:"userId,omitempty"`
	// Public key ID
	PublicKeyId *string `json:"publicKeyId,omitempty"`
	// Encrypted private key
	PrivateKeyEnc NullableString `json:"privateKeyEnc,omitempty"`
}

// NewAccessRequestKeyDto instantiates a new AccessRequestKeyDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestKeyDto() *AccessRequestKeyDto {
	this := AccessRequestKeyDto{}
	return &this
}

// NewAccessRequestKeyDtoWithDefaults instantiates a new AccessRequestKeyDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestKeyDtoWithDefaults() *AccessRequestKeyDto {
	this := AccessRequestKeyDto{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AccessRequestKeyDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestKeyDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AccessRequestKeyDto) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AccessRequestKeyDto) SetUserId(v string) {
	o.UserId = &v
}

// GetPublicKeyId returns the PublicKeyId field value if set, zero value otherwise.
func (o *AccessRequestKeyDto) GetPublicKeyId() string {
	if o == nil || IsNil(o.PublicKeyId) {
		var ret string
		return ret
	}
	return *o.PublicKeyId
}

// GetPublicKeyIdOk returns a tuple with the PublicKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestKeyDto) GetPublicKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKeyId) {
		return nil, false
	}
	return o.PublicKeyId, true
}

// HasPublicKeyId returns a boolean if a field has been set.
func (o *AccessRequestKeyDto) IsPublicKeyIdSet() bool {
	if o != nil && !IsNil(o.PublicKeyId) {
		return true
	}

	return false
}

// SetPublicKeyId gets a reference to the given string and assigns it to the PublicKeyId field.
func (o *AccessRequestKeyDto) SetPublicKeyId(v string) {
	o.PublicKeyId = &v
}

// GetPrivateKeyEnc returns the PrivateKeyEnc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestKeyDto) GetPrivateKeyEnc() string {
	if o == nil || IsNil(o.PrivateKeyEnc.Get()) {
		var ret string
		return ret
	}
	return *o.PrivateKeyEnc.Get()
}

// GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestKeyDto) GetPrivateKeyEncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKeyEnc.Get(), o.PrivateKeyEnc.IsSet()
}

// HasPrivateKeyEnc returns a boolean if a field has been set.
func (o *AccessRequestKeyDto) IsPrivateKeyEncSet() bool {
	if o != nil && o.PrivateKeyEnc.IsSet() {
		return true
	}

	return false
}

// SetPrivateKeyEnc gets a reference to the given NullableString and assigns it to the PrivateKeyEnc field.
func (o *AccessRequestKeyDto) SetPrivateKeyEnc(v string) {
	o.PrivateKeyEnc.Set(&v)
}
// SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil
func (o *AccessRequestKeyDto) SetPrivateKeyEncNil() {
	o.PrivateKeyEnc.Set(nil)
}

// UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil
func (o *AccessRequestKeyDto) UnsetPrivateKeyEnc() {
	o.PrivateKeyEnc.Unset()
}

func (o AccessRequestKeyDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestKeyDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.PublicKeyId) {
		toSerialize["publicKeyId"] = o.PublicKeyId
	}
	if o.PrivateKeyEnc.IsSet() {
		toSerialize["privateKeyEnc"] = o.PrivateKeyEnc.Get()
	}
	return toSerialize, nil
}

type NullableAccessRequestKeyDto struct {
	value *AccessRequestKeyDto
	isSet bool
}

func (v NullableAccessRequestKeyDto) Get() *AccessRequestKeyDto {
	return v.value
}

func (v *NullableAccessRequestKeyDto) Set(val *AccessRequestKeyDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestKeyDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestKeyDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestKeyDto(val *AccessRequestKeyDto) *NullableAccessRequestKeyDto {
	return &NullableAccessRequestKeyDto{value: val, isSet: true}
}

func (v NullableAccessRequestKeyDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestKeyDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

