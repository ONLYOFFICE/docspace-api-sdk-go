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

// checks if the CultureSpecificExternalResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CultureSpecificExternalResource{}

// CultureSpecificExternalResource The external resource parameters.
type CultureSpecificExternalResource struct {
	// The external resource domain.
	Domain NullableString `json:"domain,omitempty"`
	// The external resource entries.
	Entries map[string]string `json:"entries,omitempty"`
}

// NewCultureSpecificExternalResource instantiates a new CultureSpecificExternalResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCultureSpecificExternalResource() *CultureSpecificExternalResource {
	this := CultureSpecificExternalResource{}
	return &this
}

// NewCultureSpecificExternalResourceWithDefaults instantiates a new CultureSpecificExternalResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCultureSpecificExternalResourceWithDefaults() *CultureSpecificExternalResource {
	this := CultureSpecificExternalResource{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CultureSpecificExternalResource) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CultureSpecificExternalResource) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *CultureSpecificExternalResource) IsDomainSet() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *CultureSpecificExternalResource) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *CultureSpecificExternalResource) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *CultureSpecificExternalResource) UnsetDomain() {
	o.Domain.Unset()
}

// GetEntries returns the Entries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CultureSpecificExternalResource) GetEntries() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CultureSpecificExternalResource) GetEntriesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return &o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *CultureSpecificExternalResource) IsEntriesSet() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given map[string]string and assigns it to the Entries field.
func (o *CultureSpecificExternalResource) SetEntries(v map[string]string) {
	o.Entries = v
}

func (o CultureSpecificExternalResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CultureSpecificExternalResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return toSerialize, nil
}

type NullableCultureSpecificExternalResource struct {
	value *CultureSpecificExternalResource
	isSet bool
}

func (v NullableCultureSpecificExternalResource) Get() *CultureSpecificExternalResource {
	return v.value
}

func (v *NullableCultureSpecificExternalResource) Set(val *CultureSpecificExternalResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCultureSpecificExternalResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCultureSpecificExternalResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCultureSpecificExternalResource(val *CultureSpecificExternalResource) *NullableCultureSpecificExternalResource {
	return &NullableCultureSpecificExternalResource{value: val, isSet: true}
}

func (v NullableCultureSpecificExternalResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCultureSpecificExternalResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

