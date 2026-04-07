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

// checks if the UploadResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadResultDto{}

// UploadResultDto The upload result parameters.
type UploadResultDto struct {
	// Specifies if the upload operation is successful or not.
	Success *bool `json:"success,omitempty"`
	// The uploaded data.
	Data interface{} `json:"data,omitempty"`
	// The message sent after the successful upload operation.
	Message NullableString `json:"message,omitempty"`
}

// NewUploadResultDto instantiates a new UploadResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadResultDto() *UploadResultDto {
	this := UploadResultDto{}
	return &this
}

// NewUploadResultDtoWithDefaults instantiates a new UploadResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadResultDtoWithDefaults() *UploadResultDto {
	this := UploadResultDto{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UploadResultDto) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadResultDto) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UploadResultDto) IsSuccessSet() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UploadResultDto) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadResultDto) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadResultDto) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UploadResultDto) IsDataSet() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *UploadResultDto) SetData(v interface{}) {
	o.Data = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadResultDto) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadResultDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *UploadResultDto) IsMessageSet() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *UploadResultDto) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *UploadResultDto) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *UploadResultDto) UnsetMessage() {
	o.Message.Unset()
}

func (o UploadResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableUploadResultDto struct {
	value *UploadResultDto
	isSet bool
}

func (v NullableUploadResultDto) Get() *UploadResultDto {
	return v.value
}

func (v *NullableUploadResultDto) Set(val *UploadResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadResultDto(val *UploadResultDto) *NullableUploadResultDto {
	return &NullableUploadResultDto{value: val, isSet: true}
}

func (v NullableUploadResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

