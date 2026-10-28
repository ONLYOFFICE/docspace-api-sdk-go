# ConversationResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The identifier of the conversion entry. The portal leaves it empty for file conversions, so a caller follows  its own conversion by the file it queued rather than by this value. | 
**Operation** | [**FileOperationType**](FileOperationType.md) | Tells which kind of file operation the entry describes, so that a conversion can be told apart from the copy,  move and download entries that share this envelope. A conversion entry reports the conversion type. | 
**Progress** | **int32** | How far the conversion has got, counted in percent from 0 while it is only queued to 100 once it is over -  whether it ended with a converted file or with an error. 100 is the value a polling caller waits for. | 
**Source** | Pointer to **NullableString** | Describes what is being converted: the identifier of the source file, the version that was taken and whether  an existing result may be overwritten, packed as a JSON object inside a string. It is what identifies the  entry when several conversions of the same caller are in flight. | [optional] 
**Result** | Pointer to **interface{}** |  | [optional] 
**Error** | Pointer to **NullableString** | The reason the conversion stopped, in the language of the caller, and empty while it is running and after it  has succeeded. `progress` reaches 100 for a failure as well, so this field is what separates a converted file  from a broken conversion; a conversion still unfinished after ten minutes ends with a timeout reported here. | [optional] 
**Processed** | Pointer to **NullableString** | Reports whether the portal has taken the entry as far as it goes: `1` once the conversion has finished or  failed, and empty while it is still queued or still being converted. It is the bookkeeping of the conversion  queue rather than a result - what happened is in `progress`, `error` and `result`. | [optional] 

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


