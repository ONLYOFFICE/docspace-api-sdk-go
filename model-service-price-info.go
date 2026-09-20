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

// checks if the ServicePriceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePriceInfo{}

// ServicePriceInfo Represents a price of the service.
type ServicePriceInfo struct {
	// The price unique identifier.
	Id *int32 `json:"id,omitempty"`
	// The account number.
	AccountNumber *int32 `json:"accountNumber,omitempty"`
	// The service ID.
	ServiceId *int32 `json:"serviceId,omitempty"`
	// The time unit the price is bound to.
	TimeUnit *PriceTimeUnit `json:"timeUnit,omitempty"`
	// The cost price.
	CostPrice *float64 `json:"costPrice,omitempty"`
	// The extra charge added to the cost price.
	ExtraCharge *float64 `json:"extraCharge,omitempty"`
	// The resulting service price.
	ServicePrice *float64 `json:"servicePrice,omitempty"`
	// The quota the price is set for.
	Quota NullableFloat64 `json:"quota,omitempty"`
	// The period the price is effective in.
	TimeBound *TimeBound `json:"timeBound,omitempty"`
	// The price status.
	Status *PriceStatus `json:"status,omitempty"`
	// The date and time when the price was created.
	Created *time.Time `json:"created,omitempty"`
	// The discount category ID.
	DiscountCategoryId NullableInt32 `json:"discountCategoryId,omitempty"`
	// The discount category.
	DiscountCategory *DiscountCategory `json:"discountCategory,omitempty"`
}

// NewServicePriceInfo instantiates a new ServicePriceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePriceInfo() *ServicePriceInfo {
	this := ServicePriceInfo{}
	return &this
}

// NewServicePriceInfoWithDefaults instantiates a new ServicePriceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePriceInfoWithDefaults() *ServicePriceInfo {
	this := ServicePriceInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ServicePriceInfo) SetId(v int32) {
	o.Id = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetAccountNumber() int32 {
	if o == nil || IsNil(o.AccountNumber) {
		var ret int32
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetAccountNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsAccountNumberSet() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given int32 and assigns it to the AccountNumber field.
func (o *ServicePriceInfo) SetAccountNumber(v int32) {
	o.AccountNumber = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetServiceId() int32 {
	if o == nil || IsNil(o.ServiceId) {
		var ret int32
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetServiceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsServiceIdSet() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given int32 and assigns it to the ServiceId field.
func (o *ServicePriceInfo) SetServiceId(v int32) {
	o.ServiceId = &v
}

// GetTimeUnit returns the TimeUnit field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetTimeUnit() PriceTimeUnit {
	if o == nil || IsNil(o.TimeUnit) {
		var ret PriceTimeUnit
		return ret
	}
	return *o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetTimeUnitOk() (*PriceTimeUnit, bool) {
	if o == nil || IsNil(o.TimeUnit) {
		return nil, false
	}
	return o.TimeUnit, true
}

// HasTimeUnit returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsTimeUnitSet() bool {
	if o != nil && !IsNil(o.TimeUnit) {
		return true
	}

	return false
}

// SetTimeUnit gets a reference to the given PriceTimeUnit and assigns it to the TimeUnit field.
func (o *ServicePriceInfo) SetTimeUnit(v PriceTimeUnit) {
	o.TimeUnit = &v
}

// GetCostPrice returns the CostPrice field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetCostPrice() float64 {
	if o == nil || IsNil(o.CostPrice) {
		var ret float64
		return ret
	}
	return *o.CostPrice
}

// GetCostPriceOk returns a tuple with the CostPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetCostPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPrice) {
		return nil, false
	}
	return o.CostPrice, true
}

// HasCostPrice returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsCostPriceSet() bool {
	if o != nil && !IsNil(o.CostPrice) {
		return true
	}

	return false
}

// SetCostPrice gets a reference to the given float64 and assigns it to the CostPrice field.
func (o *ServicePriceInfo) SetCostPrice(v float64) {
	o.CostPrice = &v
}

// GetExtraCharge returns the ExtraCharge field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetExtraCharge() float64 {
	if o == nil || IsNil(o.ExtraCharge) {
		var ret float64
		return ret
	}
	return *o.ExtraCharge
}

// GetExtraChargeOk returns a tuple with the ExtraCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetExtraChargeOk() (*float64, bool) {
	if o == nil || IsNil(o.ExtraCharge) {
		return nil, false
	}
	return o.ExtraCharge, true
}

// HasExtraCharge returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsExtraChargeSet() bool {
	if o != nil && !IsNil(o.ExtraCharge) {
		return true
	}

	return false
}

// SetExtraCharge gets a reference to the given float64 and assigns it to the ExtraCharge field.
func (o *ServicePriceInfo) SetExtraCharge(v float64) {
	o.ExtraCharge = &v
}

// GetServicePrice returns the ServicePrice field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetServicePrice() float64 {
	if o == nil || IsNil(o.ServicePrice) {
		var ret float64
		return ret
	}
	return *o.ServicePrice
}

// GetServicePriceOk returns a tuple with the ServicePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetServicePriceOk() (*float64, bool) {
	if o == nil || IsNil(o.ServicePrice) {
		return nil, false
	}
	return o.ServicePrice, true
}

// HasServicePrice returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsServicePriceSet() bool {
	if o != nil && !IsNil(o.ServicePrice) {
		return true
	}

	return false
}

// SetServicePrice gets a reference to the given float64 and assigns it to the ServicePrice field.
func (o *ServicePriceInfo) SetServicePrice(v float64) {
	o.ServicePrice = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePriceInfo) GetQuota() float64 {
	if o == nil || IsNil(o.Quota.Get()) {
		var ret float64
		return ret
	}
	return *o.Quota.Get()
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePriceInfo) GetQuotaOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quota.Get(), o.Quota.IsSet()
}

