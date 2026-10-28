# AiEditorToolsCall200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | What the tool produced, as text. A structured result is JSON-encoded, and a tool that failed reports its error here rather than through a status code. | 

## Methods

### NewAiEditorToolsCall200Response

`func NewAiEditorToolsCall200Response(result string, ) *AiEditorToolsCall200Response`

NewAiEditorToolsCall200Response instantiates a new AiEditorToolsCall200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEditorToolsCall200ResponseWithDefaults

`func NewAiEditorToolsCall200ResponseWithDefaults() *AiEditorToolsCall200Response`

NewAiEditorToolsCall200ResponseWithDefaults instantiates a new AiEditorToolsCall200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *AiEditorToolsCall200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AiEditorToolsCall200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AiEditorToolsCall200Response) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


