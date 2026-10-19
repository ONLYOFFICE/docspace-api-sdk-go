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

// checks if the MetadataOperationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataOperationDto{}

// MetadataOperationDto The cascade metadata assignment operation status.
type MetadataOperationDto struct {
	// The operation ID.
	Id NullableString `json:"id,omitempty"`
	// The operation progress percentage.
	Progress *float64 `json:"progress,omitempty"`
	// Specifies if the operation is completed.
	IsCompleted *bool `json:"isCompleted,omitempty"`
	// The operation error message.
	Error NullableString `json:"error,omitempty"`
	// The IDs of the metadata templates the operation propagates to the subtree. Empty when there is no operation to  report, so the answer describes nothing in progress.
	TemplateIds []int32 `json:"templateIds,omitempty"`
}

// NewMetadataOperationDto instantiates a new MetadataOperationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataOperationDto() *MetadataOperationDto {
	this := MetadataOperationDto{}
	return &this
}

// NewMetadataOperationDtoWithDefaults instantiates a new MetadataOperationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataOperationDtoWithDefaults() *MetadataOperationDto {
	this := MetadataOperationDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOperationDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOperationDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MetadataOperationDto) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MetadataOperationDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MetadataOperationDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MetadataOperationDto) UnsetId() {
	o.Id.Unset()
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *MetadataOperationDto) GetProgress() float64 {
	if o == nil || IsNil(o.Progress) {
		var ret float64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOperationDto) GetProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *MetadataOperationDto) IsProgressSet() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float64 and assigns it to the Progress field.
func (o *MetadataOperationDto) SetProgress(v float64) {
	o.Progress = &v
}

// GetIsCompleted returns the IsCompleted field value if set, zero value otherwise.
func (o *MetadataOperationDto) GetIsCompleted() bool {
	if o == nil || IsNil(o.IsCompleted) {
		var ret bool
		return ret
	}
	return *o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOperationDto) GetIsCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompleted) {
		return nil, false
	}
	return o.IsCompleted, true
}

// HasIsCompleted returns a boolean if a field has been set.
func (o *MetadataOperationDto) IsIsCompletedSet() bool {
	if o != nil && !IsNil(o.IsCompleted) {
		return true
	}

	return false
}

// SetIsCompleted gets a reference to the given bool and assigns it to the IsCompleted field.
func (o *MetadataOperationDto) SetIsCompleted(v bool) {
	o.IsCompleted = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOperationDto) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOperationDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *MetadataOperationDto) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *MetadataOperationDto) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *MetadataOperationDto) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *MetadataOperationDto) UnsetError() {
	o.Error.Unset()
}

// GetTemplateIds returns the TemplateIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOperationDto) GetTemplateIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.TemplateIds
}

// GetTemplateIdsOk returns a tuple with the TemplateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOperationDto) GetTemplateIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.TemplateIds) {
		return nil, false
	}
	return o.TemplateIds, true
}

// HasTemplateIds returns a boolean if a field has been set.
func (o *MetadataOperationDto) IsTemplateIdsSet() bool {
	if o != nil && !IsNil(o.TemplateIds) {
		return true
	}

	return false
}

// SetTemplateIds gets a reference to the given []int32 and assigns it to the TemplateIds field.
func (o *MetadataOperationDto) SetTemplateIds(v []int32) {
	o.TemplateIds = v
}

func (o MetadataOperationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataOperationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.IsCompleted) {
		toSerialize["isCompleted"] = o.IsCompleted
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.TemplateIds != nil {
		toSerialize["templateIds"] = o.TemplateIds
	}
	return toSerialize, nil
}

type NullableMetadataOperationDto struct {
	value *MetadataOperationDto
	isSet bool
}

func (v NullableMetadataOperationDto) Get() *MetadataOperationDto {
	return v.value
}

func (v *NullableMetadataOperationDto) Set(val *MetadataOperationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataOperationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataOperationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataOperationDto(val *MetadataOperationDto) *NullableMetadataOperationDto {
	return &NullableMetadataOperationDto{value: val, isSet: true}
}

func (v NullableMetadataOperationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataOperationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

