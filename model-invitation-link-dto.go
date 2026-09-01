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

// checks if the InvitationLinkDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitationLinkDto{}

// InvitationLinkDto The invitation link parameters.
type InvitationLinkDto struct {
	// The ID of the invitation link.
	Id *string `json:"id,omitempty"`
	// The type of employee role for the invitation link.
	EmployeeType EmployeeType `json:"employeeType"`
	// The expiration date of the invitation link.
	Expiration NullableTime `json:"expiration,omitempty"`
	// Indicates whether the invitation link has expired.
	IsExpired *bool `json:"isExpired,omitempty"`
	// The maximum number of times the invitation link can be used.
	MaxUseCount NullableInt32 `json:"maxUseCount,omitempty"`
	// The current number of times the invitation link has been used.
	CurrentUseCount *int32 `json:"currentUseCount,omitempty"`
	// The URL of the invitation link.
	Url NullableString `json:"url,omitempty"`
}

type _InvitationLinkDto InvitationLinkDto

// NewInvitationLinkDto instantiates a new InvitationLinkDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationLinkDto(employeeType EmployeeType) *InvitationLinkDto {
	this := InvitationLinkDto{}
	this.EmployeeType = employeeType
	return &this
}

// NewInvitationLinkDtoWithDefaults instantiates a new InvitationLinkDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationLinkDtoWithDefaults() *InvitationLinkDto {
	this := InvitationLinkDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvitationLinkDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationLinkDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvitationLinkDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InvitationLinkDto) SetId(v string) {
	o.Id = &v
}

// GetEmployeeType returns the EmployeeType field value
func (o *InvitationLinkDto) GetEmployeeType() EmployeeType {
	if o == nil {
		var ret EmployeeType
		return ret
	}

	return o.EmployeeType
}

// GetEmployeeTypeOk returns a tuple with the EmployeeType field value
// and a boolean to check if the value has been set.
func (o *InvitationLinkDto) GetEmployeeTypeOk() (*EmployeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmployeeType, true
}

// SetEmployeeType sets field value
func (o *InvitationLinkDto) SetEmployeeType(v EmployeeType) {
	o.EmployeeType = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationLinkDto) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationLinkDto) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *InvitationLinkDto) IsExpirationSet() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *InvitationLinkDto) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}
// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *InvitationLinkDto) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *InvitationLinkDto) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise.
func (o *InvitationLinkDto) GetIsExpired() bool {
	if o == nil || IsNil(o.IsExpired) {
		var ret bool
		return ret
	}
	return *o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationLinkDto) GetIsExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExpired) {
		return nil, false
	}
	return o.IsExpired, true
}

// HasIsExpired returns a boolean if a field has been set.
func (o *InvitationLinkDto) IsIsExpiredSet() bool {
	if o != nil && !IsNil(o.IsExpired) {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given bool and assigns it to the IsExpired field.
func (o *InvitationLinkDto) SetIsExpired(v bool) {
	o.IsExpired = &v
}

// GetMaxUseCount returns the MaxUseCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationLinkDto) GetMaxUseCount() int32 {
	if o == nil || IsNil(o.MaxUseCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxUseCount.Get()
}

// GetMaxUseCountOk returns a tuple with the MaxUseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationLinkDto) GetMaxUseCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxUseCount.Get(), o.MaxUseCount.IsSet()
}

// HasMaxUseCount returns a boolean if a field has been set.
func (o *InvitationLinkDto) IsMaxUseCountSet() bool {
	if o != nil && o.MaxUseCount.IsSet() {
		return true
	}

	return false
}

// SetMaxUseCount gets a reference to the given NullableInt32 and assigns it to the MaxUseCount field.
func (o *InvitationLinkDto) SetMaxUseCount(v int32) {
	o.MaxUseCount.Set(&v)
}
// SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil
func (o *InvitationLinkDto) SetMaxUseCountNil() {
	o.MaxUseCount.Set(nil)
}

// UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil
func (o *InvitationLinkDto) UnsetMaxUseCount() {
	o.MaxUseCount.Unset()
}

// GetCurrentUseCount returns the CurrentUseCount field value if set, zero value otherwise.
func (o *InvitationLinkDto) GetCurrentUseCount() int32 {
	if o == nil || IsNil(o.CurrentUseCount) {
		var ret int32
		return ret
	}
	return *o.CurrentUseCount
}

// GetCurrentUseCountOk returns a tuple with the CurrentUseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationLinkDto) GetCurrentUseCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentUseCount) {
		return nil, false
	}
	return o.CurrentUseCount, true
}

// HasCurrentUseCount returns a boolean if a field has been set.
func (o *InvitationLinkDto) IsCurrentUseCountSet() bool {
	if o != nil && !IsNil(o.CurrentUseCount) {
		return true
	}

	return false
}

// SetCurrentUseCount gets a reference to the given int32 and assigns it to the CurrentUseCount field.
func (o *InvitationLinkDto) SetCurrentUseCount(v int32) {
	o.CurrentUseCount = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationLinkDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationLinkDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *InvitationLinkDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *InvitationLinkDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *InvitationLinkDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *InvitationLinkDto) UnsetUrl() {
	o.Url.Unset()
}

func (o InvitationLinkDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitationLinkDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["employeeType"] = o.EmployeeType
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if !IsNil(o.IsExpired) {
		toSerialize["isExpired"] = o.IsExpired
	}
	if o.MaxUseCount.IsSet() {
		toSerialize["maxUseCount"] = o.MaxUseCount.Get()
	}
	if !IsNil(o.CurrentUseCount) {
		toSerialize["currentUseCount"] = o.CurrentUseCount
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

func (o *InvitationLinkDto) UnmarshalJSON(data []byte) (err error) {
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

	varInvitationLinkDto := _InvitationLinkDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInvitationLinkDto)

	if err != nil {
		return err
	}

	*o = InvitationLinkDto(varInvitationLinkDto)

	return err
}

type NullableInvitationLinkDto struct {
	value *InvitationLinkDto
	isSet bool
}

func (v NullableInvitationLinkDto) Get() *InvitationLinkDto {
	return v.value
}

func (v *NullableInvitationLinkDto) Set(val *InvitationLinkDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationLinkDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationLinkDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationLinkDto(val *InvitationLinkDto) *NullableInvitationLinkDto {
	return &NullableInvitationLinkDto{value: val, isSet: true}
}

func (v NullableInvitationLinkDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationLinkDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

