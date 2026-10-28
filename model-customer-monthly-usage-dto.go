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

// checks if the CustomerMonthlyUsageDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerMonthlyUsageDto{}

// CustomerMonthlyUsageDto What the portal spent from its wallet in one calendar month, added up across every service.
type CustomerMonthlyUsageDto struct {
	// The year the month belongs to. Months are cut in the portal time zone, so a movement at the edge of a  month falls where the portal sees it and not where UTC does.
	Year *int32 `json:"year,omitempty"`
	// The month itself, January being 1. Only months that had spending appear at all, so a gap in the list is a  month with nothing in it rather than missing data.
	Month *int32 `json:"month,omitempty"`
	// The currency `totalAmount` is expressed in, as a three-letter ISO 4217 code - the accounting currency of  the wallet.
	Currency NullableString `json:"currency,omitempty"`
	// What the month came to across every service, as a positive amount spent rather than a signed balance.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// How many separate movements that total was added up from, for a client that wants to show the weight  behind a figure. The movements themselves are in `GET api/2.0/portal/payment/customer/operations`.
	OperationCount *int32 `json:"operationCount,omitempty"`
}

// NewCustomerMonthlyUsageDto instantiates a new CustomerMonthlyUsageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerMonthlyUsageDto() *CustomerMonthlyUsageDto {
	this := CustomerMonthlyUsageDto{}
	return &this
}

// NewCustomerMonthlyUsageDtoWithDefaults instantiates a new CustomerMonthlyUsageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerMonthlyUsageDtoWithDefaults() *CustomerMonthlyUsageDto {
	this := CustomerMonthlyUsageDto{}
	return &this
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *CustomerMonthlyUsageDto) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerMonthlyUsageDto) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *CustomerMonthlyUsageDto) IsYearSet() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *CustomerMonthlyUsageDto) SetYear(v int32) {
	o.Year = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *CustomerMonthlyUsageDto) GetMonth() int32 {
	if o == nil || IsNil(o.Month) {
		var ret int32
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerMonthlyUsageDto) GetMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *CustomerMonthlyUsageDto) IsMonthSet() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given int32 and assigns it to the Month field.
func (o *CustomerMonthlyUsageDto) SetMonth(v int32) {
	o.Month = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerMonthlyUsageDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerMonthlyUsageDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *CustomerMonthlyUsageDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *CustomerMonthlyUsageDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *CustomerMonthlyUsageDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *CustomerMonthlyUsageDto) UnsetCurrency() {
	o.Currency.Unset()
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *CustomerMonthlyUsageDto) GetTotalAmount() float64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret float64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerMonthlyUsageDto) GetTotalAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *CustomerMonthlyUsageDto) IsTotalAmountSet() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float64 and assigns it to the TotalAmount field.
func (o *CustomerMonthlyUsageDto) SetTotalAmount(v float64) {
	o.TotalAmount = &v
}

// GetOperationCount returns the OperationCount field value if set, zero value otherwise.
func (o *CustomerMonthlyUsageDto) GetOperationCount() int32 {
	if o == nil || IsNil(o.OperationCount) {
		var ret int32
		return ret
	}
	return *o.OperationCount
}

// GetOperationCountOk returns a tuple with the OperationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerMonthlyUsageDto) GetOperationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OperationCount) {
		return nil, false
	}
	return o.OperationCount, true
}

// HasOperationCount returns a boolean if a field has been set.
func (o *CustomerMonthlyUsageDto) IsOperationCountSet() bool {
	if o != nil && !IsNil(o.OperationCount) {
		return true
	}

	return false
}

// SetOperationCount gets a reference to the given int32 and assigns it to the OperationCount field.
func (o *CustomerMonthlyUsageDto) SetOperationCount(v int32) {
	o.OperationCount = &v
}

func (o CustomerMonthlyUsageDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerMonthlyUsageDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.OperationCount) {
		toSerialize["operationCount"] = o.OperationCount
	}
	return toSerialize, nil
}

type NullableCustomerMonthlyUsageDto struct {
	value *CustomerMonthlyUsageDto
	isSet bool
}

func (v NullableCustomerMonthlyUsageDto) Get() *CustomerMonthlyUsageDto {
	return v.value
}

func (v *NullableCustomerMonthlyUsageDto) Set(val *CustomerMonthlyUsageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerMonthlyUsageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerMonthlyUsageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerMonthlyUsageDto(val *CustomerMonthlyUsageDto) *NullableCustomerMonthlyUsageDto {
	return &NullableCustomerMonthlyUsageDto{value: val, isSet: true}
}

func (v NullableCustomerMonthlyUsageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerMonthlyUsageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

