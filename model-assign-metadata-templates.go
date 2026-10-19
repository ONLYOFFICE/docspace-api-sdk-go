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

// checks if the AssignMetadataTemplates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignMetadataTemplates{}

// AssignMetadataTemplates The parameters for assigning metadata templates.
type AssignMetadataTemplates struct {
	// The metadata template IDs.
	TemplateIds []int32 `json:"templateIds"`
	// Specifies if the templates are propagated to the folder sub-entries.
	Cascade *bool `json:"cascade,omitempty"`
	// The conflict resolve type for the cascade assignment.
	ConflictResolveType *MetadataConflictResolveType `json:"conflictResolveType,omitempty"`
}

type _AssignMetadataTemplates AssignMetadataTemplates

// NewAssignMetadataTemplates instantiates a new AssignMetadataTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignMetadataTemplates(templateIds []int32) *AssignMetadataTemplates {
	this := AssignMetadataTemplates{}
	this.TemplateIds = templateIds
	return &this
}

// NewAssignMetadataTemplatesWithDefaults instantiates a new AssignMetadataTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignMetadataTemplatesWithDefaults() *AssignMetadataTemplates {
	this := AssignMetadataTemplates{}
	return &this
}

// GetTemplateIds returns the TemplateIds field value
// If the value is explicit nil, the zero value for []int32 will be returned
func (o *AssignMetadataTemplates) GetTemplateIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.TemplateIds
}

// GetTemplateIdsOk returns a tuple with the TemplateIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssignMetadataTemplates) GetTemplateIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.TemplateIds) {
		return nil, false
	}
	return o.TemplateIds, true
}

// SetTemplateIds sets field value
func (o *AssignMetadataTemplates) SetTemplateIds(v []int32) {
	o.TemplateIds = v
}

// GetCascade returns the Cascade field value if set, zero value otherwise.
func (o *AssignMetadataTemplates) GetCascade() bool {
	if o == nil || IsNil(o.Cascade) {
		var ret bool
		return ret
	}
	return *o.Cascade
}

// GetCascadeOk returns a tuple with the Cascade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignMetadataTemplates) GetCascadeOk() (*bool, bool) {
	if o == nil || IsNil(o.Cascade) {
		return nil, false
	}
	return o.Cascade, true
}

// HasCascade returns a boolean if a field has been set.
func (o *AssignMetadataTemplates) IsCascadeSet() bool {
	if o != nil && !IsNil(o.Cascade) {
		return true
	}

	return false
}

// SetCascade gets a reference to the given bool and assigns it to the Cascade field.
func (o *AssignMetadataTemplates) SetCascade(v bool) {
	o.Cascade = &v
}

// GetConflictResolveType returns the ConflictResolveType field value if set, zero value otherwise.
func (o *AssignMetadataTemplates) GetConflictResolveType() MetadataConflictResolveType {
	if o == nil || IsNil(o.ConflictResolveType) {
		var ret MetadataConflictResolveType
		return ret
	}
	return *o.ConflictResolveType
}

// GetConflictResolveTypeOk returns a tuple with the ConflictResolveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignMetadataTemplates) GetConflictResolveTypeOk() (*MetadataConflictResolveType, bool) {
	if o == nil || IsNil(o.ConflictResolveType) {
		return nil, false
	}
	return o.ConflictResolveType, true
}

// HasConflictResolveType returns a boolean if a field has been set.
func (o *AssignMetadataTemplates) IsConflictResolveTypeSet() bool {
	if o != nil && !IsNil(o.ConflictResolveType) {
		return true
	}

	return false
}

// SetConflictResolveType gets a reference to the given MetadataConflictResolveType and assigns it to the ConflictResolveType field.
func (o *AssignMetadataTemplates) SetConflictResolveType(v MetadataConflictResolveType) {
	o.ConflictResolveType = &v
}

func (o AssignMetadataTemplates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignMetadataTemplates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateIds != nil {
		toSerialize["templateIds"] = o.TemplateIds
	}
	if !IsNil(o.Cascade) {
		toSerialize["cascade"] = o.Cascade
	}
	if !IsNil(o.ConflictResolveType) {
		toSerialize["conflictResolveType"] = o.ConflictResolveType
	}
	return toSerialize, nil
}

func (o *AssignMetadataTemplates) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateIds",
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

	varAssignMetadataTemplates := _AssignMetadataTemplates{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAssignMetadataTemplates)

	if err != nil {
		return err
	}

	*o = AssignMetadataTemplates(varAssignMetadataTemplates)

	return err
}

type NullableAssignMetadataTemplates struct {
	value *AssignMetadataTemplates
	isSet bool
}

func (v NullableAssignMetadataTemplates) Get() *AssignMetadataTemplates {
	return v.value
}

func (v *NullableAssignMetadataTemplates) Set(val *AssignMetadataTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignMetadataTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignMetadataTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignMetadataTemplates(val *AssignMetadataTemplates) *NullableAssignMetadataTemplates {
	return &NullableAssignMetadataTemplates{value: val, isSet: true}
}

func (v NullableAssignMetadataTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignMetadataTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

