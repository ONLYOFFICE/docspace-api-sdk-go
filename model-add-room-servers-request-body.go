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

// checks if the AddRoomServersRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddRoomServersRequestBody{}

// AddRoomServersRequestBody Parameters specifying which MCP servers to assign to the room.
type AddRoomServersRequestBody struct {
	// Set of unique identifiers of MCP servers to associate with the room. A maximum of 5 servers can be assigned per room.
	Servers []string `json:"servers"`
}

type _AddRoomServersRequestBody AddRoomServersRequestBody

// NewAddRoomServersRequestBody instantiates a new AddRoomServersRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddRoomServersRequestBody(servers []string) *AddRoomServersRequestBody {
	this := AddRoomServersRequestBody{}
	this.Servers = servers
	return &this
}

// NewAddRoomServersRequestBodyWithDefaults instantiates a new AddRoomServersRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddRoomServersRequestBodyWithDefaults() *AddRoomServersRequestBody {
	this := AddRoomServersRequestBody{}
	return &this
}

// GetServers returns the Servers field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AddRoomServersRequestBody) GetServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddRoomServersRequestBody) GetServersOk() ([]string, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// SetServers sets field value
func (o *AddRoomServersRequestBody) SetServers(v []string) {
	o.Servers = v
}

func (o AddRoomServersRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddRoomServersRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

func (o *AddRoomServersRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"servers",
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

	varAddRoomServersRequestBody := _AddRoomServersRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAddRoomServersRequestBody)

	if err != nil {
		return err
	}

	*o = AddRoomServersRequestBody(varAddRoomServersRequestBody)

	return err
}

type NullableAddRoomServersRequestBody struct {
	value *AddRoomServersRequestBody
	isSet bool
}

func (v NullableAddRoomServersRequestBody) Get() *AddRoomServersRequestBody {
	return v.value
}

func (v *NullableAddRoomServersRequestBody) Set(val *AddRoomServersRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRoomServersRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRoomServersRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRoomServersRequestBody(val *AddRoomServersRequestBody) *NullableAddRoomServersRequestBody {
	return &NullableAddRoomServersRequestBody{value: val, isSet: true}
}

func (v NullableAddRoomServersRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRoomServersRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

