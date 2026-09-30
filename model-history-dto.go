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

// checks if the HistoryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryDto{}

// HistoryDto One record of the activity log of a file or a folder.
type HistoryDto struct {
	// The identifier of the record, which tells two records of the same action apart and stays stable as long as the  portal keeps the log.
	Id int32 `json:"id"`
	// What happened - the kind of event the record stands for, such as a file being uploaded, renamed, moved or  shared - with the key a client can key its own wording off.
	Action HistoryAction `json:"action"`
	// Who caused the event. For an event caused by a visitor following an external link only the name they gave is  filled in, the account fields staying empty.
	Initiator EmployeeDto `json:"initiator"`
	// When the event happened, written with the offset of the portal's time zone.
	Date ApiDateTime `json:"date"`
	// The history data. Absent for actions that carry no payload of their own - changing a room's  logo, icon colour or cover, whose interpreter returns no data (see  `RoomLogoChangedInterpreter`). It used to be declared required, which put it in the  OpenAPI document's required list while the null-dropping serializer left it out of the  response, so a generated client threw on any history page holding one of those entries.
	Data *HistoryData `json:"data,omitempty"`
	// The records folded into this one because they belong to the same action, the separate files of one upload for  instance. It is empty when the record stands alone, and the records inside it carry no further nesting.
	Related []HistoryDto `json:"related,omitempty"`
}

type _HistoryDto HistoryDto

// NewHistoryDto instantiates a new HistoryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryDto(id int32, action HistoryAction, initiator EmployeeDto, date ApiDateTime) *HistoryDto {
	this := HistoryDto{}
	this.Id = id
	this.Action = action
	this.Initiator = initiator
	this.Date = date
	return &this
}

// NewHistoryDtoWithDefaults instantiates a new HistoryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryDtoWithDefaults() *HistoryDto {
	this := HistoryDto{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryDto) SetId(v int32) {
	o.Id = v
}

// GetAction returns the Action field value
func (o *HistoryDto) GetAction() HistoryAction {
	if o == nil {
		var ret HistoryAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryDto) GetActionOk() (*HistoryAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryDto) SetAction(v HistoryAction) {
	o.Action = v
}

// GetInitiator returns the Initiator field value
func (o *HistoryDto) GetInitiator() EmployeeDto {
	if o == nil {
		var ret EmployeeDto
		return ret
	}

	return o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value
// and a boolean to check if the value has been set.
func (o *HistoryDto) GetInitiatorOk() (*EmployeeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Initiator, true
}

// SetInitiator sets field value
func (o *HistoryDto) SetInitiator(v EmployeeDto) {
	o.Initiator = v
}

// GetDate returns the Date field value
func (o *HistoryDto) GetDate() ApiDateTime {
	if o == nil {
		var ret ApiDateTime
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *HistoryDto) GetDateOk() (*ApiDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *HistoryDto) SetDate(v ApiDateTime) {
	o.Date = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HistoryDto) GetData() HistoryData {
	if o == nil || IsNil(o.Data) {
		var ret HistoryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryDto) GetDataOk() (*HistoryData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoryDto) IsDataSet() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given HistoryData and assigns it to the Data field.
func (o *HistoryDto) SetData(v HistoryData) {
	o.Data = &v
}

// GetRelated returns the Related field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryDto) GetRelated() []HistoryDto {
	if o == nil {
		var ret []HistoryDto
		return ret
	}
	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryDto) GetRelatedOk() ([]HistoryDto, bool) {
	if o == nil || IsNil(o.Related) {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *HistoryDto) IsRelatedSet() bool {
	if o != nil && !IsNil(o.Related) {
		return true
	}

	return false
}

// SetRelated gets a reference to the given []HistoryDto and assigns it to the Related field.
func (o *HistoryDto) SetRelated(v []HistoryDto) {
	o.Related = v
}

func (o HistoryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["action"] = o.Action
	toSerialize["initiator"] = o.Initiator
	toSerialize["date"] = o.Date
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Related != nil {
		toSerialize["related"] = o.Related
	}
	return toSerialize, nil
}

func (o *HistoryDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"action",
		"initiator",
		"date",
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

	varHistoryDto := _HistoryDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varHistoryDto)

	if err != nil {
		return err
	}

	*o = HistoryDto(varHistoryDto)

	return err
}

type NullableHistoryDto struct {
	value *HistoryDto
	isSet bool
}

func (v NullableHistoryDto) Get() *HistoryDto {
	return v.value
}

func (v *NullableHistoryDto) Set(val *HistoryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryDto(val *HistoryDto) *NullableHistoryDto {
	return &NullableHistoryDto{value: val, isSet: true}
}

func (v NullableHistoryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

