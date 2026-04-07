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

// checks if the ChunkedUploadSessionResponseWrapperInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChunkedUploadSessionResponseWrapperInteger{}

// ChunkedUploadSessionResponseWrapperInteger Represents a wrapper for the response of a chunked upload session operation.
type ChunkedUploadSessionResponseWrapperInteger struct {
	// Gets or sets a value indicating whether the operation was successful.
	Success *bool `json:"success,omitempty"`
	Data *ChunkedUploadSessionResponseInteger `json:"data,omitempty"`
}

// NewChunkedUploadSessionResponseWrapperInteger instantiates a new ChunkedUploadSessionResponseWrapperInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChunkedUploadSessionResponseWrapperInteger() *ChunkedUploadSessionResponseWrapperInteger {
	this := ChunkedUploadSessionResponseWrapperInteger{}
	return &this
}

// NewChunkedUploadSessionResponseWrapperIntegerWithDefaults instantiates a new ChunkedUploadSessionResponseWrapperInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChunkedUploadSessionResponseWrapperIntegerWithDefaults() *ChunkedUploadSessionResponseWrapperInteger {
	this := ChunkedUploadSessionResponseWrapperInteger{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ChunkedUploadSessionResponseWrapperInteger) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadSessionResponseWrapperInteger) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ChunkedUploadSessionResponseWrapperInteger) IsSuccessSet() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ChunkedUploadSessionResponseWrapperInteger) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ChunkedUploadSessionResponseWrapperInteger) GetData() ChunkedUploadSessionResponseInteger {
	if o == nil || IsNil(o.Data) {
		var ret ChunkedUploadSessionResponseInteger
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadSessionResponseWrapperInteger) GetDataOk() (*ChunkedUploadSessionResponseInteger, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ChunkedUploadSessionResponseWrapperInteger) IsDataSet() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ChunkedUploadSessionResponseInteger and assigns it to the Data field.
func (o *ChunkedUploadSessionResponseWrapperInteger) SetData(v ChunkedUploadSessionResponseInteger) {
	o.Data = &v
}

func (o ChunkedUploadSessionResponseWrapperInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChunkedUploadSessionResponseWrapperInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableChunkedUploadSessionResponseWrapperInteger struct {
	value *ChunkedUploadSessionResponseWrapperInteger
	isSet bool
}

func (v NullableChunkedUploadSessionResponseWrapperInteger) Get() *ChunkedUploadSessionResponseWrapperInteger {
	return v.value
}

func (v *NullableChunkedUploadSessionResponseWrapperInteger) Set(val *ChunkedUploadSessionResponseWrapperInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableChunkedUploadSessionResponseWrapperInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableChunkedUploadSessionResponseWrapperInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChunkedUploadSessionResponseWrapperInteger(val *ChunkedUploadSessionResponseWrapperInteger) *NullableChunkedUploadSessionResponseWrapperInteger {
	return &NullableChunkedUploadSessionResponseWrapperInteger{value: val, isSet: true}
}

func (v NullableChunkedUploadSessionResponseWrapperInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChunkedUploadSessionResponseWrapperInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

