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

// checks if the TenantQuotaDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantQuotaDto{}

// TenantQuotaDto The current tenant quota.
type TenantQuotaDto struct {
	// The tenant ID.
	TenantId *int32 `json:"tenantId,omitempty"`
	// The tenant name.
	Name NullableString `json:"name,omitempty"`
	// The tenant price.
	Price *float64 `json:"price,omitempty"`
	// The tenant price currency symbol.
	PriceCurrencySymbol NullableString `json:"priceCurrencySymbol,omitempty"`
	// The tenant price three-character ISO 4217 currency symbol.
	PriceISOCurrencySymbol NullableString `json:"priceISOCurrencySymbol,omitempty"`
	// The tenant product ID.
	ProductId NullableString `json:"productId,omitempty"`
	// The service name.
	ServiceName NullableString `json:"serviceName,omitempty"`
	// The service group.
	ServiceGroup NullableString `json:"serviceGroup,omitempty"`
	// Specifies if the tenant quota is visible or not.
	Visible *bool `json:"visible,omitempty"`
	// Specifies if the tenant quota applies to the wallet or not
	Wallet *bool `json:"wallet,omitempty"`
	// Specifies if the tenant quota is primary or additional.
	Additional *bool `json:"additional,omitempty"`
	// The quota due date.
	DueDate NullableTime `json:"dueDate,omitempty"`
	// The tenant quota features.
	Features NullableString `json:"features,omitempty"`
	// The tenant maximum file size.
	MaxFileSize *int64 `json:"maxFileSize,omitempty"`
	// The tenant maximum total size.
	MaxTotalSize *int64 `json:"maxTotalSize,omitempty"`
	// The number of portal users.
	CountUser *int32 `json:"countUser,omitempty"`
	// The number of portal room administrators.
	CountRoomAdmin *int32 `json:"countRoomAdmin,omitempty"`
	// The number of room users.
	UsersInRoom *int32 `json:"usersInRoom,omitempty"`
	// The number of rooms.
	CountRoom *int32 `json:"countRoom,omitempty"`
	// Specifies if the tenant quota is nonprofit or not.
	NonProfit *bool `json:"nonProfit,omitempty"`
	// Specifies if the tenant quota is trial or not.
	Trial *bool `json:"trial,omitempty"`
	// Specifies if the tenant quota is free or not.
	Free *bool `json:"free,omitempty"`
	// Specifies if the tenant quota is updated or not.
	Update *bool `json:"update,omitempty"`
	// Specifies if the audit trail is available or not.
	Audit *bool `json:"audit,omitempty"`
	// Specifies if ONLYOFFICE Docs is included in the tenant quota or not.
	DocsEdition *bool `json:"docsEdition,omitempty"`
	// Specifies if the LDAP settings are available or not.
	Ldap *bool `json:"ldap,omitempty"`
	// Specifies if the SSO settings are available or not.
	Sso *bool `json:"sso,omitempty"`
	// Specifies if the statistics settings are available or not.
	Statistic *bool `json:"statistic,omitempty"`
	// Specifies if the branding settings are available or not.
	Branding *bool `json:"branding,omitempty"`
	// Specifies if the customization settings are available or not.
	Customization *bool `json:"customization,omitempty"`
	// Specifies if the license has the lifetime settings or not.
	Lifetime *bool `json:"lifetime,omitempty"`
	// Specifies if the Automation API is available or not.
	AutomationApi *bool `json:"automationApi,omitempty"`
	// Specifies if the custom domain URL is available or not.
	Custom *bool `json:"custom,omitempty"`
	// Specifies if the restore is enabled or not.
	Restore *bool `json:"restore,omitempty"`
	// Specifies if Oauth is available or not.
	Oauth *bool `json:"oauth,omitempty"`
	// Specifies if the content search is available or not.
	ContentSearch *bool `json:"contentSearch,omitempty"`
	// Specifies if the third-party accounts linking is available or not.
	ThirdParty *bool `json:"thirdParty,omitempty"`
	// Specifies if the tenant quota is yearly subscription or not.
	Year *bool `json:"year,omitempty"`
	// The number of free backups within a month.
	CountFreeBackup *int32 `json:"countFreeBackup,omitempty"`
	// Specifies if the backup enabled as a wallet service or not.
	Backup *bool `json:"backup,omitempty"`
	// The number of AI agents.
	CountAIAgent *int32 `json:"countAIAgent,omitempty"`
	// Specifies if the AI tools enabled as a wallet service or not.
	AiTools *bool `json:"aiTools,omitempty"`
	// Specifies if the AI search enabled as a wallet service or not.
	AiSearch *bool `json:"aiSearch,omitempty"`
	// The number of Docs Connect users.
	DocsCloud *int32 `json:"docsCloud,omitempty"`
	// Specifies if the Docs Connect Dev Pack enabled or not.
	DocsCloudDevPack *bool `json:"docsCloudDevPack,omitempty"`
	// Specifies if the Docs Connect trial enabled or not.
	DocsCloudTrial *bool `json:"docsCloudTrial,omitempty"`
}

