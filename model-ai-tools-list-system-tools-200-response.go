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

// checks if the AiToolsListSystemTools200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolsListSystemTools200Response{}

// AiToolsListSystemTools200Response struct for AiToolsListSystemTools200Response
type AiToolsListSystemTools200Response struct {
	// Tools by server name, covering both the host-configured system servers and the custom MCP servers registered for this scope.
	Groups map[string][]AiTMCPItem `json:"groups"`
	// Why a registered custom server could not be reached, keyed by server name. A server that answered is absent from this map.
	Errors map[string]string `json:"errors"`
	// Names of the host-configured system servers among the keys of `groups`; everything else there was registered as a custom server.
	System []string `json:"system"`
}

type _AiToolsListSystemTools200Response AiToolsListSystemTools200Response

// NewAiToolsListSystemTools200Response instantiates a new AiToolsListSystemTools200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolsListSystemTools200Response(groups map[string][]AiTMCPItem, errors map[string]string, system []string) *AiToolsListSystemTools200Response {
	this := AiToolsListSystemTools200Response{}
	this.Groups = groups
	this.Errors = errors
	this.System = system
	return &this
}

// NewAiToolsListSystemTools200ResponseWithDefaults instantiates a new AiToolsListSystemTools200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolsListSystemTools200ResponseWithDefaults() *AiToolsListSystemTools200Response {
	this := AiToolsListSystemTools200Response{}
	return &this
}

// GetGroups returns the Groups field value
func (o *AiToolsListSystemTools200Response) GetGroups() map[string][]AiTMCPItem {
	if o == nil {
		var ret map[string][]AiTMCPItem
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *AiToolsListSystemTools200Response) GetGroupsOk() (map[string][]AiTMCPItem, bool) {
	if o == nil {
		return map[string][]AiTMCPItem{}, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *AiToolsListSystemTools200Response) SetGroups(v map[string][]AiTMCPItem) {
	o.Groups = v
}

// GetErrors returns the Errors field value
func (o *AiToolsListSystemTools200Response) GetErrors() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *AiToolsListSystemTools200Response) GetErrorsOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *AiToolsListSystemTools200Response) SetErrors(v map[string]string) {
	o.Errors = v
}

// GetSystem returns the System field value
func (o *AiToolsListSystemTools200Response) GetSystem() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *AiToolsListSystemTools200Response) GetSystemOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.System, true
}

// SetSystem sets field value
func (o *AiToolsListSystemTools200Response) SetSystem(v []string) {
	o.System = v
}

func (o AiToolsListSystemTools200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolsListSystemTools200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groups"] = o.Groups
	toSerialize["errors"] = o.Errors
	toSerialize["system"] = o.System
	return toSerialize, nil
}

func (o *AiToolsListSystemTools200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groups",
		"errors",
		"system",
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

	varAiToolsListSystemTools200Response := _AiToolsListSystemTools200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolsListSystemTools200Response)

	if err != nil {
		return err
	}

	*o = AiToolsListSystemTools200Response(varAiToolsListSystemTools200Response)

	return err
}

type NullableAiToolsListSystemTools200Response struct {
	value *AiToolsListSystemTools200Response
	isSet bool
}

func (v NullableAiToolsListSystemTools200Response) Get() *AiToolsListSystemTools200Response {
	return v.value
}

func (v *NullableAiToolsListSystemTools200Response) Set(val *AiToolsListSystemTools200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolsListSystemTools200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolsListSystemTools200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolsListSystemTools200Response(val *AiToolsListSystemTools200Response) *NullableAiToolsListSystemTools200Response {
	return &NullableAiToolsListSystemTools200Response{value: val, isSet: true}
}

func (v NullableAiToolsListSystemTools200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolsListSystemTools200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

