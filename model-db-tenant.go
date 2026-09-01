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

// checks if the DbTenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbTenant{}

// DbTenant The database tenant parameters.
type DbTenant struct {
	// The tenant ID.
	Id *int32 `json:"id,omitempty"`
	// The tenant name.
	Name NullableString `json:"name,omitempty"`
	// The tenant alias.
	Alias NullableString `json:"alias,omitempty"`
	// Mapped domain
	MappedDomain NullableString `json:"mappedDomain,omitempty"`
	// The tenant version.
	Version *int32 `json:"version,omitempty"`
	// The Version_changed field.
	VersionChangedField NullableTime `json:"version_Changed,omitempty"`
	// The date and time when the version was changed.
	VersionChanged *time.Time `json:"versionChanged,omitempty"`
	// The tenant language.
	Language NullableString `json:"language,omitempty"`
	// The tenant time zone.
	TimeZone NullableString `json:"timeZone,omitempty"`
	// The tenant trusted domains raw.
	TrustedDomainsRaw NullableString `json:"trustedDomainsRaw,omitempty"`
	// The type of the tenant trusted domains.
	TrustedDomainsEnabled *TenantTrustedDomainsType `json:"trustedDomainsEnabled,omitempty"`
	// The tenant status.
	Status *TenantStatus `json:"status,omitempty"`
	// The date and time when the tenant status was changed.
	StatusChanged NullableTime `json:"statusChanged,omitempty"`
	// The hacked date and time when the tenant status was changed.
	StatusChangedHack *time.Time `json:"statusChangedHack,omitempty"`
	// The tenant creation date.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
	// The tenant owner ID.
	OwnerId NullableString `json:"ownerId,omitempty"`
	// The tenant payment ID.
	PaymentId NullableString `json:"paymentId,omitempty"`
	// The tenant industry.
	Industry *TenantIndustry `json:"industry,omitempty"`
	// The date and time when the tenant was last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
	// Specifies if the calls are available for the current tenant or not.
	Calls *bool `json:"calls,omitempty"`
	// The database tenant partner parameters.
	Partner *DbTenantPartner `json:"partner,omitempty"`
}

// NewDbTenant instantiates a new DbTenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbTenant() *DbTenant {
	this := DbTenant{}
	return &this
}

// NewDbTenantWithDefaults instantiates a new DbTenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbTenantWithDefaults() *DbTenant {
	this := DbTenant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DbTenant) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DbTenant) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DbTenant) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DbTenant) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DbTenant) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DbTenant) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DbTenant) UnsetName() {
	o.Name.Unset()
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetAlias() string {
	if o == nil || IsNil(o.Alias.Get()) {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *DbTenant) IsAliasSet() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *DbTenant) SetAlias(v string) {
	o.Alias.Set(&v)
}
// SetAliasNil sets the value for Alias to be an explicit nil
func (o *DbTenant) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *DbTenant) UnsetAlias() {
	o.Alias.Unset()
}

// GetMappedDomain returns the MappedDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetMappedDomain() string {
	if o == nil || IsNil(o.MappedDomain.Get()) {
		var ret string
		return ret
	}
	return *o.MappedDomain.Get()
}

// GetMappedDomainOk returns a tuple with the MappedDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetMappedDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MappedDomain.Get(), o.MappedDomain.IsSet()
}

// HasMappedDomain returns a boolean if a field has been set.
func (o *DbTenant) IsMappedDomainSet() bool {
	if o != nil && o.MappedDomain.IsSet() {
		return true
	}

	return false
}

// SetMappedDomain gets a reference to the given NullableString and assigns it to the MappedDomain field.
func (o *DbTenant) SetMappedDomain(v string) {
	o.MappedDomain.Set(&v)
}
// SetMappedDomainNil sets the value for MappedDomain to be an explicit nil
func (o *DbTenant) SetMappedDomainNil() {
	o.MappedDomain.Set(nil)
}

// UnsetMappedDomain ensures that no value is present for MappedDomain, not even an explicit nil
func (o *DbTenant) UnsetMappedDomain() {
	o.MappedDomain.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DbTenant) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DbTenant) IsVersionSet() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DbTenant) SetVersion(v int32) {
	o.Version = &v
}

