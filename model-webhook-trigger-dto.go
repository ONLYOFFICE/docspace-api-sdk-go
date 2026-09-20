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

// checks if the WebhookTriggerDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookTriggerDto{}

// WebhookTriggerDto One event a webhook can listen to, with the bit that selects it and whether the caller may subscribe to it.
type WebhookTriggerDto struct {
	// The event name exactly as it appears in a delivered payload, so a receiver can match on it. The entry  named `*` is not an event but the catch-all.
	Name NullableString `json:"name,omitempty"`
	// The bit that stands for this event in the `triggers` bitmask of a subscription. Add the bits of the wanted  events together; the catch-all entry has the value `0` and is used on its own rather than added to  anything.
	Id *int64 `json:"id,omitempty"`
	// Whether the caller's own role may subscribe to this event - a plain member cannot subscribe to user, group  or room creation, where a room administrator can. An unavailable event is listed all the same, and sending  its bit to `POST api/2.0/settings/webhook` is refused as an invalid request.
	Available *bool `json:"available,omitempty"`
}

// NewWebhookTriggerDto instantiates a new WebhookTriggerDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookTriggerDto() *WebhookTriggerDto {
	this := WebhookTriggerDto{}
	return &this
}

// NewWebhookTriggerDtoWithDefaults instantiates a new WebhookTriggerDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookTriggerDtoWithDefaults() *WebhookTriggerDto {
	this := WebhookTriggerDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookTriggerDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookTriggerDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WebhookTriggerDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WebhookTriggerDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WebhookTriggerDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WebhookTriggerDto) UnsetName() {
	o.Name.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookTriggerDto) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookTriggerDto) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookTriggerDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *WebhookTriggerDto) SetId(v int64) {
	o.Id = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *WebhookTriggerDto) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookTriggerDto) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *WebhookTriggerDto) IsAvailableSet() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *WebhookTriggerDto) SetAvailable(v bool) {
	o.Available = &v
}

func (o WebhookTriggerDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookTriggerDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	return toSerialize, nil
}

type NullableWebhookTriggerDto struct {
	value *WebhookTriggerDto
	isSet bool
}

func (v NullableWebhookTriggerDto) Get() *WebhookTriggerDto {
	return v.value
}

func (v *NullableWebhookTriggerDto) Set(val *WebhookTriggerDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookTriggerDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookTriggerDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookTriggerDto(val *WebhookTriggerDto) *NullableWebhookTriggerDto {
	return &NullableWebhookTriggerDto{value: val, isSet: true}
}

func (v NullableWebhookTriggerDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookTriggerDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

