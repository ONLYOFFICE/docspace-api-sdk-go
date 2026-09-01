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
	"time"
	"bytes"
	"fmt"
)

// checks if the HistoryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryDto{}

// HistoryDto The file history information.
type HistoryDto struct {
	// The unique identifier for the file history entry.
	Id int32 `json:"id"`
	// The action performed on the file.
	Action HistoryAction `json:"action"`
	// The action initiator.
	Initiator EmployeeDto `json:"initiator"`
	// The date and time when an action on the file was performed.
	Date NullableTime `json:"date"`
	// The history data.
	Data HistoryData `json:"data"`
	// The list of related history.
	Related []HistoryDto `json:"related,omitempty"`
}

type _HistoryDto HistoryDto

// NewHistoryDto instantiates a new HistoryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryDto(id int32, action HistoryAction, initiator EmployeeDto, date NullableTime, data HistoryData) *HistoryDto {
	this := HistoryDto{}
	this.Id = id
	this.Action = action
	this.Initiator = initiator
	this.Date = date
	this.Data = data
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
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *HistoryDto) GetDate() time.Time {
	if o == nil || o.Date.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryDto) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// SetDate sets field value
func (o *HistoryDto) SetDate(v time.Time) {
	o.Date.Set(&v)
}

// GetData returns the Data field value
func (o *HistoryDto) GetData() HistoryData {
	if o == nil {
		var ret HistoryData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *HistoryDto) GetDataOk() (*HistoryData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *HistoryDto) SetData(v HistoryData) {
	o.Data = v
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
	toSerialize["date"] = o.Date.Get()
	toSerialize["data"] = o.Data
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
		"data",
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

