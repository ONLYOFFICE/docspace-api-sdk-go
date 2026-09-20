# CreateRoomFromTemplateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **int32** | The room template to copy. Templates live in their own section and are listed by `GET api/2.0/files/rooms`  with a search area of 4; an ordinary room id is rejected here. | 
**Title** | **NullableString** | The name of the room to create. It is sanitised and truncated the way a room title is, and a blank value is  rejected; the title of the template is not reused. | 
**Logo** | Pointer to [**LogoRequest**](LogoRequest.md) | The picture to use as the room logo, named by the path that `POST api/2.0/files/logos` returned for an image  uploaded beforehand, plus the crop to take from it. Leaving the field out keeps the room on its cover and  colour. It is ignored when the logo of the template is copied instead. | [optional] 
**CopyLogo** | Pointer to **bool** | Whether the new room keeps the logo of the template. With it on the uploaded picture is ignored; with it off  the room starts with no logo unless one is supplied. | [optional] 
**Tags** | Pointer to **[]string** | The labels to attach to the room, by name. Names the portal tag catalogue does not hold yet are added to it,  and `GET api/2.0/files/tags` lists what already exists. Leaving the field out keeps the tags of the template. | [optional] 
**Color** | Pointer to **NullableString** | The background colour the room is drawn with while it has no logo, as six hexadecimal digits with no leading  number sign. An empty value restores the default colour of the room type. | [optional] 
**Cover** | Pointer to **NullableString** | The picture drawn on the room while it has no logo, named by an identifier from  `GET api/2.0/files/rooms/covers`. Any other value is rejected, and an empty value leaves the room without a  cover. | [optional] 
**Quota** | Pointer to **NullableInt64** | The storage the room may take, in bytes. It is accepted only while the per-room quota feature is on for the  portal and must stay inside the portal own limit; leaving it out lets the room follow the portal default. | [optional] 
**Indexing** | Pointer to **NullableBool** | Whether the room keeps a manual order of its contents. With it on every file and folder carries a position  that listings follow and that `PUT api/2.0/files/rooms/{id}/reorder` compacts; with it off the contents are  ordered by the sorting of the request. Leaving it out keeps the setting of the template. | [optional] 
**DenyDownload** | Pointer to **NullableBool** | Whether members without editing rights are stopped from downloading and printing the contents of the room.  They can still open the documents in the editor. Leaving it out keeps the setting of the template. | [optional] 
**Lifetime** | Pointer to [**RoomDataLifetimeDto**](RoomDataLifetimeDto.md) | How long files may stay in the room before they are deleted automatically. The countdown starts when the  setting is saved, and leaving the field out keeps the files forever. Leaving the field out keeps the setting  of the template. | [optional] 
**Watermark** | Pointer to [**WatermarkRequestDto**](WatermarkRequestDto.md) | The watermark drawn over documents opened in the room. Leaving the field out adds no watermark, and sending it  with the switch turned off removes the one the room has. Leaving the field out keeps the setting of the  template. | [optional] 
**Private** | Pointer to **NullableBool** | Whether the room is end-to-end encrypted. Its files can then be opened only in the desktop application by  members whose encryption keys are set up, and the flag cannot be changed after the room is created. | [optional] 

## Methods

### NewCreateRoomFromTemplateDto

`func NewCreateRoomFromTemplateDto(templateId int32, title NullableString, ) *CreateRoomFromTemplateDto`

NewCreateRoomFromTemplateDto instantiates a new CreateRoomFromTemplateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoomFromTemplateDtoWithDefaults

`func NewCreateRoomFromTemplateDtoWithDefaults() *CreateRoomFromTemplateDto`

NewCreateRoomFromTemplateDtoWithDefaults instantiates a new CreateRoomFromTemplateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *CreateRoomFromTemplateDto) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateRoomFromTemplateDto) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateRoomFromTemplateDto) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.


### GetTitle

