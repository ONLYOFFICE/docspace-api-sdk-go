# AiWebSearchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Provider identifier (e.g. `exa`). | 
**Key** | Pointer to **string** | API key for the provider. Optional for self-hosted or keyless setups. | [optional] 
**BaseUrl** | Pointer to **string** | Optional override for the provider's base URL. | [optional] 
**IsCloudProvider** | Pointer to **bool** | Whether this provider is cloud-hosted (vs. self-hosted). | [optional] 
**Headers** | Pointer to **map[string]string** | Extra HTTP headers sent with each request to the ONLYOFFICE / cloud backend (e.g. `X-Tenant`). Merged after the derived `Authorization` header, so a custom header of the same name wins. | [optional] 

## Methods

### NewAiWebSearchConfig

`func NewAiWebSearchConfig(provider string, ) *AiWebSearchConfig`

NewAiWebSearchConfig instantiates a new AiWebSearchConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiWebSearchConfigWithDefaults

`func NewAiWebSearchConfigWithDefaults() *AiWebSearchConfig`

NewAiWebSearchConfigWithDefaults instantiates a new AiWebSearchConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AiWebSearchConfig) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiWebSearchConfig) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiWebSearchConfig) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetKey

`func (o *AiWebSearchConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AiWebSearchConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AiWebSearchConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AiWebSearchConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetBaseUrl

`func (o *AiWebSearchConfig) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AiWebSearchConfig) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AiWebSearchConfig) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *AiWebSearchConfig) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetIsCloudProvider

`func (o *AiWebSearchConfig) GetIsCloudProvider() bool`

GetIsCloudProvider returns the IsCloudProvider field if non-nil, zero value otherwise.

### GetIsCloudProviderOk

`func (o *AiWebSearchConfig) GetIsCloudProviderOk() (*bool, bool)`

GetIsCloudProviderOk returns a tuple with the IsCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudProvider

`func (o *AiWebSearchConfig) SetIsCloudProvider(v bool)`

SetIsCloudProvider sets IsCloudProvider field to given value.

### HasIsCloudProvider

`func (o *AiWebSearchConfig) HasIsCloudProvider() bool`

HasIsCloudProvider returns a boolean if a field has been set.

### GetHeaders

`func (o *AiWebSearchConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AiWebSearchConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AiWebSearchConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AiWebSearchConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


