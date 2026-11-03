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

// checks if the AiActionArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiActionArgs{}

// AiActionArgs Per-request options of an AI round: extra tools, the depth of reasoning, a system prompt override and the rooms and folder connected as context.
type AiActionArgs struct {
	// Extra tools offered to the model for this request.
	Tools []AiMCPItem `json:"tools,omitempty"`
	// Legacy extended-thinking switch; stands for `medium`. `reasoningLevel` wins when both are set.
	IsReasoning *bool `json:"isReasoning,omitempty"`
	// Depth of extended thinking for the round; providers clamp it to what the model accepts.
	ReasoningLevel *AiReasoningLevel `json:"reasoningLevel,omitempty"`
	Prompt *AiActionArgsPrompt `json:"prompt,omitempty"`
	// Usage notes of the MCP servers the caller holds the connection to, keyed by server type; folded into the system prompt as advisory text.
	ToolInstructions map[string]string `json:"toolInstructions,omitempty"`
	// Instructions rooms connected to the chat as context, each with the skills the round may use.
	ContextRooms []AiContextRoomRef `json:"contextRooms,omitempty"`
	// Knowledge rooms connected to the chat; the knowledge search tool of the round is limited to them. Each `id` has to be a DocSpace room id.
	KnowledgeRooms []AiKnowledgeRoomRef `json:"knowledgeRooms,omitempty"`
	// A single knowledge room, as an older client sends it. Read only when `knowledgeRooms` is absent; send `knowledgeRooms` instead.
	// Deprecated
	KnowledgeRoom *AiKnowledgeRoomRef `json:"knowledgeRoom,omitempty"`
	// The folder connected to the chat as the round's context. It becomes the location the model is told the user works in, and the entry the round is attributed to.
	FolderContext *AiFolderContextRef `json:"folderContext,omitempty"`
	// The entry the round is attributed to. A hint, resolved against the caller's own access.
	EntityId *string `json:"entityId,omitempty"`
}

// NewAiActionArgs instantiates a new AiActionArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiActionArgs() *AiActionArgs {
	this := AiActionArgs{}
	return &this
}

// NewAiActionArgsWithDefaults instantiates a new AiActionArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiActionArgsWithDefaults() *AiActionArgs {
	this := AiActionArgs{}
	return &this
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *AiActionArgs) GetTools() []AiMCPItem {
	if o == nil || IsNil(o.Tools) {
		var ret []AiMCPItem
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetToolsOk() ([]AiMCPItem, bool) {
	if o == nil || IsNil(o.Tools) {
		return nil, false
	}
	return o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *AiActionArgs) IsToolsSet() bool {
	if o != nil && !IsNil(o.Tools) {
		return true
	}

	return false
}

// SetTools gets a reference to the given []AiMCPItem and assigns it to the Tools field.
func (o *AiActionArgs) SetTools(v []AiMCPItem) {
	o.Tools = v
}

// GetIsReasoning returns the IsReasoning field value if set, zero value otherwise.
func (o *AiActionArgs) GetIsReasoning() bool {
	if o == nil || IsNil(o.IsReasoning) {
		var ret bool
		return ret
	}
	return *o.IsReasoning
}

// GetIsReasoningOk returns a tuple with the IsReasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetIsReasoningOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReasoning) {
		return nil, false
	}
	return o.IsReasoning, true
}

// HasIsReasoning returns a boolean if a field has been set.
func (o *AiActionArgs) IsIsReasoningSet() bool {
	if o != nil && !IsNil(o.IsReasoning) {
		return true
	}

	return false
}

// SetIsReasoning gets a reference to the given bool and assigns it to the IsReasoning field.
func (o *AiActionArgs) SetIsReasoning(v bool) {
	o.IsReasoning = &v
}

// GetReasoningLevel returns the ReasoningLevel field value if set, zero value otherwise.
func (o *AiActionArgs) GetReasoningLevel() AiReasoningLevel {
	if o == nil || IsNil(o.ReasoningLevel) {
		var ret AiReasoningLevel
		return ret
	}
	return *o.ReasoningLevel
}

