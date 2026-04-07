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

// checks if the SsoIdpSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoIdpSettings{}

// SsoIdpSettings The SSO IdP settings.
type SsoIdpSettings struct {
	// The entity ID.
	EntityId NullableString `json:"entityId,omitempty"`
	// The SSO URL.
	SsoUrl NullableString `json:"ssoUrl,omitempty"`
	// The SSO binding.
	SsoBinding NullableString `json:"ssoBinding,omitempty"`
	// The SLO URL.
	SloUrl NullableString `json:"sloUrl,omitempty"`
	// The SLO binding.
	SloBinding NullableString `json:"sloBinding,omitempty"`
	// The name ID format.
	NameIdFormat NullableString `json:"nameIdFormat,omitempty"`
}

// NewSsoIdpSettings instantiates a new SsoIdpSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoIdpSettings() *SsoIdpSettings {
	this := SsoIdpSettings{}
	return &this
}

// NewSsoIdpSettingsWithDefaults instantiates a new SsoIdpSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoIdpSettingsWithDefaults() *SsoIdpSettings {
	this := SsoIdpSettings{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettings) GetEntityId() string {
	if o == nil || IsNil(o.EntityId.Get()) {
		var ret string
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettings) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field has been set.
func (o *SsoIdpSettings) IsEntityIdSet() bool {
	if o != nil && o.EntityId.IsSet() {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableString and assigns it to the EntityId field.
func (o *SsoIdpSettings) SetEntityId(v string) {
	o.EntityId.Set(&v)
}
// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *SsoIdpSettings) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *SsoIdpSettings) UnsetEntityId() {
	o.EntityId.Unset()
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettings) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SsoUrl.Get()
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettings) GetSsoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoUrl.Get(), o.SsoUrl.IsSet()
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *SsoIdpSettings) IsSsoUrlSet() bool {
	if o != nil && o.SsoUrl.IsSet() {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given NullableString and assigns it to the SsoUrl field.
func (o *SsoIdpSettings) SetSsoUrl(v string) {
	o.SsoUrl.Set(&v)
}
// SetSsoUrlNil sets the value for SsoUrl to be an explicit nil
func (o *SsoIdpSettings) SetSsoUrlNil() {
	o.SsoUrl.Set(nil)
}

// UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil
func (o *SsoIdpSettings) UnsetSsoUrl() {
	o.SsoUrl.Unset()
}

// GetSsoBinding returns the SsoBinding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettings) GetSsoBinding() string {
	if o == nil || IsNil(o.SsoBinding.Get()) {
		var ret string
		return ret
	}
	return *o.SsoBinding.Get()
}

// GetSsoBindingOk returns a tuple with the SsoBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettings) GetSsoBindingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoBinding.Get(), o.SsoBinding.IsSet()
}

// HasSsoBinding returns a boolean if a field has been set.
func (o *SsoIdpSettings) IsSsoBindingSet() bool {
	if o != nil && o.SsoBinding.IsSet() {
		return true
	}

	return false
}

// SetSsoBinding gets a reference to the given NullableString and assigns it to the SsoBinding field.
func (o *SsoIdpSettings) SetSsoBinding(v string) {
	o.SsoBinding.Set(&v)
}
// SetSsoBindingNil sets the value for SsoBinding to be an explicit nil
func (o *SsoIdpSettings) SetSsoBindingNil() {
	o.SsoBinding.Set(nil)
}

// UnsetSsoBinding ensures that no value is present for SsoBinding, not even an explicit nil
func (o *SsoIdpSettings) UnsetSsoBinding() {
	o.SsoBinding.Unset()
}

// GetSloUrl returns the SloUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettings) GetSloUrl() string {
	if o == nil || IsNil(o.SloUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SloUrl.Get()
}

// GetSloUrlOk returns a tuple with the SloUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettings) GetSloUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SloUrl.Get(), o.SloUrl.IsSet()
}

// HasSloUrl returns a boolean if a field has been set.
func (o *SsoIdpSettings) IsSloUrlSet() bool {
	if o != nil && o.SloUrl.IsSet() {
		return true
	}

	return false
}

