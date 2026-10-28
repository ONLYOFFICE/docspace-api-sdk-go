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

// checks if the SsoFieldMappingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoFieldMappingDto{}

// SsoFieldMappingDto Which SAML attributes fill the profile fields of a user who signs in through SSO.
type SsoFieldMappingDto struct {
	// The first name.
	FirstName NullableString `json:"firstName,omitempty"`
	// The last name.
	LastName NullableString `json:"lastName,omitempty"`
	// The email address.
	Email NullableString `json:"email,omitempty"`
	// The title.
	Title NullableString `json:"title,omitempty"`
	// The location.
	Location NullableString `json:"location,omitempty"`
	// The phone number.
	Phone NullableString `json:"phone,omitempty"`
}

// NewSsoFieldMappingDto instantiates a new SsoFieldMappingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoFieldMappingDto() *SsoFieldMappingDto {
	this := SsoFieldMappingDto{}
	return &this
}

// NewSsoFieldMappingDtoWithDefaults instantiates a new SsoFieldMappingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoFieldMappingDtoWithDefaults() *SsoFieldMappingDto {
	this := SsoFieldMappingDto{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoFieldMappingDto) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoFieldMappingDto) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *SsoFieldMappingDto) IsFirstNameSet() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *SsoFieldMappingDto) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *SsoFieldMappingDto) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *SsoFieldMappingDto) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoFieldMappingDto) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoFieldMappingDto) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *SsoFieldMappingDto) IsLastNameSet() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *SsoFieldMappingDto) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *SsoFieldMappingDto) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *SsoFieldMappingDto) UnsetLastName() {
	o.LastName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoFieldMappingDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoFieldMappingDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *SsoFieldMappingDto) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *SsoFieldMappingDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *SsoFieldMappingDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *SsoFieldMappingDto) UnsetEmail() {
	o.Email.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoFieldMappingDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoFieldMappingDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *SsoFieldMappingDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *SsoFieldMappingDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *SsoFieldMappingDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *SsoFieldMappingDto) UnsetTitle() {
	o.Title.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoFieldMappingDto) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoFieldMappingDto) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *SsoFieldMappingDto) IsLocationSet() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *SsoFieldMappingDto) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *SsoFieldMappingDto) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *SsoFieldMappingDto) UnsetLocation() {
	o.Location.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoFieldMappingDto) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoFieldMappingDto) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *SsoFieldMappingDto) IsPhoneSet() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *SsoFieldMappingDto) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *SsoFieldMappingDto) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *SsoFieldMappingDto) UnsetPhone() {
	o.Phone.Unset()
}

func (o SsoFieldMappingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoFieldMappingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	return toSerialize, nil
}

type NullableSsoFieldMappingDto struct {
	value *SsoFieldMappingDto
	isSet bool
}

func (v NullableSsoFieldMappingDto) Get() *SsoFieldMappingDto {
	return v.value
}

func (v *NullableSsoFieldMappingDto) Set(val *SsoFieldMappingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoFieldMappingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoFieldMappingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoFieldMappingDto(val *SsoFieldMappingDto) *NullableSsoFieldMappingDto {
	return &NullableSsoFieldMappingDto{value: val, isSet: true}
}

func (v NullableSsoFieldMappingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoFieldMappingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

