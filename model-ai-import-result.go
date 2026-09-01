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

// checks if the AiImportResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiImportResult{}

// AiImportResult Outcome of `PromptsEngine.importBundle`. Either every entry persisted with counts, or no entries persisted plus a per-entry error report.
type AiImportResult struct {
	// True when the whole bundle was imported.
	Success bool `json:"success"`
	Imported *AiImportResultImported `json:"imported,omitempty"`
	// What was rejected, per entry. Present on failure - and then nothing was imported.
	Errors []AiImportError `json:"errors,omitempty"`
}

type _AiImportResult AiImportResult

// NewAiImportResult instantiates a new AiImportResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiImportResult(success bool) *AiImportResult {
	this := AiImportResult{}
	this.Success = success
	return &this
}

// NewAiImportResultWithDefaults instantiates a new AiImportResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiImportResultWithDefaults() *AiImportResult {
	this := AiImportResult{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiImportResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiImportResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiImportResult) SetSuccess(v bool) {
	o.Success = v
}

// GetImported returns the Imported field value if set, zero value otherwise.
func (o *AiImportResult) GetImported() AiImportResultImported {
	if o == nil || IsNil(o.Imported) {
		var ret AiImportResultImported
		return ret
	}
	return *o.Imported
}

// GetImportedOk returns a tuple with the Imported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiImportResult) GetImportedOk() (*AiImportResultImported, bool) {
	if o == nil || IsNil(o.Imported) {
		return nil, false
	}
	return o.Imported, true
}

// HasImported returns a boolean if a field has been set.
func (o *AiImportResult) IsImportedSet() bool {
	if o != nil && !IsNil(o.Imported) {
		return true
	}

	return false
}

// SetImported gets a reference to the given AiImportResultImported and assigns it to the Imported field.
func (o *AiImportResult) SetImported(v AiImportResultImported) {
	o.Imported = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AiImportResult) GetErrors() []AiImportError {
	if o == nil || IsNil(o.Errors) {
		var ret []AiImportError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiImportResult) GetErrorsOk() ([]AiImportError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AiImportResult) IsErrorsSet() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AiImportError and assigns it to the Errors field.
func (o *AiImportResult) SetErrors(v []AiImportError) {
	o.Errors = v
}

func (o AiImportResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiImportResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Imported) {
		toSerialize["imported"] = o.Imported
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

func (o *AiImportResult) UnmarshalJSON(data []byte) (err error) {
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

	varAiImportResult := _AiImportResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiImportResult)

	if err != nil {
		return err
	}

	*o = AiImportResult(varAiImportResult)

	return err
}

type NullableAiImportResult struct {
	value *AiImportResult
	isSet bool
}

func (v NullableAiImportResult) Get() *AiImportResult {
	return v.value
}

func (v *NullableAiImportResult) Set(val *AiImportResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiImportResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiImportResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiImportResult(val *AiImportResult) *NullableAiImportResult {
	return &NullableAiImportResult{value: val, isSet: true}
}

func (v NullableAiImportResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiImportResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

