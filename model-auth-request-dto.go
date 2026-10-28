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

// checks if the AuthRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthRequestDto{}

// AuthRequestDto The credentials a sign-in is attempted with: a portal password, a confirmation key, or a third-party account.
type AuthRequestDto struct {
	// The account signing in, given as its email address or its portal user name. It is required for a password  sign-in and ignored when the credentials are a confirmation key or a third-party account.
	UserName NullableString `json:"userName,omitempty"`
	// The password in the clear. Send either this or `passwordHash`, never both; hashing it in the client with the  parameters from `GET api/2.0/settings?withpassword=true` and sending `passwordHash` instead keeps the plain  password off the wire.
	Password NullableString `json:"password,omitempty"`
	// The password already hashed in the client. It has to be produced with the `salt`, iteration count and hash  size that `GET api/2.0/settings?withpassword=true` publishes, or the portal cannot recognise it; a value sent  here takes the place of `password`.
	PasswordHash NullableString `json:"passwordHash,omitempty"`
	// The third-party identity provider the account is being signed in through, by its internal key such as  `google` or `linkedin`. Sending it switches the call to a third-party sign-in, which needs `accessToken` or  `serializedProfile` and is only allowed on a self-hosted installation or a tariff that includes third-party  sign-in.
	Provider NullableString `json:"provider,omitempty"`
	// The access token the provider named in `provider` issued for the account, passed on unchanged for the portal  to verify with that provider. The portal then matches the address it gets back against its own accounts, so a  valid token for an address unknown here is answered as no such user.
	AccessToken NullableString `json:"accessToken,omitempty"`
	// The third-party profile already fetched and serialised by the caller, as an alternative to `accessToken` for  a provider whose profile the client holds. It identifies the account by the address it carries.
	SerializedProfile NullableString `json:"serializedProfile,omitempty"`
	// The OAuth authorization code obtained from the provider, for a flow that has not been exchanged for an access  token yet. It is recorded with the sign-in rather than replacing `accessToken`.
	CodeOAuth NullableString `json:"codeOAuth,omitempty"`
	// Whether the issued token is tied to the browser session. When it is, the answer carries no `expires` and the  token dies with the session; otherwise it lives for the portal session lifetime.
	Session *bool `json:"session,omitempty"`
	// The confirmation link data, as a third way to identify the account beside a password and a third-party  account. Send it when the sign-in comes from a link the portal mailed, in which case `userName` and the  password fields are not read.
	ConfirmData *ConfirmData `json:"confirmData,omitempty"`
	// Which CAPTCHA service the proof in `recaptchaResponse` came from. It has to match the service the  installation is configured with, which `GET api/2.0/settings` publishes together with the site key.
	RecaptchaType *RecaptchaType `json:"recaptchaType,omitempty"`
	// The token the CAPTCHA widget produced in the browser, passed on unchanged for the portal to verify. It is  only demanded once repeated failures have made the portal ask for a challenge, and it is single-use, so a  retry needs a freshly solved one.
	RecaptchaResponse NullableString `json:"recaptchaResponse,omitempty"`
	// The language the sign-in messages and any letter that follows are written in, as a culture name such as  `en-US`. A culture the installation does not have falls back to the portal language.
	Culture NullableString `json:"culture,omitempty"`
}

// NewAuthRequestDto instantiates a new AuthRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRequestDto() *AuthRequestDto {
	this := AuthRequestDto{}
	return &this
}

// NewAuthRequestDtoWithDefaults instantiates a new AuthRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRequestDtoWithDefaults() *AuthRequestDto {
	this := AuthRequestDto{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestDto) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *AuthRequestDto) IsUserNameSet() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *AuthRequestDto) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *AuthRequestDto) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *AuthRequestDto) UnsetUserName() {
	o.UserName.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *AuthRequestDto) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *AuthRequestDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *AuthRequestDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *AuthRequestDto) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestDto) GetPasswordHash() string {
	if o == nil || IsNil(o.PasswordHash.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestDto) GetPasswordHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *AuthRequestDto) IsPasswordHashSet() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *AuthRequestDto) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *AuthRequestDto) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *AuthRequestDto) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestDto) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestDto) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *AuthRequestDto) IsProviderSet() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *AuthRequestDto) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *AuthRequestDto) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *AuthRequestDto) UnsetProvider() {
	o.Provider.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestDto) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestDto) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthRequestDto) IsAccessTokenSet() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *AuthRequestDto) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *AuthRequestDto) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *AuthRequestDto) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetSerializedProfile returns the SerializedProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestDto) GetSerializedProfile() string {
	if o == nil || IsNil(o.SerializedProfile.Get()) {
		var ret string
		return ret
	}
	return *o.SerializedProfile.Get()
}

