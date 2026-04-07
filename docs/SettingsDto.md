# SettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **NullableString** | The time zone. | [optional] 
**TrustedDomains** | Pointer to **[]string** | The list of the trusted domains. | [optional] 
**TrustedDomainsType** | Pointer to [**TenantTrustedDomainsType**](TenantTrustedDomainsType.md) |  | [optional] 
**Culture** | **NullableString** | The language. | 
**UtcOffset** | Pointer to **string** | The UTC offset in the TimeSpan format. | [optional] 
**UtcHoursOffset** | Pointer to **float64** | The UTC offset in hours. | [optional] 
**GreetingSettings** | Pointer to **NullableString** | The greeting settings. | [optional] 
**OwnerId** | Pointer to **string** | The owner ID. | [optional] 
**NameSchemaId** | Pointer to **NullableString** | The team template ID. | [optional] 
**EnabledJoin** | Pointer to **NullableBool** | Specifies if a user can join the portal or not. | [optional] 
**EnableAdmMess** | Pointer to **NullableBool** | Specifies if a user can send a message to the administrator when accessing the DocSpace portal or not. | [optional] 
**ThirdpartyEnable** | Pointer to **NullableBool** | Specifies if a user can connect third-party providers to the portal or not. | [optional] 
**DocSpace** | Pointer to **bool** | Specifies if this portal is a DocSpace portal or not. | [optional] 
**Standalone** | Pointer to **bool** | Indicates whether the system is running in standalone mode. | [optional] 
**IsAmi** | Pointer to **bool** | Specifies if this portal is the AMI instance or not. | [optional] 
**BaseDomain** | **NullableString** | The base domain. | 
**WizardToken** | Pointer to **NullableString** | The wizard token. | [optional] 
**PasswordHash** | Pointer to [**PasswordHasher**](PasswordHasher.md) |  | [optional] 
**Firebase** | Pointer to [**FirebaseDto**](FirebaseDto.md) |  | [optional] 
**Version** | Pointer to **NullableString** | The portal version. | [optional] 
**RecaptchaType** | Pointer to [**RecaptchaType**](RecaptchaType.md) |  | [optional] 
**RecaptchaPublicKey** | Pointer to **NullableString** | The ReCAPTCHA public key. | [optional] 
**DebugInfo** | Pointer to **bool** | Specifies if the debug information will be sent or not. | [optional] 
**SocketUrl** | Pointer to **NullableString** | The socket URL. | [optional] 
**TenantStatus** | Pointer to [**TenantStatus**](TenantStatus.md) |  | [optional] 
**TenantAlias** | Pointer to **NullableString** | The tenant alias. | [optional] 
**DisplayAbout** | Pointer to **bool** | Specifies whether to display the About portal section. | [optional] 
**DomainValidator** | Pointer to [**TenantDomainValidator**](TenantDomainValidator.md) |  | [optional] 
**ZendeskKey** | Pointer to **NullableString** | The Zendesk key. | [optional] 
**TagManagerId** | Pointer to **NullableString** | The tag manager ID. | [optional] 
**CookieSettingsEnabled** | **bool** | Specifies whether the cookie settings are enabled. | 
**LimitedAccessSpace** | Pointer to **bool** | Specifies whether the access to the space management is limited or not. | [optional] 
**LimitedAccessDevToolsForUsers** | Pointer to **bool** | Specifies whether the access to the Developer Tools is limited for users or not. | [optional] 
**DisplayBanners** | Pointer to **bool** | Specifies whether to display the promotional banners. | [optional] 
**AiEnabled** | Pointer to **bool** | Specifies whether AI functionality (chat, agents, vectorization) is enabled for the current tenant.  When `false`, all AI features are disabled and the AI Agents folder is hidden. | [optional] 
**UserNameRegex** | Pointer to **NullableString** | The user name validation regex. | [optional] 
**InvitationLimit** | Pointer to **NullableInt32** | The maximum number of invitations to the portal. | [optional] 
**Plugins** | Pointer to [**PluginsDto**](PluginsDto.md) |  | [optional] 
**DeepLink** | [**DeepLinkDto**](DeepLinkDto.md) |  | 
**FormGallery** | Pointer to [**FormGalleryDto**](FormGalleryDto.md) |  | [optional] 
**MaxImageUploadSize** | Pointer to **int64** | The maximum image upload size. | [optional] 
**LogoText** | Pointer to **NullableString** | The white label logo text. | [optional] 
**ExternalResources** | Pointer to [**CultureSpecificExternalResources**](CultureSpecificExternalResources.md) |  | [optional] 
**DefaultFolderType** | Pointer to [**FolderType**](FolderType.md) |  | [optional] 
**ExternalDbEnabled** | Pointer to **bool** | Specifies if an external database is connected for storing form results. | [optional] 

