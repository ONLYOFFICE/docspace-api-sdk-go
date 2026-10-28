# EditorToolCallStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolName** | **NullableString** | Which generation to run, which also decides the shape of the parameters below. | 
**Parameters** | [**EditorToolCallParametersDto**](EditorToolCallParametersDto.md) | The arguments of the generation named above. | 

## Methods

### NewEditorToolCallStateDto

`func NewEditorToolCallStateDto(toolName NullableString, parameters EditorToolCallParametersDto, ) *EditorToolCallStateDto`

NewEditorToolCallStateDto instantiates a new EditorToolCallStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditorToolCallStateDtoWithDefaults

`func NewEditorToolCallStateDtoWithDefaults() *EditorToolCallStateDto`

NewEditorToolCallStateDtoWithDefaults instantiates a new EditorToolCallStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolName

`func (o *EditorToolCallStateDto) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *EditorToolCallStateDto) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *EditorToolCallStateDto) SetToolName(v string)`

SetToolName sets ToolName field to given value.


### SetToolNameNil

`func (o *EditorToolCallStateDto) SetToolNameNil(b bool)`

 SetToolNameNil sets the value for ToolName to be an explicit nil

### UnsetToolName
`func (o *EditorToolCallStateDto) UnsetToolName()`

UnsetToolName ensures that no value is present for ToolName, not even an explicit nil
### GetParameters

`func (o *EditorToolCallStateDto) GetParameters() EditorToolCallParametersDto`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *EditorToolCallStateDto) GetParametersOk() (*EditorToolCallParametersDto, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *EditorToolCallStateDto) SetParameters(v EditorToolCallParametersDto)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


