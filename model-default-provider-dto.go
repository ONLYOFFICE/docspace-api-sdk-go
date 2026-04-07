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

// checks if the DefaultProviderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultProviderDto{}

// DefaultProviderDto Default AI provider information.
type DefaultProviderDto struct {
	// AI provider identifier.
	ProviderId *int32 `json:"providerId,omitempty"`
	// Default model identifier used with this provider.
	DefaultModel NullableString `json:"defaultModel"`
	// AI provider title.
	ProviderTitle NullableString `json:"providerTitle,omitempty"`
}

type _DefaultProviderDto DefaultProviderDto

// NewDefaultProviderDto instantiates a new DefaultProviderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultProviderDto(defaultModel NullableString) *DefaultProviderDto {
	this := DefaultProviderDto{}
	this.DefaultModel = defaultModel
	return &this
}

// NewDefaultProviderDtoWithDefaults instantiates a new DefaultProviderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultProviderDtoWithDefaults() *DefaultProviderDto {
	this := DefaultProviderDto{}
	return &this
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *DefaultProviderDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultProviderDto) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *DefaultProviderDto) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *DefaultProviderDto) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetDefaultModel returns the DefaultModel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DefaultProviderDto) GetDefaultModel() string {
	if o == nil || o.DefaultModel.Get() == nil {
		var ret string
		return ret
	}

	return *o.DefaultModel.Get()
}

// GetDefaultModelOk returns a tuple with the DefaultModel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultProviderDto) GetDefaultModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultModel.Get(), o.DefaultModel.IsSet()
}

// SetDefaultModel sets field value
func (o *DefaultProviderDto) SetDefaultModel(v string) {
	o.DefaultModel.Set(&v)
}

// GetProviderTitle returns the ProviderTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultProviderDto) GetProviderTitle() string {
	if o == nil || IsNil(o.ProviderTitle.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderTitle.Get()
}

// GetProviderTitleOk returns a tuple with the ProviderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultProviderDto) GetProviderTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderTitle.Get(), o.ProviderTitle.IsSet()
}

// HasProviderTitle returns a boolean if a field has been set.
func (o *DefaultProviderDto) IsProviderTitleSet() bool {
	if o != nil && o.ProviderTitle.IsSet() {
		return true
	}

	return false
}

// SetProviderTitle gets a reference to the given NullableString and assigns it to the ProviderTitle field.
func (o *DefaultProviderDto) SetProviderTitle(v string) {
	o.ProviderTitle.Set(&v)
}
// SetProviderTitleNil sets the value for ProviderTitle to be an explicit nil
func (o *DefaultProviderDto) SetProviderTitleNil() {
	o.ProviderTitle.Set(nil)
}

// UnsetProviderTitle ensures that no value is present for ProviderTitle, not even an explicit nil
func (o *DefaultProviderDto) UnsetProviderTitle() {
	o.ProviderTitle.Unset()
}

func (o DefaultProviderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultProviderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	toSerialize["defaultModel"] = o.DefaultModel.Get()
	if o.ProviderTitle.IsSet() {
		toSerialize["providerTitle"] = o.ProviderTitle.Get()
	}
	return toSerialize, nil
}

func (o *DefaultProviderDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defaultModel",
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

	varDefaultProviderDto := _DefaultProviderDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDefaultProviderDto)

	if err != nil {
		return err
	}

	*o = DefaultProviderDto(varDefaultProviderDto)

	return err
}

type NullableDefaultProviderDto struct {
	value *DefaultProviderDto
	isSet bool
}

func (v NullableDefaultProviderDto) Get() *DefaultProviderDto {
	return v.value
}

func (v *NullableDefaultProviderDto) Set(val *DefaultProviderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultProviderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultProviderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultProviderDto(val *DefaultProviderDto) *NullableDefaultProviderDto {
	return &NullableDefaultProviderDto{value: val, isSet: true}
}

func (v NullableDefaultProviderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultProviderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

