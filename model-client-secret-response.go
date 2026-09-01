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

// checks if the ClientSecretResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientSecretResponse{}

// ClientSecretResponse The response containing the regenerated client secret.
type ClientSecretResponse struct {
	// The newly generated client secret.
	ClientSecret *string `json:"client_secret,omitempty"`
}

// NewClientSecretResponse instantiates a new ClientSecretResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientSecretResponse() *ClientSecretResponse {
	this := ClientSecretResponse{}
	return &this
}

// NewClientSecretResponseWithDefaults instantiates a new ClientSecretResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientSecretResponseWithDefaults() *ClientSecretResponse {
	this := ClientSecretResponse{}
	return &this
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ClientSecretResponse) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientSecretResponse) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ClientSecretResponse) IsClientSecretSet() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ClientSecretResponse) SetClientSecret(v string) {
	o.ClientSecret = &v
}

func (o ClientSecretResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientSecretResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
	}
	return toSerialize, nil
}

type NullableClientSecretResponse struct {
	value *ClientSecretResponse
	isSet bool
}

func (v NullableClientSecretResponse) Get() *ClientSecretResponse {
	return v.value
}

func (v *NullableClientSecretResponse) Set(val *ClientSecretResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientSecretResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientSecretResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientSecretResponse(val *ClientSecretResponse) *NullableClientSecretResponse {
	return &NullableClientSecretResponse{value: val, isSet: true}
}

func (v NullableClientSecretResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientSecretResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

