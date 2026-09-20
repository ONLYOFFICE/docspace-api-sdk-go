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

// checks if the LogoConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogoConfigDto{}

// LogoConfigDto The logo the editor shows, resolved for the file type and the layout of this opening.
type LogoConfigDto struct {
	// The logo for the current layout and file type, as the portal branding defines it.
	Image NullableString `json:"image,omitempty"`
	// The variant for a dark interface theme.
	ImageDark NullableString `json:"imageDark,omitempty"`
	// The variant for a light interface theme.
	ImageLight NullableString `json:"imageLight,omitempty"`
	// The variant for the framed viewer. It is empty in every layout but the embedded one.
	ImageEmbedded NullableString `json:"imageEmbedded,omitempty"`
	// Where clicking the logo takes the user.
	Url NullableString `json:"url,omitempty"`
	// Whether the logo is shown at all; the mobile layout hides it.
	Visible *bool `json:"visible,omitempty"`
}

// NewLogoConfigDto instantiates a new LogoConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogoConfigDto() *LogoConfigDto {
	this := LogoConfigDto{}
	return &this
}

// NewLogoConfigDtoWithDefaults instantiates a new LogoConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoConfigDtoWithDefaults() *LogoConfigDto {
	this := LogoConfigDto{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogoConfigDto) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoConfigDto) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *LogoConfigDto) IsImageSet() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *LogoConfigDto) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *LogoConfigDto) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *LogoConfigDto) UnsetImage() {
	o.Image.Unset()
}

// GetImageDark returns the ImageDark field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogoConfigDto) GetImageDark() string {
	if o == nil || IsNil(o.ImageDark.Get()) {
		var ret string
		return ret
	}
	return *o.ImageDark.Get()
}

// GetImageDarkOk returns a tuple with the ImageDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoConfigDto) GetImageDarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageDark.Get(), o.ImageDark.IsSet()
}

// HasImageDark returns a boolean if a field has been set.
func (o *LogoConfigDto) IsImageDarkSet() bool {
	if o != nil && o.ImageDark.IsSet() {
		return true
	}

	return false
}

// SetImageDark gets a reference to the given NullableString and assigns it to the ImageDark field.
func (o *LogoConfigDto) SetImageDark(v string) {
	o.ImageDark.Set(&v)
}
// SetImageDarkNil sets the value for ImageDark to be an explicit nil
func (o *LogoConfigDto) SetImageDarkNil() {
	o.ImageDark.Set(nil)
}

// UnsetImageDark ensures that no value is present for ImageDark, not even an explicit nil
func (o *LogoConfigDto) UnsetImageDark() {
	o.ImageDark.Unset()
}

// GetImageLight returns the ImageLight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogoConfigDto) GetImageLight() string {
	if o == nil || IsNil(o.ImageLight.Get()) {
		var ret string
		return ret
	}
	return *o.ImageLight.Get()
}

// GetImageLightOk returns a tuple with the ImageLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoConfigDto) GetImageLightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageLight.Get(), o.ImageLight.IsSet()
}

// HasImageLight returns a boolean if a field has been set.
func (o *LogoConfigDto) IsImageLightSet() bool {
	if o != nil && o.ImageLight.IsSet() {
		return true
	}

	return false
}

// SetImageLight gets a reference to the given NullableString and assigns it to the ImageLight field.
func (o *LogoConfigDto) SetImageLight(v string) {
	o.ImageLight.Set(&v)
}
// SetImageLightNil sets the value for ImageLight to be an explicit nil
func (o *LogoConfigDto) SetImageLightNil() {
	o.ImageLight.Set(nil)
}

// UnsetImageLight ensures that no value is present for ImageLight, not even an explicit nil
func (o *LogoConfigDto) UnsetImageLight() {
	o.ImageLight.Unset()
}

// GetImageEmbedded returns the ImageEmbedded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogoConfigDto) GetImageEmbedded() string {
	if o == nil || IsNil(o.ImageEmbedded.Get()) {
		var ret string
		return ret
	}
	return *o.ImageEmbedded.Get()
}

// GetImageEmbeddedOk returns a tuple with the ImageEmbedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoConfigDto) GetImageEmbeddedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageEmbedded.Get(), o.ImageEmbedded.IsSet()
}

// HasImageEmbedded returns a boolean if a field has been set.
func (o *LogoConfigDto) IsImageEmbeddedSet() bool {
	if o != nil && o.ImageEmbedded.IsSet() {
		return true
	}

	return false
}

// SetImageEmbedded gets a reference to the given NullableString and assigns it to the ImageEmbedded field.
func (o *LogoConfigDto) SetImageEmbedded(v string) {
	o.ImageEmbedded.Set(&v)
}
// SetImageEmbeddedNil sets the value for ImageEmbedded to be an explicit nil
func (o *LogoConfigDto) SetImageEmbeddedNil() {
	o.ImageEmbedded.Set(nil)
}

// UnsetImageEmbedded ensures that no value is present for ImageEmbedded, not even an explicit nil
func (o *LogoConfigDto) UnsetImageEmbedded() {
	o.ImageEmbedded.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogoConfigDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoConfigDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *LogoConfigDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *LogoConfigDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *LogoConfigDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *LogoConfigDto) UnsetUrl() {
	o.Url.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *LogoConfigDto) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogoConfigDto) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *LogoConfigDto) IsVisibleSet() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *LogoConfigDto) SetVisible(v bool) {
	o.Visible = &v
}

func (o LogoConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogoConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.ImageDark.IsSet() {
		toSerialize["imageDark"] = o.ImageDark.Get()
	}
	if o.ImageLight.IsSet() {
		toSerialize["imageLight"] = o.ImageLight.Get()
	}
	if o.ImageEmbedded.IsSet() {
		toSerialize["imageEmbedded"] = o.ImageEmbedded.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableLogoConfigDto struct {
	value *LogoConfigDto
	isSet bool
}

func (v NullableLogoConfigDto) Get() *LogoConfigDto {
	return v.value
}

func (v *NullableLogoConfigDto) Set(val *LogoConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLogoConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLogoConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogoConfigDto(val *LogoConfigDto) *NullableLogoConfigDto {
	return &NullableLogoConfigDto{value: val, isSet: true}
}

func (v NullableLogoConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogoConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

