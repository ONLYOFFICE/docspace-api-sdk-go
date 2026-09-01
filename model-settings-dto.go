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

// checks if the SettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsDto{}

// SettingsDto The settings information.
type SettingsDto struct {
	// The time zone.
	Timezone NullableString `json:"timezone,omitempty"`
	// The list of the trusted domains.
	TrustedDomains []string `json:"trustedDomains,omitempty"`
	// The type of the trusted domains.
	TrustedDomainsType *TenantTrustedDomainsType `json:"trustedDomainsType,omitempty"`
	// The language.
	Culture NullableString `json:"culture"`
	// The UTC offset in the TimeSpan format.
	UtcOffset *string `json:"utcOffset,omitempty"`
	// The UTC offset in hours.
	UtcHoursOffset *float64 `json:"utcHoursOffset,omitempty"`
	// The greeting settings.
	GreetingSettings NullableString `json:"greetingSettings,omitempty"`
	// The owner ID.
	OwnerId *string `json:"ownerId,omitempty"`
	// The team template ID.
	NameSchemaId NullableString `json:"nameSchemaId,omitempty"`
	// Specifies if a user can join the portal or not.
	EnabledJoin NullableBool `json:"enabledJoin,omitempty"`
	// Specifies if a user can send a message to the administrator when accessing the DocSpace portal or not.
	EnableAdmMess NullableBool `json:"enableAdmMess,omitempty"`
	// Specifies if a user can connect third-party providers to the portal or not.
	ThirdpartyEnable NullableBool `json:"thirdpartyEnable,omitempty"`
	// Specifies if this portal is a DocSpace portal or not.
	DocSpace *bool `json:"docSpace,omitempty"`
	// Indicates whether the system is running in standalone mode.
	Standalone *bool `json:"standalone,omitempty"`
	// Specifies if this portal is the AMI instance or not.
	IsAmi *bool `json:"isAmi,omitempty"`
	// The base domain.
	BaseDomain NullableString `json:"baseDomain"`
	// The wizard token.
	WizardToken NullableString `json:"wizardToken,omitempty"`
	// The password hash.
	PasswordHash *PasswordHasher `json:"passwordHash,omitempty"`
	// The Firebase parameters.
	Firebase *FirebaseDto `json:"firebase,omitempty"`
	// The portal version.
	Version NullableString `json:"version,omitempty"`
	// The type of CAPTCHA validation used.
	RecaptchaType *RecaptchaType `json:"recaptchaType,omitempty"`
	// The ReCAPTCHA public key.
	RecaptchaPublicKey NullableString `json:"recaptchaPublicKey,omitempty"`
	// Specifies if the debug information will be sent or not.
	DebugInfo *bool `json:"debugInfo,omitempty"`
	// The socket URL.
	SocketUrl NullableString `json:"socketUrl,omitempty"`
	// The tenant status.
	TenantStatus *TenantStatus `json:"tenantStatus,omitempty"`
	// The tenant alias.
	TenantAlias NullableString `json:"tenantAlias,omitempty"`
	// Specifies whether to display the About portal section.
	DisplayAbout *bool `json:"displayAbout,omitempty"`
	// The domain validator.
	DomainValidator *TenantDomainValidator `json:"domainValidator,omitempty"`
	// The Zendesk key.
	ZendeskKey NullableString `json:"zendeskKey,omitempty"`
	// The tag manager ID.
	TagManagerId NullableString `json:"tagManagerId,omitempty"`
	// Specifies whether the cookie settings are enabled.
	CookieSettingsEnabled bool `json:"cookieSettingsEnabled"`
	// Specifies whether the access to the space management is limited or not.
	LimitedAccessSpace *bool `json:"limitedAccessSpace,omitempty"`
	// Specifies whether the access to the Developer Tools is limited for users or not.
	LimitedAccessDevToolsForUsers *bool `json:"limitedAccessDevToolsForUsers,omitempty"`
	// Specifies whether to display the promotional banners.
	DisplayBanners *bool `json:"displayBanners,omitempty"`
	// Specifies whether AI functionality (chat, agents, vectorization) is enabled for the current tenant.  When `false`, all AI features are disabled and the AI Agents folder is hidden.
	AiEnabled *bool `json:"aiEnabled,omitempty"`
	// Specifies whether the tenant wallet balance is currently below the low-balance threshold. Only returned to portal administrators.
	WalletLowBalance NullableBool `json:"walletLowBalance,omitempty"`
	// The user name validation regex.
	UserNameRegex NullableString `json:"userNameRegex,omitempty"`
	// The maximum number of invitations to the portal.
	InvitationLimit NullableInt32 `json:"invitationLimit,omitempty"`
	// The plugins settings.
	Plugins *PluginsDto `json:"plugins,omitempty"`
	// The deep link settings.
	DeepLink DeepLinkDto `json:"deepLink"`
	// The form gallery settings.
	FormGallery *FormGalleryDto `json:"formGallery,omitempty"`
	// The maximum image upload size.
	MaxImageUploadSize *int64 `json:"maxImageUploadSize,omitempty"`
	// The white label logo text.
	LogoText NullableString `json:"logoText,omitempty"`
	// The external resources settings.
	ExternalResources *CultureSpecificExternalResources `json:"externalResources,omitempty"`
	// Specifies the default folder type for the current settings.
	DefaultFolderType *FolderType `json:"defaultFolderType,omitempty"`
	// Specifies if an external database is connected for storing form results.
	ExternalDbEnabled *bool `json:"externalDbEnabled,omitempty"`
}

type _SettingsDto SettingsDto

// NewSettingsDto instantiates a new SettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsDto(culture NullableString, baseDomain NullableString, cookieSettingsEnabled bool, deepLink DeepLinkDto) *SettingsDto {
	this := SettingsDto{}
	this.Culture = culture
	this.BaseDomain = baseDomain
	this.CookieSettingsEnabled = cookieSettingsEnabled
	this.DeepLink = deepLink
	return &this
}

// NewSettingsDtoWithDefaults instantiates a new SettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsDtoWithDefaults() *SettingsDto {
	this := SettingsDto{}
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetTimezone() string {
	if o == nil || IsNil(o.Timezone.Get()) {
		var ret string
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *SettingsDto) IsTimezoneSet() bool {
	if o != nil && o.Timezone.IsSet() {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NullableString and assigns it to the Timezone field.
func (o *SettingsDto) SetTimezone(v string) {
	o.Timezone.Set(&v)
}
// SetTimezoneNil sets the value for Timezone to be an explicit nil
func (o *SettingsDto) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
func (o *SettingsDto) UnsetTimezone() {
	o.Timezone.Unset()
}

// GetTrustedDomains returns the TrustedDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetTrustedDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TrustedDomains
}

// GetTrustedDomainsOk returns a tuple with the TrustedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetTrustedDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.TrustedDomains) {
		return nil, false
	}
	return o.TrustedDomains, true
}

