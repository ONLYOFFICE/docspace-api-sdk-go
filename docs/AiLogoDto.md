# AiLogoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Original** | **NullableString** | The original logo. | 
**Large** | **NullableString** | The large logo. | 
**Medium** | **NullableString** | The medium logo. | 
**Small** | **NullableString** | The small logo. | 
**Color** | Pointer to **NullableString** | The logo color. | [optional] 
**Cover** | Pointer to [**AiLogoCoverDto**](AiLogoCoverDto.md) | The logo cover. | [optional] 

## Methods

### NewAiLogoDto

`func NewAiLogoDto(original NullableString, large NullableString, medium NullableString, small NullableString, ) *AiLogoDto`

NewAiLogoDto instantiates a new AiLogoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiLogoDtoWithDefaults

`func NewAiLogoDtoWithDefaults() *AiLogoDto`

NewAiLogoDtoWithDefaults instantiates a new AiLogoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginal

`func (o *AiLogoDto) GetOriginal() string`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *AiLogoDto) GetOriginalOk() (*string, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *AiLogoDto) SetOriginal(v string)`

SetOriginal sets Original field to given value.


### SetOriginalNil

`func (o *AiLogoDto) SetOriginalNil(b bool)`

 SetOriginalNil sets the value for Original to be an explicit nil

### UnsetOriginal
`func (o *AiLogoDto) UnsetOriginal()`

UnsetOriginal ensures that no value is present for Original, not even an explicit nil
### GetLarge

`func (o *AiLogoDto) GetLarge() string`

GetLarge returns the Large field if non-nil, zero value otherwise.

### GetLargeOk

`func (o *AiLogoDto) GetLargeOk() (*string, bool)`

GetLargeOk returns a tuple with the Large field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLarge

`func (o *AiLogoDto) SetLarge(v string)`

SetLarge sets Large field to given value.


### SetLargeNil

`func (o *AiLogoDto) SetLargeNil(b bool)`

 SetLargeNil sets the value for Large to be an explicit nil

### UnsetLarge
`func (o *AiLogoDto) UnsetLarge()`

UnsetLarge ensures that no value is present for Large, not even an explicit nil
### GetMedium

`func (o *AiLogoDto) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *AiLogoDto) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *AiLogoDto) SetMedium(v string)`

SetMedium sets Medium field to given value.


### SetMediumNil

`func (o *AiLogoDto) SetMediumNil(b bool)`

 SetMediumNil sets the value for Medium to be an explicit nil

### UnsetMedium
`func (o *AiLogoDto) UnsetMedium()`

UnsetMedium ensures that no value is present for Medium, not even an explicit nil
### GetSmall

`func (o *AiLogoDto) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *AiLogoDto) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *AiLogoDto) SetSmall(v string)`

SetSmall sets Small field to given value.


### SetSmallNil

`func (o *AiLogoDto) SetSmallNil(b bool)`

 SetSmallNil sets the value for Small to be an explicit nil

### UnsetSmall
`func (o *AiLogoDto) UnsetSmall()`

UnsetSmall ensures that no value is present for Small, not even an explicit nil
### GetColor

`func (o *AiLogoDto) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *AiLogoDto) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *AiLogoDto) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *AiLogoDto) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *AiLogoDto) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *AiLogoDto) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *AiLogoDto) GetCover() AiLogoCoverDto`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *AiLogoDto) GetCoverOk() (*AiLogoCoverDto, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *AiLogoDto) SetCover(v AiLogoCoverDto)`

SetCover sets Cover field to given value.

### HasCover

`func (o *AiLogoDto) HasCover() bool`

HasCover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


