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

// checks if the AiOpenOrCreateResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenOrCreateResult{}

// AiOpenOrCreateResult Resolved thread state returned by `ThreadsEngine.openOrCreate`.
type AiOpenOrCreateResult struct {
	// The thread that was opened, or the one just created.
	ThreadId string `json:"threadId"`
	// Empty string for existing threads — the engine doesn't re-fetch.
	Title string `json:"title"`
	// The messages already in the thread - empty for a thread that was just created.
	PriorMessages []AiThreadMessageLike `json:"priorMessages"`
}

type _AiOpenOrCreateResult AiOpenOrCreateResult

// NewAiOpenOrCreateResult instantiates a new AiOpenOrCreateResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenOrCreateResult(threadId string, title string, priorMessages []AiThreadMessageLike) *AiOpenOrCreateResult {
	this := AiOpenOrCreateResult{}
	this.ThreadId = threadId
	this.Title = title
	this.PriorMessages = priorMessages
	return &this
}

// NewAiOpenOrCreateResultWithDefaults instantiates a new AiOpenOrCreateResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenOrCreateResultWithDefaults() *AiOpenOrCreateResult {
	this := AiOpenOrCreateResult{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *AiOpenOrCreateResult) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiOpenOrCreateResult) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiOpenOrCreateResult) SetThreadId(v string) {
	o.ThreadId = v
}

// GetTitle returns the Title field value
func (o *AiOpenOrCreateResult) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiOpenOrCreateResult) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AiOpenOrCreateResult) SetTitle(v string) {
	o.Title = v
}

// GetPriorMessages returns the PriorMessages field value
func (o *AiOpenOrCreateResult) GetPriorMessages() []AiThreadMessageLike {
	if o == nil {
		var ret []AiThreadMessageLike
		return ret
	}

	return o.PriorMessages
}

// GetPriorMessagesOk returns a tuple with the PriorMessages field value
// and a boolean to check if the value has been set.
func (o *AiOpenOrCreateResult) GetPriorMessagesOk() ([]AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriorMessages, true
}

// SetPriorMessages sets field value
func (o *AiOpenOrCreateResult) SetPriorMessages(v []AiThreadMessageLike) {
	o.PriorMessages = v
}

func (o AiOpenOrCreateResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenOrCreateResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threadId"] = o.ThreadId
	toSerialize["title"] = o.Title
	toSerialize["priorMessages"] = o.PriorMessages
	return toSerialize, nil
}

func (o *AiOpenOrCreateResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threadId",
		"title",
		"priorMessages",
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

	varAiOpenOrCreateResult := _AiOpenOrCreateResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiOpenOrCreateResult)

	if err != nil {
		return err
	}

	*o = AiOpenOrCreateResult(varAiOpenOrCreateResult)

	return err
}

type NullableAiOpenOrCreateResult struct {
	value *AiOpenOrCreateResult
	isSet bool
}

func (v NullableAiOpenOrCreateResult) Get() *AiOpenOrCreateResult {
	return v.value
}

func (v *NullableAiOpenOrCreateResult) Set(val *AiOpenOrCreateResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenOrCreateResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenOrCreateResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenOrCreateResult(val *AiOpenOrCreateResult) *NullableAiOpenOrCreateResult {
	return &NullableAiOpenOrCreateResult{value: val, isSet: true}
}

func (v NullableAiOpenOrCreateResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenOrCreateResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

