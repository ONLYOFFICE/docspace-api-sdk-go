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

// checks if the ContentType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentType{}

// ContentType struct for ContentType
type ContentType struct {
	Boundary NullableString `json:"boundary,omitempty"`
	CharSet NullableString `json:"charSet,omitempty"`
	MediaType NullableString `json:"mediaType,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Parameters []interface{} `json:"parameters,omitempty"`
}

// NewContentType instantiates a new ContentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentType() *ContentType {
	this := ContentType{}
	return &this
}

// NewContentTypeWithDefaults instantiates a new ContentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentTypeWithDefaults() *ContentType {
	this := ContentType{}
	return &this
}

// GetBoundary returns the Boundary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetBoundary() string {
	if o == nil || IsNil(o.Boundary.Get()) {
		var ret string
		return ret
	}
	return *o.Boundary.Get()
}

// GetBoundaryOk returns a tuple with the Boundary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetBoundaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Boundary.Get(), o.Boundary.IsSet()
}

// HasBoundary returns a boolean if a field has been set.
func (o *ContentType) IsBoundarySet() bool {
	if o != nil && o.Boundary.IsSet() {
		return true
	}

	return false
}

// SetBoundary gets a reference to the given NullableString and assigns it to the Boundary field.
func (o *ContentType) SetBoundary(v string) {
	o.Boundary.Set(&v)
}
// SetBoundaryNil sets the value for Boundary to be an explicit nil
func (o *ContentType) SetBoundaryNil() {
	o.Boundary.Set(nil)
}

// UnsetBoundary ensures that no value is present for Boundary, not even an explicit nil
func (o *ContentType) UnsetBoundary() {
	o.Boundary.Unset()
}

// GetCharSet returns the CharSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetCharSet() string {
	if o == nil || IsNil(o.CharSet.Get()) {
		var ret string
		return ret
	}
	return *o.CharSet.Get()
}

// GetCharSetOk returns a tuple with the CharSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetCharSetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CharSet.Get(), o.CharSet.IsSet()
}

// HasCharSet returns a boolean if a field has been set.
func (o *ContentType) IsCharSetSet() bool {
	if o != nil && o.CharSet.IsSet() {
		return true
	}

	return false
}

// SetCharSet gets a reference to the given NullableString and assigns it to the CharSet field.
func (o *ContentType) SetCharSet(v string) {
	o.CharSet.Set(&v)
}
// SetCharSetNil sets the value for CharSet to be an explicit nil
func (o *ContentType) SetCharSetNil() {
	o.CharSet.Set(nil)
}

// UnsetCharSet ensures that no value is present for CharSet, not even an explicit nil
func (o *ContentType) UnsetCharSet() {
	o.CharSet.Unset()
}

// GetMediaType returns the MediaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetMediaType() string {
	if o == nil || IsNil(o.MediaType.Get()) {
		var ret string
		return ret
	}
	return *o.MediaType.Get()
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaType.Get(), o.MediaType.IsSet()
}

// HasMediaType returns a boolean if a field has been set.
func (o *ContentType) IsMediaTypeSet() bool {
	if o != nil && o.MediaType.IsSet() {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given NullableString and assigns it to the MediaType field.
func (o *ContentType) SetMediaType(v string) {
	o.MediaType.Set(&v)
}
// SetMediaTypeNil sets the value for MediaType to be an explicit nil
func (o *ContentType) SetMediaTypeNil() {
	o.MediaType.Set(nil)
}

// UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
func (o *ContentType) UnsetMediaType() {
	o.MediaType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ContentType) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ContentType) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ContentType) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ContentType) UnsetName() {
	o.Name.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetParameters() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetParametersOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ContentType) IsParametersSet() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []interface{} and assigns it to the Parameters field.
func (o *ContentType) SetParameters(v []interface{}) {
	o.Parameters = v
}

func (o ContentType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Boundary.IsSet() {
		toSerialize["boundary"] = o.Boundary.Get()
	}
	if o.CharSet.IsSet() {
		toSerialize["charSet"] = o.CharSet.Get()
	}
	if o.MediaType.IsSet() {
		toSerialize["mediaType"] = o.MediaType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableContentType struct {
	value *ContentType
	isSet bool
}

func (v NullableContentType) Get() *ContentType {
	return v.value
}

func (v *NullableContentType) Set(val *ContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentType(val *ContentType) *NullableContentType {
	return &NullableContentType{value: val, isSet: true}
}

func (v NullableContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

