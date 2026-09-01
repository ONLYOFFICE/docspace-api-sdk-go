# FolderDtoString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The file entry title. | [optional] 
**Access** | Pointer to [**FileShare**](FileShare.md) | The access rights to the file entry. | [optional] 
**SharedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | Provides information about the employee who shared the file or folder. | [optional] 
**OwnedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The information about the employee who owns the file entry. | [optional] 
**Shared** | Pointer to **bool** | Specifies if the file entry is shared via link or not. | [optional] 
**SharedForUser** | Pointer to **bool** | Specifies if the file entry is shared for user or not. | [optional] 
**SharedExternal** | Pointer to **bool** | Specifies if the file entry is shared via a public (non-internal) external link. | [optional] 
**ParentShared** | Pointer to **bool** | Indicates whether the parent entity is shared. | [optional] 
**ShortWebUrl** | Pointer to **string** | The short Web URL. | [optional] 
**Created** | Pointer to **time.Time** | The creation date and time of the file entry. | [optional] 
**CreatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The file entry author. | [optional] 
**Updated** | Pointer to **time.Time** | The last date and time when the file entry was updated. | [optional] 
**AutoDelete** | Pointer to **time.Time** | The date and time when the file entry will be automatically deleted. | [optional] 
**RootFolderType** | Pointer to [**FolderType**](FolderType.md) | The root folder type of the file entry. | [optional] 
**ParentRoomType** | Pointer to [**FolderType**](FolderType.md) | The parent room type of the file entry. | [optional] 
**UpdatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The user who updated the file entry. | [optional] 
**ProviderItem** | Pointer to **bool** | Specifies if the file entry provider is specified or not. | [optional] 
**ProviderKey** | Pointer to **string** | The provider key of the file entry. | [optional] 
**ProviderId** | Pointer to **int32** | The provider ID of the file entry. | [optional] 
**Order** | Pointer to **string** | The order of the file entry. | [optional] 
**IsFavorite** | Pointer to **bool** | Specifies if the file is a favorite or not. | [optional] 
**FileEntryType** | Pointer to [**FileEntryType**](FileEntryType.md) | The file entry type. | [optional] 
**Id** | Pointer to **string** | The file entry ID. | [optional] 
**RootFolderId** | Pointer to **string** | The root folder ID of the file entry. | [optional] 
**OriginId** | Pointer to **string** | The origin ID of the file entry. | [optional] 
**OriginRoomId** | Pointer to **string** | The origin room ID of the file entry. | [optional] 
**OriginTitle** | Pointer to **string** | The origin title of the file entry. | [optional] 
**OriginRoomTitle** | Pointer to **string** | The origin room title of the file entry. | [optional] 
**CanShare** | Pointer to **bool** | Specifies if the file entry can be shared or not. | [optional] 
**ShareSettings** | Pointer to [**NullableFileEntryDtoIntegerAllOfShareSettings**](FileEntryDtoIntegerAllOfShareSettings.md) |  | [optional] 
**Security** | Pointer to [**NullableFileEntryDtoIntegerAllOfSecurity**](FileEntryDtoIntegerAllOfSecurity.md) |  | [optional] 
**AvailableShareRights** | Pointer to [**NullableFileEntryDtoIntegerAllOfAvailableShareRights**](FileEntryDtoIntegerAllOfAvailableShareRights.md) |  | [optional] 
**RequestToken** | Pointer to **string** | The request token of the file entry. | [optional] 
**External** | Pointer to **bool** | Specifies if the folder can be accessed via an external link or not. | [optional] 
**ExpirationDate** | Pointer to **time.Time** | Represents the expiration date of the file entry. | [optional] 
**IsLinkExpired** | Pointer to **bool** | Indicates whether the shareable link associated with the file or folder has expired. | [optional] 
**ParentId** | Pointer to **NullableString** | The parent folder ID of the folder. | [optional] 
**FilesCount** | Pointer to **int32** | The number of files that the folder contains. | [optional] 
**FoldersCount** | Pointer to **int32** | The number of folders that the folder contains. | [optional] 
**IsShareable** | Pointer to **NullableBool** | Specifies if the folder can be shared or not. | [optional] 
**New** | Pointer to **int32** | The new element index in the folder. | [optional] 
**Mute** | Pointer to **bool** | Specifies if the folder notifications are enabled or not. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of the folder. | [optional] 
**Logo** | Pointer to [**Logo**](Logo.md) | The folder logo. | [optional] 
**Pinned** | Pointer to **bool** | Specifies if the folder is pinned or not. | [optional] 
**RoomType** | Pointer to [**RoomType**](RoomType.md) | The room type of the folder. | [optional] 
**Private** | Pointer to **bool** | Specifies if the folder is private or not. | [optional] 
**Indexing** | Pointer to **bool** | Specifies if the folder is indexed or not. | [optional] 
**DenyDownload** | Pointer to **bool** | Specifies if the folder can be downloaded or not. | [optional] 
**Lifetime** | Pointer to [**RoomDataLifetimeDto**](RoomDataLifetimeDto.md) | The room data lifetime settings of the folder. | [optional] 
**Watermark** | Pointer to [**WatermarkDto**](WatermarkDto.md) | The watermark settings of the folder. | [optional] 
**Type** | Pointer to [**FolderType**](FolderType.md) | The folder type. | [optional] 
**InRoom** | Pointer to **NullableBool** | Specifies if the folder is placed in the room or not. | [optional] 
**QuotaLimit** | Pointer to **NullableInt64** | The folder quota limit. | [optional] 
**IsCustomQuota** | Pointer to **NullableBool** | Specifies if the folder room has a custom quota or not. | [optional] 
**UsedSpace** | Pointer to **NullableInt64** | How much folder space is used (counter). | [optional] 
**PasswordProtected** | Pointer to **NullableBool** | Specifies if the folder is password protected or not. | [optional] 
**Expired** | Pointer to **NullableBool** | Specifies if an external link to the folder is expired or not. | [optional] 
**ChatSettings** | Pointer to [**ChatSettingsDto**](ChatSettingsDto.md) | The AI chat settings for the folder room. Contains configuration for AI provider, model selection, and custom prompts.  Only applicable to rooms with AI chat functionality enabled. Null if the room does not have chat settings configured. | [optional] 
**RootRoomType** | Pointer to [**RoomType**](RoomType.md) | The room type of the root folder. Indicates the type of the parent room if the current folder is nested within a room hierarchy.  This property helps identify the context in which a nested folder exists. | [optional] 
**SaveFormAsXLSX** | Pointer to **NullableBool** | Specifies whether to save form data as XLSX file. | [optional] 
**SendFormToExternalDB** | Pointer to **NullableBool** | Specifies whether to send form data to external database. | [optional] 
**OriginalFormId** | Pointer to **NullableInt32** | The original form ID that corresponds to this FormFillingFolderDone folder. | [optional] 