// GetSerializedProfileOk returns a tuple with the SerializedProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestDto) GetSerializedProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SerializedProfile.Get(), o.SerializedProfile.IsSet()
}

// HasSerializedProfile returns a boolean if a field has been set.
func (o *AuthRequestDto) IsSerializedProfileSet() bool {
	if o != nil && o.SerializedProfile.IsSet() {
		return true
	}

	return false
}

// SetSerializedProfile gets a reference to the given NullableString and assigns it to the SerializedProfile field.
func (o *AuthRequestDto) SetSerializedProfile(v string) {
	o.SerializedProfile.Set(&v)
}
// SetSerializedProfileNil sets the value for SerializedProfile to be an explicit nil
func (o *AuthRequestDto) SetSerializedProfileNil() {
	o.SerializedProfile.Set(nil)
}

// UnsetSerializedProfile ensures that no value is present for SerializedProfile, not even an explicit nil
func (o *AuthRequestDto) UnsetSerializedProfile() {
	o.SerializedProfile.Unset()
}

// GetCodeOAuth returns the CodeOAuth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestDto) GetCodeOAuth() string {
	if o == nil || IsNil(o.CodeOAuth.Get()) {
		var ret string
		return ret
	}
	return *o.CodeOAuth.Get()
}

// GetCodeOAuthOk returns a tuple with the CodeOAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestDto) GetCodeOAuthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeOAuth.Get(), o.CodeOAuth.IsSet()
}

// HasCodeOAuth returns a boolean if a field has been set.
func (o *AuthRequestDto) IsCodeOAuthSet() bool {
	if o != nil && o.CodeOAuth.IsSet() {
		return true
	}

	return false
}

// SetCodeOAuth gets a reference to the given NullableString and assigns it to the CodeOAuth field.
func (o *AuthRequestDto) SetCodeOAuth(v string) {
	o.CodeOAuth.Set(&v)
}
// SetCodeOAuthNil sets the value for CodeOAuth to be an explicit nil
func (o *AuthRequestDto) SetCodeOAuthNil() {
	o.CodeOAuth.Set(nil)
}

// UnsetCodeOAuth ensures that no value is present for CodeOAuth, not even an explicit nil
func (o *AuthRequestDto) UnsetCodeOAuth() {
	o.CodeOAuth.Unset()
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *AuthRequestDto) GetSession() bool {
	if o == nil || IsNil(o.Session) {
		var ret bool
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestDto) GetSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *AuthRequestDto) IsSessionSet() bool {
	if o != nil && !IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given bool and assigns it to the Session field.
func (o *AuthRequestDto) SetSession(v bool) {
	o.Session = &v
}

// GetConfirmData returns the ConfirmData field value if set, zero value otherwise.
func (o *AuthRequestDto) GetConfirmData() ConfirmData {
	if o == nil || IsNil(o.ConfirmData) {
		var ret ConfirmData
		return ret
	}
	return *o.ConfirmData
}

// GetConfirmDataOk returns a tuple with the ConfirmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestDto) GetConfirmDataOk() (*ConfirmData, bool) {
	if o == nil || IsNil(o.ConfirmData) {
		return nil, false
	}
	return o.ConfirmData, true
}

// HasConfirmData returns a boolean if a field has been set.
func (o *AuthRequestDto) IsConfirmDataSet() bool {
	if o != nil && !IsNil(o.ConfirmData) {
		return true
	}

	return false
}

