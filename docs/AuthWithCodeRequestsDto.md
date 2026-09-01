# AuthWithCodeRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** | The username or email used for authentication. | [optional] 
**Password** | Pointer to **string** | The password in plain text for user authentication. | [optional] 
**PasswordHash** | Pointer to **string** | The hashed password for secure verification. | [optional] 
**Provider** | Pointer to **string** | The type of authentication provider (e.g., internal, Google, Azure). | [optional] 
**AccessToken** | Pointer to **string** | The access token used for authentication with external providers. | [optional] 
**SerializedProfile** | Pointer to **string** | The serialized user profile data, if applicable. | [optional] 
**CodeOAuth** | Pointer to **string** | The authorization code used for obtaining OAuth tokens. | [optional] 
**Session** | Pointer to **bool** | Specifies whether the authentication is session-based. | [optional] 
**ConfirmData** | Pointer to [**ConfirmData**](ConfirmData.md) | The additional confirmation data required for authentication. | [optional] 
**RecaptchaType** | Pointer to [**RecaptchaType**](RecaptchaType.md) | The type of CAPTCHA validation used. | [optional] 
**RecaptchaResponse** | Pointer to **string** | The user's response to the CAPTCHA challenge. | [optional] 
**Culture** | Pointer to **string** | The culture code for localization during authentication. | [optional] 
**Code** | Pointer to **NullableString** | The code for two-factor authentication. | [optional] 

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


