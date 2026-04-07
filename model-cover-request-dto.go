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

// checks if the CoverRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoverRequestDto{}

// CoverRequestDto The request parameters to change the room cover.
type CoverRequestDto struct {
	// The cover color.
	Color NullableString `json:"color,omitempty" validate:"regexp=^([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"`
	// The cover name.
	Cover NullableString `json:"cover,omitempty"`
}

// NewCoverRequestDto instantiates a new CoverRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoverRequestDto() *CoverRequestDto {
	this := CoverRequestDto{}
	return &this
}

// NewCoverRequestDtoWithDefaults instantiates a new CoverRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoverRequestDtoWithDefaults() *CoverRequestDto {
	this := CoverRequestDto{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoverRequestDto) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoverRequestDto) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *CoverRequestDto) IsColorSet() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *CoverRequestDto) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *CoverRequestDto) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *CoverRequestDto) UnsetColor() {
	o.Color.Unset()
}

// GetCover returns the Cover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoverRequestDto) GetCover() string {
	if o == nil || IsNil(o.Cover.Get()) {
		var ret string
		return ret
	}
	return *o.Cover.Get()
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoverRequestDto) GetCoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cover.Get(), o.Cover.IsSet()
}

// HasCover returns a boolean if a field has been set.
func (o *CoverRequestDto) IsCoverSet() bool {
	if o != nil && o.Cover.IsSet() {
		return true
	}

	return false
}

// SetCover gets a reference to the given NullableString and assigns it to the Cover field.
func (o *CoverRequestDto) SetCover(v string) {
	o.Cover.Set(&v)
}
// SetCoverNil sets the value for Cover to be an explicit nil
func (o *CoverRequestDto) SetCoverNil() {
	o.Cover.Set(nil)
}

// UnsetCover ensures that no value is present for Cover, not even an explicit nil
func (o *CoverRequestDto) UnsetCover() {
	o.Cover.Unset()
}

func (o CoverRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoverRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Cover.IsSet() {
		toSerialize["cover"] = o.Cover.Get()
	}
	return toSerialize, nil
}

type NullableCoverRequestDto struct {
	value *CoverRequestDto
	isSet bool
}

func (v NullableCoverRequestDto) Get() *CoverRequestDto {
	return v.value
}

func (v *NullableCoverRequestDto) Set(val *CoverRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCoverRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCoverRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoverRequestDto(val *CoverRequestDto) *NullableCoverRequestDto {
	return &NullableCoverRequestDto{value: val, isSet: true}
}

func (v NullableCoverRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoverRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

