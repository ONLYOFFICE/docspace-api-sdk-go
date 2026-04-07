# TenantQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **int32** | The tenant ID. | [optional] 
**Name** | Pointer to **NullableString** | The tenant name. | [optional] 
**Price** | Pointer to **float64** | The tenant price. | [optional] 
**PriceCurrencySymbol** | Pointer to **NullableString** | The tenant price currency symbol. | [optional] 
**PriceISOCurrencySymbol** | Pointer to **NullableString** | The tenant price three-character ISO 4217 currency symbol. | [optional] 
**ProductId** | Pointer to **NullableString** | The tenant product ID. | [optional] 
**ServiceName** | Pointer to **NullableString** | The service name. | [optional] 
**ServiceGroup** | Pointer to **NullableString** | The service group. | [optional] 
**Visible** | Pointer to **bool** | Specifies if the tenant quota is visible or not. | [optional] 
**Wallet** | Pointer to **bool** | Specifies if the tenant quota applies to the wallet or not | [optional] 
**DueDate** | Pointer to **NullableTime** | The quota due date. | [optional] 
**Features** | Pointer to **NullableString** | The tenant quota features. | [optional] 
**MaxFileSize** | Pointer to **int64** | The tenant maximum file size. | [optional] 
**MaxTotalSize** | Pointer to **int64** | The tenant maximum total size. | [optional] 
**CountUser** | Pointer to **int32** | The number of portal users. | [optional] 
**CountRoomAdmin** | Pointer to **int32** | The number of portal room administrators. | [optional] 
**UsersInRoom** | Pointer to **int32** | The number of room users. | [optional] 
**CountRoom** | Pointer to **int32** | The number of rooms. | [optional] 
**NonProfit** | Pointer to **bool** | Specifies if the tenant quota is nonprofit or not. | [optional] 
**Trial** | Pointer to **bool** | Specifies if the tenant quota is trial or not. | [optional] 
**Free** | Pointer to **bool** | Specifies if the tenant quota is free or not. | [optional] 
**Update** | Pointer to **bool** | Specifies if the tenant quota is updated or not. | [optional] 
**Audit** | Pointer to **bool** | Specifies if the audit trail is available or not. | [optional] 
**DocsEdition** | Pointer to **bool** | Specifies if ONLYOFFICE Docs is included in the tenant quota or not. | [optional] 
**Ldap** | Pointer to **bool** | Specifies if the LDAP settings are available or not. | [optional] 
**Sso** | Pointer to **bool** | Specifies if the SSO settings are available or not. | [optional] 
**Statistic** | Pointer to **bool** | Specifies if the statistics settings are available or not. | [optional] 
**Branding** | Pointer to **bool** | Specifies if the branding settings are available or not. | [optional] 
**Customization** | Pointer to **bool** | Specifies if the customization settings are available or not. | [optional] 
**Lifetime** | Pointer to **bool** | Specifies if the license has the lifetime settings or not. | [optional] 
**AutomationApi** | Pointer to **bool** | Specifies if the Automation API is available or not. | [optional] 
**Custom** | Pointer to **bool** | Specifies if the custom domain URL is available or not. | [optional] 
**Restore** | Pointer to **bool** | Specifies if the restore is enabled or not. | [optional] 
**Oauth** | Pointer to **bool** | Specifies if Oauth is available or not. | [optional] 
**ContentSearch** | Pointer to **bool** | Specifies if the content search is available or not. | [optional] 
**ThirdParty** | Pointer to **bool** | Specifies if the third-party accounts linking is available or not. | [optional] 
**Year** | Pointer to **bool** | Specifies if the tenant quota is yearly subscription or not. | [optional] 
**CountFreeBackup** | Pointer to **int32** | The number of free backups within a month. | [optional] 
**Backup** | Pointer to **bool** | Specifies if the backup enabled as a wallet service or not. | [optional] 
**CountAIAgent** | Pointer to **int32** | The number of AI agents. | [optional] 
**AiTools** | Pointer to **bool** | Specifies if the AI tools enabled as a wallet service or not. | [optional] 

## Methods

### NewTenantQuota

`func NewTenantQuota() *TenantQuota`

