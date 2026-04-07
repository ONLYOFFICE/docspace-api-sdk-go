# Tariff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The tariff ID. | [optional] 
**State** | Pointer to [**TariffState**](TariffState.md) |  | [optional] 
**DueDate** | **time.Time** | The tariff due date. | 
**DelayDueDate** | Pointer to **time.Time** | The tariff delay due date. | [optional] 
**LicenseDate** | Pointer to **time.Time** | The tariff license date. | [optional] 
**CustomerId** | Pointer to **NullableString** | The tariff customer ID. | [optional] 
**Quotas** | [**[]Quota**](Quota.md) | The list of tariff quotas. | 
**OverdueQuotas** | Pointer to [**[]Quota**](Quota.md) | The list of overdue tariff quotas. | [optional] 

## Methods

### NewTariff

`func NewTariff(dueDate time.Time, quotas []Quota, ) *Tariff`

NewTariff instantiates a new Tariff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffWithDefaults

`func NewTariffWithDefaults() *Tariff`

NewTariffWithDefaults instantiates a new Tariff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tariff) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tariff) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tariff) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Tariff) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *Tariff) GetState() TariffState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Tariff) GetStateOk() (*TariffState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Tariff) SetState(v TariffState)`

SetState sets State field to given value.

### HasState

`func (o *Tariff) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDueDate

`func (o *Tariff) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Tariff) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Tariff) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.


### GetDelayDueDate

`func (o *Tariff) GetDelayDueDate() time.Time`

GetDelayDueDate returns the DelayDueDate field if non-nil, zero value otherwise.

### GetDelayDueDateOk

`func (o *Tariff) GetDelayDueDateOk() (*time.Time, bool)`

GetDelayDueDateOk returns a tuple with the DelayDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDueDate

`func (o *Tariff) SetDelayDueDate(v time.Time)`

SetDelayDueDate sets DelayDueDate field to given value.

### HasDelayDueDate

`func (o *Tariff) HasDelayDueDate() bool`

HasDelayDueDate returns a boolean if a field has been set.

### GetLicenseDate

`func (o *Tariff) GetLicenseDate() time.Time`

GetLicenseDate returns the LicenseDate field if non-nil, zero value otherwise.

### GetLicenseDateOk

`func (o *Tariff) GetLicenseDateOk() (*time.Time, bool)`

GetLicenseDateOk returns a tuple with the LicenseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseDate

`func (o *Tariff) SetLicenseDate(v time.Time)`

SetLicenseDate sets LicenseDate field to given value.

### HasLicenseDate

`func (o *Tariff) HasLicenseDate() bool`

HasLicenseDate returns a boolean if a field has been set.

### GetCustomerId

`func (o *Tariff) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Tariff) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Tariff) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Tariff) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *Tariff) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *Tariff) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetQuotas

`func (o *Tariff) GetQuotas() []Quota`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *Tariff) GetQuotasOk() (*[]Quota, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *Tariff) SetQuotas(v []Quota)`

SetQuotas sets Quotas field to given value.


### SetQuotasNil

`func (o *Tariff) SetQuotasNil(b bool)`

 SetQuotasNil sets the value for Quotas to be an explicit nil

### UnsetQuotas
`func (o *Tariff) UnsetQuotas()`

UnsetQuotas ensures that no value is present for Quotas, not even an explicit nil
### GetOverdueQuotas

`func (o *Tariff) GetOverdueQuotas() []Quota`

GetOverdueQuotas returns the OverdueQuotas field if non-nil, zero value otherwise.

### GetOverdueQuotasOk

`func (o *Tariff) GetOverdueQuotasOk() (*[]Quota, bool)`

GetOverdueQuotasOk returns a tuple with the OverdueQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueQuotas

`func (o *Tariff) SetOverdueQuotas(v []Quota)`

SetOverdueQuotas sets OverdueQuotas field to given value.

### HasOverdueQuotas

`func (o *Tariff) HasOverdueQuotas() bool`

HasOverdueQuotas returns a boolean if a field has been set.

### SetOverdueQuotasNil

`func (o *Tariff) SetOverdueQuotasNil(b bool)`

 SetOverdueQuotasNil sets the value for OverdueQuotas to be an explicit nil

### UnsetOverdueQuotas
`func (o *Tariff) UnsetOverdueQuotas()`

UnsetOverdueQuotas ensures that no value is present for OverdueQuotas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


