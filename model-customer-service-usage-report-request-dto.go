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

// checks if the CustomerServiceUsageReportRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerServiceUsageReportRequestDto{}

// CustomerServiceUsageReportRequestDto The filters that select which wallet service consumption is reported: the services, the period, the participant,  the outcome, the usage metadata and the ordering.
type CustomerServiceUsageReportRequestDto struct {
	// The wallet services whose consumption is reported, named the way the billing catalogue names them -  `backup`, `ai-tools`, `ai-search`, `disk-storage`, `docscloud`. Take the values from the `serviceName` field  of `GET api/2.0/portal/payment/walletservices`; the match ignores case, a name this installation does not  sell fails the call with 404, and an omitted list reports every service. A bare string is accepted in place  of an array for backward compatibility.
	ServiceName []string `json:"serviceName,omitempty"`
	// The beginning of the reported period, inclusive. Read in the portal time zone rather than in UTC, and  defaults to the portal creation date.
	StartDate NullableTime `json:"startDate,omitempty"`
	// The end of the reported period, inclusive. Read in the portal time zone rather than in UTC, and defaults to  the moment the call is made.
	EndDate NullableTime `json:"endDate,omitempty"`
	// The participant whose consumption is reported - the account the accounting service records as the consumer.  Consumption caused by a portal user carries that user ID here; surrounding whitespace is trimmed, and an  omitted value reports every participant.
	ParticipantName NullableString `json:"participantName,omitempty"`
	// The outcome to keep. Consumption that is still being settled is reported as pending and may change later,  while the other outcomes are final; every outcome is reported when this is omitted.
	Status *OperationStatus `json:"status,omitempty"`
	// The usage annotations a wallet service records alongside its consumption, as the key and value pairs that  must all match for a record to be reported. The keys are chosen by the service that writes them, so read  them off the `metadata` of the records returned by `GET api/2.0/portal/payment/customer/usage` rather than  guessing; an omitted map reports every record.
	Metadata map[string]*string `json:"metadata,omitempty"`
	// The name of the field the per-service totals are sorted by, spelled as the accounting service names it, such  as `ServiceName` or `StartDate`. Surrounding whitespace is trimmed, and the accounting service applies its  own ordering when this is omitted.
	OrderBy NullableString `json:"orderBy,omitempty"`
	// The direction the field named in `orderBy` is sorted in. Newest or largest first is what the accounting  service does by default, so leaving this out sorts the same way as asking for descending explicitly.
	OrderType *OperationOrderType `json:"orderType,omitempty"`
}

// NewCustomerServiceUsageReportRequestDto instantiates a new CustomerServiceUsageReportRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerServiceUsageReportRequestDto() *CustomerServiceUsageReportRequestDto {
	this := CustomerServiceUsageReportRequestDto{}
	return &this
}

// NewCustomerServiceUsageReportRequestDtoWithDefaults instantiates a new CustomerServiceUsageReportRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerServiceUsageReportRequestDtoWithDefaults() *CustomerServiceUsageReportRequestDto {
	this := CustomerServiceUsageReportRequestDto{}
	return &this
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerServiceUsageReportRequestDto) GetServiceName() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerServiceUsageReportRequestDto) GetServiceNameOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *CustomerServiceUsageReportRequestDto) IsServiceNameSet() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given []string and assigns it to the ServiceName field.
func (o *CustomerServiceUsageReportRequestDto) SetServiceName(v []string) {
	o.ServiceName = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerServiceUsageReportRequestDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerServiceUsageReportRequestDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *CustomerServiceUsageReportRequestDto) IsStartDateSet() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *CustomerServiceUsageReportRequestDto) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *CustomerServiceUsageReportRequestDto) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *CustomerServiceUsageReportRequestDto) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerServiceUsageReportRequestDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerServiceUsageReportRequestDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *CustomerServiceUsageReportRequestDto) IsEndDateSet() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *CustomerServiceUsageReportRequestDto) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *CustomerServiceUsageReportRequestDto) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *CustomerServiceUsageReportRequestDto) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetParticipantName returns the ParticipantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerServiceUsageReportRequestDto) GetParticipantName() string {
	if o == nil || IsNil(o.ParticipantName.Get()) {
		var ret string
		return ret
	}
	return *o.ParticipantName.Get()
}

