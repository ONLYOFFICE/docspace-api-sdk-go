# AiToolsSetAllowAlwaysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerType** | **string** |  | 
**ToolName** | **string** |  | 
**Value** | **bool** | Whether the tool is always allowed. | 
**EntityId** | Pointer to **string** |  | [optional] 

## Methods

### NewAiToolsSetAllowAlwaysRequest

`func NewAiToolsSetAllowAlwaysRequest(serverType string, toolName string, value bool, ) *AiToolsSetAllowAlwaysRequest`

NewAiToolsSetAllowAlwaysRequest instantiates a new AiToolsSetAllowAlwaysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiToolsSetAllowAlwaysRequestWithDefaults

`func NewAiToolsSetAllowAlwaysRequestWithDefaults() *AiToolsSetAllowAlwaysRequest`

NewAiToolsSetAllowAlwaysRequestWithDefaults instantiates a new AiToolsSetAllowAlwaysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerType

`func (o *AiToolsSetAllowAlwaysRequest) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *AiToolsSetAllowAlwaysRequest) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *AiToolsSetAllowAlwaysRequest) SetServerType(v string)`

SetServerType sets ServerType field to given value.


### GetToolName

`func (o *AiToolsSetAllowAlwaysRequest) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *AiToolsSetAllowAlwaysRequest) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *AiToolsSetAllowAlwaysRequest) SetToolName(v string)`

SetToolName sets ToolName field to given value.


### GetValue

`func (o *AiToolsSetAllowAlwaysRequest) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AiToolsSetAllowAlwaysRequest) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AiToolsSetAllowAlwaysRequest) SetValue(v bool)`

SetValue sets Value field to given value.


### GetEntityId

`func (o *AiToolsSetAllowAlwaysRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiToolsSetAllowAlwaysRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiToolsSetAllowAlwaysRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiToolsSetAllowAlwaysRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


