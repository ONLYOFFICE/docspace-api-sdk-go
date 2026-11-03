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

// checks if the AiFolderContextRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFolderContextRef{}

// AiFolderContextRef A folder connected as the round's context - any folder of the portal, not only a room - named by its account and ID.
type AiFolderContextRef struct {
	// The account the folder belongs to; `docspace` for a folder of this portal.
	Cloud string `json:"cloud"`
	// The folder ID.
	Id string `json:"id"`
}

type _AiFolderContextRef AiFolderContextRef

// NewAiFolderContextRef instantiates a new AiFolderContextRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFolderContextRef(cloud string, id string) *AiFolderContextRef {
	this := AiFolderContextRef{}
	this.Cloud = cloud
	this.Id = id
	return &this
}

// NewAiFolderContextRefWithDefaults instantiates a new AiFolderContextRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFolderContextRefWithDefaults() *AiFolderContextRef {
	this := AiFolderContextRef{}
	return &this
}

// GetCloud returns the Cloud field value
func (o *AiFolderContextRef) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *AiFolderContextRef) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *AiFolderContextRef) SetCloud(v string) {
	o.Cloud = v
}

// GetId returns the Id field value
func (o *AiFolderContextRef) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiFolderContextRef) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiFolderContextRef) SetId(v string) {
	o.Id = v
}

func (o AiFolderContextRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFolderContextRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud"] = o.Cloud
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AiFolderContextRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud",
		"id",
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

	varAiFolderContextRef := _AiFolderContextRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiFolderContextRef)

	if err != nil {
		return err
	}

	*o = AiFolderContextRef(varAiFolderContextRef)

	return err
}

type NullableAiFolderContextRef struct {
	value *AiFolderContextRef
	isSet bool
}

func (v NullableAiFolderContextRef) Get() *AiFolderContextRef {
	return v.value
}

func (v *NullableAiFolderContextRef) Set(val *AiFolderContextRef) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFolderContextRef) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFolderContextRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFolderContextRef(val *AiFolderContextRef) *NullableAiFolderContextRef {
	return &NullableAiFolderContextRef{value: val, isSet: true}
}

func (v NullableAiFolderContextRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFolderContextRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

