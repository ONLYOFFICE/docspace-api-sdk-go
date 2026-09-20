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

// checks if the CustomerOperationsReportRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerOperationsReportRequestDto{}

// CustomerOperationsReportRequestDto The filters that select which wallet movements are reported: the services, the period, the participant, the  direction and the outcome of the movement, and the ordering.
type CustomerOperationsReportRequestDto struct {
	// The wallet services whose movements are kept, named the way the billing catalogue names them - `backup`,  `ai-tools`, `ai-search`, `disk-storage`, `docscloud`. Take the values from the `serviceName` field of  `GET api/2.0/portal/payment/walletservices`; the match ignores case, a name this installation does not sell  fails the call with 404, and an omitted list keeps every service. A bare string is accepted in place of an  array for backward compatibility.
	ServiceName []string `json:"serviceName,omitempty"`
	// The beginning of the reported period, inclusive. Read in the portal time zone rather than in UTC, so a  movement at the edge of the period falls where the portal sees it; defaults to the portal creation date.
	StartDate NullableTime `json:"startDate,omitempty"`
	// The end of the reported period, inclusive. Read in the portal time zone rather than in UTC, and defaults to  the moment the call is made.
	EndDate NullableTime `json:"endDate,omitempty"`
	// The participant whose movements are kept - the account the accounting service records as the cause of a  movement. A movement caused by a portal user carries that user ID here, and one caused by the portal itself  carries the customer name; surrounding whitespace is trimmed, and an omitted value keeps every participant.
	ParticipantName NullableString `json:"participantName,omitempty"`
	// Whether movements that add money to the wallet - top-ups, refunds and corrections in the portal's favour -  are kept. Both directions are reported when neither this nor `debit` is given.
	Credit NullableBool `json:"credit,omitempty"`
	// Whether movements that take money out of the wallet - the charges of the wallet services - are kept. Both  directions are reported when neither this nor `credit` is given.
	Debit NullableBool `json:"debit,omitempty"`
	// The kind of movement to keep, which says what caused the money to move rather than how it ended. Every kind  is reported when it is omitted.
	Type *OperationType `json:"type,omitempty"`
	// The outcome to keep. A movement that is still being settled is reported as pending and may change later,  while the other outcomes are final; every outcome is reported when this is omitted.
	Status *OperationStatus `json:"status,omitempty"`
	// The name of the field the movements are sorted by, spelled as the accounting service names it, such as  `StartDate` or `ServiceName`. Surrounding whitespace is trimmed, and the accounting service applies its own  ordering when this is omitted.
	OrderBy NullableString `json:"orderBy,omitempty"`
	// The direction the field named in `orderBy` is sorted in. Newest or largest first is what the accounting  service does by default, so leaving this out sorts the same way as asking for descending explicitly.
	OrderType *OperationOrderType `json:"orderType,omitempty"`
}

// NewCustomerOperationsReportRequestDto instantiates a new CustomerOperationsReportRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerOperationsReportRequestDto() *CustomerOperationsReportRequestDto {
	this := CustomerOperationsReportRequestDto{}
	return &this
}

// NewCustomerOperationsReportRequestDtoWithDefaults instantiates a new CustomerOperationsReportRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerOperationsReportRequestDtoWithDefaults() *CustomerOperationsReportRequestDto {
	this := CustomerOperationsReportRequestDto{}
	return &this
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerOperationsReportRequestDto) GetServiceName() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerOperationsReportRequestDto) GetServiceNameOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsServiceNameSet() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given []string and assigns it to the ServiceName field.
func (o *CustomerOperationsReportRequestDto) SetServiceName(v []string) {
	o.ServiceName = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerOperationsReportRequestDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerOperationsReportRequestDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsStartDateSet() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *CustomerOperationsReportRequestDto) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *CustomerOperationsReportRequestDto) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *CustomerOperationsReportRequestDto) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerOperationsReportRequestDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerOperationsReportRequestDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsEndDateSet() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *CustomerOperationsReportRequestDto) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *CustomerOperationsReportRequestDto) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *CustomerOperationsReportRequestDto) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetParticipantName returns the ParticipantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerOperationsReportRequestDto) GetParticipantName() string {
	if o == nil || IsNil(o.ParticipantName.Get()) {
		var ret string
		return ret
	}
	return *o.ParticipantName.Get()
}

// GetParticipantNameOk returns a tuple with the ParticipantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerOperationsReportRequestDto) GetParticipantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParticipantName.Get(), o.ParticipantName.IsSet()
}

// HasParticipantName returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsParticipantNameSet() bool {
	if o != nil && o.ParticipantName.IsSet() {
		return true
	}

	return false
}

// SetParticipantName gets a reference to the given NullableString and assigns it to the ParticipantName field.
func (o *CustomerOperationsReportRequestDto) SetParticipantName(v string) {
	o.ParticipantName.Set(&v)
}
// SetParticipantNameNil sets the value for ParticipantName to be an explicit nil
func (o *CustomerOperationsReportRequestDto) SetParticipantNameNil() {
	o.ParticipantName.Set(nil)
}

// UnsetParticipantName ensures that no value is present for ParticipantName, not even an explicit nil
func (o *CustomerOperationsReportRequestDto) UnsetParticipantName() {
	o.ParticipantName.Unset()
}

