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

// checks if the EntryTemplateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntryTemplateDto{}

// EntryTemplateDto A metadata template assigned to an entry: the template with every field of it, each field carrying its value on the entry.
type EntryTemplateDto struct {
	// The template ID.
	Id *int32 `json:"id,omitempty"`
	// The template name.
	Name NullableString `json:"name,omitempty"`
	// Specifies if the template is visible in the UI pickers.
	Visible *bool `json:"visible,omitempty"`
	// Whether the template cascades from this entry to the folders and files below it. Only a folder or a room can  cascade; on a file the value is always false.
	Cascade *bool `json:"cascade,omitempty"`
	// How the cascade of this entry treats a value an entry below already holds. Null while the template does not  cascade from this entry.
	ConflictResolveType *MetadataConflictResolveType `json:"conflictResolveType,omitempty"`
	// The template fields with their values on the entry.
	Fields []EntryFieldDto `json:"fields,omitempty"`
}

// NewEntryTemplateDto instantiates a new EntryTemplateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryTemplateDto() *EntryTemplateDto {
	this := EntryTemplateDto{}
	return &this
}

// NewEntryTemplateDtoWithDefaults instantiates a new EntryTemplateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryTemplateDtoWithDefaults() *EntryTemplateDto {
	this := EntryTemplateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntryTemplateDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplateDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntryTemplateDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EntryTemplateDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntryTemplateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntryTemplateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EntryTemplateDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EntryTemplateDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EntryTemplateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EntryTemplateDto) UnsetName() {
	o.Name.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *EntryTemplateDto) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplateDto) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *EntryTemplateDto) IsVisibleSet() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *EntryTemplateDto) SetVisible(v bool) {
	o.Visible = &v
}

// GetCascade returns the Cascade field value if set, zero value otherwise.
func (o *EntryTemplateDto) GetCascade() bool {
	if o == nil || IsNil(o.Cascade) {
		var ret bool
		return ret
	}
	return *o.Cascade
}

// GetCascadeOk returns a tuple with the Cascade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplateDto) GetCascadeOk() (*bool, bool) {
	if o == nil || IsNil(o.Cascade) {
		return nil, false
	}
	return o.Cascade, true
}

// HasCascade returns a boolean if a field has been set.
func (o *EntryTemplateDto) IsCascadeSet() bool {
	if o != nil && !IsNil(o.Cascade) {
		return true
	}

	return false
}

// SetCascade gets a reference to the given bool and assigns it to the Cascade field.
func (o *EntryTemplateDto) SetCascade(v bool) {
	o.Cascade = &v
}

// GetConflictResolveType returns the ConflictResolveType field value if set, zero value otherwise.
func (o *EntryTemplateDto) GetConflictResolveType() MetadataConflictResolveType {
	if o == nil || IsNil(o.ConflictResolveType) {
		var ret MetadataConflictResolveType
		return ret
	}
	return *o.ConflictResolveType
}

// GetConflictResolveTypeOk returns a tuple with the ConflictResolveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplateDto) GetConflictResolveTypeOk() (*MetadataConflictResolveType, bool) {
	if o == nil || IsNil(o.ConflictResolveType) {
		return nil, false
	}
	return o.ConflictResolveType, true
}

// HasConflictResolveType returns a boolean if a field has been set.
func (o *EntryTemplateDto) IsConflictResolveTypeSet() bool {
	if o != nil && !IsNil(o.ConflictResolveType) {
		return true
	}

	return false
}

// SetConflictResolveType gets a reference to the given MetadataConflictResolveType and assigns it to the ConflictResolveType field.
func (o *EntryTemplateDto) SetConflictResolveType(v MetadataConflictResolveType) {
	o.ConflictResolveType = &v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntryTemplateDto) GetFields() []EntryFieldDto {
	if o == nil {
		var ret []EntryFieldDto
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntryTemplateDto) GetFieldsOk() ([]EntryFieldDto, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *EntryTemplateDto) IsFieldsSet() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []EntryFieldDto and assigns it to the Fields field.
func (o *EntryTemplateDto) SetFields(v []EntryFieldDto) {
	o.Fields = v
}

func (o EntryTemplateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntryTemplateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if !IsNil(o.Cascade) {
		toSerialize["cascade"] = o.Cascade
	}
	if !IsNil(o.ConflictResolveType) {
		toSerialize["conflictResolveType"] = o.ConflictResolveType
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableEntryTemplateDto struct {
	value *EntryTemplateDto
	isSet bool
}

func (v NullableEntryTemplateDto) Get() *EntryTemplateDto {
	return v.value
}

func (v *NullableEntryTemplateDto) Set(val *EntryTemplateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryTemplateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryTemplateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryTemplateDto(val *EntryTemplateDto) *NullableEntryTemplateDto {
	return &NullableEntryTemplateDto{value: val, isSet: true}
}

func (v NullableEntryTemplateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryTemplateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

