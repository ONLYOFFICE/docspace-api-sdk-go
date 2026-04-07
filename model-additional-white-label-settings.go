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

// checks if the AdditionalWhiteLabelSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalWhiteLabelSettings{}

// AdditionalWhiteLabelSettings The additional white label settings.
type AdditionalWhiteLabelSettings struct {
	// Specifies if the sample documents are displayed or hidden.
	StartDocsEnabled *bool `json:"startDocsEnabled,omitempty"`
	// Specifies if the Help Center link is available or not.
	HelpCenterEnabled *bool `json:"helpCenterEnabled,omitempty"`
	// Specifies if the Feedback & Support link is available or not.
	FeedbackAndSupportEnabled *bool `json:"feedbackAndSupportEnabled,omitempty"`
	// Specifies if the user forum is available or not.
	UserForumEnabled *bool `json:"userForumEnabled,omitempty"`
	// Specifies if the Video Guides link is available or not.
	VideoGuidesEnabled *bool `json:"videoGuidesEnabled,omitempty"`
	// Specifies if the License Agreements link is available or not.
	LicenseAgreementsEnabled *bool `json:"licenseAgreementsEnabled,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewAdditionalWhiteLabelSettings instantiates a new AdditionalWhiteLabelSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalWhiteLabelSettings() *AdditionalWhiteLabelSettings {
	this := AdditionalWhiteLabelSettings{}
	return &this
}

// NewAdditionalWhiteLabelSettingsWithDefaults instantiates a new AdditionalWhiteLabelSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalWhiteLabelSettingsWithDefaults() *AdditionalWhiteLabelSettings {
	this := AdditionalWhiteLabelSettings{}
	return &this
}

// GetStartDocsEnabled returns the StartDocsEnabled field value if set, zero value otherwise.
func (o *AdditionalWhiteLabelSettings) GetStartDocsEnabled() bool {
	if o == nil || IsNil(o.StartDocsEnabled) {
		var ret bool
		return ret
	}
	return *o.StartDocsEnabled
}

// GetStartDocsEnabledOk returns a tuple with the StartDocsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettings) GetStartDocsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StartDocsEnabled) {
		return nil, false
	}
	return o.StartDocsEnabled, true
}

// HasStartDocsEnabled returns a boolean if a field has been set.
func (o *AdditionalWhiteLabelSettings) IsStartDocsEnabledSet() bool {
	if o != nil && !IsNil(o.StartDocsEnabled) {
		return true
	}

	return false
}

// SetStartDocsEnabled gets a reference to the given bool and assigns it to the StartDocsEnabled field.
func (o *AdditionalWhiteLabelSettings) SetStartDocsEnabled(v bool) {
	o.StartDocsEnabled = &v
}

// GetHelpCenterEnabled returns the HelpCenterEnabled field value if set, zero value otherwise.
func (o *AdditionalWhiteLabelSettings) GetHelpCenterEnabled() bool {
	if o == nil || IsNil(o.HelpCenterEnabled) {
		var ret bool
		return ret
	}
	return *o.HelpCenterEnabled
}

// GetHelpCenterEnabledOk returns a tuple with the HelpCenterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettings) GetHelpCenterEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HelpCenterEnabled) {
		return nil, false
	}
	return o.HelpCenterEnabled, true
}

// HasHelpCenterEnabled returns a boolean if a field has been set.
func (o *AdditionalWhiteLabelSettings) IsHelpCenterEnabledSet() bool {
	if o != nil && !IsNil(o.HelpCenterEnabled) {
		return true
	}

	return false
}

// SetHelpCenterEnabled gets a reference to the given bool and assigns it to the HelpCenterEnabled field.
func (o *AdditionalWhiteLabelSettings) SetHelpCenterEnabled(v bool) {
	o.HelpCenterEnabled = &v
}

// GetFeedbackAndSupportEnabled returns the FeedbackAndSupportEnabled field value if set, zero value otherwise.
func (o *AdditionalWhiteLabelSettings) GetFeedbackAndSupportEnabled() bool {
	if o == nil || IsNil(o.FeedbackAndSupportEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedbackAndSupportEnabled
}

// GetFeedbackAndSupportEnabledOk returns a tuple with the FeedbackAndSupportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettings) GetFeedbackAndSupportEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedbackAndSupportEnabled) {
		return nil, false
	}
	return o.FeedbackAndSupportEnabled, true
}

// HasFeedbackAndSupportEnabled returns a boolean if a field has been set.
func (o *AdditionalWhiteLabelSettings) IsFeedbackAndSupportEnabledSet() bool {
	if o != nil && !IsNil(o.FeedbackAndSupportEnabled) {
		return true
	}

	return false
}

// SetFeedbackAndSupportEnabled gets a reference to the given bool and assigns it to the FeedbackAndSupportEnabled field.
func (o *AdditionalWhiteLabelSettings) SetFeedbackAndSupportEnabled(v bool) {
	o.FeedbackAndSupportEnabled = &v
}

// GetUserForumEnabled returns the UserForumEnabled field value if set, zero value otherwise.
func (o *AdditionalWhiteLabelSettings) GetUserForumEnabled() bool {
	if o == nil || IsNil(o.UserForumEnabled) {
		var ret bool
		return ret
	}
	return *o.UserForumEnabled
}

// GetUserForumEnabledOk returns a tuple with the UserForumEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettings) GetUserForumEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UserForumEnabled) {
		return nil, false
	}
	return o.UserForumEnabled, true
}

// HasUserForumEnabled returns a boolean if a field has been set.
func (o *AdditionalWhiteLabelSettings) IsUserForumEnabledSet() bool {
	if o != nil && !IsNil(o.UserForumEnabled) {
		return true
	}

	return false
}

// SetUserForumEnabled gets a reference to the given bool and assigns it to the UserForumEnabled field.
func (o *AdditionalWhiteLabelSettings) SetUserForumEnabled(v bool) {
	o.UserForumEnabled = &v
}

// GetVideoGuidesEnabled returns the VideoGuidesEnabled field value if set, zero value otherwise.
func (o *AdditionalWhiteLabelSettings) GetVideoGuidesEnabled() bool {
	if o == nil || IsNil(o.VideoGuidesEnabled) {
		var ret bool
		return ret
	}
	return *o.VideoGuidesEnabled
}

// GetVideoGuidesEnabledOk returns a tuple with the VideoGuidesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettings) GetVideoGuidesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VideoGuidesEnabled) {
		return nil, false
	}
	return o.VideoGuidesEnabled, true
}

// HasVideoGuidesEnabled returns a boolean if a field has been set.
func (o *AdditionalWhiteLabelSettings) IsVideoGuidesEnabledSet() bool {
	if o != nil && !IsNil(o.VideoGuidesEnabled) {
		return true
	}

	return false
}

// SetVideoGuidesEnabled gets a reference to the given bool and assigns it to the VideoGuidesEnabled field.
func (o *AdditionalWhiteLabelSettings) SetVideoGuidesEnabled(v bool) {
	o.VideoGuidesEnabled = &v
}

// GetLicenseAgreementsEnabled returns the LicenseAgreementsEnabled field value if set, zero value otherwise.
func (o *AdditionalWhiteLabelSettings) GetLicenseAgreementsEnabled() bool {
	if o == nil || IsNil(o.LicenseAgreementsEnabled) {
		var ret bool
		return ret
	}
	return *o.LicenseAgreementsEnabled
}

// GetLicenseAgreementsEnabledOk returns a tuple with the LicenseAgreementsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettings) GetLicenseAgreementsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LicenseAgreementsEnabled) {
		return nil, false
	}
	return o.LicenseAgreementsEnabled, true
}

// HasLicenseAgreementsEnabled returns a boolean if a field has been set.
func (o *AdditionalWhiteLabelSettings) IsLicenseAgreementsEnabledSet() bool {
	if o != nil && !IsNil(o.LicenseAgreementsEnabled) {
		return true
	}

	return false
}

// SetLicenseAgreementsEnabled gets a reference to the given bool and assigns it to the LicenseAgreementsEnabled field.
func (o *AdditionalWhiteLabelSettings) SetLicenseAgreementsEnabled(v bool) {
	o.LicenseAgreementsEnabled = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *AdditionalWhiteLabelSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalWhiteLabelSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *AdditionalWhiteLabelSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *AdditionalWhiteLabelSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o AdditionalWhiteLabelSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalWhiteLabelSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDocsEnabled) {
		toSerialize["startDocsEnabled"] = o.StartDocsEnabled
	}
	if !IsNil(o.HelpCenterEnabled) {
		toSerialize["helpCenterEnabled"] = o.HelpCenterEnabled
	}
	if !IsNil(o.FeedbackAndSupportEnabled) {
		toSerialize["feedbackAndSupportEnabled"] = o.FeedbackAndSupportEnabled
	}
	if !IsNil(o.UserForumEnabled) {
		toSerialize["userForumEnabled"] = o.UserForumEnabled
	}
	if !IsNil(o.VideoGuidesEnabled) {
		toSerialize["videoGuidesEnabled"] = o.VideoGuidesEnabled
	}
	if !IsNil(o.LicenseAgreementsEnabled) {
		toSerialize["licenseAgreementsEnabled"] = o.LicenseAgreementsEnabled
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableAdditionalWhiteLabelSettings struct {
	value *AdditionalWhiteLabelSettings
	isSet bool
}

func (v NullableAdditionalWhiteLabelSettings) Get() *AdditionalWhiteLabelSettings {
	return v.value
}

func (v *NullableAdditionalWhiteLabelSettings) Set(val *AdditionalWhiteLabelSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalWhiteLabelSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalWhiteLabelSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalWhiteLabelSettings(val *AdditionalWhiteLabelSettings) *NullableAdditionalWhiteLabelSettings {
	return &NullableAdditionalWhiteLabelSettings{value: val, isSet: true}
}

func (v NullableAdditionalWhiteLabelSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalWhiteLabelSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