## Methods

### NewSettingsDto

`func NewSettingsDto(culture NullableString, baseDomain NullableString, cookieSettingsEnabled bool, deepLink DeepLinkDto, ) *SettingsDto`

NewSettingsDto instantiates a new SettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsDtoWithDefaults

`func NewSettingsDtoWithDefaults() *SettingsDto`

NewSettingsDtoWithDefaults instantiates a new SettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *SettingsDto) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SettingsDto) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SettingsDto) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SettingsDto) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *SettingsDto) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *SettingsDto) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetTrustedDomains

`func (o *SettingsDto) GetTrustedDomains() []string`

GetTrustedDomains returns the TrustedDomains field if non-nil, zero value otherwise.

### GetTrustedDomainsOk

`func (o *SettingsDto) GetTrustedDomainsOk() (*[]string, bool)`

GetTrustedDomainsOk returns a tuple with the TrustedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomains

`func (o *SettingsDto) SetTrustedDomains(v []string)`

SetTrustedDomains sets TrustedDomains field to given value.

### HasTrustedDomains

`func (o *SettingsDto) HasTrustedDomains() bool`

HasTrustedDomains returns a boolean if a field has been set.

### SetTrustedDomainsNil

`func (o *SettingsDto) SetTrustedDomainsNil(b bool)`

 SetTrustedDomainsNil sets the value for TrustedDomains to be an explicit nil

### UnsetTrustedDomains
`func (o *SettingsDto) UnsetTrustedDomains()`

UnsetTrustedDomains ensures that no value is present for TrustedDomains, not even an explicit nil
### GetTrustedDomainsType

`func (o *SettingsDto) GetTrustedDomainsType() TenantTrustedDomainsType`

GetTrustedDomainsType returns the TrustedDomainsType field if non-nil, zero value otherwise.

### GetTrustedDomainsTypeOk

`func (o *SettingsDto) GetTrustedDomainsTypeOk() (*TenantTrustedDomainsType, bool)`

GetTrustedDomainsTypeOk returns a tuple with the TrustedDomainsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainsType

`func (o *SettingsDto) SetTrustedDomainsType(v TenantTrustedDomainsType)`

SetTrustedDomainsType sets TrustedDomainsType field to given value.

### HasTrustedDomainsType

`func (o *SettingsDto) HasTrustedDomainsType() bool`

HasTrustedDomainsType returns a boolean if a field has been set.

### GetCulture

`func (o *SettingsDto) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *SettingsDto) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *SettingsDto) SetCulture(v string)`

SetCulture sets Culture field to given value.


### SetCultureNil

`func (o *SettingsDto) SetCultureNil(b bool)`

 SetCultureNil sets the value for Culture to be an explicit nil

### UnsetCulture
`func (o *SettingsDto) UnsetCulture()`

UnsetCulture ensures that no value is present for Culture, not even an explicit nil
### GetUtcOffset

`func (o *SettingsDto) GetUtcOffset() string`

GetUtcOffset returns the UtcOffset field if non-nil, zero value otherwise.

### GetUtcOffsetOk

`func (o *SettingsDto) GetUtcOffsetOk() (*string, bool)`

GetUtcOffsetOk returns a tuple with the UtcOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffset

`func (o *SettingsDto) SetUtcOffset(v string)`

SetUtcOffset sets UtcOffset field to given value.

### HasUtcOffset

`func (o *SettingsDto) HasUtcOffset() bool`

HasUtcOffset returns a boolean if a field has been set.

### GetUtcHoursOffset

`func (o *SettingsDto) GetUtcHoursOffset() float64`

GetUtcHoursOffset returns the UtcHoursOffset field if non-nil, zero value otherwise.

### GetUtcHoursOffsetOk

`func (o *SettingsDto) GetUtcHoursOffsetOk() (*float64, bool)`

GetUtcHoursOffsetOk returns a tuple with the UtcHoursOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcHoursOffset

`func (o *SettingsDto) SetUtcHoursOffset(v float64)`

SetUtcHoursOffset sets UtcHoursOffset field to given value.

### HasUtcHoursOffset

`func (o *SettingsDto) HasUtcHoursOffset() bool`

HasUtcHoursOffset returns a boolean if a field has been set.

### GetGreetingSettings

`func (o *SettingsDto) GetGreetingSettings() string`

GetGreetingSettings returns the GreetingSettings field if non-nil, zero value otherwise.

### GetGreetingSettingsOk

`func (o *SettingsDto) GetGreetingSettingsOk() (*string, bool)`

GetGreetingSettingsOk returns a tuple with the GreetingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingSettings

