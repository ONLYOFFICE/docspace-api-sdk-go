# CustomColorThemesSettingsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The custom color theme ID. | [optional] 
**Name** | Pointer to **NullableString** | The custom color theme name. | [optional] 
**Main** | Pointer to [**CustomColorThemesSettingsColorItem**](CustomColorThemesSettingsColorItem.md) |  | [optional] 
**Text** | Pointer to [**CustomColorThemesSettingsColorItem**](CustomColorThemesSettingsColorItem.md) |  | [optional] 

## Methods

### NewCustomColorThemesSettingsItem

`func NewCustomColorThemesSettingsItem() *CustomColorThemesSettingsItem`

NewCustomColorThemesSettingsItem instantiates a new CustomColorThemesSettingsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomColorThemesSettingsItemWithDefaults

`func NewCustomColorThemesSettingsItemWithDefaults() *CustomColorThemesSettingsItem`

NewCustomColorThemesSettingsItemWithDefaults instantiates a new CustomColorThemesSettingsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomColorThemesSettingsItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomColorThemesSettingsItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomColorThemesSettingsItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomColorThemesSettingsItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomColorThemesSettingsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomColorThemesSettingsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomColorThemesSettingsItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomColorThemesSettingsItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomColorThemesSettingsItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomColorThemesSettingsItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMain

`func (o *CustomColorThemesSettingsItem) GetMain() CustomColorThemesSettingsColorItem`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *CustomColorThemesSettingsItem) GetMainOk() (*CustomColorThemesSettingsColorItem, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *CustomColorThemesSettingsItem) SetMain(v CustomColorThemesSettingsColorItem)`

SetMain sets Main field to given value.

### HasMain

`func (o *CustomColorThemesSettingsItem) HasMain() bool`

HasMain returns a boolean if a field has been set.

### GetText

`func (o *CustomColorThemesSettingsItem) GetText() CustomColorThemesSettingsColorItem`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CustomColorThemesSettingsItem) GetTextOk() (*CustomColorThemesSettingsColorItem, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CustomColorThemesSettingsItem) SetText(v CustomColorThemesSettingsColorItem)`

SetText sets Text field to given value.

### HasText

`func (o *CustomColorThemesSettingsItem) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


