# MigrationStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | Pointer to **float64** | The share of the job that is done, from 0 to 100. It advances unevenly, since the stages differ in  length, so poll `isCompleted` rather than waiting for this to reach 100. | [optional] 
**Error** | Pointer to **NullableString** | The message that ended the job, in the portal language. It stays empty while nothing has gone wrong, so  once `isCompleted` is `true` this field is what tells success from failure. | [optional] 
**ParseResult** | Pointer to [**MigrationApiInfo**](MigrationApiInfo.md) | What the migrator has read so far. After a parse pass it holds the users, the groups and the archives it  could not read, which is the body to edit and post to `POST api/2.0/migration/migrate`; during an import it  also carries the accounts that were created and the ones that failed. Its own `operation` field, `parse`  or `migration`, is what tells the two stages apart. | [optional] 
**IsCompleted** | Pointer to **bool** | Whether the job has stopped, successfully or not. It is the field to poll on; the whole body comes back  empty instead when the portal has no job at all, which is not an error. | [optional] 

## Methods

### NewMigrationStatusDto

`func NewMigrationStatusDto() *MigrationStatusDto`

NewMigrationStatusDto instantiates a new MigrationStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationStatusDtoWithDefaults

`func NewMigrationStatusDtoWithDefaults() *MigrationStatusDto`

NewMigrationStatusDtoWithDefaults instantiates a new MigrationStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgress

`func (o *MigrationStatusDto) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *MigrationStatusDto) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *MigrationStatusDto) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *MigrationStatusDto) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetError

`func (o *MigrationStatusDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MigrationStatusDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MigrationStatusDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MigrationStatusDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MigrationStatusDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MigrationStatusDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetParseResult

`func (o *MigrationStatusDto) GetParseResult() MigrationApiInfo`

GetParseResult returns the ParseResult field if non-nil, zero value otherwise.

### GetParseResultOk

`func (o *MigrationStatusDto) GetParseResultOk() (*MigrationApiInfo, bool)`

GetParseResultOk returns a tuple with the ParseResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseResult

`func (o *MigrationStatusDto) SetParseResult(v MigrationApiInfo)`

SetParseResult sets ParseResult field to given value.

### HasParseResult

`func (o *MigrationStatusDto) HasParseResult() bool`

HasParseResult returns a boolean if a field has been set.

### GetIsCompleted

`func (o *MigrationStatusDto) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *MigrationStatusDto) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *MigrationStatusDto) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *MigrationStatusDto) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


