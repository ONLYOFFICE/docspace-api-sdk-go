# UpdateMetadataFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The new field name. | [optional] 
**Type** | Pointer to [**MetadataFieldType**](MetadataFieldType.md) | The new field type. The type can be changed only while the field has no values. | [optional] 
**Options** | Pointer to [**[]MetadataFieldOptionRequest**](MetadataFieldOptionRequest.md) | The new choice options of the field. The options in use cannot be removed. | [optional] 
**Order** | Pointer to **NullableInt32** | The new display position of the field inside the template: the fields are shown by it ascending. | [optional] 

## Methods

### NewUpdateMetadataFieldRequest

`func NewUpdateMetadataFieldRequest() *UpdateMetadataFieldRequest`

NewUpdateMetadataFieldRequest instantiates a new UpdateMetadataFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMetadataFieldRequestWithDefaults

`func NewUpdateMetadataFieldRequestWithDefaults() *UpdateMetadataFieldRequest`

NewUpdateMetadataFieldRequestWithDefaults instantiates a new UpdateMetadataFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMetadataFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMetadataFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMetadataFieldRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMetadataFieldRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateMetadataFieldRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateMetadataFieldRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *UpdateMetadataFieldRequest) GetType() MetadataFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateMetadataFieldRequest) GetTypeOk() (*MetadataFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateMetadataFieldRequest) SetType(v MetadataFieldType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateMetadataFieldRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateMetadataFieldRequest) GetOptions() []MetadataFieldOptionRequest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateMetadataFieldRequest) GetOptionsOk() (*[]MetadataFieldOptionRequest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateMetadataFieldRequest) SetOptions(v []MetadataFieldOptionRequest)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateMetadataFieldRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *UpdateMetadataFieldRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *UpdateMetadataFieldRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetOrder

`func (o *UpdateMetadataFieldRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UpdateMetadataFieldRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UpdateMetadataFieldRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *UpdateMetadataFieldRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *UpdateMetadataFieldRequest) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *UpdateMetadataFieldRequest) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


