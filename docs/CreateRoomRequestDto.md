# CreateRoomRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **NullableString** | The room name. | 
**Quota** | Pointer to **NullableInt64** | The room quota. | [optional] 
**Indexing** | Pointer to **NullableBool** | Specifies whether to create a room with indexing. | [optional] 
**DenyDownload** | Pointer to **NullableBool** | Specifies whether to deny downloads from the room. | [optional] 
**Lifetime** | Pointer to [**RoomDataLifetimeDto**](RoomDataLifetimeDto.md) | The room data lifetime information. | [optional] 
**Watermark** | Pointer to [**WatermarkRequestDto**](WatermarkRequestDto.md) | The watermark settings. | [optional] 
**Logo** | Pointer to [**LogoRequest**](LogoRequest.md) | The room logo. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags. | [optional] 
**Color** | Pointer to **NullableString** | The room color, as a six-digit hexadecimal value without a leading '#'. | [optional] 
**Cover** | Pointer to **NullableString** | The room cover. | [optional] 
**RoomType** | [**RoomType**](RoomType.md) | The room type. | 
**Private** | Pointer to **bool** | Specifies whether the room to be created is private or not. | [optional] 
**Share** | Pointer to [**[]FileShareParams**](FileShareParams.md) | The collection of sharing parameters. | [optional] 
**ChatSettings** | Pointer to [**ChatSettings**](ChatSettings.md) | The chat settings. | [optional] 
**SendFormToExternalDB** | Pointer to **NullableBool** | Specifies whether to send form data to external database. | [optional] 
**SaveFormAsXLSX** | Pointer to **NullableBool** | Specifies whether to save form data as XLSX file. | [optional] 

## Methods

### NewCreateRoomRequestDto

`func NewCreateRoomRequestDto(title NullableString, roomType RoomType, ) *CreateRoomRequestDto`

NewCreateRoomRequestDto instantiates a new CreateRoomRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoomRequestDtoWithDefaults

`func NewCreateRoomRequestDtoWithDefaults() *CreateRoomRequestDto`

NewCreateRoomRequestDtoWithDefaults instantiates a new CreateRoomRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateRoomRequestDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateRoomRequestDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateRoomRequestDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CreateRoomRequestDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateRoomRequestDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetQuota

`func (o *CreateRoomRequestDto) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CreateRoomRequestDto) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CreateRoomRequestDto) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CreateRoomRequestDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *CreateRoomRequestDto) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *CreateRoomRequestDto) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetIndexing

`func (o *CreateRoomRequestDto) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *CreateRoomRequestDto) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *CreateRoomRequestDto) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *CreateRoomRequestDto) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### SetIndexingNil

`func (o *CreateRoomRequestDto) SetIndexingNil(b bool)`

 SetIndexingNil sets the value for Indexing to be an explicit nil

### UnsetIndexing
`func (o *CreateRoomRequestDto) UnsetIndexing()`

UnsetIndexing ensures that no value is present for Indexing, not even an explicit nil
### GetDenyDownload

`func (o *CreateRoomRequestDto) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *CreateRoomRequestDto) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *CreateRoomRequestDto) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *CreateRoomRequestDto) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### SetDenyDownloadNil

`func (o *CreateRoomRequestDto) SetDenyDownloadNil(b bool)`

 SetDenyDownloadNil sets the value for DenyDownload to be an explicit nil

### UnsetDenyDownload
`func (o *CreateRoomRequestDto) UnsetDenyDownload()`

UnsetDenyDownload ensures that no value is present for DenyDownload, not even an explicit nil
### GetLifetime

`func (o *CreateRoomRequestDto) GetLifetime() RoomDataLifetimeDto`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *CreateRoomRequestDto) GetLifetimeOk() (*RoomDataLifetimeDto, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *CreateRoomRequestDto) SetLifetime(v RoomDataLifetimeDto)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *CreateRoomRequestDto) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetWatermark

`func (o *CreateRoomRequestDto) GetWatermark() WatermarkRequestDto`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *CreateRoomRequestDto) GetWatermarkOk() (*WatermarkRequestDto, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *CreateRoomRequestDto) SetWatermark(v WatermarkRequestDto)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *CreateRoomRequestDto) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetLogo

`func (o *CreateRoomRequestDto) GetLogo() LogoRequest`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CreateRoomRequestDto) GetLogoOk() (*LogoRequest, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CreateRoomRequestDto) SetLogo(v LogoRequest)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CreateRoomRequestDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetTags

`func (o *CreateRoomRequestDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRoomRequestDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRoomRequestDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRoomRequestDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateRoomRequestDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateRoomRequestDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetColor

`func (o *CreateRoomRequestDto) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateRoomRequestDto) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateRoomRequestDto) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateRoomRequestDto) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateRoomRequestDto) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateRoomRequestDto) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *CreateRoomRequestDto) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *CreateRoomRequestDto) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *CreateRoomRequestDto) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *CreateRoomRequestDto) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *CreateRoomRequestDto) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *CreateRoomRequestDto) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil
### GetRoomType

`func (o *CreateRoomRequestDto) GetRoomType() RoomType`

GetRoomType returns the RoomType field if non-nil, zero value otherwise.

### GetRoomTypeOk

`func (o *CreateRoomRequestDto) GetRoomTypeOk() (*RoomType, bool)`

GetRoomTypeOk returns a tuple with the RoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomType

`func (o *CreateRoomRequestDto) SetRoomType(v RoomType)`

SetRoomType sets RoomType field to given value.


### GetPrivate

`func (o *CreateRoomRequestDto) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CreateRoomRequestDto) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CreateRoomRequestDto) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *CreateRoomRequestDto) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetShare

`func (o *CreateRoomRequestDto) GetShare() []FileShareParams`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *CreateRoomRequestDto) GetShareOk() (*[]FileShareParams, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *CreateRoomRequestDto) SetShare(v []FileShareParams)`

SetShare sets Share field to given value.

### HasShare

`func (o *CreateRoomRequestDto) HasShare() bool`

HasShare returns a boolean if a field has been set.

### SetShareNil

`func (o *CreateRoomRequestDto) SetShareNil(b bool)`

 SetShareNil sets the value for Share to be an explicit nil

### UnsetShare
`func (o *CreateRoomRequestDto) UnsetShare()`

UnsetShare ensures that no value is present for Share, not even an explicit nil
### GetChatSettings

`func (o *CreateRoomRequestDto) GetChatSettings() ChatSettings`

GetChatSettings returns the ChatSettings field if non-nil, zero value otherwise.

### GetChatSettingsOk

`func (o *CreateRoomRequestDto) GetChatSettingsOk() (*ChatSettings, bool)`

GetChatSettingsOk returns a tuple with the ChatSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSettings

`func (o *CreateRoomRequestDto) SetChatSettings(v ChatSettings)`

SetChatSettings sets ChatSettings field to given value.

### HasChatSettings

`func (o *CreateRoomRequestDto) HasChatSettings() bool`

HasChatSettings returns a boolean if a field has been set.

### GetSendFormToExternalDB

`func (o *CreateRoomRequestDto) GetSendFormToExternalDB() bool`

GetSendFormToExternalDB returns the SendFormToExternalDB field if non-nil, zero value otherwise.

### GetSendFormToExternalDBOk

`func (o *CreateRoomRequestDto) GetSendFormToExternalDBOk() (*bool, bool)`

GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFormToExternalDB

`func (o *CreateRoomRequestDto) SetSendFormToExternalDB(v bool)`

SetSendFormToExternalDB sets SendFormToExternalDB field to given value.

### HasSendFormToExternalDB

`func (o *CreateRoomRequestDto) HasSendFormToExternalDB() bool`

HasSendFormToExternalDB returns a boolean if a field has been set.

### SetSendFormToExternalDBNil

`func (o *CreateRoomRequestDto) SetSendFormToExternalDBNil(b bool)`

 SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil

### UnsetSendFormToExternalDB
`func (o *CreateRoomRequestDto) UnsetSendFormToExternalDB()`

UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
### GetSaveFormAsXLSX

`func (o *CreateRoomRequestDto) GetSaveFormAsXLSX() bool`

GetSaveFormAsXLSX returns the SaveFormAsXLSX field if non-nil, zero value otherwise.

### GetSaveFormAsXLSXOk

`func (o *CreateRoomRequestDto) GetSaveFormAsXLSXOk() (*bool, bool)`

GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveFormAsXLSX

`func (o *CreateRoomRequestDto) SetSaveFormAsXLSX(v bool)`

SetSaveFormAsXLSX sets SaveFormAsXLSX field to given value.

### HasSaveFormAsXLSX

`func (o *CreateRoomRequestDto) HasSaveFormAsXLSX() bool`

HasSaveFormAsXLSX returns a boolean if a field has been set.

### SetSaveFormAsXLSXNil

`func (o *CreateRoomRequestDto) SetSaveFormAsXLSXNil(b bool)`

 SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil

### UnsetSaveFormAsXLSX
`func (o *CreateRoomRequestDto) UnsetSaveFormAsXLSX()`

UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


