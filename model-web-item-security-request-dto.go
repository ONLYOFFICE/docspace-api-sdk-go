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

// checks if the WebItemSecurityRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebItemSecurityRequestDto{}

// WebItemSecurityRequestDto The access rule stored for one portal module: whether it may be opened, and by whom.
type WebItemSecurityRequestDto struct {
	// The module the rule applies to, given as a GUID. A value that is not a GUID fails the request as invalid.
	Id NullableString `json:"id"`
	// Whether the module may be opened. It decides the outcome only while `subjects` names somebody: an empty  `subjects` array is stored as access for everyone whatever this flag says.
	Enabled *bool `json:"enabled,omitempty"`
	// The users and groups the rule is stored for, given by their IDs. This is the whole allow-list that is to hold  afterwards and not a list of additions - what was stored before is dropped. Leaving it out applies `enabled`  to everyone and skips the audit trail entry, while sending it empty stores access for everyone.
	Subjects []string `json:"subjects,omitempty"`
}

type _WebItemSecurityRequestDto WebItemSecurityRequestDto

// NewWebItemSecurityRequestDto instantiates a new WebItemSecurityRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebItemSecurityRequestDto(id NullableString) *WebItemSecurityRequestDto {
	this := WebItemSecurityRequestDto{}
	this.Id = id
	return &this
}

// NewWebItemSecurityRequestDtoWithDefaults instantiates a new WebItemSecurityRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebItemSecurityRequestDtoWithDefaults() *WebItemSecurityRequestDto {
	this := WebItemSecurityRequestDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebItemSecurityRequestDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebItemSecurityRequestDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *WebItemSecurityRequestDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WebItemSecurityRequestDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebItemSecurityRequestDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WebItemSecurityRequestDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WebItemSecurityRequestDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebItemSecurityRequestDto) GetSubjects() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebItemSecurityRequestDto) GetSubjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subjects) {
		return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *WebItemSecurityRequestDto) IsSubjectsSet() bool {
	if o != nil && !IsNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []string and assigns it to the Subjects field.
func (o *WebItemSecurityRequestDto) SetSubjects(v []string) {
	o.Subjects = v
}

func (o WebItemSecurityRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebItemSecurityRequestDto) ToMap() (map[string]interface{}, error) {
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

func (o *WebItemSecurityRequestDto) UnmarshalJSON(data []byte) (err error) {
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

	varWebItemSecurityRequestDto := _WebItemSecurityRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWebItemSecurityRequestDto)

	if err != nil {
		return err
	}

	*o = WebItemSecurityRequestDto(varWebItemSecurityRequestDto)

	return err
}

type NullableWebItemSecurityRequestDto struct {
	value *WebItemSecurityRequestDto
	isSet bool
}

func (v NullableWebItemSecurityRequestDto) Get() *WebItemSecurityRequestDto {
	return v.value
}

func (v *NullableWebItemSecurityRequestDto) Set(val *WebItemSecurityRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebItemSecurityRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebItemSecurityRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebItemSecurityRequestDto(val *WebItemSecurityRequestDto) *NullableWebItemSecurityRequestDto {
	return &NullableWebItemSecurityRequestDto{value: val, isSet: true}
}

func (v NullableWebItemSecurityRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebItemSecurityRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