NewTenantQuota instantiates a new TenantQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantQuotaWithDefaults

`func NewTenantQuotaWithDefaults() *TenantQuota`

NewTenantQuotaWithDefaults instantiates a new TenantQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantQuota) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantQuota) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantQuota) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantQuota) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *TenantQuota) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantQuota) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantQuota) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantQuota) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantQuota) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantQuota) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrice

`func (o *TenantQuota) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TenantQuota) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TenantQuota) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TenantQuota) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceCurrencySymbol

`func (o *TenantQuota) GetPriceCurrencySymbol() string`

GetPriceCurrencySymbol returns the PriceCurrencySymbol field if non-nil, zero value otherwise.

### GetPriceCurrencySymbolOk

`func (o *TenantQuota) GetPriceCurrencySymbolOk() (*string, bool)`

GetPriceCurrencySymbolOk returns a tuple with the PriceCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCurrencySymbol

`func (o *TenantQuota) SetPriceCurrencySymbol(v string)`

SetPriceCurrencySymbol sets PriceCurrencySymbol field to given value.

### HasPriceCurrencySymbol

`func (o *TenantQuota) HasPriceCurrencySymbol() bool`

HasPriceCurrencySymbol returns a boolean if a field has been set.

### SetPriceCurrencySymbolNil

`func (o *TenantQuota) SetPriceCurrencySymbolNil(b bool)`

 SetPriceCurrencySymbolNil sets the value for PriceCurrencySymbol to be an explicit nil

### UnsetPriceCurrencySymbol
`func (o *TenantQuota) UnsetPriceCurrencySymbol()`

UnsetPriceCurrencySymbol ensures that no value is present for PriceCurrencySymbol, not even an explicit nil
### GetPriceISOCurrencySymbol

`func (o *TenantQuota) GetPriceISOCurrencySymbol() string`

GetPriceISOCurrencySymbol returns the PriceISOCurrencySymbol field if non-nil, zero value otherwise.

### GetPriceISOCurrencySymbolOk

`func (o *TenantQuota) GetPriceISOCurrencySymbolOk() (*string, bool)`

GetPriceISOCurrencySymbolOk returns a tuple with the PriceISOCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceISOCurrencySymbol

`func (o *TenantQuota) SetPriceISOCurrencySymbol(v string)`

SetPriceISOCurrencySymbol sets PriceISOCurrencySymbol field to given value.

### HasPriceISOCurrencySymbol

`func (o *TenantQuota) HasPriceISOCurrencySymbol() bool`

HasPriceISOCurrencySymbol returns a boolean if a field has been set.

### SetPriceISOCurrencySymbolNil

`func (o *TenantQuota) SetPriceISOCurrencySymbolNil(b bool)`

 SetPriceISOCurrencySymbolNil sets the value for PriceISOCurrencySymbol to be an explicit nil

### UnsetPriceISOCurrencySymbol
`func (o *TenantQuota) UnsetPriceISOCurrencySymbol()`

UnsetPriceISOCurrencySymbol ensures that no value is present for PriceISOCurrencySymbol, not even an explicit nil
### GetProductId

`func (o *TenantQuota) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *TenantQuota) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *TenantQuota) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *TenantQuota) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *TenantQuota) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *TenantQuota) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetServiceName

`func (o *TenantQuota) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *TenantQuota) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *TenantQuota) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *TenantQuota) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *TenantQuota) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *TenantQuota) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetServiceGroup

`func (o *TenantQuota) GetServiceGroup() string`

GetServiceGroup returns the ServiceGroup field if non-nil, zero value otherwise.

### GetServiceGroupOk

`func (o *TenantQuota) GetServiceGroupOk() (*string, bool)`

GetServiceGroupOk returns a tuple with the ServiceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGroup

`func (o *TenantQuota) SetServiceGroup(v string)`

SetServiceGroup sets ServiceGroup field to given value.

### HasServiceGroup

`func (o *TenantQuota) HasServiceGroup() bool`

HasServiceGroup returns a boolean if a field has been set.

### SetServiceGroupNil

`func (o *TenantQuota) SetServiceGroupNil(b bool)`

 SetServiceGroupNil sets the value for ServiceGroup to be an explicit nil

### UnsetServiceGroup
`func (o *TenantQuota) UnsetServiceGroup()`

UnsetServiceGroup ensures that no value is present for ServiceGroup, not even an explicit nil
### GetVisible

`func (o *TenantQuota) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *TenantQuota) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *TenantQuota) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *TenantQuota) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetWallet

