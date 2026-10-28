# RunBuilderScriptRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **NullableString** | The document builder script. It names a portal file by its id where the document builder documentation  writes an address - `builder.OpenFile(1234)`; an address written out in the script is refused. | 
**FolderId** | Pointer to **NullableInt32** | The id of the folder for the produced files that `outputs` does not list, as reported by a folder listing  such as `GET api/2.0/files/{folderId}`. | [optional] 
**Outputs** | Pointer to [**map[string]RunBuilderScriptOutputRequest**](RunBuilderScriptOutputRequest.md) | Destinations of the produced files, keyed by the name the script saves each one under - the second argument of  `builder.SaveFile(docx, result.docx)`. | [optional] 
**Argument** | Pointer to [**RunBuilderScriptRequestDtoArgument**](RunBuilderScriptRequestDtoArgument.md) |  | [optional] 

## Methods

### NewRunBuilderScriptRequestDto

`func NewRunBuilderScriptRequestDto(script NullableString, ) *RunBuilderScriptRequestDto`

NewRunBuilderScriptRequestDto instantiates a new RunBuilderScriptRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunBuilderScriptRequestDtoWithDefaults

`func NewRunBuilderScriptRequestDtoWithDefaults() *RunBuilderScriptRequestDto`

NewRunBuilderScriptRequestDtoWithDefaults instantiates a new RunBuilderScriptRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *RunBuilderScriptRequestDto) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *RunBuilderScriptRequestDto) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *RunBuilderScriptRequestDto) SetScript(v string)`

SetScript sets Script field to given value.


### SetScriptNil

`func (o *RunBuilderScriptRequestDto) SetScriptNil(b bool)`

 SetScriptNil sets the value for Script to be an explicit nil

### UnsetScript
`func (o *RunBuilderScriptRequestDto) UnsetScript()`

UnsetScript ensures that no value is present for Script, not even an explicit nil
### GetFolderId

`func (o *RunBuilderScriptRequestDto) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RunBuilderScriptRequestDto) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RunBuilderScriptRequestDto) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RunBuilderScriptRequestDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *RunBuilderScriptRequestDto) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *RunBuilderScriptRequestDto) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetOutputs

`func (o *RunBuilderScriptRequestDto) GetOutputs() map[string]RunBuilderScriptOutputRequest`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *RunBuilderScriptRequestDto) GetOutputsOk() (*map[string]RunBuilderScriptOutputRequest, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *RunBuilderScriptRequestDto) SetOutputs(v map[string]RunBuilderScriptOutputRequest)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *RunBuilderScriptRequestDto) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetArgument

`func (o *RunBuilderScriptRequestDto) GetArgument() RunBuilderScriptRequestDtoArgument`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *RunBuilderScriptRequestDto) GetArgumentOk() (*RunBuilderScriptRequestDtoArgument, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *RunBuilderScriptRequestDto) SetArgument(v RunBuilderScriptRequestDtoArgument)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *RunBuilderScriptRequestDto) HasArgument() bool`

HasArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


