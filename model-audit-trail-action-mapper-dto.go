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

// checks if the AuditTrailActionMapperDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditTrailActionMapperDto{}

// AuditTrailActionMapperDto One audit trail action, with the kind of change it stands for and the kind of object it applies to.
type AuditTrailActionMapperDto struct {
	// The action name to send as the `action` filter of `GET api/2.0/security/audit/events/filter`, and the value  that comes back as `actionId` on an event.
	MessageAction NullableString `json:"messageAction,omitempty"`
	// The kind of change the action makes, accepted by the `actionType` filter of the same operation.
	ActionType NullableString `json:"actionType,omitempty"`
	// The kind of object the action applies to, accepted by the `entryType` filter. It is `None` for an action  that targets no object, such as a settings change, and an action with a second object type reports only the  first one here.
	Entity NullableString `json:"entity,omitempty"`
}

// NewAuditTrailActionMapperDto instantiates a new AuditTrailActionMapperDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditTrailActionMapperDto() *AuditTrailActionMapperDto {
	this := AuditTrailActionMapperDto{}
	return &this
}

// NewAuditTrailActionMapperDtoWithDefaults instantiates a new AuditTrailActionMapperDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditTrailActionMapperDtoWithDefaults() *AuditTrailActionMapperDto {
	this := AuditTrailActionMapperDto{}
	return &this
}

// GetMessageAction returns the MessageAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailActionMapperDto) GetMessageAction() string {
	if o == nil || IsNil(o.MessageAction.Get()) {
		var ret string
		return ret
	}
	return *o.MessageAction.Get()
}

// GetMessageActionOk returns a tuple with the MessageAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailActionMapperDto) GetMessageActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageAction.Get(), o.MessageAction.IsSet()
}

// HasMessageAction returns a boolean if a field has been set.
func (o *AuditTrailActionMapperDto) IsMessageActionSet() bool {
	if o != nil && o.MessageAction.IsSet() {
		return true
	}

	return false
}

// SetMessageAction gets a reference to the given NullableString and assigns it to the MessageAction field.
func (o *AuditTrailActionMapperDto) SetMessageAction(v string) {
	o.MessageAction.Set(&v)
}
// SetMessageActionNil sets the value for MessageAction to be an explicit nil
func (o *AuditTrailActionMapperDto) SetMessageActionNil() {
	o.MessageAction.Set(nil)
}

// UnsetMessageAction ensures that no value is present for MessageAction, not even an explicit nil
func (o *AuditTrailActionMapperDto) UnsetMessageAction() {
	o.MessageAction.Unset()
}

// GetActionType returns the ActionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailActionMapperDto) GetActionType() string {
	if o == nil || IsNil(o.ActionType.Get()) {
		var ret string
		return ret
	}
	return *o.ActionType.Get()
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailActionMapperDto) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionType.Get(), o.ActionType.IsSet()
}

// HasActionType returns a boolean if a field has been set.
func (o *AuditTrailActionMapperDto) IsActionTypeSet() bool {
	if o != nil && o.ActionType.IsSet() {
		return true
	}

	return false
}

// SetActionType gets a reference to the given NullableString and assigns it to the ActionType field.
func (o *AuditTrailActionMapperDto) SetActionType(v string) {
	o.ActionType.Set(&v)
}
// SetActionTypeNil sets the value for ActionType to be an explicit nil
func (o *AuditTrailActionMapperDto) SetActionTypeNil() {
	o.ActionType.Set(nil)
}

// UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
func (o *AuditTrailActionMapperDto) UnsetActionType() {
	o.ActionType.Unset()
}

// GetEntity returns the Entity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailActionMapperDto) GetEntity() string {
	if o == nil || IsNil(o.Entity.Get()) {
		var ret string
		return ret
	}
	return *o.Entity.Get()
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailActionMapperDto) GetEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entity.Get(), o.Entity.IsSet()
}

// HasEntity returns a boolean if a field has been set.
func (o *AuditTrailActionMapperDto) IsEntitySet() bool {
	if o != nil && o.Entity.IsSet() {
		return true
	}

	return false
}

// SetEntity gets a reference to the given NullableString and assigns it to the Entity field.
func (o *AuditTrailActionMapperDto) SetEntity(v string) {
	o.Entity.Set(&v)
}
// SetEntityNil sets the value for Entity to be an explicit nil
func (o *AuditTrailActionMapperDto) SetEntityNil() {
	o.Entity.Set(nil)
}

// UnsetEntity ensures that no value is present for Entity, not even an explicit nil
func (o *AuditTrailActionMapperDto) UnsetEntity() {
	o.Entity.Unset()
}

func (o AuditTrailActionMapperDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditTrailActionMapperDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MessageAction.IsSet() {
		toSerialize["messageAction"] = o.MessageAction.Get()
	}
	if o.ActionType.IsSet() {
		toSerialize["actionType"] = o.ActionType.Get()
	}
	if o.Entity.IsSet() {
		toSerialize["entity"] = o.Entity.Get()
	}
	return toSerialize, nil
}

type NullableAuditTrailActionMapperDto struct {
	value *AuditTrailActionMapperDto
	isSet bool
}

func (v NullableAuditTrailActionMapperDto) Get() *AuditTrailActionMapperDto {
	return v.value
}

func (v *NullableAuditTrailActionMapperDto) Set(val *AuditTrailActionMapperDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditTrailActionMapperDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditTrailActionMapperDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditTrailActionMapperDto(val *AuditTrailActionMapperDto) *NullableAuditTrailActionMapperDto {
	return &NullableAuditTrailActionMapperDto{value: val, isSet: true}
}

func (v NullableAuditTrailActionMapperDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditTrailActionMapperDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

