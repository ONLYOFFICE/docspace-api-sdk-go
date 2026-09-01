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

// checks if the CustomerServiceUsageDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerServiceUsageDto{}

// CustomerServiceUsageDto Aggregated customer usage statistics for a service over a period.
type CustomerServiceUsageDto struct {
	// The name of the service.
	Service NullableString `json:"service,omitempty"`
	// The title of the service.
	Title NullableString `json:"title,omitempty"`
	// The unit of measurement for the service.
	ServiceUnit NullableString `json:"serviceUnit,omitempty"`
	// The three-character ISO 4217 currency symbol of the amounts.
	Currency NullableString `json:"currency,omitempty"`
	// The total number of units consumed.
	TotalQuantity *int32 `json:"totalQuantity,omitempty"`
	// The total amount charged for the service.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// The number of individual purchase operations.
	OperationCount *int32 `json:"operationCount,omitempty"`
	// The price of the service.
	Price *float64 `json:"price,omitempty"`
	// Indicates whether the service is subscription-based.
	Subscription *bool `json:"subscription,omitempty"`
}

// NewCustomerServiceUsageDto instantiates a new CustomerServiceUsageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerServiceUsageDto() *CustomerServiceUsageDto {
	this := CustomerServiceUsageDto{}
	return &this
}

// NewCustomerServiceUsageDtoWithDefaults instantiates a new CustomerServiceUsageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerServiceUsageDtoWithDefaults() *CustomerServiceUsageDto {
	this := CustomerServiceUsageDto{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerServiceUsageDto) GetService() string {
	if o == nil || IsNil(o.Service.Get()) {
		var ret string
		return ret
	}
	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerServiceUsageDto) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// HasService returns a boolean if a field has been set.
func (o *CustomerServiceUsageDto) IsServiceSet() bool {
	if o != nil && o.Service.IsSet() {
		return true
	}

	return false
}

// SetService gets a reference to the given NullableString and assigns it to the Service field.
func (o *CustomerServiceUsageDto) SetService(v string) {
	o.Service.Set(&v)
}
// SetServiceNil sets the value for Service to be an explicit nil
func (o *CustomerServiceUsageDto) SetServiceNil() {
	o.Service.Set(nil)
}

// UnsetService ensures that no value is present for Service, not even an explicit nil
func (o *CustomerServiceUsageDto) UnsetService() {
	o.Service.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerServiceUsageDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerServiceUsageDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *CustomerServiceUsageDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *CustomerServiceUsageDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *CustomerServiceUsageDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *CustomerServiceUsageDto) UnsetTitle() {
	o.Title.Unset()
}

// GetServiceUnit returns the ServiceUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerServiceUsageDto) GetServiceUnit() string {
	if o == nil || IsNil(o.ServiceUnit.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceUnit.Get()
}

// GetServiceUnitOk returns a tuple with the ServiceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerServiceUsageDto) GetServiceUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceUnit.Get(), o.ServiceUnit.IsSet()
}

// HasServiceUnit returns a boolean if a field has been set.
func (o *CustomerServiceUsageDto) IsServiceUnitSet() bool {
	if o != nil && o.ServiceUnit.IsSet() {
		return true
	}

	return false
}

// SetServiceUnit gets a reference to the given NullableString and assigns it to the ServiceUnit field.
func (o *CustomerServiceUsageDto) SetServiceUnit(v string) {
	o.ServiceUnit.Set(&v)
}
// SetServiceUnitNil sets the value for ServiceUnit to be an explicit nil
func (o *CustomerServiceUsageDto) SetServiceUnitNil() {
	o.ServiceUnit.Set(nil)
}

// UnsetServiceUnit ensures that no value is present for ServiceUnit, not even an explicit nil
func (o *CustomerServiceUsageDto) UnsetServiceUnit() {
	o.ServiceUnit.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerServiceUsageDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerServiceUsageDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *CustomerServiceUsageDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *CustomerServiceUsageDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *CustomerServiceUsageDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *CustomerServiceUsageDto) UnsetCurrency() {
	o.Currency.Unset()
}

// GetTotalQuantity returns the TotalQuantity field value if set, zero value otherwise.
func (o *CustomerServiceUsageDto) GetTotalQuantity() int32 {
	if o == nil || IsNil(o.TotalQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalQuantity
}

// GetTotalQuantityOk returns a tuple with the TotalQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceUsageDto) GetTotalQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalQuantity) {
		return nil, false
	}
	return o.TotalQuantity, true
}

// HasTotalQuantity returns a boolean if a field has been set.
func (o *CustomerServiceUsageDto) IsTotalQuantitySet() bool {
	if o != nil && !IsNil(o.TotalQuantity) {
		return true
	}

	return false
}

// SetTotalQuantity gets a reference to the given int32 and assigns it to the TotalQuantity field.
func (o *CustomerServiceUsageDto) SetTotalQuantity(v int32) {
	o.TotalQuantity = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *CustomerServiceUsageDto) GetTotalAmount() float64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret float64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceUsageDto) GetTotalAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *CustomerServiceUsageDto) IsTotalAmountSet() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float64 and assigns it to the TotalAmount field.
func (o *CustomerServiceUsageDto) SetTotalAmount(v float64) {
	o.TotalAmount = &v
}

// GetOperationCount returns the OperationCount field value if set, zero value otherwise.
func (o *CustomerServiceUsageDto) GetOperationCount() int32 {
	if o == nil || IsNil(o.OperationCount) {
		var ret int32
		return ret
	}
	return *o.OperationCount
}

// GetOperationCountOk returns a tuple with the OperationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceUsageDto) GetOperationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OperationCount) {
		return nil, false
	}
	return o.OperationCount, true
}

// HasOperationCount returns a boolean if a field has been set.
func (o *CustomerServiceUsageDto) IsOperationCountSet() bool {
	if o != nil && !IsNil(o.OperationCount) {
		return true
	}

	return false
}

// SetOperationCount gets a reference to the given int32 and assigns it to the OperationCount field.
func (o *CustomerServiceUsageDto) SetOperationCount(v int32) {
	o.OperationCount = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CustomerServiceUsageDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceUsageDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CustomerServiceUsageDto) IsPriceSet() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *CustomerServiceUsageDto) SetPrice(v float64) {
	o.Price = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *CustomerServiceUsageDto) GetSubscription() bool {
	if o == nil || IsNil(o.Subscription) {
		var ret bool
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceUsageDto) GetSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *CustomerServiceUsageDto) IsSubscriptionSet() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given bool and assigns it to the Subscription field.
func (o *CustomerServiceUsageDto) SetSubscription(v bool) {
	o.Subscription = &v
}

func (o CustomerServiceUsageDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerServiceUsageDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Service.IsSet() {
		toSerialize["service"] = o.Service.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.ServiceUnit.IsSet() {
		toSerialize["serviceUnit"] = o.ServiceUnit.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.TotalQuantity) {
		toSerialize["totalQuantity"] = o.TotalQuantity
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.OperationCount) {
		toSerialize["operationCount"] = o.OperationCount
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableCustomerServiceUsageDto struct {
	value *CustomerServiceUsageDto
	isSet bool
}

func (v NullableCustomerServiceUsageDto) Get() *CustomerServiceUsageDto {
	return v.value
}

func (v *NullableCustomerServiceUsageDto) Set(val *CustomerServiceUsageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerServiceUsageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerServiceUsageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerServiceUsageDto(val *CustomerServiceUsageDto) *NullableCustomerServiceUsageDto {
	return &NullableCustomerServiceUsageDto{value: val, isSet: true}
}

func (v NullableCustomerServiceUsageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerServiceUsageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

