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

// checks if the LogoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogoRequest{}

// LogoRequest The part of an uploaded picture to use as the logo.
type LogoRequest struct {
	// The picture to cut the logo out of, named by the path that `POST api/2.0/files/logos` returned for it. The  path may be used once and only by the account that uploaded it.
	TmpFile string `json:"tmpFile"`
	// The left edge of the rectangle cut out of the uploaded picture, counted in pixels from its left side. The  picture itself was already scaled down to fit 1280 by 1280 pixels when it was uploaded.
	X *int32 `json:"x,omitempty"`
	// The top edge of the rectangle cut out of the uploaded picture, counted in pixels from its top.
	Y *int32 `json:"y,omitempty"`
	// How wide a piece of the uploaded picture to cut out, in pixels. It has to be sent together with the height,  and the portal builds the four logo sizes out of the piece.
	Width *int32 `json:"width,omitempty"`
	// How tall a piece of the uploaded picture to cut out, in pixels. It has to be sent together with the width.
	Height *int32 `json:"height,omitempty"`
}

type _LogoRequest LogoRequest

// NewLogoRequest instantiates a new LogoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogoRequest(tmpFile string) *LogoRequest {
	this := LogoRequest{}
	this.TmpFile = tmpFile
	return &this
}

// NewLogoRequestWithDefaults instantiates a new LogoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoRequestWithDefaults() *LogoRequest {
	this := LogoRequest{}
	return &this
}

// GetTmpFile returns the TmpFile field value
func (o *LogoRequest) GetTmpFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TmpFile
}

// GetTmpFileOk returns a tuple with the TmpFile field value
// and a boolean to check if the value has been set.
func (o *LogoRequest) GetTmpFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TmpFile, true
}

// SetTmpFile sets field value
func (o *LogoRequest) SetTmpFile(v string) {
	o.TmpFile = v
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
	toSerialize["tmpFile"] = o.TmpFile
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

func (o *LogoRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tmpFile",
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

	varLogoRequest := _LogoRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varLogoRequest)

	if err != nil {
		return err
	}

	*o = LogoRequest(varLogoRequest)

	return err
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

