# ServicePriceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The price unique identifier. | [optional] 
**AccountNumber** | Pointer to **int32** | The account number. | [optional] 
**ServiceId** | Pointer to **int32** | The service ID. | [optional] 
**TimeUnit** | Pointer to [**PriceTimeUnit**](PriceTimeUnit.md) | The time unit the price is bound to. | [optional] 
**CostPrice** | Pointer to **float64** | The cost price. | [optional] 
**ExtraCharge** | Pointer to **float64** | The extra charge added to the cost price. | [optional] 
**ServicePrice** | Pointer to **float64** | The resulting service price. | [optional] 
**Quota** | Pointer to **NullableFloat64** | The quota the price is set for. | [optional] 
**TimeBound** | Pointer to [**TimeBound**](TimeBound.md) | The period the price is effective in. | [optional] 
**Status** | Pointer to [**PriceStatus**](PriceStatus.md) | The price status. | [optional] 
**Created** | Pointer to **time.Time** | The date and time when the price was created. | [optional] 
**DiscountCategoryId** | Pointer to **NullableInt32** | The discount category ID. | [optional] 
**DiscountCategory** | Pointer to [**DiscountCategory**](DiscountCategory.md) | The discount category. | [optional] 

## Methods

### NewServicePriceInfo

`func NewServicePriceInfo() *ServicePriceInfo`

NewServicePriceInfo instantiates a new ServicePriceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePriceInfoWithDefaults

`func NewServicePriceInfoWithDefaults() *ServicePriceInfo`

NewServicePriceInfoWithDefaults instantiates a new ServicePriceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServicePriceInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicePriceInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicePriceInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServicePriceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *ServicePriceInfo) GetAccountNumber() int32`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *ServicePriceInfo) GetAccountNumberOk() (*int32, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *ServicePriceInfo) SetAccountNumber(v int32)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *ServicePriceInfo) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetServiceId

`func (o *ServicePriceInfo) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServicePriceInfo) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServicePriceInfo) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServicePriceInfo) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTimeUnit

`func (o *ServicePriceInfo) GetTimeUnit() PriceTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *ServicePriceInfo) GetTimeUnitOk() (*PriceTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *ServicePriceInfo) SetTimeUnit(v PriceTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *ServicePriceInfo) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.

### GetCostPrice

`func (o *ServicePriceInfo) GetCostPrice() float64`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ServicePriceInfo) GetCostPriceOk() (*float64, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ServicePriceInfo) SetCostPrice(v float64)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ServicePriceInfo) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetExtraCharge

`func (o *ServicePriceInfo) GetExtraCharge() float64`

GetExtraCharge returns the ExtraCharge field if non-nil, zero value otherwise.

### GetExtraChargeOk

`func (o *ServicePriceInfo) GetExtraChargeOk() (*float64, bool)`

GetExtraChargeOk returns a tuple with the ExtraCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraCharge

`func (o *ServicePriceInfo) SetExtraCharge(v float64)`

SetExtraCharge sets ExtraCharge field to given value.

### HasExtraCharge

`func (o *ServicePriceInfo) HasExtraCharge() bool`

HasExtraCharge returns a boolean if a field has been set.

### GetServicePrice

`func (o *ServicePriceInfo) GetServicePrice() float64`

GetServicePrice returns the ServicePrice field if non-nil, zero value otherwise.

### GetServicePriceOk

`func (o *ServicePriceInfo) GetServicePriceOk() (*float64, bool)`

GetServicePriceOk returns a tuple with the ServicePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrice

`func (o *ServicePriceInfo) SetServicePrice(v float64)`

SetServicePrice sets ServicePrice field to given value.

### HasServicePrice

`func (o *ServicePriceInfo) HasServicePrice() bool`

HasServicePrice returns a boolean if a field has been set.

### GetQuota

`func (o *ServicePriceInfo) GetQuota() float64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ServicePriceInfo) GetQuotaOk() (*float64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ServicePriceInfo) SetQuota(v float64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *ServicePriceInfo) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *ServicePriceInfo) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *ServicePriceInfo) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetTimeBound

`func (o *ServicePriceInfo) GetTimeBound() TimeBound`

GetTimeBound returns the TimeBound field if non-nil, zero value otherwise.

### GetTimeBoundOk

`func (o *ServicePriceInfo) GetTimeBoundOk() (*TimeBound, bool)`

GetTimeBoundOk returns a tuple with the TimeBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBound

`func (o *ServicePriceInfo) SetTimeBound(v TimeBound)`

SetTimeBound sets TimeBound field to given value.

### HasTimeBound

`func (o *ServicePriceInfo) HasTimeBound() bool`

HasTimeBound returns a boolean if a field has been set.

### GetStatus

`func (o *ServicePriceInfo) GetStatus() PriceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServicePriceInfo) GetStatusOk() (*PriceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServicePriceInfo) SetStatus(v PriceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServicePriceInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *ServicePriceInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServicePriceInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServicePriceInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServicePriceInfo) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDiscountCategoryId

`func (o *ServicePriceInfo) GetDiscountCategoryId() int32`

GetDiscountCategoryId returns the DiscountCategoryId field if non-nil, zero value otherwise.

### GetDiscountCategoryIdOk

`func (o *ServicePriceInfo) GetDiscountCategoryIdOk() (*int32, bool)`

GetDiscountCategoryIdOk returns a tuple with the DiscountCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCategoryId

`func (o *ServicePriceInfo) SetDiscountCategoryId(v int32)`

SetDiscountCategoryId sets DiscountCategoryId field to given value.

### HasDiscountCategoryId

`func (o *ServicePriceInfo) HasDiscountCategoryId() bool`

HasDiscountCategoryId returns a boolean if a field has been set.

### SetDiscountCategoryIdNil

`func (o *ServicePriceInfo) SetDiscountCategoryIdNil(b bool)`

 SetDiscountCategoryIdNil sets the value for DiscountCategoryId to be an explicit nil

### UnsetDiscountCategoryId
`func (o *ServicePriceInfo) UnsetDiscountCategoryId()`

UnsetDiscountCategoryId ensures that no value is present for DiscountCategoryId, not even an explicit nil
### GetDiscountCategory

`func (o *ServicePriceInfo) GetDiscountCategory() DiscountCategory`

GetDiscountCategory returns the DiscountCategory field if non-nil, zero value otherwise.

### GetDiscountCategoryOk

`func (o *ServicePriceInfo) GetDiscountCategoryOk() (*DiscountCategory, bool)`

GetDiscountCategoryOk returns a tuple with the DiscountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCategory

`func (o *ServicePriceInfo) SetDiscountCategory(v DiscountCategory)`

SetDiscountCategory sets DiscountCategory field to given value.

### HasDiscountCategory

`func (o *ServicePriceInfo) HasDiscountCategory() bool`

HasDiscountCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


