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

// checks if the AiMetaGenerate200ResponseSkippedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaGenerate200ResponseSkippedInner{}

// AiMetaGenerate200ResponseSkippedInner struct for AiMetaGenerate200ResponseSkippedInner
type AiMetaGenerate200ResponseSkippedInner struct {
	Id string `json:"id"`
	Title *string `json:"title,omitempty"`
	Reason string `json:"reason"`
}

type _AiMetaGenerate200ResponseSkippedInner AiMetaGenerate200ResponseSkippedInner

// NewAiMetaGenerate200ResponseSkippedInner instantiates a new AiMetaGenerate200ResponseSkippedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaGenerate200ResponseSkippedInner(id string, reason string) *AiMetaGenerate200ResponseSkippedInner {
	this := AiMetaGenerate200ResponseSkippedInner{}
	this.Id = id
	this.Reason = reason
	return &this
}

// NewAiMetaGenerate200ResponseSkippedInnerWithDefaults instantiates a new AiMetaGenerate200ResponseSkippedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaGenerate200ResponseSkippedInnerWithDefaults() *AiMetaGenerate200ResponseSkippedInner {
	this := AiMetaGenerate200ResponseSkippedInner{}
	return &this
}

// GetId returns the Id field value
func (o *AiMetaGenerate200ResponseSkippedInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200ResponseSkippedInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiMetaGenerate200ResponseSkippedInner) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiMetaGenerate200ResponseSkippedInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200ResponseSkippedInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiMetaGenerate200ResponseSkippedInner) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiMetaGenerate200ResponseSkippedInner) SetTitle(v string) {
	o.Title = &v
}

// GetReason returns the Reason field value
func (o *AiMetaGenerate200ResponseSkippedInner) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200ResponseSkippedInner) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *AiMetaGenerate200ResponseSkippedInner) SetReason(v string) {
	o.Reason = v
}

func (o AiMetaGenerate200ResponseSkippedInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaGenerate200ResponseSkippedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *AiMetaGenerate200ResponseSkippedInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"reason",
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

	varAiMetaGenerate200ResponseSkippedInner := _AiMetaGenerate200ResponseSkippedInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaGenerate200ResponseSkippedInner)

	if err != nil {
		return err
	}

	*o = AiMetaGenerate200ResponseSkippedInner(varAiMetaGenerate200ResponseSkippedInner)

	return err
}

type NullableAiMetaGenerate200ResponseSkippedInner struct {
	value *AiMetaGenerate200ResponseSkippedInner
	isSet bool
}

func (v NullableAiMetaGenerate200ResponseSkippedInner) Get() *AiMetaGenerate200ResponseSkippedInner {
	return v.value
}

func (v *NullableAiMetaGenerate200ResponseSkippedInner) Set(val *AiMetaGenerate200ResponseSkippedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaGenerate200ResponseSkippedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaGenerate200ResponseSkippedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaGenerate200ResponseSkippedInner(val *AiMetaGenerate200ResponseSkippedInner) *NullableAiMetaGenerate200ResponseSkippedInner {
	return &NullableAiMetaGenerate200ResponseSkippedInner{value: val, isSet: true}
}

func (v NullableAiMetaGenerate200ResponseSkippedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaGenerate200ResponseSkippedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

