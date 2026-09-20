# UpcomingPaymentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The quota that is going to be charged. When a switch to another quota is scheduled, this is the quota  being switched to, so it can differ from what `GET api/2.0/portal/tariff` reports for today. | [optional] 
**Name** | Pointer to **NullableString** | The quota's stable key, which is the same identifier the wallet operations use for a service. | [optional] 
**Title** | Pointer to **NullableString** | The quota name in the portal language, meant to be printed on an invoice preview. | [optional] 
**UnitOfMeasure** | Pointer to **NullableString** | What `quantity` counts, in the portal language - seats, administrators, gigabytes. It is empty for a quota  that is simply on or off. | [optional] 
**Quantity** | Pointer to **int32** | How much is going to be charged for, which is the quantity scheduled for the next period when one has been  scheduled and today's quantity otherwise. | [optional] 
**Wallet** | Pointer to **bool** | Whether the charge is paid out of the portal wallet rather than from the subscription. | [optional] 
**DueDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the charge falls due, in the portal time zone. | [optional] 
**Amount** | Pointer to **float64** | What the charge comes to: the unit price of the quota multiplied by `quantity`. Taxes are not part of it,  and a quota with no price of its own is not listed at all rather than listed with a zero. | [optional] 
**Currency** | Pointer to **NullableString** | The currency `amount` is expressed in, as a three-letter ISO 4217 code. It follows the portal's billing  account, so every entry of one answer carries the same code. | [optional] 

## Methods

### NewUpcomingPaymentDto

`func NewUpcomingPaymentDto() *UpcomingPaymentDto`

NewUpcomingPaymentDto instantiates a new UpcomingPaymentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpcomingPaymentDtoWithDefaults

`func NewUpcomingPaymentDtoWithDefaults() *UpcomingPaymentDto`

NewUpcomingPaymentDtoWithDefaults instantiates a new UpcomingPaymentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpcomingPaymentDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpcomingPaymentDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpcomingPaymentDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpcomingPaymentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpcomingPaymentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpcomingPaymentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpcomingPaymentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpcomingPaymentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpcomingPaymentDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpcomingPaymentDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *UpcomingPaymentDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpcomingPaymentDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpcomingPaymentDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpcomingPaymentDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpcomingPaymentDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpcomingPaymentDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUnitOfMeasure

`func (o *UpcomingPaymentDto) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *UpcomingPaymentDto) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *UpcomingPaymentDto) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *UpcomingPaymentDto) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### SetUnitOfMeasureNil

`func (o *UpcomingPaymentDto) SetUnitOfMeasureNil(b bool)`

 SetUnitOfMeasureNil sets the value for UnitOfMeasure to be an explicit nil

### UnsetUnitOfMeasure
`func (o *UpcomingPaymentDto) UnsetUnitOfMeasure()`

UnsetUnitOfMeasure ensures that no value is present for UnitOfMeasure, not even an explicit nil
### GetQuantity

`func (o *UpcomingPaymentDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UpcomingPaymentDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UpcomingPaymentDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UpcomingPaymentDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetWallet

`func (o *UpcomingPaymentDto) GetWallet() bool`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *UpcomingPaymentDto) GetWalletOk() (*bool, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *UpcomingPaymentDto) SetWallet(v bool)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *UpcomingPaymentDto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetDueDate

`func (o *UpcomingPaymentDto) GetDueDate() ApiDateTime`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *UpcomingPaymentDto) GetDueDateOk() (*ApiDateTime, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *UpcomingPaymentDto) SetDueDate(v ApiDateTime)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *UpcomingPaymentDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetAmount

`func (o *UpcomingPaymentDto) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpcomingPaymentDto) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpcomingPaymentDto) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UpcomingPaymentDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *UpcomingPaymentDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpcomingPaymentDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpcomingPaymentDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpcomingPaymentDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *UpcomingPaymentDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *UpcomingPaymentDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


