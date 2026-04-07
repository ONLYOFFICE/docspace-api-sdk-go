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

// checks if the McpServerShortDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &McpServerShortDto{}

// McpServerShortDto Compact MCP server summary without sensitive details like endpoint URL or authentication headers.
type McpServerShortDto struct {
	// Unique identifier of the MCP server.
	Id *string `json:"id,omitempty"`
	// Display name of the MCP server.
	Name NullableString `json:"name,omitempty"`
	ServerType *ServerType `json:"serverType,omitempty"`
	// Indicates whether the server is currently enabled and available for room assignment.
	Enabled *bool `json:"enabled,omitempty"`
	Icon *Icon `json:"icon,omitempty"`
	// Indicates whether the server requires a configuration reset due to connectivity or credential issues.
	NeedReset *bool `json:"needReset,omitempty"`
}

// NewMcpServerShortDto instantiates a new McpServerShortDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcpServerShortDto() *McpServerShortDto {
	this := McpServerShortDto{}
	return &this
}

// NewMcpServerShortDtoWithDefaults instantiates a new McpServerShortDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcpServerShortDtoWithDefaults() *McpServerShortDto {
	this := McpServerShortDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *McpServerShortDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerShortDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *McpServerShortDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *McpServerShortDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McpServerShortDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McpServerShortDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *McpServerShortDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *McpServerShortDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *McpServerShortDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *McpServerShortDto) UnsetName() {
	o.Name.Unset()
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *McpServerShortDto) GetServerType() ServerType {
	if o == nil || IsNil(o.ServerType) {
		var ret ServerType
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerShortDto) GetServerTypeOk() (*ServerType, bool) {
	if o == nil || IsNil(o.ServerType) {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *McpServerShortDto) IsServerTypeSet() bool {
	if o != nil && !IsNil(o.ServerType) {
		return true
	}

	return false
}

// SetServerType gets a reference to the given ServerType and assigns it to the ServerType field.
func (o *McpServerShortDto) SetServerType(v ServerType) {
	o.ServerType = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *McpServerShortDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerShortDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *McpServerShortDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *McpServerShortDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *McpServerShortDto) GetIcon() Icon {
	if o == nil || IsNil(o.Icon) {
		var ret Icon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerShortDto) GetIconOk() (*Icon, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *McpServerShortDto) IsIconSet() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given Icon and assigns it to the Icon field.
func (o *McpServerShortDto) SetIcon(v Icon) {
	o.Icon = &v
}

// GetNeedReset returns the NeedReset field value if set, zero value otherwise.
func (o *McpServerShortDto) GetNeedReset() bool {
	if o == nil || IsNil(o.NeedReset) {
		var ret bool
		return ret
	}
	return *o.NeedReset
}

// GetNeedResetOk returns a tuple with the NeedReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerShortDto) GetNeedResetOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedReset) {
		return nil, false
	}
	return o.NeedReset, true
}

// HasNeedReset returns a boolean if a field has been set.
func (o *McpServerShortDto) IsNeedResetSet() bool {
	if o != nil && !IsNil(o.NeedReset) {
		return true
	}

	return false
}

// SetNeedReset gets a reference to the given bool and assigns it to the NeedReset field.
func (o *McpServerShortDto) SetNeedReset(v bool) {
	o.NeedReset = &v
}

func (o McpServerShortDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o McpServerShortDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.ServerType) {
		toSerialize["serverType"] = o.ServerType
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

type NullableMcpServerShortDto struct {
	value *McpServerShortDto
	isSet bool
}

func (v NullableMcpServerShortDto) Get() *McpServerShortDto {
	return v.value
}

func (v *NullableMcpServerShortDto) Set(val *McpServerShortDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMcpServerShortDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMcpServerShortDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcpServerShortDto(val *McpServerShortDto) *NullableMcpServerShortDto {
	return &NullableMcpServerShortDto{value: val, isSet: true}
}

func (v NullableMcpServerShortDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcpServerShortDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

