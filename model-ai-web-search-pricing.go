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

// AiWebSearchPricing struct for AiWebSearchPricing
type AiWebSearchPricing struct {
	Provider NullableString `json:"provider,omitempty"`
	Search *float64 `json:"search,omitempty"`
	Contents *float64 `json:"contents,omitempty"`
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

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *AiWebSearchPricing) GetSearch() float64 {
	if o == nil || IsNil(o.Search) {
		var ret float64
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchPricing) GetSearchOk() (*float64, bool) {
	if o == nil || IsNil(o.Search) {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *AiWebSearchPricing) IsSearchSet() bool {
	if o != nil && !IsNil(o.Search) {
		return true
	}

	return false
}

// SetSearch gets a reference to the given float64 and assigns it to the Search field.
func (o *AiWebSearchPricing) SetSearch(v float64) {
	o.Search = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *AiWebSearchPricing) GetContents() float64 {
	if o == nil || IsNil(o.Contents) {
		var ret float64
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchPricing) GetContentsOk() (*float64, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *AiWebSearchPricing) IsContentsSet() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given float64 and assigns it to the Contents field.
func (o *AiWebSearchPricing) SetContents(v float64) {
	o.Contents = &v
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
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if !IsNil(o.Search) {
		toSerialize["search"] = o.Search
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
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

