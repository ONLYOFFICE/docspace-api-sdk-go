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
)

// checks if the UpdateApiKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApiKeyRequest{}

// UpdateApiKeyRequest The request parameters for updating an existing API key.
type UpdateApiKeyRequest struct {
	// The new name for the API key.
	Name NullableString `json:"name,omitempty"`
	// The new list of permissions for the API key.
	Permissions []string `json:"permissions,omitempty"`
	// Indicates whether the API key should be active or not.
	IsActive NullableBool `json:"isActive,omitempty"`
}

// NewUpdateApiKeyRequest instantiates a new UpdateApiKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApiKeyRequest() *UpdateApiKeyRequest {
	this := UpdateApiKeyRequest{}
	return &this
}

// NewUpdateApiKeyRequestWithDefaults instantiates a new UpdateApiKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApiKeyRequestWithDefaults() *UpdateApiKeyRequest {
	this := UpdateApiKeyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateApiKeyRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateApiKeyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateApiKeyRequest) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateApiKeyRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateApiKeyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateApiKeyRequest) UnsetName() {
	o.Name.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateApiKeyRequest) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateApiKeyRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UpdateApiKeyRequest) IsPermissionsSet() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *UpdateApiKeyRequest) SetPermissions(v []string) {
	o.Permissions = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateApiKeyRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateApiKeyRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *UpdateApiKeyRequest) IsIsActiveSet() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *UpdateApiKeyRequest) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *UpdateApiKeyRequest) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *UpdateApiKeyRequest) UnsetIsActive() {
	o.IsActive.Unset()
}

func (o UpdateApiKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApiKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	return toSerialize, nil
}

type NullableUpdateApiKeyRequest struct {
	value *UpdateApiKeyRequest
	isSet bool
}

func (v NullableUpdateApiKeyRequest) Get() *UpdateApiKeyRequest {
	return v.value
}

func (v *NullableUpdateApiKeyRequest) Set(val *UpdateApiKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApiKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApiKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApiKeyRequest(val *UpdateApiKeyRequest) *NullableUpdateApiKeyRequest {
	return &NullableUpdateApiKeyRequest{value: val, isSet: true}
}

func (v NullableUpdateApiKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApiKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

