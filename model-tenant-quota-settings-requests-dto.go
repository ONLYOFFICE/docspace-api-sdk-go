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

// checks if the TenantQuotaSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantQuotaSettingsRequestsDto{}

// TenantQuotaSettingsRequestsDto The storage limit set on one tenant of a self-hosted installation.
type TenantQuotaSettingsRequestsDto struct {
	// The tenant the limit applies to, by tenant ID. Only a self-hosted installation has more than one, which is  why the operation is refused on SaaS.
	TenantId int32 `json:"tenantId"`
	// The limit in bytes. A negative value is not a smaller limit but the absence of one: it removes whatever limit  the tenant had. The value is a ceiling on stored data and says nothing about how much of it is already used.
	Quota *int64 `json:"quota,omitempty"`
}

type _TenantQuotaSettingsRequestsDto TenantQuotaSettingsRequestsDto

// NewTenantQuotaSettingsRequestsDto instantiates a new TenantQuotaSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantQuotaSettingsRequestsDto(tenantId int32) *TenantQuotaSettingsRequestsDto {
	this := TenantQuotaSettingsRequestsDto{}
	this.TenantId = tenantId
	return &this
}

// NewTenantQuotaSettingsRequestsDtoWithDefaults instantiates a new TenantQuotaSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantQuotaSettingsRequestsDtoWithDefaults() *TenantQuotaSettingsRequestsDto {
	this := TenantQuotaSettingsRequestsDto{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *TenantQuotaSettingsRequestsDto) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *TenantQuotaSettingsRequestsDto) GetTenantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *TenantQuotaSettingsRequestsDto) SetTenantId(v int32) {
	o.TenantId = v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *TenantQuotaSettingsRequestsDto) GetQuota() int64 {
	if o == nil || IsNil(o.Quota) {
		var ret int64
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaSettingsRequestsDto) GetQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *TenantQuotaSettingsRequestsDto) IsQuotaSet() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int64 and assigns it to the Quota field.
func (o *TenantQuotaSettingsRequestsDto) SetQuota(v int64) {
	o.Quota = &v
}

func (o TenantQuotaSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantQuotaSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	return toSerialize, nil
}

func (o *TenantQuotaSettingsRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
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

	varTenantQuotaSettingsRequestsDto := _TenantQuotaSettingsRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTenantQuotaSettingsRequestsDto)

	if err != nil {
		return err
	}

	*o = TenantQuotaSettingsRequestsDto(varTenantQuotaSettingsRequestsDto)

	return err
}

type NullableTenantQuotaSettingsRequestsDto struct {
	value *TenantQuotaSettingsRequestsDto
	isSet bool
}

func (v NullableTenantQuotaSettingsRequestsDto) Get() *TenantQuotaSettingsRequestsDto {
	return v.value
}

func (v *NullableTenantQuotaSettingsRequestsDto) Set(val *TenantQuotaSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantQuotaSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantQuotaSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantQuotaSettingsRequestsDto(val *TenantQuotaSettingsRequestsDto) *NullableTenantQuotaSettingsRequestsDto {
	return &NullableTenantQuotaSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableTenantQuotaSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantQuotaSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

