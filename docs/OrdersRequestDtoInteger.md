# OrdersRequestDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]OrdersItemRequestDtoInteger**](OrdersItemRequestDtoInteger.md) | The list of items with their ordering information. | 

## Methods

### NewOrdersRequestDtoInteger

`func NewOrdersRequestDtoInteger(items []OrdersItemRequestDtoInteger, ) *OrdersRequestDtoInteger`

NewOrdersRequestDtoInteger instantiates a new OrdersRequestDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersRequestDtoIntegerWithDefaults

`func NewOrdersRequestDtoIntegerWithDefaults() *OrdersRequestDtoInteger`

NewOrdersRequestDtoIntegerWithDefaults instantiates a new OrdersRequestDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *OrdersRequestDtoInteger) GetItems() []OrdersItemRequestDtoInteger`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrdersRequestDtoInteger) GetItemsOk() (*[]OrdersItemRequestDtoInteger, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrdersRequestDtoInteger) SetItems(v []OrdersItemRequestDtoInteger)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *OrdersRequestDtoInteger) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *OrdersRequestDtoInteger) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


