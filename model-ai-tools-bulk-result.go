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

// checks if the AiToolsBulkResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolsBulkResult{}

// AiToolsBulkResult Outcome of `ToolsEngine.replaceAllCustomServers` — either every entry persisted, or no entries persisted plus a per-key error report.
type AiToolsBulkResult struct {
	// True when every custom MCP server was persisted.
	Success bool `json:"success"`
	// What was rejected, per server. Present on failure - and then no server was persisted.
	Errors []AiToolsBulkResultErrorsInner `json:"errors,omitempty"`
}

type _AiToolsBulkResult AiToolsBulkResult

// NewAiToolsBulkResult instantiates a new AiToolsBulkResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolsBulkResult(success bool) *AiToolsBulkResult {
	this := AiToolsBulkResult{}
	this.Success = success
	return &this
}

// NewAiToolsBulkResultWithDefaults instantiates a new AiToolsBulkResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolsBulkResultWithDefaults() *AiToolsBulkResult {
	this := AiToolsBulkResult{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiToolsBulkResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiToolsBulkResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiToolsBulkResult) SetSuccess(v bool) {
	o.Success = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AiToolsBulkResult) GetErrors() []AiToolsBulkResultErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []AiToolsBulkResultErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolsBulkResult) GetErrorsOk() ([]AiToolsBulkResultErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AiToolsBulkResult) IsErrorsSet() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AiToolsBulkResultErrorsInner and assigns it to the Errors field.
func (o *AiToolsBulkResult) SetErrors(v []AiToolsBulkResultErrorsInner) {
	o.Errors = v
}

func (o AiToolsBulkResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolsBulkResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

func (o *AiToolsBulkResult) UnmarshalJSON(data []byte) (err error) {
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

	varAiToolsBulkResult := _AiToolsBulkResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolsBulkResult)

	if err != nil {
		return err
	}

	*o = AiToolsBulkResult(varAiToolsBulkResult)

	return err
}

type NullableAiToolsBulkResult struct {
	value *AiToolsBulkResult
	isSet bool
}

func (v NullableAiToolsBulkResult) Get() *AiToolsBulkResult {
	return v.value
}

func (v *NullableAiToolsBulkResult) Set(val *AiToolsBulkResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolsBulkResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolsBulkResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolsBulkResult(val *AiToolsBulkResult) *NullableAiToolsBulkResult {
	return &NullableAiToolsBulkResult{value: val, isSet: true}
}

func (v NullableAiToolsBulkResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolsBulkResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

