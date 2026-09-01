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

// checks if the TenantAuditSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantAuditSettings{}

// TenantAuditSettings The tenant audit settings parameters.
type TenantAuditSettings struct {
	// The login history lifetime.
	LoginHistoryLifeTime *int32 `json:"loginHistoryLifeTime,omitempty"`
	// The audit trail lifetime.
	AuditTrailLifeTime *int32 `json:"auditTrailLifeTime,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewTenantAuditSettings instantiates a new TenantAuditSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantAuditSettings() *TenantAuditSettings {
	this := TenantAuditSettings{}
	return &this
}

// NewTenantAuditSettingsWithDefaults instantiates a new TenantAuditSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantAuditSettingsWithDefaults() *TenantAuditSettings {
	this := TenantAuditSettings{}
	return &this
}

// GetLoginHistoryLifeTime returns the LoginHistoryLifeTime field value if set, zero value otherwise.
func (o *TenantAuditSettings) GetLoginHistoryLifeTime() int32 {
	if o == nil || IsNil(o.LoginHistoryLifeTime) {
		var ret int32
		return ret
	}
	return *o.LoginHistoryLifeTime
}

// GetLoginHistoryLifeTimeOk returns a tuple with the LoginHistoryLifeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAuditSettings) GetLoginHistoryLifeTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.LoginHistoryLifeTime) {
		return nil, false
	}
	return o.LoginHistoryLifeTime, true
}

// HasLoginHistoryLifeTime returns a boolean if a field has been set.
func (o *TenantAuditSettings) IsLoginHistoryLifeTimeSet() bool {
	if o != nil && !IsNil(o.LoginHistoryLifeTime) {
		return true
	}

	return false
}

// SetLoginHistoryLifeTime gets a reference to the given int32 and assigns it to the LoginHistoryLifeTime field.
func (o *TenantAuditSettings) SetLoginHistoryLifeTime(v int32) {
	o.LoginHistoryLifeTime = &v
}

// GetAuditTrailLifeTime returns the AuditTrailLifeTime field value if set, zero value otherwise.
func (o *TenantAuditSettings) GetAuditTrailLifeTime() int32 {
	if o == nil || IsNil(o.AuditTrailLifeTime) {
		var ret int32
		return ret
	}
	return *o.AuditTrailLifeTime
}

// GetAuditTrailLifeTimeOk returns a tuple with the AuditTrailLifeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAuditSettings) GetAuditTrailLifeTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditTrailLifeTime) {
		return nil, false
	}
	return o.AuditTrailLifeTime, true
}

// HasAuditTrailLifeTime returns a boolean if a field has been set.
func (o *TenantAuditSettings) IsAuditTrailLifeTimeSet() bool {
	if o != nil && !IsNil(o.AuditTrailLifeTime) {
		return true
	}

	return false
}

// SetAuditTrailLifeTime gets a reference to the given int32 and assigns it to the AuditTrailLifeTime field.
func (o *TenantAuditSettings) SetAuditTrailLifeTime(v int32) {
	o.AuditTrailLifeTime = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *TenantAuditSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAuditSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *TenantAuditSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *TenantAuditSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o TenantAuditSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantAuditSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoginHistoryLifeTime) {
		toSerialize["loginHistoryLifeTime"] = o.LoginHistoryLifeTime
	}
	if !IsNil(o.AuditTrailLifeTime) {
		toSerialize["auditTrailLifeTime"] = o.AuditTrailLifeTime
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableTenantAuditSettings struct {
	value *TenantAuditSettings
	isSet bool
}

func (v NullableTenantAuditSettings) Get() *TenantAuditSettings {
	return v.value
}

func (v *NullableTenantAuditSettings) Set(val *TenantAuditSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantAuditSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantAuditSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantAuditSettings(val *TenantAuditSettings) *NullableTenantAuditSettings {
	return &NullableTenantAuditSettings{value: val, isSet: true}
}

func (v NullableTenantAuditSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantAuditSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

