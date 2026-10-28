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

// checks if the AiThreadMessageLike type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadMessageLike{}

// AiThreadMessageLike A single chat message as it travels on the wire.
type AiThreadMessageLike struct {
	// Storage-assigned message id (absent on inbound drafts).
	Id *string `json:"id,omitempty"`
	// Message author role.
	Role string `json:"role"`
	Content AiThreadMessageLikeContent `json:"content"`
	// Creation timestamp, ISO-8601 on the wire.
	CreatedAt *string `json:"createdAt,omitempty"`
	Status *AiThreadMessageLikeStatus `json:"status,omitempty"`
	// Arbitrary per-message metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Attachments linked to the message.
	Attachments []map[string]interface{} `json:"attachments,omitempty"`
}

type _AiThreadMessageLike AiThreadMessageLike

// NewAiThreadMessageLike instantiates a new AiThreadMessageLike object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadMessageLike(role string, content AiThreadMessageLikeContent) *AiThreadMessageLike {
	this := AiThreadMessageLike{}
	this.Role = role
	this.Content = content
	return &this
}

// NewAiThreadMessageLikeWithDefaults instantiates a new AiThreadMessageLike object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadMessageLikeWithDefaults() *AiThreadMessageLike {
	this := AiThreadMessageLike{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiThreadMessageLike) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLike) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiThreadMessageLike) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AiThreadMessageLike) SetId(v string) {
	o.Id = &v
}

// GetRole returns the Role field value
func (o *AiThreadMessageLike) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLike) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *AiThreadMessageLike) SetRole(v string) {
	o.Role = v
}

// GetContent returns the Content field value
func (o *AiThreadMessageLike) GetContent() AiThreadMessageLikeContent {
	if o == nil {
		var ret AiThreadMessageLikeContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLike) GetContentOk() (*AiThreadMessageLikeContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *AiThreadMessageLike) SetContent(v AiThreadMessageLikeContent) {
	o.Content = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiThreadMessageLike) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLike) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiThreadMessageLike) IsCreatedAtSet() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AiThreadMessageLike) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiThreadMessageLike) GetStatus() AiThreadMessageLikeStatus {
	if o == nil || IsNil(o.Status) {
		var ret AiThreadMessageLikeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLike) GetStatusOk() (*AiThreadMessageLikeStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiThreadMessageLike) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AiThreadMessageLikeStatus and assigns it to the Status field.
func (o *AiThreadMessageLike) SetStatus(v AiThreadMessageLikeStatus) {
	o.Status = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AiThreadMessageLike) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLike) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AiThreadMessageLike) IsMetadataSet() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AiThreadMessageLike) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *AiThreadMessageLike) GetAttachments() []map[string]interface{} {
	if o == nil || IsNil(o.Attachments) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLike) GetAttachmentsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AiThreadMessageLike) IsAttachmentsSet() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []map[string]interface{} and assigns it to the Attachments field.
func (o *AiThreadMessageLike) SetAttachments(v []map[string]interface{}) {
	o.Attachments = v
}

func (o AiThreadMessageLike) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadMessageLike) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["role"] = o.Role
	toSerialize["content"] = o.Content
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

func (o *AiThreadMessageLike) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role",
		"content",
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

	varAiThreadMessageLike := _AiThreadMessageLike{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadMessageLike)

	if err != nil {
		return err
	}

	*o = AiThreadMessageLike(varAiThreadMessageLike)

	return err
}

type NullableAiThreadMessageLike struct {
	value *AiThreadMessageLike
	isSet bool
}

func (v NullableAiThreadMessageLike) Get() *AiThreadMessageLike {
	return v.value
}

func (v *NullableAiThreadMessageLike) Set(val *AiThreadMessageLike) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadMessageLike) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadMessageLike) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadMessageLike(val *AiThreadMessageLike) *NullableAiThreadMessageLike {
	return &NullableAiThreadMessageLike{value: val, isSet: true}
}

func (v NullableAiThreadMessageLike) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadMessageLike) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

