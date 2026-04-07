# McpServerStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the MCP server. | [optional] 
**Name** | **NullableString** | Display name of the MCP server. | 
**ServerType** | Pointer to [**ServerType**](ServerType.md) |  | [optional] 
**Connected** | Pointer to **bool** | Indicates whether the current user has an active connection to this server. For direct-connection servers this is always true; for OAuth-based servers it reflects whether the user has completed authorization. | [optional] 
**Icon** | Pointer to [**Icon**](Icon.md) |  | [optional] 
**NeedReset** | Pointer to **bool** | Indicates whether the server requires a configuration reset due to connectivity or credential issues. | [optional] 

## Methods

### NewMcpServerStatusDto

`func NewMcpServerStatusDto(name NullableString, ) *McpServerStatusDto`

NewMcpServerStatusDto instantiates a new McpServerStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerStatusDtoWithDefaults

`func NewMcpServerStatusDtoWithDefaults() *McpServerStatusDto`

NewMcpServerStatusDtoWithDefaults instantiates a new McpServerStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McpServerStatusDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McpServerStatusDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McpServerStatusDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *McpServerStatusDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *McpServerStatusDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerStatusDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerStatusDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *McpServerStatusDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *McpServerStatusDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServerType

`func (o *McpServerStatusDto) GetServerType() ServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *McpServerStatusDto) GetServerTypeOk() (*ServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *McpServerStatusDto) SetServerType(v ServerType)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *McpServerStatusDto) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetConnected

`func (o *McpServerStatusDto) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *McpServerStatusDto) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *McpServerStatusDto) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *McpServerStatusDto) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetIcon

`func (o *McpServerStatusDto) GetIcon() Icon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *McpServerStatusDto) GetIconOk() (*Icon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *McpServerStatusDto) SetIcon(v Icon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *McpServerStatusDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetNeedReset

`func (o *McpServerStatusDto) GetNeedReset() bool`

GetNeedReset returns the NeedReset field if non-nil, zero value otherwise.

### GetNeedResetOk

`func (o *McpServerStatusDto) GetNeedResetOk() (*bool, bool)`

GetNeedResetOk returns a tuple with the NeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedReset

`func (o *McpServerStatusDto) SetNeedReset(v bool)`

SetNeedReset sets NeedReset field to given value.

### HasNeedReset

`func (o *McpServerStatusDto) HasNeedReset() bool`

HasNeedReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


