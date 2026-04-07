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
)

// checks if the SsoCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoCertificate{}

// SsoCertificate The SSO certificate parameters.
type SsoCertificate struct {
	// Specifies if a certificate is self-signed or not.
	SelfSigned *bool `json:"selfSigned,omitempty"`
	// The CRT certificate file.
	Crt NullableString `json:"crt,omitempty"`
	// The certificate key.
	Key NullableString `json:"key,omitempty"`
	// The certificate action.
	Action NullableString `json:"action,omitempty"`
	// The certificate domain name.
	DomainName NullableString `json:"domainName,omitempty"`
	// The certificate start date.
	StartDate *time.Time `json:"startDate,omitempty"`
	// The certificate expiration date.
	ExpiredDate *time.Time `json:"expiredDate,omitempty"`
}

// NewSsoCertificate instantiates a new SsoCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoCertificate() *SsoCertificate {
	this := SsoCertificate{}
	return &this
}

// NewSsoCertificateWithDefaults instantiates a new SsoCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoCertificateWithDefaults() *SsoCertificate {
	this := SsoCertificate{}
	return &this
}

// GetSelfSigned returns the SelfSigned field value if set, zero value otherwise.
func (o *SsoCertificate) GetSelfSigned() bool {
	if o == nil || IsNil(o.SelfSigned) {
		var ret bool
		return ret
	}
	return *o.SelfSigned
}

// GetSelfSignedOk returns a tuple with the SelfSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoCertificate) GetSelfSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfSigned) {
		return nil, false
	}
	return o.SelfSigned, true
}

// HasSelfSigned returns a boolean if a field has been set.
func (o *SsoCertificate) IsSelfSignedSet() bool {
	if o != nil && !IsNil(o.SelfSigned) {
		return true
	}

	return false
}

// SetSelfSigned gets a reference to the given bool and assigns it to the SelfSigned field.
func (o *SsoCertificate) SetSelfSigned(v bool) {
	o.SelfSigned = &v
}

// GetCrt returns the Crt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoCertificate) GetCrt() string {
	if o == nil || IsNil(o.Crt.Get()) {
		var ret string
		return ret
	}
	return *o.Crt.Get()
}

// GetCrtOk returns a tuple with the Crt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoCertificate) GetCrtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Crt.Get(), o.Crt.IsSet()
}

// HasCrt returns a boolean if a field has been set.
func (o *SsoCertificate) IsCrtSet() bool {
	if o != nil && o.Crt.IsSet() {
		return true
	}

	return false
}

// SetCrt gets a reference to the given NullableString and assigns it to the Crt field.
func (o *SsoCertificate) SetCrt(v string) {
	o.Crt.Set(&v)
}
// SetCrtNil sets the value for Crt to be an explicit nil
func (o *SsoCertificate) SetCrtNil() {
	o.Crt.Set(nil)
}

// UnsetCrt ensures that no value is present for Crt, not even an explicit nil
func (o *SsoCertificate) UnsetCrt() {
	o.Crt.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoCertificate) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoCertificate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *SsoCertificate) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *SsoCertificate) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *SsoCertificate) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *SsoCertificate) UnsetKey() {
	o.Key.Unset()
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoCertificate) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoCertificate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *SsoCertificate) IsActionSet() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *SsoCertificate) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *SsoCertificate) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *SsoCertificate) UnsetAction() {
	o.Action.Unset()
}

// GetDomainName returns the DomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoCertificate) GetDomainName() string {
	if o == nil || IsNil(o.DomainName.Get()) {
		var ret string
		return ret
	}
	return *o.DomainName.Get()
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoCertificate) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DomainName.Get(), o.DomainName.IsSet()
}

// HasDomainName returns a boolean if a field has been set.
func (o *SsoCertificate) IsDomainNameSet() bool {
	if o != nil && o.DomainName.IsSet() {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given NullableString and assigns it to the DomainName field.
func (o *SsoCertificate) SetDomainName(v string) {
	o.DomainName.Set(&v)
}
// SetDomainNameNil sets the value for DomainName to be an explicit nil
func (o *SsoCertificate) SetDomainNameNil() {
	o.DomainName.Set(nil)
}

// UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
func (o *SsoCertificate) UnsetDomainName() {
	o.DomainName.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SsoCertificate) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoCertificate) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SsoCertificate) IsStartDateSet() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *SsoCertificate) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiredDate returns the ExpiredDate field value if set, zero value otherwise.
func (o *SsoCertificate) GetExpiredDate() time.Time {
	if o == nil || IsNil(o.ExpiredDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpiredDate
}

// GetExpiredDateOk returns a tuple with the ExpiredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoCertificate) GetExpiredDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiredDate) {
		return nil, false
	}
	return o.ExpiredDate, true
}

// HasExpiredDate returns a boolean if a field has been set.
func (o *SsoCertificate) IsExpiredDateSet() bool {
	if o != nil && !IsNil(o.ExpiredDate) {
		return true
	}

	return false
}

// SetExpiredDate gets a reference to the given time.Time and assigns it to the ExpiredDate field.
func (o *SsoCertificate) SetExpiredDate(v time.Time) {
	o.ExpiredDate = &v
}

func (o SsoCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SelfSigned) {
		toSerialize["selfSigned"] = o.SelfSigned
	}
	if o.Crt.IsSet() {
		toSerialize["crt"] = o.Crt.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.DomainName.IsSet() {
		toSerialize["domainName"] = o.DomainName.Get()
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.ExpiredDate) {
		toSerialize["expiredDate"] = o.ExpiredDate
	}
	return toSerialize, nil
}

type NullableSsoCertificate struct {
	value *SsoCertificate
	isSet bool
}

func (v NullableSsoCertificate) Get() *SsoCertificate {
	return v.value
}

func (v *NullableSsoCertificate) Set(val *SsoCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoCertificate(val *SsoCertificate) *NullableSsoCertificate {
	return &NullableSsoCertificate{value: val, isSet: true}
}

func (v NullableSsoCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

