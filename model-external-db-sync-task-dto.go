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
	"bytes"
	"fmt"
)

// checks if the ExternalDbSyncTaskDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalDbSyncTaskDto{}

// ExternalDbSyncTaskDto The external DB synchronization task parameters.
type ExternalDbSyncTaskDto struct {
	// The task ID.
	Id NullableString `json:"id"`
	// The error message if the synchronization failed.
	Error NullableString `json:"error,omitempty"`
	// The progress percentage of the synchronization.
	Percentage int32 `json:"percentage"`
	// Specifies whether the synchronization is completed or not.
	IsCompleted bool `json:"isCompleted"`
	// The status of the synchronization task.
	Status DistributedTaskStatus `json:"status"`
	// The synchronization results for all original forms in the room.
	Forms []ExternalDbSyncFormResultDto `json:"forms"`
}

type _ExternalDbSyncTaskDto ExternalDbSyncTaskDto

// NewExternalDbSyncTaskDto instantiates a new ExternalDbSyncTaskDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalDbSyncTaskDto(id NullableString, percentage int32, isCompleted bool, status DistributedTaskStatus, forms []ExternalDbSyncFormResultDto) *ExternalDbSyncTaskDto {
	this := ExternalDbSyncTaskDto{}
	this.Id = id
	this.Percentage = percentage
	this.IsCompleted = isCompleted
	this.Status = status
	this.Forms = forms
	return &this
}

// NewExternalDbSyncTaskDtoWithDefaults instantiates a new ExternalDbSyncTaskDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalDbSyncTaskDtoWithDefaults() *ExternalDbSyncTaskDto {
	this := ExternalDbSyncTaskDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalDbSyncTaskDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDbSyncTaskDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *ExternalDbSyncTaskDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalDbSyncTaskDto) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDbSyncTaskDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ExternalDbSyncTaskDto) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ExternalDbSyncTaskDto) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ExternalDbSyncTaskDto) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ExternalDbSyncTaskDto) UnsetError() {
	o.Error.Unset()
}

// GetPercentage returns the Percentage field value
func (o *ExternalDbSyncTaskDto) GetPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *ExternalDbSyncTaskDto) GetPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value
func (o *ExternalDbSyncTaskDto) SetPercentage(v int32) {
	o.Percentage = v
}

// GetIsCompleted returns the IsCompleted field value
func (o *ExternalDbSyncTaskDto) GetIsCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value
// and a boolean to check if the value has been set.
func (o *ExternalDbSyncTaskDto) GetIsCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCompleted, true
}

// SetIsCompleted sets field value
func (o *ExternalDbSyncTaskDto) SetIsCompleted(v bool) {
	o.IsCompleted = v
}

// GetStatus returns the Status field value
func (o *ExternalDbSyncTaskDto) GetStatus() DistributedTaskStatus {
	if o == nil {
		var ret DistributedTaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExternalDbSyncTaskDto) GetStatusOk() (*DistributedTaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExternalDbSyncTaskDto) SetStatus(v DistributedTaskStatus) {
	o.Status = v
}

// GetForms returns the Forms field value
// If the value is explicit nil, the zero value for []ExternalDbSyncFormResultDto will be returned
func (o *ExternalDbSyncTaskDto) GetForms() []ExternalDbSyncFormResultDto {
	if o == nil {
		var ret []ExternalDbSyncFormResultDto
		return ret
	}

	return o.Forms
}

// GetFormsOk returns a tuple with the Forms field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDbSyncTaskDto) GetFormsOk() ([]ExternalDbSyncFormResultDto, bool) {
	if o == nil || IsNil(o.Forms) {
		return nil, false
	}
	return o.Forms, true
}

// SetForms sets field value
func (o *ExternalDbSyncTaskDto) SetForms(v []ExternalDbSyncFormResultDto) {
	o.Forms = v
}

func (o ExternalDbSyncTaskDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalDbSyncTaskDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	toSerialize["percentage"] = o.Percentage
	toSerialize["isCompleted"] = o.IsCompleted
	toSerialize["status"] = o.Status
	if o.Forms != nil {
		toSerialize["forms"] = o.Forms
	}
	return toSerialize, nil
}

func (o *ExternalDbSyncTaskDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"percentage",
		"isCompleted",
		"status",
		"forms",
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

	varExternalDbSyncTaskDto := _ExternalDbSyncTaskDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varExternalDbSyncTaskDto)

	if err != nil {
		return err
	}

	*o = ExternalDbSyncTaskDto(varExternalDbSyncTaskDto)

	return err
}

type NullableExternalDbSyncTaskDto struct {
	value *ExternalDbSyncTaskDto
	isSet bool
}

func (v NullableExternalDbSyncTaskDto) Get() *ExternalDbSyncTaskDto {
	return v.value
}

func (v *NullableExternalDbSyncTaskDto) Set(val *ExternalDbSyncTaskDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalDbSyncTaskDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalDbSyncTaskDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalDbSyncTaskDto(val *ExternalDbSyncTaskDto) *NullableExternalDbSyncTaskDto {
	return &NullableExternalDbSyncTaskDto{value: val, isSet: true}
}

func (v NullableExternalDbSyncTaskDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalDbSyncTaskDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

