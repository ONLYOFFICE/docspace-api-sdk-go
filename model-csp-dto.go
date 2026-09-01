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

// checks if the CspDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CspDto{}

// CspDto The CSP (Content Security Policy) parameters.
type CspDto struct {
	// The list of CSP domains.
	Domains []string `json:"domains"`
	// The CSP header.
	Header NullableString `json:"header"`
}

type _CspDto CspDto

// NewCspDto instantiates a new CspDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCspDto(domains []string, header NullableString) *CspDto {
	this := CspDto{}
	this.Domains = domains
	this.Header = header
	return &this
}

// NewCspDtoWithDefaults instantiates a new CspDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCspDtoWithDefaults() *CspDto {
	this := CspDto{}
	return &this
}

// GetDomains returns the Domains field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *CspDto) GetDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CspDto) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *CspDto) SetDomains(v []string) {
	o.Domains = v
}

// GetHeader returns the Header field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CspDto) GetHeader() string {
	if o == nil || o.Header.Get() == nil {
		var ret string
		return ret
	}

	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CspDto) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// SetHeader sets field value
func (o *CspDto) SetHeader(v string) {
	o.Header.Set(&v)
}

func (o CspDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CspDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	toSerialize["header"] = o.Header.Get()
	return toSerialize, nil
}

func (o *CspDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domains",
		"header",
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

	varCspDto := _CspDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCspDto)

	if err != nil {
		return err
	}

	*o = CspDto(varCspDto)

	return err
}

type NullableCspDto struct {
	value *CspDto
	isSet bool
}

func (v NullableCspDto) Get() *CspDto {
	return v.value
}

func (v *NullableCspDto) Set(val *CspDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCspDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCspDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCspDto(val *CspDto) *NullableCspDto {
	return &NullableCspDto{value: val, isSet: true}
}

func (v NullableCspDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCspDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

