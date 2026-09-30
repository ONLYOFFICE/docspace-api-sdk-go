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

// checks if the ExternalDbSyncFormResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalDbSyncFormResultDto{}

// ExternalDbSyncFormResultDto What happened to one original form while the room was being exported to the external database.
type ExternalDbSyncFormResultDto struct {
	// The file of the original form whose collected data was exported. It is the form itself, not one of the filled  copies, so the same id can be read with the file operations of the portal.
	Id *int32 `json:"id,omitempty"`
	// The name of that form file at the moment of the export. It is empty when the form file no longer exists, which  is also the case in which the export of that entry fails.
	Title NullableString `json:"title,omitempty"`
	// Whether the data of this form reached the external database. One rejected form does not stop the others, so a  finished job can hold both successful and failed entries.
	Success *bool `json:"success,omitempty"`
	// Why this form was not exported. It is empty for a successful entry, and for a failed one it carries either the  message of the underlying failure or the generic export error of the portal.
	Error NullableString `json:"error,omitempty"`
}

// NewExternalDbSyncFormResultDto instantiates a new ExternalDbSyncFormResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalDbSyncFormResultDto() *ExternalDbSyncFormResultDto {
	this := ExternalDbSyncFormResultDto{}
	return &this
}

// NewExternalDbSyncFormResultDtoWithDefaults instantiates a new ExternalDbSyncFormResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalDbSyncFormResultDtoWithDefaults() *ExternalDbSyncFormResultDto {
	this := ExternalDbSyncFormResultDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExternalDbSyncFormResultDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalDbSyncFormResultDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExternalDbSyncFormResultDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ExternalDbSyncFormResultDto) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalDbSyncFormResultDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDbSyncFormResultDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ExternalDbSyncFormResultDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ExternalDbSyncFormResultDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ExternalDbSyncFormResultDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ExternalDbSyncFormResultDto) UnsetTitle() {
	o.Title.Unset()
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ExternalDbSyncFormResultDto) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalDbSyncFormResultDto) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ExternalDbSyncFormResultDto) IsSuccessSet() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ExternalDbSyncFormResultDto) SetSuccess(v bool) {
	o.Success = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalDbSyncFormResultDto) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDbSyncFormResultDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ExternalDbSyncFormResultDto) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ExternalDbSyncFormResultDto) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ExternalDbSyncFormResultDto) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ExternalDbSyncFormResultDto) UnsetError() {
	o.Error.Unset()
}

func (o ExternalDbSyncFormResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalDbSyncFormResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	return toSerialize, nil
}

type NullableExternalDbSyncFormResultDto struct {
	value *ExternalDbSyncFormResultDto
	isSet bool
}

func (v NullableExternalDbSyncFormResultDto) Get() *ExternalDbSyncFormResultDto {
	return v.value
}

func (v *NullableExternalDbSyncFormResultDto) Set(val *ExternalDbSyncFormResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalDbSyncFormResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalDbSyncFormResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalDbSyncFormResultDto(val *ExternalDbSyncFormResultDto) *NullableExternalDbSyncFormResultDto {
	return &NullableExternalDbSyncFormResultDto{value: val, isSet: true}
}

func (v NullableExternalDbSyncFormResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalDbSyncFormResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

