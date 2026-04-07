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

// checks if the ConnectServerRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectServerRequestBody{}

// ConnectServerRequestBody Parameters for completing an OAuth connection to an MCP server.
type ConnectServerRequestBody struct {
	// OAuth authorization code received from the provider's redirect. Used to exchange for access and refresh tokens.
	Code NullableString `json:"code"`
}

type _ConnectServerRequestBody ConnectServerRequestBody

// NewConnectServerRequestBody instantiates a new ConnectServerRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectServerRequestBody(code NullableString) *ConnectServerRequestBody {
	this := ConnectServerRequestBody{}
	this.Code = code
	return &this
}

// NewConnectServerRequestBodyWithDefaults instantiates a new ConnectServerRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectServerRequestBodyWithDefaults() *ConnectServerRequestBody {
	this := ConnectServerRequestBody{}
	return &this
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConnectServerRequestBody) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}

	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectServerRequestBody) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// SetCode sets field value
func (o *ConnectServerRequestBody) SetCode(v string) {
	o.Code.Set(&v)
}

func (o ConnectServerRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectServerRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code.Get()
	return toSerialize, nil
}

func (o *ConnectServerRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varConnectServerRequestBody := _ConnectServerRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varConnectServerRequestBody)

	if err != nil {
		return err
	}

	*o = ConnectServerRequestBody(varConnectServerRequestBody)

	return err
}

type NullableConnectServerRequestBody struct {
	value *ConnectServerRequestBody
	isSet bool
}

func (v NullableConnectServerRequestBody) Get() *ConnectServerRequestBody {
	return v.value
}

func (v *NullableConnectServerRequestBody) Set(val *ConnectServerRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectServerRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectServerRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectServerRequestBody(val *ConnectServerRequestBody) *NullableConnectServerRequestBody {
	return &NullableConnectServerRequestBody{value: val, isSet: true}
}

func (v NullableConnectServerRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectServerRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

