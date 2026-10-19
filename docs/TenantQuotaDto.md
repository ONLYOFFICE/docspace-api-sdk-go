# TenantQuotaDto

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
**Additional** | Pointer to **bool** | Specifies if the tenant quota is primary or additional. | [optional] 
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
**AiSearch** | Pointer to **bool** | Specifies if the AI search enabled as a wallet service or not. | [optional] 
**DocsCloud** | Pointer to **int32** | The number of Docs Connect users. | [optional] 
**DocsCloudDevPack** | Pointer to **bool** | Specifies if the Docs Connect Dev Pack enabled or not. | [optional] 
**DocsCloudTrial** | Pointer to **bool** | Specifies if the Docs Connect trial enabled or not. | [optional] 

## Methods

### NewTenantQuotaDto

`func NewTenantQuotaDto() *TenantQuotaDto`

NewTenantQuotaDto instantiates a new TenantQuotaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantQuotaDtoWithDefaults

`func NewTenantQuotaDtoWithDefaults() *TenantQuotaDto`

NewTenantQuotaDtoWithDefaults instantiates a new TenantQuotaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantQuotaDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantQuotaDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantQuotaDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantQuotaDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *TenantQuotaDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantQuotaDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantQuotaDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantQuotaDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantQuotaDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantQuotaDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrice

`func (o *TenantQuotaDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TenantQuotaDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TenantQuotaDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TenantQuotaDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceCurrencySymbol

`func (o *TenantQuotaDto) GetPriceCurrencySymbol() string`

GetPriceCurrencySymbol returns the PriceCurrencySymbol field if non-nil, zero value otherwise.

### GetPriceCurrencySymbolOk

`func (o *TenantQuotaDto) GetPriceCurrencySymbolOk() (*string, bool)`

GetPriceCurrencySymbolOk returns a tuple with the PriceCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCurrencySymbol

`func (o *TenantQuotaDto) SetPriceCurrencySymbol(v string)`

SetPriceCurrencySymbol sets PriceCurrencySymbol field to given value.

### HasPriceCurrencySymbol

`func (o *TenantQuotaDto) HasPriceCurrencySymbol() bool`

HasPriceCurrencySymbol returns a boolean if a field has been set.

### SetPriceCurrencySymbolNil

`func (o *TenantQuotaDto) SetPriceCurrencySymbolNil(b bool)`

 SetPriceCurrencySymbolNil sets the value for PriceCurrencySymbol to be an explicit nil

### UnsetPriceCurrencySymbol
`func (o *TenantQuotaDto) UnsetPriceCurrencySymbol()`

UnsetPriceCurrencySymbol ensures that no value is present for PriceCurrencySymbol, not even an explicit nil
### GetPriceISOCurrencySymbol

`func (o *TenantQuotaDto) GetPriceISOCurrencySymbol() string`

GetPriceISOCurrencySymbol returns the PriceISOCurrencySymbol field if non-nil, zero value otherwise.

### GetPriceISOCurrencySymbolOk

`func (o *TenantQuotaDto) GetPriceISOCurrencySymbolOk() (*string, bool)`

GetPriceISOCurrencySymbolOk returns a tuple with the PriceISOCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceISOCurrencySymbol

`func (o *TenantQuotaDto) SetPriceISOCurrencySymbol(v string)`

SetPriceISOCurrencySymbol sets PriceISOCurrencySymbol field to given value.

### HasPriceISOCurrencySymbol

`func (o *TenantQuotaDto) HasPriceISOCurrencySymbol() bool`

HasPriceISOCurrencySymbol returns a boolean if a field has been set.

### SetPriceISOCurrencySymbolNil

`func (o *TenantQuotaDto) SetPriceISOCurrencySymbolNil(b bool)`

 SetPriceISOCurrencySymbolNil sets the value for PriceISOCurrencySymbol to be an explicit nil

### UnsetPriceISOCurrencySymbol
`func (o *TenantQuotaDto) UnsetPriceISOCurrencySymbol()`

UnsetPriceISOCurrencySymbol ensures that no value is present for PriceISOCurrencySymbol, not even an explicit nil
### GetProductId

`func (o *TenantQuotaDto) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *TenantQuotaDto) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *TenantQuotaDto) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *TenantQuotaDto) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *TenantQuotaDto) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *TenantQuotaDto) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetServiceName

