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

// checks if the ConvertDocumentRendererRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertDocumentRendererRequest{}

// ConvertDocumentRendererRequest How the text of a pdf, xps or oxps source is read back when converting from it.
type ConvertDocumentRendererRequest struct {
	// The rendering mode: blockChar reads the text by single characters, blockLine by separate lines, plainLine as  plain text with a paragraph per line, plainParagraph as plain text with lines combined into paragraphs.
	TextAssociation NullableString `json:"textAssociation,omitempty"`
}

// NewConvertDocumentRendererRequest instantiates a new ConvertDocumentRendererRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertDocumentRendererRequest() *ConvertDocumentRendererRequest {
	this := ConvertDocumentRendererRequest{}
	return &this
}

// NewConvertDocumentRendererRequestWithDefaults instantiates a new ConvertDocumentRendererRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertDocumentRendererRequestWithDefaults() *ConvertDocumentRendererRequest {
	this := ConvertDocumentRendererRequest{}
	return &this
}

// GetTextAssociation returns the TextAssociation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertDocumentRendererRequest) GetTextAssociation() string {
	if o == nil || IsNil(o.TextAssociation.Get()) {
		var ret string
		return ret
	}
	return *o.TextAssociation.Get()
}

// GetTextAssociationOk returns a tuple with the TextAssociation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertDocumentRendererRequest) GetTextAssociationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextAssociation.Get(), o.TextAssociation.IsSet()
}

// HasTextAssociation returns a boolean if a field has been set.
func (o *ConvertDocumentRendererRequest) IsTextAssociationSet() bool {
	if o != nil && o.TextAssociation.IsSet() {
		return true
	}

	return false
}

// SetTextAssociation gets a reference to the given NullableString and assigns it to the TextAssociation field.
func (o *ConvertDocumentRendererRequest) SetTextAssociation(v string) {
	o.TextAssociation.Set(&v)
}
// SetTextAssociationNil sets the value for TextAssociation to be an explicit nil
func (o *ConvertDocumentRendererRequest) SetTextAssociationNil() {
	o.TextAssociation.Set(nil)
}

// UnsetTextAssociation ensures that no value is present for TextAssociation, not even an explicit nil
func (o *ConvertDocumentRendererRequest) UnsetTextAssociation() {
	o.TextAssociation.Unset()
}

func (o ConvertDocumentRendererRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertDocumentRendererRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TextAssociation.IsSet() {
		toSerialize["textAssociation"] = o.TextAssociation.Get()
	}
	return toSerialize, nil
}

type NullableConvertDocumentRendererRequest struct {
	value *ConvertDocumentRendererRequest
	isSet bool
}

func (v NullableConvertDocumentRendererRequest) Get() *ConvertDocumentRendererRequest {
	return v.value
}

func (v *NullableConvertDocumentRendererRequest) Set(val *ConvertDocumentRendererRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertDocumentRendererRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertDocumentRendererRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertDocumentRendererRequest(val *ConvertDocumentRendererRequest) *NullableConvertDocumentRendererRequest {
	return &NullableConvertDocumentRendererRequest{value: val, isSet: true}
}

func (v NullableConvertDocumentRendererRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertDocumentRendererRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

