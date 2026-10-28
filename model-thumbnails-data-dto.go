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

// checks if the ThumbnailsDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThumbnailsDataDto{}

// ThumbnailsDataDto The thumbnails data parameters.
type ThumbnailsDataDto struct {
	// The thumbnail original photo.
	Original NullableString `json:"original,omitempty"`
	// The thumbnail retina.
	Retina NullableString `json:"retina,omitempty"`
	// The thumbnail maximum size photo.
	Max NullableString `json:"max,omitempty"`
	// The thumbnail big size photo.
	Big NullableString `json:"big,omitempty"`
	// The thumbnail medium size photo.
	Medium NullableString `json:"medium,omitempty"`
	// The thumbnail small size photo.
	Small NullableString `json:"small,omitempty"`
}

// NewThumbnailsDataDto instantiates a new ThumbnailsDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThumbnailsDataDto() *ThumbnailsDataDto {
	this := ThumbnailsDataDto{}
	return &this
}

// NewThumbnailsDataDtoWithDefaults instantiates a new ThumbnailsDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThumbnailsDataDtoWithDefaults() *ThumbnailsDataDto {
	this := ThumbnailsDataDto{}
	return &this
}

// GetOriginal returns the Original field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThumbnailsDataDto) GetOriginal() string {
	if o == nil || IsNil(o.Original.Get()) {
		var ret string
		return ret
	}
	return *o.Original.Get()
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThumbnailsDataDto) GetOriginalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Original.Get(), o.Original.IsSet()
}

// HasOriginal returns a boolean if a field has been set.
func (o *ThumbnailsDataDto) IsOriginalSet() bool {
	if o != nil && o.Original.IsSet() {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given NullableString and assigns it to the Original field.
func (o *ThumbnailsDataDto) SetOriginal(v string) {
	o.Original.Set(&v)
}
// SetOriginalNil sets the value for Original to be an explicit nil
func (o *ThumbnailsDataDto) SetOriginalNil() {
	o.Original.Set(nil)
}

// UnsetOriginal ensures that no value is present for Original, not even an explicit nil
func (o *ThumbnailsDataDto) UnsetOriginal() {
	o.Original.Unset()
}

// GetRetina returns the Retina field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThumbnailsDataDto) GetRetina() string {
	if o == nil || IsNil(o.Retina.Get()) {
		var ret string
		return ret
	}
	return *o.Retina.Get()
}

// GetRetinaOk returns a tuple with the Retina field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThumbnailsDataDto) GetRetinaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Retina.Get(), o.Retina.IsSet()
}

// HasRetina returns a boolean if a field has been set.
func (o *ThumbnailsDataDto) IsRetinaSet() bool {
	if o != nil && o.Retina.IsSet() {
		return true
	}

	return false
}

// SetRetina gets a reference to the given NullableString and assigns it to the Retina field.
func (o *ThumbnailsDataDto) SetRetina(v string) {
	o.Retina.Set(&v)
}
// SetRetinaNil sets the value for Retina to be an explicit nil
func (o *ThumbnailsDataDto) SetRetinaNil() {
	o.Retina.Set(nil)
}

// UnsetRetina ensures that no value is present for Retina, not even an explicit nil
func (o *ThumbnailsDataDto) UnsetRetina() {
	o.Retina.Unset()
}

// GetMax returns the Max field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThumbnailsDataDto) GetMax() string {
	if o == nil || IsNil(o.Max.Get()) {
		var ret string
		return ret
	}
	return *o.Max.Get()
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThumbnailsDataDto) GetMaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Max.Get(), o.Max.IsSet()
}

// HasMax returns a boolean if a field has been set.
func (o *ThumbnailsDataDto) IsMaxSet() bool {
	if o != nil && o.Max.IsSet() {
		return true
	}

	return false
}

// SetMax gets a reference to the given NullableString and assigns it to the Max field.
func (o *ThumbnailsDataDto) SetMax(v string) {
	o.Max.Set(&v)
}
// SetMaxNil sets the value for Max to be an explicit nil
func (o *ThumbnailsDataDto) SetMaxNil() {
	o.Max.Set(nil)
}

// UnsetMax ensures that no value is present for Max, not even an explicit nil
func (o *ThumbnailsDataDto) UnsetMax() {
	o.Max.Unset()
}

