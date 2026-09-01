# ExternalDatabaseSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseType** | Pointer to **NullableString** | The engine of the external database. | [optional] 
**DatabaseTypeEnum** | Pointer to [**ExternalDatabaseType**](ExternalDatabaseType.md) | The engine of an external database. | [optional] 
**DbHost** | Pointer to **NullableString** | The host name or the IP address of the database server. | [optional] 
**DbPort** | Pointer to **int32** | The port the database server listens on. | [optional] 
**DbName** | Pointer to **NullableString** | The name of the database to connect to. | [optional] 
**DbUser** | Pointer to **NullableString** | The user name to connect with. | [optional] 
**DbPassword** | Pointer to **NullableString** | The password to connect with. | [optional] 
**DbSsl** | Pointer to **bool** | Specifies whether the connection to the database is secured with SSL. | [optional] 
**SqliteFilePath** | Pointer to **NullableString** | The path to the database file, used by the SQLite engine only. | [optional] 

## Methods

### NewExternalDatabaseSettings

`func NewExternalDatabaseSettings() *ExternalDatabaseSettings`

NewExternalDatabaseSettings instantiates a new ExternalDatabaseSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalDatabaseSettingsWithDefaults

`func NewExternalDatabaseSettingsWithDefaults() *ExternalDatabaseSettings`

NewExternalDatabaseSettingsWithDefaults instantiates a new ExternalDatabaseSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseType

`func (o *ExternalDatabaseSettings) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *ExternalDatabaseSettings) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *ExternalDatabaseSettings) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *ExternalDatabaseSettings) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### SetDatabaseTypeNil

`func (o *ExternalDatabaseSettings) SetDatabaseTypeNil(b bool)`

 SetDatabaseTypeNil sets the value for DatabaseType to be an explicit nil

### UnsetDatabaseType
`func (o *ExternalDatabaseSettings) UnsetDatabaseType()`

UnsetDatabaseType ensures that no value is present for DatabaseType, not even an explicit nil
### GetDatabaseTypeEnum

`func (o *ExternalDatabaseSettings) GetDatabaseTypeEnum() ExternalDatabaseType`

GetDatabaseTypeEnum returns the DatabaseTypeEnum field if non-nil, zero value otherwise.

### GetDatabaseTypeEnumOk

`func (o *ExternalDatabaseSettings) GetDatabaseTypeEnumOk() (*ExternalDatabaseType, bool)`

GetDatabaseTypeEnumOk returns a tuple with the DatabaseTypeEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseTypeEnum

`func (o *ExternalDatabaseSettings) SetDatabaseTypeEnum(v ExternalDatabaseType)`

SetDatabaseTypeEnum sets DatabaseTypeEnum field to given value.

### HasDatabaseTypeEnum

`func (o *ExternalDatabaseSettings) HasDatabaseTypeEnum() bool`

HasDatabaseTypeEnum returns a boolean if a field has been set.

### GetDbHost

`func (o *ExternalDatabaseSettings) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *ExternalDatabaseSettings) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *ExternalDatabaseSettings) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.

### HasDbHost

`func (o *ExternalDatabaseSettings) HasDbHost() bool`

HasDbHost returns a boolean if a field has been set.

### SetDbHostNil

`func (o *ExternalDatabaseSettings) SetDbHostNil(b bool)`

 SetDbHostNil sets the value for DbHost to be an explicit nil

### UnsetDbHost
`func (o *ExternalDatabaseSettings) UnsetDbHost()`

UnsetDbHost ensures that no value is present for DbHost, not even an explicit nil
### GetDbPort

`func (o *ExternalDatabaseSettings) GetDbPort() int32`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *ExternalDatabaseSettings) GetDbPortOk() (*int32, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *ExternalDatabaseSettings) SetDbPort(v int32)`

SetDbPort sets DbPort field to given value.

### HasDbPort

`func (o *ExternalDatabaseSettings) HasDbPort() bool`

HasDbPort returns a boolean if a field has been set.

### GetDbName

`func (o *ExternalDatabaseSettings) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *ExternalDatabaseSettings) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *ExternalDatabaseSettings) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *ExternalDatabaseSettings) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *ExternalDatabaseSettings) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *ExternalDatabaseSettings) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetDbUser

`func (o *ExternalDatabaseSettings) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *ExternalDatabaseSettings) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *ExternalDatabaseSettings) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.

### HasDbUser

`func (o *ExternalDatabaseSettings) HasDbUser() bool`

HasDbUser returns a boolean if a field has been set.

### SetDbUserNil

`func (o *ExternalDatabaseSettings) SetDbUserNil(b bool)`

 SetDbUserNil sets the value for DbUser to be an explicit nil

### UnsetDbUser
`func (o *ExternalDatabaseSettings) UnsetDbUser()`

UnsetDbUser ensures that no value is present for DbUser, not even an explicit nil
### GetDbPassword

`func (o *ExternalDatabaseSettings) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *ExternalDatabaseSettings) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *ExternalDatabaseSettings) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.

### HasDbPassword

`func (o *ExternalDatabaseSettings) HasDbPassword() bool`

HasDbPassword returns a boolean if a field has been set.

### SetDbPasswordNil

`func (o *ExternalDatabaseSettings) SetDbPasswordNil(b bool)`

 SetDbPasswordNil sets the value for DbPassword to be an explicit nil

### UnsetDbPassword
`func (o *ExternalDatabaseSettings) UnsetDbPassword()`

UnsetDbPassword ensures that no value is present for DbPassword, not even an explicit nil
### GetDbSsl

`func (o *ExternalDatabaseSettings) GetDbSsl() bool`

GetDbSsl returns the DbSsl field if non-nil, zero value otherwise.

### GetDbSslOk

`func (o *ExternalDatabaseSettings) GetDbSslOk() (*bool, bool)`

GetDbSslOk returns a tuple with the DbSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbSsl

`func (o *ExternalDatabaseSettings) SetDbSsl(v bool)`

SetDbSsl sets DbSsl field to given value.

### HasDbSsl

`func (o *ExternalDatabaseSettings) HasDbSsl() bool`

HasDbSsl returns a boolean if a field has been set.

### GetSqliteFilePath

`func (o *ExternalDatabaseSettings) GetSqliteFilePath() string`

GetSqliteFilePath returns the SqliteFilePath field if non-nil, zero value otherwise.

### GetSqliteFilePathOk

`func (o *ExternalDatabaseSettings) GetSqliteFilePathOk() (*string, bool)`

GetSqliteFilePathOk returns a tuple with the SqliteFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqliteFilePath

`func (o *ExternalDatabaseSettings) SetSqliteFilePath(v string)`

SetSqliteFilePath sets SqliteFilePath field to given value.

### HasSqliteFilePath

`func (o *ExternalDatabaseSettings) HasSqliteFilePath() bool`

HasSqliteFilePath returns a boolean if a field has been set.

### SetSqliteFilePathNil

`func (o *ExternalDatabaseSettings) SetSqliteFilePathNil(b bool)`

 SetSqliteFilePathNil sets the value for SqliteFilePath to be an explicit nil

### UnsetSqliteFilePath
`func (o *ExternalDatabaseSettings) UnsetSqliteFilePath()`

UnsetSqliteFilePath ensures that no value is present for SqliteFilePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


