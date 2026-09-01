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

// checks if the AiToolsMutationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolsMutationResult{}

// AiToolsMutationResult Outcome of an MCP-server CRUD call. Either success or a field-scoped error suitable for the settings form.
type AiToolsMutationResult struct {
	// True when the MCP server was persisted.
	Success bool `json:"success"`
	// Why the MCP server was rejected. Present on failure.
	Error *AiTErrorData `json:"error,omitempty"`
}

type _AiToolsMutationResult AiToolsMutationResult

// NewAiToolsMutationResult instantiates a new AiToolsMutationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolsMutationResult(success bool) *AiToolsMutationResult {
	this := AiToolsMutationResult{}
	this.Success = success
	return &this
}

// NewAiToolsMutationResultWithDefaults instantiates a new AiToolsMutationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolsMutationResultWithDefaults() *AiToolsMutationResult {
	this := AiToolsMutationResult{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiToolsMutationResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiToolsMutationResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiToolsMutationResult) SetSuccess(v bool) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AiToolsMutationResult) GetError() AiTErrorData {
	if o == nil || IsNil(o.Error) {
		var ret AiTErrorData
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolsMutationResult) GetErrorOk() (*AiTErrorData, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AiToolsMutationResult) IsErrorSet() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AiTErrorData and assigns it to the Error field.
func (o *AiToolsMutationResult) SetError(v AiTErrorData) {
	o.Error = &v
}

func (o AiToolsMutationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolsMutationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *AiToolsMutationResult) UnmarshalJSON(data []byte) (err error) {
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

	varAiToolsMutationResult := _AiToolsMutationResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolsMutationResult)

	if err != nil {
		return err
	}

	*o = AiToolsMutationResult(varAiToolsMutationResult)

	return err
}

type NullableAiToolsMutationResult struct {
	value *AiToolsMutationResult
	isSet bool
}

func (v NullableAiToolsMutationResult) Get() *AiToolsMutationResult {
	return v.value
}

func (v *NullableAiToolsMutationResult) Set(val *AiToolsMutationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolsMutationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolsMutationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolsMutationResult(val *AiToolsMutationResult) *NullableAiToolsMutationResult {
	return &NullableAiToolsMutationResult{value: val, isSet: true}
}

func (v NullableAiToolsMutationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolsMutationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

