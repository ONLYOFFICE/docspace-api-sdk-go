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

// checks if the AiProfilesTestConnection200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProfilesTestConnection200ResponseAnyOf{}

// AiProfilesTestConnection200ResponseAnyOf struct for AiProfilesTestConnection200ResponseAnyOf
type AiProfilesTestConnection200ResponseAnyOf struct {
	Message *string `json:"message,omitempty"`
}

// NewAiProfilesTestConnection200ResponseAnyOf instantiates a new AiProfilesTestConnection200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProfilesTestConnection200ResponseAnyOf() *AiProfilesTestConnection200ResponseAnyOf {
	this := AiProfilesTestConnection200ResponseAnyOf{}
	return &this
}

// NewAiProfilesTestConnection200ResponseAnyOfWithDefaults instantiates a new AiProfilesTestConnection200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProfilesTestConnection200ResponseAnyOfWithDefaults() *AiProfilesTestConnection200ResponseAnyOf {
	this := AiProfilesTestConnection200ResponseAnyOf{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AiProfilesTestConnection200ResponseAnyOf) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesTestConnection200ResponseAnyOf) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AiProfilesTestConnection200ResponseAnyOf) IsMessageSet() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AiProfilesTestConnection200ResponseAnyOf) SetMessage(v string) {
	o.Message = &v
}

func (o AiProfilesTestConnection200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProfilesTestConnection200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableAiProfilesTestConnection200ResponseAnyOf struct {
	value *AiProfilesTestConnection200ResponseAnyOf
	isSet bool
}

func (v NullableAiProfilesTestConnection200ResponseAnyOf) Get() *AiProfilesTestConnection200ResponseAnyOf {
	return v.value
}

func (v *NullableAiProfilesTestConnection200ResponseAnyOf) Set(val *AiProfilesTestConnection200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProfilesTestConnection200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProfilesTestConnection200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProfilesTestConnection200ResponseAnyOf(val *AiProfilesTestConnection200ResponseAnyOf) *NullableAiProfilesTestConnection200ResponseAnyOf {
	return &NullableAiProfilesTestConnection200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableAiProfilesTestConnection200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProfilesTestConnection200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

