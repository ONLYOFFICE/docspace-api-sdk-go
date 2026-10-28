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

// checks if the WizardRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WizardRequestDto{}

// WizardRequestDto What the initial setup wizard needs to finish a new portal: the owner credentials and the portal locale.
type WizardRequestDto struct {
	// The address the portal owner account is created with, which is also the address every administrative letter  goes to afterwards. It has to be a well-formed email address; a malformed one leaves the wizard unfinished.
	Email NullableString `json:"email"`
	// The owner password, already hashed in the client rather than sent in the clear. Hash it with the `salt`,  iteration count and hash size that `GET api/2.0/settings?withpassword=true` publishes, so the portal can  recognise it later; an empty value leaves the wizard unfinished.
	PasswordHash NullableString `json:"passwordHash"`
	// The portal interface language, as a culture name such as `en-US`. It has to be one of the cultures enabled  for the installation, and an unknown one leaves the shipped default in place instead of failing the wizard.
	Lng NullableString `json:"lng,omitempty"`
	// The time zone every portal date is rendered in, as an IANA identifier such as `Europe/Riga`. A value that  matches nothing falls back to UTC rather than failing the wizard.
	TimeZone NullableString `json:"timeZone,omitempty"`
	// The identifier of the Amazon Machine Image the portal was launched from, for an installation started from an  AWS image. It is recorded for the installation record only and changes nothing about the portal; leave it out  anywhere else.
	AmiId NullableString `json:"amiId,omitempty"`
	// Whether the owner agrees to receive product news at the address in `email`. It is a mailing consent and has  no bearing on the portal notifications, which are subscribed separately.
	SubscribeFromSite *bool `json:"subscribeFromSite,omitempty"`
}

type _WizardRequestDto WizardRequestDto

// NewWizardRequestDto instantiates a new WizardRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWizardRequestDto(email NullableString, passwordHash NullableString) *WizardRequestDto {
	this := WizardRequestDto{}
	this.Email = email
	this.PasswordHash = passwordHash
	return &this
}

// NewWizardRequestDtoWithDefaults instantiates a new WizardRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWizardRequestDtoWithDefaults() *WizardRequestDto {
	this := WizardRequestDto{}
	return &this
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WizardRequestDto) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WizardRequestDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *WizardRequestDto) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetPasswordHash returns the PasswordHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WizardRequestDto) GetPasswordHash() string {
	if o == nil || o.PasswordHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WizardRequestDto) GetPasswordHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// SetPasswordHash sets field value
func (o *WizardRequestDto) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}

// GetLng returns the Lng field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WizardRequestDto) GetLng() string {
	if o == nil || IsNil(o.Lng.Get()) {
		var ret string
		return ret
	}
	return *o.Lng.Get()
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WizardRequestDto) GetLngOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lng.Get(), o.Lng.IsSet()
}

// HasLng returns a boolean if a field has been set.
func (o *WizardRequestDto) IsLngSet() bool {
	if o != nil && o.Lng.IsSet() {
		return true
	}

	return false
}

// SetLng gets a reference to the given NullableString and assigns it to the Lng field.
func (o *WizardRequestDto) SetLng(v string) {
	o.Lng.Set(&v)
}
// SetLngNil sets the value for Lng to be an explicit nil
func (o *WizardRequestDto) SetLngNil() {
	o.Lng.Set(nil)
}

// UnsetLng ensures that no value is present for Lng, not even an explicit nil
func (o *WizardRequestDto) UnsetLng() {
	o.Lng.Unset()
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WizardRequestDto) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone.Get()) {
		var ret string
		return ret
	}
	return *o.TimeZone.Get()
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WizardRequestDto) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeZone.Get(), o.TimeZone.IsSet()
}

