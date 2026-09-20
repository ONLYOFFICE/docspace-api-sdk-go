# OrdersItemRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **int32** | The file or folder to move. | 
**EntryType** | [**FileEntryType**](FileEntryType.md) | Which of the two the identifier names, because a file and a folder may carry the same number. | 
**Order** | **int32** | The position the entry is to take, counting from 1. The entry that held it, and everything after it, is  shifted to make room. A dotted path such as 1.2.3 is accepted as well, of which only the last segment is  read. | 

## Methods

### NewOrdersItemRequestDto

`func NewOrdersItemRequestDto(entryId int32, entryType FileEntryType, order int32, ) *OrdersItemRequestDto`

NewOrdersItemRequestDto instantiates a new OrdersItemRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersItemRequestDtoWithDefaults

`func NewOrdersItemRequestDtoWithDefaults() *OrdersItemRequestDto`

NewOrdersItemRequestDtoWithDefaults instantiates a new OrdersItemRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *OrdersItemRequestDto) GetEntryId() int32`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *OrdersItemRequestDto) GetEntryIdOk() (*int32, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *OrdersItemRequestDto) SetEntryId(v int32)`

SetEntryId sets EntryId field to given value.


### GetEntryType

`func (o *OrdersItemRequestDto) GetEntryType() FileEntryType`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *OrdersItemRequestDto) GetEntryTypeOk() (*FileEntryType, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *OrdersItemRequestDto) SetEntryType(v FileEntryType)`

SetEntryType sets EntryType field to given value.


### GetOrder

`func (o *OrdersItemRequestDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrdersItemRequestDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrdersItemRequestDto) SetOrder(v int32)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


