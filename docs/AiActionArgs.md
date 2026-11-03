# AiActionArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | Pointer to [**[]AiMCPItem**](AiMCPItem.md) | Extra tools offered to the model for this request. | [optional] 
**IsReasoning** | Pointer to **bool** | Legacy extended-thinking switch; stands for `medium`. `reasoningLevel` wins when both are set. | [optional] 
**ReasoningLevel** | Pointer to [**AiReasoningLevel**](AiReasoningLevel.md) | Depth of extended thinking for the round; providers clamp it to what the model accepts. | [optional] 
**Prompt** | Pointer to [**AiActionArgsPrompt**](AiActionArgsPrompt.md) |  | [optional] 
**ToolInstructions** | Pointer to **map[string]string** | Usage notes of the MCP servers the caller holds the connection to, keyed by server type; folded into the system prompt as advisory text. | [optional] 
**ContextRooms** | Pointer to [**[]AiContextRoomRef**](AiContextRoomRef.md) | Instructions rooms connected to the chat as context, each with the skills the round may use. | [optional] 
**KnowledgeRooms** | Pointer to [**[]AiKnowledgeRoomRef**](AiKnowledgeRoomRef.md) | Knowledge rooms connected to the chat; the knowledge search tool of the round is limited to them. Each `id` has to be a DocSpace room id. | [optional] 
**KnowledgeRoom** | Pointer to [**AiKnowledgeRoomRef**](AiKnowledgeRoomRef.md) | A single knowledge room, as an older client sends it. Read only when `knowledgeRooms` is absent; send `knowledgeRooms` instead. | [optional] 
**FolderContext** | Pointer to [**AiFolderContextRef**](AiFolderContextRef.md) | The folder connected to the chat as the round's context. It becomes the location the model is told the user works in, and the entry the round is attributed to. | [optional] 
**EntityId** | Pointer to **string** | The entry the round is attributed to. A hint, resolved against the caller's own access. | [optional] 

## Methods

### NewAiActionArgs

`func NewAiActionArgs() *AiActionArgs`

NewAiActionArgs instantiates a new AiActionArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiActionArgsWithDefaults

`func NewAiActionArgsWithDefaults() *AiActionArgs`

