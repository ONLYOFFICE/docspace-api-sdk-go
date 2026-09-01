# LogoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TmpFile** | **string** | The path to the temporary image file. | 
**X** | Pointer to **int32** | The X coordinate of the rectangle starting point. | [optional] 
**Y** | Pointer to **int32** | The Y coordinate of the rectangle starting point. | [optional] 
**Width** | Pointer to **int32** | The rectangle width. | [optional] 
**Height** | Pointer to **int32** | The rectangle height. | [optional] 

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


