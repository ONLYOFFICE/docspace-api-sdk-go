# FileDtoInteger

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
**FolderId** | Pointer to **int32** | The folder ID where the file is located. | [optional] 
**Version** | Pointer to **int32** | The file version. | [optional] 
**VersionGroup** | Pointer to **int32** | The version group of the file. | [optional] 
**ContentLength** | Pointer to **NullableString** | The content length of the file. | [optional] 
**PureContentLength** | Pointer to **NullableInt64** | The pure content length of the file. | [optional] 
**FileStatus** | Pointer to [**FileStatus**](FileStatus.md) |  | [optional] 
**EditingBy** | Pointer to **map[string]string** | The list of users editing the file. | [optional] 
**Mute** | Pointer to **bool** | Specifies if the file is muted or not. | [optional] 
**ViewUrl** | Pointer to **NullableString** | The URL link to view the file. | [optional] 
**WebUrl** | Pointer to **NullableString** | The Web URL link to the file. | [optional] 
**FileType** | Pointer to [**FileType**](FileType.md) |  | [optional] 
**FileExst** | Pointer to **NullableString** | The file extension. | [optional] 
**Comment** | Pointer to **NullableString** | The comment to the file. | [optional] 
**Encrypted** | Pointer to **NullableBool** | Specifies if the file is encrypted or not. | [optional] 
**ThumbnailUrl** | Pointer to **NullableString** | The thumbnail URL of the file. | [optional] 
**ThumbnailStatus** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**Locked** | Pointer to **NullableBool** | Specifies if the file is locked or not. | [optional] 
**LockedBy** | Pointer to **NullableString** | The user ID of the person who locked the file. | [optional] 
**HasDraft** | Pointer to **NullableBool** | Specifies if the file has a draft or not. | [optional] 
**FormFillingStatus** | Pointer to [**FormFillingStatus**](FormFillingStatus.md) |  | [optional] 
**IsForm** | Pointer to **NullableBool** | Specifies if the file is a form or not. | [optional] 
**CustomFilterEnabled** | Pointer to **NullableBool** | Specifies if the Custom Filter editing mode is enabled for a file or not. | [optional] 
**CustomFilterEnabledBy** | Pointer to **NullableString** | The name of the user who enabled a Custom Filter editing mode for a file. | [optional] 
**StartFilling** | Pointer to **NullableBool** | Specifies if the filling has started or not. | [optional] 
**IsFillingPreparing** | Pointer to **NullableBool** | Specifies if the form filling has started but the file is still being saved by the document editor. Filling and editing are not allowed. | [optional] 
**InProcessFolderId** | Pointer to **NullableInt32** | The InProcess folder ID of the file. | [optional] 
**InProcessFolderTitle** | Pointer to **NullableString** | The InProcess folder title of the file. | [optional] 
**ResultsFolderId** | Pointer to **NullableInt32** | The ID of the FormFillingFolderDone folder that corresponds to this original form. | [optional] 
**DraftLocation** | Pointer to [**DraftLocationInteger**](DraftLocationInteger.md) |  | [optional] 
**ViewAccessibility** | Pointer to [**NullableFileDtoIntegerAllOfViewAccessibility**](FileDtoIntegerAllOfViewAccessibility.md) |  | [optional] 
**LastOpened** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**Expired** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**VectorizationStatus** | Pointer to [**VectorizationStatus**](VectorizationStatus.md) |  | [optional] 
**ExternalDbTableName** | Pointer to **NullableString** | The name of the table in the external database that corresponds to this form. | [optional] 
**Dimensions** | Pointer to [**Size**](Size.md) |  | [optional] 

## Methods

### NewFileDtoInteger

`func NewFileDtoInteger() *FileDtoInteger`

NewFileDtoInteger instantiates a new FileDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDtoIntegerWithDefaults

`func NewFileDtoIntegerWithDefaults() *FileDtoInteger`

NewFileDtoIntegerWithDefaults instantiates a new FileDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FileDtoInteger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FileDtoInteger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FileDtoInteger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FileDtoInteger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FileDtoInteger) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FileDtoInteger) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAccess

`func (o *FileDtoInteger) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FileDtoInteger) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FileDtoInteger) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FileDtoInteger) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *FileDtoInteger) GetSharedBy() EmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *FileDtoInteger) GetSharedByOk() (*EmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *FileDtoInteger) SetSharedBy(v EmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *FileDtoInteger) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *FileDtoInteger) GetOwnedBy() EmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *FileDtoInteger) GetOwnedByOk() (*EmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *FileDtoInteger) SetOwnedBy(v EmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *FileDtoInteger) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *FileDtoInteger) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *FileDtoInteger) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *FileDtoInteger) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *FileDtoInteger) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *FileDtoInteger) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *FileDtoInteger) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *FileDtoInteger) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *FileDtoInteger) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *FileDtoInteger) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *FileDtoInteger) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *FileDtoInteger) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *FileDtoInteger) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *FileDtoInteger) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *FileDtoInteger) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *FileDtoInteger) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *FileDtoInteger) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *FileDtoInteger) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *FileDtoInteger) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *FileDtoInteger) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *FileDtoInteger) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### SetShortWebUrlNil

`func (o *FileDtoInteger) SetShortWebUrlNil(b bool)`

 SetShortWebUrlNil sets the value for ShortWebUrl to be an explicit nil

### UnsetShortWebUrl
`func (o *FileDtoInteger) UnsetShortWebUrl()`

UnsetShortWebUrl ensures that no value is present for ShortWebUrl, not even an explicit nil
### GetCreated

`func (o *FileDtoInteger) GetCreated() ApiDateTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FileDtoInteger) GetCreatedOk() (*ApiDateTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FileDtoInteger) SetCreated(v ApiDateTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FileDtoInteger) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FileDtoInteger) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FileDtoInteger) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FileDtoInteger) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FileDtoInteger) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *FileDtoInteger) GetUpdated() ApiDateTime`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FileDtoInteger) GetUpdatedOk() (*ApiDateTime, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FileDtoInteger) SetUpdated(v ApiDateTime)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FileDtoInteger) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAutoDelete

`func (o *FileDtoInteger) GetAutoDelete() ApiDateTime`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *FileDtoInteger) GetAutoDeleteOk() (*ApiDateTime, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *FileDtoInteger) SetAutoDelete(v ApiDateTime)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *FileDtoInteger) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetRootFolderType

`func (o *FileDtoInteger) GetRootFolderType() FolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *FileDtoInteger) GetRootFolderTypeOk() (*FolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *FileDtoInteger) SetRootFolderType(v FolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *FileDtoInteger) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *FileDtoInteger) GetParentRoomType() FolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *FileDtoInteger) GetParentRoomTypeOk() (*FolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *FileDtoInteger) SetParentRoomType(v FolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *FileDtoInteger) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FileDtoInteger) GetUpdatedBy() EmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FileDtoInteger) GetUpdatedByOk() (*EmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FileDtoInteger) SetUpdatedBy(v EmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FileDtoInteger) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *FileDtoInteger) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *FileDtoInteger) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *FileDtoInteger) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *FileDtoInteger) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### SetProviderItemNil

`func (o *FileDtoInteger) SetProviderItemNil(b bool)`

 SetProviderItemNil sets the value for ProviderItem to be an explicit nil

### UnsetProviderItem
`func (o *FileDtoInteger) UnsetProviderItem()`

UnsetProviderItem ensures that no value is present for ProviderItem, not even an explicit nil
### GetProviderKey

`func (o *FileDtoInteger) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *FileDtoInteger) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *FileDtoInteger) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *FileDtoInteger) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *FileDtoInteger) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *FileDtoInteger) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetProviderId

`func (o *FileDtoInteger) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *FileDtoInteger) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *FileDtoInteger) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *FileDtoInteger) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *FileDtoInteger) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *FileDtoInteger) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetOrder

`func (o *FileDtoInteger) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FileDtoInteger) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FileDtoInteger) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FileDtoInteger) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *FileDtoInteger) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *FileDtoInteger) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetIsFavorite

`func (o *FileDtoInteger) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *FileDtoInteger) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *FileDtoInteger) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *FileDtoInteger) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *FileDtoInteger) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *FileDtoInteger) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetFileEntryType

`func (o *FileDtoInteger) GetFileEntryType() FileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *FileDtoInteger) GetFileEntryTypeOk() (*FileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *FileDtoInteger) SetFileEntryType(v FileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *FileDtoInteger) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.

### GetId

`func (o *FileDtoInteger) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileDtoInteger) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileDtoInteger) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FileDtoInteger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootFolderId

`func (o *FileDtoInteger) GetRootFolderId() int32`

GetRootFolderId returns the RootFolderId field if non-nil, zero value otherwise.

### GetRootFolderIdOk

`func (o *FileDtoInteger) GetRootFolderIdOk() (*int32, bool)`

GetRootFolderIdOk returns a tuple with the RootFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderId

`func (o *FileDtoInteger) SetRootFolderId(v int32)`

SetRootFolderId sets RootFolderId field to given value.

### HasRootFolderId

`func (o *FileDtoInteger) HasRootFolderId() bool`

HasRootFolderId returns a boolean if a field has been set.

### GetOriginId

`func (o *FileDtoInteger) GetOriginId() int32`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *FileDtoInteger) GetOriginIdOk() (*int32, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *FileDtoInteger) SetOriginId(v int32)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *FileDtoInteger) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginRoomId

`func (o *FileDtoInteger) GetOriginRoomId() int32`

GetOriginRoomId returns the OriginRoomId field if non-nil, zero value otherwise.

### GetOriginRoomIdOk

`func (o *FileDtoInteger) GetOriginRoomIdOk() (*int32, bool)`

GetOriginRoomIdOk returns a tuple with the OriginRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomId

`func (o *FileDtoInteger) SetOriginRoomId(v int32)`

SetOriginRoomId sets OriginRoomId field to given value.

### HasOriginRoomId

