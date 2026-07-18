# PaymentUrlRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackUrl** | **string** | The URL where the user will be redirected after payment cancellation. | 
**SuccessUrl** | **string** | The URL where the user will be redirected after successful payment. | 
**Quantity** | Pointer to **map[string]int32** | The payment quantity. | [optional] 

## Methods

### NewPaymentUrlRequestDto

`func NewPaymentUrlRequestDto(backUrl string, successUrl string, ) *PaymentUrlRequestDto`

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

### HasQuantity

`func (o *PaymentUrlRequestDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *PaymentUrlRequestDto) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *PaymentUrlRequestDto) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


