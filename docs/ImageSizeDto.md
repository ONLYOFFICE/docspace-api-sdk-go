# ImageSizeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int32** | The height of the image in pixels, read from the stored file rather than from any display setting. | [optional] 
**Width** | Pointer to **int32** | The width of the image in pixels, read from the stored file rather than from any display setting. | [optional] 

## Methods

### NewImageSizeDto

`func NewImageSizeDto() *ImageSizeDto`

NewImageSizeDto instantiates a new ImageSizeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageSizeDtoWithDefaults

`func NewImageSizeDtoWithDefaults() *ImageSizeDto`

NewImageSizeDtoWithDefaults instantiates a new ImageSizeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *ImageSizeDto) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ImageSizeDto) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ImageSizeDto) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ImageSizeDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *ImageSizeDto) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ImageSizeDto) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ImageSizeDto) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ImageSizeDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


