# UpdateClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the client | [optional] 
**Description** | Pointer to **string** | The description of the client | [optional] 
**Logo** | Pointer to **string** | The logo of the client in base64 format | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**AllowPkce** | Pointer to **bool** | Indicates whether PKCE is allowed for the client | [optional] 
**IsPublic** | Pointer to **bool** | Indicates whether client is accessible by third-party tenants | [optional] 
**AllowedOrigins** | Pointer to **[]string** | The allowed origins for the client | [optional] 

## Methods

### NewUpdateClientRequest

`func NewUpdateClientRequest() *UpdateClientRequest`

NewUpdateClientRequest instantiates a new UpdateClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClientRequestWithDefaults

`func NewUpdateClientRequestWithDefaults() *UpdateClientRequest`

NewUpdateClientRequestWithDefaults instantiates a new UpdateClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClientRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClientRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateClientRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateClientRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateClientRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateClientRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogo

`func (o *UpdateClientRequest) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *UpdateClientRequest) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *UpdateClientRequest) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *UpdateClientRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPublic

`func (o *UpdateClientRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *UpdateClientRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *UpdateClientRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *UpdateClientRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetAllowPkce

`func (o *UpdateClientRequest) GetAllowPkce() bool`

GetAllowPkce returns the AllowPkce field if non-nil, zero value otherwise.

### GetAllowPkceOk

`func (o *UpdateClientRequest) GetAllowPkceOk() (*bool, bool)`

GetAllowPkceOk returns a tuple with the AllowPkce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPkce

`func (o *UpdateClientRequest) SetAllowPkce(v bool)`

SetAllowPkce sets AllowPkce field to given value.

### HasAllowPkce

`func (o *UpdateClientRequest) HasAllowPkce() bool`

HasAllowPkce returns a boolean if a field has been set.

### GetIsPublic

`func (o *UpdateClientRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *UpdateClientRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *UpdateClientRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *UpdateClientRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *UpdateClientRequest) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *UpdateClientRequest) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *UpdateClientRequest) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *UpdateClientRequest) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


