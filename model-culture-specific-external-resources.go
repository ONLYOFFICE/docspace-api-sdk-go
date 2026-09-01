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

// checks if the CultureSpecificExternalResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CultureSpecificExternalResources{}

// CultureSpecificExternalResources The external resources settings.
type CultureSpecificExternalResources struct {
	// The link to the product API.
	Api *CultureSpecificExternalResource `json:"api,omitempty"`
	// The link to the common product information.
	Common *CultureSpecificExternalResource `json:"common,omitempty"`
	// The link to the forum.
	Forum *CultureSpecificExternalResource `json:"forum,omitempty"`
	// The link to the Help Center.
	Helpcenter *CultureSpecificExternalResource `json:"helpcenter,omitempty"`
	// The link to the product integrations.
	Integrations *CultureSpecificExternalResource `json:"integrations,omitempty"`
	// The link to the product website.
	Site *CultureSpecificExternalResource `json:"site,omitempty"`
	// The link to the product social nerworks.
	SocialNetworks *CultureSpecificExternalResource `json:"socialNetworks,omitempty"`
	// The link to the product support.
	Support *CultureSpecificExternalResource `json:"support,omitempty"`
	// The link to the video guides.
	Videoguides *CultureSpecificExternalResource `json:"videoguides,omitempty"`
}

// NewCultureSpecificExternalResources instantiates a new CultureSpecificExternalResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCultureSpecificExternalResources() *CultureSpecificExternalResources {
	this := CultureSpecificExternalResources{}
	return &this
}

// NewCultureSpecificExternalResourcesWithDefaults instantiates a new CultureSpecificExternalResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCultureSpecificExternalResourcesWithDefaults() *CultureSpecificExternalResources {
	this := CultureSpecificExternalResources{}
	return &this
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *CultureSpecificExternalResources) GetApi() CultureSpecificExternalResource {
	if o == nil || IsNil(o.Api) {
		var ret CultureSpecificExternalResource
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureSpecificExternalResources) GetApiOk() (*CultureSpecificExternalResource, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *CultureSpecificExternalResources) IsApiSet() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given CultureSpecificExternalResource and assigns it to the Api field.
func (o *CultureSpecificExternalResources) SetApi(v CultureSpecificExternalResource) {
	o.Api = &v
}

// GetCommon returns the Common field value if set, zero value otherwise.
func (o *CultureSpecificExternalResources) GetCommon() CultureSpecificExternalResource {
	if o == nil || IsNil(o.Common) {
		var ret CultureSpecificExternalResource
		return ret
	}
	return *o.Common
}

// GetCommonOk returns a tuple with the Common field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureSpecificExternalResources) GetCommonOk() (*CultureSpecificExternalResource, bool) {
	if o == nil || IsNil(o.Common) {
		return nil, false
	}
	return o.Common, true
}

// HasCommon returns a boolean if a field has been set.
func (o *CultureSpecificExternalResources) IsCommonSet() bool {
	if o != nil && !IsNil(o.Common) {
		return true
	}

	return false
}

// SetCommon gets a reference to the given CultureSpecificExternalResource and assigns it to the Common field.
func (o *CultureSpecificExternalResources) SetCommon(v CultureSpecificExternalResource) {
	o.Common = &v
}

// GetForum returns the Forum field value if set, zero value otherwise.
func (o *CultureSpecificExternalResources) GetForum() CultureSpecificExternalResource {
	if o == nil || IsNil(o.Forum) {
		var ret CultureSpecificExternalResource
		return ret
	}
	return *o.Forum
}

// GetForumOk returns a tuple with the Forum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureSpecificExternalResources) GetForumOk() (*CultureSpecificExternalResource, bool) {
	if o == nil || IsNil(o.Forum) {
		return nil, false
	}
	return o.Forum, true
}

// HasForum returns a boolean if a field has been set.
func (o *CultureSpecificExternalResources) IsForumSet() bool {
	if o != nil && !IsNil(o.Forum) {
		return true
	}

	return false
}

// SetForum gets a reference to the given CultureSpecificExternalResource and assigns it to the Forum field.
func (o *CultureSpecificExternalResources) SetForum(v CultureSpecificExternalResource) {
	o.Forum = &v
}

// GetHelpcenter returns the Helpcenter field value if set, zero value otherwise.
func (o *CultureSpecificExternalResources) GetHelpcenter() CultureSpecificExternalResource {
	if o == nil || IsNil(o.Helpcenter) {
		var ret CultureSpecificExternalResource
		return ret
	}
	return *o.Helpcenter
}

// GetHelpcenterOk returns a tuple with the Helpcenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureSpecificExternalResources) GetHelpcenterOk() (*CultureSpecificExternalResource, bool) {
	if o == nil || IsNil(o.Helpcenter) {
		return nil, false
	}
	return o.Helpcenter, true
}

// HasHelpcenter returns a boolean if a field has been set.
func (o *CultureSpecificExternalResources) IsHelpcenterSet() bool {
	if o != nil && !IsNil(o.Helpcenter) {
		return true
	}

	return false
}

