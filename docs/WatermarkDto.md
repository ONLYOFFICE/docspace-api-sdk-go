# WatermarkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Additions** | [**WatermarkAdditions**](WatermarkAdditions.md) |  | 
**Text** | Pointer to **NullableString** | The watermark text. | [optional] 
**Rotate** | **int32** | The watermark text and image rotate. | 
**ImageScale** | **int32** | The watermark image scale. | 
**ImageUrl** | Pointer to **NullableString** | The watermark image url. | [optional] 
**ImageHeight** | **float64** | The watermark image height. | 
**ImageWidth** | **float64** | The watermark image width. | 

## Methods

### NewWatermarkDto

`func NewWatermarkDto(additions WatermarkAdditions, rotate int32, imageScale int32, imageHeight float64, imageWidth float64, ) *WatermarkDto`

NewWatermarkDto instantiates a new WatermarkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkDtoWithDefaults

`func NewWatermarkDtoWithDefaults() *WatermarkDto`

NewWatermarkDtoWithDefaults instantiates a new WatermarkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditions

`func (o *WatermarkDto) GetAdditions() WatermarkAdditions`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *WatermarkDto) GetAdditionsOk() (*WatermarkAdditions, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *WatermarkDto) SetAdditions(v WatermarkAdditions)`

SetAdditions sets Additions field to given value.


### GetText

`func (o *WatermarkDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WatermarkDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WatermarkDto) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WatermarkDto) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *WatermarkDto) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *WatermarkDto) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetRotate

`func (o *WatermarkDto) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *WatermarkDto) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *WatermarkDto) SetRotate(v int32)`

SetRotate sets Rotate field to given value.


### GetImageScale

`func (o *WatermarkDto) GetImageScale() int32`

GetImageScale returns the ImageScale field if non-nil, zero value otherwise.

### GetImageScaleOk

`func (o *WatermarkDto) GetImageScaleOk() (*int32, bool)`

GetImageScaleOk returns a tuple with the ImageScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScale

`func (o *WatermarkDto) SetImageScale(v int32)`

SetImageScale sets ImageScale field to given value.


### GetImageUrl

`func (o *WatermarkDto) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *WatermarkDto) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *WatermarkDto) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *WatermarkDto) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *WatermarkDto) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *WatermarkDto) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetImageHeight

`func (o *WatermarkDto) GetImageHeight() float64`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *WatermarkDto) GetImageHeightOk() (*float64, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *WatermarkDto) SetImageHeight(v float64)`

SetImageHeight sets ImageHeight field to given value.


### GetImageWidth

`func (o *WatermarkDto) GetImageWidth() float64`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *WatermarkDto) GetImageWidthOk() (*float64, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *WatermarkDto) SetImageWidth(v float64)`

SetImageWidth sets ImageWidth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


