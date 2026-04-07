# CustomColorThemesSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Theme** | Pointer to [**CustomColorThemesSettingsItem**](CustomColorThemesSettingsItem.md) |  | [optional] 
**Selected** | Pointer to **NullableInt32** | Specifies the optional value indicating the selected custom color theme. | [optional] 

## Methods

### NewCustomColorThemesSettingsRequestsDto

`func NewCustomColorThemesSettingsRequestsDto() *CustomColorThemesSettingsRequestsDto`

NewCustomColorThemesSettingsRequestsDto instantiates a new CustomColorThemesSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomColorThemesSettingsRequestsDtoWithDefaults

`func NewCustomColorThemesSettingsRequestsDtoWithDefaults() *CustomColorThemesSettingsRequestsDto`

NewCustomColorThemesSettingsRequestsDtoWithDefaults instantiates a new CustomColorThemesSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTheme

`func (o *CustomColorThemesSettingsRequestsDto) GetTheme() CustomColorThemesSettingsItem`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *CustomColorThemesSettingsRequestsDto) GetThemeOk() (*CustomColorThemesSettingsItem, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *CustomColorThemesSettingsRequestsDto) SetTheme(v CustomColorThemesSettingsItem)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *CustomColorThemesSettingsRequestsDto) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetSelected

`func (o *CustomColorThemesSettingsRequestsDto) GetSelected() int32`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *CustomColorThemesSettingsRequestsDto) GetSelectedOk() (*int32, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *CustomColorThemesSettingsRequestsDto) SetSelected(v int32)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *CustomColorThemesSettingsRequestsDto) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### SetSelectedNil

`func (o *CustomColorThemesSettingsRequestsDto) SetSelectedNil(b bool)`

 SetSelectedNil sets the value for Selected to be an explicit nil

### UnsetSelected
`func (o *CustomColorThemesSettingsRequestsDto) UnsetSelected()`

UnsetSelected ensures that no value is present for Selected, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


