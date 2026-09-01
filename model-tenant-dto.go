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

// checks if the TenantDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantDto{}

// TenantDto The tenant parameters.
type TenantDto struct {
	// The affiliate ID.
	AffiliateId NullableString `json:"affiliateId,omitempty"`
	// The tenant alias.
	TenantAlias NullableString `json:"tenantAlias,omitempty"`
	// Specifies if the calls are available for this tenant or not.
	Calls *bool `json:"calls,omitempty"`
	// The tenant campaign.
	Campaign NullableString `json:"campaign,omitempty"`
	// The tenant creation date and time.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
	// The hosted region.
	HostedRegion NullableString `json:"hostedRegion,omitempty"`
	// The tenant ID.
	TenantId *int32 `json:"tenantId,omitempty"`
	// The tenant industry.
	Industry *TenantIndustry `json:"industry,omitempty"`
	// The tenant language.
	Language NullableString `json:"language,omitempty"`
	// The date and time when the tenant was last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
	// The tenant mapped domain.
	MappedDomain NullableString `json:"mappedDomain,omitempty"`
	// The tenant name.
	Name NullableString `json:"name,omitempty"`
	// The tenant owner ID.
	OwnerId *string `json:"ownerId,omitempty"`
	// The tenant payment ID.
	PaymentId NullableString `json:"paymentId,omitempty"`
	// Specifies if the ONLYOFFICE newsletter is allowed or not.
	Spam *bool `json:"spam,omitempty"`
	// The tenant status.
	Status *TenantStatus `json:"status,omitempty"`
	// The date and time when the tenant status was changed.
	StatusChangeDate *time.Time `json:"statusChangeDate,omitempty"`
	// The tenant time zone.
	TimeZone NullableString `json:"timeZone,omitempty"`
	// The list of tenant trusted domains.
	TrustedDomains []string `json:"trustedDomains,omitempty"`
	// The tenant trusted domains in the string format.
	TrustedDomainsRaw NullableString `json:"trustedDomainsRaw,omitempty"`
	// The type of the tenant trusted domains.
	TrustedDomainsType *TenantTrustedDomainsType `json:"trustedDomainsType,omitempty"`
	// The tenant version
	Version *int32 `json:"version,omitempty"`
	// The date and time when the tenant version was changed.
	VersionChanged *time.Time `json:"versionChanged,omitempty"`
	// The tenant AWS region.
	Region NullableString `json:"region,omitempty"`
}

// NewTenantDto instantiates a new TenantDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantDto() *TenantDto {
	this := TenantDto{}
	return &this
}

// NewTenantDtoWithDefaults instantiates a new TenantDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantDtoWithDefaults() *TenantDto {
	this := TenantDto{}
	return &this
}

// GetAffiliateId returns the AffiliateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetAffiliateId() string {
	if o == nil || IsNil(o.AffiliateId.Get()) {
		var ret string
		return ret
	}
	return *o.AffiliateId.Get()
}

// GetAffiliateIdOk returns a tuple with the AffiliateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetAffiliateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AffiliateId.Get(), o.AffiliateId.IsSet()
}

// HasAffiliateId returns a boolean if a field has been set.
func (o *TenantDto) IsAffiliateIdSet() bool {
	if o != nil && o.AffiliateId.IsSet() {
		return true
	}

	return false
}

// SetAffiliateId gets a reference to the given NullableString and assigns it to the AffiliateId field.
func (o *TenantDto) SetAffiliateId(v string) {
	o.AffiliateId.Set(&v)
}
// SetAffiliateIdNil sets the value for AffiliateId to be an explicit nil
func (o *TenantDto) SetAffiliateIdNil() {
	o.AffiliateId.Set(nil)
}

// UnsetAffiliateId ensures that no value is present for AffiliateId, not even an explicit nil
func (o *TenantDto) UnsetAffiliateId() {
	o.AffiliateId.Unset()
}

