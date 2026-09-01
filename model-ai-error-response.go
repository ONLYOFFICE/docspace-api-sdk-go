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

// checks if the AiErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiErrorResponse{}

// AiErrorResponse Error body — a single human-readable message.
type AiErrorResponse struct {
	// The error message, ready to be shown to the caller.
	Error string `json:"error"`
}

type _AiErrorResponse AiErrorResponse

// NewAiErrorResponse instantiates a new AiErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiErrorResponse(error_ string) *AiErrorResponse {
	this := AiErrorResponse{}
	this.Error = error_
	return &this
}

// NewAiErrorResponseWithDefaults instantiates a new AiErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiErrorResponseWithDefaults() *AiErrorResponse {
	this := AiErrorResponse{}
	return &this
}

// GetError returns the Error field value
func (o *AiErrorResponse) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AiErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AiErrorResponse) SetError(v string) {
	o.Error = v
}

func (o AiErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *AiErrorResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAiErrorResponse := _AiErrorResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiErrorResponse)

	if err != nil {
		return err
	}

	*o = AiErrorResponse(varAiErrorResponse)

	return err
}

type NullableAiErrorResponse struct {
	value *AiErrorResponse
	isSet bool
}

func (v NullableAiErrorResponse) Get() *AiErrorResponse {
	return v.value
}

func (v *NullableAiErrorResponse) Set(val *AiErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAiErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAiErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiErrorResponse(val *AiErrorResponse) *NullableAiErrorResponse {
	return &NullableAiErrorResponse{value: val, isSet: true}
}

func (v NullableAiErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

