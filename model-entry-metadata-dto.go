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

// checks if the EntryMetadataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntryMetadataDto{}

// EntryMetadataDto The metadata of an entry: the assigned templates with their values, and the custom fields holding a value.
type EntryMetadataDto struct {
	// The assigned metadata templates, each field carrying its value on the entry.
	Templates []EntryTemplateDto `json:"templates,omitempty"`
	// The custom fields with their values.
	CustomFields []CustomFieldValueDto `json:"customFields,omitempty"`
}

// NewEntryMetadataDto instantiates a new EntryMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryMetadataDto() *EntryMetadataDto {
	this := EntryMetadataDto{}
	return &this
}

// NewEntryMetadataDtoWithDefaults instantiates a new EntryMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryMetadataDtoWithDefaults() *EntryMetadataDto {
	this := EntryMetadataDto{}
	return &this
}

// GetTemplates returns the Templates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntryMetadataDto) GetTemplates() []EntryTemplateDto {
	if o == nil {
		var ret []EntryTemplateDto
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntryMetadataDto) GetTemplatesOk() ([]EntryTemplateDto, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *EntryMetadataDto) IsTemplatesSet() bool {
	if o != nil && !IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []EntryTemplateDto and assigns it to the Templates field.
func (o *EntryMetadataDto) SetTemplates(v []EntryTemplateDto) {
	o.Templates = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntryMetadataDto) GetCustomFields() []CustomFieldValueDto {
	if o == nil {
		var ret []CustomFieldValueDto
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntryMetadataDto) GetCustomFieldsOk() ([]CustomFieldValueDto, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *EntryMetadataDto) IsCustomFieldsSet() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomFieldValueDto and assigns it to the CustomFields field.
func (o *EntryMetadataDto) SetCustomFields(v []CustomFieldValueDto) {
	o.CustomFields = v
}

func (o EntryMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntryMetadataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if o.CustomFields != nil {
		toSerialize["customFields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableEntryMetadataDto struct {
	value *EntryMetadataDto
	isSet bool
}

func (v NullableEntryMetadataDto) Get() *EntryMetadataDto {
	return v.value
}

func (v *NullableEntryMetadataDto) Set(val *EntryMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryMetadataDto(val *EntryMetadataDto) *NullableEntryMetadataDto {
	return &NullableEntryMetadataDto{value: val, isSet: true}
}

func (v NullableEntryMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

