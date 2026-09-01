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

// checks if the AiThreadsCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadsCreateRequest{}

// AiThreadsCreateRequest struct for AiThreadsCreateRequest
type AiThreadsCreateRequest struct {
	// Thread title.
	Title string `json:"title"`
	// Optional profile to bind.
	ProfileId *string `json:"profileId,omitempty"`
	// Optional entity (room) scope.
	EntityId *string `json:"entityId,omitempty"`
}

type _AiThreadsCreateRequest AiThreadsCreateRequest

// NewAiThreadsCreateRequest instantiates a new AiThreadsCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadsCreateRequest(title string) *AiThreadsCreateRequest {
	this := AiThreadsCreateRequest{}
	this.Title = title
	return &this
}

// NewAiThreadsCreateRequestWithDefaults instantiates a new AiThreadsCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadsCreateRequestWithDefaults() *AiThreadsCreateRequest {
	this := AiThreadsCreateRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *AiThreadsCreateRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiThreadsCreateRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AiThreadsCreateRequest) SetTitle(v string) {
	o.Title = v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiThreadsCreateRequest) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadsCreateRequest) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiThreadsCreateRequest) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiThreadsCreateRequest) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiThreadsCreateRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadsCreateRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiThreadsCreateRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiThreadsCreateRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiThreadsCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadsCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiThreadsCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varAiThreadsCreateRequest := _AiThreadsCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadsCreateRequest)

	if err != nil {
		return err
	}

	*o = AiThreadsCreateRequest(varAiThreadsCreateRequest)

	return err
}

type NullableAiThreadsCreateRequest struct {
	value *AiThreadsCreateRequest
	isSet bool
}

func (v NullableAiThreadsCreateRequest) Get() *AiThreadsCreateRequest {
	return v.value
}

func (v *NullableAiThreadsCreateRequest) Set(val *AiThreadsCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadsCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadsCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadsCreateRequest(val *AiThreadsCreateRequest) *NullableAiThreadsCreateRequest {
	return &NullableAiThreadsCreateRequest{value: val, isSet: true}
}

func (v NullableAiThreadsCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadsCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

