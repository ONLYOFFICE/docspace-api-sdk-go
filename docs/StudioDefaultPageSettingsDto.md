# StudioDefaultPageSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultFolderType** | Pointer to [**FolderType**](FolderType.md) | Specifies the type of the default folder associated with the settings. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewStudioDefaultPageSettingsDto

`func NewStudioDefaultPageSettingsDto() *StudioDefaultPageSettingsDto`

NewStudioDefaultPageSettingsDto instantiates a new StudioDefaultPageSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioDefaultPageSettingsDtoWithDefaults

`func NewStudioDefaultPageSettingsDtoWithDefaults() *StudioDefaultPageSettingsDto`

NewStudioDefaultPageSettingsDtoWithDefaults instantiates a new StudioDefaultPageSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultFolderType

`func (o *StudioDefaultPageSettingsDto) GetDefaultFolderType() FolderType`

GetDefaultFolderType returns the DefaultFolderType field if non-nil, zero value otherwise.

### GetDefaultFolderTypeOk

`func (o *StudioDefaultPageSettingsDto) GetDefaultFolderTypeOk() (*FolderType, bool)`

GetDefaultFolderTypeOk returns a tuple with the DefaultFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolderType

`func (o *StudioDefaultPageSettingsDto) SetDefaultFolderType(v FolderType)`

SetDefaultFolderType sets DefaultFolderType field to given value.

### HasDefaultFolderType

`func (o *StudioDefaultPageSettingsDto) HasDefaultFolderType() bool`

HasDefaultFolderType returns a boolean if a field has been set.

### GetLastModified

`func (o *StudioDefaultPageSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *StudioDefaultPageSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *StudioDefaultPageSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *StudioDefaultPageSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


