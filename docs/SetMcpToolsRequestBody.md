# SetMcpToolsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisabledTools** | **[]string** | List of tool names to disable. Tools not included in this list will remain enabled. Pass an empty list to enable all tools. | 

## Methods

### NewSetMcpToolsRequestBody

`func NewSetMcpToolsRequestBody(disabledTools []string, ) *SetMcpToolsRequestBody`

NewSetMcpToolsRequestBody instantiates a new SetMcpToolsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetMcpToolsRequestBodyWithDefaults

`func NewSetMcpToolsRequestBodyWithDefaults() *SetMcpToolsRequestBody`

NewSetMcpToolsRequestBodyWithDefaults instantiates a new SetMcpToolsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabledTools

`func (o *SetMcpToolsRequestBody) GetDisabledTools() []string`

GetDisabledTools returns the DisabledTools field if non-nil, zero value otherwise.

### GetDisabledToolsOk

`func (o *SetMcpToolsRequestBody) GetDisabledToolsOk() (*[]string, bool)`

GetDisabledToolsOk returns a tuple with the DisabledTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledTools

`func (o *SetMcpToolsRequestBody) SetDisabledTools(v []string)`

SetDisabledTools sets DisabledTools field to given value.


### SetDisabledToolsNil

`func (o *SetMcpToolsRequestBody) SetDisabledToolsNil(b bool)`

 SetDisabledToolsNil sets the value for DisabledTools to be an explicit nil

### UnsetDisabledTools
`func (o *SetMcpToolsRequestBody) UnsetDisabledTools()`

UnsetDisabledTools ensures that no value is present for DisabledTools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


