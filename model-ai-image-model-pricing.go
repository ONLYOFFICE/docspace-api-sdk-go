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

// checks if the AiImageModelPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiImageModelPricing{}

// AiImageModelPricing The pricing of a single image model.
type AiImageModelPricing struct {
	// The identifier of the model, as the provider expects it on the wire.
	Id NullableString `json:"id"`
	// The display name of the model.
	Alias NullableString `json:"alias,omitempty"`
	// The owner of the model, as reported by the provider.
	OwnedBy NullableString `json:"ownedBy,omitempty"`
	// The provider that serves the model.
	Provider NullableString `json:"provider,omitempty"`
	// The link to the pricing page of the model.
	Link NullableString `json:"link,omitempty"`
	// The price of an image model: per prompt token and per generated image.
	Price AiImagePrice `json:"price"`
}

type _AiImageModelPricing AiImageModelPricing

// NewAiImageModelPricing instantiates a new AiImageModelPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiImageModelPricing(id NullableString, price AiImagePrice) *AiImageModelPricing {
	this := AiImageModelPricing{}
	this.Id = id
	this.Price = price
	return &this
}

// NewAiImageModelPricingWithDefaults instantiates a new AiImageModelPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiImageModelPricingWithDefaults() *AiImageModelPricing {
	this := AiImageModelPricing{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiImageModelPricing) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiImageModelPricing) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *AiImageModelPricing) SetId(v string) {
	o.Id.Set(&v)
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiImageModelPricing) GetAlias() string {
	if o == nil || IsNil(o.Alias.Get()) {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiImageModelPricing) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *AiImageModelPricing) IsAliasSet() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *AiImageModelPricing) SetAlias(v string) {
	o.Alias.Set(&v)
}
// SetAliasNil sets the value for Alias to be an explicit nil
func (o *AiImageModelPricing) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *AiImageModelPricing) UnsetAlias() {
	o.Alias.Unset()
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiImageModelPricing) GetOwnedBy() string {
	if o == nil || IsNil(o.OwnedBy.Get()) {
		var ret string
		return ret
	}
	return *o.OwnedBy.Get()
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiImageModelPricing) GetOwnedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnedBy.Get(), o.OwnedBy.IsSet()
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *AiImageModelPricing) IsOwnedBySet() bool {
	if o != nil && o.OwnedBy.IsSet() {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given NullableString and assigns it to the OwnedBy field.
func (o *AiImageModelPricing) SetOwnedBy(v string) {
	o.OwnedBy.Set(&v)
}
// SetOwnedByNil sets the value for OwnedBy to be an explicit nil
func (o *AiImageModelPricing) SetOwnedByNil() {
	o.OwnedBy.Set(nil)
}

// UnsetOwnedBy ensures that no value is present for OwnedBy, not even an explicit nil
func (o *AiImageModelPricing) UnsetOwnedBy() {
	o.OwnedBy.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiImageModelPricing) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiImageModelPricing) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *AiImageModelPricing) IsProviderSet() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *AiImageModelPricing) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *AiImageModelPricing) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *AiImageModelPricing) UnsetProvider() {
	o.Provider.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiImageModelPricing) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiImageModelPricing) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *AiImageModelPricing) IsLinkSet() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *AiImageModelPricing) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *AiImageModelPricing) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *AiImageModelPricing) UnsetLink() {
	o.Link.Unset()
}

// GetPrice returns the Price field value
func (o *AiImageModelPricing) GetPrice() AiImagePrice {
	if o == nil {
		var ret AiImagePrice
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *AiImageModelPricing) GetPriceOk() (*AiImagePrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *AiImageModelPricing) SetPrice(v AiImagePrice) {
	o.Price = v
}

func (o AiImageModelPricing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiImageModelPricing) ToMap() (map[string]interface{}, error) {
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
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

func (o *AiImageModelPricing) UnmarshalJSON(data []byte) (err error) {
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

	varAiImageModelPricing := _AiImageModelPricing{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiImageModelPricing)

	if err != nil {
		return err
	}

	*o = AiImageModelPricing(varAiImageModelPricing)

	return err
}

type NullableAiImageModelPricing struct {
	value *AiImageModelPricing
	isSet bool
}

func (v NullableAiImageModelPricing) Get() *AiImageModelPricing {
	return v.value
}

func (v *NullableAiImageModelPricing) Set(val *AiImageModelPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableAiImageModelPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableAiImageModelPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiImageModelPricing(val *AiImageModelPricing) *NullableAiImageModelPricing {
	return &NullableAiImageModelPricing{value: val, isSet: true}
}

func (v NullableAiImageModelPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiImageModelPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