`func (o *SettingsDto) SetGreetingSettings(v string)`

SetGreetingSettings sets GreetingSettings field to given value.

### HasGreetingSettings

`func (o *SettingsDto) HasGreetingSettings() bool`

HasGreetingSettings returns a boolean if a field has been set.

### SetGreetingSettingsNil

`func (o *SettingsDto) SetGreetingSettingsNil(b bool)`

 SetGreetingSettingsNil sets the value for GreetingSettings to be an explicit nil

### UnsetGreetingSettings
`func (o *SettingsDto) UnsetGreetingSettings()`

UnsetGreetingSettings ensures that no value is present for GreetingSettings, not even an explicit nil
### GetOwnerId

`func (o *SettingsDto) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SettingsDto) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SettingsDto) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SettingsDto) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetNameSchemaId

`func (o *SettingsDto) GetNameSchemaId() string`

GetNameSchemaId returns the NameSchemaId field if non-nil, zero value otherwise.

### GetNameSchemaIdOk

`func (o *SettingsDto) GetNameSchemaIdOk() (*string, bool)`

GetNameSchemaIdOk returns a tuple with the NameSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSchemaId

`func (o *SettingsDto) SetNameSchemaId(v string)`

SetNameSchemaId sets NameSchemaId field to given value.

### HasNameSchemaId

`func (o *SettingsDto) HasNameSchemaId() bool`

HasNameSchemaId returns a boolean if a field has been set.

### SetNameSchemaIdNil

`func (o *SettingsDto) SetNameSchemaIdNil(b bool)`

 SetNameSchemaIdNil sets the value for NameSchemaId to be an explicit nil

### UnsetNameSchemaId
`func (o *SettingsDto) UnsetNameSchemaId()`

UnsetNameSchemaId ensures that no value is present for NameSchemaId, not even an explicit nil
### GetEnabledJoin

`func (o *SettingsDto) GetEnabledJoin() bool`

GetEnabledJoin returns the EnabledJoin field if non-nil, zero value otherwise.

### GetEnabledJoinOk

`func (o *SettingsDto) GetEnabledJoinOk() (*bool, bool)`

GetEnabledJoinOk returns a tuple with the EnabledJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledJoin

`func (o *SettingsDto) SetEnabledJoin(v bool)`

SetEnabledJoin sets EnabledJoin field to given value.

### HasEnabledJoin

`func (o *SettingsDto) HasEnabledJoin() bool`

HasEnabledJoin returns a boolean if a field has been set.

### SetEnabledJoinNil

`func (o *SettingsDto) SetEnabledJoinNil(b bool)`

 SetEnabledJoinNil sets the value for EnabledJoin to be an explicit nil

### UnsetEnabledJoin
`func (o *SettingsDto) UnsetEnabledJoin()`

UnsetEnabledJoin ensures that no value is present for EnabledJoin, not even an explicit nil
### GetEnableAdmMess

`func (o *SettingsDto) GetEnableAdmMess() bool`

GetEnableAdmMess returns the EnableAdmMess field if non-nil, zero value otherwise.

### GetEnableAdmMessOk

`func (o *SettingsDto) GetEnableAdmMessOk() (*bool, bool)`

GetEnableAdmMessOk returns a tuple with the EnableAdmMess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdmMess

`func (o *SettingsDto) SetEnableAdmMess(v bool)`

SetEnableAdmMess sets EnableAdmMess field to given value.

### HasEnableAdmMess

`func (o *SettingsDto) HasEnableAdmMess() bool`

HasEnableAdmMess returns a boolean if a field has been set.

### SetEnableAdmMessNil

`func (o *SettingsDto) SetEnableAdmMessNil(b bool)`

 SetEnableAdmMessNil sets the value for EnableAdmMess to be an explicit nil

### UnsetEnableAdmMess
`func (o *SettingsDto) UnsetEnableAdmMess()`

UnsetEnableAdmMess ensures that no value is present for EnableAdmMess, not even an explicit nil
### GetThirdpartyEnable

`func (o *SettingsDto) GetThirdpartyEnable() bool`

GetThirdpartyEnable returns the ThirdpartyEnable field if non-nil, zero value otherwise.

### GetThirdpartyEnableOk

`func (o *SettingsDto) GetThirdpartyEnableOk() (*bool, bool)`

GetThirdpartyEnableOk returns a tuple with the ThirdpartyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdpartyEnable

`func (o *SettingsDto) SetThirdpartyEnable(v bool)`

SetThirdpartyEnable sets ThirdpartyEnable field to given value.

### HasThirdpartyEnable

`func (o *SettingsDto) HasThirdpartyEnable() bool`

HasThirdpartyEnable returns a boolean if a field has been set.

### SetThirdpartyEnableNil

`func (o *SettingsDto) SetThirdpartyEnableNil(b bool)`

 SetThirdpartyEnableNil sets the value for ThirdpartyEnable to be an explicit nil

### UnsetThirdpartyEnable
`func (o *SettingsDto) UnsetThirdpartyEnable()`

UnsetThirdpartyEnable ensures that no value is present for ThirdpartyEnable, not even an explicit nil
### GetDocSpace

`func (o *SettingsDto) GetDocSpace() bool`

GetDocSpace returns the DocSpace field if non-nil, zero value otherwise.

### GetDocSpaceOk

`func (o *SettingsDto) GetDocSpaceOk() (*bool, bool)`

GetDocSpaceOk returns a tuple with the DocSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocSpace

`func (o *SettingsDto) SetDocSpace(v bool)`

SetDocSpace sets DocSpace field to given value.

### HasDocSpace

`func (o *SettingsDto) HasDocSpace() bool`

HasDocSpace returns a boolean if a field has been set.

### GetStandalone

`func (o *SettingsDto) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *SettingsDto) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *SettingsDto) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *SettingsDto) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetIsAmi

`func (o *SettingsDto) GetIsAmi() bool`

GetIsAmi returns the IsAmi field if non-nil, zero value otherwise.

### GetIsAmiOk

`func (o *SettingsDto) GetIsAmiOk() (*bool, bool)`

GetIsAmiOk returns a tuple with the IsAmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmi

`func (o *SettingsDto) SetIsAmi(v bool)`

SetIsAmi sets IsAmi field to given value.

### HasIsAmi

`func (o *SettingsDto) HasIsAmi() bool`

HasIsAmi returns a boolean if a field has been set.

### GetBaseDomain

`func (o *SettingsDto) GetBaseDomain() string`

GetBaseDomain returns the BaseDomain field if non-nil, zero value otherwise.

### GetBaseDomainOk

`func (o *SettingsDto) GetBaseDomainOk() (*string, bool)`

GetBaseDomainOk returns a tuple with the BaseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDomain

`func (o *SettingsDto) SetBaseDomain(v string)`

SetBaseDomain sets BaseDomain field to given value.


### SetBaseDomainNil

`func (o *SettingsDto) SetBaseDomainNil(b bool)`

 SetBaseDomainNil sets the value for BaseDomain to be an explicit nil

### UnsetBaseDomain
`func (o *SettingsDto) UnsetBaseDomain()`

UnsetBaseDomain ensures that no value is present for BaseDomain, not even an explicit nil
### GetWizardToken

`func (o *SettingsDto) GetWizardToken() string`

GetWizardToken returns the WizardToken field if non-nil, zero value otherwise.

### GetWizardTokenOk

`func (o *SettingsDto) GetWizardTokenOk() (*string, bool)`

GetWizardTokenOk returns a tuple with the WizardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWizardToken

`func (o *SettingsDto) SetWizardToken(v string)`

SetWizardToken sets WizardToken field to given value.

### HasWizardToken

`func (o *SettingsDto) HasWizardToken() bool`

HasWizardToken returns a boolean if a field has been set.

### SetWizardTokenNil

`func (o *SettingsDto) SetWizardTokenNil(b bool)`

 SetWizardTokenNil sets the value for WizardToken to be an explicit nil

### UnsetWizardToken
`func (o *SettingsDto) UnsetWizardToken()`

UnsetWizardToken ensures that no value is present for WizardToken, not even an explicit nil
### GetPasswordHash

`func (o *SettingsDto) GetPasswordHash() PasswordHasher`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *SettingsDto) GetPasswordHashOk() (*PasswordHasher, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *SettingsDto) SetPasswordHash(v PasswordHasher)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *SettingsDto) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetFirebase

`func (o *SettingsDto) GetFirebase() FirebaseDto`

GetFirebase returns the Firebase field if non-nil, zero value otherwise.

### GetFirebaseOk

`func (o *SettingsDto) GetFirebaseOk() (*FirebaseDto, bool)`

GetFirebaseOk returns a tuple with the Firebase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirebase

`func (o *SettingsDto) SetFirebase(v FirebaseDto)`

SetFirebase sets Firebase field to given value.

### HasFirebase

`func (o *SettingsDto) HasFirebase() bool`

HasFirebase returns a boolean if a field has been set.

### GetVersion

