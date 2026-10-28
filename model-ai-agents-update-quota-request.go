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

// checks if the AiAgentsUpdateQuotaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAgentsUpdateQuotaRequest{}

// AiAgentsUpdateQuotaRequest struct for AiAgentsUpdateQuotaRequest
type AiAgentsUpdateQuotaRequest struct {
	// Agent (room) ids to update.
	RoomIds []AiAgentsUpdateQuotaRequestRoomIdsInner `json:"roomIds"`
	// New quota in bytes; a negative value disables the custom quota.
	Quota float32 `json:"quota"`
}

type _AiAgentsUpdateQuotaRequest AiAgentsUpdateQuotaRequest

// NewAiAgentsUpdateQuotaRequest instantiates a new AiAgentsUpdateQuotaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAgentsUpdateQuotaRequest(roomIds []AiAgentsUpdateQuotaRequestRoomIdsInner, quota float32) *AiAgentsUpdateQuotaRequest {
	this := AiAgentsUpdateQuotaRequest{}
	this.RoomIds = roomIds
	this.Quota = quota
	return &this
}

// NewAiAgentsUpdateQuotaRequestWithDefaults instantiates a new AiAgentsUpdateQuotaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAgentsUpdateQuotaRequestWithDefaults() *AiAgentsUpdateQuotaRequest {
	this := AiAgentsUpdateQuotaRequest{}
	return &this
}

// GetRoomIds returns the RoomIds field value
func (o *AiAgentsUpdateQuotaRequest) GetRoomIds() []AiAgentsUpdateQuotaRequestRoomIdsInner {
	if o == nil {
		var ret []AiAgentsUpdateQuotaRequestRoomIdsInner
		return ret
	}

	return o.RoomIds
}

// GetRoomIdsOk returns a tuple with the RoomIds field value
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateQuotaRequest) GetRoomIdsOk() ([]AiAgentsUpdateQuotaRequestRoomIdsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomIds, true
}

// SetRoomIds sets field value
func (o *AiAgentsUpdateQuotaRequest) SetRoomIds(v []AiAgentsUpdateQuotaRequestRoomIdsInner) {
	o.RoomIds = v
}

// GetQuota returns the Quota field value
func (o *AiAgentsUpdateQuotaRequest) GetQuota() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateQuotaRequest) GetQuotaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quota, true
}

// SetQuota sets field value
func (o *AiAgentsUpdateQuotaRequest) SetQuota(v float32) {
	o.Quota = v
}

func (o AiAgentsUpdateQuotaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAgentsUpdateQuotaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roomIds"] = o.RoomIds
	toSerialize["quota"] = o.Quota
	return toSerialize, nil
}

func (o *AiAgentsUpdateQuotaRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roomIds",
		"quota",
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

	varAiAgentsUpdateQuotaRequest := _AiAgentsUpdateQuotaRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAgentsUpdateQuotaRequest)

	if err != nil {
		return err
	}

	*o = AiAgentsUpdateQuotaRequest(varAiAgentsUpdateQuotaRequest)

	return err
}

type NullableAiAgentsUpdateQuotaRequest struct {
	value *AiAgentsUpdateQuotaRequest
	isSet bool
}

func (v NullableAiAgentsUpdateQuotaRequest) Get() *AiAgentsUpdateQuotaRequest {
	return v.value
}

func (v *NullableAiAgentsUpdateQuotaRequest) Set(val *AiAgentsUpdateQuotaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAgentsUpdateQuotaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAgentsUpdateQuotaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAgentsUpdateQuotaRequest(val *AiAgentsUpdateQuotaRequest) *NullableAiAgentsUpdateQuotaRequest {
	return &NullableAiAgentsUpdateQuotaRequest{value: val, isSet: true}
}

func (v NullableAiAgentsUpdateQuotaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAgentsUpdateQuotaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

