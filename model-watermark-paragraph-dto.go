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

// checks if the WatermarkParagraphDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatermarkParagraphDto{}

// WatermarkParagraphDto One paragraph of the watermark text.
type WatermarkParagraphDto struct {
	// The paragraph align.
	Align *int32 `json:"align,omitempty"`
	// The list of text runs from the paragraph.
	Runs []WatermarkTextRunDto `json:"runs,omitempty"`
}

// NewWatermarkParagraphDto instantiates a new WatermarkParagraphDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatermarkParagraphDto() *WatermarkParagraphDto {
	this := WatermarkParagraphDto{}
	return &this
}

// NewWatermarkParagraphDtoWithDefaults instantiates a new WatermarkParagraphDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatermarkParagraphDtoWithDefaults() *WatermarkParagraphDto {
	this := WatermarkParagraphDto{}
	return &this
}

// GetAlign returns the Align field value if set, zero value otherwise.
func (o *WatermarkParagraphDto) GetAlign() int32 {
	if o == nil || IsNil(o.Align) {
		var ret int32
		return ret
	}
	return *o.Align
}

// GetAlignOk returns a tuple with the Align field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatermarkParagraphDto) GetAlignOk() (*int32, bool) {
	if o == nil || IsNil(o.Align) {
		return nil, false
	}
	return o.Align, true
}

// HasAlign returns a boolean if a field has been set.
func (o *WatermarkParagraphDto) IsAlignSet() bool {
	if o != nil && !IsNil(o.Align) {
		return true
	}

	return false
}

// SetAlign gets a reference to the given int32 and assigns it to the Align field.
func (o *WatermarkParagraphDto) SetAlign(v int32) {
	o.Align = &v
}

// GetRuns returns the Runs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatermarkParagraphDto) GetRuns() []WatermarkTextRunDto {
	if o == nil {
		var ret []WatermarkTextRunDto
		return ret
	}
	return o.Runs
}

// GetRunsOk returns a tuple with the Runs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatermarkParagraphDto) GetRunsOk() ([]WatermarkTextRunDto, bool) {
	if o == nil || IsNil(o.Runs) {
		return nil, false
	}
	return o.Runs, true
}

// HasRuns returns a boolean if a field has been set.
func (o *WatermarkParagraphDto) IsRunsSet() bool {
	if o != nil && !IsNil(o.Runs) {
		return true
	}

	return false
}

// SetRuns gets a reference to the given []WatermarkTextRunDto and assigns it to the Runs field.
func (o *WatermarkParagraphDto) SetRuns(v []WatermarkTextRunDto) {
	o.Runs = v
}

func (o WatermarkParagraphDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatermarkParagraphDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Align) {
		toSerialize["align"] = o.Align
	}
	if o.Runs != nil {
		toSerialize["runs"] = o.Runs
	}
	return toSerialize, nil
}

type NullableWatermarkParagraphDto struct {
	value *WatermarkParagraphDto
	isSet bool
}

func (v NullableWatermarkParagraphDto) Get() *WatermarkParagraphDto {
	return v.value
}

func (v *NullableWatermarkParagraphDto) Set(val *WatermarkParagraphDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWatermarkParagraphDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWatermarkParagraphDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatermarkParagraphDto(val *WatermarkParagraphDto) *NullableWatermarkParagraphDto {
	return &NullableWatermarkParagraphDto{value: val, isSet: true}
}

func (v NullableWatermarkParagraphDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatermarkParagraphDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

