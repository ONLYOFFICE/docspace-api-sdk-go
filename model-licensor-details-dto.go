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

// checks if the LicensorDetailsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicensorDetailsDto{}

// LicensorDetailsDto The vendor details of the installation in the shape the licensor listing and the reset of the company details  return, with the licensor flag spelled `IsLicensor`.
type LicensorDetailsDto struct {
	// The vendor name the About page shows and the letters sign off with. Until details are saved it holds  whatever the installation ships as its built-in vendor, and it is empty on an installation that ships none.
	CompanyName NullableString `json:"companyName,omitempty"`
	// The address the vendor name links to, as an absolute URL with its scheme. Empty under the same conditions  as `companyName`.
	Site NullableString `json:"site,omitempty"`
	// The mailbox the About page offers for reaching the vendor. It is not the portal's own support address, and  it is empty under the same conditions as `companyName`.
	Email NullableString `json:"email,omitempty"`
	// The postal address of the vendor as one free-form line, in the shape it was saved in - no structure is  imposed on it.
	Address NullableString `json:"address,omitempty"`
	// The telephone number of the vendor in the shape it was saved in, with no dialling format enforced.
	Phone NullableString `json:"phone,omitempty"`
	// Whether these details are those of the licensor of the product itself rather than of a reseller. Saving  through `POST api/2.0/settings/rebranding/company` always clears it, so only details that came with the  installation can report `true`. The name starts with a capital letter, unlike the other fields.
	IsLicensor *bool `json:"IsLicensor,omitempty"`
	// Whether the About page is hidden from the interface. A plan that does not include branding cannot switch it  on: the value is stored as `false` in that case, so it can come back different from what was saved.
	HideAbout *bool `json:"hideAbout,omitempty"`
	// When these details were last stored. Details that were never stored report the moment they were read; the  built-in ONLYOFFICE entry and the answer of the reset operation report `0001-01-01T00:00:00`.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewLicensorDetailsDto instantiates a new LicensorDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensorDetailsDto() *LicensorDetailsDto {
	this := LicensorDetailsDto{}
	return &this
}

// NewLicensorDetailsDtoWithDefaults instantiates a new LicensorDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensorDetailsDtoWithDefaults() *LicensorDetailsDto {
	this := LicensorDetailsDto{}
	return &this
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensorDetailsDto) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyName.Get()
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensorDetailsDto) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyName.Get(), o.CompanyName.IsSet()
}

// HasCompanyName returns a boolean if a field has been set.
func (o *LicensorDetailsDto) IsCompanyNameSet() bool {
	if o != nil && o.CompanyName.IsSet() {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given NullableString and assigns it to the CompanyName field.
func (o *LicensorDetailsDto) SetCompanyName(v string) {
	o.CompanyName.Set(&v)
}
// SetCompanyNameNil sets the value for CompanyName to be an explicit nil
func (o *LicensorDetailsDto) SetCompanyNameNil() {
	o.CompanyName.Set(nil)
}

// UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
func (o *LicensorDetailsDto) UnsetCompanyName() {
	o.CompanyName.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensorDetailsDto) GetSite() string {
	if o == nil || IsNil(o.Site.Get()) {
		var ret string
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensorDetailsDto) GetSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *LicensorDetailsDto) IsSiteSet() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableString and assigns it to the Site field.
func (o *LicensorDetailsDto) SetSite(v string) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *LicensorDetailsDto) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *LicensorDetailsDto) UnsetSite() {
	o.Site.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensorDetailsDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensorDetailsDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *LicensorDetailsDto) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *LicensorDetailsDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *LicensorDetailsDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *LicensorDetailsDto) UnsetEmail() {
	o.Email.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensorDetailsDto) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensorDetailsDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *LicensorDetailsDto) IsAddressSet() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *LicensorDetailsDto) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *LicensorDetailsDto) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *LicensorDetailsDto) UnsetAddress() {
	o.Address.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensorDetailsDto) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensorDetailsDto) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *LicensorDetailsDto) IsPhoneSet() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *LicensorDetailsDto) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *LicensorDetailsDto) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *LicensorDetailsDto) UnsetPhone() {
	o.Phone.Unset()
}

// GetIsLicensor returns the IsLicensor field value if set, zero value otherwise.
func (o *LicensorDetailsDto) GetIsLicensor() bool {
	if o == nil || IsNil(o.IsLicensor) {
		var ret bool
		return ret
	}
	return *o.IsLicensor
}

// GetIsLicensorOk returns a tuple with the IsLicensor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensorDetailsDto) GetIsLicensorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLicensor) {
		return nil, false
	}
	return o.IsLicensor, true
}

// HasIsLicensor returns a boolean if a field has been set.
func (o *LicensorDetailsDto) IsIsLicensorSet() bool {
	if o != nil && !IsNil(o.IsLicensor) {
		return true
	}

	return false
}

// SetIsLicensor gets a reference to the given bool and assigns it to the IsLicensor field.
func (o *LicensorDetailsDto) SetIsLicensor(v bool) {
	o.IsLicensor = &v
}

// GetHideAbout returns the HideAbout field value if set, zero value otherwise.
func (o *LicensorDetailsDto) GetHideAbout() bool {
	if o == nil || IsNil(o.HideAbout) {
		var ret bool
		return ret
	}
	return *o.HideAbout
}

// GetHideAboutOk returns a tuple with the HideAbout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensorDetailsDto) GetHideAboutOk() (*bool, bool) {
	if o == nil || IsNil(o.HideAbout) {
		return nil, false
	}
	return o.HideAbout, true
}

// HasHideAbout returns a boolean if a field has been set.
func (o *LicensorDetailsDto) IsHideAboutSet() bool {
	if o != nil && !IsNil(o.HideAbout) {
		return true
	}

	return false
}

// SetHideAbout gets a reference to the given bool and assigns it to the HideAbout field.
func (o *LicensorDetailsDto) SetHideAbout(v bool) {
	o.HideAbout = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *LicensorDetailsDto) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensorDetailsDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *LicensorDetailsDto) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *LicensorDetailsDto) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o LicensorDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicensorDetailsDto) ToMap() (map[string]interface{}, error) {
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

type NullableLicensorDetailsDto struct {
	value *LicensorDetailsDto
	isSet bool
}

func (v NullableLicensorDetailsDto) Get() *LicensorDetailsDto {
	return v.value
}

func (v *NullableLicensorDetailsDto) Set(val *LicensorDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensorDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensorDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensorDetailsDto(val *LicensorDetailsDto) *NullableLicensorDetailsDto {
	return &NullableLicensorDetailsDto{value: val, isSet: true}
}

func (v NullableLicensorDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensorDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

