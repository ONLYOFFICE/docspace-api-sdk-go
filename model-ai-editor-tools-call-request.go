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

// checks if the AiEditorToolsCallRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEditorToolsCallRequest{}

// AiEditorToolsCallRequest struct for AiEditorToolsCallRequest
type AiEditorToolsCallRequest struct {
	// Name of the tool to run, as listed by the tools endpoint. A name that is unknown or excluded from the editor is rejected with 400.
	Name string `json:"name"`
	// Arguments for the tool, shaped by that tool's own input schema. Treated as empty when it is not an object.
	Arguments map[string]*interface{} `json:"arguments,omitempty"`
	// Room the call is scoped to. Left out for a portal-wide call.
	EntityId *string `json:"entityId,omitempty"`
}

type _AiEditorToolsCallRequest AiEditorToolsCallRequest

// NewAiEditorToolsCallRequest instantiates a new AiEditorToolsCallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEditorToolsCallRequest(name string) *AiEditorToolsCallRequest {
	this := AiEditorToolsCallRequest{}
	this.Name = name
	return &this
}

// NewAiEditorToolsCallRequestWithDefaults instantiates a new AiEditorToolsCallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEditorToolsCallRequestWithDefaults() *AiEditorToolsCallRequest {
	this := AiEditorToolsCallRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AiEditorToolsCallRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiEditorToolsCallRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiEditorToolsCallRequest) SetName(v string) {
	o.Name = v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *AiEditorToolsCallRequest) GetArguments() map[string]*interface{} {
	if o == nil || IsNil(o.Arguments) {
		var ret map[string]*interface{}
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiEditorToolsCallRequest) GetArgumentsOk() (map[string]*interface{}, bool) {
	if o == nil || IsNil(o.Arguments) {
		return map[string]*interface{}{}, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *AiEditorToolsCallRequest) IsArgumentsSet() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given map[string]*interface{} and assigns it to the Arguments field.
func (o *AiEditorToolsCallRequest) SetArguments(v map[string]*interface{}) {
	o.Arguments = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiEditorToolsCallRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiEditorToolsCallRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiEditorToolsCallRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiEditorToolsCallRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiEditorToolsCallRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEditorToolsCallRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiEditorToolsCallRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAiEditorToolsCallRequest := _AiEditorToolsCallRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEditorToolsCallRequest)

	if err != nil {
		return err
	}

	*o = AiEditorToolsCallRequest(varAiEditorToolsCallRequest)

	return err
}

type NullableAiEditorToolsCallRequest struct {
	value *AiEditorToolsCallRequest
	isSet bool
}

func (v NullableAiEditorToolsCallRequest) Get() *AiEditorToolsCallRequest {
	return v.value
}

func (v *NullableAiEditorToolsCallRequest) Set(val *AiEditorToolsCallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEditorToolsCallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEditorToolsCallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEditorToolsCallRequest(val *AiEditorToolsCallRequest) *NullableAiEditorToolsCallRequest {
	return &NullableAiEditorToolsCallRequest{value: val, isSet: true}
}

func (v NullableAiEditorToolsCallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEditorToolsCallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

