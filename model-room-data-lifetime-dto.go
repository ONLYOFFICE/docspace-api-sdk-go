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

// checks if the RoomDataLifetimeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomDataLifetimeDto{}

// RoomDataLifetimeDto The room data lifetime information.
type RoomDataLifetimeDto struct {
	// Specifies whether to permanently delete the room data or not.
	DeletePermanently *bool `json:"deletePermanently,omitempty"`
	Period *RoomDataLifetimePeriod `json:"period,omitempty"`
	// Specifies the time period value of the room data lifetime.
	Value NullableInt32 `json:"value,omitempty"`
	// Specifies whether the room data lifetime setting is enabled or not.
	Enabled NullableBool `json:"enabled,omitempty"`
}

// NewRoomDataLifetimeDto instantiates a new RoomDataLifetimeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomDataLifetimeDto() *RoomDataLifetimeDto {
	this := RoomDataLifetimeDto{}
	return &this
}

// NewRoomDataLifetimeDtoWithDefaults instantiates a new RoomDataLifetimeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomDataLifetimeDtoWithDefaults() *RoomDataLifetimeDto {
	this := RoomDataLifetimeDto{}
	return &this
}

// GetDeletePermanently returns the DeletePermanently field value if set, zero value otherwise.
func (o *RoomDataLifetimeDto) GetDeletePermanently() bool {
	if o == nil || IsNil(o.DeletePermanently) {
		var ret bool
		return ret
	}
	return *o.DeletePermanently
}

// GetDeletePermanentlyOk returns a tuple with the DeletePermanently field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomDataLifetimeDto) GetDeletePermanentlyOk() (*bool, bool) {
	if o == nil || IsNil(o.DeletePermanently) {
		return nil, false
	}
	return o.DeletePermanently, true
}

// HasDeletePermanently returns a boolean if a field has been set.
func (o *RoomDataLifetimeDto) IsDeletePermanentlySet() bool {
	if o != nil && !IsNil(o.DeletePermanently) {
		return true
	}

	return false
}

// SetDeletePermanently gets a reference to the given bool and assigns it to the DeletePermanently field.
func (o *RoomDataLifetimeDto) SetDeletePermanently(v bool) {
	o.DeletePermanently = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *RoomDataLifetimeDto) GetPeriod() RoomDataLifetimePeriod {
	if o == nil || IsNil(o.Period) {
		var ret RoomDataLifetimePeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomDataLifetimeDto) GetPeriodOk() (*RoomDataLifetimePeriod, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *RoomDataLifetimeDto) IsPeriodSet() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given RoomDataLifetimePeriod and assigns it to the Period field.
func (o *RoomDataLifetimeDto) SetPeriod(v RoomDataLifetimePeriod) {
	o.Period = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomDataLifetimeDto) GetValue() int32 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret int32
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomDataLifetimeDto) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *RoomDataLifetimeDto) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableInt32 and assigns it to the Value field.
func (o *RoomDataLifetimeDto) SetValue(v int32) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *RoomDataLifetimeDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *RoomDataLifetimeDto) UnsetValue() {
	o.Value.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomDataLifetimeDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomDataLifetimeDto) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *RoomDataLifetimeDto) IsEnabledSet() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *RoomDataLifetimeDto) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *RoomDataLifetimeDto) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *RoomDataLifetimeDto) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o RoomDataLifetimeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomDataLifetimeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeletePermanently) {
		toSerialize["deletePermanently"] = o.DeletePermanently
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

type NullableRoomDataLifetimeDto struct {
	value *RoomDataLifetimeDto
	isSet bool
}

func (v NullableRoomDataLifetimeDto) Get() *RoomDataLifetimeDto {
	return v.value
}

func (v *NullableRoomDataLifetimeDto) Set(val *RoomDataLifetimeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomDataLifetimeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomDataLifetimeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomDataLifetimeDto(val *RoomDataLifetimeDto) *NullableRoomDataLifetimeDto {
	return &NullableRoomDataLifetimeDto{value: val, isSet: true}
}

func (v NullableRoomDataLifetimeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomDataLifetimeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

