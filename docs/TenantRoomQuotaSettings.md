# TenantRoomQuotaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableQuota** | Pointer to **bool** | Specifies if the quota is enabled for the tenant entity or not. | [optional] 
**DefaultQuota** | Pointer to **int64** | The default quota of the tenant entity. | [optional] 
**LastRecalculateDate** | Pointer to **NullableTime** | The date of the last quota recalculation. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewTenantRoomQuotaSettings

`func NewTenantRoomQuotaSettings() *TenantRoomQuotaSettings`

NewTenantRoomQuotaSettings instantiates a new TenantRoomQuotaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantRoomQuotaSettingsWithDefaults

`func NewTenantRoomQuotaSettingsWithDefaults() *TenantRoomQuotaSettings`

NewTenantRoomQuotaSettingsWithDefaults instantiates a new TenantRoomQuotaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableQuota

`func (o *TenantRoomQuotaSettings) GetEnableQuota() bool`

GetEnableQuota returns the EnableQuota field if non-nil, zero value otherwise.

### GetEnableQuotaOk

`func (o *TenantRoomQuotaSettings) GetEnableQuotaOk() (*bool, bool)`

GetEnableQuotaOk returns a tuple with the EnableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuota

`func (o *TenantRoomQuotaSettings) SetEnableQuota(v bool)`

SetEnableQuota sets EnableQuota field to given value.

### HasEnableQuota

`func (o *TenantRoomQuotaSettings) HasEnableQuota() bool`

HasEnableQuota returns a boolean if a field has been set.

### GetDefaultQuota

`func (o *TenantRoomQuotaSettings) GetDefaultQuota() int64`

GetDefaultQuota returns the DefaultQuota field if non-nil, zero value otherwise.

### GetDefaultQuotaOk

`func (o *TenantRoomQuotaSettings) GetDefaultQuotaOk() (*int64, bool)`

GetDefaultQuotaOk returns a tuple with the DefaultQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuota

`func (o *TenantRoomQuotaSettings) SetDefaultQuota(v int64)`

SetDefaultQuota sets DefaultQuota field to given value.

### HasDefaultQuota

`func (o *TenantRoomQuotaSettings) HasDefaultQuota() bool`

HasDefaultQuota returns a boolean if a field has been set.

### GetLastRecalculateDate

`func (o *TenantRoomQuotaSettings) GetLastRecalculateDate() time.Time`

GetLastRecalculateDate returns the LastRecalculateDate field if non-nil, zero value otherwise.

### GetLastRecalculateDateOk

`func (o *TenantRoomQuotaSettings) GetLastRecalculateDateOk() (*time.Time, bool)`

GetLastRecalculateDateOk returns a tuple with the LastRecalculateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecalculateDate

`func (o *TenantRoomQuotaSettings) SetLastRecalculateDate(v time.Time)`

SetLastRecalculateDate sets LastRecalculateDate field to given value.

### HasLastRecalculateDate

`func (o *TenantRoomQuotaSettings) HasLastRecalculateDate() bool`

HasLastRecalculateDate returns a boolean if a field has been set.

### SetLastRecalculateDateNil

`func (o *TenantRoomQuotaSettings) SetLastRecalculateDateNil(b bool)`

 SetLastRecalculateDateNil sets the value for LastRecalculateDate to be an explicit nil

### UnsetLastRecalculateDate
`func (o *TenantRoomQuotaSettings) UnsetLastRecalculateDate()`

UnsetLastRecalculateDate ensures that no value is present for LastRecalculateDate, not even an explicit nil
### GetLastModified

`func (o *TenantRoomQuotaSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantRoomQuotaSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantRoomQuotaSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantRoomQuotaSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


