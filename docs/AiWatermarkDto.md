# AiWatermarkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Additions** | [**AiWatermarkAdditions**](AiWatermarkAdditions.md) | Which details of the reader and of the room are stamped alongside the text. The values combine, so a number  that is not a member on its own is the sum of several of them, and 0 means that only the text is stamped. | 
**Text** | Pointer to **NullableString** | The fixed line drawn over the document, printed before the details selected alongside it. Empty when the room  stamps an image instead. | [optional] 
**Rotate** | **int32** | How far the stamp is turned, in degrees, with negative values turning it anticlockwise and 0 drawing it  horizontally. | 
**ImageScale** | **int32** | How large the image is drawn, as a percentage of its own size. It is 0 for a text watermark, where nothing is  scaled. | 
**ImageUrl** | Pointer to **NullableString** | The address the stamped picture is served from, inside the storage of the room. Empty for a text watermark. | [optional] 
**ImageHeight** | **float64** | The height the picture is drawn with, in pixels, kept together with the width so that the proportions survive.  It is 0 for a text watermark. | 
**ImageWidth** | **float64** | The width the picture is drawn with, in pixels, kept together with the height so that the proportions survive.  It is 0 for a text watermark. | 

## Methods

### NewAiWatermarkDto

`func NewAiWatermarkDto(additions AiWatermarkAdditions, rotate int32, imageScale int32, imageHeight float64, imageWidth float64, ) *AiWatermarkDto`

NewAiWatermarkDto instantiates a new AiWatermarkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiWatermarkDtoWithDefaults

`func NewAiWatermarkDtoWithDefaults() *AiWatermarkDto`

NewAiWatermarkDtoWithDefaults instantiates a new AiWatermarkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditions

`func (o *AiWatermarkDto) GetAdditions() AiWatermarkAdditions`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *AiWatermarkDto) GetAdditionsOk() (*AiWatermarkAdditions, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *AiWatermarkDto) SetAdditions(v AiWatermarkAdditions)`

SetAdditions sets Additions field to given value.


### GetText

`func (o *AiWatermarkDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AiWatermarkDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AiWatermarkDto) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AiWatermarkDto) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *AiWatermarkDto) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *AiWatermarkDto) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetRotate

`func (o *AiWatermarkDto) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *AiWatermarkDto) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *AiWatermarkDto) SetRotate(v int32)`

SetRotate sets Rotate field to given value.


### GetImageScale

`func (o *AiWatermarkDto) GetImageScale() int32`

GetImageScale returns the ImageScale field if non-nil, zero value otherwise.

### GetImageScaleOk

`func (o *AiWatermarkDto) GetImageScaleOk() (*int32, bool)`

GetImageScaleOk returns a tuple with the ImageScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScale

`func (o *AiWatermarkDto) SetImageScale(v int32)`

SetImageScale sets ImageScale field to given value.


### GetImageUrl

`func (o *AiWatermarkDto) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *AiWatermarkDto) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *AiWatermarkDto) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *AiWatermarkDto) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *AiWatermarkDto) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *AiWatermarkDto) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetImageHeight

`func (o *AiWatermarkDto) GetImageHeight() float64`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *AiWatermarkDto) GetImageHeightOk() (*float64, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *AiWatermarkDto) SetImageHeight(v float64)`

SetImageHeight sets ImageHeight field to given value.


### GetImageWidth

`func (o *AiWatermarkDto) GetImageWidth() float64`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *AiWatermarkDto) GetImageWidthOk() (*float64, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *AiWatermarkDto) SetImageWidth(v float64)`

SetImageWidth sets ImageWidth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