// HasTrustedDomains returns a boolean if a field has been set.
func (o *SettingsDto) IsTrustedDomainsSet() bool {
	if o != nil && !IsNil(o.TrustedDomains) {
		return true
	}

	return false
}

// SetTrustedDomains gets a reference to the given []string and assigns it to the TrustedDomains field.
func (o *SettingsDto) SetTrustedDomains(v []string) {
	o.TrustedDomains = v
}

// GetTrustedDomainsType returns the TrustedDomainsType field value if set, zero value otherwise.
func (o *SettingsDto) GetTrustedDomainsType() TenantTrustedDomainsType {
	if o == nil || IsNil(o.TrustedDomainsType) {
		var ret TenantTrustedDomainsType
		return ret
	}
	return *o.TrustedDomainsType
}

// GetTrustedDomainsTypeOk returns a tuple with the TrustedDomainsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetTrustedDomainsTypeOk() (*TenantTrustedDomainsType, bool) {
	if o == nil || IsNil(o.TrustedDomainsType) {
		return nil, false
	}
	return o.TrustedDomainsType, true
}

// HasTrustedDomainsType returns a boolean if a field has been set.
func (o *SettingsDto) IsTrustedDomainsTypeSet() bool {
	if o != nil && !IsNil(o.TrustedDomainsType) {
		return true
	}

	return false
}

// SetTrustedDomainsType gets a reference to the given TenantTrustedDomainsType and assigns it to the TrustedDomainsType field.
func (o *SettingsDto) SetTrustedDomainsType(v TenantTrustedDomainsType) {
	o.TrustedDomainsType = &v
}

// GetCulture returns the Culture field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SettingsDto) GetCulture() string {
	if o == nil || o.Culture.Get() == nil {
		var ret string
		return ret
	}

	return *o.Culture.Get()
}

// GetCultureOk returns a tuple with the Culture field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetCultureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Culture.Get(), o.Culture.IsSet()
}

// SetCulture sets field value
func (o *SettingsDto) SetCulture(v string) {
	o.Culture.Set(&v)
}

// GetUtcOffset returns the UtcOffset field value if set, zero value otherwise.
func (o *SettingsDto) GetUtcOffset() string {
	if o == nil || IsNil(o.UtcOffset) {
		var ret string
		return ret
	}
	return *o.UtcOffset
}

// GetUtcOffsetOk returns a tuple with the UtcOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetUtcOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.UtcOffset) {
		return nil, false
	}
	return o.UtcOffset, true
}

// HasUtcOffset returns a boolean if a field has been set.
func (o *SettingsDto) IsUtcOffsetSet() bool {
	if o != nil && !IsNil(o.UtcOffset) {
		return true
	}

	return false
}

// SetUtcOffset gets a reference to the given string and assigns it to the UtcOffset field.
func (o *SettingsDto) SetUtcOffset(v string) {
	o.UtcOffset = &v
}

// GetUtcHoursOffset returns the UtcHoursOffset field value if set, zero value otherwise.
func (o *SettingsDto) GetUtcHoursOffset() float64 {
	if o == nil || IsNil(o.UtcHoursOffset) {
		var ret float64
		return ret
	}
	return *o.UtcHoursOffset
}

// GetUtcHoursOffsetOk returns a tuple with the UtcHoursOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetUtcHoursOffsetOk() (*float64, bool) {
	if o == nil || IsNil(o.UtcHoursOffset) {
		return nil, false
	}
	return o.UtcHoursOffset, true
}

// HasUtcHoursOffset returns a boolean if a field has been set.
func (o *SettingsDto) IsUtcHoursOffsetSet() bool {
	if o != nil && !IsNil(o.UtcHoursOffset) {
		return true
	}

	return false
}

// SetUtcHoursOffset gets a reference to the given float64 and assigns it to the UtcHoursOffset field.
func (o *SettingsDto) SetUtcHoursOffset(v float64) {
	o.UtcHoursOffset = &v
}

// GetGreetingSettings returns the GreetingSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetGreetingSettings() string {
	if o == nil || IsNil(o.GreetingSettings.Get()) {
		var ret string
		return ret
	}
	return *o.GreetingSettings.Get()
}

// GetGreetingSettingsOk returns a tuple with the GreetingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetGreetingSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GreetingSettings.Get(), o.GreetingSettings.IsSet()
}

// HasGreetingSettings returns a boolean if a field has been set.
func (o *SettingsDto) IsGreetingSettingsSet() bool {
	if o != nil && o.GreetingSettings.IsSet() {
		return true
	}

	return false
}

// SetGreetingSettings gets a reference to the given NullableString and assigns it to the GreetingSettings field.
func (o *SettingsDto) SetGreetingSettings(v string) {
	o.GreetingSettings.Set(&v)
}
// SetGreetingSettingsNil sets the value for GreetingSettings to be an explicit nil
func (o *SettingsDto) SetGreetingSettingsNil() {
	o.GreetingSettings.Set(nil)
}

// UnsetGreetingSettings ensures that no value is present for GreetingSettings, not even an explicit nil
func (o *SettingsDto) UnsetGreetingSettings() {
	o.GreetingSettings.Unset()
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *SettingsDto) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *SettingsDto) IsOwnerIdSet() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *SettingsDto) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetNameSchemaId returns the NameSchemaId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetNameSchemaId() string {
	if o == nil || IsNil(o.NameSchemaId.Get()) {
		var ret string
		return ret
	}
	return *o.NameSchemaId.Get()
}

// GetNameSchemaIdOk returns a tuple with the NameSchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetNameSchemaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameSchemaId.Get(), o.NameSchemaId.IsSet()
}

// HasNameSchemaId returns a boolean if a field has been set.
func (o *SettingsDto) IsNameSchemaIdSet() bool {
	if o != nil && o.NameSchemaId.IsSet() {
		return true
	}

	return false
}

// SetNameSchemaId gets a reference to the given NullableString and assigns it to the NameSchemaId field.
func (o *SettingsDto) SetNameSchemaId(v string) {
	o.NameSchemaId.Set(&v)
}
// SetNameSchemaIdNil sets the value for NameSchemaId to be an explicit nil
func (o *SettingsDto) SetNameSchemaIdNil() {
	o.NameSchemaId.Set(nil)
}

// UnsetNameSchemaId ensures that no value is present for NameSchemaId, not even an explicit nil
func (o *SettingsDto) UnsetNameSchemaId() {
	o.NameSchemaId.Unset()
}

// GetEnabledJoin returns the EnabledJoin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetEnabledJoin() bool {
	if o == nil || IsNil(o.EnabledJoin.Get()) {
		var ret bool
		return ret
	}
	return *o.EnabledJoin.Get()
}

// GetEnabledJoinOk returns a tuple with the EnabledJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetEnabledJoinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnabledJoin.Get(), o.EnabledJoin.IsSet()
}

