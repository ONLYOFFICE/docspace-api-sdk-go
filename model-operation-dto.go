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

// checks if the OperationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationDto{}

// OperationDto Represents an operation.
type OperationDto struct {
	Date *ApiDateTime `json:"date,omitempty"`
	// The service related to the operation.
	Service NullableString `json:"service,omitempty"`
	// The brief operation description.
	Description NullableString `json:"description,omitempty"`
	// The detailed information about the operation.
	Details NullableString `json:"details,omitempty"`
	// The service unit.
	ServiceUnit NullableString `json:"serviceUnit,omitempty"`
	// The quantity of the service used.
	Quantity *int32 `json:"quantity,omitempty"`
	// The three-character ISO 4217 currency symbol of the operation.
	Currency NullableString `json:"currency,omitempty"`
	// The credit amount of the operation.
	Credit *float64 `json:"credit,omitempty"`
	// The debit amount of the operation.
	Debit *float64 `json:"debit,omitempty"`
	// The participant original name.
	ParticipantName NullableString `json:"participantName,omitempty"`
	// The participant display name.
	ParticipantDisplayName NullableString `json:"participantDisplayName,omitempty"`
}

// NewOperationDto instantiates a new OperationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationDto() *OperationDto {
	this := OperationDto{}
	return &this
}

// NewOperationDtoWithDefaults instantiates a new OperationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationDtoWithDefaults() *OperationDto {
	this := OperationDto{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *OperationDto) GetDate() ApiDateTime {
	if o == nil || IsNil(o.Date) {
		var ret ApiDateTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationDto) GetDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *OperationDto) IsDateSet() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given ApiDateTime and assigns it to the Date field.
func (o *OperationDto) SetDate(v ApiDateTime) {
	o.Date = &v
}

// GetService returns the Service field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationDto) GetService() string {
	if o == nil || IsNil(o.Service.Get()) {
		var ret string
		return ret
	}
	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationDto) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// HasService returns a boolean if a field has been set.
func (o *OperationDto) IsServiceSet() bool {
	if o != nil && o.Service.IsSet() {
		return true
	}

	return false
}

// SetService gets a reference to the given NullableString and assigns it to the Service field.
func (o *OperationDto) SetService(v string) {
	o.Service.Set(&v)
}
// SetServiceNil sets the value for Service to be an explicit nil
func (o *OperationDto) SetServiceNil() {
	o.Service.Set(nil)
}

// UnsetService ensures that no value is present for Service, not even an explicit nil
func (o *OperationDto) UnsetService() {
	o.Service.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OperationDto) IsDescriptionSet() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OperationDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OperationDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OperationDto) UnsetDescription() {
	o.Description.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationDto) GetDetails() string {
	if o == nil || IsNil(o.Details.Get()) {
		var ret string
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationDto) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *OperationDto) IsDetailsSet() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableString and assigns it to the Details field.
func (o *OperationDto) SetDetails(v string) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *OperationDto) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *OperationDto) UnsetDetails() {
	o.Details.Unset()
}

// GetServiceUnit returns the ServiceUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationDto) GetServiceUnit() string {
	if o == nil || IsNil(o.ServiceUnit.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceUnit.Get()
}

// GetServiceUnitOk returns a tuple with the ServiceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationDto) GetServiceUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceUnit.Get(), o.ServiceUnit.IsSet()
}

// HasServiceUnit returns a boolean if a field has been set.
func (o *OperationDto) IsServiceUnitSet() bool {
	if o != nil && o.ServiceUnit.IsSet() {
		return true
	}

	return false
}

// SetServiceUnit gets a reference to the given NullableString and assigns it to the ServiceUnit field.
func (o *OperationDto) SetServiceUnit(v string) {
	o.ServiceUnit.Set(&v)
}
// SetServiceUnitNil sets the value for ServiceUnit to be an explicit nil
func (o *OperationDto) SetServiceUnitNil() {
	o.ServiceUnit.Set(nil)
}

