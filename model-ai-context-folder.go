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

// checks if the AiContextFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiContextFolder{}

// AiContextFolder The rooms of one account the caller can connect as context, grouped under the account.
type AiContextFolder struct {
	// The account the rooms belong to; `docspace` for this portal.
	Cloud string `json:"cloud"`
	// The rooms of that account the caller can connect.
	Rooms []AiContextRoom `json:"rooms"`
}

type _AiContextFolder AiContextFolder

// NewAiContextFolder instantiates a new AiContextFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiContextFolder(cloud string, rooms []AiContextRoom) *AiContextFolder {
	this := AiContextFolder{}
	this.Cloud = cloud
	this.Rooms = rooms
	return &this
}

// NewAiContextFolderWithDefaults instantiates a new AiContextFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiContextFolderWithDefaults() *AiContextFolder {
	this := AiContextFolder{}
	return &this
}

// GetCloud returns the Cloud field value
func (o *AiContextFolder) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *AiContextFolder) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *AiContextFolder) SetCloud(v string) {
	o.Cloud = v
}

// GetRooms returns the Rooms field value
func (o *AiContextFolder) GetRooms() []AiContextRoom {
	if o == nil {
		var ret []AiContextRoom
		return ret
	}

	return o.Rooms
}

// GetRoomsOk returns a tuple with the Rooms field value
// and a boolean to check if the value has been set.
func (o *AiContextFolder) GetRoomsOk() ([]AiContextRoom, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rooms, true
}

// SetRooms sets field value
func (o *AiContextFolder) SetRooms(v []AiContextRoom) {
	o.Rooms = v
}

func (o AiContextFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiContextFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud"] = o.Cloud
	toSerialize["rooms"] = o.Rooms
	return toSerialize, nil
}

func (o *AiContextFolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud",
		"rooms",
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

	varAiContextFolder := _AiContextFolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiContextFolder)

	if err != nil {
		return err
	}

	*o = AiContextFolder(varAiContextFolder)

	return err
}

type NullableAiContextFolder struct {
	value *AiContextFolder
	isSet bool
}

func (v NullableAiContextFolder) Get() *AiContextFolder {
	return v.value
}

func (v *NullableAiContextFolder) Set(val *AiContextFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableAiContextFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableAiContextFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiContextFolder(val *AiContextFolder) *NullableAiContextFolder {
	return &NullableAiContextFolder{value: val, isSet: true}
}

func (v NullableAiContextFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiContextFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

