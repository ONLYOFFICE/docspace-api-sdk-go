# DeleteServersRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | **[]string** | Set of unique identifiers of the MCP servers to permanently remove. All room associations and connection data will also be deleted. | 

## Methods

### NewDeleteServersRequestBody

`func NewDeleteServersRequestBody(servers []string, ) *DeleteServersRequestBody`

NewDeleteServersRequestBody instantiates a new DeleteServersRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteServersRequestBodyWithDefaults

`func NewDeleteServersRequestBodyWithDefaults() *DeleteServersRequestBody`

NewDeleteServersRequestBodyWithDefaults instantiates a new DeleteServersRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *DeleteServersRequestBody) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DeleteServersRequestBody) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DeleteServersRequestBody) SetServers(v []string)`

SetServers sets Servers field to given value.


### SetServersNil

`func (o *DeleteServersRequestBody) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *DeleteServersRequestBody) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


