# AiFileEntryBaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | The file entry title. | [optional] 
**Access** | Pointer to [**AiFileShare**](AiFileShare.md) | The access rights to the file entry. | [optional] 
**SharedBy** | Pointer to [**AiEmployeeDto**](AiEmployeeDto.md) | Provides information about the employee who shared the file or folder. | [optional] 
**OwnedBy** | Pointer to [**AiEmployeeDto**](AiEmployeeDto.md) | The information about the employee who owns the file entry. | [optional] 
**Shared** | Pointer to **bool** | Specifies if the file entry is shared via link or not. | [optional] 
**SharedForUser** | Pointer to **bool** | Specifies if the file entry is shared for user or not. | [optional] 
**SharedExternal** | Pointer to **bool** | Specifies if the file entry is shared via a public (non-internal) external link. | [optional] 
**ParentShared** | Pointer to **bool** | Indicates whether the parent entity is shared. | [optional] 
**ShortWebUrl** | Pointer to **NullableString** | The short Web URL. | [optional] 
**Created** | Pointer to **NullableTime** | The creation date and time of the file entry. | [optional] 
**CreatedBy** | Pointer to [**AiEmployeeDto**](AiEmployeeDto.md) | The file entry author. | [optional] 
**Updated** | Pointer to **NullableTime** | The last date and time when the file entry was updated. | [optional] 
**AutoDelete** | Pointer to **NullableTime** | The date and time when the file entry will be automatically deleted. | [optional] 
**RootFolderType** | Pointer to [**AiFolderType**](AiFolderType.md) | The root folder type of the file entry. | [optional] 
**ParentRoomType** | Pointer to [**AiFolderType**](AiFolderType.md) | The parent room type of the file entry. | [optional] 
**UpdatedBy** | Pointer to [**AiEmployeeDto**](AiEmployeeDto.md) | The user who updated the file entry. | [optional] 
**ProviderItem** | Pointer to **NullableBool** | Specifies if the file entry provider is specified or not. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The provider key of the file entry. | [optional] 
**ProviderId** | Pointer to **NullableInt32** | The provider ID of the file entry. | [optional] 
**Order** | Pointer to **NullableString** | The order of the file entry. | [optional] 
**IsFavorite** | Pointer to **NullableBool** | Specifies if the file is a favorite or not. | [optional] 
**FileEntryType** | Pointer to [**AiFileEntryType**](AiFileEntryType.md) | The file entry type. | [optional] 

## Methods

### NewAiFileEntryBaseDto

`func NewAiFileEntryBaseDto() *AiFileEntryBaseDto`

NewAiFileEntryBaseDto instantiates a new AiFileEntryBaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiFileEntryBaseDtoWithDefaults

`func NewAiFileEntryBaseDtoWithDefaults() *AiFileEntryBaseDto`

NewAiFileEntryBaseDtoWithDefaults instantiates a new AiFileEntryBaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AiFileEntryBaseDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiFileEntryBaseDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiFileEntryBaseDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AiFileEntryBaseDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AiFileEntryBaseDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AiFileEntryBaseDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAccess

`func (o *AiFileEntryBaseDto) GetAccess() AiFileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AiFileEntryBaseDto) GetAccessOk() (*AiFileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AiFileEntryBaseDto) SetAccess(v AiFileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AiFileEntryBaseDto) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *AiFileEntryBaseDto) GetSharedBy() AiEmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *AiFileEntryBaseDto) GetSharedByOk() (*AiEmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *AiFileEntryBaseDto) SetSharedBy(v AiEmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *AiFileEntryBaseDto) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *AiFileEntryBaseDto) GetOwnedBy() AiEmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *AiFileEntryBaseDto) GetOwnedByOk() (*AiEmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *AiFileEntryBaseDto) SetOwnedBy(v AiEmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *AiFileEntryBaseDto) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *AiFileEntryBaseDto) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *AiFileEntryBaseDto) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *AiFileEntryBaseDto) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *AiFileEntryBaseDto) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *AiFileEntryBaseDto) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *AiFileEntryBaseDto) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *AiFileEntryBaseDto) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *AiFileEntryBaseDto) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *AiFileEntryBaseDto) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *AiFileEntryBaseDto) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *AiFileEntryBaseDto) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *AiFileEntryBaseDto) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *AiFileEntryBaseDto) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *AiFileEntryBaseDto) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *AiFileEntryBaseDto) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *AiFileEntryBaseDto) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *AiFileEntryBaseDto) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *AiFileEntryBaseDto) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *AiFileEntryBaseDto) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *AiFileEntryBaseDto) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### SetShortWebUrlNil

`func (o *AiFileEntryBaseDto) SetShortWebUrlNil(b bool)`

 SetShortWebUrlNil sets the value for ShortWebUrl to be an explicit nil

### UnsetShortWebUrl
`func (o *AiFileEntryBaseDto) UnsetShortWebUrl()`

UnsetShortWebUrl ensures that no value is present for ShortWebUrl, not even an explicit nil
### GetCreated

`func (o *AiFileEntryBaseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiFileEntryBaseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiFileEntryBaseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AiFileEntryBaseDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *AiFileEntryBaseDto) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *AiFileEntryBaseDto) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetCreatedBy

