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

// checks if the IconRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IconRequest{}

// IconRequest struct for IconRequest
type IconRequest struct {
	// Group icon
	Icon NullableString `json:"icon,omitempty"`
}

// NewIconRequest instantiates a new IconRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIconRequest() *IconRequest {
	this := IconRequest{}
	return &this
}

// NewIconRequestWithDefaults instantiates a new IconRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIconRequestWithDefaults() *IconRequest {
	this := IconRequest{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IconRequest) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IconRequest) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *IconRequest) IsIconSet() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *IconRequest) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *IconRequest) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *IconRequest) UnsetIcon() {
	o.Icon.Unset()
}

func (o IconRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IconRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	return toSerialize, nil
}

type NullableIconRequest struct {
	value *IconRequest
	isSet bool
}

func (v NullableIconRequest) Get() *IconRequest {
	return v.value
}

func (v *NullableIconRequest) Set(val *IconRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIconRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIconRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIconRequest(val *IconRequest) *NullableIconRequest {
	return &NullableIconRequest{value: val, isSet: true}
}

func (v NullableIconRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIconRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

