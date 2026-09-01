# WhiteLabelItemSizeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectRatio** | Pointer to **bool** | Specifies whether the size is an aspect ratio. | [optional] 
**FillArea** | Pointer to **bool** | Specifies whether the logo is resized based on the smallest fitting dimension. | [optional] 
**Greater** | Pointer to **bool** | Specifies whether the logo is resized only if it is greater than the size. | [optional] 
**Height** | Pointer to **int32** | The logo height, in pixels. | [optional] 
**IgnoreAspectRatio** | Pointer to **bool** | Specifies whether the logo is resized without preserving the aspect ratio. | [optional] 
**IsPercentage** | Pointer to **bool** | Specifies whether the width and height are expressed as percentages. | [optional] 
**Less** | Pointer to **bool** | Specifies whether the logo is resized only if it is less than the size. | [optional] 
**LimitPixels** | Pointer to **bool** | Specifies whether the logo is resized using a pixel area count limit. | [optional] 
**Width** | Pointer to **int32** | The logo width, in pixels. | [optional] 
**X** | Pointer to **int32** | The X offset from the origin, in pixels. | [optional] 
**Y** | Pointer to **int32** | The Y offset from the origin, in pixels. | [optional] 

## Methods

### NewWhiteLabelItemSizeDto

`func NewWhiteLabelItemSizeDto() *WhiteLabelItemSizeDto`

NewWhiteLabelItemSizeDto instantiates a new WhiteLabelItemSizeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhiteLabelItemSizeDtoWithDefaults

`func NewWhiteLabelItemSizeDtoWithDefaults() *WhiteLabelItemSizeDto`

NewWhiteLabelItemSizeDtoWithDefaults instantiates a new WhiteLabelItemSizeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectRatio

`func (o *WhiteLabelItemSizeDto) GetAspectRatio() bool`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *WhiteLabelItemSizeDto) GetAspectRatioOk() (*bool, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *WhiteLabelItemSizeDto) SetAspectRatio(v bool)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *WhiteLabelItemSizeDto) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### GetFillArea

`func (o *WhiteLabelItemSizeDto) GetFillArea() bool`

GetFillArea returns the FillArea field if non-nil, zero value otherwise.

### GetFillAreaOk

`func (o *WhiteLabelItemSizeDto) GetFillAreaOk() (*bool, bool)`

GetFillAreaOk returns a tuple with the FillArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillArea

`func (o *WhiteLabelItemSizeDto) SetFillArea(v bool)`

SetFillArea sets FillArea field to given value.

### HasFillArea

`func (o *WhiteLabelItemSizeDto) HasFillArea() bool`

HasFillArea returns a boolean if a field has been set.

### GetGreater

`func (o *WhiteLabelItemSizeDto) GetGreater() bool`

GetGreater returns the Greater field if non-nil, zero value otherwise.

### GetGreaterOk

`func (o *WhiteLabelItemSizeDto) GetGreaterOk() (*bool, bool)`

GetGreaterOk returns a tuple with the Greater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreater

`func (o *WhiteLabelItemSizeDto) SetGreater(v bool)`

SetGreater sets Greater field to given value.

### HasGreater

`func (o *WhiteLabelItemSizeDto) HasGreater() bool`

HasGreater returns a boolean if a field has been set.

### GetHeight

`func (o *WhiteLabelItemSizeDto) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *WhiteLabelItemSizeDto) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *WhiteLabelItemSizeDto) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *WhiteLabelItemSizeDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetIgnoreAspectRatio

`func (o *WhiteLabelItemSizeDto) GetIgnoreAspectRatio() bool`

GetIgnoreAspectRatio returns the IgnoreAspectRatio field if non-nil, zero value otherwise.

### GetIgnoreAspectRatioOk

`func (o *WhiteLabelItemSizeDto) GetIgnoreAspectRatioOk() (*bool, bool)`

GetIgnoreAspectRatioOk returns a tuple with the IgnoreAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreAspectRatio

`func (o *WhiteLabelItemSizeDto) SetIgnoreAspectRatio(v bool)`

SetIgnoreAspectRatio sets IgnoreAspectRatio field to given value.

### HasIgnoreAspectRatio

`func (o *WhiteLabelItemSizeDto) HasIgnoreAspectRatio() bool`

HasIgnoreAspectRatio returns a boolean if a field has been set.

### GetIsPercentage

`func (o *WhiteLabelItemSizeDto) GetIsPercentage() bool`

GetIsPercentage returns the IsPercentage field if non-nil, zero value otherwise.

### GetIsPercentageOk

`func (o *WhiteLabelItemSizeDto) GetIsPercentageOk() (*bool, bool)`

GetIsPercentageOk returns a tuple with the IsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPercentage

`func (o *WhiteLabelItemSizeDto) SetIsPercentage(v bool)`

SetIsPercentage sets IsPercentage field to given value.

### HasIsPercentage

`func (o *WhiteLabelItemSizeDto) HasIsPercentage() bool`

HasIsPercentage returns a boolean if a field has been set.

### GetLess

`func (o *WhiteLabelItemSizeDto) GetLess() bool`

GetLess returns the Less field if non-nil, zero value otherwise.

### GetLessOk

`func (o *WhiteLabelItemSizeDto) GetLessOk() (*bool, bool)`

GetLessOk returns a tuple with the Less field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLess

`func (o *WhiteLabelItemSizeDto) SetLess(v bool)`

SetLess sets Less field to given value.

### HasLess

`func (o *WhiteLabelItemSizeDto) HasLess() bool`

HasLess returns a boolean if a field has been set.

### GetLimitPixels

`func (o *WhiteLabelItemSizeDto) GetLimitPixels() bool`

GetLimitPixels returns the LimitPixels field if non-nil, zero value otherwise.

### GetLimitPixelsOk

`func (o *WhiteLabelItemSizeDto) GetLimitPixelsOk() (*bool, bool)`

GetLimitPixelsOk returns a tuple with the LimitPixels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPixels

`func (o *WhiteLabelItemSizeDto) SetLimitPixels(v bool)`

SetLimitPixels sets LimitPixels field to given value.

### HasLimitPixels

`func (o *WhiteLabelItemSizeDto) HasLimitPixels() bool`

HasLimitPixels returns a boolean if a field has been set.

### GetWidth

`func (o *WhiteLabelItemSizeDto) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *WhiteLabelItemSizeDto) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *WhiteLabelItemSizeDto) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *WhiteLabelItemSizeDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetX

`func (o *WhiteLabelItemSizeDto) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *WhiteLabelItemSizeDto) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *WhiteLabelItemSizeDto) SetX(v int32)`

SetX sets X field to given value.

### HasX

`func (o *WhiteLabelItemSizeDto) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *WhiteLabelItemSizeDto) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *WhiteLabelItemSizeDto) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *WhiteLabelItemSizeDto) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *WhiteLabelItemSizeDto) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


