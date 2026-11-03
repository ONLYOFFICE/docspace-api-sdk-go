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

// checks if the AiThreadMessageLikeContentAnyOfInnerInterrupt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadMessageLikeContentAnyOfInnerInterrupt{}

// AiThreadMessageLikeContentAnyOfInnerInterrupt A `tool-call` part's pending request for human input.
type AiThreadMessageLikeContentAnyOfInnerInterrupt struct {
	// Interrupt kind; always `human`.
	Type string `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

type _AiThreadMessageLikeContentAnyOfInnerInterrupt AiThreadMessageLikeContentAnyOfInnerInterrupt

// NewAiThreadMessageLikeContentAnyOfInnerInterrupt instantiates a new AiThreadMessageLikeContentAnyOfInnerInterrupt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadMessageLikeContentAnyOfInnerInterrupt(type_ string) *AiThreadMessageLikeContentAnyOfInnerInterrupt {
	this := AiThreadMessageLikeContentAnyOfInnerInterrupt{}
	this.Type = type_
	return &this
}

// NewAiThreadMessageLikeContentAnyOfInnerInterruptWithDefaults instantiates a new AiThreadMessageLikeContentAnyOfInnerInterrupt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadMessageLikeContentAnyOfInnerInterruptWithDefaults() *AiThreadMessageLikeContentAnyOfInnerInterrupt {
	this := AiThreadMessageLikeContentAnyOfInnerInterrupt{}
	return &this
}

// GetType returns the Type field value
func (o *AiThreadMessageLikeContentAnyOfInnerInterrupt) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInnerInterrupt) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiThreadMessageLikeContentAnyOfInnerInterrupt) SetType(v string) {
	o.Type = v
}

// GetPayload returns the Payload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiThreadMessageLikeContentAnyOfInnerInterrupt) GetPayload() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiThreadMessageLikeContentAnyOfInnerInterrupt) GetPayloadOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInnerInterrupt) IsPayloadSet() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given interface{} and assigns it to the Payload field.
func (o *AiThreadMessageLikeContentAnyOfInnerInterrupt) SetPayload(v interface{}) {
	o.Payload = v
}

func (o AiThreadMessageLikeContentAnyOfInnerInterrupt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadMessageLikeContentAnyOfInnerInterrupt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

func (o *AiThreadMessageLikeContentAnyOfInnerInterrupt) UnmarshalJSON(data []byte) (err error) {
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

	varAiThreadMessageLikeContentAnyOfInnerInterrupt := _AiThreadMessageLikeContentAnyOfInnerInterrupt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadMessageLikeContentAnyOfInnerInterrupt)

	if err != nil {
		return err
	}

	*o = AiThreadMessageLikeContentAnyOfInnerInterrupt(varAiThreadMessageLikeContentAnyOfInnerInterrupt)

	return err
}

type NullableAiThreadMessageLikeContentAnyOfInnerInterrupt struct {
	value *AiThreadMessageLikeContentAnyOfInnerInterrupt
	isSet bool
}

func (v NullableAiThreadMessageLikeContentAnyOfInnerInterrupt) Get() *AiThreadMessageLikeContentAnyOfInnerInterrupt {
	return v.value
}

func (v *NullableAiThreadMessageLikeContentAnyOfInnerInterrupt) Set(val *AiThreadMessageLikeContentAnyOfInnerInterrupt) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadMessageLikeContentAnyOfInnerInterrupt) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadMessageLikeContentAnyOfInnerInterrupt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadMessageLikeContentAnyOfInnerInterrupt(val *AiThreadMessageLikeContentAnyOfInnerInterrupt) *NullableAiThreadMessageLikeContentAnyOfInnerInterrupt {
	return &NullableAiThreadMessageLikeContentAnyOfInnerInterrupt{value: val, isSet: true}
}

func (v NullableAiThreadMessageLikeContentAnyOfInnerInterrupt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadMessageLikeContentAnyOfInnerInterrupt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

