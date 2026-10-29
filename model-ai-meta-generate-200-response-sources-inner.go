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

// checks if the AiMetaGenerate200ResponseSourcesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaGenerate200ResponseSourcesInner{}

// AiMetaGenerate200ResponseSourcesInner struct for AiMetaGenerate200ResponseSourcesInner
type AiMetaGenerate200ResponseSourcesInner struct {
	Id string `json:"id"`
	Title string `json:"title"`
	// Whether the document was cut to fit the model's budget.
	Truncated bool `json:"truncated"`
}

type _AiMetaGenerate200ResponseSourcesInner AiMetaGenerate200ResponseSourcesInner

// NewAiMetaGenerate200ResponseSourcesInner instantiates a new AiMetaGenerate200ResponseSourcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaGenerate200ResponseSourcesInner(id string, title string, truncated bool) *AiMetaGenerate200ResponseSourcesInner {
	this := AiMetaGenerate200ResponseSourcesInner{}
	this.Id = id
	this.Title = title
	this.Truncated = truncated
	return &this
}

// NewAiMetaGenerate200ResponseSourcesInnerWithDefaults instantiates a new AiMetaGenerate200ResponseSourcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaGenerate200ResponseSourcesInnerWithDefaults() *AiMetaGenerate200ResponseSourcesInner {
	this := AiMetaGenerate200ResponseSourcesInner{}
	return &this
}

// GetId returns the Id field value
func (o *AiMetaGenerate200ResponseSourcesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200ResponseSourcesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiMetaGenerate200ResponseSourcesInner) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *AiMetaGenerate200ResponseSourcesInner) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200ResponseSourcesInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AiMetaGenerate200ResponseSourcesInner) SetTitle(v string) {
	o.Title = v
}

// GetTruncated returns the Truncated field value
func (o *AiMetaGenerate200ResponseSourcesInner) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200ResponseSourcesInner) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value
func (o *AiMetaGenerate200ResponseSourcesInner) SetTruncated(v bool) {
	o.Truncated = v
}

func (o AiMetaGenerate200ResponseSourcesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaGenerate200ResponseSourcesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["truncated"] = o.Truncated
	return toSerialize, nil
}

func (o *AiMetaGenerate200ResponseSourcesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"truncated",
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

	varAiMetaGenerate200ResponseSourcesInner := _AiMetaGenerate200ResponseSourcesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaGenerate200ResponseSourcesInner)

	if err != nil {
		return err
	}

	*o = AiMetaGenerate200ResponseSourcesInner(varAiMetaGenerate200ResponseSourcesInner)

	return err
}

type NullableAiMetaGenerate200ResponseSourcesInner struct {
	value *AiMetaGenerate200ResponseSourcesInner
	isSet bool
}

func (v NullableAiMetaGenerate200ResponseSourcesInner) Get() *AiMetaGenerate200ResponseSourcesInner {
	return v.value
}

func (v *NullableAiMetaGenerate200ResponseSourcesInner) Set(val *AiMetaGenerate200ResponseSourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaGenerate200ResponseSourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaGenerate200ResponseSourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaGenerate200ResponseSourcesInner(val *AiMetaGenerate200ResponseSourcesInner) *NullableAiMetaGenerate200ResponseSourcesInner {
	return &NullableAiMetaGenerate200ResponseSourcesInner{value: val, isSet: true}
}

func (v NullableAiMetaGenerate200ResponseSourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaGenerate200ResponseSourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

