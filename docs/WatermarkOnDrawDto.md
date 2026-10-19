# WatermarkOnDrawDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Width** | Pointer to **float64** | Defines the watermark width measured in millimeters. | [optional] 
**Height** | Pointer to **float64** | Defines the watermark height measured in millimeters. | [optional] 
**Margins** | Pointer to **[]int32** | Defines the watermark margins measured in millimeters. | [optional] 
**Fill** | Pointer to **NullableString** | Defines the watermark fill color. | [optional] 
**Rotate** | Pointer to **int32** | Defines the watermark rotation angle. | [optional] 
**Transparent** | Pointer to **float64** | Defines the watermark transparency percentage. | [optional] 
**Paragraphs** | Pointer to [**[]WatermarkParagraphDto**](WatermarkParagraphDto.md) | The list of paragraphs of the watermark. | [optional] 

## Methods

### NewWatermarkOnDrawDto

`func NewWatermarkOnDrawDto() *WatermarkOnDrawDto`

NewWatermarkOnDrawDto instantiates a new WatermarkOnDrawDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkOnDrawDtoWithDefaults

`func NewWatermarkOnDrawDtoWithDefaults() *WatermarkOnDrawDto`

NewWatermarkOnDrawDtoWithDefaults instantiates a new WatermarkOnDrawDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidth

`func (o *WatermarkOnDrawDto) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *WatermarkOnDrawDto) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *WatermarkOnDrawDto) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *WatermarkOnDrawDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *WatermarkOnDrawDto) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *WatermarkOnDrawDto) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *WatermarkOnDrawDto) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *WatermarkOnDrawDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMargins

`func (o *WatermarkOnDrawDto) GetMargins() []int32`

GetMargins returns the Margins field if non-nil, zero value otherwise.

### GetMarginsOk

`func (o *WatermarkOnDrawDto) GetMarginsOk() (*[]int32, bool)`

GetMarginsOk returns a tuple with the Margins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargins

`func (o *WatermarkOnDrawDto) SetMargins(v []int32)`

SetMargins sets Margins field to given value.

### HasMargins

`func (o *WatermarkOnDrawDto) HasMargins() bool`

HasMargins returns a boolean if a field has been set.

### SetMarginsNil

`func (o *WatermarkOnDrawDto) SetMarginsNil(b bool)`

 SetMarginsNil sets the value for Margins to be an explicit nil

### UnsetMargins
`func (o *WatermarkOnDrawDto) UnsetMargins()`

UnsetMargins ensures that no value is present for Margins, not even an explicit nil
### GetFill

`func (o *WatermarkOnDrawDto) GetFill() string`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *WatermarkOnDrawDto) GetFillOk() (*string, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *WatermarkOnDrawDto) SetFill(v string)`

SetFill sets Fill field to given value.

### HasFill

`func (o *WatermarkOnDrawDto) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *WatermarkOnDrawDto) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *WatermarkOnDrawDto) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetRotate

`func (o *WatermarkOnDrawDto) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *WatermarkOnDrawDto) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *WatermarkOnDrawDto) SetRotate(v int32)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *WatermarkOnDrawDto) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetTransparent

`func (o *WatermarkOnDrawDto) GetTransparent() float64`

GetTransparent returns the Transparent field if non-nil, zero value otherwise.

### GetTransparentOk

`func (o *WatermarkOnDrawDto) GetTransparentOk() (*float64, bool)`

GetTransparentOk returns a tuple with the Transparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparent

`func (o *WatermarkOnDrawDto) SetTransparent(v float64)`

SetTransparent sets Transparent field to given value.

### HasTransparent

`func (o *WatermarkOnDrawDto) HasTransparent() bool`

HasTransparent returns a boolean if a field has been set.

### GetParagraphs

`func (o *WatermarkOnDrawDto) GetParagraphs() []WatermarkParagraphDto`

GetParagraphs returns the Paragraphs field if non-nil, zero value otherwise.

### GetParagraphsOk

`func (o *WatermarkOnDrawDto) GetParagraphsOk() (*[]WatermarkParagraphDto, bool)`

GetParagraphsOk returns a tuple with the Paragraphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParagraphs

`func (o *WatermarkOnDrawDto) SetParagraphs(v []WatermarkParagraphDto)`

SetParagraphs sets Paragraphs field to given value.

### HasParagraphs

`func (o *WatermarkOnDrawDto) HasParagraphs() bool`

HasParagraphs returns a boolean if a field has been set.

### SetParagraphsNil

`func (o *WatermarkOnDrawDto) SetParagraphsNil(b bool)`

 SetParagraphsNil sets the value for Paragraphs to be an explicit nil

### UnsetParagraphs
`func (o *WatermarkOnDrawDto) UnsetParagraphs()`

UnsetParagraphs ensures that no value is present for Paragraphs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


