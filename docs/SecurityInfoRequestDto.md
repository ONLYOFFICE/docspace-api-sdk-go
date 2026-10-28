# SecurityInfoRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderIds** | Pointer to [**[]DuplicateRequestDtoAllOfFileIds**](DuplicateRequestDtoAllOfFileIds.md) | The folders and rooms whose rights are being changed, identified as a listing operation returns them - a  number on the portal, a string on a connected third-party account. | [optional] 
**FileIds** | Pointer to [**[]DuplicateRequestDtoAllOfFileIds**](DuplicateRequestDtoAllOfFileIds.md) | The files whose rights are being changed, identified as a listing operation returns them - a number on the  portal, a string on a connected third-party account. | [optional] 
**Share** | Pointer to [**[]FileShareParams**](FileShareParams.md) | One record per account or group whose rights are being set, each naming the subject and the level it gets on  all of the listed entries; a level of `None` takes the access away. An empty collection makes the call change  nothing. | [optional] 
**Notify** | Pointer to **bool** | Set to true to have every account named in `share` emailed about the access it just received; false changes  the rights without telling anyone. | [optional] 
**SharingMessage** | Pointer to **NullableString** | The text put into that email, ignored while `notify` is false. Markup is stripped before sending, so only the  plain text of the value survives. | [optional] 

## Methods

### NewSecurityInfoRequestDto

`func NewSecurityInfoRequestDto() *SecurityInfoRequestDto`

NewSecurityInfoRequestDto instantiates a new SecurityInfoRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityInfoRequestDtoWithDefaults

`func NewSecurityInfoRequestDtoWithDefaults() *SecurityInfoRequestDto`

NewSecurityInfoRequestDtoWithDefaults instantiates a new SecurityInfoRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderIds

`func (o *SecurityInfoRequestDto) GetFolderIds() []DuplicateRequestDtoAllOfFileIds`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *SecurityInfoRequestDto) GetFolderIdsOk() (*[]DuplicateRequestDtoAllOfFileIds, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *SecurityInfoRequestDto) SetFolderIds(v []DuplicateRequestDtoAllOfFileIds)`

SetFolderIds sets FolderIds field to given value.

### HasFolderIds

`func (o *SecurityInfoRequestDto) HasFolderIds() bool`

HasFolderIds returns a boolean if a field has been set.

### SetFolderIdsNil

`func (o *SecurityInfoRequestDto) SetFolderIdsNil(b bool)`

 SetFolderIdsNil sets the value for FolderIds to be an explicit nil

### UnsetFolderIds
`func (o *SecurityInfoRequestDto) UnsetFolderIds()`

UnsetFolderIds ensures that no value is present for FolderIds, not even an explicit nil
### GetFileIds

`func (o *SecurityInfoRequestDto) GetFileIds() []DuplicateRequestDtoAllOfFileIds`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *SecurityInfoRequestDto) GetFileIdsOk() (*[]DuplicateRequestDtoAllOfFileIds, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *SecurityInfoRequestDto) SetFileIds(v []DuplicateRequestDtoAllOfFileIds)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *SecurityInfoRequestDto) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### SetFileIdsNil

`func (o *SecurityInfoRequestDto) SetFileIdsNil(b bool)`

 SetFileIdsNil sets the value for FileIds to be an explicit nil

### UnsetFileIds
`func (o *SecurityInfoRequestDto) UnsetFileIds()`

UnsetFileIds ensures that no value is present for FileIds, not even an explicit nil
### GetShare

`func (o *SecurityInfoRequestDto) GetShare() []FileShareParams`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *SecurityInfoRequestDto) GetShareOk() (*[]FileShareParams, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *SecurityInfoRequestDto) SetShare(v []FileShareParams)`

SetShare sets Share field to given value.

### HasShare

`func (o *SecurityInfoRequestDto) HasShare() bool`

HasShare returns a boolean if a field has been set.

### SetShareNil

`func (o *SecurityInfoRequestDto) SetShareNil(b bool)`

 SetShareNil sets the value for Share to be an explicit nil

### UnsetShare
`func (o *SecurityInfoRequestDto) UnsetShare()`

UnsetShare ensures that no value is present for Share, not even an explicit nil
### GetNotify

`func (o *SecurityInfoRequestDto) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *SecurityInfoRequestDto) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *SecurityInfoRequestDto) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *SecurityInfoRequestDto) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetSharingMessage

`func (o *SecurityInfoRequestDto) GetSharingMessage() string`

GetSharingMessage returns the SharingMessage field if non-nil, zero value otherwise.

### GetSharingMessageOk

`func (o *SecurityInfoRequestDto) GetSharingMessageOk() (*string, bool)`

GetSharingMessageOk returns a tuple with the SharingMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingMessage

`func (o *SecurityInfoRequestDto) SetSharingMessage(v string)`

SetSharingMessage sets SharingMessage field to given value.

### HasSharingMessage

`func (o *SecurityInfoRequestDto) HasSharingMessage() bool`

HasSharingMessage returns a boolean if a field has been set.

### SetSharingMessageNil

`func (o *SecurityInfoRequestDto) SetSharingMessageNil(b bool)`

 SetSharingMessageNil sets the value for SharingMessage to be an explicit nil

### UnsetSharingMessage
`func (o *SecurityInfoRequestDto) UnsetSharingMessage()`

UnsetSharingMessage ensures that no value is present for SharingMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


