# TenantAuditSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginHistoryLifeTime** | Pointer to **int32** | The login history lifetime. | [optional] 
**AuditTrailLifeTime** | Pointer to **int32** | The audit trail lifetime. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewTenantAuditSettings

`func NewTenantAuditSettings() *TenantAuditSettings`

NewTenantAuditSettings instantiates a new TenantAuditSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAuditSettingsWithDefaults

`func NewTenantAuditSettingsWithDefaults() *TenantAuditSettings`

NewTenantAuditSettingsWithDefaults instantiates a new TenantAuditSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginHistoryLifeTime

`func (o *TenantAuditSettings) GetLoginHistoryLifeTime() int32`

GetLoginHistoryLifeTime returns the LoginHistoryLifeTime field if non-nil, zero value otherwise.

### GetLoginHistoryLifeTimeOk

`func (o *TenantAuditSettings) GetLoginHistoryLifeTimeOk() (*int32, bool)`

GetLoginHistoryLifeTimeOk returns a tuple with the LoginHistoryLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHistoryLifeTime

`func (o *TenantAuditSettings) SetLoginHistoryLifeTime(v int32)`

SetLoginHistoryLifeTime sets LoginHistoryLifeTime field to given value.

### HasLoginHistoryLifeTime

`func (o *TenantAuditSettings) HasLoginHistoryLifeTime() bool`

HasLoginHistoryLifeTime returns a boolean if a field has been set.

### GetAuditTrailLifeTime

`func (o *TenantAuditSettings) GetAuditTrailLifeTime() int32`

GetAuditTrailLifeTime returns the AuditTrailLifeTime field if non-nil, zero value otherwise.

### GetAuditTrailLifeTimeOk

`func (o *TenantAuditSettings) GetAuditTrailLifeTimeOk() (*int32, bool)`

GetAuditTrailLifeTimeOk returns a tuple with the AuditTrailLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditTrailLifeTime

`func (o *TenantAuditSettings) SetAuditTrailLifeTime(v int32)`

SetAuditTrailLifeTime sets AuditTrailLifeTime field to given value.

### HasAuditTrailLifeTime

`func (o *TenantAuditSettings) HasAuditTrailLifeTime() bool`

HasAuditTrailLifeTime returns a boolean if a field has been set.

### GetLastModified

`func (o *TenantAuditSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantAuditSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantAuditSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantAuditSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


