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

// checks if the UpdateFolderMetadataTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFolderMetadataTemplateRequest{}

// UpdateFolderMetadataTemplateRequest The new state of a template assignment on a folder: whether the template cascades to the content of the folder and how.
type UpdateFolderMetadataTemplateRequest struct {
	// Whether the template cascades from the folder to the folders and files below it: true turns the cascade on,  false turns it off while the template stays on the folder.
	Cascade *bool `json:"cascade,omitempty"`
	// How the cascade treats a value an entry below already holds.
	ConflictResolveType *MetadataConflictResolveType `json:"conflictResolveType,omitempty"`
}

// NewUpdateFolderMetadataTemplateRequest instantiates a new UpdateFolderMetadataTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFolderMetadataTemplateRequest() *UpdateFolderMetadataTemplateRequest {
	this := UpdateFolderMetadataTemplateRequest{}
	return &this
}

// NewUpdateFolderMetadataTemplateRequestWithDefaults instantiates a new UpdateFolderMetadataTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFolderMetadataTemplateRequestWithDefaults() *UpdateFolderMetadataTemplateRequest {
	this := UpdateFolderMetadataTemplateRequest{}
	return &this
}

// GetCascade returns the Cascade field value if set, zero value otherwise.
func (o *UpdateFolderMetadataTemplateRequest) GetCascade() bool {
	if o == nil || IsNil(o.Cascade) {
		var ret bool
		return ret
	}
	return *o.Cascade
}

// GetCascadeOk returns a tuple with the Cascade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFolderMetadataTemplateRequest) GetCascadeOk() (*bool, bool) {
	if o == nil || IsNil(o.Cascade) {
		return nil, false
	}
	return o.Cascade, true
}

// HasCascade returns a boolean if a field has been set.
func (o *UpdateFolderMetadataTemplateRequest) IsCascadeSet() bool {
	if o != nil && !IsNil(o.Cascade) {
		return true
	}

	return false
}

// SetCascade gets a reference to the given bool and assigns it to the Cascade field.
func (o *UpdateFolderMetadataTemplateRequest) SetCascade(v bool) {
	o.Cascade = &v
}

// GetConflictResolveType returns the ConflictResolveType field value if set, zero value otherwise.
func (o *UpdateFolderMetadataTemplateRequest) GetConflictResolveType() MetadataConflictResolveType {
	if o == nil || IsNil(o.ConflictResolveType) {
		var ret MetadataConflictResolveType
		return ret
	}
	return *o.ConflictResolveType
}

// GetConflictResolveTypeOk returns a tuple with the ConflictResolveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFolderMetadataTemplateRequest) GetConflictResolveTypeOk() (*MetadataConflictResolveType, bool) {
	if o == nil || IsNil(o.ConflictResolveType) {
		return nil, false
	}
	return o.ConflictResolveType, true
}

// HasConflictResolveType returns a boolean if a field has been set.
func (o *UpdateFolderMetadataTemplateRequest) IsConflictResolveTypeSet() bool {
	if o != nil && !IsNil(o.ConflictResolveType) {
		return true
	}

	return false
}

// SetConflictResolveType gets a reference to the given MetadataConflictResolveType and assigns it to the ConflictResolveType field.
func (o *UpdateFolderMetadataTemplateRequest) SetConflictResolveType(v MetadataConflictResolveType) {
	o.ConflictResolveType = &v
}

func (o UpdateFolderMetadataTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFolderMetadataTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cascade) {
		toSerialize["cascade"] = o.Cascade
	}
	if !IsNil(o.ConflictResolveType) {
		toSerialize["conflictResolveType"] = o.ConflictResolveType
	}
	return toSerialize, nil
}

type NullableUpdateFolderMetadataTemplateRequest struct {
	value *UpdateFolderMetadataTemplateRequest
	isSet bool
}

func (v NullableUpdateFolderMetadataTemplateRequest) Get() *UpdateFolderMetadataTemplateRequest {
	return v.value
}

func (v *NullableUpdateFolderMetadataTemplateRequest) Set(val *UpdateFolderMetadataTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFolderMetadataTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFolderMetadataTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFolderMetadataTemplateRequest(val *UpdateFolderMetadataTemplateRequest) *NullableUpdateFolderMetadataTemplateRequest {
	return &NullableUpdateFolderMetadataTemplateRequest{value: val, isSet: true}
}

func (v NullableUpdateFolderMetadataTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFolderMetadataTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

