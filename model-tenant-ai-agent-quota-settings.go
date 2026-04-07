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

// checks if the TenantAiAgentQuotaSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantAiAgentQuotaSettings{}

// TenantAiAgentQuotaSettings The AI agent quota settings.
type TenantAiAgentQuotaSettings struct {
	// Specifies if the quota is enabled for the tenant entity or not.
	EnableQuota *bool `json:"enableQuota,omitempty"`
	// The default quota of the tenant entity.
	DefaultQuota *int64 `json:"defaultQuota,omitempty"`
	// The date of the last quota recalculation.
	LastRecalculateDate NullableTime `json:"lastRecalculateDate,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewTenantAiAgentQuotaSettings instantiates a new TenantAiAgentQuotaSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantAiAgentQuotaSettings() *TenantAiAgentQuotaSettings {
	this := TenantAiAgentQuotaSettings{}
	return &this
}

// NewTenantAiAgentQuotaSettingsWithDefaults instantiates a new TenantAiAgentQuotaSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantAiAgentQuotaSettingsWithDefaults() *TenantAiAgentQuotaSettings {
	this := TenantAiAgentQuotaSettings{}
	return &this
}

// GetEnableQuota returns the EnableQuota field value if set, zero value otherwise.
func (o *TenantAiAgentQuotaSettings) GetEnableQuota() bool {
	if o == nil || IsNil(o.EnableQuota) {
		var ret bool
		return ret
	}
	return *o.EnableQuota
}

// GetEnableQuotaOk returns a tuple with the EnableQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAiAgentQuotaSettings) GetEnableQuotaOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableQuota) {
		return nil, false
	}
	return o.EnableQuota, true
}

// HasEnableQuota returns a boolean if a field has been set.
func (o *TenantAiAgentQuotaSettings) IsEnableQuotaSet() bool {
	if o != nil && !IsNil(o.EnableQuota) {
		return true
	}

	return false
}

// SetEnableQuota gets a reference to the given bool and assigns it to the EnableQuota field.
func (o *TenantAiAgentQuotaSettings) SetEnableQuota(v bool) {
	o.EnableQuota = &v
}

// GetDefaultQuota returns the DefaultQuota field value if set, zero value otherwise.
func (o *TenantAiAgentQuotaSettings) GetDefaultQuota() int64 {
	if o == nil || IsNil(o.DefaultQuota) {
		var ret int64
		return ret
	}
	return *o.DefaultQuota
}

// GetDefaultQuotaOk returns a tuple with the DefaultQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAiAgentQuotaSettings) GetDefaultQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultQuota) {
		return nil, false
	}
	return o.DefaultQuota, true
}

// HasDefaultQuota returns a boolean if a field has been set.
func (o *TenantAiAgentQuotaSettings) IsDefaultQuotaSet() bool {
	if o != nil && !IsNil(o.DefaultQuota) {
		return true
	}

	return false
}

// SetDefaultQuota gets a reference to the given int64 and assigns it to the DefaultQuota field.
func (o *TenantAiAgentQuotaSettings) SetDefaultQuota(v int64) {
	o.DefaultQuota = &v
}

// GetLastRecalculateDate returns the LastRecalculateDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantAiAgentQuotaSettings) GetLastRecalculateDate() time.Time {
	if o == nil || IsNil(o.LastRecalculateDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastRecalculateDate.Get()
}

// GetLastRecalculateDateOk returns a tuple with the LastRecalculateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantAiAgentQuotaSettings) GetLastRecalculateDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastRecalculateDate.Get(), o.LastRecalculateDate.IsSet()
}

// HasLastRecalculateDate returns a boolean if a field has been set.
func (o *TenantAiAgentQuotaSettings) IsLastRecalculateDateSet() bool {
	if o != nil && o.LastRecalculateDate.IsSet() {
		return true
	}

	return false
}

// SetLastRecalculateDate gets a reference to the given NullableTime and assigns it to the LastRecalculateDate field.
func (o *TenantAiAgentQuotaSettings) SetLastRecalculateDate(v time.Time) {
	o.LastRecalculateDate.Set(&v)
}
// SetLastRecalculateDateNil sets the value for LastRecalculateDate to be an explicit nil
func (o *TenantAiAgentQuotaSettings) SetLastRecalculateDateNil() {
	o.LastRecalculateDate.Set(nil)
}

// UnsetLastRecalculateDate ensures that no value is present for LastRecalculateDate, not even an explicit nil
func (o *TenantAiAgentQuotaSettings) UnsetLastRecalculateDate() {
	o.LastRecalculateDate.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *TenantAiAgentQuotaSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAiAgentQuotaSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *TenantAiAgentQuotaSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *TenantAiAgentQuotaSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o TenantAiAgentQuotaSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantAiAgentQuotaSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableQuota) {
		toSerialize["enableQuota"] = o.EnableQuota
	}
	if !IsNil(o.DefaultQuota) {
		toSerialize["defaultQuota"] = o.DefaultQuota
	}
	if o.LastRecalculateDate.IsSet() {
		toSerialize["lastRecalculateDate"] = o.LastRecalculateDate.Get()
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableTenantAiAgentQuotaSettings struct {
	value *TenantAiAgentQuotaSettings
	isSet bool
}

func (v NullableTenantAiAgentQuotaSettings) Get() *TenantAiAgentQuotaSettings {
	return v.value
}

func (v *NullableTenantAiAgentQuotaSettings) Set(val *TenantAiAgentQuotaSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantAiAgentQuotaSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantAiAgentQuotaSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantAiAgentQuotaSettings(val *TenantAiAgentQuotaSettings) *NullableTenantAiAgentQuotaSettings {
	return &NullableTenantAiAgentQuotaSettings{value: val, isSet: true}
}

func (v NullableTenantAiAgentQuotaSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantAiAgentQuotaSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

