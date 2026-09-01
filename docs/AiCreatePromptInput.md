# AiCreatePromptInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The prompt name. | 
**Text** | **string** | The prompt body. | 
**FolderId** | Pointer to **NullableString** | The folder to file the prompt under. Omit or send null to leave it outside any folder. | [optional] 

## Methods

### NewAiCreatePromptInput

`func NewAiCreatePromptInput(name string, text string, ) *AiCreatePromptInput`

NewAiCreatePromptInput instantiates a new AiCreatePromptInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiCreatePromptInputWithDefaults

`func NewAiCreatePromptInputWithDefaults() *AiCreatePromptInput`

NewAiCreatePromptInputWithDefaults instantiates a new AiCreatePromptInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiCreatePromptInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiCreatePromptInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiCreatePromptInput) SetName(v string)`

SetName sets Name field to given value.


### GetText

`func (o *AiCreatePromptInput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AiCreatePromptInput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AiCreatePromptInput) SetText(v string)`

SetText sets Text field to given value.


### GetFolderId

`func (o *AiCreatePromptInput) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AiCreatePromptInput) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AiCreatePromptInput) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AiCreatePromptInput) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *AiCreatePromptInput) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *AiCreatePromptInput) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


