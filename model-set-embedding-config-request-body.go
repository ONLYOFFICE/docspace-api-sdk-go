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

// checks if the SetEmbeddingConfigRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetEmbeddingConfigRequestBody{}

// SetEmbeddingConfigRequestBody Parameters for configuring the embedding provider.
type SetEmbeddingConfigRequestBody struct {
	Type *EmbeddingProviderType `json:"type,omitempty"`
	// The API key for the selected embedding provider. Pass null to keep the existing key unchanged.
	Key NullableString `json:"key,omitempty"`
}

// NewSetEmbeddingConfigRequestBody instantiates a new SetEmbeddingConfigRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetEmbeddingConfigRequestBody() *SetEmbeddingConfigRequestBody {
	this := SetEmbeddingConfigRequestBody{}
	return &this
}

// NewSetEmbeddingConfigRequestBodyWithDefaults instantiates a new SetEmbeddingConfigRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetEmbeddingConfigRequestBodyWithDefaults() *SetEmbeddingConfigRequestBody {
	this := SetEmbeddingConfigRequestBody{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SetEmbeddingConfigRequestBody) GetType() EmbeddingProviderType {
	if o == nil || IsNil(o.Type) {
		var ret EmbeddingProviderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEmbeddingConfigRequestBody) GetTypeOk() (*EmbeddingProviderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SetEmbeddingConfigRequestBody) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EmbeddingProviderType and assigns it to the Type field.
func (o *SetEmbeddingConfigRequestBody) SetType(v EmbeddingProviderType) {
	o.Type = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetEmbeddingConfigRequestBody) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetEmbeddingConfigRequestBody) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *SetEmbeddingConfigRequestBody) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *SetEmbeddingConfigRequestBody) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *SetEmbeddingConfigRequestBody) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *SetEmbeddingConfigRequestBody) UnsetKey() {
	o.Key.Unset()
}

func (o SetEmbeddingConfigRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetEmbeddingConfigRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	return toSerialize, nil
}

type NullableSetEmbeddingConfigRequestBody struct {
	value *SetEmbeddingConfigRequestBody
	isSet bool
}

func (v NullableSetEmbeddingConfigRequestBody) Get() *SetEmbeddingConfigRequestBody {
	return v.value
}

func (v *NullableSetEmbeddingConfigRequestBody) Set(val *SetEmbeddingConfigRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetEmbeddingConfigRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetEmbeddingConfigRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetEmbeddingConfigRequestBody(val *SetEmbeddingConfigRequestBody) *NullableSetEmbeddingConfigRequestBody {
	return &NullableSetEmbeddingConfigRequestBody{value: val, isSet: true}
}

func (v NullableSetEmbeddingConfigRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetEmbeddingConfigRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

