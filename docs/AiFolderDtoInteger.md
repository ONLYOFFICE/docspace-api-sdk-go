# AiFolderDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The file entry title. | [optional] 
**Access** | Pointer to [**AiFileShare**](AiFileShare.md) | The access rights to the file entry. | [optional] 
**SharedBy** | Pointer to [**AiEmployeeDto**](AiEmployeeDto.md) | Provides information about the employee who shared the file or folder. | [optional] 
**OwnedBy** | Pointer to [**AiEmployeeDto**](AiEmployeeDto.md) | The information about the employee who owns the file entry. | [optional] 
**Shared** | Pointer to **bool** | Specifies if the file entry is shared via link or not. | [optional] 
**SharedForUser** | Pointer to **bool** | Specifies if the file entry is shared for user or not. | [optional] 
**SharedExternal** | Pointer to **bool** | Specifies if the file entry is shared via a public (non-internal) external link. | [optional] 
**ParentShared** | Pointer to **bool** | Indicates whether the parent entity is shared. | [optional] 
**ShortWebUrl** | Pointer to **string** | The short Web URL. | [optional] 
**Created** | Pointer to **time.Time** | The creation date and time of the file entry. | [optional] 
**CreatedBy** | Pointer to [**AiEmployeeDto**](AiEmployeeDto.md) | The file entry author. | [optional] 
**Updated** | Pointer to **time.Time** | The last date and time when the file entry was updated. | [optional] 
**AutoDelete** | Pointer to **time.Time** | The date and time when the file entry will be automatically deleted. | [optional] 
**RootFolderType** | Pointer to [**AiFolderType**](AiFolderType.md) | The root folder type of the file entry. | [optional] 
**ParentRoomType** | Pointer to [**AiFolderType**](AiFolderType.md) | The parent room type of the file entry. | [optional] 
**UpdatedBy** | Pointer to [**AiEmployeeDto**](AiEmployeeDto.md) | The user who updated the file entry. | [optional] 
**ProviderItem** | Pointer to **bool** | Specifies if the file entry provider is specified or not. | [optional] 
**ProviderKey** | Pointer to **string** | The provider key of the file entry. | [optional] 
**ProviderId** | Pointer to **int32** | The provider ID of the file entry. | [optional] 
**Order** | Pointer to **string** | The order of the file entry. | [optional] 
**IsFavorite** | Pointer to **bool** | Specifies if the file is a favorite or not. | [optional] 
**FileEntryType** | Pointer to [**AiFileEntryType**](AiFileEntryType.md) | The file entry type. | [optional] 
**Id** | Pointer to **int32** | The file entry ID. | [optional] 
**RootFolderId** | Pointer to **int32** | The root folder ID of the file entry. | [optional] 
**OriginId** | Pointer to **int32** | The origin ID of the file entry. | [optional] 
**OriginRoomId** | Pointer to **int32** | The origin room ID of the file entry. | [optional] 
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
**ParentId** | Pointer to **int32** | The parent folder ID of the folder. | [optional] 
**FilesCount** | Pointer to **int32** | The number of files that the folder contains. | [optional] 
**FoldersCount** | Pointer to **int32** | The number of folders that the folder contains. | [optional] 
**IsShareable** | Pointer to **NullableBool** | Specifies if the folder can be shared or not. | [optional] 
**New** | Pointer to **int32** | The new element index in the folder. | [optional] 
**Mute** | Pointer to **bool** | Specifies if the folder notifications are enabled or not. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of the folder. | [optional] 
**Logo** | Pointer to [**AiLogo**](AiLogo.md) | The folder logo. | [optional] 
**Pinned** | Pointer to **bool** | Specifies if the folder is pinned or not. | [optional] 
**RoomType** | Pointer to [**AiRoomType**](AiRoomType.md) | The room type of the folder. | [optional] 
**Private** | Pointer to **bool** | Specifies if the folder is private or not. | [optional] 
**Indexing** | Pointer to **bool** | Specifies if the folder is indexed or not. | [optional] 
**DenyDownload** | Pointer to **bool** | Specifies if the folder can be downloaded or not. | [optional] 
**Lifetime** | Pointer to [**AiRoomDataLifetimeDto**](AiRoomDataLifetimeDto.md) | The room data lifetime settings of the folder. | [optional] 
**Watermark** | Pointer to [**AiWatermarkDto**](AiWatermarkDto.md) | The watermark settings of the folder. | [optional] 
**Type** | Pointer to [**AiFolderType**](AiFolderType.md) | The folder type. | [optional] 
**InRoom** | Pointer to **NullableBool** | Specifies if the folder is placed in the room or not. | [optional] 
**QuotaLimit** | Pointer to **NullableInt64** | The folder quota limit. | [optional] 
**IsCustomQuota** | Pointer to **NullableBool** | Specifies if the folder room has a custom quota or not. | [optional] 
**UsedSpace** | Pointer to **NullableInt64** | How much folder space is used (counter). | [optional] 
**PasswordProtected** | Pointer to **NullableBool** | Specifies if the folder is password protected or not. | [optional] 
**Expired** | Pointer to **NullableBool** | Specifies if an external link to the folder is expired or not. | [optional] 
**ChatSettings** | Pointer to [**AiChatSettingsDto**](AiChatSettingsDto.md) | The AI chat settings for the folder room. Contains configuration for AI provider, model selection, and custom prompts.  Only applicable to rooms with AI chat functionality enabled. Null if the room does not have chat settings configured. | [optional] 
**RootRoomType** | Pointer to [**AiRoomType**](AiRoomType.md) | The room type of the root folder. Indicates the type of the parent room if the current folder is nested within a room hierarchy.  This property helps identify the context in which a nested folder exists. | [optional] 
**SaveFormAsXLSX** | Pointer to **NullableBool** | Specifies whether to save form data as XLSX file. | [optional] 
**SendFormToExternalDB** | Pointer to **NullableBool** | Specifies whether to send form data to external database. | [optional] 
**OriginalFormId** | Pointer to **NullableInt32** | The original form ID that corresponds to this FormFillingFolderDone folder. | [optional] 

