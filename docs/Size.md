# Size

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int32** | Gets or sets the height dimension of an object, typically measured in pixels or other unit.  It defines the vertical size of the object. | [optional] 
**Width** | Pointer to **int32** | Gets or sets the width dimension of an object, typically measured in pixels or other unit. | [optional] 

## Methods

### NewSize

`func NewSize() *Size`

NewSize instantiates a new Size object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeWithDefaults

`func NewSizeWithDefaults() *Size`

NewSizeWithDefaults instantiates a new Size object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *Size) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Size) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Size) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Size) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *Size) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Size) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Size) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Size) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


