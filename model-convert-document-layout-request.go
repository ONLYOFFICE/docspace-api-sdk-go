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

// checks if the ConvertDocumentLayoutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertDocumentLayoutRequest{}

// ConvertDocumentLayoutRequest The layout of forms printed as pdf documents or images.
type ConvertDocumentLayoutRequest struct {
	// Whether an empty form field shows its placeholder text in the result; false leaves it blank.
	DrawPlaceHolders NullableBool `json:"drawPlaceHolders,omitempty"`
	// Whether form fields keep their highlight colour in the result; false draws them without it.
	DrawFormHighlight NullableBool `json:"drawFormHighlight,omitempty"`
	// Whether the print mode is turned on. It only applies to a docx converted into pdf: with the print mode off  the highlight flag does nothing and the placeholder flag saves the forms in the pdf.
	IsPrint NullableBool `json:"isPrint,omitempty"`
}

// NewConvertDocumentLayoutRequest instantiates a new ConvertDocumentLayoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertDocumentLayoutRequest() *ConvertDocumentLayoutRequest {
	this := ConvertDocumentLayoutRequest{}
	return &this
}

// NewConvertDocumentLayoutRequestWithDefaults instantiates a new ConvertDocumentLayoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertDocumentLayoutRequestWithDefaults() *ConvertDocumentLayoutRequest {
	this := ConvertDocumentLayoutRequest{}
	return &this
}

// GetDrawPlaceHolders returns the DrawPlaceHolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertDocumentLayoutRequest) GetDrawPlaceHolders() bool {
	if o == nil || IsNil(o.DrawPlaceHolders.Get()) {
		var ret bool
		return ret
	}
	return *o.DrawPlaceHolders.Get()
}

// GetDrawPlaceHoldersOk returns a tuple with the DrawPlaceHolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertDocumentLayoutRequest) GetDrawPlaceHoldersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DrawPlaceHolders.Get(), o.DrawPlaceHolders.IsSet()
}

// HasDrawPlaceHolders returns a boolean if a field has been set.
func (o *ConvertDocumentLayoutRequest) IsDrawPlaceHoldersSet() bool {
	if o != nil && o.DrawPlaceHolders.IsSet() {
		return true
	}

	return false
}

// SetDrawPlaceHolders gets a reference to the given NullableBool and assigns it to the DrawPlaceHolders field.
func (o *ConvertDocumentLayoutRequest) SetDrawPlaceHolders(v bool) {
	o.DrawPlaceHolders.Set(&v)
}
// SetDrawPlaceHoldersNil sets the value for DrawPlaceHolders to be an explicit nil
func (o *ConvertDocumentLayoutRequest) SetDrawPlaceHoldersNil() {
	o.DrawPlaceHolders.Set(nil)
}

// UnsetDrawPlaceHolders ensures that no value is present for DrawPlaceHolders, not even an explicit nil
func (o *ConvertDocumentLayoutRequest) UnsetDrawPlaceHolders() {
	o.DrawPlaceHolders.Unset()
}

// GetDrawFormHighlight returns the DrawFormHighlight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertDocumentLayoutRequest) GetDrawFormHighlight() bool {
	if o == nil || IsNil(o.DrawFormHighlight.Get()) {
		var ret bool
		return ret
	}
	return *o.DrawFormHighlight.Get()
}

// GetDrawFormHighlightOk returns a tuple with the DrawFormHighlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertDocumentLayoutRequest) GetDrawFormHighlightOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DrawFormHighlight.Get(), o.DrawFormHighlight.IsSet()
}

// HasDrawFormHighlight returns a boolean if a field has been set.
func (o *ConvertDocumentLayoutRequest) IsDrawFormHighlightSet() bool {
	if o != nil && o.DrawFormHighlight.IsSet() {
		return true
	}

	return false
}

// SetDrawFormHighlight gets a reference to the given NullableBool and assigns it to the DrawFormHighlight field.
func (o *ConvertDocumentLayoutRequest) SetDrawFormHighlight(v bool) {
	o.DrawFormHighlight.Set(&v)
}
// SetDrawFormHighlightNil sets the value for DrawFormHighlight to be an explicit nil
func (o *ConvertDocumentLayoutRequest) SetDrawFormHighlightNil() {
	o.DrawFormHighlight.Set(nil)
}

// UnsetDrawFormHighlight ensures that no value is present for DrawFormHighlight, not even an explicit nil
func (o *ConvertDocumentLayoutRequest) UnsetDrawFormHighlight() {
	o.DrawFormHighlight.Unset()
}

// GetIsPrint returns the IsPrint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertDocumentLayoutRequest) GetIsPrint() bool {
	if o == nil || IsNil(o.IsPrint.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPrint.Get()
}

// GetIsPrintOk returns a tuple with the IsPrint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertDocumentLayoutRequest) GetIsPrintOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPrint.Get(), o.IsPrint.IsSet()
}

// HasIsPrint returns a boolean if a field has been set.
func (o *ConvertDocumentLayoutRequest) IsIsPrintSet() bool {
	if o != nil && o.IsPrint.IsSet() {
		return true
	}

	return false
}

// SetIsPrint gets a reference to the given NullableBool and assigns it to the IsPrint field.
func (o *ConvertDocumentLayoutRequest) SetIsPrint(v bool) {
	o.IsPrint.Set(&v)
}
// SetIsPrintNil sets the value for IsPrint to be an explicit nil
func (o *ConvertDocumentLayoutRequest) SetIsPrintNil() {
	o.IsPrint.Set(nil)
}

// UnsetIsPrint ensures that no value is present for IsPrint, not even an explicit nil
func (o *ConvertDocumentLayoutRequest) UnsetIsPrint() {
	o.IsPrint.Unset()
}

func (o ConvertDocumentLayoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertDocumentLayoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DrawPlaceHolders.IsSet() {
		toSerialize["drawPlaceHolders"] = o.DrawPlaceHolders.Get()
	}
	if o.DrawFormHighlight.IsSet() {
		toSerialize["drawFormHighlight"] = o.DrawFormHighlight.Get()
	}
	if o.IsPrint.IsSet() {
		toSerialize["isPrint"] = o.IsPrint.Get()
	}
	return toSerialize, nil
}

type NullableConvertDocumentLayoutRequest struct {
	value *ConvertDocumentLayoutRequest
	isSet bool
}

func (v NullableConvertDocumentLayoutRequest) Get() *ConvertDocumentLayoutRequest {
	return v.value
}

func (v *NullableConvertDocumentLayoutRequest) Set(val *ConvertDocumentLayoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertDocumentLayoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertDocumentLayoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertDocumentLayoutRequest(val *ConvertDocumentLayoutRequest) *NullableConvertDocumentLayoutRequest {
	return &NullableConvertDocumentLayoutRequest{value: val, isSet: true}
}

func (v NullableConvertDocumentLayoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertDocumentLayoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

