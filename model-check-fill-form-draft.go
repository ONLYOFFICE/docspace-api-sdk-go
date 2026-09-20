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

// checks if the CheckFillFormDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckFillFormDraft{}

// CheckFillFormDraft The revision of the form to open and what the caller intends to do with it.
type CheckFillFormDraft struct {
	// The revision of the form to open. Pass 0 for the current revision; a positive number addresses that entry of  the file history and is accepted only from a caller who may read the history, so a member who only has  fill-forms access must send 0.
	Version int32 `json:"version"`
	// What the caller intends to do with the form. `view` asks for a read-only address and `embedded` for an address  to be shown inside a frame; both only resolve the address and leave the file untouched. Leave it out to enter  the filling flow, where the personal draft is created or reused. The value is matched case-insensitively, and  anything else behaves like an empty value.
	Action NullableString `json:"action,omitempty"`
	// Whether the caller asked for a read-only address. The server derives it from `action` being `view` and ignores  any value sent with the request.
	RequestView *bool `json:"requestView,omitempty"`
	// Whether the caller asked for an address to be shown inside a frame. The server derives it from `action` being  `embedded` and ignores any value sent with the request.
	RequestEmbedded *bool `json:"requestEmbedded,omitempty"`
}

type _CheckFillFormDraft CheckFillFormDraft

// NewCheckFillFormDraft instantiates a new CheckFillFormDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckFillFormDraft(version int32) *CheckFillFormDraft {
	this := CheckFillFormDraft{}
	this.Version = version
	return &this
}

// NewCheckFillFormDraftWithDefaults instantiates a new CheckFillFormDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckFillFormDraftWithDefaults() *CheckFillFormDraft {
	this := CheckFillFormDraft{}
	return &this
}

// GetVersion returns the Version field value
func (o *CheckFillFormDraft) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CheckFillFormDraft) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CheckFillFormDraft) SetVersion(v int32) {
	o.Version = v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckFillFormDraft) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckFillFormDraft) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *CheckFillFormDraft) IsActionSet() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *CheckFillFormDraft) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *CheckFillFormDraft) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *CheckFillFormDraft) UnsetAction() {
	o.Action.Unset()
}

// GetRequestView returns the RequestView field value if set, zero value otherwise.
func (o *CheckFillFormDraft) GetRequestView() bool {
	if o == nil || IsNil(o.RequestView) {
		var ret bool
		return ret
	}
	return *o.RequestView
}

// GetRequestViewOk returns a tuple with the RequestView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckFillFormDraft) GetRequestViewOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestView) {
		return nil, false
	}
	return o.RequestView, true
}

// HasRequestView returns a boolean if a field has been set.
func (o *CheckFillFormDraft) IsRequestViewSet() bool {
	if o != nil && !IsNil(o.RequestView) {
		return true
	}

	return false
}

// SetRequestView gets a reference to the given bool and assigns it to the RequestView field.
func (o *CheckFillFormDraft) SetRequestView(v bool) {
	o.RequestView = &v
}

// GetRequestEmbedded returns the RequestEmbedded field value if set, zero value otherwise.
func (o *CheckFillFormDraft) GetRequestEmbedded() bool {
	if o == nil || IsNil(o.RequestEmbedded) {
		var ret bool
		return ret
	}
	return *o.RequestEmbedded
}

// GetRequestEmbeddedOk returns a tuple with the RequestEmbedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckFillFormDraft) GetRequestEmbeddedOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestEmbedded) {
		return nil, false
	}
	return o.RequestEmbedded, true
}

// HasRequestEmbedded returns a boolean if a field has been set.
func (o *CheckFillFormDraft) IsRequestEmbeddedSet() bool {
	if o != nil && !IsNil(o.RequestEmbedded) {
		return true
	}

	return false
}

// SetRequestEmbedded gets a reference to the given bool and assigns it to the RequestEmbedded field.
func (o *CheckFillFormDraft) SetRequestEmbedded(v bool) {
	o.RequestEmbedded = &v
}

func (o CheckFillFormDraft) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckFillFormDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if !IsNil(o.RequestView) {
		toSerialize["requestView"] = o.RequestView
	}
	if !IsNil(o.RequestEmbedded) {
		toSerialize["requestEmbedded"] = o.RequestEmbedded
	}
	return toSerialize, nil
}

func (o *CheckFillFormDraft) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
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

	varCheckFillFormDraft := _CheckFillFormDraft{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCheckFillFormDraft)

	if err != nil {
		return err
	}

	*o = CheckFillFormDraft(varCheckFillFormDraft)

	return err
}

type NullableCheckFillFormDraft struct {
	value *CheckFillFormDraft
	isSet bool
}

func (v NullableCheckFillFormDraft) Get() *CheckFillFormDraft {
	return v.value
}

func (v *NullableCheckFillFormDraft) Set(val *CheckFillFormDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckFillFormDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckFillFormDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckFillFormDraft(val *CheckFillFormDraft) *NullableCheckFillFormDraft {
	return &NullableCheckFillFormDraft{value: val, isSet: true}
}

func (v NullableCheckFillFormDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckFillFormDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