// HasEnabledJoin returns a boolean if a field has been set.
func (o *SettingsDto) IsEnabledJoinSet() bool {
	if o != nil && o.EnabledJoin.IsSet() {
		return true
	}

	return false
}

// SetEnabledJoin gets a reference to the given NullableBool and assigns it to the EnabledJoin field.
func (o *SettingsDto) SetEnabledJoin(v bool) {
	o.EnabledJoin.Set(&v)
}
// SetEnabledJoinNil sets the value for EnabledJoin to be an explicit nil
func (o *SettingsDto) SetEnabledJoinNil() {
	o.EnabledJoin.Set(nil)
}

// UnsetEnabledJoin ensures that no value is present for EnabledJoin, not even an explicit nil
func (o *SettingsDto) UnsetEnabledJoin() {
	o.EnabledJoin.Unset()
}

// GetEnableAdmMess returns the EnableAdmMess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetEnableAdmMess() bool {
	if o == nil || IsNil(o.EnableAdmMess.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableAdmMess.Get()
}

// GetEnableAdmMessOk returns a tuple with the EnableAdmMess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetEnableAdmMessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableAdmMess.Get(), o.EnableAdmMess.IsSet()
}

// HasEnableAdmMess returns a boolean if a field has been set.
func (o *SettingsDto) IsEnableAdmMessSet() bool {
	if o != nil && o.EnableAdmMess.IsSet() {
		return true
	}

	return false
}

// SetEnableAdmMess gets a reference to the given NullableBool and assigns it to the EnableAdmMess field.
func (o *SettingsDto) SetEnableAdmMess(v bool) {
	o.EnableAdmMess.Set(&v)
}
// SetEnableAdmMessNil sets the value for EnableAdmMess to be an explicit nil
func (o *SettingsDto) SetEnableAdmMessNil() {
	o.EnableAdmMess.Set(nil)
}

// UnsetEnableAdmMess ensures that no value is present for EnableAdmMess, not even an explicit nil
func (o *SettingsDto) UnsetEnableAdmMess() {
	o.EnableAdmMess.Unset()
}

// GetThirdpartyEnable returns the ThirdpartyEnable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetThirdpartyEnable() bool {
	if o == nil || IsNil(o.ThirdpartyEnable.Get()) {
		var ret bool
		return ret
	}
	return *o.ThirdpartyEnable.Get()
}

// GetThirdpartyEnableOk returns a tuple with the ThirdpartyEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetThirdpartyEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdpartyEnable.Get(), o.ThirdpartyEnable.IsSet()
}

// HasThirdpartyEnable returns a boolean if a field has been set.
func (o *SettingsDto) IsThirdpartyEnableSet() bool {
	if o != nil && o.ThirdpartyEnable.IsSet() {
		return true
	}

	return false
}

// SetThirdpartyEnable gets a reference to the given NullableBool and assigns it to the ThirdpartyEnable field.
func (o *SettingsDto) SetThirdpartyEnable(v bool) {
	o.ThirdpartyEnable.Set(&v)
}
// SetThirdpartyEnableNil sets the value for ThirdpartyEnable to be an explicit nil
func (o *SettingsDto) SetThirdpartyEnableNil() {
	o.ThirdpartyEnable.Set(nil)
}

// UnsetThirdpartyEnable ensures that no value is present for ThirdpartyEnable, not even an explicit nil
func (o *SettingsDto) UnsetThirdpartyEnable() {
	o.ThirdpartyEnable.Unset()
}

// GetDocSpace returns the DocSpace field value if set, zero value otherwise.
func (o *SettingsDto) GetDocSpace() bool {
	if o == nil || IsNil(o.DocSpace) {
		var ret bool
		return ret
	}
	return *o.DocSpace
}

// GetDocSpaceOk returns a tuple with the DocSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetDocSpaceOk() (*bool, bool) {
	if o == nil || IsNil(o.DocSpace) {
		return nil, false
	}
	return o.DocSpace, true
}

// HasDocSpace returns a boolean if a field has been set.
func (o *SettingsDto) IsDocSpaceSet() bool {
	if o != nil && !IsNil(o.DocSpace) {
		return true
	}

	return false
}

// SetDocSpace gets a reference to the given bool and assigns it to the DocSpace field.
func (o *SettingsDto) SetDocSpace(v bool) {
	o.DocSpace = &v
}

// GetStandalone returns the Standalone field value if set, zero value otherwise.
func (o *SettingsDto) GetStandalone() bool {
	if o == nil || IsNil(o.Standalone) {
		var ret bool
		return ret
	}
	return *o.Standalone
}

// GetStandaloneOk returns a tuple with the Standalone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetStandaloneOk() (*bool, bool) {
	if o == nil || IsNil(o.Standalone) {
		return nil, false
	}
	return o.Standalone, true
}

// HasStandalone returns a boolean if a field has been set.
func (o *SettingsDto) IsStandaloneSet() bool {
	if o != nil && !IsNil(o.Standalone) {
		return true
	}

	return false
}

// SetStandalone gets a reference to the given bool and assigns it to the Standalone field.
func (o *SettingsDto) SetStandalone(v bool) {
	o.Standalone = &v
}

// GetIsAmi returns the IsAmi field value if set, zero value otherwise.
func (o *SettingsDto) GetIsAmi() bool {
	if o == nil || IsNil(o.IsAmi) {
		var ret bool
		return ret
	}
	return *o.IsAmi
}

// GetIsAmiOk returns a tuple with the IsAmi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetIsAmiOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmi) {
		return nil, false
	}
	return o.IsAmi, true
}

// HasIsAmi returns a boolean if a field has been set.
func (o *SettingsDto) IsIsAmiSet() bool {
	if o != nil && !IsNil(o.IsAmi) {
		return true
	}

	return false
}

// SetIsAmi gets a reference to the given bool and assigns it to the IsAmi field.
func (o *SettingsDto) SetIsAmi(v bool) {
	o.IsAmi = &v
}

// GetBaseDomain returns the BaseDomain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SettingsDto) GetBaseDomain() string {
	if o == nil || o.BaseDomain.Get() == nil {
		var ret string
		return ret
	}

	return *o.BaseDomain.Get()
}

// GetBaseDomainOk returns a tuple with the BaseDomain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetBaseDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseDomain.Get(), o.BaseDomain.IsSet()
}

// SetBaseDomain sets field value
func (o *SettingsDto) SetBaseDomain(v string) {
	o.BaseDomain.Set(&v)
}

// GetWizardToken returns the WizardToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetWizardToken() string {
	if o == nil || IsNil(o.WizardToken.Get()) {
		var ret string
		return ret
	}
	return *o.WizardToken.Get()
}

// GetWizardTokenOk returns a tuple with the WizardToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetWizardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WizardToken.Get(), o.WizardToken.IsSet()
}

// HasWizardToken returns a boolean if a field has been set.
func (o *SettingsDto) IsWizardTokenSet() bool {
	if o != nil && o.WizardToken.IsSet() {
		return true
	}

	return false
}

