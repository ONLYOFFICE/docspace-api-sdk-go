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

// checks if the WhiteLabelItemDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WhiteLabelItemDto{}

// WhiteLabelItemDto The white label item parameters.
type WhiteLabelItemDto struct {
	Type *WhiteLabelLogoType `json:"type,omitempty"`
	// The white label file name.
	Name NullableString `json:"name,omitempty"`
	Size *IMagickGeometry `json:"size,omitempty"`
	Path *WhiteLabelItemPathDto `json:"path,omitempty"`
}

// NewWhiteLabelItemDto instantiates a new WhiteLabelItemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhiteLabelItemDto() *WhiteLabelItemDto {
	this := WhiteLabelItemDto{}
	return &this
}

// NewWhiteLabelItemDtoWithDefaults instantiates a new WhiteLabelItemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhiteLabelItemDtoWithDefaults() *WhiteLabelItemDto {
	this := WhiteLabelItemDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhiteLabelItemDto) GetType() WhiteLabelLogoType {
	if o == nil || IsNil(o.Type) {
		var ret WhiteLabelLogoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemDto) GetTypeOk() (*WhiteLabelLogoType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhiteLabelItemDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given WhiteLabelLogoType and assigns it to the Type field.
func (o *WhiteLabelItemDto) SetType(v WhiteLabelLogoType) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WhiteLabelItemDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WhiteLabelItemDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WhiteLabelItemDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WhiteLabelItemDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WhiteLabelItemDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WhiteLabelItemDto) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *WhiteLabelItemDto) GetSize() IMagickGeometry {
	if o == nil || IsNil(o.Size) {
		var ret IMagickGeometry
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemDto) GetSizeOk() (*IMagickGeometry, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *WhiteLabelItemDto) IsSizeSet() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given IMagickGeometry and assigns it to the Size field.
func (o *WhiteLabelItemDto) SetSize(v IMagickGeometry) {
	o.Size = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *WhiteLabelItemDto) GetPath() WhiteLabelItemPathDto {
	if o == nil || IsNil(o.Path) {
		var ret WhiteLabelItemPathDto
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelItemDto) GetPathOk() (*WhiteLabelItemPathDto, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *WhiteLabelItemDto) IsPathSet() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given WhiteLabelItemPathDto and assigns it to the Path field.
func (o *WhiteLabelItemDto) SetPath(v WhiteLabelItemPathDto) {
	o.Path = &v
}

func (o WhiteLabelItemDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WhiteLabelItemDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableWhiteLabelItemDto struct {
	value *WhiteLabelItemDto
	isSet bool
}

func (v NullableWhiteLabelItemDto) Get() *WhiteLabelItemDto {
	return v.value
}

func (v *NullableWhiteLabelItemDto) Set(val *WhiteLabelItemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWhiteLabelItemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWhiteLabelItemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhiteLabelItemDto(val *WhiteLabelItemDto) *NullableWhiteLabelItemDto {
	return &NullableWhiteLabelItemDto{value: val, isSet: true}
}

func (v NullableWhiteLabelItemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhiteLabelItemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

