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

// checks if the AiFolderMutationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFolderMutationResult{}

// AiFolderMutationResult Outcome of `createFolder` / `renameFolder` — either the persisted folder or a field-scoped error.
type AiFolderMutationResult struct {
	// True when the folder was persisted.
	Success bool `json:"success"`
	// The persisted folder. Present on success.
	Folder *AiPromptFolder `json:"folder,omitempty"`
	// Why the folder was rejected. Present on failure.
	Error *AiTErrorData `json:"error,omitempty"`
}

type _AiFolderMutationResult AiFolderMutationResult

// NewAiFolderMutationResult instantiates a new AiFolderMutationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFolderMutationResult(success bool) *AiFolderMutationResult {
	this := AiFolderMutationResult{}
	this.Success = success
	return &this
}

// NewAiFolderMutationResultWithDefaults instantiates a new AiFolderMutationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFolderMutationResultWithDefaults() *AiFolderMutationResult {
	this := AiFolderMutationResult{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiFolderMutationResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiFolderMutationResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiFolderMutationResult) SetSuccess(v bool) {
	o.Success = v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *AiFolderMutationResult) GetFolder() AiPromptFolder {
	if o == nil || IsNil(o.Folder) {
		var ret AiPromptFolder
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderMutationResult) GetFolderOk() (*AiPromptFolder, bool) {
	if o == nil || IsNil(o.Folder) {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *AiFolderMutationResult) IsFolderSet() bool {
	if o != nil && !IsNil(o.Folder) {
		return true
	}

	return false
}

// SetFolder gets a reference to the given AiPromptFolder and assigns it to the Folder field.
func (o *AiFolderMutationResult) SetFolder(v AiPromptFolder) {
	o.Folder = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AiFolderMutationResult) GetError() AiTErrorData {
	if o == nil || IsNil(o.Error) {
		var ret AiTErrorData
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderMutationResult) GetErrorOk() (*AiTErrorData, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AiFolderMutationResult) IsErrorSet() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AiTErrorData and assigns it to the Error field.
func (o *AiFolderMutationResult) SetError(v AiTErrorData) {
	o.Error = &v
}

func (o AiFolderMutationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFolderMutationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Folder) {
		toSerialize["folder"] = o.Folder
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *AiFolderMutationResult) UnmarshalJSON(data []byte) (err error) {
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

	varAiFolderMutationResult := _AiFolderMutationResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiFolderMutationResult)

	if err != nil {
		return err
	}

	*o = AiFolderMutationResult(varAiFolderMutationResult)

	return err
}

type NullableAiFolderMutationResult struct {
	value *AiFolderMutationResult
	isSet bool
}

func (v NullableAiFolderMutationResult) Get() *AiFolderMutationResult {
	return v.value
}

func (v *NullableAiFolderMutationResult) Set(val *AiFolderMutationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFolderMutationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFolderMutationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFolderMutationResult(val *AiFolderMutationResult) *NullableAiFolderMutationResult {
	return &NullableAiFolderMutationResult{value: val, isSet: true}
}

func (v NullableAiFolderMutationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFolderMutationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