`func (o *CreateRoomFromTemplateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateRoomFromTemplateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateRoomFromTemplateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CreateRoomFromTemplateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateRoomFromTemplateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLogo

`func (o *CreateRoomFromTemplateDto) GetLogo() LogoRequest`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CreateRoomFromTemplateDto) GetLogoOk() (*LogoRequest, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CreateRoomFromTemplateDto) SetLogo(v LogoRequest)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CreateRoomFromTemplateDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetCopyLogo

`func (o *CreateRoomFromTemplateDto) GetCopyLogo() bool`

GetCopyLogo returns the CopyLogo field if non-nil, zero value otherwise.

### GetCopyLogoOk

`func (o *CreateRoomFromTemplateDto) GetCopyLogoOk() (*bool, bool)`

GetCopyLogoOk returns a tuple with the CopyLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyLogo

`func (o *CreateRoomFromTemplateDto) SetCopyLogo(v bool)`

SetCopyLogo sets CopyLogo field to given value.

### HasCopyLogo

`func (o *CreateRoomFromTemplateDto) HasCopyLogo() bool`

HasCopyLogo returns a boolean if a field has been set.

### GetTags

`func (o *CreateRoomFromTemplateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRoomFromTemplateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRoomFromTemplateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRoomFromTemplateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateRoomFromTemplateDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateRoomFromTemplateDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetColor

`func (o *CreateRoomFromTemplateDto) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateRoomFromTemplateDto) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateRoomFromTemplateDto) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateRoomFromTemplateDto) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateRoomFromTemplateDto) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateRoomFromTemplateDto) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *CreateRoomFromTemplateDto) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *CreateRoomFromTemplateDto) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *CreateRoomFromTemplateDto) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *CreateRoomFromTemplateDto) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *CreateRoomFromTemplateDto) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *CreateRoomFromTemplateDto) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil
### GetQuota

`func (o *CreateRoomFromTemplateDto) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CreateRoomFromTemplateDto) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CreateRoomFromTemplateDto) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CreateRoomFromTemplateDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *CreateRoomFromTemplateDto) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *CreateRoomFromTemplateDto) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetIndexing

`func (o *CreateRoomFromTemplateDto) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *CreateRoomFromTemplateDto) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *CreateRoomFromTemplateDto) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *CreateRoomFromTemplateDto) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### SetIndexingNil

`func (o *CreateRoomFromTemplateDto) SetIndexingNil(b bool)`

 SetIndexingNil sets the value for Indexing to be an explicit nil

### UnsetIndexing
`func (o *CreateRoomFromTemplateDto) UnsetIndexing()`

UnsetIndexing ensures that no value is present for Indexing, not even an explicit nil
### GetDenyDownload

`func (o *CreateRoomFromTemplateDto) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *CreateRoomFromTemplateDto) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *CreateRoomFromTemplateDto) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *CreateRoomFromTemplateDto) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### SetDenyDownloadNil

`func (o *CreateRoomFromTemplateDto) SetDenyDownloadNil(b bool)`

 SetDenyDownloadNil sets the value for DenyDownload to be an explicit nil

### UnsetDenyDownload
`func (o *CreateRoomFromTemplateDto) UnsetDenyDownload()`

UnsetDenyDownload ensures that no value is present for DenyDownload, not even an explicit nil
### GetLifetime

`func (o *CreateRoomFromTemplateDto) GetLifetime() RoomDataLifetimeDto`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *CreateRoomFromTemplateDto) GetLifetimeOk() (*RoomDataLifetimeDto, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *CreateRoomFromTemplateDto) SetLifetime(v RoomDataLifetimeDto)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *CreateRoomFromTemplateDto) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetWatermark

`func (o *CreateRoomFromTemplateDto) GetWatermark() WatermarkRequestDto`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *CreateRoomFromTemplateDto) GetWatermarkOk() (*WatermarkRequestDto, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *CreateRoomFromTemplateDto) SetWatermark(v WatermarkRequestDto)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *CreateRoomFromTemplateDto) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetPrivate

`func (o *CreateRoomFromTemplateDto) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CreateRoomFromTemplateDto) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CreateRoomFromTemplateDto) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *CreateRoomFromTemplateDto) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### SetPrivateNil

`func (o *CreateRoomFromTemplateDto) SetPrivateNil(b bool)`

 SetPrivateNil sets the value for Private to be an explicit nil

### UnsetPrivate
`func (o *CreateRoomFromTemplateDto) UnsetPrivate()`

UnsetPrivate ensures that no value is present for Private, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


