# CustomerMonthlyUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | Pointer to **int32** | The year the month belongs to. Months are cut in the portal time zone, so a movement at the edge of a  month falls where the portal sees it and not where UTC does. | [optional] 
**Month** | Pointer to **int32** | The month itself, January being 1. Only months that had spending appear at all, so a gap in the list is a  month with nothing in it rather than missing data. | [optional] 
**Currency** | Pointer to **NullableString** | The currency `totalAmount` is expressed in, as a three-letter ISO 4217 code - the accounting currency of  the wallet. | [optional] 
**TotalAmount** | Pointer to **float64** | What the month came to across every service, as a positive amount spent rather than a signed balance. | [optional] 
**OperationCount** | Pointer to **int32** | How many separate movements that total was added up from, for a client that wants to show the weight  behind a figure. The movements themselves are in `GET api/2.0/portal/payment/customer/operations`. | [optional] 

## Methods

### NewCustomerMonthlyUsageDto

`func NewCustomerMonthlyUsageDto() *CustomerMonthlyUsageDto`

NewCustomerMonthlyUsageDto instantiates a new CustomerMonthlyUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerMonthlyUsageDtoWithDefaults

`func NewCustomerMonthlyUsageDtoWithDefaults() *CustomerMonthlyUsageDto`

NewCustomerMonthlyUsageDtoWithDefaults instantiates a new CustomerMonthlyUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *CustomerMonthlyUsageDto) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *CustomerMonthlyUsageDto) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *CustomerMonthlyUsageDto) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *CustomerMonthlyUsageDto) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *CustomerMonthlyUsageDto) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *CustomerMonthlyUsageDto) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *CustomerMonthlyUsageDto) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *CustomerMonthlyUsageDto) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetCurrency

`func (o *CustomerMonthlyUsageDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerMonthlyUsageDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerMonthlyUsageDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CustomerMonthlyUsageDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CustomerMonthlyUsageDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CustomerMonthlyUsageDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetTotalAmount

`func (o *CustomerMonthlyUsageDto) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *CustomerMonthlyUsageDto) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *CustomerMonthlyUsageDto) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *CustomerMonthlyUsageDto) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetOperationCount

`func (o *CustomerMonthlyUsageDto) GetOperationCount() int32`

GetOperationCount returns the OperationCount field if non-nil, zero value otherwise.

### GetOperationCountOk

`func (o *CustomerMonthlyUsageDto) GetOperationCountOk() (*int32, bool)`

GetOperationCountOk returns a tuple with the OperationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCount

`func (o *CustomerMonthlyUsageDto) SetOperationCount(v int32)`

SetOperationCount sets OperationCount field to given value.

### HasOperationCount

`func (o *CustomerMonthlyUsageDto) HasOperationCount() bool`

HasOperationCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


