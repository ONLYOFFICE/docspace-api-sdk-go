# AiMCPItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Tool name as registered on the MCP server (e.g. `web_search`, `insert_text`). | 
**Description** | **string** | Human-readable description shown to the AI model and in the tools list UI. | 
**InputSchema** | **map[string]interface{}** | JSON Schema describing the tool's input parameters. | 
**Enabled** | Pointer to **bool** | Whether this tool is currently enabled. Disabled tools are hidden from the AI model. | [optional] 
**ServerType** | Pointer to **string** | The server or tool group the tool belongs to (an MCP server name or a host tool group ID): the key a disabled tool is stored under, so a tool sent in `actionArgs.tools` is still matched to its group. | [optional] 
**RequireApproval** | Pointer to **bool** | Whether the tool asks for approval before it runs, under the `auto` tool permission mode: `false` runs it without asking; `true` or absent leaves it to the always-allow list and the tool's MCP annotations (see `annotations`). Under `ask` every call asks and under `allow` none does, whatever this says. Host tools default to `false`; MCP and custom-server tools leave it unset. | [optional] 
**Annotations** | Pointer to [**AiToolAnnotations**](AiToolAnnotations.md) | The MCP tool annotations as the server declared them in `tools/list`; never sent to the model. Under the `auto` tool permission mode a tool that cannot destroy anything runs without approval - one annotated `readOnlyHint: true`, or one with an explicit `destructiveHint: false`; a destructive tool, and one with no annotations (the protocol's default is destructive), keeps asking. | [optional] 

## Methods

### NewAiMCPItem

`func NewAiMCPItem(name string, description string, inputSchema map[string]interface{}, ) *AiMCPItem`

NewAiMCPItem instantiates a new AiMCPItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMCPItemWithDefaults

`func NewAiMCPItemWithDefaults() *AiMCPItem`

NewAiMCPItemWithDefaults instantiates a new AiMCPItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiMCPItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiMCPItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiMCPItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AiMCPItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiMCPItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiMCPItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInputSchema

`func (o *AiMCPItem) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *AiMCPItem) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *AiMCPItem) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetEnabled

`func (o *AiMCPItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AiMCPItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AiMCPItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AiMCPItem) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServerType

`func (o *AiMCPItem) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *AiMCPItem) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *AiMCPItem) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *AiMCPItem) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetRequireApproval

`func (o *AiMCPItem) GetRequireApproval() bool`

GetRequireApproval returns the RequireApproval field if non-nil, zero value otherwise.

### GetRequireApprovalOk

`func (o *AiMCPItem) GetRequireApprovalOk() (*bool, bool)`

GetRequireApprovalOk returns a tuple with the RequireApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireApproval

`func (o *AiMCPItem) SetRequireApproval(v bool)`

SetRequireApproval sets RequireApproval field to given value.

### HasRequireApproval

`func (o *AiMCPItem) HasRequireApproval() bool`

HasRequireApproval returns a boolean if a field has been set.

### GetAnnotations

`func (o *AiMCPItem) GetAnnotations() AiToolAnnotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AiMCPItem) GetAnnotationsOk() (*AiToolAnnotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AiMCPItem) SetAnnotations(v AiToolAnnotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AiMCPItem) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


