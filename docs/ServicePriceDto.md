# ServicePriceDto

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
**TimeBound** | Pointer to [**TimeBoundDto**](TimeBoundDto.md) | The period the price is effective in. | [optional] 
**Status** | Pointer to [**PriceStatus**](PriceStatus.md) | The price status. | [optional] 
**Created** | Pointer to **time.Time** | The date and time when the price was created. | [optional] 
**DiscountCategoryId** | Pointer to **NullableInt32** | The discount category ID. | [optional] 
**DiscountCategory** | Pointer to [**DiscountCategoryDto**](DiscountCategoryDto.md) | The discount category. | [optional] 

## Methods

### NewServicePriceDto

`func NewServicePriceDto() *ServicePriceDto`

NewServicePriceDto instantiates a new ServicePriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePriceDtoWithDefaults

`func NewServicePriceDtoWithDefaults() *ServicePriceDto`

NewServicePriceDtoWithDefaults instantiates a new ServicePriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServicePriceDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicePriceDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicePriceDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServicePriceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *ServicePriceDto) GetAccountNumber() int32`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *ServicePriceDto) GetAccountNumberOk() (*int32, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *ServicePriceDto) SetAccountNumber(v int32)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *ServicePriceDto) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetServiceId

`func (o *ServicePriceDto) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServicePriceDto) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServicePriceDto) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServicePriceDto) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTimeUnit

`func (o *ServicePriceDto) GetTimeUnit() PriceTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *ServicePriceDto) GetTimeUnitOk() (*PriceTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *ServicePriceDto) SetTimeUnit(v PriceTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *ServicePriceDto) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.

### GetCostPrice

`func (o *ServicePriceDto) GetCostPrice() float64`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ServicePriceDto) GetCostPriceOk() (*float64, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ServicePriceDto) SetCostPrice(v float64)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ServicePriceDto) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetExtraCharge

`func (o *ServicePriceDto) GetExtraCharge() float64`

GetExtraCharge returns the ExtraCharge field if non-nil, zero value otherwise.

### GetExtraChargeOk

`func (o *ServicePriceDto) GetExtraChargeOk() (*float64, bool)`

GetExtraChargeOk returns a tuple with the ExtraCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraCharge

`func (o *ServicePriceDto) SetExtraCharge(v float64)`

SetExtraCharge sets ExtraCharge field to given value.

### HasExtraCharge

`func (o *ServicePriceDto) HasExtraCharge() bool`

HasExtraCharge returns a boolean if a field has been set.

### GetServicePrice

`func (o *ServicePriceDto) GetServicePrice() float64`

GetServicePrice returns the ServicePrice field if non-nil, zero value otherwise.

### GetServicePriceOk

`func (o *ServicePriceDto) GetServicePriceOk() (*float64, bool)`

GetServicePriceOk returns a tuple with the ServicePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrice

`func (o *ServicePriceDto) SetServicePrice(v float64)`

SetServicePrice sets ServicePrice field to given value.

### HasServicePrice

`func (o *ServicePriceDto) HasServicePrice() bool`

HasServicePrice returns a boolean if a field has been set.

### GetQuota

`func (o *ServicePriceDto) GetQuota() float64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ServicePriceDto) GetQuotaOk() (*float64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ServicePriceDto) SetQuota(v float64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *ServicePriceDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *ServicePriceDto) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *ServicePriceDto) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetTimeBound

`func (o *ServicePriceDto) GetTimeBound() TimeBoundDto`

GetTimeBound returns the TimeBound field if non-nil, zero value otherwise.

### GetTimeBoundOk

`func (o *ServicePriceDto) GetTimeBoundOk() (*TimeBoundDto, bool)`

GetTimeBoundOk returns a tuple with the TimeBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBound

`func (o *ServicePriceDto) SetTimeBound(v TimeBoundDto)`

SetTimeBound sets TimeBound field to given value.

### HasTimeBound

`func (o *ServicePriceDto) HasTimeBound() bool`

HasTimeBound returns a boolean if a field has been set.

### GetStatus

`func (o *ServicePriceDto) GetStatus() PriceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServicePriceDto) GetStatusOk() (*PriceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServicePriceDto) SetStatus(v PriceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServicePriceDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *ServicePriceDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServicePriceDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServicePriceDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServicePriceDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDiscountCategoryId

`func (o *ServicePriceDto) GetDiscountCategoryId() int32`

GetDiscountCategoryId returns the DiscountCategoryId field if non-nil, zero value otherwise.

### GetDiscountCategoryIdOk

`func (o *ServicePriceDto) GetDiscountCategoryIdOk() (*int32, bool)`

GetDiscountCategoryIdOk returns a tuple with the DiscountCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCategoryId

`func (o *ServicePriceDto) SetDiscountCategoryId(v int32)`

SetDiscountCategoryId sets DiscountCategoryId field to given value.

### HasDiscountCategoryId

`func (o *ServicePriceDto) HasDiscountCategoryId() bool`

HasDiscountCategoryId returns a boolean if a field has been set.

### SetDiscountCategoryIdNil

`func (o *ServicePriceDto) SetDiscountCategoryIdNil(b bool)`

 SetDiscountCategoryIdNil sets the value for DiscountCategoryId to be an explicit nil

### UnsetDiscountCategoryId
`func (o *ServicePriceDto) UnsetDiscountCategoryId()`

UnsetDiscountCategoryId ensures that no value is present for DiscountCategoryId, not even an explicit nil
### GetDiscountCategory

`func (o *ServicePriceDto) GetDiscountCategory() DiscountCategoryDto`

GetDiscountCategory returns the DiscountCategory field if non-nil, zero value otherwise.

### GetDiscountCategoryOk

`func (o *ServicePriceDto) GetDiscountCategoryOk() (*DiscountCategoryDto, bool)`

GetDiscountCategoryOk returns a tuple with the DiscountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCategory

`func (o *ServicePriceDto) SetDiscountCategory(v DiscountCategoryDto)`

SetDiscountCategory sets DiscountCategory field to given value.

### HasDiscountCategory

`func (o *ServicePriceDto) HasDiscountCategory() bool`

HasDiscountCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


