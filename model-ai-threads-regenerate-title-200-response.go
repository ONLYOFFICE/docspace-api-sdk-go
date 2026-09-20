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

// checks if the AiThreadsRegenerateTitle200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadsRegenerateTitle200Response{}

// AiThreadsRegenerateTitle200Response struct for AiThreadsRegenerateTitle200Response
type AiThreadsRegenerateTitle200Response struct {
	// The regenerated thread title.
	Title string `json:"title"`
}

type _AiThreadsRegenerateTitle200Response AiThreadsRegenerateTitle200Response

// NewAiThreadsRegenerateTitle200Response instantiates a new AiThreadsRegenerateTitle200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadsRegenerateTitle200Response(title string) *AiThreadsRegenerateTitle200Response {
	this := AiThreadsRegenerateTitle200Response{}
	this.Title = title
	return &this
}

// NewAiThreadsRegenerateTitle200ResponseWithDefaults instantiates a new AiThreadsRegenerateTitle200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadsRegenerateTitle200ResponseWithDefaults() *AiThreadsRegenerateTitle200Response {
	this := AiThreadsRegenerateTitle200Response{}
	return &this
}

// GetTitle returns the Title field value
func (o *AiThreadsRegenerateTitle200Response) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiThreadsRegenerateTitle200Response) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AiThreadsRegenerateTitle200Response) SetTitle(v string) {
	o.Title = v
}

func (o AiThreadsRegenerateTitle200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadsRegenerateTitle200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *AiThreadsRegenerateTitle200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varAiThreadsRegenerateTitle200Response := _AiThreadsRegenerateTitle200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadsRegenerateTitle200Response)

	if err != nil {
		return err
	}

	*o = AiThreadsRegenerateTitle200Response(varAiThreadsRegenerateTitle200Response)

	return err
}

type NullableAiThreadsRegenerateTitle200Response struct {
	value *AiThreadsRegenerateTitle200Response
	isSet bool
}

func (v NullableAiThreadsRegenerateTitle200Response) Get() *AiThreadsRegenerateTitle200Response {
	return v.value
}

func (v *NullableAiThreadsRegenerateTitle200Response) Set(val *AiThreadsRegenerateTitle200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadsRegenerateTitle200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadsRegenerateTitle200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadsRegenerateTitle200Response(val *AiThreadsRegenerateTitle200Response) *NullableAiThreadsRegenerateTitle200Response {
	return &NullableAiThreadsRegenerateTitle200Response{value: val, isSet: true}
}

func (v NullableAiThreadsRegenerateTitle200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadsRegenerateTitle200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

