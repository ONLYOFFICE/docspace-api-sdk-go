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

// checks if the UpcomingPaymentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpcomingPaymentDto{}

// UpcomingPaymentDto One charge the portal is going to be billed for at the start of the next period.
type UpcomingPaymentDto struct {
	// The quota that is going to be charged. When a switch to another quota is scheduled, this is the quota  being switched to, so it can differ from what `GET api/2.0/portal/tariff` reports for today.
	Id *int32 `json:"id,omitempty"`
	// The quota's stable key, which is the same identifier the wallet operations use for a service.
	Name NullableString `json:"name,omitempty"`
	// The quota name in the portal language, meant to be printed on an invoice preview.
	Title NullableString `json:"title,omitempty"`
	// What `quantity` counts, in the portal language - seats, administrators, gigabytes. It is empty for a quota  that is simply on or off.
	UnitOfMeasure NullableString `json:"unitOfMeasure,omitempty"`
	// How much is going to be charged for, which is the quantity scheduled for the next period when one has been  scheduled and today's quantity otherwise.
	Quantity *int32 `json:"quantity,omitempty"`
	// Whether the charge is paid out of the portal wallet rather than from the subscription.
	Wallet *bool `json:"wallet,omitempty"`
	// When the charge falls due, in the portal time zone.
	DueDate *ApiDateTime `json:"dueDate,omitempty"`
	// What the charge comes to: the unit price of the quota multiplied by `quantity`. Taxes are not part of it,  and a quota with no price of its own is not listed at all rather than listed with a zero.
	Amount *float64 `json:"amount,omitempty"`
	// The currency `amount` is expressed in, as a three-letter ISO 4217 code. It follows the portal's billing  account, so every entry of one answer carries the same code.
	Currency NullableString `json:"currency,omitempty"`
}

// NewUpcomingPaymentDto instantiates a new UpcomingPaymentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpcomingPaymentDto() *UpcomingPaymentDto {
	this := UpcomingPaymentDto{}
	return &this
}

// NewUpcomingPaymentDtoWithDefaults instantiates a new UpcomingPaymentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpcomingPaymentDtoWithDefaults() *UpcomingPaymentDto {
	this := UpcomingPaymentDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpcomingPaymentDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpcomingPaymentDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpcomingPaymentDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpcomingPaymentDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpcomingPaymentDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpcomingPaymentDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpcomingPaymentDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpcomingPaymentDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpcomingPaymentDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpcomingPaymentDto) UnsetName() {
	o.Name.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpcomingPaymentDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpcomingPaymentDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UpcomingPaymentDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UpcomingPaymentDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UpcomingPaymentDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UpcomingPaymentDto) UnsetTitle() {
	o.Title.Unset()
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpcomingPaymentDto) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure.Get()) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure.Get()
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpcomingPaymentDto) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitOfMeasure.Get(), o.UnitOfMeasure.IsSet()
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *UpcomingPaymentDto) IsUnitOfMeasureSet() bool {
	if o != nil && o.UnitOfMeasure.IsSet() {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given NullableString and assigns it to the UnitOfMeasure field.
func (o *UpcomingPaymentDto) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure.Set(&v)
}
// SetUnitOfMeasureNil sets the value for UnitOfMeasure to be an explicit nil
func (o *UpcomingPaymentDto) SetUnitOfMeasureNil() {
	o.UnitOfMeasure.Set(nil)
}

// UnsetUnitOfMeasure ensures that no value is present for UnitOfMeasure, not even an explicit nil
func (o *UpcomingPaymentDto) UnsetUnitOfMeasure() {
	o.UnitOfMeasure.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UpcomingPaymentDto) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpcomingPaymentDto) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UpcomingPaymentDto) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *UpcomingPaymentDto) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *UpcomingPaymentDto) GetWallet() bool {
	if o == nil || IsNil(o.Wallet) {
		var ret bool
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpcomingPaymentDto) GetWalletOk() (*bool, bool) {
	if o == nil || IsNil(o.Wallet) {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *UpcomingPaymentDto) IsWalletSet() bool {
	if o != nil && !IsNil(o.Wallet) {
		return true
	}

	return false
}

// SetWallet gets a reference to the given bool and assigns it to the Wallet field.
func (o *UpcomingPaymentDto) SetWallet(v bool) {
	o.Wallet = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *UpcomingPaymentDto) GetDueDate() ApiDateTime {
	if o == nil || IsNil(o.DueDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpcomingPaymentDto) GetDueDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *UpcomingPaymentDto) IsDueDateSet() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given ApiDateTime and assigns it to the DueDate field.
func (o *UpcomingPaymentDto) SetDueDate(v ApiDateTime) {
	o.DueDate = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UpcomingPaymentDto) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpcomingPaymentDto) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UpcomingPaymentDto) IsAmountSet() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *UpcomingPaymentDto) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpcomingPaymentDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpcomingPaymentDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *UpcomingPaymentDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *UpcomingPaymentDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *UpcomingPaymentDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *UpcomingPaymentDto) UnsetCurrency() {
	o.Currency.Unset()
}

func (o UpcomingPaymentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpcomingPaymentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.UnitOfMeasure.IsSet() {
		toSerialize["unitOfMeasure"] = o.UnitOfMeasure.Get()
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Wallet) {
		toSerialize["wallet"] = o.Wallet
	}
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	return toSerialize, nil
}

type NullableUpcomingPaymentDto struct {
	value *UpcomingPaymentDto
	isSet bool
}

func (v NullableUpcomingPaymentDto) Get() *UpcomingPaymentDto {
	return v.value
}

func (v *NullableUpcomingPaymentDto) Set(val *UpcomingPaymentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpcomingPaymentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpcomingPaymentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpcomingPaymentDto(val *UpcomingPaymentDto) *NullableUpcomingPaymentDto {
	return &NullableUpcomingPaymentDto{value: val, isSet: true}
}

func (v NullableUpcomingPaymentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpcomingPaymentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

