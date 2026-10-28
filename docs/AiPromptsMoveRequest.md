# AiPromptsMoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Prompt id to move. | 
**FolderId** | **NullableString** | Target folder id, or `null` for root. | 

## Methods

### NewAiPromptsMoveRequest

`func NewAiPromptsMoveRequest(id string, folderId NullableString, ) *AiPromptsMoveRequest`

NewAiPromptsMoveRequest instantiates a new AiPromptsMoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPromptsMoveRequestWithDefaults

`func NewAiPromptsMoveRequestWithDefaults() *AiPromptsMoveRequest`

NewAiPromptsMoveRequestWithDefaults instantiates a new AiPromptsMoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiPromptsMoveRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiPromptsMoveRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiPromptsMoveRequest) SetId(v string)`

SetId sets Id field to given value.


### GetFolderId

`func (o *AiPromptsMoveRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AiPromptsMoveRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AiPromptsMoveRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### SetFolderIdNil

`func (o *AiPromptsMoveRequest) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *AiPromptsMoveRequest) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


