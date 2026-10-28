# AiPromptBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **float32** | The bundle format version, so an import can migrate an older export. | 
**Folders** | [**[]AiPromptFolder**](AiPromptFolder.md) | Every exported prompt folder. | 
**Prompts** | [**[]AiPrompt**](AiPrompt.md) | Every exported prompt. | 

## Methods

### NewAiPromptBundle

`func NewAiPromptBundle(version float32, folders []AiPromptFolder, prompts []AiPrompt, ) *AiPromptBundle`

NewAiPromptBundle instantiates a new AiPromptBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPromptBundleWithDefaults

`func NewAiPromptBundleWithDefaults() *AiPromptBundle`

NewAiPromptBundleWithDefaults instantiates a new AiPromptBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *AiPromptBundle) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AiPromptBundle) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AiPromptBundle) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetFolders

`func (o *AiPromptBundle) GetFolders() []AiPromptFolder`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *AiPromptBundle) GetFoldersOk() (*[]AiPromptFolder, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *AiPromptBundle) SetFolders(v []AiPromptFolder)`

SetFolders sets Folders field to given value.


### GetPrompts

`func (o *AiPromptBundle) GetPrompts() []AiPrompt`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *AiPromptBundle) GetPromptsOk() (*[]AiPrompt, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *AiPromptBundle) SetPrompts(v []AiPrompt)`

SetPrompts sets Prompts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


