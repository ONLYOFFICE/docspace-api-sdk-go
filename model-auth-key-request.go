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

// checks if the AuthKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthKeyRequest{}

// AuthKeyRequest One key of a provider and the value to store for it.
type AuthKeyRequest struct {
	// The key name, as `GET api/2.0/settings/authservice` lists it in `props`.
	Name NullableString `json:"name"`
	// The value to store. An empty string clears the key.
	Value NullableString `json:"value"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	Title NullableString `json:"title,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	Type NullableString `json:"type,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	Options []string `json:"options,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	DependsOn NullableString `json:"dependsOn,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	DependsOnValue NullableString `json:"dependsOnValue,omitempty"`
}

type _AuthKeyRequest AuthKeyRequest

// NewAuthKeyRequest instantiates a new AuthKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthKeyRequest(name NullableString, value NullableString) *AuthKeyRequest {
	this := AuthKeyRequest{}
	this.Name = name
	this.Value = value
	return &this
}

// NewAuthKeyRequestWithDefaults instantiates a new AuthKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthKeyRequestWithDefaults() *AuthKeyRequest {
	this := AuthKeyRequest{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuthKeyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *AuthKeyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuthKeyRequest) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *AuthKeyRequest) SetValue(v string) {
	o.Value.Set(&v)
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AuthKeyRequest) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AuthKeyRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AuthKeyRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AuthKeyRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyRequest) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *AuthKeyRequest) IsTypeSet() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *AuthKeyRequest) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *AuthKeyRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *AuthKeyRequest) UnsetType() {
	o.Type.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyRequest) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyRequest) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AuthKeyRequest) IsOptionsSet() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *AuthKeyRequest) SetOptions(v []string) {
	o.Options = v
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyRequest) GetDependsOn() string {
	if o == nil || IsNil(o.DependsOn.Get()) {
		var ret string
		return ret
	}
	return *o.DependsOn.Get()
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyRequest) GetDependsOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependsOn.Get(), o.DependsOn.IsSet()
}

// HasDependsOn returns a boolean if a field has been set.
func (o *AuthKeyRequest) IsDependsOnSet() bool {
	if o != nil && o.DependsOn.IsSet() {
		return true
	}

	return false
}

// SetDependsOn gets a reference to the given NullableString and assigns it to the DependsOn field.
func (o *AuthKeyRequest) SetDependsOn(v string) {
	o.DependsOn.Set(&v)
}
// SetDependsOnNil sets the value for DependsOn to be an explicit nil
func (o *AuthKeyRequest) SetDependsOnNil() {
	o.DependsOn.Set(nil)
}

// UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
func (o *AuthKeyRequest) UnsetDependsOn() {
	o.DependsOn.Unset()
}

// GetDependsOnValue returns the DependsOnValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthKeyRequest) GetDependsOnValue() string {
	if o == nil || IsNil(o.DependsOnValue.Get()) {
		var ret string
		return ret
	}
	return *o.DependsOnValue.Get()
}

// GetDependsOnValueOk returns a tuple with the DependsOnValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthKeyRequest) GetDependsOnValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependsOnValue.Get(), o.DependsOnValue.IsSet()
}

// HasDependsOnValue returns a boolean if a field has been set.
func (o *AuthKeyRequest) IsDependsOnValueSet() bool {
	if o != nil && o.DependsOnValue.IsSet() {
		return true
	}

	return false
}

// SetDependsOnValue gets a reference to the given NullableString and assigns it to the DependsOnValue field.
func (o *AuthKeyRequest) SetDependsOnValue(v string) {
	o.DependsOnValue.Set(&v)
}
// SetDependsOnValueNil sets the value for DependsOnValue to be an explicit nil
func (o *AuthKeyRequest) SetDependsOnValueNil() {
	o.DependsOnValue.Set(nil)
}

// UnsetDependsOnValue ensures that no value is present for DependsOnValue, not even an explicit nil
func (o *AuthKeyRequest) UnsetDependsOnValue() {
	o.DependsOnValue.Unset()
}

func (o AuthKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthKeyRequest) ToMap() (map[string]interface{}, error) {
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

func (o *AuthKeyRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAuthKeyRequest := _AuthKeyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAuthKeyRequest)

	if err != nil {
		return err
	}

	*o = AuthKeyRequest(varAuthKeyRequest)

	return err
}

type NullableAuthKeyRequest struct {
	value *AuthKeyRequest
	isSet bool
}

func (v NullableAuthKeyRequest) Get() *AuthKeyRequest {
	return v.value
}

func (v *NullableAuthKeyRequest) Set(val *AuthKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthKeyRequest(val *AuthKeyRequest) *NullableAuthKeyRequest {
	return &NullableAuthKeyRequest{value: val, isSet: true}
}

func (v NullableAuthKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

