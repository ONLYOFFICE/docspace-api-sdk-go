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

// checks if the GobackConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GobackConfig{}

// GobackConfig The settings for the Open file location menu button and upper right corner button.
type GobackConfig struct {
	// Where the user is taken when they leave the document, normally the folder or the room it lies in. It is empty  when there is nowhere to return to, as in a framed opening.
	Url NullableString `json:"url,omitempty"`
}

// NewGobackConfig instantiates a new GobackConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGobackConfig() *GobackConfig {
	this := GobackConfig{}
	return &this
}

// NewGobackConfigWithDefaults instantiates a new GobackConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGobackConfigWithDefaults() *GobackConfig {
	this := GobackConfig{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GobackConfig) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GobackConfig) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *GobackConfig) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *GobackConfig) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *GobackConfig) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *GobackConfig) UnsetUrl() {
	o.Url.Unset()
}

func (o GobackConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GobackConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableGobackConfig struct {
	value *GobackConfig
	isSet bool
}

func (v NullableGobackConfig) Get() *GobackConfig {
	return v.value
}

func (v *NullableGobackConfig) Set(val *GobackConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGobackConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGobackConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGobackConfig(val *GobackConfig) *NullableGobackConfig {
	return &NullableGobackConfig{value: val, isSet: true}
}

func (v NullableGobackConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGobackConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

