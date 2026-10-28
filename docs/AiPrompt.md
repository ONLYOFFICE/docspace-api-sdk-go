# AiPrompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique prompt identifier (UUID). | 
**Name** | **string** | Prompt display name shown in the prompt picker. | 
**Text** | **string** | Prompt template text. May contain placeholder tokens. | 
**FolderId** | Pointer to **string** | Optional parent folder ID. `undefined` means the prompt is at the root level. | [optional] 
**CreatedAt** | **float32** | Timestamp (ms since epoch) when the prompt was created. | 
**UpdatedAt** | **float32** | Timestamp (ms since epoch) of the last prompt modification. | 

## Methods

### NewAiPrompt

`func NewAiPrompt(id string, name string, text string, createdAt float32, updatedAt float32, ) *AiPrompt`

NewAiPrompt instantiates a new AiPrompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPromptWithDefaults

`func NewAiPromptWithDefaults() *AiPrompt`

NewAiPromptWithDefaults instantiates a new AiPrompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiPrompt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiPrompt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiPrompt) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AiPrompt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiPrompt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiPrompt) SetName(v string)`

SetName sets Name field to given value.


### GetText

`func (o *AiPrompt) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AiPrompt) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AiPrompt) SetText(v string)`

SetText sets Text field to given value.


### GetFolderId

`func (o *AiPrompt) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AiPrompt) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AiPrompt) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AiPrompt) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AiPrompt) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiPrompt) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiPrompt) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AiPrompt) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AiPrompt) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AiPrompt) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


