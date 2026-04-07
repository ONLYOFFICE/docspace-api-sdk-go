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

// checks if the VectorizationSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VectorizationSettingsDto{}

// VectorizationSettingsDto The vectorization settings.
type VectorizationSettingsDto struct {
	Type *EmbeddingProviderType `json:"type,omitempty"`
	// Indicates whether the embedding provider API key needs to be reconfigured.
	NeedReset *bool `json:"needReset,omitempty"`
}

// NewVectorizationSettingsDto instantiates a new VectorizationSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVectorizationSettingsDto() *VectorizationSettingsDto {
	this := VectorizationSettingsDto{}
	return &this
}

// NewVectorizationSettingsDtoWithDefaults instantiates a new VectorizationSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVectorizationSettingsDtoWithDefaults() *VectorizationSettingsDto {
	this := VectorizationSettingsDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VectorizationSettingsDto) GetType() EmbeddingProviderType {
	if o == nil || IsNil(o.Type) {
		var ret EmbeddingProviderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VectorizationSettingsDto) GetTypeOk() (*EmbeddingProviderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VectorizationSettingsDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EmbeddingProviderType and assigns it to the Type field.
func (o *VectorizationSettingsDto) SetType(v EmbeddingProviderType) {
	o.Type = &v
}

// GetNeedReset returns the NeedReset field value if set, zero value otherwise.
func (o *VectorizationSettingsDto) GetNeedReset() bool {
	if o == nil || IsNil(o.NeedReset) {
		var ret bool
		return ret
	}
	return *o.NeedReset
}

// GetNeedResetOk returns a tuple with the NeedReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VectorizationSettingsDto) GetNeedResetOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedReset) {
		return nil, false
	}
	return o.NeedReset, true
}

// HasNeedReset returns a boolean if a field has been set.
func (o *VectorizationSettingsDto) IsNeedResetSet() bool {
	if o != nil && !IsNil(o.NeedReset) {
		return true
	}

	return false
}

// SetNeedReset gets a reference to the given bool and assigns it to the NeedReset field.
func (o *VectorizationSettingsDto) SetNeedReset(v bool) {
	o.NeedReset = &v
}

func (o VectorizationSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VectorizationSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.NeedReset) {
		toSerialize["needReset"] = o.NeedReset
	}
	return toSerialize, nil
}

type NullableVectorizationSettingsDto struct {
	value *VectorizationSettingsDto
	isSet bool
}

func (v NullableVectorizationSettingsDto) Get() *VectorizationSettingsDto {
	return v.value
}

func (v *NullableVectorizationSettingsDto) Set(val *VectorizationSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableVectorizationSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableVectorizationSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVectorizationSettingsDto(val *VectorizationSettingsDto) *NullableVectorizationSettingsDto {
	return &NullableVectorizationSettingsDto{value: val, isSet: true}
}

func (v NullableVectorizationSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVectorizationSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

