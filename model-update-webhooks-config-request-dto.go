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
	"bytes"
	"fmt"
)

// checks if the UpdateWebhooksConfigRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebhooksConfigRequestDto{}

// UpdateWebhooksConfigRequestDto The webhook subscription being changed, with the parameters it is to have afterwards.
type UpdateWebhooksConfigRequestDto struct {
	// The label the subscription is listed under. It is for the administrator reading the list and is never sent to  the target; it does not have to be unique.
	Name string `json:"name"`
	// The address the portal posts the event payload to. It has to be an absolute `http` or `https` address outside  the installation own network, and it is probed before anything is stored: it must answer a HEAD request with  a success code, and a redirect does not count as one.
	Uri string `json:"uri"`
	// The shared secret the payload signature is computed with, so the receiver can tell a genuine call from a  forged one. It has to satisfy the portal password rules published by  `GET api/2.0/settings/security/password`, and it is never echoed back by any operation. On an update an empty  value keeps the secret already stored.
	SecretKey *string `json:"secretKey,omitempty"`
	// Whether the subscription delivers at all. While it is off the matching events are dropped rather than queued,  so nothing from that period arrives once it is switched on again.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether the target certificate is verified. `true` demands an `https` address with a valid certificate; `false`  delivers without checking the certificate, and is what a plain `http` target needs.
	Ssl *bool `json:"ssl,omitempty"`
	// The events the subscription listens for, as a bitmask adding together the `id` values of the entries of `GET  api/2.0/settings/webhook/triggers`. 0 is the catch-all, not none: it stands for every event the  subscription creator's role may receive, not for every event of the portal.
	Triggers *WebhookTrigger `json:"triggers,omitempty"`
	// The single entity the subscription is narrowed to, by its identifier in the API that owns it. For user events  it is the GUID of a user (`id` of `GET api/2.0/people`), for group events the GUID of a group (`id` of `GET  api/2.0/group`), and either must be one the subscription's creator may see; for file, folder, room and form  events it is the entry ID from the Files API (`id` of `GET api/2.0/files/rooms` or of a folder listing), a  number for the portal's own storage and a string for third-party storage. Leaving it out delivers events  about every entity the subscribed triggers cover.
	TargetId *string `json:"targetId,omitempty"`
	// The subscription to act on, by the `id` that `GET api/2.0/settings/webhook` reports. It travels in the body  rather than in the path, and an id that exists in no portal subscription answers 404.
	Id int32 `json:"id"`
}

type _UpdateWebhooksConfigRequestDto UpdateWebhooksConfigRequestDto

// NewUpdateWebhooksConfigRequestDto instantiates a new UpdateWebhooksConfigRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebhooksConfigRequestDto(name string, uri string, id int32) *UpdateWebhooksConfigRequestDto {
	this := UpdateWebhooksConfigRequestDto{}
	this.Name = name
	this.Uri = uri
	this.Id = id
	return &this
}

// NewUpdateWebhooksConfigRequestDtoWithDefaults instantiates a new UpdateWebhooksConfigRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebhooksConfigRequestDtoWithDefaults() *UpdateWebhooksConfigRequestDto {
	this := UpdateWebhooksConfigRequestDto{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateWebhooksConfigRequestDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateWebhooksConfigRequestDto) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
func (o *UpdateWebhooksConfigRequestDto) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestDto) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UpdateWebhooksConfigRequestDto) SetUri(v string) {
	o.Uri = v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestDto) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestDto) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestDto) IsSecretKeySet() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *UpdateWebhooksConfigRequestDto) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateWebhooksConfigRequestDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestDto) GetSsl() bool {
	if o == nil || IsNil(o.Ssl) {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestDto) GetSslOk() (*bool, bool) {
	if o == nil || IsNil(o.Ssl) {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestDto) IsSslSet() bool {
	if o != nil && !IsNil(o.Ssl) {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *UpdateWebhooksConfigRequestDto) SetSsl(v bool) {
	o.Ssl = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestDto) GetTriggers() WebhookTrigger {
	if o == nil || IsNil(o.Triggers) {
		var ret WebhookTrigger
		return ret
	}
	return *o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestDto) GetTriggersOk() (*WebhookTrigger, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestDto) IsTriggersSet() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given WebhookTrigger and assigns it to the Triggers field.
func (o *UpdateWebhooksConfigRequestDto) SetTriggers(v WebhookTrigger) {
	o.Triggers = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestDto) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestDto) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestDto) IsTargetIdSet() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *UpdateWebhooksConfigRequestDto) SetTargetId(v string) {
	o.TargetId = &v
}

// GetId returns the Id field value
func (o *UpdateWebhooksConfigRequestDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateWebhooksConfigRequestDto) SetId(v int32) {
	o.Id = v
}

func (o UpdateWebhooksConfigRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebhooksConfigRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["uri"] = o.Uri
	if !IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Ssl) {
		toSerialize["ssl"] = o.Ssl
	}
	if !IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UpdateWebhooksConfigRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"uri",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateWebhooksConfigRequestDto := _UpdateWebhooksConfigRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varUpdateWebhooksConfigRequestDto)

	if err != nil {
		return err
	}

	*o = UpdateWebhooksConfigRequestDto(varUpdateWebhooksConfigRequestDto)

	return err
}

type NullableUpdateWebhooksConfigRequestDto struct {
	value *UpdateWebhooksConfigRequestDto
	isSet bool
}

func (v NullableUpdateWebhooksConfigRequestDto) Get() *UpdateWebhooksConfigRequestDto {
	return v.value
}

func (v *NullableUpdateWebhooksConfigRequestDto) Set(val *UpdateWebhooksConfigRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebhooksConfigRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebhooksConfigRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebhooksConfigRequestDto(val *UpdateWebhooksConfigRequestDto) *NullableUpdateWebhooksConfigRequestDto {
	return &NullableUpdateWebhooksConfigRequestDto{value: val, isSet: true}
}

func (v NullableUpdateWebhooksConfigRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebhooksConfigRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

