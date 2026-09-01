# AiPricesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**[]AiChatModelPricing**](AiChatModelPricing.md) | The pricing of every available chat model. | 
**Embedding** | [**[]AiEmbeddingModelPricing**](AiEmbeddingModelPricing.md) | The pricing of every available embedding model. | 
**Image** | [**[]AiImageModelPricing**](AiImageModelPricing.md) | The pricing of every available image model. | 
**Search** | [**[]AiWebSearchPricing**](AiWebSearchPricing.md) | The pricing of every available web search provider. | 
**Currency** | [**CurrencyInfo**](CurrencyInfo.md) | The currency the AI prices are quoted in. | 

## Methods

### NewAiPricesResponse

`func NewAiPricesResponse(chat []AiChatModelPricing, embedding []AiEmbeddingModelPricing, image []AiImageModelPricing, search []AiWebSearchPricing, currency CurrencyInfo, ) *AiPricesResponse`

NewAiPricesResponse instantiates a new AiPricesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPricesResponseWithDefaults

`func NewAiPricesResponseWithDefaults() *AiPricesResponse`

NewAiPricesResponseWithDefaults instantiates a new AiPricesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *AiPricesResponse) GetChat() []AiChatModelPricing`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *AiPricesResponse) GetChatOk() (*[]AiChatModelPricing, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *AiPricesResponse) SetChat(v []AiChatModelPricing)`

SetChat sets Chat field to given value.


### SetChatNil

`func (o *AiPricesResponse) SetChatNil(b bool)`

 SetChatNil sets the value for Chat to be an explicit nil

### UnsetChat
`func (o *AiPricesResponse) UnsetChat()`

UnsetChat ensures that no value is present for Chat, not even an explicit nil
### GetEmbedding

`func (o *AiPricesResponse) GetEmbedding() []AiEmbeddingModelPricing`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *AiPricesResponse) GetEmbeddingOk() (*[]AiEmbeddingModelPricing, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *AiPricesResponse) SetEmbedding(v []AiEmbeddingModelPricing)`

SetEmbedding sets Embedding field to given value.


### SetEmbeddingNil

`func (o *AiPricesResponse) SetEmbeddingNil(b bool)`

 SetEmbeddingNil sets the value for Embedding to be an explicit nil

### UnsetEmbedding
`func (o *AiPricesResponse) UnsetEmbedding()`

UnsetEmbedding ensures that no value is present for Embedding, not even an explicit nil
### GetImage

`func (o *AiPricesResponse) GetImage() []AiImageModelPricing`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AiPricesResponse) GetImageOk() (*[]AiImageModelPricing, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AiPricesResponse) SetImage(v []AiImageModelPricing)`

SetImage sets Image field to given value.


### SetImageNil

`func (o *AiPricesResponse) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *AiPricesResponse) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetSearch

`func (o *AiPricesResponse) GetSearch() []AiWebSearchPricing`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *AiPricesResponse) GetSearchOk() (*[]AiWebSearchPricing, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *AiPricesResponse) SetSearch(v []AiWebSearchPricing)`

SetSearch sets Search field to given value.


### SetSearchNil

`func (o *AiPricesResponse) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *AiPricesResponse) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetCurrency

`func (o *AiPricesResponse) GetCurrency() CurrencyInfo`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AiPricesResponse) GetCurrencyOk() (*CurrencyInfo, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AiPricesResponse) SetCurrency(v CurrencyInfo)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


