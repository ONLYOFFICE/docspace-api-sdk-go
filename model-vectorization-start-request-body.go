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

// checks if the VectorizationStartRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VectorizationStartRequestBody{}

// VectorizationStartRequestBody Parameters for submitting files for vectorization.
type VectorizationStartRequestBody struct {
	// The set of file identifiers to submit for vectorization.
	Files []int32 `json:"files"`
}

type _VectorizationStartRequestBody VectorizationStartRequestBody

// NewVectorizationStartRequestBody instantiates a new VectorizationStartRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVectorizationStartRequestBody(files []int32) *VectorizationStartRequestBody {
	this := VectorizationStartRequestBody{}
	this.Files = files
	return &this
}

// NewVectorizationStartRequestBodyWithDefaults instantiates a new VectorizationStartRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVectorizationStartRequestBodyWithDefaults() *VectorizationStartRequestBody {
	this := VectorizationStartRequestBody{}
	return &this
}

// GetFiles returns the Files field value
// If the value is explicit nil, the zero value for []int32 will be returned
func (o *VectorizationStartRequestBody) GetFiles() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VectorizationStartRequestBody) GetFilesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *VectorizationStartRequestBody) SetFiles(v []int32) {
	o.Files = v
}

func (o VectorizationStartRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VectorizationStartRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return toSerialize, nil
}

func (o *VectorizationStartRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varVectorizationStartRequestBody := _VectorizationStartRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varVectorizationStartRequestBody)

	if err != nil {
		return err
	}

	*o = VectorizationStartRequestBody(varVectorizationStartRequestBody)

	return err
}

type NullableVectorizationStartRequestBody struct {
	value *VectorizationStartRequestBody
	isSet bool
}

func (v NullableVectorizationStartRequestBody) Get() *VectorizationStartRequestBody {
	return v.value
}

func (v *NullableVectorizationStartRequestBody) Set(val *VectorizationStartRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableVectorizationStartRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableVectorizationStartRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVectorizationStartRequestBody(val *VectorizationStartRequestBody) *NullableVectorizationStartRequestBody {
	return &NullableVectorizationStartRequestBody{value: val, isSet: true}
}

func (v NullableVectorizationStartRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVectorizationStartRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

