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

// checks if the ConfirmDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmDto{}

// ConfirmDto Whether a confirmation link may still be used, and what it leads to when it invites into a room.
type ConfirmDto struct {
	// The outcome of the check. Only `Ok` means the action behind the link may be carried out: `Invalid` and  `Expired` fault the key itself, while `UserExisted`, `UserExcluded`, `TariffLimit` and `QuotaFailed` mean  the key is sound but the invitation behind it cannot be accepted as it stands.
	Result ValidationResult `json:"result"`
	// The room the invitation leads into - a numeric folder ID for a room of the portal, a provider-specific  string for a third-party one. It is empty for an invitation to the portal as a whole, for a room that has  been removed or that the invited account may not see, and whenever `result` is neither `Ok` nor  `UserExisted`.
	RoomId NullableString `json:"roomId,omitempty"`
	// The title of that room, present exactly when `roomId` is and meant to be shown on the confirmation page.
	Title NullableString `json:"title,omitempty"`
	// The address the link was issued for, echoed back only when `result` is `Ok` so that a sign-up form can be  prefilled with it. Every other outcome leaves it empty, `UserExisted` included.
	Email NullableString `json:"email,omitempty"`
	// Whether the room behind the link is an AI room rather than an ordinary one, which decides where the invited  person is taken. It is `false` whenever `roomId` is empty.
	IsAgent *bool `json:"isAgent,omitempty"`
}

type _ConfirmDto ConfirmDto

// NewConfirmDto instantiates a new ConfirmDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmDto(result ValidationResult) *ConfirmDto {
	this := ConfirmDto{}
	this.Result = result
	return &this
}

// NewConfirmDtoWithDefaults instantiates a new ConfirmDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmDtoWithDefaults() *ConfirmDto {
	this := ConfirmDto{}
	return &this
}

// GetResult returns the Result field value
func (o *ConfirmDto) GetResult() ValidationResult {
	if o == nil {
		var ret ValidationResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ConfirmDto) GetResultOk() (*ValidationResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ConfirmDto) SetResult(v ValidationResult) {
	o.Result = v
}

// GetRoomId returns the RoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfirmDto) GetRoomId() string {
	if o == nil || IsNil(o.RoomId.Get()) {
		var ret string
		return ret
	}
	return *o.RoomId.Get()
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfirmDto) GetRoomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomId.Get(), o.RoomId.IsSet()
}

// HasRoomId returns a boolean if a field has been set.
func (o *ConfirmDto) IsRoomIdSet() bool {
	if o != nil && o.RoomId.IsSet() {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given NullableString and assigns it to the RoomId field.
func (o *ConfirmDto) SetRoomId(v string) {
	o.RoomId.Set(&v)
}
// SetRoomIdNil sets the value for RoomId to be an explicit nil
func (o *ConfirmDto) SetRoomIdNil() {
	o.RoomId.Set(nil)
}

// UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
func (o *ConfirmDto) UnsetRoomId() {
	o.RoomId.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfirmDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfirmDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ConfirmDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ConfirmDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ConfirmDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ConfirmDto) UnsetTitle() {
	o.Title.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfirmDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfirmDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ConfirmDto) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *ConfirmDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *ConfirmDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *ConfirmDto) UnsetEmail() {
	o.Email.Unset()
}

// GetIsAgent returns the IsAgent field value if set, zero value otherwise.
func (o *ConfirmDto) GetIsAgent() bool {
	if o == nil || IsNil(o.IsAgent) {
		var ret bool
		return ret
	}
	return *o.IsAgent
}

// GetIsAgentOk returns a tuple with the IsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmDto) GetIsAgentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAgent) {
		return nil, false
	}
	return o.IsAgent, true
}

// HasIsAgent returns a boolean if a field has been set.
func (o *ConfirmDto) IsIsAgentSet() bool {
	if o != nil && !IsNil(o.IsAgent) {
		return true
	}

	return false
}

// SetIsAgent gets a reference to the given bool and assigns it to the IsAgent field.
func (o *ConfirmDto) SetIsAgent(v bool) {
	o.IsAgent = &v
}

func (o ConfirmDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfirmDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	if o.RoomId.IsSet() {
		toSerialize["roomId"] = o.RoomId.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.IsAgent) {
		toSerialize["isAgent"] = o.IsAgent
	}
	return toSerialize, nil
}

func (o *ConfirmDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varConfirmDto := _ConfirmDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varConfirmDto)

	if err != nil {
		return err
	}

	*o = ConfirmDto(varConfirmDto)

	return err
}

type NullableConfirmDto struct {
	value *ConfirmDto
	isSet bool
}

func (v NullableConfirmDto) Get() *ConfirmDto {
	return v.value
}

func (v *NullableConfirmDto) Set(val *ConfirmDto) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmDto) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmDto(val *ConfirmDto) *NullableConfirmDto {
	return &NullableConfirmDto{value: val, isSet: true}
}

func (v NullableConfirmDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

