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

// checks if the UpdateMetadataTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMetadataTemplate{}

// UpdateMetadataTemplate The parameters for updating a metadata template.
type UpdateMetadataTemplate struct {
	// The new template name.
	Name NullableString `json:"name,omitempty"`
	// Specifies if the template is visible in the UI pickers.
	Visible NullableBool `json:"visible,omitempty"`
}

// NewUpdateMetadataTemplate instantiates a new UpdateMetadataTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMetadataTemplate() *UpdateMetadataTemplate {
	this := UpdateMetadataTemplate{}
	return &this
}

// NewUpdateMetadataTemplateWithDefaults instantiates a new UpdateMetadataTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMetadataTemplateWithDefaults() *UpdateMetadataTemplate {
	this := UpdateMetadataTemplate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMetadataTemplate) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMetadataTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateMetadataTemplate) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateMetadataTemplate) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateMetadataTemplate) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateMetadataTemplate) UnsetName() {
	o.Name.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMetadataTemplate) GetVisible() bool {
	if o == nil || IsNil(o.Visible.Get()) {
		var ret bool
		return ret
	}
	return *o.Visible.Get()
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMetadataTemplate) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Visible.Get(), o.Visible.IsSet()
}

// HasVisible returns a boolean if a field has been set.
func (o *UpdateMetadataTemplate) IsVisibleSet() bool {
	if o != nil && o.Visible.IsSet() {
		return true
	}

	return false
}

// SetVisible gets a reference to the given NullableBool and assigns it to the Visible field.
func (o *UpdateMetadataTemplate) SetVisible(v bool) {
	o.Visible.Set(&v)
}
// SetVisibleNil sets the value for Visible to be an explicit nil
func (o *UpdateMetadataTemplate) SetVisibleNil() {
	o.Visible.Set(nil)
}

// UnsetVisible ensures that no value is present for Visible, not even an explicit nil
func (o *UpdateMetadataTemplate) UnsetVisible() {
	o.Visible.Unset()
}

func (o UpdateMetadataTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMetadataTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Visible.IsSet() {
		toSerialize["visible"] = o.Visible.Get()
	}
	return toSerialize, nil
}

type NullableUpdateMetadataTemplate struct {
	value *UpdateMetadataTemplate
	isSet bool
}

func (v NullableUpdateMetadataTemplate) Get() *UpdateMetadataTemplate {
	return v.value
}

func (v *NullableUpdateMetadataTemplate) Set(val *UpdateMetadataTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMetadataTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMetadataTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMetadataTemplate(val *UpdateMetadataTemplate) *NullableUpdateMetadataTemplate {
	return &NullableUpdateMetadataTemplate{value: val, isSet: true}
}

func (v NullableUpdateMetadataTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMetadataTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

