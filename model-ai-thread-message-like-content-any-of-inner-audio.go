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

// checks if the AiThreadMessageLikeContentAnyOfInnerAudio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadMessageLikeContentAnyOfInnerAudio{}

// AiThreadMessageLikeContentAnyOfInnerAudio An `audio` part's sound.
type AiThreadMessageLikeContentAnyOfInnerAudio struct {
	// The audio data, base64.
	Data string `json:"data"`
	// Audio format: `mp3` or `wav`.
	Format string `json:"format"`
}

type _AiThreadMessageLikeContentAnyOfInnerAudio AiThreadMessageLikeContentAnyOfInnerAudio

// NewAiThreadMessageLikeContentAnyOfInnerAudio instantiates a new AiThreadMessageLikeContentAnyOfInnerAudio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadMessageLikeContentAnyOfInnerAudio(data string, format string) *AiThreadMessageLikeContentAnyOfInnerAudio {
	this := AiThreadMessageLikeContentAnyOfInnerAudio{}
	this.Data = data
	this.Format = format
	return &this
}

// NewAiThreadMessageLikeContentAnyOfInnerAudioWithDefaults instantiates a new AiThreadMessageLikeContentAnyOfInnerAudio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadMessageLikeContentAnyOfInnerAudioWithDefaults() *AiThreadMessageLikeContentAnyOfInnerAudio {
	this := AiThreadMessageLikeContentAnyOfInnerAudio{}
	return &this
}

// GetData returns the Data field value
func (o *AiThreadMessageLikeContentAnyOfInnerAudio) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInnerAudio) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AiThreadMessageLikeContentAnyOfInnerAudio) SetData(v string) {
	o.Data = v
}

// GetFormat returns the Format field value
func (o *AiThreadMessageLikeContentAnyOfInnerAudio) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInnerAudio) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *AiThreadMessageLikeContentAnyOfInnerAudio) SetFormat(v string) {
	o.Format = v
}

func (o AiThreadMessageLikeContentAnyOfInnerAudio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadMessageLikeContentAnyOfInnerAudio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["format"] = o.Format
	return toSerialize, nil
}

func (o *AiThreadMessageLikeContentAnyOfInnerAudio) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"format",
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

	varAiThreadMessageLikeContentAnyOfInnerAudio := _AiThreadMessageLikeContentAnyOfInnerAudio{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadMessageLikeContentAnyOfInnerAudio)

	if err != nil {
		return err
	}

	*o = AiThreadMessageLikeContentAnyOfInnerAudio(varAiThreadMessageLikeContentAnyOfInnerAudio)

	return err
}

type NullableAiThreadMessageLikeContentAnyOfInnerAudio struct {
	value *AiThreadMessageLikeContentAnyOfInnerAudio
	isSet bool
}

func (v NullableAiThreadMessageLikeContentAnyOfInnerAudio) Get() *AiThreadMessageLikeContentAnyOfInnerAudio {
	return v.value
}

func (v *NullableAiThreadMessageLikeContentAnyOfInnerAudio) Set(val *AiThreadMessageLikeContentAnyOfInnerAudio) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadMessageLikeContentAnyOfInnerAudio) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadMessageLikeContentAnyOfInnerAudio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadMessageLikeContentAnyOfInnerAudio(val *AiThreadMessageLikeContentAnyOfInnerAudio) *NullableAiThreadMessageLikeContentAnyOfInnerAudio {
	return &NullableAiThreadMessageLikeContentAnyOfInnerAudio{value: val, isSet: true}
}

func (v NullableAiThreadMessageLikeContentAnyOfInnerAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadMessageLikeContentAnyOfInnerAudio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