// GetVersionChangedField returns the VersionChangedField field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetVersionChangedField() time.Time {
	if o == nil || IsNil(o.VersionChangedField.Get()) {
		var ret time.Time
		return ret
	}
	return *o.VersionChangedField.Get()
}

// GetVersionChangedFieldOk returns a tuple with the VersionChangedField field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetVersionChangedFieldOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.VersionChangedField.Get(), o.VersionChangedField.IsSet()
}

// HasVersionChangedField returns a boolean if a field has been set.
func (o *DbTenant) IsVersionChangedFieldSet() bool {
	if o != nil && o.VersionChangedField.IsSet() {
		return true
	}

	return false
}

// SetVersionChangedField gets a reference to the given NullableTime and assigns it to the VersionChangedField field.
func (o *DbTenant) SetVersionChangedField(v time.Time) {
	o.VersionChangedField.Set(&v)
}
// SetVersionChangedFieldNil sets the value for VersionChangedField to be an explicit nil
func (o *DbTenant) SetVersionChangedFieldNil() {
	o.VersionChangedField.Set(nil)
}

// UnsetVersionChangedField ensures that no value is present for VersionChangedField, not even an explicit nil
func (o *DbTenant) UnsetVersionChangedField() {
	o.VersionChangedField.Unset()
}

// GetVersionChanged returns the VersionChanged field value if set, zero value otherwise.
func (o *DbTenant) GetVersionChanged() time.Time {
	if o == nil || IsNil(o.VersionChanged) {
		var ret time.Time
		return ret
	}
	return *o.VersionChanged
}

// GetVersionChangedOk returns a tuple with the VersionChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetVersionChangedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VersionChanged) {
		return nil, false
	}
	return o.VersionChanged, true
}

// HasVersionChanged returns a boolean if a field has been set.
func (o *DbTenant) IsVersionChangedSet() bool {
	if o != nil && !IsNil(o.VersionChanged) {
		return true
	}

	return false
}

// SetVersionChanged gets a reference to the given time.Time and assigns it to the VersionChanged field.
func (o *DbTenant) SetVersionChanged(v time.Time) {
	o.VersionChanged = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetLanguage() string {
	if o == nil || IsNil(o.Language.Get()) {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *DbTenant) IsLanguageSet() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *DbTenant) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *DbTenant) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *DbTenant) UnsetLanguage() {
	o.Language.Unset()
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone.Get()) {
		var ret string
		return ret
	}
	return *o.TimeZone.Get()
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeZone.Get(), o.TimeZone.IsSet()
}

// HasTimeZone returns a boolean if a field has been set.
func (o *DbTenant) IsTimeZoneSet() bool {
	if o != nil && o.TimeZone.IsSet() {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given NullableString and assigns it to the TimeZone field.
func (o *DbTenant) SetTimeZone(v string) {
	o.TimeZone.Set(&v)
}
// SetTimeZoneNil sets the value for TimeZone to be an explicit nil
func (o *DbTenant) SetTimeZoneNil() {
	o.TimeZone.Set(nil)
}

// UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
func (o *DbTenant) UnsetTimeZone() {
	o.TimeZone.Unset()
}

// GetTrustedDomainsRaw returns the TrustedDomainsRaw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetTrustedDomainsRaw() string {
	if o == nil || IsNil(o.TrustedDomainsRaw.Get()) {
		var ret string
		return ret
	}
	return *o.TrustedDomainsRaw.Get()
}

// GetTrustedDomainsRawOk returns a tuple with the TrustedDomainsRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetTrustedDomainsRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrustedDomainsRaw.Get(), o.TrustedDomainsRaw.IsSet()
}

// HasTrustedDomainsRaw returns a boolean if a field has been set.
func (o *DbTenant) IsTrustedDomainsRawSet() bool {
	if o != nil && o.TrustedDomainsRaw.IsSet() {
		return true
	}

	return false
}

// SetTrustedDomainsRaw gets a reference to the given NullableString and assigns it to the TrustedDomainsRaw field.
func (o *DbTenant) SetTrustedDomainsRaw(v string) {
	o.TrustedDomainsRaw.Set(&v)
}
// SetTrustedDomainsRawNil sets the value for TrustedDomainsRaw to be an explicit nil
func (o *DbTenant) SetTrustedDomainsRawNil() {
	o.TrustedDomainsRaw.Set(nil)
}

