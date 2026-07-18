# AiSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebSearchEnabled** | Pointer to **bool** | Indicates whether web search is enabled for AI chat sessions. | [optional] 
**WebSearchNeedReset** | Pointer to **bool** | Indicates whether the web search API key needs to be reconfigured. | [optional] 
**VectorizationEnabled** | Pointer to **bool** | Indicates whether document vectorization is enabled. | [optional] 
**VectorizationNeedReset** | Pointer to **bool** | Indicates whether the embedding provider API key needs to be reconfigured. | [optional] 
**AiReady** | Pointer to **bool** | Indicates whether the AI subsystem is fully configured and operational. | [optional] 
**AiReadyNeedReset** | Pointer to **bool** | Indicates whether the AI provider API key needs to be reconfigured. | [optional] 
**PortalMcpServerId** | Pointer to **NullableString** | The unique identifier of the portal-level MCP server, if configured. | [optional] 
**EmbeddingModel** | **NullableString** | The name of the embedding model used for document vectorization. | 
**ModelAliases** | **map[string]string** | Mapping of model identifiers to human-readable aliases. | 
**KnowledgeSearchToolName** | **NullableString** | The tool name used by the AI assistant for knowledge base search. | 
**WebSearchToolName** | **NullableString** | The tool name used by the AI assistant for web search. | 
**WebCrawlingToolName** | **NullableString** | The tool name used by the AI assistant for web page crawling. | 
**GenerateDocxToolName** | **NullableString** | The tool name used by the AI to launch docx creation in the editor. | 
**GenerateFormToolName** | **NullableString** | The tool name used by the AI assistant to launch form creation in the editor. | 
**GeneratePresentationToolName** | **NullableString** | The tool name used by the AI assistant to launch presentation creation in the editor. | 
**SystemAiEnabled** | Pointer to **bool** | Indicates whether the system-level AI provider is enabled. | [optional] 
**RecommendedModelForForms** | Pointer to **NullableString** | The identifier of the model recommended for form generation. | [optional] 

## Methods

### NewAiSettingsDto

`func NewAiSettingsDto(embeddingModel NullableString, modelAliases map[string]string, knowledgeSearchToolName NullableString, webSearchToolName NullableString, webCrawlingToolName NullableString, generateDocxToolName NullableString, generateFormToolName NullableString, generatePresentationToolName NullableString, ) *AiSettingsDto`

NewAiSettingsDto instantiates a new AiSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSettingsDtoWithDefaults

`func NewAiSettingsDtoWithDefaults() *AiSettingsDto`

NewAiSettingsDtoWithDefaults instantiates a new AiSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebSearchEnabled

`func (o *AiSettingsDto) GetWebSearchEnabled() bool`

GetWebSearchEnabled returns the WebSearchEnabled field if non-nil, zero value otherwise.

### GetWebSearchEnabledOk

`func (o *AiSettingsDto) GetWebSearchEnabledOk() (*bool, bool)`

GetWebSearchEnabledOk returns a tuple with the WebSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchEnabled

`func (o *AiSettingsDto) SetWebSearchEnabled(v bool)`

SetWebSearchEnabled sets WebSearchEnabled field to given value.

### HasWebSearchEnabled

`func (o *AiSettingsDto) HasWebSearchEnabled() bool`

HasWebSearchEnabled returns a boolean if a field has been set.

### GetWebSearchNeedReset

`func (o *AiSettingsDto) GetWebSearchNeedReset() bool`

GetWebSearchNeedReset returns the WebSearchNeedReset field if non-nil, zero value otherwise.

### GetWebSearchNeedResetOk

`func (o *AiSettingsDto) GetWebSearchNeedResetOk() (*bool, bool)`

GetWebSearchNeedResetOk returns a tuple with the WebSearchNeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchNeedReset

`func (o *AiSettingsDto) SetWebSearchNeedReset(v bool)`

SetWebSearchNeedReset sets WebSearchNeedReset field to given value.

### HasWebSearchNeedReset

`func (o *AiSettingsDto) HasWebSearchNeedReset() bool`

HasWebSearchNeedReset returns a boolean if a field has been set.

### GetVectorizationEnabled

`func (o *AiSettingsDto) GetVectorizationEnabled() bool`

GetVectorizationEnabled returns the VectorizationEnabled field if non-nil, zero value otherwise.

### GetVectorizationEnabledOk

`func (o *AiSettingsDto) GetVectorizationEnabledOk() (*bool, bool)`

GetVectorizationEnabledOk returns a tuple with the VectorizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorizationEnabled

`func (o *AiSettingsDto) SetVectorizationEnabled(v bool)`

SetVectorizationEnabled sets VectorizationEnabled field to given value.

### HasVectorizationEnabled

`func (o *AiSettingsDto) HasVectorizationEnabled() bool`

HasVectorizationEnabled returns a boolean if a field has been set.

### GetVectorizationNeedReset

`func (o *AiSettingsDto) GetVectorizationNeedReset() bool`

GetVectorizationNeedReset returns the VectorizationNeedReset field if non-nil, zero value otherwise.

### GetVectorizationNeedResetOk

`func (o *AiSettingsDto) GetVectorizationNeedResetOk() (*bool, bool)`

GetVectorizationNeedResetOk returns a tuple with the VectorizationNeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorizationNeedReset

`func (o *AiSettingsDto) SetVectorizationNeedReset(v bool)`

SetVectorizationNeedReset sets VectorizationNeedReset field to given value.

### HasVectorizationNeedReset

`func (o *AiSettingsDto) HasVectorizationNeedReset() bool`

HasVectorizationNeedReset returns a boolean if a field has been set.

### GetAiReady

`func (o *AiSettingsDto) GetAiReady() bool`

GetAiReady returns the AiReady field if non-nil, zero value otherwise.

### GetAiReadyOk

`func (o *AiSettingsDto) GetAiReadyOk() (*bool, bool)`

GetAiReadyOk returns a tuple with the AiReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiReady

`func (o *AiSettingsDto) SetAiReady(v bool)`

SetAiReady sets AiReady field to given value.

### HasAiReady

`func (o *AiSettingsDto) HasAiReady() bool`

HasAiReady returns a boolean if a field has been set.

### GetAiReadyNeedReset

`func (o *AiSettingsDto) GetAiReadyNeedReset() bool`

GetAiReadyNeedReset returns the AiReadyNeedReset field if non-nil, zero value otherwise.

### GetAiReadyNeedResetOk

`func (o *AiSettingsDto) GetAiReadyNeedResetOk() (*bool, bool)`

GetAiReadyNeedResetOk returns a tuple with the AiReadyNeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiReadyNeedReset

`func (o *AiSettingsDto) SetAiReadyNeedReset(v bool)`

SetAiReadyNeedReset sets AiReadyNeedReset field to given value.

### HasAiReadyNeedReset

`func (o *AiSettingsDto) HasAiReadyNeedReset() bool`

HasAiReadyNeedReset returns a boolean if a field has been set.

### GetPortalMcpServerId

`func (o *AiSettingsDto) GetPortalMcpServerId() string`

GetPortalMcpServerId returns the PortalMcpServerId field if non-nil, zero value otherwise.

### GetPortalMcpServerIdOk

`func (o *AiSettingsDto) GetPortalMcpServerIdOk() (*string, bool)`

GetPortalMcpServerIdOk returns a tuple with the PortalMcpServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalMcpServerId

`func (o *AiSettingsDto) SetPortalMcpServerId(v string)`

SetPortalMcpServerId sets PortalMcpServerId field to given value.

### HasPortalMcpServerId

`func (o *AiSettingsDto) HasPortalMcpServerId() bool`

HasPortalMcpServerId returns a boolean if a field has been set.

### SetPortalMcpServerIdNil

`func (o *AiSettingsDto) SetPortalMcpServerIdNil(b bool)`

 SetPortalMcpServerIdNil sets the value for PortalMcpServerId to be an explicit nil

### UnsetPortalMcpServerId
`func (o *AiSettingsDto) UnsetPortalMcpServerId()`

UnsetPortalMcpServerId ensures that no value is present for PortalMcpServerId, not even an explicit nil
### GetEmbeddingModel

