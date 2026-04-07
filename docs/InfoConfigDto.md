# InfoConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Favorite** | Pointer to **NullableBool** | Specifies if the file is favorite or not. | [optional] 
**Folder** | Pointer to **NullableString** | The folder of the file. | [optional] 
**Owner** | Pointer to **NullableString** | The file owner. | [optional] 
**SharingSettings** | Pointer to [**[]AceShortWrapper**](AceShortWrapper.md) | The sharing settings of the file. | [optional] 
**Type** | Pointer to [**EditorType**](EditorType.md) |  | [optional] 
**Uploaded** | Pointer to **NullableString** | The uploaded file. | [optional] 

## Methods

### NewInfoConfigDto

`func NewInfoConfigDto() *InfoConfigDto`

NewInfoConfigDto instantiates a new InfoConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoConfigDtoWithDefaults

`func NewInfoConfigDtoWithDefaults() *InfoConfigDto`

NewInfoConfigDtoWithDefaults instantiates a new InfoConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavorite

`func (o *InfoConfigDto) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *InfoConfigDto) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *InfoConfigDto) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *InfoConfigDto) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### SetFavoriteNil

`func (o *InfoConfigDto) SetFavoriteNil(b bool)`

 SetFavoriteNil sets the value for Favorite to be an explicit nil

### UnsetFavorite
`func (o *InfoConfigDto) UnsetFavorite()`

UnsetFavorite ensures that no value is present for Favorite, not even an explicit nil
### GetFolder

`func (o *InfoConfigDto) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *InfoConfigDto) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *InfoConfigDto) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *InfoConfigDto) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *InfoConfigDto) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *InfoConfigDto) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetOwner

`func (o *InfoConfigDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InfoConfigDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InfoConfigDto) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InfoConfigDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *InfoConfigDto) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *InfoConfigDto) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSharingSettings

`func (o *InfoConfigDto) GetSharingSettings() []AceShortWrapper`

GetSharingSettings returns the SharingSettings field if non-nil, zero value otherwise.

### GetSharingSettingsOk

`func (o *InfoConfigDto) GetSharingSettingsOk() (*[]AceShortWrapper, bool)`

GetSharingSettingsOk returns a tuple with the SharingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingSettings

`func (o *InfoConfigDto) SetSharingSettings(v []AceShortWrapper)`

SetSharingSettings sets SharingSettings field to given value.

### HasSharingSettings

`func (o *InfoConfigDto) HasSharingSettings() bool`

HasSharingSettings returns a boolean if a field has been set.

### SetSharingSettingsNil

`func (o *InfoConfigDto) SetSharingSettingsNil(b bool)`

 SetSharingSettingsNil sets the value for SharingSettings to be an explicit nil

### UnsetSharingSettings
`func (o *InfoConfigDto) UnsetSharingSettings()`

UnsetSharingSettings ensures that no value is present for SharingSettings, not even an explicit nil
### GetType

`func (o *InfoConfigDto) GetType() EditorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InfoConfigDto) GetTypeOk() (*EditorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InfoConfigDto) SetType(v EditorType)`

SetType sets Type field to given value.

### HasType

`func (o *InfoConfigDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUploaded

`func (o *InfoConfigDto) GetUploaded() string`

GetUploaded returns the Uploaded field if non-nil, zero value otherwise.

### GetUploadedOk

`func (o *InfoConfigDto) GetUploadedOk() (*string, bool)`

GetUploadedOk returns a tuple with the Uploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaded

`func (o *InfoConfigDto) SetUploaded(v string)`

SetUploaded sets Uploaded field to given value.

### HasUploaded

`func (o *InfoConfigDto) HasUploaded() bool`

HasUploaded returns a boolean if a field has been set.

### SetUploadedNil

`func (o *InfoConfigDto) SetUploadedNil(b bool)`

 SetUploadedNil sets the value for Uploaded to be an explicit nil

### UnsetUploaded
`func (o *InfoConfigDto) UnsetUploaded()`

UnsetUploaded ensures that no value is present for Uploaded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