// GetTenantAlias returns the TenantAlias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetTenantAlias() string {
	if o == nil || IsNil(o.TenantAlias.Get()) {
		var ret string
		return ret
	}
	return *o.TenantAlias.Get()
}

// GetTenantAliasOk returns a tuple with the TenantAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetTenantAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantAlias.Get(), o.TenantAlias.IsSet()
}

// HasTenantAlias returns a boolean if a field has been set.
func (o *TenantDto) IsTenantAliasSet() bool {
	if o != nil && o.TenantAlias.IsSet() {
		return true
	}

	return false
}

// SetTenantAlias gets a reference to the given NullableString and assigns it to the TenantAlias field.
func (o *TenantDto) SetTenantAlias(v string) {
	o.TenantAlias.Set(&v)
}
// SetTenantAliasNil sets the value for TenantAlias to be an explicit nil
func (o *TenantDto) SetTenantAliasNil() {
	o.TenantAlias.Set(nil)
}

// UnsetTenantAlias ensures that no value is present for TenantAlias, not even an explicit nil
func (o *TenantDto) UnsetTenantAlias() {
	o.TenantAlias.Unset()
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *TenantDto) GetCalls() bool {
	if o == nil || IsNil(o.Calls) {
		var ret bool
		return ret
	}
	return *o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetCallsOk() (*bool, bool) {
	if o == nil || IsNil(o.Calls) {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *TenantDto) IsCallsSet() bool {
	if o != nil && !IsNil(o.Calls) {
		return true
	}

	return false
}

// SetCalls gets a reference to the given bool and assigns it to the Calls field.
func (o *TenantDto) SetCalls(v bool) {
	o.Calls = &v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetCampaign() string {
	if o == nil || IsNil(o.Campaign.Get()) {
		var ret string
		return ret
	}
	return *o.Campaign.Get()
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetCampaignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaign.Get(), o.Campaign.IsSet()
}

// HasCampaign returns a boolean if a field has been set.
func (o *TenantDto) IsCampaignSet() bool {
	if o != nil && o.Campaign.IsSet() {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given NullableString and assigns it to the Campaign field.
func (o *TenantDto) SetCampaign(v string) {
	o.Campaign.Set(&v)
}
// SetCampaignNil sets the value for Campaign to be an explicit nil
func (o *TenantDto) SetCampaignNil() {
	o.Campaign.Set(nil)
}

// UnsetCampaign ensures that no value is present for Campaign, not even an explicit nil
func (o *TenantDto) UnsetCampaign() {
	o.Campaign.Unset()
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *TenantDto) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *TenantDto) IsCreationDateTimeSet() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *TenantDto) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

// GetHostedRegion returns the HostedRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetHostedRegion() string {
	if o == nil || IsNil(o.HostedRegion.Get()) {
		var ret string
		return ret
	}
	return *o.HostedRegion.Get()
}

// GetHostedRegionOk returns a tuple with the HostedRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetHostedRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostedRegion.Get(), o.HostedRegion.IsSet()
}

// HasHostedRegion returns a boolean if a field has been set.
func (o *TenantDto) IsHostedRegionSet() bool {
	if o != nil && o.HostedRegion.IsSet() {
		return true
	}

	return false
}

// SetHostedRegion gets a reference to the given NullableString and assigns it to the HostedRegion field.
func (o *TenantDto) SetHostedRegion(v string) {
	o.HostedRegion.Set(&v)
}
// SetHostedRegionNil sets the value for HostedRegion to be an explicit nil
func (o *TenantDto) SetHostedRegionNil() {
	o.HostedRegion.Set(nil)
}

// UnsetHostedRegion ensures that no value is present for HostedRegion, not even an explicit nil
func (o *TenantDto) UnsetHostedRegion() {
	o.HostedRegion.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *TenantDto) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *TenantDto) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *TenantDto) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
