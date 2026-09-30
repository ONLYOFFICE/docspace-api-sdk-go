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

// checks if the WatermarkRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatermarkRequestDto{}

// WatermarkRequestDto The watermark drawn over the documents of a room.
type WatermarkRequestDto struct {
	// Whether the room draws a watermark at all. Sending the object with this turned off removes the watermark the  room has, and the rest of the fields are then irrelevant.
	Enabled NullableBool `json:"enabled,omitempty"`
	// Which details of the reader and of the room are stamped into the watermark alongside the text. The values  combine, so several of them can be added together to stamp more than one.
	Additions *WatermarkAdditions `json:"additions,omitempty"`
	// The fixed line drawn over the document, shown before the details selected alongside it. It is the whole  watermark when no details are added.
	Text NullableString `json:"text,omitempty"`
	// How far the watermark is turned, in degrees, with negative values turning it anticlockwise. Zero draws it  horizontally across the page.
	Rotate *int32 `json:"rotate,omitempty"`
	// How large the watermark image is drawn, as a percentage of its own size. It applies to the image form of the  watermark only.
	ImageScale *int32 `json:"imageScale,omitempty"`
	// The picture to use instead of a text watermark, named by the path that `POST api/2.0/files/logos` returned for  an image uploaded beforehand. The portal copies it into the room when the setting is saved.
	ImageUrl NullableString `json:"imageUrl,omitempty"`
	// The height the watermark image is drawn with, in pixels, used together with the width to keep its proportions.
	ImageHeight *float64 `json:"imageHeight,omitempty"`
	// The width the watermark image is drawn with, in pixels, used together with the height to keep its proportions.
	ImageWidth *float64 `json:"imageWidth,omitempty"`
}

// NewWatermarkRequestDto instantiates a new WatermarkRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatermarkRequestDto() *WatermarkRequestDto {
	this := WatermarkRequestDto{}
	return &this
}

// NewWatermarkRequestDtoWithDefaults instantiates a new WatermarkRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatermarkRequestDtoWithDefaults() *WatermarkRequestDto {
	this := WatermarkRequestDto{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatermarkRequestDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatermarkRequestDto) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *WatermarkRequestDto) IsEnabledSet() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *WatermarkRequestDto) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *WatermarkRequestDto) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *WatermarkRequestDto) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetAdditions returns the Additions field value if set, zero value otherwise.
func (o *WatermarkRequestDto) GetAdditions() WatermarkAdditions {
	if o == nil || IsNil(o.Additions) {
		var ret WatermarkAdditions
		return ret
	}
	return *o.Additions
}

// GetAdditionsOk returns a tuple with the Additions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkRequestDto) GetAdditionsOk() (*WatermarkAdditions, bool) {
	if o == nil || IsNil(o.Additions) {
		return nil, false
	}
	return o.Additions, true
}

// HasAdditions returns a boolean if a field has been set.
func (o *WatermarkRequestDto) IsAdditionsSet() bool {
	if o != nil && !IsNil(o.Additions) {
		return true
	}

	return false
}

// SetAdditions gets a reference to the given WatermarkAdditions and assigns it to the Additions field.
func (o *WatermarkRequestDto) SetAdditions(v WatermarkAdditions) {
	o.Additions = &v
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatermarkRequestDto) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatermarkRequestDto) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *WatermarkRequestDto) IsTextSet() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *WatermarkRequestDto) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *WatermarkRequestDto) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *WatermarkRequestDto) UnsetText() {
	o.Text.Unset()
}

// GetRotate returns the Rotate field value if set, zero value otherwise.
func (o *WatermarkRequestDto) GetRotate() int32 {
	if o == nil || IsNil(o.Rotate) {
		var ret int32
		return ret
	}
	return *o.Rotate
}

// GetRotateOk returns a tuple with the Rotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkRequestDto) GetRotateOk() (*int32, bool) {
	if o == nil || IsNil(o.Rotate) {
		return nil, false
	}
	return o.Rotate, true
}

// HasRotate returns a boolean if a field has been set.
func (o *WatermarkRequestDto) IsRotateSet() bool {
	if o != nil && !IsNil(o.Rotate) {
		return true
	}

	return false
}

// SetRotate gets a reference to the given int32 and assigns it to the Rotate field.
func (o *WatermarkRequestDto) SetRotate(v int32) {
	o.Rotate = &v
}

