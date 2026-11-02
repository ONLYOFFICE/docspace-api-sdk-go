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
)

// checks if the AiProjectModelDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProjectModelDto{}

// AiProjectModelDto The AI model selected for an AI project's chat.
type AiProjectModelDto struct {
	// ID of the AI model profile selected for the project, as listed by `GET api/2.0/ai/profiles/list`. Null when no model has been selected for the project, or the selected profile no longer exists.
	ProfileId NullableString `json:"profileId,omitempty"`
}

// NewAiProjectModelDto instantiates a new AiProjectModelDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProjectModelDto() *AiProjectModelDto {
	this := AiProjectModelDto{}
	return &this
}

// NewAiProjectModelDtoWithDefaults instantiates a new AiProjectModelDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProjectModelDtoWithDefaults() *AiProjectModelDto {
	this := AiProjectModelDto{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiProjectModelDto) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileId.Get()
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiProjectModelDto) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileId.Get(), o.ProfileId.IsSet()
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiProjectModelDto) IsProfileIdSet() bool {
	if o != nil && o.ProfileId.IsSet() {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given NullableString and assigns it to the ProfileId field.
func (o *AiProjectModelDto) SetProfileId(v string) {
	o.ProfileId.Set(&v)
}
// SetProfileIdNil sets the value for ProfileId to be an explicit nil
func (o *AiProjectModelDto) SetProfileIdNil() {
	o.ProfileId.Set(nil)
}

// UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
func (o *AiProjectModelDto) UnsetProfileId() {
	o.ProfileId.Unset()
}

func (o AiProjectModelDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProjectModelDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProfileId.IsSet() {
		toSerialize["profileId"] = o.ProfileId.Get()
	}
	return toSerialize, nil
}

type NullableAiProjectModelDto struct {
	value *AiProjectModelDto
	isSet bool
}

func (v NullableAiProjectModelDto) Get() *AiProjectModelDto {
	return v.value
}

func (v *NullableAiProjectModelDto) Set(val *AiProjectModelDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProjectModelDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProjectModelDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProjectModelDto(val *AiProjectModelDto) *NullableAiProjectModelDto {
	return &NullableAiProjectModelDto{value: val, isSet: true}
}

func (v NullableAiProjectModelDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProjectModelDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

