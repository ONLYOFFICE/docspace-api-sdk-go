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

// checks if the FormMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormMetadata{}

// FormMetadata The metadata of a single form field.
type FormMetadata struct {
	// The form field key.
	Key NullableString `json:"key,omitempty"`
	// The form field type.
	Type NullableString `json:"type,omitempty"`
	// The form field format.
	Format NullableString `json:"format,omitempty"`
	// The list of possible values for the form field.
	PossibleValues []string `json:"possibleValues,omitempty"`
}

// NewFormMetadata instantiates a new FormMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormMetadata() *FormMetadata {
	this := FormMetadata{}
	return &this
}

// NewFormMetadataWithDefaults instantiates a new FormMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormMetadataWithDefaults() *FormMetadata {
	this := FormMetadata{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormMetadata) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormMetadata) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *FormMetadata) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *FormMetadata) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *FormMetadata) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *FormMetadata) UnsetKey() {
	o.Key.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormMetadata) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormMetadata) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *FormMetadata) IsTypeSet() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *FormMetadata) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *FormMetadata) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *FormMetadata) UnsetType() {
	o.Type.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormMetadata) GetFormat() string {
	if o == nil || IsNil(o.Format.Get()) {
		var ret string
		return ret
	}
	return *o.Format.Get()
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormMetadata) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Format.Get(), o.Format.IsSet()
}

// HasFormat returns a boolean if a field has been set.
func (o *FormMetadata) IsFormatSet() bool {
	if o != nil && o.Format.IsSet() {
		return true
	}

	return false
}

// SetFormat gets a reference to the given NullableString and assigns it to the Format field.
func (o *FormMetadata) SetFormat(v string) {
	o.Format.Set(&v)
}
// SetFormatNil sets the value for Format to be an explicit nil
func (o *FormMetadata) SetFormatNil() {
	o.Format.Set(nil)
}

// UnsetFormat ensures that no value is present for Format, not even an explicit nil
func (o *FormMetadata) UnsetFormat() {
	o.Format.Unset()
}

// GetPossibleValues returns the PossibleValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormMetadata) GetPossibleValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PossibleValues
}

// GetPossibleValuesOk returns a tuple with the PossibleValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormMetadata) GetPossibleValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.PossibleValues) {
		return nil, false
	}
	return o.PossibleValues, true
}

// HasPossibleValues returns a boolean if a field has been set.
func (o *FormMetadata) IsPossibleValuesSet() bool {
	if o != nil && !IsNil(o.PossibleValues) {
		return true
	}

	return false
}

// SetPossibleValues gets a reference to the given []string and assigns it to the PossibleValues field.
func (o *FormMetadata) SetPossibleValues(v []string) {
	o.PossibleValues = v
}

func (o FormMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Format.IsSet() {
		toSerialize["format"] = o.Format.Get()
	}
	if o.PossibleValues != nil {
		toSerialize["possibleValues"] = o.PossibleValues
	}
	return toSerialize, nil
}

type NullableFormMetadata struct {
	value *FormMetadata
	isSet bool
}

func (v NullableFormMetadata) Get() *FormMetadata {
	return v.value
}

func (v *NullableFormMetadata) Set(val *FormMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableFormMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableFormMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormMetadata(val *FormMetadata) *NullableFormMetadata {
	return &NullableFormMetadata{value: val, isSet: true}
}

func (v NullableFormMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

