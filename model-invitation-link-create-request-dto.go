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

// checks if the InvitationLinkCreateRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitationLinkCreateRequestDto{}

// InvitationLinkCreateRequestDto The role a new invitation link grants, and the limits placed on it.
type InvitationLinkCreateRequestDto struct {
	// The role whoever follows the link joins with. Only `DocSpaceAdmin`, `RoomAdmin` and `User` are accepted, and  the role cannot be changed afterwards - delete the link and create one for the other role instead.
	EmployeeType EmployeeType `json:"employeeType"`
	// When the link stops letting anyone in, read in the portal time zone. It has to lie in the future; leaving it  out creates a link with no deadline at all.
	Expiration NullableTime `json:"expiration,omitempty"`
	// How many accounts may join through the link in total. Leaving it out creates a link with no use limit; the  uses spent so far are reported as `currentUseCount`.
	MaxUseCount NullableInt32 `json:"maxUseCount,omitempty"`
}

type _InvitationLinkCreateRequestDto InvitationLinkCreateRequestDto

// NewInvitationLinkCreateRequestDto instantiates a new InvitationLinkCreateRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationLinkCreateRequestDto(employeeType EmployeeType) *InvitationLinkCreateRequestDto {
	this := InvitationLinkCreateRequestDto{}
	this.EmployeeType = employeeType
	return &this
}

// NewInvitationLinkCreateRequestDtoWithDefaults instantiates a new InvitationLinkCreateRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationLinkCreateRequestDtoWithDefaults() *InvitationLinkCreateRequestDto {
	this := InvitationLinkCreateRequestDto{}
	return &this
}

// GetEmployeeType returns the EmployeeType field value
func (o *InvitationLinkCreateRequestDto) GetEmployeeType() EmployeeType {
	if o == nil {
		var ret EmployeeType
		return ret
	}

	return o.EmployeeType
}

// GetEmployeeTypeOk returns a tuple with the EmployeeType field value
// and a boolean to check if the value has been set.
func (o *InvitationLinkCreateRequestDto) GetEmployeeTypeOk() (*EmployeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmployeeType, true
}

// SetEmployeeType sets field value
func (o *InvitationLinkCreateRequestDto) SetEmployeeType(v EmployeeType) {
	o.EmployeeType = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationLinkCreateRequestDto) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationLinkCreateRequestDto) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *InvitationLinkCreateRequestDto) IsExpirationSet() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *InvitationLinkCreateRequestDto) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}
// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *InvitationLinkCreateRequestDto) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *InvitationLinkCreateRequestDto) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetMaxUseCount returns the MaxUseCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationLinkCreateRequestDto) GetMaxUseCount() int32 {
	if o == nil || IsNil(o.MaxUseCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxUseCount.Get()
}

// GetMaxUseCountOk returns a tuple with the MaxUseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationLinkCreateRequestDto) GetMaxUseCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxUseCount.Get(), o.MaxUseCount.IsSet()
}

// HasMaxUseCount returns a boolean if a field has been set.
func (o *InvitationLinkCreateRequestDto) IsMaxUseCountSet() bool {
	if o != nil && o.MaxUseCount.IsSet() {
		return true
	}

	return false
}

// SetMaxUseCount gets a reference to the given NullableInt32 and assigns it to the MaxUseCount field.
func (o *InvitationLinkCreateRequestDto) SetMaxUseCount(v int32) {
	o.MaxUseCount.Set(&v)
}
// SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil
func (o *InvitationLinkCreateRequestDto) SetMaxUseCountNil() {
	o.MaxUseCount.Set(nil)
}

// UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil
func (o *InvitationLinkCreateRequestDto) UnsetMaxUseCount() {
	o.MaxUseCount.Unset()
}

func (o InvitationLinkCreateRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitationLinkCreateRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["employeeType"] = o.EmployeeType
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if o.MaxUseCount.IsSet() {
		toSerialize["maxUseCount"] = o.MaxUseCount.Get()
	}
	return toSerialize, nil
}

func (o *InvitationLinkCreateRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"employeeType",
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

	varInvitationLinkCreateRequestDto := _InvitationLinkCreateRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInvitationLinkCreateRequestDto)

	if err != nil {
		return err
	}

	*o = InvitationLinkCreateRequestDto(varInvitationLinkCreateRequestDto)

	return err
}

type NullableInvitationLinkCreateRequestDto struct {
	value *InvitationLinkCreateRequestDto
	isSet bool
}

func (v NullableInvitationLinkCreateRequestDto) Get() *InvitationLinkCreateRequestDto {
	return v.value
}

func (v *NullableInvitationLinkCreateRequestDto) Set(val *InvitationLinkCreateRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationLinkCreateRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationLinkCreateRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationLinkCreateRequestDto(val *InvitationLinkCreateRequestDto) *NullableInvitationLinkCreateRequestDto {
	return &NullableInvitationLinkCreateRequestDto{value: val, isSet: true}
}

func (v NullableInvitationLinkCreateRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationLinkCreateRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

