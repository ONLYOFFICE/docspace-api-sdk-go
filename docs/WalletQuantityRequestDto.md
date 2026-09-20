# WalletQuantityRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **map[string]int32** | The wallet service and the number of units of it, as a single pair. The key is the `serviceName` of a service  from `GET api/2.0/portal/payment/walletservices`, and the value is read according to  `productQuantityType`: the units to add, or the total the service is to have in the next period. Minimum  quantities apply per service - disk storage starts at 100 units, the DocsCloud developer pack at 10, and the  administrators may not be fewer than the portal already has. Exactly one pair is accepted, and a null or zero  value cancels a change scheduled earlier rather than buying nothing. | 
**ProductQuantityType** | Pointer to [**ProductQuantityType**](ProductQuantityType.md) | How the number in `quantity` is applied. `Add` buys the units straight away and charges them to the portal  wallet, while `Set` charges nothing now and records the quantity the service is to have from the next period.  Only these two are accepted here; `Sub` and `Renew` are refused with 400. | [optional] 

## Methods

### NewWalletQuantityRequestDto

`func NewWalletQuantityRequestDto(quantity map[string]*int32, ) *WalletQuantityRequestDto`

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

`func (o *WalletQuantityRequestDto) GetQuantity() map[string]*int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *WalletQuantityRequestDto) GetQuantityOk() (*map[string]*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *WalletQuantityRequestDto) SetQuantity(v map[string]*int32)`

SetQuantity sets Quantity field to given value.


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


