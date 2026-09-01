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

// checks if the AiAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAttachment{}

// AiAttachment Persistent record for a single attachment (file or image) referenced from a user message. Files carry extracted text in `content`; images carry base64 data in `base64`. Metadata (`title`, `path`, `type`) is always present for display purposes regardless of whether the heavy payload is loaded.
type AiAttachment struct {
	// Storage-assigned UUID.
	Id string `json:"id"`
	// file | image.
	Kind string `json:"kind"`
	// Origin of the attachment. `user` — uploaded by the user in the composer (the default when unset, for backward compatibility). `tool` — produced by a tool call (e.g. `generate_image`). Lets the integrator's adapter route or apply policies (separate bucket, quotas, TTL, CDN) per source.
	Source *string `json:"source,omitempty"`
	// Display label (filename or user-visible title).
	Title string `json:"title"`
	// Extracted text for files.
	Content *string `json:"content,omitempty"`
	// Base64 data URL for images.
	Base64 *string `json:"base64,omitempty"`
	// Original host file path (for files).
	Path *string `json:"path,omitempty"`
	// ONLYOFFICE file type code (for files).
	Type *float32 `json:"type,omitempty"`
	// Owning message id once linked. Unset while the attachment is a draft.
	MessageId *string `json:"messageId,omitempty"`
	// Owning thread id once linked. Unset while the attachment is a draft.
	ThreadId *string `json:"threadId,omitempty"`
	// Opaque scope token (entity / room) the attachment was created in. Drafts carry it so an entity switch keeps in-flight composer state isolated; once linked to a message the field is redundant with the thread's own entity binding.
	EntityId *string `json:"entityId,omitempty"`
	// Storage-assigned creation timestamp.
	CreatedAt float32 `json:"createdAt"`
	// Whether the attached form can be analyzed.
	CanAnalyze *bool `json:"canAnalyze,omitempty"`
	// Keys of the fields inside the form. `key` is the field identifier, `text` its human-readable label.
	FormKeys []AiAttachmentFormKeysInner `json:"formKeys,omitempty"`
}

type _AiAttachment AiAttachment

// NewAiAttachment instantiates a new AiAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAttachment(id string, kind string, title string, createdAt float32) *AiAttachment {
	this := AiAttachment{}
	this.Id = id
	this.Kind = kind
	this.Title = title
	this.CreatedAt = createdAt
	return &this
}

// NewAiAttachmentWithDefaults instantiates a new AiAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAttachmentWithDefaults() *AiAttachment {
	this := AiAttachment{}
	return &this
}

// GetId returns the Id field value
func (o *AiAttachment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiAttachment) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *AiAttachment) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AiAttachment) SetKind(v string) {
	o.Kind = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AiAttachment) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AiAttachment) IsSourceSet() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AiAttachment) SetSource(v string) {
	o.Source = &v
}

// GetTitle returns the Title field value
func (o *AiAttachment) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AiAttachment) SetTitle(v string) {
	o.Title = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AiAttachment) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AiAttachment) IsContentSet() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AiAttachment) SetContent(v string) {
	o.Content = &v
}

// GetBase64 returns the Base64 field value if set, zero value otherwise.
func (o *AiAttachment) GetBase64() string {
	if o == nil || IsNil(o.Base64) {
		var ret string
		return ret
	}
	return *o.Base64
}

// GetBase64Ok returns a tuple with the Base64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.Base64) {
		return nil, false
	}
	return o.Base64, true
}

// HasBase64 returns a boolean if a field has been set.
func (o *AiAttachment) IsBase64Set() bool {
	if o != nil && !IsNil(o.Base64) {
		return true
	}

	return false
}