`func (o *FileDtoInteger) HasOriginRoomId() bool`

HasOriginRoomId returns a boolean if a field has been set.

### GetOriginTitle

`func (o *FileDtoInteger) GetOriginTitle() string`

GetOriginTitle returns the OriginTitle field if non-nil, zero value otherwise.

### GetOriginTitleOk

`func (o *FileDtoInteger) GetOriginTitleOk() (*string, bool)`

GetOriginTitleOk returns a tuple with the OriginTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTitle

`func (o *FileDtoInteger) SetOriginTitle(v string)`

SetOriginTitle sets OriginTitle field to given value.

### HasOriginTitle

`func (o *FileDtoInteger) HasOriginTitle() bool`

HasOriginTitle returns a boolean if a field has been set.

### SetOriginTitleNil

`func (o *FileDtoInteger) SetOriginTitleNil(b bool)`

 SetOriginTitleNil sets the value for OriginTitle to be an explicit nil

### UnsetOriginTitle
`func (o *FileDtoInteger) UnsetOriginTitle()`

UnsetOriginTitle ensures that no value is present for OriginTitle, not even an explicit nil
### GetOriginRoomTitle

`func (o *FileDtoInteger) GetOriginRoomTitle() string`

GetOriginRoomTitle returns the OriginRoomTitle field if non-nil, zero value otherwise.

### GetOriginRoomTitleOk

`func (o *FileDtoInteger) GetOriginRoomTitleOk() (*string, bool)`

GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomTitle

`func (o *FileDtoInteger) SetOriginRoomTitle(v string)`

SetOriginRoomTitle sets OriginRoomTitle field to given value.

### HasOriginRoomTitle

`func (o *FileDtoInteger) HasOriginRoomTitle() bool`

HasOriginRoomTitle returns a boolean if a field has been set.

### SetOriginRoomTitleNil

`func (o *FileDtoInteger) SetOriginRoomTitleNil(b bool)`

 SetOriginRoomTitleNil sets the value for OriginRoomTitle to be an explicit nil

### UnsetOriginRoomTitle
`func (o *FileDtoInteger) UnsetOriginRoomTitle()`

UnsetOriginRoomTitle ensures that no value is present for OriginRoomTitle, not even an explicit nil
### GetCanShare

`func (o *FileDtoInteger) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *FileDtoInteger) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *FileDtoInteger) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *FileDtoInteger) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### GetShareSettings

`func (o *FileDtoInteger) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings`

GetShareSettings returns the ShareSettings field if non-nil, zero value otherwise.

### GetShareSettingsOk

`func (o *FileDtoInteger) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool)`

GetShareSettingsOk returns a tuple with the ShareSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareSettings

`func (o *FileDtoInteger) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings)`

SetShareSettings sets ShareSettings field to given value.

### HasShareSettings

`func (o *FileDtoInteger) HasShareSettings() bool`

HasShareSettings returns a boolean if a field has been set.

### SetShareSettingsNil

`func (o *FileDtoInteger) SetShareSettingsNil(b bool)`

 SetShareSettingsNil sets the value for ShareSettings to be an explicit nil

### UnsetShareSettings
`func (o *FileDtoInteger) UnsetShareSettings()`

UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
### GetSecurity

`func (o *FileDtoInteger) GetSecurity() FileEntryDtoIntegerAllOfSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *FileDtoInteger) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *FileDtoInteger) SetSecurity(v FileEntryDtoIntegerAllOfSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *FileDtoInteger) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *FileDtoInteger) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *FileDtoInteger) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetAvailableShareRights

`func (o *FileDtoInteger) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights`

GetAvailableShareRights returns the AvailableShareRights field if non-nil, zero value otherwise.

### GetAvailableShareRightsOk

`func (o *FileDtoInteger) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool)`

GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShareRights

`func (o *FileDtoInteger) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights)`

SetAvailableShareRights sets AvailableShareRights field to given value.

### HasAvailableShareRights

`func (o *FileDtoInteger) HasAvailableShareRights() bool`

HasAvailableShareRights returns a boolean if a field has been set.

### SetAvailableShareRightsNil

`func (o *FileDtoInteger) SetAvailableShareRightsNil(b bool)`

 SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil

### UnsetAvailableShareRights
`func (o *FileDtoInteger) UnsetAvailableShareRights()`

UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
### GetRequestToken

`func (o *FileDtoInteger) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *FileDtoInteger) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *FileDtoInteger) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *FileDtoInteger) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### SetRequestTokenNil

`func (o *FileDtoInteger) SetRequestTokenNil(b bool)`

 SetRequestTokenNil sets the value for RequestToken to be an explicit nil

### UnsetRequestToken
`func (o *FileDtoInteger) UnsetRequestToken()`

UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
### GetExternal

`func (o *FileDtoInteger) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *FileDtoInteger) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *FileDtoInteger) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *FileDtoInteger) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### SetExternalNil

`func (o *FileDtoInteger) SetExternalNil(b bool)`

 SetExternalNil sets the value for External to be an explicit nil

