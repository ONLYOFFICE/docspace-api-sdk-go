# AiFileEntryDtoInteger

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

### NewAiFileEntryDtoInteger

`func NewAiFileEntryDtoInteger() *AiFileEntryDtoInteger`

NewAiFileEntryDtoInteger instantiates a new AiFileEntryDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiFileEntryDtoIntegerWithDefaults

`func NewAiFileEntryDtoIntegerWithDefaults() *AiFileEntryDtoInteger`

NewAiFileEntryDtoIntegerWithDefaults instantiates a new AiFileEntryDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AiFileEntryDtoInteger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiFileEntryDtoInteger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiFileEntryDtoInteger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AiFileEntryDtoInteger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAccess

`func (o *AiFileEntryDtoInteger) GetAccess() AiFileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AiFileEntryDtoInteger) GetAccessOk() (*AiFileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AiFileEntryDtoInteger) SetAccess(v AiFileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AiFileEntryDtoInteger) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *AiFileEntryDtoInteger) GetSharedBy() AiEmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *AiFileEntryDtoInteger) GetSharedByOk() (*AiEmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *AiFileEntryDtoInteger) SetSharedBy(v AiEmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *AiFileEntryDtoInteger) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *AiFileEntryDtoInteger) GetOwnedBy() AiEmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *AiFileEntryDtoInteger) GetOwnedByOk() (*AiEmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *AiFileEntryDtoInteger) SetOwnedBy(v AiEmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *AiFileEntryDtoInteger) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *AiFileEntryDtoInteger) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *AiFileEntryDtoInteger) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *AiFileEntryDtoInteger) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *AiFileEntryDtoInteger) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *AiFileEntryDtoInteger) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *AiFileEntryDtoInteger) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *AiFileEntryDtoInteger) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *AiFileEntryDtoInteger) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *AiFileEntryDtoInteger) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *AiFileEntryDtoInteger) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *AiFileEntryDtoInteger) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *AiFileEntryDtoInteger) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *AiFileEntryDtoInteger) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *AiFileEntryDtoInteger) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *AiFileEntryDtoInteger) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *AiFileEntryDtoInteger) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *AiFileEntryDtoInteger) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *AiFileEntryDtoInteger) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *AiFileEntryDtoInteger) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *AiFileEntryDtoInteger) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### GetCreated

`func (o *AiFileEntryDtoInteger) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiFileEntryDtoInteger) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiFileEntryDtoInteger) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AiFileEntryDtoInteger) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AiFileEntryDtoInteger) GetCreatedBy() AiEmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AiFileEntryDtoInteger) GetCreatedByOk() (*AiEmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AiFileEntryDtoInteger) SetCreatedBy(v AiEmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AiFileEntryDtoInteger) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *AiFileEntryDtoInteger) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AiFileEntryDtoInteger) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AiFileEntryDtoInteger) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AiFileEntryDtoInteger) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAutoDelete

`func (o *AiFileEntryDtoInteger) GetAutoDelete() time.Time`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *AiFileEntryDtoInteger) GetAutoDeleteOk() (*time.Time, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *AiFileEntryDtoInteger) SetAutoDelete(v time.Time)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *AiFileEntryDtoInteger) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetRootFolderType

`func (o *AiFileEntryDtoInteger) GetRootFolderType() AiFolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *AiFileEntryDtoInteger) GetRootFolderTypeOk() (*AiFolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *AiFileEntryDtoInteger) SetRootFolderType(v AiFolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *AiFileEntryDtoInteger) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *AiFileEntryDtoInteger) GetParentRoomType() AiFolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *AiFileEntryDtoInteger) GetParentRoomTypeOk() (*AiFolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *AiFileEntryDtoInteger) SetParentRoomType(v AiFolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *AiFileEntryDtoInteger) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AiFileEntryDtoInteger) GetUpdatedBy() AiEmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AiFileEntryDtoInteger) GetUpdatedByOk() (*AiEmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AiFileEntryDtoInteger) SetUpdatedBy(v AiEmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AiFileEntryDtoInteger) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *AiFileEntryDtoInteger) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *AiFileEntryDtoInteger) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *AiFileEntryDtoInteger) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *AiFileEntryDtoInteger) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### GetProviderKey

`func (o *AiFileEntryDtoInteger) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *AiFileEntryDtoInteger) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *AiFileEntryDtoInteger) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *AiFileEntryDtoInteger) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### GetProviderId

`func (o *AiFileEntryDtoInteger) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *AiFileEntryDtoInteger) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *AiFileEntryDtoInteger) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *AiFileEntryDtoInteger) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetOrder

`func (o *AiFileEntryDtoInteger) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AiFileEntryDtoInteger) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AiFileEntryDtoInteger) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AiFileEntryDtoInteger) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIsFavorite

`func (o *AiFileEntryDtoInteger) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *AiFileEntryDtoInteger) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *AiFileEntryDtoInteger) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *AiFileEntryDtoInteger) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetFileEntryType

`func (o *AiFileEntryDtoInteger) GetFileEntryType() AiFileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *AiFileEntryDtoInteger) GetFileEntryTypeOk() (*AiFileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *AiFileEntryDtoInteger) SetFileEntryType(v AiFileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *AiFileEntryDtoInteger) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.

### GetId

`func (o *AiFileEntryDtoInteger) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiFileEntryDtoInteger) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiFileEntryDtoInteger) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AiFileEntryDtoInteger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootFolderId

`func (o *AiFileEntryDtoInteger) GetRootFolderId() int32`

GetRootFolderId returns the RootFolderId field if non-nil, zero value otherwise.

### GetRootFolderIdOk

`func (o *AiFileEntryDtoInteger) GetRootFolderIdOk() (*int32, bool)`

GetRootFolderIdOk returns a tuple with the RootFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderId

`func (o *AiFileEntryDtoInteger) SetRootFolderId(v int32)`

SetRootFolderId sets RootFolderId field to given value.

### HasRootFolderId

`func (o *AiFileEntryDtoInteger) HasRootFolderId() bool`

HasRootFolderId returns a boolean if a field has been set.

### GetOriginId

`func (o *AiFileEntryDtoInteger) GetOriginId() int32`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *AiFileEntryDtoInteger) GetOriginIdOk() (*int32, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *AiFileEntryDtoInteger) SetOriginId(v int32)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *AiFileEntryDtoInteger) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginRoomId

`func (o *AiFileEntryDtoInteger) GetOriginRoomId() int32`

GetOriginRoomId returns the OriginRoomId field if non-nil, zero value otherwise.

### GetOriginRoomIdOk

`func (o *AiFileEntryDtoInteger) GetOriginRoomIdOk() (*int32, bool)`

GetOriginRoomIdOk returns a tuple with the OriginRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomId

`func (o *AiFileEntryDtoInteger) SetOriginRoomId(v int32)`

SetOriginRoomId sets OriginRoomId field to given value.

### HasOriginRoomId

`func (o *AiFileEntryDtoInteger) HasOriginRoomId() bool`

HasOriginRoomId returns a boolean if a field has been set.

### GetOriginTitle

`func (o *AiFileEntryDtoInteger) GetOriginTitle() string`

GetOriginTitle returns the OriginTitle field if non-nil, zero value otherwise.

### GetOriginTitleOk

`func (o *AiFileEntryDtoInteger) GetOriginTitleOk() (*string, bool)`

GetOriginTitleOk returns a tuple with the OriginTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTitle

`func (o *AiFileEntryDtoInteger) SetOriginTitle(v string)`

SetOriginTitle sets OriginTitle field to given value.

### HasOriginTitle

`func (o *AiFileEntryDtoInteger) HasOriginTitle() bool`

HasOriginTitle returns a boolean if a field has been set.

### SetOriginTitleNil

`func (o *AiFileEntryDtoInteger) SetOriginTitleNil(b bool)`

 SetOriginTitleNil sets the value for OriginTitle to be an explicit nil

### UnsetOriginTitle
`func (o *AiFileEntryDtoInteger) UnsetOriginTitle()`

UnsetOriginTitle ensures that no value is present for OriginTitle, not even an explicit nil
### GetOriginRoomTitle

`func (o *AiFileEntryDtoInteger) GetOriginRoomTitle() string`

GetOriginRoomTitle returns the OriginRoomTitle field if non-nil, zero value otherwise.

### GetOriginRoomTitleOk

`func (o *AiFileEntryDtoInteger) GetOriginRoomTitleOk() (*string, bool)`

GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomTitle

`func (o *AiFileEntryDtoInteger) SetOriginRoomTitle(v string)`

SetOriginRoomTitle sets OriginRoomTitle field to given value.

### HasOriginRoomTitle

`func (o *AiFileEntryDtoInteger) HasOriginRoomTitle() bool`

HasOriginRoomTitle returns a boolean if a field has been set.

### SetOriginRoomTitleNil

`func (o *AiFileEntryDtoInteger) SetOriginRoomTitleNil(b bool)`

 SetOriginRoomTitleNil sets the value for OriginRoomTitle to be an explicit nil

### UnsetOriginRoomTitle
`func (o *AiFileEntryDtoInteger) UnsetOriginRoomTitle()`

UnsetOriginRoomTitle ensures that no value is present for OriginRoomTitle, not even an explicit nil
### GetCanShare

`func (o *AiFileEntryDtoInteger) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *AiFileEntryDtoInteger) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *AiFileEntryDtoInteger) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *AiFileEntryDtoInteger) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### GetShareSettings

