# TenantDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffiliateId** | Pointer to **NullableString** | The affiliate ID. | [optional] 
**TenantAlias** | Pointer to **NullableString** | The tenant alias. | [optional] 
**Calls** | Pointer to **bool** | Specifies if the calls are available for this tenant or not. | [optional] 
**Campaign** | Pointer to **NullableString** | The tenant campaign. | [optional] 
**CreationDateTime** | Pointer to **time.Time** | The tenant creation date and time. | [optional] [readonly] 
**HostedRegion** | Pointer to **NullableString** | The hosted region. | [optional] 
**TenantId** | Pointer to **int32** | The tenant ID. | [optional] [readonly] 
**Industry** | Pointer to [**TenantIndustry**](TenantIndustry.md) | The tenant industry. | [optional] 
**Language** | Pointer to **NullableString** | The tenant language. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the tenant was last modified. | [optional] 
**MappedDomain** | Pointer to **NullableString** | The tenant mapped domain. | [optional] 
**Name** | Pointer to **NullableString** | The tenant name. | [optional] 
**OwnerId** | Pointer to **string** | The tenant owner ID. | [optional] 
**PaymentId** | Pointer to **NullableString** | The tenant payment ID. | [optional] 
**Spam** | Pointer to **bool** | Specifies if the ONLYOFFICE newsletter is allowed or not. | [optional] 
**Status** | Pointer to [**TenantStatus**](TenantStatus.md) | The tenant status. | [optional] 
**StatusChangeDate** | Pointer to **time.Time** | The date and time when the tenant status was changed. | [optional] [readonly] 
**TimeZone** | Pointer to **NullableString** | The tenant time zone. | [optional] 
**TrustedDomains** | Pointer to **[]string** | The list of tenant trusted domains. | [optional] 
**TrustedDomainsRaw** | Pointer to **NullableString** | The tenant trusted domains in the string format. | [optional] 
**TrustedDomainsType** | Pointer to [**TenantTrustedDomainsType**](TenantTrustedDomainsType.md) | The type of the tenant trusted domains. | [optional] 
**Version** | Pointer to **int32** | The tenant version | [optional] 
**VersionChanged** | Pointer to **time.Time** | The date and time when the tenant version was changed. | [optional] 
**Region** | Pointer to **NullableString** | The tenant AWS region. | [optional] 

## Methods

### NewTenantDto

`func NewTenantDto() *TenantDto`

NewTenantDto instantiates a new TenantDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDtoWithDefaults

`func NewTenantDtoWithDefaults() *TenantDto`

NewTenantDtoWithDefaults instantiates a new TenantDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliateId

`func (o *TenantDto) GetAffiliateId() string`

GetAffiliateId returns the AffiliateId field if non-nil, zero value otherwise.

### GetAffiliateIdOk

`func (o *TenantDto) GetAffiliateIdOk() (*string, bool)`

GetAffiliateIdOk returns a tuple with the AffiliateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateId

`func (o *TenantDto) SetAffiliateId(v string)`

SetAffiliateId sets AffiliateId field to given value.

### HasAffiliateId

`func (o *TenantDto) HasAffiliateId() bool`

HasAffiliateId returns a boolean if a field has been set.

### SetAffiliateIdNil

`func (o *TenantDto) SetAffiliateIdNil(b bool)`

 SetAffiliateIdNil sets the value for AffiliateId to be an explicit nil

### UnsetAffiliateId
`func (o *TenantDto) UnsetAffiliateId()`

UnsetAffiliateId ensures that no value is present for AffiliateId, not even an explicit nil
### GetTenantAlias

`func (o *TenantDto) GetTenantAlias() string`

GetTenantAlias returns the TenantAlias field if non-nil, zero value otherwise.

### GetTenantAliasOk

`func (o *TenantDto) GetTenantAliasOk() (*string, bool)`

GetTenantAliasOk returns a tuple with the TenantAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantAlias

`func (o *TenantDto) SetTenantAlias(v string)`

SetTenantAlias sets TenantAlias field to given value.

### HasTenantAlias

`func (o *TenantDto) HasTenantAlias() bool`

HasTenantAlias returns a boolean if a field has been set.

### SetTenantAliasNil

`func (o *TenantDto) SetTenantAliasNil(b bool)`

 SetTenantAliasNil sets the value for TenantAlias to be an explicit nil

