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

// checks if the AnonymousConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnonymousConfigDto{}

// AnonymousConfigDto How the editors treat a participant who opened the document without an account.
type AnonymousConfigDto struct {
	// Whether the editors ask an anonymous participant for a display name before letting them in. It follows the  chat permission of the document, since a nameless participant cannot take part in one.
	Request bool `json:"request"`
}

type _AnonymousConfigDto AnonymousConfigDto

// NewAnonymousConfigDto instantiates a new AnonymousConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnonymousConfigDto(request bool) *AnonymousConfigDto {
	this := AnonymousConfigDto{}
	this.Request = request
	return &this
}

// NewAnonymousConfigDtoWithDefaults instantiates a new AnonymousConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnonymousConfigDtoWithDefaults() *AnonymousConfigDto {
	this := AnonymousConfigDto{}
	return &this
}

// GetRequest returns the Request field value
func (o *AnonymousConfigDto) GetRequest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *AnonymousConfigDto) GetRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *AnonymousConfigDto) SetRequest(v bool) {
	o.Request = v
}

func (o AnonymousConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnonymousConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request"] = o.Request
	return toSerialize, nil
}

func (o *AnonymousConfigDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request",
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

	varAnonymousConfigDto := _AnonymousConfigDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAnonymousConfigDto)

	if err != nil {
		return err
	}

	*o = AnonymousConfigDto(varAnonymousConfigDto)

	return err
}

type NullableAnonymousConfigDto struct {
	value *AnonymousConfigDto
	isSet bool
}

func (v NullableAnonymousConfigDto) Get() *AnonymousConfigDto {
	return v.value
}

func (v *NullableAnonymousConfigDto) Set(val *AnonymousConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAnonymousConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAnonymousConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnonymousConfigDto(val *AnonymousConfigDto) *NullableAnonymousConfigDto {
	return &NullableAnonymousConfigDto{value: val, isSet: true}
}

func (v NullableAnonymousConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnonymousConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

