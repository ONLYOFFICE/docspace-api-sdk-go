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
	"time"
	"bytes"
	"fmt"
)

// checks if the ApiKeyResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyResponseDto{}

// ApiKeyResponseDto The response data for the API key operations.
type ApiKeyResponseDto struct {
	// The API key unique identifier.
	Id string `json:"id"`
	// The API key name.
	Name NullableString `json:"name"`
	// The full API key value (only returned when creating a new key).
	Key NullableString `json:"key"`
	// The API key postfix (used for identification).
	KeyPostfix NullableString `json:"keyPostfix,omitempty"`
	// The list of permissions granted to the API key.
	Permissions []string `json:"permissions"`
	// The date and time when the API key was last used.
	LastUsed NullableTime `json:"lastUsed,omitempty"`
	// The date and time when the API key was created.
	CreateOn NullableTime `json:"createOn,omitempty"`
	// The identifier of the user who created the API key.
	CreateBy *EmployeeDto `json:"createBy,omitempty"`
	// The date and time when the API key expires.
	ExpiresAt NullableTime `json:"expiresAt,omitempty"`
	// Indicates whether the API key is active or not.
	IsActive bool `json:"isActive"`
}

type _ApiKeyResponseDto ApiKeyResponseDto

// NewApiKeyResponseDto instantiates a new ApiKeyResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyResponseDto(id string, name NullableString, key NullableString, permissions []string, isActive bool) *ApiKeyResponseDto {
	this := ApiKeyResponseDto{}
	this.Id = id
	this.Name = name
	this.Key = key
	this.Permissions = permissions
	this.IsActive = isActive
	return &this
}

// NewApiKeyResponseDtoWithDefaults instantiates a new ApiKeyResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyResponseDtoWithDefaults() *ApiKeyResponseDto {
	this := ApiKeyResponseDto{}
	return &this
}

// GetId returns the Id field value
func (o *ApiKeyResponseDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiKeyResponseDto) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiKeyResponseDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyResponseDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *ApiKeyResponseDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiKeyResponseDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyResponseDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *ApiKeyResponseDto) SetKey(v string) {
	o.Key.Set(&v)
}

// GetKeyPostfix returns the KeyPostfix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyResponseDto) GetKeyPostfix() string {
	if o == nil || IsNil(o.KeyPostfix.Get()) {
		var ret string
		return ret
	}
	return *o.KeyPostfix.Get()
}

// GetKeyPostfixOk returns a tuple with the KeyPostfix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyResponseDto) GetKeyPostfixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyPostfix.Get(), o.KeyPostfix.IsSet()
}

// HasKeyPostfix returns a boolean if a field has been set.
func (o *ApiKeyResponseDto) IsKeyPostfixSet() bool {
	if o != nil && o.KeyPostfix.IsSet() {
		return true
	}

	return false
}

// SetKeyPostfix gets a reference to the given NullableString and assigns it to the KeyPostfix field.
func (o *ApiKeyResponseDto) SetKeyPostfix(v string) {
	o.KeyPostfix.Set(&v)
}
// SetKeyPostfixNil sets the value for KeyPostfix to be an explicit nil
func (o *ApiKeyResponseDto) SetKeyPostfixNil() {
	o.KeyPostfix.Set(nil)
}

// UnsetKeyPostfix ensures that no value is present for KeyPostfix, not even an explicit nil
func (o *ApiKeyResponseDto) UnsetKeyPostfix() {
	o.KeyPostfix.Unset()
}

// GetPermissions returns the Permissions field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ApiKeyResponseDto) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyResponseDto) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *ApiKeyResponseDto) SetPermissions(v []string) {
	o.Permissions = v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyResponseDto) GetLastUsed() time.Time {
	if o == nil || IsNil(o.LastUsed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUsed.Get()
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyResponseDto) GetLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsed.Get(), o.LastUsed.IsSet()
}