## Methods

### NewAiFolderDtoInteger

`func NewAiFolderDtoInteger() *AiFolderDtoInteger`

NewAiFolderDtoInteger instantiates a new AiFolderDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiFolderDtoIntegerWithDefaults

`func NewAiFolderDtoIntegerWithDefaults() *AiFolderDtoInteger`

NewAiFolderDtoIntegerWithDefaults instantiates a new AiFolderDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AiFolderDtoInteger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiFolderDtoInteger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiFolderDtoInteger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AiFolderDtoInteger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAccess

`func (o *AiFolderDtoInteger) GetAccess() AiFileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AiFolderDtoInteger) GetAccessOk() (*AiFileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AiFolderDtoInteger) SetAccess(v AiFileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AiFolderDtoInteger) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *AiFolderDtoInteger) GetSharedBy() AiEmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *AiFolderDtoInteger) GetSharedByOk() (*AiEmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *AiFolderDtoInteger) SetSharedBy(v AiEmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *AiFolderDtoInteger) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *AiFolderDtoInteger) GetOwnedBy() AiEmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *AiFolderDtoInteger) GetOwnedByOk() (*AiEmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *AiFolderDtoInteger) SetOwnedBy(v AiEmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *AiFolderDtoInteger) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *AiFolderDtoInteger) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *AiFolderDtoInteger) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *AiFolderDtoInteger) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *AiFolderDtoInteger) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *AiFolderDtoInteger) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *AiFolderDtoInteger) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *AiFolderDtoInteger) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *AiFolderDtoInteger) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *AiFolderDtoInteger) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *AiFolderDtoInteger) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *AiFolderDtoInteger) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *AiFolderDtoInteger) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *AiFolderDtoInteger) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *AiFolderDtoInteger) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *AiFolderDtoInteger) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *AiFolderDtoInteger) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *AiFolderDtoInteger) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *AiFolderDtoInteger) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *AiFolderDtoInteger) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *AiFolderDtoInteger) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### GetCreated

`func (o *AiFolderDtoInteger) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiFolderDtoInteger) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiFolderDtoInteger) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AiFolderDtoInteger) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AiFolderDtoInteger) GetCreatedBy() AiEmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AiFolderDtoInteger) GetCreatedByOk() (*AiEmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AiFolderDtoInteger) SetCreatedBy(v AiEmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AiFolderDtoInteger) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *AiFolderDtoInteger) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AiFolderDtoInteger) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AiFolderDtoInteger) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AiFolderDtoInteger) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAutoDelete

`func (o *AiFolderDtoInteger) GetAutoDelete() time.Time`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *AiFolderDtoInteger) GetAutoDeleteOk() (*time.Time, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *AiFolderDtoInteger) SetAutoDelete(v time.Time)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *AiFolderDtoInteger) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetRootFolderType

`func (o *AiFolderDtoInteger) GetRootFolderType() AiFolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *AiFolderDtoInteger) GetRootFolderTypeOk() (*AiFolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *AiFolderDtoInteger) SetRootFolderType(v AiFolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *AiFolderDtoInteger) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *AiFolderDtoInteger) GetParentRoomType() AiFolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *AiFolderDtoInteger) GetParentRoomTypeOk() (*AiFolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *AiFolderDtoInteger) SetParentRoomType(v AiFolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *AiFolderDtoInteger) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AiFolderDtoInteger) GetUpdatedBy() AiEmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AiFolderDtoInteger) GetUpdatedByOk() (*AiEmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AiFolderDtoInteger) SetUpdatedBy(v AiEmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AiFolderDtoInteger) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *AiFolderDtoInteger) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *AiFolderDtoInteger) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *AiFolderDtoInteger) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *AiFolderDtoInteger) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### GetProviderKey

`func (o *AiFolderDtoInteger) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *AiFolderDtoInteger) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *AiFolderDtoInteger) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *AiFolderDtoInteger) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### GetProviderId

`func (o *AiFolderDtoInteger) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *AiFolderDtoInteger) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *AiFolderDtoInteger) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *AiFolderDtoInteger) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetOrder

`func (o *AiFolderDtoInteger) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AiFolderDtoInteger) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AiFolderDtoInteger) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AiFolderDtoInteger) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIsFavorite

`func (o *AiFolderDtoInteger) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *AiFolderDtoInteger) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *AiFolderDtoInteger) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *AiFolderDtoInteger) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetFileEntryType

`func (o *AiFolderDtoInteger) GetFileEntryType() AiFileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *AiFolderDtoInteger) GetFileEntryTypeOk() (*AiFileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *AiFolderDtoInteger) SetFileEntryType(v AiFileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *AiFolderDtoInteger) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.

### GetId