### UnsetExternal
`func (o *FileDtoInteger) UnsetExternal()`

UnsetExternal ensures that no value is present for External, not even an explicit nil
### GetExpirationDate

`func (o *FileDtoInteger) GetExpirationDate() ApiDateTime`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FileDtoInteger) GetExpirationDateOk() (*ApiDateTime, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FileDtoInteger) SetExpirationDate(v ApiDateTime)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FileDtoInteger) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIsLinkExpired

`func (o *FileDtoInteger) GetIsLinkExpired() bool`

GetIsLinkExpired returns the IsLinkExpired field if non-nil, zero value otherwise.

### GetIsLinkExpiredOk

`func (o *FileDtoInteger) GetIsLinkExpiredOk() (*bool, bool)`

GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkExpired

`func (o *FileDtoInteger) SetIsLinkExpired(v bool)`

SetIsLinkExpired sets IsLinkExpired field to given value.

### HasIsLinkExpired

`func (o *FileDtoInteger) HasIsLinkExpired() bool`

HasIsLinkExpired returns a boolean if a field has been set.

### SetIsLinkExpiredNil

`func (o *FileDtoInteger) SetIsLinkExpiredNil(b bool)`

 SetIsLinkExpiredNil sets the value for IsLinkExpired to be an explicit nil

### UnsetIsLinkExpired
`func (o *FileDtoInteger) UnsetIsLinkExpired()`

UnsetIsLinkExpired ensures that no value is present for IsLinkExpired, not even an explicit nil
### GetFolderId

`func (o *FileDtoInteger) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *FileDtoInteger) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *FileDtoInteger) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *FileDtoInteger) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetVersion

`func (o *FileDtoInteger) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FileDtoInteger) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FileDtoInteger) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FileDtoInteger) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionGroup

`func (o *FileDtoInteger) GetVersionGroup() int32`

GetVersionGroup returns the VersionGroup field if non-nil, zero value otherwise.

### GetVersionGroupOk

`func (o *FileDtoInteger) GetVersionGroupOk() (*int32, bool)`

GetVersionGroupOk returns a tuple with the VersionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroup

`func (o *FileDtoInteger) SetVersionGroup(v int32)`

SetVersionGroup sets VersionGroup field to given value.

### HasVersionGroup

`func (o *FileDtoInteger) HasVersionGroup() bool`

HasVersionGroup returns a boolean if a field has been set.

### GetContentLength

`func (o *FileDtoInteger) GetContentLength() string`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *FileDtoInteger) GetContentLengthOk() (*string, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *FileDtoInteger) SetContentLength(v string)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *FileDtoInteger) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *FileDtoInteger) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *FileDtoInteger) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetPureContentLength

`func (o *FileDtoInteger) GetPureContentLength() int64`

GetPureContentLength returns the PureContentLength field if non-nil, zero value otherwise.

### GetPureContentLengthOk

`func (o *FileDtoInteger) GetPureContentLengthOk() (*int64, bool)`

GetPureContentLengthOk returns a tuple with the PureContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureContentLength

`func (o *FileDtoInteger) SetPureContentLength(v int64)`

SetPureContentLength sets PureContentLength field to given value.

### HasPureContentLength

`func (o *FileDtoInteger) HasPureContentLength() bool`

HasPureContentLength returns a boolean if a field has been set.

### SetPureContentLengthNil

`func (o *FileDtoInteger) SetPureContentLengthNil(b bool)`

 SetPureContentLengthNil sets the value for PureContentLength to be an explicit nil

### UnsetPureContentLength
`func (o *FileDtoInteger) UnsetPureContentLength()`

UnsetPureContentLength ensures that no value is present for PureContentLength, not even an explicit nil
### GetFileStatus

`func (o *FileDtoInteger) GetFileStatus() FileStatus`

GetFileStatus returns the FileStatus field if non-nil, zero value otherwise.

### GetFileStatusOk

`func (o *FileDtoInteger) GetFileStatusOk() (*FileStatus, bool)`

GetFileStatusOk returns a tuple with the FileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStatus

`func (o *FileDtoInteger) SetFileStatus(v FileStatus)`

SetFileStatus sets FileStatus field to given value.

### HasFileStatus

`func (o *FileDtoInteger) HasFileStatus() bool`

HasFileStatus returns a boolean if a field has been set.

### GetEditingBy

`func (o *FileDtoInteger) GetEditingBy() map[string]string`

GetEditingBy returns the EditingBy field if non-nil, zero value otherwise.

### GetEditingByOk

`func (o *FileDtoInteger) GetEditingByOk() (*map[string]string, bool)`

GetEditingByOk returns a tuple with the EditingBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditingBy

`func (o *FileDtoInteger) SetEditingBy(v map[string]string)`

SetEditingBy sets EditingBy field to given value.

### HasEditingBy

`func (o *FileDtoInteger) HasEditingBy() bool`

HasEditingBy returns a boolean if a field has been set.

### SetEditingByNil

`func (o *FileDtoInteger) SetEditingByNil(b bool)`

 SetEditingByNil sets the value for EditingBy to be an explicit nil

