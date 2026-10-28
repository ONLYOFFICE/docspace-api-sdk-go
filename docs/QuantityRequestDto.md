# QuantityRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **map[string]int32** | The plan and the number of units it is to cover, as a single pair. While the portal is on a priced plan the  key has to be the `name` of that same plan, which `GET api/2.0/portal/payment/quota` reports, because the  subscription is resized rather than swapped; the value is the total the subscription is to have afterwards,  not the difference. Exactly one pair is accepted, and a value that is already in effect is refused with 400. | 

## Methods

### NewQuantityRequestDto

`func NewQuantityRequestDto(quantity map[string]int32, ) *QuantityRequestDto`

NewQuantityRequestDto instantiates a new QuantityRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuantityRequestDtoWithDefaults

`func NewQuantityRequestDtoWithDefaults() *QuantityRequestDto`

NewQuantityRequestDtoWithDefaults instantiates a new QuantityRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *QuantityRequestDto) GetQuantity() map[string]int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuantityRequestDto) GetQuantityOk() (*map[string]int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuantityRequestDto) SetQuantity(v map[string]int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