// SetSloUrl gets a reference to the given NullableString and assigns it to the SloUrl field.
func (o *SsoIdpSettings) SetSloUrl(v string) {
	o.SloUrl.Set(&v)
}
// SetSloUrlNil sets the value for SloUrl to be an explicit nil
func (o *SsoIdpSettings) SetSloUrlNil() {
	o.SloUrl.Set(nil)
}

// UnsetSloUrl ensures that no value is present for SloUrl, not even an explicit nil
func (o *SsoIdpSettings) UnsetSloUrl() {
	o.SloUrl.Unset()
}

// GetSloBinding returns the SloBinding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettings) GetSloBinding() string {
	if o == nil || IsNil(o.SloBinding.Get()) {
		var ret string
		return ret
	}
	return *o.SloBinding.Get()
}

// GetSloBindingOk returns a tuple with the SloBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettings) GetSloBindingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SloBinding.Get(), o.SloBinding.IsSet()
}

// HasSloBinding returns a boolean if a field has been set.
func (o *SsoIdpSettings) IsSloBindingSet() bool {
	if o != nil && o.SloBinding.IsSet() {
		return true
	}

	return false
}

// SetSloBinding gets a reference to the given NullableString and assigns it to the SloBinding field.
func (o *SsoIdpSettings) SetSloBinding(v string) {
	o.SloBinding.Set(&v)
}
// SetSloBindingNil sets the value for SloBinding to be an explicit nil
func (o *SsoIdpSettings) SetSloBindingNil() {
	o.SloBinding.Set(nil)
}

// UnsetSloBinding ensures that no value is present for SloBinding, not even an explicit nil
func (o *SsoIdpSettings) UnsetSloBinding() {
	o.SloBinding.Unset()
}

// GetNameIdFormat returns the NameIdFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettings) GetNameIdFormat() string {
	if o == nil || IsNil(o.NameIdFormat.Get()) {
		var ret string
		return ret
	}
	return *o.NameIdFormat.Get()
}

// GetNameIdFormatOk returns a tuple with the NameIdFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettings) GetNameIdFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameIdFormat.Get(), o.NameIdFormat.IsSet()
}

// HasNameIdFormat returns a boolean if a field has been set.
func (o *SsoIdpSettings) IsNameIdFormatSet() bool {
	if o != nil && o.NameIdFormat.IsSet() {
		return true
	}

	return false
}

// SetNameIdFormat gets a reference to the given NullableString and assigns it to the NameIdFormat field.
func (o *SsoIdpSettings) SetNameIdFormat(v string) {
	o.NameIdFormat.Set(&v)
}
// SetNameIdFormatNil sets the value for NameIdFormat to be an explicit nil
func (o *SsoIdpSettings) SetNameIdFormatNil() {
	o.NameIdFormat.Set(nil)
}

// UnsetNameIdFormat ensures that no value is present for NameIdFormat, not even an explicit nil
func (o *SsoIdpSettings) UnsetNameIdFormat() {
	o.NameIdFormat.Unset()
}

func (o SsoIdpSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoIdpSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityId.IsSet() {
		toSerialize["entityId"] = o.EntityId.Get()
	}
	if o.SsoUrl.IsSet() {
		toSerialize["ssoUrl"] = o.SsoUrl.Get()
	}
	if o.SsoBinding.IsSet() {
		toSerialize["ssoBinding"] = o.SsoBinding.Get()
	}
	if o.SloUrl.IsSet() {
		toSerialize["sloUrl"] = o.SloUrl.Get()
	}
	if o.SloBinding.IsSet() {
		toSerialize["sloBinding"] = o.SloBinding.Get()
	}
	if o.NameIdFormat.IsSet() {
		toSerialize["nameIdFormat"] = o.NameIdFormat.Get()
	}
	return toSerialize, nil
}

type NullableSsoIdpSettings struct {
	value *SsoIdpSettings
	isSet bool
}

func (v NullableSsoIdpSettings) Get() *SsoIdpSettings {
	return v.value
}

func (v *NullableSsoIdpSettings) Set(val *SsoIdpSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoIdpSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoIdpSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoIdpSettings(val *SsoIdpSettings) *NullableSsoIdpSettings {
	return &NullableSsoIdpSettings{value: val, isSet: true}
}

func (v NullableSsoIdpSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoIdpSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

