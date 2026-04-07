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

// checks if the AuthRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthRequestsDto{}

// AuthRequestsDto The parameters required for the user authentication requests.
type AuthRequestsDto struct {
	// The username or email used for authentication.
	UserName NullableString `json:"userName,omitempty"`
	// The password in plain text for user authentication.
	Password NullableString `json:"password,omitempty"`
	// The hashed password for secure verification.
	PasswordHash NullableString `json:"passwordHash,omitempty"`
	// The type of authentication provider (e.g., internal, Google, Azure).
	Provider NullableString `json:"provider,omitempty"`
	// The access token used for authentication with external providers.
	AccessToken NullableString `json:"accessToken,omitempty"`
	// The serialized user profile data, if applicable.
	SerializedProfile NullableString `json:"serializedProfile,omitempty"`
	// The authorization code used for obtaining OAuth tokens.
	CodeOAuth NullableString `json:"codeOAuth,omitempty"`
	// Specifies whether the authentication is session-based.
	Session *bool `json:"session,omitempty"`
	ConfirmData *ConfirmData `json:"confirmData,omitempty"`
	RecaptchaType *RecaptchaType `json:"recaptchaType,omitempty"`
	// The user's response to the CAPTCHA challenge.
	RecaptchaResponse NullableString `json:"recaptchaResponse,omitempty"`
	// The culture code for localization during authentication.
	Culture NullableString `json:"culture,omitempty"`
}

// NewAuthRequestsDto instantiates a new AuthRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRequestsDto() *AuthRequestsDto {
	this := AuthRequestsDto{}
	return &this
}

// NewAuthRequestsDtoWithDefaults instantiates a new AuthRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRequestsDtoWithDefaults() *AuthRequestsDto {
	this := AuthRequestsDto{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestsDto) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestsDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsUserNameSet() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *AuthRequestsDto) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *AuthRequestsDto) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *AuthRequestsDto) UnsetUserName() {
	o.UserName.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestsDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestsDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *AuthRequestsDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *AuthRequestsDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *AuthRequestsDto) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestsDto) GetPasswordHash() string {
	if o == nil || IsNil(o.PasswordHash.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestsDto) GetPasswordHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsPasswordHashSet() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *AuthRequestsDto) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *AuthRequestsDto) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *AuthRequestsDto) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestsDto) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestsDto) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsProviderSet() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *AuthRequestsDto) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *AuthRequestsDto) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *AuthRequestsDto) UnsetProvider() {
	o.Provider.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestsDto) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestsDto) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsAccessTokenSet() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *AuthRequestsDto) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *AuthRequestsDto) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *AuthRequestsDto) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetSerializedProfile returns the SerializedProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestsDto) GetSerializedProfile() string {
	if o == nil || IsNil(o.SerializedProfile.Get()) {
		var ret string
		return ret
	}
	return *o.SerializedProfile.Get()
}

// GetSerializedProfileOk returns a tuple with the SerializedProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestsDto) GetSerializedProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SerializedProfile.Get(), o.SerializedProfile.IsSet()
}

// HasSerializedProfile returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsSerializedProfileSet() bool {
	if o != nil && o.SerializedProfile.IsSet() {
		return true
	}

	return false
}

// SetSerializedProfile gets a reference to the given NullableString and assigns it to the SerializedProfile field.
func (o *AuthRequestsDto) SetSerializedProfile(v string) {
	o.SerializedProfile.Set(&v)
}
// SetSerializedProfileNil sets the value for SerializedProfile to be an explicit nil
func (o *AuthRequestsDto) SetSerializedProfileNil() {
	o.SerializedProfile.Set(nil)
}

// UnsetSerializedProfile ensures that no value is present for SerializedProfile, not even an explicit nil
func (o *AuthRequestsDto) UnsetSerializedProfile() {
	o.SerializedProfile.Unset()
}

// GetCodeOAuth returns the CodeOAuth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestsDto) GetCodeOAuth() string {
	if o == nil || IsNil(o.CodeOAuth.Get()) {
		var ret string
		return ret
	}
	return *o.CodeOAuth.Get()
}

// GetCodeOAuthOk returns a tuple with the CodeOAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestsDto) GetCodeOAuthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeOAuth.Get(), o.CodeOAuth.IsSet()
}

// HasCodeOAuth returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsCodeOAuthSet() bool {
	if o != nil && o.CodeOAuth.IsSet() {
		return true
	}

	return false
}

// SetCodeOAuth gets a reference to the given NullableString and assigns it to the CodeOAuth field.
func (o *AuthRequestsDto) SetCodeOAuth(v string) {
	o.CodeOAuth.Set(&v)
}
// SetCodeOAuthNil sets the value for CodeOAuth to be an explicit nil
func (o *AuthRequestsDto) SetCodeOAuthNil() {
	o.CodeOAuth.Set(nil)
}

// UnsetCodeOAuth ensures that no value is present for CodeOAuth, not even an explicit nil
func (o *AuthRequestsDto) UnsetCodeOAuth() {
	o.CodeOAuth.Unset()
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *AuthRequestsDto) GetSession() bool {
	if o == nil || IsNil(o.Session) {
		var ret bool
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestsDto) GetSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsSessionSet() bool {
	if o != nil && !IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given bool and assigns it to the Session field.
func (o *AuthRequestsDto) SetSession(v bool) {
	o.Session = &v
}

// GetConfirmData returns the ConfirmData field value if set, zero value otherwise.
func (o *AuthRequestsDto) GetConfirmData() ConfirmData {
	if o == nil || IsNil(o.ConfirmData) {
		var ret ConfirmData
		return ret
	}
	return *o.ConfirmData
}

// GetConfirmDataOk returns a tuple with the ConfirmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestsDto) GetConfirmDataOk() (*ConfirmData, bool) {
	if o == nil || IsNil(o.ConfirmData) {
		return nil, false
	}
	return o.ConfirmData, true
}

// HasConfirmData returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsConfirmDataSet() bool {
	if o != nil && !IsNil(o.ConfirmData) {
		return true
	}

	return false
}

