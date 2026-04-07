# Quota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The quota ID. | [optional] 
**Quantity** | Pointer to **int32** | The quota quantity. | [optional] 
**Wallet** | Pointer to **bool** | The quota applies to the wallet or not | [optional] 
**DueDate** | Pointer to **NullableTime** | The quota due date. | [optional] 
**NextQuantity** | Pointer to **NullableInt32** | The quota next quantity. | [optional] 
**State** | Pointer to [**QuotaState**](QuotaState.md) |  | [optional] 

## Methods

### NewQuota

`func NewQuota() *Quota`

NewQuota instantiates a new Quota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaWithDefaults

`func NewQuotaWithDefaults() *Quota`

NewQuotaWithDefaults instantiates a new Quota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Quota) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Quota) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Quota) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Quota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuantity

`func (o *Quota) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Quota) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Quota) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Quota) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetWallet

`func (o *Quota) GetWallet() bool`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *Quota) GetWalletOk() (*bool, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *Quota) SetWallet(v bool)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *Quota) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetDueDate

`func (o *Quota) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Quota) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Quota) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Quota) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *Quota) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *Quota) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetNextQuantity

`func (o *Quota) GetNextQuantity() int32`

GetNextQuantity returns the NextQuantity field if non-nil, zero value otherwise.

### GetNextQuantityOk

`func (o *Quota) GetNextQuantityOk() (*int32, bool)`

GetNextQuantityOk returns a tuple with the NextQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextQuantity

`func (o *Quota) SetNextQuantity(v int32)`

SetNextQuantity sets NextQuantity field to given value.

### HasNextQuantity

`func (o *Quota) HasNextQuantity() bool`

HasNextQuantity returns a boolean if a field has been set.

### SetNextQuantityNil

`func (o *Quota) SetNextQuantityNil(b bool)`

 SetNextQuantityNil sets the value for NextQuantity to be an explicit nil

### UnsetNextQuantity
`func (o *Quota) UnsetNextQuantity()`

UnsetNextQuantity ensures that no value is present for NextQuantity, not even an explicit nil
### GetState

`func (o *Quota) GetState() QuotaState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Quota) GetStateOk() (*QuotaState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Quota) SetState(v QuotaState)`

SetState sets State field to given value.

### HasState

`func (o *Quota) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


