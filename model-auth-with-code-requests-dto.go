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

// checks if the AuthWithCodeRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthWithCodeRequestsDto{}

// AuthWithCodeRequestsDto The same credentials as an ordinary sign-in, plus the one-time code that completes it.
type AuthWithCodeRequestsDto struct {
	// The account signing in, given as its email address or its portal user name. It is required for a password  sign-in and ignored when the credentials are a confirmation key or a third-party account.
	UserName *string `json:"userName,omitempty"`
	// The password in the clear. Send either this or `passwordHash`, never both; hashing it in the client with the  parameters from `GET api/2.0/settings?withpassword=true` and sending `passwordHash` instead keeps the plain  password off the wire.
	Password *string `json:"password,omitempty"`
	// The password already hashed in the client. It has to be produced with the `salt`, iteration count and hash  size that `GET api/2.0/settings?withpassword=true` publishes, or the portal cannot recognise it; a value sent  here takes the place of `password`.
	PasswordHash *string `json:"passwordHash,omitempty"`
	// The third-party identity provider the account is being signed in through, by its internal key such as  `google` or `linkedin`. Sending it switches the call to a third-party sign-in, which needs `accessToken` or  `serializedProfile` and is only allowed on a self-hosted installation or a tariff that includes third-party  sign-in.
	Provider *string `json:"provider,omitempty"`
	// The access token the provider named in `provider` issued for the account, passed on unchanged for the portal  to verify with that provider. The portal then matches the address it gets back against its own accounts, so a  valid token for an address unknown here is answered as no such user.
	AccessToken *string `json:"accessToken,omitempty"`
	// The third-party profile already fetched and serialised by the caller, as an alternative to `accessToken` for  a provider whose profile the client holds. It identifies the account by the address it carries.
	SerializedProfile *string `json:"serializedProfile,omitempty"`
	// The OAuth authorization code obtained from the provider, for a flow that has not been exchanged for an access  token yet. It is recorded with the sign-in rather than replacing `accessToken`.
	CodeOAuth *string `json:"codeOAuth,omitempty"`
	// Whether the issued token is tied to the browser session. When it is, the answer carries no `expires` and the  token dies with the session; otherwise it lives for the portal session lifetime.
	Session *bool `json:"session,omitempty"`
	// The confirmation link data, as a third way to identify the account beside a password and a third-party  account. Send it when the sign-in comes from a link the portal mailed, in which case `userName` and the  password fields are not read.
	ConfirmData *ConfirmData `json:"confirmData,omitempty"`
	// Which CAPTCHA service the proof in `recaptchaResponse` came from. It has to match the service the  installation is configured with, which `GET api/2.0/settings` publishes together with the site key.
	RecaptchaType *RecaptchaType `json:"recaptchaType,omitempty"`
	// The token the CAPTCHA widget produced in the browser, passed on unchanged for the portal to verify. It is  only demanded once repeated failures have made the portal ask for a challenge, and it is single-use, so a  retry needs a freshly solved one.
	RecaptchaResponse *string `json:"recaptchaResponse,omitempty"`
	// The language the sign-in messages and any letter that follows are written in, as a culture name such as  `en-US`. A culture the installation does not have falls back to the portal language.
	Culture *string `json:"culture,omitempty"`
	// The one-time code from the SMS the portal sent or from the authenticator app, whichever second factor the  portal has enabled for this user. It is single-use and expires; a wrong, empty or expired value fails the  sign-in and counts against the brute-force limit.
	Code NullableString `json:"code,omitempty"`
}

// NewAuthWithCodeRequestsDto instantiates a new AuthWithCodeRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthWithCodeRequestsDto() *AuthWithCodeRequestsDto {
	this := AuthWithCodeRequestsDto{}
	return &this
}

