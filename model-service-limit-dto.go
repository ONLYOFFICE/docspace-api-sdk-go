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

// checks if the ServiceLimitDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceLimitDto{}

// ServiceLimitDto A spending limit on the AI services: for the portal as a whole, for every user, or for one user.
type ServiceLimitDto struct {
	// The service limit ID, used to read, change, switch off or delete the limit. In the list of the users the same  ID repeats for every user the limit for every user applies to.
	Id *int32 `json:"id,omitempty"`
	// The limited service group; the AI services are `ai-services`.
	ServiceGroupName NullableString `json:"serviceGroupName,omitempty"`
	// The number of the billing account of the portal the limit belongs to.
	CustomerAccountNumber *int32 `json:"customerAccountNumber,omitempty"`
	// Whom the limit applies to.
	Scope *ServiceLimitScope `json:"scope,omitempty"`
	// The ID of the portal user the limit applies to: the user of a personal limit, or, in the list of the users, the  user a limit for every user applies to. Null for a limit of the portal as a whole, and for a limit for every  user read on its own.
	UserId NullableString `json:"userId,omitempty"`
	// The most money the AI services may cost per period, in `currency`; null when the limit has no money threshold.
	AmountValue NullableFloat64 `json:"amountValue,omitempty"`
	// The currency `amountValue` is expressed in; null when the limit has no money threshold.
	Currency NullableString `json:"currency,omitempty"`
	// The largest quantity the AI services may consume per period, in `serviceUnit`; null when the limit has no  quantity threshold.
	QuantityValue NullableInt32 `json:"quantityValue,omitempty"`
	// The unit `quantityValue` is expressed in; null when the limit has no quantity threshold.
	ServiceUnit NullableString `json:"serviceUnit,omitempty"`
	// The window the thresholds apply to. A day, week or month resets at the start of the next one in UTC (a week on  Monday), not a fixed time after the first charge; a lifetime limit never resets.
	Period *ServiceLimitPeriod `json:"period,omitempty"`
	// The last moment of the current period, in UTC, after which the spending starts over from zero; null for a  lifetime limit, which never resets. Filled for every limit, the limit for every user included, whose spending is  counted per user.
	PeriodEnd NullableTime `json:"periodEnd,omitempty"`
	// The moment a lifetime limit counts from, in UTC; null for the other periods.
	UsageFrom NullableTime `json:"usageFrom,omitempty"`
	// Whether the limit is in force. A switched-off limit no longer restricts the AI services and can be switched  back on.
	Enabled *bool `json:"enabled,omitempty"`
	// The date and time when the limit was created, in UTC.
	Created *time.Time `json:"created,omitempty"`
	// The date and time when the limit was last changed, in UTC.
	Modified *time.Time `json:"modified,omitempty"`
	// What was spent against the limit in its current period and what is left. Filled for the limit of the portal as  a whole and in the list of the users; null where the spending is not read, such as for a limit read by its ID  or the limit for every user, whose spending is counted per user.
	Usage *ServiceLimitUsageDto `json:"usage,omitempty"`
}

// NewServiceLimitDto instantiates a new ServiceLimitDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLimitDto() *ServiceLimitDto {
	this := ServiceLimitDto{}
	return &this
}

// NewServiceLimitDtoWithDefaults instantiates a new ServiceLimitDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLimitDtoWithDefaults() *ServiceLimitDto {
	this := ServiceLimitDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceLimitDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ServiceLimitDto) SetId(v int32) {
	o.Id = &v
}

// GetServiceGroupName returns the ServiceGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitDto) GetServiceGroupName() string {
	if o == nil || IsNil(o.ServiceGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceGroupName.Get()
}

// GetServiceGroupNameOk returns a tuple with the ServiceGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitDto) GetServiceGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceGroupName.Get(), o.ServiceGroupName.IsSet()
}

// HasServiceGroupName returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsServiceGroupNameSet() bool {
	if o != nil && o.ServiceGroupName.IsSet() {
		return true
	}

	return false
}

// SetServiceGroupName gets a reference to the given NullableString and assigns it to the ServiceGroupName field.
func (o *ServiceLimitDto) SetServiceGroupName(v string) {
	o.ServiceGroupName.Set(&v)
}
// SetServiceGroupNameNil sets the value for ServiceGroupName to be an explicit nil
func (o *ServiceLimitDto) SetServiceGroupNameNil() {
	o.ServiceGroupName.Set(nil)
}

// UnsetServiceGroupName ensures that no value is present for ServiceGroupName, not even an explicit nil
func (o *ServiceLimitDto) UnsetServiceGroupName() {
	o.ServiceGroupName.Unset()
}

