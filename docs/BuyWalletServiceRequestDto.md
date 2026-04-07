# BuyWalletServiceRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** | Number of services provided. | [optional] 
**ServiceName** | Pointer to **NullableString** | The service name. | [optional] 

## Methods

### NewBuyWalletServiceRequestDto

`func NewBuyWalletServiceRequestDto() *BuyWalletServiceRequestDto`

NewBuyWalletServiceRequestDto instantiates a new BuyWalletServiceRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyWalletServiceRequestDtoWithDefaults

`func NewBuyWalletServiceRequestDtoWithDefaults() *BuyWalletServiceRequestDto`

NewBuyWalletServiceRequestDtoWithDefaults instantiates a new BuyWalletServiceRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *BuyWalletServiceRequestDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BuyWalletServiceRequestDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BuyWalletServiceRequestDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BuyWalletServiceRequestDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetServiceName

`func (o *BuyWalletServiceRequestDto) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *BuyWalletServiceRequestDto) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *BuyWalletServiceRequestDto) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *BuyWalletServiceRequestDto) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *BuyWalletServiceRequestDto) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *BuyWalletServiceRequestDto) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


