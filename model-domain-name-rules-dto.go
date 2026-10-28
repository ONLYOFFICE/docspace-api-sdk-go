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

// checks if the DomainNameRulesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainNameRulesDto{}

// DomainNameRulesDto The rules a portal name is checked against.
type DomainNameRulesDto struct {
	// The pattern the portal name has to match.
	Regex NullableString `json:"regex,omitempty"`
	// The shortest portal name accepted.
	MinLength *int32 `json:"minLength,omitempty"`
	// The longest portal name accepted.
	MaxLength *int32 `json:"maxLength,omitempty"`
}

// NewDomainNameRulesDto instantiates a new DomainNameRulesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainNameRulesDto() *DomainNameRulesDto {
	this := DomainNameRulesDto{}
	return &this
}

// NewDomainNameRulesDtoWithDefaults instantiates a new DomainNameRulesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainNameRulesDtoWithDefaults() *DomainNameRulesDto {
	this := DomainNameRulesDto{}
	return &this
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainNameRulesDto) GetRegex() string {
	if o == nil || IsNil(o.Regex.Get()) {
		var ret string
		return ret
	}
	return *o.Regex.Get()
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainNameRulesDto) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regex.Get(), o.Regex.IsSet()
}

// HasRegex returns a boolean if a field has been set.
func (o *DomainNameRulesDto) IsRegexSet() bool {
	if o != nil && o.Regex.IsSet() {
		return true
	}

	return false
}

// SetRegex gets a reference to the given NullableString and assigns it to the Regex field.
func (o *DomainNameRulesDto) SetRegex(v string) {
	o.Regex.Set(&v)
}
// SetRegexNil sets the value for Regex to be an explicit nil
func (o *DomainNameRulesDto) SetRegexNil() {
	o.Regex.Set(nil)
}

// UnsetRegex ensures that no value is present for Regex, not even an explicit nil
func (o *DomainNameRulesDto) UnsetRegex() {
	o.Regex.Unset()
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *DomainNameRulesDto) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainNameRulesDto) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *DomainNameRulesDto) IsMinLengthSet() bool {
	if o != nil && !IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *DomainNameRulesDto) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *DomainNameRulesDto) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength) {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainNameRulesDto) GetMaxLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLength) {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *DomainNameRulesDto) IsMaxLengthSet() bool {
	if o != nil && !IsNil(o.MaxLength) {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *DomainNameRulesDto) SetMaxLength(v int32) {
	o.MaxLength = &v
}

func (o DomainNameRulesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainNameRulesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Regex.IsSet() {
		toSerialize["regex"] = o.Regex.Get()
	}
	if !IsNil(o.MinLength) {
		toSerialize["minLength"] = o.MinLength
	}
	if !IsNil(o.MaxLength) {
		toSerialize["maxLength"] = o.MaxLength
	}
	return toSerialize, nil
}

type NullableDomainNameRulesDto struct {
	value *DomainNameRulesDto
	isSet bool
}

func (v NullableDomainNameRulesDto) Get() *DomainNameRulesDto {
	return v.value
}

func (v *NullableDomainNameRulesDto) Set(val *DomainNameRulesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainNameRulesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainNameRulesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainNameRulesDto(val *DomainNameRulesDto) *NullableDomainNameRulesDto {
	return &NullableDomainNameRulesDto{value: val, isSet: true}
}

func (v NullableDomainNameRulesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainNameRulesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