// NewTenantQuotaDto instantiates a new TenantQuotaDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantQuotaDto() *TenantQuotaDto {
	this := TenantQuotaDto{}
	return &this
}

// NewTenantQuotaDtoWithDefaults instantiates a new TenantQuotaDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantQuotaDtoWithDefaults() *TenantQuotaDto {
	this := TenantQuotaDto{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *TenantQuotaDto) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TenantQuotaDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TenantQuotaDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TenantQuotaDto) UnsetName() {
	o.Name.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsPriceSet() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *TenantQuotaDto) SetPrice(v float64) {
	o.Price = &v
}

// GetPriceCurrencySymbol returns the PriceCurrencySymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaDto) GetPriceCurrencySymbol() string {
	if o == nil || IsNil(o.PriceCurrencySymbol.Get()) {
		var ret string
		return ret
	}
	return *o.PriceCurrencySymbol.Get()
}

// GetPriceCurrencySymbolOk returns a tuple with the PriceCurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaDto) GetPriceCurrencySymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceCurrencySymbol.Get(), o.PriceCurrencySymbol.IsSet()
}

// HasPriceCurrencySymbol returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsPriceCurrencySymbolSet() bool {
	if o != nil && o.PriceCurrencySymbol.IsSet() {
		return true
	}

	return false
}

// SetPriceCurrencySymbol gets a reference to the given NullableString and assigns it to the PriceCurrencySymbol field.
func (o *TenantQuotaDto) SetPriceCurrencySymbol(v string) {
	o.PriceCurrencySymbol.Set(&v)
}
// SetPriceCurrencySymbolNil sets the value for PriceCurrencySymbol to be an explicit nil
func (o *TenantQuotaDto) SetPriceCurrencySymbolNil() {
	o.PriceCurrencySymbol.Set(nil)
}

// UnsetPriceCurrencySymbol ensures that no value is present for PriceCurrencySymbol, not even an explicit nil
func (o *TenantQuotaDto) UnsetPriceCurrencySymbol() {
	o.PriceCurrencySymbol.Unset()
}

// GetPriceISOCurrencySymbol returns the PriceISOCurrencySymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaDto) GetPriceISOCurrencySymbol() string {
	if o == nil || IsNil(o.PriceISOCurrencySymbol.Get()) {
		var ret string
		return ret
	}
	return *o.PriceISOCurrencySymbol.Get()
}

// GetPriceISOCurrencySymbolOk returns a tuple with the PriceISOCurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaDto) GetPriceISOCurrencySymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceISOCurrencySymbol.Get(), o.PriceISOCurrencySymbol.IsSet()
}

// HasPriceISOCurrencySymbol returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsPriceISOCurrencySymbolSet() bool {
	if o != nil && o.PriceISOCurrencySymbol.IsSet() {
		return true
	}

	return false
}

// SetPriceISOCurrencySymbol gets a reference to the given NullableString and assigns it to the PriceISOCurrencySymbol field.
func (o *TenantQuotaDto) SetPriceISOCurrencySymbol(v string) {
	o.PriceISOCurrencySymbol.Set(&v)
}
// SetPriceISOCurrencySymbolNil sets the value for PriceISOCurrencySymbol to be an explicit nil
func (o *TenantQuotaDto) SetPriceISOCurrencySymbolNil() {
	o.PriceISOCurrencySymbol.Set(nil)
}

// UnsetPriceISOCurrencySymbol ensures that no value is present for PriceISOCurrencySymbol, not even an explicit nil
func (o *TenantQuotaDto) UnsetPriceISOCurrencySymbol() {
	o.PriceISOCurrencySymbol.Unset()
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaDto) GetProductId() string {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret string
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaDto) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsProductIdSet() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableString and assigns it to the ProductId field.
func (o *TenantQuotaDto) SetProductId(v string) {
	o.ProductId.Set(&v)
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *TenantQuotaDto) SetProductIdNil() {
	o.ProductId.Set(nil)
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *TenantQuotaDto) UnsetProductId() {
	o.ProductId.Unset()
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaDto) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceName.Get()
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaDto) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceName.Get(), o.ServiceName.IsSet()
}

// HasServiceName returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsServiceNameSet() bool {
	if o != nil && o.ServiceName.IsSet() {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given NullableString and assigns it to the ServiceName field.
func (o *TenantQuotaDto) SetServiceName(v string) {
	o.ServiceName.Set(&v)
}
// SetServiceNameNil sets the value for ServiceName to be an explicit nil
func (o *TenantQuotaDto) SetServiceNameNil() {
	o.ServiceName.Set(nil)
}

// UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
func (o *TenantQuotaDto) UnsetServiceName() {
	o.ServiceName.Unset()
}

// GetServiceGroup returns the ServiceGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaDto) GetServiceGroup() string {
	if o == nil || IsNil(o.ServiceGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceGroup.Get()
}

// GetServiceGroupOk returns a tuple with the ServiceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaDto) GetServiceGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceGroup.Get(), o.ServiceGroup.IsSet()
}

