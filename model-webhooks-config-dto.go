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
	"time"
	"bytes"
	"fmt"
)

// checks if the WebhooksConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksConfigDto{}

// WebhooksConfigDto The webhook configuration parameters.
type WebhooksConfigDto struct {
	// The webhook ID.
	Id int32 `json:"id"`
	// The webhook name.
	Name NullableString `json:"name,omitempty"`
	// The webhook URI.
	Uri NullableString `json:"uri,omitempty"`
	// Specifies if the webhooks are enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// The webhook SSL verification (enabled or not).
	Ssl *bool `json:"ssl,omitempty"`
	Triggers *WebhookTrigger `json:"triggers,omitempty"`
	// The webhook target ID.
	TargetId NullableString `json:"targetId,omitempty"`
	CreatedBy *EmployeeDto `json:"createdBy,omitempty"`
	// The date and time when the webhook was created.
	CreatedOn NullableTime `json:"createdOn,omitempty"`
	ModifiedBy *EmployeeDto `json:"modifiedBy,omitempty"`
	// The date and time when the webhook was modified.
	ModifiedOn NullableTime `json:"modifiedOn,omitempty"`
	// The date and time of the webhook last failure.
	LastFailureOn NullableTime `json:"lastFailureOn,omitempty"`
	// The webhook last failure content.
	LastFailureContent NullableString `json:"lastFailureContent,omitempty"`
	// The date and time of the webhook last success.
	LastSuccessOn NullableTime `json:"lastSuccessOn,omitempty"`
}

type _WebhooksConfigDto WebhooksConfigDto

// NewWebhooksConfigDto instantiates a new WebhooksConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksConfigDto(id int32) *WebhooksConfigDto {
	this := WebhooksConfigDto{}
	this.Id = id
	return &this
}

// NewWebhooksConfigDtoWithDefaults instantiates a new WebhooksConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksConfigDtoWithDefaults() *WebhooksConfigDto {
	this := WebhooksConfigDto{}
	return &this
}