`func (o *SettingsDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SettingsDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SettingsDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SettingsDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SettingsDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SettingsDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetRecaptchaType

`func (o *SettingsDto) GetRecaptchaType() RecaptchaType`

GetRecaptchaType returns the RecaptchaType field if non-nil, zero value otherwise.

### GetRecaptchaTypeOk

`func (o *SettingsDto) GetRecaptchaTypeOk() (*RecaptchaType, bool)`

GetRecaptchaTypeOk returns a tuple with the RecaptchaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaType

`func (o *SettingsDto) SetRecaptchaType(v RecaptchaType)`

SetRecaptchaType sets RecaptchaType field to given value.

### HasRecaptchaType

`func (o *SettingsDto) HasRecaptchaType() bool`

HasRecaptchaType returns a boolean if a field has been set.

### GetRecaptchaPublicKey

`func (o *SettingsDto) GetRecaptchaPublicKey() string`

GetRecaptchaPublicKey returns the RecaptchaPublicKey field if non-nil, zero value otherwise.

### GetRecaptchaPublicKeyOk

`func (o *SettingsDto) GetRecaptchaPublicKeyOk() (*string, bool)`

GetRecaptchaPublicKeyOk returns a tuple with the RecaptchaPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaPublicKey

`func (o *SettingsDto) SetRecaptchaPublicKey(v string)`

SetRecaptchaPublicKey sets RecaptchaPublicKey field to given value.

### HasRecaptchaPublicKey

`func (o *SettingsDto) HasRecaptchaPublicKey() bool`

HasRecaptchaPublicKey returns a boolean if a field has been set.

### SetRecaptchaPublicKeyNil

`func (o *SettingsDto) SetRecaptchaPublicKeyNil(b bool)`

 SetRecaptchaPublicKeyNil sets the value for RecaptchaPublicKey to be an explicit nil

### UnsetRecaptchaPublicKey
`func (o *SettingsDto) UnsetRecaptchaPublicKey()`

UnsetRecaptchaPublicKey ensures that no value is present for RecaptchaPublicKey, not even an explicit nil
### GetDebugInfo

`func (o *SettingsDto) GetDebugInfo() bool`

GetDebugInfo returns the DebugInfo field if non-nil, zero value otherwise.

### GetDebugInfoOk

`func (o *SettingsDto) GetDebugInfoOk() (*bool, bool)`

GetDebugInfoOk returns a tuple with the DebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugInfo

`func (o *SettingsDto) SetDebugInfo(v bool)`

SetDebugInfo sets DebugInfo field to given value.

### HasDebugInfo

`func (o *SettingsDto) HasDebugInfo() bool`

HasDebugInfo returns a boolean if a field has been set.

### GetSocketUrl

`func (o *SettingsDto) GetSocketUrl() string`

GetSocketUrl returns the SocketUrl field if non-nil, zero value otherwise.

### GetSocketUrlOk

`func (o *SettingsDto) GetSocketUrlOk() (*string, bool)`

GetSocketUrlOk returns a tuple with the SocketUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketUrl

`func (o *SettingsDto) SetSocketUrl(v string)`

SetSocketUrl sets SocketUrl field to given value.

### HasSocketUrl

`func (o *SettingsDto) HasSocketUrl() bool`

HasSocketUrl returns a boolean if a field has been set.

### SetSocketUrlNil

`func (o *SettingsDto) SetSocketUrlNil(b bool)`

 SetSocketUrlNil sets the value for SocketUrl to be an explicit nil

### UnsetSocketUrl
`func (o *SettingsDto) UnsetSocketUrl()`

UnsetSocketUrl ensures that no value is present for SocketUrl, not even an explicit nil
### GetTenantStatus

`func (o *SettingsDto) GetTenantStatus() TenantStatus`

GetTenantStatus returns the TenantStatus field if non-nil, zero value otherwise.

### GetTenantStatusOk

`func (o *SettingsDto) GetTenantStatusOk() (*TenantStatus, bool)`

GetTenantStatusOk returns a tuple with the TenantStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantStatus

`func (o *SettingsDto) SetTenantStatus(v TenantStatus)`

SetTenantStatus sets TenantStatus field to given value.

### HasTenantStatus

`func (o *SettingsDto) HasTenantStatus() bool`

HasTenantStatus returns a boolean if a field has been set.

### GetTenantAlias

`func (o *SettingsDto) GetTenantAlias() string`

GetTenantAlias returns the TenantAlias field if non-nil, zero value otherwise.

### GetTenantAliasOk

`func (o *SettingsDto) GetTenantAliasOk() (*string, bool)`

GetTenantAliasOk returns a tuple with the TenantAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantAlias

`func (o *SettingsDto) SetTenantAlias(v string)`

SetTenantAlias sets TenantAlias field to given value.

### HasTenantAlias

`func (o *SettingsDto) HasTenantAlias() bool`

HasTenantAlias returns a boolean if a field has been set.

### SetTenantAliasNil

`func (o *SettingsDto) SetTenantAliasNil(b bool)`

 SetTenantAliasNil sets the value for TenantAlias to be an explicit nil

### UnsetTenantAlias
`func (o *SettingsDto) UnsetTenantAlias()`

UnsetTenantAlias ensures that no value is present for TenantAlias, not even an explicit nil
### GetDisplayAbout

`func (o *SettingsDto) GetDisplayAbout() bool`

GetDisplayAbout returns the DisplayAbout field if non-nil, zero value otherwise.

### GetDisplayAboutOk

`func (o *SettingsDto) GetDisplayAboutOk() (*bool, bool)`

GetDisplayAboutOk returns a tuple with the DisplayAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAbout

`func (o *SettingsDto) SetDisplayAbout(v bool)`

SetDisplayAbout sets DisplayAbout field to given value.

### HasDisplayAbout

`func (o *SettingsDto) HasDisplayAbout() bool`

HasDisplayAbout returns a boolean if a field has been set.

### GetDomainValidator

`func (o *SettingsDto) GetDomainValidator() TenantDomainValidator`

GetDomainValidator returns the DomainValidator field if non-nil, zero value otherwise.

### GetDomainValidatorOk

`func (o *SettingsDto) GetDomainValidatorOk() (*TenantDomainValidator, bool)`

GetDomainValidatorOk returns a tuple with the DomainValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainValidator

`func (o *SettingsDto) SetDomainValidator(v TenantDomainValidator)`

SetDomainValidator sets DomainValidator field to given value.

### HasDomainValidator

`func (o *SettingsDto) HasDomainValidator() bool`

HasDomainValidator returns a boolean if a field has been set.

### GetZendeskKey

`func (o *SettingsDto) GetZendeskKey() string`

GetZendeskKey returns the ZendeskKey field if non-nil, zero value otherwise.

### GetZendeskKeyOk

`func (o *SettingsDto) GetZendeskKeyOk() (*string, bool)`

GetZendeskKeyOk returns a tuple with the ZendeskKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskKey

`func (o *SettingsDto) SetZendeskKey(v string)`

SetZendeskKey sets ZendeskKey field to given value.

### HasZendeskKey

`func (o *SettingsDto) HasZendeskKey() bool`

HasZendeskKey returns a boolean if a field has been set.

### SetZendeskKeyNil

`func (o *SettingsDto) SetZendeskKeyNil(b bool)`

 SetZendeskKeyNil sets the value for ZendeskKey to be an explicit nil

### UnsetZendeskKey
`func (o *SettingsDto) UnsetZendeskKey()`

UnsetZendeskKey ensures that no value is present for ZendeskKey, not even an explicit nil
### GetTagManagerId

`func (o *SettingsDto) GetTagManagerId() string`

GetTagManagerId returns the TagManagerId field if non-nil, zero value otherwise.

### GetTagManagerIdOk

`func (o *SettingsDto) GetTagManagerIdOk() (*string, bool)`

GetTagManagerIdOk returns a tuple with the TagManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagManagerId

`func (o *SettingsDto) SetTagManagerId(v string)`

SetTagManagerId sets TagManagerId field to given value.

### HasTagManagerId

`func (o *SettingsDto) HasTagManagerId() bool`

HasTagManagerId returns a boolean if a field has been set.

### SetTagManagerIdNil

`func (o *SettingsDto) SetTagManagerIdNil(b bool)`

 SetTagManagerIdNil sets the value for TagManagerId to be an explicit nil

### UnsetTagManagerId
`func (o *SettingsDto) UnsetTagManagerId()`

UnsetTagManagerId ensures that no value is present for TagManagerId, not even an explicit nil
### GetCookieSettingsEnabled

`func (o *SettingsDto) GetCookieSettingsEnabled() bool`

GetCookieSettingsEnabled returns the CookieSettingsEnabled field if non-nil, zero value otherwise.

### GetCookieSettingsEnabledOk

`func (o *SettingsDto) GetCookieSettingsEnabledOk() (*bool, bool)`

GetCookieSettingsEnabledOk returns a tuple with the CookieSettingsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieSettingsEnabled

`func (o *SettingsDto) SetCookieSettingsEnabled(v bool)`

SetCookieSettingsEnabled sets CookieSettingsEnabled field to given value.


### GetLimitedAccessSpace

`func (o *SettingsDto) GetLimitedAccessSpace() bool`

GetLimitedAccessSpace returns the LimitedAccessSpace field if non-nil, zero value otherwise.

### GetLimitedAccessSpaceOk

`func (o *SettingsDto) GetLimitedAccessSpaceOk() (*bool, bool)`

GetLimitedAccessSpaceOk returns a tuple with the LimitedAccessSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedAccessSpace

`func (o *SettingsDto) SetLimitedAccessSpace(v bool)`

SetLimitedAccessSpace sets LimitedAccessSpace field to given value.

### HasLimitedAccessSpace

`func (o *SettingsDto) HasLimitedAccessSpace() bool`

HasLimitedAccessSpace returns a boolean if a field has been set.

### GetLimitedAccessDevToolsForUsers

`func (o *SettingsDto) GetLimitedAccessDevToolsForUsers() bool`

GetLimitedAccessDevToolsForUsers returns the LimitedAccessDevToolsForUsers field if non-nil, zero value otherwise.

### GetLimitedAccessDevToolsForUsersOk

`func (o *SettingsDto) GetLimitedAccessDevToolsForUsersOk() (*bool, bool)`

GetLimitedAccessDevToolsForUsersOk returns a tuple with the LimitedAccessDevToolsForUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedAccessDevToolsForUsers

`func (o *SettingsDto) SetLimitedAccessDevToolsForUsers(v bool)`

SetLimitedAccessDevToolsForUsers sets LimitedAccessDevToolsForUsers field to given value.

### HasLimitedAccessDevToolsForUsers

`func (o *SettingsDto) HasLimitedAccessDevToolsForUsers() bool`

HasLimitedAccessDevToolsForUsers returns a boolean if a field has been set.

### GetDisplayBanners

`func (o *SettingsDto) GetDisplayBanners() bool`

GetDisplayBanners returns the DisplayBanners field if non-nil, zero value otherwise.

### GetDisplayBannersOk

`func (o *SettingsDto) GetDisplayBannersOk() (*bool, bool)`

GetDisplayBannersOk returns a tuple with the DisplayBanners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayBanners

`func (o *SettingsDto) SetDisplayBanners(v bool)`

SetDisplayBanners sets DisplayBanners field to given value.

### HasDisplayBanners

`func (o *SettingsDto) HasDisplayBanners() bool`

HasDisplayBanners returns a boolean if a field has been set.

### GetAiEnabled

`func (o *SettingsDto) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *SettingsDto) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *SettingsDto) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.

