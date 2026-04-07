# StudioDefaultPageSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultFolderType** | Pointer to [**FolderType**](FolderType.md) |  | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewStudioDefaultPageSettings

`func NewStudioDefaultPageSettings() *StudioDefaultPageSettings`

NewStudioDefaultPageSettings instantiates a new StudioDefaultPageSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioDefaultPageSettingsWithDefaults

`func NewStudioDefaultPageSettingsWithDefaults() *StudioDefaultPageSettings`

NewStudioDefaultPageSettingsWithDefaults instantiates a new StudioDefaultPageSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultFolderType

`func (o *StudioDefaultPageSettings) GetDefaultFolderType() FolderType`

GetDefaultFolderType returns the DefaultFolderType field if non-nil, zero value otherwise.

### GetDefaultFolderTypeOk

`func (o *StudioDefaultPageSettings) GetDefaultFolderTypeOk() (*FolderType, bool)`

GetDefaultFolderTypeOk returns a tuple with the DefaultFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolderType

`func (o *StudioDefaultPageSettings) SetDefaultFolderType(v FolderType)`

SetDefaultFolderType sets DefaultFolderType field to given value.

### HasDefaultFolderType

`func (o *StudioDefaultPageSettings) HasDefaultFolderType() bool`

HasDefaultFolderType returns a boolean if a field has been set.

### GetLastModified

`func (o *StudioDefaultPageSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *StudioDefaultPageSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *StudioDefaultPageSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *StudioDefaultPageSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


