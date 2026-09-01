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

// checks if the AiProfileMutationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProfileMutationResult{}

// AiProfileMutationResult Outcome of `create` / `update` — either a success carrying the persisted profile, or a failure with a field-level error description from the name check or the provider credential check.
type AiProfileMutationResult struct {
	// True when the profile was persisted.
	Success bool `json:"success"`
	// The persisted profile. Present on success.
	Profile *AiProfile `json:"profile,omitempty"`
	// Why the profile was rejected - the name check or the provider credential check. Present on failure.
	Error *AiTErrorData `json:"error,omitempty"`
}

type _AiProfileMutationResult AiProfileMutationResult

// NewAiProfileMutationResult instantiates a new AiProfileMutationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProfileMutationResult(success bool) *AiProfileMutationResult {
	this := AiProfileMutationResult{}
	this.Success = success
	return &this
}

// NewAiProfileMutationResultWithDefaults instantiates a new AiProfileMutationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProfileMutationResultWithDefaults() *AiProfileMutationResult {
	this := AiProfileMutationResult{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiProfileMutationResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiProfileMutationResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiProfileMutationResult) SetSuccess(v bool) {
	o.Success = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AiProfileMutationResult) GetProfile() AiProfile {
	if o == nil || IsNil(o.Profile) {
		var ret AiProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfileMutationResult) GetProfileOk() (*AiProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AiProfileMutationResult) IsProfileSet() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given AiProfile and assigns it to the Profile field.
func (o *AiProfileMutationResult) SetProfile(v AiProfile) {
	o.Profile = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AiProfileMutationResult) GetError() AiTErrorData {
	if o == nil || IsNil(o.Error) {
		var ret AiTErrorData
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfileMutationResult) GetErrorOk() (*AiTErrorData, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AiProfileMutationResult) IsErrorSet() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AiTErrorData and assigns it to the Error field.
func (o *AiProfileMutationResult) SetError(v AiTErrorData) {
	o.Error = &v
}

func (o AiProfileMutationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProfileMutationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *AiProfileMutationResult) UnmarshalJSON(data []byte) (err error) {
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

	varAiProfileMutationResult := _AiProfileMutationResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProfileMutationResult)

	if err != nil {
		return err
	}

	*o = AiProfileMutationResult(varAiProfileMutationResult)

	return err
}

type NullableAiProfileMutationResult struct {
	value *AiProfileMutationResult
	isSet bool
}

func (v NullableAiProfileMutationResult) Get() *AiProfileMutationResult {
	return v.value
}

func (v *NullableAiProfileMutationResult) Set(val *AiProfileMutationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProfileMutationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProfileMutationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProfileMutationResult(val *AiProfileMutationResult) *NullableAiProfileMutationResult {
	return &NullableAiProfileMutationResult{value: val, isSet: true}
}

func (v NullableAiProfileMutationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProfileMutationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

