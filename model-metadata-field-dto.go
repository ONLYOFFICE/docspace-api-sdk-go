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

// checks if the MetadataFieldDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataFieldDto{}

// MetadataFieldDto The metadata field information.
type MetadataFieldDto struct {
	// The field ID.
	Id *int32 `json:"id,omitempty"`
	// The ID of the template the field belongs to.
	TemplateId *int32 `json:"templateId,omitempty"`
	// The field name.
	Name NullableString `json:"name,omitempty"`
	// The field type.
	Type *MetadataFieldType `json:"type,omitempty"`
	// The choice options of the field.
	Options []MetadataFieldOptionDto `json:"options,omitempty"`
	// The field display order inside the template.
	Order *int32 `json:"order,omitempty"`
}

// NewMetadataFieldDto instantiates a new MetadataFieldDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataFieldDto() *MetadataFieldDto {
	this := MetadataFieldDto{}
	return &this
}

// NewMetadataFieldDtoWithDefaults instantiates a new MetadataFieldDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataFieldDtoWithDefaults() *MetadataFieldDto {
	this := MetadataFieldDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetadataFieldDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataFieldDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetadataFieldDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MetadataFieldDto) SetId(v int32) {
	o.Id = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *MetadataFieldDto) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataFieldDto) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *MetadataFieldDto) IsTemplateIdSet() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *MetadataFieldDto) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFieldDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFieldDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MetadataFieldDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MetadataFieldDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MetadataFieldDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MetadataFieldDto) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetadataFieldDto) GetType() MetadataFieldType {
	if o == nil || IsNil(o.Type) {
		var ret MetadataFieldType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataFieldDto) GetTypeOk() (*MetadataFieldType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetadataFieldDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given MetadataFieldType and assigns it to the Type field.
func (o *MetadataFieldDto) SetType(v MetadataFieldType) {
	o.Type = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFieldDto) GetOptions() []MetadataFieldOptionDto {
	if o == nil {
		var ret []MetadataFieldOptionDto
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFieldDto) GetOptionsOk() ([]MetadataFieldOptionDto, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *MetadataFieldDto) IsOptionsSet() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []MetadataFieldOptionDto and assigns it to the Options field.
func (o *MetadataFieldDto) SetOptions(v []MetadataFieldOptionDto) {
	o.Options = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *MetadataFieldDto) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataFieldDto) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *MetadataFieldDto) IsOrderSet() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *MetadataFieldDto) SetOrder(v int32) {
	o.Order = &v
}

func (o MetadataFieldDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataFieldDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
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
	return toSerialize, nil
}

type NullableMetadataFieldDto struct {
	value *MetadataFieldDto
	isSet bool
}

func (v NullableMetadataFieldDto) Get() *MetadataFieldDto {
	return v.value
}

func (v *NullableMetadataFieldDto) Set(val *MetadataFieldDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataFieldDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataFieldDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataFieldDto(val *MetadataFieldDto) *NullableMetadataFieldDto {
	return &NullableMetadataFieldDto{value: val, isSet: true}
}

func (v NullableMetadataFieldDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataFieldDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

