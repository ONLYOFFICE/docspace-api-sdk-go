# WhiteLabelRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogoText** | Pointer to **NullableString** | The wordmark printed next to or instead of a logo image, on the login page, in the editors and in  notification letters. An empty or blank value, and the built-in `ONLYOFFICE` itself, clear the setting rather  than store it. The text is not rendered into the logo images, which carry their own wordmark. | [optional] 
**Logo** | Pointer to [**[]ItemKeyValuePairStringLogoRequestDto**](ItemKeyValuePairStringLogoRequestDto.md) | The logo images to store, each entry naming a logo slot in its `key` - the numeric `type` published by  `GET api/2.0/settings/whitelabel/logos` - and carrying the two theme images in its value. A slot left out of  the list keeps the image it has, so this is a partial update rather than a replacement of the whole branding.  Saving the login-page slot also rebuilds the notification logo from it. | [optional] 

## Methods

### NewWhiteLabelRequestDto

`func NewWhiteLabelRequestDto() *WhiteLabelRequestDto`

NewWhiteLabelRequestDto instantiates a new WhiteLabelRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhiteLabelRequestDtoWithDefaults

`func NewWhiteLabelRequestDtoWithDefaults() *WhiteLabelRequestDto`

NewWhiteLabelRequestDtoWithDefaults instantiates a new WhiteLabelRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogoText

`func (o *WhiteLabelRequestDto) GetLogoText() string`

GetLogoText returns the LogoText field if non-nil, zero value otherwise.

### GetLogoTextOk

`func (o *WhiteLabelRequestDto) GetLogoTextOk() (*string, bool)`

GetLogoTextOk returns a tuple with the LogoText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoText

`func (o *WhiteLabelRequestDto) SetLogoText(v string)`

SetLogoText sets LogoText field to given value.

### HasLogoText

`func (o *WhiteLabelRequestDto) HasLogoText() bool`

HasLogoText returns a boolean if a field has been set.

### SetLogoTextNil

`func (o *WhiteLabelRequestDto) SetLogoTextNil(b bool)`

 SetLogoTextNil sets the value for LogoText to be an explicit nil

### UnsetLogoText
`func (o *WhiteLabelRequestDto) UnsetLogoText()`

UnsetLogoText ensures that no value is present for LogoText, not even an explicit nil
### GetLogo

`func (o *WhiteLabelRequestDto) GetLogo() []ItemKeyValuePairStringLogoRequestDto`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *WhiteLabelRequestDto) GetLogoOk() (*[]ItemKeyValuePairStringLogoRequestDto, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *WhiteLabelRequestDto) SetLogo(v []ItemKeyValuePairStringLogoRequestDto)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *WhiteLabelRequestDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *WhiteLabelRequestDto) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *WhiteLabelRequestDto) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