// NewAuthWithCodeRequestsDtoWithDefaults instantiates a new AuthWithCodeRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthWithCodeRequestsDtoWithDefaults() *AuthWithCodeRequestsDto {
	this := AuthWithCodeRequestsDto{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsUserNameSet() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AuthWithCodeRequestsDto) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsPasswordSet() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AuthWithCodeRequestsDto) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetPasswordHash() string {
	if o == nil || IsNil(o.PasswordHash) {
		var ret string
		return ret
	}
	return *o.PasswordHash
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetPasswordHashOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordHash) {
		return nil, false
	}
	return o.PasswordHash, true
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsPasswordHashSet() bool {
	if o != nil && !IsNil(o.PasswordHash) {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given string and assigns it to the PasswordHash field.
func (o *AuthWithCodeRequestsDto) SetPasswordHash(v string) {
	o.PasswordHash = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsProviderSet() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *AuthWithCodeRequestsDto) SetProvider(v string) {
	o.Provider = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsAccessTokenSet() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AuthWithCodeRequestsDto) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetSerializedProfile returns the SerializedProfile field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetSerializedProfile() string {
	if o == nil || IsNil(o.SerializedProfile) {
		var ret string
		return ret
	}
	return *o.SerializedProfile
}

// GetSerializedProfileOk returns a tuple with the SerializedProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetSerializedProfileOk() (*string, bool) {
	if o == nil || IsNil(o.SerializedProfile) {
		return nil, false
	}
	return o.SerializedProfile, true
}

// HasSerializedProfile returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsSerializedProfileSet() bool {
	if o != nil && !IsNil(o.SerializedProfile) {
		return true
	}

	return false
}

// SetSerializedProfile gets a reference to the given string and assigns it to the SerializedProfile field.
func (o *AuthWithCodeRequestsDto) SetSerializedProfile(v string) {
	o.SerializedProfile = &v
}

// GetCodeOAuth returns the CodeOAuth field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetCodeOAuth() string {
	if o == nil || IsNil(o.CodeOAuth) {
		var ret string
		return ret
	}
	return *o.CodeOAuth
}

// GetCodeOAuthOk returns a tuple with the CodeOAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetCodeOAuthOk() (*string, bool) {
	if o == nil || IsNil(o.CodeOAuth) {
		return nil, false
	}
	return o.CodeOAuth, true
}

// HasCodeOAuth returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsCodeOAuthSet() bool {
	if o != nil && !IsNil(o.CodeOAuth) {
		return true
	}

	return false
}

// SetCodeOAuth gets a reference to the given string and assigns it to the CodeOAuth field.
func (o *AuthWithCodeRequestsDto) SetCodeOAuth(v string) {
	o.CodeOAuth = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetSession() bool {
	if o == nil || IsNil(o.Session) {
		var ret bool
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsSessionSet() bool {
	if o != nil && !IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given bool and assigns it to the Session field.
func (o *AuthWithCodeRequestsDto) SetSession(v bool) {
	o.Session = &v
}

// GetConfirmData returns the ConfirmData field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetConfirmData() ConfirmData {
	if o == nil || IsNil(o.ConfirmData) {
		var ret ConfirmData
		return ret
	}
	return *o.ConfirmData
}

// GetConfirmDataOk returns a tuple with the ConfirmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetConfirmDataOk() (*ConfirmData, bool) {
	if o == nil || IsNil(o.ConfirmData) {
		return nil, false
	}
	return o.ConfirmData, true
}

// HasConfirmData returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsConfirmDataSet() bool {
	if o != nil && !IsNil(o.ConfirmData) {
		return true
	}

	return false
}

// SetConfirmData gets a reference to the given ConfirmData and assigns it to the ConfirmData field.
func (o *AuthWithCodeRequestsDto) SetConfirmData(v ConfirmData) {
	o.ConfirmData = &v
}

// GetRecaptchaType returns the RecaptchaType field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetRecaptchaType() RecaptchaType {
	if o == nil || IsNil(o.RecaptchaType) {
		var ret RecaptchaType
		return ret
	}
	return *o.RecaptchaType
}

