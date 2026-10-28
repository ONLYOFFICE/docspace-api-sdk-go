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

// checks if the InvitationLinkUpdateRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitationLinkUpdateRequestDto{}

// InvitationLinkUpdateRequestDto The invitation link being changed, with the deadline and use limit it is to have afterwards.
type InvitationLinkUpdateRequestDto struct {
	// The link to change, by the `id` that creating or reading it returned. The role behind that id cannot be  changed here.
	Id string `json:"id"`
	// The new deadline, read in the portal time zone. The body is applied as a whole, so leaving it out clears the  deadline rather than keeping the current one; a moment in the past is refused.
	Expiration NullableTime `json:"expiration,omitempty"`
	// The new total number of accounts that may join through the link. It may not be lower than the uses already  spent, which the link reports as `currentUseCount`, and leaving it out removes the limit rather than keeping  the current one.
	MaxUseCount NullableInt32 `json:"maxUseCount,omitempty"`
}

type _InvitationLinkUpdateRequestDto InvitationLinkUpdateRequestDto

// NewInvitationLinkUpdateRequestDto instantiates a new InvitationLinkUpdateRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationLinkUpdateRequestDto(id string) *InvitationLinkUpdateRequestDto {
	this := InvitationLinkUpdateRequestDto{}
	this.Id = id
	return &this
}

// NewInvitationLinkUpdateRequestDtoWithDefaults instantiates a new InvitationLinkUpdateRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationLinkUpdateRequestDtoWithDefaults() *InvitationLinkUpdateRequestDto {
	this := InvitationLinkUpdateRequestDto{}
	return &this
}

// GetId returns the Id field value
func (o *InvitationLinkUpdateRequestDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvitationLinkUpdateRequestDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvitationLinkUpdateRequestDto) SetId(v string) {
	o.Id = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationLinkUpdateRequestDto) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationLinkUpdateRequestDto) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *InvitationLinkUpdateRequestDto) IsExpirationSet() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *InvitationLinkUpdateRequestDto) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}
// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *InvitationLinkUpdateRequestDto) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *InvitationLinkUpdateRequestDto) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetMaxUseCount returns the MaxUseCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationLinkUpdateRequestDto) GetMaxUseCount() int32 {
	if o == nil || IsNil(o.MaxUseCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxUseCount.Get()
}

// GetMaxUseCountOk returns a tuple with the MaxUseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationLinkUpdateRequestDto) GetMaxUseCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxUseCount.Get(), o.MaxUseCount.IsSet()
}

// HasMaxUseCount returns a boolean if a field has been set.
func (o *InvitationLinkUpdateRequestDto) IsMaxUseCountSet() bool {
	if o != nil && o.MaxUseCount.IsSet() {
		return true
	}

	return false
}

// SetMaxUseCount gets a reference to the given NullableInt32 and assigns it to the MaxUseCount field.
func (o *InvitationLinkUpdateRequestDto) SetMaxUseCount(v int32) {
	o.MaxUseCount.Set(&v)
}
// SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil
func (o *InvitationLinkUpdateRequestDto) SetMaxUseCountNil() {
	o.MaxUseCount.Set(nil)
}

// UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil
func (o *InvitationLinkUpdateRequestDto) UnsetMaxUseCount() {
	o.MaxUseCount.Unset()
}

func (o InvitationLinkUpdateRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitationLinkUpdateRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if o.MaxUseCount.IsSet() {
		toSerialize["maxUseCount"] = o.MaxUseCount.Get()
	}
	return toSerialize, nil
}

func (o *InvitationLinkUpdateRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varInvitationLinkUpdateRequestDto := _InvitationLinkUpdateRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInvitationLinkUpdateRequestDto)

	if err != nil {
		return err
	}

	*o = InvitationLinkUpdateRequestDto(varInvitationLinkUpdateRequestDto)

	return err
}

type NullableInvitationLinkUpdateRequestDto struct {
	value *InvitationLinkUpdateRequestDto
	isSet bool
}

func (v NullableInvitationLinkUpdateRequestDto) Get() *InvitationLinkUpdateRequestDto {
	return v.value
}

func (v *NullableInvitationLinkUpdateRequestDto) Set(val *InvitationLinkUpdateRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationLinkUpdateRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationLinkUpdateRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationLinkUpdateRequestDto(val *InvitationLinkUpdateRequestDto) *NullableInvitationLinkUpdateRequestDto {
	return &NullableInvitationLinkUpdateRequestDto{value: val, isSet: true}
}

func (v NullableInvitationLinkUpdateRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationLinkUpdateRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

