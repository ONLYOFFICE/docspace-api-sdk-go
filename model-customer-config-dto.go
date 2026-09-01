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

// checks if the CustomerConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerConfigDto{}

// CustomerConfigDto The customer config parameters.
type CustomerConfigDto struct {
	// The address of the customer configuration.
	Address NullableString `json:"address,omitempty"`
	// The logo of the customer configuration.
	Logo NullableString `json:"logo,omitempty"`
	// The dark logo of the customer configuration.
	LogoDark NullableString `json:"logoDark,omitempty"`
	// The mail address of the customer configuration.
	Mail NullableString `json:"mail,omitempty"`
	// The name of the customer configuration.
	Name NullableString `json:"name,omitempty"`
	// The site web address of the customer configuration.
	Www NullableString `json:"www,omitempty"`
}

// NewCustomerConfigDto instantiates a new CustomerConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerConfigDto() *CustomerConfigDto {
	this := CustomerConfigDto{}
	return &this
}

// NewCustomerConfigDtoWithDefaults instantiates a new CustomerConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerConfigDtoWithDefaults() *CustomerConfigDto {
	this := CustomerConfigDto{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerConfigDto) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerConfigDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *CustomerConfigDto) IsAddressSet() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *CustomerConfigDto) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *CustomerConfigDto) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *CustomerConfigDto) UnsetAddress() {
	o.Address.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerConfigDto) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerConfigDto) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *CustomerConfigDto) IsLogoSet() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *CustomerConfigDto) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *CustomerConfigDto) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *CustomerConfigDto) UnsetLogo() {
	o.Logo.Unset()
}

// GetLogoDark returns the LogoDark field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerConfigDto) GetLogoDark() string {
	if o == nil || IsNil(o.LogoDark.Get()) {
		var ret string
		return ret
	}
	return *o.LogoDark.Get()
}

// GetLogoDarkOk returns a tuple with the LogoDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerConfigDto) GetLogoDarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoDark.Get(), o.LogoDark.IsSet()
}

// HasLogoDark returns a boolean if a field has been set.
func (o *CustomerConfigDto) IsLogoDarkSet() bool {
	if o != nil && o.LogoDark.IsSet() {
		return true
	}

	return false
}

// SetLogoDark gets a reference to the given NullableString and assigns it to the LogoDark field.
func (o *CustomerConfigDto) SetLogoDark(v string) {
	o.LogoDark.Set(&v)
}
// SetLogoDarkNil sets the value for LogoDark to be an explicit nil
func (o *CustomerConfigDto) SetLogoDarkNil() {
	o.LogoDark.Set(nil)
}

// UnsetLogoDark ensures that no value is present for LogoDark, not even an explicit nil
func (o *CustomerConfigDto) UnsetLogoDark() {
	o.LogoDark.Unset()
}

// GetMail returns the Mail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerConfigDto) GetMail() string {
	if o == nil || IsNil(o.Mail.Get()) {
		var ret string
		return ret
	}
	return *o.Mail.Get()
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerConfigDto) GetMailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mail.Get(), o.Mail.IsSet()
}

// HasMail returns a boolean if a field has been set.
func (o *CustomerConfigDto) IsMailSet() bool {
	if o != nil && o.Mail.IsSet() {
		return true
	}

	return false
}

// SetMail gets a reference to the given NullableString and assigns it to the Mail field.
func (o *CustomerConfigDto) SetMail(v string) {
	o.Mail.Set(&v)
}
// SetMailNil sets the value for Mail to be an explicit nil
func (o *CustomerConfigDto) SetMailNil() {
	o.Mail.Set(nil)
}

// UnsetMail ensures that no value is present for Mail, not even an explicit nil
func (o *CustomerConfigDto) UnsetMail() {
	o.Mail.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerConfigDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerConfigDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CustomerConfigDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CustomerConfigDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CustomerConfigDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CustomerConfigDto) UnsetName() {
	o.Name.Unset()
}

// GetWww returns the Www field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerConfigDto) GetWww() string {
	if o == nil || IsNil(o.Www.Get()) {
		var ret string
		return ret
	}
	return *o.Www.Get()
}

// GetWwwOk returns a tuple with the Www field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerConfigDto) GetWwwOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Www.Get(), o.Www.IsSet()
}

// HasWww returns a boolean if a field has been set.
func (o *CustomerConfigDto) IsWwwSet() bool {
	if o != nil && o.Www.IsSet() {
		return true
	}

	return false
}

// SetWww gets a reference to the given NullableString and assigns it to the Www field.
func (o *CustomerConfigDto) SetWww(v string) {
	o.Www.Set(&v)
}
// SetWwwNil sets the value for Www to be an explicit nil
func (o *CustomerConfigDto) SetWwwNil() {
	o.Www.Set(nil)
}

// UnsetWww ensures that no value is present for Www, not even an explicit nil
func (o *CustomerConfigDto) UnsetWww() {
	o.Www.Unset()
}

func (o CustomerConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	if o.LogoDark.IsSet() {
		toSerialize["logoDark"] = o.LogoDark.Get()
	}
	if o.Mail.IsSet() {
		toSerialize["mail"] = o.Mail.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Www.IsSet() {
		toSerialize["www"] = o.Www.Get()
	}
	return toSerialize, nil
}

type NullableCustomerConfigDto struct {
	value *CustomerConfigDto
	isSet bool
}

func (v NullableCustomerConfigDto) Get() *CustomerConfigDto {
	return v.value
}

func (v *NullableCustomerConfigDto) Set(val *CustomerConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerConfigDto(val *CustomerConfigDto) *NullableCustomerConfigDto {
	return &NullableCustomerConfigDto{value: val, isSet: true}
}

func (v NullableCustomerConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