`func (o *AiFolderDtoInteger) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiFolderDtoInteger) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiFolderDtoInteger) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AiFolderDtoInteger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootFolderId

`func (o *AiFolderDtoInteger) GetRootFolderId() int32`

GetRootFolderId returns the RootFolderId field if non-nil, zero value otherwise.

### GetRootFolderIdOk

`func (o *AiFolderDtoInteger) GetRootFolderIdOk() (*int32, bool)`

GetRootFolderIdOk returns a tuple with the RootFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderId

`func (o *AiFolderDtoInteger) SetRootFolderId(v int32)`

SetRootFolderId sets RootFolderId field to given value.

### HasRootFolderId

`func (o *AiFolderDtoInteger) HasRootFolderId() bool`

HasRootFolderId returns a boolean if a field has been set.

### GetOriginId

`func (o *AiFolderDtoInteger) GetOriginId() int32`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *AiFolderDtoInteger) GetOriginIdOk() (*int32, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *AiFolderDtoInteger) SetOriginId(v int32)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *AiFolderDtoInteger) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginRoomId

`func (o *AiFolderDtoInteger) GetOriginRoomId() int32`

GetOriginRoomId returns the OriginRoomId field if non-nil, zero value otherwise.

### GetOriginRoomIdOk

`func (o *AiFolderDtoInteger) GetOriginRoomIdOk() (*int32, bool)`

GetOriginRoomIdOk returns a tuple with the OriginRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomId

`func (o *AiFolderDtoInteger) SetOriginRoomId(v int32)`

SetOriginRoomId sets OriginRoomId field to given value.

### HasOriginRoomId

`func (o *AiFolderDtoInteger) HasOriginRoomId() bool`

HasOriginRoomId returns a boolean if a field has been set.

### GetOriginTitle

`func (o *AiFolderDtoInteger) GetOriginTitle() string`

GetOriginTitle returns the OriginTitle field if non-nil, zero value otherwise.

### GetOriginTitleOk

`func (o *AiFolderDtoInteger) GetOriginTitleOk() (*string, bool)`

GetOriginTitleOk returns a tuple with the OriginTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTitle

`func (o *AiFolderDtoInteger) SetOriginTitle(v string)`

SetOriginTitle sets OriginTitle field to given value.

### HasOriginTitle

`func (o *AiFolderDtoInteger) HasOriginTitle() bool`

HasOriginTitle returns a boolean if a field has been set.

### GetOriginRoomTitle

`func (o *AiFolderDtoInteger) GetOriginRoomTitle() string`

GetOriginRoomTitle returns the OriginRoomTitle field if non-nil, zero value otherwise.

### GetOriginRoomTitleOk

`func (o *AiFolderDtoInteger) GetOriginRoomTitleOk() (*string, bool)`

GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomTitle

`func (o *AiFolderDtoInteger) SetOriginRoomTitle(v string)`

SetOriginRoomTitle sets OriginRoomTitle field to given value.

### HasOriginRoomTitle

`func (o *AiFolderDtoInteger) HasOriginRoomTitle() bool`

HasOriginRoomTitle returns a boolean if a field has been set.

### GetCanShare

`func (o *AiFolderDtoInteger) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *AiFolderDtoInteger) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *AiFolderDtoInteger) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *AiFolderDtoInteger) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### GetShareSettings

`func (o *AiFolderDtoInteger) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings`

GetShareSettings returns the ShareSettings field if non-nil, zero value otherwise.

### GetShareSettingsOk

`func (o *AiFolderDtoInteger) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool)`

GetShareSettingsOk returns a tuple with the ShareSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareSettings

`func (o *AiFolderDtoInteger) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings)`

SetShareSettings sets ShareSettings field to given value.

### HasShareSettings

`func (o *AiFolderDtoInteger) HasShareSettings() bool`

HasShareSettings returns a boolean if a field has been set.

### SetShareSettingsNil

`func (o *AiFolderDtoInteger) SetShareSettingsNil(b bool)`

 SetShareSettingsNil sets the value for ShareSettings to be an explicit nil

### UnsetShareSettings
`func (o *AiFolderDtoInteger) UnsetShareSettings()`

UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
### GetSecurity

`func (o *AiFolderDtoInteger) GetSecurity() FileEntryDtoIntegerAllOfSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *AiFolderDtoInteger) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *AiFolderDtoInteger) SetSecurity(v FileEntryDtoIntegerAllOfSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *AiFolderDtoInteger) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *AiFolderDtoInteger) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *AiFolderDtoInteger) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetAvailableShareRights

`func (o *AiFolderDtoInteger) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights`

GetAvailableShareRights returns the AvailableShareRights field if non-nil, zero value otherwise.

### GetAvailableShareRightsOk

`func (o *AiFolderDtoInteger) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool)`

GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShareRights

`func (o *AiFolderDtoInteger) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights)`

SetAvailableShareRights sets AvailableShareRights field to given value.

### HasAvailableShareRights

`func (o *AiFolderDtoInteger) HasAvailableShareRights() bool`

HasAvailableShareRights returns a boolean if a field has been set.

### SetAvailableShareRightsNil

`func (o *AiFolderDtoInteger) SetAvailableShareRightsNil(b bool)`

 SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil

### UnsetAvailableShareRights
`func (o *AiFolderDtoInteger) UnsetAvailableShareRights()`

UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
### GetRequestToken

`func (o *AiFolderDtoInteger) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *AiFolderDtoInteger) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *AiFolderDtoInteger) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *AiFolderDtoInteger) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### GetExternal

