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
)

// checks if the Balance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Balance{}

// Balance Represents a balance with an account number and a list of sub-accounts.
type Balance struct {
	// The account number.
	AccountNumber *int32 `json:"accountNumber,omitempty"`
	// The sub-account number.
	SubAccountNumber *int32 `json:"subAccountNumber,omitempty"`
	// The account name.
	AccountName NullableString `json:"accountName,omitempty"`
	// The account currency.
	AccountCurrency NullableString `json:"accountCurrency,omitempty"`
	// A list of sub-accounts.
	SubAccounts []SubAccount `json:"subAccounts,omitempty"`
	LastCredit *TransactionInfo `json:"lastCredit,omitempty"`
}

// NewBalance instantiates a new Balance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalance() *Balance {
	this := Balance{}
	return &this
}

// NewBalanceWithDefaults instantiates a new Balance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceWithDefaults() *Balance {
	this := Balance{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *Balance) GetAccountNumber() int32 {
	if o == nil || IsNil(o.AccountNumber) {
		var ret int32
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetAccountNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *Balance) IsAccountNumberSet() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given int32 and assigns it to the AccountNumber field.
func (o *Balance) SetAccountNumber(v int32) {
	o.AccountNumber = &v
}

// GetSubAccountNumber returns the SubAccountNumber field value if set, zero value otherwise.
func (o *Balance) GetSubAccountNumber() int32 {
	if o == nil || IsNil(o.SubAccountNumber) {
		var ret int32
		return ret
	}
	return *o.SubAccountNumber
}

// GetSubAccountNumberOk returns a tuple with the SubAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetSubAccountNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.SubAccountNumber) {
		return nil, false
	}
	return o.SubAccountNumber, true
}

// HasSubAccountNumber returns a boolean if a field has been set.
func (o *Balance) IsSubAccountNumberSet() bool {
	if o != nil && !IsNil(o.SubAccountNumber) {
		return true
	}

	return false
}

// SetSubAccountNumber gets a reference to the given int32 and assigns it to the SubAccountNumber field.
func (o *Balance) SetSubAccountNumber(v int32) {
	o.SubAccountNumber = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetAccountName() string {
	if o == nil || IsNil(o.AccountName.Get()) {
		var ret string
		return ret
	}
	return *o.AccountName.Get()
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountName.Get(), o.AccountName.IsSet()
}

// HasAccountName returns a boolean if a field has been set.
func (o *Balance) IsAccountNameSet() bool {
	if o != nil && o.AccountName.IsSet() {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given NullableString and assigns it to the AccountName field.
func (o *Balance) SetAccountName(v string) {
	o.AccountName.Set(&v)
}
// SetAccountNameNil sets the value for AccountName to be an explicit nil
func (o *Balance) SetAccountNameNil() {
	o.AccountName.Set(nil)
}

// UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
func (o *Balance) UnsetAccountName() {
	o.AccountName.Unset()
}

// GetAccountCurrency returns the AccountCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetAccountCurrency() string {
	if o == nil || IsNil(o.AccountCurrency.Get()) {
		var ret string
		return ret
	}
	return *o.AccountCurrency.Get()
}

// GetAccountCurrencyOk returns a tuple with the AccountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetAccountCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountCurrency.Get(), o.AccountCurrency.IsSet()
}

// HasAccountCurrency returns a boolean if a field has been set.
func (o *Balance) IsAccountCurrencySet() bool {
	if o != nil && o.AccountCurrency.IsSet() {
		return true
	}

	return false
}

// SetAccountCurrency gets a reference to the given NullableString and assigns it to the AccountCurrency field.
func (o *Balance) SetAccountCurrency(v string) {
	o.AccountCurrency.Set(&v)
}
// SetAccountCurrencyNil sets the value for AccountCurrency to be an explicit nil
func (o *Balance) SetAccountCurrencyNil() {
	o.AccountCurrency.Set(nil)
}

// UnsetAccountCurrency ensures that no value is present for AccountCurrency, not even an explicit nil
func (o *Balance) UnsetAccountCurrency() {
	o.AccountCurrency.Unset()
}

// GetSubAccounts returns the SubAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetSubAccounts() []SubAccount {
	if o == nil {
		var ret []SubAccount
		return ret
	}
	return o.SubAccounts
}

// GetSubAccountsOk returns a tuple with the SubAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetSubAccountsOk() ([]SubAccount, bool) {
	if o == nil || IsNil(o.SubAccounts) {
		return nil, false
	}
	return o.SubAccounts, true
}

// HasSubAccounts returns a boolean if a field has been set.
func (o *Balance) IsSubAccountsSet() bool {
	if o != nil && !IsNil(o.SubAccounts) {
		return true
	}

	return false
}

// SetSubAccounts gets a reference to the given []SubAccount and assigns it to the SubAccounts field.
func (o *Balance) SetSubAccounts(v []SubAccount) {
	o.SubAccounts = v
}

// GetLastCredit returns the LastCredit field value if set, zero value otherwise.
func (o *Balance) GetLastCredit() TransactionInfo {
	if o == nil || IsNil(o.LastCredit) {
		var ret TransactionInfo
		return ret
	}
	return *o.LastCredit
}

// GetLastCreditOk returns a tuple with the LastCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetLastCreditOk() (*TransactionInfo, bool) {
	if o == nil || IsNil(o.LastCredit) {
		return nil, false
	}
	return o.LastCredit, true
}

// HasLastCredit returns a boolean if a field has been set.
func (o *Balance) IsLastCreditSet() bool {
	if o != nil && !IsNil(o.LastCredit) {
		return true
	}

	return false
}

// SetLastCredit gets a reference to the given TransactionInfo and assigns it to the LastCredit field.
func (o *Balance) SetLastCredit(v TransactionInfo) {
	o.LastCredit = &v
}

func (o Balance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Balance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.SubAccountNumber) {
		toSerialize["subAccountNumber"] = o.SubAccountNumber
	}
	if o.AccountName.IsSet() {
		toSerialize["accountName"] = o.AccountName.Get()
	}
	if o.AccountCurrency.IsSet() {
		toSerialize["accountCurrency"] = o.AccountCurrency.Get()
	}
	if o.SubAccounts != nil {
		toSerialize["subAccounts"] = o.SubAccounts
	}
	if !IsNil(o.LastCredit) {
		toSerialize["lastCredit"] = o.LastCredit
	}
	return toSerialize, nil
}

type NullableBalance struct {
	value *Balance
	isSet bool
}

func (v NullableBalance) Get() *Balance {
	return v.value
}

func (v *NullableBalance) Set(val *Balance) {
	v.value = val
	v.isSet = true
}

func (v NullableBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalance(val *Balance) *NullableBalance {
	return &NullableBalance{value: val, isSet: true}
}

func (v NullableBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

