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

// checks if the MetadataTemplateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataTemplateDto{}

// MetadataTemplateDto The metadata template information.
type MetadataTemplateDto struct {
	// The template ID.
	Id *int32 `json:"id,omitempty"`
	// The template name.
	Name NullableString `json:"name,omitempty"`
	// Specifies if the template is visible in the UI pickers.
	Visible *bool `json:"visible,omitempty"`
	// The user who created the template. A user who has since left the portal comes back as the lost user placeholder.
	CreateBy *EmployeeDto `json:"createBy,omitempty"`
	// The template creation date.
	CreateOn *ApiDateTime `json:"createOn,omitempty"`
	// The user who changed the template itself last: its name or its visibility. The creator until somebody changes  it; a change of a field is recorded on the field, not here.
	ModifiedBy *EmployeeDto `json:"modifiedBy,omitempty"`
	// The date when the template was modified last.
	ModifiedOn *ApiDateTime `json:"modifiedOn,omitempty"`
	// The template metadata fields.
	Fields []MetadataFieldDto `json:"fields,omitempty"`
}

// NewMetadataTemplateDto instantiates a new MetadataTemplateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataTemplateDto() *MetadataTemplateDto {
	this := MetadataTemplateDto{}
	return &this
}

// NewMetadataTemplateDtoWithDefaults instantiates a new MetadataTemplateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataTemplateDtoWithDefaults() *MetadataTemplateDto {
	this := MetadataTemplateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetadataTemplateDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataTemplateDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetadataTemplateDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MetadataTemplateDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataTemplateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataTemplateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MetadataTemplateDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MetadataTemplateDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MetadataTemplateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MetadataTemplateDto) UnsetName() {
	o.Name.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *MetadataTemplateDto) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataTemplateDto) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *MetadataTemplateDto) IsVisibleSet() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *MetadataTemplateDto) SetVisible(v bool) {
	o.Visible = &v
}

// GetCreateBy returns the CreateBy field value if set, zero value otherwise.
func (o *MetadataTemplateDto) GetCreateBy() EmployeeDto {
	if o == nil || IsNil(o.CreateBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreateBy
}

// GetCreateByOk returns a tuple with the CreateBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataTemplateDto) GetCreateByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreateBy) {
		return nil, false
	}
	return o.CreateBy, true
}

// HasCreateBy returns a boolean if a field has been set.
func (o *MetadataTemplateDto) IsCreateBySet() bool {
	if o != nil && !IsNil(o.CreateBy) {
		return true
	}

	return false
}

// SetCreateBy gets a reference to the given EmployeeDto and assigns it to the CreateBy field.
func (o *MetadataTemplateDto) SetCreateBy(v EmployeeDto) {
	o.CreateBy = &v
}

// GetCreateOn returns the CreateOn field value if set, zero value otherwise.
func (o *MetadataTemplateDto) GetCreateOn() ApiDateTime {
	if o == nil || IsNil(o.CreateOn) {
		var ret ApiDateTime
		return ret
	}
	return *o.CreateOn
}

// GetCreateOnOk returns a tuple with the CreateOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataTemplateDto) GetCreateOnOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.CreateOn) {
		return nil, false
	}
	return o.CreateOn, true
}

// HasCreateOn returns a boolean if a field has been set.
func (o *MetadataTemplateDto) IsCreateOnSet() bool {
	if o != nil && !IsNil(o.CreateOn) {
		return true
	}

	return false
}

// SetCreateOn gets a reference to the given ApiDateTime and assigns it to the CreateOn field.
func (o *MetadataTemplateDto) SetCreateOn(v ApiDateTime) {
	o.CreateOn = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *MetadataTemplateDto) GetModifiedBy() EmployeeDto {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataTemplateDto) GetModifiedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *MetadataTemplateDto) IsModifiedBySet() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given EmployeeDto and assigns it to the ModifiedBy field.
func (o *MetadataTemplateDto) SetModifiedBy(v EmployeeDto) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *MetadataTemplateDto) GetModifiedOn() ApiDateTime {
	if o == nil || IsNil(o.ModifiedOn) {
		var ret ApiDateTime
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataTemplateDto) GetModifiedOnOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ModifiedOn) {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *MetadataTemplateDto) IsModifiedOnSet() bool {
	if o != nil && !IsNil(o.ModifiedOn) {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given ApiDateTime and assigns it to the ModifiedOn field.
func (o *MetadataTemplateDto) SetModifiedOn(v ApiDateTime) {
	o.ModifiedOn = &v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataTemplateDto) GetFields() []MetadataFieldDto {
	if o == nil {
		var ret []MetadataFieldDto
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataTemplateDto) GetFieldsOk() ([]MetadataFieldDto, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MetadataTemplateDto) IsFieldsSet() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []MetadataFieldDto and assigns it to the Fields field.
func (o *MetadataTemplateDto) SetFields(v []MetadataFieldDto) {
	o.Fields = v
}

func (o MetadataTemplateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataTemplateDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CreateBy) {
		toSerialize["createBy"] = o.CreateBy
	}
	if !IsNil(o.CreateOn) {
		toSerialize["createOn"] = o.CreateOn
	}
	if !IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if !IsNil(o.ModifiedOn) {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableMetadataTemplateDto struct {
	value *MetadataTemplateDto
	isSet bool
}

func (v NullableMetadataTemplateDto) Get() *MetadataTemplateDto {
	return v.value
}

func (v *NullableMetadataTemplateDto) Set(val *MetadataTemplateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataTemplateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataTemplateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataTemplateDto(val *MetadataTemplateDto) *NullableMetadataTemplateDto {
	return &NullableMetadataTemplateDto{value: val, isSet: true}
}

func (v NullableMetadataTemplateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataTemplateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