// UnsetTrustedDomainsRaw ensures that no value is present for TrustedDomainsRaw, not even an explicit nil
func (o *DbTenant) UnsetTrustedDomainsRaw() {
	o.TrustedDomainsRaw.Unset()
}

// GetTrustedDomainsEnabled returns the TrustedDomainsEnabled field value if set, zero value otherwise.
func (o *DbTenant) GetTrustedDomainsEnabled() TenantTrustedDomainsType {
	if o == nil || IsNil(o.TrustedDomainsEnabled) {
		var ret TenantTrustedDomainsType
		return ret
	}
	return *o.TrustedDomainsEnabled
}

// GetTrustedDomainsEnabledOk returns a tuple with the TrustedDomainsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetTrustedDomainsEnabledOk() (*TenantTrustedDomainsType, bool) {
	if o == nil || IsNil(o.TrustedDomainsEnabled) {
		return nil, false
	}
	return o.TrustedDomainsEnabled, true
}

// HasTrustedDomainsEnabled returns a boolean if a field has been set.
func (o *DbTenant) IsTrustedDomainsEnabledSet() bool {
	if o != nil && !IsNil(o.TrustedDomainsEnabled) {
		return true
	}

	return false
}

// SetTrustedDomainsEnabled gets a reference to the given TenantTrustedDomainsType and assigns it to the TrustedDomainsEnabled field.
func (o *DbTenant) SetTrustedDomainsEnabled(v TenantTrustedDomainsType) {
	o.TrustedDomainsEnabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DbTenant) GetStatus() TenantStatus {
	if o == nil || IsNil(o.Status) {
		var ret TenantStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetStatusOk() (*TenantStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DbTenant) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TenantStatus and assigns it to the Status field.
func (o *DbTenant) SetStatus(v TenantStatus) {
	o.Status = &v
}

// GetStatusChanged returns the StatusChanged field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetStatusChanged() time.Time {
	if o == nil || IsNil(o.StatusChanged.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StatusChanged.Get()
}

// GetStatusChangedOk returns a tuple with the StatusChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetStatusChangedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusChanged.Get(), o.StatusChanged.IsSet()
}

// HasStatusChanged returns a boolean if a field has been set.
func (o *DbTenant) IsStatusChangedSet() bool {
	if o != nil && o.StatusChanged.IsSet() {
		return true
	}

	return false
}

// SetStatusChanged gets a reference to the given NullableTime and assigns it to the StatusChanged field.
func (o *DbTenant) SetStatusChanged(v time.Time) {
	o.StatusChanged.Set(&v)
}
// SetStatusChangedNil sets the value for StatusChanged to be an explicit nil
func (o *DbTenant) SetStatusChangedNil() {
	o.StatusChanged.Set(nil)
}

// UnsetStatusChanged ensures that no value is present for StatusChanged, not even an explicit nil
func (o *DbTenant) UnsetStatusChanged() {
	o.StatusChanged.Unset()
}

// GetStatusChangedHack returns the StatusChangedHack field value if set, zero value otherwise.
func (o *DbTenant) GetStatusChangedHack() time.Time {
	if o == nil || IsNil(o.StatusChangedHack) {
		var ret time.Time
		return ret
	}
	return *o.StatusChangedHack
}

// GetStatusChangedHackOk returns a tuple with the StatusChangedHack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetStatusChangedHackOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StatusChangedHack) {
		return nil, false
	}
	return o.StatusChangedHack, true
}

// HasStatusChangedHack returns a boolean if a field has been set.
func (o *DbTenant) IsStatusChangedHackSet() bool {
	if o != nil && !IsNil(o.StatusChangedHack) {
		return true
	}

	return false
}

// SetStatusChangedHack gets a reference to the given time.Time and assigns it to the StatusChangedHack field.
func (o *DbTenant) SetStatusChangedHack(v time.Time) {
	o.StatusChangedHack = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *DbTenant) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *DbTenant) IsCreationDateTimeSet() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *DbTenant) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *DbTenant) IsOwnerIdSet() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *DbTenant) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *DbTenant) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *DbTenant) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbTenant) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbTenant) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// HasPaymentId returns a boolean if a field has been set.
