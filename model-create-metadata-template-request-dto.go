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
	"bytes"
	"fmt"
)

// checks if the CreateMetadataTemplateRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMetadataTemplateRequestDto{}

// CreateMetadataTemplateRequestDto The request parameters for creating a metadata template.
type CreateMetadataTemplateRequestDto struct {
	// The template name.
	Name NullableString `json:"name"`
	// Specifies if the template is visible in the UI pickers.
	Visible *bool `json:"visible,omitempty"`
	// The template metadata fields.
	Fields []MetadataFieldRequest `json:"fields,omitempty"`
}

type _CreateMetadataTemplateRequestDto CreateMetadataTemplateRequestDto

// NewCreateMetadataTemplateRequestDto instantiates a new CreateMetadataTemplateRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMetadataTemplateRequestDto(name NullableString) *CreateMetadataTemplateRequestDto {
	this := CreateMetadataTemplateRequestDto{}
	this.Name = name
	return &this
}

// NewCreateMetadataTemplateRequestDtoWithDefaults instantiates a new CreateMetadataTemplateRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMetadataTemplateRequestDtoWithDefaults() *CreateMetadataTemplateRequestDto {
	this := CreateMetadataTemplateRequestDto{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateMetadataTemplateRequestDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMetadataTemplateRequestDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CreateMetadataTemplateRequestDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *CreateMetadataTemplateRequestDto) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetadataTemplateRequestDto) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *CreateMetadataTemplateRequestDto) IsVisibleSet() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *CreateMetadataTemplateRequestDto) SetVisible(v bool) {
	o.Visible = &v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMetadataTemplateRequestDto) GetFields() []MetadataFieldRequest {
	if o == nil {
		var ret []MetadataFieldRequest
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMetadataTemplateRequestDto) GetFieldsOk() ([]MetadataFieldRequest, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CreateMetadataTemplateRequestDto) IsFieldsSet() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []MetadataFieldRequest and assigns it to the Fields field.
func (o *CreateMetadataTemplateRequestDto) SetFields(v []MetadataFieldRequest) {
	o.Fields = v
}

func (o CreateMetadataTemplateRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMetadataTemplateRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

func (o *CreateMetadataTemplateRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateMetadataTemplateRequestDto := _CreateMetadataTemplateRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateMetadataTemplateRequestDto)

	if err != nil {
		return err
	}

	*o = CreateMetadataTemplateRequestDto(varCreateMetadataTemplateRequestDto)

	return err
}

type NullableCreateMetadataTemplateRequestDto struct {
	value *CreateMetadataTemplateRequestDto
	isSet bool
}

func (v NullableCreateMetadataTemplateRequestDto) Get() *CreateMetadataTemplateRequestDto {
	return v.value
}

func (v *NullableCreateMetadataTemplateRequestDto) Set(val *CreateMetadataTemplateRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMetadataTemplateRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMetadataTemplateRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMetadataTemplateRequestDto(val *CreateMetadataTemplateRequestDto) *NullableCreateMetadataTemplateRequestDto {
	return &NullableCreateMetadataTemplateRequestDto{value: val, isSet: true}
}

func (v NullableCreateMetadataTemplateRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMetadataTemplateRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

