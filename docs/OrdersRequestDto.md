# OrdersRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]OrdersItemRequestDto**](OrdersItemRequestDto.md) | The entries to move, applied one after another in the order they are sent, so each of them shifts the  neighbours the ones before it left behind. | 

## Methods

### NewOrdersRequestDto

`func NewOrdersRequestDto(items []OrdersItemRequestDto, ) *OrdersRequestDto`

NewOrdersRequestDto instantiates a new OrdersRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersRequestDtoWithDefaults

`func NewOrdersRequestDtoWithDefaults() *OrdersRequestDto`

NewOrdersRequestDtoWithDefaults instantiates a new OrdersRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *OrdersRequestDto) GetItems() []OrdersItemRequestDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrdersRequestDto) GetItemsOk() (*[]OrdersItemRequestDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrdersRequestDto) SetItems(v []OrdersItemRequestDto)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *OrdersRequestDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *OrdersRequestDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


