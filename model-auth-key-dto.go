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

// checks if the AuthKeyDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthKeyDto{}

// AuthKeyDto One key of an authorization provider or a storage, with how the settings form shows it.
type AuthKeyDto struct {
	// The authorization key name.
	Name NullableString `json:"name"`
	// The authorization key value.
	Value NullableString `json:"value"`
	// The authorization key title.
	Title NullableString `json:"title,omitempty"`
	// The field type: text, password, select, toggle.
	Type NullableString `json:"type,omitempty"`
	// The list of options for select type fields.
	Options []string `json:"options,omitempty"`
	// The name of another key this field depends on for visibility.
	DependsOn NullableString `json:"dependsOn,omitempty"`
	// The value of the `dependsOn` key that makes this field visible.
	DependsOnValue NullableString `json:"dependsOnValue,omitempty"`
}

type _AuthKeyDto AuthKeyDto

// NewAuthKeyDto instantiates a new AuthKeyDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthKeyDto(name NullableString, value NullableString) *AuthKeyDto {
	this := AuthKeyDto{}
	this.Name = name
	this.Value = value
	return &this
}

// NewAuthKeyDtoWithDefaults instantiates a new AuthKeyDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthKeyDtoWithDefaults() *AuthKeyDto {
	this := AuthKeyDto{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuthKeyDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *AuthKeyDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuthKeyDto) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *AuthKeyDto) SetValue(v string) {
	o.Value.Set(&v)
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AuthKeyDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AuthKeyDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AuthKeyDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AuthKeyDto) UnsetTitle() {
	o.Title.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *AuthKeyDto) IsTypeSet() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *AuthKeyDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *AuthKeyDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *AuthKeyDto) UnsetType() {
	o.Type.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyDto) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyDto) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AuthKeyDto) IsOptionsSet() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *AuthKeyDto) SetOptions(v []string) {
	o.Options = v
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyDto) GetDependsOn() string {
	if o == nil || IsNil(o.DependsOn.Get()) {
		var ret string
		return ret
	}
	return *o.DependsOn.Get()
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyDto) GetDependsOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependsOn.Get(), o.DependsOn.IsSet()
}

// HasDependsOn returns a boolean if a field has been set.
func (o *AuthKeyDto) IsDependsOnSet() bool {
	if o != nil && o.DependsOn.IsSet() {
		return true
	}

	return false
}

// SetDependsOn gets a reference to the given NullableString and assigns it to the DependsOn field.
func (o *AuthKeyDto) SetDependsOn(v string) {
	o.DependsOn.Set(&v)
}
// SetDependsOnNil sets the value for DependsOn to be an explicit nil
func (o *AuthKeyDto) SetDependsOnNil() {
	o.DependsOn.Set(nil)
}

// UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
func (o *AuthKeyDto) UnsetDependsOn() {
	o.DependsOn.Unset()
}

// GetDependsOnValue returns the DependsOnValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyDto) GetDependsOnValue() string {
	if o == nil || IsNil(o.DependsOnValue.Get()) {
		var ret string
		return ret
	}
	return *o.DependsOnValue.Get()
}

// GetDependsOnValueOk returns a tuple with the DependsOnValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyDto) GetDependsOnValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependsOnValue.Get(), o.DependsOnValue.IsSet()
}

// HasDependsOnValue returns a boolean if a field has been set.
func (o *AuthKeyDto) IsDependsOnValueSet() bool {
	if o != nil && o.DependsOnValue.IsSet() {
		return true
	}

	return false
}

// SetDependsOnValue gets a reference to the given NullableString and assigns it to the DependsOnValue field.
func (o *AuthKeyDto) SetDependsOnValue(v string) {
	o.DependsOnValue.Set(&v)
}
// SetDependsOnValueNil sets the value for DependsOnValue to be an explicit nil
func (o *AuthKeyDto) SetDependsOnValueNil() {
	o.DependsOnValue.Set(nil)
}

// UnsetDependsOnValue ensures that no value is present for DependsOnValue, not even an explicit nil
func (o *AuthKeyDto) UnsetDependsOnValue() {
	o.DependsOnValue.Unset()
}

func (o AuthKeyDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthKeyDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["value"] = o.Value.Get()
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.DependsOn.IsSet() {
		toSerialize["dependsOn"] = o.DependsOn.Get()
	}
	if o.DependsOnValue.IsSet() {
		toSerialize["dependsOnValue"] = o.DependsOnValue.Get()
	}
	return toSerialize, nil
}

func (o *AuthKeyDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varAuthKeyDto := _AuthKeyDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAuthKeyDto)

	if err != nil {
		return err
	}

	*o = AuthKeyDto(varAuthKeyDto)

	return err
}

type NullableAuthKeyDto struct {
	value *AuthKeyDto
	isSet bool
}

func (v NullableAuthKeyDto) Get() *AuthKeyDto {
	return v.value
}

func (v *NullableAuthKeyDto) Set(val *AuthKeyDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthKeyDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthKeyDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthKeyDto(val *AuthKeyDto) *NullableAuthKeyDto {
	return &NullableAuthKeyDto{value: val, isSet: true}
}

func (v NullableAuthKeyDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthKeyDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