// GetRecaptchaTypeOk returns a tuple with the RecaptchaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetRecaptchaTypeOk() (*RecaptchaType, bool) {
	if o == nil || IsNil(o.RecaptchaType) {
		return nil, false
	}
	return o.RecaptchaType, true
}

// HasRecaptchaType returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsRecaptchaTypeSet() bool {
	if o != nil && !IsNil(o.RecaptchaType) {
		return true
	}

	return false
}

// SetRecaptchaType gets a reference to the given RecaptchaType and assigns it to the RecaptchaType field.
func (o *AuthWithCodeRequestsDto) SetRecaptchaType(v RecaptchaType) {
	o.RecaptchaType = &v
}

// GetRecaptchaResponse returns the RecaptchaResponse field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetRecaptchaResponse() string {
	if o == nil || IsNil(o.RecaptchaResponse) {
		var ret string
		return ret
	}
	return *o.RecaptchaResponse
}

// GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetRecaptchaResponseOk() (*string, bool) {
	if o == nil || IsNil(o.RecaptchaResponse) {
		return nil, false
	}
	return o.RecaptchaResponse, true
}

// HasRecaptchaResponse returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsRecaptchaResponseSet() bool {
	if o != nil && !IsNil(o.RecaptchaResponse) {
		return true
	}

	return false
}

// SetRecaptchaResponse gets a reference to the given string and assigns it to the RecaptchaResponse field.
func (o *AuthWithCodeRequestsDto) SetRecaptchaResponse(v string) {
	o.RecaptchaResponse = &v
}

// GetCulture returns the Culture field value if set, zero value otherwise.
func (o *AuthWithCodeRequestsDto) GetCulture() string {
	if o == nil || IsNil(o.Culture) {
		var ret string
		return ret
	}
	return *o.Culture
}

// GetCultureOk returns a tuple with the Culture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWithCodeRequestsDto) GetCultureOk() (*string, bool) {
	if o == nil || IsNil(o.Culture) {
		return nil, false
	}
	return o.Culture, true
}

// HasCulture returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsCultureSet() bool {
	if o != nil && !IsNil(o.Culture) {
		return true
	}

	return false
}

// SetCulture gets a reference to the given string and assigns it to the Culture field.
func (o *AuthWithCodeRequestsDto) SetCulture(v string) {
	o.Culture = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthWithCodeRequestsDto) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthWithCodeRequestsDto) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *AuthWithCodeRequestsDto) IsCodeSet() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *AuthWithCodeRequestsDto) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *AuthWithCodeRequestsDto) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *AuthWithCodeRequestsDto) UnsetCode() {
	o.Code.Unset()
}

func (o AuthWithCodeRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthWithCodeRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PasswordHash) {
		toSerialize["passwordHash"] = o.PasswordHash
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.SerializedProfile) {
		toSerialize["serializedProfile"] = o.SerializedProfile
	}
	if !IsNil(o.CodeOAuth) {
		toSerialize["codeOAuth"] = o.CodeOAuth
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
	if !IsNil(o.RecaptchaResponse) {
		toSerialize["recaptchaResponse"] = o.RecaptchaResponse
	}
	if !IsNil(o.Culture) {
		toSerialize["culture"] = o.Culture
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	return toSerialize, nil
}

type NullableAuthWithCodeRequestsDto struct {
	value *AuthWithCodeRequestsDto
	isSet bool
}

func (v NullableAuthWithCodeRequestsDto) Get() *AuthWithCodeRequestsDto {
	return v.value
}

func (v *NullableAuthWithCodeRequestsDto) Set(val *AuthWithCodeRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthWithCodeRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthWithCodeRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthWithCodeRequestsDto(val *AuthWithCodeRequestsDto) *NullableAuthWithCodeRequestsDto {
	return &NullableAuthWithCodeRequestsDto{value: val, isSet: true}
}

func (v NullableAuthWithCodeRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthWithCodeRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

