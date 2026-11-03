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

// checks if the AiChatErrorPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiChatErrorPayload{}

// AiChatErrorPayload Why a round failed or stopped early, as the message status and the stream report it: a stable `code`, a message safe to show the user, and the cause when there is one.
type AiChatErrorPayload struct {
	// Stable machine-readable code. One of a canonical set — `auth`, `model_not_found`, `rate_limit`, `timeout`, `network`, `bad_request`, `server`, `not_found`, `insufficient_funds`, `aborted`, `unknown` — or a provider-specific string when the SDK supplies one verbatim.
	Code string `json:"code"`
	// Human-readable, end-user-safe message.
	Message string `json:"message"`
	// Optional nested cause, normalized through the same helper.
	Cause *AiChatErrorPayload `json:"cause,omitempty"`
}

type _AiChatErrorPayload AiChatErrorPayload

// NewAiChatErrorPayload instantiates a new AiChatErrorPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiChatErrorPayload(code string, message string) *AiChatErrorPayload {
	this := AiChatErrorPayload{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAiChatErrorPayloadWithDefaults instantiates a new AiChatErrorPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiChatErrorPayloadWithDefaults() *AiChatErrorPayload {
	this := AiChatErrorPayload{}
	return &this
}

// GetCode returns the Code field value
func (o *AiChatErrorPayload) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AiChatErrorPayload) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AiChatErrorPayload) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *AiChatErrorPayload) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiChatErrorPayload) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiChatErrorPayload) SetMessage(v string) {
	o.Message = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *AiChatErrorPayload) GetCause() AiChatErrorPayload {
	if o == nil || IsNil(o.Cause) {
		var ret AiChatErrorPayload
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatErrorPayload) GetCauseOk() (*AiChatErrorPayload, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *AiChatErrorPayload) IsCauseSet() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given AiChatErrorPayload and assigns it to the Cause field.
func (o *AiChatErrorPayload) SetCause(v AiChatErrorPayload) {
	o.Cause = &v
}

func (o AiChatErrorPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiChatErrorPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	return toSerialize, nil
}

func (o *AiChatErrorPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varAiChatErrorPayload := _AiChatErrorPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiChatErrorPayload)

	if err != nil {
		return err
	}

	*o = AiChatErrorPayload(varAiChatErrorPayload)

	return err
}

type NullableAiChatErrorPayload struct {
	value *AiChatErrorPayload
	isSet bool
}

func (v NullableAiChatErrorPayload) Get() *AiChatErrorPayload {
	return v.value
}

func (v *NullableAiChatErrorPayload) Set(val *AiChatErrorPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAiChatErrorPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAiChatErrorPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiChatErrorPayload(val *AiChatErrorPayload) *NullableAiChatErrorPayload {
	return &NullableAiChatErrorPayload{value: val, isSet: true}
}

func (v NullableAiChatErrorPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiChatErrorPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

