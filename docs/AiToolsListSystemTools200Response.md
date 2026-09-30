# AiToolsListSystemTools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**map[string][]AiTMCPItem**](array.md) | Tools by server name, covering both the host-configured system servers and the custom MCP servers registered for this scope. | 
**Errors** | **map[string]string** | Why a registered custom server could not be reached, keyed by server name. A server that answered is absent from this map. | 
**System** | **[]string** | Names of the host-configured system servers among the keys of `groups`; everything else there was registered as a custom server. | 

## Methods

### NewAiToolsListSystemTools200Response

`func NewAiToolsListSystemTools200Response(groups map[string][]AiTMCPItem, errors map[string]string, system []string, ) *AiToolsListSystemTools200Response`

NewAiToolsListSystemTools200Response instantiates a new AiToolsListSystemTools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiToolsListSystemTools200ResponseWithDefaults

`func NewAiToolsListSystemTools200ResponseWithDefaults() *AiToolsListSystemTools200Response`

NewAiToolsListSystemTools200ResponseWithDefaults instantiates a new AiToolsListSystemTools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *AiToolsListSystemTools200Response) GetGroups() map[string][]AiTMCPItem`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AiToolsListSystemTools200Response) GetGroupsOk() (*map[string][]AiTMCPItem, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AiToolsListSystemTools200Response) SetGroups(v map[string][]AiTMCPItem)`

SetGroups sets Groups field to given value.


### GetErrors

`func (o *AiToolsListSystemTools200Response) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AiToolsListSystemTools200Response) GetErrorsOk() (*map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AiToolsListSystemTools200Response) SetErrors(v map[string]string)`

SetErrors sets Errors field to given value.


### GetSystem

`func (o *AiToolsListSystemTools200Response) GetSystem() []string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AiToolsListSystemTools200Response) GetSystemOk() (*[]string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AiToolsListSystemTools200Response) SetSystem(v []string)`

SetSystem sets System field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


