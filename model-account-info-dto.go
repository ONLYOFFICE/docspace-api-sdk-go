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

// checks if the AccountInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountInfoDto{}

// AccountInfoDto The account information parameters.
type AccountInfoDto struct {
	// The name of the identity provider, in lowercase, as every other operation of this group expects it: `google`,  `zoom`, `linkedin`, `facebook`, `twitter`, `microsoft`, `appleid`, `weixin` or `nextcloud`.
	Provider NullableString `json:"provider"`
	// The URL that starts the login with this provider. Open it as it is - it already carries the provider and the  popup or redirect mode the request asked for.
	Url NullableString `json:"url"`
	// Whether this provider is already linked to the calling profile. It is always false for an anonymous caller,  because there is no profile to compare against.
	Linked bool `json:"linked"`
}

type _AccountInfoDto AccountInfoDto

// NewAccountInfoDto instantiates a new AccountInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfoDto(provider NullableString, url NullableString, linked bool) *AccountInfoDto {
	this := AccountInfoDto{}
	this.Provider = provider
	this.Url = url
	this.Linked = linked
	return &this
}

// NewAccountInfoDtoWithDefaults instantiates a new AccountInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoDtoWithDefaults() *AccountInfoDto {
	this := AccountInfoDto{}
	return &this
}

// GetProvider returns the Provider field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountInfoDto) GetProvider() string {
	if o == nil || o.Provider.Get() == nil {
		var ret string
		return ret
	}

	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfoDto) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// SetProvider sets field value
func (o *AccountInfoDto) SetProvider(v string) {
	o.Provider.Set(&v)
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountInfoDto) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfoDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *AccountInfoDto) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetLinked returns the Linked field value
func (o *AccountInfoDto) GetLinked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Linked
}

// GetLinkedOk returns a tuple with the Linked field value
// and a boolean to check if the value has been set.
func (o *AccountInfoDto) GetLinkedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Linked, true
}

// SetLinked sets field value
func (o *AccountInfoDto) SetLinked(v bool) {
	o.Linked = v
}

func (o AccountInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider.Get()
	toSerialize["url"] = o.Url.Get()
	toSerialize["linked"] = o.Linked
	return toSerialize, nil
}

func (o *AccountInfoDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider",
		"url",
		"linked",
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

	varAccountInfoDto := _AccountInfoDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAccountInfoDto)

	if err != nil {
		return err
	}

	*o = AccountInfoDto(varAccountInfoDto)

	return err
}

type NullableAccountInfoDto struct {
	value *AccountInfoDto
	isSet bool
}

func (v NullableAccountInfoDto) Get() *AccountInfoDto {
	return v.value
}

func (v *NullableAccountInfoDto) Set(val *AccountInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfoDto(val *AccountInfoDto) *NullableAccountInfoDto {
	return &NullableAccountInfoDto{value: val, isSet: true}
}

func (v NullableAccountInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

