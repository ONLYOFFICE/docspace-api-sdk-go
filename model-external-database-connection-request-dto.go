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

// checks if the ExternalDatabaseConnectionRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalDatabaseConnectionRequestDto{}

// ExternalDatabaseConnectionRequestDto The connection parameters of an external database to test.
type ExternalDatabaseConnectionRequestDto struct {
	// The engine of the external database: `mysql`, or `sqlite` on a Standalone installation.
	DatabaseType NullableString `json:"databaseType,omitempty"`
	// The engine named by `databaseType`, as the server reads it. A value sent here is ignored.
	DatabaseTypeEnum *ExternalDatabaseType `json:"databaseTypeEnum,omitempty"`
	// The host name or the IP address of the database server.
	DbHost NullableString `json:"dbHost,omitempty"`
	// The port the database server listens on.
	DbPort *int32 `json:"dbPort,omitempty"`
	// The name of the database to connect to.
	DbName NullableString `json:"dbName,omitempty"`
	// The user name to connect with.
	DbUser NullableString `json:"dbUser,omitempty"`
	// The password to connect with.
	DbPassword NullableString `json:"dbPassword,omitempty"`
	// Whether the connection is secured with SSL.
	DbSsl *bool `json:"dbSsl,omitempty"`
	// The path to the database file, used by the SQLite engine only.
	SqliteFilePath NullableString `json:"sqliteFilePath,omitempty"`
}

// NewExternalDatabaseConnectionRequestDto instantiates a new ExternalDatabaseConnectionRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalDatabaseConnectionRequestDto() *ExternalDatabaseConnectionRequestDto {
	this := ExternalDatabaseConnectionRequestDto{}
	return &this
}

// NewExternalDatabaseConnectionRequestDtoWithDefaults instantiates a new ExternalDatabaseConnectionRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalDatabaseConnectionRequestDtoWithDefaults() *ExternalDatabaseConnectionRequestDto {
	this := ExternalDatabaseConnectionRequestDto{}
	return &this
}

// GetDatabaseType returns the DatabaseType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalDatabaseConnectionRequestDto) GetDatabaseType() string {
	if o == nil || IsNil(o.DatabaseType.Get()) {
		var ret string
		return ret
	}
	return *o.DatabaseType.Get()
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDatabaseConnectionRequestDto) GetDatabaseTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatabaseType.Get(), o.DatabaseType.IsSet()
}

// HasDatabaseType returns a boolean if a field has been set.
func (o *ExternalDatabaseConnectionRequestDto) IsDatabaseTypeSet() bool {
	if o != nil && o.DatabaseType.IsSet() {
		return true
	}

	return false
}

// SetDatabaseType gets a reference to the given NullableString and assigns it to the DatabaseType field.
func (o *ExternalDatabaseConnectionRequestDto) SetDatabaseType(v string) {
	o.DatabaseType.Set(&v)
}
// SetDatabaseTypeNil sets the value for DatabaseType to be an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) SetDatabaseTypeNil() {
	o.DatabaseType.Set(nil)
}

// UnsetDatabaseType ensures that no value is present for DatabaseType, not even an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) UnsetDatabaseType() {
	o.DatabaseType.Unset()
}

// GetDatabaseTypeEnum returns the DatabaseTypeEnum field value if set, zero value otherwise.
func (o *ExternalDatabaseConnectionRequestDto) GetDatabaseTypeEnum() ExternalDatabaseType {
	if o == nil || IsNil(o.DatabaseTypeEnum) {
		var ret ExternalDatabaseType
		return ret
	}
	return *o.DatabaseTypeEnum
}

// GetDatabaseTypeEnumOk returns a tuple with the DatabaseTypeEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalDatabaseConnectionRequestDto) GetDatabaseTypeEnumOk() (*ExternalDatabaseType, bool) {
	if o == nil || IsNil(o.DatabaseTypeEnum) {
		return nil, false
	}
	return o.DatabaseTypeEnum, true
}

// HasDatabaseTypeEnum returns a boolean if a field has been set.
func (o *ExternalDatabaseConnectionRequestDto) IsDatabaseTypeEnumSet() bool {
	if o != nil && !IsNil(o.DatabaseTypeEnum) {
		return true
	}

	return false
}

// SetDatabaseTypeEnum gets a reference to the given ExternalDatabaseType and assigns it to the DatabaseTypeEnum field.
func (o *ExternalDatabaseConnectionRequestDto) SetDatabaseTypeEnum(v ExternalDatabaseType) {
	o.DatabaseTypeEnum = &v
}

// GetDbHost returns the DbHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalDatabaseConnectionRequestDto) GetDbHost() string {
	if o == nil || IsNil(o.DbHost.Get()) {
		var ret string
		return ret
	}
	return *o.DbHost.Get()
}

