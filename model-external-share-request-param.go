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

// checks if the ExternalShareRequestParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalShareRequestParam{}

// ExternalShareRequestParam The password that unlocks a protected external share link.
type ExternalShareRequestParam struct {
	// The password chosen by the member who shared the entry, spelled exactly as they typed it. It is compared  against the stored value and never returned back; a mismatch is reported through the answer's status instead  of an error.
	Password NullableString `json:"password,omitempty"`
}

// NewExternalShareRequestParam instantiates a new ExternalShareRequestParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalShareRequestParam() *ExternalShareRequestParam {
	this := ExternalShareRequestParam{}
	return &this
}

// NewExternalShareRequestParamWithDefaults instantiates a new ExternalShareRequestParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalShareRequestParamWithDefaults() *ExternalShareRequestParam {
	this := ExternalShareRequestParam{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalShareRequestParam) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalShareRequestParam) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *ExternalShareRequestParam) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *ExternalShareRequestParam) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *ExternalShareRequestParam) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *ExternalShareRequestParam) UnsetPassword() {
	o.Password.Unset()
}

func (o ExternalShareRequestParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalShareRequestParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return toSerialize, nil
}

type NullableExternalShareRequestParam struct {
	value *ExternalShareRequestParam
	isSet bool
}

func (v NullableExternalShareRequestParam) Get() *ExternalShareRequestParam {
	return v.value
}

func (v *NullableExternalShareRequestParam) Set(val *ExternalShareRequestParam) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalShareRequestParam) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalShareRequestParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalShareRequestParam(val *ExternalShareRequestParam) *NullableExternalShareRequestParam {
	return &NullableExternalShareRequestParam{value: val, isSet: true}
}

func (v NullableExternalShareRequestParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalShareRequestParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

