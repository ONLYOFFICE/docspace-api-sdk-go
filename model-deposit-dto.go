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

// checks if the DepositDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepositDto{}

// DepositDto One payment into the portal wallet: when it was made, how much it brought in, how it ended, and whether an  invoice stands behind it.
type DepositDto struct {
	// The identifier of the top-up, which is what `GET api/2.0/portal/payment/customer/deposits/{id}/invoice` takes  to read its invoice.
	Id *int64 `json:"id,omitempty"`
	// When the top-up was booked, as an ISO 8601 timestamp that carries its own offset: that of the portal's time  zone, written as `Z` for a portal in UTC - the zone in which `startDate` and `endDate` without an offset are  read.
	Date *ApiDateTime `json:"date,omitempty"`
	// The amount that went into the wallet, in `currency`.
	Amount *float64 `json:"amount,omitempty"`
	// The currency `amount` is expressed in, as a three-letter ISO 4217 code - the accounting currency of the  wallet.
	Currency NullableString `json:"currency,omitempty"`
	// How the top-up ended. A pending one is still being settled and may change later, so only a completed one has  actually raised the balance.
	Status *OperationStatus `json:"status,omitempty"`
	// Whether the billing provider holds an invoice for the payment behind the top-up, which is when  `GET api/2.0/portal/payment/customer/deposits/{id}/invoice` can open it. It is `false` for a top-up whose  payment came without one, and asking for its invoice then fails with 404.
	HasInvoice *bool `json:"hasInvoice,omitempty"`
}

// NewDepositDto instantiates a new DepositDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositDto() *DepositDto {
	this := DepositDto{}
	return &this
}

// NewDepositDtoWithDefaults instantiates a new DepositDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositDtoWithDefaults() *DepositDto {
	this := DepositDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DepositDto) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDto) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DepositDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DepositDto) SetId(v int64) {
	o.Id = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *DepositDto) GetDate() ApiDateTime {
	if o == nil || IsNil(o.Date) {
		var ret ApiDateTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDto) GetDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *DepositDto) IsDateSet() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given ApiDateTime and assigns it to the Date field.
func (o *DepositDto) SetDate(v ApiDateTime) {
	o.Date = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DepositDto) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDto) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DepositDto) IsAmountSet() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *DepositDto) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *DepositDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *DepositDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *DepositDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *DepositDto) UnsetCurrency() {
	o.Currency.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DepositDto) GetStatus() OperationStatus {
	if o == nil || IsNil(o.Status) {
		var ret OperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDto) GetStatusOk() (*OperationStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DepositDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OperationStatus and assigns it to the Status field.
func (o *DepositDto) SetStatus(v OperationStatus) {
	o.Status = &v
}

// GetHasInvoice returns the HasInvoice field value if set, zero value otherwise.
func (o *DepositDto) GetHasInvoice() bool {
	if o == nil || IsNil(o.HasInvoice) {
		var ret bool
		return ret
	}
	return *o.HasInvoice
}

// GetHasInvoiceOk returns a tuple with the HasInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDto) GetHasInvoiceOk() (*bool, bool) {
	if o == nil || IsNil(o.HasInvoice) {
		return nil, false
	}
	return o.HasInvoice, true
}

// HasHasInvoice returns a boolean if a field has been set.
func (o *DepositDto) IsHasInvoiceSet() bool {
	if o != nil && !IsNil(o.HasInvoice) {
		return true
	}

	return false
}

// SetHasInvoice gets a reference to the given bool and assigns it to the HasInvoice field.
func (o *DepositDto) SetHasInvoice(v bool) {
	o.HasInvoice = &v
}

func (o DepositDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.HasInvoice) {
		toSerialize["hasInvoice"] = o.HasInvoice
	}
	return toSerialize, nil
}

type NullableDepositDto struct {
	value *DepositDto
	isSet bool
}

func (v NullableDepositDto) Get() *DepositDto {
	return v.value
}

func (v *NullableDepositDto) Set(val *DepositDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositDto(val *DepositDto) *NullableDepositDto {
	return &NullableDepositDto{value: val, isSet: true}
}

func (v NullableDepositDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

