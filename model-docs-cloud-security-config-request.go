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

// checks if the DocsCloudSecurityConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudSecurityConfigRequest{}

// DocsCloudSecurityConfigRequest Represents the security configuration of a Docs Connect tenant.
type DocsCloudSecurityConfigRequest struct {
	// The security secret.
	Secret NullableString `json:"secret,omitempty"`
	// The security header name.
	Header NullableString `json:"header,omitempty"`
}

// NewDocsCloudSecurityConfigRequest instantiates a new DocsCloudSecurityConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudSecurityConfigRequest() *DocsCloudSecurityConfigRequest {
	this := DocsCloudSecurityConfigRequest{}
	return &this
}

// NewDocsCloudSecurityConfigRequestWithDefaults instantiates a new DocsCloudSecurityConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudSecurityConfigRequestWithDefaults() *DocsCloudSecurityConfigRequest {
	this := DocsCloudSecurityConfigRequest{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudSecurityConfigRequest) GetSecret() string {
	if o == nil || IsNil(o.Secret.Get()) {
		var ret string
		return ret
	}
	return *o.Secret.Get()
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudSecurityConfigRequest) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secret.Get(), o.Secret.IsSet()
}

// HasSecret returns a boolean if a field has been set.
func (o *DocsCloudSecurityConfigRequest) IsSecretSet() bool {
	if o != nil && o.Secret.IsSet() {
		return true
	}

	return false
}

// SetSecret gets a reference to the given NullableString and assigns it to the Secret field.
func (o *DocsCloudSecurityConfigRequest) SetSecret(v string) {
	o.Secret.Set(&v)
}
// SetSecretNil sets the value for Secret to be an explicit nil
func (o *DocsCloudSecurityConfigRequest) SetSecretNil() {
	o.Secret.Set(nil)
}

// UnsetSecret ensures that no value is present for Secret, not even an explicit nil
func (o *DocsCloudSecurityConfigRequest) UnsetSecret() {
	o.Secret.Unset()
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudSecurityConfigRequest) GetHeader() string {
	if o == nil || IsNil(o.Header.Get()) {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudSecurityConfigRequest) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *DocsCloudSecurityConfigRequest) IsHeaderSet() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
func (o *DocsCloudSecurityConfigRequest) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *DocsCloudSecurityConfigRequest) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *DocsCloudSecurityConfigRequest) UnsetHeader() {
	o.Header.Unset()
}

func (o DocsCloudSecurityConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudSecurityConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret.IsSet() {
		toSerialize["secret"] = o.Secret.Get()
	}
	if o.Header.IsSet() {
		toSerialize["header"] = o.Header.Get()
	}
	return toSerialize, nil
}

type NullableDocsCloudSecurityConfigRequest struct {
	value *DocsCloudSecurityConfigRequest
	isSet bool
}

func (v NullableDocsCloudSecurityConfigRequest) Get() *DocsCloudSecurityConfigRequest {
	return v.value
}

func (v *NullableDocsCloudSecurityConfigRequest) Set(val *DocsCloudSecurityConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudSecurityConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudSecurityConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudSecurityConfigRequest(val *DocsCloudSecurityConfigRequest) *NullableDocsCloudSecurityConfigRequest {
	return &NullableDocsCloudSecurityConfigRequest{value: val, isSet: true}
}

func (v NullableDocsCloudSecurityConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudSecurityConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