// SetHelpcenter gets a reference to the given CultureSpecificExternalResource and assigns it to the Helpcenter field.
func (o *CultureSpecificExternalResources) SetHelpcenter(v CultureSpecificExternalResource) {
	o.Helpcenter = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *CultureSpecificExternalResources) GetIntegrations() CultureSpecificExternalResource {
	if o == nil || IsNil(o.Integrations) {
		var ret CultureSpecificExternalResource
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureSpecificExternalResources) GetIntegrationsOk() (*CultureSpecificExternalResource, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *CultureSpecificExternalResources) IsIntegrationsSet() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given CultureSpecificExternalResource and assigns it to the Integrations field.
func (o *CultureSpecificExternalResources) SetIntegrations(v CultureSpecificExternalResource) {
	o.Integrations = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *CultureSpecificExternalResources) GetSite() CultureSpecificExternalResource {
	if o == nil || IsNil(o.Site) {
		var ret CultureSpecificExternalResource
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureSpecificExternalResources) GetSiteOk() (*CultureSpecificExternalResource, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *CultureSpecificExternalResources) IsSiteSet() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given CultureSpecificExternalResource and assigns it to the Site field.
func (o *CultureSpecificExternalResources) SetSite(v CultureSpecificExternalResource) {
	o.Site = &v
}

// GetSocialNetworks returns the SocialNetworks field value if set, zero value otherwise.
func (o *CultureSpecificExternalResources) GetSocialNetworks() CultureSpecificExternalResource {
	if o == nil || IsNil(o.SocialNetworks) {
		var ret CultureSpecificExternalResource
		return ret
	}
	return *o.SocialNetworks
}

// GetSocialNetworksOk returns a tuple with the SocialNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureSpecificExternalResources) GetSocialNetworksOk() (*CultureSpecificExternalResource, bool) {
	if o == nil || IsNil(o.SocialNetworks) {
		return nil, false
	}
	return o.SocialNetworks, true
}

// HasSocialNetworks returns a boolean if a field has been set.
func (o *CultureSpecificExternalResources) IsSocialNetworksSet() bool {
	if o != nil && !IsNil(o.SocialNetworks) {
		return true
	}

	return false
}

// SetSocialNetworks gets a reference to the given CultureSpecificExternalResource and assigns it to the SocialNetworks field.
func (o *CultureSpecificExternalResources) SetSocialNetworks(v CultureSpecificExternalResource) {
	o.SocialNetworks = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *CultureSpecificExternalResources) GetSupport() CultureSpecificExternalResource {
	if o == nil || IsNil(o.Support) {
		var ret CultureSpecificExternalResource
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureSpecificExternalResources) GetSupportOk() (*CultureSpecificExternalResource, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *CultureSpecificExternalResources) IsSupportSet() bool {
	if o != nil && !IsNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given CultureSpecificExternalResource and assigns it to the Support field.
func (o *CultureSpecificExternalResources) SetSupport(v CultureSpecificExternalResource) {
	o.Support = &v
}

// GetVideoguides returns the Videoguides field value if set, zero value otherwise.
func (o *CultureSpecificExternalResources) GetVideoguides() CultureSpecificExternalResource {
	if o == nil || IsNil(o.Videoguides) {
		var ret CultureSpecificExternalResource
		return ret
	}
	return *o.Videoguides
}

// GetVideoguidesOk returns a tuple with the Videoguides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureSpecificExternalResources) GetVideoguidesOk() (*CultureSpecificExternalResource, bool) {
	if o == nil || IsNil(o.Videoguides) {
		return nil, false
	}
	return o.Videoguides, true
}

// HasVideoguides returns a boolean if a field has been set.
func (o *CultureSpecificExternalResources) IsVideoguidesSet() bool {
	if o != nil && !IsNil(o.Videoguides) {
		return true
	}

	return false
}

// SetVideoguides gets a reference to the given CultureSpecificExternalResource and assigns it to the Videoguides field.
func (o *CultureSpecificExternalResources) SetVideoguides(v CultureSpecificExternalResource) {
	o.Videoguides = &v
}

func (o CultureSpecificExternalResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CultureSpecificExternalResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !IsNil(o.Common) {
		toSerialize["common"] = o.Common
	}
	if !IsNil(o.Forum) {
		toSerialize["forum"] = o.Forum
	}
	if !IsNil(o.Helpcenter) {
		toSerialize["helpcenter"] = o.Helpcenter
	}
	if !IsNil(o.Integrations) {
		toSerialize["integrations"] = o.Integrations
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.SocialNetworks) {
		toSerialize["socialNetworks"] = o.SocialNetworks
	}
	if !IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	if !IsNil(o.Videoguides) {
		toSerialize["videoguides"] = o.Videoguides
	}
	return toSerialize, nil
}

type NullableCultureSpecificExternalResources struct {
	value *CultureSpecificExternalResources
	isSet bool
}

func (v NullableCultureSpecificExternalResources) Get() *CultureSpecificExternalResources {
	return v.value
}

func (v *NullableCultureSpecificExternalResources) Set(val *CultureSpecificExternalResources) {
	v.value = val
	v.isSet = true
}

func (v NullableCultureSpecificExternalResources) IsSet() bool {
	return v.isSet
}

func (v *NullableCultureSpecificExternalResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCultureSpecificExternalResources(val *CultureSpecificExternalResources) *NullableCultureSpecificExternalResources {
	return &NullableCultureSpecificExternalResources{value: val, isSet: true}
}

func (v NullableCultureSpecificExternalResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCultureSpecificExternalResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

