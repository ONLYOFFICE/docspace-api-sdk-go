# DbTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The tenant ID. | [optional] 
**Name** | Pointer to **NullableString** | The tenant name. | [optional] 
**Alias** | Pointer to **NullableString** | The tenant alias. | [optional] 
**MappedDomain** | Pointer to **NullableString** | Mapped domain | [optional] 
**Version** | Pointer to **int32** | The tenant version. | [optional] 
**VersionChangedField** | Pointer to **NullableTime** | The Version_changed field. | [optional] 
**VersionChanged** | Pointer to **time.Time** | The date and time when the version was changed. | [optional] 
**Language** | Pointer to **NullableString** | The tenant language. | [optional] 
**TimeZone** | Pointer to **NullableString** | The tenant time zone. | [optional] 
**TrustedDomainsRaw** | Pointer to **NullableString** | The tenant trusted domains raw. | [optional] 
**TrustedDomainsEnabled** | Pointer to [**TenantTrustedDomainsType**](TenantTrustedDomainsType.md) | The type of the tenant trusted domains. | [optional] 
**Status** | Pointer to [**TenantStatus**](TenantStatus.md) | The tenant status. | [optional] 
**StatusChanged** | Pointer to **NullableTime** | The date and time when the tenant status was changed. | [optional] 
**StatusChangedHack** | Pointer to **time.Time** | The hacked date and time when the tenant status was changed. | [optional] 
**CreationDateTime** | Pointer to **time.Time** | The tenant creation date. | [optional] 
**OwnerId** | Pointer to **NullableString** | The tenant owner ID. | [optional] 
**PaymentId** | Pointer to **NullableString** | The tenant payment ID. | [optional] 
**Industry** | Pointer to [**TenantIndustry**](TenantIndustry.md) | The tenant industry. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the tenant was last modified. | [optional] 
**Calls** | Pointer to **bool** | Specifies if the calls are available for the current tenant or not. | [optional] 
**Partner** | Pointer to [**DbTenantPartner**](DbTenantPartner.md) | The database tenant partner parameters. | [optional] 

## Methods

### NewDbTenant

`func NewDbTenant() *DbTenant`

NewDbTenant instantiates a new DbTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbTenantWithDefaults

`func NewDbTenantWithDefaults() *DbTenant`

NewDbTenantWithDefaults instantiates a new DbTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbTenant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbTenant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbTenant) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DbTenant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DbTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbTenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbTenant) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DbTenant) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DbTenant) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAlias

`func (o *DbTenant) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DbTenant) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DbTenant) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DbTenant) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *DbTenant) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *DbTenant) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetMappedDomain

`func (o *DbTenant) GetMappedDomain() string`

GetMappedDomain returns the MappedDomain field if non-nil, zero value otherwise.

### GetMappedDomainOk

`func (o *DbTenant) GetMappedDomainOk() (*string, bool)`

GetMappedDomainOk returns a tuple with the MappedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedDomain

`func (o *DbTenant) SetMappedDomain(v string)`

SetMappedDomain sets MappedDomain field to given value.

### HasMappedDomain

`func (o *DbTenant) HasMappedDomain() bool`

HasMappedDomain returns a boolean if a field has been set.

### SetMappedDomainNil

`func (o *DbTenant) SetMappedDomainNil(b bool)`

 SetMappedDomainNil sets the value for MappedDomain to be an explicit nil

### UnsetMappedDomain
`func (o *DbTenant) UnsetMappedDomain()`

UnsetMappedDomain ensures that no value is present for MappedDomain, not even an explicit nil
### GetVersion

`func (o *DbTenant) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DbTenant) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DbTenant) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DbTenant) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionChangedField

`func (o *DbTenant) GetVersionChangedField() time.Time`

GetVersionChangedField returns the VersionChangedField field if non-nil, zero value otherwise.

### GetVersionChangedFieldOk

`func (o *DbTenant) GetVersionChangedFieldOk() (*time.Time, bool)`

GetVersionChangedFieldOk returns a tuple with the VersionChangedField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionChangedField

`func (o *DbTenant) SetVersionChangedField(v time.Time)`

SetVersionChangedField sets VersionChangedField field to given value.

### HasVersionChangedField

`func (o *DbTenant) HasVersionChangedField() bool`

HasVersionChangedField returns a boolean if a field has been set.

### SetVersionChangedFieldNil

`func (o *DbTenant) SetVersionChangedFieldNil(b bool)`

 SetVersionChangedFieldNil sets the value for VersionChangedField to be an explicit nil

### UnsetVersionChangedField
`func (o *DbTenant) UnsetVersionChangedField()`

UnsetVersionChangedField ensures that no value is present for VersionChangedField, not even an explicit nil
### GetVersionChanged

`func (o *DbTenant) GetVersionChanged() time.Time`

GetVersionChanged returns the VersionChanged field if non-nil, zero value otherwise.

### GetVersionChangedOk

`func (o *DbTenant) GetVersionChangedOk() (*time.Time, bool)`

GetVersionChangedOk returns a tuple with the VersionChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionChanged

`func (o *DbTenant) SetVersionChanged(v time.Time)`

SetVersionChanged sets VersionChanged field to given value.

### HasVersionChanged

