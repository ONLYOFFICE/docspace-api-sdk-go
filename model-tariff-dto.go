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

// checks if the TariffDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TariffDto{}

// TariffDto The subscription this portal runs on: its state, the end of the current period, and the quotas it is made of.
type TariffDto struct {
	// Whether the installation runs the open-source build, which has no paid plan at all. This flag and the two  below describe the build rather than the subscription, and all three are left empty for a caller without  the portal-settings right.
	OpenSource NullableBool `json:"openSource,omitempty"`
	// Whether the installation runs on an Enterprise licence file, which is what makes the licence operations  under `api/2.0/settings/license` usable.
	Enterprise NullableBool `json:"enterprise,omitempty"`
	// Whether the installation runs on a Developer licence, an Enterprise licence meant for embedding rather  than for production use.
	Developer NullableBool `json:"developer,omitempty"`
	// The identifier of the subscription record itself, for quoting when a charge has to be traced. It is filled  in for a caller with the portal-settings right only, and nothing accepts it as an argument.
	Id *int32 `json:"id,omitempty"`
	// How the subscription stands: on trial, paid, inside the grace period that follows the due date, or unpaid.  It is the one field every caller gets, whatever their role, so a client can warn about payment without  needing administrator rights.
	State *TariffState `json:"state,omitempty"`
	// When the current period ends, in the portal time zone. It is filled in for a room or DocSpace  administrator only, and set to the largest value a date can hold for a subscription that never ends.
	DueDate *ApiDateTime `json:"dueDate,omitempty"`
	// When the grace period after `dueDate` runs out and the portal is cut off, in the portal time zone. Filled  in under the same conditions as `dueDate`, and equal to it when the plan grants no grace period.
	DelayDueDate *ApiDateTime `json:"delayDueDate,omitempty"`
	// When the licence file behind the subscription was issued, in the portal time zone. It is meaningful on a  server installation and filled in for a caller with the portal-settings right only.
	LicenseDate *ApiDateTime `json:"licenseDate,omitempty"`
	// The account in the billing system the subscription is charged to, empty for a portal that has never been  billed. Filled in for a caller with the portal-settings right only.
	CustomerId NullableString `json:"customerId,omitempty"`
	// The quotas the subscription is made of - the plan itself and its add-ons - with the overdue ones listed  alongside the current ones, so an entry here is not proof that it is still being paid for; read each  entry's own `state` for that. Filled in for a caller with the portal-settings right only.
	Quotas []TariffQuotaDto `json:"quotas,omitempty"`
}

// NewTariffDto instantiates a new TariffDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTariffDto() *TariffDto {
	this := TariffDto{}
	return &this
}

// NewTariffDtoWithDefaults instantiates a new TariffDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTariffDtoWithDefaults() *TariffDto {
	this := TariffDto{}
	return &this
}

// GetOpenSource returns the OpenSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TariffDto) GetOpenSource() bool {
	if o == nil || IsNil(o.OpenSource.Get()) {
		var ret bool
		return ret
	}
	return *o.OpenSource.Get()
}

// GetOpenSourceOk returns a tuple with the OpenSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TariffDto) GetOpenSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenSource.Get(), o.OpenSource.IsSet()
}

// HasOpenSource returns a boolean if a field has been set.
func (o *TariffDto) IsOpenSourceSet() bool {
	if o != nil && o.OpenSource.IsSet() {
		return true
	}

	return false
}

// SetOpenSource gets a reference to the given NullableBool and assigns it to the OpenSource field.
func (o *TariffDto) SetOpenSource(v bool) {
	o.OpenSource.Set(&v)
}
// SetOpenSourceNil sets the value for OpenSource to be an explicit nil
func (o *TariffDto) SetOpenSourceNil() {
	o.OpenSource.Set(nil)
}

// UnsetOpenSource ensures that no value is present for OpenSource, not even an explicit nil
func (o *TariffDto) UnsetOpenSource() {
	o.OpenSource.Unset()
}

// GetEnterprise returns the Enterprise field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TariffDto) GetEnterprise() bool {
	if o == nil || IsNil(o.Enterprise.Get()) {
		var ret bool
		return ret
	}
	return *o.Enterprise.Get()
}

// GetEnterpriseOk returns a tuple with the Enterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TariffDto) GetEnterpriseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enterprise.Get(), o.Enterprise.IsSet()
}

// HasEnterprise returns a boolean if a field has been set.
func (o *TariffDto) IsEnterpriseSet() bool {
	if o != nil && o.Enterprise.IsSet() {
		return true
	}

	return false
}

// SetEnterprise gets a reference to the given NullableBool and assigns it to the Enterprise field.
func (o *TariffDto) SetEnterprise(v bool) {
	o.Enterprise.Set(&v)
}
// SetEnterpriseNil sets the value for Enterprise to be an explicit nil
func (o *TariffDto) SetEnterpriseNil() {
	o.Enterprise.Set(nil)
}

// UnsetEnterprise ensures that no value is present for Enterprise, not even an explicit nil
func (o *TariffDto) UnsetEnterprise() {
	o.Enterprise.Unset()
}

