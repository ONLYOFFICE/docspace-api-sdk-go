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

// checks if the ConvertSpreadsheetLayoutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertSpreadsheetLayoutRequest{}

// ConvertSpreadsheetLayoutRequest How a spreadsheet converted to PDF is laid out on the page. A field left out takes the document service default.
type ConvertSpreadsheetLayoutRequest struct {
	// Whether the print area chosen in the spreadsheet is ignored; false converts only that area.
	IgnorePrintArea NullableBool `json:"ignorePrintArea,omitempty"`
	// The page orientation of the output PDF file: portrait or landscape.
	Orientation NullableString `json:"orientation,omitempty"`
	// The number of pages the converted area is fitted into in height; 0 does not fit it.
	FitToHeight NullableInt32 `json:"fitToHeight,omitempty"`
	// The scale of the output PDF file, in percent.
	Scale NullableInt32 `json:"scale,omitempty"`
	// The number of pages the converted area is fitted into in width; 0 does not fit it.
	FitToWidth NullableInt32 `json:"fitToWidth,omitempty"`
	// Whether the row and column headings are printed in the output PDF file.
	Headings NullableBool `json:"headings,omitempty"`
	// Whether the grid lines are printed in the output PDF file.
	GridLines NullableBool `json:"gridLines,omitempty"`
	// The margins of the output PDF file.
	Margins *ConvertSpreadsheetMarginsRequest `json:"margins,omitempty"`
	// The page size of the output PDF file.
	PageSize *ConvertSpreadsheetPageSizeRequest `json:"pageSize,omitempty"`
}

// NewConvertSpreadsheetLayoutRequest instantiates a new ConvertSpreadsheetLayoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertSpreadsheetLayoutRequest() *ConvertSpreadsheetLayoutRequest {
	this := ConvertSpreadsheetLayoutRequest{}
	return &this
}

// NewConvertSpreadsheetLayoutRequestWithDefaults instantiates a new ConvertSpreadsheetLayoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertSpreadsheetLayoutRequestWithDefaults() *ConvertSpreadsheetLayoutRequest {
	this := ConvertSpreadsheetLayoutRequest{}
	return &this
}

// GetIgnorePrintArea returns the IgnorePrintArea field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetLayoutRequest) GetIgnorePrintArea() bool {
	if o == nil || IsNil(o.IgnorePrintArea.Get()) {
		var ret bool
		return ret
	}
	return *o.IgnorePrintArea.Get()
}

// GetIgnorePrintAreaOk returns a tuple with the IgnorePrintArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetLayoutRequest) GetIgnorePrintAreaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IgnorePrintArea.Get(), o.IgnorePrintArea.IsSet()
}

// HasIgnorePrintArea returns a boolean if a field has been set.
func (o *ConvertSpreadsheetLayoutRequest) IsIgnorePrintAreaSet() bool {
	if o != nil && o.IgnorePrintArea.IsSet() {
		return true
	}

	return false
}

// SetIgnorePrintArea gets a reference to the given NullableBool and assigns it to the IgnorePrintArea field.
func (o *ConvertSpreadsheetLayoutRequest) SetIgnorePrintArea(v bool) {
	o.IgnorePrintArea.Set(&v)
}
// SetIgnorePrintAreaNil sets the value for IgnorePrintArea to be an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) SetIgnorePrintAreaNil() {
	o.IgnorePrintArea.Set(nil)
}

// UnsetIgnorePrintArea ensures that no value is present for IgnorePrintArea, not even an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) UnsetIgnorePrintArea() {
	o.IgnorePrintArea.Unset()
}

// GetOrientation returns the Orientation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetLayoutRequest) GetOrientation() string {
	if o == nil || IsNil(o.Orientation.Get()) {
		var ret string
		return ret
	}
	return *o.Orientation.Get()
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetLayoutRequest) GetOrientationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Orientation.Get(), o.Orientation.IsSet()
}

// HasOrientation returns a boolean if a field has been set.
func (o *ConvertSpreadsheetLayoutRequest) IsOrientationSet() bool {
	if o != nil && o.Orientation.IsSet() {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given NullableString and assigns it to the Orientation field.
func (o *ConvertSpreadsheetLayoutRequest) SetOrientation(v string) {
	o.Orientation.Set(&v)
}
// SetOrientationNil sets the value for Orientation to be an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) SetOrientationNil() {
	o.Orientation.Set(nil)
}

// UnsetOrientation ensures that no value is present for Orientation, not even an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) UnsetOrientation() {
	o.Orientation.Unset()
}

