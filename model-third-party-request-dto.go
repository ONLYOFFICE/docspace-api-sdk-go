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

// checks if the ThirdPartyRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyRequestDto{}

// ThirdPartyRequestDto The third-party request parameters.
type ThirdPartyRequestDto struct {
	// The connection URL for the sharepoint.
	Url NullableString `json:"url,omitempty"`
	// The third-party request login.
	Login NullableString `json:"login,omitempty"`
	// The third-party request password.
	Password NullableString `json:"password,omitempty"`
	// The authentication token.
	Token NullableString `json:"token,omitempty"`
	// The customer title.
	CustomerTitle NullableString `json:"customerTitle"`
	// The provider key.
	ProviderKey NullableString `json:"providerKey"`
	// The provider ID.
	ProviderId NullableInt32 `json:"providerId,omitempty"`
}

type _ThirdPartyRequestDto ThirdPartyRequestDto

// NewThirdPartyRequestDto instantiates a new ThirdPartyRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyRequestDto(customerTitle NullableString, providerKey NullableString) *ThirdPartyRequestDto {
	this := ThirdPartyRequestDto{}
	this.CustomerTitle = customerTitle
	this.ProviderKey = providerKey
	return &this
}

// NewThirdPartyRequestDtoWithDefaults instantiates a new ThirdPartyRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyRequestDtoWithDefaults() *ThirdPartyRequestDto {
	this := ThirdPartyRequestDto{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyRequestDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyRequestDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ThirdPartyRequestDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ThirdPartyRequestDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ThirdPartyRequestDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ThirdPartyRequestDto) UnsetUrl() {
	o.Url.Unset()
}

// GetLogin returns the Login field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyRequestDto) GetLogin() string {
	if o == nil || IsNil(o.Login.Get()) {
		var ret string
		return ret
	}
	return *o.Login.Get()
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyRequestDto) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Login.Get(), o.Login.IsSet()
}

// HasLogin returns a boolean if a field has been set.
func (o *ThirdPartyRequestDto) IsLoginSet() bool {
	if o != nil && o.Login.IsSet() {
		return true
	}

	return false
}

// SetLogin gets a reference to the given NullableString and assigns it to the Login field.
func (o *ThirdPartyRequestDto) SetLogin(v string) {
	o.Login.Set(&v)
}
// SetLoginNil sets the value for Login to be an explicit nil
func (o *ThirdPartyRequestDto) SetLoginNil() {
	o.Login.Set(nil)
}

// UnsetLogin ensures that no value is present for Login, not even an explicit nil
func (o *ThirdPartyRequestDto) UnsetLogin() {
	o.Login.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyRequestDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyRequestDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *ThirdPartyRequestDto) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *ThirdPartyRequestDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *ThirdPartyRequestDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *ThirdPartyRequestDto) UnsetPassword() {
	o.Password.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyRequestDto) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyRequestDto) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *ThirdPartyRequestDto) IsTokenSet() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *ThirdPartyRequestDto) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *ThirdPartyRequestDto) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *ThirdPartyRequestDto) UnsetToken() {
	o.Token.Unset()
}

// GetCustomerTitle returns the CustomerTitle field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThirdPartyRequestDto) GetCustomerTitle() string {
	if o == nil || o.CustomerTitle.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomerTitle.Get()
}

// GetCustomerTitleOk returns a tuple with the CustomerTitle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyRequestDto) GetCustomerTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerTitle.Get(), o.CustomerTitle.IsSet()
}

// SetCustomerTitle sets field value
func (o *ThirdPartyRequestDto) SetCustomerTitle(v string) {
	o.CustomerTitle.Set(&v)
}

// GetProviderKey returns the ProviderKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThirdPartyRequestDto) GetProviderKey() string {
	if o == nil || o.ProviderKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProviderKey.Get()
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyRequestDto) GetProviderKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderKey.Get(), o.ProviderKey.IsSet()
}

// SetProviderKey sets field value
func (o *ThirdPartyRequestDto) SetProviderKey(v string) {
	o.ProviderKey.Set(&v)
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyRequestDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProviderId.Get()
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyRequestDto) GetProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderId.Get(), o.ProviderId.IsSet()
}

// HasProviderId returns a boolean if a field has been set.
func (o *ThirdPartyRequestDto) IsProviderIdSet() bool {
	if o != nil && o.ProviderId.IsSet() {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given NullableInt32 and assigns it to the ProviderId field.
func (o *ThirdPartyRequestDto) SetProviderId(v int32) {
	o.ProviderId.Set(&v)
}
// SetProviderIdNil sets the value for ProviderId to be an explicit nil
func (o *ThirdPartyRequestDto) SetProviderIdNil() {
	o.ProviderId.Set(nil)
}

// UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
func (o *ThirdPartyRequestDto) UnsetProviderId() {
	o.ProviderId.Unset()
}

func (o ThirdPartyRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyRequestDto) ToMap() (map[string]interface{}, error) {
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
	toSerialize["customerTitle"] = o.CustomerTitle.Get()
	toSerialize["providerKey"] = o.ProviderKey.Get()
	if o.ProviderId.IsSet() {
		toSerialize["providerId"] = o.ProviderId.Get()
	}
	return toSerialize, nil
}

func (o *ThirdPartyRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerTitle",
		"providerKey",
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

	varThirdPartyRequestDto := _ThirdPartyRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varThirdPartyRequestDto)

	if err != nil {
		return err
	}

	*o = ThirdPartyRequestDto(varThirdPartyRequestDto)

	return err
}

type NullableThirdPartyRequestDto struct {
	value *ThirdPartyRequestDto
	isSet bool
}

func (v NullableThirdPartyRequestDto) Get() *ThirdPartyRequestDto {
	return v.value
}

func (v *NullableThirdPartyRequestDto) Set(val *ThirdPartyRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyRequestDto(val *ThirdPartyRequestDto) *NullableThirdPartyRequestDto {
	return &NullableThirdPartyRequestDto{value: val, isSet: true}
}

func (v NullableThirdPartyRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

