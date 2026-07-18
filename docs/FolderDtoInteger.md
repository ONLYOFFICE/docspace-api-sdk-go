# FolderDtoInteger

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
**Id** | Pointer to **int32** | The file entry ID. | [optional] 
**RootFolderId** | Pointer to **int32** | The root folder ID of the file entry. | [optional] 
**OriginId** | Pointer to **int32** | The origin ID of the file entry. | [optional] 
**OriginRoomId** | Pointer to **int32** | The origin room ID of the file entry. | [optional] 
**OriginTitle** | Pointer to **NullableString** | The origin title of the file entry. | [optional] 
**OriginRoomTitle** | Pointer to **NullableString** | The origin room title of the file entry. | [optional] 
**CanShare** | Pointer to **bool** | Specifies if the file entry can be shared or not. | [optional] 
**ShareSettings** | Pointer to [**NullableFileEntryDtoIntegerAllOfShareSettings**](FileEntryDtoIntegerAllOfShareSettings.md) |  | [optional] 
**Security** | Pointer to [**NullableFileEntryDtoIntegerAllOfSecurity**](FileEntryDtoIntegerAllOfSecurity.md) |  | [optional] 
**AvailableShareRights** | Pointer to [**NullableFileEntryDtoIntegerAllOfAvailableShareRights**](FileEntryDtoIntegerAllOfAvailableShareRights.md) |  | [optional] 
**RequestToken** | Pointer to **NullableString** | The request token of the file entry. | [optional] 
**External** | Pointer to **NullableBool** | Specifies if the folder can be accessed via an external link or not. | [optional] 
**ExpirationDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**IsLinkExpired** | Pointer to **NullableBool** | Indicates whether the shareable link associated with the file or folder has expired. | [optional] 
**ParentId** | Pointer to **int32** | The parent folder ID of the folder. | [optional] 
**FilesCount** | Pointer to **int32** | The number of files that the folder contains. | [optional] 
**FoldersCount** | Pointer to **int32** | The number of folders that the folder contains. | [optional] 
**IsShareable** | Pointer to **NullableBool** | Specifies if the folder can be shared or not. | [optional] 
**New** | Pointer to **int32** | The new element index in the folder. | [optional] 
**Mute** | Pointer to **bool** | Specifies if the folder notifications are enabled or not. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of the folder. | [optional] 
**Logo** | Pointer to [**Logo**](Logo.md) |  | [optional] 
**Pinned** | Pointer to **bool** | Specifies if the folder is pinned or not. | [optional] 
**RoomType** | Pointer to [**RoomType**](RoomType.md) |  | [optional] 
**Private** | Pointer to **bool** | Specifies if the folder is private or not. | [optional] 
**Indexing** | Pointer to **bool** | Specifies if the folder is indexed or not. | [optional] 
**DenyDownload** | Pointer to **bool** | Specifies if the folder can be downloaded or not. | [optional] 
**Lifetime** | Pointer to [**RoomDataLifetimeDto**](RoomDataLifetimeDto.md) |  | [optional] 
**Watermark** | Pointer to [**WatermarkDto**](WatermarkDto.md) |  | [optional] 
**Type** | Pointer to [**FolderType**](FolderType.md) |  | [optional] 
**InRoom** | Pointer to **NullableBool** | Specifies if the folder is placed in the room or not. | [optional] 
**QuotaLimit** | Pointer to **NullableInt64** | The folder quota limit. | [optional] 
**IsCustomQuota** | Pointer to **NullableBool** | Specifies if the folder room has a custom quota or not. | [optional] 
**UsedSpace** | Pointer to **NullableInt64** | How much folder space is used (counter). | [optional] 
**PasswordProtected** | Pointer to **NullableBool** | Specifies if the folder is password protected or not. | [optional] 
**Expired** | Pointer to **NullableBool** | Specifies if an external link to the folder is expired or not. | [optional] 
**ChatSettings** | Pointer to [**ChatSettingsDto**](ChatSettingsDto.md) |  | [optional] 
**RootRoomType** | Pointer to [**RoomType**](RoomType.md) |  | [optional] 
**SaveFormAsXLSX** | Pointer to **NullableBool** | Specifies whether to save form data as XLSX file. | [optional] 
**SendFormToExternalDB** | Pointer to **NullableBool** | Specifies whether to send form data to external database. | [optional] 
**OriginalFormId** | Pointer to **NullableInt32** | The original form ID that corresponds to this FormFillingFolderDone folder. | [optional] 

