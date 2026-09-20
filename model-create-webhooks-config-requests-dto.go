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

// checks if the CreateWebhooksConfigRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebhooksConfigRequestsDto{}

// CreateWebhooksConfigRequestsDto The target a webhook subscription calls, the events it listens for, and the secret it signs with.
type CreateWebhooksConfigRequestsDto struct {
	// The label the subscription is listed under. It is for the administrator reading the list and is never sent to  the target; it does not have to be unique.
	Name string `json:"name"`
	// The address the portal posts the event payload to. It has to be an absolute `http` or `https` address outside  the installation own network, and it is probed before anything is stored: it must answer a HEAD request with  a success code, and a redirect does not count as one.
	Uri string `json:"uri"`
	// The shared secret the payload signature is computed with, so the receiver can tell a genuine call from a  forged one. It has to satisfy the portal password rules published by  `GET api/2.0/settings/security/password`, and it is never echoed back by any operation. On an update an empty  value keeps the secret already stored.
	SecretKey NullableString `json:"secretKey,omitempty"`
	// Whether the subscription delivers at all. While it is off the matching events are dropped rather than queued,  so nothing from that period arrives once it is switched on again.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether the target certificate is verified. Setting it demands an `https` target with a valid certificate;  leaving it off delivers without checking the certificate at all.
	Ssl *bool `json:"ssl,omitempty"`
	// The events the subscription listens for, as a bitmask combining the flags; 0 subscribes to all of them. Take  the flags the caller role is allowed to use from `GET api/2.0/settings/webhook/triggers`, since a flag beyond  that set is refused with 400. A subscription still only fires for events its creator may see.
	Triggers *WebhookTrigger `json:"triggers,omitempty"`
	// The single entity the subscription is narrowed to, by its identifier - a room or a file, for instance.  Leaving it out delivers events about every entity the subscribed triggers cover.
	TargetId NullableString `json:"targetId,omitempty"`
}

type _CreateWebhooksConfigRequestsDto CreateWebhooksConfigRequestsDto

// NewCreateWebhooksConfigRequestsDto instantiates a new CreateWebhooksConfigRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhooksConfigRequestsDto(name string, uri string) *CreateWebhooksConfigRequestsDto {
	this := CreateWebhooksConfigRequestsDto{}
	this.Name = name
	this.Uri = uri
	return &this
}

// NewCreateWebhooksConfigRequestsDtoWithDefaults instantiates a new CreateWebhooksConfigRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhooksConfigRequestsDtoWithDefaults() *CreateWebhooksConfigRequestsDto {
	this := CreateWebhooksConfigRequestsDto{}
	return &this
}

// GetName returns the Name field value
func (o *CreateWebhooksConfigRequestsDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWebhooksConfigRequestsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWebhooksConfigRequestsDto) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
func (o *CreateWebhooksConfigRequestsDto) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *CreateWebhooksConfigRequestsDto) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *CreateWebhooksConfigRequestsDto) SetUri(v string) {
	o.Uri = v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWebhooksConfigRequestsDto) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey.Get()) {
		var ret string
		return ret
	}
	return *o.SecretKey.Get()
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWebhooksConfigRequestsDto) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretKey.Get(), o.SecretKey.IsSet()
}

// HasSecretKey returns a boolean if a field has been set.
func (o *CreateWebhooksConfigRequestsDto) IsSecretKeySet() bool {
	if o != nil && o.SecretKey.IsSet() {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given NullableString and assigns it to the SecretKey field.
func (o *CreateWebhooksConfigRequestsDto) SetSecretKey(v string) {
	o.SecretKey.Set(&v)
}
// SetSecretKeyNil sets the value for SecretKey to be an explicit nil
func (o *CreateWebhooksConfigRequestsDto) SetSecretKeyNil() {
	o.SecretKey.Set(nil)
}

// UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
func (o *CreateWebhooksConfigRequestsDto) UnsetSecretKey() {
	o.SecretKey.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateWebhooksConfigRequestsDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhooksConfigRequestsDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateWebhooksConfigRequestsDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateWebhooksConfigRequestsDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *CreateWebhooksConfigRequestsDto) GetSsl() bool {
	if o == nil || IsNil(o.Ssl) {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhooksConfigRequestsDto) GetSslOk() (*bool, bool) {
	if o == nil || IsNil(o.Ssl) {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *CreateWebhooksConfigRequestsDto) IsSslSet() bool {
	if o != nil && !IsNil(o.Ssl) {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *CreateWebhooksConfigRequestsDto) SetSsl(v bool) {
	o.Ssl = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *CreateWebhooksConfigRequestsDto) GetTriggers() WebhookTrigger {
	if o == nil || IsNil(o.Triggers) {
		var ret WebhookTrigger
		return ret
	}
	return *o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhooksConfigRequestsDto) GetTriggersOk() (*WebhookTrigger, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *CreateWebhooksConfigRequestsDto) IsTriggersSet() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given WebhookTrigger and assigns it to the Triggers field.
func (o *CreateWebhooksConfigRequestsDto) SetTriggers(v WebhookTrigger) {
	o.Triggers = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWebhooksConfigRequestsDto) GetTargetId() string {
	if o == nil || IsNil(o.TargetId.Get()) {
		var ret string
		return ret
	}
	return *o.TargetId.Get()
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWebhooksConfigRequestsDto) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetId.Get(), o.TargetId.IsSet()
}

// HasTargetId returns a boolean if a field has been set.
func (o *CreateWebhooksConfigRequestsDto) IsTargetIdSet() bool {
	if o != nil && o.TargetId.IsSet() {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given NullableString and assigns it to the TargetId field.
func (o *CreateWebhooksConfigRequestsDto) SetTargetId(v string) {
	o.TargetId.Set(&v)
}
// SetTargetIdNil sets the value for TargetId to be an explicit nil
func (o *CreateWebhooksConfigRequestsDto) SetTargetIdNil() {
	o.TargetId.Set(nil)
}

// UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
func (o *CreateWebhooksConfigRequestsDto) UnsetTargetId() {
	o.TargetId.Unset()
}

func (o CreateWebhooksConfigRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebhooksConfigRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["uri"] = o.Uri
	if o.SecretKey.IsSet() {
		toSerialize["secretKey"] = o.SecretKey.Get()
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
	if o.TargetId.IsSet() {
		toSerialize["targetId"] = o.TargetId.Get()
	}
	return toSerialize, nil
}

func (o *CreateWebhooksConfigRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"uri",
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

	varCreateWebhooksConfigRequestsDto := _CreateWebhooksConfigRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateWebhooksConfigRequestsDto)

	if err != nil {
		return err
	}

	*o = CreateWebhooksConfigRequestsDto(varCreateWebhooksConfigRequestsDto)

	return err
}

type NullableCreateWebhooksConfigRequestsDto struct {
	value *CreateWebhooksConfigRequestsDto
	isSet bool
}

func (v NullableCreateWebhooksConfigRequestsDto) Get() *CreateWebhooksConfigRequestsDto {
	return v.value
}

func (v *NullableCreateWebhooksConfigRequestsDto) Set(val *CreateWebhooksConfigRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhooksConfigRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhooksConfigRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhooksConfigRequestsDto(val *CreateWebhooksConfigRequestsDto) *NullableCreateWebhooksConfigRequestsDto {
	return &NullableCreateWebhooksConfigRequestsDto{value: val, isSet: true}
}

func (v NullableCreateWebhooksConfigRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhooksConfigRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

