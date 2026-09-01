# AiAttachmentsSaveFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**AiAttachmentsSaveFileRequestInput**](AiAttachmentsSaveFileRequestInput.md) |  | 
**EntityId** | Pointer to **string** | Optional entity (room) scope. | [optional] 

## Methods

### NewAiAttachmentsSaveFileRequest

`func NewAiAttachmentsSaveFileRequest(input AiAttachmentsSaveFileRequestInput, ) *AiAttachmentsSaveFileRequest`

NewAiAttachmentsSaveFileRequest instantiates a new AiAttachmentsSaveFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAttachmentsSaveFileRequestWithDefaults

`func NewAiAttachmentsSaveFileRequestWithDefaults() *AiAttachmentsSaveFileRequest`

NewAiAttachmentsSaveFileRequestWithDefaults instantiates a new AiAttachmentsSaveFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *AiAttachmentsSaveFileRequest) GetInput() AiAttachmentsSaveFileRequestInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AiAttachmentsSaveFileRequest) GetInputOk() (*AiAttachmentsSaveFileRequestInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AiAttachmentsSaveFileRequest) SetInput(v AiAttachmentsSaveFileRequestInput)`

SetInput sets Input field to given value.


### GetEntityId

`func (o *AiAttachmentsSaveFileRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiAttachmentsSaveFileRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiAttachmentsSaveFileRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiAttachmentsSaveFileRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