func (o *DbTenant) IsPaymentIdSet() bool {
	if o != nil && o.PaymentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given NullableString and assigns it to the PaymentId field.
func (o *DbTenant) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}
// SetPaymentIdNil sets the value for PaymentId to be an explicit nil
func (o *DbTenant) SetPaymentIdNil() {
	o.PaymentId.Set(nil)
}

// UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
func (o *DbTenant) UnsetPaymentId() {
	o.PaymentId.Unset()
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
func (o *DbTenant) GetIndustry() TenantIndustry {
	if o == nil || IsNil(o.Industry) {
		var ret TenantIndustry
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetIndustryOk() (*TenantIndustry, bool) {
	if o == nil || IsNil(o.Industry) {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *DbTenant) IsIndustrySet() bool {
	if o != nil && !IsNil(o.Industry) {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given TenantIndustry and assigns it to the Industry field.
func (o *DbTenant) SetIndustry(v TenantIndustry) {
	o.Industry = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *DbTenant) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *DbTenant) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *DbTenant) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *DbTenant) GetCalls() bool {
	if o == nil || IsNil(o.Calls) {
		var ret bool
		return ret
	}
	return *o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetCallsOk() (*bool, bool) {
	if o == nil || IsNil(o.Calls) {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *DbTenant) IsCallsSet() bool {
	if o != nil && !IsNil(o.Calls) {
		return true
	}

	return false
}

// SetCalls gets a reference to the given bool and assigns it to the Calls field.
func (o *DbTenant) SetCalls(v bool) {
	o.Calls = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *DbTenant) GetPartner() DbTenantPartner {
	if o == nil || IsNil(o.Partner) {
		var ret DbTenantPartner
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTenant) GetPartnerOk() (*DbTenantPartner, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *DbTenant) IsPartnerSet() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given DbTenantPartner and assigns it to the Partner field.
func (o *DbTenant) SetPartner(v DbTenantPartner) {
	o.Partner = &v
}

func (o DbTenant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbTenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if o.MappedDomain.IsSet() {
		toSerialize["mappedDomain"] = o.MappedDomain.Get()
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if o.VersionChangedField.IsSet() {
		toSerialize["versionChangedField"] = o.VersionChangedField.Get()
	}
	if !IsNil(o.VersionChanged) {
		toSerialize["versionChanged"] = o.VersionChanged
	}
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if o.TimeZone.IsSet() {
		toSerialize["timeZone"] = o.TimeZone.Get()
	}
	if o.TrustedDomainsRaw.IsSet() {
		toSerialize["trustedDomainsRaw"] = o.TrustedDomainsRaw.Get()
	}
	if !IsNil(o.TrustedDomainsEnabled) {
		toSerialize["trustedDomainsEnabled"] = o.TrustedDomainsEnabled
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.StatusChanged.IsSet() {
		toSerialize["statusChanged"] = o.StatusChanged.Get()
	}
	if !IsNil(o.StatusChangedHack) {
		toSerialize["statusChangedHack"] = o.StatusChangedHack
	}
	if !IsNil(o.CreationDateTime) {
		toSerialize["creationDateTime"] = o.CreationDateTime
	}
	if o.OwnerId.IsSet() {
		toSerialize["ownerId"] = o.OwnerId.Get()
	}
	if o.PaymentId.IsSet() {
		toSerialize["paymentId"] = o.PaymentId.Get()
	}
	if !IsNil(o.Industry) {
		toSerialize["industry"] = o.Industry
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.Calls) {
		toSerialize["calls"] = o.Calls
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	return toSerialize, nil
}

type NullableDbTenant struct {
	value *DbTenant
	isSet bool
}

func (v NullableDbTenant) Get() *DbTenant {
	return v.value
}

func (v *NullableDbTenant) Set(val *DbTenant) {
	v.value = val
	v.isSet = true
}

func (v NullableDbTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableDbTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbTenant(val *DbTenant) *NullableDbTenant {
	return &NullableDbTenant{value: val, isSet: true}
}

func (v NullableDbTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

