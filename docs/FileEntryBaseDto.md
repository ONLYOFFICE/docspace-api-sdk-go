# FileEntryBaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | The file entry title. | [optional] 
**Access** | Pointer to [**FileShare**](FileShare.md) |  | [optional] 
**SharedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) |  | [optional] 
**OwnedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) |  | [optional] 
**Shared** | Pointer to **bool** | Specifies if the file entry is shared via link or not. | [optional] 
**SharedForUser** | Pointer to **bool** | Specifies if the file entry is shared for user or not. | [optional] 
**SharedExternal** | Pointer to **bool** | Specifies if the file entry is shared via a public (non-internal) external link. | [optional] 
**ParentShared** | Pointer to **bool** | Indicates whether the parent entity is shared. | [optional] 
**ShortWebUrl** | Pointer to **NullableString** | The short Web URL. | [optional] 
**Created** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**CreatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) |  | [optional] 
**Updated** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**AutoDelete** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**RootFolderType** | Pointer to [**FolderType**](FolderType.md) |  | [optional] 
**ParentRoomType** | Pointer to [**FolderType**](FolderType.md) |  | [optional] 
**UpdatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) |  | [optional] 
**ProviderItem** | Pointer to **NullableBool** | Specifies if the file entry provider is specified or not. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The provider key of the file entry. | [optional] 
**ProviderId** | Pointer to **NullableInt32** | The provider ID of the file entry. | [optional] 
**Order** | Pointer to **NullableString** | The order of the file entry. | [optional] 
**IsFavorite** | Pointer to **NullableBool** | Specifies if the file is a favorite or not. | [optional] 
**FileEntryType** | Pointer to [**FileEntryType**](FileEntryType.md) |  | [optional] 

## Methods

### NewFileEntryBaseDto

`func NewFileEntryBaseDto() *FileEntryBaseDto`

NewFileEntryBaseDto instantiates a new FileEntryBaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileEntryBaseDtoWithDefaults

`func NewFileEntryBaseDtoWithDefaults() *FileEntryBaseDto`

NewFileEntryBaseDtoWithDefaults instantiates a new FileEntryBaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FileEntryBaseDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FileEntryBaseDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FileEntryBaseDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FileEntryBaseDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FileEntryBaseDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FileEntryBaseDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAccess

`func (o *FileEntryBaseDto) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FileEntryBaseDto) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FileEntryBaseDto) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FileEntryBaseDto) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *FileEntryBaseDto) GetSharedBy() EmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *FileEntryBaseDto) GetSharedByOk() (*EmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *FileEntryBaseDto) SetSharedBy(v EmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *FileEntryBaseDto) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *FileEntryBaseDto) GetOwnedBy() EmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *FileEntryBaseDto) GetOwnedByOk() (*EmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *FileEntryBaseDto) SetOwnedBy(v EmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *FileEntryBaseDto) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *FileEntryBaseDto) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *FileEntryBaseDto) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *FileEntryBaseDto) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *FileEntryBaseDto) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *FileEntryBaseDto) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *FileEntryBaseDto) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *FileEntryBaseDto) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *FileEntryBaseDto) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *FileEntryBaseDto) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *FileEntryBaseDto) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *FileEntryBaseDto) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *FileEntryBaseDto) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *FileEntryBaseDto) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *FileEntryBaseDto) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *FileEntryBaseDto) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *FileEntryBaseDto) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *FileEntryBaseDto) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *FileEntryBaseDto) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *FileEntryBaseDto) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *FileEntryBaseDto) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### SetShortWebUrlNil

`func (o *FileEntryBaseDto) SetShortWebUrlNil(b bool)`

 SetShortWebUrlNil sets the value for ShortWebUrl to be an explicit nil

### UnsetShortWebUrl
`func (o *FileEntryBaseDto) UnsetShortWebUrl()`

UnsetShortWebUrl ensures that no value is present for ShortWebUrl, not even an explicit nil
### GetCreated

`func (o *FileEntryBaseDto) GetCreated() ApiDateTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FileEntryBaseDto) GetCreatedOk() (*ApiDateTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FileEntryBaseDto) SetCreated(v ApiDateTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FileEntryBaseDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FileEntryBaseDto) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FileEntryBaseDto) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FileEntryBaseDto) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FileEntryBaseDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *FileEntryBaseDto) GetUpdated() ApiDateTime`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FileEntryBaseDto) GetUpdatedOk() (*ApiDateTime, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FileEntryBaseDto) SetUpdated(v ApiDateTime)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FileEntryBaseDto) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAutoDelete

`func (o *FileEntryBaseDto) GetAutoDelete() ApiDateTime`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *FileEntryBaseDto) GetAutoDeleteOk() (*ApiDateTime, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *FileEntryBaseDto) SetAutoDelete(v ApiDateTime)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *FileEntryBaseDto) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetRootFolderType

`func (o *FileEntryBaseDto) GetRootFolderType() FolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *FileEntryBaseDto) GetRootFolderTypeOk() (*FolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *FileEntryBaseDto) SetRootFolderType(v FolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *FileEntryBaseDto) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *FileEntryBaseDto) GetParentRoomType() FolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *FileEntryBaseDto) GetParentRoomTypeOk() (*FolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *FileEntryBaseDto) SetParentRoomType(v FolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *FileEntryBaseDto) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FileEntryBaseDto) GetUpdatedBy() EmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FileEntryBaseDto) GetUpdatedByOk() (*EmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FileEntryBaseDto) SetUpdatedBy(v EmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FileEntryBaseDto) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *FileEntryBaseDto) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *FileEntryBaseDto) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *FileEntryBaseDto) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *FileEntryBaseDto) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### SetProviderItemNil

`func (o *FileEntryBaseDto) SetProviderItemNil(b bool)`

 SetProviderItemNil sets the value for ProviderItem to be an explicit nil

### UnsetProviderItem
`func (o *FileEntryBaseDto) UnsetProviderItem()`

UnsetProviderItem ensures that no value is present for ProviderItem, not even an explicit nil
### GetProviderKey

`func (o *FileEntryBaseDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *FileEntryBaseDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *FileEntryBaseDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *FileEntryBaseDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *FileEntryBaseDto) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *FileEntryBaseDto) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetProviderId

`func (o *FileEntryBaseDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *FileEntryBaseDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *FileEntryBaseDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *FileEntryBaseDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *FileEntryBaseDto) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *FileEntryBaseDto) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetOrder

`func (o *FileEntryBaseDto) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FileEntryBaseDto) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FileEntryBaseDto) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FileEntryBaseDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *FileEntryBaseDto) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *FileEntryBaseDto) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetIsFavorite

`func (o *FileEntryBaseDto) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *FileEntryBaseDto) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *FileEntryBaseDto) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *FileEntryBaseDto) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *FileEntryBaseDto) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *FileEntryBaseDto) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetFileEntryType

`func (o *FileEntryBaseDto) GetFileEntryType() FileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *FileEntryBaseDto) GetFileEntryTypeOk() (*FileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *FileEntryBaseDto) SetFileEntryType(v FileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *FileEntryBaseDto) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


