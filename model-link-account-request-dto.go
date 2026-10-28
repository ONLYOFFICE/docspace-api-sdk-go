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
)

// checks if the LinkAccountRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkAccountRequestDto{}

// LinkAccountRequestDto The request parameters for linking accounts.
type LinkAccountRequestDto struct {
	// The profile a completed provider authorization produced, in the serialized form the login flow hands back.  Pass that value unchanged; it carries the provider, the third-party account ID and the authorization result,  and a hand-written object is not accepted.
	SerializedProfile NullableString `json:"serializedProfile,omitempty"`
}

// NewLinkAccountRequestDto instantiates a new LinkAccountRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkAccountRequestDto() *LinkAccountRequestDto {
	this := LinkAccountRequestDto{}
	return &this
}

// NewLinkAccountRequestDtoWithDefaults instantiates a new LinkAccountRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkAccountRequestDtoWithDefaults() *LinkAccountRequestDto {
	this := LinkAccountRequestDto{}
	return &this
}

// GetSerializedProfile returns the SerializedProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkAccountRequestDto) GetSerializedProfile() string {
	if o == nil || IsNil(o.SerializedProfile.Get()) {
		var ret string
		return ret
	}
	return *o.SerializedProfile.Get()
}

// GetSerializedProfileOk returns a tuple with the SerializedProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkAccountRequestDto) GetSerializedProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SerializedProfile.Get(), o.SerializedProfile.IsSet()
}

// HasSerializedProfile returns a boolean if a field has been set.
func (o *LinkAccountRequestDto) IsSerializedProfileSet() bool {
	if o != nil && o.SerializedProfile.IsSet() {
		return true
	}

	return false
}

// SetSerializedProfile gets a reference to the given NullableString and assigns it to the SerializedProfile field.
func (o *LinkAccountRequestDto) SetSerializedProfile(v string) {
	o.SerializedProfile.Set(&v)
}
// SetSerializedProfileNil sets the value for SerializedProfile to be an explicit nil
func (o *LinkAccountRequestDto) SetSerializedProfileNil() {
	o.SerializedProfile.Set(nil)
}

// UnsetSerializedProfile ensures that no value is present for SerializedProfile, not even an explicit nil
func (o *LinkAccountRequestDto) UnsetSerializedProfile() {
	o.SerializedProfile.Unset()
}

func (o LinkAccountRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkAccountRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SerializedProfile.IsSet() {
		toSerialize["serializedProfile"] = o.SerializedProfile.Get()
	}
	return toSerialize, nil
}

type NullableLinkAccountRequestDto struct {
	value *LinkAccountRequestDto
	isSet bool
}

func (v NullableLinkAccountRequestDto) Get() *LinkAccountRequestDto {
	return v.value
}

func (v *NullableLinkAccountRequestDto) Set(val *LinkAccountRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkAccountRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkAccountRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkAccountRequestDto(val *LinkAccountRequestDto) *NullableLinkAccountRequestDto {
	return &NullableLinkAccountRequestDto{value: val, isSet: true}
}

func (v NullableLinkAccountRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkAccountRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

