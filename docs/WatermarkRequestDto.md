# WatermarkRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Specifies whether watermarks are on or off. | [optional] 
**Additions** | Pointer to [**WatermarkAdditions**](WatermarkAdditions.md) |  | [optional] 
**Text** | Pointer to **NullableString** | The watermark text. | [optional] 
**Rotate** | Pointer to **int32** | The watermark text and image rotate angle. | [optional] 
**ImageScale** | Pointer to **int32** | The watermark image scale. | [optional] 
**ImageUrl** | Pointer to **NullableString** | The path to the temporary image file. | [optional] 
**ImageHeight** | Pointer to **float64** | The watermark image height. | [optional] 
**ImageWidth** | Pointer to **float64** | The watermark image width. | [optional] 

## Methods

### NewWatermarkRequestDto

`func NewWatermarkRequestDto() *WatermarkRequestDto`

NewWatermarkRequestDto instantiates a new WatermarkRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkRequestDtoWithDefaults

`func NewWatermarkRequestDtoWithDefaults() *WatermarkRequestDto`

NewWatermarkRequestDtoWithDefaults instantiates a new WatermarkRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WatermarkRequestDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WatermarkRequestDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WatermarkRequestDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WatermarkRequestDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *WatermarkRequestDto) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *WatermarkRequestDto) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetAdditions

`func (o *WatermarkRequestDto) GetAdditions() WatermarkAdditions`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *WatermarkRequestDto) GetAdditionsOk() (*WatermarkAdditions, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *WatermarkRequestDto) SetAdditions(v WatermarkAdditions)`

SetAdditions sets Additions field to given value.

### HasAdditions

`func (o *WatermarkRequestDto) HasAdditions() bool`

HasAdditions returns a boolean if a field has been set.

### GetText

`func (o *WatermarkRequestDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WatermarkRequestDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WatermarkRequestDto) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WatermarkRequestDto) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *WatermarkRequestDto) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *WatermarkRequestDto) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetRotate

`func (o *WatermarkRequestDto) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *WatermarkRequestDto) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *WatermarkRequestDto) SetRotate(v int32)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *WatermarkRequestDto) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetImageScale

`func (o *WatermarkRequestDto) GetImageScale() int32`

GetImageScale returns the ImageScale field if non-nil, zero value otherwise.

### GetImageScaleOk

`func (o *WatermarkRequestDto) GetImageScaleOk() (*int32, bool)`

GetImageScaleOk returns a tuple with the ImageScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScale

`func (o *WatermarkRequestDto) SetImageScale(v int32)`

SetImageScale sets ImageScale field to given value.

### HasImageScale

`func (o *WatermarkRequestDto) HasImageScale() bool`

HasImageScale returns a boolean if a field has been set.

### GetImageUrl

`func (o *WatermarkRequestDto) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *WatermarkRequestDto) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *WatermarkRequestDto) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *WatermarkRequestDto) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *WatermarkRequestDto) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *WatermarkRequestDto) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetImageHeight

`func (o *WatermarkRequestDto) GetImageHeight() float64`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *WatermarkRequestDto) GetImageHeightOk() (*float64, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *WatermarkRequestDto) SetImageHeight(v float64)`

SetImageHeight sets ImageHeight field to given value.

### HasImageHeight

`func (o *WatermarkRequestDto) HasImageHeight() bool`

HasImageHeight returns a boolean if a field has been set.

### GetImageWidth

`func (o *WatermarkRequestDto) GetImageWidth() float64`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *WatermarkRequestDto) GetImageWidthOk() (*float64, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *WatermarkRequestDto) SetImageWidth(v float64)`

SetImageWidth sets ImageWidth field to given value.

### HasImageWidth

`func (o *WatermarkRequestDto) HasImageWidth() bool`

HasImageWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


