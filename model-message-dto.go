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

// checks if the MessageDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDto{}

// MessageDto The chat message information.
type MessageDto struct {
	// The unique identifier of the message.
	Id *int64 `json:"id,omitempty"`
	Role *Role `json:"role,omitempty"`
	// The ordered collection of content blocks that make up the message body (text, tool calls, or attachments).
	Contents []MessageContentDto `json:"contents,omitempty"`
	CreatedOn *ApiDateTime `json:"createdOn,omitempty"`
}

// NewMessageDto instantiates a new MessageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDto() *MessageDto {
	this := MessageDto{}
	return &this
}

// NewMessageDtoWithDefaults instantiates a new MessageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDtoWithDefaults() *MessageDto {
	this := MessageDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MessageDto) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDto) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MessageDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MessageDto) SetId(v int64) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MessageDto) GetRole() Role {
	if o == nil || IsNil(o.Role) {
		var ret Role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDto) GetRoleOk() (*Role, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MessageDto) IsRoleSet() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given Role and assigns it to the Role field.
func (o *MessageDto) SetRole(v Role) {
	o.Role = &v
}

// GetContents returns the Contents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageDto) GetContents() []MessageContentDto {
	if o == nil {
		var ret []MessageContentDto
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageDto) GetContentsOk() ([]MessageContentDto, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *MessageDto) IsContentsSet() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []MessageContentDto and assigns it to the Contents field.
func (o *MessageDto) SetContents(v []MessageContentDto) {
	o.Contents = v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *MessageDto) GetCreatedOn() ApiDateTime {
	if o == nil || IsNil(o.CreatedOn) {
		var ret ApiDateTime
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDto) GetCreatedOnOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *MessageDto) IsCreatedOnSet() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given ApiDateTime and assigns it to the CreatedOn field.
func (o *MessageDto) SetCreatedOn(v ApiDateTime) {
	o.CreatedOn = &v
}

func (o MessageDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	return toSerialize, nil
}

type NullableMessageDto struct {
	value *MessageDto
	isSet bool
}

func (v NullableMessageDto) Get() *MessageDto {
	return v.value
}

func (v *NullableMessageDto) Set(val *MessageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDto(val *MessageDto) *NullableMessageDto {
	return &NullableMessageDto{value: val, isSet: true}
}

func (v NullableMessageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