`func (o *DbTenant) HasVersionChanged() bool`

HasVersionChanged returns a boolean if a field has been set.

### GetLanguage

`func (o *DbTenant) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *DbTenant) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *DbTenant) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *DbTenant) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *DbTenant) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *DbTenant) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetTimeZone

`func (o *DbTenant) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *DbTenant) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *DbTenant) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *DbTenant) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *DbTenant) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *DbTenant) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetTrustedDomainsRaw

`func (o *DbTenant) GetTrustedDomainsRaw() string`

GetTrustedDomainsRaw returns the TrustedDomainsRaw field if non-nil, zero value otherwise.

### GetTrustedDomainsRawOk

`func (o *DbTenant) GetTrustedDomainsRawOk() (*string, bool)`

GetTrustedDomainsRawOk returns a tuple with the TrustedDomainsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainsRaw

`func (o *DbTenant) SetTrustedDomainsRaw(v string)`

SetTrustedDomainsRaw sets TrustedDomainsRaw field to given value.

### HasTrustedDomainsRaw

`func (o *DbTenant) HasTrustedDomainsRaw() bool`

HasTrustedDomainsRaw returns a boolean if a field has been set.

### SetTrustedDomainsRawNil

`func (o *DbTenant) SetTrustedDomainsRawNil(b bool)`

 SetTrustedDomainsRawNil sets the value for TrustedDomainsRaw to be an explicit nil

### UnsetTrustedDomainsRaw
`func (o *DbTenant) UnsetTrustedDomainsRaw()`

UnsetTrustedDomainsRaw ensures that no value is present for TrustedDomainsRaw, not even an explicit nil
### GetTrustedDomainsEnabled

`func (o *DbTenant) GetTrustedDomainsEnabled() TenantTrustedDomainsType`

GetTrustedDomainsEnabled returns the TrustedDomainsEnabled field if non-nil, zero value otherwise.

### GetTrustedDomainsEnabledOk

`func (o *DbTenant) GetTrustedDomainsEnabledOk() (*TenantTrustedDomainsType, bool)`

GetTrustedDomainsEnabledOk returns a tuple with the TrustedDomainsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainsEnabled

`func (o *DbTenant) SetTrustedDomainsEnabled(v TenantTrustedDomainsType)`

SetTrustedDomainsEnabled sets TrustedDomainsEnabled field to given value.

### HasTrustedDomainsEnabled

`func (o *DbTenant) HasTrustedDomainsEnabled() bool`

HasTrustedDomainsEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *DbTenant) GetStatus() TenantStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DbTenant) GetStatusOk() (*TenantStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DbTenant) SetStatus(v TenantStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DbTenant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusChanged

`func (o *DbTenant) GetStatusChanged() time.Time`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *DbTenant) GetStatusChangedOk() (*time.Time, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *DbTenant) SetStatusChanged(v time.Time)`

SetStatusChanged sets StatusChanged field to given value.

### HasStatusChanged

`func (o *DbTenant) HasStatusChanged() bool`

HasStatusChanged returns a boolean if a field has been set.

### SetStatusChangedNil

`func (o *DbTenant) SetStatusChangedNil(b bool)`

 SetStatusChangedNil sets the value for StatusChanged to be an explicit nil

### UnsetStatusChanged
`func (o *DbTenant) UnsetStatusChanged()`

UnsetStatusChanged ensures that no value is present for StatusChanged, not even an explicit nil
### GetStatusChangedHack

`func (o *DbTenant) GetStatusChangedHack() time.Time`

GetStatusChangedHack returns the StatusChangedHack field if non-nil, zero value otherwise.

### GetStatusChangedHackOk

`func (o *DbTenant) GetStatusChangedHackOk() (*time.Time, bool)`

GetStatusChangedHackOk returns a tuple with the StatusChangedHack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangedHack

`func (o *DbTenant) SetStatusChangedHack(v time.Time)`

SetStatusChangedHack sets StatusChangedHack field to given value.

### HasStatusChangedHack

`func (o *DbTenant) HasStatusChangedHack() bool`

HasStatusChangedHack returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *DbTenant) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *DbTenant) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *DbTenant) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *DbTenant) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetOwnerId

`func (o *DbTenant) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DbTenant) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DbTenant) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DbTenant) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *DbTenant) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *DbTenant) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetPaymentId

`func (o *DbTenant) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *DbTenant) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *DbTenant) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *DbTenant) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *DbTenant) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *DbTenant) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetIndustry

`func (o *DbTenant) GetIndustry() TenantIndustry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *DbTenant) GetIndustryOk() (*TenantIndustry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *DbTenant) SetIndustry(v TenantIndustry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *DbTenant) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetLastModified

`func (o *DbTenant) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DbTenant) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DbTenant) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *DbTenant) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetCalls

`func (o *DbTenant) GetCalls() bool`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *DbTenant) GetCallsOk() (*bool, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *DbTenant) SetCalls(v bool)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *DbTenant) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetPartner

`func (o *DbTenant) GetPartner() DbTenantPartner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *DbTenant) GetPartnerOk() (*DbTenantPartner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *DbTenant) SetPartner(v DbTenantPartner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *DbTenant) HasPartner() bool`

HasPartner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


