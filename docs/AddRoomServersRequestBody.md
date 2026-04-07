# AddRoomServersRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | **[]string** | Set of unique identifiers of MCP servers to associate with the room. A maximum of 5 servers can be assigned per room. | 

## Methods

### NewAddRoomServersRequestBody

`func NewAddRoomServersRequestBody(servers []string, ) *AddRoomServersRequestBody`

NewAddRoomServersRequestBody instantiates a new AddRoomServersRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRoomServersRequestBodyWithDefaults

`func NewAddRoomServersRequestBodyWithDefaults() *AddRoomServersRequestBody`

NewAddRoomServersRequestBodyWithDefaults instantiates a new AddRoomServersRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *AddRoomServersRequestBody) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *AddRoomServersRequestBody) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *AddRoomServersRequestBody) SetServers(v []string)`

SetServers sets Servers field to given value.


### SetServersNil

`func (o *AddRoomServersRequestBody) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *AddRoomServersRequestBody) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


