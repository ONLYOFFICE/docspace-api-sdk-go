# FileEntryDtoInteger

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
**ExpirationDate** | Pointer to **NullableTime** | Represents the expiration date of the file entry. | [optional] 
**IsLinkExpired** | Pointer to **NullableBool** | Indicates whether the shareable link associated with the file or folder has expired. | [optional] 

## Methods

### NewFileEntryDtoInteger

`func NewFileEntryDtoInteger() *FileEntryDtoInteger`

NewFileEntryDtoInteger instantiates a new FileEntryDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileEntryDtoIntegerWithDefaults

`func NewFileEntryDtoIntegerWithDefaults() *FileEntryDtoInteger`

NewFileEntryDtoIntegerWithDefaults instantiates a new FileEntryDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FileEntryDtoInteger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FileEntryDtoInteger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FileEntryDtoInteger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FileEntryDtoInteger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAccess

`func (o *FileEntryDtoInteger) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FileEntryDtoInteger) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FileEntryDtoInteger) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FileEntryDtoInteger) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *FileEntryDtoInteger) GetSharedBy() EmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *FileEntryDtoInteger) GetSharedByOk() (*EmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *FileEntryDtoInteger) SetSharedBy(v EmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *FileEntryDtoInteger) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *FileEntryDtoInteger) GetOwnedBy() EmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *FileEntryDtoInteger) GetOwnedByOk() (*EmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *FileEntryDtoInteger) SetOwnedBy(v EmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *FileEntryDtoInteger) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *FileEntryDtoInteger) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *FileEntryDtoInteger) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *FileEntryDtoInteger) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *FileEntryDtoInteger) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *FileEntryDtoInteger) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *FileEntryDtoInteger) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *FileEntryDtoInteger) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *FileEntryDtoInteger) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *FileEntryDtoInteger) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *FileEntryDtoInteger) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *FileEntryDtoInteger) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *FileEntryDtoInteger) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *FileEntryDtoInteger) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *FileEntryDtoInteger) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *FileEntryDtoInteger) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *FileEntryDtoInteger) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *FileEntryDtoInteger) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *FileEntryDtoInteger) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *FileEntryDtoInteger) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *FileEntryDtoInteger) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### GetCreated

`func (o *FileEntryDtoInteger) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FileEntryDtoInteger) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FileEntryDtoInteger) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FileEntryDtoInteger) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FileEntryDtoInteger) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FileEntryDtoInteger) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FileEntryDtoInteger) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FileEntryDtoInteger) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *FileEntryDtoInteger) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FileEntryDtoInteger) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FileEntryDtoInteger) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FileEntryDtoInteger) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAutoDelete

`func (o *FileEntryDtoInteger) GetAutoDelete() time.Time`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *FileEntryDtoInteger) GetAutoDeleteOk() (*time.Time, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *FileEntryDtoInteger) SetAutoDelete(v time.Time)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *FileEntryDtoInteger) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetRootFolderType

`func (o *FileEntryDtoInteger) GetRootFolderType() FolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *FileEntryDtoInteger) GetRootFolderTypeOk() (*FolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *FileEntryDtoInteger) SetRootFolderType(v FolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *FileEntryDtoInteger) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *FileEntryDtoInteger) GetParentRoomType() FolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *FileEntryDtoInteger) GetParentRoomTypeOk() (*FolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *FileEntryDtoInteger) SetParentRoomType(v FolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *FileEntryDtoInteger) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FileEntryDtoInteger) GetUpdatedBy() EmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FileEntryDtoInteger) GetUpdatedByOk() (*EmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FileEntryDtoInteger) SetUpdatedBy(v EmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FileEntryDtoInteger) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *FileEntryDtoInteger) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *FileEntryDtoInteger) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *FileEntryDtoInteger) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *FileEntryDtoInteger) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### GetProviderKey

`func (o *FileEntryDtoInteger) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *FileEntryDtoInteger) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *FileEntryDtoInteger) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *FileEntryDtoInteger) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### GetProviderId

`func (o *FileEntryDtoInteger) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *FileEntryDtoInteger) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *FileEntryDtoInteger) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *FileEntryDtoInteger) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetOrder

`func (o *FileEntryDtoInteger) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FileEntryDtoInteger) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FileEntryDtoInteger) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FileEntryDtoInteger) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIsFavorite

