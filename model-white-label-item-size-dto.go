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

// checks if the WhiteLabelItemSizeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WhiteLabelItemSizeDto{}

// WhiteLabelItemSizeDto The pixel box a logo slot is drawn in, in the shape the imaging library reports a geometry.
type WhiteLabelItemSizeDto struct {
	// Whether the numbers are to be read as an aspect ratio rather than as pixels. Always `false` on the sizes  this API reports.
	AspectRatio *bool `json:"aspectRatio,omitempty"`
	// Whether an image would be scaled to cover the box rather than to fit inside it. Always `false` here.
	FillArea *bool `json:"fillArea,omitempty"`
	// Whether scaling would apply only to an image larger than the box. Always `false` here.
	Greater *bool `json:"greater,omitempty"`
	// The height of the box in pixels - one of the two fields of this object that carry information.
	Height *int32 `json:"height,omitempty"`
	// Whether scaling would be allowed to distort the image. Always `false` here.
	IgnoreAspectRatio *bool `json:"ignoreAspectRatio,omitempty"`
	// Whether `width` and `height` are to be read as percentages. Always `false` here, so both are pixels.
	IsPercentage *bool `json:"isPercentage,omitempty"`
	// Whether scaling would apply only to an image smaller than the box. Always `false` here.
	Less *bool `json:"less,omitempty"`
	// Whether the box is to be read as a total pixel-area budget instead of as two dimensions. Always `false`  here.
	LimitPixels *bool `json:"limitPixels,omitempty"`
	// The width of the box in pixels - the other field of this object that carries information.
	Width *int32 `json:"width,omitempty"`
	// The horizontal offset of the box from the origin. Always `0` here.
	X *int32 `json:"x,omitempty"`
	// The vertical offset of the box from the origin. Always `0` here.
	Y *int32 `json:"y,omitempty"`
}

// NewWhiteLabelItemSizeDto instantiates a new WhiteLabelItemSizeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhiteLabelItemSizeDto() *WhiteLabelItemSizeDto {
	this := WhiteLabelItemSizeDto{}
	return &this
}

// NewWhiteLabelItemSizeDtoWithDefaults instantiates a new WhiteLabelItemSizeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhiteLabelItemSizeDtoWithDefaults() *WhiteLabelItemSizeDto {
	this := WhiteLabelItemSizeDto{}
	return &this
}

// GetAspectRatio returns the AspectRatio field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetAspectRatio() bool {
	if o == nil || IsNil(o.AspectRatio) {
		var ret bool
		return ret
	}
	return *o.AspectRatio
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetAspectRatioOk() (*bool, bool) {
	if o == nil || IsNil(o.AspectRatio) {
		return nil, false
	}
	return o.AspectRatio, true
}

// HasAspectRatio returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsAspectRatioSet() bool {
	if o != nil && !IsNil(o.AspectRatio) {
		return true
	}

	return false
}

// SetAspectRatio gets a reference to the given bool and assigns it to the AspectRatio field.
func (o *WhiteLabelItemSizeDto) SetAspectRatio(v bool) {
	o.AspectRatio = &v
}

// GetFillArea returns the FillArea field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetFillArea() bool {
	if o == nil || IsNil(o.FillArea) {
		var ret bool
		return ret
	}
	return *o.FillArea
}

// GetFillAreaOk returns a tuple with the FillArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetFillAreaOk() (*bool, bool) {
	if o == nil || IsNil(o.FillArea) {
		return nil, false
	}
	return o.FillArea, true
}

// HasFillArea returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsFillAreaSet() bool {
	if o != nil && !IsNil(o.FillArea) {
		return true
	}

	return false
}

// SetFillArea gets a reference to the given bool and assigns it to the FillArea field.
func (o *WhiteLabelItemSizeDto) SetFillArea(v bool) {
	o.FillArea = &v
}

// GetGreater returns the Greater field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetGreater() bool {
	if o == nil || IsNil(o.Greater) {
		var ret bool
		return ret
	}
	return *o.Greater
}

// GetGreaterOk returns a tuple with the Greater field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetGreaterOk() (*bool, bool) {
	if o == nil || IsNil(o.Greater) {
		return nil, false
	}
	return o.Greater, true
}

// HasGreater returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsGreaterSet() bool {
	if o != nil && !IsNil(o.Greater) {
		return true
	}

	return false
}

// SetGreater gets a reference to the given bool and assigns it to the Greater field.
func (o *WhiteLabelItemSizeDto) SetGreater(v bool) {
	o.Greater = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsHeightSet() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *WhiteLabelItemSizeDto) SetHeight(v int32) {
	o.Height = &v
}

// GetIgnoreAspectRatio returns the IgnoreAspectRatio field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetIgnoreAspectRatio() bool {
	if o == nil || IsNil(o.IgnoreAspectRatio) {
		var ret bool
		return ret
	}
	return *o.IgnoreAspectRatio
}

