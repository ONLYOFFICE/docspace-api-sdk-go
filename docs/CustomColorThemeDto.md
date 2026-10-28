# CustomColorThemeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The theme id; the built-in themes have the lowest ids. | [optional] 
**Name** | Pointer to **NullableString** | The theme name; empty for a custom theme. | [optional] 
**Main** | Pointer to [**ColorThemeColorsDto**](ColorThemeColorsDto.md) | The accent and button colours of the interface. | [optional] 
**Text** | Pointer to [**ColorThemeColorsDto**](ColorThemeColorsDto.md) | The colours of the text shown on the accent and on the buttons. | [optional] 

## Methods

### NewCustomColorThemeDto

`func NewCustomColorThemeDto() *CustomColorThemeDto`

NewCustomColorThemeDto instantiates a new CustomColorThemeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomColorThemeDtoWithDefaults

`func NewCustomColorThemeDtoWithDefaults() *CustomColorThemeDto`

NewCustomColorThemeDtoWithDefaults instantiates a new CustomColorThemeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomColorThemeDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomColorThemeDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomColorThemeDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomColorThemeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomColorThemeDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomColorThemeDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomColorThemeDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomColorThemeDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomColorThemeDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomColorThemeDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMain

`func (o *CustomColorThemeDto) GetMain() ColorThemeColorsDto`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *CustomColorThemeDto) GetMainOk() (*ColorThemeColorsDto, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *CustomColorThemeDto) SetMain(v ColorThemeColorsDto)`

SetMain sets Main field to given value.

### HasMain

`func (o *CustomColorThemeDto) HasMain() bool`

HasMain returns a boolean if a field has been set.

### GetText

`func (o *CustomColorThemeDto) GetText() ColorThemeColorsDto`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CustomColorThemeDto) GetTextOk() (*ColorThemeColorsDto, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CustomColorThemeDto) SetText(v ColorThemeColorsDto)`

SetText sets Text field to given value.

### HasText

`func (o *CustomColorThemeDto) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


