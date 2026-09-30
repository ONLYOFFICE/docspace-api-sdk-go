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

// checks if the AuthenticationTokenDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationTokenDto{}

// AuthenticationTokenDto The outcome of a sign-in attempt: either the authentication token, or the second factor still to be passed.
type AuthenticationTokenDto struct {
	// The token to put in the `Authorization` header of later calls. It is empty whenever a second factor is  still outstanding, which is what `sms` or `tfa` then says; the same token is also set as a portal cookie by  the call that issued it, so a browser client does not have to carry it itself.
	Token NullableString `json:"token,omitempty"`
	// When the token stops being accepted. It stays at its zero value when `session=true` tied the token to the  browser session instead of to a fixed moment. On the two operations that only send an SMS it carries a  different meaning: there is no token, and this is the moment the code that was just sent expires.
	Expires *time.Time `json:"expires,omitempty"`
	// Whether an SMS code is the second factor in play. Next to an empty `token` it means the code has to be sent  to `POST api/2.0/authentication/{code}` before a token is issued; next to a filled `token` it means the  code just accepted was an SMS one.
	Sms *bool `json:"sms,omitempty"`
	// The stored phone number with its middle digits masked, filled in only while `sms` is set and a number is  already activated for the user. It is there to be shown to the person signing in, not to be sent back.
	PhoneNoise NullableString `json:"phoneNoise,omitempty"`
	// Whether an authenticator app is the second factor in play, with the same two readings as `sms`.
	Tfa *bool `json:"tfa,omitempty"`
	// The secret to enrol in an authenticator app, in the manual-entry form. It is filled in only while `tfa` is  set and the app has not been connected yet, which is the one moment the secret is handed out; once the app  is connected it stays empty. `GET api/2.0/settings/tfaapp/setup` returns the same secret with a QR code.
	TfaKey NullableString `json:"tfaKey,omitempty"`
	// The confirmation link the client has to open to get past the second factor. It points at phone activation  while no number is activated, at authenticator-app activation while the app is not connected, and at the  plain code prompt once either is in place. It is empty in an answer that already carries a token.
	ConfirmUrl NullableString `json:"confirmUrl,omitempty"`
}

// NewAuthenticationTokenDto instantiates a new AuthenticationTokenDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationTokenDto() *AuthenticationTokenDto {
	this := AuthenticationTokenDto{}
	return &this
}

// NewAuthenticationTokenDtoWithDefaults instantiates a new AuthenticationTokenDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationTokenDtoWithDefaults() *AuthenticationTokenDto {
	this := AuthenticationTokenDto{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationTokenDto) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationTokenDto) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *AuthenticationTokenDto) IsTokenSet() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *AuthenticationTokenDto) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *AuthenticationTokenDto) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *AuthenticationTokenDto) UnsetToken() {
	o.Token.Unset()
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *AuthenticationTokenDto) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationTokenDto) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *AuthenticationTokenDto) IsExpiresSet() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *AuthenticationTokenDto) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *AuthenticationTokenDto) GetSms() bool {
	if o == nil || IsNil(o.Sms) {
		var ret bool
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationTokenDto) GetSmsOk() (*bool, bool) {
	if o == nil || IsNil(o.Sms) {
		return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *AuthenticationTokenDto) IsSmsSet() bool {
	if o != nil && !IsNil(o.Sms) {
		return true
	}

	return false
}

// SetSms gets a reference to the given bool and assigns it to the Sms field.
func (o *AuthenticationTokenDto) SetSms(v bool) {
	o.Sms = &v
}

// GetPhoneNoise returns the PhoneNoise field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationTokenDto) GetPhoneNoise() string {
	if o == nil || IsNil(o.PhoneNoise.Get()) {
		var ret string
		return ret
	}
	return *o.PhoneNoise.Get()
}

// GetPhoneNoiseOk returns a tuple with the PhoneNoise field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationTokenDto) GetPhoneNoiseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNoise.Get(), o.PhoneNoise.IsSet()
}

// HasPhoneNoise returns a boolean if a field has been set.
func (o *AuthenticationTokenDto) IsPhoneNoiseSet() bool {
	if o != nil && o.PhoneNoise.IsSet() {
		return true
	}

	return false
}

// SetPhoneNoise gets a reference to the given NullableString and assigns it to the PhoneNoise field.
func (o *AuthenticationTokenDto) SetPhoneNoise(v string) {
	o.PhoneNoise.Set(&v)
}
// SetPhoneNoiseNil sets the value for PhoneNoise to be an explicit nil
func (o *AuthenticationTokenDto) SetPhoneNoiseNil() {
	o.PhoneNoise.Set(nil)
}

