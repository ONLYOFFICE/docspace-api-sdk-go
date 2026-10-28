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

// checks if the DocsCloudTenantDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudTenantDto{}

// DocsCloudTenantDto Represents a Docs Connect tenant of a portal.
type DocsCloudTenantDto struct {
	// The external ID of the dedicated resource the tenant is hosted on.
	DedicatedResourceExId *int32 `json:"dedicatedResourceExId,omitempty"`
	// The tenant alias.
	Alias NullableString `json:"alias,omitempty"`
	// The tenant name.
	Name NullableString `json:"name,omitempty"`
	// The date and time when the tenant was last modified.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// The customer ID.
	CustomerId NullableString `json:"customerId,omitempty"`
	// The customer name.
	CustomerName NullableString `json:"customerName,omitempty"`
	// The date and time when the tenant subscription ends.
	EndDate *time.Time `json:"endDate,omitempty"`
	// The resource type.
	ResourceType *int32 `json:"resourceType,omitempty"`
	// Whether the tenant is active (the end date is in the future).
	IsActive *bool `json:"isActive,omitempty"`
	// The tenant address.
	Address NullableString `json:"address,omitempty"`
	// The tenant payment information.
	Payment *DocsCloudPaymentDto `json:"payment,omitempty"`
}

// NewDocsCloudTenantDto instantiates a new DocsCloudTenantDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudTenantDto() *DocsCloudTenantDto {
	this := DocsCloudTenantDto{}
	return &this
}

// NewDocsCloudTenantDtoWithDefaults instantiates a new DocsCloudTenantDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudTenantDtoWithDefaults() *DocsCloudTenantDto {
	this := DocsCloudTenantDto{}
	return &this
}

// GetDedicatedResourceExId returns the DedicatedResourceExId field value if set, zero value otherwise.
func (o *DocsCloudTenantDto) GetDedicatedResourceExId() int32 {
	if o == nil || IsNil(o.DedicatedResourceExId) {
		var ret int32
		return ret
	}
	return *o.DedicatedResourceExId
}

// GetDedicatedResourceExIdOk returns a tuple with the DedicatedResourceExId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantDto) GetDedicatedResourceExIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DedicatedResourceExId) {
		return nil, false
	}
	return o.DedicatedResourceExId, true
}

// HasDedicatedResourceExId returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsDedicatedResourceExIdSet() bool {
	if o != nil && !IsNil(o.DedicatedResourceExId) {
		return true
	}

	return false
}

// SetDedicatedResourceExId gets a reference to the given int32 and assigns it to the DedicatedResourceExId field.
func (o *DocsCloudTenantDto) SetDedicatedResourceExId(v int32) {
	o.DedicatedResourceExId = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudTenantDto) GetAlias() string {
	if o == nil || IsNil(o.Alias.Get()) {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudTenantDto) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsAliasSet() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *DocsCloudTenantDto) SetAlias(v string) {
	o.Alias.Set(&v)
}
// SetAliasNil sets the value for Alias to be an explicit nil
func (o *DocsCloudTenantDto) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *DocsCloudTenantDto) UnsetAlias() {
	o.Alias.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudTenantDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudTenantDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DocsCloudTenantDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DocsCloudTenantDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DocsCloudTenantDto) UnsetName() {
	o.Name.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *DocsCloudTenantDto) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantDto) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsModifiedDateSet() bool {
	if o != nil && !IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *DocsCloudTenantDto) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudTenantDto) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudTenantDto) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsCustomerIdSet() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *DocsCloudTenantDto) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *DocsCloudTenantDto) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *DocsCloudTenantDto) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudTenantDto) GetCustomerName() string {
	if o == nil || IsNil(o.CustomerName.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerName.Get()
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudTenantDto) GetCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerName.Get(), o.CustomerName.IsSet()
}

// HasCustomerName returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsCustomerNameSet() bool {
	if o != nil && o.CustomerName.IsSet() {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given NullableString and assigns it to the CustomerName field.
func (o *DocsCloudTenantDto) SetCustomerName(v string) {
	o.CustomerName.Set(&v)
}
// SetCustomerNameNil sets the value for CustomerName to be an explicit nil
func (o *DocsCloudTenantDto) SetCustomerNameNil() {
	o.CustomerName.Set(nil)
}

// UnsetCustomerName ensures that no value is present for CustomerName, not even an explicit nil
func (o *DocsCloudTenantDto) UnsetCustomerName() {
	o.CustomerName.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DocsCloudTenantDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsEndDateSet() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *DocsCloudTenantDto) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *DocsCloudTenantDto) GetResourceType() int32 {
	if o == nil || IsNil(o.ResourceType) {
		var ret int32
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantDto) GetResourceTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsResourceTypeSet() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given int32 and assigns it to the ResourceType field.
func (o *DocsCloudTenantDto) SetResourceType(v int32) {
	o.ResourceType = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *DocsCloudTenantDto) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantDto) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsIsActiveSet() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *DocsCloudTenantDto) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudTenantDto) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudTenantDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsAddressSet() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *DocsCloudTenantDto) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *DocsCloudTenantDto) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *DocsCloudTenantDto) UnsetAddress() {
	o.Address.Unset()
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *DocsCloudTenantDto) GetPayment() DocsCloudPaymentDto {
	if o == nil || IsNil(o.Payment) {
		var ret DocsCloudPaymentDto
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantDto) GetPaymentOk() (*DocsCloudPaymentDto, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *DocsCloudTenantDto) IsPaymentSet() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given DocsCloudPaymentDto and assigns it to the Payment field.
func (o *DocsCloudTenantDto) SetPayment(v DocsCloudPaymentDto) {
	o.Payment = &v
}

func (o DocsCloudTenantDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudTenantDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DedicatedResourceExId) {
		toSerialize["dedicatedResourceExId"] = o.DedicatedResourceExId
	}
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if o.CustomerId.IsSet() {
		toSerialize["customerId"] = o.CustomerId.Get()
	}
	if o.CustomerName.IsSet() {
		toSerialize["customerName"] = o.CustomerName.Get()
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	return toSerialize, nil
}

type NullableDocsCloudTenantDto struct {
	value *DocsCloudTenantDto
	isSet bool
}

func (v NullableDocsCloudTenantDto) Get() *DocsCloudTenantDto {
	return v.value
}

func (v *NullableDocsCloudTenantDto) Set(val *DocsCloudTenantDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudTenantDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudTenantDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudTenantDto(val *DocsCloudTenantDto) *NullableDocsCloudTenantDto {
	return &NullableDocsCloudTenantDto{value: val, isSet: true}
}

func (v NullableDocsCloudTenantDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudTenantDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

