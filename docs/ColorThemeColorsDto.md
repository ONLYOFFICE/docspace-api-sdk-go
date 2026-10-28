# ColorThemeColorsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accent** | Pointer to **NullableString** | The accent colour, as a hex code. | [optional] 
**Buttons** | Pointer to **NullableString** | The button colour, as a hex code. | [optional] 

## Methods

### NewColorThemeColorsDto

`func NewColorThemeColorsDto() *ColorThemeColorsDto`

NewColorThemeColorsDto instantiates a new ColorThemeColorsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColorThemeColorsDtoWithDefaults

`func NewColorThemeColorsDtoWithDefaults() *ColorThemeColorsDto`

NewColorThemeColorsDtoWithDefaults instantiates a new ColorThemeColorsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccent

`func (o *ColorThemeColorsDto) GetAccent() string`

GetAccent returns the Accent field if non-nil, zero value otherwise.

### GetAccentOk

`func (o *ColorThemeColorsDto) GetAccentOk() (*string, bool)`

GetAccentOk returns a tuple with the Accent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccent

`func (o *ColorThemeColorsDto) SetAccent(v string)`

SetAccent sets Accent field to given value.

### HasAccent

`func (o *ColorThemeColorsDto) HasAccent() bool`

HasAccent returns a boolean if a field has been set.

### SetAccentNil

`func (o *ColorThemeColorsDto) SetAccentNil(b bool)`

 SetAccentNil sets the value for Accent to be an explicit nil

### UnsetAccent
`func (o *ColorThemeColorsDto) UnsetAccent()`

UnsetAccent ensures that no value is present for Accent, not even an explicit nil
### GetButtons

`func (o *ColorThemeColorsDto) GetButtons() string`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *ColorThemeColorsDto) GetButtonsOk() (*string, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *ColorThemeColorsDto) SetButtons(v string)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *ColorThemeColorsDto) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### SetButtonsNil

`func (o *ColorThemeColorsDto) SetButtonsNil(b bool)`

 SetButtonsNil sets the value for Buttons to be an explicit nil

### UnsetButtons
`func (o *ColorThemeColorsDto) UnsetButtons()`

UnsetButtons ensures that no value is present for Buttons, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


