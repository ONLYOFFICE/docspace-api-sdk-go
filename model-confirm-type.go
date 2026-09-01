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
	"fmt"
)

// ConfirmType [EmpInvite - Emp invite, LinkInvite - Link invite, PortalSuspend - Portal suspend, PortalContinue - Portal continue, PortalRemove - Portal remove, DnsChange - Dns change, PortalOwnerChange - Portal owner change, Activation - Activation, EmailChange - Email change, EmailActivation - Email activation, PasswordChange - Password change, ProfileRemove - Profile remove, PhoneActivation - Phone activation, PhoneAuth - Phone auth, Auth - Auth, TfaActivation - Tfa activation, TfaAuth - Tfa auth, Wizard - Wizard, GuestShareLink - Guest share link]
type ConfirmType string

// List of ConfirmType
const (
	CONFIRMTYPE_EMP_INVITE ConfirmType = "EmpInvite"
	CONFIRMTYPE_LINK_INVITE ConfirmType = "LinkInvite"
	CONFIRMTYPE_PORTAL_SUSPEND ConfirmType = "PortalSuspend"
	CONFIRMTYPE_PORTAL_CONTINUE ConfirmType = "PortalContinue"
	CONFIRMTYPE_PORTAL_REMOVE ConfirmType = "PortalRemove"
	CONFIRMTYPE_DNS_CHANGE ConfirmType = "DnsChange"
	CONFIRMTYPE_PORTAL_OWNER_CHANGE ConfirmType = "PortalOwnerChange"
	CONFIRMTYPE_ACTIVATION ConfirmType = "Activation"
	CONFIRMTYPE_EMAIL_CHANGE ConfirmType = "EmailChange"
	CONFIRMTYPE_EMAIL_ACTIVATION ConfirmType = "EmailActivation"
	CONFIRMTYPE_PASSWORD_CHANGE ConfirmType = "PasswordChange"
	CONFIRMTYPE_PROFILE_REMOVE ConfirmType = "ProfileRemove"
	CONFIRMTYPE_PHONE_ACTIVATION ConfirmType = "PhoneActivation"
	CONFIRMTYPE_PHONE_AUTH ConfirmType = "PhoneAuth"
	CONFIRMTYPE_AUTH ConfirmType = "Auth"
	CONFIRMTYPE_TFA_ACTIVATION ConfirmType = "TfaActivation"
	CONFIRMTYPE_TFA_AUTH ConfirmType = "TfaAuth"
	CONFIRMTYPE_WIZARD ConfirmType = "Wizard"
	CONFIRMTYPE_GUEST_SHARE_LINK ConfirmType = "GuestShareLink"
)

// All allowed values of ConfirmType enum
var AllowedConfirmTypeEnumValues = []ConfirmType{
	"EmpInvite",
	"LinkInvite",
	"PortalSuspend",
	"PortalContinue",
	"PortalRemove",
	"DnsChange",
	"PortalOwnerChange",
	"Activation",
	"EmailChange",
	"EmailActivation",
	"PasswordChange",
	"ProfileRemove",
	"PhoneActivation",
	"PhoneAuth",
	"Auth",
	"TfaActivation",
	"TfaAuth",
	"Wizard",
	"GuestShareLink",
}

func (v *ConfirmType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConfirmType(value)
	for _, existing := range AllowedConfirmTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConfirmType", value)
}

// NewConfirmTypeFromValue returns a pointer to a valid ConfirmType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConfirmTypeFromValue(v string) (*ConfirmType, error) {
	ev := ConfirmType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConfirmType: valid values are %v", v, AllowedConfirmTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConfirmType) IsValid() bool {
	for _, existing := range AllowedConfirmTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfirmType value
func (v ConfirmType) Ptr() *ConfirmType {
	return &v
}

type NullableConfirmType struct {
	value *ConfirmType
	isSet bool
}

func (v NullableConfirmType) Get() *ConfirmType {
	return v.value
}

func (v *NullableConfirmType) Set(val *ConfirmType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmType(val *ConfirmType) *NullableConfirmType {
	return &NullableConfirmType{value: val, isSet: true}
}

func (v NullableConfirmType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