// GetReasoningLevelOk returns a tuple with the ReasoningLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetReasoningLevelOk() (*AiReasoningLevel, bool) {
	if o == nil || IsNil(o.ReasoningLevel) {
		return nil, false
	}
	return o.ReasoningLevel, true
}

// HasReasoningLevel returns a boolean if a field has been set.
func (o *AiActionArgs) IsReasoningLevelSet() bool {
	if o != nil && !IsNil(o.ReasoningLevel) {
		return true
	}

	return false
}

// SetReasoningLevel gets a reference to the given AiReasoningLevel and assigns it to the ReasoningLevel field.
func (o *AiActionArgs) SetReasoningLevel(v AiReasoningLevel) {
	o.ReasoningLevel = &v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiActionArgs) GetPrompt() AiActionArgsPrompt {
	if o == nil || IsNil(o.Prompt) {
		var ret AiActionArgsPrompt
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetPromptOk() (*AiActionArgsPrompt, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiActionArgs) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given AiActionArgsPrompt and assigns it to the Prompt field.
func (o *AiActionArgs) SetPrompt(v AiActionArgsPrompt) {
	o.Prompt = &v
}

// GetToolInstructions returns the ToolInstructions field value if set, zero value otherwise.
func (o *AiActionArgs) GetToolInstructions() map[string]string {
	if o == nil || IsNil(o.ToolInstructions) {
		var ret map[string]string
		return ret
	}
	return o.ToolInstructions
}

// GetToolInstructionsOk returns a tuple with the ToolInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetToolInstructionsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.ToolInstructions) {
		return map[string]string{}, false
	}
	return o.ToolInstructions, true
}

// HasToolInstructions returns a boolean if a field has been set.
func (o *AiActionArgs) IsToolInstructionsSet() bool {
	if o != nil && !IsNil(o.ToolInstructions) {
		return true
	}

	return false
}

// SetToolInstructions gets a reference to the given map[string]string and assigns it to the ToolInstructions field.
func (o *AiActionArgs) SetToolInstructions(v map[string]string) {
	o.ToolInstructions = v
}

// GetContextRooms returns the ContextRooms field value if set, zero value otherwise.
func (o *AiActionArgs) GetContextRooms() []AiContextRoomRef {
	if o == nil || IsNil(o.ContextRooms) {
		var ret []AiContextRoomRef
		return ret
	}
	return o.ContextRooms
}

// GetContextRoomsOk returns a tuple with the ContextRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetContextRoomsOk() ([]AiContextRoomRef, bool) {
	if o == nil || IsNil(o.ContextRooms) {
		return nil, false
	}
	return o.ContextRooms, true
}

// HasContextRooms returns a boolean if a field has been set.
func (o *AiActionArgs) IsContextRoomsSet() bool {
	if o != nil && !IsNil(o.ContextRooms) {
		return true
	}

	return false
}

// SetContextRooms gets a reference to the given []AiContextRoomRef and assigns it to the ContextRooms field.
func (o *AiActionArgs) SetContextRooms(v []AiContextRoomRef) {
	o.ContextRooms = v
}

// GetKnowledgeRooms returns the KnowledgeRooms field value if set, zero value otherwise.
func (o *AiActionArgs) GetKnowledgeRooms() []AiKnowledgeRoomRef {
	if o == nil || IsNil(o.KnowledgeRooms) {
		var ret []AiKnowledgeRoomRef
		return ret
	}
	return o.KnowledgeRooms
}

// GetKnowledgeRoomsOk returns a tuple with the KnowledgeRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetKnowledgeRoomsOk() ([]AiKnowledgeRoomRef, bool) {
	if o == nil || IsNil(o.KnowledgeRooms) {
		return nil, false
	}
	return o.KnowledgeRooms, true
}

// HasKnowledgeRooms returns a boolean if a field has been set.
func (o *AiActionArgs) IsKnowledgeRoomsSet() bool {
	if o != nil && !IsNil(o.KnowledgeRooms) {
		return true
	}

	return false
}

