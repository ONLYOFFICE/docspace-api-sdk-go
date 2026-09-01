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

// checks if the TenantWalletSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantWalletSettings{}

// TenantWalletSettings The tenant wallet settings.
type TenantWalletSettings struct {
	// Specifies whether automatic top-up for the tenant wallet is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The minimum wallet balance at which automatic top-up will be triggered. Must be between 5 and 1000.
	MinBalance *int32 `json:"minBalance,omitempty"`
	// The maximum wallet balance at which automatic top-up will be triggered. Must be between 6 and 5000.
	UpToBalance *int32 `json:"upToBalance,omitempty"`
	// The three-character ISO 4217 currency symbol.
	Currency NullableString `json:"currency,omitempty"`
	// The wallet balance below which a low-balance notification is sent. Set internally, not user-configurable.
	LowBalanceThreshold *int32 `json:"lowBalanceThreshold,omitempty"`
	// Specifies whether a low-balance notification has already been sent for the current dip below ASC.Core.Tenants.TenantWalletSettings.LowBalanceThreshold.
	LowBalanceNotified *bool `json:"lowBalanceNotified,omitempty"`
	// The date and time when the tenant wallet settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewTenantWalletSettings instantiates a new TenantWalletSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantWalletSettings() *TenantWalletSettings {
	this := TenantWalletSettings{}
	return &this
}

// NewTenantWalletSettingsWithDefaults instantiates a new TenantWalletSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWalletSettingsWithDefaults() *TenantWalletSettings {
	this := TenantWalletSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TenantWalletSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWalletSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TenantWalletSettings) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TenantWalletSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMinBalance returns the MinBalance field value if set, zero value otherwise.
func (o *TenantWalletSettings) GetMinBalance() int32 {
	if o == nil || IsNil(o.MinBalance) {
		var ret int32
		return ret
	}
	return *o.MinBalance
}

// GetMinBalanceOk returns a tuple with the MinBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWalletSettings) GetMinBalanceOk() (*int32, bool) {
	if o == nil || IsNil(o.MinBalance) {
		return nil, false
	}
	return o.MinBalance, true
}

// HasMinBalance returns a boolean if a field has been set.
func (o *TenantWalletSettings) IsMinBalanceSet() bool {
	if o != nil && !IsNil(o.MinBalance) {
		return true
	}

	return false
}

// SetMinBalance gets a reference to the given int32 and assigns it to the MinBalance field.
func (o *TenantWalletSettings) SetMinBalance(v int32) {
	o.MinBalance = &v
}

// GetUpToBalance returns the UpToBalance field value if set, zero value otherwise.
func (o *TenantWalletSettings) GetUpToBalance() int32 {
	if o == nil || IsNil(o.UpToBalance) {
		var ret int32
		return ret
	}
	return *o.UpToBalance
}

// GetUpToBalanceOk returns a tuple with the UpToBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWalletSettings) GetUpToBalanceOk() (*int32, bool) {
	if o == nil || IsNil(o.UpToBalance) {
		return nil, false
	}
	return o.UpToBalance, true
}

// HasUpToBalance returns a boolean if a field has been set.
func (o *TenantWalletSettings) IsUpToBalanceSet() bool {
	if o != nil && !IsNil(o.UpToBalance) {
		return true
	}

	return false
}

// SetUpToBalance gets a reference to the given int32 and assigns it to the UpToBalance field.
func (o *TenantWalletSettings) SetUpToBalance(v int32) {
	o.UpToBalance = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantWalletSettings) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantWalletSettings) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *TenantWalletSettings) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *TenantWalletSettings) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *TenantWalletSettings) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *TenantWalletSettings) UnsetCurrency() {
	o.Currency.Unset()
}

// GetLowBalanceThreshold returns the LowBalanceThreshold field value if set, zero value otherwise.
func (o *TenantWalletSettings) GetLowBalanceThreshold() int32 {
	if o == nil || IsNil(o.LowBalanceThreshold) {
		var ret int32
		return ret
	}
	return *o.LowBalanceThreshold
}

// GetLowBalanceThresholdOk returns a tuple with the LowBalanceThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWalletSettings) GetLowBalanceThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.LowBalanceThreshold) {
		return nil, false
	}
	return o.LowBalanceThreshold, true
}

// HasLowBalanceThreshold returns a boolean if a field has been set.
func (o *TenantWalletSettings) IsLowBalanceThresholdSet() bool {
	if o != nil && !IsNil(o.LowBalanceThreshold) {
		return true
	}

	return false
}

// SetLowBalanceThreshold gets a reference to the given int32 and assigns it to the LowBalanceThreshold field.
func (o *TenantWalletSettings) SetLowBalanceThreshold(v int32) {
	o.LowBalanceThreshold = &v
}

// GetLowBalanceNotified returns the LowBalanceNotified field value if set, zero value otherwise.
func (o *TenantWalletSettings) GetLowBalanceNotified() bool {
	if o == nil || IsNil(o.LowBalanceNotified) {
		var ret bool
		return ret
	}
	return *o.LowBalanceNotified
}

// GetLowBalanceNotifiedOk returns a tuple with the LowBalanceNotified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWalletSettings) GetLowBalanceNotifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.LowBalanceNotified) {
		return nil, false
	}
	return o.LowBalanceNotified, true
}

// HasLowBalanceNotified returns a boolean if a field has been set.
func (o *TenantWalletSettings) IsLowBalanceNotifiedSet() bool {
	if o != nil && !IsNil(o.LowBalanceNotified) {
		return true
	}

	return false
}

// SetLowBalanceNotified gets a reference to the given bool and assigns it to the LowBalanceNotified field.
func (o *TenantWalletSettings) SetLowBalanceNotified(v bool) {
	o.LowBalanceNotified = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *TenantWalletSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWalletSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *TenantWalletSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *TenantWalletSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o TenantWalletSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantWalletSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MinBalance) {
		toSerialize["minBalance"] = o.MinBalance
	}
	if !IsNil(o.UpToBalance) {
		toSerialize["upToBalance"] = o.UpToBalance
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.LowBalanceThreshold) {
		toSerialize["lowBalanceThreshold"] = o.LowBalanceThreshold
	}
	if !IsNil(o.LowBalanceNotified) {
		toSerialize["lowBalanceNotified"] = o.LowBalanceNotified
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableTenantWalletSettings struct {
	value *TenantWalletSettings
	isSet bool
}

func (v NullableTenantWalletSettings) Get() *TenantWalletSettings {
	return v.value
}

func (v *NullableTenantWalletSettings) Set(val *TenantWalletSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantWalletSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantWalletSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantWalletSettings(val *TenantWalletSettings) *NullableTenantWalletSettings {
	return &NullableTenantWalletSettings{value: val, isSet: true}
}

func (v NullableTenantWalletSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantWalletSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

