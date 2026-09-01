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

// checks if the AiSuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSuccessResponse{}

// AiSuccessResponse Generic success acknowledgement for mutations that return no data.
type AiSuccessResponse struct {
	// Always true — the mutation completed.
	Success bool `json:"success"`
}

type _AiSuccessResponse AiSuccessResponse

// NewAiSuccessResponse instantiates a new AiSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSuccessResponse(success bool) *AiSuccessResponse {
	this := AiSuccessResponse{}
	this.Success = success
	return &this
}

// NewAiSuccessResponseWithDefaults instantiates a new AiSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSuccessResponseWithDefaults() *AiSuccessResponse {
	this := AiSuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiSuccessResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiSuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiSuccessResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o AiSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *AiSuccessResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAiSuccessResponse := _AiSuccessResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiSuccessResponse)

	if err != nil {
		return err
	}

	*o = AiSuccessResponse(varAiSuccessResponse)

	return err
}

type NullableAiSuccessResponse struct {
	value *AiSuccessResponse
	isSet bool
}

func (v NullableAiSuccessResponse) Get() *AiSuccessResponse {
	return v.value
}

func (v *NullableAiSuccessResponse) Set(val *AiSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSuccessResponse(val *AiSuccessResponse) *NullableAiSuccessResponse {
	return &NullableAiSuccessResponse{value: val, isSet: true}
}

func (v NullableAiSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

