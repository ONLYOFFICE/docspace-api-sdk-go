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

// checks if the ConvertSpreadsheetPageSizeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertSpreadsheetPageSizeRequest{}

// ConvertSpreadsheetPageSizeRequest The page size of the output PDF file, each side a length with its unit: mm, cm or in.
type ConvertSpreadsheetPageSizeRequest struct {
	// The page height of the output PDF file.
	Height NullableString `json:"height,omitempty"`
	// The page width of the output PDF file.
	Width NullableString `json:"width,omitempty"`
}

// NewConvertSpreadsheetPageSizeRequest instantiates a new ConvertSpreadsheetPageSizeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertSpreadsheetPageSizeRequest() *ConvertSpreadsheetPageSizeRequest {
	this := ConvertSpreadsheetPageSizeRequest{}
	return &this
}

// NewConvertSpreadsheetPageSizeRequestWithDefaults instantiates a new ConvertSpreadsheetPageSizeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertSpreadsheetPageSizeRequestWithDefaults() *ConvertSpreadsheetPageSizeRequest {
	this := ConvertSpreadsheetPageSizeRequest{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetPageSizeRequest) GetHeight() string {
	if o == nil || IsNil(o.Height.Get()) {
		var ret string
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetPageSizeRequest) GetHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *ConvertSpreadsheetPageSizeRequest) IsHeightSet() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableString and assigns it to the Height field.
func (o *ConvertSpreadsheetPageSizeRequest) SetHeight(v string) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *ConvertSpreadsheetPageSizeRequest) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *ConvertSpreadsheetPageSizeRequest) UnsetHeight() {
	o.Height.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetPageSizeRequest) GetWidth() string {
	if o == nil || IsNil(o.Width.Get()) {
		var ret string
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetPageSizeRequest) GetWidthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *ConvertSpreadsheetPageSizeRequest) IsWidthSet() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableString and assigns it to the Width field.
func (o *ConvertSpreadsheetPageSizeRequest) SetWidth(v string) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *ConvertSpreadsheetPageSizeRequest) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *ConvertSpreadsheetPageSizeRequest) UnsetWidth() {
	o.Width.Unset()
}

func (o ConvertSpreadsheetPageSizeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertSpreadsheetPageSizeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}
	return toSerialize, nil
}

type NullableConvertSpreadsheetPageSizeRequest struct {
	value *ConvertSpreadsheetPageSizeRequest
	isSet bool
}

func (v NullableConvertSpreadsheetPageSizeRequest) Get() *ConvertSpreadsheetPageSizeRequest {
	return v.value
}

func (v *NullableConvertSpreadsheetPageSizeRequest) Set(val *ConvertSpreadsheetPageSizeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertSpreadsheetPageSizeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertSpreadsheetPageSizeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertSpreadsheetPageSizeRequest(val *ConvertSpreadsheetPageSizeRequest) *NullableConvertSpreadsheetPageSizeRequest {
	return &NullableConvertSpreadsheetPageSizeRequest{value: val, isSet: true}
}

func (v NullableConvertSpreadsheetPageSizeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertSpreadsheetPageSizeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