// UnsetPhoneNoise ensures that no value is present for PhoneNoise, not even an explicit nil
func (o *AuthenticationTokenDto) UnsetPhoneNoise() {
	o.PhoneNoise.Unset()
}

// GetTfa returns the Tfa field value if set, zero value otherwise.
func (o *AuthenticationTokenDto) GetTfa() bool {
	if o == nil || IsNil(o.Tfa) {
		var ret bool
		return ret
	}
	return *o.Tfa
}

// GetTfaOk returns a tuple with the Tfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationTokenDto) GetTfaOk() (*bool, bool) {
	if o == nil || IsNil(o.Tfa) {
		return nil, false
	}
	return o.Tfa, true
}

// HasTfa returns a boolean if a field has been set.
func (o *AuthenticationTokenDto) IsTfaSet() bool {
	if o != nil && !IsNil(o.Tfa) {
		return true
	}

	return false
}

// SetTfa gets a reference to the given bool and assigns it to the Tfa field.
func (o *AuthenticationTokenDto) SetTfa(v bool) {
	o.Tfa = &v
}

// GetTfaKey returns the TfaKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationTokenDto) GetTfaKey() string {
	if o == nil || IsNil(o.TfaKey.Get()) {
		var ret string
		return ret
	}
	return *o.TfaKey.Get()
}

// GetTfaKeyOk returns a tuple with the TfaKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationTokenDto) GetTfaKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TfaKey.Get(), o.TfaKey.IsSet()
}

// HasTfaKey returns a boolean if a field has been set.
func (o *AuthenticationTokenDto) IsTfaKeySet() bool {
	if o != nil && o.TfaKey.IsSet() {
		return true
	}

	return false
}

// SetTfaKey gets a reference to the given NullableString and assigns it to the TfaKey field.
func (o *AuthenticationTokenDto) SetTfaKey(v string) {
	o.TfaKey.Set(&v)
}
// SetTfaKeyNil sets the value for TfaKey to be an explicit nil
func (o *AuthenticationTokenDto) SetTfaKeyNil() {
	o.TfaKey.Set(nil)
}

// UnsetTfaKey ensures that no value is present for TfaKey, not even an explicit nil
func (o *AuthenticationTokenDto) UnsetTfaKey() {
	o.TfaKey.Unset()
}

// GetConfirmUrl returns the ConfirmUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationTokenDto) GetConfirmUrl() string {
	if o == nil || IsNil(o.ConfirmUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ConfirmUrl.Get()
}

// GetConfirmUrlOk returns a tuple with the ConfirmUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationTokenDto) GetConfirmUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfirmUrl.Get(), o.ConfirmUrl.IsSet()
}

// HasConfirmUrl returns a boolean if a field has been set.
func (o *AuthenticationTokenDto) IsConfirmUrlSet() bool {
	if o != nil && o.ConfirmUrl.IsSet() {
		return true
	}

	return false
}

// SetConfirmUrl gets a reference to the given NullableString and assigns it to the ConfirmUrl field.
func (o *AuthenticationTokenDto) SetConfirmUrl(v string) {
	o.ConfirmUrl.Set(&v)
}
// SetConfirmUrlNil sets the value for ConfirmUrl to be an explicit nil
func (o *AuthenticationTokenDto) SetConfirmUrlNil() {
	o.ConfirmUrl.Set(nil)
}

// UnsetConfirmUrl ensures that no value is present for ConfirmUrl, not even an explicit nil
func (o *AuthenticationTokenDto) UnsetConfirmUrl() {
	o.ConfirmUrl.Unset()
}

func (o AuthenticationTokenDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationTokenDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	if !IsNil(o.Sms) {
		toSerialize["sms"] = o.Sms
	}
	if o.PhoneNoise.IsSet() {
		toSerialize["phoneNoise"] = o.PhoneNoise.Get()
	}
	if !IsNil(o.Tfa) {
		toSerialize["tfa"] = o.Tfa
	}
	if o.TfaKey.IsSet() {
		toSerialize["tfaKey"] = o.TfaKey.Get()
	}
	if o.ConfirmUrl.IsSet() {
		toSerialize["confirmUrl"] = o.ConfirmUrl.Get()
	}
	return toSerialize, nil
}

type NullableAuthenticationTokenDto struct {
	value *AuthenticationTokenDto
	isSet bool
}

func (v NullableAuthenticationTokenDto) Get() *AuthenticationTokenDto {
	return v.value
}

func (v *NullableAuthenticationTokenDto) Set(val *AuthenticationTokenDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationTokenDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationTokenDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationTokenDto(val *AuthenticationTokenDto) *NullableAuthenticationTokenDto {
	return &NullableAuthenticationTokenDto{value: val, isSet: true}
}

func (v NullableAuthenticationTokenDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationTokenDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

