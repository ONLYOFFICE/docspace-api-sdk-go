# AiMetaGenerateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The entry's table. Only `file` for now: a folder or a room answers 400. | 
**Id** | [**AiMetaGenerateRequestId**](AiMetaGenerateRequestId.md) |  | 
**Prompt** | Pointer to **string** | What kind of metadata fields the caller is interested in; a hint for the model. | [optional] 

## Methods

### NewAiMetaGenerateRequest

`func NewAiMetaGenerateRequest(kind string, id AiMetaGenerateRequestId, ) *AiMetaGenerateRequest`

NewAiMetaGenerateRequest instantiates a new AiMetaGenerateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMetaGenerateRequestWithDefaults

`func NewAiMetaGenerateRequestWithDefaults() *AiMetaGenerateRequest`

NewAiMetaGenerateRequestWithDefaults instantiates a new AiMetaGenerateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AiMetaGenerateRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiMetaGenerateRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiMetaGenerateRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *AiMetaGenerateRequest) GetId() AiMetaGenerateRequestId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiMetaGenerateRequest) GetIdOk() (*AiMetaGenerateRequestId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiMetaGenerateRequest) SetId(v AiMetaGenerateRequestId)`

SetId sets Id field to given value.


### GetPrompt

`func (o *AiMetaGenerateRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiMetaGenerateRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiMetaGenerateRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiMetaGenerateRequest) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