`func (o *TenantQuotaDto) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *TenantQuotaDto) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *TenantQuotaDto) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *TenantQuotaDto) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *TenantQuotaDto) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *TenantQuotaDto) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetServiceGroup

`func (o *TenantQuotaDto) GetServiceGroup() string`

GetServiceGroup returns the ServiceGroup field if non-nil, zero value otherwise.

### GetServiceGroupOk

`func (o *TenantQuotaDto) GetServiceGroupOk() (*string, bool)`

GetServiceGroupOk returns a tuple with the ServiceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGroup

`func (o *TenantQuotaDto) SetServiceGroup(v string)`

SetServiceGroup sets ServiceGroup field to given value.

### HasServiceGroup

`func (o *TenantQuotaDto) HasServiceGroup() bool`

HasServiceGroup returns a boolean if a field has been set.

### SetServiceGroupNil

`func (o *TenantQuotaDto) SetServiceGroupNil(b bool)`

 SetServiceGroupNil sets the value for ServiceGroup to be an explicit nil

### UnsetServiceGroup
`func (o *TenantQuotaDto) UnsetServiceGroup()`

UnsetServiceGroup ensures that no value is present for ServiceGroup, not even an explicit nil
### GetVisible

`func (o *TenantQuotaDto) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *TenantQuotaDto) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *TenantQuotaDto) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *TenantQuotaDto) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetWallet

`func (o *TenantQuotaDto) GetWallet() bool`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *TenantQuotaDto) GetWalletOk() (*bool, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *TenantQuotaDto) SetWallet(v bool)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *TenantQuotaDto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetAdditional

`func (o *TenantQuotaDto) GetAdditional() bool`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *TenantQuotaDto) GetAdditionalOk() (*bool, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *TenantQuotaDto) SetAdditional(v bool)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *TenantQuotaDto) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.

### GetDueDate

`func (o *TenantQuotaDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TenantQuotaDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TenantQuotaDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TenantQuotaDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *TenantQuotaDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *TenantQuotaDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetFeatures

`func (o *TenantQuotaDto) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TenantQuotaDto) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TenantQuotaDto) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *TenantQuotaDto) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *TenantQuotaDto) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *TenantQuotaDto) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetMaxFileSize

`func (o *TenantQuotaDto) GetMaxFileSize() int64`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *TenantQuotaDto) GetMaxFileSizeOk() (*int64, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *TenantQuotaDto) SetMaxFileSize(v int64)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *TenantQuotaDto) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMaxTotalSize

`func (o *TenantQuotaDto) GetMaxTotalSize() int64`

GetMaxTotalSize returns the MaxTotalSize field if non-nil, zero value otherwise.

### GetMaxTotalSizeOk

`func (o *TenantQuotaDto) GetMaxTotalSizeOk() (*int64, bool)`

GetMaxTotalSizeOk returns a tuple with the MaxTotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalSize

`func (o *TenantQuotaDto) SetMaxTotalSize(v int64)`

SetMaxTotalSize sets MaxTotalSize field to given value.

### HasMaxTotalSize

`func (o *TenantQuotaDto) HasMaxTotalSize() bool`

HasMaxTotalSize returns a boolean if a field has been set.

### GetCountUser

`func (o *TenantQuotaDto) GetCountUser() int32`

GetCountUser returns the CountUser field if non-nil, zero value otherwise.

### GetCountUserOk

`func (o *TenantQuotaDto) GetCountUserOk() (*int32, bool)`

GetCountUserOk returns a tuple with the CountUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountUser

`func (o *TenantQuotaDto) SetCountUser(v int32)`

SetCountUser sets CountUser field to given value.

### HasCountUser

`func (o *TenantQuotaDto) HasCountUser() bool`

HasCountUser returns a boolean if a field has been set.

### GetCountRoomAdmin

`func (o *TenantQuotaDto) GetCountRoomAdmin() int32`

GetCountRoomAdmin returns the CountRoomAdmin field if non-nil, zero value otherwise.

### GetCountRoomAdminOk

`func (o *TenantQuotaDto) GetCountRoomAdminOk() (*int32, bool)`

GetCountRoomAdminOk returns a tuple with the CountRoomAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountRoomAdmin

`func (o *TenantQuotaDto) SetCountRoomAdmin(v int32)`

SetCountRoomAdmin sets CountRoomAdmin field to given value.

### HasCountRoomAdmin

