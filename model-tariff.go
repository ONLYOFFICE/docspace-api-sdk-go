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
	"bytes"
	"fmt"
)

// checks if the Tariff type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tariff{}

// Tariff The tariff parameters.
type Tariff struct {
	// The tariff ID.
	Id *int32 `json:"id,omitempty"`
	State *TariffState `json:"state,omitempty"`
	// The tariff due date.
	DueDate time.Time `json:"dueDate"`
	// The tariff delay due date.
	DelayDueDate *time.Time `json:"delayDueDate,omitempty"`
	// The tariff license date.
	LicenseDate *time.Time `json:"licenseDate,omitempty"`
	// The tariff customer ID.
	CustomerId NullableString `json:"customerId,omitempty"`
	// The list of tariff quotas.
	Quotas []Quota `json:"quotas"`
	// The list of overdue tariff quotas.
	OverdueQuotas []Quota `json:"overdueQuotas,omitempty"`
}

type _Tariff Tariff

// NewTariff instantiates a new Tariff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTariff(dueDate time.Time, quotas []Quota) *Tariff {
	this := Tariff{}
	this.DueDate = dueDate
	this.Quotas = quotas
	return &this
}

// NewTariffWithDefaults instantiates a new Tariff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTariffWithDefaults() *Tariff {
	this := Tariff{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tariff) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tariff) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Tariff) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Tariff) SetId(v int32) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Tariff) GetState() TariffState {
	if o == nil || IsNil(o.State) {
		var ret TariffState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tariff) GetStateOk() (*TariffState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Tariff) IsStateSet() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given TariffState and assigns it to the State field.
func (o *Tariff) SetState(v TariffState) {
	o.State = &v
}

// GetDueDate returns the DueDate field value
func (o *Tariff) GetDueDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value
// and a boolean to check if the value has been set.
func (o *Tariff) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDate, true
}

// SetDueDate sets field value
func (o *Tariff) SetDueDate(v time.Time) {
	o.DueDate = v
}

// GetDelayDueDate returns the DelayDueDate field value if set, zero value otherwise.
func (o *Tariff) GetDelayDueDate() time.Time {
	if o == nil || IsNil(o.DelayDueDate) {
		var ret time.Time
		return ret
	}
	return *o.DelayDueDate
}

// GetDelayDueDateOk returns a tuple with the DelayDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tariff) GetDelayDueDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DelayDueDate) {
		return nil, false
	}
	return o.DelayDueDate, true
}

// HasDelayDueDate returns a boolean if a field has been set.
func (o *Tariff) IsDelayDueDateSet() bool {
	if o != nil && !IsNil(o.DelayDueDate) {
		return true
	}

	return false
}

// SetDelayDueDate gets a reference to the given time.Time and assigns it to the DelayDueDate field.
func (o *Tariff) SetDelayDueDate(v time.Time) {
	o.DelayDueDate = &v
}

// GetLicenseDate returns the LicenseDate field value if set, zero value otherwise.
func (o *Tariff) GetLicenseDate() time.Time {
	if o == nil || IsNil(o.LicenseDate) {
		var ret time.Time
		return ret
	}
	return *o.LicenseDate
}

// GetLicenseDateOk returns a tuple with the LicenseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tariff) GetLicenseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LicenseDate) {
		return nil, false
	}
	return o.LicenseDate, true
}

// HasLicenseDate returns a boolean if a field has been set.
func (o *Tariff) IsLicenseDateSet() bool {
	if o != nil && !IsNil(o.LicenseDate) {
		return true
	}

	return false
}

// SetLicenseDate gets a reference to the given time.Time and assigns it to the LicenseDate field.
func (o *Tariff) SetLicenseDate(v time.Time) {
	o.LicenseDate = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tariff) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tariff) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Tariff) IsCustomerIdSet() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *Tariff) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *Tariff) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *Tariff) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetQuotas returns the Quotas field value
// If the value is explicit nil, the zero value for []Quota will be returned
func (o *Tariff) GetQuotas() []Quota {
	if o == nil {
		var ret []Quota
		return ret
	}

	return o.Quotas
}

// GetQuotasOk returns a tuple with the Quotas field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tariff) GetQuotasOk() ([]Quota, bool) {
	if o == nil || IsNil(o.Quotas) {
		return nil, false
	}
	return o.Quotas, true
}

// SetQuotas sets field value
func (o *Tariff) SetQuotas(v []Quota) {
	o.Quotas = v
}

// GetOverdueQuotas returns the OverdueQuotas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tariff) GetOverdueQuotas() []Quota {
	if o == nil {
		var ret []Quota
		return ret
	}
	return o.OverdueQuotas
}

// GetOverdueQuotasOk returns a tuple with the OverdueQuotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tariff) GetOverdueQuotasOk() ([]Quota, bool) {
	if o == nil || IsNil(o.OverdueQuotas) {
		return nil, false
	}
	return o.OverdueQuotas, true
}

// HasOverdueQuotas returns a boolean if a field has been set.
func (o *Tariff) IsOverdueQuotasSet() bool {
	if o != nil && !IsNil(o.OverdueQuotas) {
		return true
	}

	return false
}

// SetOverdueQuotas gets a reference to the given []Quota and assigns it to the OverdueQuotas field.
func (o *Tariff) SetOverdueQuotas(v []Quota) {
	o.OverdueQuotas = v
}

func (o Tariff) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tariff) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["dueDate"] = o.DueDate
	if !IsNil(o.DelayDueDate) {
		toSerialize["delayDueDate"] = o.DelayDueDate
	}
	if !IsNil(o.LicenseDate) {
		toSerialize["licenseDate"] = o.LicenseDate
	}
	if o.CustomerId.IsSet() {
		toSerialize["customerId"] = o.CustomerId.Get()
	}
	if o.Quotas != nil {
		toSerialize["quotas"] = o.Quotas
	}
	if o.OverdueQuotas != nil {
		toSerialize["overdueQuotas"] = o.OverdueQuotas
	}
	return toSerialize, nil
}

func (o *Tariff) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dueDate",
		"quotas",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTariff := _Tariff{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTariff)

	if err != nil {
		return err
	}

	*o = Tariff(varTariff)

	return err
}

type NullableTariff struct {
	value *Tariff
	isSet bool
}

func (v NullableTariff) Get() *Tariff {
	return v.value
}

func (v *NullableTariff) Set(val *Tariff) {
	v.value = val
	v.isSet = true
}

func (v NullableTariff) IsSet() bool {
	return v.isSet
}

func (v *NullableTariff) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTariff(val *Tariff) *NullableTariff {
	return &NullableTariff{value: val, isSet: true}
}

func (v NullableTariff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTariff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