// HasServiceGroup returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsServiceGroupSet() bool {
	if o != nil && o.ServiceGroup.IsSet() {
		return true
	}

	return false
}

// SetServiceGroup gets a reference to the given NullableString and assigns it to the ServiceGroup field.
func (o *TenantQuotaDto) SetServiceGroup(v string) {
	o.ServiceGroup.Set(&v)
}
// SetServiceGroupNil sets the value for ServiceGroup to be an explicit nil
func (o *TenantQuotaDto) SetServiceGroupNil() {
	o.ServiceGroup.Set(nil)
}

// UnsetServiceGroup ensures that no value is present for ServiceGroup, not even an explicit nil
func (o *TenantQuotaDto) UnsetServiceGroup() {
	o.ServiceGroup.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsVisibleSet() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *TenantQuotaDto) SetVisible(v bool) {
	o.Visible = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetWallet() bool {
	if o == nil || IsNil(o.Wallet) {
		var ret bool
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetWalletOk() (*bool, bool) {
	if o == nil || IsNil(o.Wallet) {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsWalletSet() bool {
	if o != nil && !IsNil(o.Wallet) {
		return true
	}

	return false
}

// SetWallet gets a reference to the given bool and assigns it to the Wallet field.
func (o *TenantQuotaDto) SetWallet(v bool) {
	o.Wallet = &v
}

// GetAdditional returns the Additional field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetAdditional() bool {
	if o == nil || IsNil(o.Additional) {
		var ret bool
		return ret
	}
	return *o.Additional
}

// GetAdditionalOk returns a tuple with the Additional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetAdditionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Additional) {
		return nil, false
	}
	return o.Additional, true
}

// HasAdditional returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsAdditionalSet() bool {
	if o != nil && !IsNil(o.Additional) {
		return true
	}

	return false
}

// SetAdditional gets a reference to the given bool and assigns it to the Additional field.
func (o *TenantQuotaDto) SetAdditional(v bool) {
	o.Additional = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaDto) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaDto) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsDueDateSet() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableTime and assigns it to the DueDate field.
func (o *TenantQuotaDto) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *TenantQuotaDto) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *TenantQuotaDto) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetFeatures returns the Features field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaDto) GetFeatures() string {
	if o == nil || IsNil(o.Features.Get()) {
		var ret string
		return ret
	}
	return *o.Features.Get()
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaDto) GetFeaturesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features.Get(), o.Features.IsSet()
}

// HasFeatures returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsFeaturesSet() bool {
	if o != nil && o.Features.IsSet() {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given NullableString and assigns it to the Features field.
func (o *TenantQuotaDto) SetFeatures(v string) {
	o.Features.Set(&v)
}
// SetFeaturesNil sets the value for Features to be an explicit nil
func (o *TenantQuotaDto) SetFeaturesNil() {
	o.Features.Set(nil)
}

// UnsetFeatures ensures that no value is present for Features, not even an explicit nil
func (o *TenantQuotaDto) UnsetFeatures() {
	o.Features.Unset()
}

// GetMaxFileSize returns the MaxFileSize field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetMaxFileSize() int64 {
	if o == nil || IsNil(o.MaxFileSize) {
		var ret int64
		return ret
	}
	return *o.MaxFileSize
}

// GetMaxFileSizeOk returns a tuple with the MaxFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetMaxFileSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxFileSize) {
		return nil, false
	}
	return o.MaxFileSize, true
}

// HasMaxFileSize returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsMaxFileSizeSet() bool {
	if o != nil && !IsNil(o.MaxFileSize) {
		return true
	}

	return false
}

// SetMaxFileSize gets a reference to the given int64 and assigns it to the MaxFileSize field.
func (o *TenantQuotaDto) SetMaxFileSize(v int64) {
	o.MaxFileSize = &v
}

// GetMaxTotalSize returns the MaxTotalSize field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetMaxTotalSize() int64 {
	if o == nil || IsNil(o.MaxTotalSize) {
		var ret int64
		return ret
	}
	return *o.MaxTotalSize
}

// GetMaxTotalSizeOk returns a tuple with the MaxTotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetMaxTotalSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxTotalSize) {
		return nil, false
	}
	return o.MaxTotalSize, true
}

// HasMaxTotalSize returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsMaxTotalSizeSet() bool {
	if o != nil && !IsNil(o.MaxTotalSize) {
		return true
	}

	return false
}

// SetMaxTotalSize gets a reference to the given int64 and assigns it to the MaxTotalSize field.
func (o *TenantQuotaDto) SetMaxTotalSize(v int64) {
	o.MaxTotalSize = &v
}

// GetCountUser returns the CountUser field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetCountUser() int32 {
	if o == nil || IsNil(o.CountUser) {
		var ret int32
		return ret
	}
	return *o.CountUser
}

