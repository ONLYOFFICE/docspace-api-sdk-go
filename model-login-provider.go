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

// LoginProvider [0 - Facebook, 1 - Google, 2 - Dropbox, 3 - Docusign, 4 - Box, 5 - OneDrive, 6 - GosUslugi, 7 - LinkedIn, 8 - MailRu, 9 - VK, 10 - Wordpress, 11 - Yahoo, 12 - Yandex, 13 - Github]
type LoginProvider int32

// List of LoginProvider
const (
	LOGINPROVIDER_Facebook LoginProvider = 0
	LOGINPROVIDER_Google LoginProvider = 1
	LOGINPROVIDER_Dropbox LoginProvider = 2
	LOGINPROVIDER_Docusign LoginProvider = 3
	LOGINPROVIDER_Box LoginProvider = 4
	LOGINPROVIDER_OneDrive LoginProvider = 5
	LOGINPROVIDER_GosUslugi LoginProvider = 6
	LOGINPROVIDER_LinkedIn LoginProvider = 7
	LOGINPROVIDER_MailRu LoginProvider = 8
	LOGINPROVIDER_VK LoginProvider = 9
	LOGINPROVIDER_Wordpress LoginProvider = 10
	LOGINPROVIDER_Yahoo LoginProvider = 11
	LOGINPROVIDER_Yandex LoginProvider = 12
	LOGINPROVIDER_Github LoginProvider = 13
)

// All allowed values of LoginProvider enum
var AllowedLoginProviderEnumValues = []LoginProvider{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
}

func (v *LoginProvider) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoginProvider(value)
	for _, existing := range AllowedLoginProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoginProvider", value)
}

// NewLoginProviderFromValue returns a pointer to a valid LoginProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoginProviderFromValue(v int32) (*LoginProvider, error) {
	ev := LoginProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LoginProvider: valid values are %v", v, AllowedLoginProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoginProvider) IsValid() bool {
	for _, existing := range AllowedLoginProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoginProvider value
func (v LoginProvider) Ptr() *LoginProvider {
	return &v
}

type NullableLoginProvider struct {
	value *LoginProvider
	isSet bool
}

func (v NullableLoginProvider) Get() *LoginProvider {
	return v.value
}

func (v *NullableLoginProvider) Set(val *LoginProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginProvider(val *LoginProvider) *NullableLoginProvider {
	return &NullableLoginProvider{value: val, isSet: true}
}

func (v NullableLoginProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

