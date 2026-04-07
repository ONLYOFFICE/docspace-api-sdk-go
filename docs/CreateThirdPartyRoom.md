# CreateThirdPartyRoom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateAsNewFolder** | Pointer to **bool** | Specifies whether to create a third-party room as a new folder or not. | [optional] 
**Title** | **NullableString** | The third-party room name to be created. | 
**RoomType** | [**RoomType**](RoomType.md) |  | 
**Private** | Pointer to **bool** | Specifies whether to create the private third-party room or not. | [optional] 
**Indexing** | Pointer to **bool** | Specifies whether to create the third-party room with indexing. | [optional] 
**DenyDownload** | Pointer to **bool** | Specifies whether to deny downloads from the third-party room. | [optional] 
**Color** | Pointer to **NullableString** | The color of the third-party room. | [optional] 
**Cover** | Pointer to **NullableString** | The cover of the third-party room. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of the third-party room. | [optional] 
**Logo** | Pointer to [**LogoRequest**](LogoRequest.md) |  | [optional] 

## Methods

### NewCreateThirdPartyRoom

`func NewCreateThirdPartyRoom(title NullableString, roomType RoomType, ) *CreateThirdPartyRoom`

NewCreateThirdPartyRoom instantiates a new CreateThirdPartyRoom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateThirdPartyRoomWithDefaults

`func NewCreateThirdPartyRoomWithDefaults() *CreateThirdPartyRoom`

NewCreateThirdPartyRoomWithDefaults instantiates a new CreateThirdPartyRoom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateAsNewFolder

`func (o *CreateThirdPartyRoom) GetCreateAsNewFolder() bool`

GetCreateAsNewFolder returns the CreateAsNewFolder field if non-nil, zero value otherwise.

### GetCreateAsNewFolderOk

`func (o *CreateThirdPartyRoom) GetCreateAsNewFolderOk() (*bool, bool)`

GetCreateAsNewFolderOk returns a tuple with the CreateAsNewFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAsNewFolder

`func (o *CreateThirdPartyRoom) SetCreateAsNewFolder(v bool)`

SetCreateAsNewFolder sets CreateAsNewFolder field to given value.

### HasCreateAsNewFolder

`func (o *CreateThirdPartyRoom) HasCreateAsNewFolder() bool`

HasCreateAsNewFolder returns a boolean if a field has been set.

### GetTitle

`func (o *CreateThirdPartyRoom) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateThirdPartyRoom) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateThirdPartyRoom) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CreateThirdPartyRoom) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateThirdPartyRoom) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetRoomType

`func (o *CreateThirdPartyRoom) GetRoomType() RoomType`

GetRoomType returns the RoomType field if non-nil, zero value otherwise.

### GetRoomTypeOk

`func (o *CreateThirdPartyRoom) GetRoomTypeOk() (*RoomType, bool)`

GetRoomTypeOk returns a tuple with the RoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomType

`func (o *CreateThirdPartyRoom) SetRoomType(v RoomType)`

SetRoomType sets RoomType field to given value.


### GetPrivate

`func (o *CreateThirdPartyRoom) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CreateThirdPartyRoom) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CreateThirdPartyRoom) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *CreateThirdPartyRoom) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetIndexing

`func (o *CreateThirdPartyRoom) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *CreateThirdPartyRoom) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *CreateThirdPartyRoom) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *CreateThirdPartyRoom) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### GetDenyDownload

`func (o *CreateThirdPartyRoom) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *CreateThirdPartyRoom) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *CreateThirdPartyRoom) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *CreateThirdPartyRoom) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetColor

`func (o *CreateThirdPartyRoom) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateThirdPartyRoom) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateThirdPartyRoom) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateThirdPartyRoom) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateThirdPartyRoom) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateThirdPartyRoom) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *CreateThirdPartyRoom) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *CreateThirdPartyRoom) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *CreateThirdPartyRoom) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *CreateThirdPartyRoom) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *CreateThirdPartyRoom) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *CreateThirdPartyRoom) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil
### GetTags

`func (o *CreateThirdPartyRoom) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateThirdPartyRoom) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateThirdPartyRoom) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateThirdPartyRoom) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateThirdPartyRoom) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateThirdPartyRoom) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetLogo

`func (o *CreateThirdPartyRoom) GetLogo() LogoRequest`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CreateThirdPartyRoom) GetLogoOk() (*LogoRequest, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CreateThirdPartyRoom) SetLogo(v LogoRequest)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CreateThirdPartyRoom) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


