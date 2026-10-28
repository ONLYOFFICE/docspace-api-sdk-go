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

// checks if the AiEntryPricingDtoAiImagePriceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEntryPricingDtoAiImagePriceDto{}

// AiEntryPricingDtoAiImagePriceDto One AI model or service on the price list: how to name it, who provides it, and what it costs.
type AiEntryPricingDtoAiImagePriceDto struct {
	// The model identifier to send to the AI operations. It is the value to branch on, while `alias` is for display  only.
	Id NullableString `json:"id"`
	// The model name as the vendor writes it, meant to be shown to a person rather than matched on.
	Alias NullableString `json:"alias"`
	// Who runs the model. Two entries can share a provider, and one provider's models can be priced quite  differently, so the price always belongs to the entry and never to the provider.
	Provider NullableString `json:"provider"`
	// The absolute URL of the provider's icon, for rendering next to the entry.
	Image NullableString `json:"image"`
	// What the entry costs, in the currency the answer names. Amounts per token are normalised per million  tokens, so they are not the price of a single call.
	Price AiImagePriceDto `json:"price"`
	// The provider's own page for the model, for a person to read the model's terms. It is empty when the  provider publishes none.
	Link NullableString `json:"link"`
}

type _AiEntryPricingDtoAiImagePriceDto AiEntryPricingDtoAiImagePriceDto

// NewAiEntryPricingDtoAiImagePriceDto instantiates a new AiEntryPricingDtoAiImagePriceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEntryPricingDtoAiImagePriceDto(id NullableString, alias NullableString, provider NullableString, image NullableString, price AiImagePriceDto, link NullableString) *AiEntryPricingDtoAiImagePriceDto {
	this := AiEntryPricingDtoAiImagePriceDto{}
	this.Id = id
	this.Alias = alias
	this.Provider = provider
	this.Image = image
	this.Price = price
	this.Link = link
	return &this
}

// NewAiEntryPricingDtoAiImagePriceDtoWithDefaults instantiates a new AiEntryPricingDtoAiImagePriceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEntryPricingDtoAiImagePriceDtoWithDefaults() *AiEntryPricingDtoAiImagePriceDto {
	this := AiEntryPricingDtoAiImagePriceDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *AiEntryPricingDtoAiImagePriceDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetAlias returns the Alias field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetAlias() string {
	if o == nil || o.Alias.Get() == nil {
		var ret string
		return ret
	}

	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// SetAlias sets field value
func (o *AiEntryPricingDtoAiImagePriceDto) SetAlias(v string) {
	o.Alias.Set(&v)
}

// GetProvider returns the Provider field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetProvider() string {
	if o == nil || o.Provider.Get() == nil {
		var ret string
		return ret
	}

	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// SetProvider sets field value
func (o *AiEntryPricingDtoAiImagePriceDto) SetProvider(v string) {
	o.Provider.Set(&v)
}

// GetImage returns the Image field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetImage() string {
	if o == nil || o.Image.Get() == nil {
		var ret string
		return ret
	}

	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// SetImage sets field value
func (o *AiEntryPricingDtoAiImagePriceDto) SetImage(v string) {
	o.Image.Set(&v)
}

// GetPrice returns the Price field value
func (o *AiEntryPricingDtoAiImagePriceDto) GetPrice() AiImagePriceDto {
	if o == nil {
		var ret AiImagePriceDto
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *AiEntryPricingDtoAiImagePriceDto) GetPriceOk() (*AiImagePriceDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *AiEntryPricingDtoAiImagePriceDto) SetPrice(v AiImagePriceDto) {
	o.Price = v
}

// GetLink returns the Link field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetLink() string {
	if o == nil || o.Link.Get() == nil {
		var ret string
		return ret
	}

	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEntryPricingDtoAiImagePriceDto) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// SetLink sets field value
func (o *AiEntryPricingDtoAiImagePriceDto) SetLink(v string) {
	o.Link.Set(&v)
}

func (o AiEntryPricingDtoAiImagePriceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEntryPricingDtoAiImagePriceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["alias"] = o.Alias.Get()
	toSerialize["provider"] = o.Provider.Get()
	toSerialize["image"] = o.Image.Get()
	toSerialize["price"] = o.Price
	toSerialize["link"] = o.Link.Get()
	return toSerialize, nil
}

func (o *AiEntryPricingDtoAiImagePriceDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"alias",
		"provider",
		"image",
		"price",
		"link",
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

	varAiEntryPricingDtoAiImagePriceDto := _AiEntryPricingDtoAiImagePriceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEntryPricingDtoAiImagePriceDto)

	if err != nil {
		return err
	}

	*o = AiEntryPricingDtoAiImagePriceDto(varAiEntryPricingDtoAiImagePriceDto)

	return err
}

type NullableAiEntryPricingDtoAiImagePriceDto struct {
	value *AiEntryPricingDtoAiImagePriceDto
	isSet bool
}

func (v NullableAiEntryPricingDtoAiImagePriceDto) Get() *AiEntryPricingDtoAiImagePriceDto {
	return v.value
}

func (v *NullableAiEntryPricingDtoAiImagePriceDto) Set(val *AiEntryPricingDtoAiImagePriceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEntryPricingDtoAiImagePriceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEntryPricingDtoAiImagePriceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEntryPricingDtoAiImagePriceDto(val *AiEntryPricingDtoAiImagePriceDto) *NullableAiEntryPricingDtoAiImagePriceDto {
	return &NullableAiEntryPricingDtoAiImagePriceDto{value: val, isSet: true}
}

func (v NullableAiEntryPricingDtoAiImagePriceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEntryPricingDtoAiImagePriceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

