# LockFileParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockFile** | Pointer to **bool** | The state to reach: `true` locks the file, which blocks editing, renaming and deleting for everybody but the  account that locked it and the room admins, and drops the others out of a running editing session; `false`  releases the lock. | [optional] 

## Methods

### NewLockFileParameters

`func NewLockFileParameters() *LockFileParameters`

NewLockFileParameters instantiates a new LockFileParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockFileParametersWithDefaults

`func NewLockFileParametersWithDefaults() *LockFileParameters`

NewLockFileParametersWithDefaults instantiates a new LockFileParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockFile

`func (o *LockFileParameters) GetLockFile() bool`

GetLockFile returns the LockFile field if non-nil, zero value otherwise.

### GetLockFileOk

`func (o *LockFileParameters) GetLockFileOk() (*bool, bool)`

GetLockFileOk returns a tuple with the LockFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockFile

`func (o *LockFileParameters) SetLockFile(v bool)`

SetLockFile sets LockFile field to given value.

### HasLockFile

`func (o *LockFileParameters) HasLockFile() bool`

HasLockFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


