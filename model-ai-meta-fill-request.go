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

// checks if the AiMetaFillRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaFillRequest{}

// AiMetaFillRequest struct for AiMetaFillRequest
type AiMetaFillRequest struct {
	// The entry's table. Only `file` for now: a folder or a room answers 400.
	Kind string `json:"kind"`
	Id AiMetaGenerateRequestId `json:"id"`
	// What to pay attention to; a hint for the model.
	Prompt *string `json:"prompt,omitempty"`
	// Also replace the fields that already carry a value. Default false.
	Overwrite *bool `json:"overwrite,omitempty"`
	// Resolve the values but write nothing; the caller writes what the user confirms. Default true: a write needs an explicit false.
	DryRun *bool `json:"dryRun,omitempty"`
}

type _AiMetaFillRequest AiMetaFillRequest

// NewAiMetaFillRequest instantiates a new AiMetaFillRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaFillRequest(kind string, id AiMetaGenerateRequestId) *AiMetaFillRequest {
	this := AiMetaFillRequest{}
	this.Kind = kind
	this.Id = id
	return &this
}

// NewAiMetaFillRequestWithDefaults instantiates a new AiMetaFillRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaFillRequestWithDefaults() *AiMetaFillRequest {
	this := AiMetaFillRequest{}
	return &this
}

// GetKind returns the Kind field value
func (o *AiMetaFillRequest) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AiMetaFillRequest) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AiMetaFillRequest) SetKind(v string) {
	o.Kind = v
}

// GetId returns the Id field value
func (o *AiMetaFillRequest) GetId() AiMetaGenerateRequestId {
	if o == nil {
		var ret AiMetaGenerateRequestId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiMetaFillRequest) GetIdOk() (*AiMetaGenerateRequestId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiMetaFillRequest) SetId(v AiMetaGenerateRequestId) {
	o.Id = v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiMetaFillRequest) GetPrompt() string {
	if o == nil || IsNil(o.Prompt) {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetaFillRequest) GetPromptOk() (*string, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiMetaFillRequest) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *AiMetaFillRequest) SetPrompt(v string) {
	o.Prompt = &v
}

// GetOverwrite returns the Overwrite field value if set, zero value otherwise.
func (o *AiMetaFillRequest) GetOverwrite() bool {
	if o == nil || IsNil(o.Overwrite) {
		var ret bool
		return ret
	}
	return *o.Overwrite
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetaFillRequest) GetOverwriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Overwrite) {
		return nil, false
	}
	return o.Overwrite, true
}

// HasOverwrite returns a boolean if a field has been set.
func (o *AiMetaFillRequest) IsOverwriteSet() bool {
	if o != nil && !IsNil(o.Overwrite) {
		return true
	}

	return false
}

// SetOverwrite gets a reference to the given bool and assigns it to the Overwrite field.
func (o *AiMetaFillRequest) SetOverwrite(v bool) {
	o.Overwrite = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AiMetaFillRequest) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetaFillRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AiMetaFillRequest) IsDryRunSet() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AiMetaFillRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o AiMetaFillRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaFillRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["id"] = o.Id
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.Overwrite) {
		toSerialize["overwrite"] = o.Overwrite
	}
	if !IsNil(o.DryRun) {
		toSerialize["dryRun"] = o.DryRun
	}
	return toSerialize, nil
}

func (o *AiMetaFillRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"id",
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

	varAiMetaFillRequest := _AiMetaFillRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaFillRequest)

	if err != nil {
		return err
	}

	*o = AiMetaFillRequest(varAiMetaFillRequest)

	return err
}

type NullableAiMetaFillRequest struct {
	value *AiMetaFillRequest
	isSet bool
}

func (v NullableAiMetaFillRequest) Get() *AiMetaFillRequest {
	return v.value
}

func (v *NullableAiMetaFillRequest) Set(val *AiMetaFillRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaFillRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaFillRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaFillRequest(val *AiMetaFillRequest) *NullableAiMetaFillRequest {
	return &NullableAiMetaFillRequest{value: val, isSet: true}
}

func (v NullableAiMetaFillRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaFillRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

