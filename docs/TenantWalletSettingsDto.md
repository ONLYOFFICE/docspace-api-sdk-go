# TenantWalletSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the payment method on file is charged automatically when the wallet balance runs low. | [optional] 
**MinBalance** | Pointer to **int32** | The balance below which a top-up is charged, in `currency`; 0 while top-up has never been configured. | [optional] 
**UpToBalance** | Pointer to **int32** | The balance a top-up brings the wallet up to, in `currency`; 0 while top-up has never been configured. | [optional] 
**Currency** | Pointer to **NullableString** | The three-letter ISO 4217 code both amounts are expressed in, or `null` while top-up has never been configured. | [optional] 
**LowBalanceThreshold** | Pointer to **int32** | The wallet balance below which the portal sends its low-balance warning. The portal maintains it; it cannot be  set by a request. | [optional] 
**LowBalanceNotified** | Pointer to **bool** | Whether the low-balance warning has already been sent for the current dip below `lowBalanceThreshold`. The  portal maintains it, and switching top-up on re-arms it. | [optional] 
**LastModified** | Pointer to **time.Time** | When the settings were last stored; when they were never stored, the moment they were read instead. | [optional] 

## Methods

### NewTenantWalletSettingsDto

`func NewTenantWalletSettingsDto() *TenantWalletSettingsDto`

NewTenantWalletSettingsDto instantiates a new TenantWalletSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWalletSettingsDtoWithDefaults

`func NewTenantWalletSettingsDtoWithDefaults() *TenantWalletSettingsDto`

NewTenantWalletSettingsDtoWithDefaults instantiates a new TenantWalletSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TenantWalletSettingsDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TenantWalletSettingsDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TenantWalletSettingsDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TenantWalletSettingsDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinBalance

`func (o *TenantWalletSettingsDto) GetMinBalance() int32`

GetMinBalance returns the MinBalance field if non-nil, zero value otherwise.

### GetMinBalanceOk

`func (o *TenantWalletSettingsDto) GetMinBalanceOk() (*int32, bool)`

GetMinBalanceOk returns a tuple with the MinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBalance

`func (o *TenantWalletSettingsDto) SetMinBalance(v int32)`

SetMinBalance sets MinBalance field to given value.

### HasMinBalance

`func (o *TenantWalletSettingsDto) HasMinBalance() bool`

HasMinBalance returns a boolean if a field has been set.

### GetUpToBalance

`func (o *TenantWalletSettingsDto) GetUpToBalance() int32`

GetUpToBalance returns the UpToBalance field if non-nil, zero value otherwise.

### GetUpToBalanceOk

`func (o *TenantWalletSettingsDto) GetUpToBalanceOk() (*int32, bool)`

GetUpToBalanceOk returns a tuple with the UpToBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToBalance

`func (o *TenantWalletSettingsDto) SetUpToBalance(v int32)`

SetUpToBalance sets UpToBalance field to given value.

### HasUpToBalance

`func (o *TenantWalletSettingsDto) HasUpToBalance() bool`

HasUpToBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *TenantWalletSettingsDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TenantWalletSettingsDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TenantWalletSettingsDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TenantWalletSettingsDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *TenantWalletSettingsDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *TenantWalletSettingsDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetLowBalanceThreshold

`func (o *TenantWalletSettingsDto) GetLowBalanceThreshold() int32`

GetLowBalanceThreshold returns the LowBalanceThreshold field if non-nil, zero value otherwise.

### GetLowBalanceThresholdOk

`func (o *TenantWalletSettingsDto) GetLowBalanceThresholdOk() (*int32, bool)`

GetLowBalanceThresholdOk returns a tuple with the LowBalanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBalanceThreshold

`func (o *TenantWalletSettingsDto) SetLowBalanceThreshold(v int32)`

SetLowBalanceThreshold sets LowBalanceThreshold field to given value.

### HasLowBalanceThreshold

`func (o *TenantWalletSettingsDto) HasLowBalanceThreshold() bool`

HasLowBalanceThreshold returns a boolean if a field has been set.

### GetLowBalanceNotified

`func (o *TenantWalletSettingsDto) GetLowBalanceNotified() bool`

GetLowBalanceNotified returns the LowBalanceNotified field if non-nil, zero value otherwise.

### GetLowBalanceNotifiedOk

`func (o *TenantWalletSettingsDto) GetLowBalanceNotifiedOk() (*bool, bool)`

GetLowBalanceNotifiedOk returns a tuple with the LowBalanceNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBalanceNotified

`func (o *TenantWalletSettingsDto) SetLowBalanceNotified(v bool)`

SetLowBalanceNotified sets LowBalanceNotified field to given value.

### HasLowBalanceNotified

`func (o *TenantWalletSettingsDto) HasLowBalanceNotified() bool`

HasLowBalanceNotified returns a boolean if a field has been set.

### GetLastModified

`func (o *TenantWalletSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantWalletSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantWalletSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantWalletSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


