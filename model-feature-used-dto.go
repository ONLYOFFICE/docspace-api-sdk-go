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

// checks if the FeatureUsedDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureUsedDto{}

// FeatureUsedDto How much of one quota feature the portal has already consumed.
type FeatureUsedDto struct {
	Value interface{} `json:"value"`
	// The same figure as a sentence in the portal language, ready to print. It is empty when this build ships no  wording for the feature.
	Title NullableString `json:"title,omitempty"`
}

type _FeatureUsedDto FeatureUsedDto

// NewFeatureUsedDto instantiates a new FeatureUsedDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureUsedDto(value interface{}) *FeatureUsedDto {
	this := FeatureUsedDto{}
	this.Value = value
	return &this
}

// NewFeatureUsedDtoWithDefaults instantiates a new FeatureUsedDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureUsedDtoWithDefaults() *FeatureUsedDto {
	this := FeatureUsedDto{}
	return &this
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureUsedDto) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureUsedDto) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FeatureUsedDto) SetValue(v interface{}) {
	o.Value = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureUsedDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureUsedDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FeatureUsedDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FeatureUsedDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FeatureUsedDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FeatureUsedDto) UnsetTitle() {
	o.Title.Unset()
}

func (o FeatureUsedDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureUsedDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

func (o *FeatureUsedDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varFeatureUsedDto := _FeatureUsedDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFeatureUsedDto)

	if err != nil {
		return err
	}

	*o = FeatureUsedDto(varFeatureUsedDto)

	return err
}

type NullableFeatureUsedDto struct {
	value *FeatureUsedDto
	isSet bool
}

func (v NullableFeatureUsedDto) Get() *FeatureUsedDto {
	return v.value
}

func (v *NullableFeatureUsedDto) Set(val *FeatureUsedDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureUsedDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureUsedDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureUsedDto(val *FeatureUsedDto) *NullableFeatureUsedDto {
	return &NullableFeatureUsedDto{value: val, isSet: true}
}

func (v NullableFeatureUsedDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureUsedDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

