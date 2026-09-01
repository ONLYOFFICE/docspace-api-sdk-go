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

// checks if the ThirdPartyBackupRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyBackupRequestDto{}

// ThirdPartyBackupRequestDto The third-party backup request parameters.
type ThirdPartyBackupRequestDto struct {
	// The connection URL for the sharepoint.
	Url NullableString `json:"url,omitempty"`
	// The login.
	Login NullableString `json:"login,omitempty"`
	// The password.
	Password NullableString `json:"password,omitempty"`
	// The authentication token.
	Token NullableString `json:"token,omitempty"`
	// The customer title.
	CustomerTitle NullableString `json:"customerTitle,omitempty"`
	// The provider key.
	ProviderKey NullableString `json:"providerKey,omitempty"`
}

// NewThirdPartyBackupRequestDto instantiates a new ThirdPartyBackupRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyBackupRequestDto() *ThirdPartyBackupRequestDto {
	this := ThirdPartyBackupRequestDto{}
	return &this
}

// NewThirdPartyBackupRequestDtoWithDefaults instantiates a new ThirdPartyBackupRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyBackupRequestDtoWithDefaults() *ThirdPartyBackupRequestDto {
	this := ThirdPartyBackupRequestDto{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyBackupRequestDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyBackupRequestDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ThirdPartyBackupRequestDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ThirdPartyBackupRequestDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ThirdPartyBackupRequestDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ThirdPartyBackupRequestDto) UnsetUrl() {
	o.Url.Unset()
}

// GetLogin returns the Login field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyBackupRequestDto) GetLogin() string {
	if o == nil || IsNil(o.Login.Get()) {
		var ret string
		return ret
	}
	return *o.Login.Get()
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyBackupRequestDto) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Login.Get(), o.Login.IsSet()
}

// HasLogin returns a boolean if a field has been set.
func (o *ThirdPartyBackupRequestDto) IsLoginSet() bool {
	if o != nil && o.Login.IsSet() {
		return true
	}

	return false
}

// SetLogin gets a reference to the given NullableString and assigns it to the Login field.
func (o *ThirdPartyBackupRequestDto) SetLogin(v string) {
	o.Login.Set(&v)
}
// SetLoginNil sets the value for Login to be an explicit nil
func (o *ThirdPartyBackupRequestDto) SetLoginNil() {
	o.Login.Set(nil)
}

// UnsetLogin ensures that no value is present for Login, not even an explicit nil
func (o *ThirdPartyBackupRequestDto) UnsetLogin() {
	o.Login.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyBackupRequestDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyBackupRequestDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *ThirdPartyBackupRequestDto) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *ThirdPartyBackupRequestDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *ThirdPartyBackupRequestDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *ThirdPartyBackupRequestDto) UnsetPassword() {
	o.Password.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyBackupRequestDto) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyBackupRequestDto) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *ThirdPartyBackupRequestDto) IsTokenSet() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *ThirdPartyBackupRequestDto) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *ThirdPartyBackupRequestDto) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *ThirdPartyBackupRequestDto) UnsetToken() {
	o.Token.Unset()
}

// GetCustomerTitle returns the CustomerTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyBackupRequestDto) GetCustomerTitle() string {
	if o == nil || IsNil(o.CustomerTitle.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerTitle.Get()
}

// GetCustomerTitleOk returns a tuple with the CustomerTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyBackupRequestDto) GetCustomerTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerTitle.Get(), o.CustomerTitle.IsSet()
}

// HasCustomerTitle returns a boolean if a field has been set.
func (o *ThirdPartyBackupRequestDto) IsCustomerTitleSet() bool {
	if o != nil && o.CustomerTitle.IsSet() {
		return true
	}

	return false
}

// SetCustomerTitle gets a reference to the given NullableString and assigns it to the CustomerTitle field.
func (o *ThirdPartyBackupRequestDto) SetCustomerTitle(v string) {
	o.CustomerTitle.Set(&v)
}
// SetCustomerTitleNil sets the value for CustomerTitle to be an explicit nil
func (o *ThirdPartyBackupRequestDto) SetCustomerTitleNil() {
	o.CustomerTitle.Set(nil)
}

// UnsetCustomerTitle ensures that no value is present for CustomerTitle, not even an explicit nil
func (o *ThirdPartyBackupRequestDto) UnsetCustomerTitle() {
	o.CustomerTitle.Unset()
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyBackupRequestDto) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderKey.Get()
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyBackupRequestDto) GetProviderKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderKey.Get(), o.ProviderKey.IsSet()
}

// HasProviderKey returns a boolean if a field has been set.
func (o *ThirdPartyBackupRequestDto) IsProviderKeySet() bool {
	if o != nil && o.ProviderKey.IsSet() {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given NullableString and assigns it to the ProviderKey field.
func (o *ThirdPartyBackupRequestDto) SetProviderKey(v string) {
	o.ProviderKey.Set(&v)
}
// SetProviderKeyNil sets the value for ProviderKey to be an explicit nil
func (o *ThirdPartyBackupRequestDto) SetProviderKeyNil() {
	o.ProviderKey.Set(nil)
}

// UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
func (o *ThirdPartyBackupRequestDto) UnsetProviderKey() {
	o.ProviderKey.Unset()
}

func (o ThirdPartyBackupRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyBackupRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Login.IsSet() {
		toSerialize["login"] = o.Login.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.CustomerTitle.IsSet() {
		toSerialize["customerTitle"] = o.CustomerTitle.Get()
	}
	if o.ProviderKey.IsSet() {
		toSerialize["providerKey"] = o.ProviderKey.Get()
	}
	return toSerialize, nil
}

type NullableThirdPartyBackupRequestDto struct {
	value *ThirdPartyBackupRequestDto
	isSet bool
}

func (v NullableThirdPartyBackupRequestDto) Get() *ThirdPartyBackupRequestDto {
	return v.value
}

func (v *NullableThirdPartyBackupRequestDto) Set(val *ThirdPartyBackupRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyBackupRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyBackupRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyBackupRequestDto(val *ThirdPartyBackupRequestDto) *NullableThirdPartyBackupRequestDto {
	return &NullableThirdPartyBackupRequestDto{value: val, isSet: true}
}

func (v NullableThirdPartyBackupRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyBackupRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