## Methods

### NewFolderDtoInteger

`func NewFolderDtoInteger() *FolderDtoInteger`

NewFolderDtoInteger instantiates a new FolderDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderDtoIntegerWithDefaults

`func NewFolderDtoIntegerWithDefaults() *FolderDtoInteger`

NewFolderDtoIntegerWithDefaults instantiates a new FolderDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FolderDtoInteger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FolderDtoInteger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FolderDtoInteger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FolderDtoInteger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FolderDtoInteger) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FolderDtoInteger) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAccess

`func (o *FolderDtoInteger) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FolderDtoInteger) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FolderDtoInteger) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FolderDtoInteger) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *FolderDtoInteger) GetSharedBy() EmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *FolderDtoInteger) GetSharedByOk() (*EmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *FolderDtoInteger) SetSharedBy(v EmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *FolderDtoInteger) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *FolderDtoInteger) GetOwnedBy() EmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *FolderDtoInteger) GetOwnedByOk() (*EmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *FolderDtoInteger) SetOwnedBy(v EmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *FolderDtoInteger) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *FolderDtoInteger) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *FolderDtoInteger) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *FolderDtoInteger) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *FolderDtoInteger) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *FolderDtoInteger) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *FolderDtoInteger) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *FolderDtoInteger) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *FolderDtoInteger) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *FolderDtoInteger) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *FolderDtoInteger) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *FolderDtoInteger) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *FolderDtoInteger) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *FolderDtoInteger) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *FolderDtoInteger) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *FolderDtoInteger) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *FolderDtoInteger) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *FolderDtoInteger) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *FolderDtoInteger) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *FolderDtoInteger) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *FolderDtoInteger) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### SetShortWebUrlNil

`func (o *FolderDtoInteger) SetShortWebUrlNil(b bool)`

 SetShortWebUrlNil sets the value for ShortWebUrl to be an explicit nil

### UnsetShortWebUrl
`func (o *FolderDtoInteger) UnsetShortWebUrl()`

UnsetShortWebUrl ensures that no value is present for ShortWebUrl, not even an explicit nil
### GetCreated

`func (o *FolderDtoInteger) GetCreated() ApiDateTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FolderDtoInteger) GetCreatedOk() (*ApiDateTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FolderDtoInteger) SetCreated(v ApiDateTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FolderDtoInteger) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FolderDtoInteger) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FolderDtoInteger) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FolderDtoInteger) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FolderDtoInteger) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *FolderDtoInteger) GetUpdated() ApiDateTime`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FolderDtoInteger) GetUpdatedOk() (*ApiDateTime, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FolderDtoInteger) SetUpdated(v ApiDateTime)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FolderDtoInteger) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAutoDelete

`func (o *FolderDtoInteger) GetAutoDelete() ApiDateTime`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *FolderDtoInteger) GetAutoDeleteOk() (*ApiDateTime, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *FolderDtoInteger) SetAutoDelete(v ApiDateTime)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *FolderDtoInteger) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetRootFolderType

`func (o *FolderDtoInteger) GetRootFolderType() FolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *FolderDtoInteger) GetRootFolderTypeOk() (*FolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *FolderDtoInteger) SetRootFolderType(v FolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *FolderDtoInteger) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *FolderDtoInteger) GetParentRoomType() FolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *FolderDtoInteger) GetParentRoomTypeOk() (*FolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *FolderDtoInteger) SetParentRoomType(v FolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *FolderDtoInteger) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FolderDtoInteger) GetUpdatedBy() EmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FolderDtoInteger) GetUpdatedByOk() (*EmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FolderDtoInteger) SetUpdatedBy(v EmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FolderDtoInteger) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *FolderDtoInteger) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *FolderDtoInteger) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *FolderDtoInteger) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *FolderDtoInteger) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### SetProviderItemNil

