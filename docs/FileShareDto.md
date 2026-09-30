# FileShareDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**FileShare**](FileShare.md) | The level the subject holds on the entry. On a link entry it is the level the link hands to whoever opens it,  and in a batch answer `Varies` means the subject holds different levels on the listed entries. | [optional] 
**SharedTo** | Pointer to **interface{}** |  | [optional] 
**SharedToUser** | Pointer to [**EmployeeFullDto**](EmployeeFullDto.md) | The account the entry belongs to. It is filled in only when `subjectType` says an account, and is null for a  group entry and for a link. | [optional] 
**SharedToGroup** | Pointer to [**GroupSummaryDto**](GroupSummaryDto.md) | The portal group the entry belongs to, which hands the level to everybody in it. It is filled in only for a  group entry, and is null otherwise. | [optional] 
**SharedLink** | Pointer to [**FileShareLink**](FileShareLink.md) | The sharing link the entry stands for, together with everything set on it. It is filled in only for a link  entry, and is null for an account or a group. | [optional] 
**IsLocked** | **bool** | Whether this entry is the caller's own, which is why they cannot change its level. Link entries never report  it. | 
**IsOwner** | **bool** | Whether the subject created the entry the access is given on, and so cannot be removed from it. | 
**CanEditAccess** | **bool** | Whether the caller may change the level of this entry. It is false on the caller's own entry, on every link,  and whenever the caller may not hand out access at all. | 
**CanEditInternal** | **bool** | Whether the caller may switch this link between being open to anybody and asking the visitor to sign in to the  portal first. | 
**CanEditDenyDownload** | **bool** | Whether the caller may forbid downloading through this link. Only a link of a virtual data room reports true,  and only while the room itself still allows downloads. | 
**CanEditExpirationDate** | **bool** | Whether the caller may move the moment this link stops working. | 
**CanRevoke** | **bool** | Whether the caller may take this entry away altogether, which for a link means deleting the link. | 
**SubjectType** | [**SubjectType**](SubjectType.md) | What the entry was given to, which tells which of the three subject fields is filled in: an account, a group,  or one of the kinds of link. | 

## Methods

### NewFileShareDto

`func NewFileShareDto(isLocked bool, isOwner bool, canEditAccess bool, canEditInternal bool, canEditDenyDownload bool, canEditExpirationDate bool, canRevoke bool, subjectType SubjectType, ) *FileShareDto`

NewFileShareDto instantiates a new FileShareDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileShareDtoWithDefaults

`func NewFileShareDtoWithDefaults() *FileShareDto`

NewFileShareDtoWithDefaults instantiates a new FileShareDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *FileShareDto) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FileShareDto) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FileShareDto) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FileShareDto) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedTo

`func (o *FileShareDto) GetSharedTo() interface{}`

GetSharedTo returns the SharedTo field if non-nil, zero value otherwise.

### GetSharedToOk

`func (o *FileShareDto) GetSharedToOk() (*interface{}, bool)`

GetSharedToOk returns a tuple with the SharedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedTo

`func (o *FileShareDto) SetSharedTo(v interface{})`

SetSharedTo sets SharedTo field to given value.

### HasSharedTo

`func (o *FileShareDto) HasSharedTo() bool`

HasSharedTo returns a boolean if a field has been set.

### SetSharedToNil

`func (o *FileShareDto) SetSharedToNil(b bool)`

 SetSharedToNil sets the value for SharedTo to be an explicit nil

### UnsetSharedTo
`func (o *FileShareDto) UnsetSharedTo()`

UnsetSharedTo ensures that no value is present for SharedTo, not even an explicit nil
### GetSharedToUser

`func (o *FileShareDto) GetSharedToUser() EmployeeFullDto`

GetSharedToUser returns the SharedToUser field if non-nil, zero value otherwise.

### GetSharedToUserOk

`func (o *FileShareDto) GetSharedToUserOk() (*EmployeeFullDto, bool)`

GetSharedToUserOk returns a tuple with the SharedToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedToUser

`func (o *FileShareDto) SetSharedToUser(v EmployeeFullDto)`

SetSharedToUser sets SharedToUser field to given value.

### HasSharedToUser

`func (o *FileShareDto) HasSharedToUser() bool`

HasSharedToUser returns a boolean if a field has been set.

### GetSharedToGroup

`func (o *FileShareDto) GetSharedToGroup() GroupSummaryDto`