func (o *TenantDto) GetIndustry() TenantIndustry {
	if o == nil || IsNil(o.Industry) {
		var ret TenantIndustry
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetIndustryOk() (*TenantIndustry, bool) {
	if o == nil || IsNil(o.Industry) {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *TenantDto) IsIndustrySet() bool {
	if o != nil && !IsNil(o.Industry) {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given TenantIndustry and assigns it to the Industry field.
func (o *TenantDto) SetIndustry(v TenantIndustry) {
	o.Industry = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetLanguage() string {
	if o == nil || IsNil(o.Language.Get()) {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *TenantDto) IsLanguageSet() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *TenantDto) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *TenantDto) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *TenantDto) UnsetLanguage() {
	o.Language.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *TenantDto) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *TenantDto) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *TenantDto) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetMappedDomain returns the MappedDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetMappedDomain() string {
	if o == nil || IsNil(o.MappedDomain.Get()) {
		var ret string
		return ret
	}
	return *o.MappedDomain.Get()
}

// GetMappedDomainOk returns a tuple with the MappedDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetMappedDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MappedDomain.Get(), o.MappedDomain.IsSet()
}

// HasMappedDomain returns a boolean if a field has been set.
func (o *TenantDto) IsMappedDomainSet() bool {
	if o != nil && o.MappedDomain.IsSet() {
		return true
	}

	return false
}

// SetMappedDomain gets a reference to the given NullableString and assigns it to the MappedDomain field.
func (o *TenantDto) SetMappedDomain(v string) {
	o.MappedDomain.Set(&v)
}
// SetMappedDomainNil sets the value for MappedDomain to be an explicit nil
func (o *TenantDto) SetMappedDomainNil() {
	o.MappedDomain.Set(nil)
}

// UnsetMappedDomain ensures that no value is present for MappedDomain, not even an explicit nil
func (o *TenantDto) UnsetMappedDomain() {
	o.MappedDomain.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TenantDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TenantDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TenantDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TenantDto) UnsetName() {
	o.Name.Unset()
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *TenantDto) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *TenantDto) IsOwnerIdSet() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *TenantDto) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// HasPaymentId returns a boolean if a field has been set.
func (o *TenantDto) IsPaymentIdSet() bool {
	if o != nil && o.PaymentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given NullableString and assigns it to the PaymentId field.
func (o *TenantDto) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}
// SetPaymentIdNil sets the value for PaymentId to be an explicit nil
func (o *TenantDto) SetPaymentIdNil() {
	o.PaymentId.Set(nil)
}

// UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
func (o *TenantDto) UnsetPaymentId() {
	o.PaymentId.Unset()
}

// GetSpam returns the Spam field value if set, zero value otherwise.
func (o *TenantDto) GetSpam() bool {
	if o == nil || IsNil(o.Spam) {
		var ret bool
		return ret
	}
	return *o.Spam
}

// GetSpamOk returns a tuple with the Spam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetSpamOk() (*bool, bool) {
	if o == nil || IsNil(o.Spam) {
		return nil, false
	}
	return o.Spam, true
}

// HasSpam returns a boolean if a field has been set.
func (o *TenantDto) IsSpamSet() bool {
	if o != nil && !IsNil(o.Spam) {
		return true
	}

	return false
}

// SetSpam gets a reference to the given bool and assigns it to the Spam field.
func (o *TenantDto) SetSpam(v bool) {
	o.Spam = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TenantDto) GetStatus() TenantStatus {
	if o == nil || IsNil(o.Status) {
		var ret TenantStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetStatusOk() (*TenantStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TenantDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TenantStatus and assigns it to the Status field.
func (o *TenantDto) SetStatus(v TenantStatus) {
	o.Status = &v
}

// GetStatusChangeDate returns the StatusChangeDate field value if set, zero value otherwise.
func (o *TenantDto) GetStatusChangeDate() time.Time {
	if o == nil || IsNil(o.StatusChangeDate) {
		var ret time.Time
		return ret
	}
	return *o.StatusChangeDate
}

// GetStatusChangeDateOk returns a tuple with the StatusChangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetStatusChangeDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StatusChangeDate) {
		return nil, false
	}
	return o.StatusChangeDate, true
}

// HasStatusChangeDate returns a boolean if a field has been set.
func (o *TenantDto) IsStatusChangeDateSet() bool {
	if o != nil && !IsNil(o.StatusChangeDate) {
		return true
	}

	return false
}

// SetStatusChangeDate gets a reference to the given time.Time and assigns it to the StatusChangeDate field.
func (o *TenantDto) SetStatusChangeDate(v time.Time) {
	o.StatusChangeDate = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone.Get()) {
		var ret string
		return ret
	}
	return *o.TimeZone.Get()
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeZone.Get(), o.TimeZone.IsSet()
}