`func (o *FolderDtoInteger) SetProviderItemNil(b bool)`

 SetProviderItemNil sets the value for ProviderItem to be an explicit nil

### UnsetProviderItem
`func (o *FolderDtoInteger) UnsetProviderItem()`

UnsetProviderItem ensures that no value is present for ProviderItem, not even an explicit nil
### GetProviderKey

`func (o *FolderDtoInteger) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *FolderDtoInteger) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *FolderDtoInteger) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *FolderDtoInteger) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *FolderDtoInteger) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *FolderDtoInteger) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetProviderId

`func (o *FolderDtoInteger) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *FolderDtoInteger) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *FolderDtoInteger) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *FolderDtoInteger) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *FolderDtoInteger) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *FolderDtoInteger) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetOrder

`func (o *FolderDtoInteger) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FolderDtoInteger) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FolderDtoInteger) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FolderDtoInteger) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *FolderDtoInteger) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *FolderDtoInteger) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetIsFavorite

`func (o *FolderDtoInteger) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *FolderDtoInteger) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *FolderDtoInteger) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *FolderDtoInteger) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *FolderDtoInteger) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *FolderDtoInteger) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetFileEntryType

`func (o *FolderDtoInteger) GetFileEntryType() FileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *FolderDtoInteger) GetFileEntryTypeOk() (*FileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *FolderDtoInteger) SetFileEntryType(v FileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *FolderDtoInteger) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.

### GetId

`func (o *FolderDtoInteger) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FolderDtoInteger) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FolderDtoInteger) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FolderDtoInteger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootFolderId

`func (o *FolderDtoInteger) GetRootFolderId() int32`

GetRootFolderId returns the RootFolderId field if non-nil, zero value otherwise.

### GetRootFolderIdOk

`func (o *FolderDtoInteger) GetRootFolderIdOk() (*int32, bool)`

GetRootFolderIdOk returns a tuple with the RootFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderId

`func (o *FolderDtoInteger) SetRootFolderId(v int32)`

SetRootFolderId sets RootFolderId field to given value.

### HasRootFolderId

`func (o *FolderDtoInteger) HasRootFolderId() bool`

HasRootFolderId returns a boolean if a field has been set.

### GetOriginId

`func (o *FolderDtoInteger) GetOriginId() int32`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *FolderDtoInteger) GetOriginIdOk() (*int32, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *FolderDtoInteger) SetOriginId(v int32)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *FolderDtoInteger) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginRoomId

`func (o *FolderDtoInteger) GetOriginRoomId() int32`

GetOriginRoomId returns the OriginRoomId field if non-nil, zero value otherwise.

### GetOriginRoomIdOk

`func (o *FolderDtoInteger) GetOriginRoomIdOk() (*int32, bool)`

GetOriginRoomIdOk returns a tuple with the OriginRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomId

`func (o *FolderDtoInteger) SetOriginRoomId(v int32)`

SetOriginRoomId sets OriginRoomId field to given value.

### HasOriginRoomId

`func (o *FolderDtoInteger) HasOriginRoomId() bool`

HasOriginRoomId returns a boolean if a field has been set.

### GetOriginTitle

`func (o *FolderDtoInteger) GetOriginTitle() string`

GetOriginTitle returns the OriginTitle field if non-nil, zero value otherwise.

### GetOriginTitleOk

`func (o *FolderDtoInteger) GetOriginTitleOk() (*string, bool)`

GetOriginTitleOk returns a tuple with the OriginTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTitle

`func (o *FolderDtoInteger) SetOriginTitle(v string)`

SetOriginTitle sets OriginTitle field to given value.

### HasOriginTitle

`func (o *FolderDtoInteger) HasOriginTitle() bool`

HasOriginTitle returns a boolean if a field has been set.

### SetOriginTitleNil

