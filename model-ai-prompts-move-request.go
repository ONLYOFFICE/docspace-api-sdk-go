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

// checks if the AiPromptsMoveRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPromptsMoveRequest{}

// AiPromptsMoveRequest struct for AiPromptsMoveRequest
type AiPromptsMoveRequest struct {
	// Prompt id to move.
	Id string `json:"id"`
	// Target folder id, or `null` for root.
	FolderId NullableString `json:"folderId"`
}

type _AiPromptsMoveRequest AiPromptsMoveRequest

// NewAiPromptsMoveRequest instantiates a new AiPromptsMoveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPromptsMoveRequest(id string, folderId NullableString) *AiPromptsMoveRequest {
	this := AiPromptsMoveRequest{}
	this.Id = id
	this.FolderId = folderId
	return &this
}

// NewAiPromptsMoveRequestWithDefaults instantiates a new AiPromptsMoveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPromptsMoveRequestWithDefaults() *AiPromptsMoveRequest {
	this := AiPromptsMoveRequest{}
	return &this
}

// GetId returns the Id field value
func (o *AiPromptsMoveRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiPromptsMoveRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiPromptsMoveRequest) SetId(v string) {
	o.Id = v
}

// GetFolderId returns the FolderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiPromptsMoveRequest) GetFolderId() string {
	if o == nil || o.FolderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPromptsMoveRequest) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// SetFolderId sets field value
func (o *AiPromptsMoveRequest) SetFolderId(v string) {
	o.FolderId.Set(&v)
}

func (o AiPromptsMoveRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPromptsMoveRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["folderId"] = o.FolderId.Get()
	return toSerialize, nil
}

func (o *AiPromptsMoveRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"folderId",
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

	varAiPromptsMoveRequest := _AiPromptsMoveRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPromptsMoveRequest)

	if err != nil {
		return err
	}

	*o = AiPromptsMoveRequest(varAiPromptsMoveRequest)

	return err
}

type NullableAiPromptsMoveRequest struct {
	value *AiPromptsMoveRequest
	isSet bool
}

func (v NullableAiPromptsMoveRequest) Get() *AiPromptsMoveRequest {
	return v.value
}

func (v *NullableAiPromptsMoveRequest) Set(val *AiPromptsMoveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPromptsMoveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPromptsMoveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPromptsMoveRequest(val *AiPromptsMoveRequest) *NullableAiPromptsMoveRequest {
	return &NullableAiPromptsMoveRequest{value: val, isSet: true}
}

func (v NullableAiPromptsMoveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPromptsMoveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

