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

// checks if the TfaRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TfaRequestDto{}

// TfaRequestDto The portal two-factor policy: which method is in force, who must pass it, and from where it is waived.
type TfaRequestDto struct {
	// The second factor the portal demands. The two methods are mutually exclusive, so switching one on switches  the other off, and any value outside the defined set is read as switching TFA off rather than refused.
	Type *TfaType `json:"type,omitempty"`
	// The account the request concerns, by portal user ID. Naming the portal owner is refused unless it is the  caller's own account. Where an operation detaches an authenticator application, the empty GUID and the  caller's own ID both mean the caller.
	Id *string `json:"id,omitempty"`
	// The list of IP addresses that bypass TFA verification. Each entry is a single address, an inclusive  from-to range or a CIDR block. This is the whole list that is to hold afterwards, so send the addresses  already trusted along with a new one; an entry that cannot be parsed fails the call with 400, and accounts  named as mandatory still have to pass the challenge even from a trusted address.
	TrustedIps []string `json:"trustedIps,omitempty"`
	// The accounts that must pass the challenge whatever their address, by portal user ID. This is the whole list  that is to hold afterwards - leaving it out clears it rather than keeping it - and naming the portal owner is  refused unless the caller is the owner.
	MandatoryUsers []string `json:"mandatoryUsers,omitempty"`
	// The groups whose members must pass the challenge whatever their address, by group ID. This is the whole list  that is to hold afterwards - leaving it out clears it rather than keeping it.
	MandatoryGroups []string `json:"mandatoryGroups,omitempty"`
}

// NewTfaRequestDto instantiates a new TfaRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaRequestDto() *TfaRequestDto {
	this := TfaRequestDto{}
	return &this
}

// NewTfaRequestDtoWithDefaults instantiates a new TfaRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaRequestDtoWithDefaults() *TfaRequestDto {
	this := TfaRequestDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TfaRequestDto) GetType() TfaType {
	if o == nil || IsNil(o.Type) {
		var ret TfaType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaRequestDto) GetTypeOk() (*TfaType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TfaRequestDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TfaType and assigns it to the Type field.
func (o *TfaRequestDto) SetType(v TfaType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TfaRequestDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaRequestDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TfaRequestDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TfaRequestDto) SetId(v string) {
	o.Id = &v
}

// GetTrustedIps returns the TrustedIps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaRequestDto) GetTrustedIps() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TrustedIps
}

// GetTrustedIpsOk returns a tuple with the TrustedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaRequestDto) GetTrustedIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.TrustedIps) {
		return nil, false
	}
	return o.TrustedIps, true
}

// HasTrustedIps returns a boolean if a field has been set.
func (o *TfaRequestDto) IsTrustedIpsSet() bool {
	if o != nil && !IsNil(o.TrustedIps) {
		return true
	}

	return false
}

// SetTrustedIps gets a reference to the given []string and assigns it to the TrustedIps field.
func (o *TfaRequestDto) SetTrustedIps(v []string) {
	o.TrustedIps = v
}

// GetMandatoryUsers returns the MandatoryUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaRequestDto) GetMandatoryUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MandatoryUsers
}

// GetMandatoryUsersOk returns a tuple with the MandatoryUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaRequestDto) GetMandatoryUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.MandatoryUsers) {
		return nil, false
	}
	return o.MandatoryUsers, true
}

// HasMandatoryUsers returns a boolean if a field has been set.
func (o *TfaRequestDto) IsMandatoryUsersSet() bool {
	if o != nil && !IsNil(o.MandatoryUsers) {
		return true
	}

	return false
}

// SetMandatoryUsers gets a reference to the given []string and assigns it to the MandatoryUsers field.
func (o *TfaRequestDto) SetMandatoryUsers(v []string) {
	o.MandatoryUsers = v
}

// GetMandatoryGroups returns the MandatoryGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaRequestDto) GetMandatoryGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MandatoryGroups
}

// GetMandatoryGroupsOk returns a tuple with the MandatoryGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaRequestDto) GetMandatoryGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.MandatoryGroups) {
		return nil, false
	}
	return o.MandatoryGroups, true
}

// HasMandatoryGroups returns a boolean if a field has been set.
func (o *TfaRequestDto) IsMandatoryGroupsSet() bool {
	if o != nil && !IsNil(o.MandatoryGroups) {
		return true
	}

	return false
}

// SetMandatoryGroups gets a reference to the given []string and assigns it to the MandatoryGroups field.
func (o *TfaRequestDto) SetMandatoryGroups(v []string) {
	o.MandatoryGroups = v
}

func (o TfaRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TfaRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
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

type NullableTfaRequestDto struct {
	value *TfaRequestDto
	isSet bool
}

func (v NullableTfaRequestDto) Get() *TfaRequestDto {
	return v.value
}

func (v *NullableTfaRequestDto) Set(val *TfaRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaRequestDto(val *TfaRequestDto) *NullableTfaRequestDto {
	return &NullableTfaRequestDto{value: val, isSet: true}
}

func (v NullableTfaRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

