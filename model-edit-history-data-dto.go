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

// checks if the EditHistoryDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHistoryDataDto{}

// EditHistoryDataDto The file editing history data.
type EditHistoryDataDto struct {
	// The URL address of the file with the document changes data.
	ChangesUrl NullableString `json:"changesUrl,omitempty"`
	// The document identifier used to unambiguously identify the document file.
	Key NullableString `json:"key"`
	Previous *EditHistoryUrl `json:"previous,omitempty"`
	// The encrypted signature added to the parameter in the form of a token.
	Token NullableString `json:"token,omitempty"`
	// The URL address of the current document version.
	Url NullableString `json:"url"`
	// The document version number.
	Version int32 `json:"version"`
	// The document extension.
	FileType NullableString `json:"fileType"`
}

type _EditHistoryDataDto EditHistoryDataDto

// NewEditHistoryDataDto instantiates a new EditHistoryDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHistoryDataDto(key NullableString, url NullableString, version int32, fileType NullableString) *EditHistoryDataDto {
	this := EditHistoryDataDto{}
	this.Key = key
	this.Url = url
	this.Version = version
	this.FileType = fileType
	return &this
}

// NewEditHistoryDataDtoWithDefaults instantiates a new EditHistoryDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHistoryDataDtoWithDefaults() *EditHistoryDataDto {
	this := EditHistoryDataDto{}
	return &this
}

// GetChangesUrl returns the ChangesUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryDataDto) GetChangesUrl() string {
	if o == nil || IsNil(o.ChangesUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ChangesUrl.Get()
}

// GetChangesUrlOk returns a tuple with the ChangesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryDataDto) GetChangesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangesUrl.Get(), o.ChangesUrl.IsSet()
}

// HasChangesUrl returns a boolean if a field has been set.
func (o *EditHistoryDataDto) IsChangesUrlSet() bool {
	if o != nil && o.ChangesUrl.IsSet() {
		return true
	}

	return false
}

// SetChangesUrl gets a reference to the given NullableString and assigns it to the ChangesUrl field.
func (o *EditHistoryDataDto) SetChangesUrl(v string) {
	o.ChangesUrl.Set(&v)
}
// SetChangesUrlNil sets the value for ChangesUrl to be an explicit nil
func (o *EditHistoryDataDto) SetChangesUrlNil() {
	o.ChangesUrl.Set(nil)
}

// UnsetChangesUrl ensures that no value is present for ChangesUrl, not even an explicit nil
func (o *EditHistoryDataDto) UnsetChangesUrl() {
	o.ChangesUrl.Unset()
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EditHistoryDataDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryDataDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *EditHistoryDataDto) SetKey(v string) {
	o.Key.Set(&v)
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *EditHistoryDataDto) GetPrevious() EditHistoryUrl {
	if o == nil || IsNil(o.Previous) {
		var ret EditHistoryUrl
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHistoryDataDto) GetPreviousOk() (*EditHistoryUrl, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *EditHistoryDataDto) IsPreviousSet() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given EditHistoryUrl and assigns it to the Previous field.
func (o *EditHistoryDataDto) SetPrevious(v EditHistoryUrl) {
	o.Previous = &v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryDataDto) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryDataDto) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *EditHistoryDataDto) IsTokenSet() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *EditHistoryDataDto) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *EditHistoryDataDto) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *EditHistoryDataDto) UnsetToken() {
	o.Token.Unset()
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EditHistoryDataDto) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryDataDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *EditHistoryDataDto) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetVersion returns the Version field value
func (o *EditHistoryDataDto) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EditHistoryDataDto) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EditHistoryDataDto) SetVersion(v int32) {
	o.Version = v
}

// GetFileType returns the FileType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EditHistoryDataDto) GetFileType() string {
	if o == nil || o.FileType.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileType.Get()
}

// GetFileTypeOk returns a tuple with the FileType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryDataDto) GetFileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileType.Get(), o.FileType.IsSet()
}

// SetFileType sets field value
func (o *EditHistoryDataDto) SetFileType(v string) {
	o.FileType.Set(&v)
}

func (o EditHistoryDataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHistoryDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangesUrl.IsSet() {
		toSerialize["changesUrl"] = o.ChangesUrl.Get()
	}
	toSerialize["key"] = o.Key.Get()
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	toSerialize["url"] = o.Url.Get()
	toSerialize["version"] = o.Version
	toSerialize["fileType"] = o.FileType.Get()
	return toSerialize, nil
}

func (o *EditHistoryDataDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"url",
		"version",
		"fileType",
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

	varEditHistoryDataDto := _EditHistoryDataDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varEditHistoryDataDto)

	if err != nil {
		return err
	}

	*o = EditHistoryDataDto(varEditHistoryDataDto)

	return err
}

type NullableEditHistoryDataDto struct {
	value *EditHistoryDataDto
	isSet bool
}

func (v NullableEditHistoryDataDto) Get() *EditHistoryDataDto {
	return v.value
}

func (v *NullableEditHistoryDataDto) Set(val *EditHistoryDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHistoryDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHistoryDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHistoryDataDto(val *EditHistoryDataDto) *NullableEditHistoryDataDto {
	return &NullableEditHistoryDataDto{value: val, isSet: true}
}

func (v NullableEditHistoryDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHistoryDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