## Methods

### NewFolderDtoString

`func NewFolderDtoString() *FolderDtoString`

NewFolderDtoString instantiates a new FolderDtoString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderDtoStringWithDefaults

`func NewFolderDtoStringWithDefaults() *FolderDtoString`

NewFolderDtoStringWithDefaults instantiates a new FolderDtoString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FolderDtoString) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FolderDtoString) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FolderDtoString) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FolderDtoString) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAccess

`func (o *FolderDtoString) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FolderDtoString) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FolderDtoString) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FolderDtoString) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *FolderDtoString) GetSharedBy() EmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *FolderDtoString) GetSharedByOk() (*EmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *FolderDtoString) SetSharedBy(v EmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *FolderDtoString) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *FolderDtoString) GetOwnedBy() EmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *FolderDtoString) GetOwnedByOk() (*EmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *FolderDtoString) SetOwnedBy(v EmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *FolderDtoString) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *FolderDtoString) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *FolderDtoString) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *FolderDtoString) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *FolderDtoString) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *FolderDtoString) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *FolderDtoString) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *FolderDtoString) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *FolderDtoString) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *FolderDtoString) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *FolderDtoString) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *FolderDtoString) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *FolderDtoString) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *FolderDtoString) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *FolderDtoString) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *FolderDtoString) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *FolderDtoString) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *FolderDtoString) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *FolderDtoString) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *FolderDtoString) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *FolderDtoString) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### GetCreated

`func (o *FolderDtoString) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FolderDtoString) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FolderDtoString) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FolderDtoString) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FolderDtoString) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FolderDtoString) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FolderDtoString) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FolderDtoString) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *FolderDtoString) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FolderDtoString) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FolderDtoString) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FolderDtoString) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAutoDelete

`func (o *FolderDtoString) GetAutoDelete() time.Time`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *FolderDtoString) GetAutoDeleteOk() (*time.Time, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *FolderDtoString) SetAutoDelete(v time.Time)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *FolderDtoString) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetRootFolderType

`func (o *FolderDtoString) GetRootFolderType() FolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *FolderDtoString) GetRootFolderTypeOk() (*FolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *FolderDtoString) SetRootFolderType(v FolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *FolderDtoString) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *FolderDtoString) GetParentRoomType() FolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *FolderDtoString) GetParentRoomTypeOk() (*FolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *FolderDtoString) SetParentRoomType(v FolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *FolderDtoString) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FolderDtoString) GetUpdatedBy() EmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FolderDtoString) GetUpdatedByOk() (*EmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FolderDtoString) SetUpdatedBy(v EmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FolderDtoString) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *FolderDtoString) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *FolderDtoString) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *FolderDtoString) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *FolderDtoString) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### GetProviderKey

`func (o *FolderDtoString) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *FolderDtoString) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *FolderDtoString) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *FolderDtoString) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### GetProviderId

`func (o *FolderDtoString) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *FolderDtoString) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *FolderDtoString) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *FolderDtoString) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetOrder

`func (o *FolderDtoString) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FolderDtoString) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FolderDtoString) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FolderDtoString) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIsFavorite

`func (o *FolderDtoString) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *FolderDtoString) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *FolderDtoString) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *FolderDtoString) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetFileEntryType

`func (o *FolderDtoString) GetFileEntryType() FileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *FolderDtoString) GetFileEntryTypeOk() (*FileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *FolderDtoString) SetFileEntryType(v FileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *FolderDtoString) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.

### GetId

`func (o *FolderDtoString) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FolderDtoString) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FolderDtoString) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FolderDtoString) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootFolderId

`func (o *FolderDtoString) GetRootFolderId() string`

GetRootFolderId returns the RootFolderId field if non-nil, zero value otherwise.

### GetRootFolderIdOk

`func (o *FolderDtoString) GetRootFolderIdOk() (*string, bool)`

GetRootFolderIdOk returns a tuple with the RootFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderId

`func (o *FolderDtoString) SetRootFolderId(v string)`

SetRootFolderId sets RootFolderId field to given value.

### HasRootFolderId

`func (o *FolderDtoString) HasRootFolderId() bool`

HasRootFolderId returns a boolean if a field has been set.

### GetOriginId

`func (o *FolderDtoString) GetOriginId() string`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *FolderDtoString) GetOriginIdOk() (*string, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *FolderDtoString) SetOriginId(v string)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *FolderDtoString) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginRoomId

`func (o *FolderDtoString) GetOriginRoomId() string`

GetOriginRoomId returns the OriginRoomId field if non-nil, zero value otherwise.

### GetOriginRoomIdOk

`func (o *FolderDtoString) GetOriginRoomIdOk() (*string, bool)`

GetOriginRoomIdOk returns a tuple with the OriginRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomId

`func (o *FolderDtoString) SetOriginRoomId(v string)`

SetOriginRoomId sets OriginRoomId field to given value.

### HasOriginRoomId

`func (o *FolderDtoString) HasOriginRoomId() bool`

HasOriginRoomId returns a boolean if a field has been set.

### GetOriginTitle

`func (o *FolderDtoString) GetOriginTitle() string`

GetOriginTitle returns the OriginTitle field if non-nil, zero value otherwise.

### GetOriginTitleOk

`func (o *FolderDtoString) GetOriginTitleOk() (*string, bool)`

GetOriginTitleOk returns a tuple with the OriginTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTitle

`func (o *FolderDtoString) SetOriginTitle(v string)`

SetOriginTitle sets OriginTitle field to given value.

### HasOriginTitle

`func (o *FolderDtoString) HasOriginTitle() bool`

HasOriginTitle returns a boolean if a field has been set.

### GetOriginRoomTitle

`func (o *FolderDtoString) GetOriginRoomTitle() string`

GetOriginRoomTitle returns the OriginRoomTitle field if non-nil, zero value otherwise.

### GetOriginRoomTitleOk

`func (o *FolderDtoString) GetOriginRoomTitleOk() (*string, bool)`

GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomTitle

`func (o *FolderDtoString) SetOriginRoomTitle(v string)`

SetOriginRoomTitle sets OriginRoomTitle field to given value.

### HasOriginRoomTitle

`func (o *FolderDtoString) HasOriginRoomTitle() bool`

HasOriginRoomTitle returns a boolean if a field has been set.

### GetCanShare

`func (o *FolderDtoString) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *FolderDtoString) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *FolderDtoString) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *FolderDtoString) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### GetShareSettings