// GetCountUserOk returns a tuple with the CountUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetCountUserOk() (*int32, bool) {
	if o == nil || IsNil(o.CountUser) {
		return nil, false
	}
	return o.CountUser, true
}

// HasCountUser returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsCountUserSet() bool {
	if o != nil && !IsNil(o.CountUser) {
		return true
	}

	return false
}

// SetCountUser gets a reference to the given int32 and assigns it to the CountUser field.
func (o *TenantQuotaDto) SetCountUser(v int32) {
	o.CountUser = &v
}

// GetCountRoomAdmin returns the CountRoomAdmin field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetCountRoomAdmin() int32 {
	if o == nil || IsNil(o.CountRoomAdmin) {
		var ret int32
		return ret
	}
	return *o.CountRoomAdmin
}

// GetCountRoomAdminOk returns a tuple with the CountRoomAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetCountRoomAdminOk() (*int32, bool) {
	if o == nil || IsNil(o.CountRoomAdmin) {
		return nil, false
	}
	return o.CountRoomAdmin, true
}

// HasCountRoomAdmin returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsCountRoomAdminSet() bool {
	if o != nil && !IsNil(o.CountRoomAdmin) {
		return true
	}

	return false
}

// SetCountRoomAdmin gets a reference to the given int32 and assigns it to the CountRoomAdmin field.
func (o *TenantQuotaDto) SetCountRoomAdmin(v int32) {
	o.CountRoomAdmin = &v
}

// GetUsersInRoom returns the UsersInRoom field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetUsersInRoom() int32 {
	if o == nil || IsNil(o.UsersInRoom) {
		var ret int32
		return ret
	}
	return *o.UsersInRoom
}

// GetUsersInRoomOk returns a tuple with the UsersInRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetUsersInRoomOk() (*int32, bool) {
	if o == nil || IsNil(o.UsersInRoom) {
		return nil, false
	}
	return o.UsersInRoom, true
}

// HasUsersInRoom returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsUsersInRoomSet() bool {
	if o != nil && !IsNil(o.UsersInRoom) {
		return true
	}

	return false
}

// SetUsersInRoom gets a reference to the given int32 and assigns it to the UsersInRoom field.
func (o *TenantQuotaDto) SetUsersInRoom(v int32) {
	o.UsersInRoom = &v
}

// GetCountRoom returns the CountRoom field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetCountRoom() int32 {
	if o == nil || IsNil(o.CountRoom) {
		var ret int32
		return ret
	}
	return *o.CountRoom
}

// GetCountRoomOk returns a tuple with the CountRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetCountRoomOk() (*int32, bool) {
	if o == nil || IsNil(o.CountRoom) {
		return nil, false
	}
	return o.CountRoom, true
}

// HasCountRoom returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsCountRoomSet() bool {
	if o != nil && !IsNil(o.CountRoom) {
		return true
	}

	return false
}

// SetCountRoom gets a reference to the given int32 and assigns it to the CountRoom field.
func (o *TenantQuotaDto) SetCountRoom(v int32) {
	o.CountRoom = &v
}

// GetNonProfit returns the NonProfit field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetNonProfit() bool {
	if o == nil || IsNil(o.NonProfit) {
		var ret bool
		return ret
	}
	return *o.NonProfit
}

// GetNonProfitOk returns a tuple with the NonProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetNonProfitOk() (*bool, bool) {
	if o == nil || IsNil(o.NonProfit) {
		return nil, false
	}
	return o.NonProfit, true
}

// HasNonProfit returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsNonProfitSet() bool {
	if o != nil && !IsNil(o.NonProfit) {
		return true
	}

	return false
}

// SetNonProfit gets a reference to the given bool and assigns it to the NonProfit field.
func (o *TenantQuotaDto) SetNonProfit(v bool) {
	o.NonProfit = &v
}

// GetTrial returns the Trial field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetTrial() bool {
	if o == nil || IsNil(o.Trial) {
		var ret bool
		return ret
	}
	return *o.Trial
}

// GetTrialOk returns a tuple with the Trial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetTrialOk() (*bool, bool) {
	if o == nil || IsNil(o.Trial) {
		return nil, false
	}
	return o.Trial, true
}

// HasTrial returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsTrialSet() bool {
	if o != nil && !IsNil(o.Trial) {
		return true
	}

	return false
}

