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

// checks if the ApiDateTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDateTime{}

// ApiDateTime The API date and time parameters.
type ApiDateTime struct {
	// The time in UTC format.
	UtcTime *time.Time `json:"utcTime,omitempty"`
	// The time zone offset.
	TimeZoneOffset *string `json:"timeZoneOffset,omitempty"`
}

// NewApiDateTime instantiates a new ApiDateTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDateTime() *ApiDateTime {
	this := ApiDateTime{}
	return &this
}

// NewApiDateTimeWithDefaults instantiates a new ApiDateTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDateTimeWithDefaults() *ApiDateTime {
	this := ApiDateTime{}
	return &this
}

// GetUtcTime returns the UtcTime field value if set, zero value otherwise.
func (o *ApiDateTime) GetUtcTime() time.Time {
	if o == nil || IsNil(o.UtcTime) {
		var ret time.Time
		return ret
	}
	return *o.UtcTime
}

// GetUtcTimeOk returns a tuple with the UtcTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDateTime) GetUtcTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UtcTime) {
		return nil, false
	}
	return o.UtcTime, true
}

// HasUtcTime returns a boolean if a field has been set.
func (o *ApiDateTime) IsUtcTimeSet() bool {
	if o != nil && !IsNil(o.UtcTime) {
		return true
	}

	return false
}

// SetUtcTime gets a reference to the given time.Time and assigns it to the UtcTime field.
func (o *ApiDateTime) SetUtcTime(v time.Time) {
	o.UtcTime = &v
}

// GetTimeZoneOffset returns the TimeZoneOffset field value if set, zero value otherwise.
func (o *ApiDateTime) GetTimeZoneOffset() string {
	if o == nil || IsNil(o.TimeZoneOffset) {
		var ret string
		return ret
	}
	return *o.TimeZoneOffset
}

// GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDateTime) GetTimeZoneOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZoneOffset) {
		return nil, false
	}
	return o.TimeZoneOffset, true
}

// HasTimeZoneOffset returns a boolean if a field has been set.
func (o *ApiDateTime) IsTimeZoneOffsetSet() bool {
	if o != nil && !IsNil(o.TimeZoneOffset) {
		return true
	}

	return false
}

// SetTimeZoneOffset gets a reference to the given string and assigns it to the TimeZoneOffset field.
func (o *ApiDateTime) SetTimeZoneOffset(v string) {
	o.TimeZoneOffset = &v
}

func (o *ApiDateTime) UnmarshalJSON(data []byte) error {
	// Some DocSpace endpoints return datetime as plain string instead of object.
	if string(data) == "null" {
		return nil
	}

	var asString string
	if err := json.Unmarshal(data, &asString); err == nil {
		if asString == "" {
			return nil
		}

		layouts := []string{
			time.RFC3339Nano,
			time.RFC3339,
			"2006-01-02 15:04:05",
			"2006-01-02T15:04:05",
		}

		for _, layout := range layouts {
			if parsed, parseErr := time.Parse(layout, asString); parseErr == nil {
				o.UtcTime = &parsed
				return nil
			}
		}

		// Keep backward compatibility even for unknown string formats.
		return nil
	}

	// Fallback to object format: {"utcTime":"...","timeZoneOffset":"..."}
	type alias ApiDateTime
	var aux alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	*o = ApiDateTime(aux)
	return nil
}

func (o ApiDateTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDateTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UtcTime) {
		toSerialize["utcTime"] = o.UtcTime
	}
	if !IsNil(o.TimeZoneOffset) {
		toSerialize["timeZoneOffset"] = o.TimeZoneOffset
	}
	return toSerialize, nil
}

type NullableApiDateTime struct {
	value *ApiDateTime
	isSet bool
}

func (v NullableApiDateTime) Get() *ApiDateTime {
	return v.value
}

func (v *NullableApiDateTime) Set(val *ApiDateTime) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDateTime) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDateTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDateTime(val *ApiDateTime) *NullableApiDateTime {
	return &NullableApiDateTime{value: val, isSet: true}
}

func (v NullableApiDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDateTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