`func (o *TenantQuota) GetWallet() bool`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *TenantQuota) GetWalletOk() (*bool, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *TenantQuota) SetWallet(v bool)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *TenantQuota) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetDueDate

`func (o *TenantQuota) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TenantQuota) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TenantQuota) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TenantQuota) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *TenantQuota) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *TenantQuota) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetFeatures

`func (o *TenantQuota) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TenantQuota) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TenantQuota) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *TenantQuota) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *TenantQuota) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *TenantQuota) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetMaxFileSize

`func (o *TenantQuota) GetMaxFileSize() int64`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *TenantQuota) GetMaxFileSizeOk() (*int64, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *TenantQuota) SetMaxFileSize(v int64)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *TenantQuota) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMaxTotalSize

`func (o *TenantQuota) GetMaxTotalSize() int64`

GetMaxTotalSize returns the MaxTotalSize field if non-nil, zero value otherwise.

### GetMaxTotalSizeOk

`func (o *TenantQuota) GetMaxTotalSizeOk() (*int64, bool)`

GetMaxTotalSizeOk returns a tuple with the MaxTotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalSize

`func (o *TenantQuota) SetMaxTotalSize(v int64)`

SetMaxTotalSize sets MaxTotalSize field to given value.

### HasMaxTotalSize

`func (o *TenantQuota) HasMaxTotalSize() bool`

HasMaxTotalSize returns a boolean if a field has been set.

### GetCountUser

`func (o *TenantQuota) GetCountUser() int32`

GetCountUser returns the CountUser field if non-nil, zero value otherwise.

### GetCountUserOk

`func (o *TenantQuota) GetCountUserOk() (*int32, bool)`

GetCountUserOk returns a tuple with the CountUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountUser

`func (o *TenantQuota) SetCountUser(v int32)`

SetCountUser sets CountUser field to given value.

### HasCountUser

`func (o *TenantQuota) HasCountUser() bool`

HasCountUser returns a boolean if a field has been set.

### GetCountRoomAdmin

`func (o *TenantQuota) GetCountRoomAdmin() int32`

GetCountRoomAdmin returns the CountRoomAdmin field if non-nil, zero value otherwise.

### GetCountRoomAdminOk

`func (o *TenantQuota) GetCountRoomAdminOk() (*int32, bool)`

GetCountRoomAdminOk returns a tuple with the CountRoomAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountRoomAdmin

`func (o *TenantQuota) SetCountRoomAdmin(v int32)`

SetCountRoomAdmin sets CountRoomAdmin field to given value.

### HasCountRoomAdmin

`func (o *TenantQuota) HasCountRoomAdmin() bool`

HasCountRoomAdmin returns a boolean if a field has been set.

### GetUsersInRoom

`func (o *TenantQuota) GetUsersInRoom() int32`

GetUsersInRoom returns the UsersInRoom field if non-nil, zero value otherwise.

### GetUsersInRoomOk

`func (o *TenantQuota) GetUsersInRoomOk() (*int32, bool)`

GetUsersInRoomOk returns a tuple with the UsersInRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersInRoom

`func (o *TenantQuota) SetUsersInRoom(v int32)`

SetUsersInRoom sets UsersInRoom field to given value.

### HasUsersInRoom

`func (o *TenantQuota) HasUsersInRoom() bool`

HasUsersInRoom returns a boolean if a field has been set.

### GetCountRoom

`func (o *TenantQuota) GetCountRoom() int32`

GetCountRoom returns the CountRoom field if non-nil, zero value otherwise.

### GetCountRoomOk

`func (o *TenantQuota) GetCountRoomOk() (*int32, bool)`

GetCountRoomOk returns a tuple with the CountRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountRoom

`func (o *TenantQuota) SetCountRoom(v int32)`

