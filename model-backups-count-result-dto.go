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

// checks if the BackupsCountResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupsCountResultDto{}

// BackupsCountResultDto The number of backups.
type BackupsCountResultDto struct {
	// The number of free backups.
	Free *int32 `json:"free,omitempty"`
	// The number of paid backups.
	Paid *int32 `json:"paid,omitempty"`
}

// NewBackupsCountResultDto instantiates a new BackupsCountResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupsCountResultDto() *BackupsCountResultDto {
	this := BackupsCountResultDto{}
	return &this
}

// NewBackupsCountResultDtoWithDefaults instantiates a new BackupsCountResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupsCountResultDtoWithDefaults() *BackupsCountResultDto {
	this := BackupsCountResultDto{}
	return &this
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *BackupsCountResultDto) GetFree() int32 {
	if o == nil || IsNil(o.Free) {
		var ret int32
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsCountResultDto) GetFreeOk() (*int32, bool) {
	if o == nil || IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *BackupsCountResultDto) IsFreeSet() bool {
	if o != nil && !IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given int32 and assigns it to the Free field.
func (o *BackupsCountResultDto) SetFree(v int32) {
	o.Free = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *BackupsCountResultDto) GetPaid() int32 {
	if o == nil || IsNil(o.Paid) {
		var ret int32
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsCountResultDto) GetPaidOk() (*int32, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *BackupsCountResultDto) IsPaidSet() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given int32 and assigns it to the Paid field.
func (o *BackupsCountResultDto) SetPaid(v int32) {
	o.Paid = &v
}

func (o BackupsCountResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupsCountResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !IsNil(o.Paid) {
		toSerialize["paid"] = o.Paid
	}
	return toSerialize, nil
}

type NullableBackupsCountResultDto struct {
	value *BackupsCountResultDto
	isSet bool
}

func (v NullableBackupsCountResultDto) Get() *BackupsCountResultDto {
	return v.value
}

func (v *NullableBackupsCountResultDto) Set(val *BackupsCountResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupsCountResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupsCountResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupsCountResultDto(val *BackupsCountResultDto) *NullableBackupsCountResultDto {
	return &NullableBackupsCountResultDto{value: val, isSet: true}
}

func (v NullableBackupsCountResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupsCountResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

