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

// checks if the AiImportError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiImportError{}

// AiImportError Per-entry error reported by `PromptsEngine.importBundle`.
type AiImportError struct {
	// `folder` or `prompt`, plus the offending name or id.
	Kind string `json:"kind"`
	// The offending entry - its name or its id.
	Ref string `json:"ref"`
	// Why the entry was rejected.
	Error AiTErrorData `json:"error"`
}

type _AiImportError AiImportError

// NewAiImportError instantiates a new AiImportError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiImportError(kind string, ref string, error_ AiTErrorData) *AiImportError {
	this := AiImportError{}
	this.Kind = kind
	this.Ref = ref
	this.Error = error_
	return &this
}

// NewAiImportErrorWithDefaults instantiates a new AiImportError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiImportErrorWithDefaults() *AiImportError {
	this := AiImportError{}
	return &this
}

// GetKind returns the Kind field value
func (o *AiImportError) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AiImportError) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AiImportError) SetKind(v string) {
	o.Kind = v
}

// GetRef returns the Ref field value
func (o *AiImportError) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *AiImportError) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *AiImportError) SetRef(v string) {
	o.Ref = v
}

// GetError returns the Error field value
func (o *AiImportError) GetError() AiTErrorData {
	if o == nil {
		var ret AiTErrorData
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AiImportError) GetErrorOk() (*AiTErrorData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AiImportError) SetError(v AiTErrorData) {
	o.Error = v
}

func (o AiImportError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiImportError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["ref"] = o.Ref
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *AiImportError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"ref",
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

	varAiImportError := _AiImportError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiImportError)

	if err != nil {
		return err
	}

	*o = AiImportError(varAiImportError)

	return err
}

type NullableAiImportError struct {
	value *AiImportError
	isSet bool
}

func (v NullableAiImportError) Get() *AiImportError {
	return v.value
}

func (v *NullableAiImportError) Set(val *AiImportError) {
	v.value = val
	v.isSet = true
}

func (v NullableAiImportError) IsSet() bool {
	return v.isSet
}

func (v *NullableAiImportError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiImportError(val *AiImportError) *NullableAiImportError {
	return &NullableAiImportError{value: val, isSet: true}
}

func (v NullableAiImportError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiImportError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

