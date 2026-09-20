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

// checks if the CustomerInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerInfoDto{}

// CustomerInfoDto The billing customer behind the portal, and which portal member pays for it.
type CustomerInfoDto struct {
	// The portal's identifier in the billing system, which is what support and invoices refer to. It is not the  portal alias.
	PortalId NullableString `json:"portalId,omitempty"`
	// Whether a payment method is stored for the account and usable. Without one the portal can hold a wallet  balance but cannot be charged automatically.
	PaymentMethodStatus *PaymentMethodStatus `json:"paymentMethodStatus,omitempty"`
	// The customer's payment method type.
	PaymentMethodType NullableString `json:"paymentMethodType,omitempty"`
	// Indicates whether the customer's payment method is delayed, i.e. the money reaches the wallet only after  the transfer settles rather than immediately.
	IsDelayedPaymentMethod *bool `json:"isDelayedPaymentMethod,omitempty"`
	// The address the billing account is registered to, lower-cased. It need not belong to a portal member,  which is exactly when `payer` stays empty.
	Email NullableString `json:"email,omitempty"`
	// The portal member whose account is behind the billing address. It is empty when `email` matches no member  of this portal, and while it is empty every operation of this group that only the payer may call is out  of reach for everybody.
	Payer *EmployeeDto `json:"payer,omitempty"`
}

// NewCustomerInfoDto instantiates a new CustomerInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerInfoDto() *CustomerInfoDto {
	this := CustomerInfoDto{}
	return &this
}

// NewCustomerInfoDtoWithDefaults instantiates a new CustomerInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerInfoDtoWithDefaults() *CustomerInfoDto {
	this := CustomerInfoDto{}
	return &this
}

// GetPortalId returns the PortalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerInfoDto) GetPortalId() string {
	if o == nil || IsNil(o.PortalId.Get()) {
		var ret string
		return ret
	}
	return *o.PortalId.Get()
}

// GetPortalIdOk returns a tuple with the PortalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerInfoDto) GetPortalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortalId.Get(), o.PortalId.IsSet()
}

// HasPortalId returns a boolean if a field has been set.
func (o *CustomerInfoDto) IsPortalIdSet() bool {
	if o != nil && o.PortalId.IsSet() {
		return true
	}

	return false
}

// SetPortalId gets a reference to the given NullableString and assigns it to the PortalId field.
func (o *CustomerInfoDto) SetPortalId(v string) {
	o.PortalId.Set(&v)
}
// SetPortalIdNil sets the value for PortalId to be an explicit nil
func (o *CustomerInfoDto) SetPortalIdNil() {
	o.PortalId.Set(nil)
}

// UnsetPortalId ensures that no value is present for PortalId, not even an explicit nil
func (o *CustomerInfoDto) UnsetPortalId() {
	o.PortalId.Unset()
}

// GetPaymentMethodStatus returns the PaymentMethodStatus field value if set, zero value otherwise.
func (o *CustomerInfoDto) GetPaymentMethodStatus() PaymentMethodStatus {
	if o == nil || IsNil(o.PaymentMethodStatus) {
		var ret PaymentMethodStatus
		return ret
	}
	return *o.PaymentMethodStatus
}

// GetPaymentMethodStatusOk returns a tuple with the PaymentMethodStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfoDto) GetPaymentMethodStatusOk() (*PaymentMethodStatus, bool) {
	if o == nil || IsNil(o.PaymentMethodStatus) {
		return nil, false
	}
	return o.PaymentMethodStatus, true
}

// HasPaymentMethodStatus returns a boolean if a field has been set.
func (o *CustomerInfoDto) IsPaymentMethodStatusSet() bool {
	if o != nil && !IsNil(o.PaymentMethodStatus) {
		return true
	}

	return false
}

