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

// checks if the BalanceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceDto{}

// BalanceDto Represents a balance with an account number and a list of sub-accounts.
type BalanceDto struct {
	// The account number.
	AccountNumber *int32 `json:"accountNumber,omitempty"`
	// The sub-account number.
	SubAccountNumber *int32 `json:"subAccountNumber,omitempty"`
	// The account name.
	AccountName NullableString `json:"accountName,omitempty"`
	// The account currency.
	AccountCurrency NullableString `json:"accountCurrency,omitempty"`
	// A list of sub-accounts.
	SubAccounts []SubAccountDto `json:"subAccounts,omitempty"`
	// The most recent credit transaction applied to the account.
	LastCredit *TransactionInfoDto `json:"lastCredit,omitempty"`
}

// NewBalanceDto instantiates a new BalanceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceDto() *BalanceDto {
	this := BalanceDto{}
	return &this
}

// NewBalanceDtoWithDefaults instantiates a new BalanceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceDtoWithDefaults() *BalanceDto {
	this := BalanceDto{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *BalanceDto) GetAccountNumber() int32 {
	if o == nil || IsNil(o.AccountNumber) {
		var ret int32
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceDto) GetAccountNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *BalanceDto) IsAccountNumberSet() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given int32 and assigns it to the AccountNumber field.
func (o *BalanceDto) SetAccountNumber(v int32) {
	o.AccountNumber = &v
}

// GetSubAccountNumber returns the SubAccountNumber field value if set, zero value otherwise.
func (o *BalanceDto) GetSubAccountNumber() int32 {
	if o == nil || IsNil(o.SubAccountNumber) {
		var ret int32
		return ret
	}
	return *o.SubAccountNumber
}

// GetSubAccountNumberOk returns a tuple with the SubAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceDto) GetSubAccountNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.SubAccountNumber) {
		return nil, false
	}
	return o.SubAccountNumber, true
}

// HasSubAccountNumber returns a boolean if a field has been set.
func (o *BalanceDto) IsSubAccountNumberSet() bool {
	if o != nil && !IsNil(o.SubAccountNumber) {
		return true
	}

	return false
}

// SetSubAccountNumber gets a reference to the given int32 and assigns it to the SubAccountNumber field.
func (o *BalanceDto) SetSubAccountNumber(v int32) {
	o.SubAccountNumber = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceDto) GetAccountName() string {
	if o == nil || IsNil(o.AccountName.Get()) {
		var ret string
		return ret
	}
	return *o.AccountName.Get()
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceDto) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountName.Get(), o.AccountName.IsSet()
}

// HasAccountName returns a boolean if a field has been set.
func (o *BalanceDto) IsAccountNameSet() bool {
	if o != nil && o.AccountName.IsSet() {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given NullableString and assigns it to the AccountName field.
func (o *BalanceDto) SetAccountName(v string) {
	o.AccountName.Set(&v)
}
// SetAccountNameNil sets the value for AccountName to be an explicit nil
func (o *BalanceDto) SetAccountNameNil() {
	o.AccountName.Set(nil)
}

// UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
func (o *BalanceDto) UnsetAccountName() {
	o.AccountName.Unset()
}

// GetAccountCurrency returns the AccountCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceDto) GetAccountCurrency() string {
	if o == nil || IsNil(o.AccountCurrency.Get()) {
		var ret string
		return ret
	}
	return *o.AccountCurrency.Get()
}

// GetAccountCurrencyOk returns a tuple with the AccountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceDto) GetAccountCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountCurrency.Get(), o.AccountCurrency.IsSet()
}

// HasAccountCurrency returns a boolean if a field has been set.
func (o *BalanceDto) IsAccountCurrencySet() bool {
	if o != nil && o.AccountCurrency.IsSet() {
		return true
	}

	return false
}

// SetAccountCurrency gets a reference to the given NullableString and assigns it to the AccountCurrency field.
func (o *BalanceDto) SetAccountCurrency(v string) {
	o.AccountCurrency.Set(&v)
}
// SetAccountCurrencyNil sets the value for AccountCurrency to be an explicit nil
func (o *BalanceDto) SetAccountCurrencyNil() {
	o.AccountCurrency.Set(nil)
}

// UnsetAccountCurrency ensures that no value is present for AccountCurrency, not even an explicit nil
func (o *BalanceDto) UnsetAccountCurrency() {
	o.AccountCurrency.Unset()
}

// GetSubAccounts returns the SubAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceDto) GetSubAccounts() []SubAccountDto {
	if o == nil {
		var ret []SubAccountDto
		return ret
	}
	return o.SubAccounts
}

// GetSubAccountsOk returns a tuple with the SubAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceDto) GetSubAccountsOk() ([]SubAccountDto, bool) {
	if o == nil || IsNil(o.SubAccounts) {
		return nil, false
	}
	return o.SubAccounts, true
}

// HasSubAccounts returns a boolean if a field has been set.
func (o *BalanceDto) IsSubAccountsSet() bool {
	if o != nil && !IsNil(o.SubAccounts) {
		return true
	}

	return false
}

// SetSubAccounts gets a reference to the given []SubAccountDto and assigns it to the SubAccounts field.
func (o *BalanceDto) SetSubAccounts(v []SubAccountDto) {
	o.SubAccounts = v
}

// GetLastCredit returns the LastCredit field value if set, zero value otherwise.
func (o *BalanceDto) GetLastCredit() TransactionInfoDto {
	if o == nil || IsNil(o.LastCredit) {
		var ret TransactionInfoDto
		return ret
	}
	return *o.LastCredit
}

// GetLastCreditOk returns a tuple with the LastCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceDto) GetLastCreditOk() (*TransactionInfoDto, bool) {
	if o == nil || IsNil(o.LastCredit) {
		return nil, false
	}
	return o.LastCredit, true
}

// HasLastCredit returns a boolean if a field has been set.
func (o *BalanceDto) IsLastCreditSet() bool {
	if o != nil && !IsNil(o.LastCredit) {
		return true
	}

	return false
}

// SetLastCredit gets a reference to the given TransactionInfoDto and assigns it to the LastCredit field.
func (o *BalanceDto) SetLastCredit(v TransactionInfoDto) {
	o.LastCredit = &v
}

func (o BalanceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceDto) ToMap() (map[string]interface{}, error) {
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

type NullableBalanceDto struct {
	value *BalanceDto
	isSet bool
}

func (v NullableBalanceDto) Get() *BalanceDto {
	return v.value
}

func (v *NullableBalanceDto) Set(val *BalanceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceDto(val *BalanceDto) *NullableBalanceDto {
	return &NullableBalanceDto{value: val, isSet: true}
}

func (v NullableBalanceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

