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

// checks if the EditorToolDecisionRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditorToolDecisionRequestBody{}

// EditorToolDecisionRequestBody Parameters for an editor file-generation tool decision.
type EditorToolDecisionRequestBody struct {
	// Whether the user approved creating the file.
	Allow *bool `json:"allow,omitempty"`
}

// NewEditorToolDecisionRequestBody instantiates a new EditorToolDecisionRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditorToolDecisionRequestBody() *EditorToolDecisionRequestBody {
	this := EditorToolDecisionRequestBody{}
	return &this
}

// NewEditorToolDecisionRequestBodyWithDefaults instantiates a new EditorToolDecisionRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditorToolDecisionRequestBodyWithDefaults() *EditorToolDecisionRequestBody {
	this := EditorToolDecisionRequestBody{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *EditorToolDecisionRequestBody) GetAllow() bool {
	if o == nil || IsNil(o.Allow) {
		var ret bool
		return ret
	}
	return *o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditorToolDecisionRequestBody) GetAllowOk() (*bool, bool) {
	if o == nil || IsNil(o.Allow) {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *EditorToolDecisionRequestBody) IsAllowSet() bool {
	if o != nil && !IsNil(o.Allow) {
		return true
	}

	return false
}

// SetAllow gets a reference to the given bool and assigns it to the Allow field.
func (o *EditorToolDecisionRequestBody) SetAllow(v bool) {
	o.Allow = &v
}

func (o EditorToolDecisionRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditorToolDecisionRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allow) {
		toSerialize["allow"] = o.Allow
	}
	return toSerialize, nil
}

type NullableEditorToolDecisionRequestBody struct {
	value *EditorToolDecisionRequestBody
	isSet bool
}

func (v NullableEditorToolDecisionRequestBody) Get() *EditorToolDecisionRequestBody {
	return v.value
}

func (v *NullableEditorToolDecisionRequestBody) Set(val *EditorToolDecisionRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEditorToolDecisionRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEditorToolDecisionRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditorToolDecisionRequestBody(val *EditorToolDecisionRequestBody) *NullableEditorToolDecisionRequestBody {
	return &NullableEditorToolDecisionRequestBody{value: val, isSet: true}
}

func (v NullableEditorToolDecisionRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditorToolDecisionRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

