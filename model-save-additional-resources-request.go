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

// checks if the SaveAdditionalResourcesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveAdditionalResourcesRequest{}

// SaveAdditionalResourcesRequest Which help and community resources the interface links to.
type SaveAdditionalResourcesRequest struct {
	// Whether the getting-started documents are offered.
	StartDocsEnabled *bool `json:"startDocsEnabled,omitempty"`
	// Whether the help center is linked.
	HelpCenterEnabled *bool `json:"helpCenterEnabled,omitempty"`
	// Whether the feedback and support link is shown.
	FeedbackAndSupportEnabled *bool `json:"feedbackAndSupportEnabled,omitempty"`
	// Whether the user forum is linked.
	UserForumEnabled *bool `json:"userForumEnabled,omitempty"`
	// Whether the video guides are linked.
	VideoGuidesEnabled *bool `json:"videoGuidesEnabled,omitempty"`
	// Whether the license agreements are linked.
	LicenseAgreementsEnabled *bool `json:"licenseAgreementsEnabled,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewSaveAdditionalResourcesRequest instantiates a new SaveAdditionalResourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveAdditionalResourcesRequest() *SaveAdditionalResourcesRequest {
	this := SaveAdditionalResourcesRequest{}
	return &this
}

// NewSaveAdditionalResourcesRequestWithDefaults instantiates a new SaveAdditionalResourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveAdditionalResourcesRequestWithDefaults() *SaveAdditionalResourcesRequest {
	this := SaveAdditionalResourcesRequest{}
	return &this
}

// GetStartDocsEnabled returns the StartDocsEnabled field value if set, zero value otherwise.
func (o *SaveAdditionalResourcesRequest) GetStartDocsEnabled() bool {
	if o == nil || IsNil(o.StartDocsEnabled) {
		var ret bool
		return ret
	}
	return *o.StartDocsEnabled
}

// GetStartDocsEnabledOk returns a tuple with the StartDocsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAdditionalResourcesRequest) GetStartDocsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StartDocsEnabled) {
		return nil, false
	}
	return o.StartDocsEnabled, true
}

// HasStartDocsEnabled returns a boolean if a field has been set.
func (o *SaveAdditionalResourcesRequest) IsStartDocsEnabledSet() bool {
	if o != nil && !IsNil(o.StartDocsEnabled) {
		return true
	}

	return false
}

// SetStartDocsEnabled gets a reference to the given bool and assigns it to the StartDocsEnabled field.
func (o *SaveAdditionalResourcesRequest) SetStartDocsEnabled(v bool) {
	o.StartDocsEnabled = &v
}

// GetHelpCenterEnabled returns the HelpCenterEnabled field value if set, zero value otherwise.
func (o *SaveAdditionalResourcesRequest) GetHelpCenterEnabled() bool {
	if o == nil || IsNil(o.HelpCenterEnabled) {
		var ret bool
		return ret
	}
	return *o.HelpCenterEnabled
}

// GetHelpCenterEnabledOk returns a tuple with the HelpCenterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAdditionalResourcesRequest) GetHelpCenterEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HelpCenterEnabled) {
		return nil, false
	}
	return o.HelpCenterEnabled, true
}

// HasHelpCenterEnabled returns a boolean if a field has been set.
func (o *SaveAdditionalResourcesRequest) IsHelpCenterEnabledSet() bool {
	if o != nil && !IsNil(o.HelpCenterEnabled) {
		return true
	}

	return false
}

// SetHelpCenterEnabled gets a reference to the given bool and assigns it to the HelpCenterEnabled field.
func (o *SaveAdditionalResourcesRequest) SetHelpCenterEnabled(v bool) {
	o.HelpCenterEnabled = &v
}

// GetFeedbackAndSupportEnabled returns the FeedbackAndSupportEnabled field value if set, zero value otherwise.
func (o *SaveAdditionalResourcesRequest) GetFeedbackAndSupportEnabled() bool {
	if o == nil || IsNil(o.FeedbackAndSupportEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedbackAndSupportEnabled
}

// GetFeedbackAndSupportEnabledOk returns a tuple with the FeedbackAndSupportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAdditionalResourcesRequest) GetFeedbackAndSupportEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedbackAndSupportEnabled) {
		return nil, false
	}
	return o.FeedbackAndSupportEnabled, true
}

// HasFeedbackAndSupportEnabled returns a boolean if a field has been set.
func (o *SaveAdditionalResourcesRequest) IsFeedbackAndSupportEnabledSet() bool {
	if o != nil && !IsNil(o.FeedbackAndSupportEnabled) {
		return true
	}

	return false
}

// SetFeedbackAndSupportEnabled gets a reference to the given bool and assigns it to the FeedbackAndSupportEnabled field.
func (o *SaveAdditionalResourcesRequest) SetFeedbackAndSupportEnabled(v bool) {
	o.FeedbackAndSupportEnabled = &v
}

// GetUserForumEnabled returns the UserForumEnabled field value if set, zero value otherwise.
func (o *SaveAdditionalResourcesRequest) GetUserForumEnabled() bool {
	if o == nil || IsNil(o.UserForumEnabled) {
		var ret bool
		return ret
	}
	return *o.UserForumEnabled
}

// GetUserForumEnabledOk returns a tuple with the UserForumEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAdditionalResourcesRequest) GetUserForumEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UserForumEnabled) {
		return nil, false
	}
	return o.UserForumEnabled, true
}

// HasUserForumEnabled returns a boolean if a field has been set.
func (o *SaveAdditionalResourcesRequest) IsUserForumEnabledSet() bool {
	if o != nil && !IsNil(o.UserForumEnabled) {
		return true
	}

	return false
}

// SetUserForumEnabled gets a reference to the given bool and assigns it to the UserForumEnabled field.
func (o *SaveAdditionalResourcesRequest) SetUserForumEnabled(v bool) {
	o.UserForumEnabled = &v
}

// GetVideoGuidesEnabled returns the VideoGuidesEnabled field value if set, zero value otherwise.
func (o *SaveAdditionalResourcesRequest) GetVideoGuidesEnabled() bool {
	if o == nil || IsNil(o.VideoGuidesEnabled) {
		var ret bool
		return ret
	}
	return *o.VideoGuidesEnabled
}

// GetVideoGuidesEnabledOk returns a tuple with the VideoGuidesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAdditionalResourcesRequest) GetVideoGuidesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VideoGuidesEnabled) {
		return nil, false
	}
	return o.VideoGuidesEnabled, true
}

// HasVideoGuidesEnabled returns a boolean if a field has been set.
func (o *SaveAdditionalResourcesRequest) IsVideoGuidesEnabledSet() bool {
	if o != nil && !IsNil(o.VideoGuidesEnabled) {
		return true
	}

	return false
}

// SetVideoGuidesEnabled gets a reference to the given bool and assigns it to the VideoGuidesEnabled field.
func (o *SaveAdditionalResourcesRequest) SetVideoGuidesEnabled(v bool) {
	o.VideoGuidesEnabled = &v
}

// GetLicenseAgreementsEnabled returns the LicenseAgreementsEnabled field value if set, zero value otherwise.
func (o *SaveAdditionalResourcesRequest) GetLicenseAgreementsEnabled() bool {
	if o == nil || IsNil(o.LicenseAgreementsEnabled) {
		var ret bool
		return ret
	}
	return *o.LicenseAgreementsEnabled
}

// GetLicenseAgreementsEnabledOk returns a tuple with the LicenseAgreementsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAdditionalResourcesRequest) GetLicenseAgreementsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LicenseAgreementsEnabled) {
		return nil, false
	}
	return o.LicenseAgreementsEnabled, true
}

// HasLicenseAgreementsEnabled returns a boolean if a field has been set.
func (o *SaveAdditionalResourcesRequest) IsLicenseAgreementsEnabledSet() bool {
	if o != nil && !IsNil(o.LicenseAgreementsEnabled) {
		return true
	}

	return false
}

// SetLicenseAgreementsEnabled gets a reference to the given bool and assigns it to the LicenseAgreementsEnabled field.
func (o *SaveAdditionalResourcesRequest) SetLicenseAgreementsEnabled(v bool) {
	o.LicenseAgreementsEnabled = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *SaveAdditionalResourcesRequest) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAdditionalResourcesRequest) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *SaveAdditionalResourcesRequest) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *SaveAdditionalResourcesRequest) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o SaveAdditionalResourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveAdditionalResourcesRequest) ToMap() (map[string]interface{}, error) {
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

type NullableSaveAdditionalResourcesRequest struct {
	value *SaveAdditionalResourcesRequest
	isSet bool
}

func (v NullableSaveAdditionalResourcesRequest) Get() *SaveAdditionalResourcesRequest {
	return v.value
}

func (v *NullableSaveAdditionalResourcesRequest) Set(val *SaveAdditionalResourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveAdditionalResourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveAdditionalResourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveAdditionalResourcesRequest(val *SaveAdditionalResourcesRequest) *NullableSaveAdditionalResourcesRequest {
	return &NullableSaveAdditionalResourcesRequest{value: val, isSet: true}
}

func (v NullableSaveAdditionalResourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveAdditionalResourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

