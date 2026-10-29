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

// checks if the AiMetaGenerate402Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaGenerate402Response{}

// AiMetaGenerate402Response struct for AiMetaGenerate402Response
type AiMetaGenerate402Response struct {
	// Short reason, ready to be shown to the caller.
	Error string `json:"error"`
	// Stable classification of the failure: `rate_limit`, `insufficient_funds`, `auth`, `model_not_found`, `timeout`, `network`, `bad_request`, `server`, `unknown`, or a provider-specific code relayed as it came.
	Code string `json:"code"`
	// The provider's own error message, cut to 500 characters. Absent when the failure never reached the provider, or was a timeout.
	ProviderMessage *string `json:"providerMessage,omitempty"`
}

type _AiMetaGenerate402Response AiMetaGenerate402Response

// NewAiMetaGenerate402Response instantiates a new AiMetaGenerate402Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaGenerate402Response(error_ string, code string) *AiMetaGenerate402Response {
	this := AiMetaGenerate402Response{}
	this.Error = error_
	this.Code = code
	return &this
}

// NewAiMetaGenerate402ResponseWithDefaults instantiates a new AiMetaGenerate402Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaGenerate402ResponseWithDefaults() *AiMetaGenerate402Response {
	this := AiMetaGenerate402Response{}
	return &this
}

// GetError returns the Error field value
func (o *AiMetaGenerate402Response) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate402Response) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AiMetaGenerate402Response) SetError(v string) {
	o.Error = v
}

// GetCode returns the Code field value
func (o *AiMetaGenerate402Response) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate402Response) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AiMetaGenerate402Response) SetCode(v string) {
	o.Code = v
}

// GetProviderMessage returns the ProviderMessage field value if set, zero value otherwise.
func (o *AiMetaGenerate402Response) GetProviderMessage() string {
	if o == nil || IsNil(o.ProviderMessage) {
		var ret string
		return ret
	}
	return *o.ProviderMessage
}

// GetProviderMessageOk returns a tuple with the ProviderMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate402Response) GetProviderMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderMessage) {
		return nil, false
	}
	return o.ProviderMessage, true
}

// HasProviderMessage returns a boolean if a field has been set.
func (o *AiMetaGenerate402Response) IsProviderMessageSet() bool {
	if o != nil && !IsNil(o.ProviderMessage) {
		return true
	}

	return false
}

// SetProviderMessage gets a reference to the given string and assigns it to the ProviderMessage field.
func (o *AiMetaGenerate402Response) SetProviderMessage(v string) {
	o.ProviderMessage = &v
}

func (o AiMetaGenerate402Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaGenerate402Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["code"] = o.Code
	if !IsNil(o.ProviderMessage) {
		toSerialize["providerMessage"] = o.ProviderMessage
	}
	return toSerialize, nil
}

func (o *AiMetaGenerate402Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varAiMetaGenerate402Response := _AiMetaGenerate402Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaGenerate402Response)

	if err != nil {
		return err
	}

	*o = AiMetaGenerate402Response(varAiMetaGenerate402Response)

	return err
}

type NullableAiMetaGenerate402Response struct {
	value *AiMetaGenerate402Response
	isSet bool
}

func (v NullableAiMetaGenerate402Response) Get() *AiMetaGenerate402Response {
	return v.value
}

func (v *NullableAiMetaGenerate402Response) Set(val *AiMetaGenerate402Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaGenerate402Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaGenerate402Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaGenerate402Response(val *AiMetaGenerate402Response) *NullableAiMetaGenerate402Response {
	return &NullableAiMetaGenerate402Response{value: val, isSet: true}
}

func (v NullableAiMetaGenerate402Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaGenerate402Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