`func (o *TenantQuotaDto) HasCountRoomAdmin() bool`

HasCountRoomAdmin returns a boolean if a field has been set.

### GetUsersInRoom

`func (o *TenantQuotaDto) GetUsersInRoom() int32`

GetUsersInRoom returns the UsersInRoom field if non-nil, zero value otherwise.

### GetUsersInRoomOk

`func (o *TenantQuotaDto) GetUsersInRoomOk() (*int32, bool)`

GetUsersInRoomOk returns a tuple with the UsersInRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersInRoom

`func (o *TenantQuotaDto) SetUsersInRoom(v int32)`

SetUsersInRoom sets UsersInRoom field to given value.

### HasUsersInRoom

`func (o *TenantQuotaDto) HasUsersInRoom() bool`

HasUsersInRoom returns a boolean if a field has been set.

### GetCountRoom

`func (o *TenantQuotaDto) GetCountRoom() int32`

GetCountRoom returns the CountRoom field if non-nil, zero value otherwise.

### GetCountRoomOk

`func (o *TenantQuotaDto) GetCountRoomOk() (*int32, bool)`

GetCountRoomOk returns a tuple with the CountRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountRoom

`func (o *TenantQuotaDto) SetCountRoom(v int32)`

SetCountRoom sets CountRoom field to given value.

### HasCountRoom

`func (o *TenantQuotaDto) HasCountRoom() bool`

HasCountRoom returns a boolean if a field has been set.

### GetNonProfit

`func (o *TenantQuotaDto) GetNonProfit() bool`

GetNonProfit returns the NonProfit field if non-nil, zero value otherwise.

### GetNonProfitOk

`func (o *TenantQuotaDto) GetNonProfitOk() (*bool, bool)`

GetNonProfitOk returns a tuple with the NonProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonProfit

`func (o *TenantQuotaDto) SetNonProfit(v bool)`

SetNonProfit sets NonProfit field to given value.

### HasNonProfit

`func (o *TenantQuotaDto) HasNonProfit() bool`

HasNonProfit returns a boolean if a field has been set.

### GetTrial

`func (o *TenantQuotaDto) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *TenantQuotaDto) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *TenantQuotaDto) SetTrial(v bool)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *TenantQuotaDto) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetFree

`func (o *TenantQuotaDto) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *TenantQuotaDto) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *TenantQuotaDto) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *TenantQuotaDto) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetUpdate

`func (o *TenantQuotaDto) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *TenantQuotaDto) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *TenantQuotaDto) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *TenantQuotaDto) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAudit

`func (o *TenantQuotaDto) GetAudit() bool`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *TenantQuotaDto) GetAuditOk() (*bool, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *TenantQuotaDto) SetAudit(v bool)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *TenantQuotaDto) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetDocsEdition

`func (o *TenantQuotaDto) GetDocsEdition() bool`

GetDocsEdition returns the DocsEdition field if non-nil, zero value otherwise.

### GetDocsEditionOk

`func (o *TenantQuotaDto) GetDocsEditionOk() (*bool, bool)`

GetDocsEditionOk returns a tuple with the DocsEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsEdition

`func (o *TenantQuotaDto) SetDocsEdition(v bool)`

SetDocsEdition sets DocsEdition field to given value.

### HasDocsEdition

`func (o *TenantQuotaDto) HasDocsEdition() bool`

HasDocsEdition returns a boolean if a field has been set.

### GetLdap

`func (o *TenantQuotaDto) GetLdap() bool`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *TenantQuotaDto) GetLdapOk() (*bool, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *TenantQuotaDto) SetLdap(v bool)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *TenantQuotaDto) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetSso

`func (o *TenantQuotaDto) GetSso() bool`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *TenantQuotaDto) GetSsoOk() (*bool, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *TenantQuotaDto) SetSso(v bool)`

SetSso sets Sso field to given value.

### HasSso