### HasAiEnabled

`func (o *SettingsDto) HasAiEnabled() bool`

HasAiEnabled returns a boolean if a field has been set.

### GetUserNameRegex

`func (o *SettingsDto) GetUserNameRegex() string`

GetUserNameRegex returns the UserNameRegex field if non-nil, zero value otherwise.

### GetUserNameRegexOk

`func (o *SettingsDto) GetUserNameRegexOk() (*string, bool)`

GetUserNameRegexOk returns a tuple with the UserNameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameRegex

`func (o *SettingsDto) SetUserNameRegex(v string)`

SetUserNameRegex sets UserNameRegex field to given value.

### HasUserNameRegex

`func (o *SettingsDto) HasUserNameRegex() bool`

HasUserNameRegex returns a boolean if a field has been set.

### SetUserNameRegexNil

`func (o *SettingsDto) SetUserNameRegexNil(b bool)`

 SetUserNameRegexNil sets the value for UserNameRegex to be an explicit nil

### UnsetUserNameRegex
`func (o *SettingsDto) UnsetUserNameRegex()`

UnsetUserNameRegex ensures that no value is present for UserNameRegex, not even an explicit nil
### GetInvitationLimit

`func (o *SettingsDto) GetInvitationLimit() int32`

GetInvitationLimit returns the InvitationLimit field if non-nil, zero value otherwise.

