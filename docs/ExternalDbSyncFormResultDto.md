# ExternalDbSyncFormResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The file of the original form whose collected data was exported. It is the form itself, not one of the filled  copies, so the same id can be read with the file operations of the portal. | [optional] 
**Title** | Pointer to **NullableString** | The name of that form file at the moment of the export. It is empty when the form file no longer exists, which  is also the case in which the export of that entry fails. | [optional] 
**Success** | Pointer to **bool** | Whether the data of this form reached the external database. One rejected form does not stop the others, so a  finished job can hold both successful and failed entries. | [optional] 
**Error** | Pointer to **NullableString** | Why this form was not exported. It is empty for a successful entry, and for a failed one it carries either the  message of the underlying failure or the generic export error of the portal. | [optional] 

## Methods

### NewExternalDbSyncFormResultDto

`func NewExternalDbSyncFormResultDto() *ExternalDbSyncFormResultDto`

NewExternalDbSyncFormResultDto instantiates a new ExternalDbSyncFormResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalDbSyncFormResultDtoWithDefaults

`func NewExternalDbSyncFormResultDtoWithDefaults() *ExternalDbSyncFormResultDto`

NewExternalDbSyncFormResultDtoWithDefaults instantiates a new ExternalDbSyncFormResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalDbSyncFormResultDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalDbSyncFormResultDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalDbSyncFormResultDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalDbSyncFormResultDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ExternalDbSyncFormResultDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExternalDbSyncFormResultDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExternalDbSyncFormResultDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExternalDbSyncFormResultDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ExternalDbSyncFormResultDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ExternalDbSyncFormResultDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSuccess

`func (o *ExternalDbSyncFormResultDto) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ExternalDbSyncFormResultDto) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ExternalDbSyncFormResultDto) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ExternalDbSyncFormResultDto) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *ExternalDbSyncFormResultDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExternalDbSyncFormResultDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExternalDbSyncFormResultDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ExternalDbSyncFormResultDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ExternalDbSyncFormResultDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ExternalDbSyncFormResultDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


