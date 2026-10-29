# ServiceLimitUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodStart** | Pointer to **time.Time** | The start of the current period, in UTC. | [optional] 
**PeriodEnd** | Pointer to **NullableTime** | The end of the current period, in UTC: the moment the spending starts again from zero. Null for a lifetime  limit, which never resets. | [optional] 
**AmountConsumed** | Pointer to **float64** | The money charged in the current period, in the currency of the limit. | [optional] 
**AmountAvailable** | Pointer to **NullableFloat64** | The money still available in the current period: the threshold less what was charged and what is reserved for  operations in progress, never below zero. Null when the limit has no money threshold. | [optional] 
**QuantityConsumed** | Pointer to **int32** | The quantity consumed in the current period, in the unit of the limit. | [optional] 
**QuantityAvailable** | Pointer to **NullableInt32** | The quantity still available in the current period, never below zero; null when the limit has no quantity  threshold. | [optional] 

## Methods

### NewServiceLimitUsageDto

`func NewServiceLimitUsageDto() *ServiceLimitUsageDto`

NewServiceLimitUsageDto instantiates a new ServiceLimitUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLimitUsageDtoWithDefaults

`func NewServiceLimitUsageDtoWithDefaults() *ServiceLimitUsageDto`

NewServiceLimitUsageDtoWithDefaults instantiates a new ServiceLimitUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodStart

`func (o *ServiceLimitUsageDto) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *ServiceLimitUsageDto) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *ServiceLimitUsageDto) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *ServiceLimitUsageDto) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *ServiceLimitUsageDto) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ServiceLimitUsageDto) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ServiceLimitUsageDto) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *ServiceLimitUsageDto) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### SetPeriodEndNil

`func (o *ServiceLimitUsageDto) SetPeriodEndNil(b bool)`

 SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil

### UnsetPeriodEnd
`func (o *ServiceLimitUsageDto) UnsetPeriodEnd()`

UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
### GetAmountConsumed

`func (o *ServiceLimitUsageDto) GetAmountConsumed() float64`

GetAmountConsumed returns the AmountConsumed field if non-nil, zero value otherwise.

### GetAmountConsumedOk

`func (o *ServiceLimitUsageDto) GetAmountConsumedOk() (*float64, bool)`

GetAmountConsumedOk returns a tuple with the AmountConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountConsumed

`func (o *ServiceLimitUsageDto) SetAmountConsumed(v float64)`

SetAmountConsumed sets AmountConsumed field to given value.

### HasAmountConsumed

`func (o *ServiceLimitUsageDto) HasAmountConsumed() bool`

HasAmountConsumed returns a boolean if a field has been set.

### GetAmountAvailable

`func (o *ServiceLimitUsageDto) GetAmountAvailable() float64`

GetAmountAvailable returns the AmountAvailable field if non-nil, zero value otherwise.

### GetAmountAvailableOk

`func (o *ServiceLimitUsageDto) GetAmountAvailableOk() (*float64, bool)`

GetAmountAvailableOk returns a tuple with the AmountAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAvailable

`func (o *ServiceLimitUsageDto) SetAmountAvailable(v float64)`

SetAmountAvailable sets AmountAvailable field to given value.

### HasAmountAvailable

`func (o *ServiceLimitUsageDto) HasAmountAvailable() bool`

HasAmountAvailable returns a boolean if a field has been set.

### SetAmountAvailableNil

`func (o *ServiceLimitUsageDto) SetAmountAvailableNil(b bool)`

 SetAmountAvailableNil sets the value for AmountAvailable to be an explicit nil

### UnsetAmountAvailable
`func (o *ServiceLimitUsageDto) UnsetAmountAvailable()`

UnsetAmountAvailable ensures that no value is present for AmountAvailable, not even an explicit nil
### GetQuantityConsumed

`func (o *ServiceLimitUsageDto) GetQuantityConsumed() int32`

GetQuantityConsumed returns the QuantityConsumed field if non-nil, zero value otherwise.

### GetQuantityConsumedOk

`func (o *ServiceLimitUsageDto) GetQuantityConsumedOk() (*int32, bool)`

GetQuantityConsumedOk returns a tuple with the QuantityConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityConsumed

`func (o *ServiceLimitUsageDto) SetQuantityConsumed(v int32)`

SetQuantityConsumed sets QuantityConsumed field to given value.

### HasQuantityConsumed

`func (o *ServiceLimitUsageDto) HasQuantityConsumed() bool`

HasQuantityConsumed returns a boolean if a field has been set.

### GetQuantityAvailable

`func (o *ServiceLimitUsageDto) GetQuantityAvailable() int32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *ServiceLimitUsageDto) GetQuantityAvailableOk() (*int32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *ServiceLimitUsageDto) SetQuantityAvailable(v int32)`

SetQuantityAvailable sets QuantityAvailable field to given value.

### HasQuantityAvailable

`func (o *ServiceLimitUsageDto) HasQuantityAvailable() bool`

HasQuantityAvailable returns a boolean if a field has been set.

### SetQuantityAvailableNil

`func (o *ServiceLimitUsageDto) SetQuantityAvailableNil(b bool)`

 SetQuantityAvailableNil sets the value for QuantityAvailable to be an explicit nil

### UnsetQuantityAvailable
`func (o *ServiceLimitUsageDto) UnsetQuantityAvailable()`

UnsetQuantityAvailable ensures that no value is present for QuantityAvailable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


