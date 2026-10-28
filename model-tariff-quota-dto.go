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

// checks if the TariffQuotaDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TariffQuotaDto{}

// TariffQuotaDto One quota the subscription is made of - the plan itself or an add-on - with its quantity and its own deadline.
type TariffQuotaDto struct {
	// The quota this entry stands for. `GET api/2.0/portal/payment/quotas` describes the quota behind the ID,  including what its `quantity` counts; a negative ID belongs to a built-in quota rather than a purchased  one.
	Id *int32 `json:"id,omitempty"`
	// How much of the quota the portal holds, in whatever the quota itself is measured in - seats for a plan,  gigabytes for storage. It is `1` for a quota that is simply on or off.
	Quantity *int32 `json:"quantity,omitempty"`
	// Whether the quota is paid for out of the portal wallet as it is consumed, rather than being part of the  subscription charged per period.
	Wallet *bool `json:"wallet,omitempty"`
	// Whether this is an add-on bought on top of the plan rather than the plan itself. Exactly one entry of  `quotas` is the plan, and the rest are add-ons.
	Additional *bool `json:"additional,omitempty"`
	// When this quota runs out, in the portal time zone. An add-on can end earlier or later than the  subscription; a quota with no deadline of its own reports the subscription's `dueDate` instead of an empty  value.
	DueDate *ApiDateTime `json:"dueDate,omitempty"`
	// The quantity the next period is going to be charged for, when a change has been scheduled. It is empty  while `quantity` simply carries over.
	NextQuantity NullableInt32 `json:"nextQuantity,omitempty"`
	// The quota this one is scheduled to be replaced by at the start of the next period, empty when no such  switch is planned. `GET api/2.0/portal/tariff/upcoming` already reports the charge for the replacement.
	NextQuota NullableInt32 `json:"nextQuota,omitempty"`
	// Whether the quota is still running or its deadline has passed. It is empty for a quota that has no  deadline of its own, which means it lasts as long as the subscription does.
	State *QuotaState `json:"state,omitempty"`
}

// NewTariffQuotaDto instantiates a new TariffQuotaDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTariffQuotaDto() *TariffQuotaDto {
	this := TariffQuotaDto{}
	return &this
}

// NewTariffQuotaDtoWithDefaults instantiates a new TariffQuotaDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTariffQuotaDtoWithDefaults() *TariffQuotaDto {
	this := TariffQuotaDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TariffQuotaDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffQuotaDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TariffQuotaDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TariffQuotaDto) SetId(v int32) {
	o.Id = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *TariffQuotaDto) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffQuotaDto) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *TariffQuotaDto) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *TariffQuotaDto) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *TariffQuotaDto) GetWallet() bool {
	if o == nil || IsNil(o.Wallet) {
		var ret bool
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffQuotaDto) GetWalletOk() (*bool, bool) {
	if o == nil || IsNil(o.Wallet) {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *TariffQuotaDto) IsWalletSet() bool {
	if o != nil && !IsNil(o.Wallet) {
		return true
	}

	return false
}

// SetWallet gets a reference to the given bool and assigns it to the Wallet field.
func (o *TariffQuotaDto) SetWallet(v bool) {
	o.Wallet = &v
}

// GetAdditional returns the Additional field value if set, zero value otherwise.
func (o *TariffQuotaDto) GetAdditional() bool {
	if o == nil || IsNil(o.Additional) {
		var ret bool
		return ret
	}
	return *o.Additional
}

// GetAdditionalOk returns a tuple with the Additional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffQuotaDto) GetAdditionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Additional) {
		return nil, false
	}
	return o.Additional, true
}

// HasAdditional returns a boolean if a field has been set.
func (o *TariffQuotaDto) IsAdditionalSet() bool {
	if o != nil && !IsNil(o.Additional) {
		return true
	}

	return false
}

// SetAdditional gets a reference to the given bool and assigns it to the Additional field.
func (o *TariffQuotaDto) SetAdditional(v bool) {
	o.Additional = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *TariffQuotaDto) GetDueDate() ApiDateTime {
	if o == nil || IsNil(o.DueDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffQuotaDto) GetDueDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *TariffQuotaDto) IsDueDateSet() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given ApiDateTime and assigns it to the DueDate field.
