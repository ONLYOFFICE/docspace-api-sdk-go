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

// checks if the TfaRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TfaRequestsDto{}

// TfaRequestsDto The request parameters for configuring the Two-Factor Authentication (TFA) settings.
type TfaRequestsDto struct {
	Type *TfaRequestsDtoType `json:"type,omitempty"`
	// The ID of the user for whom the TFA settings are being configured.
	Id *string `json:"id,omitempty"`
	// The list of IP addresses that bypass TFA verification.
	TrustedIps []string `json:"trustedIps,omitempty"`
	// The list of user IDs for whom TFA is mandatory.
	MandatoryUsers []string `json:"mandatoryUsers,omitempty"`
	// The list group IDs whose members must use TFA.
	MandatoryGroups []string `json:"mandatoryGroups,omitempty"`
}

// NewTfaRequestsDto instantiates a new TfaRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaRequestsDto() *TfaRequestsDto {
	this := TfaRequestsDto{}
	return &this
}

// NewTfaRequestsDtoWithDefaults instantiates a new TfaRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaRequestsDtoWithDefaults() *TfaRequestsDto {
	this := TfaRequestsDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TfaRequestsDto) GetType() TfaRequestsDtoType {
	if o == nil || IsNil(o.Type) {
		var ret TfaRequestsDtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaRequestsDto) GetTypeOk() (*TfaRequestsDtoType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TfaRequestsDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TfaRequestsDtoType and assigns it to the Type field.
func (o *TfaRequestsDto) SetType(v TfaRequestsDtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TfaRequestsDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaRequestsDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TfaRequestsDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TfaRequestsDto) SetId(v string) {
	o.Id = &v
}

// GetTrustedIps returns the TrustedIps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaRequestsDto) GetTrustedIps() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TrustedIps
}

// GetTrustedIpsOk returns a tuple with the TrustedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaRequestsDto) GetTrustedIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.TrustedIps) {
		return nil, false
	}
	return o.TrustedIps, true
}

// HasTrustedIps returns a boolean if a field has been set.
func (o *TfaRequestsDto) IsTrustedIpsSet() bool {
	if o != nil && !IsNil(o.TrustedIps) {
		return true
	}

	return false
}

// SetTrustedIps gets a reference to the given []string and assigns it to the TrustedIps field.
func (o *TfaRequestsDto) SetTrustedIps(v []string) {
	o.TrustedIps = v
}

// GetMandatoryUsers returns the MandatoryUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaRequestsDto) GetMandatoryUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MandatoryUsers
}

// GetMandatoryUsersOk returns a tuple with the MandatoryUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaRequestsDto) GetMandatoryUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.MandatoryUsers) {
		return nil, false
	}
	return o.MandatoryUsers, true
}

// HasMandatoryUsers returns a boolean if a field has been set.
func (o *TfaRequestsDto) IsMandatoryUsersSet() bool {
	if o != nil && !IsNil(o.MandatoryUsers) {
		return true
	}

	return false
}

// SetMandatoryUsers gets a reference to the given []string and assigns it to the MandatoryUsers field.
func (o *TfaRequestsDto) SetMandatoryUsers(v []string) {
	o.MandatoryUsers = v
}

// GetMandatoryGroups returns the MandatoryGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaRequestsDto) GetMandatoryGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MandatoryGroups
}

// GetMandatoryGroupsOk returns a tuple with the MandatoryGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaRequestsDto) GetMandatoryGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.MandatoryGroups) {
		return nil, false
	}
	return o.MandatoryGroups, true
}

// HasMandatoryGroups returns a boolean if a field has been set.
func (o *TfaRequestsDto) IsMandatoryGroupsSet() bool {
	if o != nil && !IsNil(o.MandatoryGroups) {
		return true
	}

	return false
}

// SetMandatoryGroups gets a reference to the given []string and assigns it to the MandatoryGroups field.
func (o *TfaRequestsDto) SetMandatoryGroups(v []string) {
	o.MandatoryGroups = v
}

func (o TfaRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TfaRequestsDto) ToMap() (map[string]interface{}, error) {
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

type NullableTfaRequestsDto struct {
	value *TfaRequestsDto
	isSet bool
}

func (v NullableTfaRequestsDto) Get() *TfaRequestsDto {
	return v.value
}

func (v *NullableTfaRequestsDto) Set(val *TfaRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaRequestsDto(val *TfaRequestsDto) *NullableTfaRequestsDto {
	return &NullableTfaRequestsDto{value: val, isSet: true}
}

func (v NullableTfaRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

