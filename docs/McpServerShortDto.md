# McpServerShortDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the MCP server. | [optional] 
**Name** | Pointer to **NullableString** | Display name of the MCP server. | [optional] 
**ServerType** | Pointer to [**ServerType**](ServerType.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the server is currently enabled and available for room assignment. | [optional] 
**Icon** | Pointer to [**Icon**](Icon.md) |  | [optional] 
**NeedReset** | Pointer to **bool** | Indicates whether the server requires a configuration reset due to connectivity or credential issues. | [optional] 

## Methods

### NewMcpServerShortDto

`func NewMcpServerShortDto() *McpServerShortDto`

NewMcpServerShortDto instantiates a new McpServerShortDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerShortDtoWithDefaults

`func NewMcpServerShortDtoWithDefaults() *McpServerShortDto`

NewMcpServerShortDtoWithDefaults instantiates a new McpServerShortDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McpServerShortDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McpServerShortDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McpServerShortDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *McpServerShortDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *McpServerShortDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerShortDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerShortDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *McpServerShortDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *McpServerShortDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *McpServerShortDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServerType

`func (o *McpServerShortDto) GetServerType() ServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *McpServerShortDto) GetServerTypeOk() (*ServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *McpServerShortDto) SetServerType(v ServerType)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *McpServerShortDto) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetEnabled

`func (o *McpServerShortDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *McpServerShortDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *McpServerShortDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *McpServerShortDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIcon

`func (o *McpServerShortDto) GetIcon() Icon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *McpServerShortDto) GetIconOk() (*Icon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *McpServerShortDto) SetIcon(v Icon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *McpServerShortDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetNeedReset

`func (o *McpServerShortDto) GetNeedReset() bool`

GetNeedReset returns the NeedReset field if non-nil, zero value otherwise.

### GetNeedResetOk

`func (o *McpServerShortDto) GetNeedResetOk() (*bool, bool)`

GetNeedResetOk returns a tuple with the NeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedReset

`func (o *McpServerShortDto) SetNeedReset(v bool)`

SetNeedReset sets NeedReset field to given value.

### HasNeedReset

`func (o *McpServerShortDto) HasNeedReset() bool`

HasNeedReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