func (o *TariffQuotaDto) SetDueDate(v ApiDateTime) {
	o.DueDate = &v
}

// GetNextQuantity returns the NextQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TariffQuotaDto) GetNextQuantity() int32 {
	if o == nil || IsNil(o.NextQuantity.Get()) {
		var ret int32
		return ret
	}
	return *o.NextQuantity.Get()
}

// GetNextQuantityOk returns a tuple with the NextQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TariffQuotaDto) GetNextQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextQuantity.Get(), o.NextQuantity.IsSet()
}

// HasNextQuantity returns a boolean if a field has been set.
func (o *TariffQuotaDto) IsNextQuantitySet() bool {
	if o != nil && o.NextQuantity.IsSet() {
		return true
	}

	return false
}

// SetNextQuantity gets a reference to the given NullableInt32 and assigns it to the NextQuantity field.
func (o *TariffQuotaDto) SetNextQuantity(v int32) {
	o.NextQuantity.Set(&v)
}
// SetNextQuantityNil sets the value for NextQuantity to be an explicit nil
func (o *TariffQuotaDto) SetNextQuantityNil() {
	o.NextQuantity.Set(nil)
}

// UnsetNextQuantity ensures that no value is present for NextQuantity, not even an explicit nil
func (o *TariffQuotaDto) UnsetNextQuantity() {
	o.NextQuantity.Unset()
}

// GetNextQuota returns the NextQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TariffQuotaDto) GetNextQuota() int32 {
	if o == nil || IsNil(o.NextQuota.Get()) {
		var ret int32
		return ret
	}
	return *o.NextQuota.Get()
}

// GetNextQuotaOk returns a tuple with the NextQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TariffQuotaDto) GetNextQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextQuota.Get(), o.NextQuota.IsSet()
}

// HasNextQuota returns a boolean if a field has been set.
func (o *TariffQuotaDto) IsNextQuotaSet() bool {
	if o != nil && o.NextQuota.IsSet() {
		return true
	}

	return false
}

// SetNextQuota gets a reference to the given NullableInt32 and assigns it to the NextQuota field.
func (o *TariffQuotaDto) SetNextQuota(v int32) {
	o.NextQuota.Set(&v)
}
// SetNextQuotaNil sets the value for NextQuota to be an explicit nil
func (o *TariffQuotaDto) SetNextQuotaNil() {
	o.NextQuota.Set(nil)
}

// UnsetNextQuota ensures that no value is present for NextQuota, not even an explicit nil
func (o *TariffQuotaDto) UnsetNextQuota() {
	o.NextQuota.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TariffQuotaDto) GetState() QuotaState {
	if o == nil || IsNil(o.State) {
		var ret QuotaState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffQuotaDto) GetStateOk() (*QuotaState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TariffQuotaDto) IsStateSet() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given QuotaState and assigns it to the State field.
func (o *TariffQuotaDto) SetState(v QuotaState) {
	o.State = &v
}

func (o TariffQuotaDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TariffQuotaDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Wallet) {
		toSerialize["wallet"] = o.Wallet
	}
	if !IsNil(o.Additional) {
		toSerialize["additional"] = o.Additional
	}
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if o.NextQuantity.IsSet() {
		toSerialize["nextQuantity"] = o.NextQuantity.Get()
	}
	if o.NextQuota.IsSet() {
		toSerialize["nextQuota"] = o.NextQuota.Get()
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableTariffQuotaDto struct {
	value *TariffQuotaDto
	isSet bool
}

func (v NullableTariffQuotaDto) Get() *TariffQuotaDto {
	return v.value
}

func (v *NullableTariffQuotaDto) Set(val *TariffQuotaDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTariffQuotaDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTariffQuotaDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTariffQuotaDto(val *TariffQuotaDto) *NullableTariffQuotaDto {
	return &NullableTariffQuotaDto{value: val, isSet: true}
}

func (v NullableTariffQuotaDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTariffQuotaDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

