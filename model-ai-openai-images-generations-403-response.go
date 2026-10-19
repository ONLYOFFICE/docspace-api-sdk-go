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

// checks if the AiOpenaiImagesGenerations403Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenaiImagesGenerations403Response{}

// AiOpenaiImagesGenerations403Response struct for AiOpenaiImagesGenerations403Response
type AiOpenaiImagesGenerations403Response struct {
	Error AiOpenaiImagesGenerations403ResponseError `json:"error"`
}

type _AiOpenaiImagesGenerations403Response AiOpenaiImagesGenerations403Response

// NewAiOpenaiImagesGenerations403Response instantiates a new AiOpenaiImagesGenerations403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenaiImagesGenerations403Response(error_ AiOpenaiImagesGenerations403ResponseError) *AiOpenaiImagesGenerations403Response {
	this := AiOpenaiImagesGenerations403Response{}
	this.Error = error_
	return &this
}

// NewAiOpenaiImagesGenerations403ResponseWithDefaults instantiates a new AiOpenaiImagesGenerations403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenaiImagesGenerations403ResponseWithDefaults() *AiOpenaiImagesGenerations403Response {
	this := AiOpenaiImagesGenerations403Response{}
	return &this
}

// GetError returns the Error field value
func (o *AiOpenaiImagesGenerations403Response) GetError() AiOpenaiImagesGenerations403ResponseError {
	if o == nil {
		var ret AiOpenaiImagesGenerations403ResponseError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AiOpenaiImagesGenerations403Response) GetErrorOk() (*AiOpenaiImagesGenerations403ResponseError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AiOpenaiImagesGenerations403Response) SetError(v AiOpenaiImagesGenerations403ResponseError) {
	o.Error = v
}

func (o AiOpenaiImagesGenerations403Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenaiImagesGenerations403Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *AiOpenaiImagesGenerations403Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varAiOpenaiImagesGenerations403Response := _AiOpenaiImagesGenerations403Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiOpenaiImagesGenerations403Response)

	if err != nil {
		return err
	}

	*o = AiOpenaiImagesGenerations403Response(varAiOpenaiImagesGenerations403Response)

	return err
}

type NullableAiOpenaiImagesGenerations403Response struct {
	value *AiOpenaiImagesGenerations403Response
	isSet bool
}

func (v NullableAiOpenaiImagesGenerations403Response) Get() *AiOpenaiImagesGenerations403Response {
	return v.value
}

func (v *NullableAiOpenaiImagesGenerations403Response) Set(val *AiOpenaiImagesGenerations403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenaiImagesGenerations403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenaiImagesGenerations403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenaiImagesGenerations403Response(val *AiOpenaiImagesGenerations403Response) *NullableAiOpenaiImagesGenerations403Response {
	return &NullableAiOpenaiImagesGenerations403Response{value: val, isSet: true}
}

func (v NullableAiOpenaiImagesGenerations403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenaiImagesGenerations403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

