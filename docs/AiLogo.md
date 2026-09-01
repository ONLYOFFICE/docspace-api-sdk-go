# AiLogo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Original** | **NullableString** | The original logo. | 
**Large** | **NullableString** | The large logo. | 
**Medium** | **NullableString** | The medium logo. | 
**Small** | **NullableString** | The small logo. | 
**Color** | Pointer to **NullableString** | The logo color. | [optional] 
**Cover** | Pointer to [**AiLogoCover**](AiLogoCover.md) | The logo cover. | [optional] 

## Methods

### NewAiLogo

`func NewAiLogo(original NullableString, large NullableString, medium NullableString, small NullableString, ) *AiLogo`

NewAiLogo instantiates a new AiLogo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiLogoWithDefaults

`func NewAiLogoWithDefaults() *AiLogo`

NewAiLogoWithDefaults instantiates a new AiLogo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginal

`func (o *AiLogo) GetOriginal() string`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *AiLogo) GetOriginalOk() (*string, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *AiLogo) SetOriginal(v string)`

SetOriginal sets Original field to given value.


### SetOriginalNil

`func (o *AiLogo) SetOriginalNil(b bool)`

 SetOriginalNil sets the value for Original to be an explicit nil

### UnsetOriginal
`func (o *AiLogo) UnsetOriginal()`

UnsetOriginal ensures that no value is present for Original, not even an explicit nil
### GetLarge

`func (o *AiLogo) GetLarge() string`

GetLarge returns the Large field if non-nil, zero value otherwise.

### GetLargeOk

`func (o *AiLogo) GetLargeOk() (*string, bool)`

GetLargeOk returns a tuple with the Large field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLarge

`func (o *AiLogo) SetLarge(v string)`

SetLarge sets Large field to given value.


### SetLargeNil

`func (o *AiLogo) SetLargeNil(b bool)`

 SetLargeNil sets the value for Large to be an explicit nil

### UnsetLarge
`func (o *AiLogo) UnsetLarge()`

UnsetLarge ensures that no value is present for Large, not even an explicit nil
### GetMedium

`func (o *AiLogo) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *AiLogo) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *AiLogo) SetMedium(v string)`

SetMedium sets Medium field to given value.


### SetMediumNil

`func (o *AiLogo) SetMediumNil(b bool)`

 SetMediumNil sets the value for Medium to be an explicit nil

### UnsetMedium
`func (o *AiLogo) UnsetMedium()`

UnsetMedium ensures that no value is present for Medium, not even an explicit nil
### GetSmall

`func (o *AiLogo) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *AiLogo) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *AiLogo) SetSmall(v string)`

SetSmall sets Small field to given value.


### SetSmallNil

`func (o *AiLogo) SetSmallNil(b bool)`

 SetSmallNil sets the value for Small to be an explicit nil

### UnsetSmall
`func (o *AiLogo) UnsetSmall()`

UnsetSmall ensures that no value is present for Small, not even an explicit nil
### GetColor

`func (o *AiLogo) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *AiLogo) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *AiLogo) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *AiLogo) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *AiLogo) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *AiLogo) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *AiLogo) GetCover() AiLogoCover`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *AiLogo) GetCoverOk() (*AiLogoCover, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *AiLogo) SetCover(v AiLogoCover)`

SetCover sets Cover field to given value.

### HasCover

`func (o *AiLogo) HasCover() bool`

HasCover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


