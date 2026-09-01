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

// checks if the UpdateWebhooksConfigRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebhooksConfigRequestsDto{}

// UpdateWebhooksConfigRequestsDto The request parameters for updating the webhook configuration.
type UpdateWebhooksConfigRequestsDto struct {
	// The human-readable name of the webhook configuration.
	Name string `json:"name"`
	// The destination URL where the webhook events will be sent.
	Uri string `json:"uri"`
	// The webhook secret key used to sign the webhook payloads for the security verification.
	SecretKey *string `json:"secretKey,omitempty"`
	// Specifies whether the webhook configuration is active or not.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies whether the SSL certificate verification is required or not.
	Ssl *bool `json:"ssl,omitempty"`
	// Defines which events will trigger webhook notifications.
	Triggers *WebhookTrigger `json:"triggers,omitempty"`
	// Target ID
	TargetId *string `json:"targetId,omitempty"`
	// The webhook configuration ID.
	Id int32 `json:"id"`
}

type _UpdateWebhooksConfigRequestsDto UpdateWebhooksConfigRequestsDto

// NewUpdateWebhooksConfigRequestsDto instantiates a new UpdateWebhooksConfigRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebhooksConfigRequestsDto(name string, uri string, id int32) *UpdateWebhooksConfigRequestsDto {
	this := UpdateWebhooksConfigRequestsDto{}
	this.Name = name
	this.Uri = uri
	this.Id = id
	return &this
}

// NewUpdateWebhooksConfigRequestsDtoWithDefaults instantiates a new UpdateWebhooksConfigRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebhooksConfigRequestsDtoWithDefaults() *UpdateWebhooksConfigRequestsDto {
	this := UpdateWebhooksConfigRequestsDto{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateWebhooksConfigRequestsDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateWebhooksConfigRequestsDto) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
func (o *UpdateWebhooksConfigRequestsDto) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestsDto) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UpdateWebhooksConfigRequestsDto) SetUri(v string) {
	o.Uri = v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestsDto) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestsDto) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestsDto) IsSecretKeySet() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *UpdateWebhooksConfigRequestsDto) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestsDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestsDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestsDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateWebhooksConfigRequestsDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestsDto) GetSsl() bool {
	if o == nil || IsNil(o.Ssl) {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestsDto) GetSslOk() (*bool, bool) {
	if o == nil || IsNil(o.Ssl) {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestsDto) IsSslSet() bool {
	if o != nil && !IsNil(o.Ssl) {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *UpdateWebhooksConfigRequestsDto) SetSsl(v bool) {
	o.Ssl = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestsDto) GetTriggers() WebhookTrigger {
	if o == nil || IsNil(o.Triggers) {
		var ret WebhookTrigger
		return ret
	}
	return *o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestsDto) GetTriggersOk() (*WebhookTrigger, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestsDto) IsTriggersSet() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given WebhookTrigger and assigns it to the Triggers field.
func (o *UpdateWebhooksConfigRequestsDto) SetTriggers(v WebhookTrigger) {
	o.Triggers = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *UpdateWebhooksConfigRequestsDto) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestsDto) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *UpdateWebhooksConfigRequestsDto) IsTargetIdSet() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *UpdateWebhooksConfigRequestsDto) SetTargetId(v string) {
	o.TargetId = &v
}

// GetId returns the Id field value
func (o *UpdateWebhooksConfigRequestsDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhooksConfigRequestsDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateWebhooksConfigRequestsDto) SetId(v int32) {
	o.Id = v
}

func (o UpdateWebhooksConfigRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebhooksConfigRequestsDto) ToMap() (map[string]interface{}, error) {
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

func (o *UpdateWebhooksConfigRequestsDto) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateWebhooksConfigRequestsDto := _UpdateWebhooksConfigRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varUpdateWebhooksConfigRequestsDto)

	if err != nil {
		return err
	}

	*o = UpdateWebhooksConfigRequestsDto(varUpdateWebhooksConfigRequestsDto)

	return err
}

type NullableUpdateWebhooksConfigRequestsDto struct {
	value *UpdateWebhooksConfigRequestsDto
	isSet bool
}

func (v NullableUpdateWebhooksConfigRequestsDto) Get() *UpdateWebhooksConfigRequestsDto {
	return v.value
}

func (v *NullableUpdateWebhooksConfigRequestsDto) Set(val *UpdateWebhooksConfigRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebhooksConfigRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebhooksConfigRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebhooksConfigRequestsDto(val *UpdateWebhooksConfigRequestsDto) *NullableUpdateWebhooksConfigRequestsDto {
	return &NullableUpdateWebhooksConfigRequestsDto{value: val, isSet: true}
}

func (v NullableUpdateWebhooksConfigRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebhooksConfigRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

