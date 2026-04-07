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

// checks if the Quota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Quota{}

// Quota The quota parameters.  <example>  {    id: 1,    quantity: 50,    wallet: false,    dueDate: 2026-03-31T00:00:00Z,    nextQuantity: 100,    state: Active  }  </example>
type Quota struct {
	// The quota ID.
	Id *int32 `json:"id,omitempty"`
	// The quota quantity.
	Quantity *int32 `json:"quantity,omitempty"`
	// The quota applies to the wallet or not
	Wallet *bool `json:"wallet,omitempty"`
	// The quota due date.
	DueDate NullableTime `json:"dueDate,omitempty"`
	// The quota next quantity.
	NextQuantity NullableInt32 `json:"nextQuantity,omitempty"`
	State *QuotaState `json:"state,omitempty"`
}

// NewQuota instantiates a new Quota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuota() *Quota {
	this := Quota{}
	return &this
}

// NewQuotaWithDefaults instantiates a new Quota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaWithDefaults() *Quota {
	this := Quota{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Quota) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Quota) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Quota) SetId(v int32) {
	o.Id = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Quota) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Quota) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *Quota) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *Quota) GetWallet() bool {
	if o == nil || IsNil(o.Wallet) {
		var ret bool
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetWalletOk() (*bool, bool) {
	if o == nil || IsNil(o.Wallet) {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *Quota) IsWalletSet() bool {
	if o != nil && !IsNil(o.Wallet) {
		return true
	}

	return false
}

// SetWallet gets a reference to the given bool and assigns it to the Wallet field.
func (o *Quota) SetWallet(v bool) {
	o.Wallet = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Quota) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Quota) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *Quota) IsDueDateSet() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableTime and assigns it to the DueDate field.
func (o *Quota) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *Quota) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *Quota) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetNextQuantity returns the NextQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Quota) GetNextQuantity() int32 {
	if o == nil || IsNil(o.NextQuantity.Get()) {
		var ret int32
		return ret
	}
	return *o.NextQuantity.Get()
}

// GetNextQuantityOk returns a tuple with the NextQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Quota) GetNextQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextQuantity.Get(), o.NextQuantity.IsSet()
}

// HasNextQuantity returns a boolean if a field has been set.
func (o *Quota) IsNextQuantitySet() bool {
	if o != nil && o.NextQuantity.IsSet() {
		return true
	}

	return false
}

// SetNextQuantity gets a reference to the given NullableInt32 and assigns it to the NextQuantity field.
func (o *Quota) SetNextQuantity(v int32) {
	o.NextQuantity.Set(&v)
}
// SetNextQuantityNil sets the value for NextQuantity to be an explicit nil
func (o *Quota) SetNextQuantityNil() {
	o.NextQuantity.Set(nil)
}

// UnsetNextQuantity ensures that no value is present for NextQuantity, not even an explicit nil
func (o *Quota) UnsetNextQuantity() {
	o.NextQuantity.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Quota) GetState() QuotaState {
	if o == nil || IsNil(o.State) {
		var ret QuotaState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetStateOk() (*QuotaState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Quota) IsStateSet() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given QuotaState and assigns it to the State field.
func (o *Quota) SetState(v QuotaState) {
	o.State = &v
}

func (o Quota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Quota) ToMap() (map[string]interface{}, error) {
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
	if o.DueDate.IsSet() {
		toSerialize["dueDate"] = o.DueDate.Get()
	}
	if o.NextQuantity.IsSet() {
		toSerialize["nextQuantity"] = o.NextQuantity.Get()
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableQuota struct {
	value *Quota
	isSet bool
}

func (v NullableQuota) Get() *Quota {
	return v.value
}

func (v *NullableQuota) Set(val *Quota) {
	v.value = val
	v.isSet = true
}

func (v NullableQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuota(val *Quota) *NullableQuota {
	return &NullableQuota{value: val, isSet: true}
}

func (v NullableQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

