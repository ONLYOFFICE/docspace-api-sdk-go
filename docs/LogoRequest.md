# LogoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TmpFile** | **string** | The picture to cut the logo out of, named by the path that `POST api/2.0/files/logos` returned for it. The  path may be used once and only by the account that uploaded it. | 
**X** | Pointer to **int32** | The left edge of the rectangle cut out of the uploaded picture, counted in pixels from its left side. The  picture itself was already scaled down to fit 1280 by 1280 pixels when it was uploaded. | [optional] 
**Y** | Pointer to **int32** | The top edge of the rectangle cut out of the uploaded picture, counted in pixels from its top. | [optional] 
**Width** | Pointer to **int32** | How wide a piece of the uploaded picture to cut out, in pixels. It has to be sent together with the height,  and the portal builds the four logo sizes out of the piece. | [optional] 
**Height** | Pointer to **int32** | How tall a piece of the uploaded picture to cut out, in pixels. It has to be sent together with the width. | [optional] 

## Methods

### NewLogoRequest

`func NewLogoRequest(tmpFile string, ) *LogoRequest`

NewLogoRequest instantiates a new LogoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogoRequestWithDefaults

`func NewLogoRequestWithDefaults() *LogoRequest`

NewLogoRequestWithDefaults instantiates a new LogoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTmpFile

`func (o *LogoRequest) GetTmpFile() string`

GetTmpFile returns the TmpFile field if non-nil, zero value otherwise.

### GetTmpFileOk

`func (o *LogoRequest) GetTmpFileOk() (*string, bool)`

GetTmpFileOk returns a tuple with the TmpFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpFile

`func (o *LogoRequest) SetTmpFile(v string)`

SetTmpFile sets TmpFile field to given value.


### GetX

`func (o *LogoRequest) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *LogoRequest) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *LogoRequest) SetX(v int32)`

SetX sets X field to given value.

### HasX

`func (o *LogoRequest) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *LogoRequest) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *LogoRequest) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *LogoRequest) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *LogoRequest) HasY() bool`

HasY returns a boolean if a field has been set.

### GetWidth

`func (o *LogoRequest) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *LogoRequest) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *LogoRequest) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *LogoRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *LogoRequest) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *LogoRequest) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *LogoRequest) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *LogoRequest) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


