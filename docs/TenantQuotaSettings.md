# TenantQuotaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableQuota** | Pointer to **bool** | Specifies if the tenant quota is enabled or not. | [optional] 
**Quota** | Pointer to **int64** | The tenant quota. | [optional] 
**LastRecalculateDate** | Pointer to **NullableTime** | The date of the last tenant quota recalculation. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewTenantQuotaSettings

`func NewTenantQuotaSettings() *TenantQuotaSettings`

NewTenantQuotaSettings instantiates a new TenantQuotaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantQuotaSettingsWithDefaults

`func NewTenantQuotaSettingsWithDefaults() *TenantQuotaSettings`

NewTenantQuotaSettingsWithDefaults instantiates a new TenantQuotaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableQuota

`func (o *TenantQuotaSettings) GetEnableQuota() bool`

GetEnableQuota returns the EnableQuota field if non-nil, zero value otherwise.

### GetEnableQuotaOk

`func (o *TenantQuotaSettings) GetEnableQuotaOk() (*bool, bool)`

GetEnableQuotaOk returns a tuple with the EnableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuota

`func (o *TenantQuotaSettings) SetEnableQuota(v bool)`

SetEnableQuota sets EnableQuota field to given value.

### HasEnableQuota

`func (o *TenantQuotaSettings) HasEnableQuota() bool`

HasEnableQuota returns a boolean if a field has been set.

### GetQuota

`func (o *TenantQuotaSettings) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *TenantQuotaSettings) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *TenantQuotaSettings) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *TenantQuotaSettings) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetLastRecalculateDate

`func (o *TenantQuotaSettings) GetLastRecalculateDate() time.Time`

GetLastRecalculateDate returns the LastRecalculateDate field if non-nil, zero value otherwise.

### GetLastRecalculateDateOk

`func (o *TenantQuotaSettings) GetLastRecalculateDateOk() (*time.Time, bool)`

GetLastRecalculateDateOk returns a tuple with the LastRecalculateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecalculateDate

`func (o *TenantQuotaSettings) SetLastRecalculateDate(v time.Time)`

SetLastRecalculateDate sets LastRecalculateDate field to given value.

### HasLastRecalculateDate

`func (o *TenantQuotaSettings) HasLastRecalculateDate() bool`

HasLastRecalculateDate returns a boolean if a field has been set.

### SetLastRecalculateDateNil

`func (o *TenantQuotaSettings) SetLastRecalculateDateNil(b bool)`

 SetLastRecalculateDateNil sets the value for LastRecalculateDate to be an explicit nil

### UnsetLastRecalculateDate
`func (o *TenantQuotaSettings) UnsetLastRecalculateDate()`

UnsetLastRecalculateDate ensures that no value is present for LastRecalculateDate, not even an explicit nil
### GetLastModified

`func (o *TenantQuotaSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantQuotaSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantQuotaSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantQuotaSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


