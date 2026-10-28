# AiEditorToolsList200ResponseToolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Tool name, as it is passed back to the call endpoint. | 
**Description** | **string** | What the tool does, empty when the server declares nothing. | 
**InputSchema** | **map[string]interface{}** | JSON Schema of the tool arguments. | 
**RequireApproval** | **bool** | Whether the editor has to ask the user before running the tool. Read-only operations arrive with this off. | 

## Methods

### NewAiEditorToolsList200ResponseToolsInner

`func NewAiEditorToolsList200ResponseToolsInner(name string, description string, inputSchema map[string]*interface{}, requireApproval bool, ) *AiEditorToolsList200ResponseToolsInner`

NewAiEditorToolsList200ResponseToolsInner instantiates a new AiEditorToolsList200ResponseToolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEditorToolsList200ResponseToolsInnerWithDefaults

`func NewAiEditorToolsList200ResponseToolsInnerWithDefaults() *AiEditorToolsList200ResponseToolsInner`

NewAiEditorToolsList200ResponseToolsInnerWithDefaults instantiates a new AiEditorToolsList200ResponseToolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiEditorToolsList200ResponseToolsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiEditorToolsList200ResponseToolsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiEditorToolsList200ResponseToolsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AiEditorToolsList200ResponseToolsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiEditorToolsList200ResponseToolsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiEditorToolsList200ResponseToolsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInputSchema

`func (o *AiEditorToolsList200ResponseToolsInner) GetInputSchema() map[string]*interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *AiEditorToolsList200ResponseToolsInner) GetInputSchemaOk() (*map[string]*interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *AiEditorToolsList200ResponseToolsInner) SetInputSchema(v map[string]*interface{})`

SetInputSchema sets InputSchema field to given value.


### GetRequireApproval

`func (o *AiEditorToolsList200ResponseToolsInner) GetRequireApproval() bool`

GetRequireApproval returns the RequireApproval field if non-nil, zero value otherwise.

### GetRequireApprovalOk

`func (o *AiEditorToolsList200ResponseToolsInner) GetRequireApprovalOk() (*bool, bool)`

GetRequireApprovalOk returns a tuple with the RequireApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireApproval

`func (o *AiEditorToolsList200ResponseToolsInner) SetRequireApproval(v bool)`

SetRequireApproval sets RequireApproval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


