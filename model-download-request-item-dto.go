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

// checks if the DownloadRequestItemDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadRequestItemDto{}

// DownloadRequestItemDto One file of a bulk download, together with the format it is converted to.
type DownloadRequestItemDto struct {
	Key DownloadRequestItemDtoKey `json:"key"`
	// The format the file is converted to before it is packed, as a file extension without a leading dot.
	Value NullableString `json:"value"`
	// The password that opens the source file, for a file protected with one; a protected file cannot be converted  without it.
	Password NullableString `json:"password,omitempty"`
}

type _DownloadRequestItemDto DownloadRequestItemDto

// NewDownloadRequestItemDto instantiates a new DownloadRequestItemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadRequestItemDto(key DownloadRequestItemDtoKey, value NullableString) *DownloadRequestItemDto {
	this := DownloadRequestItemDto{}
	this.Key = key
	this.Value = value
	return &this
}

// NewDownloadRequestItemDtoWithDefaults instantiates a new DownloadRequestItemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadRequestItemDtoWithDefaults() *DownloadRequestItemDto {
	this := DownloadRequestItemDto{}
	return &this
}

// GetKey returns the Key field value
func (o *DownloadRequestItemDto) GetKey() DownloadRequestItemDtoKey {
	if o == nil {
		var ret DownloadRequestItemDtoKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DownloadRequestItemDto) GetKeyOk() (*DownloadRequestItemDtoKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DownloadRequestItemDto) SetKey(v DownloadRequestItemDtoKey) {
	o.Key = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DownloadRequestItemDto) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadRequestItemDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *DownloadRequestItemDto) SetValue(v string) {
	o.Value.Set(&v)
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadRequestItemDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadRequestItemDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *DownloadRequestItemDto) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *DownloadRequestItemDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *DownloadRequestItemDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *DownloadRequestItemDto) UnsetPassword() {
	o.Password.Unset()
}

func (o DownloadRequestItemDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadRequestItemDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value.Get()
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return toSerialize, nil
}

func (o *DownloadRequestItemDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varDownloadRequestItemDto := _DownloadRequestItemDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDownloadRequestItemDto)

	if err != nil {
		return err
	}

	*o = DownloadRequestItemDto(varDownloadRequestItemDto)

	return err
}

type NullableDownloadRequestItemDto struct {
	value *DownloadRequestItemDto
	isSet bool
}

func (v NullableDownloadRequestItemDto) Get() *DownloadRequestItemDto {
	return v.value
}

func (v *NullableDownloadRequestItemDto) Set(val *DownloadRequestItemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadRequestItemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadRequestItemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadRequestItemDto(val *DownloadRequestItemDto) *NullableDownloadRequestItemDto {
	return &NullableDownloadRequestItemDto{value: val, isSet: true}
}

func (v NullableDownloadRequestItemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadRequestItemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