`func (o *AiFolderDtoInteger) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *AiFolderDtoInteger) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *AiFolderDtoInteger) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *AiFolderDtoInteger) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AiFolderDtoInteger) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AiFolderDtoInteger) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AiFolderDtoInteger) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AiFolderDtoInteger) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIsLinkExpired

`func (o *AiFolderDtoInteger) GetIsLinkExpired() bool`

GetIsLinkExpired returns the IsLinkExpired field if non-nil, zero value otherwise.

### GetIsLinkExpiredOk

`func (o *AiFolderDtoInteger) GetIsLinkExpiredOk() (*bool, bool)`

GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkExpired

`func (o *AiFolderDtoInteger) SetIsLinkExpired(v bool)`

SetIsLinkExpired sets IsLinkExpired field to given value.

### HasIsLinkExpired

`func (o *AiFolderDtoInteger) HasIsLinkExpired() bool`

HasIsLinkExpired returns a boolean if a field has been set.

### GetParentId

`func (o *AiFolderDtoInteger) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AiFolderDtoInteger) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AiFolderDtoInteger) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *AiFolderDtoInteger) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetFilesCount

`func (o *AiFolderDtoInteger) GetFilesCount() int32`

GetFilesCount returns the FilesCount field if non-nil, zero value otherwise.

### GetFilesCountOk

`func (o *AiFolderDtoInteger) GetFilesCountOk() (*int32, bool)`

GetFilesCountOk returns a tuple with the FilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesCount

`func (o *AiFolderDtoInteger) SetFilesCount(v int32)`

SetFilesCount sets FilesCount field to given value.

### HasFilesCount

`func (o *AiFolderDtoInteger) HasFilesCount() bool`

HasFilesCount returns a boolean if a field has been set.

### GetFoldersCount

`func (o *AiFolderDtoInteger) GetFoldersCount() int32`

GetFoldersCount returns the FoldersCount field if non-nil, zero value otherwise.

### GetFoldersCountOk

`func (o *AiFolderDtoInteger) GetFoldersCountOk() (*int32, bool)`

GetFoldersCountOk returns a tuple with the FoldersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldersCount

`func (o *AiFolderDtoInteger) SetFoldersCount(v int32)`

SetFoldersCount sets FoldersCount field to given value.

### HasFoldersCount

`func (o *AiFolderDtoInteger) HasFoldersCount() bool`

HasFoldersCount returns a boolean if a field has been set.

### GetIsShareable

`func (o *AiFolderDtoInteger) GetIsShareable() bool`

GetIsShareable returns the IsShareable field if non-nil, zero value otherwise.

### GetIsShareableOk

`func (o *AiFolderDtoInteger) GetIsShareableOk() (*bool, bool)`

GetIsShareableOk returns a tuple with the IsShareable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShareable

`func (o *AiFolderDtoInteger) SetIsShareable(v bool)`

SetIsShareable sets IsShareable field to given value.

### HasIsShareable

`func (o *AiFolderDtoInteger) HasIsShareable() bool`

HasIsShareable returns a boolean if a field has been set.

### SetIsShareableNil

`func (o *AiFolderDtoInteger) SetIsShareableNil(b bool)`

 SetIsShareableNil sets the value for IsShareable to be an explicit nil

### UnsetIsShareable
`func (o *AiFolderDtoInteger) UnsetIsShareable()`

UnsetIsShareable ensures that no value is present for IsShareable, not even an explicit nil
### GetNew

`func (o *AiFolderDtoInteger) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *AiFolderDtoInteger) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *AiFolderDtoInteger) SetNew(v int32)`

SetNew sets New field to given value.

### HasNew

`func (o *AiFolderDtoInteger) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetMute

`func (o *AiFolderDtoInteger) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *AiFolderDtoInteger) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *AiFolderDtoInteger) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *AiFolderDtoInteger) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetTags

`func (o *AiFolderDtoInteger) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AiFolderDtoInteger) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AiFolderDtoInteger) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AiFolderDtoInteger) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AiFolderDtoInteger) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AiFolderDtoInteger) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetLogo

`func (o *AiFolderDtoInteger) GetLogo() AiLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AiFolderDtoInteger) GetLogoOk() (*AiLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AiFolderDtoInteger) SetLogo(v AiLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AiFolderDtoInteger) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPinned

`func (o *AiFolderDtoInteger) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *AiFolderDtoInteger) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *AiFolderDtoInteger) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *AiFolderDtoInteger) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetRoomType

`func (o *AiFolderDtoInteger) GetRoomType() AiRoomType`

GetRoomType returns the RoomType field if non-nil, zero value otherwise.

### GetRoomTypeOk

`func (o *AiFolderDtoInteger) GetRoomTypeOk() (*AiRoomType, bool)`

GetRoomTypeOk returns a tuple with the RoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomType

`func (o *AiFolderDtoInteger) SetRoomType(v AiRoomType)`

SetRoomType sets RoomType field to given value.

### HasRoomType

`func (o *AiFolderDtoInteger) HasRoomType() bool`

HasRoomType returns a boolean if a field has been set.

### GetPrivate

`func (o *AiFolderDtoInteger) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *AiFolderDtoInteger) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *AiFolderDtoInteger) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *AiFolderDtoInteger) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetIndexing

`func (o *AiFolderDtoInteger) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *AiFolderDtoInteger) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *AiFolderDtoInteger) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *AiFolderDtoInteger) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### GetDenyDownload

`func (o *AiFolderDtoInteger) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *AiFolderDtoInteger) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *AiFolderDtoInteger) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *AiFolderDtoInteger) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetLifetime

`func (o *AiFolderDtoInteger) GetLifetime() AiRoomDataLifetimeDto`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *AiFolderDtoInteger) GetLifetimeOk() (*AiRoomDataLifetimeDto, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *AiFolderDtoInteger) SetLifetime(v AiRoomDataLifetimeDto)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *AiFolderDtoInteger) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetWatermark

`func (o *AiFolderDtoInteger) GetWatermark() AiWatermarkDto`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *AiFolderDtoInteger) GetWatermarkOk() (*AiWatermarkDto, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *AiFolderDtoInteger) SetWatermark(v AiWatermarkDto)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *AiFolderDtoInteger) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetType

`func (o *AiFolderDtoInteger) GetType() AiFolderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiFolderDtoInteger) GetTypeOk() (*AiFolderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiFolderDtoInteger) SetType(v AiFolderType)`

SetType sets Type field to given value.

### HasType

`func (o *AiFolderDtoInteger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInRoom

`func (o *AiFolderDtoInteger) GetInRoom() bool`

GetInRoom returns the InRoom field if non-nil, zero value otherwise.

### GetInRoomOk

`func (o *AiFolderDtoInteger) GetInRoomOk() (*bool, bool)`

GetInRoomOk returns a tuple with the InRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInRoom

`func (o *AiFolderDtoInteger) SetInRoom(v bool)`

SetInRoom sets InRoom field to given value.

### HasInRoom

`func (o *AiFolderDtoInteger) HasInRoom() bool`

HasInRoom returns a boolean if a field has been set.

### SetInRoomNil

`func (o *AiFolderDtoInteger) SetInRoomNil(b bool)`

 SetInRoomNil sets the value for InRoom to be an explicit nil

### UnsetInRoom
`func (o *AiFolderDtoInteger) UnsetInRoom()`

UnsetInRoom ensures that no value is present for InRoom, not even an explicit nil
### GetQuotaLimit

`func (o *AiFolderDtoInteger) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *AiFolderDtoInteger) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *AiFolderDtoInteger) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *AiFolderDtoInteger) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### SetQuotaLimitNil

`func (o *AiFolderDtoInteger) SetQuotaLimitNil(b bool)`

 SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil

### UnsetQuotaLimit
`func (o *AiFolderDtoInteger) UnsetQuotaLimit()`

UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
### GetIsCustomQuota

`func (o *AiFolderDtoInteger) GetIsCustomQuota() bool`

GetIsCustomQuota returns the IsCustomQuota field if non-nil, zero value otherwise.

### GetIsCustomQuotaOk

`func (o *AiFolderDtoInteger) GetIsCustomQuotaOk() (*bool, bool)`

GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomQuota

`func (o *AiFolderDtoInteger) SetIsCustomQuota(v bool)`

SetIsCustomQuota sets IsCustomQuota field to given value.

### HasIsCustomQuota

`func (o *AiFolderDtoInteger) HasIsCustomQuota() bool`

HasIsCustomQuota returns a boolean if a field has been set.

### SetIsCustomQuotaNil

`func (o *AiFolderDtoInteger) SetIsCustomQuotaNil(b bool)`

 SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil

### UnsetIsCustomQuota
`func (o *AiFolderDtoInteger) UnsetIsCustomQuota()`

UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
### GetUsedSpace

`func (o *AiFolderDtoInteger) GetUsedSpace() int64`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *AiFolderDtoInteger) GetUsedSpaceOk() (*int64, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *AiFolderDtoInteger) SetUsedSpace(v int64)`

SetUsedSpace sets UsedSpace field to given value.

### HasUsedSpace

`func (o *AiFolderDtoInteger) HasUsedSpace() bool`

HasUsedSpace returns a boolean if a field has been set.

### SetUsedSpaceNil

`func (o *AiFolderDtoInteger) SetUsedSpaceNil(b bool)`

 SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil

### UnsetUsedSpace
`func (o *AiFolderDtoInteger) UnsetUsedSpace()`

UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
### GetPasswordProtected

`func (o *AiFolderDtoInteger) GetPasswordProtected() bool`

GetPasswordProtected returns the PasswordProtected field if non-nil, zero value otherwise.

### GetPasswordProtectedOk

`func (o *AiFolderDtoInteger) GetPasswordProtectedOk() (*bool, bool)`

GetPasswordProtectedOk returns a tuple with the PasswordProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProtected

`func (o *AiFolderDtoInteger) SetPasswordProtected(v bool)`

SetPasswordProtected sets PasswordProtected field to given value.

### HasPasswordProtected

`func (o *AiFolderDtoInteger) HasPasswordProtected() bool`

HasPasswordProtected returns a boolean if a field has been set.

### SetPasswordProtectedNil

`func (o *AiFolderDtoInteger) SetPasswordProtectedNil(b bool)`

 SetPasswordProtectedNil sets the value for PasswordProtected to be an explicit nil

### UnsetPasswordProtected
`func (o *AiFolderDtoInteger) UnsetPasswordProtected()`

UnsetPasswordProtected ensures that no value is present for PasswordProtected, not even an explicit nil
### GetExpired

`func (o *AiFolderDtoInteger) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *AiFolderDtoInteger) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *AiFolderDtoInteger) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *AiFolderDtoInteger) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### SetExpiredNil

`func (o *AiFolderDtoInteger) SetExpiredNil(b bool)`

 SetExpiredNil sets the value for Expired to be an explicit nil

### UnsetExpired
`func (o *AiFolderDtoInteger) UnsetExpired()`

UnsetExpired ensures that no value is present for Expired, not even an explicit nil
### GetChatSettings

`func (o *AiFolderDtoInteger) GetChatSettings() AiChatSettingsDto`

GetChatSettings returns the ChatSettings field if non-nil, zero value otherwise.

### GetChatSettingsOk

`func (o *AiFolderDtoInteger) GetChatSettingsOk() (*AiChatSettingsDto, bool)`

GetChatSettingsOk returns a tuple with the ChatSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSettings

`func (o *AiFolderDtoInteger) SetChatSettings(v AiChatSettingsDto)`

SetChatSettings sets ChatSettings field to given value.

### HasChatSettings

