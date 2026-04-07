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

// checks if the MigrationStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrationStatusDto{}

// MigrationStatusDto The migration status parameters.
type MigrationStatusDto struct {
	// The migration progress.
	Progress *float64 `json:"progress,omitempty"`
	// The migration error.
	Error NullableString `json:"error,omitempty"`
	ParseResult *MigrationApiInfo `json:"parseResult,omitempty"`
	// Specifies whether the migration is completed or not.
	IsCompleted *bool `json:"isCompleted,omitempty"`
}

// NewMigrationStatusDto instantiates a new MigrationStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationStatusDto() *MigrationStatusDto {
	this := MigrationStatusDto{}
	return &this
}

// NewMigrationStatusDtoWithDefaults instantiates a new MigrationStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationStatusDtoWithDefaults() *MigrationStatusDto {
	this := MigrationStatusDto{}
	return &this
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *MigrationStatusDto) GetProgress() float64 {
	if o == nil || IsNil(o.Progress) {
		var ret float64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusDto) GetProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *MigrationStatusDto) IsProgressSet() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float64 and assigns it to the Progress field.
func (o *MigrationStatusDto) SetProgress(v float64) {
	o.Progress = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationStatusDto) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationStatusDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *MigrationStatusDto) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *MigrationStatusDto) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *MigrationStatusDto) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *MigrationStatusDto) UnsetError() {
	o.Error.Unset()
}

// GetParseResult returns the ParseResult field value if set, zero value otherwise.
func (o *MigrationStatusDto) GetParseResult() MigrationApiInfo {
	if o == nil || IsNil(o.ParseResult) {
		var ret MigrationApiInfo
		return ret
	}
	return *o.ParseResult
}

// GetParseResultOk returns a tuple with the ParseResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusDto) GetParseResultOk() (*MigrationApiInfo, bool) {
	if o == nil || IsNil(o.ParseResult) {
		return nil, false
	}
	return o.ParseResult, true
}

// HasParseResult returns a boolean if a field has been set.
func (o *MigrationStatusDto) IsParseResultSet() bool {
	if o != nil && !IsNil(o.ParseResult) {
		return true
	}

	return false
}

// SetParseResult gets a reference to the given MigrationApiInfo and assigns it to the ParseResult field.
func (o *MigrationStatusDto) SetParseResult(v MigrationApiInfo) {
	o.ParseResult = &v
}

// GetIsCompleted returns the IsCompleted field value if set, zero value otherwise.
func (o *MigrationStatusDto) GetIsCompleted() bool {
	if o == nil || IsNil(o.IsCompleted) {
		var ret bool
		return ret
	}
	return *o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusDto) GetIsCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompleted) {
		return nil, false
	}
	return o.IsCompleted, true
}

// HasIsCompleted returns a boolean if a field has been set.
func (o *MigrationStatusDto) IsIsCompletedSet() bool {
	if o != nil && !IsNil(o.IsCompleted) {
		return true
	}

	return false
}

// SetIsCompleted gets a reference to the given bool and assigns it to the IsCompleted field.
func (o *MigrationStatusDto) SetIsCompleted(v bool) {
	o.IsCompleted = &v
}

func (o MigrationStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrationStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if !IsNil(o.ParseResult) {
		toSerialize["parseResult"] = o.ParseResult
	}
	if !IsNil(o.IsCompleted) {
		toSerialize["isCompleted"] = o.IsCompleted
	}
	return toSerialize, nil
}

type NullableMigrationStatusDto struct {
	value *MigrationStatusDto
	isSet bool
}

func (v NullableMigrationStatusDto) Get() *MigrationStatusDto {
	return v.value
}

func (v *NullableMigrationStatusDto) Set(val *MigrationStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationStatusDto(val *MigrationStatusDto) *NullableMigrationStatusDto {
	return &NullableMigrationStatusDto{value: val, isSet: true}
}

func (v NullableMigrationStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