// SetConfirmData gets a reference to the given ConfirmData and assigns it to the ConfirmData field.
func (o *AuthRequestsDto) SetConfirmData(v ConfirmData) {
	o.ConfirmData = &v
}

// GetRecaptchaType returns the RecaptchaType field value if set, zero value otherwise.
func (o *AuthRequestsDto) GetRecaptchaType() RecaptchaType {
	if o == nil || IsNil(o.RecaptchaType) {
		var ret RecaptchaType
		return ret
	}
	return *o.RecaptchaType
}

// GetRecaptchaTypeOk returns a tuple with the RecaptchaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestsDto) GetRecaptchaTypeOk() (*RecaptchaType, bool) {
	if o == nil || IsNil(o.RecaptchaType) {
		return nil, false
	}
	return o.RecaptchaType, true
}

// HasRecaptchaType returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsRecaptchaTypeSet() bool {
	if o != nil && !IsNil(o.RecaptchaType) {
		return true
	}

	return false
}

// SetRecaptchaType gets a reference to the given RecaptchaType and assigns it to the RecaptchaType field.
func (o *AuthRequestsDto) SetRecaptchaType(v RecaptchaType) {
	o.RecaptchaType = &v
}

// GetRecaptchaResponse returns the RecaptchaResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestsDto) GetRecaptchaResponse() string {
	if o == nil || IsNil(o.RecaptchaResponse.Get()) {
		var ret string
		return ret
	}
	return *o.RecaptchaResponse.Get()
}

// GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestsDto) GetRecaptchaResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecaptchaResponse.Get(), o.RecaptchaResponse.IsSet()
}

// HasRecaptchaResponse returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsRecaptchaResponseSet() bool {
	if o != nil && o.RecaptchaResponse.IsSet() {
		return true
	}

	return false
}

// SetRecaptchaResponse gets a reference to the given NullableString and assigns it to the RecaptchaResponse field.
func (o *AuthRequestsDto) SetRecaptchaResponse(v string) {
	o.RecaptchaResponse.Set(&v)
}
// SetRecaptchaResponseNil sets the value for RecaptchaResponse to be an explicit nil
func (o *AuthRequestsDto) SetRecaptchaResponseNil() {
	o.RecaptchaResponse.Set(nil)
}

// UnsetRecaptchaResponse ensures that no value is present for RecaptchaResponse, not even an explicit nil
func (o *AuthRequestsDto) UnsetRecaptchaResponse() {
	o.RecaptchaResponse.Unset()
}

// GetCulture returns the Culture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestsDto) GetCulture() string {
	if o == nil || IsNil(o.Culture.Get()) {
		var ret string
		return ret
	}
	return *o.Culture.Get()
}

// GetCultureOk returns a tuple with the Culture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestsDto) GetCultureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Culture.Get(), o.Culture.IsSet()
}

// HasCulture returns a boolean if a field has been set.
func (o *AuthRequestsDto) IsCultureSet() bool {
	if o != nil && o.Culture.IsSet() {
		return true
	}

	return false
}

// SetCulture gets a reference to the given NullableString and assigns it to the Culture field.
func (o *AuthRequestsDto) SetCulture(v string) {
	o.Culture.Set(&v)
}
// SetCultureNil sets the value for Culture to be an explicit nil
func (o *AuthRequestsDto) SetCultureNil() {
	o.Culture.Set(nil)
}

// UnsetCulture ensures that no value is present for Culture, not even an explicit nil
func (o *AuthRequestsDto) UnsetCulture() {
	o.Culture.Unset()
}

func (o AuthRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.PasswordHash.IsSet() {
		toSerialize["passwordHash"] = o.PasswordHash.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if o.AccessToken.IsSet() {
		toSerialize["accessToken"] = o.AccessToken.Get()
	}
	if o.SerializedProfile.IsSet() {
		toSerialize["serializedProfile"] = o.SerializedProfile.Get()
	}
	if o.CodeOAuth.IsSet() {
		toSerialize["codeOAuth"] = o.CodeOAuth.Get()
	}
	if !IsNil(o.Session) {
		toSerialize["session"] = o.Session
	}
	if !IsNil(o.ConfirmData) {
		toSerialize["confirmData"] = o.ConfirmData
	}
	if !IsNil(o.RecaptchaType) {
		toSerialize["recaptchaType"] = o.RecaptchaType
	}
	if o.RecaptchaResponse.IsSet() {
		toSerialize["recaptchaResponse"] = o.RecaptchaResponse.Get()
	}
	if o.Culture.IsSet() {
		toSerialize["culture"] = o.Culture.Get()
	}
	return toSerialize, nil
}

type NullableAuthRequestsDto struct {
	value *AuthRequestsDto
	isSet bool
}

func (v NullableAuthRequestsDto) Get() *AuthRequestsDto {
	return v.value
}

func (v *NullableAuthRequestsDto) Set(val *AuthRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRequestsDto(val *AuthRequestsDto) *NullableAuthRequestsDto {
	return &NullableAuthRequestsDto{value: val, isSet: true}
}

func (v NullableAuthRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