// SetWizardToken gets a reference to the given NullableString and assigns it to the WizardToken field.
func (o *SettingsDto) SetWizardToken(v string) {
	o.WizardToken.Set(&v)
}
// SetWizardTokenNil sets the value for WizardToken to be an explicit nil
func (o *SettingsDto) SetWizardTokenNil() {
	o.WizardToken.Set(nil)
}

// UnsetWizardToken ensures that no value is present for WizardToken, not even an explicit nil
func (o *SettingsDto) UnsetWizardToken() {
	o.WizardToken.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise.
func (o *SettingsDto) GetPasswordHash() PasswordHasher {
	if o == nil || IsNil(o.PasswordHash) {
		var ret PasswordHasher
		return ret
	}
	return *o.PasswordHash
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetPasswordHashOk() (*PasswordHasher, bool) {
	if o == nil || IsNil(o.PasswordHash) {
		return nil, false
	}
	return o.PasswordHash, true
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *SettingsDto) IsPasswordHashSet() bool {
	if o != nil && !IsNil(o.PasswordHash) {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given PasswordHasher and assigns it to the PasswordHash field.
func (o *SettingsDto) SetPasswordHash(v PasswordHasher) {
	o.PasswordHash = &v
}

// GetFirebase returns the Firebase field value if set, zero value otherwise.
func (o *SettingsDto) GetFirebase() FirebaseDto {
	if o == nil || IsNil(o.Firebase) {
		var ret FirebaseDto
		return ret
	}
	return *o.Firebase
}

// GetFirebaseOk returns a tuple with the Firebase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetFirebaseOk() (*FirebaseDto, bool) {
	if o == nil || IsNil(o.Firebase) {
		return nil, false
	}
	return o.Firebase, true
}

// HasFirebase returns a boolean if a field has been set.
func (o *SettingsDto) IsFirebaseSet() bool {
	if o != nil && !IsNil(o.Firebase) {
		return true
	}

	return false
}

// SetFirebase gets a reference to the given FirebaseDto and assigns it to the Firebase field.
func (o *SettingsDto) SetFirebase(v FirebaseDto) {
	o.Firebase = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *SettingsDto) IsVersionSet() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *SettingsDto) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *SettingsDto) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *SettingsDto) UnsetVersion() {
	o.Version.Unset()
}

// GetRecaptchaType returns the RecaptchaType field value if set, zero value otherwise.
func (o *SettingsDto) GetRecaptchaType() RecaptchaType {
	if o == nil || IsNil(o.RecaptchaType) {
		var ret RecaptchaType
		return ret
	}
	return *o.RecaptchaType
}

// GetRecaptchaTypeOk returns a tuple with the RecaptchaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetRecaptchaTypeOk() (*RecaptchaType, bool) {
	if o == nil || IsNil(o.RecaptchaType) {
		return nil, false
	}
	return o.RecaptchaType, true
}

// HasRecaptchaType returns a boolean if a field has been set.
func (o *SettingsDto) IsRecaptchaTypeSet() bool {
	if o != nil && !IsNil(o.RecaptchaType) {
		return true
	}

	return false
}

// SetRecaptchaType gets a reference to the given RecaptchaType and assigns it to the RecaptchaType field.
func (o *SettingsDto) SetRecaptchaType(v RecaptchaType) {
	o.RecaptchaType = &v
}

// GetRecaptchaPublicKey returns the RecaptchaPublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetRecaptchaPublicKey() string {
	if o == nil || IsNil(o.RecaptchaPublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.RecaptchaPublicKey.Get()
}

// GetRecaptchaPublicKeyOk returns a tuple with the RecaptchaPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetRecaptchaPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecaptchaPublicKey.Get(), o.RecaptchaPublicKey.IsSet()
}

// HasRecaptchaPublicKey returns a boolean if a field has been set.
func (o *SettingsDto) IsRecaptchaPublicKeySet() bool {
	if o != nil && o.RecaptchaPublicKey.IsSet() {
		return true
	}

	return false
}

// SetRecaptchaPublicKey gets a reference to the given NullableString and assigns it to the RecaptchaPublicKey field.
func (o *SettingsDto) SetRecaptchaPublicKey(v string) {
	o.RecaptchaPublicKey.Set(&v)
}
// SetRecaptchaPublicKeyNil sets the value for RecaptchaPublicKey to be an explicit nil
func (o *SettingsDto) SetRecaptchaPublicKeyNil() {
	o.RecaptchaPublicKey.Set(nil)
}

// UnsetRecaptchaPublicKey ensures that no value is present for RecaptchaPublicKey, not even an explicit nil
func (o *SettingsDto) UnsetRecaptchaPublicKey() {
	o.RecaptchaPublicKey.Unset()
}

// GetDebugInfo returns the DebugInfo field value if set, zero value otherwise.
func (o *SettingsDto) GetDebugInfo() bool {
	if o == nil || IsNil(o.DebugInfo) {
		var ret bool
		return ret
	}
	return *o.DebugInfo
}

// GetDebugInfoOk returns a tuple with the DebugInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetDebugInfoOk() (*bool, bool) {
	if o == nil || IsNil(o.DebugInfo) {
		return nil, false
	}
	return o.DebugInfo, true
}

// HasDebugInfo returns a boolean if a field has been set.
func (o *SettingsDto) IsDebugInfoSet() bool {
	if o != nil && !IsNil(o.DebugInfo) {
		return true
	}

	return false
}

// SetDebugInfo gets a reference to the given bool and assigns it to the DebugInfo field.
func (o *SettingsDto) SetDebugInfo(v bool) {
	o.DebugInfo = &v
}

// GetSocketUrl returns the SocketUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetSocketUrl() string {
	if o == nil || IsNil(o.SocketUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SocketUrl.Get()
}

// GetSocketUrlOk returns a tuple with the SocketUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetSocketUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SocketUrl.Get(), o.SocketUrl.IsSet()
}

// HasSocketUrl returns a boolean if a field has been set.
func (o *SettingsDto) IsSocketUrlSet() bool {
	if o != nil && o.SocketUrl.IsSet() {
		return true
	}

	return false
}

// SetSocketUrl gets a reference to the given NullableString and assigns it to the SocketUrl field.
func (o *SettingsDto) SetSocketUrl(v string) {
	o.SocketUrl.Set(&v)
}
// SetSocketUrlNil sets the value for SocketUrl to be an explicit nil
func (o *SettingsDto) SetSocketUrlNil() {
	o.SocketUrl.Set(nil)
}

// UnsetSocketUrl ensures that no value is present for SocketUrl, not even an explicit nil
func (o *SettingsDto) UnsetSocketUrl() {
	o.SocketUrl.Unset()
}