`func (o *FolderDtoInteger) SetOriginTitleNil(b bool)`

 SetOriginTitleNil sets the value for OriginTitle to be an explicit nil

### UnsetOriginTitle
`func (o *FolderDtoInteger) UnsetOriginTitle()`

UnsetOriginTitle ensures that no value is present for OriginTitle, not even an explicit nil
### GetOriginRoomTitle

`func (o *FolderDtoInteger) GetOriginRoomTitle() string`

GetOriginRoomTitle returns the OriginRoomTitle field if non-nil, zero value otherwise.

### GetOriginRoomTitleOk

`func (o *FolderDtoInteger) GetOriginRoomTitleOk() (*string, bool)`

GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomTitle

`func (o *FolderDtoInteger) SetOriginRoomTitle(v string)`

SetOriginRoomTitle sets OriginRoomTitle field to given value.

### HasOriginRoomTitle

`func (o *FolderDtoInteger) HasOriginRoomTitle() bool`

HasOriginRoomTitle returns a boolean if a field has been set.

### SetOriginRoomTitleNil

`func (o *FolderDtoInteger) SetOriginRoomTitleNil(b bool)`

 SetOriginRoomTitleNil sets the value for OriginRoomTitle to be an explicit nil

### UnsetOriginRoomTitle
`func (o *FolderDtoInteger) UnsetOriginRoomTitle()`

UnsetOriginRoomTitle ensures that no value is present for OriginRoomTitle, not even an explicit nil
### GetCanShare

`func (o *FolderDtoInteger) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *FolderDtoInteger) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *FolderDtoInteger) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *FolderDtoInteger) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### GetShareSettings

`func (o *FolderDtoInteger) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings`

GetShareSettings returns the ShareSettings field if non-nil, zero value otherwise.

### GetShareSettingsOk

`func (o *FolderDtoInteger) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool)`

GetShareSettingsOk returns a tuple with the ShareSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareSettings

`func (o *FolderDtoInteger) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings)`

SetShareSettings sets ShareSettings field to given value.

### HasShareSettings

`func (o *FolderDtoInteger) HasShareSettings() bool`

HasShareSettings returns a boolean if a field has been set.

### SetShareSettingsNil

`func (o *FolderDtoInteger) SetShareSettingsNil(b bool)`

 SetShareSettingsNil sets the value for ShareSettings to be an explicit nil

### UnsetShareSettings
`func (o *FolderDtoInteger) UnsetShareSettings()`

UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
### GetSecurity

`func (o *FolderDtoInteger) GetSecurity() FileEntryDtoIntegerAllOfSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *FolderDtoInteger) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *FolderDtoInteger) SetSecurity(v FileEntryDtoIntegerAllOfSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *FolderDtoInteger) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *FolderDtoInteger) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *FolderDtoInteger) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetAvailableShareRights

`func (o *FolderDtoInteger) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights`

GetAvailableShareRights returns the AvailableShareRights field if non-nil, zero value otherwise.

### GetAvailableShareRightsOk

`func (o *FolderDtoInteger) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool)`

GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShareRights

`func (o *FolderDtoInteger) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights)`

SetAvailableShareRights sets AvailableShareRights field to given value.

### HasAvailableShareRights

`func (o *FolderDtoInteger) HasAvailableShareRights() bool`

HasAvailableShareRights returns a boolean if a field has been set.

### SetAvailableShareRightsNil

`func (o *FolderDtoInteger) SetAvailableShareRightsNil(b bool)`

 SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil

### UnsetAvailableShareRights
`func (o *FolderDtoInteger) UnsetAvailableShareRights()`

UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
### GetRequestToken

`func (o *FolderDtoInteger) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *FolderDtoInteger) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *FolderDtoInteger) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *FolderDtoInteger) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### SetRequestTokenNil

`func (o *FolderDtoInteger) SetRequestTokenNil(b bool)`

 SetRequestTokenNil sets the value for RequestToken to be an explicit nil

### UnsetRequestToken
`func (o *FolderDtoInteger) UnsetRequestToken()`

UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
### GetExternal

`func (o *FolderDtoInteger) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *FolderDtoInteger) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *FolderDtoInteger) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *FolderDtoInteger) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### SetExternalNil

