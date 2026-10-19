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

// checks if the CreateFolderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFolderRequest{}

// CreateFolderRequest The title a folder is created with or renamed to.
type CreateFolderRequest struct {
	// The title the folder is given. It is trimmed before it is stored and may not be blank or consist of spaces  alone; it need not differ from the titles of the neighbouring folders, so the same title may appear twice in  one parent.
	Title NullableString `json:"title"`
}

type _CreateFolderRequest CreateFolderRequest

// NewCreateFolderRequest instantiates a new CreateFolderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFolderRequest(title NullableString) *CreateFolderRequest {
	this := CreateFolderRequest{}
	this.Title = title
	return &this
}

// NewCreateFolderRequestWithDefaults instantiates a new CreateFolderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFolderRequestWithDefaults() *CreateFolderRequest {
	this := CreateFolderRequest{}
	return &this
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateFolderRequest) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFolderRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateFolderRequest) SetTitle(v string) {
	o.Title.Set(&v)
}

func (o CreateFolderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFolderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title.Get()
	return toSerialize, nil
}

func (o *CreateFolderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varCreateFolderRequest := _CreateFolderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateFolderRequest)

	if err != nil {
		return err
	}

	*o = CreateFolderRequest(varCreateFolderRequest)

	return err
}

type NullableCreateFolderRequest struct {
	value *CreateFolderRequest
	isSet bool
}

func (v NullableCreateFolderRequest) Get() *CreateFolderRequest {
	return v.value
}

func (v *NullableCreateFolderRequest) Set(val *CreateFolderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFolderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFolderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFolderRequest(val *CreateFolderRequest) *NullableCreateFolderRequest {
	return &NullableCreateFolderRequest{value: val, isSet: true}
}

func (v NullableCreateFolderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFolderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

