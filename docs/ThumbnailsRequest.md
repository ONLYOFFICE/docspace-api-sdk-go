# ThumbnailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TmpFile** | Pointer to **NullableString** | The path to the temporary thumbnail file. | [optional] 
**X** | Pointer to **int32** | The thumbnail horizontal coordinate. | [optional] 
**Y** | Pointer to **int32** | The thumbnail vertical coordinate. | [optional] 
**Width** | Pointer to **int32** | The thumbnail width. | [optional] 
**Height** | Pointer to **int32** | The thumbnail height. | [optional] 

## Methods

### NewThumbnailsRequest

`func NewThumbnailsRequest() *ThumbnailsRequest`

NewThumbnailsRequest instantiates a new ThumbnailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbnailsRequestWithDefaults

`func NewThumbnailsRequestWithDefaults() *ThumbnailsRequest`

NewThumbnailsRequestWithDefaults instantiates a new ThumbnailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTmpFile

`func (o *ThumbnailsRequest) GetTmpFile() string`

GetTmpFile returns the TmpFile field if non-nil, zero value otherwise.

### GetTmpFileOk

`func (o *ThumbnailsRequest) GetTmpFileOk() (*string, bool)`

GetTmpFileOk returns a tuple with the TmpFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpFile

`func (o *ThumbnailsRequest) SetTmpFile(v string)`

SetTmpFile sets TmpFile field to given value.

### HasTmpFile

`func (o *ThumbnailsRequest) HasTmpFile() bool`

HasTmpFile returns a boolean if a field has been set.

### SetTmpFileNil

`func (o *ThumbnailsRequest) SetTmpFileNil(b bool)`

 SetTmpFileNil sets the value for TmpFile to be an explicit nil

### UnsetTmpFile
`func (o *ThumbnailsRequest) UnsetTmpFile()`

UnsetTmpFile ensures that no value is present for TmpFile, not even an explicit nil
### GetX

`func (o *ThumbnailsRequest) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ThumbnailsRequest) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ThumbnailsRequest) SetX(v int32)`

SetX sets X field to given value.

### HasX

`func (o *ThumbnailsRequest) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ThumbnailsRequest) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ThumbnailsRequest) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ThumbnailsRequest) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *ThumbnailsRequest) HasY() bool`

HasY returns a boolean if a field has been set.

### GetWidth

`func (o *ThumbnailsRequest) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ThumbnailsRequest) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ThumbnailsRequest) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ThumbnailsRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ThumbnailsRequest) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ThumbnailsRequest) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ThumbnailsRequest) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ThumbnailsRequest) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


