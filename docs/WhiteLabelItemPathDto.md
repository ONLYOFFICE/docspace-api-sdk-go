# WhiteLabelItemPathDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Light** | Pointer to **NullableString** | The absolute URL of the image to render on a light background. It is filled in unless the request asked  for the dark theme alone with `isDark=true`, in which case only `dark` comes back. | [optional] 
**Dark** | Pointer to **NullableString** | The absolute URL of the image to render on a dark background. When both themes are asked for it comes back  empty for a slot that has no separate dark image, meaning the light one is to be used for both; when  `isDark=false` was passed it is left out entirely. | [optional] 

## Methods

### NewWhiteLabelItemPathDto

`func NewWhiteLabelItemPathDto() *WhiteLabelItemPathDto`

NewWhiteLabelItemPathDto instantiates a new WhiteLabelItemPathDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhiteLabelItemPathDtoWithDefaults

`func NewWhiteLabelItemPathDtoWithDefaults() *WhiteLabelItemPathDto`

NewWhiteLabelItemPathDtoWithDefaults instantiates a new WhiteLabelItemPathDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLight

`func (o *WhiteLabelItemPathDto) GetLight() string`

GetLight returns the Light field if non-nil, zero value otherwise.

### GetLightOk

`func (o *WhiteLabelItemPathDto) GetLightOk() (*string, bool)`

GetLightOk returns a tuple with the Light field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLight

`func (o *WhiteLabelItemPathDto) SetLight(v string)`

SetLight sets Light field to given value.

### HasLight

`func (o *WhiteLabelItemPathDto) HasLight() bool`

HasLight returns a boolean if a field has been set.

### SetLightNil

`func (o *WhiteLabelItemPathDto) SetLightNil(b bool)`

 SetLightNil sets the value for Light to be an explicit nil

### UnsetLight
`func (o *WhiteLabelItemPathDto) UnsetLight()`

UnsetLight ensures that no value is present for Light, not even an explicit nil
### GetDark

`func (o *WhiteLabelItemPathDto) GetDark() string`

GetDark returns the Dark field if non-nil, zero value otherwise.

### GetDarkOk

`func (o *WhiteLabelItemPathDto) GetDarkOk() (*string, bool)`

GetDarkOk returns a tuple with the Dark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDark

`func (o *WhiteLabelItemPathDto) SetDark(v string)`

SetDark sets Dark field to given value.

### HasDark

`func (o *WhiteLabelItemPathDto) HasDark() bool`

HasDark returns a boolean if a field has been set.

### SetDarkNil

`func (o *WhiteLabelItemPathDto) SetDarkNil(b bool)`

 SetDarkNil sets the value for Dark to be an explicit nil

### UnsetDark
`func (o *WhiteLabelItemPathDto) UnsetDark()`

UnsetDark ensures that no value is present for Dark, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