// HasLastUsed returns a boolean if a field has been set.
func (o *ApiKeyResponseDto) IsLastUsedSet() bool {
	if o != nil && o.LastUsed.IsSet() {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given NullableTime and assigns it to the LastUsed field.
func (o *ApiKeyResponseDto) SetLastUsed(v time.Time) {
	o.LastUsed.Set(&v)
}
// SetLastUsedNil sets the value for LastUsed to be an explicit nil
func (o *ApiKeyResponseDto) SetLastUsedNil() {
	o.LastUsed.Set(nil)
}

// UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
func (o *ApiKeyResponseDto) UnsetLastUsed() {
	o.LastUsed.Unset()
}

// GetCreateOn returns the CreateOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyResponseDto) GetCreateOn() time.Time {
	if o == nil || IsNil(o.CreateOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreateOn.Get()
}

// GetCreateOnOk returns a tuple with the CreateOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyResponseDto) GetCreateOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateOn.Get(), o.CreateOn.IsSet()
}

// HasCreateOn returns a boolean if a field has been set.
func (o *ApiKeyResponseDto) IsCreateOnSet() bool {
	if o != nil && o.CreateOn.IsSet() {
		return true
	}

	return false
}

// SetCreateOn gets a reference to the given NullableTime and assigns it to the CreateOn field.
func (o *ApiKeyResponseDto) SetCreateOn(v time.Time) {
	o.CreateOn.Set(&v)
}
// SetCreateOnNil sets the value for CreateOn to be an explicit nil
func (o *ApiKeyResponseDto) SetCreateOnNil() {
	o.CreateOn.Set(nil)
}

// UnsetCreateOn ensures that no value is present for CreateOn, not even an explicit nil
func (o *ApiKeyResponseDto) UnsetCreateOn() {
	o.CreateOn.Unset()
}

// GetCreateBy returns the CreateBy field value if set, zero value otherwise.
func (o *ApiKeyResponseDto) GetCreateBy() EmployeeDto {
	if o == nil || IsNil(o.CreateBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreateBy
}

// GetCreateByOk returns a tuple with the CreateBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseDto) GetCreateByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreateBy) {
		return nil, false
	}
	return o.CreateBy, true
}

// HasCreateBy returns a boolean if a field has been set.
func (o *ApiKeyResponseDto) IsCreateBySet() bool {
	if o != nil && !IsNil(o.CreateBy) {
		return true
	}

	return false
}

// SetCreateBy gets a reference to the given EmployeeDto and assigns it to the CreateBy field.
func (o *ApiKeyResponseDto) SetCreateBy(v EmployeeDto) {
	o.CreateBy = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyResponseDto) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyResponseDto) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ApiKeyResponseDto) IsExpiresAtSet() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *ApiKeyResponseDto) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *ApiKeyResponseDto) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *ApiKeyResponseDto) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetIsActive returns the IsActive field value
func (o *ApiKeyResponseDto) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseDto) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *ApiKeyResponseDto) SetIsActive(v bool) {
	o.IsActive = v
}

func (o ApiKeyResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["key"] = o.Key.Get()
	if o.KeyPostfix.IsSet() {
		toSerialize["keyPostfix"] = o.KeyPostfix.Get()
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.LastUsed.IsSet() {
		toSerialize["lastUsed"] = o.LastUsed.Get()
	}
	if o.CreateOn.IsSet() {
		toSerialize["createOn"] = o.CreateOn.Get()
	}
	if !IsNil(o.CreateBy) {
		toSerialize["createBy"] = o.CreateBy
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expiresAt"] = o.ExpiresAt.Get()
	}
	toSerialize["isActive"] = o.IsActive
	return toSerialize, nil
}

func (o *ApiKeyResponseDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"key",
		"permissions",
		"isActive",
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

	varApiKeyResponseDto := _ApiKeyResponseDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varApiKeyResponseDto)

	if err != nil {
		return err
	}

	*o = ApiKeyResponseDto(varApiKeyResponseDto)

	return err
}

type NullableApiKeyResponseDto struct {
	value *ApiKeyResponseDto
	isSet bool
}

func (v NullableApiKeyResponseDto) Get() *ApiKeyResponseDto {
	return v.value
}

func (v *NullableApiKeyResponseDto) Set(val *ApiKeyResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyResponseDto(val *ApiKeyResponseDto) *NullableApiKeyResponseDto {
	return &NullableApiKeyResponseDto{value: val, isSet: true}
}

func (v NullableApiKeyResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

