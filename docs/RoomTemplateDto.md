# RoomTemplateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | **int32** | The room template ID. | 
**Title** | **string** | The room template title. | 
**Logo** | Pointer to [**LogoRequest**](LogoRequest.md) | The room template logo. | [optional] 
**CopyLogo** | Pointer to **bool** | Specifies whether to copy room logo or not. | [optional] 
**Share** | Pointer to **[]string** | The collection of email addresses of users with whom to share a room. | [optional] 
**Groups** | Pointer to **[]string** | The collection of groups with whom to share a room. | [optional] 
**Public** | Pointer to **bool** | Specifies whether the room template is public or not. | [optional] 
**Tags** | Pointer to **[]string** | The collection of tags. | [optional] 
**Color** | Pointer to **NullableString** | The color of the room template. | [optional] 
**Cover** | Pointer to **NullableString** | The cover of the room template. | [optional] 
**Quota** | Pointer to **NullableInt64** | Room quota | [optional] 

## Methods

### NewRoomTemplateDto

`func NewRoomTemplateDto(roomId int32, title string, ) *RoomTemplateDto`

NewRoomTemplateDto instantiates a new RoomTemplateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomTemplateDtoWithDefaults

`func NewRoomTemplateDtoWithDefaults() *RoomTemplateDto`

NewRoomTemplateDtoWithDefaults instantiates a new RoomTemplateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *RoomTemplateDto) GetRoomId() int32`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *RoomTemplateDto) GetRoomIdOk() (*int32, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *RoomTemplateDto) SetRoomId(v int32)`

SetRoomId sets RoomId field to given value.


### GetTitle

`func (o *RoomTemplateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RoomTemplateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RoomTemplateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetLogo

`func (o *RoomTemplateDto) GetLogo() LogoRequest`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *RoomTemplateDto) GetLogoOk() (*LogoRequest, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *RoomTemplateDto) SetLogo(v LogoRequest)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *RoomTemplateDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetCopyLogo

`func (o *RoomTemplateDto) GetCopyLogo() bool`

GetCopyLogo returns the CopyLogo field if non-nil, zero value otherwise.

### GetCopyLogoOk

`func (o *RoomTemplateDto) GetCopyLogoOk() (*bool, bool)`

GetCopyLogoOk returns a tuple with the CopyLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyLogo

`func (o *RoomTemplateDto) SetCopyLogo(v bool)`

SetCopyLogo sets CopyLogo field to given value.

### HasCopyLogo

`func (o *RoomTemplateDto) HasCopyLogo() bool`

HasCopyLogo returns a boolean if a field has been set.

### GetShare

`func (o *RoomTemplateDto) GetShare() []string`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *RoomTemplateDto) GetShareOk() (*[]string, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *RoomTemplateDto) SetShare(v []string)`

SetShare sets Share field to given value.

### HasShare

`func (o *RoomTemplateDto) HasShare() bool`

HasShare returns a boolean if a field has been set.

### SetShareNil

`func (o *RoomTemplateDto) SetShareNil(b bool)`

 SetShareNil sets the value for Share to be an explicit nil

### UnsetShare
`func (o *RoomTemplateDto) UnsetShare()`

UnsetShare ensures that no value is present for Share, not even an explicit nil
### GetGroups

`func (o *RoomTemplateDto) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RoomTemplateDto) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RoomTemplateDto) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *RoomTemplateDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *RoomTemplateDto) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *RoomTemplateDto) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetPublic

`func (o *RoomTemplateDto) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *RoomTemplateDto) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *RoomTemplateDto) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *RoomTemplateDto) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetTags

`func (o *RoomTemplateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoomTemplateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoomTemplateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoomTemplateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *RoomTemplateDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *RoomTemplateDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetColor

`func (o *RoomTemplateDto) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *RoomTemplateDto) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *RoomTemplateDto) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *RoomTemplateDto) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *RoomTemplateDto) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *RoomTemplateDto) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *RoomTemplateDto) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *RoomTemplateDto) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *RoomTemplateDto) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *RoomTemplateDto) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *RoomTemplateDto) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *RoomTemplateDto) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil
### GetQuota

`func (o *RoomTemplateDto) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *RoomTemplateDto) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *RoomTemplateDto) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *RoomTemplateDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *RoomTemplateDto) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *RoomTemplateDto) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


