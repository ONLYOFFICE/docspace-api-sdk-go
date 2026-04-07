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

// checks if the DeleteRoomServersRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRoomServersRequestBody{}

// DeleteRoomServersRequestBody Parameters specifying which MCP servers to detach from the room.
type DeleteRoomServersRequestBody struct {
	// Set of unique identifiers of MCP servers to remove from the room. Associated connections and tool configurations will also be cleaned up.
	Servers []string `json:"servers"`
}

type _DeleteRoomServersRequestBody DeleteRoomServersRequestBody

// NewDeleteRoomServersRequestBody instantiates a new DeleteRoomServersRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRoomServersRequestBody(servers []string) *DeleteRoomServersRequestBody {
	this := DeleteRoomServersRequestBody{}
	this.Servers = servers
	return &this
}

// NewDeleteRoomServersRequestBodyWithDefaults instantiates a new DeleteRoomServersRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRoomServersRequestBodyWithDefaults() *DeleteRoomServersRequestBody {
	this := DeleteRoomServersRequestBody{}
	return &this
}

// GetServers returns the Servers field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *DeleteRoomServersRequestBody) GetServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteRoomServersRequestBody) GetServersOk() ([]string, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// SetServers sets field value
func (o *DeleteRoomServersRequestBody) SetServers(v []string) {
	o.Servers = v
}

func (o DeleteRoomServersRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRoomServersRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

func (o *DeleteRoomServersRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteRoomServersRequestBody := _DeleteRoomServersRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDeleteRoomServersRequestBody)

	if err != nil {
		return err
	}

	*o = DeleteRoomServersRequestBody(varDeleteRoomServersRequestBody)

	return err
}

type NullableDeleteRoomServersRequestBody struct {
	value *DeleteRoomServersRequestBody
	isSet bool
}

func (v NullableDeleteRoomServersRequestBody) Get() *DeleteRoomServersRequestBody {
	return v.value
}

func (v *NullableDeleteRoomServersRequestBody) Set(val *DeleteRoomServersRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRoomServersRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRoomServersRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRoomServersRequestBody(val *DeleteRoomServersRequestBody) *NullableDeleteRoomServersRequestBody {
	return &NullableDeleteRoomServersRequestBody{value: val, isSet: true}
}

func (v NullableDeleteRoomServersRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRoomServersRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

