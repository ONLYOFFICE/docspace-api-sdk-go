# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **int32** | The account number. | [optional] 
**SubAccountNumber** | Pointer to **int32** | The sub-account number. | [optional] 
**AccountName** | Pointer to **NullableString** | The account name. | [optional] 
**AccountCurrency** | Pointer to **NullableString** | The account currency. | [optional] 
**SubAccounts** | Pointer to [**[]SubAccount**](SubAccount.md) | A list of sub-accounts. | [optional] 
**LastCredit** | Pointer to [**TransactionInfo**](TransactionInfo.md) | The most recent credit transaction applied to the account. | [optional] 

## Methods

### NewBalance

`func NewBalance() *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *Balance) GetAccountNumber() int32`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Balance) GetAccountNumberOk() (*int32, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Balance) SetAccountNumber(v int32)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Balance) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetSubAccountNumber

`func (o *Balance) GetSubAccountNumber() int32`

GetSubAccountNumber returns the SubAccountNumber field if non-nil, zero value otherwise.

### GetSubAccountNumberOk

`func (o *Balance) GetSubAccountNumberOk() (*int32, bool)`

GetSubAccountNumberOk returns a tuple with the SubAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountNumber

`func (o *Balance) SetSubAccountNumber(v int32)`

SetSubAccountNumber sets SubAccountNumber field to given value.

### HasSubAccountNumber

`func (o *Balance) HasSubAccountNumber() bool`

HasSubAccountNumber returns a boolean if a field has been set.

### GetAccountName

`func (o *Balance) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Balance) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Balance) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Balance) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *Balance) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *Balance) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetAccountCurrency

`func (o *Balance) GetAccountCurrency() string`

GetAccountCurrency returns the AccountCurrency field if non-nil, zero value otherwise.

### GetAccountCurrencyOk

`func (o *Balance) GetAccountCurrencyOk() (*string, bool)`

GetAccountCurrencyOk returns a tuple with the AccountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCurrency

`func (o *Balance) SetAccountCurrency(v string)`

SetAccountCurrency sets AccountCurrency field to given value.

### HasAccountCurrency

`func (o *Balance) HasAccountCurrency() bool`

HasAccountCurrency returns a boolean if a field has been set.

### SetAccountCurrencyNil

`func (o *Balance) SetAccountCurrencyNil(b bool)`

 SetAccountCurrencyNil sets the value for AccountCurrency to be an explicit nil

### UnsetAccountCurrency
`func (o *Balance) UnsetAccountCurrency()`

UnsetAccountCurrency ensures that no value is present for AccountCurrency, not even an explicit nil
### GetSubAccounts

`func (o *Balance) GetSubAccounts() []SubAccount`

GetSubAccounts returns the SubAccounts field if non-nil, zero value otherwise.

### GetSubAccountsOk

`func (o *Balance) GetSubAccountsOk() (*[]SubAccount, bool)`

GetSubAccountsOk returns a tuple with the SubAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccounts

`func (o *Balance) SetSubAccounts(v []SubAccount)`

SetSubAccounts sets SubAccounts field to given value.

### HasSubAccounts

`func (o *Balance) HasSubAccounts() bool`

HasSubAccounts returns a boolean if a field has been set.

### SetSubAccountsNil

`func (o *Balance) SetSubAccountsNil(b bool)`

 SetSubAccountsNil sets the value for SubAccounts to be an explicit nil

### UnsetSubAccounts
`func (o *Balance) UnsetSubAccounts()`

UnsetSubAccounts ensures that no value is present for SubAccounts, not even an explicit nil
### GetLastCredit

`func (o *Balance) GetLastCredit() TransactionInfo`

GetLastCredit returns the LastCredit field if non-nil, zero value otherwise.

### GetLastCreditOk

`func (o *Balance) GetLastCreditOk() (*TransactionInfo, bool)`

GetLastCreditOk returns a tuple with the LastCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCredit

`func (o *Balance) SetLastCredit(v TransactionInfo)`

SetLastCredit sets LastCredit field to given value.

### HasLastCredit

`func (o *Balance) HasLastCredit() bool`

HasLastCredit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


