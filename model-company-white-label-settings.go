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
	"time"
)

// checks if the CompanyWhiteLabelSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyWhiteLabelSettings{}

// CompanyWhiteLabelSettings The company white label settings.
type CompanyWhiteLabelSettings struct {
	// The company name.
	CompanyName NullableString `json:"companyName,omitempty"`
	// The company site.
	Site NullableString `json:"site,omitempty"`
	// The company email address.
	Email NullableString `json:"email,omitempty"`
	// The company address.
	Address NullableString `json:"address,omitempty"`
	// The company phone number.
	Phone NullableString `json:"phone,omitempty"`
	// Specifies if a company is a licensor or not.
	IsLicensor *bool `json:"IsLicensor,omitempty"`
	// Specifies if the About page is visible or not
	HideAbout *bool `json:"hideAbout,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewCompanyWhiteLabelSettings instantiates a new CompanyWhiteLabelSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyWhiteLabelSettings() *CompanyWhiteLabelSettings {
	this := CompanyWhiteLabelSettings{}
	return &this
}

// NewCompanyWhiteLabelSettingsWithDefaults instantiates a new CompanyWhiteLabelSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWhiteLabelSettingsWithDefaults() *CompanyWhiteLabelSettings {
	this := CompanyWhiteLabelSettings{}
	return &this
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyWhiteLabelSettings) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyName.Get()
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettings) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyName.Get(), o.CompanyName.IsSet()
}

// HasCompanyName returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettings) IsCompanyNameSet() bool {
	if o != nil && o.CompanyName.IsSet() {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given NullableString and assigns it to the CompanyName field.
func (o *CompanyWhiteLabelSettings) SetCompanyName(v string) {
	o.CompanyName.Set(&v)
}
// SetCompanyNameNil sets the value for CompanyName to be an explicit nil
func (o *CompanyWhiteLabelSettings) SetCompanyNameNil() {
	o.CompanyName.Set(nil)
}

// UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
func (o *CompanyWhiteLabelSettings) UnsetCompanyName() {
	o.CompanyName.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyWhiteLabelSettings) GetSite() string {
	if o == nil || IsNil(o.Site.Get()) {
		var ret string
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettings) GetSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettings) IsSiteSet() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableString and assigns it to the Site field.
func (o *CompanyWhiteLabelSettings) SetSite(v string) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *CompanyWhiteLabelSettings) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *CompanyWhiteLabelSettings) UnsetSite() {
	o.Site.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyWhiteLabelSettings) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettings) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettings) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *CompanyWhiteLabelSettings) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *CompanyWhiteLabelSettings) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *CompanyWhiteLabelSettings) UnsetEmail() {
	o.Email.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyWhiteLabelSettings) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettings) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettings) IsAddressSet() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *CompanyWhiteLabelSettings) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *CompanyWhiteLabelSettings) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *CompanyWhiteLabelSettings) UnsetAddress() {
	o.Address.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyWhiteLabelSettings) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettings) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettings) IsPhoneSet() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *CompanyWhiteLabelSettings) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *CompanyWhiteLabelSettings) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *CompanyWhiteLabelSettings) UnsetPhone() {
	o.Phone.Unset()
}

// GetIsLicensor returns the IsLicensor field value if set, zero value otherwise.
func (o *CompanyWhiteLabelSettings) GetIsLicensor() bool {
	if o == nil || IsNil(o.IsLicensor) {
		var ret bool
		return ret
	}
	return *o.IsLicensor
}

// GetIsLicensorOk returns a tuple with the IsLicensor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyWhiteLabelSettings) GetIsLicensorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLicensor) {
		return nil, false
	}
	return o.IsLicensor, true
}

// HasIsLicensor returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettings) IsIsLicensorSet() bool {
	if o != nil && !IsNil(o.IsLicensor) {
		return true
	}

	return false
}

// SetIsLicensor gets a reference to the given bool and assigns it to the IsLicensor field.
func (o *CompanyWhiteLabelSettings) SetIsLicensor(v bool) {
	o.IsLicensor = &v
}

// GetHideAbout returns the HideAbout field value if set, zero value otherwise.
func (o *CompanyWhiteLabelSettings) GetHideAbout() bool {
	if o == nil || IsNil(o.HideAbout) {
		var ret bool
		return ret
	}
	return *o.HideAbout
}

// GetHideAboutOk returns a tuple with the HideAbout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyWhiteLabelSettings) GetHideAboutOk() (*bool, bool) {
	if o == nil || IsNil(o.HideAbout) {
		return nil, false
	}
	return o.HideAbout, true
}

// HasHideAbout returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettings) IsHideAboutSet() bool {
	if o != nil && !IsNil(o.HideAbout) {
		return true
	}

	return false
}

// SetHideAbout gets a reference to the given bool and assigns it to the HideAbout field.
func (o *CompanyWhiteLabelSettings) SetHideAbout(v bool) {
	o.HideAbout = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *CompanyWhiteLabelSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyWhiteLabelSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *CompanyWhiteLabelSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o CompanyWhiteLabelSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyWhiteLabelSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyName.IsSet() {
		toSerialize["companyName"] = o.CompanyName.Get()
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if !IsNil(o.IsLicensor) {
		toSerialize["IsLicensor"] = o.IsLicensor
	}
	if !IsNil(o.HideAbout) {
		toSerialize["hideAbout"] = o.HideAbout
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableCompanyWhiteLabelSettings struct {
	value *CompanyWhiteLabelSettings
	isSet bool
}

func (v NullableCompanyWhiteLabelSettings) Get() *CompanyWhiteLabelSettings {
	return v.value
}

func (v *NullableCompanyWhiteLabelSettings) Set(val *CompanyWhiteLabelSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyWhiteLabelSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyWhiteLabelSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyWhiteLabelSettings(val *CompanyWhiteLabelSettings) *NullableCompanyWhiteLabelSettings {
	return &NullableCompanyWhiteLabelSettings{value: val, isSet: true}
}

func (v NullableCompanyWhiteLabelSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyWhiteLabelSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

