# AiToolsAddCustomServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Server name (unique within scope). | 
**Config** | **map[string]interface{}** | Server transport configuration. | 
**EntityId** | Pointer to **string** |  | [optional] 

## Methods

### NewAiToolsAddCustomServerRequest

`func NewAiToolsAddCustomServerRequest(name string, config map[string]interface{}, ) *AiToolsAddCustomServerRequest`

NewAiToolsAddCustomServerRequest instantiates a new AiToolsAddCustomServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiToolsAddCustomServerRequestWithDefaults

`func NewAiToolsAddCustomServerRequestWithDefaults() *AiToolsAddCustomServerRequest`

NewAiToolsAddCustomServerRequestWithDefaults instantiates a new AiToolsAddCustomServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiToolsAddCustomServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiToolsAddCustomServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiToolsAddCustomServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *AiToolsAddCustomServerRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AiToolsAddCustomServerRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AiToolsAddCustomServerRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetEntityId

`func (o *AiToolsAddCustomServerRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiToolsAddCustomServerRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiToolsAddCustomServerRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiToolsAddCustomServerRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


