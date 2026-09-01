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
)

// checks if the AiWebSearchPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiWebSearchPricing{}

// AiWebSearchPricing The pricing of a single web search provider, per request.
type AiWebSearchPricing struct {
	// The identifier of the web search provider.
	Id NullableString `json:"id,omitempty"`
	// The provider that serves the web search requests.
	Provider NullableString `json:"provider,omitempty"`
	// The price of a single web search request.
	Price *float64 `json:"price,omitempty"`
	// The link to the pricing page of the provider.
	Link NullableString `json:"link,omitempty"`
}

// NewAiWebSearchPricing instantiates a new AiWebSearchPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiWebSearchPricing() *AiWebSearchPricing {
	this := AiWebSearchPricing{}
	return &this
}

// NewAiWebSearchPricingWithDefaults instantiates a new AiWebSearchPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiWebSearchPricingWithDefaults() *AiWebSearchPricing {
	this := AiWebSearchPricing{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiWebSearchPricing) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiWebSearchPricing) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AiWebSearchPricing) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *AiWebSearchPricing) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *AiWebSearchPricing) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AiWebSearchPricing) UnsetId() {
	o.Id.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiWebSearchPricing) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiWebSearchPricing) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *AiWebSearchPricing) IsProviderSet() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *AiWebSearchPricing) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *AiWebSearchPricing) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *AiWebSearchPricing) UnsetProvider() {
	o.Provider.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AiWebSearchPricing) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchPricing) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AiWebSearchPricing) IsPriceSet() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *AiWebSearchPricing) SetPrice(v float64) {
	o.Price = &v
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiWebSearchPricing) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiWebSearchPricing) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *AiWebSearchPricing) IsLinkSet() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *AiWebSearchPricing) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *AiWebSearchPricing) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *AiWebSearchPricing) UnsetLink() {
	o.Link.Unset()
}

func (o AiWebSearchPricing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiWebSearchPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	return toSerialize, nil
}

type NullableAiWebSearchPricing struct {
	value *AiWebSearchPricing
	isSet bool
}

func (v NullableAiWebSearchPricing) Get() *AiWebSearchPricing {
	return v.value
}

func (v *NullableAiWebSearchPricing) Set(val *AiWebSearchPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableAiWebSearchPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableAiWebSearchPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiWebSearchPricing(val *AiWebSearchPricing) *NullableAiWebSearchPricing {
	return &NullableAiWebSearchPricing{value: val, isSet: true}
}

func (v NullableAiWebSearchPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiWebSearchPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

