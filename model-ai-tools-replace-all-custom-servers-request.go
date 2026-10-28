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

// checks if the AiToolsReplaceAllCustomServersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolsReplaceAllCustomServersRequest{}

// AiToolsReplaceAllCustomServersRequest struct for AiToolsReplaceAllCustomServersRequest
type AiToolsReplaceAllCustomServersRequest struct {
	// Full replacement set, keyed by server name.
	Map map[string]map[string]interface{} `json:"map"`
	EntityId *string `json:"entityId,omitempty"`
}

type _AiToolsReplaceAllCustomServersRequest AiToolsReplaceAllCustomServersRequest

// NewAiToolsReplaceAllCustomServersRequest instantiates a new AiToolsReplaceAllCustomServersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolsReplaceAllCustomServersRequest(map_ map[string]map[string]interface{}) *AiToolsReplaceAllCustomServersRequest {
	this := AiToolsReplaceAllCustomServersRequest{}
	this.Map = map_
	return &this
}

// NewAiToolsReplaceAllCustomServersRequestWithDefaults instantiates a new AiToolsReplaceAllCustomServersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolsReplaceAllCustomServersRequestWithDefaults() *AiToolsReplaceAllCustomServersRequest {
	this := AiToolsReplaceAllCustomServersRequest{}
	return &this
}

// GetMap returns the Map field value
func (o *AiToolsReplaceAllCustomServersRequest) GetMap() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Map
}

// GetMapOk returns a tuple with the Map field value
// and a boolean to check if the value has been set.
func (o *AiToolsReplaceAllCustomServersRequest) GetMapOk() (map[string]map[string]interface{}, bool) {
	if o == nil {
		return map[string]map[string]interface{}{}, false
	}
	return o.Map, true
}

// SetMap sets field value
func (o *AiToolsReplaceAllCustomServersRequest) SetMap(v map[string]map[string]interface{}) {
	o.Map = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiToolsReplaceAllCustomServersRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolsReplaceAllCustomServersRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiToolsReplaceAllCustomServersRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiToolsReplaceAllCustomServersRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiToolsReplaceAllCustomServersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolsReplaceAllCustomServersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["map"] = o.Map
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiToolsReplaceAllCustomServersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"map",
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

	varAiToolsReplaceAllCustomServersRequest := _AiToolsReplaceAllCustomServersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolsReplaceAllCustomServersRequest)

	if err != nil {
		return err
	}

	*o = AiToolsReplaceAllCustomServersRequest(varAiToolsReplaceAllCustomServersRequest)

	return err
}

type NullableAiToolsReplaceAllCustomServersRequest struct {
	value *AiToolsReplaceAllCustomServersRequest
	isSet bool
}

func (v NullableAiToolsReplaceAllCustomServersRequest) Get() *AiToolsReplaceAllCustomServersRequest {
	return v.value
}

func (v *NullableAiToolsReplaceAllCustomServersRequest) Set(val *AiToolsReplaceAllCustomServersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolsReplaceAllCustomServersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolsReplaceAllCustomServersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolsReplaceAllCustomServersRequest(val *AiToolsReplaceAllCustomServersRequest) *NullableAiToolsReplaceAllCustomServersRequest {
	return &NullableAiToolsReplaceAllCustomServersRequest{value: val, isSet: true}
}

func (v NullableAiToolsReplaceAllCustomServersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolsReplaceAllCustomServersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

