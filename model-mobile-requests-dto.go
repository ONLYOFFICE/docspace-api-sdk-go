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

// checks if the MobileRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileRequestsDto{}

// MobileRequestsDto The parameters required for the mobile phone verification.
type MobileRequestsDto struct {
	// The user's mobile phone number.
	MobilePhone NullableString `json:"mobilePhone,omitempty"`
}

// NewMobileRequestsDto instantiates a new MobileRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileRequestsDto() *MobileRequestsDto {
	this := MobileRequestsDto{}
	return &this
}

// NewMobileRequestsDtoWithDefaults instantiates a new MobileRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileRequestsDtoWithDefaults() *MobileRequestsDto {
	this := MobileRequestsDto{}
	return &this
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MobileRequestsDto) GetMobilePhone() string {
	if o == nil || IsNil(o.MobilePhone.Get()) {
		var ret string
		return ret
	}
	return *o.MobilePhone.Get()
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MobileRequestsDto) GetMobilePhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobilePhone.Get(), o.MobilePhone.IsSet()
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *MobileRequestsDto) IsMobilePhoneSet() bool {
	if o != nil && o.MobilePhone.IsSet() {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given NullableString and assigns it to the MobilePhone field.
func (o *MobileRequestsDto) SetMobilePhone(v string) {
	o.MobilePhone.Set(&v)
}
// SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil
func (o *MobileRequestsDto) SetMobilePhoneNil() {
	o.MobilePhone.Set(nil)
}

// UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
func (o *MobileRequestsDto) UnsetMobilePhone() {
	o.MobilePhone.Unset()
}

func (o MobileRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MobilePhone.IsSet() {
		toSerialize["mobilePhone"] = o.MobilePhone.Get()
	}
	return toSerialize, nil
}

type NullableMobileRequestsDto struct {
	value *MobileRequestsDto
	isSet bool
}

func (v NullableMobileRequestsDto) Get() *MobileRequestsDto {
	return v.value
}

func (v *NullableMobileRequestsDto) Set(val *MobileRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileRequestsDto(val *MobileRequestsDto) *NullableMobileRequestsDto {
	return &NullableMobileRequestsDto{value: val, isSet: true}
}

func (v NullableMobileRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

