// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docspace_api_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AiPricesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPricesDto{}

// AiPricesDto What the AI features cost out of the portal wallet, grouped by the kind of model, in one currency.
type AiPricesDto struct {
	// The chat models on offer, each priced per million prompt and completion tokens. A model listed here is one  the installation can bill for, not necessarily one this portal may use -  `GET api/2.0/portal/payment/ai-model/restrictions` says which are allowed.
	Chat []AiEntryPricingDtoAiChatPriceDto `json:"chat"`
	// The embedding models on offer, priced per million tokens of input; an embedding model has no completion  side, so its price object carries `prompt` alone.
	Embedding []AiEntryPricingDtoAiEmbeddingPriceDto `json:"embedding"`
	// The image models on offer, priced per million prompt and completion tokens plus a price for each image  produced.
	Image []AiEntryPricingDtoAiImagePriceDto `json:"image"`
	// The web search providers on offer. Their `price` is a bare number - the cost of one search - rather than  an object, because there are no tokens to distinguish.
	WebSearch []AiEntryPricingDtoDecimal `json:"webSearch"`
	// The currency every price above is expressed in, with its ISO code and symbol. One answer never mixes  currencies, so this is the only place to read it.
	Currency CurrencyInfo `json:"currency"`
}

type _AiPricesDto AiPricesDto

// NewAiPricesDto instantiates a new AiPricesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPricesDto(chat []AiEntryPricingDtoAiChatPriceDto, embedding []AiEntryPricingDtoAiEmbeddingPriceDto, image []AiEntryPricingDtoAiImagePriceDto, webSearch []AiEntryPricingDtoDecimal, currency CurrencyInfo) *AiPricesDto {
	this := AiPricesDto{}
	this.Chat = chat
	this.Embedding = embedding
	this.Image = image
	this.WebSearch = webSearch
	this.Currency = currency
	return &this
}

// NewAiPricesDtoWithDefaults instantiates a new AiPricesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPricesDtoWithDefaults() *AiPricesDto {
	this := AiPricesDto{}
	return &this
}

// GetChat returns the Chat field value
// If the value is explicit nil, the zero value for []AiEntryPricingDtoAiChatPriceDto will be returned
func (o *AiPricesDto) GetChat() []AiEntryPricingDtoAiChatPriceDto {
	if o == nil {
		var ret []AiEntryPricingDtoAiChatPriceDto
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPricesDto) GetChatOk() ([]AiEntryPricingDtoAiChatPriceDto, bool) {
	if o == nil || IsNil(o.Chat) {
		return nil, false
	}
	return o.Chat, true
}

// SetChat sets field value
func (o *AiPricesDto) SetChat(v []AiEntryPricingDtoAiChatPriceDto) {
	o.Chat = v
}

// GetEmbedding returns the Embedding field value
// If the value is explicit nil, the zero value for []AiEntryPricingDtoAiEmbeddingPriceDto will be returned
func (o *AiPricesDto) GetEmbedding() []AiEntryPricingDtoAiEmbeddingPriceDto {
	if o == nil {
		var ret []AiEntryPricingDtoAiEmbeddingPriceDto
		return ret
	}

	return o.Embedding
}

// GetEmbeddingOk returns a tuple with the Embedding field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPricesDto) GetEmbeddingOk() ([]AiEntryPricingDtoAiEmbeddingPriceDto, bool) {
	if o == nil || IsNil(o.Embedding) {
		return nil, false
	}
	return o.Embedding, true
}

// SetEmbedding sets field value
func (o *AiPricesDto) SetEmbedding(v []AiEntryPricingDtoAiEmbeddingPriceDto) {
	o.Embedding = v
}

// GetImage returns the Image field value
// If the value is explicit nil, the zero value for []AiEntryPricingDtoAiImagePriceDto will be returned
func (o *AiPricesDto) GetImage() []AiEntryPricingDtoAiImagePriceDto {
	if o == nil {
		var ret []AiEntryPricingDtoAiImagePriceDto
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPricesDto) GetImageOk() ([]AiEntryPricingDtoAiImagePriceDto, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// SetImage sets field value
func (o *AiPricesDto) SetImage(v []AiEntryPricingDtoAiImagePriceDto) {
	o.Image = v
}

// GetWebSearch returns the WebSearch field value
// If the value is explicit nil, the zero value for []AiEntryPricingDtoDecimal will be returned
func (o *AiPricesDto) GetWebSearch() []AiEntryPricingDtoDecimal {
	if o == nil {
		var ret []AiEntryPricingDtoDecimal
		return ret
	}

	return o.WebSearch
}

// GetWebSearchOk returns a tuple with the WebSearch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPricesDto) GetWebSearchOk() ([]AiEntryPricingDtoDecimal, bool) {
	if o == nil || IsNil(o.WebSearch) {
		return nil, false
	}
	return o.WebSearch, true
}

// SetWebSearch sets field value
func (o *AiPricesDto) SetWebSearch(v []AiEntryPricingDtoDecimal) {
	o.WebSearch = v
}

// GetCurrency returns the Currency field value
func (o *AiPricesDto) GetCurrency() CurrencyInfo {
	if o == nil {
		var ret CurrencyInfo
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *AiPricesDto) GetCurrencyOk() (*CurrencyInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *AiPricesDto) SetCurrency(v CurrencyInfo) {
	o.Currency = v
}

func (o AiPricesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPricesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Chat != nil {
		toSerialize["chat"] = o.Chat
	}
	if o.Embedding != nil {
		toSerialize["embedding"] = o.Embedding
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.WebSearch != nil {
		toSerialize["webSearch"] = o.WebSearch
	}
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *AiPricesDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat",
		"embedding",
		"image",
		"webSearch",
		"currency",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAiPricesDto := _AiPricesDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPricesDto)

	if err != nil {
		return err
	}

	*o = AiPricesDto(varAiPricesDto)

	return err
}

type NullableAiPricesDto struct {
	value *AiPricesDto
	isSet bool
}

func (v NullableAiPricesDto) Get() *AiPricesDto {
	return v.value
}

func (v *NullableAiPricesDto) Set(val *AiPricesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPricesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPricesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPricesDto(val *AiPricesDto) *NullableAiPricesDto {
	return &NullableAiPricesDto{value: val, isSet: true}
}

func (v NullableAiPricesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPricesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

