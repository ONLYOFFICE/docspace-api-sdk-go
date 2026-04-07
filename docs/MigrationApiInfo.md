# MigrationApiInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigratorName** | Pointer to **NullableString** |  | [optional] 
**Operation** | Pointer to **NullableString** |  | [optional] 
**FailedArchives** | Pointer to **[]string** |  | [optional] 
**Users** | Pointer to [**[]MigratingApiUser**](MigratingApiUser.md) |  | [optional] 
**WithoutEmailUsers** | Pointer to [**[]MigratingApiUser**](MigratingApiUser.md) |  | [optional] 
**ExistUsers** | Pointer to [**[]MigratingApiUser**](MigratingApiUser.md) |  | [optional] 
**Groups** | Pointer to [**[]MigratingApiGroup**](MigratingApiGroup.md) |  | [optional] 
**ImportPersonalFiles** | Pointer to **bool** |  | [optional] 
**ImportSharedFiles** | Pointer to **bool** |  | [optional] 
**ImportSharedFolders** | Pointer to **bool** |  | [optional] 
**ImportCommonFiles** | Pointer to **bool** |  | [optional] 
**ImportProjectFiles** | Pointer to **bool** |  | [optional] 
**ImportGroups** | Pointer to **bool** |  | [optional] 
**SuccessedUsers** | Pointer to **int32** |  | [optional] 
**FailedUsers** | Pointer to **int32** |  | [optional] 
**Files** | Pointer to **[]string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMigrationApiInfo

`func NewMigrationApiInfo() *MigrationApiInfo`

NewMigrationApiInfo instantiates a new MigrationApiInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationApiInfoWithDefaults

`func NewMigrationApiInfoWithDefaults() *MigrationApiInfo`

NewMigrationApiInfoWithDefaults instantiates a new MigrationApiInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigratorName

`func (o *MigrationApiInfo) GetMigratorName() string`

GetMigratorName returns the MigratorName field if non-nil, zero value otherwise.

### GetMigratorNameOk

`func (o *MigrationApiInfo) GetMigratorNameOk() (*string, bool)`

GetMigratorNameOk returns a tuple with the MigratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratorName

`func (o *MigrationApiInfo) SetMigratorName(v string)`

SetMigratorName sets MigratorName field to given value.

### HasMigratorName

`func (o *MigrationApiInfo) HasMigratorName() bool`

HasMigratorName returns a boolean if a field has been set.

### SetMigratorNameNil

`func (o *MigrationApiInfo) SetMigratorNameNil(b bool)`

 SetMigratorNameNil sets the value for MigratorName to be an explicit nil

### UnsetMigratorName
`func (o *MigrationApiInfo) UnsetMigratorName()`

UnsetMigratorName ensures that no value is present for MigratorName, not even an explicit nil
### GetOperation

`func (o *MigrationApiInfo) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MigrationApiInfo) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *MigrationApiInfo) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *MigrationApiInfo) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *MigrationApiInfo) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *MigrationApiInfo) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetFailedArchives

`func (o *MigrationApiInfo) GetFailedArchives() []string`

GetFailedArchives returns the FailedArchives field if non-nil, zero value otherwise.

### GetFailedArchivesOk

`func (o *MigrationApiInfo) GetFailedArchivesOk() (*[]string, bool)`

GetFailedArchivesOk returns a tuple with the FailedArchives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedArchives

`func (o *MigrationApiInfo) SetFailedArchives(v []string)`

SetFailedArchives sets FailedArchives field to given value.

### HasFailedArchives

`func (o *MigrationApiInfo) HasFailedArchives() bool`

HasFailedArchives returns a boolean if a field has been set.

### SetFailedArchivesNil

`func (o *MigrationApiInfo) SetFailedArchivesNil(b bool)`

 SetFailedArchivesNil sets the value for FailedArchives to be an explicit nil

### UnsetFailedArchives
`func (o *MigrationApiInfo) UnsetFailedArchives()`

UnsetFailedArchives ensures that no value is present for FailedArchives, not even an explicit nil
### GetUsers

`func (o *MigrationApiInfo) GetUsers() []MigratingApiUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MigrationApiInfo) GetUsersOk() (*[]MigratingApiUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *MigrationApiInfo) SetUsers(v []MigratingApiUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *MigrationApiInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *MigrationApiInfo) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *MigrationApiInfo) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetWithoutEmailUsers

`func (o *MigrationApiInfo) GetWithoutEmailUsers() []MigratingApiUser`

GetWithoutEmailUsers returns the WithoutEmailUsers field if non-nil, zero value otherwise.

### GetWithoutEmailUsersOk

`func (o *MigrationApiInfo) GetWithoutEmailUsersOk() (*[]MigratingApiUser, bool)`

GetWithoutEmailUsersOk returns a tuple with the WithoutEmailUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutEmailUsers

`func (o *MigrationApiInfo) SetWithoutEmailUsers(v []MigratingApiUser)`

SetWithoutEmailUsers sets WithoutEmailUsers field to given value.

### HasWithoutEmailUsers

`func (o *MigrationApiInfo) HasWithoutEmailUsers() bool`

HasWithoutEmailUsers returns a boolean if a field has been set.

### SetWithoutEmailUsersNil

`func (o *MigrationApiInfo) SetWithoutEmailUsersNil(b bool)`

 SetWithoutEmailUsersNil sets the value for WithoutEmailUsers to be an explicit nil

### UnsetWithoutEmailUsers
`func (o *MigrationApiInfo) UnsetWithoutEmailUsers()`

UnsetWithoutEmailUsers ensures that no value is present for WithoutEmailUsers, not even an explicit nil
### GetExistUsers

`func (o *MigrationApiInfo) GetExistUsers() []MigratingApiUser`

GetExistUsers returns the ExistUsers field if non-nil, zero value otherwise.

### GetExistUsersOk

`func (o *MigrationApiInfo) GetExistUsersOk() (*[]MigratingApiUser, bool)`

GetExistUsersOk returns a tuple with the ExistUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistUsers

`func (o *MigrationApiInfo) SetExistUsers(v []MigratingApiUser)`

SetExistUsers sets ExistUsers field to given value.

### HasExistUsers

`func (o *MigrationApiInfo) HasExistUsers() bool`

HasExistUsers returns a boolean if a field has been set.

### SetExistUsersNil

`func (o *MigrationApiInfo) SetExistUsersNil(b bool)`

 SetExistUsersNil sets the value for ExistUsers to be an explicit nil

### UnsetExistUsers
`func (o *MigrationApiInfo) UnsetExistUsers()`

UnsetExistUsers ensures that no value is present for ExistUsers, not even an explicit nil
### GetGroups

