# AiEmployeeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The user ID. | [optional] 
**DisplayName** | Pointer to **NullableString** | The HTML-encoded user's display name formatted according to the default format for the current culture. | [optional] 
**Avatar** | Pointer to **NullableString** | The user avatar. | [optional] 
**AvatarOriginal** | Pointer to **NullableString** | The user original size avatar. | [optional] 
**AvatarMax** | Pointer to **NullableString** | The user maximum size avatar. | [optional] 
**AvatarMedium** | Pointer to **NullableString** | The user medium size avatar. | [optional] 
**AvatarSmall** | Pointer to **NullableString** | The user small size avatar. | [optional] 
**ProfileUrl** | Pointer to **NullableString** | The user profile URL. | [optional] 
**HasAvatar** | Pointer to **bool** | Specifies if the user has an avatar or not. | [optional] 
**IsAnonim** | Pointer to **bool** | Specifies if the user is anonymous or not. | [optional] 

## Methods

### NewAiEmployeeDto

`func NewAiEmployeeDto() *AiEmployeeDto`

NewAiEmployeeDto instantiates a new AiEmployeeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEmployeeDtoWithDefaults

`func NewAiEmployeeDtoWithDefaults() *AiEmployeeDto`

NewAiEmployeeDtoWithDefaults instantiates a new AiEmployeeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiEmployeeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiEmployeeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiEmployeeDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiEmployeeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AiEmployeeDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AiEmployeeDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AiEmployeeDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AiEmployeeDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AiEmployeeDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AiEmployeeDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetAvatar

`func (o *AiEmployeeDto) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *AiEmployeeDto) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *AiEmployeeDto) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *AiEmployeeDto) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *AiEmployeeDto) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *AiEmployeeDto) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarOriginal

`func (o *AiEmployeeDto) GetAvatarOriginal() string`

GetAvatarOriginal returns the AvatarOriginal field if non-nil, zero value otherwise.

### GetAvatarOriginalOk

`func (o *AiEmployeeDto) GetAvatarOriginalOk() (*string, bool)`

GetAvatarOriginalOk returns a tuple with the AvatarOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarOriginal

`func (o *AiEmployeeDto) SetAvatarOriginal(v string)`

SetAvatarOriginal sets AvatarOriginal field to given value.

### HasAvatarOriginal

`func (o *AiEmployeeDto) HasAvatarOriginal() bool`

HasAvatarOriginal returns a boolean if a field has been set.

### SetAvatarOriginalNil

`func (o *AiEmployeeDto) SetAvatarOriginalNil(b bool)`

 SetAvatarOriginalNil sets the value for AvatarOriginal to be an explicit nil

### UnsetAvatarOriginal
`func (o *AiEmployeeDto) UnsetAvatarOriginal()`

UnsetAvatarOriginal ensures that no value is present for AvatarOriginal, not even an explicit nil
### GetAvatarMax

`func (o *AiEmployeeDto) GetAvatarMax() string`

GetAvatarMax returns the AvatarMax field if non-nil, zero value otherwise.

### GetAvatarMaxOk

`func (o *AiEmployeeDto) GetAvatarMaxOk() (*string, bool)`

GetAvatarMaxOk returns a tuple with the AvatarMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarMax

`func (o *AiEmployeeDto) SetAvatarMax(v string)`

SetAvatarMax sets AvatarMax field to given value.

### HasAvatarMax

`func (o *AiEmployeeDto) HasAvatarMax() bool`

HasAvatarMax returns a boolean if a field has been set.

### SetAvatarMaxNil

`func (o *AiEmployeeDto) SetAvatarMaxNil(b bool)`

 SetAvatarMaxNil sets the value for AvatarMax to be an explicit nil

### UnsetAvatarMax
`func (o *AiEmployeeDto) UnsetAvatarMax()`

UnsetAvatarMax ensures that no value is present for AvatarMax, not even an explicit nil
### GetAvatarMedium

`func (o *AiEmployeeDto) GetAvatarMedium() string`

GetAvatarMedium returns the AvatarMedium field if non-nil, zero value otherwise.

### GetAvatarMediumOk

`func (o *AiEmployeeDto) GetAvatarMediumOk() (*string, bool)`

GetAvatarMediumOk returns a tuple with the AvatarMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarMedium

`func (o *AiEmployeeDto) SetAvatarMedium(v string)`

SetAvatarMedium sets AvatarMedium field to given value.

### HasAvatarMedium

`func (o *AiEmployeeDto) HasAvatarMedium() bool`

HasAvatarMedium returns a boolean if a field has been set.

### SetAvatarMediumNil

`func (o *AiEmployeeDto) SetAvatarMediumNil(b bool)`

 SetAvatarMediumNil sets the value for AvatarMedium to be an explicit nil

### UnsetAvatarMedium
`func (o *AiEmployeeDto) UnsetAvatarMedium()`

UnsetAvatarMedium ensures that no value is present for AvatarMedium, not even an explicit nil
### GetAvatarSmall

`func (o *AiEmployeeDto) GetAvatarSmall() string`

GetAvatarSmall returns the AvatarSmall field if non-nil, zero value otherwise.

### GetAvatarSmallOk

`func (o *AiEmployeeDto) GetAvatarSmallOk() (*string, bool)`

GetAvatarSmallOk returns a tuple with the AvatarSmall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSmall

`func (o *AiEmployeeDto) SetAvatarSmall(v string)`

SetAvatarSmall sets AvatarSmall field to given value.

### HasAvatarSmall

`func (o *AiEmployeeDto) HasAvatarSmall() bool`

HasAvatarSmall returns a boolean if a field has been set.

### SetAvatarSmallNil

`func (o *AiEmployeeDto) SetAvatarSmallNil(b bool)`

 SetAvatarSmallNil sets the value for AvatarSmall to be an explicit nil

### UnsetAvatarSmall
`func (o *AiEmployeeDto) UnsetAvatarSmall()`

UnsetAvatarSmall ensures that no value is present for AvatarSmall, not even an explicit nil
### GetProfileUrl

`func (o *AiEmployeeDto) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *AiEmployeeDto) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *AiEmployeeDto) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *AiEmployeeDto) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *AiEmployeeDto) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *AiEmployeeDto) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetHasAvatar

`func (o *AiEmployeeDto) GetHasAvatar() bool`

GetHasAvatar returns the HasAvatar field if non-nil, zero value otherwise.

### GetHasAvatarOk

`func (o *AiEmployeeDto) GetHasAvatarOk() (*bool, bool)`

GetHasAvatarOk returns a tuple with the HasAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAvatar

`func (o *AiEmployeeDto) SetHasAvatar(v bool)`

SetHasAvatar sets HasAvatar field to given value.

### HasHasAvatar

`func (o *AiEmployeeDto) HasHasAvatar() bool`

HasHasAvatar returns a boolean if a field has been set.

### GetIsAnonim

`func (o *AiEmployeeDto) GetIsAnonim() bool`

GetIsAnonim returns the IsAnonim field if non-nil, zero value otherwise.

### GetIsAnonimOk

`func (o *AiEmployeeDto) GetIsAnonimOk() (*bool, bool)`

GetIsAnonimOk returns a tuple with the IsAnonim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonim

`func (o *AiEmployeeDto) SetIsAnonim(v bool)`

SetIsAnonim sets IsAnonim field to given value.

### HasIsAnonim

`func (o *AiEmployeeDto) HasIsAnonim() bool`

HasIsAnonim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


