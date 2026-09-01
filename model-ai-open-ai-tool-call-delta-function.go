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

// checks if the AiOpenAIToolCallDeltaFunction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenAIToolCallDeltaFunction{}

// AiOpenAIToolCallDeltaFunction The call itself: the function name and its JSON-encoded arguments.
type AiOpenAIToolCallDeltaFunction struct {
	Name *string `json:"name,omitempty"`
	Arguments *string `json:"arguments,omitempty"`
}

// NewAiOpenAIToolCallDeltaFunction instantiates a new AiOpenAIToolCallDeltaFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenAIToolCallDeltaFunction() *AiOpenAIToolCallDeltaFunction {
	this := AiOpenAIToolCallDeltaFunction{}
	return &this
}

// NewAiOpenAIToolCallDeltaFunctionWithDefaults instantiates a new AiOpenAIToolCallDeltaFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenAIToolCallDeltaFunctionWithDefaults() *AiOpenAIToolCallDeltaFunction {
	this := AiOpenAIToolCallDeltaFunction{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiOpenAIToolCallDeltaFunction) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiOpenAIToolCallDeltaFunction) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiOpenAIToolCallDeltaFunction) IsNameSet() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiOpenAIToolCallDeltaFunction) SetName(v string) {
	o.Name = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *AiOpenAIToolCallDeltaFunction) GetArguments() string {
	if o == nil || IsNil(o.Arguments) {
		var ret string
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiOpenAIToolCallDeltaFunction) GetArgumentsOk() (*string, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *AiOpenAIToolCallDeltaFunction) IsArgumentsSet() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given string and assigns it to the Arguments field.
func (o *AiOpenAIToolCallDeltaFunction) SetArguments(v string) {
	o.Arguments = &v
}

func (o AiOpenAIToolCallDeltaFunction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenAIToolCallDeltaFunction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	return toSerialize, nil
}

type NullableAiOpenAIToolCallDeltaFunction struct {
	value *AiOpenAIToolCallDeltaFunction
	isSet bool
}

func (v NullableAiOpenAIToolCallDeltaFunction) Get() *AiOpenAIToolCallDeltaFunction {
	return v.value
}

func (v *NullableAiOpenAIToolCallDeltaFunction) Set(val *AiOpenAIToolCallDeltaFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenAIToolCallDeltaFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenAIToolCallDeltaFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenAIToolCallDeltaFunction(val *AiOpenAIToolCallDeltaFunction) *NullableAiOpenAIToolCallDeltaFunction {
	return &NullableAiOpenAIToolCallDeltaFunction{value: val, isSet: true}
}

func (v NullableAiOpenAIToolCallDeltaFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenAIToolCallDeltaFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

