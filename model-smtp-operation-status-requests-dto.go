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

// checks if the SmtpOperationStatusRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmtpOperationStatusRequestsDto{}

// SmtpOperationStatusRequestsDto The state of the background job that sends the portal SMTP test message.
type SmtpOperationStatusRequestsDto struct {
	// Whether the job has finished. This is the field to poll; the first answer that reports it true also discards  the job, so read `error` out of that same answer rather than calling again.
	Completed *bool `json:"completed,omitempty"`
	// The identifier of the queued job. A portal only ever has one test job at a time, so it names the run rather  than selecting among several.
	Id NullableString `json:"id,omitempty"`
	// Why the test failed. It stays empty while the job runs and also once the relay has accepted the message, so  an empty value on a finished job is what success looks like; an unreachable relay is reported here after a  30-second connection timeout rather than as a failed request.
	Error NullableString `json:"error,omitempty"`
	// The step the job has reached, in words - `Connect to host` or `Send test message`, for instance. It is meant  to be shown to a person and is not a fixed set of values to branch on.
	Status NullableString `json:"status,omitempty"`
	// How far the job has got, as a percentage climbing to 100. Reaching 100 says the job ran to the end, not that  the message was accepted - that is what an empty `error` says.
	Percents *int32 `json:"percents,omitempty"`
}

// NewSmtpOperationStatusRequestsDto instantiates a new SmtpOperationStatusRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpOperationStatusRequestsDto() *SmtpOperationStatusRequestsDto {
	this := SmtpOperationStatusRequestsDto{}
	return &this
}

// NewSmtpOperationStatusRequestsDtoWithDefaults instantiates a new SmtpOperationStatusRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpOperationStatusRequestsDtoWithDefaults() *SmtpOperationStatusRequestsDto {
	this := SmtpOperationStatusRequestsDto{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *SmtpOperationStatusRequestsDto) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpOperationStatusRequestsDto) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *SmtpOperationStatusRequestsDto) IsCompletedSet() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *SmtpOperationStatusRequestsDto) SetCompleted(v bool) {
	o.Completed = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpOperationStatusRequestsDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpOperationStatusRequestsDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SmtpOperationStatusRequestsDto) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *SmtpOperationStatusRequestsDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SmtpOperationStatusRequestsDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SmtpOperationStatusRequestsDto) UnsetId() {
	o.Id.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpOperationStatusRequestsDto) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpOperationStatusRequestsDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *SmtpOperationStatusRequestsDto) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *SmtpOperationStatusRequestsDto) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *SmtpOperationStatusRequestsDto) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *SmtpOperationStatusRequestsDto) UnsetError() {
	o.Error.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpOperationStatusRequestsDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpOperationStatusRequestsDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *SmtpOperationStatusRequestsDto) IsStatusSet() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *SmtpOperationStatusRequestsDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *SmtpOperationStatusRequestsDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *SmtpOperationStatusRequestsDto) UnsetStatus() {
	o.Status.Unset()
}

// GetPercents returns the Percents field value if set, zero value otherwise.
func (o *SmtpOperationStatusRequestsDto) GetPercents() int32 {
	if o == nil || IsNil(o.Percents) {
		var ret int32
		return ret
	}
	return *o.Percents
}

// GetPercentsOk returns a tuple with the Percents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpOperationStatusRequestsDto) GetPercentsOk() (*int32, bool) {
	if o == nil || IsNil(o.Percents) {
		return nil, false
	}
	return o.Percents, true
}

// HasPercents returns a boolean if a field has been set.
func (o *SmtpOperationStatusRequestsDto) IsPercentsSet() bool {
	if o != nil && !IsNil(o.Percents) {
		return true
	}

	return false
}

// SetPercents gets a reference to the given int32 and assigns it to the Percents field.
func (o *SmtpOperationStatusRequestsDto) SetPercents(v int32) {
	o.Percents = &v
}

func (o SmtpOperationStatusRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmtpOperationStatusRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.Percents) {
		toSerialize["percents"] = o.Percents
	}
	return toSerialize, nil
}

type NullableSmtpOperationStatusRequestsDto struct {
	value *SmtpOperationStatusRequestsDto
	isSet bool
}

func (v NullableSmtpOperationStatusRequestsDto) Get() *SmtpOperationStatusRequestsDto {
	return v.value
}

func (v *NullableSmtpOperationStatusRequestsDto) Set(val *SmtpOperationStatusRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpOperationStatusRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpOperationStatusRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpOperationStatusRequestsDto(val *SmtpOperationStatusRequestsDto) *NullableSmtpOperationStatusRequestsDto {
	return &NullableSmtpOperationStatusRequestsDto{value: val, isSet: true}
}

func (v NullableSmtpOperationStatusRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpOperationStatusRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

