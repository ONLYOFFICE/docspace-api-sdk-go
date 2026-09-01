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

// checks if the FormSubmissionsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormSubmissionsDto{}

// FormSubmissionsDto All submissions of a form, together with the metadata of its fields.
type FormSubmissionsDto struct {
	// The form field metadata.
	Metadata []FormMetadata `json:"metadata,omitempty"`
	// All submissions.
	Submissions []FormResultsDto `json:"submissions,omitempty"`
}

// NewFormSubmissionsDto instantiates a new FormSubmissionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormSubmissionsDto() *FormSubmissionsDto {
	this := FormSubmissionsDto{}
	return &this
}

// NewFormSubmissionsDtoWithDefaults instantiates a new FormSubmissionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormSubmissionsDtoWithDefaults() *FormSubmissionsDto {
	this := FormSubmissionsDto{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormSubmissionsDto) GetMetadata() []FormMetadata {
	if o == nil {
		var ret []FormMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormSubmissionsDto) GetMetadataOk() ([]FormMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FormSubmissionsDto) IsMetadataSet() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []FormMetadata and assigns it to the Metadata field.
func (o *FormSubmissionsDto) SetMetadata(v []FormMetadata) {
	o.Metadata = v
}

// GetSubmissions returns the Submissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormSubmissionsDto) GetSubmissions() []FormResultsDto {
	if o == nil {
		var ret []FormResultsDto
		return ret
	}
	return o.Submissions
}

// GetSubmissionsOk returns a tuple with the Submissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormSubmissionsDto) GetSubmissionsOk() ([]FormResultsDto, bool) {
	if o == nil || IsNil(o.Submissions) {
		return nil, false
	}
	return o.Submissions, true
}

// HasSubmissions returns a boolean if a field has been set.
func (o *FormSubmissionsDto) IsSubmissionsSet() bool {
	if o != nil && !IsNil(o.Submissions) {
		return true
	}

	return false
}

// SetSubmissions gets a reference to the given []FormResultsDto and assigns it to the Submissions field.
func (o *FormSubmissionsDto) SetSubmissions(v []FormResultsDto) {
	o.Submissions = v
}

func (o FormSubmissionsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormSubmissionsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Submissions != nil {
		toSerialize["submissions"] = o.Submissions
	}
	return toSerialize, nil
}

type NullableFormSubmissionsDto struct {
	value *FormSubmissionsDto
	isSet bool
}

func (v NullableFormSubmissionsDto) Get() *FormSubmissionsDto {
	return v.value
}

func (v *NullableFormSubmissionsDto) Set(val *FormSubmissionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFormSubmissionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFormSubmissionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormSubmissionsDto(val *FormSubmissionsDto) *NullableFormSubmissionsDto {
	return &NullableFormSubmissionsDto{value: val, isSet: true}
}

func (v NullableFormSubmissionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormSubmissionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

