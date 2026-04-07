# WalletQuantityRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **map[string]int32** | The mapping of item identifiers to their respective quantities in the payment. | [optional] 
**ProductQuantityType** | Pointer to [**ProductQuantityType**](ProductQuantityType.md) |  | [optional] 

## Methods

### NewWalletQuantityRequestDto

`func NewWalletQuantityRequestDto() *WalletQuantityRequestDto`

NewWalletQuantityRequestDto instantiates a new WalletQuantityRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletQuantityRequestDtoWithDefaults

`func NewWalletQuantityRequestDtoWithDefaults() *WalletQuantityRequestDto`

NewWalletQuantityRequestDtoWithDefaults instantiates a new WalletQuantityRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *WalletQuantityRequestDto) GetQuantity() map[string]int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *WalletQuantityRequestDto) GetQuantityOk() (*map[string]int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *WalletQuantityRequestDto) SetQuantity(v map[string]int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *WalletQuantityRequestDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *WalletQuantityRequestDto) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *WalletQuantityRequestDto) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetProductQuantityType

`func (o *WalletQuantityRequestDto) GetProductQuantityType() ProductQuantityType`

GetProductQuantityType returns the ProductQuantityType field if non-nil, zero value otherwise.

### GetProductQuantityTypeOk

`func (o *WalletQuantityRequestDto) GetProductQuantityTypeOk() (*ProductQuantityType, bool)`

GetProductQuantityTypeOk returns a tuple with the ProductQuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductQuantityType

`func (o *WalletQuantityRequestDto) SetProductQuantityType(v ProductQuantityType)`

SetProductQuantityType sets ProductQuantityType field to given value.

### HasProductQuantityType

`func (o *WalletQuantityRequestDto) HasProductQuantityType() bool`

HasProductQuantityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


