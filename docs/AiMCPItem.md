# AiMCPItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Tool name as registered on the MCP server (e.g. `web_search`, `insert_text`). | 
**Description** | **string** | Human-readable description shown to the AI model and in the tools list UI. | 
**InputSchema** | **map[string]interface{}** | JSON Schema describing the tool's input parameters. | 
**Enabled** | Pointer to **bool** | Whether this tool is currently enabled. Disabled tools are hidden from the AI model. | [optional] 
**ServerType** | Pointer to **string** | Server type (MCP server name / host tool group id) this tool belongs to — the key the persisted disabled map is stored under. Set by the source that enumerated the tool, so a caller-supplied tool can still be attributed to its group after being flattened into a single list: that is what lets the engine apply the disabled map to `actionArgs.tools` instead of trusting the caller to pre-filter. Wire-serializable, so it survives a remote (server-side) engine. | [optional] 
**RequireApproval** | Pointer to **bool** | Whether the consumer must show an approval dialog before this tool runs. Feeds the `autoAllow` flag on a `tool-call-pending` event together with the user's tool permission mode: `false` skips the dialog under the auto mode, the default (under ask only the persisted always-allow list does), `true` and `undefined` defer to that list; allow skips it for every tool. Host tools set it and default to `false`; MCP / custom-server tools leave it unset. Wire-serializable, so it survives a remote (server-side) engine. | [optional] 
**Annotations** | Pointer to [**AiToolAnnotations**](AiToolAnnotations.md) | The MCP tool annotations as the server declared them in `tools/list` (kept verbatim on the descriptor; never sent to the model). The approval flow reads two of them under the auto permission mode: `readOnlyHint: true` and `destructiveHint: false` run without the dialog, a destructive or unannotated tool keeps asking — see `resolveAutoAllow`. Wire-serializable. | [optional] 

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


