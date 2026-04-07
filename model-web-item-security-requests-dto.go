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

// checks if the WebItemSecurityRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebItemSecurityRequestsDto{}

// WebItemSecurityRequestsDto The request parameters for configuring security settings of a single web module.
type WebItemSecurityRequestsDto struct {
	// The module ID.
	Id NullableString `json:"id"`
	// Controls whether the security restrictions are enforced for this module.
	Enabled *bool `json:"enabled,omitempty"`
	// The collection of user and group identifiers granted access to the module.
	Subjects []string `json:"subjects,omitempty"`
}

type _WebItemSecurityRequestsDto WebItemSecurityRequestsDto

// NewWebItemSecurityRequestsDto instantiates a new WebItemSecurityRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebItemSecurityRequestsDto(id NullableString) *WebItemSecurityRequestsDto {
	this := WebItemSecurityRequestsDto{}
	this.Id = id
	return &this
}

// NewWebItemSecurityRequestsDtoWithDefaults instantiates a new WebItemSecurityRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebItemSecurityRequestsDtoWithDefaults() *WebItemSecurityRequestsDto {
	this := WebItemSecurityRequestsDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebItemSecurityRequestsDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebItemSecurityRequestsDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *WebItemSecurityRequestsDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WebItemSecurityRequestsDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebItemSecurityRequestsDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WebItemSecurityRequestsDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WebItemSecurityRequestsDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebItemSecurityRequestsDto) GetSubjects() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebItemSecurityRequestsDto) GetSubjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subjects) {
		return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *WebItemSecurityRequestsDto) IsSubjectsSet() bool {
	if o != nil && !IsNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []string and assigns it to the Subjects field.
func (o *WebItemSecurityRequestsDto) SetSubjects(v []string) {
	o.Subjects = v
}

func (o WebItemSecurityRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebItemSecurityRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Subjects != nil {
		toSerialize["subjects"] = o.Subjects
	}
	return toSerialize, nil
}

func (o *WebItemSecurityRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varWebItemSecurityRequestsDto := _WebItemSecurityRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWebItemSecurityRequestsDto)

	if err != nil {
		return err
	}

	*o = WebItemSecurityRequestsDto(varWebItemSecurityRequestsDto)

	return err
}

type NullableWebItemSecurityRequestsDto struct {
	value *WebItemSecurityRequestsDto
	isSet bool
}

func (v NullableWebItemSecurityRequestsDto) Get() *WebItemSecurityRequestsDto {
	return v.value
}

func (v *NullableWebItemSecurityRequestsDto) Set(val *WebItemSecurityRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebItemSecurityRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebItemSecurityRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebItemSecurityRequestsDto(val *WebItemSecurityRequestsDto) *NullableWebItemSecurityRequestsDto {
	return &NullableWebItemSecurityRequestsDto{value: val, isSet: true}
}

func (v NullableWebItemSecurityRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebItemSecurityRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

