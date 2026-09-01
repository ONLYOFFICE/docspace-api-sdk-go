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

// checks if the TenantRoomQuotaSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantRoomQuotaSettings{}

// TenantRoomQuotaSettings The room quota settings.
type TenantRoomQuotaSettings struct {
	// Specifies if the quota is enabled for the tenant entity or not.
	EnableQuota *bool `json:"enableQuota,omitempty"`
	// The default quota of the tenant entity.
	DefaultQuota *int64 `json:"defaultQuota,omitempty"`
	// The date of the last quota recalculation.
	LastRecalculateDate *time.Time `json:"lastRecalculateDate,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewTenantRoomQuotaSettings instantiates a new TenantRoomQuotaSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantRoomQuotaSettings() *TenantRoomQuotaSettings {
	this := TenantRoomQuotaSettings{}
	return &this
}

// NewTenantRoomQuotaSettingsWithDefaults instantiates a new TenantRoomQuotaSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantRoomQuotaSettingsWithDefaults() *TenantRoomQuotaSettings {
	this := TenantRoomQuotaSettings{}
	return &this
}

// GetEnableQuota returns the EnableQuota field value if set, zero value otherwise.
func (o *TenantRoomQuotaSettings) GetEnableQuota() bool {
	if o == nil || IsNil(o.EnableQuota) {
		var ret bool
		return ret
	}
	return *o.EnableQuota
}

// GetEnableQuotaOk returns a tuple with the EnableQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRoomQuotaSettings) GetEnableQuotaOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableQuota) {
		return nil, false
	}
	return o.EnableQuota, true
}

// HasEnableQuota returns a boolean if a field has been set.
func (o *TenantRoomQuotaSettings) IsEnableQuotaSet() bool {
	if o != nil && !IsNil(o.EnableQuota) {
		return true
	}

	return false
}

// SetEnableQuota gets a reference to the given bool and assigns it to the EnableQuota field.
func (o *TenantRoomQuotaSettings) SetEnableQuota(v bool) {
	o.EnableQuota = &v
}

// GetDefaultQuota returns the DefaultQuota field value if set, zero value otherwise.
func (o *TenantRoomQuotaSettings) GetDefaultQuota() int64 {
	if o == nil || IsNil(o.DefaultQuota) {
		var ret int64
		return ret
	}
	return *o.DefaultQuota
}

// GetDefaultQuotaOk returns a tuple with the DefaultQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRoomQuotaSettings) GetDefaultQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultQuota) {
		return nil, false
	}
	return o.DefaultQuota, true
}

// HasDefaultQuota returns a boolean if a field has been set.
func (o *TenantRoomQuotaSettings) IsDefaultQuotaSet() bool {
	if o != nil && !IsNil(o.DefaultQuota) {
		return true
	}

	return false
}

// SetDefaultQuota gets a reference to the given int64 and assigns it to the DefaultQuota field.
func (o *TenantRoomQuotaSettings) SetDefaultQuota(v int64) {
	o.DefaultQuota = &v
}

// GetLastRecalculateDate returns the LastRecalculateDate field value if set, zero value otherwise.
func (o *TenantRoomQuotaSettings) GetLastRecalculateDate() time.Time {
	if o == nil || IsNil(o.LastRecalculateDate) {
		var ret time.Time
		return ret
	}
	return *o.LastRecalculateDate
}

// GetLastRecalculateDateOk returns a tuple with the LastRecalculateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRoomQuotaSettings) GetLastRecalculateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastRecalculateDate) {
		return nil, false
	}
	return o.LastRecalculateDate, true
}

// HasLastRecalculateDate returns a boolean if a field has been set.
func (o *TenantRoomQuotaSettings) IsLastRecalculateDateSet() bool {
	if o != nil && !IsNil(o.LastRecalculateDate) {
		return true
	}

	return false
}

// SetLastRecalculateDate gets a reference to the given time.Time and assigns it to the LastRecalculateDate field.
func (o *TenantRoomQuotaSettings) SetLastRecalculateDate(v time.Time) {
	o.LastRecalculateDate = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *TenantRoomQuotaSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRoomQuotaSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *TenantRoomQuotaSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *TenantRoomQuotaSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o TenantRoomQuotaSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantRoomQuotaSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableQuota) {
		toSerialize["enableQuota"] = o.EnableQuota
	}
	if !IsNil(o.DefaultQuota) {
		toSerialize["defaultQuota"] = o.DefaultQuota
	}
	if !IsNil(o.LastRecalculateDate) {
		toSerialize["lastRecalculateDate"] = o.LastRecalculateDate
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableTenantRoomQuotaSettings struct {
	value *TenantRoomQuotaSettings
	isSet bool
}

func (v NullableTenantRoomQuotaSettings) Get() *TenantRoomQuotaSettings {
	return v.value
}

func (v *NullableTenantRoomQuotaSettings) Set(val *TenantRoomQuotaSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantRoomQuotaSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantRoomQuotaSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantRoomQuotaSettings(val *TenantRoomQuotaSettings) *NullableTenantRoomQuotaSettings {
	return &NullableTenantRoomQuotaSettings{value: val, isSet: true}
}

func (v NullableTenantRoomQuotaSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantRoomQuotaSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