`func (o *MigrationApiInfo) GetGroups() []MigratingApiGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MigrationApiInfo) GetGroupsOk() (*[]MigratingApiGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MigrationApiInfo) SetGroups(v []MigratingApiGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *MigrationApiInfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *MigrationApiInfo) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *MigrationApiInfo) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetImportPersonalFiles

`func (o *MigrationApiInfo) GetImportPersonalFiles() bool`

GetImportPersonalFiles returns the ImportPersonalFiles field if non-nil, zero value otherwise.

### GetImportPersonalFilesOk

`func (o *MigrationApiInfo) GetImportPersonalFilesOk() (*bool, bool)`

GetImportPersonalFilesOk returns a tuple with the ImportPersonalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPersonalFiles

`func (o *MigrationApiInfo) SetImportPersonalFiles(v bool)`

SetImportPersonalFiles sets ImportPersonalFiles field to given value.

### HasImportPersonalFiles

`func (o *MigrationApiInfo) HasImportPersonalFiles() bool`

HasImportPersonalFiles returns a boolean if a field has been set.

### GetImportSharedFiles

`func (o *MigrationApiInfo) GetImportSharedFiles() bool`

GetImportSharedFiles returns the ImportSharedFiles field if non-nil, zero value otherwise.

### GetImportSharedFilesOk

`func (o *MigrationApiInfo) GetImportSharedFilesOk() (*bool, bool)`

GetImportSharedFilesOk returns a tuple with the ImportSharedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSharedFiles

`func (o *MigrationApiInfo) SetImportSharedFiles(v bool)`

SetImportSharedFiles sets ImportSharedFiles field to given value.

### HasImportSharedFiles

`func (o *MigrationApiInfo) HasImportSharedFiles() bool`

HasImportSharedFiles returns a boolean if a field has been set.

### GetImportSharedFolders

`func (o *MigrationApiInfo) GetImportSharedFolders() bool`

GetImportSharedFolders returns the ImportSharedFolders field if non-nil, zero value otherwise.

### GetImportSharedFoldersOk

`func (o *MigrationApiInfo) GetImportSharedFoldersOk() (*bool, bool)`

GetImportSharedFoldersOk returns a tuple with the ImportSharedFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSharedFolders

`func (o *MigrationApiInfo) SetImportSharedFolders(v bool)`

SetImportSharedFolders sets ImportSharedFolders field to given value.

### HasImportSharedFolders

`func (o *MigrationApiInfo) HasImportSharedFolders() bool`

HasImportSharedFolders returns a boolean if a field has been set.

### GetImportCommonFiles

`func (o *MigrationApiInfo) GetImportCommonFiles() bool`

GetImportCommonFiles returns the ImportCommonFiles field if non-nil, zero value otherwise.

### GetImportCommonFilesOk

`func (o *MigrationApiInfo) GetImportCommonFilesOk() (*bool, bool)`

GetImportCommonFilesOk returns a tuple with the ImportCommonFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportCommonFiles

`func (o *MigrationApiInfo) SetImportCommonFiles(v bool)`

SetImportCommonFiles sets ImportCommonFiles field to given value.

### HasImportCommonFiles

`func (o *MigrationApiInfo) HasImportCommonFiles() bool`

HasImportCommonFiles returns a boolean if a field has been set.

### GetImportProjectFiles

`func (o *MigrationApiInfo) GetImportProjectFiles() bool`

GetImportProjectFiles returns the ImportProjectFiles field if non-nil, zero value otherwise.

### GetImportProjectFilesOk

`func (o *MigrationApiInfo) GetImportProjectFilesOk() (*bool, bool)`

GetImportProjectFilesOk returns a tuple with the ImportProjectFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportProjectFiles

`func (o *MigrationApiInfo) SetImportProjectFiles(v bool)`

SetImportProjectFiles sets ImportProjectFiles field to given value.

### HasImportProjectFiles

`func (o *MigrationApiInfo) HasImportProjectFiles() bool`

HasImportProjectFiles returns a boolean if a field has been set.

### GetImportGroups

`func (o *MigrationApiInfo) GetImportGroups() bool`

GetImportGroups returns the ImportGroups field if non-nil, zero value otherwise.

### GetImportGroupsOk

`func (o *MigrationApiInfo) GetImportGroupsOk() (*bool, bool)`

GetImportGroupsOk returns a tuple with the ImportGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportGroups

`func (o *MigrationApiInfo) SetImportGroups(v bool)`

SetImportGroups sets ImportGroups field to given value.

### HasImportGroups

`func (o *MigrationApiInfo) HasImportGroups() bool`

HasImportGroups returns a boolean if a field has been set.

### GetSuccessedUsers

`func (o *MigrationApiInfo) GetSuccessedUsers() int32`

GetSuccessedUsers returns the SuccessedUsers field if non-nil, zero value otherwise.

### GetSuccessedUsersOk

`func (o *MigrationApiInfo) GetSuccessedUsersOk() (*int32, bool)`

GetSuccessedUsersOk returns a tuple with the SuccessedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessedUsers

`func (o *MigrationApiInfo) SetSuccessedUsers(v int32)`

SetSuccessedUsers sets SuccessedUsers field to given value.

### HasSuccessedUsers

`func (o *MigrationApiInfo) HasSuccessedUsers() bool`

HasSuccessedUsers returns a boolean if a field has been set.

### GetFailedUsers

`func (o *MigrationApiInfo) GetFailedUsers() int32`

GetFailedUsers returns the FailedUsers field if non-nil, zero value otherwise.

### GetFailedUsersOk

`func (o *MigrationApiInfo) GetFailedUsersOk() (*int32, bool)`

GetFailedUsersOk returns a tuple with the FailedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedUsers

`func (o *MigrationApiInfo) SetFailedUsers(v int32)`

SetFailedUsers sets FailedUsers field to given value.

### HasFailedUsers

`func (o *MigrationApiInfo) HasFailedUsers() bool`

HasFailedUsers returns a boolean if a field has been set.

### GetFiles

`func (o *MigrationApiInfo) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *MigrationApiInfo) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *MigrationApiInfo) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *MigrationApiInfo) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *MigrationApiInfo) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *MigrationApiInfo) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetErrors

`func (o *MigrationApiInfo) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MigrationApiInfo) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MigrationApiInfo) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MigrationApiInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *MigrationApiInfo) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *MigrationApiInfo) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


