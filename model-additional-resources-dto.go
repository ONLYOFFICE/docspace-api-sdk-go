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

// checks if the AdditionalResourcesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalResourcesDto{}

// AdditionalResourcesDto Which of the ONLYOFFICE help and community entries the interface may offer, installation-wide, in the shape the  reset of these flags returns.
type AdditionalResourcesDto struct {
	// Whether the sample documents that ONLYOFFICE ships may be placed in a new user's Documents. Unlike the link  flags below it depends on nothing that has to be configured, so its built-in value is always `true`.
	StartDocsEnabled *bool `json:"startDocsEnabled,omitempty"`
	// Whether the interface may offer the Help Center entry. It is `false` both when the entry was switched off  for the installation and when the installation configures no Help Center address at all; the addresses  themselves are not part of this answer and arrive in `externalResources` of `GET api/2.0/settings`.
	HelpCenterEnabled *bool `json:"helpCenterEnabled,omitempty"`
	// Whether the interface may offer the Feedback and Support entry, `false` for the same two reasons as  `helpCenterEnabled`.
	FeedbackAndSupportEnabled *bool `json:"feedbackAndSupportEnabled,omitempty"`
	// Whether the interface may offer the user forum entry, `false` for the same two reasons as  `helpCenterEnabled`.
	UserForumEnabled *bool `json:"userForumEnabled,omitempty"`
	// Whether the interface may offer the Video Guides entry, `false` for the same two reasons as  `helpCenterEnabled`.
	VideoGuidesEnabled *bool `json:"videoGuidesEnabled,omitempty"`
	// Whether the interface may offer the License Agreements entry, `false` for the same two reasons as  `helpCenterEnabled`.
	LicenseAgreementsEnabled *bool `json:"licenseAgreementsEnabled,omitempty"`
	// When these flags were last stored. Flags that were never stored report the moment they were read, and the  answer of the reset operation reports `0001-01-01T00:00:00`.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewAdditionalResourcesDto instantiates a new AdditionalResourcesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalResourcesDto() *AdditionalResourcesDto {
	this := AdditionalResourcesDto{}
	return &this
}

// NewAdditionalResourcesDtoWithDefaults instantiates a new AdditionalResourcesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalResourcesDtoWithDefaults() *AdditionalResourcesDto {
	this := AdditionalResourcesDto{}
	return &this
}

// GetStartDocsEnabled returns the StartDocsEnabled field value if set, zero value otherwise.
func (o *AdditionalResourcesDto) GetStartDocsEnabled() bool {
	if o == nil || IsNil(o.StartDocsEnabled) {
		var ret bool
		return ret
	}
	return *o.StartDocsEnabled
}

// GetStartDocsEnabledOk returns a tuple with the StartDocsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalResourcesDto) GetStartDocsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StartDocsEnabled) {
		return nil, false
	}
	return o.StartDocsEnabled, true
}

// HasStartDocsEnabled returns a boolean if a field has been set.
func (o *AdditionalResourcesDto) IsStartDocsEnabledSet() bool {
	if o != nil && !IsNil(o.StartDocsEnabled) {
		return true
	}

	return false
}

// SetStartDocsEnabled gets a reference to the given bool and assigns it to the StartDocsEnabled field.
func (o *AdditionalResourcesDto) SetStartDocsEnabled(v bool) {
	o.StartDocsEnabled = &v
}

// GetHelpCenterEnabled returns the HelpCenterEnabled field value if set, zero value otherwise.
func (o *AdditionalResourcesDto) GetHelpCenterEnabled() bool {
	if o == nil || IsNil(o.HelpCenterEnabled) {
		var ret bool
		return ret
	}
	return *o.HelpCenterEnabled
}

// GetHelpCenterEnabledOk returns a tuple with the HelpCenterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalResourcesDto) GetHelpCenterEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HelpCenterEnabled) {
		return nil, false
	}
	return o.HelpCenterEnabled, true
}

// HasHelpCenterEnabled returns a boolean if a field has been set.
func (o *AdditionalResourcesDto) IsHelpCenterEnabledSet() bool {
	if o != nil && !IsNil(o.HelpCenterEnabled) {
		return true
	}

	return false
}

// SetHelpCenterEnabled gets a reference to the given bool and assigns it to the HelpCenterEnabled field.
func (o *AdditionalResourcesDto) SetHelpCenterEnabled(v bool) {
	o.HelpCenterEnabled = &v
}

// GetFeedbackAndSupportEnabled returns the FeedbackAndSupportEnabled field value if set, zero value otherwise.
func (o *AdditionalResourcesDto) GetFeedbackAndSupportEnabled() bool {
	if o == nil || IsNil(o.FeedbackAndSupportEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedbackAndSupportEnabled
}

// GetFeedbackAndSupportEnabledOk returns a tuple with the FeedbackAndSupportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalResourcesDto) GetFeedbackAndSupportEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedbackAndSupportEnabled) {
		return nil, false
	}
	return o.FeedbackAndSupportEnabled, true
}

