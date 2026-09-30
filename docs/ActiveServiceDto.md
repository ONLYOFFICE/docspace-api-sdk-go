# ActiveServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **NullableString** | The stable key of the service, which is what `POST api/2.0/portal/payment/servicestate` takes to switch  it off again. | [optional] 
**ServiceUnit** | Pointer to **NullableString** | What `limit` and `used` count, in the portal language - gigabytes, editor seats, credits. | [optional] 
**Subscription** | Pointer to **bool** | Whether the service is billed as a standing subscription rather than per unit consumed. Only a  subscription can carry `limit` and `used`. | [optional] 
**Title** | Pointer to **NullableString** | The service name in the portal language, for printing rather than matching. | [optional] 
**Limit** | Pointer to **NullableInt32** | How much of the service the portal is entitled to. It is empty for a service whose consumption is not  counted this way, which is not the same as a service without a limit. | [optional] 
**Used** | Pointer to **NullableInt32** | How much of that allowance is in use - the editors currently active for the cloud editors, the units  already consumed for disk storage. Empty under the same conditions as `limit`. | [optional] 

## Methods

### NewActiveServiceDto

`func NewActiveServiceDto() *ActiveServiceDto`

NewActiveServiceDto instantiates a new ActiveServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveServiceDtoWithDefaults

`func NewActiveServiceDtoWithDefaults() *ActiveServiceDto`

NewActiveServiceDtoWithDefaults instantiates a new ActiveServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ActiveServiceDto) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ActiveServiceDto) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ActiveServiceDto) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ActiveServiceDto) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *ActiveServiceDto) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *ActiveServiceDto) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetServiceUnit

`func (o *ActiveServiceDto) GetServiceUnit() string`

GetServiceUnit returns the ServiceUnit field if non-nil, zero value otherwise.

### GetServiceUnitOk

`func (o *ActiveServiceDto) GetServiceUnitOk() (*string, bool)`

GetServiceUnitOk returns a tuple with the ServiceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUnit

`func (o *ActiveServiceDto) SetServiceUnit(v string)`

SetServiceUnit sets ServiceUnit field to given value.

### HasServiceUnit

`func (o *ActiveServiceDto) HasServiceUnit() bool`

HasServiceUnit returns a boolean if a field has been set.

### SetServiceUnitNil

`func (o *ActiveServiceDto) SetServiceUnitNil(b bool)`

 SetServiceUnitNil sets the value for ServiceUnit to be an explicit nil

### UnsetServiceUnit
`func (o *ActiveServiceDto) UnsetServiceUnit()`

UnsetServiceUnit ensures that no value is present for ServiceUnit, not even an explicit nil
### GetSubscription

`func (o *ActiveServiceDto) GetSubscription() bool`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ActiveServiceDto) GetSubscriptionOk() (*bool, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ActiveServiceDto) SetSubscription(v bool)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *ActiveServiceDto) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTitle

`func (o *ActiveServiceDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ActiveServiceDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ActiveServiceDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ActiveServiceDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ActiveServiceDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ActiveServiceDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLimit

`func (o *ActiveServiceDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ActiveServiceDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ActiveServiceDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ActiveServiceDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ActiveServiceDto) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ActiveServiceDto) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetUsed

`func (o *ActiveServiceDto) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *ActiveServiceDto) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *ActiveServiceDto) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *ActiveServiceDto) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### SetUsedNil

`func (o *ActiveServiceDto) SetUsedNil(b bool)`

 SetUsedNil sets the value for Used to be an explicit nil

### UnsetUsed
`func (o *ActiveServiceDto) UnsetUsed()`

UnsetUsed ensures that no value is present for Used, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


