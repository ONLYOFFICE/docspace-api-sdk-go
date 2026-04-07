# TenantEntityQuotaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableQuota** | Pointer to **bool** | Specifies if the quota is enabled for the tenant entity or not. | [optional] 
**DefaultQuota** | Pointer to **int64** | The default quota of the tenant entity. | [optional] 
**LastRecalculateDate** | Pointer to **NullableTime** | The date of the last quota recalculation. | [optional] 

## Methods

### NewTenantEntityQuotaSettings

`func NewTenantEntityQuotaSettings() *TenantEntityQuotaSettings`

NewTenantEntityQuotaSettings instantiates a new TenantEntityQuotaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantEntityQuotaSettingsWithDefaults

`func NewTenantEntityQuotaSettingsWithDefaults() *TenantEntityQuotaSettings`

NewTenantEntityQuotaSettingsWithDefaults instantiates a new TenantEntityQuotaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableQuota

`func (o *TenantEntityQuotaSettings) GetEnableQuota() bool`

GetEnableQuota returns the EnableQuota field if non-nil, zero value otherwise.

### GetEnableQuotaOk

`func (o *TenantEntityQuotaSettings) GetEnableQuotaOk() (*bool, bool)`

GetEnableQuotaOk returns a tuple with the EnableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuota

`func (o *TenantEntityQuotaSettings) SetEnableQuota(v bool)`

SetEnableQuota sets EnableQuota field to given value.

### HasEnableQuota

`func (o *TenantEntityQuotaSettings) HasEnableQuota() bool`

HasEnableQuota returns a boolean if a field has been set.

### GetDefaultQuota

`func (o *TenantEntityQuotaSettings) GetDefaultQuota() int64`

GetDefaultQuota returns the DefaultQuota field if non-nil, zero value otherwise.

### GetDefaultQuotaOk

`func (o *TenantEntityQuotaSettings) GetDefaultQuotaOk() (*int64, bool)`

GetDefaultQuotaOk returns a tuple with the DefaultQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuota

`func (o *TenantEntityQuotaSettings) SetDefaultQuota(v int64)`

SetDefaultQuota sets DefaultQuota field to given value.

### HasDefaultQuota

`func (o *TenantEntityQuotaSettings) HasDefaultQuota() bool`

HasDefaultQuota returns a boolean if a field has been set.

### GetLastRecalculateDate

`func (o *TenantEntityQuotaSettings) GetLastRecalculateDate() time.Time`

GetLastRecalculateDate returns the LastRecalculateDate field if non-nil, zero value otherwise.

### GetLastRecalculateDateOk

`func (o *TenantEntityQuotaSettings) GetLastRecalculateDateOk() (*time.Time, bool)`

GetLastRecalculateDateOk returns a tuple with the LastRecalculateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecalculateDate

`func (o *TenantEntityQuotaSettings) SetLastRecalculateDate(v time.Time)`

SetLastRecalculateDate sets LastRecalculateDate field to given value.

### HasLastRecalculateDate

`func (o *TenantEntityQuotaSettings) HasLastRecalculateDate() bool`

HasLastRecalculateDate returns a boolean if a field has been set.

### SetLastRecalculateDateNil

`func (o *TenantEntityQuotaSettings) SetLastRecalculateDateNil(b bool)`

 SetLastRecalculateDateNil sets the value for LastRecalculateDate to be an explicit nil

### UnsetLastRecalculateDate
`func (o *TenantEntityQuotaSettings) UnsetLastRecalculateDate()`

UnsetLastRecalculateDate ensures that no value is present for LastRecalculateDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


