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

// checks if the AddMcpServerRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddMcpServerRequestBody{}

// AddMcpServerRequestBody Parameters for creating a new custom MCP server.
type AddMcpServerRequestBody struct {
	// Unique display name for the server. Only letters, numbers, underscores, and hyphens are allowed. Maximum 128 characters.
	Name NullableString `json:"name"`
	// Human-readable description of the server's purpose and capabilities. Maximum 255 characters.
	Description NullableString `json:"description"`
	// Base URL of the MCP server endpoint. Must be a valid, reachable URL. The system will verify connectivity during registration.
	Endpoint NullableString `json:"endpoint"`
	// Optional HTTP headers to include with every request to the MCP server (e.g., authentication tokens or API keys).
	Headers map[string]string `json:"headers,omitempty"`
	// Optional Base64-encoded icon image for the server. Used as the visual identifier in the UI.
	Icon NullableString `json:"icon,omitempty"`
}

type _AddMcpServerRequestBody AddMcpServerRequestBody

// NewAddMcpServerRequestBody instantiates a new AddMcpServerRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMcpServerRequestBody(name NullableString, description NullableString, endpoint NullableString) *AddMcpServerRequestBody {
	this := AddMcpServerRequestBody{}
	this.Name = name
	this.Description = description
	this.Endpoint = endpoint
	return &this
}

// NewAddMcpServerRequestBodyWithDefaults instantiates a new AddMcpServerRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMcpServerRequestBodyWithDefaults() *AddMcpServerRequestBody {
	this := AddMcpServerRequestBody{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AddMcpServerRequestBody) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMcpServerRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *AddMcpServerRequestBody) SetName(v string) {
	o.Name.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AddMcpServerRequestBody) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMcpServerRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *AddMcpServerRequestBody) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetEndpoint returns the Endpoint field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AddMcpServerRequestBody) GetEndpoint() string {
	if o == nil || o.Endpoint.Get() == nil {
		var ret string
		return ret
	}

	return *o.Endpoint.Get()
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMcpServerRequestBody) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoint.Get(), o.Endpoint.IsSet()
}

// SetEndpoint sets field value
func (o *AddMcpServerRequestBody) SetEndpoint(v string) {
	o.Endpoint.Set(&v)
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddMcpServerRequestBody) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMcpServerRequestBody) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AddMcpServerRequestBody) IsHeadersSet() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *AddMcpServerRequestBody) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddMcpServerRequestBody) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddMcpServerRequestBody) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *AddMcpServerRequestBody) IsIconSet() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *AddMcpServerRequestBody) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *AddMcpServerRequestBody) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *AddMcpServerRequestBody) UnsetIcon() {
	o.Icon.Unset()
}

func (o AddMcpServerRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddMcpServerRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["endpoint"] = o.Endpoint.Get()
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	return toSerialize, nil
}

func (o *AddMcpServerRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"endpoint",
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

	varAddMcpServerRequestBody := _AddMcpServerRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAddMcpServerRequestBody)

	if err != nil {
		return err
	}

	*o = AddMcpServerRequestBody(varAddMcpServerRequestBody)

	return err
}

type NullableAddMcpServerRequestBody struct {
	value *AddMcpServerRequestBody
	isSet bool
}

func (v NullableAddMcpServerRequestBody) Get() *AddMcpServerRequestBody {
	return v.value
}

func (v *NullableAddMcpServerRequestBody) Set(val *AddMcpServerRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMcpServerRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMcpServerRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMcpServerRequestBody(val *AddMcpServerRequestBody) *NullableAddMcpServerRequestBody {
	return &NullableAddMcpServerRequestBody{value: val, isSet: true}
}

func (v NullableAddMcpServerRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMcpServerRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

