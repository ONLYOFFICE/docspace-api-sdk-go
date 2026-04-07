# ConversationResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The conversion operation ID. | 
**Operation** | [**FileOperationType**](FileOperationType.md) |  | 
**Progress** | **int32** | The conversion operation progress. | 
**Source** | Pointer to **NullableString** | The source file for the conversion. | [optional] 
**Result** | Pointer to **interface{}** | The resulting file after the conversion. | [optional] 
**Error** | Pointer to **NullableString** | The conversion operation error message. | [optional] 
**Processed** | Pointer to **NullableString** | Specifies if the conversion operation is processed or not. | [optional] 

## Methods

### NewConversationResultDto

`func NewConversationResultDto(id NullableString, operation FileOperationType, progress int32, ) *ConversationResultDto`

NewConversationResultDto instantiates a new ConversationResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationResultDtoWithDefaults

`func NewConversationResultDtoWithDefaults() *ConversationResultDto`

NewConversationResultDtoWithDefaults instantiates a new ConversationResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConversationResultDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversationResultDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversationResultDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ConversationResultDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConversationResultDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOperation

`func (o *ConversationResultDto) GetOperation() FileOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ConversationResultDto) GetOperationOk() (*FileOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ConversationResultDto) SetOperation(v FileOperationType)`

SetOperation sets Operation field to given value.


### GetProgress

`func (o *ConversationResultDto) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ConversationResultDto) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ConversationResultDto) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetSource

`func (o *ConversationResultDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConversationResultDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConversationResultDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConversationResultDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ConversationResultDto) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ConversationResultDto) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetResult

`func (o *ConversationResultDto) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ConversationResultDto) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ConversationResultDto) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *ConversationResultDto) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ConversationResultDto) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ConversationResultDto) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetError

`func (o *ConversationResultDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConversationResultDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConversationResultDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConversationResultDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ConversationResultDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ConversationResultDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetProcessed

`func (o *ConversationResultDto) GetProcessed() string`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *ConversationResultDto) GetProcessedOk() (*string, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *ConversationResultDto) SetProcessed(v string)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *ConversationResultDto) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### SetProcessedNil

`func (o *ConversationResultDto) SetProcessedNil(b bool)`

 SetProcessedNil sets the value for Processed to be an explicit nil

### UnsetProcessed
`func (o *ConversationResultDto) UnsetProcessed()`

UnsetProcessed ensures that no value is present for Processed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


