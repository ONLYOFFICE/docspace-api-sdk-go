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

// checks if the StartNewChatBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartNewChatBody{}

// StartNewChatBody Parameters for starting a new AI chat session.
type StartNewChatBody struct {
	// The initial user message to send to the AI assistant.
	Message NullableString `json:"message"`
	// The optional collection of file identifiers to attach as context for the AI model.
	ContextFolderId NullableInt32 `json:"contextFolderId,omitempty"`
	// The list of attached files.
	Files []ContinueChatBodyFilesInner `json:"files,omitempty"`
}

type _StartNewChatBody StartNewChatBody

// NewStartNewChatBody instantiates a new StartNewChatBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartNewChatBody(message NullableString) *StartNewChatBody {
	this := StartNewChatBody{}
	this.Message = message
	return &this
}

// NewStartNewChatBodyWithDefaults instantiates a new StartNewChatBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartNewChatBodyWithDefaults() *StartNewChatBody {
	this := StartNewChatBody{}
	return &this
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StartNewChatBody) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartNewChatBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *StartNewChatBody) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetContextFolderId returns the ContextFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartNewChatBody) GetContextFolderId() int32 {
	if o == nil || IsNil(o.ContextFolderId.Get()) {
		var ret int32
		return ret
	}
	return *o.ContextFolderId.Get()
}

// GetContextFolderIdOk returns a tuple with the ContextFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartNewChatBody) GetContextFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContextFolderId.Get(), o.ContextFolderId.IsSet()
}

// HasContextFolderId returns a boolean if a field has been set.
func (o *StartNewChatBody) IsContextFolderIdSet() bool {
	if o != nil && o.ContextFolderId.IsSet() {
		return true
	}

	return false
}

// SetContextFolderId gets a reference to the given NullableInt32 and assigns it to the ContextFolderId field.
func (o *StartNewChatBody) SetContextFolderId(v int32) {
	o.ContextFolderId.Set(&v)
}
// SetContextFolderIdNil sets the value for ContextFolderId to be an explicit nil
func (o *StartNewChatBody) SetContextFolderIdNil() {
	o.ContextFolderId.Set(nil)
}

// UnsetContextFolderId ensures that no value is present for ContextFolderId, not even an explicit nil
func (o *StartNewChatBody) UnsetContextFolderId() {
	o.ContextFolderId.Unset()
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartNewChatBody) GetFiles() []ContinueChatBodyFilesInner {
	if o == nil {
		var ret []ContinueChatBodyFilesInner
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartNewChatBody) GetFilesOk() ([]ContinueChatBodyFilesInner, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *StartNewChatBody) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []ContinueChatBodyFilesInner and assigns it to the Files field.
func (o *StartNewChatBody) SetFiles(v []ContinueChatBodyFilesInner) {
	o.Files = v
}

func (o StartNewChatBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartNewChatBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message.Get()
	if o.ContextFolderId.IsSet() {
		toSerialize["contextFolderId"] = o.ContextFolderId.Get()
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return toSerialize, nil
}

func (o *StartNewChatBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varStartNewChatBody := _StartNewChatBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varStartNewChatBody)

	if err != nil {
		return err
	}

	*o = StartNewChatBody(varStartNewChatBody)

	return err
}

type NullableStartNewChatBody struct {
	value *StartNewChatBody
	isSet bool
}

func (v NullableStartNewChatBody) Get() *StartNewChatBody {
	return v.value
}

func (v *NullableStartNewChatBody) Set(val *StartNewChatBody) {
	v.value = val
	v.isSet = true
}

func (v NullableStartNewChatBody) IsSet() bool {
	return v.isSet
}

func (v *NullableStartNewChatBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartNewChatBody(val *StartNewChatBody) *NullableStartNewChatBody {
	return &NullableStartNewChatBody{value: val, isSet: true}
}

func (v NullableStartNewChatBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartNewChatBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

