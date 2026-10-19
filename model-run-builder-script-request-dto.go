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

// checks if the RunBuilderScriptRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunBuilderScriptRequestDto{}

// RunBuilderScriptRequestDto The document builder script to run and where to put what it produces.
type RunBuilderScriptRequestDto struct {
	// The document builder script. It names a portal file by its id where the document builder documentation  writes an address - `builder.OpenFile(1234)`; an address written out in the script is refused.
	Script NullableString `json:"script"`
	// The id of the folder for the produced files that `outputs` does not list, as reported by a folder listing  such as `GET api/2.0/files/{folderId}`.
	FolderId NullableInt32 `json:"folderId,omitempty"`
	// Destinations of the produced files, keyed by the name the script saves each one under - the second argument of  `builder.SaveFile(docx, result.docx)`.
	Outputs map[string]RunBuilderScriptOutputRequest `json:"outputs,omitempty"`
	Argument *RunBuilderScriptRequestDtoArgument `json:"argument,omitempty"`
}

type _RunBuilderScriptRequestDto RunBuilderScriptRequestDto

// NewRunBuilderScriptRequestDto instantiates a new RunBuilderScriptRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunBuilderScriptRequestDto(script NullableString) *RunBuilderScriptRequestDto {
	this := RunBuilderScriptRequestDto{}
	this.Script = script
	return &this
}

// NewRunBuilderScriptRequestDtoWithDefaults instantiates a new RunBuilderScriptRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunBuilderScriptRequestDtoWithDefaults() *RunBuilderScriptRequestDto {
	this := RunBuilderScriptRequestDto{}
	return &this
}

// GetScript returns the Script field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RunBuilderScriptRequestDto) GetScript() string {
	if o == nil || o.Script.Get() == nil {
		var ret string
		return ret
	}

	return *o.Script.Get()
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunBuilderScriptRequestDto) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Script.Get(), o.Script.IsSet()
}

// SetScript sets field value
func (o *RunBuilderScriptRequestDto) SetScript(v string) {
	o.Script.Set(&v)
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunBuilderScriptRequestDto) GetFolderId() int32 {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret int32
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunBuilderScriptRequestDto) GetFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *RunBuilderScriptRequestDto) IsFolderIdSet() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableInt32 and assigns it to the FolderId field.
func (o *RunBuilderScriptRequestDto) SetFolderId(v int32) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *RunBuilderScriptRequestDto) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *RunBuilderScriptRequestDto) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *RunBuilderScriptRequestDto) GetOutputs() map[string]RunBuilderScriptOutputRequest {
	if o == nil || IsNil(o.Outputs) {
		var ret map[string]RunBuilderScriptOutputRequest
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunBuilderScriptRequestDto) GetOutputsOk() (map[string]RunBuilderScriptOutputRequest, bool) {
	if o == nil || IsNil(o.Outputs) {
		return map[string]RunBuilderScriptOutputRequest{}, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *RunBuilderScriptRequestDto) IsOutputsSet() bool {
	if o != nil && !IsNil(o.Outputs) {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given map[string]RunBuilderScriptOutputRequest and assigns it to the Outputs field.
func (o *RunBuilderScriptRequestDto) SetOutputs(v map[string]RunBuilderScriptOutputRequest) {
	o.Outputs = v
}

// GetArgument returns the Argument field value if set, zero value otherwise.
func (o *RunBuilderScriptRequestDto) GetArgument() RunBuilderScriptRequestDtoArgument {
	if o == nil || IsNil(o.Argument) {
		var ret RunBuilderScriptRequestDtoArgument
		return ret
	}
	return *o.Argument
}

// GetArgumentOk returns a tuple with the Argument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunBuilderScriptRequestDto) GetArgumentOk() (*RunBuilderScriptRequestDtoArgument, bool) {
	if o == nil || IsNil(o.Argument) {
		return nil, false
	}
	return o.Argument, true
}

// HasArgument returns a boolean if a field has been set.
func (o *RunBuilderScriptRequestDto) IsArgumentSet() bool {
	if o != nil && !IsNil(o.Argument) {
		return true
	}

	return false
}

// SetArgument gets a reference to the given RunBuilderScriptRequestDtoArgument and assigns it to the Argument field.
func (o *RunBuilderScriptRequestDto) SetArgument(v RunBuilderScriptRequestDtoArgument) {
	o.Argument = &v
}

func (o RunBuilderScriptRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunBuilderScriptRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["script"] = o.Script.Get()
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	if !IsNil(o.Outputs) {
		toSerialize["outputs"] = o.Outputs
	}
	if !IsNil(o.Argument) {
		toSerialize["argument"] = o.Argument
	}
	return toSerialize, nil
}

func (o *RunBuilderScriptRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"script",
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

	varRunBuilderScriptRequestDto := _RunBuilderScriptRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRunBuilderScriptRequestDto)

	if err != nil {
		return err
	}

	*o = RunBuilderScriptRequestDto(varRunBuilderScriptRequestDto)

	return err
}

type NullableRunBuilderScriptRequestDto struct {
	value *RunBuilderScriptRequestDto
	isSet bool
}

func (v NullableRunBuilderScriptRequestDto) Get() *RunBuilderScriptRequestDto {
	return v.value
}

func (v *NullableRunBuilderScriptRequestDto) Set(val *RunBuilderScriptRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRunBuilderScriptRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRunBuilderScriptRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunBuilderScriptRequestDto(val *RunBuilderScriptRequestDto) *NullableRunBuilderScriptRequestDto {
	return &NullableRunBuilderScriptRequestDto{value: val, isSet: true}
}

func (v NullableRunBuilderScriptRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunBuilderScriptRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