### UnsetTenantAlias
`func (o *TenantDto) UnsetTenantAlias()`

UnsetTenantAlias ensures that no value is present for TenantAlias, not even an explicit nil
### GetCalls

`func (o *TenantDto) GetCalls() bool`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *TenantDto) GetCallsOk() (*bool, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *TenantDto) SetCalls(v bool)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *TenantDto) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetCampaign

`func (o *TenantDto) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *TenantDto) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *TenantDto) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *TenantDto) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### SetCampaignNil

`func (o *TenantDto) SetCampaignNil(b bool)`

 SetCampaignNil sets the value for Campaign to be an explicit nil

### UnsetCampaign
`func (o *TenantDto) UnsetCampaign()`

UnsetCampaign ensures that no value is present for Campaign, not even an explicit nil
### GetCreationDateTime

`func (o *TenantDto) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *TenantDto) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *TenantDto) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *TenantDto) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetHostedRegion

`func (o *TenantDto) GetHostedRegion() string`

GetHostedRegion returns the HostedRegion field if non-nil, zero value otherwise.

### GetHostedRegionOk

`func (o *TenantDto) GetHostedRegionOk() (*string, bool)`

GetHostedRegionOk returns a tuple with the HostedRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedRegion

`func (o *TenantDto) SetHostedRegion(v string)`

SetHostedRegion sets HostedRegion field to given value.

### HasHostedRegion

`func (o *TenantDto) HasHostedRegion() bool`

HasHostedRegion returns a boolean if a field has been set.

### SetHostedRegionNil

`func (o *TenantDto) SetHostedRegionNil(b bool)`

 SetHostedRegionNil sets the value for HostedRegion to be an explicit nil

### UnsetHostedRegion
`func (o *TenantDto) UnsetHostedRegion()`

UnsetHostedRegion ensures that no value is present for HostedRegion, not even an explicit nil
### GetTenantId

`func (o *TenantDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetIndustry

`func (o *TenantDto) GetIndustry() TenantIndustry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *TenantDto) GetIndustryOk() (*TenantIndustry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *TenantDto) SetIndustry(v TenantIndustry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *TenantDto) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetLanguage

`func (o *TenantDto) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TenantDto) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TenantDto) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *TenantDto) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *TenantDto) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *TenantDto) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetLastModified

`func (o *TenantDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMappedDomain

`func (o *TenantDto) GetMappedDomain() string`

GetMappedDomain returns the MappedDomain field if non-nil, zero value otherwise.

### GetMappedDomainOk

`func (o *TenantDto) GetMappedDomainOk() (*string, bool)`

GetMappedDomainOk returns a tuple with the MappedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedDomain

`func (o *TenantDto) SetMappedDomain(v string)`

SetMappedDomain sets MappedDomain field to given value.

### HasMappedDomain

`func (o *TenantDto) HasMappedDomain() bool`

HasMappedDomain returns a boolean if a field has been set.

### SetMappedDomainNil

`func (o *TenantDto) SetMappedDomainNil(b bool)`

 SetMappedDomainNil sets the value for MappedDomain to be an explicit nil

### UnsetMappedDomain
`func (o *TenantDto) UnsetMappedDomain()`

UnsetMappedDomain ensures that no value is present for MappedDomain, not even an explicit nil
### GetName

`func (o *TenantDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *TenantDto) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *TenantDto) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *TenantDto) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *TenantDto) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPaymentId

`func (o *TenantDto) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *TenantDto) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *TenantDto) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *TenantDto) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *TenantDto) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *TenantDto) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetSpam

