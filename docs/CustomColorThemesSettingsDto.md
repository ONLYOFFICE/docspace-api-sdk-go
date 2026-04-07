# CustomColorThemesSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Themes** | Pointer to [**[]CustomColorThemesSettingsItem**](CustomColorThemesSettingsItem.md) | The list of the custom color themes. | [optional] 
**Selected** | Pointer to **int32** | Specifies whether the custom color theme is selected. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of the custom color themes. | [optional] 

## Methods

### NewCustomColorThemesSettingsDto

`func NewCustomColorThemesSettingsDto() *CustomColorThemesSettingsDto`

NewCustomColorThemesSettingsDto instantiates a new CustomColorThemesSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomColorThemesSettingsDtoWithDefaults

`func NewCustomColorThemesSettingsDtoWithDefaults() *CustomColorThemesSettingsDto`

NewCustomColorThemesSettingsDtoWithDefaults instantiates a new CustomColorThemesSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThemes

`func (o *CustomColorThemesSettingsDto) GetThemes() []CustomColorThemesSettingsItem`

GetThemes returns the Themes field if non-nil, zero value otherwise.

### GetThemesOk

`func (o *CustomColorThemesSettingsDto) GetThemesOk() (*[]CustomColorThemesSettingsItem, bool)`

GetThemesOk returns a tuple with the Themes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemes

`func (o *CustomColorThemesSettingsDto) SetThemes(v []CustomColorThemesSettingsItem)`

SetThemes sets Themes field to given value.

### HasThemes

`func (o *CustomColorThemesSettingsDto) HasThemes() bool`

HasThemes returns a boolean if a field has been set.

### SetThemesNil

`func (o *CustomColorThemesSettingsDto) SetThemesNil(b bool)`

 SetThemesNil sets the value for Themes to be an explicit nil

### UnsetThemes
`func (o *CustomColorThemesSettingsDto) UnsetThemes()`

UnsetThemes ensures that no value is present for Themes, not even an explicit nil
### GetSelected

`func (o *CustomColorThemesSettingsDto) GetSelected() int32`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *CustomColorThemesSettingsDto) GetSelectedOk() (*int32, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *CustomColorThemesSettingsDto) SetSelected(v int32)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *CustomColorThemesSettingsDto) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### GetLimit

`func (o *CustomColorThemesSettingsDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CustomColorThemesSettingsDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CustomColorThemesSettingsDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CustomColorThemesSettingsDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


