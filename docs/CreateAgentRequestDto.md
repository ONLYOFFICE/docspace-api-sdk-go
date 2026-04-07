# CreateAgentRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **NullableString** | The room name. | 
**Quota** | Pointer to **NullableInt64** | The room quota. | [optional] 
**Indexing** | Pointer to **NullableBool** | Specifies whether to create a room with indexing. | [optional] 
**DenyDownload** | Pointer to **NullableBool** | Specifies whether to deny downloads from the room. | [optional] 
**Lifetime** | Pointer to [**RoomDataLifetimeDto**](RoomDataLifetimeDto.md) |  | [optional] 
**Watermark** | Pointer to [**WatermarkRequestDto**](WatermarkRequestDto.md) |  | [optional] 
**Logo** | Pointer to [**LogoRequest**](LogoRequest.md) |  | [optional] 
**Tags** | Pointer to **[]string** | The list of tags. | [optional] 
**Color** | Pointer to **NullableString** | The room color. | [optional] 
**Cover** | Pointer to **NullableString** | The room cover. | [optional] 
**Private** | Pointer to **bool** | Specifies whether the room to be created is private or not. | [optional] 
**Share** | Pointer to [**[]FileShareParams**](FileShareParams.md) | The collection of sharing parameters. | [optional] 
**ChatSettings** | [**ChatSettings**](ChatSettings.md) |  | 
**AttachDefaultTools** | Pointer to **bool** | Specifies whether to attach default tools to the agent or not. | [optional] 

## Methods

### NewCreateAgentRequestDto

`func NewCreateAgentRequestDto(title NullableString, chatSettings ChatSettings, ) *CreateAgentRequestDto`

NewCreateAgentRequestDto instantiates a new CreateAgentRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentRequestDtoWithDefaults

`func NewCreateAgentRequestDtoWithDefaults() *CreateAgentRequestDto`

NewCreateAgentRequestDtoWithDefaults instantiates a new CreateAgentRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateAgentRequestDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateAgentRequestDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateAgentRequestDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CreateAgentRequestDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateAgentRequestDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetQuota

`func (o *CreateAgentRequestDto) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CreateAgentRequestDto) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CreateAgentRequestDto) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CreateAgentRequestDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *CreateAgentRequestDto) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *CreateAgentRequestDto) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetIndexing

`func (o *CreateAgentRequestDto) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *CreateAgentRequestDto) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *CreateAgentRequestDto) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *CreateAgentRequestDto) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### SetIndexingNil

`func (o *CreateAgentRequestDto) SetIndexingNil(b bool)`

 SetIndexingNil sets the value for Indexing to be an explicit nil

### UnsetIndexing
`func (o *CreateAgentRequestDto) UnsetIndexing()`

UnsetIndexing ensures that no value is present for Indexing, not even an explicit nil
### GetDenyDownload

`func (o *CreateAgentRequestDto) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *CreateAgentRequestDto) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *CreateAgentRequestDto) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *CreateAgentRequestDto) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### SetDenyDownloadNil

`func (o *CreateAgentRequestDto) SetDenyDownloadNil(b bool)`

 SetDenyDownloadNil sets the value for DenyDownload to be an explicit nil

### UnsetDenyDownload
`func (o *CreateAgentRequestDto) UnsetDenyDownload()`

UnsetDenyDownload ensures that no value is present for DenyDownload, not even an explicit nil
### GetLifetime

`func (o *CreateAgentRequestDto) GetLifetime() RoomDataLifetimeDto`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *CreateAgentRequestDto) GetLifetimeOk() (*RoomDataLifetimeDto, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *CreateAgentRequestDto) SetLifetime(v RoomDataLifetimeDto)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *CreateAgentRequestDto) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetWatermark

`func (o *CreateAgentRequestDto) GetWatermark() WatermarkRequestDto`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *CreateAgentRequestDto) GetWatermarkOk() (*WatermarkRequestDto, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *CreateAgentRequestDto) SetWatermark(v WatermarkRequestDto)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *CreateAgentRequestDto) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetLogo

`func (o *CreateAgentRequestDto) GetLogo() LogoRequest`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CreateAgentRequestDto) GetLogoOk() (*LogoRequest, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CreateAgentRequestDto) SetLogo(v LogoRequest)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CreateAgentRequestDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetTags

`func (o *CreateAgentRequestDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateAgentRequestDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateAgentRequestDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateAgentRequestDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateAgentRequestDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateAgentRequestDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetColor

`func (o *CreateAgentRequestDto) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateAgentRequestDto) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateAgentRequestDto) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateAgentRequestDto) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateAgentRequestDto) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateAgentRequestDto) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *CreateAgentRequestDto) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *CreateAgentRequestDto) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *CreateAgentRequestDto) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *CreateAgentRequestDto) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *CreateAgentRequestDto) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *CreateAgentRequestDto) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil
### GetPrivate

`func (o *CreateAgentRequestDto) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CreateAgentRequestDto) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CreateAgentRequestDto) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *CreateAgentRequestDto) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetShare

`func (o *CreateAgentRequestDto) GetShare() []FileShareParams`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *CreateAgentRequestDto) GetShareOk() (*[]FileShareParams, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *CreateAgentRequestDto) SetShare(v []FileShareParams)`

SetShare sets Share field to given value.

### HasShare

`func (o *CreateAgentRequestDto) HasShare() bool`

HasShare returns a boolean if a field has been set.

### SetShareNil

`func (o *CreateAgentRequestDto) SetShareNil(b bool)`

 SetShareNil sets the value for Share to be an explicit nil

### UnsetShare
`func (o *CreateAgentRequestDto) UnsetShare()`

UnsetShare ensures that no value is present for Share, not even an explicit nil
### GetChatSettings

`func (o *CreateAgentRequestDto) GetChatSettings() ChatSettings`

GetChatSettings returns the ChatSettings field if non-nil, zero value otherwise.

### GetChatSettingsOk

`func (o *CreateAgentRequestDto) GetChatSettingsOk() (*ChatSettings, bool)`

GetChatSettingsOk returns a tuple with the ChatSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSettings

`func (o *CreateAgentRequestDto) SetChatSettings(v ChatSettings)`

SetChatSettings sets ChatSettings field to given value.


### GetAttachDefaultTools

`func (o *CreateAgentRequestDto) GetAttachDefaultTools() bool`

GetAttachDefaultTools returns the AttachDefaultTools field if non-nil, zero value otherwise.

### GetAttachDefaultToolsOk

`func (o *CreateAgentRequestDto) GetAttachDefaultToolsOk() (*bool, bool)`

GetAttachDefaultToolsOk returns a tuple with the AttachDefaultTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachDefaultTools

`func (o *CreateAgentRequestDto) SetAttachDefaultTools(v bool)`

SetAttachDefaultTools sets AttachDefaultTools field to given value.

### HasAttachDefaultTools

`func (o *CreateAgentRequestDto) HasAttachDefaultTools() bool`

HasAttachDefaultTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


