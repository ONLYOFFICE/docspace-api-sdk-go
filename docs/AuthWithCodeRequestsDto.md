# AuthWithCodeRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** | The account signing in, given as its email address or its portal user name. It is required for a password  sign-in and ignored when the credentials are a confirmation key or a third-party account. | [optional] 
**Password** | Pointer to **string** | The password in the clear. Send either this or `passwordHash`, never both; hashing it in the client with the  parameters from `GET api/2.0/settings?withpassword=true` and sending `passwordHash` instead keeps the plain  password off the wire. | [optional] 
**PasswordHash** | Pointer to **string** | The password already hashed in the client. It has to be produced with the `salt`, iteration count and hash  size that `GET api/2.0/settings?withpassword=true` publishes, or the portal cannot recognise it; a value sent  here takes the place of `password`. | [optional] 
**Provider** | Pointer to **string** | The third-party identity provider the account is being signed in through, by its internal key such as  `google` or `linkedin`. Sending it switches the call to a third-party sign-in, which needs `accessToken` or  `serializedProfile` and is only allowed on a self-hosted installation or a tariff that includes third-party  sign-in. | [optional] 
**AccessToken** | Pointer to **string** | The access token the provider named in `provider` issued for the account, passed on unchanged for the portal  to verify with that provider. The portal then matches the address it gets back against its own accounts, so a  valid token for an address unknown here is answered as no such user. | [optional] 
**SerializedProfile** | Pointer to **string** | The third-party profile already fetched and serialised by the caller, as an alternative to `accessToken` for  a provider whose profile the client holds. It identifies the account by the address it carries. | [optional] 
**CodeOAuth** | Pointer to **string** | The OAuth authorization code obtained from the provider, for a flow that has not been exchanged for an access  token yet. It is recorded with the sign-in rather than replacing `accessToken`. | [optional] 
**Session** | Pointer to **bool** | Whether the issued token is tied to the browser session. When it is, the answer carries no `expires` and the  token dies with the session; otherwise it lives for the portal session lifetime. | [optional] 
**ConfirmData** | Pointer to [**ConfirmData**](ConfirmData.md) | The confirmation link data, as a third way to identify the account beside a password and a third-party  account. Send it when the sign-in comes from a link the portal mailed, in which case `userName` and the  password fields are not read. | [optional] 
**RecaptchaType** | Pointer to [**RecaptchaType**](RecaptchaType.md) | Which CAPTCHA service the proof in `recaptchaResponse` came from. It has to match the service the  installation is configured with, which `GET api/2.0/settings` publishes together with the site key. | [optional] 
**RecaptchaResponse** | Pointer to **string** | The token the CAPTCHA widget produced in the browser, passed on unchanged for the portal to verify. It is  only demanded once repeated failures have made the portal ask for a challenge, and it is single-use, so a  retry needs a freshly solved one. | [optional] 
**Culture** | Pointer to **string** | The language the sign-in messages and any letter that follows are written in, as a culture name such as  `en-US`. A culture the installation does not have falls back to the portal language. | [optional] 
**Code** | Pointer to **NullableString** | The one-time code from the SMS the portal sent or from the authenticator app, whichever second factor the  portal has enabled for this user. It is single-use and expires; a wrong, empty or expired value fails the  sign-in and counts against the brute-force limit. | [optional] 

## Methods

### NewAuthWithCodeRequestsDto

`func NewAuthWithCodeRequestsDto() *AuthWithCodeRequestsDto`

NewAuthWithCodeRequestsDto instantiates a new AuthWithCodeRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthWithCodeRequestsDtoWithDefaults

`func NewAuthWithCodeRequestsDtoWithDefaults() *AuthWithCodeRequestsDto`

NewAuthWithCodeRequestsDtoWithDefaults instantiates a new AuthWithCodeRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *AuthWithCodeRequestsDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AuthWithCodeRequestsDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AuthWithCodeRequestsDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AuthWithCodeRequestsDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *AuthWithCodeRequestsDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthWithCodeRequestsDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthWithCodeRequestsDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthWithCodeRequestsDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *AuthWithCodeRequestsDto) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AuthWithCodeRequestsDto) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AuthWithCodeRequestsDto) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AuthWithCodeRequestsDto) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetProvider

`func (o *AuthWithCodeRequestsDto) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthWithCodeRequestsDto) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthWithCodeRequestsDto) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AuthWithCodeRequestsDto) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetAccessToken

`func (o *AuthWithCodeRequestsDto) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthWithCodeRequestsDto) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthWithCodeRequestsDto) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthWithCodeRequestsDto) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetSerializedProfile

