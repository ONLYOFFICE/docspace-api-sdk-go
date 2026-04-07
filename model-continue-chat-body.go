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

// checks if the ContinueChatBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContinueChatBody{}

// ContinueChatBody Parameters for continuing an AI chat session.
type ContinueChatBody struct {
	// The user message to append to the conversation.
	Message NullableString `json:"message"`
	// The optional collection of file identifiers to attach as context for the AI model.
	ContextFolderId *int32 `json:"contextFolderId,omitempty"`
	// The list of attached files.
	Files []ContinueChatBodyFilesInner `json:"files,omitempty"`
}

type _ContinueChatBody ContinueChatBody

// NewContinueChatBody instantiates a new ContinueChatBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinueChatBody(message NullableString) *ContinueChatBody {
	this := ContinueChatBody{}
	this.Message = message
	return &this
}

// NewContinueChatBodyWithDefaults instantiates a new ContinueChatBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinueChatBodyWithDefaults() *ContinueChatBody {
	this := ContinueChatBody{}
	return &this
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ContinueChatBody) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContinueChatBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *ContinueChatBody) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetContextFolderId returns the ContextFolderId field value if set, zero value otherwise.
func (o *ContinueChatBody) GetContextFolderId() int32 {
	if o == nil || IsNil(o.ContextFolderId) {
		var ret int32
		return ret
	}
	return *o.ContextFolderId
}

// GetContextFolderIdOk returns a tuple with the ContextFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinueChatBody) GetContextFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ContextFolderId) {
		return nil, false
	}
	return o.ContextFolderId, true
}

// HasContextFolderId returns a boolean if a field has been set.
func (o *ContinueChatBody) IsContextFolderIdSet() bool {
	if o != nil && !IsNil(o.ContextFolderId) {
		return true
	}

	return false
}

// SetContextFolderId gets a reference to the given int32 and assigns it to the ContextFolderId field.
func (o *ContinueChatBody) SetContextFolderId(v int32) {
	o.ContextFolderId = &v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContinueChatBody) GetFiles() []ContinueChatBodyFilesInner {
	if o == nil {
		var ret []ContinueChatBodyFilesInner
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContinueChatBody) GetFilesOk() ([]ContinueChatBodyFilesInner, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ContinueChatBody) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []ContinueChatBodyFilesInner and assigns it to the Files field.
func (o *ContinueChatBody) SetFiles(v []ContinueChatBodyFilesInner) {
	o.Files = v
}

func (o ContinueChatBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinueChatBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message.Get()
	if !IsNil(o.ContextFolderId) {
		toSerialize["contextFolderId"] = o.ContextFolderId
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return toSerialize, nil
}

func (o *ContinueChatBody) UnmarshalJSON(data []byte) (err error) {
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

	varContinueChatBody := _ContinueChatBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varContinueChatBody)

	if err != nil {
		return err
	}

	*o = ContinueChatBody(varContinueChatBody)

	return err
}

type NullableContinueChatBody struct {
	value *ContinueChatBody
	isSet bool
}

func (v NullableContinueChatBody) Get() *ContinueChatBody {
	return v.value
}

func (v *NullableContinueChatBody) Set(val *ContinueChatBody) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueChatBody) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueChatBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueChatBody(val *ContinueChatBody) *NullableContinueChatBody {
	return &NullableContinueChatBody{value: val, isSet: true}
}

func (v NullableContinueChatBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueChatBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