`func (o *FileEntryDtoInteger) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *FileEntryDtoInteger) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *FileEntryDtoInteger) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *FileEntryDtoInteger) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetFileEntryType

`func (o *FileEntryDtoInteger) GetFileEntryType() FileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *FileEntryDtoInteger) GetFileEntryTypeOk() (*FileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *FileEntryDtoInteger) SetFileEntryType(v FileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *FileEntryDtoInteger) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.

### GetId

`func (o *FileEntryDtoInteger) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileEntryDtoInteger) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileEntryDtoInteger) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FileEntryDtoInteger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootFolderId

`func (o *FileEntryDtoInteger) GetRootFolderId() int32`

GetRootFolderId returns the RootFolderId field if non-nil, zero value otherwise.

### GetRootFolderIdOk

`func (o *FileEntryDtoInteger) GetRootFolderIdOk() (*int32, bool)`

GetRootFolderIdOk returns a tuple with the RootFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderId

`func (o *FileEntryDtoInteger) SetRootFolderId(v int32)`

SetRootFolderId sets RootFolderId field to given value.

### HasRootFolderId

`func (o *FileEntryDtoInteger) HasRootFolderId() bool`

HasRootFolderId returns a boolean if a field has been set.

### GetOriginId

`func (o *FileEntryDtoInteger) GetOriginId() int32`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *FileEntryDtoInteger) GetOriginIdOk() (*int32, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *FileEntryDtoInteger) SetOriginId(v int32)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *FileEntryDtoInteger) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginRoomId

`func (o *FileEntryDtoInteger) GetOriginRoomId() int32`

GetOriginRoomId returns the OriginRoomId field if non-nil, zero value otherwise.

### GetOriginRoomIdOk

`func (o *FileEntryDtoInteger) GetOriginRoomIdOk() (*int32, bool)`

GetOriginRoomIdOk returns a tuple with the OriginRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomId

`func (o *FileEntryDtoInteger) SetOriginRoomId(v int32)`

SetOriginRoomId sets OriginRoomId field to given value.

### HasOriginRoomId

`func (o *FileEntryDtoInteger) HasOriginRoomId() bool`

HasOriginRoomId returns a boolean if a field has been set.

### GetOriginTitle

`func (o *FileEntryDtoInteger) GetOriginTitle() string`

GetOriginTitle returns the OriginTitle field if non-nil, zero value otherwise.

### GetOriginTitleOk

`func (o *FileEntryDtoInteger) GetOriginTitleOk() (*string, bool)`

GetOriginTitleOk returns a tuple with the OriginTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTitle

`func (o *FileEntryDtoInteger) SetOriginTitle(v string)`

SetOriginTitle sets OriginTitle field to given value.

### HasOriginTitle

`func (o *FileEntryDtoInteger) HasOriginTitle() bool`

HasOriginTitle returns a boolean if a field has been set.

### SetOriginTitleNil

`func (o *FileEntryDtoInteger) SetOriginTitleNil(b bool)`

 SetOriginTitleNil sets the value for OriginTitle to be an explicit nil

### UnsetOriginTitle
`func (o *FileEntryDtoInteger) UnsetOriginTitle()`

UnsetOriginTitle ensures that no value is present for OriginTitle, not even an explicit nil
### GetOriginRoomTitle

`func (o *FileEntryDtoInteger) GetOriginRoomTitle() string`

GetOriginRoomTitle returns the OriginRoomTitle field if non-nil, zero value otherwise.

### GetOriginRoomTitleOk

`func (o *FileEntryDtoInteger) GetOriginRoomTitleOk() (*string, bool)`

GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomTitle

`func (o *FileEntryDtoInteger) SetOriginRoomTitle(v string)`

SetOriginRoomTitle sets OriginRoomTitle field to given value.

### HasOriginRoomTitle

`func (o *FileEntryDtoInteger) HasOriginRoomTitle() bool`

HasOriginRoomTitle returns a boolean if a field has been set.

### SetOriginRoomTitleNil

`func (o *FileEntryDtoInteger) SetOriginRoomTitleNil(b bool)`

 SetOriginRoomTitleNil sets the value for OriginRoomTitle to be an explicit nil

### UnsetOriginRoomTitle
`func (o *FileEntryDtoInteger) UnsetOriginRoomTitle()`

UnsetOriginRoomTitle ensures that no value is present for OriginRoomTitle, not even an explicit nil
### GetCanShare

`func (o *FileEntryDtoInteger) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *FileEntryDtoInteger) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *FileEntryDtoInteger) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *FileEntryDtoInteger) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### GetShareSettings

