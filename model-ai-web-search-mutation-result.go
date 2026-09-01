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

// checks if the AiWebSearchMutationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiWebSearchMutationResult{}

// AiWebSearchMutationResult Outcome of `WebSearchEngine.configure` — either the persisted config or a field-scoped error suitable for the settings form.
type AiWebSearchMutationResult struct {
	// True when the configuration was persisted.
	Success bool `json:"success"`
	// The persisted web-search configuration. Present on success.
	Config *AiWebSearchConfig `json:"config,omitempty"`
	// Why the configuration was rejected. Present on failure.
	Error *AiTErrorData `json:"error,omitempty"`
}

type _AiWebSearchMutationResult AiWebSearchMutationResult

// NewAiWebSearchMutationResult instantiates a new AiWebSearchMutationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiWebSearchMutationResult(success bool) *AiWebSearchMutationResult {
	this := AiWebSearchMutationResult{}
	this.Success = success
	return &this
}

// NewAiWebSearchMutationResultWithDefaults instantiates a new AiWebSearchMutationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiWebSearchMutationResultWithDefaults() *AiWebSearchMutationResult {
	this := AiWebSearchMutationResult{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiWebSearchMutationResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiWebSearchMutationResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiWebSearchMutationResult) SetSuccess(v bool) {
	o.Success = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AiWebSearchMutationResult) GetConfig() AiWebSearchConfig {
	if o == nil || IsNil(o.Config) {
		var ret AiWebSearchConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchMutationResult) GetConfigOk() (*AiWebSearchConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *AiWebSearchMutationResult) IsConfigSet() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AiWebSearchConfig and assigns it to the Config field.
func (o *AiWebSearchMutationResult) SetConfig(v AiWebSearchConfig) {
	o.Config = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AiWebSearchMutationResult) GetError() AiTErrorData {
	if o == nil || IsNil(o.Error) {
		var ret AiTErrorData
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchMutationResult) GetErrorOk() (*AiTErrorData, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AiWebSearchMutationResult) IsErrorSet() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AiTErrorData and assigns it to the Error field.
func (o *AiWebSearchMutationResult) SetError(v AiTErrorData) {
	o.Error = &v
}

func (o AiWebSearchMutationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiWebSearchMutationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *AiWebSearchMutationResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
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

	varAiWebSearchMutationResult := _AiWebSearchMutationResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiWebSearchMutationResult)

	if err != nil {
		return err
	}

	*o = AiWebSearchMutationResult(varAiWebSearchMutationResult)

	return err
}

type NullableAiWebSearchMutationResult struct {
	value *AiWebSearchMutationResult
	isSet bool
}

func (v NullableAiWebSearchMutationResult) Get() *AiWebSearchMutationResult {
	return v.value
}

func (v *NullableAiWebSearchMutationResult) Set(val *AiWebSearchMutationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiWebSearchMutationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiWebSearchMutationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiWebSearchMutationResult(val *AiWebSearchMutationResult) *NullableAiWebSearchMutationResult {
	return &NullableAiWebSearchMutationResult{value: val, isSet: true}
}

func (v NullableAiWebSearchMutationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiWebSearchMutationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

