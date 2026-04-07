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

// checks if the OAuth20Token type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth20Token{}

// OAuth20Token struct for OAuth20Token
type OAuth20Token struct {
	AccessToken NullableString `json:"access_token,omitempty"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	ClientId NullableString `json:"client_id,omitempty"`
	ClientSecret NullableString `json:"client_secret,omitempty"`
	RedirectUri NullableString `json:"redirect_uri,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	IsExpired *bool `json:"isExpired,omitempty"`
}

// NewOAuth20Token instantiates a new OAuth20Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth20Token() *OAuth20Token {
	this := OAuth20Token{}
	return &this
}

// NewOAuth20TokenWithDefaults instantiates a new OAuth20Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth20TokenWithDefaults() *OAuth20Token {
	this := OAuth20Token{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth20Token) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth20Token) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *OAuth20Token) IsAccessTokenSet() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *OAuth20Token) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *OAuth20Token) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *OAuth20Token) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth20Token) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth20Token) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OAuth20Token) IsRefreshTokenSet() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *OAuth20Token) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *OAuth20Token) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *OAuth20Token) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *OAuth20Token) GetExpiresIn() int64 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int64
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth20Token) GetExpiresInOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *OAuth20Token) IsExpiresInSet() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int64 and assigns it to the ExpiresIn field.
func (o *OAuth20Token) SetExpiresIn(v int64) {
	o.ExpiresIn = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth20Token) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth20Token) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuth20Token) IsClientIdSet() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *OAuth20Token) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *OAuth20Token) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *OAuth20Token) UnsetClientId() {
	o.ClientId.Unset()
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth20Token) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth20Token) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAuth20Token) IsClientSecretSet() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *OAuth20Token) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}
// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *OAuth20Token) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *OAuth20Token) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth20Token) GetRedirectUri() string {
	if o == nil || IsNil(o.RedirectUri.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUri.Get()
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth20Token) GetRedirectUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUri.Get(), o.RedirectUri.IsSet()
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *OAuth20Token) IsRedirectUriSet() bool {
	if o != nil && o.RedirectUri.IsSet() {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given NullableString and assigns it to the RedirectUri field.
func (o *OAuth20Token) SetRedirectUri(v string) {
	o.RedirectUri.Set(&v)
}
// SetRedirectUriNil sets the value for RedirectUri to be an explicit nil
func (o *OAuth20Token) SetRedirectUriNil() {
	o.RedirectUri.Set(nil)
}

// UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil
func (o *OAuth20Token) UnsetRedirectUri() {
	o.RedirectUri.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OAuth20Token) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth20Token) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OAuth20Token) IsTimestampSet() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *OAuth20Token) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise.
func (o *OAuth20Token) GetIsExpired() bool {
	if o == nil || IsNil(o.IsExpired) {
		var ret bool
		return ret
	}
	return *o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth20Token) GetIsExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExpired) {
		return nil, false
	}
	return o.IsExpired, true
}

// HasIsExpired returns a boolean if a field has been set.
func (o *OAuth20Token) IsIsExpiredSet() bool {
	if o != nil && !IsNil(o.IsExpired) {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given bool and assigns it to the IsExpired field.
func (o *OAuth20Token) SetIsExpired(v bool) {
	o.IsExpired = &v
}

func (o OAuth20Token) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth20Token) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	if o.RefreshToken.IsSet() {
		toSerialize["refresh_token"] = o.RefreshToken.Get()
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.ClientId.IsSet() {
		toSerialize["client_id"] = o.ClientId.Get()
	}
	if o.ClientSecret.IsSet() {
		toSerialize["client_secret"] = o.ClientSecret.Get()
	}
	if o.RedirectUri.IsSet() {
		toSerialize["redirect_uri"] = o.RedirectUri.Get()
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.IsExpired) {
		toSerialize["isExpired"] = o.IsExpired
	}
	return toSerialize, nil
}

type NullableOAuth20Token struct {
	value *OAuth20Token
	isSet bool
}

func (v NullableOAuth20Token) Get() *OAuth20Token {
	return v.value
}

func (v *NullableOAuth20Token) Set(val *OAuth20Token) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth20Token) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth20Token) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth20Token(val *OAuth20Token) *NullableOAuth20Token {
	return &NullableOAuth20Token{value: val, isSet: true}
}

func (v NullableOAuth20Token) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth20Token) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

