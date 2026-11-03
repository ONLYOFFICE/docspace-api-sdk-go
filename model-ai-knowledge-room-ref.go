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

// checks if the AiKnowledgeRoomRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiKnowledgeRoomRef{}

// AiKnowledgeRoomRef One Knowledge room a round searches: the account and room ID from `GET api/2.0/ai/context/get-knowledge-folders`.
type AiKnowledgeRoomRef struct {
	// The account the room belongs to; `docspace` for a room of this portal.
	Cloud string `json:"cloud"`
	// The room ID; anything that is not a positive integer is ignored.
	Id string `json:"id"`
}

type _AiKnowledgeRoomRef AiKnowledgeRoomRef

// NewAiKnowledgeRoomRef instantiates a new AiKnowledgeRoomRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiKnowledgeRoomRef(cloud string, id string) *AiKnowledgeRoomRef {
	this := AiKnowledgeRoomRef{}
	this.Cloud = cloud
	this.Id = id
	return &this
}

// NewAiKnowledgeRoomRefWithDefaults instantiates a new AiKnowledgeRoomRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiKnowledgeRoomRefWithDefaults() *AiKnowledgeRoomRef {
	this := AiKnowledgeRoomRef{}
	return &this
}

// GetCloud returns the Cloud field value
func (o *AiKnowledgeRoomRef) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *AiKnowledgeRoomRef) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *AiKnowledgeRoomRef) SetCloud(v string) {
	o.Cloud = v
}

// GetId returns the Id field value
func (o *AiKnowledgeRoomRef) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiKnowledgeRoomRef) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiKnowledgeRoomRef) SetId(v string) {
	o.Id = v
}

func (o AiKnowledgeRoomRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiKnowledgeRoomRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud"] = o.Cloud
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AiKnowledgeRoomRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud",
		"id",
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

	varAiKnowledgeRoomRef := _AiKnowledgeRoomRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiKnowledgeRoomRef)

	if err != nil {
		return err
	}

	*o = AiKnowledgeRoomRef(varAiKnowledgeRoomRef)

	return err
}

type NullableAiKnowledgeRoomRef struct {
	value *AiKnowledgeRoomRef
	isSet bool
}

func (v NullableAiKnowledgeRoomRef) Get() *AiKnowledgeRoomRef {
	return v.value
}

func (v *NullableAiKnowledgeRoomRef) Set(val *AiKnowledgeRoomRef) {
	v.value = val
	v.isSet = true
}

func (v NullableAiKnowledgeRoomRef) IsSet() bool {
	return v.isSet
}

func (v *NullableAiKnowledgeRoomRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiKnowledgeRoomRef(val *AiKnowledgeRoomRef) *NullableAiKnowledgeRoomRef {
	return &NullableAiKnowledgeRoomRef{value: val, isSet: true}
}

func (v NullableAiKnowledgeRoomRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiKnowledgeRoomRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

