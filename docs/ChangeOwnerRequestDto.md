# ChangeOwnerRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderIds** | Pointer to [**[]BatchRequestDtoAllOfFileIds**](BatchRequestDtoAllOfFileIds.md) | The rooms to hand over, identified as `GET api/2.0/files/rooms` returns them - a number for a room stored on  the portal and a string for one that lives on a connected third-party account. Only rooms belong here; a  folder inside a room is refused. | [optional] 
**FileIds** | Pointer to [**[]BatchRequestDtoAllOfFileIds**](BatchRequestDtoAllOfFileIds.md) | The files to hand over, identified as a listing operation returns them - a number for a file stored on the  portal and a string for one on a connected third-party account. Only a file kept in the portal's common  section is accepted. | [optional] 
**UserId** | **string** | The account that becomes the owner of every listed entry. It has to be an active member allowed to manage  rooms, so a deactivated account, a guest or a plain member is rejected, and for a private room the account  must have set up its encryption keys beforehand. | 

## Methods

### NewChangeOwnerRequestDto

`func NewChangeOwnerRequestDto(userId string, ) *ChangeOwnerRequestDto`

NewChangeOwnerRequestDto instantiates a new ChangeOwnerRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeOwnerRequestDtoWithDefaults

`func NewChangeOwnerRequestDtoWithDefaults() *ChangeOwnerRequestDto`

NewChangeOwnerRequestDtoWithDefaults instantiates a new ChangeOwnerRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderIds

`func (o *ChangeOwnerRequestDto) GetFolderIds() []BatchRequestDtoAllOfFileIds`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *ChangeOwnerRequestDto) GetFolderIdsOk() (*[]BatchRequestDtoAllOfFileIds, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *ChangeOwnerRequestDto) SetFolderIds(v []BatchRequestDtoAllOfFileIds)`

SetFolderIds sets FolderIds field to given value.

### HasFolderIds

`func (o *ChangeOwnerRequestDto) HasFolderIds() bool`

HasFolderIds returns a boolean if a field has been set.

### SetFolderIdsNil

`func (o *ChangeOwnerRequestDto) SetFolderIdsNil(b bool)`

 SetFolderIdsNil sets the value for FolderIds to be an explicit nil

### UnsetFolderIds
`func (o *ChangeOwnerRequestDto) UnsetFolderIds()`

UnsetFolderIds ensures that no value is present for FolderIds, not even an explicit nil
### GetFileIds

`func (o *ChangeOwnerRequestDto) GetFileIds() []BatchRequestDtoAllOfFileIds`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *ChangeOwnerRequestDto) GetFileIdsOk() (*[]BatchRequestDtoAllOfFileIds, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *ChangeOwnerRequestDto) SetFileIds(v []BatchRequestDtoAllOfFileIds)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *ChangeOwnerRequestDto) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### SetFileIdsNil

`func (o *ChangeOwnerRequestDto) SetFileIdsNil(b bool)`

 SetFileIdsNil sets the value for FileIds to be an explicit nil

### UnsetFileIds
`func (o *ChangeOwnerRequestDto) UnsetFileIds()`

UnsetFileIds ensures that no value is present for FileIds, not even an explicit nil
### GetUserId

`func (o *ChangeOwnerRequestDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChangeOwnerRequestDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChangeOwnerRequestDto) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


