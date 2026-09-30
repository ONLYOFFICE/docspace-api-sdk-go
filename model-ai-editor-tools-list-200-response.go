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

// checks if the AiEditorToolsList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEditorToolsList200Response{}

// AiEditorToolsList200Response struct for AiEditorToolsList200Response
type AiEditorToolsList200Response struct {
	// The tools the editor may offer, flattened across every server.
	Tools []AiEditorToolsList200ResponseToolsInner `json:"tools"`
}

type _AiEditorToolsList200Response AiEditorToolsList200Response

// NewAiEditorToolsList200Response instantiates a new AiEditorToolsList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEditorToolsList200Response(tools []AiEditorToolsList200ResponseToolsInner) *AiEditorToolsList200Response {
	this := AiEditorToolsList200Response{}
	this.Tools = tools
	return &this
}

// NewAiEditorToolsList200ResponseWithDefaults instantiates a new AiEditorToolsList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEditorToolsList200ResponseWithDefaults() *AiEditorToolsList200Response {
	this := AiEditorToolsList200Response{}
	return &this
}

// GetTools returns the Tools field value
func (o *AiEditorToolsList200Response) GetTools() []AiEditorToolsList200ResponseToolsInner {
	if o == nil {
		var ret []AiEditorToolsList200ResponseToolsInner
		return ret
	}

	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value
// and a boolean to check if the value has been set.
func (o *AiEditorToolsList200Response) GetToolsOk() ([]AiEditorToolsList200ResponseToolsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tools, true
}

// SetTools sets field value
func (o *AiEditorToolsList200Response) SetTools(v []AiEditorToolsList200ResponseToolsInner) {
	o.Tools = v
}

func (o AiEditorToolsList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEditorToolsList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tools"] = o.Tools
	return toSerialize, nil
}

func (o *AiEditorToolsList200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tools",
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

	varAiEditorToolsList200Response := _AiEditorToolsList200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEditorToolsList200Response)

	if err != nil {
		return err
	}

	*o = AiEditorToolsList200Response(varAiEditorToolsList200Response)

	return err
}

type NullableAiEditorToolsList200Response struct {
	value *AiEditorToolsList200Response
	isSet bool
}

func (v NullableAiEditorToolsList200Response) Get() *AiEditorToolsList200Response {
	return v.value
}

func (v *NullableAiEditorToolsList200Response) Set(val *AiEditorToolsList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEditorToolsList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEditorToolsList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEditorToolsList200Response(val *AiEditorToolsList200Response) *NullableAiEditorToolsList200Response {
	return &NullableAiEditorToolsList200Response{value: val, isSet: true}
}

func (v NullableAiEditorToolsList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEditorToolsList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