// GetBig returns the Big field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThumbnailsDataDto) GetBig() string {
	if o == nil || IsNil(o.Big.Get()) {
		var ret string
		return ret
	}
	return *o.Big.Get()
}

// GetBigOk returns a tuple with the Big field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThumbnailsDataDto) GetBigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Big.Get(), o.Big.IsSet()
}

// HasBig returns a boolean if a field has been set.
func (o *ThumbnailsDataDto) IsBigSet() bool {
	if o != nil && o.Big.IsSet() {
		return true
	}

	return false
}

// SetBig gets a reference to the given NullableString and assigns it to the Big field.
func (o *ThumbnailsDataDto) SetBig(v string) {
	o.Big.Set(&v)
}
// SetBigNil sets the value for Big to be an explicit nil
func (o *ThumbnailsDataDto) SetBigNil() {
	o.Big.Set(nil)
}

// UnsetBig ensures that no value is present for Big, not even an explicit nil
func (o *ThumbnailsDataDto) UnsetBig() {
	o.Big.Unset()
}

// GetMedium returns the Medium field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThumbnailsDataDto) GetMedium() string {
	if o == nil || IsNil(o.Medium.Get()) {
		var ret string
		return ret
	}
	return *o.Medium.Get()
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThumbnailsDataDto) GetMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Medium.Get(), o.Medium.IsSet()
}

// HasMedium returns a boolean if a field has been set.
func (o *ThumbnailsDataDto) IsMediumSet() bool {
	if o != nil && o.Medium.IsSet() {
		return true
	}

	return false
}

// SetMedium gets a reference to the given NullableString and assigns it to the Medium field.
func (o *ThumbnailsDataDto) SetMedium(v string) {
	o.Medium.Set(&v)
}
// SetMediumNil sets the value for Medium to be an explicit nil
func (o *ThumbnailsDataDto) SetMediumNil() {
	o.Medium.Set(nil)
}

// UnsetMedium ensures that no value is present for Medium, not even an explicit nil
func (o *ThumbnailsDataDto) UnsetMedium() {
	o.Medium.Unset()
}

// GetSmall returns the Small field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThumbnailsDataDto) GetSmall() string {
	if o == nil || IsNil(o.Small.Get()) {
		var ret string
		return ret
	}
	return *o.Small.Get()
}

// GetSmallOk returns a tuple with the Small field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThumbnailsDataDto) GetSmallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Small.Get(), o.Small.IsSet()
}

// HasSmall returns a boolean if a field has been set.
func (o *ThumbnailsDataDto) IsSmallSet() bool {
	if o != nil && o.Small.IsSet() {
		return true
	}

	return false
}

// SetSmall gets a reference to the given NullableString and assigns it to the Small field.
func (o *ThumbnailsDataDto) SetSmall(v string) {
	o.Small.Set(&v)
}
// SetSmallNil sets the value for Small to be an explicit nil
func (o *ThumbnailsDataDto) SetSmallNil() {
	o.Small.Set(nil)
}

// UnsetSmall ensures that no value is present for Small, not even an explicit nil
func (o *ThumbnailsDataDto) UnsetSmall() {
	o.Small.Unset()
}

func (o ThumbnailsDataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThumbnailsDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Original.IsSet() {
		toSerialize["original"] = o.Original.Get()
	}
	if o.Retina.IsSet() {
		toSerialize["retina"] = o.Retina.Get()
	}
	if o.Max.IsSet() {
		toSerialize["max"] = o.Max.Get()
	}
	if o.Big.IsSet() {
		toSerialize["big"] = o.Big.Get()
	}
	if o.Medium.IsSet() {
		toSerialize["medium"] = o.Medium.Get()
	}
	if o.Small.IsSet() {
		toSerialize["small"] = o.Small.Get()
	}
	return toSerialize, nil
}

type NullableThumbnailsDataDto struct {
	value *ThumbnailsDataDto
	isSet bool
}

func (v NullableThumbnailsDataDto) Get() *ThumbnailsDataDto {
	return v.value
}

func (v *NullableThumbnailsDataDto) Set(val *ThumbnailsDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThumbnailsDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThumbnailsDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThumbnailsDataDto(val *ThumbnailsDataDto) *NullableThumbnailsDataDto {
	return &NullableThumbnailsDataDto{value: val, isSet: true}
}

func (v NullableThumbnailsDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThumbnailsDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

