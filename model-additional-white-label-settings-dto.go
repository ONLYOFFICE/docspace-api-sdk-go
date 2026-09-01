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

// checks if the AdditionalWhiteLabelSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalWhiteLabelSettingsDto{}

// AdditionalWhiteLabelSettingsDto The additional white label settings parameters.
type AdditionalWhiteLabelSettingsDto struct {
	// Specifies if the sample documents are displayed or hidden.
	StartDocsEnabled bool `json:"startDocsEnabled"`
	// Specifies if the Help Center link is available or not.
	HelpCenterEnabled bool `json:"helpCenterEnabled"`
	// Specifies if the Feedback & Support link is available or not.
	FeedbackAndSupportEnabled bool `json:"feedbackAndSupportEnabled"`
	// Specifies if the user forum is available or not.
	UserForumEnabled bool `json:"userForumEnabled"`
	// Specifies if the Video Guides link is available or not.
	VideoGuidesEnabled bool `json:"videoGuidesEnabled"`
	// Specifies if the License Agreements link is available or not.
	LicenseAgreementsEnabled bool `json:"licenseAgreementsEnabled"`
	// Specifies if the additional white label settings are default or not.
	IsDefault bool `json:"isDefault"`
}

type _AdditionalWhiteLabelSettingsDto AdditionalWhiteLabelSettingsDto

// NewAdditionalWhiteLabelSettingsDto instantiates a new AdditionalWhiteLabelSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalWhiteLabelSettingsDto(startDocsEnabled bool, helpCenterEnabled bool, feedbackAndSupportEnabled bool, userForumEnabled bool, videoGuidesEnabled bool, licenseAgreementsEnabled bool, isDefault bool) *AdditionalWhiteLabelSettingsDto {
	this := AdditionalWhiteLabelSettingsDto{}
	this.StartDocsEnabled = startDocsEnabled
	this.HelpCenterEnabled = helpCenterEnabled
	this.FeedbackAndSupportEnabled = feedbackAndSupportEnabled
	this.UserForumEnabled = userForumEnabled
	this.VideoGuidesEnabled = videoGuidesEnabled
	this.LicenseAgreementsEnabled = licenseAgreementsEnabled
	this.IsDefault = isDefault
	return &this
}

// NewAdditionalWhiteLabelSettingsDtoWithDefaults instantiates a new AdditionalWhiteLabelSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalWhiteLabelSettingsDtoWithDefaults() *AdditionalWhiteLabelSettingsDto {
	this := AdditionalWhiteLabelSettingsDto{}
	return &this
}

// GetStartDocsEnabled returns the StartDocsEnabled field value
func (o *AdditionalWhiteLabelSettingsDto) GetStartDocsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StartDocsEnabled
}

// GetStartDocsEnabledOk returns a tuple with the StartDocsEnabled field value
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettingsDto) GetStartDocsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDocsEnabled, true
}

// SetStartDocsEnabled sets field value
func (o *AdditionalWhiteLabelSettingsDto) SetStartDocsEnabled(v bool) {
	o.StartDocsEnabled = v
}

// GetHelpCenterEnabled returns the HelpCenterEnabled field value
func (o *AdditionalWhiteLabelSettingsDto) GetHelpCenterEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HelpCenterEnabled
}

// GetHelpCenterEnabledOk returns a tuple with the HelpCenterEnabled field value
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettingsDto) GetHelpCenterEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HelpCenterEnabled, true
}

// SetHelpCenterEnabled sets field value
func (o *AdditionalWhiteLabelSettingsDto) SetHelpCenterEnabled(v bool) {
	o.HelpCenterEnabled = v
}

// GetFeedbackAndSupportEnabled returns the FeedbackAndSupportEnabled field value
func (o *AdditionalWhiteLabelSettingsDto) GetFeedbackAndSupportEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FeedbackAndSupportEnabled
}

// GetFeedbackAndSupportEnabledOk returns a tuple with the FeedbackAndSupportEnabled field value
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettingsDto) GetFeedbackAndSupportEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeedbackAndSupportEnabled, true
}