### GetInvitationLimitOk

`func (o *SettingsDto) GetInvitationLimitOk() (*int32, bool)`

GetInvitationLimitOk returns a tuple with the InvitationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationLimit

`func (o *SettingsDto) SetInvitationLimit(v int32)`

SetInvitationLimit sets InvitationLimit field to given value.

### HasInvitationLimit

`func (o *SettingsDto) HasInvitationLimit() bool`

HasInvitationLimit returns a boolean if a field has been set.

### SetInvitationLimitNil

`func (o *SettingsDto) SetInvitationLimitNil(b bool)`

 SetInvitationLimitNil sets the value for InvitationLimit to be an explicit nil

### UnsetInvitationLimit
`func (o *SettingsDto) UnsetInvitationLimit()`

UnsetInvitationLimit ensures that no value is present for InvitationLimit, not even an explicit nil
### GetPlugins

`func (o *SettingsDto) GetPlugins() PluginsDto`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *SettingsDto) GetPluginsOk() (*PluginsDto, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *SettingsDto) SetPlugins(v PluginsDto)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *SettingsDto) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetDeepLink

`func (o *SettingsDto) GetDeepLink() DeepLinkDto`

GetDeepLink returns the DeepLink field if non-nil, zero value otherwise.

### GetDeepLinkOk

