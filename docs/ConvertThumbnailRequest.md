# ConvertThumbnailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aspect** | Pointer to **NullableInt32** | The mode to fit the image to the height and width specified:  0 - stretch file to fit height and width;  1 - keep the aspect for the image;  2 - convert the metric size of the page into pixels at 96 dpi. | [optional] 
**First** | Pointer to **NullableBool** | Whether a thumbnail is made of the first page only; false makes one of every page. | [optional] 
**Height** | Pointer to **NullableInt32** | The thumbnail height in pixels; not used with aspect 2. | [optional] 
**Width** | Pointer to **NullableInt32** | The thumbnail width in pixels; not used with aspect 2. | [optional] 

## Methods

### NewConvertThumbnailRequest

`func NewConvertThumbnailRequest() *ConvertThumbnailRequest`

NewConvertThumbnailRequest instantiates a new ConvertThumbnailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertThumbnailRequestWithDefaults

`func NewConvertThumbnailRequestWithDefaults() *ConvertThumbnailRequest`

NewConvertThumbnailRequestWithDefaults instantiates a new ConvertThumbnailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspect

`func (o *ConvertThumbnailRequest) GetAspect() int32`

GetAspect returns the Aspect field if non-nil, zero value otherwise.

### GetAspectOk

`func (o *ConvertThumbnailRequest) GetAspectOk() (*int32, bool)`

GetAspectOk returns a tuple with the Aspect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspect

`func (o *ConvertThumbnailRequest) SetAspect(v int32)`

SetAspect sets Aspect field to given value.

### HasAspect

`func (o *ConvertThumbnailRequest) HasAspect() bool`

HasAspect returns a boolean if a field has been set.

### SetAspectNil

`func (o *ConvertThumbnailRequest) SetAspectNil(b bool)`

 SetAspectNil sets the value for Aspect to be an explicit nil

### UnsetAspect
`func (o *ConvertThumbnailRequest) UnsetAspect()`

UnsetAspect ensures that no value is present for Aspect, not even an explicit nil
### GetFirst

`func (o *ConvertThumbnailRequest) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *ConvertThumbnailRequest) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *ConvertThumbnailRequest) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *ConvertThumbnailRequest) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### SetFirstNil

`func (o *ConvertThumbnailRequest) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *ConvertThumbnailRequest) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetHeight

`func (o *ConvertThumbnailRequest) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ConvertThumbnailRequest) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ConvertThumbnailRequest) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ConvertThumbnailRequest) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *ConvertThumbnailRequest) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ConvertThumbnailRequest) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *ConvertThumbnailRequest) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ConvertThumbnailRequest) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ConvertThumbnailRequest) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ConvertThumbnailRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *ConvertThumbnailRequest) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *ConvertThumbnailRequest) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


