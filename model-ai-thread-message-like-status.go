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
	// `running`, `requires-action`, `complete` or `incomplete`.
	Type string `json:"type"`
	// Why the message is in that state: `tool-calls` or `interrupt` for `requires-action`; `stop` or `unknown` for `complete`; `cancelled`, `tool-calls`, `length`, `content-filter`, `other` or `error` for `incomplete`. An `incomplete` message the chat stored carries `cancelled` (the round was stopped) or `error`.
	Reason *string `json:"reason,omitempty"`
	// What failed, on an `incomplete` message whose `reason` is `error`.
	Error *AiChatErrorPayload `json:"error,omitempty"`
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

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AiThreadMessageLikeStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AiThreadMessageLikeStatus) IsReasonSet() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AiThreadMessageLikeStatus) SetReason(v string) {
	o.Reason = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AiThreadMessageLikeStatus) GetError() AiChatErrorPayload {
	if o == nil || IsNil(o.Error) {
		var ret AiChatErrorPayload
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeStatus) GetErrorOk() (*AiChatErrorPayload, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AiThreadMessageLikeStatus) IsErrorSet() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AiChatErrorPayload and assigns it to the Error field.
func (o *AiThreadMessageLikeStatus) SetError(v AiChatErrorPayload) {
	o.Error = &v
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
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
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

