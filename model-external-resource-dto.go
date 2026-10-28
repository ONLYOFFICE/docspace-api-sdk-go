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

// checks if the ExternalResourceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalResourceDto{}

// ExternalResourceDto The external resource parameters.
type ExternalResourceDto struct {
	// The external resource domain.
	Domain NullableString `json:"domain,omitempty"`
	// The external resource entries.
	Entries map[string]*string `json:"entries,omitempty"`
}

// NewExternalResourceDto instantiates a new ExternalResourceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalResourceDto() *ExternalResourceDto {
	this := ExternalResourceDto{}
	return &this
}

// NewExternalResourceDtoWithDefaults instantiates a new ExternalResourceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalResourceDtoWithDefaults() *ExternalResourceDto {
	this := ExternalResourceDto{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalResourceDto) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalResourceDto) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *ExternalResourceDto) IsDomainSet() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *ExternalResourceDto) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *ExternalResourceDto) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *ExternalResourceDto) UnsetDomain() {
	o.Domain.Unset()
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *ExternalResourceDto) GetEntries() map[string]*string {
	if o == nil || IsNil(o.Entries) {
		var ret map[string]*string
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalResourceDto) GetEntriesOk() (map[string]*string, bool) {
	if o == nil || IsNil(o.Entries) {
		return map[string]*string{}, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *ExternalResourceDto) IsEntriesSet() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given map[string]*string and assigns it to the Entries field.
func (o *ExternalResourceDto) SetEntries(v map[string]*string) {
	o.Entries = v
}

func (o ExternalResourceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalResourceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	return toSerialize, nil
}

type NullableExternalResourceDto struct {
	value *ExternalResourceDto
	isSet bool
}

func (v NullableExternalResourceDto) Get() *ExternalResourceDto {
	return v.value
}

func (v *NullableExternalResourceDto) Set(val *ExternalResourceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalResourceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalResourceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalResourceDto(val *ExternalResourceDto) *NullableExternalResourceDto {
	return &NullableExternalResourceDto{value: val, isSet: true}
}

func (v NullableExternalResourceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalResourceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

