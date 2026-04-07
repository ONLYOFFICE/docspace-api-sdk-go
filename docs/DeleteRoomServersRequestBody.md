# DeleteRoomServersRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | **[]string** | Set of unique identifiers of MCP servers to remove from the room. Associated connections and tool configurations will also be cleaned up. | 

## Methods

### NewDeleteRoomServersRequestBody

`func NewDeleteRoomServersRequestBody(servers []string, ) *DeleteRoomServersRequestBody`

NewDeleteRoomServersRequestBody instantiates a new DeleteRoomServersRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRoomServersRequestBodyWithDefaults

`func NewDeleteRoomServersRequestBodyWithDefaults() *DeleteRoomServersRequestBody`

NewDeleteRoomServersRequestBodyWithDefaults instantiates a new DeleteRoomServersRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *DeleteRoomServersRequestBody) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DeleteRoomServersRequestBody) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DeleteRoomServersRequestBody) SetServers(v []string)`

SetServers sets Servers field to given value.


### SetServersNil

`func (o *DeleteRoomServersRequestBody) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *DeleteRoomServersRequestBody) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


