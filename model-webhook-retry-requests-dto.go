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

// checks if the WebhookRetryRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookRetryRequestsDto{}

// WebhookRetryRequestsDto Which past webhook deliveries are sent again.
type WebhookRetryRequestsDto struct {
	// The delivery records to send again, by the identifiers `GET api/2.0/settings/webhooks/log` reports. An  identifier that exists nowhere, and one belonging to another member subscription when the caller is not a  DocSpace administrator, is skipped in silence rather than failing the call, so compare the number of records  that come back against the number sent. An empty list is accepted and queues nothing.
	Ids []int32 `json:"ids,omitempty"`
}

// NewWebhookRetryRequestsDto instantiates a new WebhookRetryRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookRetryRequestsDto() *WebhookRetryRequestsDto {
	this := WebhookRetryRequestsDto{}
	return &this
}

// NewWebhookRetryRequestsDtoWithDefaults instantiates a new WebhookRetryRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookRetryRequestsDtoWithDefaults() *WebhookRetryRequestsDto {
	this := WebhookRetryRequestsDto{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookRetryRequestsDto) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookRetryRequestsDto) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *WebhookRetryRequestsDto) IsIdsSet() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *WebhookRetryRequestsDto) SetIds(v []int32) {
	o.Ids = v
}

func (o WebhookRetryRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookRetryRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableWebhookRetryRequestsDto struct {
	value *WebhookRetryRequestsDto
	isSet bool
}

func (v NullableWebhookRetryRequestsDto) Get() *WebhookRetryRequestsDto {
	return v.value
}

func (v *NullableWebhookRetryRequestsDto) Set(val *WebhookRetryRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookRetryRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookRetryRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookRetryRequestsDto(val *WebhookRetryRequestsDto) *NullableWebhookRetryRequestsDto {
	return &NullableWebhookRetryRequestsDto{value: val, isSet: true}
}

func (v NullableWebhookRetryRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookRetryRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

