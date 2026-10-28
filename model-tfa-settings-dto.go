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

// checks if the TfaSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TfaSettingsDto{}

// TfaSettingsDto One two-factor authentication method the portal offers, with the portal-wide state of that method.
type TfaSettingsDto struct {
	// Which method this entry describes: `sms` for a code sent by text message, `app` for a code from an  authenticator application. It is the value `PUT api/2.0/settings/tfaapp` takes as its `type`, and no other  value ever appears here.
	Id NullableString `json:"id"`
	// The label for the method in the portal language, meant for a button or a radio option. It is not stable  enough to branch on - match `id` for that.
	Title NullableString `json:"title"`
	// Whether this method is the portal's current policy. At most one entry can have it set, and none has it  while the portal challenges nobody. It says nothing about the caller's own account, which may be exempt  through `trustedIps` or forced through `mandatoryUsers`.
	Enabled bool `json:"enabled"`
	// Whether the method could be switched on at all. For `sms` it is `false` until the installation has a  working SMS provider, so a method can be offered here and still be impossible to enable; for `app` it is  always `true`.
	Available bool `json:"available"`
	// The addresses that skip the challenge, each either a single address, a `from-to` pair or a CIDR range. It  is empty when no address is exempt, which means every account is challenged.
	TrustedIps []string `json:"trustedIps,omitempty"`
	// The accounts that are challenged even from a trusted address, by user ID. Empty means the exemption in  `trustedIps` holds for everyone.
	MandatoryUsers []string `json:"mandatoryUsers,omitempty"`
	// The groups whose members are challenged even from a trusted address, by group ID, with the same reading of  an empty list as `mandatoryUsers`.
	MandatoryGroups []string `json:"mandatoryGroups,omitempty"`
}

type _TfaSettingsDto TfaSettingsDto

// NewTfaSettingsDto instantiates a new TfaSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaSettingsDto(id NullableString, title NullableString, enabled bool, available bool) *TfaSettingsDto {
	this := TfaSettingsDto{}
	this.Id = id
	this.Title = title
	this.Enabled = enabled
	this.Available = available
	return &this
}

// NewTfaSettingsDtoWithDefaults instantiates a new TfaSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaSettingsDtoWithDefaults() *TfaSettingsDto {
	this := TfaSettingsDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TfaSettingsDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaSettingsDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *TfaSettingsDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TfaSettingsDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaSettingsDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *TfaSettingsDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetEnabled returns the Enabled field value
func (o *TfaSettingsDto) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TfaSettingsDto) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TfaSettingsDto) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAvailable returns the Available field value
func (o *TfaSettingsDto) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *TfaSettingsDto) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *TfaSettingsDto) SetAvailable(v bool) {
	o.Available = v
}

// GetTrustedIps returns the TrustedIps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaSettingsDto) GetTrustedIps() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TrustedIps
}

// GetTrustedIpsOk returns a tuple with the TrustedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaSettingsDto) GetTrustedIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.TrustedIps) {
		return nil, false
	}
	return o.TrustedIps, true
}

// HasTrustedIps returns a boolean if a field has been set.
func (o *TfaSettingsDto) IsTrustedIpsSet() bool {
	if o != nil && !IsNil(o.TrustedIps) {
		return true
	}

	return false
}

// SetTrustedIps gets a reference to the given []string and assigns it to the TrustedIps field.
func (o *TfaSettingsDto) SetTrustedIps(v []string) {
	o.TrustedIps = v
}

// GetMandatoryUsers returns the MandatoryUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaSettingsDto) GetMandatoryUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MandatoryUsers
}

// GetMandatoryUsersOk returns a tuple with the MandatoryUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaSettingsDto) GetMandatoryUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.MandatoryUsers) {
		return nil, false
	}
	return o.MandatoryUsers, true
}

// HasMandatoryUsers returns a boolean if a field has been set.
func (o *TfaSettingsDto) IsMandatoryUsersSet() bool {
	if o != nil && !IsNil(o.MandatoryUsers) {
		return true
	}

	return false
}

// SetMandatoryUsers gets a reference to the given []string and assigns it to the MandatoryUsers field.
func (o *TfaSettingsDto) SetMandatoryUsers(v []string) {
	o.MandatoryUsers = v
}

// GetMandatoryGroups returns the MandatoryGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaSettingsDto) GetMandatoryGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MandatoryGroups
}

// GetMandatoryGroupsOk returns a tuple with the MandatoryGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaSettingsDto) GetMandatoryGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.MandatoryGroups) {
		return nil, false
	}
	return o.MandatoryGroups, true
}

// HasMandatoryGroups returns a boolean if a field has been set.
func (o *TfaSettingsDto) IsMandatoryGroupsSet() bool {
	if o != nil && !IsNil(o.MandatoryGroups) {
		return true
	}

	return false
}

// SetMandatoryGroups gets a reference to the given []string and assigns it to the MandatoryGroups field.
func (o *TfaSettingsDto) SetMandatoryGroups(v []string) {
	o.MandatoryGroups = v
}

func (o TfaSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TfaSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["title"] = o.Title.Get()
	toSerialize["enabled"] = o.Enabled
	toSerialize["available"] = o.Available
	if o.TrustedIps != nil {
		toSerialize["trustedIps"] = o.TrustedIps
	}
	if o.MandatoryUsers != nil {
		toSerialize["mandatoryUsers"] = o.MandatoryUsers
	}
	if o.MandatoryGroups != nil {
		toSerialize["mandatoryGroups"] = o.MandatoryGroups
	}
	return toSerialize, nil
}

func (o *TfaSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"enabled",
		"available",
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

	varTfaSettingsDto := _TfaSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTfaSettingsDto)

	if err != nil {
		return err
	}

	*o = TfaSettingsDto(varTfaSettingsDto)

	return err
}

type NullableTfaSettingsDto struct {
	value *TfaSettingsDto
	isSet bool
}

func (v NullableTfaSettingsDto) Get() *TfaSettingsDto {
	return v.value
}

func (v *NullableTfaSettingsDto) Set(val *TfaSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaSettingsDto(val *TfaSettingsDto) *NullableTfaSettingsDto {
	return &NullableTfaSettingsDto{value: val, isSet: true}
}

func (v NullableTfaSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

