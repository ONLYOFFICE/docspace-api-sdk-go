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

// checks if the AiBulkAssignmentResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiBulkAssignmentResult{}

// AiBulkAssignmentResult Outcome of `AssignmentsEngine.bulkAssign`. Either every entry persisted, or no entries persisted and a per-key error report. The engine validates first and writes second so a single bad entry never leaves the assignment table in a half-written state.
type AiBulkAssignmentResult struct {
	// True when every entry was persisted.
	Success bool `json:"success"`
	// What was rejected, per action. Present on failure - and then no entry was persisted.
	Errors []AiBulkAssignmentResultErrorsInner `json:"errors,omitempty"`
}

type _AiBulkAssignmentResult AiBulkAssignmentResult

// NewAiBulkAssignmentResult instantiates a new AiBulkAssignmentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiBulkAssignmentResult(success bool) *AiBulkAssignmentResult {
	this := AiBulkAssignmentResult{}
	this.Success = success
	return &this
}

// NewAiBulkAssignmentResultWithDefaults instantiates a new AiBulkAssignmentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiBulkAssignmentResultWithDefaults() *AiBulkAssignmentResult {
	this := AiBulkAssignmentResult{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiBulkAssignmentResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiBulkAssignmentResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiBulkAssignmentResult) SetSuccess(v bool) {
	o.Success = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AiBulkAssignmentResult) GetErrors() []AiBulkAssignmentResultErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []AiBulkAssignmentResultErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiBulkAssignmentResult) GetErrorsOk() ([]AiBulkAssignmentResultErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AiBulkAssignmentResult) IsErrorsSet() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AiBulkAssignmentResultErrorsInner and assigns it to the Errors field.
func (o *AiBulkAssignmentResult) SetErrors(v []AiBulkAssignmentResultErrorsInner) {
	o.Errors = v
}

func (o AiBulkAssignmentResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiBulkAssignmentResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

func (o *AiBulkAssignmentResult) UnmarshalJSON(data []byte) (err error) {
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

	varAiBulkAssignmentResult := _AiBulkAssignmentResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiBulkAssignmentResult)

	if err != nil {
		return err
	}

	*o = AiBulkAssignmentResult(varAiBulkAssignmentResult)

	return err
}

type NullableAiBulkAssignmentResult struct {
	value *AiBulkAssignmentResult
	isSet bool
}

func (v NullableAiBulkAssignmentResult) Get() *AiBulkAssignmentResult {
	return v.value
}

func (v *NullableAiBulkAssignmentResult) Set(val *AiBulkAssignmentResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiBulkAssignmentResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiBulkAssignmentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiBulkAssignmentResult(val *AiBulkAssignmentResult) *NullableAiBulkAssignmentResult {
	return &NullableAiBulkAssignmentResult{value: val, isSet: true}
}

func (v NullableAiBulkAssignmentResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiBulkAssignmentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