// GetTenantStatus returns the TenantStatus field value if set, zero value otherwise.
func (o *SettingsDto) GetTenantStatus() TenantStatus {
	if o == nil || IsNil(o.TenantStatus) {
		var ret TenantStatus
		return ret
	}
	return *o.TenantStatus
}

// GetTenantStatusOk returns a tuple with the TenantStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetTenantStatusOk() (*TenantStatus, bool) {
	if o == nil || IsNil(o.TenantStatus) {
		return nil, false
	}
	return o.TenantStatus, true
}

// HasTenantStatus returns a boolean if a field has been set.
func (o *SettingsDto) IsTenantStatusSet() bool {
	if o != nil && !IsNil(o.TenantStatus) {
		return true
	}

	return false
}

// SetTenantStatus gets a reference to the given TenantStatus and assigns it to the TenantStatus field.
func (o *SettingsDto) SetTenantStatus(v TenantStatus) {
	o.TenantStatus = &v
}

// GetTenantAlias returns the TenantAlias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetTenantAlias() string {
	if o == nil || IsNil(o.TenantAlias.Get()) {
		var ret string
		return ret
	}
	return *o.TenantAlias.Get()
}

// GetTenantAliasOk returns a tuple with the TenantAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetTenantAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantAlias.Get(), o.TenantAlias.IsSet()
}

// HasTenantAlias returns a boolean if a field has been set.
func (o *SettingsDto) IsTenantAliasSet() bool {
	if o != nil && o.TenantAlias.IsSet() {
		return true
	}

	return false
}

// SetTenantAlias gets a reference to the given NullableString and assigns it to the TenantAlias field.
func (o *SettingsDto) SetTenantAlias(v string) {
	o.TenantAlias.Set(&v)
}
// SetTenantAliasNil sets the value for TenantAlias to be an explicit nil
func (o *SettingsDto) SetTenantAliasNil() {
	o.TenantAlias.Set(nil)
}

// UnsetTenantAlias ensures that no value is present for TenantAlias, not even an explicit nil
func (o *SettingsDto) UnsetTenantAlias() {
	o.TenantAlias.Unset()
}

// GetDisplayAbout returns the DisplayAbout field value if set, zero value otherwise.
func (o *SettingsDto) GetDisplayAbout() bool {
	if o == nil || IsNil(o.DisplayAbout) {
		var ret bool
		return ret
	}
	return *o.DisplayAbout
}

// GetDisplayAboutOk returns a tuple with the DisplayAbout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetDisplayAboutOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayAbout) {
		return nil, false
	}
	return o.DisplayAbout, true
}

// HasDisplayAbout returns a boolean if a field has been set.
func (o *SettingsDto) IsDisplayAboutSet() bool {
	if o != nil && !IsNil(o.DisplayAbout) {
		return true
	}

	return false
}

// SetDisplayAbout gets a reference to the given bool and assigns it to the DisplayAbout field.
func (o *SettingsDto) SetDisplayAbout(v bool) {
	o.DisplayAbout = &v
}

// GetDomainValidator returns the DomainValidator field value if set, zero value otherwise.
func (o *SettingsDto) GetDomainValidator() TenantDomainValidator {
	if o == nil || IsNil(o.DomainValidator) {
		var ret TenantDomainValidator
		return ret
	}
	return *o.DomainValidator
}

// GetDomainValidatorOk returns a tuple with the DomainValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetDomainValidatorOk() (*TenantDomainValidator, bool) {
	if o == nil || IsNil(o.DomainValidator) {
		return nil, false
	}
	return o.DomainValidator, true
}

// HasDomainValidator returns a boolean if a field has been set.
func (o *SettingsDto) IsDomainValidatorSet() bool {
	if o != nil && !IsNil(o.DomainValidator) {
		return true
	}

	return false
}

// SetDomainValidator gets a reference to the given TenantDomainValidator and assigns it to the DomainValidator field.
func (o *SettingsDto) SetDomainValidator(v TenantDomainValidator) {
	o.DomainValidator = &v
}

// GetZendeskKey returns the ZendeskKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetZendeskKey() string {
	if o == nil || IsNil(o.ZendeskKey.Get()) {
		var ret string
		return ret
	}
	return *o.ZendeskKey.Get()
}

// GetZendeskKeyOk returns a tuple with the ZendeskKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetZendeskKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZendeskKey.Get(), o.ZendeskKey.IsSet()
}

// HasZendeskKey returns a boolean if a field has been set.
func (o *SettingsDto) IsZendeskKeySet() bool {
	if o != nil && o.ZendeskKey.IsSet() {
		return true
	}

	return false
}

// SetZendeskKey gets a reference to the given NullableString and assigns it to the ZendeskKey field.
func (o *SettingsDto) SetZendeskKey(v string) {
	o.ZendeskKey.Set(&v)
}
// SetZendeskKeyNil sets the value for ZendeskKey to be an explicit nil
func (o *SettingsDto) SetZendeskKeyNil() {
	o.ZendeskKey.Set(nil)
}

// UnsetZendeskKey ensures that no value is present for ZendeskKey, not even an explicit nil
func (o *SettingsDto) UnsetZendeskKey() {
	o.ZendeskKey.Unset()
}

// GetTagManagerId returns the TagManagerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetTagManagerId() string {
	if o == nil || IsNil(o.TagManagerId.Get()) {
		var ret string
		return ret
	}
	return *o.TagManagerId.Get()
}

// GetTagManagerIdOk returns a tuple with the TagManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetTagManagerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagManagerId.Get(), o.TagManagerId.IsSet()
}

// HasTagManagerId returns a boolean if a field has been set.
func (o *SettingsDto) IsTagManagerIdSet() bool {
	if o != nil && o.TagManagerId.IsSet() {
		return true
	}

	return false
}

// SetTagManagerId gets a reference to the given NullableString and assigns it to the TagManagerId field.
func (o *SettingsDto) SetTagManagerId(v string) {
	o.TagManagerId.Set(&v)
}
// SetTagManagerIdNil sets the value for TagManagerId to be an explicit nil
func (o *SettingsDto) SetTagManagerIdNil() {
	o.TagManagerId.Set(nil)
}

// UnsetTagManagerId ensures that no value is present for TagManagerId, not even an explicit nil
func (o *SettingsDto) UnsetTagManagerId() {
	o.TagManagerId.Unset()
}

// GetCookieSettingsEnabled returns the CookieSettingsEnabled field value
func (o *SettingsDto) GetCookieSettingsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CookieSettingsEnabled
}

// GetCookieSettingsEnabledOk returns a tuple with the CookieSettingsEnabled field value
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetCookieSettingsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CookieSettingsEnabled, true
}

// SetCookieSettingsEnabled sets field value
func (o *SettingsDto) SetCookieSettingsEnabled(v bool) {
	o.CookieSettingsEnabled = v
}

// GetLimitedAccessSpace returns the LimitedAccessSpace field value if set, zero value otherwise.
func (o *SettingsDto) GetLimitedAccessSpace() bool {
	if o == nil || IsNil(o.LimitedAccessSpace) {
		var ret bool
		return ret
	}
	return *o.LimitedAccessSpace
}

