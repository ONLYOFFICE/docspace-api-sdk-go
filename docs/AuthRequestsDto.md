# AuthRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **NullableString** | The account signing in, given as its email address or its portal user name. It is required for a password  sign-in and ignored when the credentials are a confirmation key or a third-party account. | [optional] 
**Password** | Pointer to **NullableString** | The password in the clear. Send either this or `passwordHash`, never both; hashing it in the client with the  parameters from `GET api/2.0/settings?withpassword=true` and sending `passwordHash` instead keeps the plain  password off the wire. | [optional] 
**PasswordHash** | Pointer to **NullableString** | The password already hashed in the client. It has to be produced with the `salt`, iteration count and hash  size that `GET api/2.0/settings?withpassword=true` publishes, or the portal cannot recognise it; a value sent  here takes the place of `password`. | [optional] 
**Provider** | Pointer to **NullableString** | The third-party identity provider the account is being signed in through, by its internal key such as  `google` or `linkedin`. Sending it switches the call to a third-party sign-in, which needs `accessToken` or  `serializedProfile` and is only allowed on a self-hosted installation or a tariff that includes third-party  sign-in. | [optional] 
**AccessToken** | Pointer to **NullableString** | The access token the provider named in `provider` issued for the account, passed on unchanged for the portal  to verify with that provider. The portal then matches the address it gets back against its own accounts, so a  valid token for an address unknown here is answered as no such user. | [optional] 
**SerializedProfile** | Pointer to **NullableString** | The third-party profile already fetched and serialised by the caller, as an alternative to `accessToken` for  a provider whose profile the client holds. It identifies the account by the address it carries. | [optional] 
**CodeOAuth** | Pointer to **NullableString** | The OAuth authorization code obtained from the provider, for a flow that has not been exchanged for an access  token yet. It is recorded with the sign-in rather than replacing `accessToken`. | [optional] 
**Session** | Pointer to **bool** | Whether the issued token is tied to the browser session. When it is, the answer carries no `expires` and the  token dies with the session; otherwise it lives for the portal session lifetime. | [optional] 
**ConfirmData** | Pointer to [**ConfirmData**](ConfirmData.md) | The confirmation link data, as a third way to identify the account beside a password and a third-party  account. Send it when the sign-in comes from a link the portal mailed, in which case `userName` and the  password fields are not read. | [optional] 
**RecaptchaType** | Pointer to [**RecaptchaType**](RecaptchaType.md) | Which CAPTCHA service the proof in `recaptchaResponse` came from. It has to match the service the  installation is configured with, which `GET api/2.0/settings` publishes together with the site key. | [optional] 
**RecaptchaResponse** | Pointer to **NullableString** | The token the CAPTCHA widget produced in the browser, passed on unchanged for the portal to verify. It is  only demanded once repeated failures have made the portal ask for a challenge, and it is single-use, so a  retry needs a freshly solved one. | [optional] 
**Culture** | Pointer to **NullableString** | The language the sign-in messages and any letter that follows are written in, as a culture name such as  `en-US`. A culture the installation does not have falls back to the portal language. | [optional] 

## Methods

### NewAuthRequestsDto

`func NewAuthRequestsDto() *AuthRequestsDto`

NewAuthRequestsDto instantiates a new AuthRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRequestsDtoWithDefaults

`func NewAuthRequestsDtoWithDefaults() *AuthRequestsDto`

NewAuthRequestsDtoWithDefaults instantiates a new AuthRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *AuthRequestsDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AuthRequestsDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AuthRequestsDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AuthRequestsDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *AuthRequestsDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *AuthRequestsDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetPassword

`func (o *AuthRequestsDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthRequestsDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthRequestsDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthRequestsDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AuthRequestsDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AuthRequestsDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *AuthRequestsDto) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AuthRequestsDto) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AuthRequestsDto) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AuthRequestsDto) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *AuthRequestsDto) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *AuthRequestsDto) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetProvider

`func (o *AuthRequestsDto) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthRequestsDto) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthRequestsDto) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AuthRequestsDto) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *AuthRequestsDto) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AuthRequestsDto) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetAccessToken

