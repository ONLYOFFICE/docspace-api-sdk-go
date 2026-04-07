# AiPricesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**[]AiChatModelPricing**](AiChatModelPricing.md) |  | 
**Embedding** | [**[]AiEmbeddingModelPricing**](AiEmbeddingModelPricing.md) |  | 
**WebSearch** | [**AiWebSearchPricing**](AiWebSearchPricing.md) |  | 
**Currency** | [**CurrencyInfo**](CurrencyInfo.md) |  | 

## Methods

### NewAiPricesResponse

`func NewAiPricesResponse(chat []AiChatModelPricing, embedding []AiEmbeddingModelPricing, webSearch AiWebSearchPricing, currency CurrencyInfo, ) *AiPricesResponse`

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
### GetWebSearch

`func (o *AiPricesResponse) GetWebSearch() AiWebSearchPricing`

GetWebSearch returns the WebSearch field if non-nil, zero value otherwise.

### GetWebSearchOk

`func (o *AiPricesResponse) GetWebSearchOk() (*AiWebSearchPricing, bool)`

GetWebSearchOk returns a tuple with the WebSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearch

`func (o *AiPricesResponse) SetWebSearch(v AiWebSearchPricing)`

SetWebSearch sets WebSearch field to given value.


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


