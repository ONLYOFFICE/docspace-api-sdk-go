# ThirdPartyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthData** | Pointer to [**AuthData**](AuthData.md) | The authentication data. | [optional] 
**Corporate** | Pointer to **bool** | Specifies if this is a corporate account or not. | [optional] 
**RoomsStorage** | Pointer to **bool** | Specifies if this is a room storage or not. | [optional] 
**CustomerTitle** | Pointer to **NullableString** | The customer title. | [optional] 
**ProviderId** | Pointer to **NullableInt32** | The provider ID. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The provider key. | [optional] 

## Methods

### NewThirdPartyParams

`func NewThirdPartyParams() *ThirdPartyParams`

NewThirdPartyParams instantiates a new ThirdPartyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyParamsWithDefaults

`func NewThirdPartyParamsWithDefaults() *ThirdPartyParams`

NewThirdPartyParamsWithDefaults instantiates a new ThirdPartyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthData

`func (o *ThirdPartyParams) GetAuthData() AuthData`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *ThirdPartyParams) GetAuthDataOk() (*AuthData, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *ThirdPartyParams) SetAuthData(v AuthData)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *ThirdPartyParams) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.

### GetCorporate

`func (o *ThirdPartyParams) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *ThirdPartyParams) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *ThirdPartyParams) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *ThirdPartyParams) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetRoomsStorage

`func (o *ThirdPartyParams) GetRoomsStorage() bool`

GetRoomsStorage returns the RoomsStorage field if non-nil, zero value otherwise.

### GetRoomsStorageOk

`func (o *ThirdPartyParams) GetRoomsStorageOk() (*bool, bool)`

GetRoomsStorageOk returns a tuple with the RoomsStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsStorage

`func (o *ThirdPartyParams) SetRoomsStorage(v bool)`

SetRoomsStorage sets RoomsStorage field to given value.

### HasRoomsStorage

`func (o *ThirdPartyParams) HasRoomsStorage() bool`

HasRoomsStorage returns a boolean if a field has been set.

### GetCustomerTitle

`func (o *ThirdPartyParams) GetCustomerTitle() string`

GetCustomerTitle returns the CustomerTitle field if non-nil, zero value otherwise.

### GetCustomerTitleOk

`func (o *ThirdPartyParams) GetCustomerTitleOk() (*string, bool)`

GetCustomerTitleOk returns a tuple with the CustomerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTitle

`func (o *ThirdPartyParams) SetCustomerTitle(v string)`

SetCustomerTitle sets CustomerTitle field to given value.

### HasCustomerTitle

`func (o *ThirdPartyParams) HasCustomerTitle() bool`

HasCustomerTitle returns a boolean if a field has been set.

### SetCustomerTitleNil

`func (o *ThirdPartyParams) SetCustomerTitleNil(b bool)`

 SetCustomerTitleNil sets the value for CustomerTitle to be an explicit nil

### UnsetCustomerTitle
`func (o *ThirdPartyParams) UnsetCustomerTitle()`

UnsetCustomerTitle ensures that no value is present for CustomerTitle, not even an explicit nil
### GetProviderId

`func (o *ThirdPartyParams) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ThirdPartyParams) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ThirdPartyParams) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ThirdPartyParams) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ThirdPartyParams) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ThirdPartyParams) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderKey

`func (o *ThirdPartyParams) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *ThirdPartyParams) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *ThirdPartyParams) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *ThirdPartyParams) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *ThirdPartyParams) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *ThirdPartyParams) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


