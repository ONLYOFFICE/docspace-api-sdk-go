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

// checks if the AiApiDateTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiApiDateTime{}

// AiApiDateTime The API date and time parameters.
type AiApiDateTime struct {
	// The time in UTC format.
	UtcTime *time.Time `json:"utcTime,omitempty"`
	// The time zone offset.
	TimeZoneOffset *string `json:"timeZoneOffset,omitempty"`
}

// NewAiApiDateTime instantiates a new AiApiDateTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiApiDateTime() *AiApiDateTime {
	this := AiApiDateTime{}
	return &this
}

// NewAiApiDateTimeWithDefaults instantiates a new AiApiDateTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiApiDateTimeWithDefaults() *AiApiDateTime {
	this := AiApiDateTime{}
	return &this
}

// GetUtcTime returns the UtcTime field value if set, zero value otherwise.
func (o *AiApiDateTime) GetUtcTime() time.Time {
	if o == nil || IsNil(o.UtcTime) {
		var ret time.Time
		return ret
	}
	return *o.UtcTime
}

// GetUtcTimeOk returns a tuple with the UtcTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiApiDateTime) GetUtcTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UtcTime) {
		return nil, false
	}
	return o.UtcTime, true
}

// HasUtcTime returns a boolean if a field has been set.
func (o *AiApiDateTime) IsUtcTimeSet() bool {
	if o != nil && !IsNil(o.UtcTime) {
		return true
	}

	return false
}

// SetUtcTime gets a reference to the given time.Time and assigns it to the UtcTime field.
func (o *AiApiDateTime) SetUtcTime(v time.Time) {
	o.UtcTime = &v
}

// GetTimeZoneOffset returns the TimeZoneOffset field value if set, zero value otherwise.
func (o *AiApiDateTime) GetTimeZoneOffset() string {
	if o == nil || IsNil(o.TimeZoneOffset) {
		var ret string
		return ret
	}
	return *o.TimeZoneOffset
}

// GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiApiDateTime) GetTimeZoneOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZoneOffset) {
		return nil, false
	}
	return o.TimeZoneOffset, true
}

// HasTimeZoneOffset returns a boolean if a field has been set.
func (o *AiApiDateTime) IsTimeZoneOffsetSet() bool {
	if o != nil && !IsNil(o.TimeZoneOffset) {
		return true
	}

	return false
}

// SetTimeZoneOffset gets a reference to the given string and assigns it to the TimeZoneOffset field.
func (o *AiApiDateTime) SetTimeZoneOffset(v string) {
	o.TimeZoneOffset = &v
}

func (o AiApiDateTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiApiDateTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UtcTime) {
		toSerialize["utcTime"] = o.UtcTime
	}
	if !IsNil(o.TimeZoneOffset) {
		toSerialize["timeZoneOffset"] = o.TimeZoneOffset
	}
	return toSerialize, nil
}

type NullableAiApiDateTime struct {
	value *AiApiDateTime
	isSet bool
}

func (v NullableAiApiDateTime) Get() *AiApiDateTime {
	return v.value
}

func (v *NullableAiApiDateTime) Set(val *AiApiDateTime) {
	v.value = val
	v.isSet = true
}

func (v NullableAiApiDateTime) IsSet() bool {
	return v.isSet
}

func (v *NullableAiApiDateTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiApiDateTime(val *AiApiDateTime) *NullableAiApiDateTime {
	return &NullableAiApiDateTime{value: val, isSet: true}
}

func (v NullableAiApiDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiApiDateTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