// GetIgnoreAspectRatioOk returns a tuple with the IgnoreAspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetIgnoreAspectRatioOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreAspectRatio) {
		return nil, false
	}
	return o.IgnoreAspectRatio, true
}

// HasIgnoreAspectRatio returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsIgnoreAspectRatioSet() bool {
	if o != nil && !IsNil(o.IgnoreAspectRatio) {
		return true
	}

	return false
}

// SetIgnoreAspectRatio gets a reference to the given bool and assigns it to the IgnoreAspectRatio field.
func (o *WhiteLabelItemSizeDto) SetIgnoreAspectRatio(v bool) {
	o.IgnoreAspectRatio = &v
}

// GetIsPercentage returns the IsPercentage field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetIsPercentage() bool {
	if o == nil || IsNil(o.IsPercentage) {
		var ret bool
		return ret
	}
	return *o.IsPercentage
}

// GetIsPercentageOk returns a tuple with the IsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetIsPercentageOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPercentage) {
		return nil, false
	}
	return o.IsPercentage, true
}

// HasIsPercentage returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsIsPercentageSet() bool {
	if o != nil && !IsNil(o.IsPercentage) {
		return true
	}

	return false
}

// SetIsPercentage gets a reference to the given bool and assigns it to the IsPercentage field.
func (o *WhiteLabelItemSizeDto) SetIsPercentage(v bool) {
	o.IsPercentage = &v
}

// GetLess returns the Less field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetLess() bool {
	if o == nil || IsNil(o.Less) {
		var ret bool
		return ret
	}
	return *o.Less
}

// GetLessOk returns a tuple with the Less field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetLessOk() (*bool, bool) {
	if o == nil || IsNil(o.Less) {
		return nil, false
	}
	return o.Less, true
}

// HasLess returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsLessSet() bool {
	if o != nil && !IsNil(o.Less) {
		return true
	}

	return false
}

// SetLess gets a reference to the given bool and assigns it to the Less field.
func (o *WhiteLabelItemSizeDto) SetLess(v bool) {
	o.Less = &v
}

// GetLimitPixels returns the LimitPixels field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetLimitPixels() bool {
	if o == nil || IsNil(o.LimitPixels) {
		var ret bool
		return ret
	}
	return *o.LimitPixels
}

// GetLimitPixelsOk returns a tuple with the LimitPixels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetLimitPixelsOk() (*bool, bool) {
	if o == nil || IsNil(o.LimitPixels) {
		return nil, false
	}
	return o.LimitPixels, true
}

// HasLimitPixels returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsLimitPixelsSet() bool {
	if o != nil && !IsNil(o.LimitPixels) {
		return true
	}

	return false
}

// SetLimitPixels gets a reference to the given bool and assigns it to the LimitPixels field.
func (o *WhiteLabelItemSizeDto) SetLimitPixels(v bool) {
	o.LimitPixels = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsWidthSet() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *WhiteLabelItemSizeDto) SetWidth(v int32) {
	o.Width = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetX() int32 {
	if o == nil || IsNil(o.X) {
		var ret int32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetXOk() (*int32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsXSet() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given int32 and assigns it to the X field.
func (o *WhiteLabelItemSizeDto) SetX(v int32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *WhiteLabelItemSizeDto) GetY() int32 {
	if o == nil || IsNil(o.Y) {
		var ret int32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemSizeDto) GetYOk() (*int32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *WhiteLabelItemSizeDto) IsYSet() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given int32 and assigns it to the Y field.
func (o *WhiteLabelItemSizeDto) SetY(v int32) {
	o.Y = &v
}

func (o WhiteLabelItemSizeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WhiteLabelItemSizeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AspectRatio) {
		toSerialize["aspectRatio"] = o.AspectRatio
	}
	if !IsNil(o.FillArea) {
		toSerialize["fillArea"] = o.FillArea
	}
	if !IsNil(o.Greater) {
		toSerialize["greater"] = o.Greater
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.IgnoreAspectRatio) {
		toSerialize["ignoreAspectRatio"] = o.IgnoreAspectRatio
	}
	if !IsNil(o.IsPercentage) {
		toSerialize["isPercentage"] = o.IsPercentage
	}
	if !IsNil(o.Less) {
		toSerialize["less"] = o.Less
	}
	if !IsNil(o.LimitPixels) {
		toSerialize["limitPixels"] = o.LimitPixels
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	return toSerialize, nil
}

type NullableWhiteLabelItemSizeDto struct {
	value *WhiteLabelItemSizeDto
	isSet bool
}

func (v NullableWhiteLabelItemSizeDto) Get() *WhiteLabelItemSizeDto {
	return v.value
}

func (v *NullableWhiteLabelItemSizeDto) Set(val *WhiteLabelItemSizeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWhiteLabelItemSizeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWhiteLabelItemSizeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhiteLabelItemSizeDto(val *WhiteLabelItemSizeDto) *NullableWhiteLabelItemSizeDto {
	return &NullableWhiteLabelItemSizeDto{value: val, isSet: true}
}

func (v NullableWhiteLabelItemSizeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhiteLabelItemSizeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

