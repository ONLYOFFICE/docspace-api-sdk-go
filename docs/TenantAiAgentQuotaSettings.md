# TenantAiAgentQuotaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableQuota** | Pointer to **bool** | Specifies if the quota is enabled for the tenant entity or not. | [optional] 
**DefaultQuota** | Pointer to **int64** | The default quota of the tenant entity. | [optional] 
**LastRecalculateDate** | Pointer to **time.Time** | The date of the last quota recalculation. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewTenantAiAgentQuotaSettings

`func NewTenantAiAgentQuotaSettings() *TenantAiAgentQuotaSettings`

NewTenantAiAgentQuotaSettings instantiates a new TenantAiAgentQuotaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAiAgentQuotaSettingsWithDefaults

`func NewTenantAiAgentQuotaSettingsWithDefaults() *TenantAiAgentQuotaSettings`

NewTenantAiAgentQuotaSettingsWithDefaults instantiates a new TenantAiAgentQuotaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableQuota

`func (o *TenantAiAgentQuotaSettings) GetEnableQuota() bool`

GetEnableQuota returns the EnableQuota field if non-nil, zero value otherwise.

### GetEnableQuotaOk

`func (o *TenantAiAgentQuotaSettings) GetEnableQuotaOk() (*bool, bool)`

GetEnableQuotaOk returns a tuple with the EnableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuota

`func (o *TenantAiAgentQuotaSettings) SetEnableQuota(v bool)`

SetEnableQuota sets EnableQuota field to given value.

### HasEnableQuota

`func (o *TenantAiAgentQuotaSettings) HasEnableQuota() bool`

HasEnableQuota returns a boolean if a field has been set.

### GetDefaultQuota

`func (o *TenantAiAgentQuotaSettings) GetDefaultQuota() int64`

GetDefaultQuota returns the DefaultQuota field if non-nil, zero value otherwise.

### GetDefaultQuotaOk

`func (o *TenantAiAgentQuotaSettings) GetDefaultQuotaOk() (*int64, bool)`

GetDefaultQuotaOk returns a tuple with the DefaultQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuota

`func (o *TenantAiAgentQuotaSettings) SetDefaultQuota(v int64)`

SetDefaultQuota sets DefaultQuota field to given value.

### HasDefaultQuota

`func (o *TenantAiAgentQuotaSettings) HasDefaultQuota() bool`

HasDefaultQuota returns a boolean if a field has been set.

### GetLastRecalculateDate

`func (o *TenantAiAgentQuotaSettings) GetLastRecalculateDate() time.Time`

GetLastRecalculateDate returns the LastRecalculateDate field if non-nil, zero value otherwise.

### GetLastRecalculateDateOk

`func (o *TenantAiAgentQuotaSettings) GetLastRecalculateDateOk() (*time.Time, bool)`

GetLastRecalculateDateOk returns a tuple with the LastRecalculateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecalculateDate

`func (o *TenantAiAgentQuotaSettings) SetLastRecalculateDate(v time.Time)`

SetLastRecalculateDate sets LastRecalculateDate field to given value.

### HasLastRecalculateDate

`func (o *TenantAiAgentQuotaSettings) HasLastRecalculateDate() bool`

HasLastRecalculateDate returns a boolean if a field has been set.

### GetLastModified

`func (o *TenantAiAgentQuotaSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantAiAgentQuotaSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantAiAgentQuotaSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantAiAgentQuotaSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


