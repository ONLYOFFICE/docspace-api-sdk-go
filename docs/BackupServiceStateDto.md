# BackupServiceStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether the paid backup service is switched on for this portal, which is a setting of its  wallet rather than the health of the backup service. While it is true, backups beyond the free  monthly allowance are charged to the wallet. | [optional] 

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


