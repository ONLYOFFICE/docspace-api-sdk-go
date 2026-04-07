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

// checks if the ReviewConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewConfig{}

// ReviewConfig Configuration for review display settings.
type ReviewConfig struct {
	// The review display string representation.
	ReviewDisplay NullableString `json:"reviewDisplay,omitempty"`
}

// NewReviewConfig instantiates a new ReviewConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewConfig() *ReviewConfig {
	this := ReviewConfig{}
	return &this
}

// NewReviewConfigWithDefaults instantiates a new ReviewConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewConfigWithDefaults() *ReviewConfig {
	this := ReviewConfig{}
	return &this
}

// GetReviewDisplay returns the ReviewDisplay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewConfig) GetReviewDisplay() string {
	if o == nil || IsNil(o.ReviewDisplay.Get()) {
		var ret string
		return ret
	}
	return *o.ReviewDisplay.Get()
}

// GetReviewDisplayOk returns a tuple with the ReviewDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewConfig) GetReviewDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewDisplay.Get(), o.ReviewDisplay.IsSet()
}

// HasReviewDisplay returns a boolean if a field has been set.
func (o *ReviewConfig) IsReviewDisplaySet() bool {
	if o != nil && o.ReviewDisplay.IsSet() {
		return true
	}

	return false
}

// SetReviewDisplay gets a reference to the given NullableString and assigns it to the ReviewDisplay field.
func (o *ReviewConfig) SetReviewDisplay(v string) {
	o.ReviewDisplay.Set(&v)
}
// SetReviewDisplayNil sets the value for ReviewDisplay to be an explicit nil
func (o *ReviewConfig) SetReviewDisplayNil() {
	o.ReviewDisplay.Set(nil)
}

// UnsetReviewDisplay ensures that no value is present for ReviewDisplay, not even an explicit nil
func (o *ReviewConfig) UnsetReviewDisplay() {
	o.ReviewDisplay.Unset()
}

func (o ReviewConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReviewDisplay.IsSet() {
		toSerialize["reviewDisplay"] = o.ReviewDisplay.Get()
	}
	return toSerialize, nil
}

type NullableReviewConfig struct {
	value *ReviewConfig
	isSet bool
}

func (v NullableReviewConfig) Get() *ReviewConfig {
	return v.value
}

func (v *NullableReviewConfig) Set(val *ReviewConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewConfig(val *ReviewConfig) *NullableReviewConfig {
	return &NullableReviewConfig{value: val, isSet: true}
}

func (v NullableReviewConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

