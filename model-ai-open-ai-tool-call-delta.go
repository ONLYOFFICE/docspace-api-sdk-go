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

// checks if the AiOpenAIToolCallDelta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenAIToolCallDelta{}

// AiOpenAIToolCallDelta The incremental part of one tool call the model requested.
type AiOpenAIToolCallDelta struct {
	// The zero-based position of the tool call within the message.
	Index float32 `json:"index"`
	// The tool call identifier, quoted back when its result is submitted.
	Id *string `json:"id,omitempty"`
	// Always `function` - the only tool kind the API defines.
	Type *string `json:"type,omitempty"`
	Function *AiOpenAIToolCallDeltaFunction `json:"function,omitempty"`
}

type _AiOpenAIToolCallDelta AiOpenAIToolCallDelta

// NewAiOpenAIToolCallDelta instantiates a new AiOpenAIToolCallDelta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenAIToolCallDelta(index float32) *AiOpenAIToolCallDelta {
	this := AiOpenAIToolCallDelta{}
	this.Index = index
	return &this
}

// NewAiOpenAIToolCallDeltaWithDefaults instantiates a new AiOpenAIToolCallDelta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenAIToolCallDeltaWithDefaults() *AiOpenAIToolCallDelta {
	this := AiOpenAIToolCallDelta{}
	return &this
}

// GetIndex returns the Index field value
func (o *AiOpenAIToolCallDelta) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIToolCallDelta) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *AiOpenAIToolCallDelta) SetIndex(v float32) {
	o.Index = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiOpenAIToolCallDelta) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiOpenAIToolCallDelta) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiOpenAIToolCallDelta) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AiOpenAIToolCallDelta) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AiOpenAIToolCallDelta) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiOpenAIToolCallDelta) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AiOpenAIToolCallDelta) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AiOpenAIToolCallDelta) SetType(v string) {
	o.Type = &v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *AiOpenAIToolCallDelta) GetFunction() AiOpenAIToolCallDeltaFunction {
	if o == nil || IsNil(o.Function) {
		var ret AiOpenAIToolCallDeltaFunction
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiOpenAIToolCallDelta) GetFunctionOk() (*AiOpenAIToolCallDeltaFunction, bool) {
	if o == nil || IsNil(o.Function) {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *AiOpenAIToolCallDelta) IsFunctionSet() bool {
	if o != nil && !IsNil(o.Function) {
		return true
	}

	return false
}

// SetFunction gets a reference to the given AiOpenAIToolCallDeltaFunction and assigns it to the Function field.
func (o *AiOpenAIToolCallDelta) SetFunction(v AiOpenAIToolCallDeltaFunction) {
	o.Function = &v
}

func (o AiOpenAIToolCallDelta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenAIToolCallDelta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Function) {
		toSerialize["function"] = o.Function
	}
	return toSerialize, nil
}

func (o *AiOpenAIToolCallDelta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
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

	varAiOpenAIToolCallDelta := _AiOpenAIToolCallDelta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiOpenAIToolCallDelta)

	if err != nil {
		return err
	}

	*o = AiOpenAIToolCallDelta(varAiOpenAIToolCallDelta)

	return err
}

type NullableAiOpenAIToolCallDelta struct {
	value *AiOpenAIToolCallDelta
	isSet bool
}

func (v NullableAiOpenAIToolCallDelta) Get() *AiOpenAIToolCallDelta {
	return v.value
}

func (v *NullableAiOpenAIToolCallDelta) Set(val *AiOpenAIToolCallDelta) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenAIToolCallDelta) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenAIToolCallDelta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenAIToolCallDelta(val *AiOpenAIToolCallDelta) *NullableAiOpenAIToolCallDelta {
	return &NullableAiOpenAIToolCallDelta{value: val, isSet: true}
}

func (v NullableAiOpenAIToolCallDelta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenAIToolCallDelta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