SetCountRoom sets CountRoom field to given value.

### HasCountRoom

`func (o *TenantQuota) HasCountRoom() bool`

HasCountRoom returns a boolean if a field has been set.

### GetNonProfit

`func (o *TenantQuota) GetNonProfit() bool`

GetNonProfit returns the NonProfit field if non-nil, zero value otherwise.

### GetNonProfitOk

`func (o *TenantQuota) GetNonProfitOk() (*bool, bool)`

GetNonProfitOk returns a tuple with the NonProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonProfit

`func (o *TenantQuota) SetNonProfit(v bool)`

SetNonProfit sets NonProfit field to given value.

### HasNonProfit

`func (o *TenantQuota) HasNonProfit() bool`

HasNonProfit returns a boolean if a field has been set.

### GetTrial

`func (o *TenantQuota) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *TenantQuota) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *TenantQuota) SetTrial(v bool)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *TenantQuota) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetFree

`func (o *TenantQuota) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *TenantQuota) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *TenantQuota) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *TenantQuota) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetUpdate

`func (o *TenantQuota) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *TenantQuota) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *TenantQuota) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *TenantQuota) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAudit

`func (o *TenantQuota) GetAudit() bool`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *TenantQuota) GetAuditOk() (*bool, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *TenantQuota) SetAudit(v bool)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *TenantQuota) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetDocsEdition

`func (o *TenantQuota) GetDocsEdition() bool`

GetDocsEdition returns the DocsEdition field if non-nil, zero value otherwise.

### GetDocsEditionOk

`func (o *TenantQuota) GetDocsEditionOk() (*bool, bool)`

GetDocsEditionOk returns a tuple with the DocsEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsEdition

`func (o *TenantQuota) SetDocsEdition(v bool)`

SetDocsEdition sets DocsEdition field to given value.

### HasDocsEdition

`func (o *TenantQuota) HasDocsEdition() bool`

HasDocsEdition returns a boolean if a field has been set.

### GetLdap

`func (o *TenantQuota) GetLdap() bool`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *TenantQuota) GetLdapOk() (*bool, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *TenantQuota) SetLdap(v bool)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *TenantQuota) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetSso

`func (o *TenantQuota) GetSso() bool`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *TenantQuota) GetSsoOk() (*bool, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *TenantQuota) SetSso(v bool)`

SetSso sets Sso field to given value.

### HasSso

`func (o *TenantQuota) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetStatistic

`func (o *TenantQuota) GetStatistic() bool`

GetStatistic returns the Statistic field if non-nil, zero value otherwise.

### GetStatisticOk

`func (o *TenantQuota) GetStatisticOk() (*bool, bool)`

GetStatisticOk returns a tuple with the Statistic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistic

`func (o *TenantQuota) SetStatistic(v bool)`

SetStatistic sets Statistic field to given value.

### HasStatistic

`func (o *TenantQuota) HasStatistic() bool`

HasStatistic returns a boolean if a field has been set.

### GetBranding

`func (o *TenantQuota) GetBranding() bool`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *TenantQuota) GetBrandingOk() (*bool, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *TenantQuota) SetBranding(v bool)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *TenantQuota) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetCustomization

`func (o *TenantQuota) GetCustomization() bool`

GetCustomization returns the Customization field if non-nil, zero value otherwise.

### GetCustomizationOk

`func (o *TenantQuota) GetCustomizationOk() (*bool, bool)`

GetCustomizationOk returns a tuple with the Customization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomization

`func (o *TenantQuota) SetCustomization(v bool)`

SetCustomization sets Customization field to given value.

### HasCustomization

`func (o *TenantQuota) HasCustomization() bool`

HasCustomization returns a boolean if a field has been set.

### GetLifetime

`func (o *TenantQuota) GetLifetime() bool`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *TenantQuota) GetLifetimeOk() (*bool, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *TenantQuota) SetLifetime(v bool)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *TenantQuota) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetAutomationApi

`func (o *TenantQuota) GetAutomationApi() bool`

GetAutomationApi returns the AutomationApi field if non-nil, zero value otherwise.

### GetAutomationApiOk

`func (o *TenantQuota) GetAutomationApiOk() (*bool, bool)`

GetAutomationApiOk returns a tuple with the AutomationApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationApi

`func (o *TenantQuota) SetAutomationApi(v bool)`

SetAutomationApi sets AutomationApi field to given value.

### HasAutomationApi

`func (o *TenantQuota) HasAutomationApi() bool`

HasAutomationApi returns a boolean if a field has been set.

### GetCustom

`func (o *TenantQuota) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *TenantQuota) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *TenantQuota) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *TenantQuota) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetRestore

`func (o *TenantQuota) GetRestore() bool`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *TenantQuota) GetRestoreOk() (*bool, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *TenantQuota) SetRestore(v bool)`

SetRestore sets Restore field to given value.

### HasRestore

`func (o *TenantQuota) HasRestore() bool`

HasRestore returns a boolean if a field has been set.

### GetOauth

`func (o *TenantQuota) GetOauth() bool`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *TenantQuota) GetOauthOk() (*bool, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *TenantQuota) SetOauth(v bool)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *TenantQuota) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetContentSearch

`func (o *TenantQuota) GetContentSearch() bool`

GetContentSearch returns the ContentSearch field if non-nil, zero value otherwise.

### GetContentSearchOk

`func (o *TenantQuota) GetContentSearchOk() (*bool, bool)`

GetContentSearchOk returns a tuple with the ContentSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSearch

`func (o *TenantQuota) SetContentSearch(v bool)`

SetContentSearch sets ContentSearch field to given value.

### HasContentSearch

`func (o *TenantQuota) HasContentSearch() bool`

HasContentSearch returns a boolean if a field has been set.

### GetThirdParty

`func (o *TenantQuota) GetThirdParty() bool`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *TenantQuota) GetThirdPartyOk() (*bool, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *TenantQuota) SetThirdParty(v bool)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *TenantQuota) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.

### GetYear

`func (o *TenantQuota) GetYear() bool`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TenantQuota) GetYearOk() (*bool, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TenantQuota) SetYear(v bool)`

SetYear sets Year field to given value.

### HasYear

`func (o *TenantQuota) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetCountFreeBackup

`func (o *TenantQuota) GetCountFreeBackup() int32`

GetCountFreeBackup returns the CountFreeBackup field if non-nil, zero value otherwise.

### GetCountFreeBackupOk

`func (o *TenantQuota) GetCountFreeBackupOk() (*int32, bool)`

GetCountFreeBackupOk returns a tuple with the CountFreeBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountFreeBackup

`func (o *TenantQuota) SetCountFreeBackup(v int32)`

SetCountFreeBackup sets CountFreeBackup field to given value.

### HasCountFreeBackup

`func (o *TenantQuota) HasCountFreeBackup() bool`

HasCountFreeBackup returns a boolean if a field has been set.

### GetBackup

`func (o *TenantQuota) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *TenantQuota) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *TenantQuota) SetBackup(v bool)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *TenantQuota) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetCountAIAgent

`func (o *TenantQuota) GetCountAIAgent() int32`

GetCountAIAgent returns the CountAIAgent field if non-nil, zero value otherwise.

### GetCountAIAgentOk

`func (o *TenantQuota) GetCountAIAgentOk() (*int32, bool)`

GetCountAIAgentOk returns a tuple with the CountAIAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAIAgent

`func (o *TenantQuota) SetCountAIAgent(v int32)`

SetCountAIAgent sets CountAIAgent field to given value.

### HasCountAIAgent

`func (o *TenantQuota) HasCountAIAgent() bool`

HasCountAIAgent returns a boolean if a field has been set.

### GetAiTools

`func (o *TenantQuota) GetAiTools() bool`

GetAiTools returns the AiTools field if non-nil, zero value otherwise.

### GetAiToolsOk

`func (o *TenantQuota) GetAiToolsOk() (*bool, bool)`

GetAiToolsOk returns a tuple with the AiTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiTools

`func (o *TenantQuota) SetAiTools(v bool)`

SetAiTools sets AiTools field to given value.

### HasAiTools

`func (o *TenantQuota) HasAiTools() bool`

HasAiTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