`func (o *FolderDtoString) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings`

GetShareSettings returns the ShareSettings field if non-nil, zero value otherwise.

### GetShareSettingsOk

`func (o *FolderDtoString) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool)`

GetShareSettingsOk returns a tuple with the ShareSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareSettings

`func (o *FolderDtoString) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings)`

SetShareSettings sets ShareSettings field to given value.

### HasShareSettings

`func (o *FolderDtoString) HasShareSettings() bool`

HasShareSettings returns a boolean if a field has been set.

### SetShareSettingsNil

`func (o *FolderDtoString) SetShareSettingsNil(b bool)`

 SetShareSettingsNil sets the value for ShareSettings to be an explicit nil

### UnsetShareSettings
`func (o *FolderDtoString) UnsetShareSettings()`

UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
### GetSecurity

`func (o *FolderDtoString) GetSecurity() FileEntryDtoIntegerAllOfSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *FolderDtoString) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *FolderDtoString) SetSecurity(v FileEntryDtoIntegerAllOfSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *FolderDtoString) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *FolderDtoString) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *FolderDtoString) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetAvailableShareRights

`func (o *FolderDtoString) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights`

GetAvailableShareRights returns the AvailableShareRights field if non-nil, zero value otherwise.

### GetAvailableShareRightsOk

`func (o *FolderDtoString) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool)`

GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShareRights

`func (o *FolderDtoString) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights)`

SetAvailableShareRights sets AvailableShareRights field to given value.

### HasAvailableShareRights

`func (o *FolderDtoString) HasAvailableShareRights() bool`

HasAvailableShareRights returns a boolean if a field has been set.

### SetAvailableShareRightsNil

`func (o *FolderDtoString) SetAvailableShareRightsNil(b bool)`

 SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil

### UnsetAvailableShareRights
`func (o *FolderDtoString) UnsetAvailableShareRights()`

UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
### GetRequestToken

`func (o *FolderDtoString) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *FolderDtoString) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *FolderDtoString) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *FolderDtoString) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### GetExternal

`func (o *FolderDtoString) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *FolderDtoString) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *FolderDtoString) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *FolderDtoString) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExpirationDate

`func (o *FolderDtoString) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FolderDtoString) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FolderDtoString) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FolderDtoString) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIsLinkExpired

`func (o *FolderDtoString) GetIsLinkExpired() bool`

GetIsLinkExpired returns the IsLinkExpired field if non-nil, zero value otherwise.

### GetIsLinkExpiredOk

`func (o *FolderDtoString) GetIsLinkExpiredOk() (*bool, bool)`

GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkExpired

`func (o *FolderDtoString) SetIsLinkExpired(v bool)`

SetIsLinkExpired sets IsLinkExpired field to given value.

### HasIsLinkExpired

`func (o *FolderDtoString) HasIsLinkExpired() bool`

HasIsLinkExpired returns a boolean if a field has been set.

### GetParentId