`func (o *AuthRequestsDto) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthRequestsDto) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthRequestsDto) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthRequestsDto) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *AuthRequestsDto) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *AuthRequestsDto) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetSerializedProfile

`func (o *AuthRequestsDto) GetSerializedProfile() string`

GetSerializedProfile returns the SerializedProfile field if non-nil, zero value otherwise.

### GetSerializedProfileOk

`func (o *AuthRequestsDto) GetSerializedProfileOk() (*string, bool)`

GetSerializedProfileOk returns a tuple with the SerializedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedProfile

`func (o *AuthRequestsDto) SetSerializedProfile(v string)`

SetSerializedProfile sets SerializedProfile field to given value.

### HasSerializedProfile

`func (o *AuthRequestsDto) HasSerializedProfile() bool`

HasSerializedProfile returns a boolean if a field has been set.

### SetSerializedProfileNil

`func (o *AuthRequestsDto) SetSerializedProfileNil(b bool)`

 SetSerializedProfileNil sets the value for SerializedProfile to be an explicit nil

### UnsetSerializedProfile
`func (o *AuthRequestsDto) UnsetSerializedProfile()`

UnsetSerializedProfile ensures that no value is present for SerializedProfile, not even an explicit nil
### GetCodeOAuth

`func (o *AuthRequestsDto) GetCodeOAuth() string`

GetCodeOAuth returns the CodeOAuth field if non-nil, zero value otherwise.

### GetCodeOAuthOk

`func (o *AuthRequestsDto) GetCodeOAuthOk() (*string, bool)`

GetCodeOAuthOk returns a tuple with the CodeOAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeOAuth

`func (o *AuthRequestsDto) SetCodeOAuth(v string)`

SetCodeOAuth sets CodeOAuth field to given value.

### HasCodeOAuth

`func (o *AuthRequestsDto) HasCodeOAuth() bool`

HasCodeOAuth returns a boolean if a field has been set.

### SetCodeOAuthNil

`func (o *AuthRequestsDto) SetCodeOAuthNil(b bool)`

 SetCodeOAuthNil sets the value for CodeOAuth to be an explicit nil

### UnsetCodeOAuth
`func (o *AuthRequestsDto) UnsetCodeOAuth()`

UnsetCodeOAuth ensures that no value is present for CodeOAuth, not even an explicit nil
### GetSession

`func (o *AuthRequestsDto) GetSession() bool`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AuthRequestsDto) GetSessionOk() (*bool, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AuthRequestsDto) SetSession(v bool)`

SetSession sets Session field to given value.

### HasSession

`func (o *AuthRequestsDto) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetConfirmData

`func (o *AuthRequestsDto) GetConfirmData() ConfirmData`

GetConfirmData returns the ConfirmData field if non-nil, zero value otherwise.

### GetConfirmDataOk

`func (o *AuthRequestsDto) GetConfirmDataOk() (*ConfirmData, bool)`

GetConfirmDataOk returns a tuple with the ConfirmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmData

`func (o *AuthRequestsDto) SetConfirmData(v ConfirmData)`

SetConfirmData sets ConfirmData field to given value.

### HasConfirmData

`func (o *AuthRequestsDto) HasConfirmData() bool`

HasConfirmData returns a boolean if a field has been set.

### GetRecaptchaType

`func (o *AuthRequestsDto) GetRecaptchaType() RecaptchaType`

GetRecaptchaType returns the RecaptchaType field if non-nil, zero value otherwise.

### GetRecaptchaTypeOk

`func (o *AuthRequestsDto) GetRecaptchaTypeOk() (*RecaptchaType, bool)`

GetRecaptchaTypeOk returns a tuple with the RecaptchaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaType

`func (o *AuthRequestsDto) SetRecaptchaType(v RecaptchaType)`

SetRecaptchaType sets RecaptchaType field to given value.

### HasRecaptchaType

`func (o *AuthRequestsDto) HasRecaptchaType() bool`

HasRecaptchaType returns a boolean if a field has been set.

### GetRecaptchaResponse

`func (o *AuthRequestsDto) GetRecaptchaResponse() string`

GetRecaptchaResponse returns the RecaptchaResponse field if non-nil, zero value otherwise.

### GetRecaptchaResponseOk

`func (o *AuthRequestsDto) GetRecaptchaResponseOk() (*string, bool)`

GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaResponse

`func (o *AuthRequestsDto) SetRecaptchaResponse(v string)`

SetRecaptchaResponse sets RecaptchaResponse field to given value.

### HasRecaptchaResponse

`func (o *AuthRequestsDto) HasRecaptchaResponse() bool`

HasRecaptchaResponse returns a boolean if a field has been set.

### SetRecaptchaResponseNil

`func (o *AuthRequestsDto) SetRecaptchaResponseNil(b bool)`

 SetRecaptchaResponseNil sets the value for RecaptchaResponse to be an explicit nil

### UnsetRecaptchaResponse
`func (o *AuthRequestsDto) UnsetRecaptchaResponse()`

UnsetRecaptchaResponse ensures that no value is present for RecaptchaResponse, not even an explicit nil
### GetCulture

`func (o *AuthRequestsDto) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *AuthRequestsDto) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *AuthRequestsDto) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *AuthRequestsDto) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### SetCultureNil

`func (o *AuthRequestsDto) SetCultureNil(b bool)`

 SetCultureNil sets the value for Culture to be an explicit nil

### UnsetCulture
`func (o *AuthRequestsDto) UnsetCulture()`

UnsetCulture ensures that no value is present for Culture, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


