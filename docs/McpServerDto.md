# McpServerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the MCP server. | [optional] 
**Name** | Pointer to **NullableString** | Display name of the MCP server. | [optional] 
**Description** | Pointer to **NullableString** | Human-readable description of the server's purpose and capabilities. | [optional] 
**Endpoint** | Pointer to **NullableString** | Base URL of the MCP server endpoint. | [optional] 
**ServerType** | Pointer to [**ServerType**](ServerType.md) |  | [optional] 
**Headers** | Pointer to **map[string]string** | HTTP headers sent with every request to the server (e.g., authentication tokens). | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the server is currently enabled and available for room assignment. | [optional] 
**Icon** | Pointer to [**Icon**](Icon.md) |  | [optional] 
**NeedReset** | Pointer to **bool** | Indicates whether the server requires a configuration reset due to connectivity or credential issues. | [optional] 

## Methods

### NewMcpServerDto

`func NewMcpServerDto() *McpServerDto`

NewMcpServerDto instantiates a new McpServerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerDtoWithDefaults

`func NewMcpServerDtoWithDefaults() *McpServerDto`

NewMcpServerDtoWithDefaults instantiates a new McpServerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McpServerDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McpServerDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McpServerDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *McpServerDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *McpServerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *McpServerDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *McpServerDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *McpServerDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *McpServerDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *McpServerDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *McpServerDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *McpServerDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *McpServerDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *McpServerDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndpoint

`func (o *McpServerDto) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *McpServerDto) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *McpServerDto) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *McpServerDto) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *McpServerDto) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *McpServerDto) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetServerType

`func (o *McpServerDto) GetServerType() ServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *McpServerDto) GetServerTypeOk() (*ServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *McpServerDto) SetServerType(v ServerType)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *McpServerDto) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetHeaders

`func (o *McpServerDto) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *McpServerDto) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *McpServerDto) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *McpServerDto) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *McpServerDto) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *McpServerDto) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetEnabled

`func (o *McpServerDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *McpServerDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *McpServerDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *McpServerDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIcon

`func (o *McpServerDto) GetIcon() Icon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *McpServerDto) GetIconOk() (*Icon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *McpServerDto) SetIcon(v Icon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *McpServerDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetNeedReset

`func (o *McpServerDto) GetNeedReset() bool`

GetNeedReset returns the NeedReset field if non-nil, zero value otherwise.

### GetNeedResetOk

`func (o *McpServerDto) GetNeedResetOk() (*bool, bool)`

GetNeedResetOk returns a tuple with the NeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedReset

`func (o *McpServerDto) SetNeedReset(v bool)`

SetNeedReset sets NeedReset field to given value.

### HasNeedReset

`func (o *McpServerDto) HasNeedReset() bool`

HasNeedReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


