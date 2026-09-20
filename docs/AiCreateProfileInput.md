# AiCreateProfileInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User-defined profile display name. | 
**ProviderType** | [**AiProviderType**](AiProviderType.md) | Provider type for this profile. Use `external` to delegate all HTTP transport to `PlatformAdapter.externalFetch` while reusing an existing provider's response parser — see `Profile.basedOn` for the format selector. | 
**BasedOn** | Pointer to [**AiBuiltinProviderType**](AiBuiltinProviderType.md) | Selects the response-format parser used by the `external` provider. Ignored for any other `providerType`.  Supported values are `openai`, `anthropic`, `mistral` and `openrouter`. Remaining values (`genai`, `stabilityai`, …) are accepted by the type but not yet implemented; passing one raises an error at request time. | [optional] 
**BaseUrl** | **string** | Base URL of the provider API. | 
**Key** | Pointer to **string** | API key or token. Optional for local providers. | [optional] 
**Headers** | Pointer to **map[string]string** | Extra HTTP headers sent with every request to this provider. Merged into the SDK client's default headers; an explicit `Authorization` here wins over the one derived from `key`. Honoured by the OpenAI-family providers. | [optional] 
**ModelId** | **string** | Selected model ID within this provider. | 
**Reasoning** | Pointer to **bool** | Whether extended thinking is enabled for this profile's model. | [optional] 
**ReasoningSupport** | Pointer to [**AiReasoningSupport**](AiReasoningSupport.md) | Extended-thinking capabilities of the selected model as reported by the provider's catalogue at save time (see `Model.reasoningSupport`). When present the composer's Effort row follows it exactly; when absent the provider's id-based table answers. Hosts persist it with the rest of the profile. | [optional] 
**Capabilities** | Pointer to **float32** | Bitmask of capabilities supported by the selected model. | [optional] 
**CanUseTool** | Pointer to **bool** | Result of the live tool-capability probe performed at create time and on changes to `modelId` / `providerType` / `baseUrl`. `undefined` means the probe has never run for this profile (legacy record). | [optional] 
**UseResponsesApi** | Pointer to **bool** | Result of the live Responses-API probe (parallel to `canUseTool`). `true` means the model speaks `/v1/responses` and the OpenAI provider must route through `client.responses.create` — required for gpt-5+ reasoning models that reject `reasoning_effort` together with `tools` on `/v1/chat/completions`. Probed at create time and whenever `modelId` / `providerType` / `baseUrl` change. `undefined` means the probe never ran (legacy record) — readers treat that as `false`. | [optional] 
**IsCloudProvider** | Pointer to **bool** | Whether this profile uses a cloud-hosted provider (e.g. ONLYOFFICE DocSpace). | [optional] 
**UseProxy** | Pointer to **bool** | Route every provider request through the host's `fetchProxy` instead of the global `fetch`. Useful when the host runs the widget in a sandbox without direct network access (CORS, custom auth, etc.). Has no effect when the `PlatformAdapter.fetchProxy` is not configured. | [optional] 

## Methods

### NewAiCreateProfileInput

`func NewAiCreateProfileInput(name string, providerType AiProviderType, baseUrl string, modelId string, ) *AiCreateProfileInput`

NewAiCreateProfileInput instantiates a new AiCreateProfileInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiCreateProfileInputWithDefaults

`func NewAiCreateProfileInputWithDefaults() *AiCreateProfileInput`

NewAiCreateProfileInputWithDefaults instantiates a new AiCreateProfileInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiCreateProfileInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiCreateProfileInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiCreateProfileInput) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *AiCreateProfileInput) GetProviderType() AiProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AiCreateProfileInput) GetProviderTypeOk() (*AiProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AiCreateProfileInput) SetProviderType(v AiProviderType)`

SetProviderType sets ProviderType field to given value.


### GetBasedOn

`func (o *AiCreateProfileInput) GetBasedOn() AiBuiltinProviderType`

GetBasedOn returns the BasedOn field if non-nil, zero value otherwise.

### GetBasedOnOk

`func (o *AiCreateProfileInput) GetBasedOnOk() (*AiBuiltinProviderType, bool)`

GetBasedOnOk returns a tuple with the BasedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasedOn

`func (o *AiCreateProfileInput) SetBasedOn(v AiBuiltinProviderType)`

SetBasedOn sets BasedOn field to given value.

### HasBasedOn

`func (o *AiCreateProfileInput) HasBasedOn() bool`

HasBasedOn returns a boolean if a field has been set.

### GetBaseUrl

`func (o *AiCreateProfileInput) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AiCreateProfileInput) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AiCreateProfileInput) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetKey

