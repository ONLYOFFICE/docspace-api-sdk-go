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

// checks if the ConvertPdfRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertPdfRequest{}

// ConvertPdfRequest The settings of a PDF result.
type ConvertPdfRequest struct {
	// Whether the result is a fillable PDF form: true keeps the form fields of the source as fields to fill in,  false saves a plain PDF.
	Form *bool `json:"form,omitempty"`
}

// NewConvertPdfRequest instantiates a new ConvertPdfRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertPdfRequest() *ConvertPdfRequest {
	this := ConvertPdfRequest{}
	return &this
}

// NewConvertPdfRequestWithDefaults instantiates a new ConvertPdfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertPdfRequestWithDefaults() *ConvertPdfRequest {
	this := ConvertPdfRequest{}
	return &this
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *ConvertPdfRequest) GetForm() bool {
	if o == nil || IsNil(o.Form) {
		var ret bool
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertPdfRequest) GetFormOk() (*bool, bool) {
	if o == nil || IsNil(o.Form) {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *ConvertPdfRequest) IsFormSet() bool {
	if o != nil && !IsNil(o.Form) {
		return true
	}

	return false
}

// SetForm gets a reference to the given bool and assigns it to the Form field.
func (o *ConvertPdfRequest) SetForm(v bool) {
	o.Form = &v
}

func (o ConvertPdfRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertPdfRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Form) {
		toSerialize["form"] = o.Form
	}
	return toSerialize, nil
}

type NullableConvertPdfRequest struct {
	value *ConvertPdfRequest
	isSet bool
}

func (v NullableConvertPdfRequest) Get() *ConvertPdfRequest {
	return v.value
}

func (v *NullableConvertPdfRequest) Set(val *ConvertPdfRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertPdfRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertPdfRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertPdfRequest(val *ConvertPdfRequest) *NullableConvertPdfRequest {
	return &NullableConvertPdfRequest{value: val, isSet: true}
}

func (v NullableConvertPdfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertPdfRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