`func (o *TenantDto) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *TenantDto) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *TenantDto) SetSpam(v bool)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *TenantDto) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetStatus

`func (o *TenantDto) GetStatus() TenantStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantDto) GetStatusOk() (*TenantStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantDto) SetStatus(v TenantStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TenantDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusChangeDate

`func (o *TenantDto) GetStatusChangeDate() time.Time`

GetStatusChangeDate returns the StatusChangeDate field if non-nil, zero value otherwise.

### GetStatusChangeDateOk

`func (o *TenantDto) GetStatusChangeDateOk() (*time.Time, bool)`

GetStatusChangeDateOk returns a tuple with the StatusChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangeDate

`func (o *TenantDto) SetStatusChangeDate(v time.Time)`

SetStatusChangeDate sets StatusChangeDate field to given value.

### HasStatusChangeDate

`func (o *TenantDto) HasStatusChangeDate() bool`

HasStatusChangeDate returns a boolean if a field has been set.

### GetTimeZone

`func (o *TenantDto) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TenantDto) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TenantDto) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TenantDto) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *TenantDto) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *TenantDto) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetTrustedDomains

`func (o *TenantDto) GetTrustedDomains() []string`

GetTrustedDomains returns the TrustedDomains field if non-nil, zero value otherwise.

### GetTrustedDomainsOk

`func (o *TenantDto) GetTrustedDomainsOk() (*[]string, bool)`

GetTrustedDomainsOk returns a tuple with the TrustedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomains

`func (o *TenantDto) SetTrustedDomains(v []string)`

SetTrustedDomains sets TrustedDomains field to given value.

### HasTrustedDomains

`func (o *TenantDto) HasTrustedDomains() bool`

HasTrustedDomains returns a boolean if a field has been set.

### SetTrustedDomainsNil

`func (o *TenantDto) SetTrustedDomainsNil(b bool)`

 SetTrustedDomainsNil sets the value for TrustedDomains to be an explicit nil

### UnsetTrustedDomains
`func (o *TenantDto) UnsetTrustedDomains()`

UnsetTrustedDomains ensures that no value is present for TrustedDomains, not even an explicit nil
### GetTrustedDomainsRaw

`func (o *TenantDto) GetTrustedDomainsRaw() string`

GetTrustedDomainsRaw returns the TrustedDomainsRaw field if non-nil, zero value otherwise.

### GetTrustedDomainsRawOk

`func (o *TenantDto) GetTrustedDomainsRawOk() (*string, bool)`

GetTrustedDomainsRawOk returns a tuple with the TrustedDomainsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainsRaw

`func (o *TenantDto) SetTrustedDomainsRaw(v string)`

SetTrustedDomainsRaw sets TrustedDomainsRaw field to given value.

### HasTrustedDomainsRaw

`func (o *TenantDto) HasTrustedDomainsRaw() bool`

HasTrustedDomainsRaw returns a boolean if a field has been set.

### SetTrustedDomainsRawNil

`func (o *TenantDto) SetTrustedDomainsRawNil(b bool)`

 SetTrustedDomainsRawNil sets the value for TrustedDomainsRaw to be an explicit nil

### UnsetTrustedDomainsRaw
`func (o *TenantDto) UnsetTrustedDomainsRaw()`

UnsetTrustedDomainsRaw ensures that no value is present for TrustedDomainsRaw, not even an explicit nil
### GetTrustedDomainsType

`func (o *TenantDto) GetTrustedDomainsType() TenantTrustedDomainsType`

GetTrustedDomainsType returns the TrustedDomainsType field if non-nil, zero value otherwise.

### GetTrustedDomainsTypeOk

`func (o *TenantDto) GetTrustedDomainsTypeOk() (*TenantTrustedDomainsType, bool)`

GetTrustedDomainsTypeOk returns a tuple with the TrustedDomainsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainsType

`func (o *TenantDto) SetTrustedDomainsType(v TenantTrustedDomainsType)`

SetTrustedDomainsType sets TrustedDomainsType field to given value.

### HasTrustedDomainsType

`func (o *TenantDto) HasTrustedDomainsType() bool`

HasTrustedDomainsType returns a boolean if a field has been set.

### GetVersion

`func (o *TenantDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TenantDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TenantDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TenantDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionChanged

`func (o *TenantDto) GetVersionChanged() time.Time`

GetVersionChanged returns the VersionChanged field if non-nil, zero value otherwise.

### GetVersionChangedOk

`func (o *TenantDto) GetVersionChangedOk() (*time.Time, bool)`

GetVersionChangedOk returns a tuple with the VersionChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionChanged

`func (o *TenantDto) SetVersionChanged(v time.Time)`

SetVersionChanged sets VersionChanged field to given value.

### HasVersionChanged

`func (o *TenantDto) HasVersionChanged() bool`

HasVersionChanged returns a boolean if a field has been set.

### GetRegion

`func (o *TenantDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TenantDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TenantDto) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TenantDto) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *TenantDto) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *TenantDto) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