// GetCredit returns the Credit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerOperationsReportRequestDto) GetCredit() bool {
	if o == nil || IsNil(o.Credit.Get()) {
		var ret bool
		return ret
	}
	return *o.Credit.Get()
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerOperationsReportRequestDto) GetCreditOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credit.Get(), o.Credit.IsSet()
}

// HasCredit returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsCreditSet() bool {
	if o != nil && o.Credit.IsSet() {
		return true
	}

	return false
}

// SetCredit gets a reference to the given NullableBool and assigns it to the Credit field.
func (o *CustomerOperationsReportRequestDto) SetCredit(v bool) {
	o.Credit.Set(&v)
}
// SetCreditNil sets the value for Credit to be an explicit nil
func (o *CustomerOperationsReportRequestDto) SetCreditNil() {
	o.Credit.Set(nil)
}

// UnsetCredit ensures that no value is present for Credit, not even an explicit nil
func (o *CustomerOperationsReportRequestDto) UnsetCredit() {
	o.Credit.Unset()
}

// GetDebit returns the Debit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerOperationsReportRequestDto) GetDebit() bool {
	if o == nil || IsNil(o.Debit.Get()) {
		var ret bool
		return ret
	}
	return *o.Debit.Get()
}

// GetDebitOk returns a tuple with the Debit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerOperationsReportRequestDto) GetDebitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Debit.Get(), o.Debit.IsSet()
}

// HasDebit returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsDebitSet() bool {
	if o != nil && o.Debit.IsSet() {
		return true
	}

	return false
}

// SetDebit gets a reference to the given NullableBool and assigns it to the Debit field.
func (o *CustomerOperationsReportRequestDto) SetDebit(v bool) {
	o.Debit.Set(&v)
}
// SetDebitNil sets the value for Debit to be an explicit nil
func (o *CustomerOperationsReportRequestDto) SetDebitNil() {
	o.Debit.Set(nil)
}

// UnsetDebit ensures that no value is present for Debit, not even an explicit nil
func (o *CustomerOperationsReportRequestDto) UnsetDebit() {
	o.Debit.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomerOperationsReportRequestDto) GetType() OperationType {
	if o == nil || IsNil(o.Type) {
		var ret OperationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOperationsReportRequestDto) GetTypeOk() (*OperationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OperationType and assigns it to the Type field.
func (o *CustomerOperationsReportRequestDto) SetType(v OperationType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomerOperationsReportRequestDto) GetStatus() OperationStatus {
	if o == nil || IsNil(o.Status) {
		var ret OperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOperationsReportRequestDto) GetStatusOk() (*OperationStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OperationStatus and assigns it to the Status field.
func (o *CustomerOperationsReportRequestDto) SetStatus(v OperationStatus) {
	o.Status = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerOperationsReportRequestDto) GetOrderBy() string {
	if o == nil || IsNil(o.OrderBy.Get()) {
		var ret string
		return ret
	}
	return *o.OrderBy.Get()
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerOperationsReportRequestDto) GetOrderByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderBy.Get(), o.OrderBy.IsSet()
}

// HasOrderBy returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsOrderBySet() bool {
	if o != nil && o.OrderBy.IsSet() {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given NullableString and assigns it to the OrderBy field.
func (o *CustomerOperationsReportRequestDto) SetOrderBy(v string) {
	o.OrderBy.Set(&v)
}
// SetOrderByNil sets the value for OrderBy to be an explicit nil
func (o *CustomerOperationsReportRequestDto) SetOrderByNil() {
	o.OrderBy.Set(nil)
}

// UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
func (o *CustomerOperationsReportRequestDto) UnsetOrderBy() {
	o.OrderBy.Unset()
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *CustomerOperationsReportRequestDto) GetOrderType() OperationOrderType {
	if o == nil || IsNil(o.OrderType) {
		var ret OperationOrderType
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOperationsReportRequestDto) GetOrderTypeOk() (*OperationOrderType, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *CustomerOperationsReportRequestDto) IsOrderTypeSet() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given OperationOrderType and assigns it to the OrderType field.
func (o *CustomerOperationsReportRequestDto) SetOrderType(v OperationOrderType) {
	o.OrderType = &v
}

func (o CustomerOperationsReportRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerOperationsReportRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceName != nil {
		toSerialize["serviceName"] = o.ServiceName
	}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if o.ParticipantName.IsSet() {
		toSerialize["participantName"] = o.ParticipantName.Get()
	}
	if o.Credit.IsSet() {
		toSerialize["credit"] = o.Credit.Get()
	}
	if o.Debit.IsSet() {
		toSerialize["debit"] = o.Debit.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.OrderBy.IsSet() {
		toSerialize["orderBy"] = o.OrderBy.Get()
	}
	if !IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	return toSerialize, nil
}

type NullableCustomerOperationsReportRequestDto struct {
	value *CustomerOperationsReportRequestDto
	isSet bool
}

func (v NullableCustomerOperationsReportRequestDto) Get() *CustomerOperationsReportRequestDto {
	return v.value
}

func (v *NullableCustomerOperationsReportRequestDto) Set(val *CustomerOperationsReportRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerOperationsReportRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerOperationsReportRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerOperationsReportRequestDto(val *CustomerOperationsReportRequestDto) *NullableCustomerOperationsReportRequestDto {
	return &NullableCustomerOperationsReportRequestDto{value: val, isSet: true}
}

func (v NullableCustomerOperationsReportRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerOperationsReportRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