// GetCustomerAccountNumber returns the CustomerAccountNumber field value if set, zero value otherwise.
func (o *ServiceLimitDto) GetCustomerAccountNumber() int32 {
	if o == nil || IsNil(o.CustomerAccountNumber) {
		var ret int32
		return ret
	}
	return *o.CustomerAccountNumber
}

// GetCustomerAccountNumberOk returns a tuple with the CustomerAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitDto) GetCustomerAccountNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomerAccountNumber) {
		return nil, false
	}
	return o.CustomerAccountNumber, true
}

// HasCustomerAccountNumber returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsCustomerAccountNumberSet() bool {
	if o != nil && !IsNil(o.CustomerAccountNumber) {
		return true
	}

	return false
}

// SetCustomerAccountNumber gets a reference to the given int32 and assigns it to the CustomerAccountNumber field.
func (o *ServiceLimitDto) SetCustomerAccountNumber(v int32) {
	o.CustomerAccountNumber = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ServiceLimitDto) GetScope() ServiceLimitScope {
	if o == nil || IsNil(o.Scope) {
		var ret ServiceLimitScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitDto) GetScopeOk() (*ServiceLimitScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsScopeSet() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given ServiceLimitScope and assigns it to the Scope field.
func (o *ServiceLimitDto) SetScope(v ServiceLimitScope) {
	o.Scope = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsUserIdSet() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *ServiceLimitDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *ServiceLimitDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *ServiceLimitDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetAmountValue returns the AmountValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitDto) GetAmountValue() float64 {
	if o == nil || IsNil(o.AmountValue.Get()) {
		var ret float64
		return ret
	}
	return *o.AmountValue.Get()
}

// GetAmountValueOk returns a tuple with the AmountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitDto) GetAmountValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountValue.Get(), o.AmountValue.IsSet()
}

// HasAmountValue returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsAmountValueSet() bool {
	if o != nil && o.AmountValue.IsSet() {
		return true
	}

	return false
}

// SetAmountValue gets a reference to the given NullableFloat64 and assigns it to the AmountValue field.
func (o *ServiceLimitDto) SetAmountValue(v float64) {
	o.AmountValue.Set(&v)
}
// SetAmountValueNil sets the value for AmountValue to be an explicit nil
func (o *ServiceLimitDto) SetAmountValueNil() {
	o.AmountValue.Set(nil)
}

// UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
func (o *ServiceLimitDto) UnsetAmountValue() {
	o.AmountValue.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *ServiceLimitDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *ServiceLimitDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *ServiceLimitDto) UnsetCurrency() {
	o.Currency.Unset()
}

// GetQuantityValue returns the QuantityValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitDto) GetQuantityValue() int32 {
	if o == nil || IsNil(o.QuantityValue.Get()) {
		var ret int32
		return ret
	}
	return *o.QuantityValue.Get()
}

// GetQuantityValueOk returns a tuple with the QuantityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitDto) GetQuantityValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuantityValue.Get(), o.QuantityValue.IsSet()
}

// HasQuantityValue returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsQuantityValueSet() bool {
	if o != nil && o.QuantityValue.IsSet() {
		return true
	}

	return false
}

// SetQuantityValue gets a reference to the given NullableInt32 and assigns it to the QuantityValue field.
func (o *ServiceLimitDto) SetQuantityValue(v int32) {
	o.QuantityValue.Set(&v)
}
// SetQuantityValueNil sets the value for QuantityValue to be an explicit nil
func (o *ServiceLimitDto) SetQuantityValueNil() {
	o.QuantityValue.Set(nil)
}

// UnsetQuantityValue ensures that no value is present for QuantityValue, not even an explicit nil
func (o *ServiceLimitDto) UnsetQuantityValue() {
	o.QuantityValue.Unset()
}

// GetServiceUnit returns the ServiceUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitDto) GetServiceUnit() string {
	if o == nil || IsNil(o.ServiceUnit.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceUnit.Get()
}

// GetServiceUnitOk returns a tuple with the ServiceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitDto) GetServiceUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceUnit.Get(), o.ServiceUnit.IsSet()
}

// HasServiceUnit returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsServiceUnitSet() bool {
	if o != nil && o.ServiceUnit.IsSet() {
		return true
	}

	return false
}

// SetServiceUnit gets a reference to the given NullableString and assigns it to the ServiceUnit field.
func (o *ServiceLimitDto) SetServiceUnit(v string) {
	o.ServiceUnit.Set(&v)
}
// SetServiceUnitNil sets the value for ServiceUnit to be an explicit nil
func (o *ServiceLimitDto) SetServiceUnitNil() {
	o.ServiceUnit.Set(nil)
}

