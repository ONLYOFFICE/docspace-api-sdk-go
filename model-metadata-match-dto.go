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

// checks if the MetadataMatchDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataMatchDto{}

// MetadataMatchDto Where a text search met the metadata of a file, a folder or a room: one string value that holds the searched text,  with the field it belongs to.
type MetadataMatchDto struct {
	// The ID of the metadata template the matching field belongs to. Null for a custom field, which belongs to no  template the API shows.
	TemplateId NullableInt32 `json:"templateId,omitempty"`
	// The ID of the matching template field. Null for a custom field, which is addressed by its name alone.
	FieldId NullableInt32 `json:"fieldId,omitempty"`
	// The name of the matching field: the template field or the custom field as it is shown in the entry metadata.
	FieldName NullableString `json:"fieldName,omitempty"`
	// The stored value that holds the searched text, as it was written, not only the matching part of it.
	Value NullableString `json:"value,omitempty"`
}

// NewMetadataMatchDto instantiates a new MetadataMatchDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataMatchDto() *MetadataMatchDto {
	this := MetadataMatchDto{}
	return &this
}

// NewMetadataMatchDtoWithDefaults instantiates a new MetadataMatchDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataMatchDtoWithDefaults() *MetadataMatchDto {
	this := MetadataMatchDto{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataMatchDto) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId.Get()) {
		var ret int32
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataMatchDto) GetTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *MetadataMatchDto) IsTemplateIdSet() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableInt32 and assigns it to the TemplateId field.
func (o *MetadataMatchDto) SetTemplateId(v int32) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *MetadataMatchDto) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *MetadataMatchDto) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetFieldId returns the FieldId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataMatchDto) GetFieldId() int32 {
	if o == nil || IsNil(o.FieldId.Get()) {
		var ret int32
		return ret
	}
	return *o.FieldId.Get()
}

// GetFieldIdOk returns a tuple with the FieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataMatchDto) GetFieldIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldId.Get(), o.FieldId.IsSet()
}

// HasFieldId returns a boolean if a field has been set.
func (o *MetadataMatchDto) IsFieldIdSet() bool {
	if o != nil && o.FieldId.IsSet() {
		return true
	}

	return false
}

// SetFieldId gets a reference to the given NullableInt32 and assigns it to the FieldId field.
func (o *MetadataMatchDto) SetFieldId(v int32) {
	o.FieldId.Set(&v)
}
// SetFieldIdNil sets the value for FieldId to be an explicit nil
func (o *MetadataMatchDto) SetFieldIdNil() {
	o.FieldId.Set(nil)
}

// UnsetFieldId ensures that no value is present for FieldId, not even an explicit nil
func (o *MetadataMatchDto) UnsetFieldId() {
	o.FieldId.Unset()
}

// GetFieldName returns the FieldName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataMatchDto) GetFieldName() string {
	if o == nil || IsNil(o.FieldName.Get()) {
		var ret string
		return ret
	}
	return *o.FieldName.Get()
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataMatchDto) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldName.Get(), o.FieldName.IsSet()
}

// HasFieldName returns a boolean if a field has been set.
func (o *MetadataMatchDto) IsFieldNameSet() bool {
	if o != nil && o.FieldName.IsSet() {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given NullableString and assigns it to the FieldName field.
func (o *MetadataMatchDto) SetFieldName(v string) {
	o.FieldName.Set(&v)
}
// SetFieldNameNil sets the value for FieldName to be an explicit nil
func (o *MetadataMatchDto) SetFieldNameNil() {
	o.FieldName.Set(nil)
}

// UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
func (o *MetadataMatchDto) UnsetFieldName() {
	o.FieldName.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataMatchDto) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataMatchDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *MetadataMatchDto) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *MetadataMatchDto) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *MetadataMatchDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *MetadataMatchDto) UnsetValue() {
	o.Value.Unset()
}

func (o MetadataMatchDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataMatchDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	if o.FieldId.IsSet() {
		toSerialize["fieldId"] = o.FieldId.Get()
	}
	if o.FieldName.IsSet() {
		toSerialize["fieldName"] = o.FieldName.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableMetadataMatchDto struct {
	value *MetadataMatchDto
	isSet bool
}

func (v NullableMetadataMatchDto) Get() *MetadataMatchDto {
	return v.value
}

func (v *NullableMetadataMatchDto) Set(val *MetadataMatchDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataMatchDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataMatchDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataMatchDto(val *MetadataMatchDto) *NullableMetadataMatchDto {
	return &NullableMetadataMatchDto{value: val, isSet: true}
}

func (v NullableMetadataMatchDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataMatchDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