// GetFitToHeight returns the FitToHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetLayoutRequest) GetFitToHeight() int32 {
	if o == nil || IsNil(o.FitToHeight.Get()) {
		var ret int32
		return ret
	}
	return *o.FitToHeight.Get()
}

// GetFitToHeightOk returns a tuple with the FitToHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetLayoutRequest) GetFitToHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FitToHeight.Get(), o.FitToHeight.IsSet()
}

// HasFitToHeight returns a boolean if a field has been set.
func (o *ConvertSpreadsheetLayoutRequest) IsFitToHeightSet() bool {
	if o != nil && o.FitToHeight.IsSet() {
		return true
	}

	return false
}

// SetFitToHeight gets a reference to the given NullableInt32 and assigns it to the FitToHeight field.
func (o *ConvertSpreadsheetLayoutRequest) SetFitToHeight(v int32) {
	o.FitToHeight.Set(&v)
}
// SetFitToHeightNil sets the value for FitToHeight to be an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) SetFitToHeightNil() {
	o.FitToHeight.Set(nil)
}

// UnsetFitToHeight ensures that no value is present for FitToHeight, not even an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) UnsetFitToHeight() {
	o.FitToHeight.Unset()
}

// GetScale returns the Scale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetLayoutRequest) GetScale() int32 {
	if o == nil || IsNil(o.Scale.Get()) {
		var ret int32
		return ret
	}
	return *o.Scale.Get()
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetLayoutRequest) GetScaleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scale.Get(), o.Scale.IsSet()
}

// HasScale returns a boolean if a field has been set.
func (o *ConvertSpreadsheetLayoutRequest) IsScaleSet() bool {
	if o != nil && o.Scale.IsSet() {
		return true
	}

	return false
}

// SetScale gets a reference to the given NullableInt32 and assigns it to the Scale field.
func (o *ConvertSpreadsheetLayoutRequest) SetScale(v int32) {
	o.Scale.Set(&v)
}
// SetScaleNil sets the value for Scale to be an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) SetScaleNil() {
	o.Scale.Set(nil)
}

// UnsetScale ensures that no value is present for Scale, not even an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) UnsetScale() {
	o.Scale.Unset()
}

// GetFitToWidth returns the FitToWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetLayoutRequest) GetFitToWidth() int32 {
	if o == nil || IsNil(o.FitToWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.FitToWidth.Get()
}

// GetFitToWidthOk returns a tuple with the FitToWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetLayoutRequest) GetFitToWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FitToWidth.Get(), o.FitToWidth.IsSet()
}

// HasFitToWidth returns a boolean if a field has been set.
func (o *ConvertSpreadsheetLayoutRequest) IsFitToWidthSet() bool {
	if o != nil && o.FitToWidth.IsSet() {
		return true
	}

	return false
}

// SetFitToWidth gets a reference to the given NullableInt32 and assigns it to the FitToWidth field.
func (o *ConvertSpreadsheetLayoutRequest) SetFitToWidth(v int32) {
	o.FitToWidth.Set(&v)
}
// SetFitToWidthNil sets the value for FitToWidth to be an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) SetFitToWidthNil() {
	o.FitToWidth.Set(nil)
}

// UnsetFitToWidth ensures that no value is present for FitToWidth, not even an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) UnsetFitToWidth() {
	o.FitToWidth.Unset()
}

// GetHeadings returns the Headings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetLayoutRequest) GetHeadings() bool {
	if o == nil || IsNil(o.Headings.Get()) {
		var ret bool
		return ret
	}
	return *o.Headings.Get()
}

// GetHeadingsOk returns a tuple with the Headings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetLayoutRequest) GetHeadingsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Headings.Get(), o.Headings.IsSet()
}

// HasHeadings returns a boolean if a field has been set.
func (o *ConvertSpreadsheetLayoutRequest) IsHeadingsSet() bool {
	if o != nil && o.Headings.IsSet() {
		return true
	}

	return false
}

// SetHeadings gets a reference to the given NullableBool and assigns it to the Headings field.
func (o *ConvertSpreadsheetLayoutRequest) SetHeadings(v bool) {
	o.Headings.Set(&v)
}
// SetHeadingsNil sets the value for Headings to be an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) SetHeadingsNil() {
	o.Headings.Set(nil)
}

// UnsetHeadings ensures that no value is present for Headings, not even an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) UnsetHeadings() {
	o.Headings.Unset()
}

// GetGridLines returns the GridLines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertSpreadsheetLayoutRequest) GetGridLines() bool {
	if o == nil || IsNil(o.GridLines.Get()) {
		var ret bool
		return ret
	}
	return *o.GridLines.Get()
}

