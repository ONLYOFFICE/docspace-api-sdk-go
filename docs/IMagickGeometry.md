# IMagickGeometry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectRatio** | Pointer to **bool** |  | [optional] [readonly] 
**FillArea** | Pointer to **bool** |  | [optional] 
**Greater** | Pointer to **bool** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**IgnoreAspectRatio** | Pointer to **bool** |  | [optional] 
**IsPercentage** | Pointer to **bool** |  | [optional] 
**Less** | Pointer to **bool** |  | [optional] 
**LimitPixels** | Pointer to **bool** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 
**X** | Pointer to **int32** |  | [optional] 
**Y** | Pointer to **int32** |  | [optional] 

## Methods

### NewIMagickGeometry

`func NewIMagickGeometry() *IMagickGeometry`

NewIMagickGeometry instantiates a new IMagickGeometry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIMagickGeometryWithDefaults

`func NewIMagickGeometryWithDefaults() *IMagickGeometry`

NewIMagickGeometryWithDefaults instantiates a new IMagickGeometry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectRatio

`func (o *IMagickGeometry) GetAspectRatio() bool`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *IMagickGeometry) GetAspectRatioOk() (*bool, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *IMagickGeometry) SetAspectRatio(v bool)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *IMagickGeometry) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### GetFillArea

`func (o *IMagickGeometry) GetFillArea() bool`

GetFillArea returns the FillArea field if non-nil, zero value otherwise.

### GetFillAreaOk

`func (o *IMagickGeometry) GetFillAreaOk() (*bool, bool)`

GetFillAreaOk returns a tuple with the FillArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillArea

`func (o *IMagickGeometry) SetFillArea(v bool)`

SetFillArea sets FillArea field to given value.

### HasFillArea

`func (o *IMagickGeometry) HasFillArea() bool`

HasFillArea returns a boolean if a field has been set.

### GetGreater

`func (o *IMagickGeometry) GetGreater() bool`

GetGreater returns the Greater field if non-nil, zero value otherwise.

### GetGreaterOk

`func (o *IMagickGeometry) GetGreaterOk() (*bool, bool)`

GetGreaterOk returns a tuple with the Greater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreater

`func (o *IMagickGeometry) SetGreater(v bool)`

SetGreater sets Greater field to given value.

### HasGreater

`func (o *IMagickGeometry) HasGreater() bool`

HasGreater returns a boolean if a field has been set.

### GetHeight

`func (o *IMagickGeometry) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *IMagickGeometry) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *IMagickGeometry) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *IMagickGeometry) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetIgnoreAspectRatio

`func (o *IMagickGeometry) GetIgnoreAspectRatio() bool`

GetIgnoreAspectRatio returns the IgnoreAspectRatio field if non-nil, zero value otherwise.

### GetIgnoreAspectRatioOk

`func (o *IMagickGeometry) GetIgnoreAspectRatioOk() (*bool, bool)`

GetIgnoreAspectRatioOk returns a tuple with the IgnoreAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreAspectRatio

`func (o *IMagickGeometry) SetIgnoreAspectRatio(v bool)`

SetIgnoreAspectRatio sets IgnoreAspectRatio field to given value.

### HasIgnoreAspectRatio

`func (o *IMagickGeometry) HasIgnoreAspectRatio() bool`

HasIgnoreAspectRatio returns a boolean if a field has been set.

### GetIsPercentage

`func (o *IMagickGeometry) GetIsPercentage() bool`

GetIsPercentage returns the IsPercentage field if non-nil, zero value otherwise.

### GetIsPercentageOk

`func (o *IMagickGeometry) GetIsPercentageOk() (*bool, bool)`

GetIsPercentageOk returns a tuple with the IsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPercentage

`func (o *IMagickGeometry) SetIsPercentage(v bool)`

SetIsPercentage sets IsPercentage field to given value.

### HasIsPercentage

`func (o *IMagickGeometry) HasIsPercentage() bool`

HasIsPercentage returns a boolean if a field has been set.

### GetLess

`func (o *IMagickGeometry) GetLess() bool`

GetLess returns the Less field if non-nil, zero value otherwise.

### GetLessOk

`func (o *IMagickGeometry) GetLessOk() (*bool, bool)`

GetLessOk returns a tuple with the Less field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLess

`func (o *IMagickGeometry) SetLess(v bool)`

SetLess sets Less field to given value.

### HasLess

`func (o *IMagickGeometry) HasLess() bool`

HasLess returns a boolean if a field has been set.

### GetLimitPixels

`func (o *IMagickGeometry) GetLimitPixels() bool`

GetLimitPixels returns the LimitPixels field if non-nil, zero value otherwise.

### GetLimitPixelsOk

`func (o *IMagickGeometry) GetLimitPixelsOk() (*bool, bool)`

GetLimitPixelsOk returns a tuple with the LimitPixels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPixels

`func (o *IMagickGeometry) SetLimitPixels(v bool)`

SetLimitPixels sets LimitPixels field to given value.

### HasLimitPixels

`func (o *IMagickGeometry) HasLimitPixels() bool`

HasLimitPixels returns a boolean if a field has been set.

### GetWidth

`func (o *IMagickGeometry) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *IMagickGeometry) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *IMagickGeometry) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *IMagickGeometry) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetX

`func (o *IMagickGeometry) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *IMagickGeometry) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *IMagickGeometry) SetX(v int32)`

SetX sets X field to given value.

### HasX

`func (o *IMagickGeometry) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *IMagickGeometry) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *IMagickGeometry) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *IMagickGeometry) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *IMagickGeometry) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


