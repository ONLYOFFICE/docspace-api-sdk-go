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

// checks if the MigratingApiUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigratingApiUser{}

// MigratingApiUser The migrating user parameters.
type MigratingApiUser struct {
	// Specifies whether the API entity should be imported.
	ShouldImport *bool `json:"shouldImport,omitempty"`
	// The user key.
	Key NullableString `json:"key,omitempty"`
	// The user email.
	Email NullableString `json:"email,omitempty"`
	// The user display name.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The user first name.
	FirstName NullableString `json:"firstName,omitempty"`
	// The user last name.
	LastName NullableString `json:"lastName,omitempty"`
	// The user type.
	UserType *EmployeeType `json:"userType,omitempty"`
	// The user's migrating files.
	MigratingFiles *MigratingApiFiles `json:"migratingFiles,omitempty"`
}

// NewMigratingApiUser instantiates a new MigratingApiUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigratingApiUser() *MigratingApiUser {
	this := MigratingApiUser{}
	return &this
}

// NewMigratingApiUserWithDefaults instantiates a new MigratingApiUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigratingApiUserWithDefaults() *MigratingApiUser {
	this := MigratingApiUser{}
	return &this
}

// GetShouldImport returns the ShouldImport field value if set, zero value otherwise.
func (o *MigratingApiUser) GetShouldImport() bool {
	if o == nil || IsNil(o.ShouldImport) {
		var ret bool
		return ret
	}
	return *o.ShouldImport
}

// GetShouldImportOk returns a tuple with the ShouldImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigratingApiUser) GetShouldImportOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldImport) {
		return nil, false
	}
	return o.ShouldImport, true
}

// HasShouldImport returns a boolean if a field has been set.
func (o *MigratingApiUser) IsShouldImportSet() bool {
	if o != nil && !IsNil(o.ShouldImport) {
		return true
	}

	return false
}

// SetShouldImport gets a reference to the given bool and assigns it to the ShouldImport field.
func (o *MigratingApiUser) SetShouldImport(v bool) {
	o.ShouldImport = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigratingApiUser) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigratingApiUser) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *MigratingApiUser) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *MigratingApiUser) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *MigratingApiUser) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *MigratingApiUser) UnsetKey() {
	o.Key.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigratingApiUser) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigratingApiUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *MigratingApiUser) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *MigratingApiUser) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *MigratingApiUser) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *MigratingApiUser) UnsetEmail() {
	o.Email.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigratingApiUser) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigratingApiUser) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MigratingApiUser) IsDisplayNameSet() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MigratingApiUser) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MigratingApiUser) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MigratingApiUser) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigratingApiUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigratingApiUser) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *MigratingApiUser) IsFirstNameSet() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *MigratingApiUser) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *MigratingApiUser) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *MigratingApiUser) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigratingApiUser) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigratingApiUser) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *MigratingApiUser) IsLastNameSet() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *MigratingApiUser) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *MigratingApiUser) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *MigratingApiUser) UnsetLastName() {
	o.LastName.Unset()
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *MigratingApiUser) GetUserType() EmployeeType {
	if o == nil || IsNil(o.UserType) {
		var ret EmployeeType
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigratingApiUser) GetUserTypeOk() (*EmployeeType, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *MigratingApiUser) IsUserTypeSet() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given EmployeeType and assigns it to the UserType field.
func (o *MigratingApiUser) SetUserType(v EmployeeType) {
	o.UserType = &v
}

// GetMigratingFiles returns the MigratingFiles field value if set, zero value otherwise.
func (o *MigratingApiUser) GetMigratingFiles() MigratingApiFiles {
	if o == nil || IsNil(o.MigratingFiles) {
		var ret MigratingApiFiles
		return ret
	}
	return *o.MigratingFiles
}

// GetMigratingFilesOk returns a tuple with the MigratingFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigratingApiUser) GetMigratingFilesOk() (*MigratingApiFiles, bool) {
	if o == nil || IsNil(o.MigratingFiles) {
		return nil, false
	}
	return o.MigratingFiles, true
}

// HasMigratingFiles returns a boolean if a field has been set.
func (o *MigratingApiUser) IsMigratingFilesSet() bool {
	if o != nil && !IsNil(o.MigratingFiles) {
		return true
	}

	return false
}

// SetMigratingFiles gets a reference to the given MigratingApiFiles and assigns it to the MigratingFiles field.
func (o *MigratingApiUser) SetMigratingFiles(v MigratingApiFiles) {
	o.MigratingFiles = &v
}

func (o MigratingApiUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigratingApiUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShouldImport) {
		toSerialize["shouldImport"] = o.ShouldImport
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if !IsNil(o.UserType) {
		toSerialize["userType"] = o.UserType
	}
	if !IsNil(o.MigratingFiles) {
		toSerialize["migratingFiles"] = o.MigratingFiles
	}
	return toSerialize, nil
}

type NullableMigratingApiUser struct {
	value *MigratingApiUser
	isSet bool
}

func (v NullableMigratingApiUser) Get() *MigratingApiUser {
	return v.value
}

func (v *NullableMigratingApiUser) Set(val *MigratingApiUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMigratingApiUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMigratingApiUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigratingApiUser(val *MigratingApiUser) *NullableMigratingApiUser {
	return &NullableMigratingApiUser{value: val, isSet: true}
}

func (v NullableMigratingApiUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigratingApiUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

