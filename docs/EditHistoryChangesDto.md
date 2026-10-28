# EditHistoryChangesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**EditHistoryAuthorDto**](EditHistoryAuthorDto.md) | The account that made this change, as the editing service reported it; an account it could not name is  reported as a guest. | [optional] 
**Created** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When this change was made, written with the offset of the portal's time zone rather than as plain UTC. | [optional] 
**DocumentSha256** | Pointer to **NullableString** | The SHA-256 hash of the document as it stood after this change, where the editing service recorded one, so  that a client can check a stored copy against the change it claims to hold. Empty when the change record  carries no hash. | [optional] 

## Methods

### NewEditHistoryChangesDto

`func NewEditHistoryChangesDto() *EditHistoryChangesDto`

NewEditHistoryChangesDto instantiates a new EditHistoryChangesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHistoryChangesDtoWithDefaults

`func NewEditHistoryChangesDtoWithDefaults() *EditHistoryChangesDto`

NewEditHistoryChangesDtoWithDefaults instantiates a new EditHistoryChangesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *EditHistoryChangesDto) GetUser() EditHistoryAuthorDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EditHistoryChangesDto) GetUserOk() (*EditHistoryAuthorDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EditHistoryChangesDto) SetUser(v EditHistoryAuthorDto)`

SetUser sets User field to given value.

### HasUser

`func (o *EditHistoryChangesDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCreated

`func (o *EditHistoryChangesDto) GetCreated() ApiDateTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EditHistoryChangesDto) GetCreatedOk() (*ApiDateTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EditHistoryChangesDto) SetCreated(v ApiDateTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EditHistoryChangesDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDocumentSha256

`func (o *EditHistoryChangesDto) GetDocumentSha256() string`

GetDocumentSha256 returns the DocumentSha256 field if non-nil, zero value otherwise.

### GetDocumentSha256Ok

`func (o *EditHistoryChangesDto) GetDocumentSha256Ok() (*string, bool)`

GetDocumentSha256Ok returns a tuple with the DocumentSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSha256

`func (o *EditHistoryChangesDto) SetDocumentSha256(v string)`

SetDocumentSha256 sets DocumentSha256 field to given value.

### HasDocumentSha256

`func (o *EditHistoryChangesDto) HasDocumentSha256() bool`

HasDocumentSha256 returns a boolean if a field has been set.

### SetDocumentSha256Nil

`func (o *EditHistoryChangesDto) SetDocumentSha256Nil(b bool)`

 SetDocumentSha256Nil sets the value for DocumentSha256 to be an explicit nil

### UnsetDocumentSha256
`func (o *EditHistoryChangesDto) UnsetDocumentSha256()`

UnsetDocumentSha256 ensures that no value is present for DocumentSha256, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


