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

// checks if the AiProjectModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProjectModel{}

// AiProjectModel The project's own model; `null` when the project runs on the portal default.
type AiProjectModel struct {
	Profile NullableAiProjectModelProfile `json:"profile"`
}

type _AiProjectModel AiProjectModel

// NewAiProjectModel instantiates a new AiProjectModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProjectModel(profile NullableAiProjectModelProfile) *AiProjectModel {
	this := AiProjectModel{}
	this.Profile = profile
	return &this
}

// NewAiProjectModelWithDefaults instantiates a new AiProjectModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProjectModelWithDefaults() *AiProjectModel {
	this := AiProjectModel{}
	return &this
}

// GetProfile returns the Profile field value
// If the value is explicit nil, the zero value for AiProjectModelProfile will be returned
func (o *AiProjectModel) GetProfile() AiProjectModelProfile {
	if o == nil || o.Profile.Get() == nil {
		var ret AiProjectModelProfile
		return ret
	}

	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiProjectModel) GetProfileOk() (*AiProjectModelProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// SetProfile sets field value
func (o *AiProjectModel) SetProfile(v AiProjectModelProfile) {
	o.Profile.Set(&v)
}

func (o AiProjectModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProjectModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profile"] = o.Profile.Get()
	return toSerialize, nil
}

func (o *AiProjectModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile",
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

	varAiProjectModel := _AiProjectModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProjectModel)

	if err != nil {
		return err
	}

	*o = AiProjectModel(varAiProjectModel)

	return err
}

type NullableAiProjectModel struct {
	value *AiProjectModel
	isSet bool
}

func (v NullableAiProjectModel) Get() *AiProjectModel {
	return v.value
}

func (v *NullableAiProjectModel) Set(val *AiProjectModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProjectModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProjectModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProjectModel(val *AiProjectModel) *NullableAiProjectModel {
	return &NullableAiProjectModel{value: val, isSet: true}
}

func (v NullableAiProjectModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProjectModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