// SetBase64 gets a reference to the given string and assigns it to the Base64 field.
func (o *AiAttachment) SetBase64(v string) {
	o.Base64 = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *AiAttachment) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *AiAttachment) IsPathSet() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *AiAttachment) SetPath(v string) {
	o.Path = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AiAttachment) GetType() float32 {
	if o == nil || IsNil(o.Type) {
		var ret float32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetTypeOk() (*float32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AiAttachment) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given float32 and assigns it to the Type field.
func (o *AiAttachment) SetType(v float32) {
	o.Type = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *AiAttachment) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *AiAttachment) IsMessageIdSet() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *AiAttachment) SetMessageId(v string) {
	o.MessageId = &v
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *AiAttachment) GetThreadId() string {
	if o == nil || IsNil(o.ThreadId) {
		var ret string
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetThreadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *AiAttachment) IsThreadIdSet() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given string and assigns it to the ThreadId field.
func (o *AiAttachment) SetThreadId(v string) {
	o.ThreadId = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiAttachment) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiAttachment) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiAttachment) SetEntityId(v string) {
	o.EntityId = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AiAttachment) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AiAttachment) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetCanAnalyze returns the CanAnalyze field value if set, zero value otherwise.
func (o *AiAttachment) GetCanAnalyze() bool {
	if o == nil || IsNil(o.CanAnalyze) {
		var ret bool
		return ret
	}
	return *o.CanAnalyze
}

// GetCanAnalyzeOk returns a tuple with the CanAnalyze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetCanAnalyzeOk() (*bool, bool) {
	if o == nil || IsNil(o.CanAnalyze) {
		return nil, false
	}
	return o.CanAnalyze, true
}

// HasCanAnalyze returns a boolean if a field has been set.
func (o *AiAttachment) IsCanAnalyzeSet() bool {
	if o != nil && !IsNil(o.CanAnalyze) {
		return true
	}

	return false
}

// SetCanAnalyze gets a reference to the given bool and assigns it to the CanAnalyze field.
func (o *AiAttachment) SetCanAnalyze(v bool) {
	o.CanAnalyze = &v
}

// GetFormKeys returns the FormKeys field value if set, zero value otherwise.
func (o *AiAttachment) GetFormKeys() []AiAttachmentFormKeysInner {
	if o == nil || IsNil(o.FormKeys) {
		var ret []AiAttachmentFormKeysInner
		return ret
	}
	return o.FormKeys
}

// GetFormKeysOk returns a tuple with the FormKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachment) GetFormKeysOk() ([]AiAttachmentFormKeysInner, bool) {
	if o == nil || IsNil(o.FormKeys) {
		return nil, false
	}
	return o.FormKeys, true
}

// HasFormKeys returns a boolean if a field has been set.
func (o *AiAttachment) IsFormKeysSet() bool {
	if o != nil && !IsNil(o.FormKeys) {
		return true
	}

	return false
}

// SetFormKeys gets a reference to the given []AiAttachmentFormKeysInner and assigns it to the FormKeys field.
func (o *AiAttachment) SetFormKeys(v []AiAttachmentFormKeysInner) {
	o.FormKeys = v
}

func (o AiAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["kind"] = o.Kind
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	toSerialize["title"] = o.Title
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Base64) {
		toSerialize["base64"] = o.Base64
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.CanAnalyze) {
		toSerialize["canAnalyze"] = o.CanAnalyze
	}
	if !IsNil(o.FormKeys) {
		toSerialize["formKeys"] = o.FormKeys
	}
	return toSerialize, nil
}

func (o *AiAttachment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"kind",
		"title",
		"createdAt",
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

	varAiAttachment := _AiAttachment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAttachment)

	if err != nil {
		return err
	}

	*o = AiAttachment(varAiAttachment)

	return err
}

type NullableAiAttachment struct {
	value *AiAttachment
	isSet bool
}

func (v NullableAiAttachment) Get() *AiAttachment {
	return v.value
}

func (v *NullableAiAttachment) Set(val *AiAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAttachment(val *AiAttachment) *NullableAiAttachment {
	return &NullableAiAttachment{value: val, isSet: true}
}

func (v NullableAiAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

