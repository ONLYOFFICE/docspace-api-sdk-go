# DepositDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The identifier of the top-up, which is what `GET api/2.0/portal/payment/customer/deposits/{id}/invoice` takes  to read its invoice. | [optional] 
**Date** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the top-up was booked, as an ISO 8601 timestamp that carries its own offset: that of the portal's time  zone, written as `Z` for a portal in UTC - the zone in which `startDate` and `endDate` without an offset are  read. | [optional] 
**Amount** | Pointer to **float64** | The amount that went into the wallet, in `currency`. | [optional] 
**Currency** | Pointer to **NullableString** | The currency `amount` is expressed in, as a three-letter ISO 4217 code - the accounting currency of the  wallet. | [optional] 
**Status** | Pointer to [**OperationStatus**](OperationStatus.md) | How the top-up ended. A pending one is still being settled and may change later, so only a completed one has  actually raised the balance. | [optional] 
**HasInvoice** | Pointer to **bool** | Whether the billing provider holds an invoice for the payment behind the top-up, which is when  `GET api/2.0/portal/payment/customer/deposits/{id}/invoice` can open it. It is `false` for a top-up whose  payment came without one, and asking for its invoice then fails with 404. | [optional] 

## Methods

### NewDepositDto

`func NewDepositDto() *DepositDto`

NewDepositDto instantiates a new DepositDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositDtoWithDefaults

`func NewDepositDtoWithDefaults() *DepositDto`

NewDepositDtoWithDefaults instantiates a new DepositDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DepositDto) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepositDto) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepositDto) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DepositDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *DepositDto) GetDate() ApiDateTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DepositDto) GetDateOk() (*ApiDateTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DepositDto) SetDate(v ApiDateTime)`

SetDate sets Date field to given value.

### HasDate

`func (o *DepositDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAmount

`func (o *DepositDto) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DepositDto) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DepositDto) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DepositDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *DepositDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DepositDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DepositDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DepositDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *DepositDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *DepositDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetStatus

`func (o *DepositDto) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DepositDto) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DepositDto) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DepositDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHasInvoice

`func (o *DepositDto) GetHasInvoice() bool`

GetHasInvoice returns the HasInvoice field if non-nil, zero value otherwise.

### GetHasInvoiceOk

`func (o *DepositDto) GetHasInvoiceOk() (*bool, bool)`

GetHasInvoiceOk returns a tuple with the HasInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInvoice

`func (o *DepositDto) SetHasInvoice(v bool)`

SetHasInvoice sets HasInvoice field to given value.

### HasHasInvoice

`func (o *DepositDto) HasHasInvoice() bool`

HasHasInvoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