// GetGridLinesOk returns a tuple with the GridLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertSpreadsheetLayoutRequest) GetGridLinesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GridLines.Get(), o.GridLines.IsSet()
}

// HasGridLines returns a boolean if a field has been set.
func (o *ConvertSpreadsheetLayoutRequest) IsGridLinesSet() bool {
	if o != nil && o.GridLines.IsSet() {
		return true
	}

	return false
}

// SetGridLines gets a reference to the given NullableBool and assigns it to the GridLines field.
func (o *ConvertSpreadsheetLayoutRequest) SetGridLines(v bool) {
	o.GridLines.Set(&v)
}
// SetGridLinesNil sets the value for GridLines to be an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) SetGridLinesNil() {
	o.GridLines.Set(nil)
}

// UnsetGridLines ensures that no value is present for GridLines, not even an explicit nil
func (o *ConvertSpreadsheetLayoutRequest) UnsetGridLines() {
	o.GridLines.Unset()
}

// GetMargins returns the Margins field value if set, zero value otherwise.
func (o *ConvertSpreadsheetLayoutRequest) GetMargins() ConvertSpreadsheetMarginsRequest {
	if o == nil || IsNil(o.Margins) {
		var ret ConvertSpreadsheetMarginsRequest
		return ret
	}
	return *o.Margins
}

// GetMarginsOk returns a tuple with the Margins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertSpreadsheetLayoutRequest) GetMarginsOk() (*ConvertSpreadsheetMarginsRequest, bool) {
	if o == nil || IsNil(o.Margins) {
		return nil, false
	}
	return o.Margins, true
}

// HasMargins returns a boolean if a field has been set.
func (o *ConvertSpreadsheetLayoutRequest) IsMarginsSet() bool {
	if o != nil && !IsNil(o.Margins) {
		return true
	}

	return false
}

// SetMargins gets a reference to the given ConvertSpreadsheetMarginsRequest and assigns it to the Margins field.
func (o *ConvertSpreadsheetLayoutRequest) SetMargins(v ConvertSpreadsheetMarginsRequest) {
	o.Margins = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ConvertSpreadsheetLayoutRequest) GetPageSize() ConvertSpreadsheetPageSizeRequest {
	if o == nil || IsNil(o.PageSize) {
		var ret ConvertSpreadsheetPageSizeRequest
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertSpreadsheetLayoutRequest) GetPageSizeOk() (*ConvertSpreadsheetPageSizeRequest, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ConvertSpreadsheetLayoutRequest) IsPageSizeSet() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given ConvertSpreadsheetPageSizeRequest and assigns it to the PageSize field.
func (o *ConvertSpreadsheetLayoutRequest) SetPageSize(v ConvertSpreadsheetPageSizeRequest) {
	o.PageSize = &v
}

func (o ConvertSpreadsheetLayoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertSpreadsheetLayoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IgnorePrintArea.IsSet() {
		toSerialize["ignorePrintArea"] = o.IgnorePrintArea.Get()
	}
	if o.Orientation.IsSet() {
		toSerialize["orientation"] = o.Orientation.Get()
	}
	if o.FitToHeight.IsSet() {
		toSerialize["fitToHeight"] = o.FitToHeight.Get()
	}
	if o.Scale.IsSet() {
		toSerialize["scale"] = o.Scale.Get()
	}
	if o.FitToWidth.IsSet() {
		toSerialize["fitToWidth"] = o.FitToWidth.Get()
	}
	if o.Headings.IsSet() {
		toSerialize["headings"] = o.Headings.Get()
	}
	if o.GridLines.IsSet() {
		toSerialize["gridLines"] = o.GridLines.Get()
	}
	if !IsNil(o.Margins) {
		toSerialize["margins"] = o.Margins
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableConvertSpreadsheetLayoutRequest struct {
	value *ConvertSpreadsheetLayoutRequest
	isSet bool
}

func (v NullableConvertSpreadsheetLayoutRequest) Get() *ConvertSpreadsheetLayoutRequest {
	return v.value
}

func (v *NullableConvertSpreadsheetLayoutRequest) Set(val *ConvertSpreadsheetLayoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertSpreadsheetLayoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertSpreadsheetLayoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertSpreadsheetLayoutRequest(val *ConvertSpreadsheetLayoutRequest) *NullableConvertSpreadsheetLayoutRequest {
	return &NullableConvertSpreadsheetLayoutRequest{value: val, isSet: true}
}

func (v NullableConvertSpreadsheetLayoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertSpreadsheetLayoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

