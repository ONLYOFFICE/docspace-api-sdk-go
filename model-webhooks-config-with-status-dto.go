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

// checks if the WebhooksConfigWithStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksConfigWithStatusDto{}

// WebhooksConfigWithStatusDto A webhook subscription together with how its last delivery ended.
type WebhooksConfigWithStatusDto struct {
	// The subscription itself. Despite the plural name it is one subscription, not a list.
	Configs *WebhooksConfigDto `json:"configs,omitempty"`
	// The HTTP status code the target answered on the last attempt. `0` means nothing has been delivered yet,  which is not the same as a failure.
	Status *int32 `json:"status,omitempty"`
}

// NewWebhooksConfigWithStatusDto instantiates a new WebhooksConfigWithStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksConfigWithStatusDto() *WebhooksConfigWithStatusDto {
	this := WebhooksConfigWithStatusDto{}
	return &this
}

// NewWebhooksConfigWithStatusDtoWithDefaults instantiates a new WebhooksConfigWithStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksConfigWithStatusDtoWithDefaults() *WebhooksConfigWithStatusDto {
	this := WebhooksConfigWithStatusDto{}
	return &this
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *WebhooksConfigWithStatusDto) GetConfigs() WebhooksConfigDto {
	if o == nil || IsNil(o.Configs) {
		var ret WebhooksConfigDto
		return ret
	}
	return *o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfigWithStatusDto) GetConfigsOk() (*WebhooksConfigDto, bool) {
	if o == nil || IsNil(o.Configs) {
		return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *WebhooksConfigWithStatusDto) IsConfigsSet() bool {
	if o != nil && !IsNil(o.Configs) {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given WebhooksConfigDto and assigns it to the Configs field.
func (o *WebhooksConfigWithStatusDto) SetConfigs(v WebhooksConfigDto) {
	o.Configs = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebhooksConfigWithStatusDto) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfigWithStatusDto) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhooksConfigWithStatusDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *WebhooksConfigWithStatusDto) SetStatus(v int32) {
	o.Status = &v
}

func (o WebhooksConfigWithStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksConfigWithStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configs) {
		toSerialize["configs"] = o.Configs
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableWebhooksConfigWithStatusDto struct {
	value *WebhooksConfigWithStatusDto
	isSet bool
}

func (v NullableWebhooksConfigWithStatusDto) Get() *WebhooksConfigWithStatusDto {
	return v.value
}

func (v *NullableWebhooksConfigWithStatusDto) Set(val *WebhooksConfigWithStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksConfigWithStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksConfigWithStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksConfigWithStatusDto(val *WebhooksConfigWithStatusDto) *NullableWebhooksConfigWithStatusDto {
	return &NullableWebhooksConfigWithStatusDto{value: val, isSet: true}
}

func (v NullableWebhooksConfigWithStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksConfigWithStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

