# AiToolsUpdateCustomServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Config** | **map[string]interface{}** | One MCP server configuration. The shape is intentionally open — MCP allows per-transport fields (`command`/`args` for stdio, `url` for HTTP, plus env, headers, etc.) and the storage layer stays agnostic to which transport is in use. | 
**EntityId** | Pointer to **string** |  | [optional] 

## Methods

### NewAiToolsUpdateCustomServerRequest

`func NewAiToolsUpdateCustomServerRequest(name string, config map[string]interface{}, ) *AiToolsUpdateCustomServerRequest`

NewAiToolsUpdateCustomServerRequest instantiates a new AiToolsUpdateCustomServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiToolsUpdateCustomServerRequestWithDefaults

`func NewAiToolsUpdateCustomServerRequestWithDefaults() *AiToolsUpdateCustomServerRequest`

NewAiToolsUpdateCustomServerRequestWithDefaults instantiates a new AiToolsUpdateCustomServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiToolsUpdateCustomServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiToolsUpdateCustomServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiToolsUpdateCustomServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *AiToolsUpdateCustomServerRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AiToolsUpdateCustomServerRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AiToolsUpdateCustomServerRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetEntityId

`func (o *AiToolsUpdateCustomServerRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiToolsUpdateCustomServerRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiToolsUpdateCustomServerRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiToolsUpdateCustomServerRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


