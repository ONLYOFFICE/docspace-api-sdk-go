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

// checks if the AiExportTextToDocx200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiExportTextToDocx200Response{}

// AiExportTextToDocx200Response Accepted-for-processing acknowledgement (conversion is asynchronous).
type AiExportTextToDocx200Response struct {
	Success bool `json:"success"`
}

type _AiExportTextToDocx200Response AiExportTextToDocx200Response

// NewAiExportTextToDocx200Response instantiates a new AiExportTextToDocx200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiExportTextToDocx200Response(success bool) *AiExportTextToDocx200Response {
	this := AiExportTextToDocx200Response{}
	this.Success = success
	return &this
}

// NewAiExportTextToDocx200ResponseWithDefaults instantiates a new AiExportTextToDocx200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiExportTextToDocx200ResponseWithDefaults() *AiExportTextToDocx200Response {
	this := AiExportTextToDocx200Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiExportTextToDocx200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiExportTextToDocx200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiExportTextToDocx200Response) SetSuccess(v bool) {
	o.Success = v
}

func (o AiExportTextToDocx200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiExportTextToDocx200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *AiExportTextToDocx200Response) UnmarshalJSON(data []byte) (err error) {
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

	varAiExportTextToDocx200Response := _AiExportTextToDocx200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiExportTextToDocx200Response)

	if err != nil {
		return err
	}

	*o = AiExportTextToDocx200Response(varAiExportTextToDocx200Response)

	return err
}

type NullableAiExportTextToDocx200Response struct {
	value *AiExportTextToDocx200Response
	isSet bool
}

func (v NullableAiExportTextToDocx200Response) Get() *AiExportTextToDocx200Response {
	return v.value
}

func (v *NullableAiExportTextToDocx200Response) Set(val *AiExportTextToDocx200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiExportTextToDocx200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiExportTextToDocx200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiExportTextToDocx200Response(val *AiExportTextToDocx200Response) *NullableAiExportTextToDocx200Response {
	return &NullableAiExportTextToDocx200Response{value: val, isSet: true}
}

func (v NullableAiExportTextToDocx200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiExportTextToDocx200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