// HasTimeZone returns a boolean if a field has been set.
func (o *WizardRequestDto) IsTimeZoneSet() bool {
	if o != nil && o.TimeZone.IsSet() {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given NullableString and assigns it to the TimeZone field.
func (o *WizardRequestDto) SetTimeZone(v string) {
	o.TimeZone.Set(&v)
}
// SetTimeZoneNil sets the value for TimeZone to be an explicit nil
func (o *WizardRequestDto) SetTimeZoneNil() {
	o.TimeZone.Set(nil)
}

// UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
func (o *WizardRequestDto) UnsetTimeZone() {
	o.TimeZone.Unset()
}

// GetAmiId returns the AmiId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WizardRequestDto) GetAmiId() string {
	if o == nil || IsNil(o.AmiId.Get()) {
		var ret string
		return ret
	}
	return *o.AmiId.Get()
}

// GetAmiIdOk returns a tuple with the AmiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WizardRequestDto) GetAmiIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmiId.Get(), o.AmiId.IsSet()
}

// HasAmiId returns a boolean if a field has been set.
func (o *WizardRequestDto) IsAmiIdSet() bool {
	if o != nil && o.AmiId.IsSet() {
		return true
	}

	return false
}

// SetAmiId gets a reference to the given NullableString and assigns it to the AmiId field.
func (o *WizardRequestDto) SetAmiId(v string) {
	o.AmiId.Set(&v)
}
// SetAmiIdNil sets the value for AmiId to be an explicit nil
func (o *WizardRequestDto) SetAmiIdNil() {
	o.AmiId.Set(nil)
}

// UnsetAmiId ensures that no value is present for AmiId, not even an explicit nil
func (o *WizardRequestDto) UnsetAmiId() {
	o.AmiId.Unset()
}

// GetSubscribeFromSite returns the SubscribeFromSite field value if set, zero value otherwise.
func (o *WizardRequestDto) GetSubscribeFromSite() bool {
	if o == nil || IsNil(o.SubscribeFromSite) {
		var ret bool
		return ret
	}
	return *o.SubscribeFromSite
}

// GetSubscribeFromSiteOk returns a tuple with the SubscribeFromSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardRequestDto) GetSubscribeFromSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.SubscribeFromSite) {
		return nil, false
	}
	return o.SubscribeFromSite, true
}

// HasSubscribeFromSite returns a boolean if a field has been set.
func (o *WizardRequestDto) IsSubscribeFromSiteSet() bool {
	if o != nil && !IsNil(o.SubscribeFromSite) {
		return true
	}

	return false
}

// SetSubscribeFromSite gets a reference to the given bool and assigns it to the SubscribeFromSite field.
func (o *WizardRequestDto) SetSubscribeFromSite(v bool) {
	o.SubscribeFromSite = &v
}

func (o WizardRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WizardRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email.Get()
	toSerialize["passwordHash"] = o.PasswordHash.Get()
	if o.Lng.IsSet() {
		toSerialize["lng"] = o.Lng.Get()
	}
	if o.TimeZone.IsSet() {
		toSerialize["timeZone"] = o.TimeZone.Get()
	}
	if o.AmiId.IsSet() {
		toSerialize["amiId"] = o.AmiId.Get()
	}
	if !IsNil(o.SubscribeFromSite) {
		toSerialize["subscribeFromSite"] = o.SubscribeFromSite
	}
	return toSerialize, nil
}

func (o *WizardRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"passwordHash",
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

	varWizardRequestDto := _WizardRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWizardRequestDto)

	if err != nil {
		return err
	}

	*o = WizardRequestDto(varWizardRequestDto)

	return err
}

type NullableWizardRequestDto struct {
	value *WizardRequestDto
	isSet bool
}

func (v NullableWizardRequestDto) Get() *WizardRequestDto {
	return v.value
}

func (v *NullableWizardRequestDto) Set(val *WizardRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWizardRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWizardRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWizardRequestDto(val *WizardRequestDto) *NullableWizardRequestDto {
	return &NullableWizardRequestDto{value: val, isSet: true}
}

func (v NullableWizardRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWizardRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

