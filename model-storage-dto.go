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

// checks if the StorageDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageDto{}

// StorageDto The storage information.
type StorageDto struct {
	// The storage ID.
	Id NullableString `json:"id"`
	// The storage title.
	Title NullableString `json:"title"`
	// The list of storage authentication keys.
	Properties []AuthKey `json:"properties,omitempty"`
	// Specifies if this is the current portal storage or not.
	Current bool `json:"current"`
	// Specifies if this storage can be set or not.
	IsSet bool `json:"isSet"`
}

type _StorageDto StorageDto

// NewStorageDto instantiates a new StorageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDto(id NullableString, title NullableString, current bool, isSet bool) *StorageDto {
	this := StorageDto{}
	this.Id = id
	this.Title = title
	this.Current = current
	this.IsSet = isSet
	return &this
}

// NewStorageDtoWithDefaults instantiates a new StorageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDtoWithDefaults() *StorageDto {
	this := StorageDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *StorageDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *StorageDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDto) GetProperties() []AuthKey {
	if o == nil {
		var ret []AuthKey
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDto) GetPropertiesOk() ([]AuthKey, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *StorageDto) IsPropertiesSet() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []AuthKey and assigns it to the Properties field.
func (o *StorageDto) SetProperties(v []AuthKey) {
	o.Properties = v
}

// GetCurrent returns the Current field value
func (o *StorageDto) GetCurrent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *StorageDto) GetCurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *StorageDto) SetCurrent(v bool) {
	o.Current = v
}

// GetIsSet returns the IsSet field value
func (o *StorageDto) GetIsSet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSet
}

// GetIsSetOk returns a tuple with the IsSet field value
// and a boolean to check if the value has been set.
func (o *StorageDto) GetIsSetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSet, true
}

// SetIsSet sets field value
func (o *StorageDto) SetIsSet(v bool) {
	o.IsSet = v
}

func (o StorageDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["title"] = o.Title.Get()
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	toSerialize["current"] = o.Current
	toSerialize["isSet"] = o.IsSet
	return toSerialize, nil
}

func (o *StorageDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"current",
		"isSet",
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

	varStorageDto := _StorageDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varStorageDto)

	if err != nil {
		return err
	}

	*o = StorageDto(varStorageDto)

	return err
}

type NullableStorageDto struct {
	value *StorageDto
	isSet bool
}

func (v NullableStorageDto) Get() *StorageDto {
	return v.value
}

func (v *NullableStorageDto) Set(val *StorageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDto(val *StorageDto) *NullableStorageDto {
	return &NullableStorageDto{value: val, isSet: true}
}

func (v NullableStorageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

