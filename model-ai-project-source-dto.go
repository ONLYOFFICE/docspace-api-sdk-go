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

// checks if the AiProjectSourceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProjectSourceDto{}

// AiProjectSourceDto A room linked to an AI project as its source.
type AiProjectSourceDto struct {
	// ID of the room.
	Id int32 `json:"id"`
	// Room title.
	Title NullableString `json:"title"`
	// Room logo.
	Logo *LogoDto `json:"logo,omitempty"`
}

type _AiProjectSourceDto AiProjectSourceDto

// NewAiProjectSourceDto instantiates a new AiProjectSourceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProjectSourceDto(id int32, title NullableString) *AiProjectSourceDto {
	this := AiProjectSourceDto{}
	this.Id = id
	this.Title = title
	return &this
}

// NewAiProjectSourceDtoWithDefaults instantiates a new AiProjectSourceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProjectSourceDtoWithDefaults() *AiProjectSourceDto {
	this := AiProjectSourceDto{}
	return &this
}

// GetId returns the Id field value
func (o *AiProjectSourceDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiProjectSourceDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiProjectSourceDto) SetId(v int32) {
	o.Id = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiProjectSourceDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiProjectSourceDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *AiProjectSourceDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *AiProjectSourceDto) GetLogo() LogoDto {
	if o == nil || IsNil(o.Logo) {
		var ret LogoDto
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProjectSourceDto) GetLogoOk() (*LogoDto, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *AiProjectSourceDto) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given LogoDto and assigns it to the Logo field.
func (o *AiProjectSourceDto) SetLogo(v LogoDto) {
	o.Logo = &v
}

func (o AiProjectSourceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProjectSourceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title.Get()
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	return toSerialize, nil
}

func (o *AiProjectSourceDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAiProjectSourceDto := _AiProjectSourceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProjectSourceDto)

	if err != nil {
		return err
	}

	*o = AiProjectSourceDto(varAiProjectSourceDto)

	return err
}

type NullableAiProjectSourceDto struct {
	value *AiProjectSourceDto
	isSet bool
}

func (v NullableAiProjectSourceDto) Get() *AiProjectSourceDto {
	return v.value
}

func (v *NullableAiProjectSourceDto) Set(val *AiProjectSourceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProjectSourceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProjectSourceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProjectSourceDto(val *AiProjectSourceDto) *NullableAiProjectSourceDto {
	return &NullableAiProjectSourceDto{value: val, isSet: true}
}

func (v NullableAiProjectSourceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProjectSourceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

