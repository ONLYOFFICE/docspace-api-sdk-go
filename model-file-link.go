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

// checks if the FileLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileLink{}

// FileLink The file link properties.
type FileLink struct {
	// The type of the file for the source viewed or edited document.
	Filetype NullableString `json:"filetype"`
	// The encrypted signature added to the config in the form of a token.
	Token NullableString `json:"token,omitempty"`
	// The absolute URL where the source viewed or edited document is stored.
	Url NullableString `json:"url"`
}

type _FileLink FileLink

// NewFileLink instantiates a new FileLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLink(filetype NullableString, url NullableString) *FileLink {
	this := FileLink{}
	this.Filetype = filetype
	this.Url = url
	return &this
}

// NewFileLinkWithDefaults instantiates a new FileLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLinkWithDefaults() *FileLink {
	this := FileLink{}
	return &this
}

// GetFiletype returns the Filetype field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FileLink) GetFiletype() string {
	if o == nil || o.Filetype.Get() == nil {
		var ret string
		return ret
	}

	return *o.Filetype.Get()
}

// GetFiletypeOk returns a tuple with the Filetype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLink) GetFiletypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filetype.Get(), o.Filetype.IsSet()
}

// SetFiletype sets field value
func (o *FileLink) SetFiletype(v string) {
	o.Filetype.Set(&v)
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLink) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLink) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *FileLink) IsTokenSet() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *FileLink) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *FileLink) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *FileLink) UnsetToken() {
	o.Token.Unset()
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FileLink) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLink) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *FileLink) SetUrl(v string) {
	o.Url.Set(&v)
}

func (o FileLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filetype"] = o.Filetype.Get()
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	toSerialize["url"] = o.Url.Get()
	return toSerialize, nil
}

func (o *FileLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filetype",
		"url",
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

	varFileLink := _FileLink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFileLink)

	if err != nil {
		return err
	}

	*o = FileLink(varFileLink)

	return err
}

type NullableFileLink struct {
	value *FileLink
	isSet bool
}

func (v NullableFileLink) Get() *FileLink {
	return v.value
}

func (v *NullableFileLink) Set(val *FileLink) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLink) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLink(val *FileLink) *NullableFileLink {
	return &NullableFileLink{value: val, isSet: true}
}

func (v NullableFileLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