// GetDeveloper returns the Developer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TariffDto) GetDeveloper() bool {
	if o == nil || IsNil(o.Developer.Get()) {
		var ret bool
		return ret
	}
	return *o.Developer.Get()
}

// GetDeveloperOk returns a tuple with the Developer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TariffDto) GetDeveloperOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Developer.Get(), o.Developer.IsSet()
}

// HasDeveloper returns a boolean if a field has been set.
func (o *TariffDto) IsDeveloperSet() bool {
	if o != nil && o.Developer.IsSet() {
		return true
	}

	return false
}

// SetDeveloper gets a reference to the given NullableBool and assigns it to the Developer field.
func (o *TariffDto) SetDeveloper(v bool) {
	o.Developer.Set(&v)
}
// SetDeveloperNil sets the value for Developer to be an explicit nil
func (o *TariffDto) SetDeveloperNil() {
	o.Developer.Set(nil)
}

// UnsetDeveloper ensures that no value is present for Developer, not even an explicit nil
func (o *TariffDto) UnsetDeveloper() {
	o.Developer.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TariffDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TariffDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TariffDto) SetId(v int32) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TariffDto) GetState() TariffState {
	if o == nil || IsNil(o.State) {
		var ret TariffState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffDto) GetStateOk() (*TariffState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TariffDto) IsStateSet() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given TariffState and assigns it to the State field.
func (o *TariffDto) SetState(v TariffState) {
	o.State = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *TariffDto) GetDueDate() ApiDateTime {
	if o == nil || IsNil(o.DueDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffDto) GetDueDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *TariffDto) IsDueDateSet() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given ApiDateTime and assigns it to the DueDate field.
func (o *TariffDto) SetDueDate(v ApiDateTime) {
	o.DueDate = &v
}

// GetDelayDueDate returns the DelayDueDate field value if set, zero value otherwise.
func (o *TariffDto) GetDelayDueDate() ApiDateTime {
	if o == nil || IsNil(o.DelayDueDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.DelayDueDate
}

// GetDelayDueDateOk returns a tuple with the DelayDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffDto) GetDelayDueDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.DelayDueDate) {
		return nil, false
	}
	return o.DelayDueDate, true
}

// HasDelayDueDate returns a boolean if a field has been set.
func (o *TariffDto) IsDelayDueDateSet() bool {
	if o != nil && !IsNil(o.DelayDueDate) {
		return true
	}

	return false
}

// SetDelayDueDate gets a reference to the given ApiDateTime and assigns it to the DelayDueDate field.
func (o *TariffDto) SetDelayDueDate(v ApiDateTime) {
	o.DelayDueDate = &v
}

// GetLicenseDate returns the LicenseDate field value if set, zero value otherwise.
func (o *TariffDto) GetLicenseDate() ApiDateTime {
	if o == nil || IsNil(o.LicenseDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.LicenseDate
}

// GetLicenseDateOk returns a tuple with the LicenseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffDto) GetLicenseDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.LicenseDate) {
		return nil, false
	}
	return o.LicenseDate, true
}

// HasLicenseDate returns a boolean if a field has been set.
func (o *TariffDto) IsLicenseDateSet() bool {
	if o != nil && !IsNil(o.LicenseDate) {
		return true
	}

	return false
}

// SetLicenseDate gets a reference to the given ApiDateTime and assigns it to the LicenseDate field.
func (o *TariffDto) SetLicenseDate(v ApiDateTime) {
	o.LicenseDate = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TariffDto) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TariffDto) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *TariffDto) IsCustomerIdSet() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *TariffDto) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *TariffDto) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *TariffDto) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetQuotas returns the Quotas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TariffDto) GetQuotas() []TariffQuotaDto {
	if o == nil {
		var ret []TariffQuotaDto
		return ret
	}
	return o.Quotas
}

// GetQuotasOk returns a tuple with the Quotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TariffDto) GetQuotasOk() ([]TariffQuotaDto, bool) {
	if o == nil || IsNil(o.Quotas) {
		return nil, false
	}
	return o.Quotas, true
}

// HasQuotas returns a boolean if a field has been set.
func (o *TariffDto) IsQuotasSet() bool {
	if o != nil && !IsNil(o.Quotas) {
		return true
	}

	return false
}

// SetQuotas gets a reference to the given []TariffQuotaDto and assigns it to the Quotas field.
func (o *TariffDto) SetQuotas(v []TariffQuotaDto) {
	o.Quotas = v
}

func (o TariffDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TariffDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OpenSource.IsSet() {
		toSerialize["openSource"] = o.OpenSource.Get()
	}
	if o.Enterprise.IsSet() {
		toSerialize["enterprise"] = o.Enterprise.Get()
	}
	if o.Developer.IsSet() {
		toSerialize["developer"] = o.Developer.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
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
	return toSerialize, nil
}

type NullableTariffDto struct {
	value *TariffDto
	isSet bool
}

func (v NullableTariffDto) Get() *TariffDto {
	return v.value
}

func (v *NullableTariffDto) Set(val *TariffDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTariffDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTariffDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTariffDto(val *TariffDto) *NullableTariffDto {
	return &NullableTariffDto{value: val, isSet: true}
}

func (v NullableTariffDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTariffDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

