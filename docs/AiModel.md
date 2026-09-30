# AiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Model identifier as used by the provider API (e.g. `gpt-4o`, `claude-sonnet-4-20250514`). | 
**Name** | **string** | Human-readable model name for display in the UI. | 
**Provider** | [**AiProviderType**](AiProviderType.md) | Provider that offers this model. | 
**Reasoning** | Pointer to **bool** | Whether this model supports extended thinking / chain-of-thought reasoning. | [optional] 
**ReasoningSupport** | Pointer to [**AiReasoningSupport**](AiReasoningSupport.md) | What the model can do with extended thinking, when the provider's catalogue says so (OpenRouter and the ONLYOFFICE route report a per-model `reasoning` object). Copied onto the profile at save time; absent, the widget falls back to the provider's id-based table. | [optional] 
**Capabilities** | Pointer to **float32** | Bitmask of model capabilities (Chat, Image, Vision, Tools, etc.). Used to filter models per `ActionType`. | [optional] 

## Methods

### NewAiModel

`func NewAiModel(id string, name string, provider AiProviderType, ) *AiModel`

NewAiModel instantiates a new AiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelWithDefaults

`func NewAiModelWithDefaults() *AiModel`

NewAiModelWithDefaults instantiates a new AiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiModel) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *AiModel) GetProvider() AiProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiModel) GetProviderOk() (*AiProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiModel) SetProvider(v AiProviderType)`

SetProvider sets Provider field to given value.


### GetReasoning

`func (o *AiModel) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *AiModel) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *AiModel) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *AiModel) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetReasoningSupport

`func (o *AiModel) GetReasoningSupport() AiReasoningSupport`

GetReasoningSupport returns the ReasoningSupport field if non-nil, zero value otherwise.

### GetReasoningSupportOk

`func (o *AiModel) GetReasoningSupportOk() (*AiReasoningSupport, bool)`

GetReasoningSupportOk returns a tuple with the ReasoningSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningSupport

`func (o *AiModel) SetReasoningSupport(v AiReasoningSupport)`

SetReasoningSupport sets ReasoningSupport field to given value.

### HasReasoningSupport

`func (o *AiModel) HasReasoningSupport() bool`

HasReasoningSupport returns a boolean if a field has been set.

### GetCapabilities

`func (o *AiModel) GetCapabilities() float32`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AiModel) GetCapabilitiesOk() (*float32, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AiModel) SetCapabilities(v float32)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AiModel) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


