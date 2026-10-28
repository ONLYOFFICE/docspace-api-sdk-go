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

// checks if the DocsCloudPaymentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudPaymentDto{}

// DocsCloudPaymentDto Represents the payment information of a Docs Connect tenant.
type DocsCloudPaymentDto struct {
	// The cart ID.
	CartId NullableString `json:"cartId,omitempty"`
	// The product ID.
	ProductId *int32 `json:"productId,omitempty"`
	// The payment status.
	Status *int32 `json:"status,omitempty"`
	// The interval unit.
	IntervalUnit *int32 `json:"intervalUnit,omitempty"`
	// Whether the payment interval is yearly.
	IsYear *bool `json:"isYear,omitempty"`
	// Whether the payment is prepaid.
	IsPrepaid *bool `json:"isPrepaid,omitempty"`
	// The quantity.
	Quantity *int32 `json:"quantity,omitempty"`
	// The three-character ISO 4217 currency symbol of the payment.
	Currency NullableString `json:"currency,omitempty"`
}

// NewDocsCloudPaymentDto instantiates a new DocsCloudPaymentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudPaymentDto() *DocsCloudPaymentDto {
	this := DocsCloudPaymentDto{}
	return &this
}

// NewDocsCloudPaymentDtoWithDefaults instantiates a new DocsCloudPaymentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudPaymentDtoWithDefaults() *DocsCloudPaymentDto {
	this := DocsCloudPaymentDto{}
	return &this
}

// GetCartId returns the CartId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudPaymentDto) GetCartId() string {
	if o == nil || IsNil(o.CartId.Get()) {
		var ret string
		return ret
	}
	return *o.CartId.Get()
}

// GetCartIdOk returns a tuple with the CartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudPaymentDto) GetCartIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CartId.Get(), o.CartId.IsSet()
}

// HasCartId returns a boolean if a field has been set.
func (o *DocsCloudPaymentDto) IsCartIdSet() bool {
	if o != nil && o.CartId.IsSet() {
		return true
	}

	return false
}

// SetCartId gets a reference to the given NullableString and assigns it to the CartId field.
func (o *DocsCloudPaymentDto) SetCartId(v string) {
	o.CartId.Set(&v)
}
// SetCartIdNil sets the value for CartId to be an explicit nil
func (o *DocsCloudPaymentDto) SetCartIdNil() {
	o.CartId.Set(nil)
}

// UnsetCartId ensures that no value is present for CartId, not even an explicit nil
func (o *DocsCloudPaymentDto) UnsetCartId() {
	o.CartId.Unset()
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *DocsCloudPaymentDto) GetProductId() int32 {
	if o == nil || IsNil(o.ProductId) {
		var ret int32
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudPaymentDto) GetProductIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *DocsCloudPaymentDto) IsProductIdSet() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int32 and assigns it to the ProductId field.
func (o *DocsCloudPaymentDto) SetProductId(v int32) {
	o.ProductId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DocsCloudPaymentDto) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudPaymentDto) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DocsCloudPaymentDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *DocsCloudPaymentDto) SetStatus(v int32) {
	o.Status = &v
}

// GetIntervalUnit returns the IntervalUnit field value if set, zero value otherwise.
func (o *DocsCloudPaymentDto) GetIntervalUnit() int32 {
	if o == nil || IsNil(o.IntervalUnit) {
		var ret int32
		return ret
	}
	return *o.IntervalUnit
}

// GetIntervalUnitOk returns a tuple with the IntervalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudPaymentDto) GetIntervalUnitOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalUnit) {
		return nil, false
	}
	return o.IntervalUnit, true
}

// HasIntervalUnit returns a boolean if a field has been set.
func (o *DocsCloudPaymentDto) IsIntervalUnitSet() bool {
	if o != nil && !IsNil(o.IntervalUnit) {
		return true
	}

	return false
}

// SetIntervalUnit gets a reference to the given int32 and assigns it to the IntervalUnit field.
func (o *DocsCloudPaymentDto) SetIntervalUnit(v int32) {
	o.IntervalUnit = &v
}

// GetIsYear returns the IsYear field value if set, zero value otherwise.
func (o *DocsCloudPaymentDto) GetIsYear() bool {
	if o == nil || IsNil(o.IsYear) {
		var ret bool
		return ret
	}
	return *o.IsYear
}

// GetIsYearOk returns a tuple with the IsYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudPaymentDto) GetIsYearOk() (*bool, bool) {
	if o == nil || IsNil(o.IsYear) {
		return nil, false
	}
	return o.IsYear, true
}

// HasIsYear returns a boolean if a field has been set.
func (o *DocsCloudPaymentDto) IsIsYearSet() bool {
	if o != nil && !IsNil(o.IsYear) {
		return true
	}

	return false
}

// SetIsYear gets a reference to the given bool and assigns it to the IsYear field.
func (o *DocsCloudPaymentDto) SetIsYear(v bool) {
	o.IsYear = &v
}

// GetIsPrepaid returns the IsPrepaid field value if set, zero value otherwise.
func (o *DocsCloudPaymentDto) GetIsPrepaid() bool {
	if o == nil || IsNil(o.IsPrepaid) {
		var ret bool
		return ret
	}
	return *o.IsPrepaid
}

// GetIsPrepaidOk returns a tuple with the IsPrepaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudPaymentDto) GetIsPrepaidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrepaid) {
		return nil, false
	}
	return o.IsPrepaid, true
}

// HasIsPrepaid returns a boolean if a field has been set.
func (o *DocsCloudPaymentDto) IsIsPrepaidSet() bool {
	if o != nil && !IsNil(o.IsPrepaid) {
		return true
	}

	return false
}

// SetIsPrepaid gets a reference to the given bool and assigns it to the IsPrepaid field.
func (o *DocsCloudPaymentDto) SetIsPrepaid(v bool) {
	o.IsPrepaid = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *DocsCloudPaymentDto) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudPaymentDto) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *DocsCloudPaymentDto) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *DocsCloudPaymentDto) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudPaymentDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudPaymentDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *DocsCloudPaymentDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *DocsCloudPaymentDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *DocsCloudPaymentDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *DocsCloudPaymentDto) UnsetCurrency() {
	o.Currency.Unset()
}

func (o DocsCloudPaymentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudPaymentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CartId.IsSet() {
		toSerialize["cartId"] = o.CartId.Get()
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.IntervalUnit) {
		toSerialize["intervalUnit"] = o.IntervalUnit
	}
	if !IsNil(o.IsYear) {
		toSerialize["isYear"] = o.IsYear
	}
	if !IsNil(o.IsPrepaid) {
		toSerialize["isPrepaid"] = o.IsPrepaid
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	return toSerialize, nil
}

type NullableDocsCloudPaymentDto struct {
	value *DocsCloudPaymentDto
	isSet bool
}

func (v NullableDocsCloudPaymentDto) Get() *DocsCloudPaymentDto {
	return v.value
}

func (v *NullableDocsCloudPaymentDto) Set(val *DocsCloudPaymentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudPaymentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudPaymentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudPaymentDto(val *DocsCloudPaymentDto) *NullableDocsCloudPaymentDto {
	return &NullableDocsCloudPaymentDto{value: val, isSet: true}
}

func (v NullableDocsCloudPaymentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudPaymentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