// UnsetServiceUnit ensures that no value is present for ServiceUnit, not even an explicit nil
func (o *OperationDto) UnsetServiceUnit() {
	o.ServiceUnit.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OperationDto) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationDto) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OperationDto) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OperationDto) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *OperationDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *OperationDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *OperationDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *OperationDto) UnsetCurrency() {
	o.Currency.Unset()
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *OperationDto) GetCredit() float64 {
	if o == nil || IsNil(o.Credit) {
		var ret float64
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationDto) GetCreditOk() (*float64, bool) {
	if o == nil || IsNil(o.Credit) {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *OperationDto) IsCreditSet() bool {
	if o != nil && !IsNil(o.Credit) {
		return true
	}

	return false
}

// SetCredit gets a reference to the given float64 and assigns it to the Credit field.
func (o *OperationDto) SetCredit(v float64) {
	o.Credit = &v
}

// GetDebit returns the Debit field value if set, zero value otherwise.
func (o *OperationDto) GetDebit() float64 {
	if o == nil || IsNil(o.Debit) {
		var ret float64
		return ret
	}
	return *o.Debit
}

// GetDebitOk returns a tuple with the Debit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationDto) GetDebitOk() (*float64, bool) {
	if o == nil || IsNil(o.Debit) {
		return nil, false
	}
	return o.Debit, true
}

// HasDebit returns a boolean if a field has been set.
func (o *OperationDto) IsDebitSet() bool {
	if o != nil && !IsNil(o.Debit) {
		return true
	}

	return false
}

// SetDebit gets a reference to the given float64 and assigns it to the Debit field.
func (o *OperationDto) SetDebit(v float64) {
	o.Debit = &v
}

// GetParticipantName returns the ParticipantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationDto) GetParticipantName() string {
	if o == nil || IsNil(o.ParticipantName.Get()) {
		var ret string
		return ret
	}
	return *o.ParticipantName.Get()
}

// GetParticipantNameOk returns a tuple with the ParticipantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationDto) GetParticipantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParticipantName.Get(), o.ParticipantName.IsSet()
}

// HasParticipantName returns a boolean if a field has been set.
func (o *OperationDto) IsParticipantNameSet() bool {
	if o != nil && o.ParticipantName.IsSet() {
		return true
	}

	return false
}

// SetParticipantName gets a reference to the given NullableString and assigns it to the ParticipantName field.
func (o *OperationDto) SetParticipantName(v string) {
	o.ParticipantName.Set(&v)
}
// SetParticipantNameNil sets the value for ParticipantName to be an explicit nil
func (o *OperationDto) SetParticipantNameNil() {
	o.ParticipantName.Set(nil)
}

// UnsetParticipantName ensures that no value is present for ParticipantName, not even an explicit nil
func (o *OperationDto) UnsetParticipantName() {
	o.ParticipantName.Unset()
}

// GetParticipantDisplayName returns the ParticipantDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationDto) GetParticipantDisplayName() string {
	if o == nil || IsNil(o.ParticipantDisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.ParticipantDisplayName.Get()
}

// GetParticipantDisplayNameOk returns a tuple with the ParticipantDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationDto) GetParticipantDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParticipantDisplayName.Get(), o.ParticipantDisplayName.IsSet()
}

// HasParticipantDisplayName returns a boolean if a field has been set.
func (o *OperationDto) IsParticipantDisplayNameSet() bool {
	if o != nil && o.ParticipantDisplayName.IsSet() {
		return true
	}

	return false
}

// SetParticipantDisplayName gets a reference to the given NullableString and assigns it to the ParticipantDisplayName field.
func (o *OperationDto) SetParticipantDisplayName(v string) {
	o.ParticipantDisplayName.Set(&v)
}
// SetParticipantDisplayNameNil sets the value for ParticipantDisplayName to be an explicit nil
func (o *OperationDto) SetParticipantDisplayNameNil() {
	o.ParticipantDisplayName.Set(nil)
}

// UnsetParticipantDisplayName ensures that no value is present for ParticipantDisplayName, not even an explicit nil
func (o *OperationDto) UnsetParticipantDisplayName() {
	o.ParticipantDisplayName.Unset()
}

func (o OperationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if o.Service.IsSet() {
		toSerialize["service"] = o.Service.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.ServiceUnit.IsSet() {
		toSerialize["serviceUnit"] = o.ServiceUnit.Get()
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.Credit) {
		toSerialize["credit"] = o.Credit
	}
	if !IsNil(o.Debit) {
		toSerialize["debit"] = o.Debit
	}
	if o.ParticipantName.IsSet() {
		toSerialize["participantName"] = o.ParticipantName.Get()
	}
	if o.ParticipantDisplayName.IsSet() {
		toSerialize["participantDisplayName"] = o.ParticipantDisplayName.Get()
	}
	return toSerialize, nil
}

type NullableOperationDto struct {
	value *OperationDto
	isSet bool
}

func (v NullableOperationDto) Get() *OperationDto {
	return v.value
}

func (v *NullableOperationDto) Set(val *OperationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationDto(val *OperationDto) *NullableOperationDto {
	return &NullableOperationDto{value: val, isSet: true}
}

func (v NullableOperationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

