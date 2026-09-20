# AiPricesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**[]AiEntryPricingDtoAiChatPriceDto**](AiEntryPricingDtoAiChatPriceDto.md) | The chat models on offer, each priced per million prompt and completion tokens. A model listed here is one  the installation can bill for, not necessarily one this portal may use -  `GET api/2.0/portal/payment/ai-model/restrictions` says which are allowed. | 
**Embedding** | [**[]AiEntryPricingDtoAiEmbeddingPriceDto**](AiEntryPricingDtoAiEmbeddingPriceDto.md) | The embedding models on offer, priced per million tokens of input; an embedding model has no completion  side, so its price object carries `prompt` alone. | 
**Image** | [**[]AiEntryPricingDtoAiImagePriceDto**](AiEntryPricingDtoAiImagePriceDto.md) | The image models on offer, priced per million prompt and completion tokens plus a price for each image  produced. | 
**WebSearch** | [**[]AiEntryPricingDtoDecimal**](AiEntryPricingDtoDecimal.md) | The web search providers on offer. Their `price` is a bare number - the cost of one search - rather than  an object, because there are no tokens to distinguish. | 
**Currency** | [**CurrencyInfo**](CurrencyInfo.md) | The currency every price above is expressed in, with its ISO code and symbol. One answer never mixes  currencies, so this is the only place to read it. | 

## Methods

### NewAiPricesDto

`func NewAiPricesDto(chat []AiEntryPricingDtoAiChatPriceDto, embedding []AiEntryPricingDtoAiEmbeddingPriceDto, image []AiEntryPricingDtoAiImagePriceDto, webSearch []AiEntryPricingDtoDecimal, currency CurrencyInfo, ) *AiPricesDto`

NewAiPricesDto instantiates a new AiPricesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPricesDtoWithDefaults

`func NewAiPricesDtoWithDefaults() *AiPricesDto`

NewAiPricesDtoWithDefaults instantiates a new AiPricesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *AiPricesDto) GetChat() []AiEntryPricingDtoAiChatPriceDto`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *AiPricesDto) GetChatOk() (*[]AiEntryPricingDtoAiChatPriceDto, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *AiPricesDto) SetChat(v []AiEntryPricingDtoAiChatPriceDto)`

SetChat sets Chat field to given value.


### SetChatNil

`func (o *AiPricesDto) SetChatNil(b bool)`

 SetChatNil sets the value for Chat to be an explicit nil

### UnsetChat
`func (o *AiPricesDto) UnsetChat()`

UnsetChat ensures that no value is present for Chat, not even an explicit nil
### GetEmbedding

`func (o *AiPricesDto) GetEmbedding() []AiEntryPricingDtoAiEmbeddingPriceDto`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *AiPricesDto) GetEmbeddingOk() (*[]AiEntryPricingDtoAiEmbeddingPriceDto, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *AiPricesDto) SetEmbedding(v []AiEntryPricingDtoAiEmbeddingPriceDto)`

SetEmbedding sets Embedding field to given value.


### SetEmbeddingNil

`func (o *AiPricesDto) SetEmbeddingNil(b bool)`

 SetEmbeddingNil sets the value for Embedding to be an explicit nil

### UnsetEmbedding
`func (o *AiPricesDto) UnsetEmbedding()`

UnsetEmbedding ensures that no value is present for Embedding, not even an explicit nil
### GetImage

`func (o *AiPricesDto) GetImage() []AiEntryPricingDtoAiImagePriceDto`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AiPricesDto) GetImageOk() (*[]AiEntryPricingDtoAiImagePriceDto, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AiPricesDto) SetImage(v []AiEntryPricingDtoAiImagePriceDto)`

SetImage sets Image field to given value.


### SetImageNil

`func (o *AiPricesDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *AiPricesDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetWebSearch

`func (o *AiPricesDto) GetWebSearch() []AiEntryPricingDtoDecimal`

GetWebSearch returns the WebSearch field if non-nil, zero value otherwise.

### GetWebSearchOk

`func (o *AiPricesDto) GetWebSearchOk() (*[]AiEntryPricingDtoDecimal, bool)`

GetWebSearchOk returns a tuple with the WebSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearch

`func (o *AiPricesDto) SetWebSearch(v []AiEntryPricingDtoDecimal)`

SetWebSearch sets WebSearch field to given value.


### SetWebSearchNil

`func (o *AiPricesDto) SetWebSearchNil(b bool)`

 SetWebSearchNil sets the value for WebSearch to be an explicit nil

### UnsetWebSearch
`func (o *AiPricesDto) UnsetWebSearch()`

UnsetWebSearch ensures that no value is present for WebSearch, not even an explicit nil
### GetCurrency

`func (o *AiPricesDto) GetCurrency() CurrencyInfo`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AiPricesDto) GetCurrencyOk() (*CurrencyInfo, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AiPricesDto) SetCurrency(v CurrencyInfo)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