NewAiActionArgsWithDefaults instantiates a new AiActionArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *AiActionArgs) GetTools() []AiMCPItem`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiActionArgs) GetToolsOk() (*[]AiMCPItem, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiActionArgs) SetTools(v []AiMCPItem)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiActionArgs) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetIsReasoning

`func (o *AiActionArgs) GetIsReasoning() bool`

GetIsReasoning returns the IsReasoning field if non-nil, zero value otherwise.

### GetIsReasoningOk

`func (o *AiActionArgs) GetIsReasoningOk() (*bool, bool)`

GetIsReasoningOk returns a tuple with the IsReasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReasoning

`func (o *AiActionArgs) SetIsReasoning(v bool)`

SetIsReasoning sets IsReasoning field to given value.

### HasIsReasoning

`func (o *AiActionArgs) HasIsReasoning() bool`

HasIsReasoning returns a boolean if a field has been set.

### GetReasoningLevel

`func (o *AiActionArgs) GetReasoningLevel() AiReasoningLevel`

GetReasoningLevel returns the ReasoningLevel field if non-nil, zero value otherwise.

### GetReasoningLevelOk

`func (o *AiActionArgs) GetReasoningLevelOk() (*AiReasoningLevel, bool)`

GetReasoningLevelOk returns a tuple with the ReasoningLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningLevel

`func (o *AiActionArgs) SetReasoningLevel(v AiReasoningLevel)`

SetReasoningLevel sets ReasoningLevel field to given value.

### HasReasoningLevel

`func (o *AiActionArgs) HasReasoningLevel() bool`

HasReasoningLevel returns a boolean if a field has been set.

### GetPrompt

`func (o *AiActionArgs) GetPrompt() AiActionArgsPrompt`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiActionArgs) GetPromptOk() (*AiActionArgsPrompt, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiActionArgs) SetPrompt(v AiActionArgsPrompt)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiActionArgs) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetToolInstructions

`func (o *AiActionArgs) GetToolInstructions() map[string]string`

GetToolInstructions returns the ToolInstructions field if non-nil, zero value otherwise.

### GetToolInstructionsOk

`func (o *AiActionArgs) GetToolInstructionsOk() (*map[string]string, bool)`

GetToolInstructionsOk returns a tuple with the ToolInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolInstructions

`func (o *AiActionArgs) SetToolInstructions(v map[string]string)`

SetToolInstructions sets ToolInstructions field to given value.

### HasToolInstructions

`func (o *AiActionArgs) HasToolInstructions() bool`

HasToolInstructions returns a boolean if a field has been set.

### GetContextRooms

`func (o *AiActionArgs) GetContextRooms() []AiContextRoomRef`

GetContextRooms returns the ContextRooms field if non-nil, zero value otherwise.

### GetContextRoomsOk

`func (o *AiActionArgs) GetContextRoomsOk() (*[]AiContextRoomRef, bool)`

GetContextRoomsOk returns a tuple with the ContextRooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRooms

`func (o *AiActionArgs) SetContextRooms(v []AiContextRoomRef)`

SetContextRooms sets ContextRooms field to given value.

### HasContextRooms

`func (o *AiActionArgs) HasContextRooms() bool`

HasContextRooms returns a boolean if a field has been set.

### GetKnowledgeRooms

`func (o *AiActionArgs) GetKnowledgeRooms() []AiKnowledgeRoomRef`

GetKnowledgeRooms returns the KnowledgeRooms field if non-nil, zero value otherwise.

### GetKnowledgeRoomsOk

`func (o *AiActionArgs) GetKnowledgeRoomsOk() (*[]AiKnowledgeRoomRef, bool)`

GetKnowledgeRoomsOk returns a tuple with the KnowledgeRooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeRooms

`func (o *AiActionArgs) SetKnowledgeRooms(v []AiKnowledgeRoomRef)`

SetKnowledgeRooms sets KnowledgeRooms field to given value.

### HasKnowledgeRooms

`func (o *AiActionArgs) HasKnowledgeRooms() bool`

HasKnowledgeRooms returns a boolean if a field has been set.

### GetKnowledgeRoom

`func (o *AiActionArgs) GetKnowledgeRoom() AiKnowledgeRoomRef`

GetKnowledgeRoom returns the KnowledgeRoom field if non-nil, zero value otherwise.

### GetKnowledgeRoomOk

`func (o *AiActionArgs) GetKnowledgeRoomOk() (*AiKnowledgeRoomRef, bool)`

GetKnowledgeRoomOk returns a tuple with the KnowledgeRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeRoom

`func (o *AiActionArgs) SetKnowledgeRoom(v AiKnowledgeRoomRef)`

SetKnowledgeRoom sets KnowledgeRoom field to given value.

### HasKnowledgeRoom

`func (o *AiActionArgs) HasKnowledgeRoom() bool`

HasKnowledgeRoom returns a boolean if a field has been set.

### GetFolderContext

`func (o *AiActionArgs) GetFolderContext() AiFolderContextRef`

GetFolderContext returns the FolderContext field if non-nil, zero value otherwise.

### GetFolderContextOk

`func (o *AiActionArgs) GetFolderContextOk() (*AiFolderContextRef, bool)`

GetFolderContextOk returns a tuple with the FolderContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderContext

`func (o *AiActionArgs) SetFolderContext(v AiFolderContextRef)`

SetFolderContext sets FolderContext field to given value.

### HasFolderContext

`func (o *AiActionArgs) HasFolderContext() bool`

HasFolderContext returns a boolean if a field has been set.

### GetEntityId

`func (o *AiActionArgs) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiActionArgs) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiActionArgs) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiActionArgs) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


