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

// checks if the CompanyWhiteLabelSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyWhiteLabelSettingsDto{}

// CompanyWhiteLabelSettingsDto The vendor details the About page and the notification letters print, shared by the whole installation.
type CompanyWhiteLabelSettingsDto struct {
	// The vendor name the About page shows and the letters sign off with. Until details are saved it holds  whatever the installation ships as its built-in vendor, and it is empty on an installation that ships none.
	CompanyName NullableString `json:"companyName"`
	// The address the vendor name links to, as an absolute URL with its scheme. Empty under the same conditions  as `companyName`.
	Site NullableString `json:"site"`
	// The mailbox the About page offers for reaching the vendor. It is not the portal's own support address, and  it is empty under the same conditions as `companyName`.
	Email NullableString `json:"email"`
	// The postal address of the vendor as one free-form line, in the shape it was saved in - no structure is  imposed on it.
	Address NullableString `json:"address"`
	// The telephone number of the vendor in the shape it was saved in, with no dialling format enforced.
	Phone NullableString `json:"phone"`
	// Whether these details are those of the licensor of the product itself rather than of a reseller. Saving  through `POST api/2.0/settings/rebranding/company` always clears it, so only details that came with the  installation can report `true`.
	IsLicensor bool `json:"isLicensor"`
	// Whether the About page is hidden from the interface. A plan that does not include branding cannot switch it  on: the value is stored as `false` in that case, so it can come back different from what was saved.
	HideAbout bool `json:"hideAbout"`
	// Whether every field above still matches the installation's built-in vendor details. It turns `false` as  soon as one of them is saved differently and `true` again after  `DELETE api/2.0/settings/rebranding/company`.
	IsDefault bool `json:"isDefault"`
}

type _CompanyWhiteLabelSettingsDto CompanyWhiteLabelSettingsDto

// NewCompanyWhiteLabelSettingsDto instantiates a new CompanyWhiteLabelSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyWhiteLabelSettingsDto(companyName NullableString, site NullableString, email NullableString, address NullableString, phone NullableString, isLicensor bool, hideAbout bool, isDefault bool) *CompanyWhiteLabelSettingsDto {
	this := CompanyWhiteLabelSettingsDto{}
	this.CompanyName = companyName
	this.Site = site
	this.Email = email
	this.Address = address
	this.Phone = phone
	this.IsLicensor = isLicensor
	this.HideAbout = hideAbout
	this.IsDefault = isDefault
	return &this
}

// NewCompanyWhiteLabelSettingsDtoWithDefaults instantiates a new CompanyWhiteLabelSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWhiteLabelSettingsDtoWithDefaults() *CompanyWhiteLabelSettingsDto {
	this := CompanyWhiteLabelSettingsDto{}
	return &this
}

// GetCompanyName returns the CompanyName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompanyWhiteLabelSettingsDto) GetCompanyName() string {
	if o == nil || o.CompanyName.Get() == nil {
		var ret string
		return ret
	}

	return *o.CompanyName.Get()
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettingsDto) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyName.Get(), o.CompanyName.IsSet()
}

// SetCompanyName sets field value
func (o *CompanyWhiteLabelSettingsDto) SetCompanyName(v string) {
	o.CompanyName.Set(&v)
}

// GetSite returns the Site field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompanyWhiteLabelSettingsDto) GetSite() string {
	if o == nil || o.Site.Get() == nil {
		var ret string
		return ret
	}

	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettingsDto) GetSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// SetSite sets field value
func (o *CompanyWhiteLabelSettingsDto) SetSite(v string) {
	o.Site.Set(&v)
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompanyWhiteLabelSettingsDto) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettingsDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *CompanyWhiteLabelSettingsDto) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompanyWhiteLabelSettingsDto) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettingsDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *CompanyWhiteLabelSettingsDto) SetAddress(v string) {
	o.Address.Set(&v)
}

// GetPhone returns the Phone field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompanyWhiteLabelSettingsDto) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}

	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyWhiteLabelSettingsDto) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// SetPhone sets field value
func (o *CompanyWhiteLabelSettingsDto) SetPhone(v string) {
	o.Phone.Set(&v)
}

// GetIsLicensor returns the IsLicensor field value
func (o *CompanyWhiteLabelSettingsDto) GetIsLicensor() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLicensor
}

// GetIsLicensorOk returns a tuple with the IsLicensor field value
// and a boolean to check if the value has been set.
func (o *CompanyWhiteLabelSettingsDto) GetIsLicensorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLicensor, true
}

// SetIsLicensor sets field value
func (o *CompanyWhiteLabelSettingsDto) SetIsLicensor(v bool) {
	o.IsLicensor = v
}

// GetHideAbout returns the HideAbout field value
func (o *CompanyWhiteLabelSettingsDto) GetHideAbout() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HideAbout
}

// GetHideAboutOk returns a tuple with the HideAbout field value
// and a boolean to check if the value has been set.
func (o *CompanyWhiteLabelSettingsDto) GetHideAboutOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideAbout, true
}

// SetHideAbout sets field value
func (o *CompanyWhiteLabelSettingsDto) SetHideAbout(v bool) {
	o.HideAbout = v
}

// GetIsDefault returns the IsDefault field value
func (o *CompanyWhiteLabelSettingsDto) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *CompanyWhiteLabelSettingsDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *CompanyWhiteLabelSettingsDto) SetIsDefault(v bool) {
	o.IsDefault = v
}

func (o CompanyWhiteLabelSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyWhiteLabelSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyName"] = o.CompanyName.Get()
	toSerialize["site"] = o.Site.Get()
	toSerialize["email"] = o.Email.Get()
	toSerialize["address"] = o.Address.Get()
	toSerialize["phone"] = o.Phone.Get()
	toSerialize["isLicensor"] = o.IsLicensor
	toSerialize["hideAbout"] = o.HideAbout
	toSerialize["isDefault"] = o.IsDefault
	return toSerialize, nil
}

func (o *CompanyWhiteLabelSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"companyName",
		"site",
		"email",
		"address",
		"phone",
		"isLicensor",
		"hideAbout",
		"isDefault",
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

	varCompanyWhiteLabelSettingsDto := _CompanyWhiteLabelSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCompanyWhiteLabelSettingsDto)

	if err != nil {
		return err
	}

	*o = CompanyWhiteLabelSettingsDto(varCompanyWhiteLabelSettingsDto)

	return err
}

type NullableCompanyWhiteLabelSettingsDto struct {
	value *CompanyWhiteLabelSettingsDto
	isSet bool
}

func (v NullableCompanyWhiteLabelSettingsDto) Get() *CompanyWhiteLabelSettingsDto {
	return v.value
}

func (v *NullableCompanyWhiteLabelSettingsDto) Set(val *CompanyWhiteLabelSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyWhiteLabelSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyWhiteLabelSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyWhiteLabelSettingsDto(val *CompanyWhiteLabelSettingsDto) *NullableCompanyWhiteLabelSettingsDto {
	return &NullableCompanyWhiteLabelSettingsDto{value: val, isSet: true}
}

func (v NullableCompanyWhiteLabelSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyWhiteLabelSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

