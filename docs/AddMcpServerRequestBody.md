# AddMcpServerRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Unique display name for the server. Only letters, numbers, underscores, and hyphens are allowed. Maximum 128 characters. | 
**Description** | **NullableString** | Human-readable description of the server's purpose and capabilities. Maximum 255 characters. | 
**Endpoint** | **NullableString** | Base URL of the MCP server endpoint. Must be a valid, reachable URL. The system will verify connectivity during registration. | 
**Headers** | Pointer to **map[string]string** | Optional HTTP headers to include with every request to the MCP server (e.g., authentication tokens or API keys). | [optional] 
**Icon** | Pointer to **NullableString** | Optional Base64-encoded icon image for the server. Used as the visual identifier in the UI. | [optional] 

## Methods

### NewAddMcpServerRequestBody

`func NewAddMcpServerRequestBody(name NullableString, description NullableString, endpoint NullableString, ) *AddMcpServerRequestBody`

NewAddMcpServerRequestBody instantiates a new AddMcpServerRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMcpServerRequestBodyWithDefaults

`func NewAddMcpServerRequestBodyWithDefaults() *AddMcpServerRequestBody`

NewAddMcpServerRequestBodyWithDefaults instantiates a new AddMcpServerRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddMcpServerRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddMcpServerRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddMcpServerRequestBody) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AddMcpServerRequestBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddMcpServerRequestBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AddMcpServerRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddMcpServerRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddMcpServerRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AddMcpServerRequestBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddMcpServerRequestBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndpoint

`func (o *AddMcpServerRequestBody) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddMcpServerRequestBody) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddMcpServerRequestBody) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *AddMcpServerRequestBody) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *AddMcpServerRequestBody) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetHeaders

`func (o *AddMcpServerRequestBody) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AddMcpServerRequestBody) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AddMcpServerRequestBody) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AddMcpServerRequestBody) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *AddMcpServerRequestBody) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *AddMcpServerRequestBody) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetIcon

`func (o *AddMcpServerRequestBody) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AddMcpServerRequestBody) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AddMcpServerRequestBody) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AddMcpServerRequestBody) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *AddMcpServerRequestBody) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *AddMcpServerRequestBody) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


