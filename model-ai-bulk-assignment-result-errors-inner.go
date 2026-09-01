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

// checks if the AiBulkAssignmentResultErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiBulkAssignmentResultErrorsInner{}

// AiBulkAssignmentResultErrorsInner struct for AiBulkAssignmentResultErrorsInner
type AiBulkAssignmentResultErrorsInner struct {
	ActionType AiActionType `json:"actionType"`
	Error AiTErrorData `json:"error"`
}

type _AiBulkAssignmentResultErrorsInner AiBulkAssignmentResultErrorsInner

// NewAiBulkAssignmentResultErrorsInner instantiates a new AiBulkAssignmentResultErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiBulkAssignmentResultErrorsInner(actionType AiActionType, error_ AiTErrorData) *AiBulkAssignmentResultErrorsInner {
	this := AiBulkAssignmentResultErrorsInner{}
	this.ActionType = actionType
	this.Error = error_
	return &this
}

// NewAiBulkAssignmentResultErrorsInnerWithDefaults instantiates a new AiBulkAssignmentResultErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiBulkAssignmentResultErrorsInnerWithDefaults() *AiBulkAssignmentResultErrorsInner {
	this := AiBulkAssignmentResultErrorsInner{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *AiBulkAssignmentResultErrorsInner) GetActionType() AiActionType {
	if o == nil {
		var ret AiActionType
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *AiBulkAssignmentResultErrorsInner) GetActionTypeOk() (*AiActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *AiBulkAssignmentResultErrorsInner) SetActionType(v AiActionType) {
	o.ActionType = v
}

// GetError returns the Error field value
func (o *AiBulkAssignmentResultErrorsInner) GetError() AiTErrorData {
	if o == nil {
		var ret AiTErrorData
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AiBulkAssignmentResultErrorsInner) GetErrorOk() (*AiTErrorData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AiBulkAssignmentResultErrorsInner) SetError(v AiTErrorData) {
	o.Error = v
}

func (o AiBulkAssignmentResultErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiBulkAssignmentResultErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionType"] = o.ActionType
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *AiBulkAssignmentResultErrorsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actionType",
		"error",
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

	varAiBulkAssignmentResultErrorsInner := _AiBulkAssignmentResultErrorsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiBulkAssignmentResultErrorsInner)

	if err != nil {
		return err
	}

	*o = AiBulkAssignmentResultErrorsInner(varAiBulkAssignmentResultErrorsInner)

	return err
}

type NullableAiBulkAssignmentResultErrorsInner struct {
	value *AiBulkAssignmentResultErrorsInner
	isSet bool
}

func (v NullableAiBulkAssignmentResultErrorsInner) Get() *AiBulkAssignmentResultErrorsInner {
	return v.value
}

func (v *NullableAiBulkAssignmentResultErrorsInner) Set(val *AiBulkAssignmentResultErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiBulkAssignmentResultErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiBulkAssignmentResultErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiBulkAssignmentResultErrorsInner(val *AiBulkAssignmentResultErrorsInner) *NullableAiBulkAssignmentResultErrorsInner {
	return &NullableAiBulkAssignmentResultErrorsInner{value: val, isSet: true}
}

func (v NullableAiBulkAssignmentResultErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiBulkAssignmentResultErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

