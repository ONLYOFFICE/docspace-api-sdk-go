# TariffQuotaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The quota this entry stands for. `GET api/2.0/portal/payment/quotas` describes the quota behind the ID,  including what its `quantity` counts; a negative ID belongs to a built-in quota rather than a purchased  one. | [optional] 
**Quantity** | Pointer to **int32** | How much of the quota the portal holds, in whatever the quota itself is measured in - seats for a plan,  gigabytes for storage. It is `1` for a quota that is simply on or off. | [optional] 
**Wallet** | Pointer to **bool** | Whether the quota is paid for out of the portal wallet as it is consumed, rather than being part of the  subscription charged per period. | [optional] 
**Additional** | Pointer to **bool** | Whether this is an add-on bought on top of the plan rather than the plan itself. Exactly one entry of  `quotas` is the plan, and the rest are add-ons. | [optional] 
**DueDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When this quota runs out, in the portal time zone. An add-on can end earlier or later than the  subscription; a quota with no deadline of its own reports the subscription's `dueDate` instead of an empty  value. | [optional] 
**NextQuantity** | Pointer to **NullableInt32** | The quantity the next period is going to be charged for, when a change has been scheduled. It is empty  while `quantity` simply carries over. | [optional] 
**NextQuota** | Pointer to **NullableInt32** | The quota this one is scheduled to be replaced by at the start of the next period, empty when no such  switch is planned. `GET api/2.0/portal/tariff/upcoming` already reports the charge for the replacement. | [optional] 
**State** | Pointer to [**QuotaState**](QuotaState.md) | Whether the quota is still running or its deadline has passed. It is empty for a quota that has no  deadline of its own, which means it lasts as long as the subscription does. | [optional] 

## Methods

### NewTariffQuotaDto

`func NewTariffQuotaDto() *TariffQuotaDto`

NewTariffQuotaDto instantiates a new TariffQuotaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffQuotaDtoWithDefaults

`func NewTariffQuotaDtoWithDefaults() *TariffQuotaDto`

NewTariffQuotaDtoWithDefaults instantiates a new TariffQuotaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TariffQuotaDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TariffQuotaDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TariffQuotaDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TariffQuotaDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuantity

`func (o *TariffQuotaDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TariffQuotaDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TariffQuotaDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *TariffQuotaDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetWallet

`func (o *TariffQuotaDto) GetWallet() bool`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *TariffQuotaDto) GetWalletOk() (*bool, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *TariffQuotaDto) SetWallet(v bool)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *TariffQuotaDto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetAdditional

`func (o *TariffQuotaDto) GetAdditional() bool`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *TariffQuotaDto) GetAdditionalOk() (*bool, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *TariffQuotaDto) SetAdditional(v bool)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *TariffQuotaDto) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.

### GetDueDate

`func (o *TariffQuotaDto) GetDueDate() ApiDateTime`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TariffQuotaDto) GetDueDateOk() (*ApiDateTime, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TariffQuotaDto) SetDueDate(v ApiDateTime)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TariffQuotaDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetNextQuantity

`func (o *TariffQuotaDto) GetNextQuantity() int32`

GetNextQuantity returns the NextQuantity field if non-nil, zero value otherwise.

### GetNextQuantityOk

`func (o *TariffQuotaDto) GetNextQuantityOk() (*int32, bool)`

GetNextQuantityOk returns a tuple with the NextQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextQuantity

`func (o *TariffQuotaDto) SetNextQuantity(v int32)`

SetNextQuantity sets NextQuantity field to given value.

### HasNextQuantity

`func (o *TariffQuotaDto) HasNextQuantity() bool`

HasNextQuantity returns a boolean if a field has been set.

### SetNextQuantityNil

`func (o *TariffQuotaDto) SetNextQuantityNil(b bool)`

 SetNextQuantityNil sets the value for NextQuantity to be an explicit nil

### UnsetNextQuantity
`func (o *TariffQuotaDto) UnsetNextQuantity()`

UnsetNextQuantity ensures that no value is present for NextQuantity, not even an explicit nil
### GetNextQuota

`func (o *TariffQuotaDto) GetNextQuota() int32`

GetNextQuota returns the NextQuota field if non-nil, zero value otherwise.

### GetNextQuotaOk

`func (o *TariffQuotaDto) GetNextQuotaOk() (*int32, bool)`

GetNextQuotaOk returns a tuple with the NextQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextQuota

`func (o *TariffQuotaDto) SetNextQuota(v int32)`

SetNextQuota sets NextQuota field to given value.

### HasNextQuota

`func (o *TariffQuotaDto) HasNextQuota() bool`

HasNextQuota returns a boolean if a field has been set.

### SetNextQuotaNil

`func (o *TariffQuotaDto) SetNextQuotaNil(b bool)`

 SetNextQuotaNil sets the value for NextQuota to be an explicit nil

### UnsetNextQuota
`func (o *TariffQuotaDto) UnsetNextQuota()`

UnsetNextQuota ensures that no value is present for NextQuota, not even an explicit nil
### GetState

`func (o *TariffQuotaDto) GetState() QuotaState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TariffQuotaDto) GetStateOk() (*QuotaState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TariffQuotaDto) SetState(v QuotaState)`

SetState sets State field to given value.

### HasState

`func (o *TariffQuotaDto) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


