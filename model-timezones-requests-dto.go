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

// checks if the TimezonesRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimezonesRequestsDto{}

// TimezonesRequestsDto One time zone the host offers, as its identifier and the label to show for it.
type TimezonesRequestsDto struct {
	// The IANA identifier of the time zone. This is the value the portal time zone is set to, so pass it on  unchanged to `PUT api/2.0/settings/timeandlanguage`.
	Id NullableString `json:"id"`
	// The label to show for the zone, carrying its UTC offset as it stood when the list was built. The offset is a  snapshot rather than a rule, so a zone observing daylight saving reads differently at other times of the  year; sort and match on `id` instead.
	DisplayName NullableString `json:"displayName"`
}

type _TimezonesRequestsDto TimezonesRequestsDto

// NewTimezonesRequestsDto instantiates a new TimezonesRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimezonesRequestsDto(id NullableString, displayName NullableString) *TimezonesRequestsDto {
	this := TimezonesRequestsDto{}
	this.Id = id
	this.DisplayName = displayName
	return &this
}

// NewTimezonesRequestsDtoWithDefaults instantiates a new TimezonesRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimezonesRequestsDtoWithDefaults() *TimezonesRequestsDto {
	this := TimezonesRequestsDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimezonesRequestsDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimezonesRequestsDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *TimezonesRequestsDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetDisplayName returns the DisplayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimezonesRequestsDto) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimezonesRequestsDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// SetDisplayName sets field value
func (o *TimezonesRequestsDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

func (o TimezonesRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimezonesRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["displayName"] = o.DisplayName.Get()
	return toSerialize, nil
}

func (o *TimezonesRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"displayName",
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

	varTimezonesRequestsDto := _TimezonesRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTimezonesRequestsDto)

	if err != nil {
		return err
	}

	*o = TimezonesRequestsDto(varTimezonesRequestsDto)

	return err
}

type NullableTimezonesRequestsDto struct {
	value *TimezonesRequestsDto
	isSet bool
}

func (v NullableTimezonesRequestsDto) Get() *TimezonesRequestsDto {
	return v.value
}

func (v *NullableTimezonesRequestsDto) Set(val *TimezonesRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTimezonesRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTimezonesRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimezonesRequestsDto(val *TimezonesRequestsDto) *NullableTimezonesRequestsDto {
	return &NullableTimezonesRequestsDto{value: val, isSet: true}
}

func (v NullableTimezonesRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimezonesRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