`func (o *FolderDtoString) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FolderDtoString) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FolderDtoString) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FolderDtoString) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *FolderDtoString) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *FolderDtoString) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetFilesCount

`func (o *FolderDtoString) GetFilesCount() int32`

GetFilesCount returns the FilesCount field if non-nil, zero value otherwise.

### GetFilesCountOk

`func (o *FolderDtoString) GetFilesCountOk() (*int32, bool)`

GetFilesCountOk returns a tuple with the FilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesCount

`func (o *FolderDtoString) SetFilesCount(v int32)`

SetFilesCount sets FilesCount field to given value.

### HasFilesCount

`func (o *FolderDtoString) HasFilesCount() bool`

HasFilesCount returns a boolean if a field has been set.

### GetFoldersCount

`func (o *FolderDtoString) GetFoldersCount() int32`

GetFoldersCount returns the FoldersCount field if non-nil, zero value otherwise.

### GetFoldersCountOk

`func (o *FolderDtoString) GetFoldersCountOk() (*int32, bool)`

GetFoldersCountOk returns a tuple with the FoldersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldersCount

`func (o *FolderDtoString) SetFoldersCount(v int32)`

SetFoldersCount sets FoldersCount field to given value.

### HasFoldersCount

`func (o *FolderDtoString) HasFoldersCount() bool`

HasFoldersCount returns a boolean if a field has been set.

### GetIsShareable

`func (o *FolderDtoString) GetIsShareable() bool`

GetIsShareable returns the IsShareable field if non-nil, zero value otherwise.

### GetIsShareableOk

`func (o *FolderDtoString) GetIsShareableOk() (*bool, bool)`

GetIsShareableOk returns a tuple with the IsShareable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShareable

`func (o *FolderDtoString) SetIsShareable(v bool)`

SetIsShareable sets IsShareable field to given value.

### HasIsShareable

`func (o *FolderDtoString) HasIsShareable() bool`

HasIsShareable returns a boolean if a field has been set.

### SetIsShareableNil

`func (o *FolderDtoString) SetIsShareableNil(b bool)`

 SetIsShareableNil sets the value for IsShareable to be an explicit nil

### UnsetIsShareable
`func (o *FolderDtoString) UnsetIsShareable()`

UnsetIsShareable ensures that no value is present for IsShareable, not even an explicit nil
### GetNew

`func (o *FolderDtoString) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *FolderDtoString) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *FolderDtoString) SetNew(v int32)`

SetNew sets New field to given value.

### HasNew

`func (o *FolderDtoString) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetMute

`func (o *FolderDtoString) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *FolderDtoString) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *FolderDtoString) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *FolderDtoString) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetTags

`func (o *FolderDtoString) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FolderDtoString) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FolderDtoString) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FolderDtoString) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *FolderDtoString) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *FolderDtoString) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetLogo

`func (o *FolderDtoString) GetLogo() Logo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *FolderDtoString) GetLogoOk() (*Logo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *FolderDtoString) SetLogo(v Logo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *FolderDtoString) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPinned

`func (o *FolderDtoString) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *FolderDtoString) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *FolderDtoString) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *FolderDtoString) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetRoomType

`func (o *FolderDtoString) GetRoomType() RoomType`

GetRoomType returns the RoomType field if non-nil, zero value otherwise.

### GetRoomTypeOk

`func (o *FolderDtoString) GetRoomTypeOk() (*RoomType, bool)`

GetRoomTypeOk returns a tuple with the RoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomType

`func (o *FolderDtoString) SetRoomType(v RoomType)`

SetRoomType sets RoomType field to given value.

### HasRoomType

`func (o *FolderDtoString) HasRoomType() bool`

HasRoomType returns a boolean if a field has been set.

### GetPrivate

`func (o *FolderDtoString) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *FolderDtoString) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *FolderDtoString) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *FolderDtoString) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetIndexing

`func (o *FolderDtoString) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *FolderDtoString) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *FolderDtoString) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *FolderDtoString) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### GetDenyDownload

`func (o *FolderDtoString) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *FolderDtoString) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *FolderDtoString) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *FolderDtoString) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetLifetime

`func (o *FolderDtoString) GetLifetime() RoomDataLifetimeDto`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *FolderDtoString) GetLifetimeOk() (*RoomDataLifetimeDto, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *FolderDtoString) SetLifetime(v RoomDataLifetimeDto)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *FolderDtoString) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetWatermark

`func (o *FolderDtoString) GetWatermark() WatermarkDto`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *FolderDtoString) GetWatermarkOk() (*WatermarkDto, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *FolderDtoString) SetWatermark(v WatermarkDto)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *FolderDtoString) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetType

`func (o *FolderDtoString) GetType() FolderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FolderDtoString) GetTypeOk() (*FolderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FolderDtoString) SetType(v FolderType)`

SetType sets Type field to given value.

### HasType

`func (o *FolderDtoString) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInRoom

`func (o *FolderDtoString) GetInRoom() bool`

GetInRoom returns the InRoom field if non-nil, zero value otherwise.

