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

// checks if the OwnerChangeInstructionsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnerChangeInstructionsDto{}

// OwnerChangeInstructionsDto The outcome of asking for the portal-ownership transfer letter to be sent.
type OwnerChangeInstructionsDto struct {
	// Whether the letter was sent: `1` that it was, `0` that the request was turned down. A refusal comes back  with HTTP 200, so this field and not the status code is what says whether anything happened - the request  is turned down when the caller is not the portal owner and when the named member is unknown or inactive.
	Status *int32 `json:"status,omitempty"`
	// The outcome spelled out in the portal language. On success it names the address the letter went to, and it  carries an HTML `mailto:` anchor rather than plain text, so it has to be rendered as markup or stripped;  on a refusal it is the localised reason.
	Message NullableString `json:"message,omitempty"`
}

// NewOwnerChangeInstructionsDto instantiates a new OwnerChangeInstructionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerChangeInstructionsDto() *OwnerChangeInstructionsDto {
	this := OwnerChangeInstructionsDto{}
	return &this
}

// NewOwnerChangeInstructionsDtoWithDefaults instantiates a new OwnerChangeInstructionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerChangeInstructionsDtoWithDefaults() *OwnerChangeInstructionsDto {
	this := OwnerChangeInstructionsDto{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OwnerChangeInstructionsDto) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerChangeInstructionsDto) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OwnerChangeInstructionsDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *OwnerChangeInstructionsDto) SetStatus(v int32) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnerChangeInstructionsDto) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OwnerChangeInstructionsDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *OwnerChangeInstructionsDto) IsMessageSet() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *OwnerChangeInstructionsDto) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *OwnerChangeInstructionsDto) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *OwnerChangeInstructionsDto) UnsetMessage() {
	o.Message.Unset()
}

func (o OwnerChangeInstructionsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnerChangeInstructionsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableOwnerChangeInstructionsDto struct {
	value *OwnerChangeInstructionsDto
	isSet bool
}

func (v NullableOwnerChangeInstructionsDto) Get() *OwnerChangeInstructionsDto {
	return v.value
}

func (v *NullableOwnerChangeInstructionsDto) Set(val *OwnerChangeInstructionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerChangeInstructionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerChangeInstructionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerChangeInstructionsDto(val *OwnerChangeInstructionsDto) *NullableOwnerChangeInstructionsDto {
	return &NullableOwnerChangeInstructionsDto{value: val, isSet: true}
}

func (v NullableOwnerChangeInstructionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerChangeInstructionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