// GetLimitedAccessSpaceOk returns a tuple with the LimitedAccessSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetLimitedAccessSpaceOk() (*bool, bool) {
	if o == nil || IsNil(o.LimitedAccessSpace) {
		return nil, false
	}
	return o.LimitedAccessSpace, true
}

// HasLimitedAccessSpace returns a boolean if a field has been set.
func (o *SettingsDto) IsLimitedAccessSpaceSet() bool {
	if o != nil && !IsNil(o.LimitedAccessSpace) {
		return true
	}

	return false
}

// SetLimitedAccessSpace gets a reference to the given bool and assigns it to the LimitedAccessSpace field.
func (o *SettingsDto) SetLimitedAccessSpace(v bool) {
	o.LimitedAccessSpace = &v
}

// GetLimitedAccessDevToolsForUsers returns the LimitedAccessDevToolsForUsers field value if set, zero value otherwise.
func (o *SettingsDto) GetLimitedAccessDevToolsForUsers() bool {
	if o == nil || IsNil(o.LimitedAccessDevToolsForUsers) {
		var ret bool
		return ret
	}
	return *o.LimitedAccessDevToolsForUsers
}

// GetLimitedAccessDevToolsForUsersOk returns a tuple with the LimitedAccessDevToolsForUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetLimitedAccessDevToolsForUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.LimitedAccessDevToolsForUsers) {
		return nil, false
	}
	return o.LimitedAccessDevToolsForUsers, true
}

// HasLimitedAccessDevToolsForUsers returns a boolean if a field has been set.
func (o *SettingsDto) IsLimitedAccessDevToolsForUsersSet() bool {
	if o != nil && !IsNil(o.LimitedAccessDevToolsForUsers) {
		return true
	}

	return false
}

// SetLimitedAccessDevToolsForUsers gets a reference to the given bool and assigns it to the LimitedAccessDevToolsForUsers field.
func (o *SettingsDto) SetLimitedAccessDevToolsForUsers(v bool) {
	o.LimitedAccessDevToolsForUsers = &v
}

// GetDisplayBanners returns the DisplayBanners field value if set, zero value otherwise.
func (o *SettingsDto) GetDisplayBanners() bool {
	if o == nil || IsNil(o.DisplayBanners) {
		var ret bool
		return ret
	}
	return *o.DisplayBanners
}

// GetDisplayBannersOk returns a tuple with the DisplayBanners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetDisplayBannersOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayBanners) {
		return nil, false
	}
	return o.DisplayBanners, true
}

// HasDisplayBanners returns a boolean if a field has been set.
func (o *SettingsDto) IsDisplayBannersSet() bool {
	if o != nil && !IsNil(o.DisplayBanners) {
		return true
	}

	return false
}

// SetDisplayBanners gets a reference to the given bool and assigns it to the DisplayBanners field.
func (o *SettingsDto) SetDisplayBanners(v bool) {
	o.DisplayBanners = &v
}

// GetAiEnabled returns the AiEnabled field value if set, zero value otherwise.
func (o *SettingsDto) GetAiEnabled() bool {
	if o == nil || IsNil(o.AiEnabled) {
		var ret bool
		return ret
	}
	return *o.AiEnabled
}

// GetAiEnabledOk returns a tuple with the AiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetAiEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AiEnabled) {
		return nil, false
	}
	return o.AiEnabled, true
}

// HasAiEnabled returns a boolean if a field has been set.
func (o *SettingsDto) IsAiEnabledSet() bool {
	if o != nil && !IsNil(o.AiEnabled) {
		return true
	}

	return false
}

// SetAiEnabled gets a reference to the given bool and assigns it to the AiEnabled field.
func (o *SettingsDto) SetAiEnabled(v bool) {
	o.AiEnabled = &v
}

// GetWalletLowBalance returns the WalletLowBalance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetWalletLowBalance() bool {
	if o == nil || IsNil(o.WalletLowBalance.Get()) {
		var ret bool
		return ret
	}
	return *o.WalletLowBalance.Get()
}

// GetWalletLowBalanceOk returns a tuple with the WalletLowBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetWalletLowBalanceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalletLowBalance.Get(), o.WalletLowBalance.IsSet()
}

// HasWalletLowBalance returns a boolean if a field has been set.
func (o *SettingsDto) IsWalletLowBalanceSet() bool {
	if o != nil && o.WalletLowBalance.IsSet() {
		return true
	}

	return false
}

// SetWalletLowBalance gets a reference to the given NullableBool and assigns it to the WalletLowBalance field.
func (o *SettingsDto) SetWalletLowBalance(v bool) {
	o.WalletLowBalance.Set(&v)
}
// SetWalletLowBalanceNil sets the value for WalletLowBalance to be an explicit nil
func (o *SettingsDto) SetWalletLowBalanceNil() {
	o.WalletLowBalance.Set(nil)
}

// UnsetWalletLowBalance ensures that no value is present for WalletLowBalance, not even an explicit nil
func (o *SettingsDto) UnsetWalletLowBalance() {
	o.WalletLowBalance.Unset()
}

// GetUserNameRegex returns the UserNameRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetUserNameRegex() string {
	if o == nil || IsNil(o.UserNameRegex.Get()) {
		var ret string
		return ret
	}
	return *o.UserNameRegex.Get()
}

// GetUserNameRegexOk returns a tuple with the UserNameRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetUserNameRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserNameRegex.Get(), o.UserNameRegex.IsSet()
}

// HasUserNameRegex returns a boolean if a field has been set.
func (o *SettingsDto) IsUserNameRegexSet() bool {
	if o != nil && o.UserNameRegex.IsSet() {
		return true
	}

	return false
}

// SetUserNameRegex gets a reference to the given NullableString and assigns it to the UserNameRegex field.
func (o *SettingsDto) SetUserNameRegex(v string) {
	o.UserNameRegex.Set(&v)
}
// SetUserNameRegexNil sets the value for UserNameRegex to be an explicit nil
func (o *SettingsDto) SetUserNameRegexNil() {
	o.UserNameRegex.Set(nil)
}

// UnsetUserNameRegex ensures that no value is present for UserNameRegex, not even an explicit nil
func (o *SettingsDto) UnsetUserNameRegex() {
	o.UserNameRegex.Unset()
}

// GetInvitationLimit returns the InvitationLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetInvitationLimit() int32 {
	if o == nil || IsNil(o.InvitationLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.InvitationLimit.Get()
}

// GetInvitationLimitOk returns a tuple with the InvitationLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetInvitationLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvitationLimit.Get(), o.InvitationLimit.IsSet()
}

// HasInvitationLimit returns a boolean if a field has been set.
func (o *SettingsDto) IsInvitationLimitSet() bool {
	if o != nil && o.InvitationLimit.IsSet() {
		return true
	}

	return false
}