// SetTrial gets a reference to the given bool and assigns it to the Trial field.
func (o *TenantQuotaDto) SetTrial(v bool) {
	o.Trial = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetFree() bool {
	if o == nil || IsNil(o.Free) {
		var ret bool
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetFreeOk() (*bool, bool) {
	if o == nil || IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsFreeSet() bool {
	if o != nil && !IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given bool and assigns it to the Free field.
func (o *TenantQuotaDto) SetFree(v bool) {
	o.Free = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetUpdate() bool {
	if o == nil || IsNil(o.Update) {
		var ret bool
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsUpdateSet() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given bool and assigns it to the Update field.
func (o *TenantQuotaDto) SetUpdate(v bool) {
	o.Update = &v
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetAudit() bool {
	if o == nil || IsNil(o.Audit) {
		var ret bool
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetAuditOk() (*bool, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsAuditSet() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given bool and assigns it to the Audit field.
func (o *TenantQuotaDto) SetAudit(v bool) {
	o.Audit = &v
}

// GetDocsEdition returns the DocsEdition field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetDocsEdition() bool {
	if o == nil || IsNil(o.DocsEdition) {
		var ret bool
		return ret
	}
	return *o.DocsEdition
}

// GetDocsEditionOk returns a tuple with the DocsEdition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetDocsEditionOk() (*bool, bool) {
	if o == nil || IsNil(o.DocsEdition) {
		return nil, false
	}
	return o.DocsEdition, true
}

// HasDocsEdition returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsDocsEditionSet() bool {
	if o != nil && !IsNil(o.DocsEdition) {
		return true
	}

	return false
}

// SetDocsEdition gets a reference to the given bool and assigns it to the DocsEdition field.
func (o *TenantQuotaDto) SetDocsEdition(v bool) {
	o.DocsEdition = &v
}

// GetLdap returns the Ldap field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetLdap() bool {
	if o == nil || IsNil(o.Ldap) {
		var ret bool
		return ret
	}
	return *o.Ldap
}

// GetLdapOk returns a tuple with the Ldap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetLdapOk() (*bool, bool) {
	if o == nil || IsNil(o.Ldap) {
		return nil, false
	}
	return o.Ldap, true
}

// HasLdap returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsLdapSet() bool {
	if o != nil && !IsNil(o.Ldap) {
		return true
	}

	return false
}

// SetLdap gets a reference to the given bool and assigns it to the Ldap field.
func (o *TenantQuotaDto) SetLdap(v bool) {
	o.Ldap = &v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetSso() bool {
	if o == nil || IsNil(o.Sso) {
		var ret bool
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetSsoOk() (*bool, bool) {
	if o == nil || IsNil(o.Sso) {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsSsoSet() bool {
	if o != nil && !IsNil(o.Sso) {
		return true
	}

	return false
}

// SetSso gets a reference to the given bool and assigns it to the Sso field.
func (o *TenantQuotaDto) SetSso(v bool) {
	o.Sso = &v
}

// GetStatistic returns the Statistic field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetStatistic() bool {
	if o == nil || IsNil(o.Statistic) {
		var ret bool
		return ret
	}
	return *o.Statistic
}

// GetStatisticOk returns a tuple with the Statistic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetStatisticOk() (*bool, bool) {
	if o == nil || IsNil(o.Statistic) {
		return nil, false
	}
	return o.Statistic, true
}

// HasStatistic returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsStatisticSet() bool {
	if o != nil && !IsNil(o.Statistic) {
		return true
	}

	return false
}

// SetStatistic gets a reference to the given bool and assigns it to the Statistic field.
func (o *TenantQuotaDto) SetStatistic(v bool) {
	o.Statistic = &v
}

// GetBranding returns the Branding field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetBranding() bool {
	if o == nil || IsNil(o.Branding) {
		var ret bool
		return ret
	}
	return *o.Branding
}

// GetBrandingOk returns a tuple with the Branding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetBrandingOk() (*bool, bool) {
	if o == nil || IsNil(o.Branding) {
		return nil, false
	}
	return o.Branding, true
}

// HasBranding returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsBrandingSet() bool {
	if o != nil && !IsNil(o.Branding) {
		return true
	}

	return false
}

// SetBranding gets a reference to the given bool and assigns it to the Branding field.
func (o *TenantQuotaDto) SetBranding(v bool) {
	o.Branding = &v
}

// GetCustomization returns the Customization field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetCustomization() bool {
	if o == nil || IsNil(o.Customization) {
		var ret bool
		return ret
	}
	return *o.Customization
}

// GetCustomizationOk returns a tuple with the Customization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetCustomizationOk() (*bool, bool) {
	if o == nil || IsNil(o.Customization) {
		return nil, false
	}
	return o.Customization, true
}

// HasCustomization returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsCustomizationSet() bool {
	if o != nil && !IsNil(o.Customization) {
		return true
	}

	return false
}

// SetCustomization gets a reference to the given bool and assigns it to the Customization field.
func (o *TenantQuotaDto) SetCustomization(v bool) {
	o.Customization = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetLifetime() bool {
	if o == nil || IsNil(o.Lifetime) {
		var ret bool
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetLifetimeOk() (*bool, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given bool and assigns it to the Lifetime field.
func (o *TenantQuotaDto) SetLifetime(v bool) {
	o.Lifetime = &v
}

// GetAutomationApi returns the AutomationApi field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetAutomationApi() bool {
	if o == nil || IsNil(o.AutomationApi) {
		var ret bool
		return ret
	}
	return *o.AutomationApi
}

// GetAutomationApiOk returns a tuple with the AutomationApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetAutomationApiOk() (*bool, bool) {
	if o == nil || IsNil(o.AutomationApi) {
		return nil, false
	}
	return o.AutomationApi, true
}

// HasAutomationApi returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsAutomationApiSet() bool {
	if o != nil && !IsNil(o.AutomationApi) {
		return true
	}

	return false
}

// SetAutomationApi gets a reference to the given bool and assigns it to the AutomationApi field.
func (o *TenantQuotaDto) SetAutomationApi(v bool) {
	o.AutomationApi = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetCustom() bool {
	if o == nil || IsNil(o.Custom) {
		var ret bool
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetCustomOk() (*bool, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsCustomSet() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given bool and assigns it to the Custom field.
func (o *TenantQuotaDto) SetCustom(v bool) {
	o.Custom = &v
}

// GetRestore returns the Restore field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetRestore() bool {
	if o == nil || IsNil(o.Restore) {
		var ret bool
		return ret
	}
	return *o.Restore
}

// GetRestoreOk returns a tuple with the Restore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetRestoreOk() (*bool, bool) {
	if o == nil || IsNil(o.Restore) {
		return nil, false
	}
	return o.Restore, true
}

// HasRestore returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsRestoreSet() bool {
	if o != nil && !IsNil(o.Restore) {
		return true
	}

	return false
}

// SetRestore gets a reference to the given bool and assigns it to the Restore field.
func (o *TenantQuotaDto) SetRestore(v bool) {
	o.Restore = &v
}

// GetOauth returns the Oauth field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetOauth() bool {
	if o == nil || IsNil(o.Oauth) {
		var ret bool
		return ret
	}
	return *o.Oauth
}

// GetOauthOk returns a tuple with the Oauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetOauthOk() (*bool, bool) {
	if o == nil || IsNil(o.Oauth) {
		return nil, false
	}
	return o.Oauth, true
}

// HasOauth returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsOauthSet() bool {
	if o != nil && !IsNil(o.Oauth) {
		return true
	}

	return false
}

// SetOauth gets a reference to the given bool and assigns it to the Oauth field.
func (o *TenantQuotaDto) SetOauth(v bool) {
	o.Oauth = &v
}

// GetContentSearch returns the ContentSearch field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetContentSearch() bool {
	if o == nil || IsNil(o.ContentSearch) {
		var ret bool
		return ret
	}
	return *o.ContentSearch
}

