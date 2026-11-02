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

// checks if the AiProjectModelProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProjectModelProfile{}

// AiProjectModelProfile struct for AiProjectModelProfile
type AiProjectModelProfile struct {
	// The profile id.
	Id string `json:"id"`
	// The profile name; empty when the profile is gone.
	Name string `json:"name"`
	// The provider type of the profile; `null` when the profile is gone.
	Provider NullableString `json:"provider"`
	// `false` when the profile is gone or cannot chat: the project's rounds are refused.
	Available bool `json:"available"`
}

type _AiProjectModelProfile AiProjectModelProfile

// NewAiProjectModelProfile instantiates a new AiProjectModelProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProjectModelProfile(id string, name string, provider NullableString, available bool) *AiProjectModelProfile {
	this := AiProjectModelProfile{}
	this.Id = id
	this.Name = name
	this.Provider = provider
	this.Available = available
	return &this
}

// NewAiProjectModelProfileWithDefaults instantiates a new AiProjectModelProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProjectModelProfileWithDefaults() *AiProjectModelProfile {
	this := AiProjectModelProfile{}
	return &this
}

// GetId returns the Id field value
func (o *AiProjectModelProfile) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiProjectModelProfile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiProjectModelProfile) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiProjectModelProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiProjectModelProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiProjectModelProfile) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiProjectModelProfile) GetProvider() string {
	if o == nil || o.Provider.Get() == nil {
		var ret string
		return ret
	}

	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiProjectModelProfile) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// SetProvider sets field value
func (o *AiProjectModelProfile) SetProvider(v string) {
	o.Provider.Set(&v)
}

// GetAvailable returns the Available field value
func (o *AiProjectModelProfile) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *AiProjectModelProfile) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *AiProjectModelProfile) SetAvailable(v bool) {
	o.Available = v
}

func (o AiProjectModelProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProjectModelProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["provider"] = o.Provider.Get()
	toSerialize["available"] = o.Available
	return toSerialize, nil
}

func (o *AiProjectModelProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"provider",
		"available",
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

	varAiProjectModelProfile := _AiProjectModelProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProjectModelProfile)

	if err != nil {
		return err
	}

	*o = AiProjectModelProfile(varAiProjectModelProfile)

	return err
}

type NullableAiProjectModelProfile struct {
	value *AiProjectModelProfile
	isSet bool
}

func (v NullableAiProjectModelProfile) Get() *AiProjectModelProfile {
	return v.value
}

func (v *NullableAiProjectModelProfile) Set(val *AiProjectModelProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProjectModelProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProjectModelProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProjectModelProfile(val *AiProjectModelProfile) *NullableAiProjectModelProfile {
	return &NullableAiProjectModelProfile{value: val, isSet: true}
}

func (v NullableAiProjectModelProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProjectModelProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

