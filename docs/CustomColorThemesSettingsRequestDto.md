# CustomColorThemesSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Theme** | Pointer to [**CustomColorThemeRequestDto**](CustomColorThemeRequestDto.md) | The theme to store, with its accent and button colours for the interface and for the text on it. An `id` that  matches a stored custom theme replaces it, an unknown `id` appends a new one, and an `id` belonging to a  built-in theme is treated as a request for a new custom theme rather than overwriting the built-in one. Once  the plan limit on custom themes is reached a new theme is silently not added, so compare the returned themes  against `limit` instead of assuming it was saved. Leave it out to change only the selection. | [optional] 
**Selected** | Pointer to **NullableInt32** | The theme the whole portal switches to, by theme ID. An ID matching no stored theme is ignored rather than  refused, and leaving it out keeps the selection as it is. | [optional] 

## Methods

### NewCustomColorThemesSettingsRequestDto

`func NewCustomColorThemesSettingsRequestDto() *CustomColorThemesSettingsRequestDto`

NewCustomColorThemesSettingsRequestDto instantiates a new CustomColorThemesSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomColorThemesSettingsRequestDtoWithDefaults

`func NewCustomColorThemesSettingsRequestDtoWithDefaults() *CustomColorThemesSettingsRequestDto`

NewCustomColorThemesSettingsRequestDtoWithDefaults instantiates a new CustomColorThemesSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTheme

`func (o *CustomColorThemesSettingsRequestDto) GetTheme() CustomColorThemeRequestDto`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *CustomColorThemesSettingsRequestDto) GetThemeOk() (*CustomColorThemeRequestDto, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *CustomColorThemesSettingsRequestDto) SetTheme(v CustomColorThemeRequestDto)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *CustomColorThemesSettingsRequestDto) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetSelected

`func (o *CustomColorThemesSettingsRequestDto) GetSelected() int32`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *CustomColorThemesSettingsRequestDto) GetSelectedOk() (*int32, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *CustomColorThemesSettingsRequestDto) SetSelected(v int32)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *CustomColorThemesSettingsRequestDto) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### SetSelectedNil

`func (o *CustomColorThemesSettingsRequestDto) SetSelectedNil(b bool)`

 SetSelectedNil sets the value for Selected to be an explicit nil

### UnsetSelected
`func (o *CustomColorThemesSettingsRequestDto) UnsetSelected()`

UnsetSelected ensures that no value is present for Selected, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


