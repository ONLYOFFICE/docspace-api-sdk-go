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

// checks if the ManageFormFillingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManageFormFillingDto{}

// ManageFormFillingDto The action to apply to the filling of a PDF form.
type ManageFormFillingDto struct {
	// The PDF form the action applies to. This is the value the operation reads, rather than the identifier in its  route, and the two are to be sent the same.
	FormId int32 `json:"formId"`
	// The action to apply.
	Action *FormFillingManageAction `json:"action,omitempty"`
}

type _ManageFormFillingDto ManageFormFillingDto

// NewManageFormFillingDto instantiates a new ManageFormFillingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageFormFillingDto(formId int32) *ManageFormFillingDto {
	this := ManageFormFillingDto{}
	this.FormId = formId
	return &this
}

// NewManageFormFillingDtoWithDefaults instantiates a new ManageFormFillingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageFormFillingDtoWithDefaults() *ManageFormFillingDto {
	this := ManageFormFillingDto{}
	return &this
}

// GetFormId returns the FormId field value
func (o *ManageFormFillingDto) GetFormId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *ManageFormFillingDto) GetFormIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value
func (o *ManageFormFillingDto) SetFormId(v int32) {
	o.FormId = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ManageFormFillingDto) GetAction() FormFillingManageAction {
	if o == nil || IsNil(o.Action) {
		var ret FormFillingManageAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageFormFillingDto) GetActionOk() (*FormFillingManageAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ManageFormFillingDto) IsActionSet() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given FormFillingManageAction and assigns it to the Action field.
func (o *ManageFormFillingDto) SetAction(v FormFillingManageAction) {
	o.Action = &v
}

func (o ManageFormFillingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManageFormFillingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["formId"] = o.FormId
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

func (o *ManageFormFillingDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"formId",
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

	varManageFormFillingDto := _ManageFormFillingDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varManageFormFillingDto)

	if err != nil {
		return err
	}

	*o = ManageFormFillingDto(varManageFormFillingDto)

	return err
}

type NullableManageFormFillingDto struct {
	value *ManageFormFillingDto
	isSet bool
}

func (v NullableManageFormFillingDto) Get() *ManageFormFillingDto {
	return v.value
}

func (v *NullableManageFormFillingDto) Set(val *ManageFormFillingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableManageFormFillingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableManageFormFillingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageFormFillingDto(val *ManageFormFillingDto) *NullableManageFormFillingDto {
	return &NullableManageFormFillingDto{value: val, isSet: true}
}

func (v NullableManageFormFillingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageFormFillingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