// GetDbHostOk returns a tuple with the DbHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDatabaseConnectionRequestDto) GetDbHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbHost.Get(), o.DbHost.IsSet()
}

// HasDbHost returns a boolean if a field has been set.
func (o *ExternalDatabaseConnectionRequestDto) IsDbHostSet() bool {
	if o != nil && o.DbHost.IsSet() {
		return true
	}

	return false
}

// SetDbHost gets a reference to the given NullableString and assigns it to the DbHost field.
func (o *ExternalDatabaseConnectionRequestDto) SetDbHost(v string) {
	o.DbHost.Set(&v)
}
// SetDbHostNil sets the value for DbHost to be an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) SetDbHostNil() {
	o.DbHost.Set(nil)
}

// UnsetDbHost ensures that no value is present for DbHost, not even an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) UnsetDbHost() {
	o.DbHost.Unset()
}

// GetDbPort returns the DbPort field value if set, zero value otherwise.
func (o *ExternalDatabaseConnectionRequestDto) GetDbPort() int32 {
	if o == nil || IsNil(o.DbPort) {
		var ret int32
		return ret
	}
	return *o.DbPort
}

// GetDbPortOk returns a tuple with the DbPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalDatabaseConnectionRequestDto) GetDbPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DbPort) {
		return nil, false
	}
	return o.DbPort, true
}

// HasDbPort returns a boolean if a field has been set.
func (o *ExternalDatabaseConnectionRequestDto) IsDbPortSet() bool {
	if o != nil && !IsNil(o.DbPort) {
		return true
	}

	return false
}

// SetDbPort gets a reference to the given int32 and assigns it to the DbPort field.
func (o *ExternalDatabaseConnectionRequestDto) SetDbPort(v int32) {
	o.DbPort = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalDatabaseConnectionRequestDto) GetDbName() string {
	if o == nil || IsNil(o.DbName.Get()) {
		var ret string
		return ret
	}
	return *o.DbName.Get()
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDatabaseConnectionRequestDto) GetDbNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbName.Get(), o.DbName.IsSet()
}

// HasDbName returns a boolean if a field has been set.
func (o *ExternalDatabaseConnectionRequestDto) IsDbNameSet() bool {
	if o != nil && o.DbName.IsSet() {
		return true
	}

	return false
}

// SetDbName gets a reference to the given NullableString and assigns it to the DbName field.
func (o *ExternalDatabaseConnectionRequestDto) SetDbName(v string) {
	o.DbName.Set(&v)
}
// SetDbNameNil sets the value for DbName to be an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) SetDbNameNil() {
	o.DbName.Set(nil)
}

// UnsetDbName ensures that no value is present for DbName, not even an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) UnsetDbName() {
	o.DbName.Unset()
}

// GetDbUser returns the DbUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalDatabaseConnectionRequestDto) GetDbUser() string {
	if o == nil || IsNil(o.DbUser.Get()) {
		var ret string
		return ret
	}
	return *o.DbUser.Get()
}

// GetDbUserOk returns a tuple with the DbUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDatabaseConnectionRequestDto) GetDbUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbUser.Get(), o.DbUser.IsSet()
}

// HasDbUser returns a boolean if a field has been set.
func (o *ExternalDatabaseConnectionRequestDto) IsDbUserSet() bool {
	if o != nil && o.DbUser.IsSet() {
		return true
	}

	return false
}

// SetDbUser gets a reference to the given NullableString and assigns it to the DbUser field.
func (o *ExternalDatabaseConnectionRequestDto) SetDbUser(v string) {
	o.DbUser.Set(&v)
}
// SetDbUserNil sets the value for DbUser to be an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) SetDbUserNil() {
	o.DbUser.Set(nil)
}

// UnsetDbUser ensures that no value is present for DbUser, not even an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) UnsetDbUser() {
	o.DbUser.Unset()
}

// GetDbPassword returns the DbPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalDatabaseConnectionRequestDto) GetDbPassword() string {
	if o == nil || IsNil(o.DbPassword.Get()) {
		var ret string
		return ret
	}
	return *o.DbPassword.Get()
}

// GetDbPasswordOk returns a tuple with the DbPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDatabaseConnectionRequestDto) GetDbPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbPassword.Get(), o.DbPassword.IsSet()
}

// HasDbPassword returns a boolean if a field has been set.
func (o *ExternalDatabaseConnectionRequestDto) IsDbPasswordSet() bool {
	if o != nil && o.DbPassword.IsSet() {
		return true
	}

	return false
}