`func (o *AiFileEntryDtoInteger) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings`

GetShareSettings returns the ShareSettings field if non-nil, zero value otherwise.

### GetShareSettingsOk

`func (o *AiFileEntryDtoInteger) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool)`

GetShareSettingsOk returns a tuple with the ShareSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareSettings

`func (o *AiFileEntryDtoInteger) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings)`

SetShareSettings sets ShareSettings field to given value.

### HasShareSettings

`func (o *AiFileEntryDtoInteger) HasShareSettings() bool`

HasShareSettings returns a boolean if a field has been set.

### SetShareSettingsNil

`func (o *AiFileEntryDtoInteger) SetShareSettingsNil(b bool)`

 SetShareSettingsNil sets the value for ShareSettings to be an explicit nil

### UnsetShareSettings
`func (o *AiFileEntryDtoInteger) UnsetShareSettings()`

UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
### GetSecurity

`func (o *AiFileEntryDtoInteger) GetSecurity() FileEntryDtoIntegerAllOfSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *AiFileEntryDtoInteger) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *AiFileEntryDtoInteger) SetSecurity(v FileEntryDtoIntegerAllOfSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *AiFileEntryDtoInteger) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *AiFileEntryDtoInteger) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *AiFileEntryDtoInteger) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetAvailableShareRights

`func (o *AiFileEntryDtoInteger) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights`

GetAvailableShareRights returns the AvailableShareRights field if non-nil, zero value otherwise.

### GetAvailableShareRightsOk

`func (o *AiFileEntryDtoInteger) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool)`

GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShareRights

`func (o *AiFileEntryDtoInteger) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights)`

SetAvailableShareRights sets AvailableShareRights field to given value.

### HasAvailableShareRights

`func (o *AiFileEntryDtoInteger) HasAvailableShareRights() bool`

HasAvailableShareRights returns a boolean if a field has been set.

### SetAvailableShareRightsNil

`func (o *AiFileEntryDtoInteger) SetAvailableShareRightsNil(b bool)`

 SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil

### UnsetAvailableShareRights
`func (o *AiFileEntryDtoInteger) UnsetAvailableShareRights()`

UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
### GetRequestToken

`func (o *AiFileEntryDtoInteger) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *AiFileEntryDtoInteger) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *AiFileEntryDtoInteger) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *AiFileEntryDtoInteger) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### SetRequestTokenNil

`func (o *AiFileEntryDtoInteger) SetRequestTokenNil(b bool)`

 SetRequestTokenNil sets the value for RequestToken to be an explicit nil

### UnsetRequestToken
`func (o *AiFileEntryDtoInteger) UnsetRequestToken()`

UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
### GetExternal

`func (o *AiFileEntryDtoInteger) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *AiFileEntryDtoInteger) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *AiFileEntryDtoInteger) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *AiFileEntryDtoInteger) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### SetExternalNil

`func (o *AiFileEntryDtoInteger) SetExternalNil(b bool)`

 SetExternalNil sets the value for External to be an explicit nil

### UnsetExternal
`func (o *AiFileEntryDtoInteger) UnsetExternal()`

UnsetExternal ensures that no value is present for External, not even an explicit nil
### GetExpirationDate

`func (o *AiFileEntryDtoInteger) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AiFileEntryDtoInteger) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AiFileEntryDtoInteger) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AiFileEntryDtoInteger) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *AiFileEntryDtoInteger) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *AiFileEntryDtoInteger) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetIsLinkExpired

`func (o *AiFileEntryDtoInteger) GetIsLinkExpired() bool`

GetIsLinkExpired returns the IsLinkExpired field if non-nil, zero value otherwise.

### GetIsLinkExpiredOk

`func (o *AiFileEntryDtoInteger) GetIsLinkExpiredOk() (*bool, bool)`

GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkExpired

`func (o *AiFileEntryDtoInteger) SetIsLinkExpired(v bool)`

SetIsLinkExpired sets IsLinkExpired field to given value.

### HasIsLinkExpired

`func (o *AiFileEntryDtoInteger) HasIsLinkExpired() bool`

HasIsLinkExpired returns a boolean if a field has been set.

### SetIsLinkExpiredNil

`func (o *AiFileEntryDtoInteger) SetIsLinkExpiredNil(b bool)`

 SetIsLinkExpiredNil sets the value for IsLinkExpired to be an explicit nil

### UnsetIsLinkExpired
`func (o *AiFileEntryDtoInteger) UnsetIsLinkExpired()`

UnsetIsLinkExpired ensures that no value is present for IsLinkExpired, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


