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

// checks if the UnlinkAiProjectSourcesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnlinkAiProjectSourcesRequest{}

// UnlinkAiProjectSourcesRequest Rooms to unlink from the AI project.
type UnlinkAiProjectSourcesRequest struct {
	// IDs of the rooms to unlink. IDs that are not linked to the project are ignored.
	RoomIds []int32 `json:"roomIds"`
}

type _UnlinkAiProjectSourcesRequest UnlinkAiProjectSourcesRequest

// NewUnlinkAiProjectSourcesRequest instantiates a new UnlinkAiProjectSourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkAiProjectSourcesRequest(roomIds []int32) *UnlinkAiProjectSourcesRequest {
	this := UnlinkAiProjectSourcesRequest{}
	this.RoomIds = roomIds
	return &this
}

// NewUnlinkAiProjectSourcesRequestWithDefaults instantiates a new UnlinkAiProjectSourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkAiProjectSourcesRequestWithDefaults() *UnlinkAiProjectSourcesRequest {
	this := UnlinkAiProjectSourcesRequest{}
	return &this
}

// GetRoomIds returns the RoomIds field value
// If the value is explicit nil, the zero value for []int32 will be returned
func (o *UnlinkAiProjectSourcesRequest) GetRoomIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.RoomIds
}

// GetRoomIdsOk returns a tuple with the RoomIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnlinkAiProjectSourcesRequest) GetRoomIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.RoomIds) {
		return nil, false
	}
	return o.RoomIds, true
}

// SetRoomIds sets field value
func (o *UnlinkAiProjectSourcesRequest) SetRoomIds(v []int32) {
	o.RoomIds = v
}

func (o UnlinkAiProjectSourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnlinkAiProjectSourcesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoomIds != nil {
		toSerialize["roomIds"] = o.RoomIds
	}
	return toSerialize, nil
}

func (o *UnlinkAiProjectSourcesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roomIds",
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

	varUnlinkAiProjectSourcesRequest := _UnlinkAiProjectSourcesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varUnlinkAiProjectSourcesRequest)

	if err != nil {
		return err
	}

	*o = UnlinkAiProjectSourcesRequest(varUnlinkAiProjectSourcesRequest)

	return err
}

type NullableUnlinkAiProjectSourcesRequest struct {
	value *UnlinkAiProjectSourcesRequest
	isSet bool
}

func (v NullableUnlinkAiProjectSourcesRequest) Get() *UnlinkAiProjectSourcesRequest {
	return v.value
}

func (v *NullableUnlinkAiProjectSourcesRequest) Set(val *UnlinkAiProjectSourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkAiProjectSourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkAiProjectSourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkAiProjectSourcesRequest(val *UnlinkAiProjectSourcesRequest) *NullableUnlinkAiProjectSourcesRequest {
	return &NullableUnlinkAiProjectSourcesRequest{value: val, isSet: true}
}

func (v NullableUnlinkAiProjectSourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkAiProjectSourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