// UnsetServiceUnit ensures that no value is present for ServiceUnit, not even an explicit nil
func (o *ServiceLimitDto) UnsetServiceUnit() {
	o.ServiceUnit.Unset()
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ServiceLimitDto) GetPeriod() ServiceLimitPeriod {
	if o == nil || IsNil(o.Period) {
		var ret ServiceLimitPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitDto) GetPeriodOk() (*ServiceLimitPeriod, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsPeriodSet() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given ServiceLimitPeriod and assigns it to the Period field.
func (o *ServiceLimitDto) SetPeriod(v ServiceLimitPeriod) {
	o.Period = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitDto) GetPeriodEnd() time.Time {
	if o == nil || IsNil(o.PeriodEnd.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PeriodEnd.Get()
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitDto) GetPeriodEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeriodEnd.Get(), o.PeriodEnd.IsSet()
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsPeriodEndSet() bool {
	if o != nil && o.PeriodEnd.IsSet() {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given NullableTime and assigns it to the PeriodEnd field.
func (o *ServiceLimitDto) SetPeriodEnd(v time.Time) {
	o.PeriodEnd.Set(&v)
}
// SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil
func (o *ServiceLimitDto) SetPeriodEndNil() {
	o.PeriodEnd.Set(nil)
}

// UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
func (o *ServiceLimitDto) UnsetPeriodEnd() {
	o.PeriodEnd.Unset()
}

// GetUsageFrom returns the UsageFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitDto) GetUsageFrom() time.Time {
	if o == nil || IsNil(o.UsageFrom.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UsageFrom.Get()
}

// GetUsageFromOk returns a tuple with the UsageFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitDto) GetUsageFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageFrom.Get(), o.UsageFrom.IsSet()
}

// HasUsageFrom returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsUsageFromSet() bool {
	if o != nil && o.UsageFrom.IsSet() {
		return true
	}

	return false
}

// SetUsageFrom gets a reference to the given NullableTime and assigns it to the UsageFrom field.
func (o *ServiceLimitDto) SetUsageFrom(v time.Time) {
	o.UsageFrom.Set(&v)
}
// SetUsageFromNil sets the value for UsageFrom to be an explicit nil
func (o *ServiceLimitDto) SetUsageFromNil() {
	o.UsageFrom.Set(nil)
}

// UnsetUsageFrom ensures that no value is present for UsageFrom, not even an explicit nil
func (o *ServiceLimitDto) UnsetUsageFrom() {
	o.UsageFrom.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ServiceLimitDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ServiceLimitDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ServiceLimitDto) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitDto) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ServiceLimitDto) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ServiceLimitDto) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitDto) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsModifiedSet() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ServiceLimitDto) SetModified(v time.Time) {
	o.Modified = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ServiceLimitDto) GetUsage() ServiceLimitUsageDto {
	if o == nil || IsNil(o.Usage) {
		var ret ServiceLimitUsageDto
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitDto) GetUsageOk() (*ServiceLimitUsageDto, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ServiceLimitDto) IsUsageSet() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given ServiceLimitUsageDto and assigns it to the Usage field.
func (o *ServiceLimitDto) SetUsage(v ServiceLimitUsageDto) {
	o.Usage = &v
}

func (o ServiceLimitDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceLimitDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ServiceGroupName.IsSet() {
		toSerialize["serviceGroupName"] = o.ServiceGroupName.Get()
	}
	if !IsNil(o.CustomerAccountNumber) {
		toSerialize["customerAccountNumber"] = o.CustomerAccountNumber
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.AmountValue.IsSet() {
		toSerialize["amountValue"] = o.AmountValue.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.QuantityValue.IsSet() {
		toSerialize["quantityValue"] = o.QuantityValue.Get()
	}
	if o.ServiceUnit.IsSet() {
		toSerialize["serviceUnit"] = o.ServiceUnit.Get()
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if o.PeriodEnd.IsSet() {
		toSerialize["periodEnd"] = o.PeriodEnd.Get()
	}
	if o.UsageFrom.IsSet() {
		toSerialize["usageFrom"] = o.UsageFrom.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableServiceLimitDto struct {
	value *ServiceLimitDto
	isSet bool
}

func (v NullableServiceLimitDto) Get() *ServiceLimitDto {
	return v.value
}

func (v *NullableServiceLimitDto) Set(val *ServiceLimitDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLimitDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLimitDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLimitDto(val *ServiceLimitDto) *NullableServiceLimitDto {
	return &NullableServiceLimitDto{value: val, isSet: true}
}

func (v NullableServiceLimitDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLimitDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

