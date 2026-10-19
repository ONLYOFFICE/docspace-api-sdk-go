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

// checks if the EmailActivationSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailActivationSettingsRequestDto{}

// EmailActivationSettingsRequestDto Whether the calling user wants to keep seeing the reminder to confirm their email address.
type EmailActivationSettingsRequestDto struct {
	// Whether the reminder is shown; send false to dismiss it for the calling user.
	Show *bool `json:"show,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewEmailActivationSettingsRequestDto instantiates a new EmailActivationSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailActivationSettingsRequestDto() *EmailActivationSettingsRequestDto {
	this := EmailActivationSettingsRequestDto{}
	return &this
}

// NewEmailActivationSettingsRequestDtoWithDefaults instantiates a new EmailActivationSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailActivationSettingsRequestDtoWithDefaults() *EmailActivationSettingsRequestDto {
	this := EmailActivationSettingsRequestDto{}
	return &this
}

// GetShow returns the Show field value if set, zero value otherwise.
func (o *EmailActivationSettingsRequestDto) GetShow() bool {
	if o == nil || IsNil(o.Show) {
		var ret bool
		return ret
	}
	return *o.Show
}

// GetShowOk returns a tuple with the Show field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailActivationSettingsRequestDto) GetShowOk() (*bool, bool) {
	if o == nil || IsNil(o.Show) {
		return nil, false
	}
	return o.Show, true
}

// HasShow returns a boolean if a field has been set.
func (o *EmailActivationSettingsRequestDto) IsShowSet() bool {
	if o != nil && !IsNil(o.Show) {
		return true
	}

	return false
}

// SetShow gets a reference to the given bool and assigns it to the Show field.
func (o *EmailActivationSettingsRequestDto) SetShow(v bool) {
	o.Show = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *EmailActivationSettingsRequestDto) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailActivationSettingsRequestDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *EmailActivationSettingsRequestDto) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *EmailActivationSettingsRequestDto) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o EmailActivationSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailActivationSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Show) {
		toSerialize["show"] = o.Show
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableEmailActivationSettingsRequestDto struct {
	value *EmailActivationSettingsRequestDto
	isSet bool
}

func (v NullableEmailActivationSettingsRequestDto) Get() *EmailActivationSettingsRequestDto {
	return v.value
}

func (v *NullableEmailActivationSettingsRequestDto) Set(val *EmailActivationSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailActivationSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailActivationSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailActivationSettingsRequestDto(val *EmailActivationSettingsRequestDto) *NullableEmailActivationSettingsRequestDto {
	return &NullableEmailActivationSettingsRequestDto{value: val, isSet: true}
}

func (v NullableEmailActivationSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailActivationSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

