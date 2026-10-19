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

// checks if the ReviewConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewConfigDto{}

// ReviewConfigDto How tracked changes are displayed when the document opens.
type ReviewConfigDto struct {
	// How the editors render tracked changes at first: with the markup, in a simplified markup, as the final text,  or as the original text. A session that may not write opens on the final text.
	ReviewDisplay NullableString `json:"reviewDisplay,omitempty"`
}

// NewReviewConfigDto instantiates a new ReviewConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewConfigDto() *ReviewConfigDto {
	this := ReviewConfigDto{}
	return &this
}

// NewReviewConfigDtoWithDefaults instantiates a new ReviewConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewConfigDtoWithDefaults() *ReviewConfigDto {
	this := ReviewConfigDto{}
	return &this
}

// GetReviewDisplay returns the ReviewDisplay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewConfigDto) GetReviewDisplay() string {
	if o == nil || IsNil(o.ReviewDisplay.Get()) {
		var ret string
		return ret
	}
	return *o.ReviewDisplay.Get()
}

// GetReviewDisplayOk returns a tuple with the ReviewDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewConfigDto) GetReviewDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewDisplay.Get(), o.ReviewDisplay.IsSet()
}

// HasReviewDisplay returns a boolean if a field has been set.
func (o *ReviewConfigDto) IsReviewDisplaySet() bool {
	if o != nil && o.ReviewDisplay.IsSet() {
		return true
	}

	return false
}

// SetReviewDisplay gets a reference to the given NullableString and assigns it to the ReviewDisplay field.
func (o *ReviewConfigDto) SetReviewDisplay(v string) {
	o.ReviewDisplay.Set(&v)
}
// SetReviewDisplayNil sets the value for ReviewDisplay to be an explicit nil
func (o *ReviewConfigDto) SetReviewDisplayNil() {
	o.ReviewDisplay.Set(nil)
}

// UnsetReviewDisplay ensures that no value is present for ReviewDisplay, not even an explicit nil
func (o *ReviewConfigDto) UnsetReviewDisplay() {
	o.ReviewDisplay.Unset()
}

func (o ReviewConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReviewDisplay.IsSet() {
		toSerialize["reviewDisplay"] = o.ReviewDisplay.Get()
	}
	return toSerialize, nil
}

type NullableReviewConfigDto struct {
	value *ReviewConfigDto
	isSet bool
}

func (v NullableReviewConfigDto) Get() *ReviewConfigDto {
	return v.value
}

func (v *NullableReviewConfigDto) Set(val *ReviewConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewConfigDto(val *ReviewConfigDto) *NullableReviewConfigDto {
	return &NullableReviewConfigDto{value: val, isSet: true}
}

func (v NullableReviewConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

