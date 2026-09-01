# TenantWalletSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether automatic top-up for the tenant wallet is enabled. | [optional] 
**MinBalance** | Pointer to **int32** | The minimum wallet balance at which automatic top-up will be triggered. Must be between 5 and 1000. | [optional] 
**UpToBalance** | Pointer to **int32** | The maximum wallet balance at which automatic top-up will be triggered. Must be between 6 and 5000. | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol. | [optional] 
**LowBalanceThreshold** | Pointer to **int32** | The wallet balance below which a low-balance notification is sent. Set internally, not user-configurable. | [optional] 
**LowBalanceNotified** | Pointer to **bool** | Specifies whether a low-balance notification has already been sent for the current dip below ASC.Core.Tenants.TenantWalletSettings.LowBalanceThreshold. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the tenant wallet settings were last modified. | [optional] 

## Methods

### NewTenantWalletSettings

`func NewTenantWalletSettings() *TenantWalletSettings`

NewTenantWalletSettings instantiates a new TenantWalletSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWalletSettingsWithDefaults

`func NewTenantWalletSettingsWithDefaults() *TenantWalletSettings`

NewTenantWalletSettingsWithDefaults instantiates a new TenantWalletSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TenantWalletSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TenantWalletSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TenantWalletSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TenantWalletSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinBalance

`func (o *TenantWalletSettings) GetMinBalance() int32`

GetMinBalance returns the MinBalance field if non-nil, zero value otherwise.

### GetMinBalanceOk

`func (o *TenantWalletSettings) GetMinBalanceOk() (*int32, bool)`

GetMinBalanceOk returns a tuple with the MinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBalance

`func (o *TenantWalletSettings) SetMinBalance(v int32)`

SetMinBalance sets MinBalance field to given value.

### HasMinBalance

`func (o *TenantWalletSettings) HasMinBalance() bool`

HasMinBalance returns a boolean if a field has been set.

### GetUpToBalance

`func (o *TenantWalletSettings) GetUpToBalance() int32`

GetUpToBalance returns the UpToBalance field if non-nil, zero value otherwise.

### GetUpToBalanceOk

`func (o *TenantWalletSettings) GetUpToBalanceOk() (*int32, bool)`

GetUpToBalanceOk returns a tuple with the UpToBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToBalance

`func (o *TenantWalletSettings) SetUpToBalance(v int32)`

SetUpToBalance sets UpToBalance field to given value.

### HasUpToBalance

`func (o *TenantWalletSettings) HasUpToBalance() bool`

HasUpToBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *TenantWalletSettings) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TenantWalletSettings) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TenantWalletSettings) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TenantWalletSettings) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *TenantWalletSettings) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *TenantWalletSettings) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetLowBalanceThreshold

`func (o *TenantWalletSettings) GetLowBalanceThreshold() int32`

GetLowBalanceThreshold returns the LowBalanceThreshold field if non-nil, zero value otherwise.

### GetLowBalanceThresholdOk

`func (o *TenantWalletSettings) GetLowBalanceThresholdOk() (*int32, bool)`

GetLowBalanceThresholdOk returns a tuple with the LowBalanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBalanceThreshold

`func (o *TenantWalletSettings) SetLowBalanceThreshold(v int32)`

SetLowBalanceThreshold sets LowBalanceThreshold field to given value.

### HasLowBalanceThreshold

`func (o *TenantWalletSettings) HasLowBalanceThreshold() bool`

HasLowBalanceThreshold returns a boolean if a field has been set.

### GetLowBalanceNotified

`func (o *TenantWalletSettings) GetLowBalanceNotified() bool`

GetLowBalanceNotified returns the LowBalanceNotified field if non-nil, zero value otherwise.

### GetLowBalanceNotifiedOk

`func (o *TenantWalletSettings) GetLowBalanceNotifiedOk() (*bool, bool)`

GetLowBalanceNotifiedOk returns a tuple with the LowBalanceNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBalanceNotified

`func (o *TenantWalletSettings) SetLowBalanceNotified(v bool)`

SetLowBalanceNotified sets LowBalanceNotified field to given value.

### HasLowBalanceNotified

`func (o *TenantWalletSettings) HasLowBalanceNotified() bool`

HasLowBalanceNotified returns a boolean if a field has been set.

### GetLastModified

`func (o *TenantWalletSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantWalletSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantWalletSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantWalletSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