// GetParticipantNameOk returns a tuple with the ParticipantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerServiceUsageReportRequestDto) GetParticipantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParticipantName.Get(), o.ParticipantName.IsSet()
}

// HasParticipantName returns a boolean if a field has been set.
func (o *CustomerServiceUsageReportRequestDto) IsParticipantNameSet() bool {
	if o != nil && o.ParticipantName.IsSet() {
		return true
	}

	return false
}

// SetParticipantName gets a reference to the given NullableString and assigns it to the ParticipantName field.
func (o *CustomerServiceUsageReportRequestDto) SetParticipantName(v string) {
	o.ParticipantName.Set(&v)
}
// SetParticipantNameNil sets the value for ParticipantName to be an explicit nil
func (o *CustomerServiceUsageReportRequestDto) SetParticipantNameNil() {
	o.ParticipantName.Set(nil)
}

// UnsetParticipantName ensures that no value is present for ParticipantName, not even an explicit nil
func (o *CustomerServiceUsageReportRequestDto) UnsetParticipantName() {
	o.ParticipantName.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomerServiceUsageReportRequestDto) GetStatus() OperationStatus {
	if o == nil || IsNil(o.Status) {
		var ret OperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceUsageReportRequestDto) GetStatusOk() (*OperationStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomerServiceUsageReportRequestDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OperationStatus and assigns it to the Status field.
func (o *CustomerServiceUsageReportRequestDto) SetStatus(v OperationStatus) {
	o.Status = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomerServiceUsageReportRequestDto) GetMetadata() map[string]*string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]*string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceUsageReportRequestDto) GetMetadataOk() (map[string]*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]*string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomerServiceUsageReportRequestDto) IsMetadataSet() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]*string and assigns it to the Metadata field.
func (o *CustomerServiceUsageReportRequestDto) SetMetadata(v map[string]*string) {
	o.Metadata = v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerServiceUsageReportRequestDto) GetOrderBy() string {
	if o == nil || IsNil(o.OrderBy.Get()) {
		var ret string
		return ret
	}
	return *o.OrderBy.Get()
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerServiceUsageReportRequestDto) GetOrderByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderBy.Get(), o.OrderBy.IsSet()
}

// HasOrderBy returns a boolean if a field has been set.
func (o *CustomerServiceUsageReportRequestDto) IsOrderBySet() bool {
	if o != nil && o.OrderBy.IsSet() {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given NullableString and assigns it to the OrderBy field.
func (o *CustomerServiceUsageReportRequestDto) SetOrderBy(v string) {
	o.OrderBy.Set(&v)
}
// SetOrderByNil sets the value for OrderBy to be an explicit nil
func (o *CustomerServiceUsageReportRequestDto) SetOrderByNil() {
	o.OrderBy.Set(nil)
}

// UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
func (o *CustomerServiceUsageReportRequestDto) UnsetOrderBy() {
	o.OrderBy.Unset()
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *CustomerServiceUsageReportRequestDto) GetOrderType() OperationOrderType {
	if o == nil || IsNil(o.OrderType) {
		var ret OperationOrderType
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceUsageReportRequestDto) GetOrderTypeOk() (*OperationOrderType, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *CustomerServiceUsageReportRequestDto) IsOrderTypeSet() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given OperationOrderType and assigns it to the OrderType field.
func (o *CustomerServiceUsageReportRequestDto) SetOrderType(v OperationOrderType) {
	o.OrderType = &v
}

func (o CustomerServiceUsageReportRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerServiceUsageReportRequestDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OrderBy.IsSet() {
		toSerialize["orderBy"] = o.OrderBy.Get()
	}
	if !IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	return toSerialize, nil
}

type NullableCustomerServiceUsageReportRequestDto struct {
	value *CustomerServiceUsageReportRequestDto
	isSet bool
}

func (v NullableCustomerServiceUsageReportRequestDto) Get() *CustomerServiceUsageReportRequestDto {
	return v.value
}

func (v *NullableCustomerServiceUsageReportRequestDto) Set(val *CustomerServiceUsageReportRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerServiceUsageReportRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerServiceUsageReportRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerServiceUsageReportRequestDto(val *CustomerServiceUsageReportRequestDto) *NullableCustomerServiceUsageReportRequestDto {
	return &NullableCustomerServiceUsageReportRequestDto{value: val, isSet: true}
}

func (v NullableCustomerServiceUsageReportRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerServiceUsageReportRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

