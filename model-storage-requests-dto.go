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

// checks if the StorageRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageRequestsDto{}

// StorageRequestsDto The request parameters for configuring the storage module settings.
type StorageRequestsDto struct {
	// The name for the storage module to be configured.
	Module NullableString `json:"module"`
	// The list of configuration key-value pairs for the storage module.
	Props []ItemKeyValuePairStringString `json:"props,omitempty"`
}

type _StorageRequestsDto StorageRequestsDto

// NewStorageRequestsDto instantiates a new StorageRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageRequestsDto(module NullableString) *StorageRequestsDto {
	this := StorageRequestsDto{}
	this.Module = module
	return &this
}

// NewStorageRequestsDtoWithDefaults instantiates a new StorageRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageRequestsDtoWithDefaults() *StorageRequestsDto {
	this := StorageRequestsDto{}
	return &this
}

// GetModule returns the Module field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageRequestsDto) GetModule() string {
	if o == nil || o.Module.Get() == nil {
		var ret string
		return ret
	}

	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageRequestsDto) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// SetModule sets field value
func (o *StorageRequestsDto) SetModule(v string) {
	o.Module.Set(&v)
}

// GetProps returns the Props field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageRequestsDto) GetProps() []ItemKeyValuePairStringString {
	if o == nil {
		var ret []ItemKeyValuePairStringString
		return ret
	}
	return o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageRequestsDto) GetPropsOk() ([]ItemKeyValuePairStringString, bool) {
	if o == nil || IsNil(o.Props) {
		return nil, false
	}
	return o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *StorageRequestsDto) IsPropsSet() bool {
	if o != nil && !IsNil(o.Props) {
		return true
	}

	return false
}

// SetProps gets a reference to the given []ItemKeyValuePairStringString and assigns it to the Props field.
func (o *StorageRequestsDto) SetProps(v []ItemKeyValuePairStringString) {
	o.Props = v
}

func (o StorageRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["module"] = o.Module.Get()
	if o.Props != nil {
		toSerialize["props"] = o.Props
	}
	return toSerialize, nil
}

func (o *StorageRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"module",
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

	varStorageRequestsDto := _StorageRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varStorageRequestsDto)

	if err != nil {
		return err
	}

	*o = StorageRequestsDto(varStorageRequestsDto)

	return err
}

type NullableStorageRequestsDto struct {
	value *StorageRequestsDto
	isSet bool
}

func (v NullableStorageRequestsDto) Get() *StorageRequestsDto {
	return v.value
}

func (v *NullableStorageRequestsDto) Set(val *StorageRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageRequestsDto(val *StorageRequestsDto) *NullableStorageRequestsDto {
	return &NullableStorageRequestsDto{value: val, isSet: true}
}

func (v NullableStorageRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

