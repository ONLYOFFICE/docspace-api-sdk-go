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

// checks if the SetDefaultProviderRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetDefaultProviderRequestDto{}

// SetDefaultProviderRequestDto Request parameters for setting the default AI provider.
type SetDefaultProviderRequestDto struct {
	// AI provider identifier.
	ProviderId *int32 `json:"providerId,omitempty"`
	// Default model identifier to use with this provider.
	DefaultModel NullableString `json:"defaultModel"`
}

type _SetDefaultProviderRequestDto SetDefaultProviderRequestDto

// NewSetDefaultProviderRequestDto instantiates a new SetDefaultProviderRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetDefaultProviderRequestDto(defaultModel NullableString) *SetDefaultProviderRequestDto {
	this := SetDefaultProviderRequestDto{}
	this.DefaultModel = defaultModel
	return &this
}

// NewSetDefaultProviderRequestDtoWithDefaults instantiates a new SetDefaultProviderRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetDefaultProviderRequestDtoWithDefaults() *SetDefaultProviderRequestDto {
	this := SetDefaultProviderRequestDto{}
	return &this
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *SetDefaultProviderRequestDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetDefaultProviderRequestDto) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *SetDefaultProviderRequestDto) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *SetDefaultProviderRequestDto) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetDefaultModel returns the DefaultModel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SetDefaultProviderRequestDto) GetDefaultModel() string {
	if o == nil || o.DefaultModel.Get() == nil {
		var ret string
		return ret
	}

	return *o.DefaultModel.Get()
}

// GetDefaultModelOk returns a tuple with the DefaultModel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetDefaultProviderRequestDto) GetDefaultModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultModel.Get(), o.DefaultModel.IsSet()
}

// SetDefaultModel sets field value
func (o *SetDefaultProviderRequestDto) SetDefaultModel(v string) {
	o.DefaultModel.Set(&v)
}

func (o SetDefaultProviderRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetDefaultProviderRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	toSerialize["defaultModel"] = o.DefaultModel.Get()
	return toSerialize, nil
}

func (o *SetDefaultProviderRequestDto) UnmarshalJSON(data []byte) (err error) {
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

	varSetDefaultProviderRequestDto := _SetDefaultProviderRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSetDefaultProviderRequestDto)

	if err != nil {
		return err
	}

	*o = SetDefaultProviderRequestDto(varSetDefaultProviderRequestDto)

	return err
}

type NullableSetDefaultProviderRequestDto struct {
	value *SetDefaultProviderRequestDto
	isSet bool
}

func (v NullableSetDefaultProviderRequestDto) Get() *SetDefaultProviderRequestDto {
	return v.value
}

func (v *NullableSetDefaultProviderRequestDto) Set(val *SetDefaultProviderRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetDefaultProviderRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetDefaultProviderRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetDefaultProviderRequestDto(val *SetDefaultProviderRequestDto) *NullableSetDefaultProviderRequestDto {
	return &NullableSetDefaultProviderRequestDto{value: val, isSet: true}
}

func (v NullableSetDefaultProviderRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetDefaultProviderRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

