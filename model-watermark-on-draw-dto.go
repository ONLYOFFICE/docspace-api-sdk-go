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

// checks if the WatermarkOnDrawDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatermarkOnDrawDto{}

// WatermarkOnDrawDto The watermark the editor draws over every page.
type WatermarkOnDrawDto struct {
	// Defines the watermark width measured in millimeters.
	Width *float64 `json:"width,omitempty"`
	// Defines the watermark height measured in millimeters.
	Height *float64 `json:"height,omitempty"`
	// Defines the watermark margins measured in millimeters.
	Margins []int32 `json:"margins,omitempty"`
	// Defines the watermark fill color.
	Fill NullableString `json:"fill,omitempty"`
	// Defines the watermark rotation angle.
	Rotate *int32 `json:"rotate,omitempty"`
	// Defines the watermark transparency percentage.
	Transparent *float64 `json:"transparent,omitempty"`
	// The list of paragraphs of the watermark.
	Paragraphs []WatermarkParagraphDto `json:"paragraphs,omitempty"`
}

// NewWatermarkOnDrawDto instantiates a new WatermarkOnDrawDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatermarkOnDrawDto() *WatermarkOnDrawDto {
	this := WatermarkOnDrawDto{}
	return &this
}

// NewWatermarkOnDrawDtoWithDefaults instantiates a new WatermarkOnDrawDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatermarkOnDrawDtoWithDefaults() *WatermarkOnDrawDto {
	this := WatermarkOnDrawDto{}
	return &this
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *WatermarkOnDrawDto) GetWidth() float64 {
	if o == nil || IsNil(o.Width) {
		var ret float64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkOnDrawDto) GetWidthOk() (*float64, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *WatermarkOnDrawDto) IsWidthSet() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float64 and assigns it to the Width field.
func (o *WatermarkOnDrawDto) SetWidth(v float64) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *WatermarkOnDrawDto) GetHeight() float64 {
	if o == nil || IsNil(o.Height) {
		var ret float64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkOnDrawDto) GetHeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *WatermarkOnDrawDto) IsHeightSet() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float64 and assigns it to the Height field.
func (o *WatermarkOnDrawDto) SetHeight(v float64) {
	o.Height = &v
}

// GetMargins returns the Margins field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatermarkOnDrawDto) GetMargins() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Margins
}

// GetMarginsOk returns a tuple with the Margins field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatermarkOnDrawDto) GetMarginsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Margins) {
		return nil, false
	}
	return o.Margins, true
}

// HasMargins returns a boolean if a field has been set.
func (o *WatermarkOnDrawDto) IsMarginsSet() bool {
	if o != nil && !IsNil(o.Margins) {
		return true
	}

	return false
}

// SetMargins gets a reference to the given []int32 and assigns it to the Margins field.
func (o *WatermarkOnDrawDto) SetMargins(v []int32) {
	o.Margins = v
}

// GetFill returns the Fill field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatermarkOnDrawDto) GetFill() string {
	if o == nil || IsNil(o.Fill.Get()) {
		var ret string
		return ret
	}
	return *o.Fill.Get()
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatermarkOnDrawDto) GetFillOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fill.Get(), o.Fill.IsSet()
}

// HasFill returns a boolean if a field has been set.
func (o *WatermarkOnDrawDto) IsFillSet() bool {
	if o != nil && o.Fill.IsSet() {
		return true
	}

	return false
}

// SetFill gets a reference to the given NullableString and assigns it to the Fill field.
func (o *WatermarkOnDrawDto) SetFill(v string) {
	o.Fill.Set(&v)
}
// SetFillNil sets the value for Fill to be an explicit nil
func (o *WatermarkOnDrawDto) SetFillNil() {
	o.Fill.Set(nil)
}

// UnsetFill ensures that no value is present for Fill, not even an explicit nil
func (o *WatermarkOnDrawDto) UnsetFill() {
	o.Fill.Unset()
}

