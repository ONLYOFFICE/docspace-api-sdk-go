# AuthData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **NullableString** | The account name at the storage service. | [optional] 
**Password** | Pointer to **NullableString** | The password of the account at the storage service. | [optional] 
**RawToken** | Pointer to **NullableString** | The token of the account, kept as the raw JSON document the storage service issued it in. | [optional] 
**Url** | Pointer to **NullableString** | The address of the storage server the account lives on. | [optional] 
**Provider** | Pointer to **NullableString** | The storage service the credentials belong to, as the provider key the account was connected with. | [optional] 
**Token** | Pointer to [**OAuth20Token**](OAuth20Token.md) | The same token as in `rawToken`, parsed into its OAuth 2.0 fields. | [optional] 

## Methods

### NewAuthData

`func NewAuthData() *AuthData`

NewAuthData instantiates a new AuthData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthDataWithDefaults

`func NewAuthDataWithDefaults() *AuthData`

NewAuthDataWithDefaults instantiates a new AuthData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *AuthData) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *AuthData) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *AuthData) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *AuthData) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### SetLoginNil

`func (o *AuthData) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *AuthData) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetPassword

`func (o *AuthData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AuthData) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AuthData) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRawToken

`func (o *AuthData) GetRawToken() string`

GetRawToken returns the RawToken field if non-nil, zero value otherwise.

### GetRawTokenOk

`func (o *AuthData) GetRawTokenOk() (*string, bool)`

GetRawTokenOk returns a tuple with the RawToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawToken

`func (o *AuthData) SetRawToken(v string)`

SetRawToken sets RawToken field to given value.

### HasRawToken

`func (o *AuthData) HasRawToken() bool`

HasRawToken returns a boolean if a field has been set.

### SetRawTokenNil

`func (o *AuthData) SetRawTokenNil(b bool)`

 SetRawTokenNil sets the value for RawToken to be an explicit nil

### UnsetRawToken
`func (o *AuthData) UnsetRawToken()`

UnsetRawToken ensures that no value is present for RawToken, not even an explicit nil
### GetUrl

`func (o *AuthData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuthData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuthData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AuthData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AuthData) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AuthData) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetProvider

`func (o *AuthData) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthData) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthData) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AuthData) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *AuthData) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AuthData) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetToken

`func (o *AuthData) GetToken() OAuth20Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthData) GetTokenOk() (*OAuth20Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthData) SetToken(v OAuth20Token)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthData) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