// SetDbPassword gets a reference to the given NullableString and assigns it to the DbPassword field.
func (o *ExternalDatabaseConnectionRequestDto) SetDbPassword(v string) {
	o.DbPassword.Set(&v)
}
// SetDbPasswordNil sets the value for DbPassword to be an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) SetDbPasswordNil() {
	o.DbPassword.Set(nil)
}

// UnsetDbPassword ensures that no value is present for DbPassword, not even an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) UnsetDbPassword() {
	o.DbPassword.Unset()
}

// GetDbSsl returns the DbSsl field value if set, zero value otherwise.
func (o *ExternalDatabaseConnectionRequestDto) GetDbSsl() bool {
	if o == nil || IsNil(o.DbSsl) {
		var ret bool
		return ret
	}
	return *o.DbSsl
}

// GetDbSslOk returns a tuple with the DbSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalDatabaseConnectionRequestDto) GetDbSslOk() (*bool, bool) {
	if o == nil || IsNil(o.DbSsl) {
		return nil, false
	}
	return o.DbSsl, true
}

// HasDbSsl returns a boolean if a field has been set.
func (o *ExternalDatabaseConnectionRequestDto) IsDbSslSet() bool {
	if o != nil && !IsNil(o.DbSsl) {
		return true
	}

	return false
}

// SetDbSsl gets a reference to the given bool and assigns it to the DbSsl field.
func (o *ExternalDatabaseConnectionRequestDto) SetDbSsl(v bool) {
	o.DbSsl = &v
}

// GetSqliteFilePath returns the SqliteFilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalDatabaseConnectionRequestDto) GetSqliteFilePath() string {
	if o == nil || IsNil(o.SqliteFilePath.Get()) {
		var ret string
		return ret
	}
	return *o.SqliteFilePath.Get()
}

// GetSqliteFilePathOk returns a tuple with the SqliteFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalDatabaseConnectionRequestDto) GetSqliteFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SqliteFilePath.Get(), o.SqliteFilePath.IsSet()
}

// HasSqliteFilePath returns a boolean if a field has been set.
func (o *ExternalDatabaseConnectionRequestDto) IsSqliteFilePathSet() bool {
	if o != nil && o.SqliteFilePath.IsSet() {
		return true
	}

	return false
}

// SetSqliteFilePath gets a reference to the given NullableString and assigns it to the SqliteFilePath field.
func (o *ExternalDatabaseConnectionRequestDto) SetSqliteFilePath(v string) {
	o.SqliteFilePath.Set(&v)
}
// SetSqliteFilePathNil sets the value for SqliteFilePath to be an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) SetSqliteFilePathNil() {
	o.SqliteFilePath.Set(nil)
}

// UnsetSqliteFilePath ensures that no value is present for SqliteFilePath, not even an explicit nil
func (o *ExternalDatabaseConnectionRequestDto) UnsetSqliteFilePath() {
	o.SqliteFilePath.Unset()
}

func (o ExternalDatabaseConnectionRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalDatabaseConnectionRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DatabaseType.IsSet() {
		toSerialize["databaseType"] = o.DatabaseType.Get()
	}
	if !IsNil(o.DatabaseTypeEnum) {
		toSerialize["databaseTypeEnum"] = o.DatabaseTypeEnum
	}
	if o.DbHost.IsSet() {
		toSerialize["dbHost"] = o.DbHost.Get()
	}
	if !IsNil(o.DbPort) {
		toSerialize["dbPort"] = o.DbPort
	}
	if o.DbName.IsSet() {
		toSerialize["dbName"] = o.DbName.Get()
	}
	if o.DbUser.IsSet() {
		toSerialize["dbUser"] = o.DbUser.Get()
	}
	if o.DbPassword.IsSet() {
		toSerialize["dbPassword"] = o.DbPassword.Get()
	}
	if !IsNil(o.DbSsl) {
		toSerialize["dbSsl"] = o.DbSsl
	}
	if o.SqliteFilePath.IsSet() {
		toSerialize["sqliteFilePath"] = o.SqliteFilePath.Get()
	}
	return toSerialize, nil
}

type NullableExternalDatabaseConnectionRequestDto struct {
	value *ExternalDatabaseConnectionRequestDto
	isSet bool
}

func (v NullableExternalDatabaseConnectionRequestDto) Get() *ExternalDatabaseConnectionRequestDto {
	return v.value
}

func (v *NullableExternalDatabaseConnectionRequestDto) Set(val *ExternalDatabaseConnectionRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalDatabaseConnectionRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalDatabaseConnectionRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalDatabaseConnectionRequestDto(val *ExternalDatabaseConnectionRequestDto) *NullableExternalDatabaseConnectionRequestDto {
	return &NullableExternalDatabaseConnectionRequestDto{value: val, isSet: true}
}

func (v NullableExternalDatabaseConnectionRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalDatabaseConnectionRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

