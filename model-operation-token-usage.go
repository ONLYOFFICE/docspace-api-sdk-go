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

// checks if the OperationTokenUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationTokenUsage{}

// OperationTokenUsage Tokens an AI operation consumed, as recorded in the operation metadata. A kind the provider did not report is `0`.
type OperationTokenUsage struct {
	// All tokens of the request: prompt plus completion.
	TotalTokens *int64 `json:"totalTokens,omitempty"`
	// Tokens sent to the model, cached ones included.
	PromptTokens *int64 `json:"promptTokens,omitempty"`
	// Tokens the model generated, reasoning ones included.
	CompletionTokens *int64 `json:"completionTokens,omitempty"`
	// Part of the prompt tokens read from the provider cache.
	CachedTokens *int64 `json:"cachedTokens,omitempty"`
	// Part of the prompt tokens written to the provider cache.
	CacheWriteTokens *int64 `json:"cacheWriteTokens,omitempty"`
	// Part of the completion tokens the model spent on reasoning.
	ReasoningTokens *int64 `json:"reasoningTokens,omitempty"`
	// Tokens spent on images.
	ImageTokens *int64 `json:"imageTokens,omitempty"`
}

// NewOperationTokenUsage instantiates a new OperationTokenUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationTokenUsage() *OperationTokenUsage {
	this := OperationTokenUsage{}
	return &this
}

// NewOperationTokenUsageWithDefaults instantiates a new OperationTokenUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationTokenUsageWithDefaults() *OperationTokenUsage {
	this := OperationTokenUsage{}
	return &this
}

// GetTotalTokens returns the TotalTokens field value if set, zero value otherwise.
func (o *OperationTokenUsage) GetTotalTokens() int64 {
	if o == nil || IsNil(o.TotalTokens) {
		var ret int64
		return ret
	}
	return *o.TotalTokens
}

// GetTotalTokensOk returns a tuple with the TotalTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationTokenUsage) GetTotalTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalTokens) {
		return nil, false
	}
	return o.TotalTokens, true
}

// HasTotalTokens returns a boolean if a field has been set.
func (o *OperationTokenUsage) IsTotalTokensSet() bool {
	if o != nil && !IsNil(o.TotalTokens) {
		return true
	}

	return false
}

// SetTotalTokens gets a reference to the given int64 and assigns it to the TotalTokens field.
func (o *OperationTokenUsage) SetTotalTokens(v int64) {
	o.TotalTokens = &v
}

// GetPromptTokens returns the PromptTokens field value if set, zero value otherwise.
func (o *OperationTokenUsage) GetPromptTokens() int64 {
	if o == nil || IsNil(o.PromptTokens) {
		var ret int64
		return ret
	}
	return *o.PromptTokens
}

// GetPromptTokensOk returns a tuple with the PromptTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationTokenUsage) GetPromptTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.PromptTokens) {
		return nil, false
	}
	return o.PromptTokens, true
}

// HasPromptTokens returns a boolean if a field has been set.
func (o *OperationTokenUsage) IsPromptTokensSet() bool {
	if o != nil && !IsNil(o.PromptTokens) {
		return true
	}

	return false
}

// SetPromptTokens gets a reference to the given int64 and assigns it to the PromptTokens field.
func (o *OperationTokenUsage) SetPromptTokens(v int64) {
	o.PromptTokens = &v
}

// GetCompletionTokens returns the CompletionTokens field value if set, zero value otherwise.
func (o *OperationTokenUsage) GetCompletionTokens() int64 {
	if o == nil || IsNil(o.CompletionTokens) {
		var ret int64
		return ret
	}
	return *o.CompletionTokens
}

// GetCompletionTokensOk returns a tuple with the CompletionTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationTokenUsage) GetCompletionTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.CompletionTokens) {
		return nil, false
	}
	return o.CompletionTokens, true
}

// HasCompletionTokens returns a boolean if a field has been set.
func (o *OperationTokenUsage) IsCompletionTokensSet() bool {
	if o != nil && !IsNil(o.CompletionTokens) {
		return true
	}

	return false
}

// SetCompletionTokens gets a reference to the given int64 and assigns it to the CompletionTokens field.
func (o *OperationTokenUsage) SetCompletionTokens(v int64) {
	o.CompletionTokens = &v
}

// GetCachedTokens returns the CachedTokens field value if set, zero value otherwise.
func (o *OperationTokenUsage) GetCachedTokens() int64 {
	if o == nil || IsNil(o.CachedTokens) {
		var ret int64
		return ret
	}
	return *o.CachedTokens
}

// GetCachedTokensOk returns a tuple with the CachedTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationTokenUsage) GetCachedTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.CachedTokens) {
		return nil, false
	}
	return o.CachedTokens, true
}

