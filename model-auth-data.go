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

// checks if the AuthData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthData{}

// AuthData The credentials of a third-party storage account. The portal takes them when an account is connected and does not  give them back afterwards.
type AuthData struct {
	// The account name at the storage service.
	Login NullableString `json:"login,omitempty"`
	// The password of the account at the storage service.
	Password NullableString `json:"password,omitempty"`
	// The token of the account, kept as the raw JSON document the storage service issued it in.
	RawToken NullableString `json:"rawToken,omitempty"`
	// The address of the storage server the account lives on.
	Url NullableString `json:"url,omitempty"`
	// The storage service the credentials belong to, as the provider key the account was connected with.
	Provider NullableString `json:"provider,omitempty"`
	// The same token as in `rawToken`, parsed into its OAuth 2.0 fields.
	Token *OAuth20Token `json:"token,omitempty"`
}

// NewAuthData instantiates a new AuthData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthData() *AuthData {
	this := AuthData{}
	return &this
}

// NewAuthDataWithDefaults instantiates a new AuthData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthDataWithDefaults() *AuthData {
	this := AuthData{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthData) GetLogin() string {
	if o == nil || IsNil(o.Login.Get()) {
		var ret string
		return ret
	}
	return *o.Login.Get()
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthData) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Login.Get(), o.Login.IsSet()
}

// HasLogin returns a boolean if a field has been set.
func (o *AuthData) IsLoginSet() bool {
	if o != nil && o.Login.IsSet() {
		return true
	}

	return false
}

// SetLogin gets a reference to the given NullableString and assigns it to the Login field.
func (o *AuthData) SetLogin(v string) {
	o.Login.Set(&v)
}
// SetLoginNil sets the value for Login to be an explicit nil
func (o *AuthData) SetLoginNil() {
	o.Login.Set(nil)
}

// UnsetLogin ensures that no value is present for Login, not even an explicit nil
func (o *AuthData) UnsetLogin() {
	o.Login.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthData) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthData) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *AuthData) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *AuthData) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *AuthData) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *AuthData) UnsetPassword() {
	o.Password.Unset()
}

// GetRawToken returns the RawToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthData) GetRawToken() string {
	if o == nil || IsNil(o.RawToken.Get()) {
		var ret string
		return ret
	}
	return *o.RawToken.Get()
}

// GetRawTokenOk returns a tuple with the RawToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthData) GetRawTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawToken.Get(), o.RawToken.IsSet()
}

// HasRawToken returns a boolean if a field has been set.
func (o *AuthData) IsRawTokenSet() bool {
	if o != nil && o.RawToken.IsSet() {
		return true
	}

	return false
}

// SetRawToken gets a reference to the given NullableString and assigns it to the RawToken field.
func (o *AuthData) SetRawToken(v string) {
	o.RawToken.Set(&v)
}
// SetRawTokenNil sets the value for RawToken to be an explicit nil
func (o *AuthData) SetRawTokenNil() {
	o.RawToken.Set(nil)
}

// UnsetRawToken ensures that no value is present for RawToken, not even an explicit nil
func (o *AuthData) UnsetRawToken() {
	o.RawToken.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthData) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthData) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *AuthData) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *AuthData) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *AuthData) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *AuthData) UnsetUrl() {
	o.Url.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthData) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthData) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *AuthData) IsProviderSet() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *AuthData) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *AuthData) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *AuthData) UnsetProvider() {
	o.Provider.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthData) GetToken() OAuth20Token {
	if o == nil || IsNil(o.Token) {
		var ret OAuth20Token
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthData) GetTokenOk() (*OAuth20Token, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthData) IsTokenSet() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given OAuth20Token and assigns it to the Token field.
func (o *AuthData) SetToken(v OAuth20Token) {
	o.Token = &v
}

func (o AuthData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Login.IsSet() {
		toSerialize["login"] = o.Login.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.RawToken.IsSet() {
		toSerialize["rawToken"] = o.RawToken.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableAuthData struct {
	value *AuthData
	isSet bool
}

func (v NullableAuthData) Get() *AuthData {
	return v.value
}

func (v *NullableAuthData) Set(val *AuthData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthData(val *AuthData) *NullableAuthData {
	return &NullableAuthData{value: val, isSet: true}
}

func (v NullableAuthData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