`func (o *AuthWithCodeRequestsDto) GetSerializedProfile() string`

GetSerializedProfile returns the SerializedProfile field if non-nil, zero value otherwise.

### GetSerializedProfileOk

`func (o *AuthWithCodeRequestsDto) GetSerializedProfileOk() (*string, bool)`

GetSerializedProfileOk returns a tuple with the SerializedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedProfile

`func (o *AuthWithCodeRequestsDto) SetSerializedProfile(v string)`

SetSerializedProfile sets SerializedProfile field to given value.

### HasSerializedProfile

`func (o *AuthWithCodeRequestsDto) HasSerializedProfile() bool`

HasSerializedProfile returns a boolean if a field has been set.

### GetCodeOAuth

`func (o *AuthWithCodeRequestsDto) GetCodeOAuth() string`

GetCodeOAuth returns the CodeOAuth field if non-nil, zero value otherwise.

### GetCodeOAuthOk

`func (o *AuthWithCodeRequestsDto) GetCodeOAuthOk() (*string, bool)`

GetCodeOAuthOk returns a tuple with the CodeOAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeOAuth

`func (o *AuthWithCodeRequestsDto) SetCodeOAuth(v string)`

SetCodeOAuth sets CodeOAuth field to given value.

### HasCodeOAuth

`func (o *AuthWithCodeRequestsDto) HasCodeOAuth() bool`

HasCodeOAuth returns a boolean if a field has been set.

### GetSession

`func (o *AuthWithCodeRequestsDto) GetSession() bool`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AuthWithCodeRequestsDto) GetSessionOk() (*bool, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AuthWithCodeRequestsDto) SetSession(v bool)`

SetSession sets Session field to given value.

### HasSession

`func (o *AuthWithCodeRequestsDto) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetConfirmData

`func (o *AuthWithCodeRequestsDto) GetConfirmData() ConfirmData`

GetConfirmData returns the ConfirmData field if non-nil, zero value otherwise.

### GetConfirmDataOk

`func (o *AuthWithCodeRequestsDto) GetConfirmDataOk() (*ConfirmData, bool)`

GetConfirmDataOk returns a tuple with the ConfirmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmData

`func (o *AuthWithCodeRequestsDto) SetConfirmData(v ConfirmData)`

SetConfirmData sets ConfirmData field to given value.

### HasConfirmData

`func (o *AuthWithCodeRequestsDto) HasConfirmData() bool`

HasConfirmData returns a boolean if a field has been set.

### GetRecaptchaType

`func (o *AuthWithCodeRequestsDto) GetRecaptchaType() RecaptchaType`

GetRecaptchaType returns the RecaptchaType field if non-nil, zero value otherwise.

### GetRecaptchaTypeOk

`func (o *AuthWithCodeRequestsDto) GetRecaptchaTypeOk() (*RecaptchaType, bool)`

GetRecaptchaTypeOk returns a tuple with the RecaptchaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaType

`func (o *AuthWithCodeRequestsDto) SetRecaptchaType(v RecaptchaType)`

SetRecaptchaType sets RecaptchaType field to given value.

### HasRecaptchaType

`func (o *AuthWithCodeRequestsDto) HasRecaptchaType() bool`

HasRecaptchaType returns a boolean if a field has been set.

### GetRecaptchaResponse

`func (o *AuthWithCodeRequestsDto) GetRecaptchaResponse() string`

GetRecaptchaResponse returns the RecaptchaResponse field if non-nil, zero value otherwise.

### GetRecaptchaResponseOk

`func (o *AuthWithCodeRequestsDto) GetRecaptchaResponseOk() (*string, bool)`

GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaResponse

`func (o *AuthWithCodeRequestsDto) SetRecaptchaResponse(v string)`

SetRecaptchaResponse sets RecaptchaResponse field to given value.

### HasRecaptchaResponse

`func (o *AuthWithCodeRequestsDto) HasRecaptchaResponse() bool`

HasRecaptchaResponse returns a boolean if a field has been set.

### GetCulture

`func (o *AuthWithCodeRequestsDto) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *AuthWithCodeRequestsDto) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *AuthWithCodeRequestsDto) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *AuthWithCodeRequestsDto) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### GetCode

`func (o *AuthWithCodeRequestsDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthWithCodeRequestsDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthWithCodeRequestsDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthWithCodeRequestsDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AuthWithCodeRequestsDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AuthWithCodeRequestsDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