### UnsetEditingBy
`func (o *FileDtoInteger) UnsetEditingBy()`

UnsetEditingBy ensures that no value is present for EditingBy, not even an explicit nil
### GetMute

`func (o *FileDtoInteger) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *FileDtoInteger) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *FileDtoInteger) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *FileDtoInteger) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetViewUrl

`func (o *FileDtoInteger) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *FileDtoInteger) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *FileDtoInteger) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *FileDtoInteger) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.

### SetViewUrlNil

`func (o *FileDtoInteger) SetViewUrlNil(b bool)`

 SetViewUrlNil sets the value for ViewUrl to be an explicit nil

### UnsetViewUrl
`func (o *FileDtoInteger) UnsetViewUrl()`

UnsetViewUrl ensures that no value is present for ViewUrl, not even an explicit nil
### GetWebUrl

`func (o *FileDtoInteger) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *FileDtoInteger) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *FileDtoInteger) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *FileDtoInteger) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *FileDtoInteger) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *FileDtoInteger) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetFileType

`func (o *FileDtoInteger) GetFileType() FileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FileDtoInteger) GetFileTypeOk() (*FileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FileDtoInteger) SetFileType(v FileType)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *FileDtoInteger) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFileExst

`func (o *FileDtoInteger) GetFileExst() string`

GetFileExst returns the FileExst field if non-nil, zero value otherwise.

### GetFileExstOk

`func (o *FileDtoInteger) GetFileExstOk() (*string, bool)`

GetFileExstOk returns a tuple with the FileExst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExst

`func (o *FileDtoInteger) SetFileExst(v string)`

SetFileExst sets FileExst field to given value.

### HasFileExst

`func (o *FileDtoInteger) HasFileExst() bool`

HasFileExst returns a boolean if a field has been set.

### SetFileExstNil

`func (o *FileDtoInteger) SetFileExstNil(b bool)`

 SetFileExstNil sets the value for FileExst to be an explicit nil

### UnsetFileExst
`func (o *FileDtoInteger) UnsetFileExst()`

UnsetFileExst ensures that no value is present for FileExst, not even an explicit nil
### GetComment

`func (o *FileDtoInteger) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FileDtoInteger) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FileDtoInteger) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FileDtoInteger) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *FileDtoInteger) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FileDtoInteger) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetEncrypted

`func (o *FileDtoInteger) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *FileDtoInteger) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *FileDtoInteger) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *FileDtoInteger) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### SetEncryptedNil

`func (o *FileDtoInteger) SetEncryptedNil(b bool)`

 SetEncryptedNil sets the value for Encrypted to be an explicit nil

### UnsetEncrypted
`func (o *FileDtoInteger) UnsetEncrypted()`

UnsetEncrypted ensures that no value is present for Encrypted, not even an explicit nil
### GetThumbnailUrl

`func (o *FileDtoInteger) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *FileDtoInteger) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *FileDtoInteger) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *FileDtoInteger) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *FileDtoInteger) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *FileDtoInteger) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetThumbnailStatus

`func (o *FileDtoInteger) GetThumbnailStatus() Thumbnail`

GetThumbnailStatus returns the ThumbnailStatus field if non-nil, zero value otherwise.

### GetThumbnailStatusOk

`func (o *FileDtoInteger) GetThumbnailStatusOk() (*Thumbnail, bool)`

GetThumbnailStatusOk returns a tuple with the ThumbnailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailStatus

`func (o *FileDtoInteger) SetThumbnailStatus(v Thumbnail)`

SetThumbnailStatus sets ThumbnailStatus field to given value.

### HasThumbnailStatus

`func (o *FileDtoInteger) HasThumbnailStatus() bool`

HasThumbnailStatus returns a boolean if a field has been set.

### GetLocked

`func (o *FileDtoInteger) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *FileDtoInteger) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *FileDtoInteger) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *FileDtoInteger) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *FileDtoInteger) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *FileDtoInteger) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetLockedBy

`func (o *FileDtoInteger) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *FileDtoInteger) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *FileDtoInteger) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *FileDtoInteger) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### SetLockedByNil

`func (o *FileDtoInteger) SetLockedByNil(b bool)`

 SetLockedByNil sets the value for LockedBy to be an explicit nil

### UnsetLockedBy
`func (o *FileDtoInteger) UnsetLockedBy()`

UnsetLockedBy ensures that no value is present for LockedBy, not even an explicit nil
### GetHasDraft

`func (o *FileDtoInteger) GetHasDraft() bool`

GetHasDraft returns the HasDraft field if non-nil, zero value otherwise.

### GetHasDraftOk

`func (o *FileDtoInteger) GetHasDraftOk() (*bool, bool)`

GetHasDraftOk returns a tuple with the HasDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDraft

`func (o *FileDtoInteger) SetHasDraft(v bool)`

SetHasDraft sets HasDraft field to given value.

### HasHasDraft

`func (o *FileDtoInteger) HasHasDraft() bool`

HasHasDraft returns a boolean if a field has been set.

### SetHasDraftNil

`func (o *FileDtoInteger) SetHasDraftNil(b bool)`

 SetHasDraftNil sets the value for HasDraft to be an explicit nil

### UnsetHasDraft
`func (o *FileDtoInteger) UnsetHasDraft()`

UnsetHasDraft ensures that no value is present for HasDraft, not even an explicit nil
### GetFormFillingStatus

`func (o *FileDtoInteger) GetFormFillingStatus() FormFillingStatus`

GetFormFillingStatus returns the FormFillingStatus field if non-nil, zero value otherwise.

### GetFormFillingStatusOk

`func (o *FileDtoInteger) GetFormFillingStatusOk() (*FormFillingStatus, bool)`

GetFormFillingStatusOk returns a tuple with the FormFillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFillingStatus

`func (o *FileDtoInteger) SetFormFillingStatus(v FormFillingStatus)`

SetFormFillingStatus sets FormFillingStatus field to given value.

### HasFormFillingStatus

`func (o *FileDtoInteger) HasFormFillingStatus() bool`

HasFormFillingStatus returns a boolean if a field has been set.

### GetIsForm

`func (o *FileDtoInteger) GetIsForm() bool`

GetIsForm returns the IsForm field if non-nil, zero value otherwise.

### GetIsFormOk

`func (o *FileDtoInteger) GetIsFormOk() (*bool, bool)`

GetIsFormOk returns a tuple with the IsForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForm

`func (o *FileDtoInteger) SetIsForm(v bool)`

SetIsForm sets IsForm field to given value.

### HasIsForm

`func (o *FileDtoInteger) HasIsForm() bool`

HasIsForm returns a boolean if a field has been set.

### SetIsFormNil

`func (o *FileDtoInteger) SetIsFormNil(b bool)`

 SetIsFormNil sets the value for IsForm to be an explicit nil

### UnsetIsForm
`func (o *FileDtoInteger) UnsetIsForm()`

UnsetIsForm ensures that no value is present for IsForm, not even an explicit nil
### GetCustomFilterEnabled

`func (o *FileDtoInteger) GetCustomFilterEnabled() bool`

GetCustomFilterEnabled returns the CustomFilterEnabled field if non-nil, zero value otherwise.

### GetCustomFilterEnabledOk

`func (o *FileDtoInteger) GetCustomFilterEnabledOk() (*bool, bool)`

GetCustomFilterEnabledOk returns a tuple with the CustomFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterEnabled

`func (o *FileDtoInteger) SetCustomFilterEnabled(v bool)`

SetCustomFilterEnabled sets CustomFilterEnabled field to given value.

### HasCustomFilterEnabled

`func (o *FileDtoInteger) HasCustomFilterEnabled() bool`

HasCustomFilterEnabled returns a boolean if a field has been set.

### SetCustomFilterEnabledNil

`func (o *FileDtoInteger) SetCustomFilterEnabledNil(b bool)`

 SetCustomFilterEnabledNil sets the value for CustomFilterEnabled to be an explicit nil

### UnsetCustomFilterEnabled
`func (o *FileDtoInteger) UnsetCustomFilterEnabled()`

UnsetCustomFilterEnabled ensures that no value is present for CustomFilterEnabled, not even an explicit nil
### GetCustomFilterEnabledBy

`func (o *FileDtoInteger) GetCustomFilterEnabledBy() string`

GetCustomFilterEnabledBy returns the CustomFilterEnabledBy field if non-nil, zero value otherwise.

### GetCustomFilterEnabledByOk

`func (o *FileDtoInteger) GetCustomFilterEnabledByOk() (*string, bool)`

GetCustomFilterEnabledByOk returns a tuple with the CustomFilterEnabledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterEnabledBy

`func (o *FileDtoInteger) SetCustomFilterEnabledBy(v string)`

SetCustomFilterEnabledBy sets CustomFilterEnabledBy field to given value.

### HasCustomFilterEnabledBy

`func (o *FileDtoInteger) HasCustomFilterEnabledBy() bool`

HasCustomFilterEnabledBy returns a boolean if a field has been set.

### SetCustomFilterEnabledByNil

`func (o *FileDtoInteger) SetCustomFilterEnabledByNil(b bool)`

 SetCustomFilterEnabledByNil sets the value for CustomFilterEnabledBy to be an explicit nil

### UnsetCustomFilterEnabledBy
`func (o *FileDtoInteger) UnsetCustomFilterEnabledBy()`

UnsetCustomFilterEnabledBy ensures that no value is present for CustomFilterEnabledBy, not even an explicit nil
### GetStartFilling

`func (o *FileDtoInteger) GetStartFilling() bool`

GetStartFilling returns the StartFilling field if non-nil, zero value otherwise.

### GetStartFillingOk

`func (o *FileDtoInteger) GetStartFillingOk() (*bool, bool)`

GetStartFillingOk returns a tuple with the StartFilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFilling

`func (o *FileDtoInteger) SetStartFilling(v bool)`

SetStartFilling sets StartFilling field to given value.

### HasStartFilling

`func (o *FileDtoInteger) HasStartFilling() bool`

HasStartFilling returns a boolean if a field has been set.

### SetStartFillingNil

`func (o *FileDtoInteger) SetStartFillingNil(b bool)`

 SetStartFillingNil sets the value for StartFilling to be an explicit nil

### UnsetStartFilling
`func (o *FileDtoInteger) UnsetStartFilling()`

UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
### GetIsFillingPreparing

`func (o *FileDtoInteger) GetIsFillingPreparing() bool`

GetIsFillingPreparing returns the IsFillingPreparing field if non-nil, zero value otherwise.

### GetIsFillingPreparingOk

`func (o *FileDtoInteger) GetIsFillingPreparingOk() (*bool, bool)`

GetIsFillingPreparingOk returns a tuple with the IsFillingPreparing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFillingPreparing

`func (o *FileDtoInteger) SetIsFillingPreparing(v bool)`

SetIsFillingPreparing sets IsFillingPreparing field to given value.

### HasIsFillingPreparing

`func (o *FileDtoInteger) HasIsFillingPreparing() bool`

HasIsFillingPreparing returns a boolean if a field has been set.

### SetIsFillingPreparingNil

`func (o *FileDtoInteger) SetIsFillingPreparingNil(b bool)`

 SetIsFillingPreparingNil sets the value for IsFillingPreparing to be an explicit nil

### UnsetIsFillingPreparing
`func (o *FileDtoInteger) UnsetIsFillingPreparing()`

UnsetIsFillingPreparing ensures that no value is present for IsFillingPreparing, not even an explicit nil
### GetInProcessFolderId

`func (o *FileDtoInteger) GetInProcessFolderId() int32`

GetInProcessFolderId returns the InProcessFolderId field if non-nil, zero value otherwise.

### GetInProcessFolderIdOk

`func (o *FileDtoInteger) GetInProcessFolderIdOk() (*int32, bool)`

GetInProcessFolderIdOk returns a tuple with the InProcessFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProcessFolderId

`func (o *FileDtoInteger) SetInProcessFolderId(v int32)`

SetInProcessFolderId sets InProcessFolderId field to given value.

### HasInProcessFolderId

`func (o *FileDtoInteger) HasInProcessFolderId() bool`

HasInProcessFolderId returns a boolean if a field has been set.

### SetInProcessFolderIdNil

`func (o *FileDtoInteger) SetInProcessFolderIdNil(b bool)`

 SetInProcessFolderIdNil sets the value for InProcessFolderId to be an explicit nil

### UnsetInProcessFolderId
`func (o *FileDtoInteger) UnsetInProcessFolderId()`

UnsetInProcessFolderId ensures that no value is present for InProcessFolderId, not even an explicit nil
### GetInProcessFolderTitle

`func (o *FileDtoInteger) GetInProcessFolderTitle() string`

GetInProcessFolderTitle returns the InProcessFolderTitle field if non-nil, zero value otherwise.

### GetInProcessFolderTitleOk

`func (o *FileDtoInteger) GetInProcessFolderTitleOk() (*string, bool)`

GetInProcessFolderTitleOk returns a tuple with the InProcessFolderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProcessFolderTitle

`func (o *FileDtoInteger) SetInProcessFolderTitle(v string)`

SetInProcessFolderTitle sets InProcessFolderTitle field to given value.

### HasInProcessFolderTitle

`func (o *FileDtoInteger) HasInProcessFolderTitle() bool`

HasInProcessFolderTitle returns a boolean if a field has been set.

### SetInProcessFolderTitleNil

`func (o *FileDtoInteger) SetInProcessFolderTitleNil(b bool)`

 SetInProcessFolderTitleNil sets the value for InProcessFolderTitle to be an explicit nil

### UnsetInProcessFolderTitle
`func (o *FileDtoInteger) UnsetInProcessFolderTitle()`

UnsetInProcessFolderTitle ensures that no value is present for InProcessFolderTitle, not even an explicit nil
### GetResultsFolderId

`func (o *FileDtoInteger) GetResultsFolderId() int32`

GetResultsFolderId returns the ResultsFolderId field if non-nil, zero value otherwise.

### GetResultsFolderIdOk

`func (o *FileDtoInteger) GetResultsFolderIdOk() (*int32, bool)`

GetResultsFolderIdOk returns a tuple with the ResultsFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsFolderId

`func (o *FileDtoInteger) SetResultsFolderId(v int32)`

SetResultsFolderId sets ResultsFolderId field to given value.

### HasResultsFolderId

`func (o *FileDtoInteger) HasResultsFolderId() bool`

HasResultsFolderId returns a boolean if a field has been set.

### SetResultsFolderIdNil

`func (o *FileDtoInteger) SetResultsFolderIdNil(b bool)`

 SetResultsFolderIdNil sets the value for ResultsFolderId to be an explicit nil

### UnsetResultsFolderId
`func (o *FileDtoInteger) UnsetResultsFolderId()`

UnsetResultsFolderId ensures that no value is present for ResultsFolderId, not even an explicit nil
### GetDraftLocation

`func (o *FileDtoInteger) GetDraftLocation() DraftLocationInteger`

GetDraftLocation returns the DraftLocation field if non-nil, zero value otherwise.

### GetDraftLocationOk

`func (o *FileDtoInteger) GetDraftLocationOk() (*DraftLocationInteger, bool)`

GetDraftLocationOk returns a tuple with the DraftLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftLocation

`func (o *FileDtoInteger) SetDraftLocation(v DraftLocationInteger)`

SetDraftLocation sets DraftLocation field to given value.

### HasDraftLocation

`func (o *FileDtoInteger) HasDraftLocation() bool`

HasDraftLocation returns a boolean if a field has been set.

### GetViewAccessibility

`func (o *FileDtoInteger) GetViewAccessibility() FileDtoIntegerAllOfViewAccessibility`

GetViewAccessibility returns the ViewAccessibility field if non-nil, zero value otherwise.

### GetViewAccessibilityOk

`func (o *FileDtoInteger) GetViewAccessibilityOk() (*FileDtoIntegerAllOfViewAccessibility, bool)`

GetViewAccessibilityOk returns a tuple with the ViewAccessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAccessibility

`func (o *FileDtoInteger) SetViewAccessibility(v FileDtoIntegerAllOfViewAccessibility)`

SetViewAccessibility sets ViewAccessibility field to given value.

### HasViewAccessibility

`func (o *FileDtoInteger) HasViewAccessibility() bool`

HasViewAccessibility returns a boolean if a field has been set.

### SetViewAccessibilityNil

`func (o *FileDtoInteger) SetViewAccessibilityNil(b bool)`

 SetViewAccessibilityNil sets the value for ViewAccessibility to be an explicit nil

### UnsetViewAccessibility
`func (o *FileDtoInteger) UnsetViewAccessibility()`

UnsetViewAccessibility ensures that no value is present for ViewAccessibility, not even an explicit nil
### GetLastOpened

`func (o *FileDtoInteger) GetLastOpened() ApiDateTime`

GetLastOpened returns the LastOpened field if non-nil, zero value otherwise.

### GetLastOpenedOk

`func (o *FileDtoInteger) GetLastOpenedOk() (*ApiDateTime, bool)`

GetLastOpenedOk returns a tuple with the LastOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpened

`func (o *FileDtoInteger) SetLastOpened(v ApiDateTime)`

SetLastOpened sets LastOpened field to given value.

### HasLastOpened

`func (o *FileDtoInteger) HasLastOpened() bool`

HasLastOpened returns a boolean if a field has been set.

### GetExpired

`func (o *FileDtoInteger) GetExpired() ApiDateTime`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *FileDtoInteger) GetExpiredOk() (*ApiDateTime, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *FileDtoInteger) SetExpired(v ApiDateTime)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *FileDtoInteger) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetVectorizationStatus

`func (o *FileDtoInteger) GetVectorizationStatus() VectorizationStatus`

GetVectorizationStatus returns the VectorizationStatus field if non-nil, zero value otherwise.

### GetVectorizationStatusOk

`func (o *FileDtoInteger) GetVectorizationStatusOk() (*VectorizationStatus, bool)`

GetVectorizationStatusOk returns a tuple with the VectorizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorizationStatus

`func (o *FileDtoInteger) SetVectorizationStatus(v VectorizationStatus)`

SetVectorizationStatus sets VectorizationStatus field to given value.

### HasVectorizationStatus

`func (o *FileDtoInteger) HasVectorizationStatus() bool`

HasVectorizationStatus returns a boolean if a field has been set.

### GetExternalDbTableName

`func (o *FileDtoInteger) GetExternalDbTableName() string`

GetExternalDbTableName returns the ExternalDbTableName field if non-nil, zero value otherwise.

### GetExternalDbTableNameOk

`func (o *FileDtoInteger) GetExternalDbTableNameOk() (*string, bool)`

GetExternalDbTableNameOk returns a tuple with the ExternalDbTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDbTableName

`func (o *FileDtoInteger) SetExternalDbTableName(v string)`

SetExternalDbTableName sets ExternalDbTableName field to given value.

### HasExternalDbTableName

`func (o *FileDtoInteger) HasExternalDbTableName() bool`

HasExternalDbTableName returns a boolean if a field has been set.

### SetExternalDbTableNameNil

`func (o *FileDtoInteger) SetExternalDbTableNameNil(b bool)`

 SetExternalDbTableNameNil sets the value for ExternalDbTableName to be an explicit nil

### UnsetExternalDbTableName
`func (o *FileDtoInteger) UnsetExternalDbTableName()`

UnsetExternalDbTableName ensures that no value is present for ExternalDbTableName, not even an explicit nil
### GetDimensions

`func (o *FileDtoInteger) GetDimensions() Size`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *FileDtoInteger) GetDimensionsOk() (*Size, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *FileDtoInteger) SetDimensions(v Size)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *FileDtoInteger) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


