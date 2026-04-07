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

// checks if the ToolDecisionRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolDecisionRequestBody{}

// ToolDecisionRequestBody Parameters for the tool execution permission decision.
type ToolDecisionRequestBody struct {
	Decision *ToolExecutionDecision `json:"decision,omitempty"`
}

// NewToolDecisionRequestBody instantiates a new ToolDecisionRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolDecisionRequestBody() *ToolDecisionRequestBody {
	this := ToolDecisionRequestBody{}
	return &this
}

// NewToolDecisionRequestBodyWithDefaults instantiates a new ToolDecisionRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolDecisionRequestBodyWithDefaults() *ToolDecisionRequestBody {
	this := ToolDecisionRequestBody{}
	return &this
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *ToolDecisionRequestBody) GetDecision() ToolExecutionDecision {
	if o == nil || IsNil(o.Decision) {
		var ret ToolExecutionDecision
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolDecisionRequestBody) GetDecisionOk() (*ToolExecutionDecision, bool) {
	if o == nil || IsNil(o.Decision) {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *ToolDecisionRequestBody) IsDecisionSet() bool {
	if o != nil && !IsNil(o.Decision) {
		return true
	}

	return false
}

// SetDecision gets a reference to the given ToolExecutionDecision and assigns it to the Decision field.
func (o *ToolDecisionRequestBody) SetDecision(v ToolExecutionDecision) {
	o.Decision = &v
}

func (o ToolDecisionRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolDecisionRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Decision) {
		toSerialize["decision"] = o.Decision
	}
	return toSerialize, nil
}

type NullableToolDecisionRequestBody struct {
	value *ToolDecisionRequestBody
	isSet bool
}

func (v NullableToolDecisionRequestBody) Get() *ToolDecisionRequestBody {
	return v.value
}

func (v *NullableToolDecisionRequestBody) Set(val *ToolDecisionRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableToolDecisionRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableToolDecisionRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolDecisionRequestBody(val *ToolDecisionRequestBody) *NullableToolDecisionRequestBody {
	return &NullableToolDecisionRequestBody{value: val, isSet: true}
}

func (v NullableToolDecisionRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolDecisionRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

