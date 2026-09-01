# ThirdPartyRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** | The connection URL for the sharepoint. | [optional] 
**Login** | Pointer to **NullableString** | The third-party request login. | [optional] 
**Password** | Pointer to **NullableString** | The third-party request password. | [optional] 
**Token** | Pointer to **NullableString** | The authentication token. | [optional] 
**CustomerTitle** | **NullableString** | The customer title. | 
**ProviderKey** | **NullableString** | The provider key. | 
**ProviderId** | Pointer to **NullableInt32** | The provider ID. | [optional] 

## Methods

### NewThirdPartyRequestDto

`func NewThirdPartyRequestDto(customerTitle NullableString, providerKey NullableString, ) *ThirdPartyRequestDto`

NewThirdPartyRequestDto instantiates a new ThirdPartyRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyRequestDtoWithDefaults

`func NewThirdPartyRequestDtoWithDefaults() *ThirdPartyRequestDto`

NewThirdPartyRequestDtoWithDefaults instantiates a new ThirdPartyRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ThirdPartyRequestDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ThirdPartyRequestDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ThirdPartyRequestDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ThirdPartyRequestDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ThirdPartyRequestDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ThirdPartyRequestDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetLogin

`func (o *ThirdPartyRequestDto) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *ThirdPartyRequestDto) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *ThirdPartyRequestDto) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *ThirdPartyRequestDto) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### SetLoginNil

`func (o *ThirdPartyRequestDto) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *ThirdPartyRequestDto) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetPassword

`func (o *ThirdPartyRequestDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ThirdPartyRequestDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ThirdPartyRequestDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ThirdPartyRequestDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ThirdPartyRequestDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ThirdPartyRequestDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetToken

`func (o *ThirdPartyRequestDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ThirdPartyRequestDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ThirdPartyRequestDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ThirdPartyRequestDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *ThirdPartyRequestDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *ThirdPartyRequestDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetCustomerTitle

`func (o *ThirdPartyRequestDto) GetCustomerTitle() string`

GetCustomerTitle returns the CustomerTitle field if non-nil, zero value otherwise.

### GetCustomerTitleOk

`func (o *ThirdPartyRequestDto) GetCustomerTitleOk() (*string, bool)`

GetCustomerTitleOk returns a tuple with the CustomerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTitle

`func (o *ThirdPartyRequestDto) SetCustomerTitle(v string)`

SetCustomerTitle sets CustomerTitle field to given value.


### SetCustomerTitleNil

`func (o *ThirdPartyRequestDto) SetCustomerTitleNil(b bool)`

 SetCustomerTitleNil sets the value for CustomerTitle to be an explicit nil

### UnsetCustomerTitle
`func (o *ThirdPartyRequestDto) UnsetCustomerTitle()`

UnsetCustomerTitle ensures that no value is present for CustomerTitle, not even an explicit nil
### GetProviderKey

`func (o *ThirdPartyRequestDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *ThirdPartyRequestDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *ThirdPartyRequestDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.


### SetProviderKeyNil

`func (o *ThirdPartyRequestDto) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *ThirdPartyRequestDto) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetProviderId

`func (o *ThirdPartyRequestDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ThirdPartyRequestDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ThirdPartyRequestDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ThirdPartyRequestDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ThirdPartyRequestDto) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ThirdPartyRequestDto) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


