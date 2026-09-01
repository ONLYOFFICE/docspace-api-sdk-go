# AiAiRegenerateStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **string** | Target thread (must already exist). | 
**ActionArgs** | Pointer to [**AiAiActionArgs**](AiAiActionArgs.md) | Per-request engine options: extra tools, reasoning, prompt override. | [optional] 
**EntityId** | Pointer to **string** | Optional entity (room) scope for profile resolution. | [optional] 
**ProfileId** | Pointer to **string** | Session-level profile override for this request only. | [optional] 

## Methods

### NewAiAiRegenerateStreamRequest

`func NewAiAiRegenerateStreamRequest(threadId string, ) *AiAiRegenerateStreamRequest`

NewAiAiRegenerateStreamRequest instantiates a new AiAiRegenerateStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAiRegenerateStreamRequestWithDefaults

`func NewAiAiRegenerateStreamRequestWithDefaults() *AiAiRegenerateStreamRequest`

NewAiAiRegenerateStreamRequestWithDefaults instantiates a new AiAiRegenerateStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiAiRegenerateStreamRequest) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiAiRegenerateStreamRequest) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiAiRegenerateStreamRequest) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetActionArgs

`func (o *AiAiRegenerateStreamRequest) GetActionArgs() AiAiActionArgs`

GetActionArgs returns the ActionArgs field if non-nil, zero value otherwise.

### GetActionArgsOk

`func (o *AiAiRegenerateStreamRequest) GetActionArgsOk() (*AiAiActionArgs, bool)`

GetActionArgsOk returns a tuple with the ActionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionArgs

`func (o *AiAiRegenerateStreamRequest) SetActionArgs(v AiAiActionArgs)`

SetActionArgs sets ActionArgs field to given value.

### HasActionArgs

`func (o *AiAiRegenerateStreamRequest) HasActionArgs() bool`

HasActionArgs returns a boolean if a field has been set.

### GetEntityId

`func (o *AiAiRegenerateStreamRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiAiRegenerateStreamRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiAiRegenerateStreamRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiAiRegenerateStreamRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetProfileId

`func (o *AiAiRegenerateStreamRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiAiRegenerateStreamRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiAiRegenerateStreamRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiAiRegenerateStreamRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


