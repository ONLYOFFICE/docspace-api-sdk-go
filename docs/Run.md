# Run

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fill** | Pointer to **[]int32** | The fill color of the text run in RGB format. | [optional] 
**Text** | Pointer to **NullableString** | The run text. | [optional] 
**FontSize** | Pointer to **NullableString** | The font size of the text run in points. | [optional] 

## Methods

### NewRun

`func NewRun() *Run`

NewRun instantiates a new Run object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunWithDefaults

`func NewRunWithDefaults() *Run`

NewRunWithDefaults instantiates a new Run object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFill

`func (o *Run) GetFill() []int32`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *Run) GetFillOk() (*[]int32, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *Run) SetFill(v []int32)`

SetFill sets Fill field to given value.

### HasFill

`func (o *Run) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *Run) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *Run) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetText

`func (o *Run) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Run) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Run) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Run) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *Run) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *Run) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetFontSize

`func (o *Run) GetFontSize() string`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *Run) GetFontSizeOk() (*string, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *Run) SetFontSize(v string)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *Run) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### SetFontSizeNil

`func (o *Run) SetFontSizeNil(b bool)`

 SetFontSizeNil sets the value for FontSize to be an explicit nil

### UnsetFontSize
`func (o *Run) UnsetFontSize()`

UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


