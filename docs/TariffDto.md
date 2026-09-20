# TariffDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenSource** | Pointer to **NullableBool** | Whether the installation runs the open-source build, which has no paid plan at all. This flag and the two  below describe the build rather than the subscription, and all three are left empty for a caller without  the portal-settings right. | [optional] 
**Enterprise** | Pointer to **NullableBool** | Whether the installation runs on an Enterprise licence file, which is what makes the licence operations  under `api/2.0/settings/license` usable. | [optional] 
**Developer** | Pointer to **NullableBool** | Whether the installation runs on a Developer licence, an Enterprise licence meant for embedding rather  than for production use. | [optional] 
**Id** | Pointer to **int32** | The identifier of the subscription record itself, for quoting when a charge has to be traced. It is filled  in for a caller with the portal-settings right only, and nothing accepts it as an argument. | [optional] 
**State** | Pointer to [**TariffState**](TariffState.md) | How the subscription stands: on trial, paid, inside the grace period that follows the due date, or unpaid.  It is the one field every caller gets, whatever their role, so a client can warn about payment without  needing administrator rights. | [optional] 
**DueDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the current period ends, in the portal time zone. It is filled in for a room or DocSpace  administrator only, and set to the largest value a date can hold for a subscription that never ends. | [optional] 
**DelayDueDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the grace period after `dueDate` runs out and the portal is cut off, in the portal time zone. Filled  in under the same conditions as `dueDate`, and equal to it when the plan grants no grace period. | [optional] 
**LicenseDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the licence file behind the subscription was issued, in the portal time zone. It is meaningful on a  server installation and filled in for a caller with the portal-settings right only. | [optional] 
**CustomerId** | Pointer to **NullableString** | The account in the billing system the subscription is charged to, empty for a portal that has never been  billed. Filled in for a caller with the portal-settings right only. | [optional] 
**Quotas** | Pointer to [**[]TariffQuotaDto**](TariffQuotaDto.md) | The quotas the subscription is made of - the plan itself and its add-ons - with the overdue ones listed  alongside the current ones, so an entry here is not proof that it is still being paid for; read each  entry's own `state` for that. Filled in for a caller with the portal-settings right only. | [optional] 

## Methods

### NewTariffDto

`func NewTariffDto() *TariffDto`

NewTariffDto instantiates a new TariffDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffDtoWithDefaults

`func NewTariffDtoWithDefaults() *TariffDto`

NewTariffDtoWithDefaults instantiates a new TariffDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenSource

`func (o *TariffDto) GetOpenSource() bool`

GetOpenSource returns the OpenSource field if non-nil, zero value otherwise.

### GetOpenSourceOk

`func (o *TariffDto) GetOpenSourceOk() (*bool, bool)`

GetOpenSourceOk returns a tuple with the OpenSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSource

`func (o *TariffDto) SetOpenSource(v bool)`

SetOpenSource sets OpenSource field to given value.

### HasOpenSource

`func (o *TariffDto) HasOpenSource() bool`

HasOpenSource returns a boolean if a field has been set.

### SetOpenSourceNil

`func (o *TariffDto) SetOpenSourceNil(b bool)`

 SetOpenSourceNil sets the value for OpenSource to be an explicit nil

### UnsetOpenSource
`func (o *TariffDto) UnsetOpenSource()`

UnsetOpenSource ensures that no value is present for OpenSource, not even an explicit nil
### GetEnterprise

`func (o *TariffDto) GetEnterprise() bool`

GetEnterprise returns the Enterprise field if non-nil, zero value otherwise.

### GetEnterpriseOk

`func (o *TariffDto) GetEnterpriseOk() (*bool, bool)`

GetEnterpriseOk returns a tuple with the Enterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprise

`func (o *TariffDto) SetEnterprise(v bool)`

SetEnterprise sets Enterprise field to given value.

### HasEnterprise

`func (o *TariffDto) HasEnterprise() bool`

HasEnterprise returns a boolean if a field has been set.

### SetEnterpriseNil

`func (o *TariffDto) SetEnterpriseNil(b bool)`

 SetEnterpriseNil sets the value for Enterprise to be an explicit nil

### UnsetEnterprise
`func (o *TariffDto) UnsetEnterprise()`

UnsetEnterprise ensures that no value is present for Enterprise, not even an explicit nil
### GetDeveloper

`func (o *TariffDto) GetDeveloper() bool`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *TariffDto) GetDeveloperOk() (*bool, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *TariffDto) SetDeveloper(v bool)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *TariffDto) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloperNil

`func (o *TariffDto) SetDeveloperNil(b bool)`

 SetDeveloperNil sets the value for Developer to be an explicit nil

### UnsetDeveloper
`func (o *TariffDto) UnsetDeveloper()`

UnsetDeveloper ensures that no value is present for Developer, not even an explicit nil
### GetId

`func (o *TariffDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TariffDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TariffDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TariffDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *TariffDto) GetState() TariffState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TariffDto) GetStateOk() (*TariffState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TariffDto) SetState(v TariffState)`

SetState sets State field to given value.

### HasState

`func (o *TariffDto) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDueDate

`func (o *TariffDto) GetDueDate() ApiDateTime`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TariffDto) GetDueDateOk() (*ApiDateTime, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TariffDto) SetDueDate(v ApiDateTime)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TariffDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetDelayDueDate

`func (o *TariffDto) GetDelayDueDate() ApiDateTime`

GetDelayDueDate returns the DelayDueDate field if non-nil, zero value otherwise.

### GetDelayDueDateOk

`func (o *TariffDto) GetDelayDueDateOk() (*ApiDateTime, bool)`

GetDelayDueDateOk returns a tuple with the DelayDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDueDate

`func (o *TariffDto) SetDelayDueDate(v ApiDateTime)`

SetDelayDueDate sets DelayDueDate field to given value.

### HasDelayDueDate

`func (o *TariffDto) HasDelayDueDate() bool`

HasDelayDueDate returns a boolean if a field has been set.

### GetLicenseDate

`func (o *TariffDto) GetLicenseDate() ApiDateTime`

GetLicenseDate returns the LicenseDate field if non-nil, zero value otherwise.

### GetLicenseDateOk

`func (o *TariffDto) GetLicenseDateOk() (*ApiDateTime, bool)`

GetLicenseDateOk returns a tuple with the LicenseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseDate

`func (o *TariffDto) SetLicenseDate(v ApiDateTime)`

SetLicenseDate sets LicenseDate field to given value.

### HasLicenseDate

`func (o *TariffDto) HasLicenseDate() bool`

HasLicenseDate returns a boolean if a field has been set.

### GetCustomerId

`func (o *TariffDto) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TariffDto) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TariffDto) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *TariffDto) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *TariffDto) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *TariffDto) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetQuotas

`func (o *TariffDto) GetQuotas() []TariffQuotaDto`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *TariffDto) GetQuotasOk() (*[]TariffQuotaDto, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *TariffDto) SetQuotas(v []TariffQuotaDto)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *TariffDto) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### SetQuotasNil

`func (o *TariffDto) SetQuotasNil(b bool)`

 SetQuotasNil sets the value for Quotas to be an explicit nil

### UnsetQuotas
`func (o *TariffDto) UnsetQuotas()`

UnsetQuotas ensures that no value is present for Quotas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


