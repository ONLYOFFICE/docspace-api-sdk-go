# UpdateRoomRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | The room title. | [optional] 
**Quota** | Pointer to **NullableInt64** | The room quota. | [optional] 
**Indexing** | Pointer to **NullableBool** | Specifies whether to create a third-party room with indexing. | [optional] 
**DenyDownload** | Pointer to **NullableBool** | Specifies whether to deny downloads from the third-party room. | [optional] 
**Lifetime** | Pointer to [**RoomDataLifetimeDto**](RoomDataLifetimeDto.md) |  | [optional] 
**Watermark** | Pointer to [**WatermarkRequestDto**](WatermarkRequestDto.md) |  | [optional] 
**Logo** | Pointer to [**LogoRequest**](LogoRequest.md) |  | [optional] 
**Tags** | Pointer to **[]string** | The list of tags. | [optional] 
**Color** | Pointer to **NullableString** | The room color. | [optional] 
**Cover** | Pointer to **NullableString** | The room cover. | [optional] 
**ChatSettings** | Pointer to [**ChatSettings**](ChatSettings.md) |  | [optional] 
**SendFormToExternalDB** | Pointer to **NullableBool** | Specifies whether to send form data to external database. | [optional] 
**SaveFormAsXLSX** | Pointer to **NullableBool** | Specifies whether to save form data as XLSX file. | [optional] 

## Methods

### NewUpdateRoomRequest

`func NewUpdateRoomRequest() *UpdateRoomRequest`

NewUpdateRoomRequest instantiates a new UpdateRoomRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoomRequestWithDefaults

`func NewUpdateRoomRequestWithDefaults() *UpdateRoomRequest`

NewUpdateRoomRequestWithDefaults instantiates a new UpdateRoomRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateRoomRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateRoomRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateRoomRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateRoomRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpdateRoomRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateRoomRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetQuota

`func (o *UpdateRoomRequest) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UpdateRoomRequest) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UpdateRoomRequest) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *UpdateRoomRequest) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *UpdateRoomRequest) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *UpdateRoomRequest) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetIndexing

`func (o *UpdateRoomRequest) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *UpdateRoomRequest) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *UpdateRoomRequest) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *UpdateRoomRequest) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### SetIndexingNil

`func (o *UpdateRoomRequest) SetIndexingNil(b bool)`

 SetIndexingNil sets the value for Indexing to be an explicit nil

### UnsetIndexing
`func (o *UpdateRoomRequest) UnsetIndexing()`

UnsetIndexing ensures that no value is present for Indexing, not even an explicit nil
### GetDenyDownload

`func (o *UpdateRoomRequest) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *UpdateRoomRequest) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *UpdateRoomRequest) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *UpdateRoomRequest) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### SetDenyDownloadNil

`func (o *UpdateRoomRequest) SetDenyDownloadNil(b bool)`

 SetDenyDownloadNil sets the value for DenyDownload to be an explicit nil

### UnsetDenyDownload
`func (o *UpdateRoomRequest) UnsetDenyDownload()`

UnsetDenyDownload ensures that no value is present for DenyDownload, not even an explicit nil
### GetLifetime

`func (o *UpdateRoomRequest) GetLifetime() RoomDataLifetimeDto`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *UpdateRoomRequest) GetLifetimeOk() (*RoomDataLifetimeDto, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *UpdateRoomRequest) SetLifetime(v RoomDataLifetimeDto)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *UpdateRoomRequest) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetWatermark

`func (o *UpdateRoomRequest) GetWatermark() WatermarkRequestDto`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *UpdateRoomRequest) GetWatermarkOk() (*WatermarkRequestDto, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *UpdateRoomRequest) SetWatermark(v WatermarkRequestDto)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *UpdateRoomRequest) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetLogo

`func (o *UpdateRoomRequest) GetLogo() LogoRequest`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *UpdateRoomRequest) GetLogoOk() (*LogoRequest, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *UpdateRoomRequest) SetLogo(v LogoRequest)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *UpdateRoomRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetTags

`func (o *UpdateRoomRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateRoomRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateRoomRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateRoomRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateRoomRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateRoomRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetColor

`func (o *UpdateRoomRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateRoomRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateRoomRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateRoomRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *UpdateRoomRequest) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *UpdateRoomRequest) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *UpdateRoomRequest) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *UpdateRoomRequest) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *UpdateRoomRequest) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *UpdateRoomRequest) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *UpdateRoomRequest) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *UpdateRoomRequest) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil
### GetChatSettings

`func (o *UpdateRoomRequest) GetChatSettings() ChatSettings`

GetChatSettings returns the ChatSettings field if non-nil, zero value otherwise.

### GetChatSettingsOk

`func (o *UpdateRoomRequest) GetChatSettingsOk() (*ChatSettings, bool)`

GetChatSettingsOk returns a tuple with the ChatSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSettings

`func (o *UpdateRoomRequest) SetChatSettings(v ChatSettings)`

SetChatSettings sets ChatSettings field to given value.

### HasChatSettings

`func (o *UpdateRoomRequest) HasChatSettings() bool`

HasChatSettings returns a boolean if a field has been set.

### GetSendFormToExternalDB

`func (o *UpdateRoomRequest) GetSendFormToExternalDB() bool`

GetSendFormToExternalDB returns the SendFormToExternalDB field if non-nil, zero value otherwise.

### GetSendFormToExternalDBOk

`func (o *UpdateRoomRequest) GetSendFormToExternalDBOk() (*bool, bool)`

GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFormToExternalDB

`func (o *UpdateRoomRequest) SetSendFormToExternalDB(v bool)`

SetSendFormToExternalDB sets SendFormToExternalDB field to given value.

### HasSendFormToExternalDB

`func (o *UpdateRoomRequest) HasSendFormToExternalDB() bool`

HasSendFormToExternalDB returns a boolean if a field has been set.

### SetSendFormToExternalDBNil

`func (o *UpdateRoomRequest) SetSendFormToExternalDBNil(b bool)`

 SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil

### UnsetSendFormToExternalDB
`func (o *UpdateRoomRequest) UnsetSendFormToExternalDB()`

UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
### GetSaveFormAsXLSX

`func (o *UpdateRoomRequest) GetSaveFormAsXLSX() bool`

GetSaveFormAsXLSX returns the SaveFormAsXLSX field if non-nil, zero value otherwise.

### GetSaveFormAsXLSXOk

`func (o *UpdateRoomRequest) GetSaveFormAsXLSXOk() (*bool, bool)`

GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveFormAsXLSX

`func (o *UpdateRoomRequest) SetSaveFormAsXLSX(v bool)`

SetSaveFormAsXLSX sets SaveFormAsXLSX field to given value.

### HasSaveFormAsXLSX

`func (o *UpdateRoomRequest) HasSaveFormAsXLSX() bool`

HasSaveFormAsXLSX returns a boolean if a field has been set.

### SetSaveFormAsXLSXNil

`func (o *UpdateRoomRequest) SetSaveFormAsXLSXNil(b bool)`

 SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil

### UnsetSaveFormAsXLSX
`func (o *UpdateRoomRequest) UnsetSaveFormAsXLSX()`

UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