// SetConfirmData gets a reference to the given ConfirmData and assigns it to the ConfirmData field.
func (o *AuthRequestDto) SetConfirmData(v ConfirmData) {
	o.ConfirmData = &v
}

// GetRecaptchaType returns the RecaptchaType field value if set, zero value otherwise.
func (o *AuthRequestDto) GetRecaptchaType() RecaptchaType {
	if o == nil || IsNil(o.RecaptchaType) {
		var ret RecaptchaType
		return ret
	}
	return *o.RecaptchaType
}

// GetRecaptchaTypeOk returns a tuple with the RecaptchaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestDto) GetRecaptchaTypeOk() (*RecaptchaType, bool) {
	if o == nil || IsNil(o.RecaptchaType) {
		return nil, false
	}
	return o.RecaptchaType, true
}

// HasRecaptchaType returns a boolean if a field has been set.
func (o *AuthRequestDto) IsRecaptchaTypeSet() bool {
	if o != nil && !IsNil(o.RecaptchaType) {
		return true
	}

	return false
}

// SetRecaptchaType gets a reference to the given RecaptchaType and assigns it to the RecaptchaType field.
func (o *AuthRequestDto) SetRecaptchaType(v RecaptchaType) {
	o.RecaptchaType = &v
}

// GetRecaptchaResponse returns the RecaptchaResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestDto) GetRecaptchaResponse() string {
	if o == nil || IsNil(o.RecaptchaResponse.Get()) {
		var ret string
		return ret
	}
	return *o.RecaptchaResponse.Get()
}

// GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestDto) GetRecaptchaResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecaptchaResponse.Get(), o.RecaptchaResponse.IsSet()
}

// HasRecaptchaResponse returns a boolean if a field has been set.
func (o *AuthRequestDto) IsRecaptchaResponseSet() bool {
	if o != nil && o.RecaptchaResponse.IsSet() {
		return true
	}

	return false
}

// SetRecaptchaResponse gets a reference to the given NullableString and assigns it to the RecaptchaResponse field.
func (o *AuthRequestDto) SetRecaptchaResponse(v string) {
	o.RecaptchaResponse.Set(&v)
}
// SetRecaptchaResponseNil sets the value for RecaptchaResponse to be an explicit nil
func (o *AuthRequestDto) SetRecaptchaResponseNil() {
	o.RecaptchaResponse.Set(nil)
}

// UnsetRecaptchaResponse ensures that no value is present for RecaptchaResponse, not even an explicit nil
func (o *AuthRequestDto) UnsetRecaptchaResponse() {
	o.RecaptchaResponse.Unset()
}

// GetCulture returns the Culture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthRequestDto) GetCulture() string {
	if o == nil || IsNil(o.Culture.Get()) {
		var ret string
		return ret
	}
	return *o.Culture.Get()
}

// GetCultureOk returns a tuple with the Culture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthRequestDto) GetCultureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Culture.Get(), o.Culture.IsSet()
}

// HasCulture returns a boolean if a field has been set.
func (o *AuthRequestDto) IsCultureSet() bool {
	if o != nil && o.Culture.IsSet() {
		return true
	}

	return false
}

// SetCulture gets a reference to the given NullableString and assigns it to the Culture field.
func (o *AuthRequestDto) SetCulture(v string) {
	o.Culture.Set(&v)
}
// SetCultureNil sets the value for Culture to be an explicit nil
func (o *AuthRequestDto) SetCultureNil() {
	o.Culture.Set(nil)
}

// UnsetCulture ensures that no value is present for Culture, not even an explicit nil
func (o *AuthRequestDto) UnsetCulture() {
	o.Culture.Unset()
}

func (o AuthRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthRequestDto) ToMap() (map[string]interface{}, error) {
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

type NullableAuthRequestDto struct {
	value *AuthRequestDto
	isSet bool
}

func (v NullableAuthRequestDto) Get() *AuthRequestDto {
	return v.value
}

func (v *NullableAuthRequestDto) Set(val *AuthRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRequestDto(val *AuthRequestDto) *NullableAuthRequestDto {
	return &NullableAuthRequestDto{value: val, isSet: true}
}

func (v NullableAuthRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