// SetPaymentMethodStatus gets a reference to the given PaymentMethodStatus and assigns it to the PaymentMethodStatus field.
func (o *CustomerInfoDto) SetPaymentMethodStatus(v PaymentMethodStatus) {
	o.PaymentMethodStatus = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerInfoDto) GetPaymentMethodType() string {
	if o == nil || IsNil(o.PaymentMethodType.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType.Get()
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerInfoDto) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodType.Get(), o.PaymentMethodType.IsSet()
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *CustomerInfoDto) IsPaymentMethodTypeSet() bool {
	if o != nil && o.PaymentMethodType.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given NullableString and assigns it to the PaymentMethodType field.
func (o *CustomerInfoDto) SetPaymentMethodType(v string) {
	o.PaymentMethodType.Set(&v)
}
// SetPaymentMethodTypeNil sets the value for PaymentMethodType to be an explicit nil
func (o *CustomerInfoDto) SetPaymentMethodTypeNil() {
	o.PaymentMethodType.Set(nil)
}

// UnsetPaymentMethodType ensures that no value is present for PaymentMethodType, not even an explicit nil
func (o *CustomerInfoDto) UnsetPaymentMethodType() {
	o.PaymentMethodType.Unset()
}

// GetIsDelayedPaymentMethod returns the IsDelayedPaymentMethod field value if set, zero value otherwise.
func (o *CustomerInfoDto) GetIsDelayedPaymentMethod() bool {
	if o == nil || IsNil(o.IsDelayedPaymentMethod) {
		var ret bool
		return ret
	}
	return *o.IsDelayedPaymentMethod
}

// GetIsDelayedPaymentMethodOk returns a tuple with the IsDelayedPaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfoDto) GetIsDelayedPaymentMethodOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDelayedPaymentMethod) {
		return nil, false
	}
	return o.IsDelayedPaymentMethod, true
}

// HasIsDelayedPaymentMethod returns a boolean if a field has been set.
func (o *CustomerInfoDto) IsIsDelayedPaymentMethodSet() bool {
	if o != nil && !IsNil(o.IsDelayedPaymentMethod) {
		return true
	}

	return false
}

// SetIsDelayedPaymentMethod gets a reference to the given bool and assigns it to the IsDelayedPaymentMethod field.
func (o *CustomerInfoDto) SetIsDelayedPaymentMethod(v bool) {
	o.IsDelayedPaymentMethod = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerInfoDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerInfoDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerInfoDto) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *CustomerInfoDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *CustomerInfoDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *CustomerInfoDto) UnsetEmail() {
	o.Email.Unset()
}

// GetPayer returns the Payer field value if set, zero value otherwise.
func (o *CustomerInfoDto) GetPayer() EmployeeDto {
	if o == nil || IsNil(o.Payer) {
		var ret EmployeeDto
		return ret
	}
	return *o.Payer
}

// GetPayerOk returns a tuple with the Payer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfoDto) GetPayerOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.Payer) {
		return nil, false
	}
	return o.Payer, true
}

// HasPayer returns a boolean if a field has been set.
func (o *CustomerInfoDto) IsPayerSet() bool {
	if o != nil && !IsNil(o.Payer) {
		return true
	}

	return false
}

// SetPayer gets a reference to the given EmployeeDto and assigns it to the Payer field.
func (o *CustomerInfoDto) SetPayer(v EmployeeDto) {
	o.Payer = &v
}

func (o CustomerInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PortalId.IsSet() {
		toSerialize["portalId"] = o.PortalId.Get()
	}
	if !IsNil(o.PaymentMethodStatus) {
		toSerialize["paymentMethodStatus"] = o.PaymentMethodStatus
	}
	if o.PaymentMethodType.IsSet() {
		toSerialize["paymentMethodType"] = o.PaymentMethodType.Get()
	}
	if !IsNil(o.IsDelayedPaymentMethod) {
		toSerialize["isDelayedPaymentMethod"] = o.IsDelayedPaymentMethod
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.Payer) {
		toSerialize["payer"] = o.Payer
	}
	return toSerialize, nil
}

type NullableCustomerInfoDto struct {
	value *CustomerInfoDto
	isSet bool
}

func (v NullableCustomerInfoDto) Get() *CustomerInfoDto {
	return v.value
}

func (v *NullableCustomerInfoDto) Set(val *CustomerInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerInfoDto(val *CustomerInfoDto) *NullableCustomerInfoDto {
	return &NullableCustomerInfoDto{value: val, isSet: true}
}

func (v NullableCustomerInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

