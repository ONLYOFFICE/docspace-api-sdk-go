# ThirdPartyBackupRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** | The address of the storage server to connect to. It is needed by the WebDAV presets whose server is not known  in advance (`WebDav`, `Nextcloud`, `ownCloud`), where it points at the WebDAV endpoint of that server, and by  `SharePoint`; the presets with a fixed address and the OAuth services ignore it. | [optional] 
**Login** | Pointer to **NullableString** | The account name at the storage service, used by the services that authenticate by login and password. A login  sent without a password is rejected as an invalid request. | [optional] 
**Password** | Pointer to **NullableString** | The password, or the application password, for `login` at the storage service. Either this or `token` has to  be sent, and the credentials are verified against the service before the account is saved. | [optional] 
**Token** | Pointer to **NullableString** | The OAuth 2.0 authorization code from the consent screen of `Box`, `DropboxV2`, `GoogleDrive` or `OneDrive` -  not an access token: the portal exchanges the code for its own token and keeps that. The client ID and  redirect URL the consent screen URL is built from come from `GET api/2.0/files/thirdparty/capabilities`. | [optional] 
**CustomerTitle** | Pointer to **NullableString** | The name the backup account is shown under in the portal. Characters that a folder title cannot hold are  replaced and the value is truncated; on the first connection a title that comes out of that empty is refused. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The storage service to connect, as the `key` of `GET api/2.0/files/thirdparty/providers`; the value is matched  case-insensitively. `Nextcloud` and `ownCloud` are presets over WebDAV and are stored and reported back as  `WebDav`. | [optional] 

## Methods

### NewThirdPartyBackupRequestDto

`func NewThirdPartyBackupRequestDto() *ThirdPartyBackupRequestDto`

NewThirdPartyBackupRequestDto instantiates a new ThirdPartyBackupRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyBackupRequestDtoWithDefaults

`func NewThirdPartyBackupRequestDtoWithDefaults() *ThirdPartyBackupRequestDto`

NewThirdPartyBackupRequestDtoWithDefaults instantiates a new ThirdPartyBackupRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ThirdPartyBackupRequestDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ThirdPartyBackupRequestDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ThirdPartyBackupRequestDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ThirdPartyBackupRequestDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ThirdPartyBackupRequestDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ThirdPartyBackupRequestDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetLogin

`func (o *ThirdPartyBackupRequestDto) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *ThirdPartyBackupRequestDto) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *ThirdPartyBackupRequestDto) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *ThirdPartyBackupRequestDto) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### SetLoginNil

`func (o *ThirdPartyBackupRequestDto) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *ThirdPartyBackupRequestDto) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetPassword

`func (o *ThirdPartyBackupRequestDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ThirdPartyBackupRequestDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ThirdPartyBackupRequestDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ThirdPartyBackupRequestDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ThirdPartyBackupRequestDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ThirdPartyBackupRequestDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetToken

`func (o *ThirdPartyBackupRequestDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ThirdPartyBackupRequestDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ThirdPartyBackupRequestDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ThirdPartyBackupRequestDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *ThirdPartyBackupRequestDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *ThirdPartyBackupRequestDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetCustomerTitle

`func (o *ThirdPartyBackupRequestDto) GetCustomerTitle() string`

GetCustomerTitle returns the CustomerTitle field if non-nil, zero value otherwise.

### GetCustomerTitleOk

`func (o *ThirdPartyBackupRequestDto) GetCustomerTitleOk() (*string, bool)`

GetCustomerTitleOk returns a tuple with the CustomerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTitle

`func (o *ThirdPartyBackupRequestDto) SetCustomerTitle(v string)`

SetCustomerTitle sets CustomerTitle field to given value.

### HasCustomerTitle

`func (o *ThirdPartyBackupRequestDto) HasCustomerTitle() bool`

HasCustomerTitle returns a boolean if a field has been set.

### SetCustomerTitleNil

`func (o *ThirdPartyBackupRequestDto) SetCustomerTitleNil(b bool)`

 SetCustomerTitleNil sets the value for CustomerTitle to be an explicit nil

### UnsetCustomerTitle
`func (o *ThirdPartyBackupRequestDto) UnsetCustomerTitle()`

UnsetCustomerTitle ensures that no value is present for CustomerTitle, not even an explicit nil
### GetProviderKey

`func (o *ThirdPartyBackupRequestDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *ThirdPartyBackupRequestDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *ThirdPartyBackupRequestDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *ThirdPartyBackupRequestDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *ThirdPartyBackupRequestDto) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *ThirdPartyBackupRequestDto) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


