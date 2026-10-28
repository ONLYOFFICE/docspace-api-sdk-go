# SetWalletTopUpSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the payment method on file is charged automatically when the wallet balance runs low. | [optional] 
**MinBalance** | Pointer to **int32** | The balance below which a top-up is charged, in `currency`. | [optional] 
**UpToBalance** | Pointer to **int32** | The balance a top-up brings the wallet up to, in `currency`. | [optional] 
**Currency** | Pointer to **NullableString** | The three-letter ISO 4217 code both amounts are expressed in; it has to be the currency of the wallet. | [optional] 
**LowBalanceThreshold** | Pointer to **int32** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**LowBalanceNotified** | Pointer to **bool** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**LastModified** | Pointer to **time.Time** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 

## Methods

### NewSetWalletTopUpSettingsRequest

`func NewSetWalletTopUpSettingsRequest() *SetWalletTopUpSettingsRequest`

NewSetWalletTopUpSettingsRequest instantiates a new SetWalletTopUpSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetWalletTopUpSettingsRequestWithDefaults

`func NewSetWalletTopUpSettingsRequestWithDefaults() *SetWalletTopUpSettingsRequest`

NewSetWalletTopUpSettingsRequestWithDefaults instantiates a new SetWalletTopUpSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SetWalletTopUpSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SetWalletTopUpSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SetWalletTopUpSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SetWalletTopUpSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinBalance

`func (o *SetWalletTopUpSettingsRequest) GetMinBalance() int32`

GetMinBalance returns the MinBalance field if non-nil, zero value otherwise.

### GetMinBalanceOk

`func (o *SetWalletTopUpSettingsRequest) GetMinBalanceOk() (*int32, bool)`

GetMinBalanceOk returns a tuple with the MinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBalance

`func (o *SetWalletTopUpSettingsRequest) SetMinBalance(v int32)`

SetMinBalance sets MinBalance field to given value.

### HasMinBalance

`func (o *SetWalletTopUpSettingsRequest) HasMinBalance() bool`

HasMinBalance returns a boolean if a field has been set.

### GetUpToBalance

`func (o *SetWalletTopUpSettingsRequest) GetUpToBalance() int32`

GetUpToBalance returns the UpToBalance field if non-nil, zero value otherwise.

### GetUpToBalanceOk

`func (o *SetWalletTopUpSettingsRequest) GetUpToBalanceOk() (*int32, bool)`

GetUpToBalanceOk returns a tuple with the UpToBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToBalance

`func (o *SetWalletTopUpSettingsRequest) SetUpToBalance(v int32)`

SetUpToBalance sets UpToBalance field to given value.

### HasUpToBalance

`func (o *SetWalletTopUpSettingsRequest) HasUpToBalance() bool`

HasUpToBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *SetWalletTopUpSettingsRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SetWalletTopUpSettingsRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SetWalletTopUpSettingsRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SetWalletTopUpSettingsRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *SetWalletTopUpSettingsRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *SetWalletTopUpSettingsRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetLowBalanceThreshold

`func (o *SetWalletTopUpSettingsRequest) GetLowBalanceThreshold() int32`

GetLowBalanceThreshold returns the LowBalanceThreshold field if non-nil, zero value otherwise.

### GetLowBalanceThresholdOk

`func (o *SetWalletTopUpSettingsRequest) GetLowBalanceThresholdOk() (*int32, bool)`

GetLowBalanceThresholdOk returns a tuple with the LowBalanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBalanceThreshold

`func (o *SetWalletTopUpSettingsRequest) SetLowBalanceThreshold(v int32)`

SetLowBalanceThreshold sets LowBalanceThreshold field to given value.

### HasLowBalanceThreshold

`func (o *SetWalletTopUpSettingsRequest) HasLowBalanceThreshold() bool`

HasLowBalanceThreshold returns a boolean if a field has been set.

### GetLowBalanceNotified

`func (o *SetWalletTopUpSettingsRequest) GetLowBalanceNotified() bool`

GetLowBalanceNotified returns the LowBalanceNotified field if non-nil, zero value otherwise.

### GetLowBalanceNotifiedOk

`func (o *SetWalletTopUpSettingsRequest) GetLowBalanceNotifiedOk() (*bool, bool)`

GetLowBalanceNotifiedOk returns a tuple with the LowBalanceNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBalanceNotified

`func (o *SetWalletTopUpSettingsRequest) SetLowBalanceNotified(v bool)`

SetLowBalanceNotified sets LowBalanceNotified field to given value.

### HasLowBalanceNotified

`func (o *SetWalletTopUpSettingsRequest) HasLowBalanceNotified() bool`

HasLowBalanceNotified returns a boolean if a field has been set.

### GetLastModified

`func (o *SetWalletTopUpSettingsRequest) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *SetWalletTopUpSettingsRequest) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *SetWalletTopUpSettingsRequest) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *SetWalletTopUpSettingsRequest) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


