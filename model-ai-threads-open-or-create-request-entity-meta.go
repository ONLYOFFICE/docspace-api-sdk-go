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

// checks if the AiThreadsOpenOrCreateRequestEntityMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadsOpenOrCreateRequestEntityMeta{}

// AiThreadsOpenOrCreateRequestEntityMeta Optional entity hint (lib 0.5.64): only `entityId` is read; the pair is re-resolved server-side before reaching the provider as metadata.
type AiThreadsOpenOrCreateRequestEntityMeta struct {
	EntityId *string `json:"entityId,omitempty"`
	EntityTitle *string `json:"entityTitle,omitempty"`
}

// NewAiThreadsOpenOrCreateRequestEntityMeta instantiates a new AiThreadsOpenOrCreateRequestEntityMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadsOpenOrCreateRequestEntityMeta() *AiThreadsOpenOrCreateRequestEntityMeta {
	this := AiThreadsOpenOrCreateRequestEntityMeta{}
	return &this
}

// NewAiThreadsOpenOrCreateRequestEntityMetaWithDefaults instantiates a new AiThreadsOpenOrCreateRequestEntityMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadsOpenOrCreateRequestEntityMetaWithDefaults() *AiThreadsOpenOrCreateRequestEntityMeta {
	this := AiThreadsOpenOrCreateRequestEntityMeta{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiThreadsOpenOrCreateRequestEntityMeta) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadsOpenOrCreateRequestEntityMeta) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiThreadsOpenOrCreateRequestEntityMeta) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiThreadsOpenOrCreateRequestEntityMeta) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityTitle returns the EntityTitle field value if set, zero value otherwise.
func (o *AiThreadsOpenOrCreateRequestEntityMeta) GetEntityTitle() string {
	if o == nil || IsNil(o.EntityTitle) {
		var ret string
		return ret
	}
	return *o.EntityTitle
}

// GetEntityTitleOk returns a tuple with the EntityTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadsOpenOrCreateRequestEntityMeta) GetEntityTitleOk() (*string, bool) {
	if o == nil || IsNil(o.EntityTitle) {
		return nil, false
	}
	return o.EntityTitle, true
}

// HasEntityTitle returns a boolean if a field has been set.
func (o *AiThreadsOpenOrCreateRequestEntityMeta) IsEntityTitleSet() bool {
	if o != nil && !IsNil(o.EntityTitle) {
		return true
	}

	return false
}

// SetEntityTitle gets a reference to the given string and assigns it to the EntityTitle field.
func (o *AiThreadsOpenOrCreateRequestEntityMeta) SetEntityTitle(v string) {
	o.EntityTitle = &v
}

func (o AiThreadsOpenOrCreateRequestEntityMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadsOpenOrCreateRequestEntityMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.EntityTitle) {
		toSerialize["entityTitle"] = o.EntityTitle
	}
	return toSerialize, nil
}

type NullableAiThreadsOpenOrCreateRequestEntityMeta struct {
	value *AiThreadsOpenOrCreateRequestEntityMeta
	isSet bool
}

func (v NullableAiThreadsOpenOrCreateRequestEntityMeta) Get() *AiThreadsOpenOrCreateRequestEntityMeta {
	return v.value
}

func (v *NullableAiThreadsOpenOrCreateRequestEntityMeta) Set(val *AiThreadsOpenOrCreateRequestEntityMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadsOpenOrCreateRequestEntityMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadsOpenOrCreateRequestEntityMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadsOpenOrCreateRequestEntityMeta(val *AiThreadsOpenOrCreateRequestEntityMeta) *NullableAiThreadsOpenOrCreateRequestEntityMeta {
	return &NullableAiThreadsOpenOrCreateRequestEntityMeta{value: val, isSet: true}
}

func (v NullableAiThreadsOpenOrCreateRequestEntityMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadsOpenOrCreateRequestEntityMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

