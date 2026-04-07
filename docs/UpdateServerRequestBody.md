# UpdateServerRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | New display name for the server. Only letters, numbers, underscores, and hyphens are allowed. Maximum 128 characters. | [optional] 
**Description** | Pointer to **NullableString** | New human-readable description of the server's purpose. Maximum 255 characters. | [optional] 
**Endpoint** | Pointer to **NullableString** | New base URL of the MCP server endpoint. If changed, the system will re-verify connectivity before saving. | [optional] 
**Headers** | Pointer to **map[string]string** | New HTTP headers to include with every request. If changed alongside the endpoint, connectivity is re-verified. | [optional] 
**UpdateIcon** | Pointer to **bool** | Set to true to update the server icon. When true, the Icon field value (or null to remove) will be applied. | [optional] 
**Icon** | Pointer to **NullableString** | New Base64-encoded icon image for the server, or null to remove the existing icon. Only applied when UpdateIcon is true. | [optional] 

## Methods

### NewUpdateServerRequestBody

`func NewUpdateServerRequestBody() *UpdateServerRequestBody`

NewUpdateServerRequestBody instantiates a new UpdateServerRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerRequestBodyWithDefaults

`func NewUpdateServerRequestBodyWithDefaults() *UpdateServerRequestBody`

NewUpdateServerRequestBodyWithDefaults instantiates a new UpdateServerRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateServerRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServerRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServerRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServerRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateServerRequestBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateServerRequestBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateServerRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServerRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServerRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServerRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateServerRequestBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateServerRequestBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndpoint

`func (o *UpdateServerRequestBody) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UpdateServerRequestBody) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UpdateServerRequestBody) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UpdateServerRequestBody) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *UpdateServerRequestBody) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *UpdateServerRequestBody) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetHeaders

`func (o *UpdateServerRequestBody) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *UpdateServerRequestBody) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *UpdateServerRequestBody) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *UpdateServerRequestBody) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *UpdateServerRequestBody) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *UpdateServerRequestBody) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetUpdateIcon

`func (o *UpdateServerRequestBody) GetUpdateIcon() bool`

GetUpdateIcon returns the UpdateIcon field if non-nil, zero value otherwise.

### GetUpdateIconOk

`func (o *UpdateServerRequestBody) GetUpdateIconOk() (*bool, bool)`

GetUpdateIconOk returns a tuple with the UpdateIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateIcon

`func (o *UpdateServerRequestBody) SetUpdateIcon(v bool)`

SetUpdateIcon sets UpdateIcon field to given value.

### HasUpdateIcon

`func (o *UpdateServerRequestBody) HasUpdateIcon() bool`

HasUpdateIcon returns a boolean if a field has been set.

### GetIcon

`func (o *UpdateServerRequestBody) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UpdateServerRequestBody) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UpdateServerRequestBody) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *UpdateServerRequestBody) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *UpdateServerRequestBody) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *UpdateServerRequestBody) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


