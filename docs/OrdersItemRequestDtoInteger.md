# OrdersItemRequestDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **int32** | The entry unique identifier (file or folder). | 
**EntryType** | [**FileEntryType**](FileEntryType.md) |  | 
**Order** | **int32** | The order value. | 

## Methods

### NewOrdersItemRequestDtoInteger

`func NewOrdersItemRequestDtoInteger(entryId int32, entryType FileEntryType, order int32, ) *OrdersItemRequestDtoInteger`

NewOrdersItemRequestDtoInteger instantiates a new OrdersItemRequestDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersItemRequestDtoIntegerWithDefaults

`func NewOrdersItemRequestDtoIntegerWithDefaults() *OrdersItemRequestDtoInteger`

NewOrdersItemRequestDtoIntegerWithDefaults instantiates a new OrdersItemRequestDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *OrdersItemRequestDtoInteger) GetEntryId() int32`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *OrdersItemRequestDtoInteger) GetEntryIdOk() (*int32, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *OrdersItemRequestDtoInteger) SetEntryId(v int32)`

SetEntryId sets EntryId field to given value.


### GetEntryType

`func (o *OrdersItemRequestDtoInteger) GetEntryType() FileEntryType`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *OrdersItemRequestDtoInteger) GetEntryTypeOk() (*FileEntryType, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *OrdersItemRequestDtoInteger) SetEntryType(v FileEntryType)`

SetEntryType sets EntryType field to given value.


### GetOrder

`func (o *OrdersItemRequestDtoInteger) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrdersItemRequestDtoInteger) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrdersItemRequestDtoInteger) SetOrder(v int32)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


