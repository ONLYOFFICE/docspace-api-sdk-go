# PaymentUrlRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackUrl** | **string** | The absolute address the hosted checkout page sends the buyer back to when the purchase is abandoned. It has  to be a well-formed URL and is carried into the checkout page as it is given, so it must be reachable by the  buyer rather than by the portal. | 
**SuccessUrl** | **string** | The absolute address the hosted checkout page sends the buyer to once the payment provider accepts the  purchase. Reaching it says the provider took the money, not that the portal has already been switched to the  new plan, so a client that lands here reads the plan back rather than assuming it. | 
**Quantity** | **map[string]int32** | The plan being bought, as a single pair of the plan name and the number of units of it. The key is the `name`  of a monthly, non-wallet quota from `GET api/2.0/portal/payment/quotas`, and the value is how many  administrators the plan is to cover, which has to be greater than zero. Exactly one pair is accepted; yearly  and wallet products are refused with 400, and wallet services are bought through  `PUT api/2.0/portal/payment/updatewallet` instead. | 

## Methods

### NewPaymentUrlRequestDto

`func NewPaymentUrlRequestDto(backUrl string, successUrl string, quantity map[string]int32, ) *PaymentUrlRequestDto`

NewPaymentUrlRequestDto instantiates a new PaymentUrlRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentUrlRequestDtoWithDefaults

`func NewPaymentUrlRequestDtoWithDefaults() *PaymentUrlRequestDto`

NewPaymentUrlRequestDtoWithDefaults instantiates a new PaymentUrlRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackUrl

`func (o *PaymentUrlRequestDto) GetBackUrl() string`

GetBackUrl returns the BackUrl field if non-nil, zero value otherwise.

### GetBackUrlOk

`func (o *PaymentUrlRequestDto) GetBackUrlOk() (*string, bool)`

GetBackUrlOk returns a tuple with the BackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUrl

`func (o *PaymentUrlRequestDto) SetBackUrl(v string)`

SetBackUrl sets BackUrl field to given value.


### GetSuccessUrl

`func (o *PaymentUrlRequestDto) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *PaymentUrlRequestDto) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *PaymentUrlRequestDto) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### GetQuantity

`func (o *PaymentUrlRequestDto) GetQuantity() map[string]int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PaymentUrlRequestDto) GetQuantityOk() (*map[string]int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PaymentUrlRequestDto) SetQuantity(v map[string]int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