// HasTimeZone returns a boolean if a field has been set.
func (o *TenantDto) IsTimeZoneSet() bool {
	if o != nil && o.TimeZone.IsSet() {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given NullableString and assigns it to the TimeZone field.
func (o *TenantDto) SetTimeZone(v string) {
	o.TimeZone.Set(&v)
}
// SetTimeZoneNil sets the value for TimeZone to be an explicit nil
func (o *TenantDto) SetTimeZoneNil() {
	o.TimeZone.Set(nil)
}

// UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
func (o *TenantDto) UnsetTimeZone() {
	o.TimeZone.Unset()
}

// GetTrustedDomains returns the TrustedDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetTrustedDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TrustedDomains
}

// GetTrustedDomainsOk returns a tuple with the TrustedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetTrustedDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.TrustedDomains) {
		return nil, false
	}
	return o.TrustedDomains, true
}

// HasTrustedDomains returns a boolean if a field has been set.
func (o *TenantDto) IsTrustedDomainsSet() bool {
	if o != nil && !IsNil(o.TrustedDomains) {
		return true
	}

	return false
}

// SetTrustedDomains gets a reference to the given []string and assigns it to the TrustedDomains field.
func (o *TenantDto) SetTrustedDomains(v []string) {
	o.TrustedDomains = v
}

// GetTrustedDomainsRaw returns the TrustedDomainsRaw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetTrustedDomainsRaw() string {
	if o == nil || IsNil(o.TrustedDomainsRaw.Get()) {
		var ret string
		return ret
	}
	return *o.TrustedDomainsRaw.Get()
}

// GetTrustedDomainsRawOk returns a tuple with the TrustedDomainsRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetTrustedDomainsRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrustedDomainsRaw.Get(), o.TrustedDomainsRaw.IsSet()
}

// HasTrustedDomainsRaw returns a boolean if a field has been set.
func (o *TenantDto) IsTrustedDomainsRawSet() bool {
	if o != nil && o.TrustedDomainsRaw.IsSet() {
		return true
	}

	return false
}

// SetTrustedDomainsRaw gets a reference to the given NullableString and assigns it to the TrustedDomainsRaw field.
func (o *TenantDto) SetTrustedDomainsRaw(v string) {
	o.TrustedDomainsRaw.Set(&v)
}
// SetTrustedDomainsRawNil sets the value for TrustedDomainsRaw to be an explicit nil
func (o *TenantDto) SetTrustedDomainsRawNil() {
	o.TrustedDomainsRaw.Set(nil)
}

// UnsetTrustedDomainsRaw ensures that no value is present for TrustedDomainsRaw, not even an explicit nil
func (o *TenantDto) UnsetTrustedDomainsRaw() {
	o.TrustedDomainsRaw.Unset()
}

// GetTrustedDomainsType returns the TrustedDomainsType field value if set, zero value otherwise.
func (o *TenantDto) GetTrustedDomainsType() TenantTrustedDomainsType {
	if o == nil || IsNil(o.TrustedDomainsType) {
		var ret TenantTrustedDomainsType
		return ret
	}
	return *o.TrustedDomainsType
}

// GetTrustedDomainsTypeOk returns a tuple with the TrustedDomainsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetTrustedDomainsTypeOk() (*TenantTrustedDomainsType, bool) {
	if o == nil || IsNil(o.TrustedDomainsType) {
		return nil, false
	}
	return o.TrustedDomainsType, true
}

