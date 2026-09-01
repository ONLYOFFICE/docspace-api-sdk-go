# ProviderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The provider name. | [optional] 
**Key** | Pointer to **NullableString** | The provider key. | [optional] 
**Connected** | Pointer to **bool** | Specifies whether the provider is connected. | [optional] 
**Oauth** | Pointer to **bool** | Specifies if the provider is OAuth. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | The provider redirect URL. | [optional] 
**RequiredConnectionUrl** | Pointer to **bool** | The required connection URL flag. | [optional] 
**ClientId** | Pointer to **NullableString** | The provider OAuth client ID. | [optional] 

## Methods

### NewProviderDto

`func NewProviderDto() *ProviderDto`

NewProviderDto instantiates a new ProviderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderDtoWithDefaults

`func NewProviderDtoWithDefaults() *ProviderDto`

NewProviderDtoWithDefaults instantiates a new ProviderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProviderDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProviderDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProviderDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKey

`func (o *ProviderDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProviderDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProviderDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProviderDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ProviderDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ProviderDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetConnected

`func (o *ProviderDto) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ProviderDto) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ProviderDto) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ProviderDto) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetOauth

`func (o *ProviderDto) GetOauth() bool`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *ProviderDto) GetOauthOk() (*bool, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *ProviderDto) SetOauth(v bool)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *ProviderDto) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *ProviderDto) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ProviderDto) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ProviderDto) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ProviderDto) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *ProviderDto) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *ProviderDto) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetRequiredConnectionUrl

`func (o *ProviderDto) GetRequiredConnectionUrl() bool`

GetRequiredConnectionUrl returns the RequiredConnectionUrl field if non-nil, zero value otherwise.

### GetRequiredConnectionUrlOk

`func (o *ProviderDto) GetRequiredConnectionUrlOk() (*bool, bool)`

GetRequiredConnectionUrlOk returns a tuple with the RequiredConnectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConnectionUrl

`func (o *ProviderDto) SetRequiredConnectionUrl(v bool)`

SetRequiredConnectionUrl sets RequiredConnectionUrl field to given value.

### HasRequiredConnectionUrl

`func (o *ProviderDto) HasRequiredConnectionUrl() bool`

HasRequiredConnectionUrl returns a boolean if a field has been set.

### GetClientId

`func (o *ProviderDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ProviderDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ProviderDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ProviderDto) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *ProviderDto) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *ProviderDto) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


