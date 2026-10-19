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

// checks if the IpRestrictionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpRestrictionDto{}

// IpRestrictionDto The IP restiction parameters.
type IpRestrictionDto struct {
	// The IP restiction ID.
	Id *int32 `json:"id,omitempty"`
	// The tenant ID.
	TenantId *int32 `json:"tenantId,omitempty"`
	// The IP address.
	Ip NullableString `json:"ip"`
	// Specifies if the IP address is for administrator users only or not.
	ForAdmin *bool `json:"forAdmin,omitempty"`
}

type _IpRestrictionDto IpRestrictionDto

// NewIpRestrictionDto instantiates a new IpRestrictionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpRestrictionDto(ip NullableString) *IpRestrictionDto {
	this := IpRestrictionDto{}
	this.Ip = ip
	return &this
}

// NewIpRestrictionDtoWithDefaults instantiates a new IpRestrictionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpRestrictionDtoWithDefaults() *IpRestrictionDto {
	this := IpRestrictionDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpRestrictionDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpRestrictionDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpRestrictionDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IpRestrictionDto) SetId(v int32) {
	o.Id = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IpRestrictionDto) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpRestrictionDto) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IpRestrictionDto) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *IpRestrictionDto) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpRestrictionDto) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpRestrictionDto) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// SetIp sets field value
func (o *IpRestrictionDto) SetIp(v string) {
	o.Ip.Set(&v)
}

// GetForAdmin returns the ForAdmin field value if set, zero value otherwise.
func (o *IpRestrictionDto) GetForAdmin() bool {
	if o == nil || IsNil(o.ForAdmin) {
		var ret bool
		return ret
	}
	return *o.ForAdmin
}

// GetForAdminOk returns a tuple with the ForAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpRestrictionDto) GetForAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.ForAdmin) {
		return nil, false
	}
	return o.ForAdmin, true
}

// HasForAdmin returns a boolean if a field has been set.
func (o *IpRestrictionDto) IsForAdminSet() bool {
	if o != nil && !IsNil(o.ForAdmin) {
		return true
	}

	return false
}

// SetForAdmin gets a reference to the given bool and assigns it to the ForAdmin field.
func (o *IpRestrictionDto) SetForAdmin(v bool) {
	o.ForAdmin = &v
}

func (o IpRestrictionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpRestrictionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	toSerialize["ip"] = o.Ip.Get()
	if !IsNil(o.ForAdmin) {
		toSerialize["forAdmin"] = o.ForAdmin
	}
	return toSerialize, nil
}

func (o *IpRestrictionDto) UnmarshalJSON(data []byte) (err error) {
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

	varIpRestrictionDto := _IpRestrictionDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varIpRestrictionDto)

	if err != nil {
		return err
	}

	*o = IpRestrictionDto(varIpRestrictionDto)

	return err
}

type NullableIpRestrictionDto struct {
	value *IpRestrictionDto
	isSet bool
}

func (v NullableIpRestrictionDto) Get() *IpRestrictionDto {
	return v.value
}

func (v *NullableIpRestrictionDto) Set(val *IpRestrictionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIpRestrictionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIpRestrictionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpRestrictionDto(val *IpRestrictionDto) *NullableIpRestrictionDto {
	return &NullableIpRestrictionDto{value: val, isSet: true}
}

func (v NullableIpRestrictionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpRestrictionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