`func (o *AiFolderDtoInteger) HasChatSettings() bool`

HasChatSettings returns a boolean if a field has been set.

### GetRootRoomType

`func (o *AiFolderDtoInteger) GetRootRoomType() AiRoomType`

GetRootRoomType returns the RootRoomType field if non-nil, zero value otherwise.

### GetRootRoomTypeOk

`func (o *AiFolderDtoInteger) GetRootRoomTypeOk() (*AiRoomType, bool)`

GetRootRoomTypeOk returns a tuple with the RootRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRoomType

`func (o *AiFolderDtoInteger) SetRootRoomType(v AiRoomType)`

SetRootRoomType sets RootRoomType field to given value.

### HasRootRoomType

`func (o *AiFolderDtoInteger) HasRootRoomType() bool`

HasRootRoomType returns a boolean if a field has been set.

### GetSaveFormAsXLSX

`func (o *AiFolderDtoInteger) GetSaveFormAsXLSX() bool`

GetSaveFormAsXLSX returns the SaveFormAsXLSX field if non-nil, zero value otherwise.

### GetSaveFormAsXLSXOk

`func (o *AiFolderDtoInteger) GetSaveFormAsXLSXOk() (*bool, bool)`

GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveFormAsXLSX

`func (o *AiFolderDtoInteger) SetSaveFormAsXLSX(v bool)`

SetSaveFormAsXLSX sets SaveFormAsXLSX field to given value.

### HasSaveFormAsXLSX

`func (o *AiFolderDtoInteger) HasSaveFormAsXLSX() bool`

HasSaveFormAsXLSX returns a boolean if a field has been set.

### SetSaveFormAsXLSXNil

`func (o *AiFolderDtoInteger) SetSaveFormAsXLSXNil(b bool)`

 SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil

### UnsetSaveFormAsXLSX
`func (o *AiFolderDtoInteger) UnsetSaveFormAsXLSX()`

UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil
### GetSendFormToExternalDB

`func (o *AiFolderDtoInteger) GetSendFormToExternalDB() bool`

GetSendFormToExternalDB returns the SendFormToExternalDB field if non-nil, zero value otherwise.

### GetSendFormToExternalDBOk

`func (o *AiFolderDtoInteger) GetSendFormToExternalDBOk() (*bool, bool)`

GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFormToExternalDB

`func (o *AiFolderDtoInteger) SetSendFormToExternalDB(v bool)`

SetSendFormToExternalDB sets SendFormToExternalDB field to given value.

### HasSendFormToExternalDB

`func (o *AiFolderDtoInteger) HasSendFormToExternalDB() bool`

HasSendFormToExternalDB returns a boolean if a field has been set.

### SetSendFormToExternalDBNil

`func (o *AiFolderDtoInteger) SetSendFormToExternalDBNil(b bool)`

 SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil

### UnsetSendFormToExternalDB
`func (o *AiFolderDtoInteger) UnsetSendFormToExternalDB()`

UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
### GetOriginalFormId

`func (o *AiFolderDtoInteger) GetOriginalFormId() int32`

GetOriginalFormId returns the OriginalFormId field if non-nil, zero value otherwise.

### GetOriginalFormIdOk

`func (o *AiFolderDtoInteger) GetOriginalFormIdOk() (*int32, bool)`

GetOriginalFormIdOk returns a tuple with the OriginalFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFormId

`func (o *AiFolderDtoInteger) SetOriginalFormId(v int32)`

SetOriginalFormId sets OriginalFormId field to given value.

### HasOriginalFormId

`func (o *AiFolderDtoInteger) HasOriginalFormId() bool`

HasOriginalFormId returns a boolean if a field has been set.

### SetOriginalFormIdNil

`func (o *AiFolderDtoInteger) SetOriginalFormIdNil(b bool)`

 SetOriginalFormIdNil sets the value for OriginalFormId to be an explicit nil

### UnsetOriginalFormId
`func (o *AiFolderDtoInteger) UnsetOriginalFormId()`

UnsetOriginalFormId ensures that no value is present for OriginalFormId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