`func (o *AiSettingsDto) GetEmbeddingModel() string`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *AiSettingsDto) GetEmbeddingModelOk() (*string, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *AiSettingsDto) SetEmbeddingModel(v string)`

SetEmbeddingModel sets EmbeddingModel field to given value.


### SetEmbeddingModelNil

`func (o *AiSettingsDto) SetEmbeddingModelNil(b bool)`

 SetEmbeddingModelNil sets the value for EmbeddingModel to be an explicit nil

### UnsetEmbeddingModel
`func (o *AiSettingsDto) UnsetEmbeddingModel()`

UnsetEmbeddingModel ensures that no value is present for EmbeddingModel, not even an explicit nil
### GetModelAliases

`func (o *AiSettingsDto) GetModelAliases() map[string]string`

GetModelAliases returns the ModelAliases field if non-nil, zero value otherwise.

### GetModelAliasesOk

`func (o *AiSettingsDto) GetModelAliasesOk() (*map[string]string, bool)`

GetModelAliasesOk returns a tuple with the ModelAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelAliases

`func (o *AiSettingsDto) SetModelAliases(v map[string]string)`

SetModelAliases sets ModelAliases field to given value.


### SetModelAliasesNil

`func (o *AiSettingsDto) SetModelAliasesNil(b bool)`

 SetModelAliasesNil sets the value for ModelAliases to be an explicit nil

### UnsetModelAliases
`func (o *AiSettingsDto) UnsetModelAliases()`

UnsetModelAliases ensures that no value is present for ModelAliases, not even an explicit nil
### GetKnowledgeSearchToolName

`func (o *AiSettingsDto) GetKnowledgeSearchToolName() string`

GetKnowledgeSearchToolName returns the KnowledgeSearchToolName field if non-nil, zero value otherwise.

### GetKnowledgeSearchToolNameOk

`func (o *AiSettingsDto) GetKnowledgeSearchToolNameOk() (*string, bool)`

GetKnowledgeSearchToolNameOk returns a tuple with the KnowledgeSearchToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeSearchToolName

`func (o *AiSettingsDto) SetKnowledgeSearchToolName(v string)`

SetKnowledgeSearchToolName sets KnowledgeSearchToolName field to given value.


### SetKnowledgeSearchToolNameNil

`func (o *AiSettingsDto) SetKnowledgeSearchToolNameNil(b bool)`

 SetKnowledgeSearchToolNameNil sets the value for KnowledgeSearchToolName to be an explicit nil

### UnsetKnowledgeSearchToolName
`func (o *AiSettingsDto) UnsetKnowledgeSearchToolName()`

UnsetKnowledgeSearchToolName ensures that no value is present for KnowledgeSearchToolName, not even an explicit nil
### GetWebSearchToolName

`func (o *AiSettingsDto) GetWebSearchToolName() string`

GetWebSearchToolName returns the WebSearchToolName field if non-nil, zero value otherwise.

### GetWebSearchToolNameOk

`func (o *AiSettingsDto) GetWebSearchToolNameOk() (*string, bool)`

GetWebSearchToolNameOk returns a tuple with the WebSearchToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchToolName

`func (o *AiSettingsDto) SetWebSearchToolName(v string)`

SetWebSearchToolName sets WebSearchToolName field to given value.


### SetWebSearchToolNameNil

`func (o *AiSettingsDto) SetWebSearchToolNameNil(b bool)`

 SetWebSearchToolNameNil sets the value for WebSearchToolName to be an explicit nil

### UnsetWebSearchToolName
`func (o *AiSettingsDto) UnsetWebSearchToolName()`

UnsetWebSearchToolName ensures that no value is present for WebSearchToolName, not even an explicit nil
### GetWebCrawlingToolName

`func (o *AiSettingsDto) GetWebCrawlingToolName() string`

GetWebCrawlingToolName returns the WebCrawlingToolName field if non-nil, zero value otherwise.

### GetWebCrawlingToolNameOk

`func (o *AiSettingsDto) GetWebCrawlingToolNameOk() (*string, bool)`

GetWebCrawlingToolNameOk returns a tuple with the WebCrawlingToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCrawlingToolName

`func (o *AiSettingsDto) SetWebCrawlingToolName(v string)`

SetWebCrawlingToolName sets WebCrawlingToolName field to given value.


### SetWebCrawlingToolNameNil

`func (o *AiSettingsDto) SetWebCrawlingToolNameNil(b bool)`

 SetWebCrawlingToolNameNil sets the value for WebCrawlingToolName to be an explicit nil

### UnsetWebCrawlingToolName
`func (o *AiSettingsDto) UnsetWebCrawlingToolName()`

UnsetWebCrawlingToolName ensures that no value is present for WebCrawlingToolName, not even an explicit nil
### GetGenerateDocxToolName

`func (o *AiSettingsDto) GetGenerateDocxToolName() string`

GetGenerateDocxToolName returns the GenerateDocxToolName field if non-nil, zero value otherwise.

### GetGenerateDocxToolNameOk

`func (o *AiSettingsDto) GetGenerateDocxToolNameOk() (*string, bool)`

GetGenerateDocxToolNameOk returns a tuple with the GenerateDocxToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateDocxToolName

`func (o *AiSettingsDto) SetGenerateDocxToolName(v string)`

SetGenerateDocxToolName sets GenerateDocxToolName field to given value.


### SetGenerateDocxToolNameNil

`func (o *AiSettingsDto) SetGenerateDocxToolNameNil(b bool)`

 SetGenerateDocxToolNameNil sets the value for GenerateDocxToolName to be an explicit nil

### UnsetGenerateDocxToolName
`func (o *AiSettingsDto) UnsetGenerateDocxToolName()`

UnsetGenerateDocxToolName ensures that no value is present for GenerateDocxToolName, not even an explicit nil
### GetGenerateFormToolName

`func (o *AiSettingsDto) GetGenerateFormToolName() string`

GetGenerateFormToolName returns the GenerateFormToolName field if non-nil, zero value otherwise.

### GetGenerateFormToolNameOk

`func (o *AiSettingsDto) GetGenerateFormToolNameOk() (*string, bool)`

GetGenerateFormToolNameOk returns a tuple with the GenerateFormToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateFormToolName

`func (o *AiSettingsDto) SetGenerateFormToolName(v string)`

SetGenerateFormToolName sets GenerateFormToolName field to given value.


### SetGenerateFormToolNameNil

`func (o *AiSettingsDto) SetGenerateFormToolNameNil(b bool)`

 SetGenerateFormToolNameNil sets the value for GenerateFormToolName to be an explicit nil

### UnsetGenerateFormToolName
`func (o *AiSettingsDto) UnsetGenerateFormToolName()`

UnsetGenerateFormToolName ensures that no value is present for GenerateFormToolName, not even an explicit nil
### GetGeneratePresentationToolName

`func (o *AiSettingsDto) GetGeneratePresentationToolName() string`

GetGeneratePresentationToolName returns the GeneratePresentationToolName field if non-nil, zero value otherwise.

### GetGeneratePresentationToolNameOk

`func (o *AiSettingsDto) GetGeneratePresentationToolNameOk() (*string, bool)`

GetGeneratePresentationToolNameOk returns a tuple with the GeneratePresentationToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratePresentationToolName

`func (o *AiSettingsDto) SetGeneratePresentationToolName(v string)`

SetGeneratePresentationToolName sets GeneratePresentationToolName field to given value.


### SetGeneratePresentationToolNameNil

`func (o *AiSettingsDto) SetGeneratePresentationToolNameNil(b bool)`

 SetGeneratePresentationToolNameNil sets the value for GeneratePresentationToolName to be an explicit nil

### UnsetGeneratePresentationToolName
`func (o *AiSettingsDto) UnsetGeneratePresentationToolName()`

UnsetGeneratePresentationToolName ensures that no value is present for GeneratePresentationToolName, not even an explicit nil
### GetSystemAiEnabled

`func (o *AiSettingsDto) GetSystemAiEnabled() bool`

GetSystemAiEnabled returns the SystemAiEnabled field if non-nil, zero value otherwise.

### GetSystemAiEnabledOk

`func (o *AiSettingsDto) GetSystemAiEnabledOk() (*bool, bool)`

GetSystemAiEnabledOk returns a tuple with the SystemAiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAiEnabled

`func (o *AiSettingsDto) SetSystemAiEnabled(v bool)`

SetSystemAiEnabled sets SystemAiEnabled field to given value.

### HasSystemAiEnabled

`func (o *AiSettingsDto) HasSystemAiEnabled() bool`

HasSystemAiEnabled returns a boolean if a field has been set.

### GetRecommendedModelForForms

`func (o *AiSettingsDto) GetRecommendedModelForForms() string`

GetRecommendedModelForForms returns the RecommendedModelForForms field if non-nil, zero value otherwise.

### GetRecommendedModelForFormsOk

`func (o *AiSettingsDto) GetRecommendedModelForFormsOk() (*string, bool)`

GetRecommendedModelForFormsOk returns a tuple with the RecommendedModelForForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedModelForForms

`func (o *AiSettingsDto) SetRecommendedModelForForms(v string)`

SetRecommendedModelForForms sets RecommendedModelForForms field to given value.

### HasRecommendedModelForForms

`func (o *AiSettingsDto) HasRecommendedModelForForms() bool`

HasRecommendedModelForForms returns a boolean if a field has been set.

### SetRecommendedModelForFormsNil

`func (o *AiSettingsDto) SetRecommendedModelForFormsNil(b bool)`

 SetRecommendedModelForFormsNil sets the value for RecommendedModelForForms to be an explicit nil

### UnsetRecommendedModelForForms
`func (o *AiSettingsDto) UnsetRecommendedModelForForms()`

UnsetRecommendedModelForForms ensures that no value is present for RecommendedModelForForms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