// SetInvitationLimit gets a reference to the given NullableInt32 and assigns it to the InvitationLimit field.
func (o *SettingsDto) SetInvitationLimit(v int32) {
	o.InvitationLimit.Set(&v)
}
// SetInvitationLimitNil sets the value for InvitationLimit to be an explicit nil
func (o *SettingsDto) SetInvitationLimitNil() {
	o.InvitationLimit.Set(nil)
}

// UnsetInvitationLimit ensures that no value is present for InvitationLimit, not even an explicit nil
func (o *SettingsDto) UnsetInvitationLimit() {
	o.InvitationLimit.Unset()
}

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *SettingsDto) GetPlugins() PluginsDto {
	if o == nil || IsNil(o.Plugins) {
		var ret PluginsDto
		return ret
	}
	return *o.Plugins
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetPluginsOk() (*PluginsDto, bool) {
	if o == nil || IsNil(o.Plugins) {
		return nil, false
	}
	return o.Plugins, true
}

// HasPlugins returns a boolean if a field has been set.
func (o *SettingsDto) IsPluginsSet() bool {
	if o != nil && !IsNil(o.Plugins) {
		return true
	}

	return false
}

// SetPlugins gets a reference to the given PluginsDto and assigns it to the Plugins field.
func (o *SettingsDto) SetPlugins(v PluginsDto) {
	o.Plugins = &v
}

// GetDeepLink returns the DeepLink field value
func (o *SettingsDto) GetDeepLink() DeepLinkDto {
	if o == nil {
		var ret DeepLinkDto
		return ret
	}

	return o.DeepLink
}

// GetDeepLinkOk returns a tuple with the DeepLink field value
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetDeepLinkOk() (*DeepLinkDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeepLink, true
}

// SetDeepLink sets field value
func (o *SettingsDto) SetDeepLink(v DeepLinkDto) {
	o.DeepLink = v
}

// GetFormGallery returns the FormGallery field value if set, zero value otherwise.
func (o *SettingsDto) GetFormGallery() FormGalleryDto {
	if o == nil || IsNil(o.FormGallery) {
		var ret FormGalleryDto
		return ret
	}
	return *o.FormGallery
}

// GetFormGalleryOk returns a tuple with the FormGallery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetFormGalleryOk() (*FormGalleryDto, bool) {
	if o == nil || IsNil(o.FormGallery) {
		return nil, false
	}
	return o.FormGallery, true
}

// HasFormGallery returns a boolean if a field has been set.
func (o *SettingsDto) IsFormGallerySet() bool {
	if o != nil && !IsNil(o.FormGallery) {
		return true
	}

	return false
}

// SetFormGallery gets a reference to the given FormGalleryDto and assigns it to the FormGallery field.
func (o *SettingsDto) SetFormGallery(v FormGalleryDto) {
	o.FormGallery = &v
}

// GetMaxImageUploadSize returns the MaxImageUploadSize field value if set, zero value otherwise.
func (o *SettingsDto) GetMaxImageUploadSize() int64 {
	if o == nil || IsNil(o.MaxImageUploadSize) {
		var ret int64
		return ret
	}
	return *o.MaxImageUploadSize
}

// GetMaxImageUploadSizeOk returns a tuple with the MaxImageUploadSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetMaxImageUploadSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxImageUploadSize) {
		return nil, false
	}
	return o.MaxImageUploadSize, true
}

// HasMaxImageUploadSize returns a boolean if a field has been set.
func (o *SettingsDto) IsMaxImageUploadSizeSet() bool {
	if o != nil && !IsNil(o.MaxImageUploadSize) {
		return true
	}

	return false
}

// SetMaxImageUploadSize gets a reference to the given int64 and assigns it to the MaxImageUploadSize field.
func (o *SettingsDto) SetMaxImageUploadSize(v int64) {
	o.MaxImageUploadSize = &v
}

// GetLogoText returns the LogoText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsDto) GetLogoText() string {
	if o == nil || IsNil(o.LogoText.Get()) {
		var ret string
		return ret
	}
	return *o.LogoText.Get()
}

// GetLogoTextOk returns a tuple with the LogoText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsDto) GetLogoTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoText.Get(), o.LogoText.IsSet()
}

// HasLogoText returns a boolean if a field has been set.
func (o *SettingsDto) IsLogoTextSet() bool {
	if o != nil && o.LogoText.IsSet() {
		return true
	}

	return false
}

// SetLogoText gets a reference to the given NullableString and assigns it to the LogoText field.
func (o *SettingsDto) SetLogoText(v string) {
	o.LogoText.Set(&v)
}
// SetLogoTextNil sets the value for LogoText to be an explicit nil
func (o *SettingsDto) SetLogoTextNil() {
	o.LogoText.Set(nil)
}

// UnsetLogoText ensures that no value is present for LogoText, not even an explicit nil
func (o *SettingsDto) UnsetLogoText() {
	o.LogoText.Unset()
}

// GetExternalResources returns the ExternalResources field value if set, zero value otherwise.
func (o *SettingsDto) GetExternalResources() CultureSpecificExternalResources {
	if o == nil || IsNil(o.ExternalResources) {
		var ret CultureSpecificExternalResources
		return ret
	}
	return *o.ExternalResources
}

// GetExternalResourcesOk returns a tuple with the ExternalResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetExternalResourcesOk() (*CultureSpecificExternalResources, bool) {
	if o == nil || IsNil(o.ExternalResources) {
		return nil, false
	}
	return o.ExternalResources, true
}

// HasExternalResources returns a boolean if a field has been set.
func (o *SettingsDto) IsExternalResourcesSet() bool {
	if o != nil && !IsNil(o.ExternalResources) {
		return true
	}

	return false
}

// SetExternalResources gets a reference to the given CultureSpecificExternalResources and assigns it to the ExternalResources field.
func (o *SettingsDto) SetExternalResources(v CultureSpecificExternalResources) {
	o.ExternalResources = &v
}

// GetDefaultFolderType returns the DefaultFolderType field value if set, zero value otherwise.
func (o *SettingsDto) GetDefaultFolderType() FolderType {
	if o == nil || IsNil(o.DefaultFolderType) {
		var ret FolderType
		return ret
	}
	return *o.DefaultFolderType
}

// GetDefaultFolderTypeOk returns a tuple with the DefaultFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetDefaultFolderTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.DefaultFolderType) {
		return nil, false
	}
	return o.DefaultFolderType, true
}

// HasDefaultFolderType returns a boolean if a field has been set.
func (o *SettingsDto) IsDefaultFolderTypeSet() bool {
	if o != nil && !IsNil(o.DefaultFolderType) {
		return true
	}

	return false
}

// SetDefaultFolderType gets a reference to the given FolderType and assigns it to the DefaultFolderType field.
func (o *SettingsDto) SetDefaultFolderType(v FolderType) {
	o.DefaultFolderType = &v
}

