# LogoRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Light** | Pointer to **NullableString** | The image used on a light background, either as a `data:image/png;base64,...` payload - `png`, `jpg` and  `svg` are accepted - or as the name of a file already put in the temporary store. | [optional] 
**Dark** | Pointer to **NullableString** | The image used on a dark background, in the same two forms as `light`. It is only stored for the slots that  have a dark variant and is ignored for the favicon and the editor logos. | [optional] 

## Methods

### NewLogoRequestDto

`func NewLogoRequestDto() *LogoRequestDto`

NewLogoRequestDto instantiates a new LogoRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogoRequestDtoWithDefaults

`func NewLogoRequestDtoWithDefaults() *LogoRequestDto`

NewLogoRequestDtoWithDefaults instantiates a new LogoRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLight

`func (o *LogoRequestDto) GetLight() string`

GetLight returns the Light field if non-nil, zero value otherwise.

### GetLightOk

`func (o *LogoRequestDto) GetLightOk() (*string, bool)`

GetLightOk returns a tuple with the Light field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLight

`func (o *LogoRequestDto) SetLight(v string)`

SetLight sets Light field to given value.

### HasLight

`func (o *LogoRequestDto) HasLight() bool`

HasLight returns a boolean if a field has been set.

### SetLightNil

`func (o *LogoRequestDto) SetLightNil(b bool)`

 SetLightNil sets the value for Light to be an explicit nil

### UnsetLight
`func (o *LogoRequestDto) UnsetLight()`

UnsetLight ensures that no value is present for Light, not even an explicit nil
### GetDark

`func (o *LogoRequestDto) GetDark() string`

GetDark returns the Dark field if non-nil, zero value otherwise.

### GetDarkOk

`func (o *LogoRequestDto) GetDarkOk() (*string, bool)`

GetDarkOk returns a tuple with the Dark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDark

`func (o *LogoRequestDto) SetDark(v string)`

SetDark sets Dark field to given value.

### HasDark

`func (o *LogoRequestDto) HasDark() bool`

HasDark returns a boolean if a field has been set.

### SetDarkNil

`func (o *LogoRequestDto) SetDarkNil(b bool)`

 SetDarkNil sets the value for Dark to be an explicit nil

### UnsetDark
`func (o *LogoRequestDto) UnsetDark()`

UnsetDark ensures that no value is present for Dark, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


