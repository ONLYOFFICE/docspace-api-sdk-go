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

// checks if the ThirdPartyChunkedUploadSessionResponseWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyChunkedUploadSessionResponseWrapper{}

// ThirdPartyChunkedUploadSessionResponseWrapper The reserved chunked upload wrapped in the envelope the two older session operations answer with.
type ThirdPartyChunkedUploadSessionResponseWrapper struct {
	// Always true in a body that reaches the caller, because a call that does not succeed answers with an error  status and no body at all. It cannot be used to tell a refusal from a success.
	Success *bool `json:"success,omitempty"`
	// The reserved upload itself, in the same shape the newer session operations answer with directly.
	Data *ThirdPartyChunkedUploadSessionResponse `json:"data,omitempty"`
}

// NewThirdPartyChunkedUploadSessionResponseWrapper instantiates a new ThirdPartyChunkedUploadSessionResponseWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyChunkedUploadSessionResponseWrapper() *ThirdPartyChunkedUploadSessionResponseWrapper {
	this := ThirdPartyChunkedUploadSessionResponseWrapper{}
	return &this
}

// NewThirdPartyChunkedUploadSessionResponseWrapperWithDefaults instantiates a new ThirdPartyChunkedUploadSessionResponseWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyChunkedUploadSessionResponseWrapperWithDefaults() *ThirdPartyChunkedUploadSessionResponseWrapper {
	this := ThirdPartyChunkedUploadSessionResponseWrapper{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ThirdPartyChunkedUploadSessionResponseWrapper) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyChunkedUploadSessionResponseWrapper) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ThirdPartyChunkedUploadSessionResponseWrapper) IsSuccessSet() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ThirdPartyChunkedUploadSessionResponseWrapper) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ThirdPartyChunkedUploadSessionResponseWrapper) GetData() ThirdPartyChunkedUploadSessionResponse {
	if o == nil || IsNil(o.Data) {
		var ret ThirdPartyChunkedUploadSessionResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyChunkedUploadSessionResponseWrapper) GetDataOk() (*ThirdPartyChunkedUploadSessionResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ThirdPartyChunkedUploadSessionResponseWrapper) IsDataSet() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ThirdPartyChunkedUploadSessionResponse and assigns it to the Data field.
func (o *ThirdPartyChunkedUploadSessionResponseWrapper) SetData(v ThirdPartyChunkedUploadSessionResponse) {
	o.Data = &v
}

func (o ThirdPartyChunkedUploadSessionResponseWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyChunkedUploadSessionResponseWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableThirdPartyChunkedUploadSessionResponseWrapper struct {
	value *ThirdPartyChunkedUploadSessionResponseWrapper
	isSet bool
}

func (v NullableThirdPartyChunkedUploadSessionResponseWrapper) Get() *ThirdPartyChunkedUploadSessionResponseWrapper {
	return v.value
}

func (v *NullableThirdPartyChunkedUploadSessionResponseWrapper) Set(val *ThirdPartyChunkedUploadSessionResponseWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyChunkedUploadSessionResponseWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyChunkedUploadSessionResponseWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyChunkedUploadSessionResponseWrapper(val *ThirdPartyChunkedUploadSessionResponseWrapper) *NullableThirdPartyChunkedUploadSessionResponseWrapper {
	return &NullableThirdPartyChunkedUploadSessionResponseWrapper{value: val, isSet: true}
}

func (v NullableThirdPartyChunkedUploadSessionResponseWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyChunkedUploadSessionResponseWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