// HasCachedTokens returns a boolean if a field has been set.
func (o *OperationTokenUsage) IsCachedTokensSet() bool {
	if o != nil && !IsNil(o.CachedTokens) {
		return true
	}

	return false
}

// SetCachedTokens gets a reference to the given int64 and assigns it to the CachedTokens field.
func (o *OperationTokenUsage) SetCachedTokens(v int64) {
	o.CachedTokens = &v
}

// GetCacheWriteTokens returns the CacheWriteTokens field value if set, zero value otherwise.
func (o *OperationTokenUsage) GetCacheWriteTokens() int64 {
	if o == nil || IsNil(o.CacheWriteTokens) {
		var ret int64
		return ret
	}
	return *o.CacheWriteTokens
}

// GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationTokenUsage) GetCacheWriteTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.CacheWriteTokens) {
		return nil, false
	}
	return o.CacheWriteTokens, true
}

// HasCacheWriteTokens returns a boolean if a field has been set.
func (o *OperationTokenUsage) IsCacheWriteTokensSet() bool {
	if o != nil && !IsNil(o.CacheWriteTokens) {
		return true
	}

	return false
}

// SetCacheWriteTokens gets a reference to the given int64 and assigns it to the CacheWriteTokens field.
func (o *OperationTokenUsage) SetCacheWriteTokens(v int64) {
	o.CacheWriteTokens = &v
}

// GetReasoningTokens returns the ReasoningTokens field value if set, zero value otherwise.
func (o *OperationTokenUsage) GetReasoningTokens() int64 {
	if o == nil || IsNil(o.ReasoningTokens) {
		var ret int64
		return ret
	}
	return *o.ReasoningTokens
}

// GetReasoningTokensOk returns a tuple with the ReasoningTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationTokenUsage) GetReasoningTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.ReasoningTokens) {
		return nil, false
	}
	return o.ReasoningTokens, true
}

// HasReasoningTokens returns a boolean if a field has been set.
func (o *OperationTokenUsage) IsReasoningTokensSet() bool {
	if o != nil && !IsNil(o.ReasoningTokens) {
		return true
	}

	return false
}

// SetReasoningTokens gets a reference to the given int64 and assigns it to the ReasoningTokens field.
func (o *OperationTokenUsage) SetReasoningTokens(v int64) {
	o.ReasoningTokens = &v
}

// GetImageTokens returns the ImageTokens field value if set, zero value otherwise.
func (o *OperationTokenUsage) GetImageTokens() int64 {
	if o == nil || IsNil(o.ImageTokens) {
		var ret int64
		return ret
	}
	return *o.ImageTokens
}

// GetImageTokensOk returns a tuple with the ImageTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationTokenUsage) GetImageTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.ImageTokens) {
		return nil, false
	}
	return o.ImageTokens, true
}

// HasImageTokens returns a boolean if a field has been set.
func (o *OperationTokenUsage) IsImageTokensSet() bool {
	if o != nil && !IsNil(o.ImageTokens) {
		return true
	}

	return false
}

// SetImageTokens gets a reference to the given int64 and assigns it to the ImageTokens field.
func (o *OperationTokenUsage) SetImageTokens(v int64) {
	o.ImageTokens = &v
}

func (o OperationTokenUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationTokenUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalTokens) {
		toSerialize["totalTokens"] = o.TotalTokens
	}
	if !IsNil(o.PromptTokens) {
		toSerialize["promptTokens"] = o.PromptTokens
	}
	if !IsNil(o.CompletionTokens) {
		toSerialize["completionTokens"] = o.CompletionTokens
	}
	if !IsNil(o.CachedTokens) {
		toSerialize["cachedTokens"] = o.CachedTokens
	}
	if !IsNil(o.CacheWriteTokens) {
		toSerialize["cacheWriteTokens"] = o.CacheWriteTokens
	}
	if !IsNil(o.ReasoningTokens) {
		toSerialize["reasoningTokens"] = o.ReasoningTokens
	}
	if !IsNil(o.ImageTokens) {
		toSerialize["imageTokens"] = o.ImageTokens
	}
	return toSerialize, nil
}

type NullableOperationTokenUsage struct {
	value *OperationTokenUsage
	isSet bool
}

func (v NullableOperationTokenUsage) Get() *OperationTokenUsage {
	return v.value
}

func (v *NullableOperationTokenUsage) Set(val *OperationTokenUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationTokenUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationTokenUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationTokenUsage(val *OperationTokenUsage) *NullableOperationTokenUsage {
	return &NullableOperationTokenUsage{value: val, isSet: true}
}

func (v NullableOperationTokenUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationTokenUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

