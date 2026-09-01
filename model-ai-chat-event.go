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

// checks if the AiChatEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiChatEvent{}

// AiChatEvent Discriminated event emitted by the streaming methods of `AIEngine`. The engine never invokes user-supplied middleware or callbacks directly — every observable side-effect is encoded as a `ChatEvent` so the same stream can be replayed over SSE, WebSocket, or in-process.  Pause point: `tool-call-pending` is the only stop. The UI must execute the tool itself (consulting `autoAllow` to decide between the silent path and the approve dialog) and resume via `AIEngine.approveToolCall` or `AIEngine.denyToolCall`.  Other variants are pure data:  - `message-start` / `message-delta` / `message-end` — assistant reply lifecycle. - `message-incomplete` — the provider returned an error or incomplete status. - `thread-title` — auto-generated title ready for a new thread.
type AiChatEvent struct {
	// Emitted once per `sendWithStream` call, immediately after the user message has been persisted by storage and before the assistant stream starts. Carries the storage-assigned `id` and `createdAt`. The UI uses it to render the user bubble — no client-side optimistic placeholder is needed, which keeps the runtime tree free of phantom nodes from index-fallback ids.
	Type string `json:"type"`
	// The message the event is about, in the state it has reached.
	Message *AiThreadMessageLike `json:"message,omitempty"`
	// The storage identifier of that message.
	MessageId *string `json:"messageId,omitempty"`
	// The zero-based position of the pending tool call within the message.
	Idx *float32 `json:"idx,omitempty"`
	// The thread the event belongs to.
	ThreadId *string `json:"threadId,omitempty"`
	// The consumer should execute the tool without prompting the user. True when the tool is in the persisted always-allow list, or the tool itself opts in via `TMCPItem.requireApproval === false` (host tools default to this). For a client-side tool with a server-side engine, this lets the engine return the pending call already flagged auto-allow so the client runs it and streams the result back without a dialog round-trip.
	AutoAllow *bool `json:"autoAllow,omitempty"`
	// Set when the tool is served by a server-side system source: the consumer must NOT execute it locally — only show the approval UI (unless `autoAllow`) and resume via `approveToolCall` (no `result` needed) / `denyToolCall`. The engine runs it in-engine.
	ServerExecuted *bool `json:"serverExecuted,omitempty"`
	// The generated thread title.
	Title *string `json:"title,omitempty"`
	// The profile that generated the title, when one was used.
	ProfileId *string `json:"profileId,omitempty"`
}

type _AiChatEvent AiChatEvent

// NewAiChatEvent instantiates a new AiChatEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiChatEvent(type_ string) *AiChatEvent {
	this := AiChatEvent{}
	this.Type = type_
	return &this
}

// NewAiChatEventWithDefaults instantiates a new AiChatEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiChatEventWithDefaults() *AiChatEvent {
	this := AiChatEvent{}
	return &this
}

// GetType returns the Type field value
func (o *AiChatEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiChatEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiChatEvent) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AiChatEvent) GetMessage() AiThreadMessageLike {
	if o == nil || IsNil(o.Message) {
		var ret AiThreadMessageLike
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatEvent) GetMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AiChatEvent) IsMessageSet() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given AiThreadMessageLike and assigns it to the Message field.
func (o *AiChatEvent) SetMessage(v AiThreadMessageLike) {
	o.Message = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *AiChatEvent) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatEvent) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *AiChatEvent) IsMessageIdSet() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *AiChatEvent) SetMessageId(v string) {
	o.MessageId = &v
}

// GetIdx returns the Idx field value if set, zero value otherwise.
func (o *AiChatEvent) GetIdx() float32 {
	if o == nil || IsNil(o.Idx) {
		var ret float32
		return ret
	}
	return *o.Idx
}

// GetIdxOk returns a tuple with the Idx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatEvent) GetIdxOk() (*float32, bool) {
	if o == nil || IsNil(o.Idx) {
		return nil, false
	}
	return o.Idx, true
}

// HasIdx returns a boolean if a field has been set.
func (o *AiChatEvent) IsIdxSet() bool {
	if o != nil && !IsNil(o.Idx) {
		return true
	}

	return false
}

// SetIdx gets a reference to the given float32 and assigns it to the Idx field.
func (o *AiChatEvent) SetIdx(v float32) {
	o.Idx = &v
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *AiChatEvent) GetThreadId() string {
	if o == nil || IsNil(o.ThreadId) {
		var ret string
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatEvent) GetThreadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *AiChatEvent) IsThreadIdSet() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given string and assigns it to the ThreadId field.
func (o *AiChatEvent) SetThreadId(v string) {
	o.ThreadId = &v
}

// GetAutoAllow returns the AutoAllow field value if set, zero value otherwise.
func (o *AiChatEvent) GetAutoAllow() bool {
	if o == nil || IsNil(o.AutoAllow) {
		var ret bool
		return ret
	}
	return *o.AutoAllow
}

// GetAutoAllowOk returns a tuple with the AutoAllow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatEvent) GetAutoAllowOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoAllow) {
		return nil, false
	}
	return o.AutoAllow, true
}

// HasAutoAllow returns a boolean if a field has been set.
func (o *AiChatEvent) IsAutoAllowSet() bool {
	if o != nil && !IsNil(o.AutoAllow) {
		return true
	}

	return false
}

// SetAutoAllow gets a reference to the given bool and assigns it to the AutoAllow field.
func (o *AiChatEvent) SetAutoAllow(v bool) {
	o.AutoAllow = &v
}

// GetServerExecuted returns the ServerExecuted field value if set, zero value otherwise.
func (o *AiChatEvent) GetServerExecuted() bool {
	if o == nil || IsNil(o.ServerExecuted) {
		var ret bool
		return ret
	}
	return *o.ServerExecuted
}

// GetServerExecutedOk returns a tuple with the ServerExecuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatEvent) GetServerExecutedOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerExecuted) {
		return nil, false
	}
	return o.ServerExecuted, true
}

// HasServerExecuted returns a boolean if a field has been set.
func (o *AiChatEvent) IsServerExecutedSet() bool {
	if o != nil && !IsNil(o.ServerExecuted) {
		return true
	}

	return false
}

// SetServerExecuted gets a reference to the given bool and assigns it to the ServerExecuted field.
func (o *AiChatEvent) SetServerExecuted(v bool) {
	o.ServerExecuted = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiChatEvent) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatEvent) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiChatEvent) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiChatEvent) SetTitle(v string) {
	o.Title = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiChatEvent) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatEvent) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiChatEvent) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiChatEvent) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiChatEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiChatEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.Idx) {
		toSerialize["idx"] = o.Idx
	}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	if !IsNil(o.AutoAllow) {
		toSerialize["autoAllow"] = o.AutoAllow
	}
	if !IsNil(o.ServerExecuted) {
		toSerialize["serverExecuted"] = o.ServerExecuted
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return toSerialize, nil
}

func (o *AiChatEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAiChatEvent := _AiChatEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiChatEvent)

	if err != nil {
		return err
	}

	*o = AiChatEvent(varAiChatEvent)

	return err
}

type NullableAiChatEvent struct {
	value *AiChatEvent
	isSet bool
}

func (v NullableAiChatEvent) Get() *AiChatEvent {
	return v.value
}

func (v *NullableAiChatEvent) Set(val *AiChatEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAiChatEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAiChatEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiChatEvent(val *AiChatEvent) *NullableAiChatEvent {
	return &NullableAiChatEvent{value: val, isSet: true}
}

func (v NullableAiChatEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiChatEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

