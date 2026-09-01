# AiTMCPItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Tool name as registered on the MCP server (e.g. `web_search`, `insert_text`). | 
**Description** | **string** | Human-readable description shown to the AI model and in the tools list UI. | 
**InputSchema** | **map[string]interface{}** | JSON Schema describing the tool's input parameters. | 
**Enabled** | Pointer to **bool** | Whether this tool is currently enabled. Disabled tools are hidden from the AI model. | [optional] 
**ServerType** | Pointer to **string** | Server type (MCP server name / host tool group id) this tool belongs to — the key the persisted disabled map is stored under. Set by the source that enumerated the tool, so a caller-supplied tool can still be attributed to its group after being flattened into a single list: that is what lets the engine apply the disabled map to `actionArgs.tools` instead of trusting the caller to pre-filter. Wire-serializable, so it survives a remote (server-side) engine. | [optional] 
**RequireApproval** | Pointer to **bool** | Whether the consumer must show an approval dialog before this tool runs. The engine reads it when deciding the `autoAllow` flag on a `tool-call-pending` event: `requireApproval === false` auto-allows the call (no dialog), `true` always prompts. `undefined` leaves the decision to the persisted always-allow list alone — so MCP / custom-server tools (which never set it) keep prompting as before, while host tools opt into auto-allow by default. Wire-serializable, so it survives a remote (server-side) engine. | [optional] 

## Methods

### NewAiTMCPItem

`func NewAiTMCPItem(name string, description string, inputSchema map[string]interface{}, ) *AiTMCPItem`

NewAiTMCPItem instantiates a new AiTMCPItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiTMCPItemWithDefaults

`func NewAiTMCPItemWithDefaults() *AiTMCPItem`

NewAiTMCPItemWithDefaults instantiates a new AiTMCPItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiTMCPItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiTMCPItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiTMCPItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AiTMCPItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiTMCPItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiTMCPItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInputSchema

`func (o *AiTMCPItem) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *AiTMCPItem) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *AiTMCPItem) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetEnabled

`func (o *AiTMCPItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AiTMCPItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AiTMCPItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AiTMCPItem) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServerType

`func (o *AiTMCPItem) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *AiTMCPItem) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *AiTMCPItem) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *AiTMCPItem) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetRequireApproval

`func (o *AiTMCPItem) GetRequireApproval() bool`

GetRequireApproval returns the RequireApproval field if non-nil, zero value otherwise.

### GetRequireApprovalOk

`func (o *AiTMCPItem) GetRequireApprovalOk() (*bool, bool)`

GetRequireApprovalOk returns a tuple with the RequireApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireApproval

`func (o *AiTMCPItem) SetRequireApproval(v bool)`

SetRequireApproval sets RequireApproval field to given value.

### HasRequireApproval

`func (o *AiTMCPItem) HasRequireApproval() bool`

HasRequireApproval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


