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

// checks if the ChatDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatDto{}

// ChatDto The chat session information.
type ChatDto struct {
	// The unique identifier of the AI chat session.
	Id *string `json:"id,omitempty"`
	// The display title of the chat session.
	Title NullableString `json:"title,omitempty"`
	CreatedOn *ApiDateTime `json:"createdOn,omitempty"`
	ModifiedOn *ApiDateTime `json:"modifiedOn,omitempty"`
	CreatedBy *EmployeeDto `json:"createdBy,omitempty"`
}

// NewChatDto instantiates a new ChatDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatDto() *ChatDto {
	this := ChatDto{}
	return &this
}

// NewChatDtoWithDefaults instantiates a new ChatDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatDtoWithDefaults() *ChatDto {
	this := ChatDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChatDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChatDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChatDto) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ChatDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ChatDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ChatDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ChatDto) UnsetTitle() {
	o.Title.Unset()
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ChatDto) GetCreatedOn() ApiDateTime {
	if o == nil || IsNil(o.CreatedOn) {
		var ret ApiDateTime
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatDto) GetCreatedOnOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ChatDto) IsCreatedOnSet() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given ApiDateTime and assigns it to the CreatedOn field.
func (o *ChatDto) SetCreatedOn(v ApiDateTime) {
	o.CreatedOn = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *ChatDto) GetModifiedOn() ApiDateTime {
	if o == nil || IsNil(o.ModifiedOn) {
		var ret ApiDateTime
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatDto) GetModifiedOnOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ModifiedOn) {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *ChatDto) IsModifiedOnSet() bool {
	if o != nil && !IsNil(o.ModifiedOn) {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given ApiDateTime and assigns it to the ModifiedOn field.
func (o *ChatDto) SetModifiedOn(v ApiDateTime) {
	o.ModifiedOn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ChatDto) GetCreatedBy() EmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatDto) GetCreatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ChatDto) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EmployeeDto and assigns it to the CreatedBy field.
func (o *ChatDto) SetCreatedBy(v EmployeeDto) {
	o.CreatedBy = &v
}

func (o ChatDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if !IsNil(o.ModifiedOn) {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	return toSerialize, nil
}

type NullableChatDto struct {
	value *ChatDto
	isSet bool
}

func (v NullableChatDto) Get() *ChatDto {
	return v.value
}

func (v *NullableChatDto) Set(val *ChatDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChatDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChatDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatDto(val *ChatDto) *NullableChatDto {
	return &NullableChatDto{value: val, isSet: true}
}

func (v NullableChatDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

