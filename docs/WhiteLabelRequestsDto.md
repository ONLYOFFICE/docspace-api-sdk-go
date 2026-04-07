# WhiteLabelRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogoText** | Pointer to **NullableString** | The text to display alongside or in place of the logo. | [optional] 
**Logo** | Pointer to [**[]ItemKeyValuePairStringLogoRequestsDto**](ItemKeyValuePairStringLogoRequestsDto.md) | The white label tenant IDs with their logos (light or dark). | [optional] 

## Methods

### NewWhiteLabelRequestsDto

`func NewWhiteLabelRequestsDto() *WhiteLabelRequestsDto`

NewWhiteLabelRequestsDto instantiates a new WhiteLabelRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhiteLabelRequestsDtoWithDefaults

`func NewWhiteLabelRequestsDtoWithDefaults() *WhiteLabelRequestsDto`

NewWhiteLabelRequestsDtoWithDefaults instantiates a new WhiteLabelRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogoText

`func (o *WhiteLabelRequestsDto) GetLogoText() string`

GetLogoText returns the LogoText field if non-nil, zero value otherwise.

### GetLogoTextOk

`func (o *WhiteLabelRequestsDto) GetLogoTextOk() (*string, bool)`

GetLogoTextOk returns a tuple with the LogoText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoText

`func (o *WhiteLabelRequestsDto) SetLogoText(v string)`

SetLogoText sets LogoText field to given value.

### HasLogoText

`func (o *WhiteLabelRequestsDto) HasLogoText() bool`

HasLogoText returns a boolean if a field has been set.

### SetLogoTextNil

`func (o *WhiteLabelRequestsDto) SetLogoTextNil(b bool)`

 SetLogoTextNil sets the value for LogoText to be an explicit nil

### UnsetLogoText
`func (o *WhiteLabelRequestsDto) UnsetLogoText()`

UnsetLogoText ensures that no value is present for LogoText, not even an explicit nil
### GetLogo

`func (o *WhiteLabelRequestsDto) GetLogo() []ItemKeyValuePairStringLogoRequestsDto`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *WhiteLabelRequestsDto) GetLogoOk() (*[]ItemKeyValuePairStringLogoRequestsDto, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *WhiteLabelRequestsDto) SetLogo(v []ItemKeyValuePairStringLogoRequestsDto)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *WhiteLabelRequestsDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *WhiteLabelRequestsDto) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *WhiteLabelRequestsDto) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