`func (o *FolderDtoInteger) SetExternalNil(b bool)`

 SetExternalNil sets the value for External to be an explicit nil

### UnsetExternal
`func (o *FolderDtoInteger) UnsetExternal()`

UnsetExternal ensures that no value is present for External, not even an explicit nil
### GetExpirationDate

`func (o *FolderDtoInteger) GetExpirationDate() ApiDateTime`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FolderDtoInteger) GetExpirationDateOk() (*ApiDateTime, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FolderDtoInteger) SetExpirationDate(v ApiDateTime)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FolderDtoInteger) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIsLinkExpired

`func (o *FolderDtoInteger) GetIsLinkExpired() bool`

GetIsLinkExpired returns the IsLinkExpired field if non-nil, zero value otherwise.

### GetIsLinkExpiredOk

`func (o *FolderDtoInteger) GetIsLinkExpiredOk() (*bool, bool)`

GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkExpired

`func (o *FolderDtoInteger) SetIsLinkExpired(v bool)`

SetIsLinkExpired sets IsLinkExpired field to given value.

### HasIsLinkExpired

`func (o *FolderDtoInteger) HasIsLinkExpired() bool`

HasIsLinkExpired returns a boolean if a field has been set.

### SetIsLinkExpiredNil

`func (o *FolderDtoInteger) SetIsLinkExpiredNil(b bool)`

 SetIsLinkExpiredNil sets the value for IsLinkExpired to be an explicit nil

### UnsetIsLinkExpired
`func (o *FolderDtoInteger) UnsetIsLinkExpired()`

UnsetIsLinkExpired ensures that no value is present for IsLinkExpired, not even an explicit nil
### GetParentId

`func (o *FolderDtoInteger) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FolderDtoInteger) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FolderDtoInteger) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FolderDtoInteger) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetFilesCount

`func (o *FolderDtoInteger) GetFilesCount() int32`

GetFilesCount returns the FilesCount field if non-nil, zero value otherwise.

### GetFilesCountOk

`func (o *FolderDtoInteger) GetFilesCountOk() (*int32, bool)`

GetFilesCountOk returns a tuple with the FilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesCount

`func (o *FolderDtoInteger) SetFilesCount(v int32)`

SetFilesCount sets FilesCount field to given value.

### HasFilesCount

`func (o *FolderDtoInteger) HasFilesCount() bool`

HasFilesCount returns a boolean if a field has been set.

### GetFoldersCount

`func (o *FolderDtoInteger) GetFoldersCount() int32`

GetFoldersCount returns the FoldersCount field if non-nil, zero value otherwise.

### GetFoldersCountOk

`func (o *FolderDtoInteger) GetFoldersCountOk() (*int32, bool)`

GetFoldersCountOk returns a tuple with the FoldersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldersCount

`func (o *FolderDtoInteger) SetFoldersCount(v int32)`

SetFoldersCount sets FoldersCount field to given value.

### HasFoldersCount

`func (o *FolderDtoInteger) HasFoldersCount() bool`

HasFoldersCount returns a boolean if a field has been set.

### GetIsShareable

`func (o *FolderDtoInteger) GetIsShareable() bool`

GetIsShareable returns the IsShareable field if non-nil, zero value otherwise.

### GetIsShareableOk

`func (o *FolderDtoInteger) GetIsShareableOk() (*bool, bool)`

GetIsShareableOk returns a tuple with the IsShareable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShareable

`func (o *FolderDtoInteger) SetIsShareable(v bool)`

SetIsShareable sets IsShareable field to given value.

### HasIsShareable

`func (o *FolderDtoInteger) HasIsShareable() bool`

HasIsShareable returns a boolean if a field has been set.

### SetIsShareableNil

`func (o *FolderDtoInteger) SetIsShareableNil(b bool)`

 SetIsShareableNil sets the value for IsShareable to be an explicit nil

### UnsetIsShareable
`func (o *FolderDtoInteger) UnsetIsShareable()`

UnsetIsShareable ensures that no value is present for IsShareable, not even an explicit nil
### GetNew

`func (o *FolderDtoInteger) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *FolderDtoInteger) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *FolderDtoInteger) SetNew(v int32)`

SetNew sets New field to given value.

### HasNew

`func (o *FolderDtoInteger) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetMute

