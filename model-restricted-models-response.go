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

// checks if the RestrictedModelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestrictedModelsResponse{}

// RestrictedModelsResponse The AI models the portal is not allowed to use.
type RestrictedModelsResponse struct {
	// The identifiers of the models the portal is not allowed to use.
	Models []string `json:"models"`
}

type _RestrictedModelsResponse RestrictedModelsResponse

// NewRestrictedModelsResponse instantiates a new RestrictedModelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestrictedModelsResponse(models []string) *RestrictedModelsResponse {
	this := RestrictedModelsResponse{}
	this.Models = models
	return &this
}

// NewRestrictedModelsResponseWithDefaults instantiates a new RestrictedModelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictedModelsResponseWithDefaults() *RestrictedModelsResponse {
	this := RestrictedModelsResponse{}
	return &this
}

// GetModels returns the Models field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *RestrictedModelsResponse) GetModels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestrictedModelsResponse) GetModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *RestrictedModelsResponse) SetModels(v []string) {
	o.Models = v
}

func (o RestrictedModelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestrictedModelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Models != nil {
		toSerialize["models"] = o.Models
	}
	return toSerialize, nil
}

func (o *RestrictedModelsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"models",
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

	varRestrictedModelsResponse := _RestrictedModelsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRestrictedModelsResponse)

	if err != nil {
		return err
	}

	*o = RestrictedModelsResponse(varRestrictedModelsResponse)

	return err
}

type NullableRestrictedModelsResponse struct {
	value *RestrictedModelsResponse
	isSet bool
}

func (v NullableRestrictedModelsResponse) Get() *RestrictedModelsResponse {
	return v.value
}

func (v *NullableRestrictedModelsResponse) Set(val *RestrictedModelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictedModelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictedModelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictedModelsResponse(val *RestrictedModelsResponse) *NullableRestrictedModelsResponse {
	return &NullableRestrictedModelsResponse{value: val, isSet: true}
}

func (v NullableRestrictedModelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictedModelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

