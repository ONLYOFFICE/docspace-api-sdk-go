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

// checks if the AiVectorizationStartTaskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiVectorizationStartTaskRequest{}

// AiVectorizationStartTaskRequest struct for AiVectorizationStartTaskRequest
type AiVectorizationStartTaskRequest struct {
	// Identifiers of the files to vectorize.
	Files []int32 `json:"files"`
}

type _AiVectorizationStartTaskRequest AiVectorizationStartTaskRequest

// NewAiVectorizationStartTaskRequest instantiates a new AiVectorizationStartTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiVectorizationStartTaskRequest(files []int32) *AiVectorizationStartTaskRequest {
	this := AiVectorizationStartTaskRequest{}
	this.Files = files
	return &this
}

// NewAiVectorizationStartTaskRequestWithDefaults instantiates a new AiVectorizationStartTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiVectorizationStartTaskRequestWithDefaults() *AiVectorizationStartTaskRequest {
	this := AiVectorizationStartTaskRequest{}
	return &this
}

// GetFiles returns the Files field value
func (o *AiVectorizationStartTaskRequest) GetFiles() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *AiVectorizationStartTaskRequest) GetFilesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *AiVectorizationStartTaskRequest) SetFiles(v []int32) {
	o.Files = v
}

func (o AiVectorizationStartTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiVectorizationStartTaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["files"] = o.Files
	return toSerialize, nil
}

func (o *AiVectorizationStartTaskRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"files",
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

	varAiVectorizationStartTaskRequest := _AiVectorizationStartTaskRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiVectorizationStartTaskRequest)

	if err != nil {
		return err
	}

	*o = AiVectorizationStartTaskRequest(varAiVectorizationStartTaskRequest)

	return err
}

type NullableAiVectorizationStartTaskRequest struct {
	value *AiVectorizationStartTaskRequest
	isSet bool
}

func (v NullableAiVectorizationStartTaskRequest) Get() *AiVectorizationStartTaskRequest {
	return v.value
}

func (v *NullableAiVectorizationStartTaskRequest) Set(val *AiVectorizationStartTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiVectorizationStartTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiVectorizationStartTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiVectorizationStartTaskRequest(val *AiVectorizationStartTaskRequest) *NullableAiVectorizationStartTaskRequest {
	return &NullableAiVectorizationStartTaskRequest{value: val, isSet: true}
}

func (v NullableAiVectorizationStartTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiVectorizationStartTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

