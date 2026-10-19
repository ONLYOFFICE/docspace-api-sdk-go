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

// checks if the StorageRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageRequestDto{}

// StorageRequestDto Which storage provider the portal is pointed at, and the credentials it needs.
type StorageRequestDto struct {
	// The storage provider to switch to, by the identifier the matching listing operation reports - `default` for  the built-in local storage. The provider has to be available on the server, which that listing reports as  `isSet`, otherwise the request is refused with 400; sending the module already in use changes nothing.
	Module NullableString `json:"module"`
	// The credentials the provider expects, as the name and value pairs it defines - a bucket, a region and an  access key for an Amazon S3 storage, for instance. Read the expected names from the entry of that provider in  the listing operation; they differ per provider, so there is no fixed set.
	Props []ItemKeyValuePairStringString `json:"props,omitempty"`
}

type _StorageRequestDto StorageRequestDto

// NewStorageRequestDto instantiates a new StorageRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageRequestDto(module NullableString) *StorageRequestDto {
	this := StorageRequestDto{}
	this.Module = module
	return &this
}

// NewStorageRequestDtoWithDefaults instantiates a new StorageRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageRequestDtoWithDefaults() *StorageRequestDto {
	this := StorageRequestDto{}
	return &this
}

// GetModule returns the Module field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageRequestDto) GetModule() string {
	if o == nil || o.Module.Get() == nil {
		var ret string
		return ret
	}

	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageRequestDto) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// SetModule sets field value
func (o *StorageRequestDto) SetModule(v string) {
	o.Module.Set(&v)
}

// GetProps returns the Props field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageRequestDto) GetProps() []ItemKeyValuePairStringString {
	if o == nil {
		var ret []ItemKeyValuePairStringString
		return ret
	}
	return o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageRequestDto) GetPropsOk() ([]ItemKeyValuePairStringString, bool) {
	if o == nil || IsNil(o.Props) {
		return nil, false
	}
	return o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *StorageRequestDto) IsPropsSet() bool {
	if o != nil && !IsNil(o.Props) {
		return true
	}

	return false
}

// SetProps gets a reference to the given []ItemKeyValuePairStringString and assigns it to the Props field.
func (o *StorageRequestDto) SetProps(v []ItemKeyValuePairStringString) {
	o.Props = v
}

func (o StorageRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["module"] = o.Module.Get()
	if o.Props != nil {
		toSerialize["props"] = o.Props
	}
	return toSerialize, nil
}

func (o *StorageRequestDto) UnmarshalJSON(data []byte) (err error) {
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

	varStorageRequestDto := _StorageRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varStorageRequestDto)

	if err != nil {
		return err
	}

	*o = StorageRequestDto(varStorageRequestDto)

	return err
}

type NullableStorageRequestDto struct {
	value *StorageRequestDto
	isSet bool
}

func (v NullableStorageRequestDto) Get() *StorageRequestDto {
	return v.value
}

func (v *NullableStorageRequestDto) Set(val *StorageRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageRequestDto(val *StorageRequestDto) *NullableStorageRequestDto {
	return &NullableStorageRequestDto{value: val, isSet: true}
}

func (v NullableStorageRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