// SetFeedbackAndSupportEnabled sets field value
func (o *AdditionalWhiteLabelSettingsDto) SetFeedbackAndSupportEnabled(v bool) {
	o.FeedbackAndSupportEnabled = v
}

// GetUserForumEnabled returns the UserForumEnabled field value
func (o *AdditionalWhiteLabelSettingsDto) GetUserForumEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UserForumEnabled
}

// GetUserForumEnabledOk returns a tuple with the UserForumEnabled field value
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettingsDto) GetUserForumEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserForumEnabled, true
}

// SetUserForumEnabled sets field value
func (o *AdditionalWhiteLabelSettingsDto) SetUserForumEnabled(v bool) {
	o.UserForumEnabled = v
}

// GetVideoGuidesEnabled returns the VideoGuidesEnabled field value
func (o *AdditionalWhiteLabelSettingsDto) GetVideoGuidesEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VideoGuidesEnabled
}

// GetVideoGuidesEnabledOk returns a tuple with the VideoGuidesEnabled field value
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettingsDto) GetVideoGuidesEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VideoGuidesEnabled, true
}

// SetVideoGuidesEnabled sets field value
func (o *AdditionalWhiteLabelSettingsDto) SetVideoGuidesEnabled(v bool) {
	o.VideoGuidesEnabled = v
}

// GetLicenseAgreementsEnabled returns the LicenseAgreementsEnabled field value
func (o *AdditionalWhiteLabelSettingsDto) GetLicenseAgreementsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LicenseAgreementsEnabled
}

// GetLicenseAgreementsEnabledOk returns a tuple with the LicenseAgreementsEnabled field value
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettingsDto) GetLicenseAgreementsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseAgreementsEnabled, true
}

// SetLicenseAgreementsEnabled sets field value
func (o *AdditionalWhiteLabelSettingsDto) SetLicenseAgreementsEnabled(v bool) {
	o.LicenseAgreementsEnabled = v
}

// GetIsDefault returns the IsDefault field value
func (o *AdditionalWhiteLabelSettingsDto) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettingsDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *AdditionalWhiteLabelSettingsDto) SetIsDefault(v bool) {
	o.IsDefault = v
}

func (o AdditionalWhiteLabelSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalWhiteLabelSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startDocsEnabled"] = o.StartDocsEnabled
	toSerialize["helpCenterEnabled"] = o.HelpCenterEnabled
	toSerialize["feedbackAndSupportEnabled"] = o.FeedbackAndSupportEnabled
	toSerialize["userForumEnabled"] = o.UserForumEnabled
	toSerialize["videoGuidesEnabled"] = o.VideoGuidesEnabled
	toSerialize["licenseAgreementsEnabled"] = o.LicenseAgreementsEnabled
	toSerialize["isDefault"] = o.IsDefault
	return toSerialize, nil
}

func (o *AdditionalWhiteLabelSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startDocsEnabled",
		"helpCenterEnabled",
		"feedbackAndSupportEnabled",
		"userForumEnabled",
		"videoGuidesEnabled",
		"licenseAgreementsEnabled",
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

	varAdditionalWhiteLabelSettingsDto := _AdditionalWhiteLabelSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAdditionalWhiteLabelSettingsDto)

	if err != nil {
		return err
	}

	*o = AdditionalWhiteLabelSettingsDto(varAdditionalWhiteLabelSettingsDto)

	return err
}

type NullableAdditionalWhiteLabelSettingsDto struct {
	value *AdditionalWhiteLabelSettingsDto
	isSet bool
}

func (v NullableAdditionalWhiteLabelSettingsDto) Get() *AdditionalWhiteLabelSettingsDto {
	return v.value
}

func (v *NullableAdditionalWhiteLabelSettingsDto) Set(val *AdditionalWhiteLabelSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalWhiteLabelSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalWhiteLabelSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalWhiteLabelSettingsDto(val *AdditionalWhiteLabelSettingsDto) *NullableAdditionalWhiteLabelSettingsDto {
	return &NullableAdditionalWhiteLabelSettingsDto{value: val, isSet: true}
}

func (v NullableAdditionalWhiteLabelSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalWhiteLabelSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

