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

// checks if the ModelDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelDto{}

// ModelDto The AI model information.
type ModelDto struct {
	// The unique identifier of the AI provider that offers this model.
	ProviderId *int32 `json:"providerId,omitempty"`
	// The human-readable display name of the AI provider (e.g., OpenAI, Anthropic).
	ProviderTitle NullableString `json:"providerTitle"`
	// The model identifier as recognized by the AI provider (e.g., gpt-4o, claude-sonnet-4-20250514).
	ModelId NullableString `json:"modelId"`
	Price *AiChatPrice `json:"price,omitempty"`
	Currency *CurrencyInfo `json:"currency,omitempty"`
}

type _ModelDto ModelDto

// NewModelDto instantiates a new ModelDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelDto(providerTitle NullableString, modelId NullableString) *ModelDto {
	this := ModelDto{}
	this.ProviderTitle = providerTitle
	this.ModelId = modelId
	return &this
}

// NewModelDtoWithDefaults instantiates a new ModelDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelDtoWithDefaults() *ModelDto {
	this := ModelDto{}
	return &this
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *ModelDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDto) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *ModelDto) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *ModelDto) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetProviderTitle returns the ProviderTitle field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ModelDto) GetProviderTitle() string {
	if o == nil || o.ProviderTitle.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProviderTitle.Get()
}

// GetProviderTitleOk returns a tuple with the ProviderTitle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelDto) GetProviderTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderTitle.Get(), o.ProviderTitle.IsSet()
}

// SetProviderTitle sets field value
func (o *ModelDto) SetProviderTitle(v string) {
	o.ProviderTitle.Set(&v)
}

// GetModelId returns the ModelId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ModelDto) GetModelId() string {
	if o == nil || o.ModelId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ModelId.Get()
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelDto) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModelId.Get(), o.ModelId.IsSet()
}

// SetModelId sets field value
func (o *ModelDto) SetModelId(v string) {
	o.ModelId.Set(&v)
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ModelDto) GetPrice() AiChatPrice {
	if o == nil || IsNil(o.Price) {
		var ret AiChatPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDto) GetPriceOk() (*AiChatPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ModelDto) IsPriceSet() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given AiChatPrice and assigns it to the Price field.
func (o *ModelDto) SetPrice(v AiChatPrice) {
	o.Price = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ModelDto) GetCurrency() CurrencyInfo {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyInfo
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDto) GetCurrencyOk() (*CurrencyInfo, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ModelDto) IsCurrencySet() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyInfo and assigns it to the Currency field.
func (o *ModelDto) SetCurrency(v CurrencyInfo) {
	o.Currency = &v
}

func (o ModelDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	toSerialize["providerTitle"] = o.ProviderTitle.Get()
	toSerialize["modelId"] = o.ModelId.Get()
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

func (o *ModelDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"providerTitle",
		"modelId",
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

	varModelDto := _ModelDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varModelDto)

	if err != nil {
		return err
	}

	*o = ModelDto(varModelDto)

	return err
}

type NullableModelDto struct {
	value *ModelDto
	isSet bool
}

func (v NullableModelDto) Get() *ModelDto {
	return v.value
}

func (v *NullableModelDto) Set(val *ModelDto) {
	v.value = val
	v.isSet = true
}

func (v NullableModelDto) IsSet() bool {
	return v.isSet
}

func (v *NullableModelDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelDto(val *ModelDto) *NullableModelDto {
	return &NullableModelDto{value: val, isSet: true}
}

func (v NullableModelDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

