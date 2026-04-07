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

// checks if the MigrationApiInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrationApiInfo{}

// MigrationApiInfo struct for MigrationApiInfo
type MigrationApiInfo struct {
	MigratorName NullableString `json:"migratorName,omitempty"`
	Operation NullableString `json:"operation,omitempty"`
	FailedArchives []string `json:"failedArchives,omitempty"`
	Users []MigratingApiUser `json:"users,omitempty"`
	WithoutEmailUsers []MigratingApiUser `json:"withoutEmailUsers,omitempty"`
	ExistUsers []MigratingApiUser `json:"existUsers,omitempty"`
	Groups []MigratingApiGroup `json:"groups,omitempty"`
	ImportPersonalFiles *bool `json:"importPersonalFiles,omitempty"`
	ImportSharedFiles *bool `json:"importSharedFiles,omitempty"`
	ImportSharedFolders *bool `json:"importSharedFolders,omitempty"`
	ImportCommonFiles *bool `json:"importCommonFiles,omitempty"`
	ImportProjectFiles *bool `json:"importProjectFiles,omitempty"`
	ImportGroups *bool `json:"importGroups,omitempty"`
	SuccessedUsers *int32 `json:"successedUsers,omitempty"`
	FailedUsers *int32 `json:"failedUsers,omitempty"`
	Files []string `json:"files,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// NewMigrationApiInfo instantiates a new MigrationApiInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationApiInfo() *MigrationApiInfo {
	this := MigrationApiInfo{}
	return &this
}

// NewMigrationApiInfoWithDefaults instantiates a new MigrationApiInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationApiInfoWithDefaults() *MigrationApiInfo {
	this := MigrationApiInfo{}
	return &this
}

// GetMigratorName returns the MigratorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationApiInfo) GetMigratorName() string {
	if o == nil || IsNil(o.MigratorName.Get()) {
		var ret string
		return ret
	}
	return *o.MigratorName.Get()
}

// GetMigratorNameOk returns a tuple with the MigratorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationApiInfo) GetMigratorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MigratorName.Get(), o.MigratorName.IsSet()
}

// HasMigratorName returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsMigratorNameSet() bool {
	if o != nil && o.MigratorName.IsSet() {
		return true
	}

	return false
}

// SetMigratorName gets a reference to the given NullableString and assigns it to the MigratorName field.
func (o *MigrationApiInfo) SetMigratorName(v string) {
	o.MigratorName.Set(&v)
}
// SetMigratorNameNil sets the value for MigratorName to be an explicit nil
func (o *MigrationApiInfo) SetMigratorNameNil() {
	o.MigratorName.Set(nil)
}

// UnsetMigratorName ensures that no value is present for MigratorName, not even an explicit nil
func (o *MigrationApiInfo) UnsetMigratorName() {
	o.MigratorName.Unset()
}

// GetOperation returns the Operation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationApiInfo) GetOperation() string {
	if o == nil || IsNil(o.Operation.Get()) {
		var ret string
		return ret
	}
	return *o.Operation.Get()
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationApiInfo) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operation.Get(), o.Operation.IsSet()
}

// HasOperation returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsOperationSet() bool {
	if o != nil && o.Operation.IsSet() {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NullableString and assigns it to the Operation field.
func (o *MigrationApiInfo) SetOperation(v string) {
	o.Operation.Set(&v)
}
// SetOperationNil sets the value for Operation to be an explicit nil
func (o *MigrationApiInfo) SetOperationNil() {
	o.Operation.Set(nil)
}

// UnsetOperation ensures that no value is present for Operation, not even an explicit nil
func (o *MigrationApiInfo) UnsetOperation() {
	o.Operation.Unset()
}

// GetFailedArchives returns the FailedArchives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationApiInfo) GetFailedArchives() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FailedArchives
}

// GetFailedArchivesOk returns a tuple with the FailedArchives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationApiInfo) GetFailedArchivesOk() ([]string, bool) {
	if o == nil || IsNil(o.FailedArchives) {
		return nil, false
	}
	return o.FailedArchives, true
}

// HasFailedArchives returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsFailedArchivesSet() bool {
	if o != nil && !IsNil(o.FailedArchives) {
		return true
	}

	return false
}

// SetFailedArchives gets a reference to the given []string and assigns it to the FailedArchives field.
func (o *MigrationApiInfo) SetFailedArchives(v []string) {
	o.FailedArchives = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationApiInfo) GetUsers() []MigratingApiUser {
	if o == nil {
		var ret []MigratingApiUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationApiInfo) GetUsersOk() ([]MigratingApiUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsUsersSet() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []MigratingApiUser and assigns it to the Users field.
func (o *MigrationApiInfo) SetUsers(v []MigratingApiUser) {
	o.Users = v
}

// GetWithoutEmailUsers returns the WithoutEmailUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationApiInfo) GetWithoutEmailUsers() []MigratingApiUser {
	if o == nil {
		var ret []MigratingApiUser
		return ret
	}
	return o.WithoutEmailUsers
}

// GetWithoutEmailUsersOk returns a tuple with the WithoutEmailUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationApiInfo) GetWithoutEmailUsersOk() ([]MigratingApiUser, bool) {
	if o == nil || IsNil(o.WithoutEmailUsers) {
		return nil, false
	}
	return o.WithoutEmailUsers, true
}

// HasWithoutEmailUsers returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsWithoutEmailUsersSet() bool {
	if o != nil && !IsNil(o.WithoutEmailUsers) {
		return true
	}

	return false
}

// SetWithoutEmailUsers gets a reference to the given []MigratingApiUser and assigns it to the WithoutEmailUsers field.
func (o *MigrationApiInfo) SetWithoutEmailUsers(v []MigratingApiUser) {
	o.WithoutEmailUsers = v
}

// GetExistUsers returns the ExistUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationApiInfo) GetExistUsers() []MigratingApiUser {
	if o == nil {
		var ret []MigratingApiUser
		return ret
	}
	return o.ExistUsers
}

// GetExistUsersOk returns a tuple with the ExistUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationApiInfo) GetExistUsersOk() ([]MigratingApiUser, bool) {
	if o == nil || IsNil(o.ExistUsers) {
		return nil, false
	}
	return o.ExistUsers, true
}

// HasExistUsers returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsExistUsersSet() bool {
	if o != nil && !IsNil(o.ExistUsers) {
		return true
	}

	return false
}

// SetExistUsers gets a reference to the given []MigratingApiUser and assigns it to the ExistUsers field.
func (o *MigrationApiInfo) SetExistUsers(v []MigratingApiUser) {
	o.ExistUsers = v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationApiInfo) GetGroups() []MigratingApiGroup {
	if o == nil {
		var ret []MigratingApiGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationApiInfo) GetGroupsOk() ([]MigratingApiGroup, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsGroupsSet() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []MigratingApiGroup and assigns it to the Groups field.
func (o *MigrationApiInfo) SetGroups(v []MigratingApiGroup) {
	o.Groups = v
}

// GetImportPersonalFiles returns the ImportPersonalFiles field value if set, zero value otherwise.
func (o *MigrationApiInfo) GetImportPersonalFiles() bool {
	if o == nil || IsNil(o.ImportPersonalFiles) {
		var ret bool
		return ret
	}
	return *o.ImportPersonalFiles
}

// GetImportPersonalFilesOk returns a tuple with the ImportPersonalFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationApiInfo) GetImportPersonalFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportPersonalFiles) {
		return nil, false
	}
	return o.ImportPersonalFiles, true
}

// HasImportPersonalFiles returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsImportPersonalFilesSet() bool {
	if o != nil && !IsNil(o.ImportPersonalFiles) {
		return true
	}

	return false
}

// SetImportPersonalFiles gets a reference to the given bool and assigns it to the ImportPersonalFiles field.
func (o *MigrationApiInfo) SetImportPersonalFiles(v bool) {
	o.ImportPersonalFiles = &v
}

// GetImportSharedFiles returns the ImportSharedFiles field value if set, zero value otherwise.
func (o *MigrationApiInfo) GetImportSharedFiles() bool {
	if o == nil || IsNil(o.ImportSharedFiles) {
		var ret bool
		return ret
	}
	return *o.ImportSharedFiles
}

// GetImportSharedFilesOk returns a tuple with the ImportSharedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationApiInfo) GetImportSharedFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportSharedFiles) {
		return nil, false
	}
	return o.ImportSharedFiles, true
}

// HasImportSharedFiles returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsImportSharedFilesSet() bool {
	if o != nil && !IsNil(o.ImportSharedFiles) {
		return true
	}

	return false
}

// SetImportSharedFiles gets a reference to the given bool and assigns it to the ImportSharedFiles field.
func (o *MigrationApiInfo) SetImportSharedFiles(v bool) {
	o.ImportSharedFiles = &v
}

// GetImportSharedFolders returns the ImportSharedFolders field value if set, zero value otherwise.
func (o *MigrationApiInfo) GetImportSharedFolders() bool {
	if o == nil || IsNil(o.ImportSharedFolders) {
		var ret bool
		return ret
	}
	return *o.ImportSharedFolders
}

// GetImportSharedFoldersOk returns a tuple with the ImportSharedFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationApiInfo) GetImportSharedFoldersOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportSharedFolders) {
		return nil, false
	}
	return o.ImportSharedFolders, true
}

// HasImportSharedFolders returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsImportSharedFoldersSet() bool {
	if o != nil && !IsNil(o.ImportSharedFolders) {
		return true
	}

	return false
}

// SetImportSharedFolders gets a reference to the given bool and assigns it to the ImportSharedFolders field.
func (o *MigrationApiInfo) SetImportSharedFolders(v bool) {
	o.ImportSharedFolders = &v
}

// GetImportCommonFiles returns the ImportCommonFiles field value if set, zero value otherwise.
func (o *MigrationApiInfo) GetImportCommonFiles() bool {
	if o == nil || IsNil(o.ImportCommonFiles) {
		var ret bool
		return ret
	}
	return *o.ImportCommonFiles
}

// GetImportCommonFilesOk returns a tuple with the ImportCommonFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationApiInfo) GetImportCommonFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportCommonFiles) {
		return nil, false
	}
	return o.ImportCommonFiles, true
}

// HasImportCommonFiles returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsImportCommonFilesSet() bool {
	if o != nil && !IsNil(o.ImportCommonFiles) {
		return true
	}

	return false
}

// SetImportCommonFiles gets a reference to the given bool and assigns it to the ImportCommonFiles field.
func (o *MigrationApiInfo) SetImportCommonFiles(v bool) {
	o.ImportCommonFiles = &v
}

// GetImportProjectFiles returns the ImportProjectFiles field value if set, zero value otherwise.
func (o *MigrationApiInfo) GetImportProjectFiles() bool {
	if o == nil || IsNil(o.ImportProjectFiles) {
		var ret bool
		return ret
	}
	return *o.ImportProjectFiles
}

// GetImportProjectFilesOk returns a tuple with the ImportProjectFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationApiInfo) GetImportProjectFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportProjectFiles) {
		return nil, false
	}
	return o.ImportProjectFiles, true
}

// HasImportProjectFiles returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsImportProjectFilesSet() bool {
	if o != nil && !IsNil(o.ImportProjectFiles) {
		return true
	}

	return false
}

// SetImportProjectFiles gets a reference to the given bool and assigns it to the ImportProjectFiles field.
func (o *MigrationApiInfo) SetImportProjectFiles(v bool) {
	o.ImportProjectFiles = &v
}

// GetImportGroups returns the ImportGroups field value if set, zero value otherwise.
func (o *MigrationApiInfo) GetImportGroups() bool {
	if o == nil || IsNil(o.ImportGroups) {
		var ret bool
		return ret
	}
	return *o.ImportGroups
}

// GetImportGroupsOk returns a tuple with the ImportGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationApiInfo) GetImportGroupsOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportGroups) {
		return nil, false
	}
	return o.ImportGroups, true
}

// HasImportGroups returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsImportGroupsSet() bool {
	if o != nil && !IsNil(o.ImportGroups) {
		return true
	}

	return false
}

// SetImportGroups gets a reference to the given bool and assigns it to the ImportGroups field.
func (o *MigrationApiInfo) SetImportGroups(v bool) {
	o.ImportGroups = &v
}

// GetSuccessedUsers returns the SuccessedUsers field value if set, zero value otherwise.
func (o *MigrationApiInfo) GetSuccessedUsers() int32 {
	if o == nil || IsNil(o.SuccessedUsers) {
		var ret int32
		return ret
	}
	return *o.SuccessedUsers
}

// GetSuccessedUsersOk returns a tuple with the SuccessedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationApiInfo) GetSuccessedUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.SuccessedUsers) {
		return nil, false
	}
	return o.SuccessedUsers, true
}

// HasSuccessedUsers returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsSuccessedUsersSet() bool {
	if o != nil && !IsNil(o.SuccessedUsers) {
		return true
	}

	return false
}

// SetSuccessedUsers gets a reference to the given int32 and assigns it to the SuccessedUsers field.
func (o *MigrationApiInfo) SetSuccessedUsers(v int32) {
	o.SuccessedUsers = &v
}

// GetFailedUsers returns the FailedUsers field value if set, zero value otherwise.
func (o *MigrationApiInfo) GetFailedUsers() int32 {
	if o == nil || IsNil(o.FailedUsers) {
		var ret int32
		return ret
	}
	return *o.FailedUsers
}

// GetFailedUsersOk returns a tuple with the FailedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationApiInfo) GetFailedUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.FailedUsers) {
		return nil, false
	}
	return o.FailedUsers, true
}

// HasFailedUsers returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsFailedUsersSet() bool {
	if o != nil && !IsNil(o.FailedUsers) {
		return true
	}

	return false
}

// SetFailedUsers gets a reference to the given int32 and assigns it to the FailedUsers field.
func (o *MigrationApiInfo) SetFailedUsers(v int32) {
	o.FailedUsers = &v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationApiInfo) GetFiles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationApiInfo) GetFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []string and assigns it to the Files field.
func (o *MigrationApiInfo) SetFiles(v []string) {
	o.Files = v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationApiInfo) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationApiInfo) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *MigrationApiInfo) IsErrorsSet() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *MigrationApiInfo) SetErrors(v []string) {
	o.Errors = v
}

func (o MigrationApiInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrationApiInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MigratorName.IsSet() {
		toSerialize["migratorName"] = o.MigratorName.Get()
	}
	if o.Operation.IsSet() {
		toSerialize["operation"] = o.Operation.Get()
	}
	if o.FailedArchives != nil {
		toSerialize["failedArchives"] = o.FailedArchives
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.WithoutEmailUsers != nil {
		toSerialize["withoutEmailUsers"] = o.WithoutEmailUsers
	}
	if o.ExistUsers != nil {
		toSerialize["existUsers"] = o.ExistUsers
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.ImportPersonalFiles) {
		toSerialize["importPersonalFiles"] = o.ImportPersonalFiles
	}
	if !IsNil(o.ImportSharedFiles) {
		toSerialize["importSharedFiles"] = o.ImportSharedFiles
	}
	if !IsNil(o.ImportSharedFolders) {
		toSerialize["importSharedFolders"] = o.ImportSharedFolders
	}
	if !IsNil(o.ImportCommonFiles) {
		toSerialize["importCommonFiles"] = o.ImportCommonFiles
	}
	if !IsNil(o.ImportProjectFiles) {
		toSerialize["importProjectFiles"] = o.ImportProjectFiles
	}
	if !IsNil(o.ImportGroups) {
		toSerialize["importGroups"] = o.ImportGroups
	}
	if !IsNil(o.SuccessedUsers) {
		toSerialize["successedUsers"] = o.SuccessedUsers
	}
	if !IsNil(o.FailedUsers) {
		toSerialize["failedUsers"] = o.FailedUsers
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableMigrationApiInfo struct {
	value *MigrationApiInfo
	isSet bool
}

func (v NullableMigrationApiInfo) Get() *MigrationApiInfo {
	return v.value
}

func (v *NullableMigrationApiInfo) Set(val *MigrationApiInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationApiInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationApiInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationApiInfo(val *MigrationApiInfo) *NullableMigrationApiInfo {
	return &NullableMigrationApiInfo{value: val, isSet: true}
}

func (v NullableMigrationApiInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationApiInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

