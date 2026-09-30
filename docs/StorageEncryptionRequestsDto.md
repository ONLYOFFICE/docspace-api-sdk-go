# StorageEncryptionRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyUsers** | Pointer to **bool** | Whether every user of every portal on the server is mailed before the encryption or decryption pass starts.  The pass runs either way; the flag only decides whether people are told that their portal is about to become  unavailable. | [optional] 

## Methods

### NewStorageEncryptionRequestsDto

`func NewStorageEncryptionRequestsDto() *StorageEncryptionRequestsDto`

NewStorageEncryptionRequestsDto instantiates a new StorageEncryptionRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEncryptionRequestsDtoWithDefaults

`func NewStorageEncryptionRequestsDtoWithDefaults() *StorageEncryptionRequestsDto`

NewStorageEncryptionRequestsDtoWithDefaults instantiates a new StorageEncryptionRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyUsers

`func (o *StorageEncryptionRequestsDto) GetNotifyUsers() bool`

GetNotifyUsers returns the NotifyUsers field if non-nil, zero value otherwise.

### GetNotifyUsersOk

`func (o *StorageEncryptionRequestsDto) GetNotifyUsersOk() (*bool, bool)`

GetNotifyUsersOk returns a tuple with the NotifyUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUsers

`func (o *StorageEncryptionRequestsDto) SetNotifyUsers(v bool)`

SetNotifyUsers sets NotifyUsers field to given value.

### HasNotifyUsers

`func (o *StorageEncryptionRequestsDto) HasNotifyUsers() bool`

HasNotifyUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


