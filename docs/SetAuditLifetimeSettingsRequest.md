# SetAuditLifetimeSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginHistoryLifeTime** | Pointer to **int32** | How many days login events are kept, from 1 to 180. | [optional] 
**AuditTrailLifeTime** | Pointer to **int32** | How many days audit trail events are kept, from 1 to 180. | [optional] 
**LastModified** | Pointer to **time.Time** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 

## Methods

### NewSetAuditLifetimeSettingsRequest

`func NewSetAuditLifetimeSettingsRequest() *SetAuditLifetimeSettingsRequest`

NewSetAuditLifetimeSettingsRequest instantiates a new SetAuditLifetimeSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAuditLifetimeSettingsRequestWithDefaults

`func NewSetAuditLifetimeSettingsRequestWithDefaults() *SetAuditLifetimeSettingsRequest`

NewSetAuditLifetimeSettingsRequestWithDefaults instantiates a new SetAuditLifetimeSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginHistoryLifeTime

`func (o *SetAuditLifetimeSettingsRequest) GetLoginHistoryLifeTime() int32`

GetLoginHistoryLifeTime returns the LoginHistoryLifeTime field if non-nil, zero value otherwise.

### GetLoginHistoryLifeTimeOk

`func (o *SetAuditLifetimeSettingsRequest) GetLoginHistoryLifeTimeOk() (*int32, bool)`

GetLoginHistoryLifeTimeOk returns a tuple with the LoginHistoryLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHistoryLifeTime

`func (o *SetAuditLifetimeSettingsRequest) SetLoginHistoryLifeTime(v int32)`

SetLoginHistoryLifeTime sets LoginHistoryLifeTime field to given value.

### HasLoginHistoryLifeTime

`func (o *SetAuditLifetimeSettingsRequest) HasLoginHistoryLifeTime() bool`

HasLoginHistoryLifeTime returns a boolean if a field has been set.

### GetAuditTrailLifeTime

`func (o *SetAuditLifetimeSettingsRequest) GetAuditTrailLifeTime() int32`

GetAuditTrailLifeTime returns the AuditTrailLifeTime field if non-nil, zero value otherwise.

### GetAuditTrailLifeTimeOk

`func (o *SetAuditLifetimeSettingsRequest) GetAuditTrailLifeTimeOk() (*int32, bool)`

GetAuditTrailLifeTimeOk returns a tuple with the AuditTrailLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditTrailLifeTime

`func (o *SetAuditLifetimeSettingsRequest) SetAuditTrailLifeTime(v int32)`

SetAuditTrailLifeTime sets AuditTrailLifeTime field to given value.

### HasAuditTrailLifeTime

`func (o *SetAuditLifetimeSettingsRequest) HasAuditTrailLifeTime() bool`

HasAuditTrailLifeTime returns a boolean if a field has been set.

### GetLastModified

`func (o *SetAuditLifetimeSettingsRequest) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *SetAuditLifetimeSettingsRequest) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *SetAuditLifetimeSettingsRequest) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *SetAuditLifetimeSettingsRequest) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