`func (o *TenantQuotaDto) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetStatistic

`func (o *TenantQuotaDto) GetStatistic() bool`

GetStatistic returns the Statistic field if non-nil, zero value otherwise.

### GetStatisticOk

`func (o *TenantQuotaDto) GetStatisticOk() (*bool, bool)`

GetStatisticOk returns a tuple with the Statistic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistic

`func (o *TenantQuotaDto) SetStatistic(v bool)`

SetStatistic sets Statistic field to given value.

### HasStatistic

`func (o *TenantQuotaDto) HasStatistic() bool`

HasStatistic returns a boolean if a field has been set.

### GetBranding

`func (o *TenantQuotaDto) GetBranding() bool`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *TenantQuotaDto) GetBrandingOk() (*bool, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *TenantQuotaDto) SetBranding(v bool)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *TenantQuotaDto) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetCustomization

`func (o *TenantQuotaDto) GetCustomization() bool`

GetCustomization returns the Customization field if non-nil, zero value otherwise.

### GetCustomizationOk

`func (o *TenantQuotaDto) GetCustomizationOk() (*bool, bool)`

GetCustomizationOk returns a tuple with the Customization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomization

`func (o *TenantQuotaDto) SetCustomization(v bool)`

SetCustomization sets Customization field to given value.

### HasCustomization

`func (o *TenantQuotaDto) HasCustomization() bool`

HasCustomization returns a boolean if a field has been set.

### GetLifetime

`func (o *TenantQuotaDto) GetLifetime() bool`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *TenantQuotaDto) GetLifetimeOk() (*bool, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *TenantQuotaDto) SetLifetime(v bool)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *TenantQuotaDto) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetAutomationApi

`func (o *TenantQuotaDto) GetAutomationApi() bool`

GetAutomationApi returns the AutomationApi field if non-nil, zero value otherwise.

### GetAutomationApiOk

`func (o *TenantQuotaDto) GetAutomationApiOk() (*bool, bool)`

GetAutomationApiOk returns a tuple with the AutomationApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationApi

`func (o *TenantQuotaDto) SetAutomationApi(v bool)`

SetAutomationApi sets AutomationApi field to given value.

### HasAutomationApi

`func (o *TenantQuotaDto) HasAutomationApi() bool`

HasAutomationApi returns a boolean if a field has been set.

### GetCustom

`func (o *TenantQuotaDto) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *TenantQuotaDto) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *TenantQuotaDto) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *TenantQuotaDto) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetRestore

`func (o *TenantQuotaDto) GetRestore() bool`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *TenantQuotaDto) GetRestoreOk() (*bool, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *TenantQuotaDto) SetRestore(v bool)`

SetRestore sets Restore field to given value.

### HasRestore

`func (o *TenantQuotaDto) HasRestore() bool`

HasRestore returns a boolean if a field has been set.

### GetOauth

`func (o *TenantQuotaDto) GetOauth() bool`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *TenantQuotaDto) GetOauthOk() (*bool, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *TenantQuotaDto) SetOauth(v bool)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *TenantQuotaDto) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetContentSearch

`func (o *TenantQuotaDto) GetContentSearch() bool`

GetContentSearch returns the ContentSearch field if non-nil, zero value otherwise.

### GetContentSearchOk

`func (o *TenantQuotaDto) GetContentSearchOk() (*bool, bool)`

GetContentSearchOk returns a tuple with the ContentSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSearch

`func (o *TenantQuotaDto) SetContentSearch(v bool)`

SetContentSearch sets ContentSearch field to given value.

### HasContentSearch

`func (o *TenantQuotaDto) HasContentSearch() bool`

HasContentSearch returns a boolean if a field has been set.

### GetThirdParty

`func (o *TenantQuotaDto) GetThirdParty() bool`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *TenantQuotaDto) GetThirdPartyOk() (*bool, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *TenantQuotaDto) SetThirdParty(v bool)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *TenantQuotaDto) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.

### GetYear

`func (o *TenantQuotaDto) GetYear() bool`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TenantQuotaDto) GetYearOk() (*bool, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TenantQuotaDto) SetYear(v bool)`

SetYear sets Year field to given value.

### HasYear

`func (o *TenantQuotaDto) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetCountFreeBackup

`func (o *TenantQuotaDto) GetCountFreeBackup() int32`

GetCountFreeBackup returns the CountFreeBackup field if non-nil, zero value otherwise.

### GetCountFreeBackupOk

`func (o *TenantQuotaDto) GetCountFreeBackupOk() (*int32, bool)`

GetCountFreeBackupOk returns a tuple with the CountFreeBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountFreeBackup

`func (o *TenantQuotaDto) SetCountFreeBackup(v int32)`

SetCountFreeBackup sets CountFreeBackup field to given value.

### HasCountFreeBackup

`func (o *TenantQuotaDto) HasCountFreeBackup() bool`

HasCountFreeBackup returns a boolean if a field has been set.

### GetBackup

`func (o *TenantQuotaDto) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *TenantQuotaDto) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *TenantQuotaDto) SetBackup(v bool)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *TenantQuotaDto) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetCountAIAgent

`func (o *TenantQuotaDto) GetCountAIAgent() int32`

GetCountAIAgent returns the CountAIAgent field if non-nil, zero value otherwise.

### GetCountAIAgentOk

`func (o *TenantQuotaDto) GetCountAIAgentOk() (*int32, bool)`

GetCountAIAgentOk returns a tuple with the CountAIAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAIAgent

`func (o *TenantQuotaDto) SetCountAIAgent(v int32)`

SetCountAIAgent sets CountAIAgent field to given value.

### HasCountAIAgent

`func (o *TenantQuotaDto) HasCountAIAgent() bool`

HasCountAIAgent returns a boolean if a field has been set.

### GetAiTools

`func (o *TenantQuotaDto) GetAiTools() bool`

GetAiTools returns the AiTools field if non-nil, zero value otherwise.

### GetAiToolsOk

`func (o *TenantQuotaDto) GetAiToolsOk() (*bool, bool)`

GetAiToolsOk returns a tuple with the AiTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiTools

`func (o *TenantQuotaDto) SetAiTools(v bool)`

SetAiTools sets AiTools field to given value.

### HasAiTools

`func (o *TenantQuotaDto) HasAiTools() bool`

HasAiTools returns a boolean if a field has been set.

### GetAiSearch

`func (o *TenantQuotaDto) GetAiSearch() bool`

GetAiSearch returns the AiSearch field if non-nil, zero value otherwise.

### GetAiSearchOk

`func (o *TenantQuotaDto) GetAiSearchOk() (*bool, bool)`

GetAiSearchOk returns a tuple with the AiSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSearch

`func (o *TenantQuotaDto) SetAiSearch(v bool)`

SetAiSearch sets AiSearch field to given value.

### HasAiSearch

`func (o *TenantQuotaDto) HasAiSearch() bool`

HasAiSearch returns a boolean if a field has been set.

### GetDocsCloud

`func (o *TenantQuotaDto) GetDocsCloud() int32`

GetDocsCloud returns the DocsCloud field if non-nil, zero value otherwise.

### GetDocsCloudOk

`func (o *TenantQuotaDto) GetDocsCloudOk() (*int32, bool)`

GetDocsCloudOk returns a tuple with the DocsCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsCloud

`func (o *TenantQuotaDto) SetDocsCloud(v int32)`

SetDocsCloud sets DocsCloud field to given value.

### HasDocsCloud

`func (o *TenantQuotaDto) HasDocsCloud() bool`

HasDocsCloud returns a boolean if a field has been set.

### GetDocsCloudDevPack

`func (o *TenantQuotaDto) GetDocsCloudDevPack() bool`

GetDocsCloudDevPack returns the DocsCloudDevPack field if non-nil, zero value otherwise.

### GetDocsCloudDevPackOk

`func (o *TenantQuotaDto) GetDocsCloudDevPackOk() (*bool, bool)`

GetDocsCloudDevPackOk returns a tuple with the DocsCloudDevPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsCloudDevPack

`func (o *TenantQuotaDto) SetDocsCloudDevPack(v bool)`

SetDocsCloudDevPack sets DocsCloudDevPack field to given value.

### HasDocsCloudDevPack

`func (o *TenantQuotaDto) HasDocsCloudDevPack() bool`

HasDocsCloudDevPack returns a boolean if a field has been set.

### GetDocsCloudTrial

`func (o *TenantQuotaDto) GetDocsCloudTrial() bool`

GetDocsCloudTrial returns the DocsCloudTrial field if non-nil, zero value otherwise.

### GetDocsCloudTrialOk

`func (o *TenantQuotaDto) GetDocsCloudTrialOk() (*bool, bool)`

GetDocsCloudTrialOk returns a tuple with the DocsCloudTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsCloudTrial

`func (o *TenantQuotaDto) SetDocsCloudTrial(v bool)`

SetDocsCloudTrial sets DocsCloudTrial field to given value.

### HasDocsCloudTrial

`func (o *TenantQuotaDto) HasDocsCloudTrial() bool`

HasDocsCloudTrial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