// HasQuota returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsQuotaSet() bool {
	if o != nil && o.Quota.IsSet() {
		return true
	}

	return false
}

// SetQuota gets a reference to the given NullableFloat64 and assigns it to the Quota field.
func (o *ServicePriceInfo) SetQuota(v float64) {
	o.Quota.Set(&v)
}
// SetQuotaNil sets the value for Quota to be an explicit nil
func (o *ServicePriceInfo) SetQuotaNil() {
	o.Quota.Set(nil)
}

// UnsetQuota ensures that no value is present for Quota, not even an explicit nil
func (o *ServicePriceInfo) UnsetQuota() {
	o.Quota.Unset()
}

// GetTimeBound returns the TimeBound field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetTimeBound() TimeBound {
	if o == nil || IsNil(o.TimeBound) {
		var ret TimeBound
		return ret
	}
	return *o.TimeBound
}

// GetTimeBoundOk returns a tuple with the TimeBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetTimeBoundOk() (*TimeBound, bool) {
	if o == nil || IsNil(o.TimeBound) {
		return nil, false
	}
	return o.TimeBound, true
}

// HasTimeBound returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsTimeBoundSet() bool {
	if o != nil && !IsNil(o.TimeBound) {
		return true
	}

	return false
}

// SetTimeBound gets a reference to the given TimeBound and assigns it to the TimeBound field.
func (o *ServicePriceInfo) SetTimeBound(v TimeBound) {
	o.TimeBound = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetStatus() PriceStatus {
	if o == nil || IsNil(o.Status) {
		var ret PriceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetStatusOk() (*PriceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PriceStatus and assigns it to the Status field.
func (o *ServicePriceInfo) SetStatus(v PriceStatus) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ServicePriceInfo) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDiscountCategoryId returns the DiscountCategoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePriceInfo) GetDiscountCategoryId() int32 {
	if o == nil || IsNil(o.DiscountCategoryId.Get()) {
		var ret int32
		return ret
	}
	return *o.DiscountCategoryId.Get()
}

// GetDiscountCategoryIdOk returns a tuple with the DiscountCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePriceInfo) GetDiscountCategoryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountCategoryId.Get(), o.DiscountCategoryId.IsSet()
}

// HasDiscountCategoryId returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsDiscountCategoryIdSet() bool {
	if o != nil && o.DiscountCategoryId.IsSet() {
		return true
	}

	return false
}

// SetDiscountCategoryId gets a reference to the given NullableInt32 and assigns it to the DiscountCategoryId field.
func (o *ServicePriceInfo) SetDiscountCategoryId(v int32) {
	o.DiscountCategoryId.Set(&v)
}
// SetDiscountCategoryIdNil sets the value for DiscountCategoryId to be an explicit nil
func (o *ServicePriceInfo) SetDiscountCategoryIdNil() {
	o.DiscountCategoryId.Set(nil)
}

// UnsetDiscountCategoryId ensures that no value is present for DiscountCategoryId, not even an explicit nil
func (o *ServicePriceInfo) UnsetDiscountCategoryId() {
	o.DiscountCategoryId.Unset()
}

// GetDiscountCategory returns the DiscountCategory field value if set, zero value otherwise.
func (o *ServicePriceInfo) GetDiscountCategory() DiscountCategory {
	if o == nil || IsNil(o.DiscountCategory) {
		var ret DiscountCategory
		return ret
	}
	return *o.DiscountCategory
}

// GetDiscountCategoryOk returns a tuple with the DiscountCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePriceInfo) GetDiscountCategoryOk() (*DiscountCategory, bool) {
	if o == nil || IsNil(o.DiscountCategory) {
		return nil, false
	}
	return o.DiscountCategory, true
}

// HasDiscountCategory returns a boolean if a field has been set.
func (o *ServicePriceInfo) IsDiscountCategorySet() bool {
	if o != nil && !IsNil(o.DiscountCategory) {
		return true
	}

	return false
}

// SetDiscountCategory gets a reference to the given DiscountCategory and assigns it to the DiscountCategory field.
func (o *ServicePriceInfo) SetDiscountCategory(v DiscountCategory) {
	o.DiscountCategory = &v
}

func (o ServicePriceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePriceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.TimeUnit) {
		toSerialize["timeUnit"] = o.TimeUnit
	}
	if !IsNil(o.CostPrice) {
		toSerialize["costPrice"] = o.CostPrice
	}
	if !IsNil(o.ExtraCharge) {
		toSerialize["extraCharge"] = o.ExtraCharge
	}
	if !IsNil(o.ServicePrice) {
		toSerialize["servicePrice"] = o.ServicePrice
	}
	if o.Quota.IsSet() {
		toSerialize["quota"] = o.Quota.Get()
	}
	if !IsNil(o.TimeBound) {
		toSerialize["timeBound"] = o.TimeBound
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if o.DiscountCategoryId.IsSet() {
		toSerialize["discountCategoryId"] = o.DiscountCategoryId.Get()
	}
	if !IsNil(o.DiscountCategory) {
		toSerialize["discountCategory"] = o.DiscountCategory
	}
	return toSerialize, nil
}

type NullableServicePriceInfo struct {
	value *ServicePriceInfo
	isSet bool
}

func (v NullableServicePriceInfo) Get() *ServicePriceInfo {
	return v.value
}

func (v *NullableServicePriceInfo) Set(val *ServicePriceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePriceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePriceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePriceInfo(val *ServicePriceInfo) *NullableServicePriceInfo {
	return &NullableServicePriceInfo{value: val, isSet: true}
}

func (v NullableServicePriceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePriceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

