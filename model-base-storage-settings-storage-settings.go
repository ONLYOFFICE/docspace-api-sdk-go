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

// checks if the BaseStorageSettingsStorageSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseStorageSettingsStorageSettings{}

// BaseStorageSettingsStorageSettings struct for BaseStorageSettingsStorageSettings
type BaseStorageSettingsStorageSettings struct {
	Module NullableString `json:"module,omitempty"`
	Props map[string]string `json:"props,omitempty"`
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewBaseStorageSettingsStorageSettings instantiates a new BaseStorageSettingsStorageSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseStorageSettingsStorageSettings() *BaseStorageSettingsStorageSettings {
	this := BaseStorageSettingsStorageSettings{}
	return &this
}

// NewBaseStorageSettingsStorageSettingsWithDefaults instantiates a new BaseStorageSettingsStorageSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseStorageSettingsStorageSettingsWithDefaults() *BaseStorageSettingsStorageSettings {
	this := BaseStorageSettingsStorageSettings{}
	return &this
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseStorageSettingsStorageSettings) GetModule() string {
	if o == nil || IsNil(o.Module.Get()) {
		var ret string
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseStorageSettingsStorageSettings) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *BaseStorageSettingsStorageSettings) IsModuleSet() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableString and assigns it to the Module field.
func (o *BaseStorageSettingsStorageSettings) SetModule(v string) {
	o.Module.Set(&v)
}
// SetModuleNil sets the value for Module to be an explicit nil
func (o *BaseStorageSettingsStorageSettings) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *BaseStorageSettingsStorageSettings) UnsetModule() {
	o.Module.Unset()
}

// GetProps returns the Props field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseStorageSettingsStorageSettings) GetProps() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseStorageSettingsStorageSettings) GetPropsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Props) {
		return nil, false
	}
	return &o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *BaseStorageSettingsStorageSettings) IsPropsSet() bool {
	if o != nil && !IsNil(o.Props) {
		return true
	}

	return false
}

// SetProps gets a reference to the given map[string]string and assigns it to the Props field.
func (o *BaseStorageSettingsStorageSettings) SetProps(v map[string]string) {
	o.Props = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *BaseStorageSettingsStorageSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStorageSettingsStorageSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *BaseStorageSettingsStorageSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *BaseStorageSettingsStorageSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o BaseStorageSettingsStorageSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseStorageSettingsStorageSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if o.Props != nil {
		toSerialize["props"] = o.Props
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableBaseStorageSettingsStorageSettings struct {
	value *BaseStorageSettingsStorageSettings
	isSet bool
}

func (v NullableBaseStorageSettingsStorageSettings) Get() *BaseStorageSettingsStorageSettings {
	return v.value
}

func (v *NullableBaseStorageSettingsStorageSettings) Set(val *BaseStorageSettingsStorageSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseStorageSettingsStorageSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseStorageSettingsStorageSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseStorageSettingsStorageSettings(val *BaseStorageSettingsStorageSettings) *NullableBaseStorageSettingsStorageSettings {
	return &NullableBaseStorageSettingsStorageSettings{value: val, isSet: true}
}

func (v NullableBaseStorageSettingsStorageSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseStorageSettingsStorageSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