// GetContentSearchOk returns a tuple with the ContentSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetContentSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.ContentSearch) {
		return nil, false
	}
	return o.ContentSearch, true
}

// HasContentSearch returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsContentSearchSet() bool {
	if o != nil && !IsNil(o.ContentSearch) {
		return true
	}

	return false
}

// SetContentSearch gets a reference to the given bool and assigns it to the ContentSearch field.
func (o *TenantQuotaDto) SetContentSearch(v bool) {
	o.ContentSearch = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetThirdParty() bool {
	if o == nil || IsNil(o.ThirdParty) {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetThirdPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.ThirdParty) {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsThirdPartySet() bool {
	if o != nil && !IsNil(o.ThirdParty) {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *TenantQuotaDto) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetYear() bool {
	if o == nil || IsNil(o.Year) {
		var ret bool
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetYearOk() (*bool, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsYearSet() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given bool and assigns it to the Year field.
func (o *TenantQuotaDto) SetYear(v bool) {
	o.Year = &v
}

// GetCountFreeBackup returns the CountFreeBackup field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetCountFreeBackup() int32 {
	if o == nil || IsNil(o.CountFreeBackup) {
		var ret int32
		return ret
	}
	return *o.CountFreeBackup
}

// GetCountFreeBackupOk returns a tuple with the CountFreeBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetCountFreeBackupOk() (*int32, bool) {
	if o == nil || IsNil(o.CountFreeBackup) {
		return nil, false
	}
	return o.CountFreeBackup, true
}

// HasCountFreeBackup returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsCountFreeBackupSet() bool {
	if o != nil && !IsNil(o.CountFreeBackup) {
		return true
	}

	return false
}

// SetCountFreeBackup gets a reference to the given int32 and assigns it to the CountFreeBackup field.
func (o *TenantQuotaDto) SetCountFreeBackup(v int32) {
	o.CountFreeBackup = &v
}

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetBackup() bool {
	if o == nil || IsNil(o.Backup) {
		var ret bool
		return ret
	}
	return *o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetBackupOk() (*bool, bool) {
	if o == nil || IsNil(o.Backup) {
		return nil, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsBackupSet() bool {
	if o != nil && !IsNil(o.Backup) {
		return true
	}

	return false
}

// SetBackup gets a reference to the given bool and assigns it to the Backup field.
func (o *TenantQuotaDto) SetBackup(v bool) {
	o.Backup = &v
}

// GetCountAIAgent returns the CountAIAgent field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetCountAIAgent() int32 {
	if o == nil || IsNil(o.CountAIAgent) {
		var ret int32
		return ret
	}
	return *o.CountAIAgent
}

// GetCountAIAgentOk returns a tuple with the CountAIAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetCountAIAgentOk() (*int32, bool) {
	if o == nil || IsNil(o.CountAIAgent) {
		return nil, false
	}
	return o.CountAIAgent, true
}

// HasCountAIAgent returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsCountAIAgentSet() bool {
	if o != nil && !IsNil(o.CountAIAgent) {
		return true
	}

	return false
}

// SetCountAIAgent gets a reference to the given int32 and assigns it to the CountAIAgent field.
func (o *TenantQuotaDto) SetCountAIAgent(v int32) {
	o.CountAIAgent = &v
}

// GetAiTools returns the AiTools field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetAiTools() bool {
	if o == nil || IsNil(o.AiTools) {
		var ret bool
		return ret
	}
	return *o.AiTools
}

// GetAiToolsOk returns a tuple with the AiTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetAiToolsOk() (*bool, bool) {
	if o == nil || IsNil(o.AiTools) {
		return nil, false
	}
	return o.AiTools, true
}

// HasAiTools returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsAiToolsSet() bool {
	if o != nil && !IsNil(o.AiTools) {
		return true
	}

	return false
}

// SetAiTools gets a reference to the given bool and assigns it to the AiTools field.
func (o *TenantQuotaDto) SetAiTools(v bool) {
	o.AiTools = &v
}

// GetAiSearch returns the AiSearch field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetAiSearch() bool {
	if o == nil || IsNil(o.AiSearch) {
		var ret bool
		return ret
	}
	return *o.AiSearch
}

