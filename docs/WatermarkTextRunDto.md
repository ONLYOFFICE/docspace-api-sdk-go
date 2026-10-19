# WatermarkTextRunDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fill** | Pointer to **[]int32** | The fill color of the text run in RGB format. | [optional] 
**Text** | Pointer to **NullableString** | The run text. | [optional] 
**FontSize** | Pointer to **NullableString** | The font size of the text run in points. | [optional] 

## Methods

### NewWatermarkTextRunDto

`func NewWatermarkTextRunDto() *WatermarkTextRunDto`

NewWatermarkTextRunDto instantiates a new WatermarkTextRunDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkTextRunDtoWithDefaults

`func NewWatermarkTextRunDtoWithDefaults() *WatermarkTextRunDto`

NewWatermarkTextRunDtoWithDefaults instantiates a new WatermarkTextRunDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFill

`func (o *WatermarkTextRunDto) GetFill() []int32`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *WatermarkTextRunDto) GetFillOk() (*[]int32, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *WatermarkTextRunDto) SetFill(v []int32)`

SetFill sets Fill field to given value.

### HasFill

`func (o *WatermarkTextRunDto) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *WatermarkTextRunDto) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *WatermarkTextRunDto) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetText

`func (o *WatermarkTextRunDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WatermarkTextRunDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WatermarkTextRunDto) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WatermarkTextRunDto) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *WatermarkTextRunDto) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *WatermarkTextRunDto) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetFontSize

`func (o *WatermarkTextRunDto) GetFontSize() string`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *WatermarkTextRunDto) GetFontSizeOk() (*string, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *WatermarkTextRunDto) SetFontSize(v string)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *WatermarkTextRunDto) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### SetFontSizeNil

`func (o *WatermarkTextRunDto) SetFontSizeNil(b bool)`

 SetFontSizeNil sets the value for FontSize to be an explicit nil

### UnsetFontSize
`func (o *WatermarkTextRunDto) UnsetFontSize()`

UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


