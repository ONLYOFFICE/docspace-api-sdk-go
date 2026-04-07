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

// checks if the AiEmbeddingModelPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEmbeddingModelPricing{}

// AiEmbeddingModelPricing struct for AiEmbeddingModelPricing
type AiEmbeddingModelPricing struct {
	Id NullableString `json:"id"`
	Alias NullableString `json:"alias,omitempty"`
	OwnedBy NullableString `json:"ownedBy,omitempty"`
	Provider NullableString `json:"provider,omitempty"`
	Price AiEmbeddingPrice `json:"price"`
}

type _AiEmbeddingModelPricing AiEmbeddingModelPricing

// NewAiEmbeddingModelPricing instantiates a new AiEmbeddingModelPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEmbeddingModelPricing(id NullableString, price AiEmbeddingPrice) *AiEmbeddingModelPricing {
	this := AiEmbeddingModelPricing{}
	this.Id = id
	this.Price = price
	return &this
}

// NewAiEmbeddingModelPricingWithDefaults instantiates a new AiEmbeddingModelPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEmbeddingModelPricingWithDefaults() *AiEmbeddingModelPricing {
	this := AiEmbeddingModelPricing{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiEmbeddingModelPricing) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmbeddingModelPricing) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *AiEmbeddingModelPricing) SetId(v string) {
	o.Id.Set(&v)
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmbeddingModelPricing) GetAlias() string {
	if o == nil || IsNil(o.Alias.Get()) {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmbeddingModelPricing) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *AiEmbeddingModelPricing) IsAliasSet() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *AiEmbeddingModelPricing) SetAlias(v string) {
	o.Alias.Set(&v)
}
// SetAliasNil sets the value for Alias to be an explicit nil
func (o *AiEmbeddingModelPricing) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *AiEmbeddingModelPricing) UnsetAlias() {
	o.Alias.Unset()
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmbeddingModelPricing) GetOwnedBy() string {
	if o == nil || IsNil(o.OwnedBy.Get()) {
		var ret string
		return ret
	}
	return *o.OwnedBy.Get()
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmbeddingModelPricing) GetOwnedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnedBy.Get(), o.OwnedBy.IsSet()
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *AiEmbeddingModelPricing) IsOwnedBySet() bool {
	if o != nil && o.OwnedBy.IsSet() {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given NullableString and assigns it to the OwnedBy field.
func (o *AiEmbeddingModelPricing) SetOwnedBy(v string) {
	o.OwnedBy.Set(&v)
}
// SetOwnedByNil sets the value for OwnedBy to be an explicit nil
func (o *AiEmbeddingModelPricing) SetOwnedByNil() {
	o.OwnedBy.Set(nil)
}

// UnsetOwnedBy ensures that no value is present for OwnedBy, not even an explicit nil
func (o *AiEmbeddingModelPricing) UnsetOwnedBy() {
	o.OwnedBy.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmbeddingModelPricing) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmbeddingModelPricing) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *AiEmbeddingModelPricing) IsProviderSet() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *AiEmbeddingModelPricing) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *AiEmbeddingModelPricing) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *AiEmbeddingModelPricing) UnsetProvider() {
	o.Provider.Unset()
}

// GetPrice returns the Price field value
func (o *AiEmbeddingModelPricing) GetPrice() AiEmbeddingPrice {
	if o == nil {
		var ret AiEmbeddingPrice
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *AiEmbeddingModelPricing) GetPriceOk() (*AiEmbeddingPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *AiEmbeddingModelPricing) SetPrice(v AiEmbeddingPrice) {
	o.Price = v
}

func (o AiEmbeddingModelPricing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEmbeddingModelPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if o.OwnedBy.IsSet() {
		toSerialize["ownedBy"] = o.OwnedBy.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

func (o *AiEmbeddingModelPricing) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"price",
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

	varAiEmbeddingModelPricing := _AiEmbeddingModelPricing{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEmbeddingModelPricing)

	if err != nil {
		return err
	}

	*o = AiEmbeddingModelPricing(varAiEmbeddingModelPricing)

	return err
}

type NullableAiEmbeddingModelPricing struct {
	value *AiEmbeddingModelPricing
	isSet bool
}

func (v NullableAiEmbeddingModelPricing) Get() *AiEmbeddingModelPricing {
	return v.value
}

func (v *NullableAiEmbeddingModelPricing) Set(val *AiEmbeddingModelPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEmbeddingModelPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEmbeddingModelPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEmbeddingModelPricing(val *AiEmbeddingModelPricing) *NullableAiEmbeddingModelPricing {
	return &NullableAiEmbeddingModelPricing{value: val, isSet: true}
}

func (v NullableAiEmbeddingModelPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEmbeddingModelPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