`func (o *SettingsDto) GetDeepLinkOk() (*DeepLinkDto, bool)`

GetDeepLinkOk returns a tuple with the DeepLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepLink

`func (o *SettingsDto) SetDeepLink(v DeepLinkDto)`

SetDeepLink sets DeepLink field to given value.


### GetFormGallery

`func (o *SettingsDto) GetFormGallery() FormGalleryDto`

GetFormGallery returns the FormGallery field if non-nil, zero value otherwise.

### GetFormGalleryOk

`func (o *SettingsDto) GetFormGalleryOk() (*FormGalleryDto, bool)`

GetFormGalleryOk returns a tuple with the FormGallery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormGallery

`func (o *SettingsDto) SetFormGallery(v FormGalleryDto)`

SetFormGallery sets FormGallery field to given value.

### HasFormGallery

`func (o *SettingsDto) HasFormGallery() bool`

HasFormGallery returns a boolean if a field has been set.

### GetMaxImageUploadSize

`func (o *SettingsDto) GetMaxImageUploadSize() int64`

GetMaxImageUploadSize returns the MaxImageUploadSize field if non-nil, zero value otherwise.

### GetMaxImageUploadSizeOk

`func (o *SettingsDto) GetMaxImageUploadSizeOk() (*int64, bool)`

GetMaxImageUploadSizeOk returns a tuple with the MaxImageUploadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxImageUploadSize

`func (o *SettingsDto) SetMaxImageUploadSize(v int64)`

SetMaxImageUploadSize sets MaxImageUploadSize field to given value.

### HasMaxImageUploadSize

`func (o *SettingsDto) HasMaxImageUploadSize() bool`

HasMaxImageUploadSize returns a boolean if a field has been set.

### GetLogoText

`func (o *SettingsDto) GetLogoText() string`

GetLogoText returns the LogoText field if non-nil, zero value otherwise.

### GetLogoTextOk

`func (o *SettingsDto) GetLogoTextOk() (*string, bool)`

GetLogoTextOk returns a tuple with the LogoText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoText

`func (o *SettingsDto) SetLogoText(v string)`

SetLogoText sets LogoText field to given value.

### HasLogoText

`func (o *SettingsDto) HasLogoText() bool`

HasLogoText returns a boolean if a field has been set.

### SetLogoTextNil

`func (o *SettingsDto) SetLogoTextNil(b bool)`

 SetLogoTextNil sets the value for LogoText to be an explicit nil

### UnsetLogoText
`func (o *SettingsDto) UnsetLogoText()`

UnsetLogoText ensures that no value is present for LogoText, not even an explicit nil
### GetExternalResources

`func (o *SettingsDto) GetExternalResources() CultureSpecificExternalResources`

GetExternalResources returns the ExternalResources field if non-nil, zero value otherwise.

### GetExternalResourcesOk

`func (o *SettingsDto) GetExternalResourcesOk() (*CultureSpecificExternalResources, bool)`

GetExternalResourcesOk returns a tuple with the ExternalResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalResources

`func (o *SettingsDto) SetExternalResources(v CultureSpecificExternalResources)`

SetExternalResources sets ExternalResources field to given value.

### HasExternalResources

`func (o *SettingsDto) HasExternalResources() bool`

HasExternalResources returns a boolean if a field has been set.

### GetDefaultFolderType

`func (o *SettingsDto) GetDefaultFolderType() FolderType`

GetDefaultFolderType returns the DefaultFolderType field if non-nil, zero value otherwise.

### GetDefaultFolderTypeOk

`func (o *SettingsDto) GetDefaultFolderTypeOk() (*FolderType, bool)`

GetDefaultFolderTypeOk returns a tuple with the DefaultFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolderType

`func (o *SettingsDto) SetDefaultFolderType(v FolderType)`

SetDefaultFolderType sets DefaultFolderType field to given value.

### HasDefaultFolderType

`func (o *SettingsDto) HasDefaultFolderType() bool`

HasDefaultFolderType returns a boolean if a field has been set.

### GetExternalDbEnabled

`func (o *SettingsDto) GetExternalDbEnabled() bool`

GetExternalDbEnabled returns the ExternalDbEnabled field if non-nil, zero value otherwise.

### GetExternalDbEnabledOk

`func (o *SettingsDto) GetExternalDbEnabledOk() (*bool, bool)`

GetExternalDbEnabledOk returns a tuple with the ExternalDbEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDbEnabled

`func (o *SettingsDto) SetExternalDbEnabled(v bool)`

SetExternalDbEnabled sets ExternalDbEnabled field to given value.

### HasExternalDbEnabled

`func (o *SettingsDto) HasExternalDbEnabled() bool`

HasExternalDbEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


