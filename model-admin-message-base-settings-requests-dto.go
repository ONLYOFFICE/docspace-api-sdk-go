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

// checks if the AdminMessageBaseSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminMessageBaseSettingsRequestsDto{}

// AdminMessageBaseSettingsRequestsDto Who is invited to join the portal, and in which language the invitation is written.
type AdminMessageBaseSettingsRequestsDto struct {
	// The address the join link is sent to. It has to be a well-formed ASCII address rather than an  internationalized one, must not already belong to a member of the portal, and, where the portal trusts named  domains only, has to end with one of them; any of these faults is refused with 400.
	Email NullableString `json:"email"`
	// The language the letter is written in, as a culture name such as `en-US`. A culture the installation does not  have falls back to the portal language rather than failing the call.
	Culture NullableString `json:"culture,omitempty"`
}

type _AdminMessageBaseSettingsRequestsDto AdminMessageBaseSettingsRequestsDto

// NewAdminMessageBaseSettingsRequestsDto instantiates a new AdminMessageBaseSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminMessageBaseSettingsRequestsDto(email NullableString) *AdminMessageBaseSettingsRequestsDto {
	this := AdminMessageBaseSettingsRequestsDto{}
	this.Email = email
	return &this
}

// NewAdminMessageBaseSettingsRequestsDtoWithDefaults instantiates a new AdminMessageBaseSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminMessageBaseSettingsRequestsDtoWithDefaults() *AdminMessageBaseSettingsRequestsDto {
	this := AdminMessageBaseSettingsRequestsDto{}
	return &this
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AdminMessageBaseSettingsRequestsDto) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminMessageBaseSettingsRequestsDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *AdminMessageBaseSettingsRequestsDto) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetCulture returns the Culture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminMessageBaseSettingsRequestsDto) GetCulture() string {
	if o == nil || IsNil(o.Culture.Get()) {
		var ret string
		return ret
	}
	return *o.Culture.Get()
}

// GetCultureOk returns a tuple with the Culture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminMessageBaseSettingsRequestsDto) GetCultureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Culture.Get(), o.Culture.IsSet()
}

// HasCulture returns a boolean if a field has been set.
func (o *AdminMessageBaseSettingsRequestsDto) IsCultureSet() bool {
	if o != nil && o.Culture.IsSet() {
		return true
	}

	return false
}

// SetCulture gets a reference to the given NullableString and assigns it to the Culture field.
func (o *AdminMessageBaseSettingsRequestsDto) SetCulture(v string) {
	o.Culture.Set(&v)
}
// SetCultureNil sets the value for Culture to be an explicit nil
func (o *AdminMessageBaseSettingsRequestsDto) SetCultureNil() {
	o.Culture.Set(nil)
}

// UnsetCulture ensures that no value is present for Culture, not even an explicit nil
func (o *AdminMessageBaseSettingsRequestsDto) UnsetCulture() {
	o.Culture.Unset()
}

func (o AdminMessageBaseSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminMessageBaseSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email.Get()
	if o.Culture.IsSet() {
		toSerialize["culture"] = o.Culture.Get()
	}
	return toSerialize, nil
}

func (o *AdminMessageBaseSettingsRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varAdminMessageBaseSettingsRequestsDto := _AdminMessageBaseSettingsRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAdminMessageBaseSettingsRequestsDto)

	if err != nil {
		return err
	}

	*o = AdminMessageBaseSettingsRequestsDto(varAdminMessageBaseSettingsRequestsDto)

	return err
}

type NullableAdminMessageBaseSettingsRequestsDto struct {
	value *AdminMessageBaseSettingsRequestsDto
	isSet bool
}

func (v NullableAdminMessageBaseSettingsRequestsDto) Get() *AdminMessageBaseSettingsRequestsDto {
	return v.value
}

func (v *NullableAdminMessageBaseSettingsRequestsDto) Set(val *AdminMessageBaseSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminMessageBaseSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminMessageBaseSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminMessageBaseSettingsRequestsDto(val *AdminMessageBaseSettingsRequestsDto) *NullableAdminMessageBaseSettingsRequestsDto {
	return &NullableAdminMessageBaseSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableAdminMessageBaseSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminMessageBaseSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