`func (o *AiCreateProfileInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AiCreateProfileInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AiCreateProfileInput) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AiCreateProfileInput) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetHeaders

`func (o *AiCreateProfileInput) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AiCreateProfileInput) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AiCreateProfileInput) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AiCreateProfileInput) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetModelId

`func (o *AiCreateProfileInput) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AiCreateProfileInput) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AiCreateProfileInput) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetReasoning

`func (o *AiCreateProfileInput) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *AiCreateProfileInput) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *AiCreateProfileInput) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *AiCreateProfileInput) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetReasoningSupport

`func (o *AiCreateProfileInput) GetReasoningSupport() AiReasoningSupport`

GetReasoningSupport returns the ReasoningSupport field if non-nil, zero value otherwise.

### GetReasoningSupportOk

`func (o *AiCreateProfileInput) GetReasoningSupportOk() (*AiReasoningSupport, bool)`

GetReasoningSupportOk returns a tuple with the ReasoningSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningSupport

`func (o *AiCreateProfileInput) SetReasoningSupport(v AiReasoningSupport)`

SetReasoningSupport sets ReasoningSupport field to given value.

### HasReasoningSupport

`func (o *AiCreateProfileInput) HasReasoningSupport() bool`

HasReasoningSupport returns a boolean if a field has been set.

### GetCapabilities

`func (o *AiCreateProfileInput) GetCapabilities() float32`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AiCreateProfileInput) GetCapabilitiesOk() (*float32, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AiCreateProfileInput) SetCapabilities(v float32)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AiCreateProfileInput) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCanUseTool

`func (o *AiCreateProfileInput) GetCanUseTool() bool`

GetCanUseTool returns the CanUseTool field if non-nil, zero value otherwise.

### GetCanUseToolOk

`func (o *AiCreateProfileInput) GetCanUseToolOk() (*bool, bool)`

GetCanUseToolOk returns a tuple with the CanUseTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUseTool

`func (o *AiCreateProfileInput) SetCanUseTool(v bool)`

SetCanUseTool sets CanUseTool field to given value.

### HasCanUseTool

`func (o *AiCreateProfileInput) HasCanUseTool() bool`

HasCanUseTool returns a boolean if a field has been set.

### GetUseResponsesApi

`func (o *AiCreateProfileInput) GetUseResponsesApi() bool`

GetUseResponsesApi returns the UseResponsesApi field if non-nil, zero value otherwise.

### GetUseResponsesApiOk

`func (o *AiCreateProfileInput) GetUseResponsesApiOk() (*bool, bool)`

GetUseResponsesApiOk returns a tuple with the UseResponsesApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResponsesApi

`func (o *AiCreateProfileInput) SetUseResponsesApi(v bool)`

SetUseResponsesApi sets UseResponsesApi field to given value.

### HasUseResponsesApi

`func (o *AiCreateProfileInput) HasUseResponsesApi() bool`

HasUseResponsesApi returns a boolean if a field has been set.

### GetIsCloudProvider

`func (o *AiCreateProfileInput) GetIsCloudProvider() bool`

GetIsCloudProvider returns the IsCloudProvider field if non-nil, zero value otherwise.

### GetIsCloudProviderOk

`func (o *AiCreateProfileInput) GetIsCloudProviderOk() (*bool, bool)`

GetIsCloudProviderOk returns a tuple with the IsCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudProvider

`func (o *AiCreateProfileInput) SetIsCloudProvider(v bool)`

SetIsCloudProvider sets IsCloudProvider field to given value.

### HasIsCloudProvider

`func (o *AiCreateProfileInput) HasIsCloudProvider() bool`

HasIsCloudProvider returns a boolean if a field has been set.

### GetUseProxy

`func (o *AiCreateProfileInput) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *AiCreateProfileInput) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *AiCreateProfileInput) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *AiCreateProfileInput) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


