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

// checks if the AiMetaFill200ResponseUnresolvedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaFill200ResponseUnresolvedInner{}

// AiMetaFill200ResponseUnresolvedInner struct for AiMetaFill200ResponseUnresolvedInner
type AiMetaFill200ResponseUnresolvedInner struct {
	FieldId *float32 `json:"fieldId,omitempty"`
	Name string `json:"name"`
	Reason string `json:"reason"`
}

type _AiMetaFill200ResponseUnresolvedInner AiMetaFill200ResponseUnresolvedInner

// NewAiMetaFill200ResponseUnresolvedInner instantiates a new AiMetaFill200ResponseUnresolvedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaFill200ResponseUnresolvedInner(name string, reason string) *AiMetaFill200ResponseUnresolvedInner {
	this := AiMetaFill200ResponseUnresolvedInner{}
	this.Name = name
	this.Reason = reason
	return &this
}

// NewAiMetaFill200ResponseUnresolvedInnerWithDefaults instantiates a new AiMetaFill200ResponseUnresolvedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaFill200ResponseUnresolvedInnerWithDefaults() *AiMetaFill200ResponseUnresolvedInner {
	this := AiMetaFill200ResponseUnresolvedInner{}
	return &this
}

// GetFieldId returns the FieldId field value if set, zero value otherwise.
func (o *AiMetaFill200ResponseUnresolvedInner) GetFieldId() float32 {
	if o == nil || IsNil(o.FieldId) {
		var ret float32
		return ret
	}
	return *o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetaFill200ResponseUnresolvedInner) GetFieldIdOk() (*float32, bool) {
	if o == nil || IsNil(o.FieldId) {
		return nil, false
	}
	return o.FieldId, true
}

// HasFieldId returns a boolean if a field has been set.
func (o *AiMetaFill200ResponseUnresolvedInner) IsFieldIdSet() bool {
	if o != nil && !IsNil(o.FieldId) {
		return true
	}

	return false
}

// SetFieldId gets a reference to the given float32 and assigns it to the FieldId field.
func (o *AiMetaFill200ResponseUnresolvedInner) SetFieldId(v float32) {
	o.FieldId = &v
}

// GetName returns the Name field value
func (o *AiMetaFill200ResponseUnresolvedInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200ResponseUnresolvedInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiMetaFill200ResponseUnresolvedInner) SetName(v string) {
	o.Name = v
}

// GetReason returns the Reason field value
func (o *AiMetaFill200ResponseUnresolvedInner) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200ResponseUnresolvedInner) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *AiMetaFill200ResponseUnresolvedInner) SetReason(v string) {
	o.Reason = v
}

func (o AiMetaFill200ResponseUnresolvedInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaFill200ResponseUnresolvedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldId) {
		toSerialize["fieldId"] = o.FieldId
	}
	toSerialize["name"] = o.Name
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *AiMetaFill200ResponseUnresolvedInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAiMetaFill200ResponseUnresolvedInner := _AiMetaFill200ResponseUnresolvedInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaFill200ResponseUnresolvedInner)

	if err != nil {
		return err
	}

	*o = AiMetaFill200ResponseUnresolvedInner(varAiMetaFill200ResponseUnresolvedInner)

	return err
}

type NullableAiMetaFill200ResponseUnresolvedInner struct {
	value *AiMetaFill200ResponseUnresolvedInner
	isSet bool
}

func (v NullableAiMetaFill200ResponseUnresolvedInner) Get() *AiMetaFill200ResponseUnresolvedInner {
	return v.value
}

func (v *NullableAiMetaFill200ResponseUnresolvedInner) Set(val *AiMetaFill200ResponseUnresolvedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaFill200ResponseUnresolvedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaFill200ResponseUnresolvedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaFill200ResponseUnresolvedInner(val *AiMetaFill200ResponseUnresolvedInner) *NullableAiMetaFill200ResponseUnresolvedInner {
	return &NullableAiMetaFill200ResponseUnresolvedInner{value: val, isSet: true}
}

func (v NullableAiMetaFill200ResponseUnresolvedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaFill200ResponseUnresolvedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