`func (o *FolderDtoInteger) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *FolderDtoInteger) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *FolderDtoInteger) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *FolderDtoInteger) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetTags

`func (o *FolderDtoInteger) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FolderDtoInteger) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FolderDtoInteger) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FolderDtoInteger) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *FolderDtoInteger) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *FolderDtoInteger) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetLogo

`func (o *FolderDtoInteger) GetLogo() Logo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *FolderDtoInteger) GetLogoOk() (*Logo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *FolderDtoInteger) SetLogo(v Logo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *FolderDtoInteger) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPinned

`func (o *FolderDtoInteger) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *FolderDtoInteger) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *FolderDtoInteger) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *FolderDtoInteger) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetRoomType

`func (o *FolderDtoInteger) GetRoomType() RoomType`

GetRoomType returns the RoomType field if non-nil, zero value otherwise.

### GetRoomTypeOk

`func (o *FolderDtoInteger) GetRoomTypeOk() (*RoomType, bool)`

GetRoomTypeOk returns a tuple with the RoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomType

`func (o *FolderDtoInteger) SetRoomType(v RoomType)`

SetRoomType sets RoomType field to given value.

### HasRoomType

`func (o *FolderDtoInteger) HasRoomType() bool`

HasRoomType returns a boolean if a field has been set.

### GetPrivate

`func (o *FolderDtoInteger) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *FolderDtoInteger) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *FolderDtoInteger) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *FolderDtoInteger) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetIndexing

`func (o *FolderDtoInteger) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *FolderDtoInteger) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *FolderDtoInteger) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *FolderDtoInteger) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### GetDenyDownload

