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

// checks if the DeleteServersRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteServersRequestBody{}

// DeleteServersRequestBody Parameters specifying which MCP servers to delete.
type DeleteServersRequestBody struct {
	// Set of unique identifiers of the MCP servers to permanently remove. All room associations and connection data will also be deleted.
	Servers []string `json:"servers"`
}

type _DeleteServersRequestBody DeleteServersRequestBody

// NewDeleteServersRequestBody instantiates a new DeleteServersRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteServersRequestBody(servers []string) *DeleteServersRequestBody {
	this := DeleteServersRequestBody{}
	this.Servers = servers
	return &this
}

// NewDeleteServersRequestBodyWithDefaults instantiates a new DeleteServersRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteServersRequestBodyWithDefaults() *DeleteServersRequestBody {
	this := DeleteServersRequestBody{}
	return &this
}

// GetServers returns the Servers field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *DeleteServersRequestBody) GetServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteServersRequestBody) GetServersOk() ([]string, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// SetServers sets field value
func (o *DeleteServersRequestBody) SetServers(v []string) {
	o.Servers = v
}

func (o DeleteServersRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteServersRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

func (o *DeleteServersRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteServersRequestBody := _DeleteServersRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDeleteServersRequestBody)

	if err != nil {
		return err
	}

	*o = DeleteServersRequestBody(varDeleteServersRequestBody)

	return err
}

type NullableDeleteServersRequestBody struct {
	value *DeleteServersRequestBody
	isSet bool
}

func (v NullableDeleteServersRequestBody) Get() *DeleteServersRequestBody {
	return v.value
}

func (v *NullableDeleteServersRequestBody) Set(val *DeleteServersRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteServersRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteServersRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteServersRequestBody(val *DeleteServersRequestBody) *NullableDeleteServersRequestBody {
	return &NullableDeleteServersRequestBody{value: val, isSet: true}
}

func (v NullableDeleteServersRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteServersRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