// HasFeedbackAndSupportEnabled returns a boolean if a field has been set.
func (o *AdditionalResourcesDto) IsFeedbackAndSupportEnabledSet() bool {
	if o != nil && !IsNil(o.FeedbackAndSupportEnabled) {
		return true
	}

	return false
}

// SetFeedbackAndSupportEnabled gets a reference to the given bool and assigns it to the FeedbackAndSupportEnabled field.
func (o *AdditionalResourcesDto) SetFeedbackAndSupportEnabled(v bool) {
	o.FeedbackAndSupportEnabled = &v
}

// GetUserForumEnabled returns the UserForumEnabled field value if set, zero value otherwise.
func (o *AdditionalResourcesDto) GetUserForumEnabled() bool {
	if o == nil || IsNil(o.UserForumEnabled) {
		var ret bool
		return ret
	}
	return *o.UserForumEnabled
}

// GetUserForumEnabledOk returns a tuple with the UserForumEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalResourcesDto) GetUserForumEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UserForumEnabled) {
		return nil, false
	}
	return o.UserForumEnabled, true
}

// HasUserForumEnabled returns a boolean if a field has been set.
func (o *AdditionalResourcesDto) IsUserForumEnabledSet() bool {
	if o != nil && !IsNil(o.UserForumEnabled) {
		return true
	}

	return false
}

// SetUserForumEnabled gets a reference to the given bool and assigns it to the UserForumEnabled field.
func (o *AdditionalResourcesDto) SetUserForumEnabled(v bool) {
	o.UserForumEnabled = &v
}

// GetVideoGuidesEnabled returns the VideoGuidesEnabled field value if set, zero value otherwise.
func (o *AdditionalResourcesDto) GetVideoGuidesEnabled() bool {
	if o == nil || IsNil(o.VideoGuidesEnabled) {
		var ret bool
		return ret
	}
	return *o.VideoGuidesEnabled
}

// GetVideoGuidesEnabledOk returns a tuple with the VideoGuidesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalResourcesDto) GetVideoGuidesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VideoGuidesEnabled) {
		return nil, false
	}
	return o.VideoGuidesEnabled, true
}

// HasVideoGuidesEnabled returns a boolean if a field has been set.
func (o *AdditionalResourcesDto) IsVideoGuidesEnabledSet() bool {
	if o != nil && !IsNil(o.VideoGuidesEnabled) {
		return true
	}

	return false
}

// SetVideoGuidesEnabled gets a reference to the given bool and assigns it to the VideoGuidesEnabled field.
func (o *AdditionalResourcesDto) SetVideoGuidesEnabled(v bool) {
	o.VideoGuidesEnabled = &v
}

// GetLicenseAgreementsEnabled returns the LicenseAgreementsEnabled field value if set, zero value otherwise.
func (o *AdditionalResourcesDto) GetLicenseAgreementsEnabled() bool {
	if o == nil || IsNil(o.LicenseAgreementsEnabled) {
		var ret bool
		return ret
	}
	return *o.LicenseAgreementsEnabled
}

// GetLicenseAgreementsEnabledOk returns a tuple with the LicenseAgreementsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalResourcesDto) GetLicenseAgreementsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LicenseAgreementsEnabled) {
		return nil, false
	}
	return o.LicenseAgreementsEnabled, true
}

// HasLicenseAgreementsEnabled returns a boolean if a field has been set.
func (o *AdditionalResourcesDto) IsLicenseAgreementsEnabledSet() bool {
	if o != nil && !IsNil(o.LicenseAgreementsEnabled) {
		return true
	}

	return false
}

// SetLicenseAgreementsEnabled gets a reference to the given bool and assigns it to the LicenseAgreementsEnabled field.
func (o *AdditionalResourcesDto) SetLicenseAgreementsEnabled(v bool) {
	o.LicenseAgreementsEnabled = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *AdditionalResourcesDto) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalResourcesDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *AdditionalResourcesDto) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *AdditionalResourcesDto) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o AdditionalResourcesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalResourcesDto) ToMap() (map[string]interface{}, error) {
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

type NullableAdditionalResourcesDto struct {
	value *AdditionalResourcesDto
	isSet bool
}

func (v NullableAdditionalResourcesDto) Get() *AdditionalResourcesDto {
	return v.value
}

func (v *NullableAdditionalResourcesDto) Set(val *AdditionalResourcesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalResourcesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalResourcesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalResourcesDto(val *AdditionalResourcesDto) *NullableAdditionalResourcesDto {
	return &NullableAdditionalResourcesDto{value: val, isSet: true}
}

func (v NullableAdditionalResourcesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalResourcesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