GetSharedToGroup returns the SharedToGroup field if non-nil, zero value otherwise.

### GetSharedToGroupOk

`func (o *FileShareDto) GetSharedToGroupOk() (*GroupSummaryDto, bool)`

GetSharedToGroupOk returns a tuple with the SharedToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedToGroup

`func (o *FileShareDto) SetSharedToGroup(v GroupSummaryDto)`

SetSharedToGroup sets SharedToGroup field to given value.

### HasSharedToGroup

`func (o *FileShareDto) HasSharedToGroup() bool`

HasSharedToGroup returns a boolean if a field has been set.

### GetSharedLink

`func (o *FileShareDto) GetSharedLink() FileShareLink`

GetSharedLink returns the SharedLink field if non-nil, zero value otherwise.

### GetSharedLinkOk

`func (o *FileShareDto) GetSharedLinkOk() (*FileShareLink, bool)`

GetSharedLinkOk returns a tuple with the SharedLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedLink

`func (o *FileShareDto) SetSharedLink(v FileShareLink)`

SetSharedLink sets SharedLink field to given value.

### HasSharedLink

`func (o *FileShareDto) HasSharedLink() bool`

HasSharedLink returns a boolean if a field has been set.

### GetIsLocked

`func (o *FileShareDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *FileShareDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *FileShareDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetIsOwner

`func (o *FileShareDto) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *FileShareDto) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *FileShareDto) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.


### GetCanEditAccess

`func (o *FileShareDto) GetCanEditAccess() bool`

GetCanEditAccess returns the CanEditAccess field if non-nil, zero value otherwise.

### GetCanEditAccessOk

`func (o *FileShareDto) GetCanEditAccessOk() (*bool, bool)`

GetCanEditAccessOk returns a tuple with the CanEditAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditAccess

`func (o *FileShareDto) SetCanEditAccess(v bool)`

SetCanEditAccess sets CanEditAccess field to given value.


### GetCanEditInternal

`func (o *FileShareDto) GetCanEditInternal() bool`

GetCanEditInternal returns the CanEditInternal field if non-nil, zero value otherwise.

### GetCanEditInternalOk

`func (o *FileShareDto) GetCanEditInternalOk() (*bool, bool)`

GetCanEditInternalOk returns a tuple with the CanEditInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditInternal

`func (o *FileShareDto) SetCanEditInternal(v bool)`

SetCanEditInternal sets CanEditInternal field to given value.


### GetCanEditDenyDownload

`func (o *FileShareDto) GetCanEditDenyDownload() bool`

GetCanEditDenyDownload returns the CanEditDenyDownload field if non-nil, zero value otherwise.

### GetCanEditDenyDownloadOk

`func (o *FileShareDto) GetCanEditDenyDownloadOk() (*bool, bool)`

GetCanEditDenyDownloadOk returns a tuple with the CanEditDenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditDenyDownload

`func (o *FileShareDto) SetCanEditDenyDownload(v bool)`

SetCanEditDenyDownload sets CanEditDenyDownload field to given value.


### GetCanEditExpirationDate

`func (o *FileShareDto) GetCanEditExpirationDate() bool`

GetCanEditExpirationDate returns the CanEditExpirationDate field if non-nil, zero value otherwise.

### GetCanEditExpirationDateOk

`func (o *FileShareDto) GetCanEditExpirationDateOk() (*bool, bool)`

GetCanEditExpirationDateOk returns a tuple with the CanEditExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditExpirationDate

`func (o *FileShareDto) SetCanEditExpirationDate(v bool)`

SetCanEditExpirationDate sets CanEditExpirationDate field to given value.


### GetCanRevoke

`func (o *FileShareDto) GetCanRevoke() bool`

GetCanRevoke returns the CanRevoke field if non-nil, zero value otherwise.

### GetCanRevokeOk

`func (o *FileShareDto) GetCanRevokeOk() (*bool, bool)`

GetCanRevokeOk returns a tuple with the CanRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRevoke

`func (o *FileShareDto) SetCanRevoke(v bool)`

SetCanRevoke sets CanRevoke field to given value.


### GetSubjectType

`func (o *FileShareDto) GetSubjectType() SubjectType`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *FileShareDto) GetSubjectTypeOk() (*SubjectType, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *FileShareDto) SetSubjectType(v SubjectType)`

SetSubjectType sets SubjectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


