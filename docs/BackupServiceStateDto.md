# BackupServiceStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies if the backup service is enabled or not. | [optional] 

## Methods

### NewBackupServiceStateDto

`func NewBackupServiceStateDto() *BackupServiceStateDto`

NewBackupServiceStateDto instantiates a new BackupServiceStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupServiceStateDtoWithDefaults

`func NewBackupServiceStateDtoWithDefaults() *BackupServiceStateDto`

NewBackupServiceStateDtoWithDefaults instantiates a new BackupServiceStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BackupServiceStateDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupServiceStateDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupServiceStateDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BackupServiceStateDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


