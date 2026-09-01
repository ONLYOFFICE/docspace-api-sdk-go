# EditHistoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The document ID. | [optional] 
**Key** | Pointer to **NullableString** | The document identifier used to unambiguously identify the document file. | [optional] 
**Version** | Pointer to **int32** | The document version number. | [optional] 
**VersionGroup** | Pointer to **int32** | The document version group. | [optional] 
**User** | Pointer to [**EditHistoryAuthor**](EditHistoryAuthor.md) | The user who updated a file. | [optional] 
**Created** | Pointer to **NullableTime** | The document version creation date. | [optional] 
**ChangesHistory** | Pointer to **NullableString** | The file history changes in the string format. | [optional] 
**Changes** | Pointer to [**[]EditHistoryChangesWrapper**](EditHistoryChangesWrapper.md) | The list of file history changes. | [optional] 
**ServerVersion** | Pointer to **NullableString** | The current server version number. | [optional] 

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

`func (o *EditHistoryDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EditHistoryDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EditHistoryDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EditHistoryDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *EditHistoryDto) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *EditHistoryDto) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
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


