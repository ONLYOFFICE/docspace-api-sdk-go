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

// checks if the DocsCloudDevPackRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudDevPackRequestDto{}

// DocsCloudDevPackRequestDto The request parameters for switching the Docs Connect subscription to Docs Connect Dev Pack, or for calculating  the cost of that switch.
type DocsCloudDevPackRequestDto struct {
	// The number of users to subscribe to Docs Connect Dev Pack for. It must be at least the number of users of  the currently purchased Docs Connect subscription, and at least the Docs Connect Dev Pack minimum configured  for the installation, which is 10 users by default; a smaller value is rejected with 400.
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewDocsCloudDevPackRequestDto instantiates a new DocsCloudDevPackRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudDevPackRequestDto() *DocsCloudDevPackRequestDto {
	this := DocsCloudDevPackRequestDto{}
	return &this
}

// NewDocsCloudDevPackRequestDtoWithDefaults instantiates a new DocsCloudDevPackRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudDevPackRequestDtoWithDefaults() *DocsCloudDevPackRequestDto {
	this := DocsCloudDevPackRequestDto{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *DocsCloudDevPackRequestDto) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudDevPackRequestDto) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *DocsCloudDevPackRequestDto) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *DocsCloudDevPackRequestDto) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o DocsCloudDevPackRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudDevPackRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableDocsCloudDevPackRequestDto struct {
	value *DocsCloudDevPackRequestDto
	isSet bool
}

func (v NullableDocsCloudDevPackRequestDto) Get() *DocsCloudDevPackRequestDto {
	return v.value
}

func (v *NullableDocsCloudDevPackRequestDto) Set(val *DocsCloudDevPackRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudDevPackRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudDevPackRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudDevPackRequestDto(val *DocsCloudDevPackRequestDto) *NullableDocsCloudDevPackRequestDto {
	return &NullableDocsCloudDevPackRequestDto{value: val, isSet: true}
}

func (v NullableDocsCloudDevPackRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudDevPackRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

