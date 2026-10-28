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
	"bytes"
	"fmt"
)

// checks if the WatermarkDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatermarkDto{}

// WatermarkDto The watermark drawn over the documents of a room while they are viewed and printed.
type WatermarkDto struct {
	// Which details of the reader and of the room are stamped alongside the text. The values combine, so a number  that is not a member on its own is the sum of several of them, and 0 means that only the text is stamped.
	Additions WatermarkAdditions `json:"additions"`
	// The fixed line drawn over the document, printed before the details selected alongside it. Empty when the room  stamps an image instead.
	Text NullableString `json:"text,omitempty"`
	// How far the stamp is turned, in degrees, with negative values turning it anticlockwise and 0 drawing it  horizontally.
	Rotate int32 `json:"rotate"`
	// How large the image is drawn, as a percentage of its own size. It is 0 for a text watermark, where nothing is  scaled.
	ImageScale int32 `json:"imageScale"`
	// The address the stamped picture is served from, inside the storage of the room. Empty for a text watermark.
	ImageUrl NullableString `json:"imageUrl,omitempty"`
	// The height the picture is drawn with, in pixels, kept together with the width so that the proportions survive.  It is 0 for a text watermark.
	ImageHeight float64 `json:"imageHeight"`
	// The width the picture is drawn with, in pixels, kept together with the height so that the proportions survive.  It is 0 for a text watermark.
	ImageWidth float64 `json:"imageWidth"`
}

type _WatermarkDto WatermarkDto

// NewWatermarkDto instantiates a new WatermarkDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatermarkDto(additions WatermarkAdditions, rotate int32, imageScale int32, imageHeight float64, imageWidth float64) *WatermarkDto {
	this := WatermarkDto{}
	this.Additions = additions
	this.Rotate = rotate
	this.ImageScale = imageScale
	this.ImageHeight = imageHeight
	this.ImageWidth = imageWidth
	return &this
}

// NewWatermarkDtoWithDefaults instantiates a new WatermarkDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatermarkDtoWithDefaults() *WatermarkDto {
	this := WatermarkDto{}
	return &this
}

// GetAdditions returns the Additions field value
func (o *WatermarkDto) GetAdditions() WatermarkAdditions {
	if o == nil {
		var ret WatermarkAdditions
		return ret
	}

	return o.Additions
}

// GetAdditionsOk returns a tuple with the Additions field value
// and a boolean to check if the value has been set.
func (o *WatermarkDto) GetAdditionsOk() (*WatermarkAdditions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Additions, true
}

// SetAdditions sets field value
func (o *WatermarkDto) SetAdditions(v WatermarkAdditions) {
	o.Additions = v
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatermarkDto) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatermarkDto) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *WatermarkDto) IsTextSet() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *WatermarkDto) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *WatermarkDto) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *WatermarkDto) UnsetText() {
	o.Text.Unset()
}

// GetRotate returns the Rotate field value
func (o *WatermarkDto) GetRotate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rotate
}

// GetRotateOk returns a tuple with the Rotate field value
// and a boolean to check if the value has been set.
func (o *WatermarkDto) GetRotateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rotate, true
}

// SetRotate sets field value
func (o *WatermarkDto) SetRotate(v int32) {
	o.Rotate = v
}

// GetImageScale returns the ImageScale field value
func (o *WatermarkDto) GetImageScale() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImageScale
}

// GetImageScaleOk returns a tuple with the ImageScale field value
// and a boolean to check if the value has been set.
func (o *WatermarkDto) GetImageScaleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageScale, true
}

// SetImageScale sets field value
func (o *WatermarkDto) SetImageScale(v int32) {
	o.ImageScale = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatermarkDto) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatermarkDto) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *WatermarkDto) IsImageUrlSet() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *WatermarkDto) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *WatermarkDto) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *WatermarkDto) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetImageHeight returns the ImageHeight field value
func (o *WatermarkDto) GetImageHeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ImageHeight
}

// GetImageHeightOk returns a tuple with the ImageHeight field value
// and a boolean to check if the value has been set.
func (o *WatermarkDto) GetImageHeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageHeight, true
}

// SetImageHeight sets field value
func (o *WatermarkDto) SetImageHeight(v float64) {
	o.ImageHeight = v
}

// GetImageWidth returns the ImageWidth field value
func (o *WatermarkDto) GetImageWidth() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ImageWidth
}

// GetImageWidthOk returns a tuple with the ImageWidth field value
// and a boolean to check if the value has been set.
func (o *WatermarkDto) GetImageWidthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageWidth, true
}

// SetImageWidth sets field value
func (o *WatermarkDto) SetImageWidth(v float64) {
	o.ImageWidth = v
}

func (o WatermarkDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatermarkDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["additions"] = o.Additions
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	toSerialize["rotate"] = o.Rotate
	toSerialize["imageScale"] = o.ImageScale
	if o.ImageUrl.IsSet() {
		toSerialize["imageUrl"] = o.ImageUrl.Get()
	}
	toSerialize["imageHeight"] = o.ImageHeight
	toSerialize["imageWidth"] = o.ImageWidth
	return toSerialize, nil
}

func (o *WatermarkDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"additions",
		"rotate",
		"imageScale",
		"imageHeight",
		"imageWidth",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWatermarkDto := _WatermarkDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWatermarkDto)

	if err != nil {
		return err
	}

	*o = WatermarkDto(varWatermarkDto)

	return err
}

type NullableWatermarkDto struct {
	value *WatermarkDto
	isSet bool
}

func (v NullableWatermarkDto) Get() *WatermarkDto {
	return v.value
}

func (v *NullableWatermarkDto) Set(val *WatermarkDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWatermarkDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWatermarkDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatermarkDto(val *WatermarkDto) *NullableWatermarkDto {
	return &NullableWatermarkDto{value: val, isSet: true}
}

func (v NullableWatermarkDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatermarkDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

