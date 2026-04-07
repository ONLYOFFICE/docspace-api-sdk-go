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

// checks if the ExchangeToken200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeToken200Response{}

// ExchangeToken200Response struct for ExchangeToken200Response
type ExchangeToken200Response struct {
	// The access token issued by the authorization server.
	AccessToken *string `json:"access_token,omitempty"`
	// The type of token issued, typically 'Bearer'.
	TokenType *string `json:"token_type,omitempty"`
	// The number of seconds until the access token expires.
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	// The token used to obtain a new access token when the current one expires.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

// NewExchangeToken200Response instantiates a new ExchangeToken200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeToken200Response() *ExchangeToken200Response {
	this := ExchangeToken200Response{}
	return &this
}

// NewExchangeToken200ResponseWithDefaults instantiates a new ExchangeToken200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeToken200ResponseWithDefaults() *ExchangeToken200Response {
	this := ExchangeToken200Response{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *ExchangeToken200Response) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeToken200Response) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *ExchangeToken200Response) IsAccessTokenSet() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *ExchangeToken200Response) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *ExchangeToken200Response) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeToken200Response) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *ExchangeToken200Response) IsTokenTypeSet() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *ExchangeToken200Response) SetTokenType(v string) {
	o.TokenType = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *ExchangeToken200Response) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeToken200Response) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *ExchangeToken200Response) IsExpiresInSet() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *ExchangeToken200Response) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *ExchangeToken200Response) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeToken200Response) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *ExchangeToken200Response) IsRefreshTokenSet() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *ExchangeToken200Response) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o ExchangeToken200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeToken200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.TokenType) {
		toSerialize["token_type"] = o.TokenType
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	return toSerialize, nil
}

type NullableExchangeToken200Response struct {
	value *ExchangeToken200Response
	isSet bool
}

func (v NullableExchangeToken200Response) Get() *ExchangeToken200Response {
	return v.value
}

func (v *NullableExchangeToken200Response) Set(val *ExchangeToken200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeToken200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeToken200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeToken200Response(val *ExchangeToken200Response) *NullableExchangeToken200Response {
	return &NullableExchangeToken200Response{value: val, isSet: true}
}

func (v NullableExchangeToken200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeToken200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