// HasTrustedDomainsType returns a boolean if a field has been set.
func (o *TenantDto) IsTrustedDomainsTypeSet() bool {
	if o != nil && !IsNil(o.TrustedDomainsType) {
		return true
	}

	return false
}

// SetTrustedDomainsType gets a reference to the given TenantTrustedDomainsType and assigns it to the TrustedDomainsType field.
func (o *TenantDto) SetTrustedDomainsType(v TenantTrustedDomainsType) {
	o.TrustedDomainsType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TenantDto) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TenantDto) IsVersionSet() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *TenantDto) SetVersion(v int32) {
	o.Version = &v
}

// GetVersionChanged returns the VersionChanged field value if set, zero value otherwise.
func (o *TenantDto) GetVersionChanged() time.Time {
	if o == nil || IsNil(o.VersionChanged) {
		var ret time.Time
		return ret
	}
	return *o.VersionChanged
}

// GetVersionChangedOk returns a tuple with the VersionChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetVersionChangedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VersionChanged) {
		return nil, false
	}
	return o.VersionChanged, true
}

// HasVersionChanged returns a boolean if a field has been set.
func (o *TenantDto) IsVersionChangedSet() bool {
	if o != nil && !IsNil(o.VersionChanged) {
		return true
	}

	return false
}

// SetVersionChanged gets a reference to the given time.Time and assigns it to the VersionChanged field.
func (o *TenantDto) SetVersionChanged(v time.Time) {
	o.VersionChanged = &v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantDto) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantDto) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *TenantDto) IsRegionSet() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *TenantDto) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *TenantDto) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *TenantDto) UnsetRegion() {
	o.Region.Unset()
}

func (o TenantDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AffiliateId.IsSet() {
		toSerialize["affiliateId"] = o.AffiliateId.Get()
	}
	if o.TenantAlias.IsSet() {
		toSerialize["tenantAlias"] = o.TenantAlias.Get()
	}
	if !IsNil(o.Calls) {
		toSerialize["calls"] = o.Calls
	}
	if o.Campaign.IsSet() {
		toSerialize["campaign"] = o.Campaign.Get()
	}
	if !IsNil(o.CreationDateTime) {
		toSerialize["creationDateTime"] = o.CreationDateTime
	}
	if o.HostedRegion.IsSet() {
		toSerialize["hostedRegion"] = o.HostedRegion.Get()
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Industry) {
		toSerialize["industry"] = o.Industry
	}
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if o.MappedDomain.IsSet() {
		toSerialize["mappedDomain"] = o.MappedDomain.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.PaymentId.IsSet() {
		toSerialize["paymentId"] = o.PaymentId.Get()
	}
	if !IsNil(o.Spam) {
		toSerialize["spam"] = o.Spam
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusChangeDate) {
		toSerialize["statusChangeDate"] = o.StatusChangeDate
	}
	if o.TimeZone.IsSet() {
		toSerialize["timeZone"] = o.TimeZone.Get()
	}
	if o.TrustedDomains != nil {
		toSerialize["trustedDomains"] = o.TrustedDomains
	}
	if o.TrustedDomainsRaw.IsSet() {
		toSerialize["trustedDomainsRaw"] = o.TrustedDomainsRaw.Get()
	}
	if !IsNil(o.TrustedDomainsType) {
		toSerialize["trustedDomainsType"] = o.TrustedDomainsType
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VersionChanged) {
		toSerialize["versionChanged"] = o.VersionChanged
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	return toSerialize, nil
}

type NullableTenantDto struct {
	value *TenantDto
	isSet bool
}

func (v NullableTenantDto) Get() *TenantDto {
	return v.value
}

func (v *NullableTenantDto) Set(val *TenantDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantDto(val *TenantDto) *NullableTenantDto {
	return &NullableTenantDto{value: val, isSet: true}
}

func (v NullableTenantDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

