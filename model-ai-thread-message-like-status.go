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

// checks if the AiThreadMessageLikeStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadMessageLikeStatus{}

// AiThreadMessageLikeStatus Delivery/generation status of the message.
type AiThreadMessageLikeStatus struct {
	Type string `json:"type"`
}

type _AiThreadMessageLikeStatus AiThreadMessageLikeStatus

// NewAiThreadMessageLikeStatus instantiates a new AiThreadMessageLikeStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadMessageLikeStatus(type_ string) *AiThreadMessageLikeStatus {
	this := AiThreadMessageLikeStatus{}
	this.Type = type_
	return &this
}

// NewAiThreadMessageLikeStatusWithDefaults instantiates a new AiThreadMessageLikeStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadMessageLikeStatusWithDefaults() *AiThreadMessageLikeStatus {
	this := AiThreadMessageLikeStatus{}
	return &this
}

// GetType returns the Type field value
func (o *AiThreadMessageLikeStatus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeStatus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiThreadMessageLikeStatus) SetType(v string) {
	o.Type = v
}

func (o AiThreadMessageLikeStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadMessageLikeStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *AiThreadMessageLikeStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varAiThreadMessageLikeStatus := _AiThreadMessageLikeStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadMessageLikeStatus)

	if err != nil {
		return err
	}

	*o = AiThreadMessageLikeStatus(varAiThreadMessageLikeStatus)

	return err
}

type NullableAiThreadMessageLikeStatus struct {
	value *AiThreadMessageLikeStatus
	isSet bool
}

func (v NullableAiThreadMessageLikeStatus) Get() *AiThreadMessageLikeStatus {
	return v.value
}

func (v *NullableAiThreadMessageLikeStatus) Set(val *AiThreadMessageLikeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadMessageLikeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadMessageLikeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadMessageLikeStatus(val *AiThreadMessageLikeStatus) *NullableAiThreadMessageLikeStatus {
	return &NullableAiThreadMessageLikeStatus{value: val, isSet: true}
}

func (v NullableAiThreadMessageLikeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadMessageLikeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

