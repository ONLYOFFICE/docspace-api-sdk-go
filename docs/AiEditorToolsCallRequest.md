# AiEditorToolsCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the tool to run, as listed by the tools endpoint. A name that is unknown or excluded from the editor is rejected with 400. | 
**Arguments** | Pointer to **map[string]interface{}** | Arguments for the tool, shaped by that tool's own input schema. Treated as empty when it is not an object. | [optional] 
**EntityId** | Pointer to **string** | Room the call is scoped to. Left out for a portal-wide call. | [optional] 

## Methods

### NewAiEditorToolsCallRequest

`func NewAiEditorToolsCallRequest(name string, ) *AiEditorToolsCallRequest`

NewAiEditorToolsCallRequest instantiates a new AiEditorToolsCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEditorToolsCallRequestWithDefaults

`func NewAiEditorToolsCallRequestWithDefaults() *AiEditorToolsCallRequest`

NewAiEditorToolsCallRequestWithDefaults instantiates a new AiEditorToolsCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiEditorToolsCallRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiEditorToolsCallRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiEditorToolsCallRequest) SetName(v string)`

SetName sets Name field to given value.


### GetArguments

`func (o *AiEditorToolsCallRequest) GetArguments() map[string]*interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *AiEditorToolsCallRequest) GetArgumentsOk() (*map[string]*interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *AiEditorToolsCallRequest) SetArguments(v map[string]*interface{})`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *AiEditorToolsCallRequest) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntityId

`func (o *AiEditorToolsCallRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiEditorToolsCallRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiEditorToolsCallRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiEditorToolsCallRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


