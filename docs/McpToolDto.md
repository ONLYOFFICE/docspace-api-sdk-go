# McpToolDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Name of the tool as reported by the MCP server. | 
**Enabled** | Pointer to **bool** | Indicates whether this tool is enabled (true) or disabled (false) for use in AI chat sessions within the room. | [optional] 

## Methods

### NewMcpToolDto

`func NewMcpToolDto(name NullableString, ) *McpToolDto`

NewMcpToolDto instantiates a new McpToolDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpToolDtoWithDefaults

`func NewMcpToolDtoWithDefaults() *McpToolDto`

NewMcpToolDtoWithDefaults instantiates a new McpToolDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *McpToolDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpToolDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpToolDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *McpToolDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *McpToolDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *McpToolDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *McpToolDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *McpToolDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *McpToolDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


