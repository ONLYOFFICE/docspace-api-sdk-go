# AiToolsSetDisabledRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerType** | **string** |  | 
**ToolNames** | **[]string** | Tool names to disable. | 
**EntityId** | Pointer to **string** |  | [optional] 

## Methods

### NewAiToolsSetDisabledRequest

`func NewAiToolsSetDisabledRequest(serverType string, toolNames []string, ) *AiToolsSetDisabledRequest`

NewAiToolsSetDisabledRequest instantiates a new AiToolsSetDisabledRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiToolsSetDisabledRequestWithDefaults

`func NewAiToolsSetDisabledRequestWithDefaults() *AiToolsSetDisabledRequest`

NewAiToolsSetDisabledRequestWithDefaults instantiates a new AiToolsSetDisabledRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerType

`func (o *AiToolsSetDisabledRequest) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *AiToolsSetDisabledRequest) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *AiToolsSetDisabledRequest) SetServerType(v string)`

SetServerType sets ServerType field to given value.


### GetToolNames

`func (o *AiToolsSetDisabledRequest) GetToolNames() []string`

GetToolNames returns the ToolNames field if non-nil, zero value otherwise.

### GetToolNamesOk

`func (o *AiToolsSetDisabledRequest) GetToolNamesOk() (*[]string, bool)`

GetToolNamesOk returns a tuple with the ToolNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolNames

`func (o *AiToolsSetDisabledRequest) SetToolNames(v []string)`

SetToolNames sets ToolNames field to given value.


### GetEntityId

`func (o *AiToolsSetDisabledRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiToolsSetDisabledRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiToolsSetDisabledRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiToolsSetDisabledRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


