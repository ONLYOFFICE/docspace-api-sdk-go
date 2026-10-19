# ThirdPartyAccountDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Corporate** | Pointer to **bool** | Whether the account is attached to the legacy Common section, which is the case only for accounts inherited  from an older portal. | [optional] 
**RoomsStorage** | Pointer to **bool** | Whether the account is attached to the Rooms section, room templates and the archive counted in. This is where  `POST api/2.0/files/thirdparty` puts every account it connects. | [optional] 
**CustomerTitle** | Pointer to **NullableString** | The name the account is shown under in the portal, as it was saved when the account was connected. | [optional] 
**ProviderId** | Pointer to **NullableInt32** | The account ID to send to `DELETE api/2.0/files/thirdparty/{providerId}`, or as `providerId` to  re-authenticate the account. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The storage service behind the account. `WebDav` stands for every WebDAV preset, so it does not tell which of  them was chosen when the account was connected. | [optional] 

## Methods

### NewThirdPartyAccountDto

`func NewThirdPartyAccountDto() *ThirdPartyAccountDto`

NewThirdPartyAccountDto instantiates a new ThirdPartyAccountDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyAccountDtoWithDefaults

`func NewThirdPartyAccountDtoWithDefaults() *ThirdPartyAccountDto`

NewThirdPartyAccountDtoWithDefaults instantiates a new ThirdPartyAccountDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorporate

`func (o *ThirdPartyAccountDto) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *ThirdPartyAccountDto) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *ThirdPartyAccountDto) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *ThirdPartyAccountDto) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetRoomsStorage

`func (o *ThirdPartyAccountDto) GetRoomsStorage() bool`

GetRoomsStorage returns the RoomsStorage field if non-nil, zero value otherwise.

### GetRoomsStorageOk

`func (o *ThirdPartyAccountDto) GetRoomsStorageOk() (*bool, bool)`

GetRoomsStorageOk returns a tuple with the RoomsStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsStorage

`func (o *ThirdPartyAccountDto) SetRoomsStorage(v bool)`

SetRoomsStorage sets RoomsStorage field to given value.

### HasRoomsStorage

`func (o *ThirdPartyAccountDto) HasRoomsStorage() bool`

HasRoomsStorage returns a boolean if a field has been set.

### GetCustomerTitle

`func (o *ThirdPartyAccountDto) GetCustomerTitle() string`

GetCustomerTitle returns the CustomerTitle field if non-nil, zero value otherwise.

### GetCustomerTitleOk

`func (o *ThirdPartyAccountDto) GetCustomerTitleOk() (*string, bool)`

GetCustomerTitleOk returns a tuple with the CustomerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTitle

`func (o *ThirdPartyAccountDto) SetCustomerTitle(v string)`

SetCustomerTitle sets CustomerTitle field to given value.

### HasCustomerTitle

`func (o *ThirdPartyAccountDto) HasCustomerTitle() bool`

HasCustomerTitle returns a boolean if a field has been set.

### SetCustomerTitleNil

`func (o *ThirdPartyAccountDto) SetCustomerTitleNil(b bool)`

 SetCustomerTitleNil sets the value for CustomerTitle to be an explicit nil

### UnsetCustomerTitle
`func (o *ThirdPartyAccountDto) UnsetCustomerTitle()`

UnsetCustomerTitle ensures that no value is present for CustomerTitle, not even an explicit nil
### GetProviderId

`func (o *ThirdPartyAccountDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ThirdPartyAccountDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ThirdPartyAccountDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ThirdPartyAccountDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ThirdPartyAccountDto) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ThirdPartyAccountDto) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderKey

`func (o *ThirdPartyAccountDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *ThirdPartyAccountDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *ThirdPartyAccountDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *ThirdPartyAccountDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *ThirdPartyAccountDto) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *ThirdPartyAccountDto) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


