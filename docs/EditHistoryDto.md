# EditHistoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The file the revision belongs to; every entry of one history carries the same value. | [optional] 
**Key** | Pointer to **NullableString** | The document key of this revision, which the editing service uses to tell the revisions of a file apart and to  reuse the copy it has cached. Hand it back unchanged when asking the editor for this revision. | [optional] 
**Version** | Pointer to **int32** | The number of the revision, counting up from 1 in the order the revisions were saved. It is the value the  operations that show the changes of a revision or restore it expect. | [optional] 
**VersionGroup** | Pointer to **int32** | Groups the revisions written by one editing session: entries sharing this number were saved while the same  session was open, which is how a client collapses a long list of revisions into the versions a person would  recognise. | [optional] 
**User** | Pointer to [**EditHistoryAuthor**](EditHistoryAuthor.md) | The account that saved the revision. A revision saved by an account that no longer exists, or through an  anonymous link, is reported as a guest. | [optional] 
**Created** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the revision was saved, written with the offset of the portal's time zone rather than as plain UTC. The  times of one history are consistent with each other, so order and display the revisions by them. | [optional] 
**ChangesHistory** | Pointer to **NullableString** | The change record the editing service stored for this revision, as the raw JSON it was written in, and empty  for a revision the portal has no record for - one uploaded as a whole file, for instance. `changes` is the  same record already parsed. | [optional] 
**Changes** | Pointer to [**[]EditHistoryChangesWrapper**](EditHistoryChangesWrapper.md) | The single changes this revision introduced - who made each of them and when - taken from the stored change  record. It comes back empty both for a revision whose changes were never recorded and for one whose record is  in a format the portal no longer reads, so an empty list is not proof that nothing changed. | [optional] 
**ServerVersion** | Pointer to **NullableString** | The build of the editing service that wrote the change record of this revision, taken from the record itself;  empty when the portal holds no record for the revision. | [optional] 

## Methods

### NewEditHistoryDto

`func NewEditHistoryDto() *EditHistoryDto`

NewEditHistoryDto instantiates a new EditHistoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHistoryDtoWithDefaults

`func NewEditHistoryDtoWithDefaults() *EditHistoryDto`

NewEditHistoryDtoWithDefaults instantiates a new EditHistoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditHistoryDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditHistoryDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditHistoryDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EditHistoryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *EditHistoryDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EditHistoryDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EditHistoryDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EditHistoryDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *EditHistoryDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *EditHistoryDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetVersion

`func (o *EditHistoryDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EditHistoryDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EditHistoryDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EditHistoryDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionGroup

`func (o *EditHistoryDto) GetVersionGroup() int32`

GetVersionGroup returns the VersionGroup field if non-nil, zero value otherwise.

### GetVersionGroupOk

`func (o *EditHistoryDto) GetVersionGroupOk() (*int32, bool)`

GetVersionGroupOk returns a tuple with the VersionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroup

`func (o *EditHistoryDto) SetVersionGroup(v int32)`

SetVersionGroup sets VersionGroup field to given value.

### HasVersionGroup

`func (o *EditHistoryDto) HasVersionGroup() bool`

HasVersionGroup returns a boolean if a field has been set.

### GetUser

`func (o *EditHistoryDto) GetUser() EditHistoryAuthor`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EditHistoryDto) GetUserOk() (*EditHistoryAuthor, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EditHistoryDto) SetUser(v EditHistoryAuthor)`

SetUser sets User field to given value.

### HasUser

`func (o *EditHistoryDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCreated

`func (o *EditHistoryDto) GetCreated() ApiDateTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EditHistoryDto) GetCreatedOk() (*ApiDateTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EditHistoryDto) SetCreated(v ApiDateTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EditHistoryDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetChangesHistory

`func (o *EditHistoryDto) GetChangesHistory() string`

GetChangesHistory returns the ChangesHistory field if non-nil, zero value otherwise.

### GetChangesHistoryOk

`func (o *EditHistoryDto) GetChangesHistoryOk() (*string, bool)`

GetChangesHistoryOk returns a tuple with the ChangesHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesHistory

`func (o *EditHistoryDto) SetChangesHistory(v string)`

SetChangesHistory sets ChangesHistory field to given value.

### HasChangesHistory

`func (o *EditHistoryDto) HasChangesHistory() bool`

HasChangesHistory returns a boolean if a field has been set.

### SetChangesHistoryNil

`func (o *EditHistoryDto) SetChangesHistoryNil(b bool)`

 SetChangesHistoryNil sets the value for ChangesHistory to be an explicit nil

### UnsetChangesHistory
`func (o *EditHistoryDto) UnsetChangesHistory()`

UnsetChangesHistory ensures that no value is present for ChangesHistory, not even an explicit nil
### GetChanges

`func (o *EditHistoryDto) GetChanges() []EditHistoryChangesWrapper`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *EditHistoryDto) GetChangesOk() (*[]EditHistoryChangesWrapper, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *EditHistoryDto) SetChanges(v []EditHistoryChangesWrapper)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *EditHistoryDto) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### SetChangesNil

`func (o *EditHistoryDto) SetChangesNil(b bool)`

 SetChangesNil sets the value for Changes to be an explicit nil

### UnsetChanges
`func (o *EditHistoryDto) UnsetChanges()`

UnsetChanges ensures that no value is present for Changes, not even an explicit nil
### GetServerVersion

`func (o *EditHistoryDto) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *EditHistoryDto) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *EditHistoryDto) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *EditHistoryDto) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### SetServerVersionNil

`func (o *EditHistoryDto) SetServerVersionNil(b bool)`

 SetServerVersionNil sets the value for ServerVersion to be an explicit nil

### UnsetServerVersion
`func (o *EditHistoryDto) UnsetServerVersion()`

UnsetServerVersion ensures that no value is present for ServerVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


