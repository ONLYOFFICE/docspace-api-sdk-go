# MigratingApiUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldImport** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**UserType** | Pointer to [**EmployeeType**](EmployeeType.md) |  | [optional] 
**MigratingFiles** | Pointer to [**MigratingApiFiles**](MigratingApiFiles.md) |  | [optional] 

## Methods

### NewMigratingApiUser

`func NewMigratingApiUser() *MigratingApiUser`

NewMigratingApiUser instantiates a new MigratingApiUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigratingApiUserWithDefaults

`func NewMigratingApiUserWithDefaults() *MigratingApiUser`

NewMigratingApiUserWithDefaults instantiates a new MigratingApiUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldImport

`func (o *MigratingApiUser) GetShouldImport() bool`

GetShouldImport returns the ShouldImport field if non-nil, zero value otherwise.

### GetShouldImportOk

`func (o *MigratingApiUser) GetShouldImportOk() (*bool, bool)`

GetShouldImportOk returns a tuple with the ShouldImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldImport

`func (o *MigratingApiUser) SetShouldImport(v bool)`

SetShouldImport sets ShouldImport field to given value.

### HasShouldImport

`func (o *MigratingApiUser) HasShouldImport() bool`

HasShouldImport returns a boolean if a field has been set.

### GetKey

`func (o *MigratingApiUser) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MigratingApiUser) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MigratingApiUser) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MigratingApiUser) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *MigratingApiUser) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *MigratingApiUser) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetEmail

`func (o *MigratingApiUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MigratingApiUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MigratingApiUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MigratingApiUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MigratingApiUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MigratingApiUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDisplayName

`func (o *MigratingApiUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MigratingApiUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MigratingApiUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MigratingApiUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MigratingApiUser) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MigratingApiUser) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFirstName

`func (o *MigratingApiUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MigratingApiUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MigratingApiUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MigratingApiUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *MigratingApiUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *MigratingApiUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *MigratingApiUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MigratingApiUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MigratingApiUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MigratingApiUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *MigratingApiUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *MigratingApiUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetUserType

`func (o *MigratingApiUser) GetUserType() EmployeeType`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *MigratingApiUser) GetUserTypeOk() (*EmployeeType, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *MigratingApiUser) SetUserType(v EmployeeType)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *MigratingApiUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetMigratingFiles

`func (o *MigratingApiUser) GetMigratingFiles() MigratingApiFiles`

GetMigratingFiles returns the MigratingFiles field if non-nil, zero value otherwise.

### GetMigratingFilesOk

`func (o *MigratingApiUser) GetMigratingFilesOk() (*MigratingApiFiles, bool)`

GetMigratingFilesOk returns a tuple with the MigratingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingFiles

`func (o *MigratingApiUser) SetMigratingFiles(v MigratingApiFiles)`

SetMigratingFiles sets MigratingFiles field to given value.

### HasMigratingFiles

`func (o *MigratingApiUser) HasMigratingFiles() bool`

HasMigratingFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