// GetId returns the Id field value
func (o *WebhooksConfigDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhooksConfigDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhooksConfigDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksConfigDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksConfigDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WebhooksConfigDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WebhooksConfigDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WebhooksConfigDto) UnsetName() {
	o.Name.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksConfigDto) GetUri() string {
	if o == nil || IsNil(o.Uri.Get()) {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksConfigDto) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsUriSet() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *WebhooksConfigDto) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *WebhooksConfigDto) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *WebhooksConfigDto) UnsetUri() {
	o.Uri.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WebhooksConfigDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfigDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WebhooksConfigDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *WebhooksConfigDto) GetSsl() bool {
	if o == nil || IsNil(o.Ssl) {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfigDto) GetSslOk() (*bool, bool) {
	if o == nil || IsNil(o.Ssl) {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsSslSet() bool {
	if o != nil && !IsNil(o.Ssl) {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *WebhooksConfigDto) SetSsl(v bool) {
	o.Ssl = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *WebhooksConfigDto) GetTriggers() WebhookTrigger {
	if o == nil || IsNil(o.Triggers) {
		var ret WebhookTrigger
		return ret
	}
	return *o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfigDto) GetTriggersOk() (*WebhookTrigger, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsTriggersSet() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given WebhookTrigger and assigns it to the Triggers field.
func (o *WebhooksConfigDto) SetTriggers(v WebhookTrigger) {
	o.Triggers = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksConfigDto) GetTargetId() string {
	if o == nil || IsNil(o.TargetId.Get()) {
		var ret string
		return ret
	}
	return *o.TargetId.Get()
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksConfigDto) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetId.Get(), o.TargetId.IsSet()
}

// HasTargetId returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsTargetIdSet() bool {
	if o != nil && o.TargetId.IsSet() {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given NullableString and assigns it to the TargetId field.
func (o *WebhooksConfigDto) SetTargetId(v string) {
	o.TargetId.Set(&v)
}
// SetTargetIdNil sets the value for TargetId to be an explicit nil
func (o *WebhooksConfigDto) SetTargetIdNil() {
	o.TargetId.Set(nil)
}

// UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
func (o *WebhooksConfigDto) UnsetTargetId() {
	o.TargetId.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WebhooksConfigDto) GetCreatedBy() EmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfigDto) GetCreatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EmployeeDto and assigns it to the CreatedBy field.
func (o *WebhooksConfigDto) SetCreatedBy(v EmployeeDto) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksConfigDto) GetCreatedOn() time.Time {
	if o == nil || IsNil(o.CreatedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn.Get()
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksConfigDto) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedOn.Get(), o.CreatedOn.IsSet()
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsCreatedOnSet() bool {
	if o != nil && o.CreatedOn.IsSet() {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given NullableTime and assigns it to the CreatedOn field.
func (o *WebhooksConfigDto) SetCreatedOn(v time.Time) {
	o.CreatedOn.Set(&v)
}
// SetCreatedOnNil sets the value for CreatedOn to be an explicit nil
func (o *WebhooksConfigDto) SetCreatedOnNil() {
	o.CreatedOn.Set(nil)
}

// UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
func (o *WebhooksConfigDto) UnsetCreatedOn() {
	o.CreatedOn.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *WebhooksConfigDto) GetModifiedBy() EmployeeDto {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfigDto) GetModifiedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsModifiedBySet() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given EmployeeDto and assigns it to the ModifiedBy field.
func (o *WebhooksConfigDto) SetModifiedBy(v EmployeeDto) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksConfigDto) GetModifiedOn() time.Time {
	if o == nil || IsNil(o.ModifiedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn.Get()
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksConfigDto) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedOn.Get(), o.ModifiedOn.IsSet()
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsModifiedOnSet() bool {
	if o != nil && o.ModifiedOn.IsSet() {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given NullableTime and assigns it to the ModifiedOn field.
func (o *WebhooksConfigDto) SetModifiedOn(v time.Time) {
	o.ModifiedOn.Set(&v)
}
// SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil
func (o *WebhooksConfigDto) SetModifiedOnNil() {
	o.ModifiedOn.Set(nil)
}

// UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
func (o *WebhooksConfigDto) UnsetModifiedOn() {
	o.ModifiedOn.Unset()
}

// GetLastFailureOn returns the LastFailureOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksConfigDto) GetLastFailureOn() time.Time {
	if o == nil || IsNil(o.LastFailureOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastFailureOn.Get()
}

// GetLastFailureOnOk returns a tuple with the LastFailureOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksConfigDto) GetLastFailureOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastFailureOn.Get(), o.LastFailureOn.IsSet()
}

// HasLastFailureOn returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsLastFailureOnSet() bool {
	if o != nil && o.LastFailureOn.IsSet() {
		return true
	}

	return false
}

// SetLastFailureOn gets a reference to the given NullableTime and assigns it to the LastFailureOn field.
func (o *WebhooksConfigDto) SetLastFailureOn(v time.Time) {
	o.LastFailureOn.Set(&v)
}
// SetLastFailureOnNil sets the value for LastFailureOn to be an explicit nil
func (o *WebhooksConfigDto) SetLastFailureOnNil() {
	o.LastFailureOn.Set(nil)
}

// UnsetLastFailureOn ensures that no value is present for LastFailureOn, not even an explicit nil
func (o *WebhooksConfigDto) UnsetLastFailureOn() {
	o.LastFailureOn.Unset()
}

// GetLastFailureContent returns the LastFailureContent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksConfigDto) GetLastFailureContent() string {
	if o == nil || IsNil(o.LastFailureContent.Get()) {
		var ret string
		return ret
	}
	return *o.LastFailureContent.Get()
}

// GetLastFailureContentOk returns a tuple with the LastFailureContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksConfigDto) GetLastFailureContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastFailureContent.Get(), o.LastFailureContent.IsSet()
}

// HasLastFailureContent returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsLastFailureContentSet() bool {
	if o != nil && o.LastFailureContent.IsSet() {
		return true
	}

	return false
}

// SetLastFailureContent gets a reference to the given NullableString and assigns it to the LastFailureContent field.
func (o *WebhooksConfigDto) SetLastFailureContent(v string) {
	o.LastFailureContent.Set(&v)
}
// SetLastFailureContentNil sets the value for LastFailureContent to be an explicit nil
func (o *WebhooksConfigDto) SetLastFailureContentNil() {
	o.LastFailureContent.Set(nil)
}

// UnsetLastFailureContent ensures that no value is present for LastFailureContent, not even an explicit nil
func (o *WebhooksConfigDto) UnsetLastFailureContent() {
	o.LastFailureContent.Unset()
}

// GetLastSuccessOn returns the LastSuccessOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksConfigDto) GetLastSuccessOn() time.Time {
	if o == nil || IsNil(o.LastSuccessOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSuccessOn.Get()
}

// GetLastSuccessOnOk returns a tuple with the LastSuccessOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksConfigDto) GetLastSuccessOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSuccessOn.Get(), o.LastSuccessOn.IsSet()
}

// HasLastSuccessOn returns a boolean if a field has been set.
func (o *WebhooksConfigDto) IsLastSuccessOnSet() bool {
	if o != nil && o.LastSuccessOn.IsSet() {
		return true
	}

	return false
}

// SetLastSuccessOn gets a reference to the given NullableTime and assigns it to the LastSuccessOn field.
func (o *WebhooksConfigDto) SetLastSuccessOn(v time.Time) {
	o.LastSuccessOn.Set(&v)
}
// SetLastSuccessOnNil sets the value for LastSuccessOn to be an explicit nil
func (o *WebhooksConfigDto) SetLastSuccessOnNil() {
	o.LastSuccessOn.Set(nil)
}

// UnsetLastSuccessOn ensures that no value is present for LastSuccessOn, not even an explicit nil
func (o *WebhooksConfigDto) UnsetLastSuccessOn() {
	o.LastSuccessOn.Unset()
}

func (o WebhooksConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
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
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedOn.IsSet() {
		toSerialize["createdOn"] = o.CreatedOn.Get()
	}
	if !IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.ModifiedOn.IsSet() {
		toSerialize["modifiedOn"] = o.ModifiedOn.Get()
	}
	if o.LastFailureOn.IsSet() {
		toSerialize["lastFailureOn"] = o.LastFailureOn.Get()
	}
	if o.LastFailureContent.IsSet() {
		toSerialize["lastFailureContent"] = o.LastFailureContent.Get()
	}
	if o.LastSuccessOn.IsSet() {
		toSerialize["lastSuccessOn"] = o.LastSuccessOn.Get()
	}
	return toSerialize, nil
}

func (o *WebhooksConfigDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varWebhooksConfigDto := _WebhooksConfigDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWebhooksConfigDto)

	if err != nil {
		return err
	}

	*o = WebhooksConfigDto(varWebhooksConfigDto)

	return err
}

type NullableWebhooksConfigDto struct {
	value *WebhooksConfigDto
	isSet bool
}

func (v NullableWebhooksConfigDto) Get() *WebhooksConfigDto {
	return v.value
}

func (v *NullableWebhooksConfigDto) Set(val *WebhooksConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksConfigDto(val *WebhooksConfigDto) *NullableWebhooksConfigDto {
	return &NullableWebhooksConfigDto{value: val, isSet: true}
}

func (v NullableWebhooksConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

