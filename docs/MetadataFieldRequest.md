# MetadataFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The field name. | [optional] 
**Type** | Pointer to [**MetadataFieldType**](MetadataFieldType.md) | The field type. | [optional] 
**Options** | Pointer to [**[]MetadataFieldOptionRequest**](MetadataFieldOptionRequest.md) | The choice options of the field. | [optional] 
**Order** | Pointer to **NullableInt32** | The display position of the field inside the template: the fields are shown by it ascending, and equal positions  keep the order of creation. | [optional] 

## Methods

### NewMetadataFieldRequest

`func NewMetadataFieldRequest() *MetadataFieldRequest`

NewMetadataFieldRequest instantiates a new MetadataFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFieldRequestWithDefaults

`func NewMetadataFieldRequestWithDefaults() *MetadataFieldRequest`

NewMetadataFieldRequestWithDefaults instantiates a new MetadataFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetadataFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataFieldRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataFieldRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetadataFieldRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetadataFieldRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *MetadataFieldRequest) GetType() MetadataFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetadataFieldRequest) GetTypeOk() (*MetadataFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetadataFieldRequest) SetType(v MetadataFieldType)`

SetType sets Type field to given value.

### HasType

`func (o *MetadataFieldRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *MetadataFieldRequest) GetOptions() []MetadataFieldOptionRequest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MetadataFieldRequest) GetOptionsOk() (*[]MetadataFieldOptionRequest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MetadataFieldRequest) SetOptions(v []MetadataFieldOptionRequest)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MetadataFieldRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MetadataFieldRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MetadataFieldRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetOrder

`func (o *MetadataFieldRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MetadataFieldRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MetadataFieldRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MetadataFieldRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *MetadataFieldRequest) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *MetadataFieldRequest) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


