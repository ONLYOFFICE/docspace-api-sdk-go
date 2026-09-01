# OrderRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to **int32** | The order value. | [optional] 

## Methods

### NewOrderRequestDto

`func NewOrderRequestDto() *OrderRequestDto`

NewOrderRequestDto instantiates a new OrderRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRequestDtoWithDefaults

`func NewOrderRequestDtoWithDefaults() *OrderRequestDto`

NewOrderRequestDtoWithDefaults instantiates a new OrderRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *OrderRequestDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrderRequestDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrderRequestDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OrderRequestDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


