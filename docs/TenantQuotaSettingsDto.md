# TenantQuotaSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableQuota** | Pointer to **bool** | Specifies if the tenant quota is enabled or not. | [optional] 
**Quota** | Pointer to **int64** | The tenant quota. | [optional] 
**LastRecalculateDate** | Pointer to **NullableTime** | The date of the last tenant quota recalculation. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewTenantQuotaSettingsDto

`func NewTenantQuotaSettingsDto() *TenantQuotaSettingsDto`

NewTenantQuotaSettingsDto instantiates a new TenantQuotaSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantQuotaSettingsDtoWithDefaults

`func NewTenantQuotaSettingsDtoWithDefaults() *TenantQuotaSettingsDto`

NewTenantQuotaSettingsDtoWithDefaults instantiates a new TenantQuotaSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableQuota

`func (o *TenantQuotaSettingsDto) GetEnableQuota() bool`

GetEnableQuota returns the EnableQuota field if non-nil, zero value otherwise.

### GetEnableQuotaOk

`func (o *TenantQuotaSettingsDto) GetEnableQuotaOk() (*bool, bool)`

GetEnableQuotaOk returns a tuple with the EnableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuota

`func (o *TenantQuotaSettingsDto) SetEnableQuota(v bool)`

SetEnableQuota sets EnableQuota field to given value.

### HasEnableQuota

`func (o *TenantQuotaSettingsDto) HasEnableQuota() bool`

HasEnableQuota returns a boolean if a field has been set.

### GetQuota

`func (o *TenantQuotaSettingsDto) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *TenantQuotaSettingsDto) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *TenantQuotaSettingsDto) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *TenantQuotaSettingsDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetLastRecalculateDate

`func (o *TenantQuotaSettingsDto) GetLastRecalculateDate() time.Time`

GetLastRecalculateDate returns the LastRecalculateDate field if non-nil, zero value otherwise.

### GetLastRecalculateDateOk

`func (o *TenantQuotaSettingsDto) GetLastRecalculateDateOk() (*time.Time, bool)`

GetLastRecalculateDateOk returns a tuple with the LastRecalculateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecalculateDate

`func (o *TenantQuotaSettingsDto) SetLastRecalculateDate(v time.Time)`

SetLastRecalculateDate sets LastRecalculateDate field to given value.

### HasLastRecalculateDate

`func (o *TenantQuotaSettingsDto) HasLastRecalculateDate() bool`

HasLastRecalculateDate returns a boolean if a field has been set.

### SetLastRecalculateDateNil

`func (o *TenantQuotaSettingsDto) SetLastRecalculateDateNil(b bool)`

 SetLastRecalculateDateNil sets the value for LastRecalculateDate to be an explicit nil

### UnsetLastRecalculateDate
`func (o *TenantQuotaSettingsDto) UnsetLastRecalculateDate()`

UnsetLastRecalculateDate ensures that no value is present for LastRecalculateDate, not even an explicit nil
### GetLastModified

`func (o *TenantQuotaSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantQuotaSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantQuotaSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantQuotaSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


