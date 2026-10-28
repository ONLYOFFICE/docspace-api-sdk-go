# CustomColorThemeRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the custom theme to replace, or an id no stored theme has to add a new one. | [optional] 
**Name** | Pointer to **NullableString** | Accepted for compatibility with earlier clients and not read: a custom theme is always stored without a name. | [optional] 
**Main** | Pointer to [**ColorThemeColorsRequestDto**](ColorThemeColorsRequestDto.md) | The accent and button colours of the interface. Left out, a stored theme keeps its own. | [optional] 
**Text** | Pointer to [**ColorThemeColorsRequestDto**](ColorThemeColorsRequestDto.md) | The colours of the text shown on the accent and on the buttons. Left out, a stored theme keeps its own. | [optional] 

## Methods

### NewCustomColorThemeRequestDto

`func NewCustomColorThemeRequestDto() *CustomColorThemeRequestDto`

NewCustomColorThemeRequestDto instantiates a new CustomColorThemeRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomColorThemeRequestDtoWithDefaults

`func NewCustomColorThemeRequestDtoWithDefaults() *CustomColorThemeRequestDto`

NewCustomColorThemeRequestDtoWithDefaults instantiates a new CustomColorThemeRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomColorThemeRequestDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomColorThemeRequestDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomColorThemeRequestDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomColorThemeRequestDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomColorThemeRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomColorThemeRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomColorThemeRequestDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomColorThemeRequestDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomColorThemeRequestDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomColorThemeRequestDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMain

`func (o *CustomColorThemeRequestDto) GetMain() ColorThemeColorsRequestDto`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *CustomColorThemeRequestDto) GetMainOk() (*ColorThemeColorsRequestDto, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *CustomColorThemeRequestDto) SetMain(v ColorThemeColorsRequestDto)`

SetMain sets Main field to given value.

### HasMain

`func (o *CustomColorThemeRequestDto) HasMain() bool`

HasMain returns a boolean if a field has been set.

### GetText

`func (o *CustomColorThemeRequestDto) GetText() ColorThemeColorsRequestDto`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CustomColorThemeRequestDto) GetTextOk() (*ColorThemeColorsRequestDto, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CustomColorThemeRequestDto) SetText(v ColorThemeColorsRequestDto)`

SetText sets Text field to given value.

### HasText

`func (o *CustomColorThemeRequestDto) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