// GetExternalDbEnabled returns the ExternalDbEnabled field value if set, zero value otherwise.
func (o *SettingsDto) GetExternalDbEnabled() bool {
	if o == nil || IsNil(o.ExternalDbEnabled) {
		var ret bool
		return ret
	}
	return *o.ExternalDbEnabled
}

// GetExternalDbEnabledOk returns a tuple with the ExternalDbEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDto) GetExternalDbEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalDbEnabled) {
		return nil, false
	}
	return o.ExternalDbEnabled, true
}

// HasExternalDbEnabled returns a boolean if a field has been set.
func (o *SettingsDto) IsExternalDbEnabledSet() bool {
	if o != nil && !IsNil(o.ExternalDbEnabled) {
		return true
	}

	return false
}

// SetExternalDbEnabled gets a reference to the given bool and assigns it to the ExternalDbEnabled field.
func (o *SettingsDto) SetExternalDbEnabled(v bool) {
	o.ExternalDbEnabled = &v
}

func (o SettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Timezone.IsSet() {
		toSerialize["timezone"] = o.Timezone.Get()
	}
	if o.TrustedDomains != nil {
		toSerialize["trustedDomains"] = o.TrustedDomains
	}
	if !IsNil(o.TrustedDomainsType) {
		toSerialize["trustedDomainsType"] = o.TrustedDomainsType
	}
	toSerialize["culture"] = o.Culture.Get()
	if !IsNil(o.UtcOffset) {
		toSerialize["utcOffset"] = o.UtcOffset
	}
	if !IsNil(o.UtcHoursOffset) {
		toSerialize["utcHoursOffset"] = o.UtcHoursOffset
	}
	if o.GreetingSettings.IsSet() {
		toSerialize["greetingSettings"] = o.GreetingSettings.Get()
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.NameSchemaId.IsSet() {
		toSerialize["nameSchemaId"] = o.NameSchemaId.Get()
	}
	if o.EnabledJoin.IsSet() {
		toSerialize["enabledJoin"] = o.EnabledJoin.Get()
	}
	if o.EnableAdmMess.IsSet() {
		toSerialize["enableAdmMess"] = o.EnableAdmMess.Get()
	}
	if o.ThirdpartyEnable.IsSet() {
		toSerialize["thirdpartyEnable"] = o.ThirdpartyEnable.Get()
	}
	if !IsNil(o.DocSpace) {
		toSerialize["docSpace"] = o.DocSpace
	}
	if !IsNil(o.Standalone) {
		toSerialize["standalone"] = o.Standalone
	}
	if !IsNil(o.IsAmi) {
		toSerialize["isAmi"] = o.IsAmi
	}
	toSerialize["baseDomain"] = o.BaseDomain.Get()
	if o.WizardToken.IsSet() {
		toSerialize["wizardToken"] = o.WizardToken.Get()
	}
	if !IsNil(o.PasswordHash) {
		toSerialize["passwordHash"] = o.PasswordHash
	}
	if !IsNil(o.Firebase) {
		toSerialize["firebase"] = o.Firebase
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if !IsNil(o.RecaptchaType) {
		toSerialize["recaptchaType"] = o.RecaptchaType
	}
	if o.RecaptchaPublicKey.IsSet() {
		toSerialize["recaptchaPublicKey"] = o.RecaptchaPublicKey.Get()
	}
	if !IsNil(o.DebugInfo) {
		toSerialize["debugInfo"] = o.DebugInfo
	}
	if o.SocketUrl.IsSet() {
		toSerialize["socketUrl"] = o.SocketUrl.Get()
	}
	if !IsNil(o.TenantStatus) {
		toSerialize["tenantStatus"] = o.TenantStatus
	}
	if o.TenantAlias.IsSet() {
		toSerialize["tenantAlias"] = o.TenantAlias.Get()
	}
	if !IsNil(o.DisplayAbout) {
		toSerialize["displayAbout"] = o.DisplayAbout
	}
	if !IsNil(o.DomainValidator) {
		toSerialize["domainValidator"] = o.DomainValidator
	}
	if o.ZendeskKey.IsSet() {
		toSerialize["zendeskKey"] = o.ZendeskKey.Get()
	}
	if o.TagManagerId.IsSet() {
		toSerialize["tagManagerId"] = o.TagManagerId.Get()
	}
	toSerialize["cookieSettingsEnabled"] = o.CookieSettingsEnabled
	if !IsNil(o.LimitedAccessSpace) {
		toSerialize["limitedAccessSpace"] = o.LimitedAccessSpace
	}
	if !IsNil(o.LimitedAccessDevToolsForUsers) {
		toSerialize["limitedAccessDevToolsForUsers"] = o.LimitedAccessDevToolsForUsers
	}
	if !IsNil(o.DisplayBanners) {
		toSerialize["displayBanners"] = o.DisplayBanners
	}
	if !IsNil(o.AiEnabled) {
		toSerialize["aiEnabled"] = o.AiEnabled
	}
	if o.WalletLowBalance.IsSet() {
		toSerialize["walletLowBalance"] = o.WalletLowBalance.Get()
	}
	if o.UserNameRegex.IsSet() {
		toSerialize["userNameRegex"] = o.UserNameRegex.Get()
	}
	if o.InvitationLimit.IsSet() {
		toSerialize["invitationLimit"] = o.InvitationLimit.Get()
	}
	if !IsNil(o.Plugins) {
		toSerialize["plugins"] = o.Plugins
	}
	toSerialize["deepLink"] = o.DeepLink
	if !IsNil(o.FormGallery) {
		toSerialize["formGallery"] = o.FormGallery
	}
	if !IsNil(o.MaxImageUploadSize) {
		toSerialize["maxImageUploadSize"] = o.MaxImageUploadSize
	}
	if o.LogoText.IsSet() {
		toSerialize["logoText"] = o.LogoText.Get()
	}
	if !IsNil(o.ExternalResources) {
		toSerialize["externalResources"] = o.ExternalResources
	}
	if !IsNil(o.DefaultFolderType) {
		toSerialize["defaultFolderType"] = o.DefaultFolderType
	}
	if !IsNil(o.ExternalDbEnabled) {
		toSerialize["externalDbEnabled"] = o.ExternalDbEnabled
	}
	return toSerialize, nil
}

func (o *SettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"culture",
		"baseDomain",
		"cookieSettingsEnabled",
		"deepLink",
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

	varSettingsDto := _SettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSettingsDto)

	if err != nil {
		return err
	}

	*o = SettingsDto(varSettingsDto)

	return err
}

type NullableSettingsDto struct {
	value *SettingsDto
	isSet bool
}

func (v NullableSettingsDto) Get() *SettingsDto {
	return v.value
}

func (v *NullableSettingsDto) Set(val *SettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsDto(val *SettingsDto) *NullableSettingsDto {
	return &NullableSettingsDto{value: val, isSet: true}
}

func (v NullableSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