### GetInRoomOk

`func (o *FolderDtoString) GetInRoomOk() (*bool, bool)`

GetInRoomOk returns a tuple with the InRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInRoom

`func (o *FolderDtoString) SetInRoom(v bool)`

SetInRoom sets InRoom field to given value.

### HasInRoom

`func (o *FolderDtoString) HasInRoom() bool`

HasInRoom returns a boolean if a field has been set.

### SetInRoomNil

`func (o *FolderDtoString) SetInRoomNil(b bool)`

 SetInRoomNil sets the value for InRoom to be an explicit nil

### UnsetInRoom
`func (o *FolderDtoString) UnsetInRoom()`

UnsetInRoom ensures that no value is present for InRoom, not even an explicit nil
### GetQuotaLimit

`func (o *FolderDtoString) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *FolderDtoString) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *FolderDtoString) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *FolderDtoString) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### SetQuotaLimitNil

`func (o *FolderDtoString) SetQuotaLimitNil(b bool)`

 SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil

### UnsetQuotaLimit
`func (o *FolderDtoString) UnsetQuotaLimit()`

UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
### GetIsCustomQuota

`func (o *FolderDtoString) GetIsCustomQuota() bool`

GetIsCustomQuota returns the IsCustomQuota field if non-nil, zero value otherwise.

### GetIsCustomQuotaOk

`func (o *FolderDtoString) GetIsCustomQuotaOk() (*bool, bool)`

GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomQuota

`func (o *FolderDtoString) SetIsCustomQuota(v bool)`

SetIsCustomQuota sets IsCustomQuota field to given value.

### HasIsCustomQuota

`func (o *FolderDtoString) HasIsCustomQuota() bool`

HasIsCustomQuota returns a boolean if a field has been set.

### SetIsCustomQuotaNil

`func (o *FolderDtoString) SetIsCustomQuotaNil(b bool)`

 SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil

### UnsetIsCustomQuota
`func (o *FolderDtoString) UnsetIsCustomQuota()`

UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
### GetUsedSpace

`func (o *FolderDtoString) GetUsedSpace() int64`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *FolderDtoString) GetUsedSpaceOk() (*int64, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *FolderDtoString) SetUsedSpace(v int64)`

SetUsedSpace sets UsedSpace field to given value.

### HasUsedSpace

`func (o *FolderDtoString) HasUsedSpace() bool`

HasUsedSpace returns a boolean if a field has been set.

### SetUsedSpaceNil

`func (o *FolderDtoString) SetUsedSpaceNil(b bool)`

 SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil

### UnsetUsedSpace
`func (o *FolderDtoString) UnsetUsedSpace()`

UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
### GetPasswordProtected

`func (o *FolderDtoString) GetPasswordProtected() bool`

GetPasswordProtected returns the PasswordProtected field if non-nil, zero value otherwise.

### GetPasswordProtectedOk

`func (o *FolderDtoString) GetPasswordProtectedOk() (*bool, bool)`

GetPasswordProtectedOk returns a tuple with the PasswordProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProtected

`func (o *FolderDtoString) SetPasswordProtected(v bool)`

SetPasswordProtected sets PasswordProtected field to given value.

### HasPasswordProtected

`func (o *FolderDtoString) HasPasswordProtected() bool`

HasPasswordProtected returns a boolean if a field has been set.

### SetPasswordProtectedNil

`func (o *FolderDtoString) SetPasswordProtectedNil(b bool)`

 SetPasswordProtectedNil sets the value for PasswordProtected to be an explicit nil

### UnsetPasswordProtected
`func (o *FolderDtoString) UnsetPasswordProtected()`

UnsetPasswordProtected ensures that no value is present for PasswordProtected, not even an explicit nil
### GetExpired

