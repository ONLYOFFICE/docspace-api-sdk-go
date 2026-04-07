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

// checks if the IPRestriction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPRestriction{}

// IPRestriction struct for IPRestriction
type IPRestriction struct {
	Ip NullableString `json:"ip"`
	ForAdmin *bool `json:"forAdmin,omitempty"`
	Id *int32 `json:"id,omitempty"`
	TenantId *int32 `json:"tenantId,omitempty"`
}

type _IPRestriction IPRestriction

// NewIPRestriction instantiates a new IPRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPRestriction(ip NullableString) *IPRestriction {
	this := IPRestriction{}
	this.Ip = ip
	return &this
}

// NewIPRestrictionWithDefaults instantiates a new IPRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPRestrictionWithDefaults() *IPRestriction {
	this := IPRestriction{}
	return &this
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IPRestriction) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPRestriction) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// SetIp sets field value
func (o *IPRestriction) SetIp(v string) {
	o.Ip.Set(&v)
}

// GetForAdmin returns the ForAdmin field value if set, zero value otherwise.
func (o *IPRestriction) GetForAdmin() bool {
	if o == nil || IsNil(o.ForAdmin) {
		var ret bool
		return ret
	}
	return *o.ForAdmin
}

// GetForAdminOk returns a tuple with the ForAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRestriction) GetForAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.ForAdmin) {
		return nil, false
	}
	return o.ForAdmin, true
}

// HasForAdmin returns a boolean if a field has been set.
func (o *IPRestriction) IsForAdminSet() bool {
	if o != nil && !IsNil(o.ForAdmin) {
		return true
	}

	return false
}

// SetForAdmin gets a reference to the given bool and assigns it to the ForAdmin field.
func (o *IPRestriction) SetForAdmin(v bool) {
	o.ForAdmin = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IPRestriction) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRestriction) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IPRestriction) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IPRestriction) SetId(v int32) {
	o.Id = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IPRestriction) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRestriction) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IPRestriction) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *IPRestriction) SetTenantId(v int32) {
	o.TenantId = &v
}

func (o IPRestriction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip.Get()
	if !IsNil(o.ForAdmin) {
		toSerialize["forAdmin"] = o.ForAdmin
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	return toSerialize, nil
}

func (o *IPRestriction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
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

	varIPRestriction := _IPRestriction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varIPRestriction)

	if err != nil {
		return err
	}

	*o = IPRestriction(varIPRestriction)

	return err
}

type NullableIPRestriction struct {
	value *IPRestriction
	isSet bool
}

func (v NullableIPRestriction) Get() *IPRestriction {
	return v.value
}

func (v *NullableIPRestriction) Set(val *IPRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableIPRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableIPRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPRestriction(val *IPRestriction) *NullableIPRestriction {
	return &NullableIPRestriction{value: val, isSet: true}
}

func (v NullableIPRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