// GetImageScale returns the ImageScale field value if set, zero value otherwise.
func (o *WatermarkRequestDto) GetImageScale() int32 {
	if o == nil || IsNil(o.ImageScale) {
		var ret int32
		return ret
	}
	return *o.ImageScale
}

// GetImageScaleOk returns a tuple with the ImageScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkRequestDto) GetImageScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.ImageScale) {
		return nil, false
	}
	return o.ImageScale, true
}

// HasImageScale returns a boolean if a field has been set.
func (o *WatermarkRequestDto) IsImageScaleSet() bool {
	if o != nil && !IsNil(o.ImageScale) {
		return true
	}

	return false
}

// SetImageScale gets a reference to the given int32 and assigns it to the ImageScale field.
func (o *WatermarkRequestDto) SetImageScale(v int32) {
	o.ImageScale = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatermarkRequestDto) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatermarkRequestDto) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *WatermarkRequestDto) IsImageUrlSet() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *WatermarkRequestDto) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *WatermarkRequestDto) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *WatermarkRequestDto) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetImageHeight returns the ImageHeight field value if set, zero value otherwise.
func (o *WatermarkRequestDto) GetImageHeight() float64 {
	if o == nil || IsNil(o.ImageHeight) {
		var ret float64
		return ret
	}
	return *o.ImageHeight
}

// GetImageHeightOk returns a tuple with the ImageHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkRequestDto) GetImageHeightOk() (*float64, bool) {
	if o == nil || IsNil(o.ImageHeight) {
		return nil, false
	}
	return o.ImageHeight, true
}

// HasImageHeight returns a boolean if a field has been set.
func (o *WatermarkRequestDto) IsImageHeightSet() bool {
	if o != nil && !IsNil(o.ImageHeight) {
		return true
	}

	return false
}

// SetImageHeight gets a reference to the given float64 and assigns it to the ImageHeight field.
func (o *WatermarkRequestDto) SetImageHeight(v float64) {
	o.ImageHeight = &v
}

// GetImageWidth returns the ImageWidth field value if set, zero value otherwise.
func (o *WatermarkRequestDto) GetImageWidth() float64 {
	if o == nil || IsNil(o.ImageWidth) {
		var ret float64
		return ret
	}
	return *o.ImageWidth
}

// GetImageWidthOk returns a tuple with the ImageWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkRequestDto) GetImageWidthOk() (*float64, bool) {
	if o == nil || IsNil(o.ImageWidth) {
		return nil, false
	}
	return o.ImageWidth, true
}

// HasImageWidth returns a boolean if a field has been set.
func (o *WatermarkRequestDto) IsImageWidthSet() bool {
	if o != nil && !IsNil(o.ImageWidth) {
		return true
	}

	return false
}

// SetImageWidth gets a reference to the given float64 and assigns it to the ImageWidth field.
func (o *WatermarkRequestDto) SetImageWidth(v float64) {
	o.ImageWidth = &v
}

func (o WatermarkRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatermarkRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if !IsNil(o.Additions) {
		toSerialize["additions"] = o.Additions
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if !IsNil(o.Rotate) {
		toSerialize["rotate"] = o.Rotate
	}
	if !IsNil(o.ImageScale) {
		toSerialize["imageScale"] = o.ImageScale
	}
	if o.ImageUrl.IsSet() {
		toSerialize["imageUrl"] = o.ImageUrl.Get()
	}
	if !IsNil(o.ImageHeight) {
		toSerialize["imageHeight"] = o.ImageHeight
	}
	if !IsNil(o.ImageWidth) {
		toSerialize["imageWidth"] = o.ImageWidth
	}
	return toSerialize, nil
}

type NullableWatermarkRequestDto struct {
	value *WatermarkRequestDto
	isSet bool
}

func (v NullableWatermarkRequestDto) Get() *WatermarkRequestDto {
	return v.value
}

func (v *NullableWatermarkRequestDto) Set(val *WatermarkRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWatermarkRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWatermarkRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatermarkRequestDto(val *WatermarkRequestDto) *NullableWatermarkRequestDto {
	return &NullableWatermarkRequestDto{value: val, isSet: true}
}

func (v NullableWatermarkRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatermarkRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

