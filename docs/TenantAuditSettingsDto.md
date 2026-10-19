# TenantAuditSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginHistoryLifeTime** | Pointer to **int32** | How many days login events are kept, from 1 to 180; 180 when the portal never changed it. | [optional] 
**AuditTrailLifeTime** | Pointer to **int32** | How many days audit trail events are kept, from 1 to 180; 180 when the portal never changed it. | [optional] 
**LastModified** | Pointer to **time.Time** | When the pair was last stored; when it never was, the moment it was read instead. | [optional] 

## Methods

### NewTenantAuditSettingsDto

`func NewTenantAuditSettingsDto() *TenantAuditSettingsDto`

NewTenantAuditSettingsDto instantiates a new TenantAuditSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAuditSettingsDtoWithDefaults

`func NewTenantAuditSettingsDtoWithDefaults() *TenantAuditSettingsDto`

NewTenantAuditSettingsDtoWithDefaults instantiates a new TenantAuditSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginHistoryLifeTime

`func (o *TenantAuditSettingsDto) GetLoginHistoryLifeTime() int32`

GetLoginHistoryLifeTime returns the LoginHistoryLifeTime field if non-nil, zero value otherwise.

### GetLoginHistoryLifeTimeOk

`func (o *TenantAuditSettingsDto) GetLoginHistoryLifeTimeOk() (*int32, bool)`

GetLoginHistoryLifeTimeOk returns a tuple with the LoginHistoryLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHistoryLifeTime

`func (o *TenantAuditSettingsDto) SetLoginHistoryLifeTime(v int32)`

SetLoginHistoryLifeTime sets LoginHistoryLifeTime field to given value.

### HasLoginHistoryLifeTime

`func (o *TenantAuditSettingsDto) HasLoginHistoryLifeTime() bool`

HasLoginHistoryLifeTime returns a boolean if a field has been set.

### GetAuditTrailLifeTime

`func (o *TenantAuditSettingsDto) GetAuditTrailLifeTime() int32`

GetAuditTrailLifeTime returns the AuditTrailLifeTime field if non-nil, zero value otherwise.

### GetAuditTrailLifeTimeOk

`func (o *TenantAuditSettingsDto) GetAuditTrailLifeTimeOk() (*int32, bool)`

GetAuditTrailLifeTimeOk returns a tuple with the AuditTrailLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditTrailLifeTime

`func (o *TenantAuditSettingsDto) SetAuditTrailLifeTime(v int32)`

SetAuditTrailLifeTime sets AuditTrailLifeTime field to given value.

### HasAuditTrailLifeTime

`func (o *TenantAuditSettingsDto) HasAuditTrailLifeTime() bool`

HasAuditTrailLifeTime returns a boolean if a field has been set.

### GetLastModified

`func (o *TenantAuditSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantAuditSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantAuditSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantAuditSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


