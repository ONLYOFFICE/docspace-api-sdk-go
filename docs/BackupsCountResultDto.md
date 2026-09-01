# BackupsCountResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Free** | Pointer to **int32** | The number of free backups. | [optional] 
**Paid** | Pointer to **int32** | The number of paid backups. | [optional] 

## Methods

### NewBackupsCountResultDto

`func NewBackupsCountResultDto() *BackupsCountResultDto`

NewBackupsCountResultDto instantiates a new BackupsCountResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupsCountResultDtoWithDefaults

`func NewBackupsCountResultDtoWithDefaults() *BackupsCountResultDto`

NewBackupsCountResultDtoWithDefaults instantiates a new BackupsCountResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFree

`func (o *BackupsCountResultDto) GetFree() int32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *BackupsCountResultDto) GetFreeOk() (*int32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *BackupsCountResultDto) SetFree(v int32)`

SetFree sets Free field to given value.

### HasFree

`func (o *BackupsCountResultDto) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetPaid

`func (o *BackupsCountResultDto) GetPaid() int32`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *BackupsCountResultDto) GetPaidOk() (*int32, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *BackupsCountResultDto) SetPaid(v int32)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *BackupsCountResultDto) HasPaid() bool`

HasPaid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