// GetAiSearchOk returns a tuple with the AiSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetAiSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.AiSearch) {
		return nil, false
	}
	return o.AiSearch, true
}

// HasAiSearch returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsAiSearchSet() bool {
	if o != nil && !IsNil(o.AiSearch) {
		return true
	}

	return false
}

// SetAiSearch gets a reference to the given bool and assigns it to the AiSearch field.
func (o *TenantQuotaDto) SetAiSearch(v bool) {
	o.AiSearch = &v
}

// GetDocsCloud returns the DocsCloud field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetDocsCloud() int32 {
	if o == nil || IsNil(o.DocsCloud) {
		var ret int32
		return ret
	}
	return *o.DocsCloud
}

// GetDocsCloudOk returns a tuple with the DocsCloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetDocsCloudOk() (*int32, bool) {
	if o == nil || IsNil(o.DocsCloud) {
		return nil, false
	}
	return o.DocsCloud, true
}

// HasDocsCloud returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsDocsCloudSet() bool {
	if o != nil && !IsNil(o.DocsCloud) {
		return true
	}

	return false
}

// SetDocsCloud gets a reference to the given int32 and assigns it to the DocsCloud field.
func (o *TenantQuotaDto) SetDocsCloud(v int32) {
	o.DocsCloud = &v
}

// GetDocsCloudDevPack returns the DocsCloudDevPack field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetDocsCloudDevPack() bool {
	if o == nil || IsNil(o.DocsCloudDevPack) {
		var ret bool
		return ret
	}
	return *o.DocsCloudDevPack
}

// GetDocsCloudDevPackOk returns a tuple with the DocsCloudDevPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetDocsCloudDevPackOk() (*bool, bool) {
	if o == nil || IsNil(o.DocsCloudDevPack) {
		return nil, false
	}
	return o.DocsCloudDevPack, true
}

// HasDocsCloudDevPack returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsDocsCloudDevPackSet() bool {
	if o != nil && !IsNil(o.DocsCloudDevPack) {
		return true
	}

	return false
}

// SetDocsCloudDevPack gets a reference to the given bool and assigns it to the DocsCloudDevPack field.
func (o *TenantQuotaDto) SetDocsCloudDevPack(v bool) {
	o.DocsCloudDevPack = &v
}

// GetDocsCloudTrial returns the DocsCloudTrial field value if set, zero value otherwise.
func (o *TenantQuotaDto) GetDocsCloudTrial() bool {
	if o == nil || IsNil(o.DocsCloudTrial) {
		var ret bool
		return ret
	}
	return *o.DocsCloudTrial
}

// GetDocsCloudTrialOk returns a tuple with the DocsCloudTrial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaDto) GetDocsCloudTrialOk() (*bool, bool) {
	if o == nil || IsNil(o.DocsCloudTrial) {
		return nil, false
	}
	return o.DocsCloudTrial, true
}

// HasDocsCloudTrial returns a boolean if a field has been set.
func (o *TenantQuotaDto) IsDocsCloudTrialSet() bool {
	if o != nil && !IsNil(o.DocsCloudTrial) {
		return true
	}

	return false
}