// SetKnowledgeRooms gets a reference to the given []AiKnowledgeRoomRef and assigns it to the KnowledgeRooms field.
func (o *AiActionArgs) SetKnowledgeRooms(v []AiKnowledgeRoomRef) {
	o.KnowledgeRooms = v
}

// GetKnowledgeRoom returns the KnowledgeRoom field value if set, zero value otherwise.
// Deprecated
func (o *AiActionArgs) GetKnowledgeRoom() AiKnowledgeRoomRef {
	if o == nil || IsNil(o.KnowledgeRoom) {
		var ret AiKnowledgeRoomRef
		return ret
	}
	return *o.KnowledgeRoom
}

// GetKnowledgeRoomOk returns a tuple with the KnowledgeRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AiActionArgs) GetKnowledgeRoomOk() (*AiKnowledgeRoomRef, bool) {
	if o == nil || IsNil(o.KnowledgeRoom) {
		return nil, false
	}
	return o.KnowledgeRoom, true
}

// HasKnowledgeRoom returns a boolean if a field has been set.
func (o *AiActionArgs) IsKnowledgeRoomSet() bool {
	if o != nil && !IsNil(o.KnowledgeRoom) {
		return true
	}

	return false
}

// SetKnowledgeRoom gets a reference to the given AiKnowledgeRoomRef and assigns it to the KnowledgeRoom field.
// Deprecated
func (o *AiActionArgs) SetKnowledgeRoom(v AiKnowledgeRoomRef) {
	o.KnowledgeRoom = &v
}

// GetFolderContext returns the FolderContext field value if set, zero value otherwise.
func (o *AiActionArgs) GetFolderContext() AiFolderContextRef {
	if o == nil || IsNil(o.FolderContext) {
		var ret AiFolderContextRef
		return ret
	}
	return *o.FolderContext
}

// GetFolderContextOk returns a tuple with the FolderContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetFolderContextOk() (*AiFolderContextRef, bool) {
	if o == nil || IsNil(o.FolderContext) {
		return nil, false
	}
	return o.FolderContext, true
}

// HasFolderContext returns a boolean if a field has been set.
func (o *AiActionArgs) IsFolderContextSet() bool {
	if o != nil && !IsNil(o.FolderContext) {
		return true
	}

	return false
}

// SetFolderContext gets a reference to the given AiFolderContextRef and assigns it to the FolderContext field.
func (o *AiActionArgs) SetFolderContext(v AiFolderContextRef) {
	o.FolderContext = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiActionArgs) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiActionArgs) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiActionArgs) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiActionArgs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiActionArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tools) {
		toSerialize["tools"] = o.Tools
	}
	if !IsNil(o.IsReasoning) {
		toSerialize["isReasoning"] = o.IsReasoning
	}
	if !IsNil(o.ReasoningLevel) {
		toSerialize["reasoningLevel"] = o.ReasoningLevel
	}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.ToolInstructions) {
		toSerialize["toolInstructions"] = o.ToolInstructions
	}
	if !IsNil(o.ContextRooms) {
		toSerialize["contextRooms"] = o.ContextRooms
	}
	if !IsNil(o.KnowledgeRooms) {
		toSerialize["knowledgeRooms"] = o.KnowledgeRooms
	}
	if !IsNil(o.KnowledgeRoom) {
		toSerialize["knowledgeRoom"] = o.KnowledgeRoom
	}
	if !IsNil(o.FolderContext) {
		toSerialize["folderContext"] = o.FolderContext
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

type NullableAiActionArgs struct {
	value *AiActionArgs
	isSet bool
}

func (v NullableAiActionArgs) Get() *AiActionArgs {
	return v.value
}

func (v *NullableAiActionArgs) Set(val *AiActionArgs) {
	v.value = val
	v.isSet = true
}

func (v NullableAiActionArgs) IsSet() bool {
	return v.isSet
}

func (v *NullableAiActionArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiActionArgs(val *AiActionArgs) *NullableAiActionArgs {
	return &NullableAiActionArgs{value: val, isSet: true}
}

func (v NullableAiActionArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiActionArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