`func (o *FolderDtoString) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *FolderDtoString) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *FolderDtoString) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *FolderDtoString) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### SetExpiredNil

`func (o *FolderDtoString) SetExpiredNil(b bool)`

 SetExpiredNil sets the value for Expired to be an explicit nil

### UnsetExpired
`func (o *FolderDtoString) UnsetExpired()`

UnsetExpired ensures that no value is present for Expired, not even an explicit nil
### GetChatSettings

`func (o *FolderDtoString) GetChatSettings() ChatSettingsDto`

GetChatSettings returns the ChatSettings field if non-nil, zero value otherwise.

### GetChatSettingsOk

`func (o *FolderDtoString) GetChatSettingsOk() (*ChatSettingsDto, bool)`

GetChatSettingsOk returns a tuple with the ChatSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSettings

`func (o *FolderDtoString) SetChatSettings(v ChatSettingsDto)`

SetChatSettings sets ChatSettings field to given value.

### HasChatSettings

`func (o *FolderDtoString) HasChatSettings() bool`

HasChatSettings returns a boolean if a field has been set.

### GetRootRoomType

`func (o *FolderDtoString) GetRootRoomType() RoomType`

GetRootRoomType returns the RootRoomType field if non-nil, zero value otherwise.

### GetRootRoomTypeOk

`func (o *FolderDtoString) GetRootRoomTypeOk() (*RoomType, bool)`

GetRootRoomTypeOk returns a tuple with the RootRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRoomType

`func (o *FolderDtoString) SetRootRoomType(v RoomType)`

SetRootRoomType sets RootRoomType field to given value.

### HasRootRoomType

`func (o *FolderDtoString) HasRootRoomType() bool`

HasRootRoomType returns a boolean if a field has been set.

### GetSaveFormAsXLSX

`func (o *FolderDtoString) GetSaveFormAsXLSX() bool`

GetSaveFormAsXLSX returns the SaveFormAsXLSX field if non-nil, zero value otherwise.

### GetSaveFormAsXLSXOk

`func (o *FolderDtoString) GetSaveFormAsXLSXOk() (*bool, bool)`

GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveFormAsXLSX

`func (o *FolderDtoString) SetSaveFormAsXLSX(v bool)`

SetSaveFormAsXLSX sets SaveFormAsXLSX field to given value.

### HasSaveFormAsXLSX

`func (o *FolderDtoString) HasSaveFormAsXLSX() bool`

HasSaveFormAsXLSX returns a boolean if a field has been set.

### SetSaveFormAsXLSXNil

`func (o *FolderDtoString) SetSaveFormAsXLSXNil(b bool)`

 SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil

### UnsetSaveFormAsXLSX
`func (o *FolderDtoString) UnsetSaveFormAsXLSX()`

UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil
### GetSendFormToExternalDB

`func (o *FolderDtoString) GetSendFormToExternalDB() bool`

GetSendFormToExternalDB returns the SendFormToExternalDB field if non-nil, zero value otherwise.

### GetSendFormToExternalDBOk

`func (o *FolderDtoString) GetSendFormToExternalDBOk() (*bool, bool)`

GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFormToExternalDB

`func (o *FolderDtoString) SetSendFormToExternalDB(v bool)`

SetSendFormToExternalDB sets SendFormToExternalDB field to given value.

### HasSendFormToExternalDB

`func (o *FolderDtoString) HasSendFormToExternalDB() bool`

HasSendFormToExternalDB returns a boolean if a field has been set.

### SetSendFormToExternalDBNil

`func (o *FolderDtoString) SetSendFormToExternalDBNil(b bool)`

 SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil

### UnsetSendFormToExternalDB
`func (o *FolderDtoString) UnsetSendFormToExternalDB()`

UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
### GetOriginalFormId

`func (o *FolderDtoString) GetOriginalFormId() int32`

GetOriginalFormId returns the OriginalFormId field if non-nil, zero value otherwise.

### GetOriginalFormIdOk

`func (o *FolderDtoString) GetOriginalFormIdOk() (*int32, bool)`

GetOriginalFormIdOk returns a tuple with the OriginalFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFormId

`func (o *FolderDtoString) SetOriginalFormId(v int32)`

SetOriginalFormId sets OriginalFormId field to given value.

### HasOriginalFormId

`func (o *FolderDtoString) HasOriginalFormId() bool`

HasOriginalFormId returns a boolean if a field has been set.

### SetOriginalFormIdNil

`func (o *FolderDtoString) SetOriginalFormIdNil(b bool)`

 SetOriginalFormIdNil sets the value for OriginalFormId to be an explicit nil

### UnsetOriginalFormId
`func (o *FolderDtoString) UnsetOriginalFormId()`

UnsetOriginalFormId ensures that no value is present for OriginalFormId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


