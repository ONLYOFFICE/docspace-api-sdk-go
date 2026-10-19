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

// checks if the SsoIdpSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoIdpSettingsDto{}

// SsoIdpSettingsDto The identity provider the portal trusts: its entity id, endpoints and bindings.
type SsoIdpSettingsDto struct {
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

// NewSsoIdpSettingsDto instantiates a new SsoIdpSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoIdpSettingsDto() *SsoIdpSettingsDto {
	this := SsoIdpSettingsDto{}
	return &this
}

// NewSsoIdpSettingsDtoWithDefaults instantiates a new SsoIdpSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoIdpSettingsDtoWithDefaults() *SsoIdpSettingsDto {
	this := SsoIdpSettingsDto{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettingsDto) GetEntityId() string {
	if o == nil || IsNil(o.EntityId.Get()) {
		var ret string
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettingsDto) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field has been set.
func (o *SsoIdpSettingsDto) IsEntityIdSet() bool {
	if o != nil && o.EntityId.IsSet() {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableString and assigns it to the EntityId field.
func (o *SsoIdpSettingsDto) SetEntityId(v string) {
	o.EntityId.Set(&v)
}
// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *SsoIdpSettingsDto) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *SsoIdpSettingsDto) UnsetEntityId() {
	o.EntityId.Unset()
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettingsDto) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SsoUrl.Get()
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettingsDto) GetSsoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoUrl.Get(), o.SsoUrl.IsSet()
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *SsoIdpSettingsDto) IsSsoUrlSet() bool {
	if o != nil && o.SsoUrl.IsSet() {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given NullableString and assigns it to the SsoUrl field.
func (o *SsoIdpSettingsDto) SetSsoUrl(v string) {
	o.SsoUrl.Set(&v)
}
// SetSsoUrlNil sets the value for SsoUrl to be an explicit nil
func (o *SsoIdpSettingsDto) SetSsoUrlNil() {
	o.SsoUrl.Set(nil)
}

// UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil
func (o *SsoIdpSettingsDto) UnsetSsoUrl() {
	o.SsoUrl.Unset()
}

// GetSsoBinding returns the SsoBinding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettingsDto) GetSsoBinding() string {
	if o == nil || IsNil(o.SsoBinding.Get()) {
		var ret string
		return ret
	}
	return *o.SsoBinding.Get()
}

// GetSsoBindingOk returns a tuple with the SsoBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettingsDto) GetSsoBindingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoBinding.Get(), o.SsoBinding.IsSet()
}

// HasSsoBinding returns a boolean if a field has been set.
func (o *SsoIdpSettingsDto) IsSsoBindingSet() bool {
	if o != nil && o.SsoBinding.IsSet() {
		return true
	}

	return false
}

// SetSsoBinding gets a reference to the given NullableString and assigns it to the SsoBinding field.
func (o *SsoIdpSettingsDto) SetSsoBinding(v string) {
	o.SsoBinding.Set(&v)
}
// SetSsoBindingNil sets the value for SsoBinding to be an explicit nil
func (o *SsoIdpSettingsDto) SetSsoBindingNil() {
	o.SsoBinding.Set(nil)
}

// UnsetSsoBinding ensures that no value is present for SsoBinding, not even an explicit nil
func (o *SsoIdpSettingsDto) UnsetSsoBinding() {
	o.SsoBinding.Unset()
}

// GetSloUrl returns the SloUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettingsDto) GetSloUrl() string {
	if o == nil || IsNil(o.SloUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SloUrl.Get()
}

// GetSloUrlOk returns a tuple with the SloUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettingsDto) GetSloUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SloUrl.Get(), o.SloUrl.IsSet()
}

// HasSloUrl returns a boolean if a field has been set.
func (o *SsoIdpSettingsDto) IsSloUrlSet() bool {
	if o != nil && o.SloUrl.IsSet() {
		return true
	}

	return false
}

// SetSloUrl gets a reference to the given NullableString and assigns it to the SloUrl field.
func (o *SsoIdpSettingsDto) SetSloUrl(v string) {
	o.SloUrl.Set(&v)
}
// SetSloUrlNil sets the value for SloUrl to be an explicit nil
func (o *SsoIdpSettingsDto) SetSloUrlNil() {
	o.SloUrl.Set(nil)
}

// UnsetSloUrl ensures that no value is present for SloUrl, not even an explicit nil
func (o *SsoIdpSettingsDto) UnsetSloUrl() {
	o.SloUrl.Unset()
}

// GetSloBinding returns the SloBinding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettingsDto) GetSloBinding() string {
	if o == nil || IsNil(o.SloBinding.Get()) {
		var ret string
		return ret
	}
	return *o.SloBinding.Get()
}

// GetSloBindingOk returns a tuple with the SloBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettingsDto) GetSloBindingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SloBinding.Get(), o.SloBinding.IsSet()
}

// HasSloBinding returns a boolean if a field has been set.
func (o *SsoIdpSettingsDto) IsSloBindingSet() bool {
	if o != nil && o.SloBinding.IsSet() {
		return true
	}

	return false
}

// SetSloBinding gets a reference to the given NullableString and assigns it to the SloBinding field.
func (o *SsoIdpSettingsDto) SetSloBinding(v string) {
	o.SloBinding.Set(&v)
}
// SetSloBindingNil sets the value for SloBinding to be an explicit nil
func (o *SsoIdpSettingsDto) SetSloBindingNil() {
	o.SloBinding.Set(nil)
}

// UnsetSloBinding ensures that no value is present for SloBinding, not even an explicit nil
func (o *SsoIdpSettingsDto) UnsetSloBinding() {
	o.SloBinding.Unset()
}

// GetNameIdFormat returns the NameIdFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpSettingsDto) GetNameIdFormat() string {
	if o == nil || IsNil(o.NameIdFormat.Get()) {
		var ret string
		return ret
	}
	return *o.NameIdFormat.Get()
}

// GetNameIdFormatOk returns a tuple with the NameIdFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpSettingsDto) GetNameIdFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameIdFormat.Get(), o.NameIdFormat.IsSet()
}

// HasNameIdFormat returns a boolean if a field has been set.
func (o *SsoIdpSettingsDto) IsNameIdFormatSet() bool {
	if o != nil && o.NameIdFormat.IsSet() {
		return true
	}

	return false
}

// SetNameIdFormat gets a reference to the given NullableString and assigns it to the NameIdFormat field.
func (o *SsoIdpSettingsDto) SetNameIdFormat(v string) {
	o.NameIdFormat.Set(&v)
}
// SetNameIdFormatNil sets the value for NameIdFormat to be an explicit nil
func (o *SsoIdpSettingsDto) SetNameIdFormatNil() {
	o.NameIdFormat.Set(nil)
}

// UnsetNameIdFormat ensures that no value is present for NameIdFormat, not even an explicit nil
func (o *SsoIdpSettingsDto) UnsetNameIdFormat() {
	o.NameIdFormat.Unset()
}

func (o SsoIdpSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoIdpSettingsDto) ToMap() (map[string]interface{}, error) {
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

type NullableSsoIdpSettingsDto struct {
	value *SsoIdpSettingsDto
	isSet bool
}

func (v NullableSsoIdpSettingsDto) Get() *SsoIdpSettingsDto {
	return v.value
}

func (v *NullableSsoIdpSettingsDto) Set(val *SsoIdpSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoIdpSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoIdpSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoIdpSettingsDto(val *SsoIdpSettingsDto) *NullableSsoIdpSettingsDto {
	return &NullableSsoIdpSettingsDto{value: val, isSet: true}
}

func (v NullableSsoIdpSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoIdpSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

