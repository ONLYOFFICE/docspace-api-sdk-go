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

// checks if the ConvertSpreadsheetMarginsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertSpreadsheetMarginsRequest{}

// ConvertSpreadsheetMarginsRequest The margins of the output PDF file, each a length with its unit: mm, cm or in.
type ConvertSpreadsheetMarginsRequest struct {
	// The left margin of the output PDF file.
	Left NullableString `json:"left,omitempty"`
	// The right margin of the output PDF file.
	Right NullableString `json:"right,omitempty"`
	// The top margin of the output PDF file.
	Top NullableString `json:"top,omitempty"`
	// The bottom margin of the output PDF file.
	Bottom NullableString `json:"bottom,omitempty"`
}

// NewConvertSpreadsheetMarginsRequest instantiates a new ConvertSpreadsheetMarginsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertSpreadsheetMarginsRequest() *ConvertSpreadsheetMarginsRequest {
	this := ConvertSpreadsheetMarginsRequest{}
	return &this
}

// NewConvertSpreadsheetMarginsRequestWithDefaults instantiates a new ConvertSpreadsheetMarginsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertSpreadsheetMarginsRequestWithDefaults() *ConvertSpreadsheetMarginsRequest {
	this := ConvertSpreadsheetMarginsRequest{}
	return &this
}

// GetLeft returns the Left field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetMarginsRequest) GetLeft() string {
	if o == nil || IsNil(o.Left.Get()) {
		var ret string
		return ret
	}
	return *o.Left.Get()
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetMarginsRequest) GetLeftOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Left.Get(), o.Left.IsSet()
}

// HasLeft returns a boolean if a field has been set.
func (o *ConvertSpreadsheetMarginsRequest) IsLeftSet() bool {
	if o != nil && o.Left.IsSet() {
		return true
	}

	return false
}

// SetLeft gets a reference to the given NullableString and assigns it to the Left field.
func (o *ConvertSpreadsheetMarginsRequest) SetLeft(v string) {
	o.Left.Set(&v)
}
// SetLeftNil sets the value for Left to be an explicit nil
func (o *ConvertSpreadsheetMarginsRequest) SetLeftNil() {
	o.Left.Set(nil)
}

// UnsetLeft ensures that no value is present for Left, not even an explicit nil
func (o *ConvertSpreadsheetMarginsRequest) UnsetLeft() {
	o.Left.Unset()
}

// GetRight returns the Right field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetMarginsRequest) GetRight() string {
	if o == nil || IsNil(o.Right.Get()) {
		var ret string
		return ret
	}
	return *o.Right.Get()
}

// GetRightOk returns a tuple with the Right field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetMarginsRequest) GetRightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Right.Get(), o.Right.IsSet()
}

// HasRight returns a boolean if a field has been set.
func (o *ConvertSpreadsheetMarginsRequest) IsRightSet() bool {
	if o != nil && o.Right.IsSet() {
		return true
	}

	return false
}

// SetRight gets a reference to the given NullableString and assigns it to the Right field.
func (o *ConvertSpreadsheetMarginsRequest) SetRight(v string) {
	o.Right.Set(&v)
}
// SetRightNil sets the value for Right to be an explicit nil
func (o *ConvertSpreadsheetMarginsRequest) SetRightNil() {
	o.Right.Set(nil)
}

// UnsetRight ensures that no value is present for Right, not even an explicit nil
func (o *ConvertSpreadsheetMarginsRequest) UnsetRight() {
	o.Right.Unset()
}

// GetTop returns the Top field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetMarginsRequest) GetTop() string {
	if o == nil || IsNil(o.Top.Get()) {
		var ret string
		return ret
	}
	return *o.Top.Get()
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetMarginsRequest) GetTopOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Top.Get(), o.Top.IsSet()
}

// HasTop returns a boolean if a field has been set.
func (o *ConvertSpreadsheetMarginsRequest) IsTopSet() bool {
	if o != nil && o.Top.IsSet() {
		return true
	}

	return false
}

// SetTop gets a reference to the given NullableString and assigns it to the Top field.
func (o *ConvertSpreadsheetMarginsRequest) SetTop(v string) {
	o.Top.Set(&v)
}
// SetTopNil sets the value for Top to be an explicit nil
func (o *ConvertSpreadsheetMarginsRequest) SetTopNil() {
	o.Top.Set(nil)
}

// UnsetTop ensures that no value is present for Top, not even an explicit nil
func (o *ConvertSpreadsheetMarginsRequest) UnsetTop() {
	o.Top.Unset()
}

// GetBottom returns the Bottom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetMarginsRequest) GetBottom() string {
	if o == nil || IsNil(o.Bottom.Get()) {
		var ret string
		return ret
	}
	return *o.Bottom.Get()
}

// GetBottomOk returns a tuple with the Bottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetMarginsRequest) GetBottomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bottom.Get(), o.Bottom.IsSet()
}

// HasBottom returns a boolean if a field has been set.
func (o *ConvertSpreadsheetMarginsRequest) IsBottomSet() bool {
	if o != nil && o.Bottom.IsSet() {
		return true
	}

	return false
}

// SetBottom gets a reference to the given NullableString and assigns it to the Bottom field.
func (o *ConvertSpreadsheetMarginsRequest) SetBottom(v string) {
	o.Bottom.Set(&v)
}
// SetBottomNil sets the value for Bottom to be an explicit nil
func (o *ConvertSpreadsheetMarginsRequest) SetBottomNil() {
	o.Bottom.Set(nil)
}

// UnsetBottom ensures that no value is present for Bottom, not even an explicit nil
func (o *ConvertSpreadsheetMarginsRequest) UnsetBottom() {
	o.Bottom.Unset()
}

func (o ConvertSpreadsheetMarginsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertSpreadsheetMarginsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Left.IsSet() {
		toSerialize["left"] = o.Left.Get()
	}
	if o.Right.IsSet() {
		toSerialize["right"] = o.Right.Get()
	}
	if o.Top.IsSet() {
		toSerialize["top"] = o.Top.Get()
	}
	if o.Bottom.IsSet() {
		toSerialize["bottom"] = o.Bottom.Get()
	}
	return toSerialize, nil
}

type NullableConvertSpreadsheetMarginsRequest struct {
	value *ConvertSpreadsheetMarginsRequest
	isSet bool
}

func (v NullableConvertSpreadsheetMarginsRequest) Get() *ConvertSpreadsheetMarginsRequest {
	return v.value
}

func (v *NullableConvertSpreadsheetMarginsRequest) Set(val *ConvertSpreadsheetMarginsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertSpreadsheetMarginsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertSpreadsheetMarginsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertSpreadsheetMarginsRequest(val *ConvertSpreadsheetMarginsRequest) *NullableConvertSpreadsheetMarginsRequest {
	return &NullableConvertSpreadsheetMarginsRequest{value: val, isSet: true}
}

func (v NullableConvertSpreadsheetMarginsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertSpreadsheetMarginsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

