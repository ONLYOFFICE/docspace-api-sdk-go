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
	"time"
)

// checks if the StudioDefaultPageSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudioDefaultPageSettings{}

// StudioDefaultPageSettings struct for StudioDefaultPageSettings
type StudioDefaultPageSettings struct {
	DefaultFolderType *FolderType `json:"defaultFolderType,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewStudioDefaultPageSettings instantiates a new StudioDefaultPageSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudioDefaultPageSettings() *StudioDefaultPageSettings {
	this := StudioDefaultPageSettings{}
	return &this
}

// NewStudioDefaultPageSettingsWithDefaults instantiates a new StudioDefaultPageSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudioDefaultPageSettingsWithDefaults() *StudioDefaultPageSettings {
	this := StudioDefaultPageSettings{}
	return &this
}

// GetDefaultFolderType returns the DefaultFolderType field value if set, zero value otherwise.
func (o *StudioDefaultPageSettings) GetDefaultFolderType() FolderType {
	if o == nil || IsNil(o.DefaultFolderType) {
		var ret FolderType
		return ret
	}
	return *o.DefaultFolderType
}

// GetDefaultFolderTypeOk returns a tuple with the DefaultFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudioDefaultPageSettings) GetDefaultFolderTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.DefaultFolderType) {
		return nil, false
	}
	return o.DefaultFolderType, true
}

// HasDefaultFolderType returns a boolean if a field has been set.
func (o *StudioDefaultPageSettings) IsDefaultFolderTypeSet() bool {
	if o != nil && !IsNil(o.DefaultFolderType) {
		return true
	}

	return false
}

// SetDefaultFolderType gets a reference to the given FolderType and assigns it to the DefaultFolderType field.
func (o *StudioDefaultPageSettings) SetDefaultFolderType(v FolderType) {
	o.DefaultFolderType = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *StudioDefaultPageSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudioDefaultPageSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *StudioDefaultPageSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *StudioDefaultPageSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o StudioDefaultPageSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudioDefaultPageSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultFolderType) {
		toSerialize["defaultFolderType"] = o.DefaultFolderType
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableStudioDefaultPageSettings struct {
	value *StudioDefaultPageSettings
	isSet bool
}

func (v NullableStudioDefaultPageSettings) Get() *StudioDefaultPageSettings {
	return v.value
}

func (v *NullableStudioDefaultPageSettings) Set(val *StudioDefaultPageSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableStudioDefaultPageSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableStudioDefaultPageSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudioDefaultPageSettings(val *StudioDefaultPageSettings) *NullableStudioDefaultPageSettings {
	return &NullableStudioDefaultPageSettings{value: val, isSet: true}
}

func (v NullableStudioDefaultPageSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudioDefaultPageSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

