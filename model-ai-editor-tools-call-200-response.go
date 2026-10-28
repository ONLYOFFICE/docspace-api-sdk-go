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

// checks if the AiEditorToolsCall200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEditorToolsCall200Response{}

// AiEditorToolsCall200Response struct for AiEditorToolsCall200Response
type AiEditorToolsCall200Response struct {
	// What the tool produced, as text. A structured result is JSON-encoded, and a tool that failed reports its error here rather than through a status code.
	Result string `json:"result"`
}

type _AiEditorToolsCall200Response AiEditorToolsCall200Response

// NewAiEditorToolsCall200Response instantiates a new AiEditorToolsCall200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEditorToolsCall200Response(result string) *AiEditorToolsCall200Response {
	this := AiEditorToolsCall200Response{}
	this.Result = result
	return &this
}

// NewAiEditorToolsCall200ResponseWithDefaults instantiates a new AiEditorToolsCall200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEditorToolsCall200ResponseWithDefaults() *AiEditorToolsCall200Response {
	this := AiEditorToolsCall200Response{}
	return &this
}

// GetResult returns the Result field value
func (o *AiEditorToolsCall200Response) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *AiEditorToolsCall200Response) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *AiEditorToolsCall200Response) SetResult(v string) {
	o.Result = v
}

func (o AiEditorToolsCall200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEditorToolsCall200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *AiEditorToolsCall200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varAiEditorToolsCall200Response := _AiEditorToolsCall200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEditorToolsCall200Response)

	if err != nil {
		return err
	}

	*o = AiEditorToolsCall200Response(varAiEditorToolsCall200Response)

	return err
}

type NullableAiEditorToolsCall200Response struct {
	value *AiEditorToolsCall200Response
	isSet bool
}

func (v NullableAiEditorToolsCall200Response) Get() *AiEditorToolsCall200Response {
	return v.value
}

func (v *NullableAiEditorToolsCall200Response) Set(val *AiEditorToolsCall200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEditorToolsCall200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEditorToolsCall200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEditorToolsCall200Response(val *AiEditorToolsCall200Response) *NullableAiEditorToolsCall200Response {
	return &NullableAiEditorToolsCall200Response{value: val, isSet: true}
}

func (v NullableAiEditorToolsCall200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEditorToolsCall200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

