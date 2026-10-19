# CreateThirdPartyRoomRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateAsNewFolder** | Pointer to **bool** | Creates a new folder named after `title` inside the folder named in the path and turns that subfolder into the  room, leaving the named folder itself untouched. When omitted, the named folder becomes the room and keeps  everything it already holds. | [optional] 
**Title** | **NullableString** | The name the room is shown under. It is stored on the connected account, so it does not have to match the name  of the folder in the storage; with `createAsNewFolder` it is also the name given to the created subfolder. | 
**RoomType** | [**RoomType**](RoomType.md) | The kind of room the folder becomes, which decides the default access rules of its members and cannot be  changed afterwards. | 
**Private** | Pointer to **bool** | Restricts the room to the members explicitly invited into it. The flag is kept on the connected storage  account rather than on the folder, so every folder read through that account reports the same value. | [optional] 
**Indexing** | Pointer to **bool** | Keeps the contents of the room in an explicit numbered order, the one reported as `order` on every entry,  instead of leaving the order to the reader. | [optional] 
**DenyDownload** | Pointer to **bool** | Forbids downloading and printing the contents of the room, which leaves the members with viewing and editing  in the editor only. | [optional] 
**Color** | Pointer to **NullableString** | The background colour drawn behind the cover of the room, as six hexadecimal digits without a leading number  sign. An empty value restores the colour the portal picks by default. | [optional] 
**Cover** | Pointer to **NullableString** | The drawing shown on the room tile, named by one of the built-in cover identifiers returned by  `GET api/2.0/files/rooms/covers`. An empty value leaves the room without a cover, and any other unknown value  is rejected as an invalid request. | [optional] 
**Tags** | Pointer to **[]string** | The tags to attach to the room, named by their text. A name that is not in the portal tag catalogue yet is  added to it, and `GET api/2.0/files/tags` lists the names already there. | [optional] 
**Logo** | Pointer to [**LogoRequest**](LogoRequest.md) | The picture to use as the room logo, which has to be uploaded with `POST api/2.0/files/logos` first; leaving  it out keeps the room on its cover and colour. | [optional] 

## Methods

### NewCreateThirdPartyRoomRequest

`func NewCreateThirdPartyRoomRequest(title NullableString, roomType RoomType, ) *CreateThirdPartyRoomRequest`

NewCreateThirdPartyRoomRequest instantiates a new CreateThirdPartyRoomRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateThirdPartyRoomRequestWithDefaults

`func NewCreateThirdPartyRoomRequestWithDefaults() *CreateThirdPartyRoomRequest`

NewCreateThirdPartyRoomRequestWithDefaults instantiates a new CreateThirdPartyRoomRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateAsNewFolder

`func (o *CreateThirdPartyRoomRequest) GetCreateAsNewFolder() bool`

GetCreateAsNewFolder returns the CreateAsNewFolder field if non-nil, zero value otherwise.

### GetCreateAsNewFolderOk

`func (o *CreateThirdPartyRoomRequest) GetCreateAsNewFolderOk() (*bool, bool)`

GetCreateAsNewFolderOk returns a tuple with the CreateAsNewFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAsNewFolder

`func (o *CreateThirdPartyRoomRequest) SetCreateAsNewFolder(v bool)`

SetCreateAsNewFolder sets CreateAsNewFolder field to given value.

### HasCreateAsNewFolder

`func (o *CreateThirdPartyRoomRequest) HasCreateAsNewFolder() bool`

HasCreateAsNewFolder returns a boolean if a field has been set.

### GetTitle

`func (o *CreateThirdPartyRoomRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateThirdPartyRoomRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateThirdPartyRoomRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CreateThirdPartyRoomRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateThirdPartyRoomRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetRoomType

`func (o *CreateThirdPartyRoomRequest) GetRoomType() RoomType`

GetRoomType returns the RoomType field if non-nil, zero value otherwise.

### GetRoomTypeOk

`func (o *CreateThirdPartyRoomRequest) GetRoomTypeOk() (*RoomType, bool)`

GetRoomTypeOk returns a tuple with the RoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomType

`func (o *CreateThirdPartyRoomRequest) SetRoomType(v RoomType)`

SetRoomType sets RoomType field to given value.


### GetPrivate

`func (o *CreateThirdPartyRoomRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CreateThirdPartyRoomRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CreateThirdPartyRoomRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *CreateThirdPartyRoomRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetIndexing

`func (o *CreateThirdPartyRoomRequest) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *CreateThirdPartyRoomRequest) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *CreateThirdPartyRoomRequest) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *CreateThirdPartyRoomRequest) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### GetDenyDownload

`func (o *CreateThirdPartyRoomRequest) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *CreateThirdPartyRoomRequest) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *CreateThirdPartyRoomRequest) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *CreateThirdPartyRoomRequest) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetColor

`func (o *CreateThirdPartyRoomRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateThirdPartyRoomRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateThirdPartyRoomRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateThirdPartyRoomRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateThirdPartyRoomRequest) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateThirdPartyRoomRequest) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *CreateThirdPartyRoomRequest) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *CreateThirdPartyRoomRequest) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *CreateThirdPartyRoomRequest) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *CreateThirdPartyRoomRequest) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *CreateThirdPartyRoomRequest) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *CreateThirdPartyRoomRequest) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil
### GetTags

`func (o *CreateThirdPartyRoomRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateThirdPartyRoomRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateThirdPartyRoomRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateThirdPartyRoomRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateThirdPartyRoomRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateThirdPartyRoomRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetLogo

`func (o *CreateThirdPartyRoomRequest) GetLogo() LogoRequest`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CreateThirdPartyRoomRequest) GetLogoOk() (*LogoRequest, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CreateThirdPartyRoomRequest) SetLogo(v LogoRequest)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CreateThirdPartyRoomRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


