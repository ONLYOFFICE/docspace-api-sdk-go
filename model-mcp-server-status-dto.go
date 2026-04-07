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

// checks if the McpServerStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &McpServerStatusDto{}

// McpServerStatusDto MCP server status within a room, reflecting the current user's connection state for OAuth-based servers.
type McpServerStatusDto struct {
	// Unique identifier of the MCP server.
	Id *string `json:"id,omitempty"`
	// Display name of the MCP server.
	Name NullableString `json:"name"`
	ServerType *ServerType `json:"serverType,omitempty"`
	// Indicates whether the current user has an active connection to this server. For direct-connection servers this is always true; for OAuth-based servers it reflects whether the user has completed authorization.
	Connected *bool `json:"connected,omitempty"`
	Icon *Icon `json:"icon,omitempty"`
	// Indicates whether the server requires a configuration reset due to connectivity or credential issues.
	NeedReset *bool `json:"needReset,omitempty"`
}

type _McpServerStatusDto McpServerStatusDto

// NewMcpServerStatusDto instantiates a new McpServerStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcpServerStatusDto(name NullableString) *McpServerStatusDto {
	this := McpServerStatusDto{}
	this.Name = name
	return &this
}

// NewMcpServerStatusDtoWithDefaults instantiates a new McpServerStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcpServerStatusDtoWithDefaults() *McpServerStatusDto {
	this := McpServerStatusDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *McpServerStatusDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerStatusDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *McpServerStatusDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *McpServerStatusDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *McpServerStatusDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McpServerStatusDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *McpServerStatusDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *McpServerStatusDto) GetServerType() ServerType {
	if o == nil || IsNil(o.ServerType) {
		var ret ServerType
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerStatusDto) GetServerTypeOk() (*ServerType, bool) {
	if o == nil || IsNil(o.ServerType) {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *McpServerStatusDto) IsServerTypeSet() bool {
	if o != nil && !IsNil(o.ServerType) {
		return true
	}

	return false
}

// SetServerType gets a reference to the given ServerType and assigns it to the ServerType field.
func (o *McpServerStatusDto) SetServerType(v ServerType) {
	o.ServerType = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *McpServerStatusDto) GetConnected() bool {
	if o == nil || IsNil(o.Connected) {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerStatusDto) GetConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *McpServerStatusDto) IsConnectedSet() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *McpServerStatusDto) SetConnected(v bool) {
	o.Connected = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *McpServerStatusDto) GetIcon() Icon {
	if o == nil || IsNil(o.Icon) {
		var ret Icon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerStatusDto) GetIconOk() (*Icon, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *McpServerStatusDto) IsIconSet() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given Icon and assigns it to the Icon field.
func (o *McpServerStatusDto) SetIcon(v Icon) {
	o.Icon = &v
}

// GetNeedReset returns the NeedReset field value if set, zero value otherwise.
func (o *McpServerStatusDto) GetNeedReset() bool {
	if o == nil || IsNil(o.NeedReset) {
		var ret bool
		return ret
	}
	return *o.NeedReset
}

// GetNeedResetOk returns a tuple with the NeedReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpServerStatusDto) GetNeedResetOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedReset) {
		return nil, false
	}
	return o.NeedReset, true
}

// HasNeedReset returns a boolean if a field has been set.
func (o *McpServerStatusDto) IsNeedResetSet() bool {
	if o != nil && !IsNil(o.NeedReset) {
		return true
	}

	return false
}

// SetNeedReset gets a reference to the given bool and assigns it to the NeedReset field.
func (o *McpServerStatusDto) SetNeedReset(v bool) {
	o.NeedReset = &v
}

func (o McpServerStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o McpServerStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name.Get()
	if !IsNil(o.ServerType) {
		toSerialize["serverType"] = o.ServerType
	}
	if !IsNil(o.Connected) {
		toSerialize["connected"] = o.Connected
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.NeedReset) {
		toSerialize["needReset"] = o.NeedReset
	}
	return toSerialize, nil
}

func (o *McpServerStatusDto) UnmarshalJSON(data []byte) (err error) {
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

	varMcpServerStatusDto := _McpServerStatusDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varMcpServerStatusDto)

	if err != nil {
		return err
	}

	*o = McpServerStatusDto(varMcpServerStatusDto)

	return err
}

type NullableMcpServerStatusDto struct {
	value *McpServerStatusDto
	isSet bool
}

func (v NullableMcpServerStatusDto) Get() *McpServerStatusDto {
	return v.value
}

func (v *NullableMcpServerStatusDto) Set(val *McpServerStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMcpServerStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMcpServerStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcpServerStatusDto(val *McpServerStatusDto) *NullableMcpServerStatusDto {
	return &NullableMcpServerStatusDto{value: val, isSet: true}
}

func (v NullableMcpServerStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcpServerStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

