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

// checks if the McpServerDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &McpServerDto{}

// McpServerDto Full MCP server configuration, including connection details and authentication headers. Returned for administrator-level operations.
type McpServerDto struct {
	// Unique identifier of the MCP server.
	Id *string `json:"id,omitempty"`
	// Display name of the MCP server.
	Name NullableString `json:"name,omitempty"`
	// Human-readable description of the server's purpose and capabilities.
	Description NullableString `json:"description,omitempty"`
	// Base URL of the MCP server endpoint.
	Endpoint NullableString `json:"endpoint,omitempty"`
	ServerType *ServerType `json:"serverType,omitempty"`
	// HTTP headers sent with every request to the server (e.g., authentication tokens).
	Headers map[string]string `json:"headers,omitempty"`
	// Indicates whether the server is currently enabled and available for room assignment.
	Enabled *bool `json:"enabled,omitempty"`
	Icon *Icon `json:"icon,omitempty"`
	// Indicates whether the server requires a configuration reset due to connectivity or credential issues.
	NeedReset *bool `json:"needReset,omitempty"`
}

// NewMcpServerDto instantiates a new McpServerDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcpServerDto() *McpServerDto {
	this := McpServerDto{}
	return &this
}

// NewMcpServerDtoWithDefaults instantiates a new McpServerDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcpServerDtoWithDefaults() *McpServerDto {
	this := McpServerDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *McpServerDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *McpServerDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *McpServerDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McpServerDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McpServerDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *McpServerDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *McpServerDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *McpServerDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *McpServerDto) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McpServerDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McpServerDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *McpServerDto) IsDescriptionSet() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *McpServerDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *McpServerDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *McpServerDto) UnsetDescription() {
	o.Description.Unset()
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McpServerDto) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint.Get()) {
		var ret string
		return ret
	}
	return *o.Endpoint.Get()
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McpServerDto) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoint.Get(), o.Endpoint.IsSet()
}

// HasEndpoint returns a boolean if a field has been set.
func (o *McpServerDto) IsEndpointSet() bool {
	if o != nil && o.Endpoint.IsSet() {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given NullableString and assigns it to the Endpoint field.
func (o *McpServerDto) SetEndpoint(v string) {
	o.Endpoint.Set(&v)
}
// SetEndpointNil sets the value for Endpoint to be an explicit nil
func (o *McpServerDto) SetEndpointNil() {
	o.Endpoint.Set(nil)
}

// UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
func (o *McpServerDto) UnsetEndpoint() {
	o.Endpoint.Unset()
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *McpServerDto) GetServerType() ServerType {
	if o == nil || IsNil(o.ServerType) {
		var ret ServerType
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerDto) GetServerTypeOk() (*ServerType, bool) {
	if o == nil || IsNil(o.ServerType) {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *McpServerDto) IsServerTypeSet() bool {
	if o != nil && !IsNil(o.ServerType) {
		return true
	}

	return false
}

// SetServerType gets a reference to the given ServerType and assigns it to the ServerType field.
func (o *McpServerDto) SetServerType(v ServerType) {
	o.ServerType = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McpServerDto) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McpServerDto) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *McpServerDto) IsHeadersSet() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *McpServerDto) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *McpServerDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *McpServerDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *McpServerDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *McpServerDto) GetIcon() Icon {
	if o == nil || IsNil(o.Icon) {
		var ret Icon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerDto) GetIconOk() (*Icon, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *McpServerDto) IsIconSet() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given Icon and assigns it to the Icon field.
func (o *McpServerDto) SetIcon(v Icon) {
	o.Icon = &v
}

// GetNeedReset returns the NeedReset field value if set, zero value otherwise.
func (o *McpServerDto) GetNeedReset() bool {
	if o == nil || IsNil(o.NeedReset) {
		var ret bool
		return ret
	}
	return *o.NeedReset
}

// GetNeedResetOk returns a tuple with the NeedReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerDto) GetNeedResetOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedReset) {
		return nil, false
	}
	return o.NeedReset, true
}

// HasNeedReset returns a boolean if a field has been set.
func (o *McpServerDto) IsNeedResetSet() bool {
	if o != nil && !IsNil(o.NeedReset) {
		return true
	}

	return false
}

// SetNeedReset gets a reference to the given bool and assigns it to the NeedReset field.
func (o *McpServerDto) SetNeedReset(v bool) {
	o.NeedReset = &v
}

func (o McpServerDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o McpServerDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Endpoint.IsSet() {
		toSerialize["endpoint"] = o.Endpoint.Get()
	}
	if !IsNil(o.ServerType) {
		toSerialize["serverType"] = o.ServerType
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.NeedReset) {
		toSerialize["needReset"] = o.NeedReset
	}
	return toSerialize, nil
}

type NullableMcpServerDto struct {
	value *McpServerDto
	isSet bool
}

func (v NullableMcpServerDto) Get() *McpServerDto {
	return v.value
}

func (v *NullableMcpServerDto) Set(val *McpServerDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMcpServerDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMcpServerDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcpServerDto(val *McpServerDto) *NullableMcpServerDto {
	return &NullableMcpServerDto{value: val, isSet: true}
}

func (v NullableMcpServerDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcpServerDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

