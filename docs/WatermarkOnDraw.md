# WatermarkOnDraw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Width** | Pointer to **float64** | Defines the watermark width measured in millimeters. | [optional] 
**Height** | Pointer to **float64** | Defines the watermark height measured in millimeters. | [optional] 
**Margins** | Pointer to **[]int32** | Defines the watermark margins measured in millimeters. | [optional] 
**Fill** | Pointer to **NullableString** | Defines the watermark fill color. | [optional] 
**Rotate** | Pointer to **int32** | Defines the watermark rotation angle. | [optional] 
**Transparent** | Pointer to **float64** | Defines the watermark transparency percentage. | [optional] 
**Paragraphs** | Pointer to [**[]Paragraph**](Paragraph.md) | The list of paragraphs of the watermark. | [optional] 

## Methods

### NewWatermarkOnDraw

`func NewWatermarkOnDraw() *WatermarkOnDraw`

NewWatermarkOnDraw instantiates a new WatermarkOnDraw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkOnDrawWithDefaults

`func NewWatermarkOnDrawWithDefaults() *WatermarkOnDraw`

NewWatermarkOnDrawWithDefaults instantiates a new WatermarkOnDraw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidth

`func (o *WatermarkOnDraw) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *WatermarkOnDraw) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *WatermarkOnDraw) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *WatermarkOnDraw) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *WatermarkOnDraw) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *WatermarkOnDraw) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *WatermarkOnDraw) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *WatermarkOnDraw) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMargins

`func (o *WatermarkOnDraw) GetMargins() []int32`

GetMargins returns the Margins field if non-nil, zero value otherwise.

### GetMarginsOk

`func (o *WatermarkOnDraw) GetMarginsOk() (*[]int32, bool)`

GetMarginsOk returns a tuple with the Margins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargins

`func (o *WatermarkOnDraw) SetMargins(v []int32)`

SetMargins sets Margins field to given value.

### HasMargins

`func (o *WatermarkOnDraw) HasMargins() bool`

HasMargins returns a boolean if a field has been set.

### SetMarginsNil

`func (o *WatermarkOnDraw) SetMarginsNil(b bool)`

 SetMarginsNil sets the value for Margins to be an explicit nil

### UnsetMargins
`func (o *WatermarkOnDraw) UnsetMargins()`

UnsetMargins ensures that no value is present for Margins, not even an explicit nil
### GetFill

`func (o *WatermarkOnDraw) GetFill() string`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *WatermarkOnDraw) GetFillOk() (*string, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *WatermarkOnDraw) SetFill(v string)`

SetFill sets Fill field to given value.

### HasFill

`func (o *WatermarkOnDraw) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *WatermarkOnDraw) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *WatermarkOnDraw) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetRotate

`func (o *WatermarkOnDraw) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *WatermarkOnDraw) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *WatermarkOnDraw) SetRotate(v int32)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *WatermarkOnDraw) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetTransparent

`func (o *WatermarkOnDraw) GetTransparent() float64`

GetTransparent returns the Transparent field if non-nil, zero value otherwise.

### GetTransparentOk

`func (o *WatermarkOnDraw) GetTransparentOk() (*float64, bool)`

GetTransparentOk returns a tuple with the Transparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparent

`func (o *WatermarkOnDraw) SetTransparent(v float64)`

SetTransparent sets Transparent field to given value.

### HasTransparent

`func (o *WatermarkOnDraw) HasTransparent() bool`

HasTransparent returns a boolean if a field has been set.

### GetParagraphs

`func (o *WatermarkOnDraw) GetParagraphs() []Paragraph`

GetParagraphs returns the Paragraphs field if non-nil, zero value otherwise.

### GetParagraphsOk

`func (o *WatermarkOnDraw) GetParagraphsOk() (*[]Paragraph, bool)`

GetParagraphsOk returns a tuple with the Paragraphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParagraphs

`func (o *WatermarkOnDraw) SetParagraphs(v []Paragraph)`

SetParagraphs sets Paragraphs field to given value.

### HasParagraphs

`func (o *WatermarkOnDraw) HasParagraphs() bool`

HasParagraphs returns a boolean if a field has been set.

### SetParagraphsNil

`func (o *WatermarkOnDraw) SetParagraphsNil(b bool)`

 SetParagraphsNil sets the value for Paragraphs to be an explicit nil

### UnsetParagraphs
`func (o *WatermarkOnDraw) UnsetParagraphs()`

UnsetParagraphs ensures that no value is present for Paragraphs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


