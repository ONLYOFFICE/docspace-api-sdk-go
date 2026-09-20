# AiProfilesGetById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique profile identifier (UUID). | 
**Name** | **string** | User-defined profile display name. | 
**ProviderType** | [**AiProviderType**](AiProviderType.md) | Provider type for this profile. Use `external` to delegate all HTTP transport to `PlatformAdapter.externalFetch` while reusing an existing provider's response parser — see `Profile.basedOn` for the format selector. | 
**BasedOn** | Pointer to [**AiBuiltinProviderType**](AiBuiltinProviderType.md) | Selects the response-format parser used by the `external` provider. Ignored for any other `providerType`.  Supported values are `openai`, `anthropic`, `mistral` and `openrouter`. Remaining values (`genai`, `stabilityai`, …) are accepted by the type but not yet implemented; passing one raises an error at request time. | [optional] 
**BaseUrl** | **string** | Base URL of the provider API. | 
**ModelId** | **string** | Selected model ID within this provider. | 
**Reasoning** | Pointer to **bool** | Whether extended thinking is enabled for this profile's model. | [optional] 
**ReasoningSupport** | Pointer to [**AiReasoningSupport**](AiReasoningSupport.md) | Extended-thinking capabilities of the selected model as reported by the provider's catalogue at save time (see `Model.reasoningSupport`). When present the composer's Effort row follows it exactly; when absent the provider's id-based table answers. Hosts persist it with the rest of the profile. | [optional] 
**Capabilities** | Pointer to **float32** | Bitmask of capabilities supported by the selected model. | [optional] 
**CanUseTool** | Pointer to **bool** | Result of the live tool-capability probe performed at create time and on changes to `modelId` / `providerType` / `baseUrl`. `undefined` means the probe has never run for this profile (legacy record). | [optional] 
**UseResponsesApi** | Pointer to **bool** | Result of the live Responses-API probe (parallel to `canUseTool`). `true` means the model speaks `/v1/responses` and the OpenAI provider must route through `client.responses.create` — required for gpt-5+ reasoning models that reject `reasoning_effort` together with `tools` on `/v1/chat/completions`. Probed at create time and whenever `modelId` / `providerType` / `baseUrl` change. `undefined` means the probe never ran (legacy record) — readers treat that as `false`. | [optional] 
**IsCloudProvider** | Pointer to **bool** | Whether this profile uses a cloud-hosted provider (e.g. ONLYOFFICE DocSpace). | [optional] 
**UseProxy** | Pointer to **bool** | Route every provider request through the host's `fetchProxy` instead of the global `fetch`. Useful when the host runs the widget in a sandbox without direct network access (CORS, custom auth, etc.). Has no effect when the `PlatformAdapter.fetchProxy` is not configured. | [optional] 
**CreatedAt** | Pointer to **float32** | Creation timestamp (ms since epoch). Used to sort the AI models list newest-first. | [optional] 

## Methods

### NewAiProfilesGetById200Response

`func NewAiProfilesGetById200Response(id string, name string, providerType AiProviderType, baseUrl string, modelId string, ) *AiProfilesGetById200Response`

NewAiProfilesGetById200Response instantiates a new AiProfilesGetById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProfilesGetById200ResponseWithDefaults

`func NewAiProfilesGetById200ResponseWithDefaults() *AiProfilesGetById200Response`

NewAiProfilesGetById200ResponseWithDefaults instantiates a new AiProfilesGetById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiProfilesGetById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiProfilesGetById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiProfilesGetById200Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AiProfilesGetById200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiProfilesGetById200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiProfilesGetById200Response) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *AiProfilesGetById200Response) GetProviderType() AiProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AiProfilesGetById200Response) GetProviderTypeOk() (*AiProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AiProfilesGetById200Response) SetProviderType(v AiProviderType)`

SetProviderType sets ProviderType field to given value.


### GetBasedOn

`func (o *AiProfilesGetById200Response) GetBasedOn() AiBuiltinProviderType`

GetBasedOn returns the BasedOn field if non-nil, zero value otherwise.

### GetBasedOnOk

`func (o *AiProfilesGetById200Response) GetBasedOnOk() (*AiBuiltinProviderType, bool)`

GetBasedOnOk returns a tuple with the BasedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasedOn

`func (o *AiProfilesGetById200Response) SetBasedOn(v AiBuiltinProviderType)`

SetBasedOn sets BasedOn field to given value.

### HasBasedOn

`func (o *AiProfilesGetById200Response) HasBasedOn() bool`

HasBasedOn returns a boolean if a field has been set.

### GetBaseUrl

`func (o *AiProfilesGetById200Response) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AiProfilesGetById200Response) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AiProfilesGetById200Response) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetModelId

`func (o *AiProfilesGetById200Response) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AiProfilesGetById200Response) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AiProfilesGetById200Response) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetReasoning

`func (o *AiProfilesGetById200Response) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *AiProfilesGetById200Response) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *AiProfilesGetById200Response) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *AiProfilesGetById200Response) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetReasoningSupport

`func (o *AiProfilesGetById200Response) GetReasoningSupport() AiReasoningSupport`

GetReasoningSupport returns the ReasoningSupport field if non-nil, zero value otherwise.

### GetReasoningSupportOk

`func (o *AiProfilesGetById200Response) GetReasoningSupportOk() (*AiReasoningSupport, bool)`

GetReasoningSupportOk returns a tuple with the ReasoningSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningSupport

`func (o *AiProfilesGetById200Response) SetReasoningSupport(v AiReasoningSupport)`

SetReasoningSupport sets ReasoningSupport field to given value.

### HasReasoningSupport

`func (o *AiProfilesGetById200Response) HasReasoningSupport() bool`

HasReasoningSupport returns a boolean if a field has been set.

### GetCapabilities

`func (o *AiProfilesGetById200Response) GetCapabilities() float32`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AiProfilesGetById200Response) GetCapabilitiesOk() (*float32, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AiProfilesGetById200Response) SetCapabilities(v float32)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AiProfilesGetById200Response) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCanUseTool

`func (o *AiProfilesGetById200Response) GetCanUseTool() bool`

GetCanUseTool returns the CanUseTool field if non-nil, zero value otherwise.

### GetCanUseToolOk

`func (o *AiProfilesGetById200Response) GetCanUseToolOk() (*bool, bool)`

GetCanUseToolOk returns a tuple with the CanUseTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUseTool

`func (o *AiProfilesGetById200Response) SetCanUseTool(v bool)`

SetCanUseTool sets CanUseTool field to given value.

### HasCanUseTool

`func (o *AiProfilesGetById200Response) HasCanUseTool() bool`

HasCanUseTool returns a boolean if a field has been set.

### GetUseResponsesApi

`func (o *AiProfilesGetById200Response) GetUseResponsesApi() bool`

GetUseResponsesApi returns the UseResponsesApi field if non-nil, zero value otherwise.

### GetUseResponsesApiOk

`func (o *AiProfilesGetById200Response) GetUseResponsesApiOk() (*bool, bool)`

GetUseResponsesApiOk returns a tuple with the UseResponsesApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResponsesApi

`func (o *AiProfilesGetById200Response) SetUseResponsesApi(v bool)`

SetUseResponsesApi sets UseResponsesApi field to given value.

### HasUseResponsesApi

`func (o *AiProfilesGetById200Response) HasUseResponsesApi() bool`

HasUseResponsesApi returns a boolean if a field has been set.

### GetIsCloudProvider

`func (o *AiProfilesGetById200Response) GetIsCloudProvider() bool`

GetIsCloudProvider returns the IsCloudProvider field if non-nil, zero value otherwise.

### GetIsCloudProviderOk

`func (o *AiProfilesGetById200Response) GetIsCloudProviderOk() (*bool, bool)`

GetIsCloudProviderOk returns a tuple with the IsCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudProvider

`func (o *AiProfilesGetById200Response) SetIsCloudProvider(v bool)`

SetIsCloudProvider sets IsCloudProvider field to given value.

### HasIsCloudProvider

`func (o *AiProfilesGetById200Response) HasIsCloudProvider() bool`

HasIsCloudProvider returns a boolean if a field has been set.

### GetUseProxy

`func (o *AiProfilesGetById200Response) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *AiProfilesGetById200Response) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *AiProfilesGetById200Response) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *AiProfilesGetById200Response) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AiProfilesGetById200Response) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiProfilesGetById200Response) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiProfilesGetById200Response) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AiProfilesGetById200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


