# AiProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique profile identifier (UUID). | 
**Name** | **string** | User-defined profile display name. | 
**ProviderType** | [**AiProviderType**](AiProviderType.md) | Provider type for this profile. Use `external` to delegate all HTTP transport to `PlatformAdapter.externalFetch` while reusing an existing provider's response parser — see `Profile.basedOn` for the format selector. | 
**BasedOn** | Pointer to [**AiBuiltinProviderType**](AiBuiltinProviderType.md) | Selects the response-format parser used by the `external` provider. Ignored for any other `providerType`.  Supported values are `openai`, `anthropic`, `mistral` and `openrouter`. Remaining values (`genai`, `stabilityai`, …) are accepted by the type but not yet implemented; passing one raises an error at request time. | [optional] 
**BaseUrl** | **string** | Base URL of the provider API. | 
**Key** | Pointer to **string** | API key or token. Optional for local providers. | [optional] 
**Headers** | Pointer to **map[string]string** | Extra HTTP headers sent with every request to this provider. Merged into the SDK client's default headers; an explicit `Authorization` here wins over the one derived from `key`. Honoured by the OpenAI-family providers. | [optional] 
**ModelId** | **string** | Selected model ID within this provider. | 
**Reasoning** | Pointer to **bool** | Whether extended thinking is enabled for this profile's model. | [optional] 
**Capabilities** | Pointer to **float32** | Bitmask of capabilities supported by the selected model. | [optional] 
**CanUseTool** | Pointer to **bool** | Result of the live tool-capability probe performed at create time and on changes to `modelId` / `providerType` / `baseUrl`. `undefined` means the probe has never run for this profile (legacy record). | [optional] 
**UseResponsesApi** | Pointer to **bool** | Result of the live Responses-API probe (parallel to `canUseTool`). `true` means the model speaks `/v1/responses` and the OpenAI provider must route through `client.responses.create` — required for gpt-5+ reasoning models that reject `reasoning_effort` together with `tools` on `/v1/chat/completions`. Probed at create time and whenever `modelId` / `providerType` / `baseUrl` change. `undefined` means the probe never ran (legacy record) — readers treat that as `false`. | [optional] 
**IsCloudProvider** | Pointer to **bool** | Whether this profile uses a cloud-hosted provider (e.g. ONLYOFFICE DocSpace). | [optional] 
**UseProxy** | Pointer to **bool** | Route every provider request through the host's `fetchProxy` instead of the global `fetch`. Useful when the host runs the widget in a sandbox without direct network access (CORS, custom auth, etc.). Has no effect when the `PlatformAdapter.fetchProxy` is not configured. | [optional] 
**CreatedAt** | Pointer to **float32** | Creation timestamp (ms since epoch). Used to sort the AI models list newest-first. | [optional] 

## Methods

### NewAiProfile

`func NewAiProfile(id string, name string, providerType AiProviderType, baseUrl string, modelId string, ) *AiProfile`

NewAiProfile instantiates a new AiProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProfileWithDefaults

`func NewAiProfileWithDefaults() *AiProfile`

NewAiProfileWithDefaults instantiates a new AiProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiProfile) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AiProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiProfile) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *AiProfile) GetProviderType() AiProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AiProfile) GetProviderTypeOk() (*AiProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AiProfile) SetProviderType(v AiProviderType)`

SetProviderType sets ProviderType field to given value.


### GetBasedOn

`func (o *AiProfile) GetBasedOn() AiBuiltinProviderType`

GetBasedOn returns the BasedOn field if non-nil, zero value otherwise.

### GetBasedOnOk

`func (o *AiProfile) GetBasedOnOk() (*AiBuiltinProviderType, bool)`

GetBasedOnOk returns a tuple with the BasedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasedOn

`func (o *AiProfile) SetBasedOn(v AiBuiltinProviderType)`

SetBasedOn sets BasedOn field to given value.

### HasBasedOn

`func (o *AiProfile) HasBasedOn() bool`

HasBasedOn returns a boolean if a field has been set.

### GetBaseUrl

`func (o *AiProfile) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AiProfile) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AiProfile) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetKey

`func (o *AiProfile) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AiProfile) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AiProfile) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AiProfile) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetHeaders

`func (o *AiProfile) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AiProfile) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AiProfile) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AiProfile) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetModelId

`func (o *AiProfile) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AiProfile) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AiProfile) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetReasoning

`func (o *AiProfile) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *AiProfile) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *AiProfile) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *AiProfile) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetCapabilities

`func (o *AiProfile) GetCapabilities() float32`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AiProfile) GetCapabilitiesOk() (*float32, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AiProfile) SetCapabilities(v float32)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AiProfile) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCanUseTool

`func (o *AiProfile) GetCanUseTool() bool`

GetCanUseTool returns the CanUseTool field if non-nil, zero value otherwise.

### GetCanUseToolOk

`func (o *AiProfile) GetCanUseToolOk() (*bool, bool)`

GetCanUseToolOk returns a tuple with the CanUseTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUseTool

`func (o *AiProfile) SetCanUseTool(v bool)`

SetCanUseTool sets CanUseTool field to given value.

### HasCanUseTool

`func (o *AiProfile) HasCanUseTool() bool`

HasCanUseTool returns a boolean if a field has been set.

### GetUseResponsesApi

`func (o *AiProfile) GetUseResponsesApi() bool`

GetUseResponsesApi returns the UseResponsesApi field if non-nil, zero value otherwise.

### GetUseResponsesApiOk

`func (o *AiProfile) GetUseResponsesApiOk() (*bool, bool)`

GetUseResponsesApiOk returns a tuple with the UseResponsesApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResponsesApi

`func (o *AiProfile) SetUseResponsesApi(v bool)`

SetUseResponsesApi sets UseResponsesApi field to given value.

### HasUseResponsesApi

`func (o *AiProfile) HasUseResponsesApi() bool`

HasUseResponsesApi returns a boolean if a field has been set.

### GetIsCloudProvider

`func (o *AiProfile) GetIsCloudProvider() bool`

GetIsCloudProvider returns the IsCloudProvider field if non-nil, zero value otherwise.

### GetIsCloudProviderOk

`func (o *AiProfile) GetIsCloudProviderOk() (*bool, bool)`

GetIsCloudProviderOk returns a tuple with the IsCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudProvider

`func (o *AiProfile) SetIsCloudProvider(v bool)`

SetIsCloudProvider sets IsCloudProvider field to given value.

### HasIsCloudProvider

`func (o *AiProfile) HasIsCloudProvider() bool`

HasIsCloudProvider returns a boolean if a field has been set.

### GetUseProxy

`func (o *AiProfile) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *AiProfile) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *AiProfile) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *AiProfile) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AiProfile) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiProfile) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiProfile) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AiProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


