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

// checks if the AiPricesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPricesResponse{}

// AiPricesResponse struct for AiPricesResponse
type AiPricesResponse struct {
	Chat []AiChatModelPricing `json:"chat"`
	Embedding []AiEmbeddingModelPricing `json:"embedding"`
	WebSearch AiWebSearchPricing `json:"webSearch"`
	Currency CurrencyInfo `json:"currency"`
}

type _AiPricesResponse AiPricesResponse

// NewAiPricesResponse instantiates a new AiPricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPricesResponse(chat []AiChatModelPricing, embedding []AiEmbeddingModelPricing, webSearch AiWebSearchPricing, currency CurrencyInfo) *AiPricesResponse {
	this := AiPricesResponse{}
	this.Chat = chat
	this.Embedding = embedding
	this.WebSearch = webSearch
	this.Currency = currency
	return &this
}

// NewAiPricesResponseWithDefaults instantiates a new AiPricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPricesResponseWithDefaults() *AiPricesResponse {
	this := AiPricesResponse{}
	return &this
}

// GetChat returns the Chat field value
// If the value is explicit nil, the zero value for []AiChatModelPricing will be returned
func (o *AiPricesResponse) GetChat() []AiChatModelPricing {
	if o == nil {
		var ret []AiChatModelPricing
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPricesResponse) GetChatOk() ([]AiChatModelPricing, bool) {
	if o == nil || IsNil(o.Chat) {
		return nil, false
	}
	return o.Chat, true
}

// SetChat sets field value
func (o *AiPricesResponse) SetChat(v []AiChatModelPricing) {
	o.Chat = v
}

// GetEmbedding returns the Embedding field value
// If the value is explicit nil, the zero value for []AiEmbeddingModelPricing will be returned
func (o *AiPricesResponse) GetEmbedding() []AiEmbeddingModelPricing {
	if o == nil {
		var ret []AiEmbeddingModelPricing
		return ret
	}

	return o.Embedding
}

// GetEmbeddingOk returns a tuple with the Embedding field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPricesResponse) GetEmbeddingOk() ([]AiEmbeddingModelPricing, bool) {
	if o == nil || IsNil(o.Embedding) {
		return nil, false
	}
	return o.Embedding, true
}

// SetEmbedding sets field value
func (o *AiPricesResponse) SetEmbedding(v []AiEmbeddingModelPricing) {
	o.Embedding = v
}

// GetWebSearch returns the WebSearch field value
func (o *AiPricesResponse) GetWebSearch() AiWebSearchPricing {
	if o == nil {
		var ret AiWebSearchPricing
		return ret
	}

	return o.WebSearch
}

// GetWebSearchOk returns a tuple with the WebSearch field value
// and a boolean to check if the value has been set.
func (o *AiPricesResponse) GetWebSearchOk() (*AiWebSearchPricing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebSearch, true
}

// SetWebSearch sets field value
func (o *AiPricesResponse) SetWebSearch(v AiWebSearchPricing) {
	o.WebSearch = v
}

// GetCurrency returns the Currency field value
func (o *AiPricesResponse) GetCurrency() CurrencyInfo {
	if o == nil {
		var ret CurrencyInfo
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *AiPricesResponse) GetCurrencyOk() (*CurrencyInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *AiPricesResponse) SetCurrency(v CurrencyInfo) {
	o.Currency = v
}

func (o AiPricesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPricesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Chat != nil {
		toSerialize["chat"] = o.Chat
	}
	if o.Embedding != nil {
		toSerialize["embedding"] = o.Embedding
	}
	toSerialize["webSearch"] = o.WebSearch
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *AiPricesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat",
		"embedding",
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

	varAiPricesResponse := _AiPricesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPricesResponse)

	if err != nil {
		return err
	}

	*o = AiPricesResponse(varAiPricesResponse)

	return err
}

type NullableAiPricesResponse struct {
	value *AiPricesResponse
	isSet bool
}

func (v NullableAiPricesResponse) Get() *AiPricesResponse {
	return v.value
}

func (v *NullableAiPricesResponse) Set(val *AiPricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPricesResponse(val *AiPricesResponse) *NullableAiPricesResponse {
	return &NullableAiPricesResponse{value: val, isSet: true}
}

func (v NullableAiPricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

