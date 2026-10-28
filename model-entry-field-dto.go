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

// checks if the EntryFieldDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntryFieldDto{}

// EntryFieldDto A metadata template field with its value on the entry.
type EntryFieldDto struct {
	// The field ID.
	Id *int32 `json:"id,omitempty"`
	// The field name.
	Name NullableString `json:"name,omitempty"`
	// The field type.
	Type *MetadataFieldType `json:"type,omitempty"`
	// The choice options of the field.
	Options []MetadataFieldOptionDto `json:"options,omitempty"`
	// The field display order inside the template.
	Order *int32 `json:"order,omitempty"`
	// The value of the field on the entry, or `null` when the entry holds no value for it.
	Value *MetadataValueDto `json:"value,omitempty"`
}

// NewEntryFieldDto instantiates a new EntryFieldDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryFieldDto() *EntryFieldDto {
	this := EntryFieldDto{}
	return &this
}

// NewEntryFieldDtoWithDefaults instantiates a new EntryFieldDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryFieldDtoWithDefaults() *EntryFieldDto {
	this := EntryFieldDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntryFieldDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryFieldDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntryFieldDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EntryFieldDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntryFieldDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntryFieldDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EntryFieldDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EntryFieldDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EntryFieldDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EntryFieldDto) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntryFieldDto) GetType() MetadataFieldType {
	if o == nil || IsNil(o.Type) {
		var ret MetadataFieldType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryFieldDto) GetTypeOk() (*MetadataFieldType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntryFieldDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given MetadataFieldType and assigns it to the Type field.
func (o *EntryFieldDto) SetType(v MetadataFieldType) {
	o.Type = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntryFieldDto) GetOptions() []MetadataFieldOptionDto {
	if o == nil {
		var ret []MetadataFieldOptionDto
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntryFieldDto) GetOptionsOk() ([]MetadataFieldOptionDto, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *EntryFieldDto) IsOptionsSet() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []MetadataFieldOptionDto and assigns it to the Options field.
func (o *EntryFieldDto) SetOptions(v []MetadataFieldOptionDto) {
	o.Options = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *EntryFieldDto) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryFieldDto) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *EntryFieldDto) IsOrderSet() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *EntryFieldDto) SetOrder(v int32) {
	o.Order = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntryFieldDto) GetValue() MetadataValueDto {
	if o == nil || IsNil(o.Value) {
		var ret MetadataValueDto
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryFieldDto) GetValueOk() (*MetadataValueDto, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntryFieldDto) IsValueSet() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given MetadataValueDto and assigns it to the Value field.
func (o *EntryFieldDto) SetValue(v MetadataValueDto) {
	o.Value = &v
}

func (o EntryFieldDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntryFieldDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableEntryFieldDto struct {
	value *EntryFieldDto
	isSet bool
}

func (v NullableEntryFieldDto) Get() *EntryFieldDto {
	return v.value
}

func (v *NullableEntryFieldDto) Set(val *EntryFieldDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryFieldDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryFieldDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryFieldDto(val *EntryFieldDto) *NullableEntryFieldDto {
	return &NullableEntryFieldDto{value: val, isSet: true}
}

func (v NullableEntryFieldDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryFieldDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

