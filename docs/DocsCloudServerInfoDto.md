# DocsCloudServerInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **NullableString** | The server version. | [optional] 
**PackageType** | Pointer to **NullableString** | The server package type (Open Source, Enterprise Edition or Developer Edition). | [optional] 
**Date** | Pointer to **time.Time** | The server build date. | [optional] 

## Methods

### NewDocsCloudServerInfoDto

`func NewDocsCloudServerInfoDto() *DocsCloudServerInfoDto`

NewDocsCloudServerInfoDto instantiates a new DocsCloudServerInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudServerInfoDtoWithDefaults

`func NewDocsCloudServerInfoDtoWithDefaults() *DocsCloudServerInfoDto`

NewDocsCloudServerInfoDtoWithDefaults instantiates a new DocsCloudServerInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DocsCloudServerInfoDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DocsCloudServerInfoDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DocsCloudServerInfoDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DocsCloudServerInfoDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DocsCloudServerInfoDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DocsCloudServerInfoDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetPackageType

`func (o *DocsCloudServerInfoDto) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *DocsCloudServerInfoDto) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *DocsCloudServerInfoDto) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *DocsCloudServerInfoDto) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### SetPackageTypeNil

`func (o *DocsCloudServerInfoDto) SetPackageTypeNil(b bool)`

 SetPackageTypeNil sets the value for PackageType to be an explicit nil

### UnsetPackageType
`func (o *DocsCloudServerInfoDto) UnsetPackageType()`

UnsetPackageType ensures that no value is present for PackageType, not even an explicit nil
### GetDate

`func (o *DocsCloudServerInfoDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DocsCloudServerInfoDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DocsCloudServerInfoDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *DocsCloudServerInfoDto) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


