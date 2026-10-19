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

// checks if the DocsCloudServerConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudServerConfigDto{}

// DocsCloudServerConfigDto Represents the server configuration of a Docs Connect tenant.
type DocsCloudServerConfigDto struct {
	// Whether anonymous access is supported.
	IsAnonymousSupport *bool `json:"isAnonymousSupport,omitempty"`
	// The maximum file size in bytes.
	FileSizeLimit *int64 `json:"fileSizeLimit,omitempty"`
}

// NewDocsCloudServerConfigDto instantiates a new DocsCloudServerConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudServerConfigDto() *DocsCloudServerConfigDto {
	this := DocsCloudServerConfigDto{}
	return &this
}

// NewDocsCloudServerConfigDtoWithDefaults instantiates a new DocsCloudServerConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudServerConfigDtoWithDefaults() *DocsCloudServerConfigDto {
	this := DocsCloudServerConfigDto{}
	return &this
}

// GetIsAnonymousSupport returns the IsAnonymousSupport field value if set, zero value otherwise.
func (o *DocsCloudServerConfigDto) GetIsAnonymousSupport() bool {
	if o == nil || IsNil(o.IsAnonymousSupport) {
		var ret bool
		return ret
	}
	return *o.IsAnonymousSupport
}

// GetIsAnonymousSupportOk returns a tuple with the IsAnonymousSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudServerConfigDto) GetIsAnonymousSupportOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnonymousSupport) {
		return nil, false
	}
	return o.IsAnonymousSupport, true
}

// HasIsAnonymousSupport returns a boolean if a field has been set.
func (o *DocsCloudServerConfigDto) IsIsAnonymousSupportSet() bool {
	if o != nil && !IsNil(o.IsAnonymousSupport) {
		return true
	}

	return false
}

// SetIsAnonymousSupport gets a reference to the given bool and assigns it to the IsAnonymousSupport field.
func (o *DocsCloudServerConfigDto) SetIsAnonymousSupport(v bool) {
	o.IsAnonymousSupport = &v
}

// GetFileSizeLimit returns the FileSizeLimit field value if set, zero value otherwise.
func (o *DocsCloudServerConfigDto) GetFileSizeLimit() int64 {
	if o == nil || IsNil(o.FileSizeLimit) {
		var ret int64
		return ret
	}
	return *o.FileSizeLimit
}

// GetFileSizeLimitOk returns a tuple with the FileSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudServerConfigDto) GetFileSizeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSizeLimit) {
		return nil, false
	}
	return o.FileSizeLimit, true
}

// HasFileSizeLimit returns a boolean if a field has been set.
func (o *DocsCloudServerConfigDto) IsFileSizeLimitSet() bool {
	if o != nil && !IsNil(o.FileSizeLimit) {
		return true
	}

	return false
}

// SetFileSizeLimit gets a reference to the given int64 and assigns it to the FileSizeLimit field.
func (o *DocsCloudServerConfigDto) SetFileSizeLimit(v int64) {
	o.FileSizeLimit = &v
}

func (o DocsCloudServerConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudServerConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAnonymousSupport) {
		toSerialize["isAnonymousSupport"] = o.IsAnonymousSupport
	}
	if !IsNil(o.FileSizeLimit) {
		toSerialize["fileSizeLimit"] = o.FileSizeLimit
	}
	return toSerialize, nil
}

type NullableDocsCloudServerConfigDto struct {
	value *DocsCloudServerConfigDto
	isSet bool
}

func (v NullableDocsCloudServerConfigDto) Get() *DocsCloudServerConfigDto {
	return v.value
}

func (v *NullableDocsCloudServerConfigDto) Set(val *DocsCloudServerConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudServerConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudServerConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudServerConfigDto(val *DocsCloudServerConfigDto) *NullableDocsCloudServerConfigDto {
	return &NullableDocsCloudServerConfigDto{value: val, isSet: true}
}

func (v NullableDocsCloudServerConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudServerConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

