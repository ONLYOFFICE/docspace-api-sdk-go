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

// checks if the SsoSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSettingsRequestsDto{}

// SsoSettingsRequestsDto The whole SAML Single Sign-On configuration of the portal, carried as a serialised JSON object.
type SsoSettingsRequestsDto struct {
	// The configuration object serialised to a JSON string, not a nested object. It is the complete configuration  rather than a patch - fields left out are stored empty - so start from `GET api/2.0/settings/ssov2` or  `GET api/2.0/settings/ssov2/default` and send back a changed copy. The identity provider entity ID and  sign-in URL are required, the sign-in and sign-out URLs have to be absolute `http` or `https` addresses, and  the attribute mapping has to name the first name, last name and email fields; the values each SAML field  accepts are listed by `GET api/2.0/settings/ssov2/constants`. An empty string, or a string that carries no  configuration object, is refused with 400.
	SerializeSettings NullableString `json:"serializeSettings"`
}

type _SsoSettingsRequestsDto SsoSettingsRequestsDto

// NewSsoSettingsRequestsDto instantiates a new SsoSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSettingsRequestsDto(serializeSettings NullableString) *SsoSettingsRequestsDto {
	this := SsoSettingsRequestsDto{}
	this.SerializeSettings = serializeSettings
	return &this
}

// NewSsoSettingsRequestsDtoWithDefaults instantiates a new SsoSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSettingsRequestsDtoWithDefaults() *SsoSettingsRequestsDto {
	this := SsoSettingsRequestsDto{}
	return &this
}

// GetSerializeSettings returns the SerializeSettings field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SsoSettingsRequestsDto) GetSerializeSettings() string {
	if o == nil || o.SerializeSettings.Get() == nil {
		var ret string
		return ret
	}

	return *o.SerializeSettings.Get()
}

// GetSerializeSettingsOk returns a tuple with the SerializeSettings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsRequestsDto) GetSerializeSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SerializeSettings.Get(), o.SerializeSettings.IsSet()
}

// SetSerializeSettings sets field value
func (o *SsoSettingsRequestsDto) SetSerializeSettings(v string) {
	o.SerializeSettings.Set(&v)
}

func (o SsoSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serializeSettings"] = o.SerializeSettings.Get()
	return toSerialize, nil
}

func (o *SsoSettingsRequestsDto) UnmarshalJSON(data []byte) (err error) {
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

	varSsoSettingsRequestsDto := _SsoSettingsRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSsoSettingsRequestsDto)

	if err != nil {
		return err
	}

	*o = SsoSettingsRequestsDto(varSsoSettingsRequestsDto)

	return err
}

type NullableSsoSettingsRequestsDto struct {
	value *SsoSettingsRequestsDto
	isSet bool
}

func (v NullableSsoSettingsRequestsDto) Get() *SsoSettingsRequestsDto {
	return v.value
}

func (v *NullableSsoSettingsRequestsDto) Set(val *SsoSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSettingsRequestsDto(val *SsoSettingsRequestsDto) *NullableSsoSettingsRequestsDto {
	return &NullableSsoSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableSsoSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

