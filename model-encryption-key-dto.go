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
	"time"
)

// checks if the EncryptionKeyDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptionKeyDto{}

// EncryptionKeyDto The encryption key pair of a user.
type EncryptionKeyDto struct {
	// The identifier of the key pair.
	Id *string `json:"id,omitempty"`
	// The identifier of the user the key pair belongs to.
	UserId *string `json:"userId,omitempty"`
	// The date and time when the key pair was created.
	Date *time.Time `json:"date,omitempty"`
	// The public key of the pair, used to encrypt the file keys.
	PublicKey NullableString `json:"publicKey,omitempty"`
	// The private key of the pair, encrypted with the user password.
	PrivateKeyEnc NullableString `json:"privateKeyEnc,omitempty"`
	// The identifier of the crypto engine the key pair was issued for.
	CryptoEngineId NullableString `json:"cryptoEngineId,omitempty"`
}

// NewEncryptionKeyDto instantiates a new EncryptionKeyDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionKeyDto() *EncryptionKeyDto {
	this := EncryptionKeyDto{}
	return &this
}

// NewEncryptionKeyDtoWithDefaults instantiates a new EncryptionKeyDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionKeyDtoWithDefaults() *EncryptionKeyDto {
	this := EncryptionKeyDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EncryptionKeyDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionKeyDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EncryptionKeyDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EncryptionKeyDto) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EncryptionKeyDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionKeyDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EncryptionKeyDto) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EncryptionKeyDto) SetUserId(v string) {
	o.UserId = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *EncryptionKeyDto) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionKeyDto) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *EncryptionKeyDto) IsDateSet() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *EncryptionKeyDto) SetDate(v time.Time) {
	o.Date = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionKeyDto) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionKeyDto) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *EncryptionKeyDto) IsPublicKeySet() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *EncryptionKeyDto) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *EncryptionKeyDto) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *EncryptionKeyDto) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetPrivateKeyEnc returns the PrivateKeyEnc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionKeyDto) GetPrivateKeyEnc() string {
	if o == nil || IsNil(o.PrivateKeyEnc.Get()) {
		var ret string
		return ret
	}
	return *o.PrivateKeyEnc.Get()
}

// GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionKeyDto) GetPrivateKeyEncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKeyEnc.Get(), o.PrivateKeyEnc.IsSet()
}

// HasPrivateKeyEnc returns a boolean if a field has been set.
func (o *EncryptionKeyDto) IsPrivateKeyEncSet() bool {
	if o != nil && o.PrivateKeyEnc.IsSet() {
		return true
	}

	return false
}

// SetPrivateKeyEnc gets a reference to the given NullableString and assigns it to the PrivateKeyEnc field.
func (o *EncryptionKeyDto) SetPrivateKeyEnc(v string) {
	o.PrivateKeyEnc.Set(&v)
}
// SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil
func (o *EncryptionKeyDto) SetPrivateKeyEncNil() {
	o.PrivateKeyEnc.Set(nil)
}

// UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil
func (o *EncryptionKeyDto) UnsetPrivateKeyEnc() {
	o.PrivateKeyEnc.Unset()
}

// GetCryptoEngineId returns the CryptoEngineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionKeyDto) GetCryptoEngineId() string {
	if o == nil || IsNil(o.CryptoEngineId.Get()) {
		var ret string
		return ret
	}
	return *o.CryptoEngineId.Get()
}

// GetCryptoEngineIdOk returns a tuple with the CryptoEngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionKeyDto) GetCryptoEngineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CryptoEngineId.Get(), o.CryptoEngineId.IsSet()
}

// HasCryptoEngineId returns a boolean if a field has been set.
func (o *EncryptionKeyDto) IsCryptoEngineIdSet() bool {
	if o != nil && o.CryptoEngineId.IsSet() {
		return true
	}

	return false
}

// SetCryptoEngineId gets a reference to the given NullableString and assigns it to the CryptoEngineId field.
func (o *EncryptionKeyDto) SetCryptoEngineId(v string) {
	o.CryptoEngineId.Set(&v)
}
// SetCryptoEngineIdNil sets the value for CryptoEngineId to be an explicit nil
func (o *EncryptionKeyDto) SetCryptoEngineIdNil() {
	o.CryptoEngineId.Set(nil)
}

// UnsetCryptoEngineId ensures that no value is present for CryptoEngineId, not even an explicit nil
func (o *EncryptionKeyDto) UnsetCryptoEngineId() {
	o.CryptoEngineId.Unset()
}

func (o EncryptionKeyDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptionKeyDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if o.PublicKey.IsSet() {
		toSerialize["publicKey"] = o.PublicKey.Get()
	}
	if o.PrivateKeyEnc.IsSet() {
		toSerialize["privateKeyEnc"] = o.PrivateKeyEnc.Get()
	}
	if o.CryptoEngineId.IsSet() {
		toSerialize["cryptoEngineId"] = o.CryptoEngineId.Get()
	}
	return toSerialize, nil
}

type NullableEncryptionKeyDto struct {
	value *EncryptionKeyDto
	isSet bool
}

func (v NullableEncryptionKeyDto) Get() *EncryptionKeyDto {
	return v.value
}

func (v *NullableEncryptionKeyDto) Set(val *EncryptionKeyDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionKeyDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionKeyDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionKeyDto(val *EncryptionKeyDto) *NullableEncryptionKeyDto {
	return &NullableEncryptionKeyDto{value: val, isSet: true}
}

func (v NullableEncryptionKeyDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionKeyDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

