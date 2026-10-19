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

// checks if the ConvertThumbnailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertThumbnailRequest{}

// ConvertThumbnailRequest How a thumbnail is fitted. A field left out takes the document service default.
type ConvertThumbnailRequest struct {
	// The mode to fit the image to the height and width specified:  0 - stretch file to fit height and width;  1 - keep the aspect for the image;  2 - convert the metric size of the page into pixels at 96 dpi.
	Aspect NullableInt32 `json:"aspect,omitempty"`
	// Whether a thumbnail is made of the first page only; false makes one of every page.
	First NullableBool `json:"first,omitempty"`
	// The thumbnail height in pixels; not used with aspect 2.
	Height NullableInt32 `json:"height,omitempty"`
	// The thumbnail width in pixels; not used with aspect 2.
	Width NullableInt32 `json:"width,omitempty"`
}

// NewConvertThumbnailRequest instantiates a new ConvertThumbnailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertThumbnailRequest() *ConvertThumbnailRequest {
	this := ConvertThumbnailRequest{}
	return &this
}

// NewConvertThumbnailRequestWithDefaults instantiates a new ConvertThumbnailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertThumbnailRequestWithDefaults() *ConvertThumbnailRequest {
	this := ConvertThumbnailRequest{}
	return &this
}

// GetAspect returns the Aspect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertThumbnailRequest) GetAspect() int32 {
	if o == nil || IsNil(o.Aspect.Get()) {
		var ret int32
		return ret
	}
	return *o.Aspect.Get()
}

// GetAspectOk returns a tuple with the Aspect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertThumbnailRequest) GetAspectOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aspect.Get(), o.Aspect.IsSet()
}

// HasAspect returns a boolean if a field has been set.
func (o *ConvertThumbnailRequest) IsAspectSet() bool {
	if o != nil && o.Aspect.IsSet() {
		return true
	}

	return false
}

// SetAspect gets a reference to the given NullableInt32 and assigns it to the Aspect field.
func (o *ConvertThumbnailRequest) SetAspect(v int32) {
	o.Aspect.Set(&v)
}
// SetAspectNil sets the value for Aspect to be an explicit nil
func (o *ConvertThumbnailRequest) SetAspectNil() {
	o.Aspect.Set(nil)
}

// UnsetAspect ensures that no value is present for Aspect, not even an explicit nil
func (o *ConvertThumbnailRequest) UnsetAspect() {
	o.Aspect.Unset()
}

// GetFirst returns the First field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertThumbnailRequest) GetFirst() bool {
	if o == nil || IsNil(o.First.Get()) {
		var ret bool
		return ret
	}
	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertThumbnailRequest) GetFirstOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// HasFirst returns a boolean if a field has been set.
func (o *ConvertThumbnailRequest) IsFirstSet() bool {
	if o != nil && o.First.IsSet() {
		return true
	}

	return false
}

// SetFirst gets a reference to the given NullableBool and assigns it to the First field.
func (o *ConvertThumbnailRequest) SetFirst(v bool) {
	o.First.Set(&v)
}
// SetFirstNil sets the value for First to be an explicit nil
func (o *ConvertThumbnailRequest) SetFirstNil() {
	o.First.Set(nil)
}

// UnsetFirst ensures that no value is present for First, not even an explicit nil
func (o *ConvertThumbnailRequest) UnsetFirst() {
	o.First.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertThumbnailRequest) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertThumbnailRequest) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *ConvertThumbnailRequest) IsHeightSet() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *ConvertThumbnailRequest) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *ConvertThumbnailRequest) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *ConvertThumbnailRequest) UnsetHeight() {
	o.Height.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertThumbnailRequest) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertThumbnailRequest) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *ConvertThumbnailRequest) IsWidthSet() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *ConvertThumbnailRequest) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *ConvertThumbnailRequest) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *ConvertThumbnailRequest) UnsetWidth() {
	o.Width.Unset()
}

func (o ConvertThumbnailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertThumbnailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Aspect.IsSet() {
		toSerialize["aspect"] = o.Aspect.Get()
	}
	if o.First.IsSet() {
		toSerialize["first"] = o.First.Get()
	}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}
	return toSerialize, nil
}

type NullableConvertThumbnailRequest struct {
	value *ConvertThumbnailRequest
	isSet bool
}

func (v NullableConvertThumbnailRequest) Get() *ConvertThumbnailRequest {
	return v.value
}

func (v *NullableConvertThumbnailRequest) Set(val *ConvertThumbnailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertThumbnailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertThumbnailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertThumbnailRequest(val *ConvertThumbnailRequest) *NullableConvertThumbnailRequest {
	return &NullableConvertThumbnailRequest{value: val, isSet: true}
}

func (v NullableConvertThumbnailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertThumbnailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