// GetRotate returns the Rotate field value if set, zero value otherwise.
func (o *WatermarkOnDrawDto) GetRotate() int32 {
	if o == nil || IsNil(o.Rotate) {
		var ret int32
		return ret
	}
	return *o.Rotate
}

// GetRotateOk returns a tuple with the Rotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkOnDrawDto) GetRotateOk() (*int32, bool) {
	if o == nil || IsNil(o.Rotate) {
		return nil, false
	}
	return o.Rotate, true
}

// HasRotate returns a boolean if a field has been set.
func (o *WatermarkOnDrawDto) IsRotateSet() bool {
	if o != nil && !IsNil(o.Rotate) {
		return true
	}

	return false
}

// SetRotate gets a reference to the given int32 and assigns it to the Rotate field.
func (o *WatermarkOnDrawDto) SetRotate(v int32) {
	o.Rotate = &v
}

// GetTransparent returns the Transparent field value if set, zero value otherwise.
func (o *WatermarkOnDrawDto) GetTransparent() float64 {
	if o == nil || IsNil(o.Transparent) {
		var ret float64
		return ret
	}
	return *o.Transparent
}

// GetTransparentOk returns a tuple with the Transparent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkOnDrawDto) GetTransparentOk() (*float64, bool) {
	if o == nil || IsNil(o.Transparent) {
		return nil, false
	}
	return o.Transparent, true
}

// HasTransparent returns a boolean if a field has been set.
func (o *WatermarkOnDrawDto) IsTransparentSet() bool {
	if o != nil && !IsNil(o.Transparent) {
		return true
	}

	return false
}

// SetTransparent gets a reference to the given float64 and assigns it to the Transparent field.
func (o *WatermarkOnDrawDto) SetTransparent(v float64) {
	o.Transparent = &v
}

// GetParagraphs returns the Paragraphs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatermarkOnDrawDto) GetParagraphs() []WatermarkParagraphDto {
	if o == nil {
		var ret []WatermarkParagraphDto
		return ret
	}
	return o.Paragraphs
}

// GetParagraphsOk returns a tuple with the Paragraphs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatermarkOnDrawDto) GetParagraphsOk() ([]WatermarkParagraphDto, bool) {
	if o == nil || IsNil(o.Paragraphs) {
		return nil, false
	}
	return o.Paragraphs, true
}

// HasParagraphs returns a boolean if a field has been set.
func (o *WatermarkOnDrawDto) IsParagraphsSet() bool {
	if o != nil && !IsNil(o.Paragraphs) {
		return true
	}

	return false
}

// SetParagraphs gets a reference to the given []WatermarkParagraphDto and assigns it to the Paragraphs field.
func (o *WatermarkOnDrawDto) SetParagraphs(v []WatermarkParagraphDto) {
	o.Paragraphs = v
}

func (o WatermarkOnDrawDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatermarkOnDrawDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if o.Margins != nil {
		toSerialize["margins"] = o.Margins
	}
	if o.Fill.IsSet() {
		toSerialize["fill"] = o.Fill.Get()
	}
	if !IsNil(o.Rotate) {
		toSerialize["rotate"] = o.Rotate
	}
	if !IsNil(o.Transparent) {
		toSerialize["transparent"] = o.Transparent
	}
	if o.Paragraphs != nil {
		toSerialize["paragraphs"] = o.Paragraphs
	}
	return toSerialize, nil
}

type NullableWatermarkOnDrawDto struct {
	value *WatermarkOnDrawDto
	isSet bool
}

func (v NullableWatermarkOnDrawDto) Get() *WatermarkOnDrawDto {
	return v.value
}

func (v *NullableWatermarkOnDrawDto) Set(val *WatermarkOnDrawDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWatermarkOnDrawDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWatermarkOnDrawDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatermarkOnDrawDto(val *WatermarkOnDrawDto) *NullableWatermarkOnDrawDto {
	return &NullableWatermarkOnDrawDto{value: val, isSet: true}
}

func (v NullableWatermarkOnDrawDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatermarkOnDrawDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

