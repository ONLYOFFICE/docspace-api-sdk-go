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

// checks if the DbTenantPartner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbTenantPartner{}

// DbTenantPartner The database tenant partner parameters.
type DbTenantPartner struct {
	// The tenant ID.
	TenantId *int32 `json:"tenantId,omitempty"`
	// The partner ID.
	PartnerId NullableString `json:"partnerId,omitempty"`
	// The affiliate ID.
	AffiliateId NullableString `json:"affiliateId,omitempty"`
	// The tenant partner campaign.
	Campaign NullableString `json:"campaign,omitempty"`
}

// NewDbTenantPartner instantiates a new DbTenantPartner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbTenantPartner() *DbTenantPartner {
	this := DbTenantPartner{}
	return &this
}

// NewDbTenantPartnerWithDefaults instantiates a new DbTenantPartner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbTenantPartnerWithDefaults() *DbTenantPartner {
	this := DbTenantPartner{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *DbTenantPartner) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenantPartner) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DbTenantPartner) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *DbTenantPartner) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenantPartner) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId.Get()) {
		var ret string
		return ret
	}
	return *o.PartnerId.Get()
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenantPartner) GetPartnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerId.Get(), o.PartnerId.IsSet()
}

// HasPartnerId returns a boolean if a field has been set.
func (o *DbTenantPartner) IsPartnerIdSet() bool {
	if o != nil && o.PartnerId.IsSet() {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given NullableString and assigns it to the PartnerId field.
func (o *DbTenantPartner) SetPartnerId(v string) {
	o.PartnerId.Set(&v)
}
// SetPartnerIdNil sets the value for PartnerId to be an explicit nil
func (o *DbTenantPartner) SetPartnerIdNil() {
	o.PartnerId.Set(nil)
}

// UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
func (o *DbTenantPartner) UnsetPartnerId() {
	o.PartnerId.Unset()
}

// GetAffiliateId returns the AffiliateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenantPartner) GetAffiliateId() string {
	if o == nil || IsNil(o.AffiliateId.Get()) {
		var ret string
		return ret
	}
	return *o.AffiliateId.Get()
}

// GetAffiliateIdOk returns a tuple with the AffiliateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenantPartner) GetAffiliateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AffiliateId.Get(), o.AffiliateId.IsSet()
}

// HasAffiliateId returns a boolean if a field has been set.
func (o *DbTenantPartner) IsAffiliateIdSet() bool {
	if o != nil && o.AffiliateId.IsSet() {
		return true
	}

	return false
}

// SetAffiliateId gets a reference to the given NullableString and assigns it to the AffiliateId field.
func (o *DbTenantPartner) SetAffiliateId(v string) {
	o.AffiliateId.Set(&v)
}
// SetAffiliateIdNil sets the value for AffiliateId to be an explicit nil
func (o *DbTenantPartner) SetAffiliateIdNil() {
	o.AffiliateId.Set(nil)
}

// UnsetAffiliateId ensures that no value is present for AffiliateId, not even an explicit nil
func (o *DbTenantPartner) UnsetAffiliateId() {
	o.AffiliateId.Unset()
}

// GetCampaign returns the Campaign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenantPartner) GetCampaign() string {
	if o == nil || IsNil(o.Campaign.Get()) {
		var ret string
		return ret
	}
	return *o.Campaign.Get()
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenantPartner) GetCampaignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaign.Get(), o.Campaign.IsSet()
}

// HasCampaign returns a boolean if a field has been set.
func (o *DbTenantPartner) IsCampaignSet() bool {
	if o != nil && o.Campaign.IsSet() {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given NullableString and assigns it to the Campaign field.
func (o *DbTenantPartner) SetCampaign(v string) {
	o.Campaign.Set(&v)
}
// SetCampaignNil sets the value for Campaign to be an explicit nil
func (o *DbTenantPartner) SetCampaignNil() {
	o.Campaign.Set(nil)
}

// UnsetCampaign ensures that no value is present for Campaign, not even an explicit nil
func (o *DbTenantPartner) UnsetCampaign() {
	o.Campaign.Unset()
}

func (o DbTenantPartner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbTenantPartner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.PartnerId.IsSet() {
		toSerialize["partnerId"] = o.PartnerId.Get()
	}
	if o.AffiliateId.IsSet() {
		toSerialize["affiliateId"] = o.AffiliateId.Get()
	}
	if o.Campaign.IsSet() {
		toSerialize["campaign"] = o.Campaign.Get()
	}
	return toSerialize, nil
}

type NullableDbTenantPartner struct {
	value *DbTenantPartner
	isSet bool
}

func (v NullableDbTenantPartner) Get() *DbTenantPartner {
	return v.value
}

func (v *NullableDbTenantPartner) Set(val *DbTenantPartner) {
	v.value = val
	v.isSet = true
}

func (v NullableDbTenantPartner) IsSet() bool {
	return v.isSet
}

func (v *NullableDbTenantPartner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbTenantPartner(val *DbTenantPartner) *NullableDbTenantPartner {
	return &NullableDbTenantPartner{value: val, isSet: true}
}

func (v NullableDbTenantPartner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbTenantPartner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

