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

// checks if the XlsxReportResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XlsxReportResponseDto{}

// XlsxReportResponseDto The XLSX report task response parameters.
type XlsxReportResponseDto struct {
	// The original form file information.
	Form *FileDtoInteger `json:"form,omitempty"`
	// The Document Builder task information.
	Task *DocumentBuilderTaskDto `json:"task,omitempty"`
	// Specifies whether the XLSX report file is newly created or an existing file will be updated.
	IsNewFile *bool `json:"isNewFile,omitempty"`
}

// NewXlsxReportResponseDto instantiates a new XlsxReportResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXlsxReportResponseDto() *XlsxReportResponseDto {
	this := XlsxReportResponseDto{}
	return &this
}

// NewXlsxReportResponseDtoWithDefaults instantiates a new XlsxReportResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXlsxReportResponseDtoWithDefaults() *XlsxReportResponseDto {
	this := XlsxReportResponseDto{}
	return &this
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *XlsxReportResponseDto) GetForm() FileDtoInteger {
	if o == nil || IsNil(o.Form) {
		var ret FileDtoInteger
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XlsxReportResponseDto) GetFormOk() (*FileDtoInteger, bool) {
	if o == nil || IsNil(o.Form) {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *XlsxReportResponseDto) IsFormSet() bool {
	if o != nil && !IsNil(o.Form) {
		return true
	}

	return false
}

// SetForm gets a reference to the given FileDtoInteger and assigns it to the Form field.
func (o *XlsxReportResponseDto) SetForm(v FileDtoInteger) {
	o.Form = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *XlsxReportResponseDto) GetTask() DocumentBuilderTaskDto {
	if o == nil || IsNil(o.Task) {
		var ret DocumentBuilderTaskDto
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XlsxReportResponseDto) GetTaskOk() (*DocumentBuilderTaskDto, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *XlsxReportResponseDto) IsTaskSet() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given DocumentBuilderTaskDto and assigns it to the Task field.
func (o *XlsxReportResponseDto) SetTask(v DocumentBuilderTaskDto) {
	o.Task = &v
}

// GetIsNewFile returns the IsNewFile field value if set, zero value otherwise.
func (o *XlsxReportResponseDto) GetIsNewFile() bool {
	if o == nil || IsNil(o.IsNewFile) {
		var ret bool
		return ret
	}
	return *o.IsNewFile
}

// GetIsNewFileOk returns a tuple with the IsNewFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XlsxReportResponseDto) GetIsNewFileOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNewFile) {
		return nil, false
	}
	return o.IsNewFile, true
}

// HasIsNewFile returns a boolean if a field has been set.
func (o *XlsxReportResponseDto) IsIsNewFileSet() bool {
	if o != nil && !IsNil(o.IsNewFile) {
		return true
	}

	return false
}

// SetIsNewFile gets a reference to the given bool and assigns it to the IsNewFile field.
func (o *XlsxReportResponseDto) SetIsNewFile(v bool) {
	o.IsNewFile = &v
}

func (o XlsxReportResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XlsxReportResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Form) {
		toSerialize["form"] = o.Form
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	if !IsNil(o.IsNewFile) {
		toSerialize["isNewFile"] = o.IsNewFile
	}
	return toSerialize, nil
}

type NullableXlsxReportResponseDto struct {
	value *XlsxReportResponseDto
	isSet bool
}

func (v NullableXlsxReportResponseDto) Get() *XlsxReportResponseDto {
	return v.value
}

func (v *NullableXlsxReportResponseDto) Set(val *XlsxReportResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableXlsxReportResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableXlsxReportResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXlsxReportResponseDto(val *XlsxReportResponseDto) *NullableXlsxReportResponseDto {
	return &NullableXlsxReportResponseDto{value: val, isSet: true}
}

func (v NullableXlsxReportResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXlsxReportResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

