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

// checks if the SsoSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSettingsRequestDto{}

// SsoSettingsRequestDto The whole SAML Single Sign-On configuration of the portal, carried as a serialised JSON object.
type SsoSettingsRequestDto struct {
	// The configuration object serialised to a JSON string, not a nested object. It is the complete configuration  rather than a patch - fields left out are stored empty - so start from `GET api/2.0/settings/ssov2` or  `GET api/2.0/settings/ssov2/default` and send back a changed copy. The identity provider entity ID and  sign-in URL are required, the sign-in and sign-out URLs have to be absolute `http` or `https` addresses, and  the attribute mapping has to name the first name, last name and email fields; the values each SAML field  accepts are listed by `GET api/2.0/settings/ssov2/constants`. An empty string, or a string that carries no  configuration object, is refused with 400.
	SerializeSettings NullableString `json:"serializeSettings"`
}

type _SsoSettingsRequestDto SsoSettingsRequestDto

// NewSsoSettingsRequestDto instantiates a new SsoSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSettingsRequestDto(serializeSettings NullableString) *SsoSettingsRequestDto {
	this := SsoSettingsRequestDto{}
	this.SerializeSettings = serializeSettings
	return &this
}

// NewSsoSettingsRequestDtoWithDefaults instantiates a new SsoSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSettingsRequestDtoWithDefaults() *SsoSettingsRequestDto {
	this := SsoSettingsRequestDto{}
	return &this
}

// GetSerializeSettings returns the SerializeSettings field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SsoSettingsRequestDto) GetSerializeSettings() string {
	if o == nil || o.SerializeSettings.Get() == nil {
		var ret string
		return ret
	}

	return *o.SerializeSettings.Get()
}

// GetSerializeSettingsOk returns a tuple with the SerializeSettings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsRequestDto) GetSerializeSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SerializeSettings.Get(), o.SerializeSettings.IsSet()
}

// SetSerializeSettings sets field value
func (o *SsoSettingsRequestDto) SetSerializeSettings(v string) {
	o.SerializeSettings.Set(&v)
}

func (o SsoSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serializeSettings"] = o.SerializeSettings.Get()
	return toSerialize, nil
}

func (o *SsoSettingsRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serializeSettings",
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

	varSsoSettingsRequestDto := _SsoSettingsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSsoSettingsRequestDto)

	if err != nil {
		return err
	}

	*o = SsoSettingsRequestDto(varSsoSettingsRequestDto)

	return err
}

type NullableSsoSettingsRequestDto struct {
	value *SsoSettingsRequestDto
	isSet bool
}

func (v NullableSsoSettingsRequestDto) Get() *SsoSettingsRequestDto {
	return v.value
}

func (v *NullableSsoSettingsRequestDto) Set(val *SsoSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSettingsRequestDto(val *SsoSettingsRequestDto) *NullableSsoSettingsRequestDto {
	return &NullableSsoSettingsRequestDto{value: val, isSet: true}
}

func (v NullableSsoSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