`func (o *FileEntryDtoInteger) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings`

GetShareSettings returns the ShareSettings field if non-nil, zero value otherwise.

### GetShareSettingsOk

`func (o *FileEntryDtoInteger) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool)`

GetShareSettingsOk returns a tuple with the ShareSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareSettings

`func (o *FileEntryDtoInteger) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings)`

SetShareSettings sets ShareSettings field to given value.

### HasShareSettings

`func (o *FileEntryDtoInteger) HasShareSettings() bool`

HasShareSettings returns a boolean if a field has been set.

### SetShareSettingsNil

`func (o *FileEntryDtoInteger) SetShareSettingsNil(b bool)`

 SetShareSettingsNil sets the value for ShareSettings to be an explicit nil

### UnsetShareSettings
`func (o *FileEntryDtoInteger) UnsetShareSettings()`

UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
### GetSecurity

`func (o *FileEntryDtoInteger) GetSecurity() FileEntryDtoIntegerAllOfSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *FileEntryDtoInteger) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *FileEntryDtoInteger) SetSecurity(v FileEntryDtoIntegerAllOfSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *FileEntryDtoInteger) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *FileEntryDtoInteger) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *FileEntryDtoInteger) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetAvailableShareRights

`func (o *FileEntryDtoInteger) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights`

GetAvailableShareRights returns the AvailableShareRights field if non-nil, zero value otherwise.

### GetAvailableShareRightsOk

`func (o *FileEntryDtoInteger) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool)`

GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShareRights

`func (o *FileEntryDtoInteger) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights)`

SetAvailableShareRights sets AvailableShareRights field to given value.

### HasAvailableShareRights

`func (o *FileEntryDtoInteger) HasAvailableShareRights() bool`

HasAvailableShareRights returns a boolean if a field has been set.

### SetAvailableShareRightsNil

`func (o *FileEntryDtoInteger) SetAvailableShareRightsNil(b bool)`

 SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil

### UnsetAvailableShareRights
`func (o *FileEntryDtoInteger) UnsetAvailableShareRights()`

UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
### GetRequestToken

`func (o *FileEntryDtoInteger) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *FileEntryDtoInteger) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *FileEntryDtoInteger) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *FileEntryDtoInteger) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### SetRequestTokenNil

`func (o *FileEntryDtoInteger) SetRequestTokenNil(b bool)`

 SetRequestTokenNil sets the value for RequestToken to be an explicit nil

### UnsetRequestToken
`func (o *FileEntryDtoInteger) UnsetRequestToken()`

UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
### GetExternal

`func (o *FileEntryDtoInteger) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *FileEntryDtoInteger) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *FileEntryDtoInteger) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *FileEntryDtoInteger) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### SetExternalNil

`func (o *FileEntryDtoInteger) SetExternalNil(b bool)`

 SetExternalNil sets the value for External to be an explicit nil

### UnsetExternal
`func (o *FileEntryDtoInteger) UnsetExternal()`

UnsetExternal ensures that no value is present for External, not even an explicit nil
### GetExpirationDate

`func (o *FileEntryDtoInteger) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FileEntryDtoInteger) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FileEntryDtoInteger) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FileEntryDtoInteger) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *FileEntryDtoInteger) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *FileEntryDtoInteger) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetIsLinkExpired

`func (o *FileEntryDtoInteger) GetIsLinkExpired() bool`

GetIsLinkExpired returns the IsLinkExpired field if non-nil, zero value otherwise.

### GetIsLinkExpiredOk

`func (o *FileEntryDtoInteger) GetIsLinkExpiredOk() (*bool, bool)`

GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkExpired

`func (o *FileEntryDtoInteger) SetIsLinkExpired(v bool)`

SetIsLinkExpired sets IsLinkExpired field to given value.

### HasIsLinkExpired

`func (o *FileEntryDtoInteger) HasIsLinkExpired() bool`

HasIsLinkExpired returns a boolean if a field has been set.

### SetIsLinkExpiredNil

`func (o *FileEntryDtoInteger) SetIsLinkExpiredNil(b bool)`

 SetIsLinkExpiredNil sets the value for IsLinkExpired to be an explicit nil

### UnsetIsLinkExpired
`func (o *FileEntryDtoInteger) UnsetIsLinkExpired()`

UnsetIsLinkExpired ensures that no value is present for IsLinkExpired, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