// SetDocsCloudTrial gets a reference to the given bool and assigns it to the DocsCloudTrial field.
func (o *TenantQuotaDto) SetDocsCloudTrial(v bool) {
	o.DocsCloudTrial = &v
}

func (o TenantQuotaDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantQuotaDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if o.PriceCurrencySymbol.IsSet() {
		toSerialize["priceCurrencySymbol"] = o.PriceCurrencySymbol.Get()
	}
	if o.PriceISOCurrencySymbol.IsSet() {
		toSerialize["priceISOCurrencySymbol"] = o.PriceISOCurrencySymbol.Get()
	}
	if o.ProductId.IsSet() {
		toSerialize["productId"] = o.ProductId.Get()
	}
	if o.ServiceName.IsSet() {
		toSerialize["serviceName"] = o.ServiceName.Get()
	}
	if o.ServiceGroup.IsSet() {
		toSerialize["serviceGroup"] = o.ServiceGroup.Get()
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if !IsNil(o.Wallet) {
		toSerialize["wallet"] = o.Wallet
	}
	if !IsNil(o.Additional) {
		toSerialize["additional"] = o.Additional
	}
	if o.DueDate.IsSet() {
		toSerialize["dueDate"] = o.DueDate.Get()
	}
	if o.Features.IsSet() {
		toSerialize["features"] = o.Features.Get()
	}
	if !IsNil(o.MaxFileSize) {
		toSerialize["maxFileSize"] = o.MaxFileSize
	}
	if !IsNil(o.MaxTotalSize) {
		toSerialize["maxTotalSize"] = o.MaxTotalSize
	}
	if !IsNil(o.CountUser) {
		toSerialize["countUser"] = o.CountUser
	}
	if !IsNil(o.CountRoomAdmin) {
		toSerialize["countRoomAdmin"] = o.CountRoomAdmin
	}
	if !IsNil(o.UsersInRoom) {
		toSerialize["usersInRoom"] = o.UsersInRoom
	}
	if !IsNil(o.CountRoom) {
		toSerialize["countRoom"] = o.CountRoom
	}
	if !IsNil(o.NonProfit) {
		toSerialize["nonProfit"] = o.NonProfit
	}
	if !IsNil(o.Trial) {
		toSerialize["trial"] = o.Trial
	}
	if !IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.DocsEdition) {
		toSerialize["docsEdition"] = o.DocsEdition
	}
	if !IsNil(o.Ldap) {
		toSerialize["ldap"] = o.Ldap
	}
	if !IsNil(o.Sso) {
		toSerialize["sso"] = o.Sso
	}
	if !IsNil(o.Statistic) {
		toSerialize["statistic"] = o.Statistic
	}
	if !IsNil(o.Branding) {
		toSerialize["branding"] = o.Branding
	}
	if !IsNil(o.Customization) {
		toSerialize["customization"] = o.Customization
	}
	if !IsNil(o.Lifetime) {
		toSerialize["lifetime"] = o.Lifetime
	}
	if !IsNil(o.AutomationApi) {
		toSerialize["automationApi"] = o.AutomationApi
	}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.Restore) {
		toSerialize["restore"] = o.Restore
	}
	if !IsNil(o.Oauth) {
		toSerialize["oauth"] = o.Oauth
	}
	if !IsNil(o.ContentSearch) {
		toSerialize["contentSearch"] = o.ContentSearch
	}
	if !IsNil(o.ThirdParty) {
		toSerialize["thirdParty"] = o.ThirdParty
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.CountFreeBackup) {
		toSerialize["countFreeBackup"] = o.CountFreeBackup
	}
	if !IsNil(o.Backup) {
		toSerialize["backup"] = o.Backup
	}
	if !IsNil(o.CountAIAgent) {
		toSerialize["countAIAgent"] = o.CountAIAgent
	}
	if !IsNil(o.AiTools) {
		toSerialize["aiTools"] = o.AiTools
	}
	if !IsNil(o.AiSearch) {
		toSerialize["aiSearch"] = o.AiSearch
	}
	if !IsNil(o.DocsCloud) {
		toSerialize["docsCloud"] = o.DocsCloud
	}
	if !IsNil(o.DocsCloudDevPack) {
		toSerialize["docsCloudDevPack"] = o.DocsCloudDevPack
	}
	if !IsNil(o.DocsCloudTrial) {
		toSerialize["docsCloudTrial"] = o.DocsCloudTrial
	}
	return toSerialize, nil
}

type NullableTenantQuotaDto struct {
	value *TenantQuotaDto
	isSet bool
}

func (v NullableTenantQuotaDto) Get() *TenantQuotaDto {
	return v.value
}

func (v *NullableTenantQuotaDto) Set(val *TenantQuotaDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantQuotaDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantQuotaDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantQuotaDto(val *TenantQuotaDto) *NullableTenantQuotaDto {
	return &NullableTenantQuotaDto{value: val, isSet: true}
}

func (v NullableTenantQuotaDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantQuotaDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