`func (o *AiFileEntryBaseDto) GetCreatedBy() AiEmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AiFileEntryBaseDto) GetCreatedByOk() (*AiEmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AiFileEntryBaseDto) SetCreatedBy(v AiEmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AiFileEntryBaseDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *AiFileEntryBaseDto) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AiFileEntryBaseDto) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AiFileEntryBaseDto) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AiFileEntryBaseDto) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *AiFileEntryBaseDto) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *AiFileEntryBaseDto) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetAutoDelete

`func (o *AiFileEntryBaseDto) GetAutoDelete() time.Time`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *AiFileEntryBaseDto) GetAutoDeleteOk() (*time.Time, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *AiFileEntryBaseDto) SetAutoDelete(v time.Time)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *AiFileEntryBaseDto) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### SetAutoDeleteNil

`func (o *AiFileEntryBaseDto) SetAutoDeleteNil(b bool)`

 SetAutoDeleteNil sets the value for AutoDelete to be an explicit nil

### UnsetAutoDelete
`func (o *AiFileEntryBaseDto) UnsetAutoDelete()`

UnsetAutoDelete ensures that no value is present for AutoDelete, not even an explicit nil
### GetRootFolderType

`func (o *AiFileEntryBaseDto) GetRootFolderType() AiFolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *AiFileEntryBaseDto) GetRootFolderTypeOk() (*AiFolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *AiFileEntryBaseDto) SetRootFolderType(v AiFolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *AiFileEntryBaseDto) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *AiFileEntryBaseDto) GetParentRoomType() AiFolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *AiFileEntryBaseDto) GetParentRoomTypeOk() (*AiFolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *AiFileEntryBaseDto) SetParentRoomType(v AiFolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *AiFileEntryBaseDto) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AiFileEntryBaseDto) GetUpdatedBy() AiEmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AiFileEntryBaseDto) GetUpdatedByOk() (*AiEmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AiFileEntryBaseDto) SetUpdatedBy(v AiEmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AiFileEntryBaseDto) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *AiFileEntryBaseDto) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *AiFileEntryBaseDto) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *AiFileEntryBaseDto) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *AiFileEntryBaseDto) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### SetProviderItemNil

`func (o *AiFileEntryBaseDto) SetProviderItemNil(b bool)`

 SetProviderItemNil sets the value for ProviderItem to be an explicit nil

### UnsetProviderItem
`func (o *AiFileEntryBaseDto) UnsetProviderItem()`

UnsetProviderItem ensures that no value is present for ProviderItem, not even an explicit nil
### GetProviderKey

`func (o *AiFileEntryBaseDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *AiFileEntryBaseDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *AiFileEntryBaseDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *AiFileEntryBaseDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *AiFileEntryBaseDto) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *AiFileEntryBaseDto) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetProviderId

`func (o *AiFileEntryBaseDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *AiFileEntryBaseDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *AiFileEntryBaseDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *AiFileEntryBaseDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *AiFileEntryBaseDto) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *AiFileEntryBaseDto) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetOrder

`func (o *AiFileEntryBaseDto) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AiFileEntryBaseDto) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AiFileEntryBaseDto) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AiFileEntryBaseDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *AiFileEntryBaseDto) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *AiFileEntryBaseDto) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetIsFavorite

`func (o *AiFileEntryBaseDto) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *AiFileEntryBaseDto) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *AiFileEntryBaseDto) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *AiFileEntryBaseDto) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *AiFileEntryBaseDto) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *AiFileEntryBaseDto) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetFileEntryType

`func (o *AiFileEntryBaseDto) GetFileEntryType() AiFileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *AiFileEntryBaseDto) GetFileEntryTypeOk() (*AiFileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *AiFileEntryBaseDto) SetFileEntryType(v AiFileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *AiFileEntryBaseDto) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