`func (o *FolderDtoInteger) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *FolderDtoInteger) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *FolderDtoInteger) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *FolderDtoInteger) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetLifetime

`func (o *FolderDtoInteger) GetLifetime() RoomDataLifetimeDto`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *FolderDtoInteger) GetLifetimeOk() (*RoomDataLifetimeDto, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *FolderDtoInteger) SetLifetime(v RoomDataLifetimeDto)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *FolderDtoInteger) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetWatermark

`func (o *FolderDtoInteger) GetWatermark() WatermarkDto`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *FolderDtoInteger) GetWatermarkOk() (*WatermarkDto, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *FolderDtoInteger) SetWatermark(v WatermarkDto)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *FolderDtoInteger) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetType

`func (o *FolderDtoInteger) GetType() FolderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FolderDtoInteger) GetTypeOk() (*FolderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FolderDtoInteger) SetType(v FolderType)`

SetType sets Type field to given value.

### HasType

`func (o *FolderDtoInteger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInRoom

`func (o *FolderDtoInteger) GetInRoom() bool`

GetInRoom returns the InRoom field if non-nil, zero value otherwise.

### GetInRoomOk

`func (o *FolderDtoInteger) GetInRoomOk() (*bool, bool)`

GetInRoomOk returns a tuple with the InRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInRoom

`func (o *FolderDtoInteger) SetInRoom(v bool)`

SetInRoom sets InRoom field to given value.

### HasInRoom

`func (o *FolderDtoInteger) HasInRoom() bool`

HasInRoom returns a boolean if a field has been set.

### SetInRoomNil

`func (o *FolderDtoInteger) SetInRoomNil(b bool)`

 SetInRoomNil sets the value for InRoom to be an explicit nil

### UnsetInRoom
`func (o *FolderDtoInteger) UnsetInRoom()`

UnsetInRoom ensures that no value is present for InRoom, not even an explicit nil
### GetQuotaLimit

`func (o *FolderDtoInteger) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *FolderDtoInteger) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *FolderDtoInteger) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *FolderDtoInteger) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### SetQuotaLimitNil

`func (o *FolderDtoInteger) SetQuotaLimitNil(b bool)`

 SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil

### UnsetQuotaLimit
`func (o *FolderDtoInteger) UnsetQuotaLimit()`

UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
### GetIsCustomQuota

`func (o *FolderDtoInteger) GetIsCustomQuota() bool`

GetIsCustomQuota returns the IsCustomQuota field if non-nil, zero value otherwise.

### GetIsCustomQuotaOk

`func (o *FolderDtoInteger) GetIsCustomQuotaOk() (*bool, bool)`

GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomQuota

`func (o *FolderDtoInteger) SetIsCustomQuota(v bool)`

SetIsCustomQuota sets IsCustomQuota field to given value.

### HasIsCustomQuota

`func (o *FolderDtoInteger) HasIsCustomQuota() bool`

HasIsCustomQuota returns a boolean if a field has been set.

### SetIsCustomQuotaNil

`func (o *FolderDtoInteger) SetIsCustomQuotaNil(b bool)`

 SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil

### UnsetIsCustomQuota
`func (o *FolderDtoInteger) UnsetIsCustomQuota()`

UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
### GetUsedSpace

`func (o *FolderDtoInteger) GetUsedSpace() int64`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *FolderDtoInteger) GetUsedSpaceOk() (*int64, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *FolderDtoInteger) SetUsedSpace(v int64)`

SetUsedSpace sets UsedSpace field to given value.

### HasUsedSpace

`func (o *FolderDtoInteger) HasUsedSpace() bool`

HasUsedSpace returns a boolean if a field has been set.

### SetUsedSpaceNil

`func (o *FolderDtoInteger) SetUsedSpaceNil(b bool)`

 SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil

### UnsetUsedSpace
`func (o *FolderDtoInteger) UnsetUsedSpace()`

UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
### GetPasswordProtected

`func (o *FolderDtoInteger) GetPasswordProtected() bool`

GetPasswordProtected returns the PasswordProtected field if non-nil, zero value otherwise.

### GetPasswordProtectedOk

`func (o *FolderDtoInteger) GetPasswordProtectedOk() (*bool, bool)`

GetPasswordProtectedOk returns a tuple with the PasswordProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProtected

`func (o *FolderDtoInteger) SetPasswordProtected(v bool)`

SetPasswordProtected sets PasswordProtected field to given value.

### HasPasswordProtected

`func (o *FolderDtoInteger) HasPasswordProtected() bool`

HasPasswordProtected returns a boolean if a field has been set.

### SetPasswordProtectedNil

`func (o *FolderDtoInteger) SetPasswordProtectedNil(b bool)`

 SetPasswordProtectedNil sets the value for PasswordProtected to be an explicit nil

### UnsetPasswordProtected
`func (o *FolderDtoInteger) UnsetPasswordProtected()`

UnsetPasswordProtected ensures that no value is present for PasswordProtected, not even an explicit nil
### GetExpired

`func (o *FolderDtoInteger) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *FolderDtoInteger) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *FolderDtoInteger) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *FolderDtoInteger) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### SetExpiredNil

`func (o *FolderDtoInteger) SetExpiredNil(b bool)`

 SetExpiredNil sets the value for Expired to be an explicit nil

### UnsetExpired
`func (o *FolderDtoInteger) UnsetExpired()`

UnsetExpired ensures that no value is present for Expired, not even an explicit nil
### GetChatSettings

`func (o *FolderDtoInteger) GetChatSettings() ChatSettingsDto`

GetChatSettings returns the ChatSettings field if non-nil, zero value otherwise.

### GetChatSettingsOk

`func (o *FolderDtoInteger) GetChatSettingsOk() (*ChatSettingsDto, bool)`

GetChatSettingsOk returns a tuple with the ChatSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSettings

`func (o *FolderDtoInteger) SetChatSettings(v ChatSettingsDto)`

SetChatSettings sets ChatSettings field to given value.

### HasChatSettings

`func (o *FolderDtoInteger) HasChatSettings() bool`

HasChatSettings returns a boolean if a field has been set.

### GetRootRoomType

`func (o *FolderDtoInteger) GetRootRoomType() RoomType`

GetRootRoomType returns the RootRoomType field if non-nil, zero value otherwise.

### GetRootRoomTypeOk

`func (o *FolderDtoInteger) GetRootRoomTypeOk() (*RoomType, bool)`

GetRootRoomTypeOk returns a tuple with the RootRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRoomType

`func (o *FolderDtoInteger) SetRootRoomType(v RoomType)`

SetRootRoomType sets RootRoomType field to given value.

### HasRootRoomType

`func (o *FolderDtoInteger) HasRootRoomType() bool`

HasRootRoomType returns a boolean if a field has been set.

### GetSaveFormAsXLSX

`func (o *FolderDtoInteger) GetSaveFormAsXLSX() bool`

GetSaveFormAsXLSX returns the SaveFormAsXLSX field if non-nil, zero value otherwise.

### GetSaveFormAsXLSXOk

`func (o *FolderDtoInteger) GetSaveFormAsXLSXOk() (*bool, bool)`

GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveFormAsXLSX

`func (o *FolderDtoInteger) SetSaveFormAsXLSX(v bool)`

SetSaveFormAsXLSX sets SaveFormAsXLSX field to given value.

### HasSaveFormAsXLSX

`func (o *FolderDtoInteger) HasSaveFormAsXLSX() bool`

HasSaveFormAsXLSX returns a boolean if a field has been set.

### SetSaveFormAsXLSXNil

`func (o *FolderDtoInteger) SetSaveFormAsXLSXNil(b bool)`

 SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil

### UnsetSaveFormAsXLSX
`func (o *FolderDtoInteger) UnsetSaveFormAsXLSX()`

UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil
### GetSendFormToExternalDB

`func (o *FolderDtoInteger) GetSendFormToExternalDB() bool`

GetSendFormToExternalDB returns the SendFormToExternalDB field if non-nil, zero value otherwise.

### GetSendFormToExternalDBOk

`func (o *FolderDtoInteger) GetSendFormToExternalDBOk() (*bool, bool)`

GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFormToExternalDB

`func (o *FolderDtoInteger) SetSendFormToExternalDB(v bool)`

SetSendFormToExternalDB sets SendFormToExternalDB field to given value.

### HasSendFormToExternalDB

`func (o *FolderDtoInteger) HasSendFormToExternalDB() bool`

HasSendFormToExternalDB returns a boolean if a field has been set.

### SetSendFormToExternalDBNil

`func (o *FolderDtoInteger) SetSendFormToExternalDBNil(b bool)`

 SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil

### UnsetSendFormToExternalDB
`func (o *FolderDtoInteger) UnsetSendFormToExternalDB()`

UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
### GetOriginalFormId

`func (o *FolderDtoInteger) GetOriginalFormId() int32`

GetOriginalFormId returns the OriginalFormId field if non-nil, zero value otherwise.

### GetOriginalFormIdOk

`func (o *FolderDtoInteger) GetOriginalFormIdOk() (*int32, bool)`

GetOriginalFormIdOk returns a tuple with the OriginalFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFormId

`func (o *FolderDtoInteger) SetOriginalFormId(v int32)`

SetOriginalFormId sets OriginalFormId field to given value.

### HasOriginalFormId

`func (o *FolderDtoInteger) HasOriginalFormId() bool`

HasOriginalFormId returns a boolean if a field has been set.

### SetOriginalFormIdNil

`func (o *FolderDtoInteger) SetOriginalFormIdNil(b bool)`

 SetOriginalFormIdNil sets the value for OriginalFormId to be an explicit nil

### UnsetOriginalFormId
`func (o *FolderDtoInteger) UnsetOriginalFormId()`

UnsetOriginalFormId ensures that no value is present for OriginalFormId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


