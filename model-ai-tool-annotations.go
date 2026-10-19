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

// checks if the AiToolAnnotations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolAnnotations{}

// AiToolAnnotations MCP tool annotations (`Tool.annotations` in the protocol). All hints are advisory and optional; the protocol's defaults are `readOnlyHint: false` and `destructiveHint: true`, which is why an unannotated tool is treated as one that may destroy state.
type AiToolAnnotations struct {
	Title *string `json:"title,omitempty"`
	ReadOnlyHint *bool `json:"readOnlyHint,omitempty"`
	DestructiveHint *bool `json:"destructiveHint,omitempty"`
	IdempotentHint *bool `json:"idempotentHint,omitempty"`
	OpenWorldHint *bool `json:"openWorldHint,omitempty"`
}

// NewAiToolAnnotations instantiates a new AiToolAnnotations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolAnnotations() *AiToolAnnotations {
	this := AiToolAnnotations{}
	return &this
}

// NewAiToolAnnotationsWithDefaults instantiates a new AiToolAnnotations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolAnnotationsWithDefaults() *AiToolAnnotations {
	this := AiToolAnnotations{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiToolAnnotations) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolAnnotations) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiToolAnnotations) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiToolAnnotations) SetTitle(v string) {
	o.Title = &v
}

// GetReadOnlyHint returns the ReadOnlyHint field value if set, zero value otherwise.
func (o *AiToolAnnotations) GetReadOnlyHint() bool {
	if o == nil || IsNil(o.ReadOnlyHint) {
		var ret bool
		return ret
	}
	return *o.ReadOnlyHint
}

// GetReadOnlyHintOk returns a tuple with the ReadOnlyHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolAnnotations) GetReadOnlyHintOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnlyHint) {
		return nil, false
	}
	return o.ReadOnlyHint, true
}

// HasReadOnlyHint returns a boolean if a field has been set.
func (o *AiToolAnnotations) IsReadOnlyHintSet() bool {
	if o != nil && !IsNil(o.ReadOnlyHint) {
		return true
	}

	return false
}

// SetReadOnlyHint gets a reference to the given bool and assigns it to the ReadOnlyHint field.
func (o *AiToolAnnotations) SetReadOnlyHint(v bool) {
	o.ReadOnlyHint = &v
}

// GetDestructiveHint returns the DestructiveHint field value if set, zero value otherwise.
func (o *AiToolAnnotations) GetDestructiveHint() bool {
	if o == nil || IsNil(o.DestructiveHint) {
		var ret bool
		return ret
	}
	return *o.DestructiveHint
}

// GetDestructiveHintOk returns a tuple with the DestructiveHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolAnnotations) GetDestructiveHintOk() (*bool, bool) {
	if o == nil || IsNil(o.DestructiveHint) {
		return nil, false
	}
	return o.DestructiveHint, true
}

// HasDestructiveHint returns a boolean if a field has been set.
func (o *AiToolAnnotations) IsDestructiveHintSet() bool {
	if o != nil && !IsNil(o.DestructiveHint) {
		return true
	}

	return false
}

// SetDestructiveHint gets a reference to the given bool and assigns it to the DestructiveHint field.
func (o *AiToolAnnotations) SetDestructiveHint(v bool) {
	o.DestructiveHint = &v
}

// GetIdempotentHint returns the IdempotentHint field value if set, zero value otherwise.
func (o *AiToolAnnotations) GetIdempotentHint() bool {
	if o == nil || IsNil(o.IdempotentHint) {
		var ret bool
		return ret
	}
	return *o.IdempotentHint
}

// GetIdempotentHintOk returns a tuple with the IdempotentHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolAnnotations) GetIdempotentHintOk() (*bool, bool) {
	if o == nil || IsNil(o.IdempotentHint) {
		return nil, false
	}
	return o.IdempotentHint, true
}

// HasIdempotentHint returns a boolean if a field has been set.
func (o *AiToolAnnotations) IsIdempotentHintSet() bool {
	if o != nil && !IsNil(o.IdempotentHint) {
		return true
	}

	return false
}

// SetIdempotentHint gets a reference to the given bool and assigns it to the IdempotentHint field.
func (o *AiToolAnnotations) SetIdempotentHint(v bool) {
	o.IdempotentHint = &v
}

// GetOpenWorldHint returns the OpenWorldHint field value if set, zero value otherwise.
func (o *AiToolAnnotations) GetOpenWorldHint() bool {
	if o == nil || IsNil(o.OpenWorldHint) {
		var ret bool
		return ret
	}
	return *o.OpenWorldHint
}

// GetOpenWorldHintOk returns a tuple with the OpenWorldHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolAnnotations) GetOpenWorldHintOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenWorldHint) {
		return nil, false
	}
	return o.OpenWorldHint, true
}

// HasOpenWorldHint returns a boolean if a field has been set.
func (o *AiToolAnnotations) IsOpenWorldHintSet() bool {
	if o != nil && !IsNil(o.OpenWorldHint) {
		return true
	}

	return false
}

// SetOpenWorldHint gets a reference to the given bool and assigns it to the OpenWorldHint field.
func (o *AiToolAnnotations) SetOpenWorldHint(v bool) {
	o.OpenWorldHint = &v
}

func (o AiToolAnnotations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolAnnotations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.ReadOnlyHint) {
		toSerialize["readOnlyHint"] = o.ReadOnlyHint
	}
	if !IsNil(o.DestructiveHint) {
		toSerialize["destructiveHint"] = o.DestructiveHint
	}
	if !IsNil(o.IdempotentHint) {
		toSerialize["idempotentHint"] = o.IdempotentHint
	}
	if !IsNil(o.OpenWorldHint) {
		toSerialize["openWorldHint"] = o.OpenWorldHint
	}
	return toSerialize, nil
}

type NullableAiToolAnnotations struct {
	value *AiToolAnnotations
	isSet bool
}

func (v NullableAiToolAnnotations) Get() *AiToolAnnotations {
	return v.value
}

func (v *NullableAiToolAnnotations) Set(val *AiToolAnnotations) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolAnnotations) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolAnnotations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolAnnotations(val *AiToolAnnotations) *NullableAiToolAnnotations {
	return &NullableAiToolAnnotations{value: val, isSet: true}
}

func (v NullableAiToolAnnotations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolAnnotations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

