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

// checks if the LogoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogoRequest{}

// LogoRequest The logo request parameters.
type LogoRequest struct {
	// The path to the temporary image file.
	TmpFile NullableString `json:"tmpFile,omitempty"`
	// The X coordinate of the rectangle starting point.
	X *int32 `json:"x,omitempty"`
	// The Y coordinate of the rectangle starting point.
	Y *int32 `json:"y,omitempty"`
	// The rectangle width.
	Width *int32 `json:"width,omitempty"`
	// The rectangle height.
	Height *int32 `json:"height,omitempty"`
}

// NewLogoRequest instantiates a new LogoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogoRequest() *LogoRequest {
	this := LogoRequest{}
	return &this
}

// NewLogoRequestWithDefaults instantiates a new LogoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoRequestWithDefaults() *LogoRequest {
	this := LogoRequest{}
	return &this
}

// GetTmpFile returns the TmpFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogoRequest) GetTmpFile() string {
	if o == nil || IsNil(o.TmpFile.Get()) {
		var ret string
		return ret
	}
	return *o.TmpFile.Get()
}

// GetTmpFileOk returns a tuple with the TmpFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoRequest) GetTmpFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TmpFile.Get(), o.TmpFile.IsSet()
}

// HasTmpFile returns a boolean if a field has been set.
func (o *LogoRequest) IsTmpFileSet() bool {
	if o != nil && o.TmpFile.IsSet() {
		return true
	}

	return false
}

// SetTmpFile gets a reference to the given NullableString and assigns it to the TmpFile field.
func (o *LogoRequest) SetTmpFile(v string) {
	o.TmpFile.Set(&v)
}
// SetTmpFileNil sets the value for TmpFile to be an explicit nil
func (o *LogoRequest) SetTmpFileNil() {
	o.TmpFile.Set(nil)
}

// UnsetTmpFile ensures that no value is present for TmpFile, not even an explicit nil
func (o *LogoRequest) UnsetTmpFile() {
	o.TmpFile.Unset()
}

// GetX returns the X field value if set, zero value otherwise.
func (o *LogoRequest) GetX() int32 {
	if o == nil || IsNil(o.X) {
		var ret int32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogoRequest) GetXOk() (*int32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *LogoRequest) IsXSet() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given int32 and assigns it to the X field.
func (o *LogoRequest) SetX(v int32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *LogoRequest) GetY() int32 {
	if o == nil || IsNil(o.Y) {
		var ret int32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogoRequest) GetYOk() (*int32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *LogoRequest) IsYSet() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given int32 and assigns it to the Y field.
func (o *LogoRequest) SetY(v int32) {
	o.Y = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *LogoRequest) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogoRequest) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *LogoRequest) IsWidthSet() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *LogoRequest) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *LogoRequest) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogoRequest) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *LogoRequest) IsHeightSet() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *LogoRequest) SetHeight(v int32) {
	o.Height = &v
}

func (o LogoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TmpFile.IsSet() {
		toSerialize["tmpFile"] = o.TmpFile.Get()
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableLogoRequest struct {
	value *LogoRequest
	isSet bool
}

func (v NullableLogoRequest) Get() *LogoRequest {
	return v.value
}

func (v *NullableLogoRequest) Set(val *LogoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLogoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLogoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogoRequest(val *LogoRequest) *NullableLogoRequest {
	return &NullableLogoRequest{value: val, isSet: true}
}

func (v NullableLogoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

