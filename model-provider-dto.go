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

// checks if the ProviderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderDto{}

// ProviderDto The provider information.
type ProviderDto struct {
	// The provider name.
	Name NullableString `json:"name,omitempty"`
	// The provider key.
	Key NullableString `json:"key,omitempty"`
	// Specifies whether the provider is connected.
	Connected *bool `json:"connected,omitempty"`
	// Specifies if the provider is OAuth.
	Oauth *bool `json:"oauth,omitempty"`
	// The provider redirect URL.
	RedirectUrl NullableString `json:"redirectUrl,omitempty"`
	// The required connection URL flag.
	RequiredConnectionUrl *bool `json:"requiredConnectionUrl,omitempty"`
	// The provider OAuth client ID.
	ClientId NullableString `json:"clientId,omitempty"`
}

// NewProviderDto instantiates a new ProviderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderDto() *ProviderDto {
	this := ProviderDto{}
	return &this
}

// NewProviderDtoWithDefaults instantiates a new ProviderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderDtoWithDefaults() *ProviderDto {
	this := ProviderDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProviderDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProviderDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProviderDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProviderDto) UnsetName() {
	o.Name.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ProviderDto) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ProviderDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *ProviderDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ProviderDto) UnsetKey() {
	o.Key.Unset()
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *ProviderDto) GetConnected() bool {
	if o == nil || IsNil(o.Connected) {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderDto) GetConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *ProviderDto) IsConnectedSet() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *ProviderDto) SetConnected(v bool) {
	o.Connected = &v
}

// GetOauth returns the Oauth field value if set, zero value otherwise.
func (o *ProviderDto) GetOauth() bool {
	if o == nil || IsNil(o.Oauth) {
		var ret bool
		return ret
	}
	return *o.Oauth
}

// GetOauthOk returns a tuple with the Oauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderDto) GetOauthOk() (*bool, bool) {
	if o == nil || IsNil(o.Oauth) {
		return nil, false
	}
	return o.Oauth, true
}

// HasOauth returns a boolean if a field has been set.
func (o *ProviderDto) IsOauthSet() bool {
	if o != nil && !IsNil(o.Oauth) {
		return true
	}

	return false
}

// SetOauth gets a reference to the given bool and assigns it to the Oauth field.
func (o *ProviderDto) SetOauth(v bool) {
	o.Oauth = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderDto) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderDto) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *ProviderDto) IsRedirectUrlSet() bool {
	if o != nil && o.RedirectUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given NullableString and assigns it to the RedirectUrl field.
func (o *ProviderDto) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}
// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil
func (o *ProviderDto) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
func (o *ProviderDto) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

// GetRequiredConnectionUrl returns the RequiredConnectionUrl field value if set, zero value otherwise.
func (o *ProviderDto) GetRequiredConnectionUrl() bool {
	if o == nil || IsNil(o.RequiredConnectionUrl) {
		var ret bool
		return ret
	}
	return *o.RequiredConnectionUrl
}

// GetRequiredConnectionUrlOk returns a tuple with the RequiredConnectionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderDto) GetRequiredConnectionUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiredConnectionUrl) {
		return nil, false
	}
	return o.RequiredConnectionUrl, true
}

// HasRequiredConnectionUrl returns a boolean if a field has been set.
func (o *ProviderDto) IsRequiredConnectionUrlSet() bool {
	if o != nil && !IsNil(o.RequiredConnectionUrl) {
		return true
	}

	return false
}

// SetRequiredConnectionUrl gets a reference to the given bool and assigns it to the RequiredConnectionUrl field.
func (o *ProviderDto) SetRequiredConnectionUrl(v bool) {
	o.RequiredConnectionUrl = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderDto) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderDto) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *ProviderDto) IsClientIdSet() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *ProviderDto) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *ProviderDto) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *ProviderDto) UnsetClientId() {
	o.ClientId.Unset()
}

func (o ProviderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !IsNil(o.Connected) {
		toSerialize["connected"] = o.Connected
	}
	if !IsNil(o.Oauth) {
		toSerialize["oauth"] = o.Oauth
	}
	if o.RedirectUrl.IsSet() {
		toSerialize["redirectUrl"] = o.RedirectUrl.Get()
	}
	if !IsNil(o.RequiredConnectionUrl) {
		toSerialize["requiredConnectionUrl"] = o.RequiredConnectionUrl
	}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	return toSerialize, nil
}

type NullableProviderDto struct {
	value *ProviderDto
	isSet bool
}

func (v NullableProviderDto) Get() *ProviderDto {
	return v.value
}

func (v *NullableProviderDto) Set(val *ProviderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderDto(val *ProviderDto) *NullableProviderDto {
	return &NullableProviderDto{value: val, isSet: true}
}

func (v NullableProviderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

