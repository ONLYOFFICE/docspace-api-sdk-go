# DocsCloudTenantInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | Pointer to [**DocsCloudLicenseInfoDto**](DocsCloudLicenseInfoDto.md) | The license information. | [optional] 
**Server** | Pointer to [**DocsCloudServerInfoDto**](DocsCloudServerInfoDto.md) | The Docs Connect server information. | [optional] 
**UsersLimit** | Pointer to [**DocsCloudUsersLimitDto**](DocsCloudUsersLimitDto.md) | The user limits of the license. | [optional] 
**Stats** | Pointer to [**DocsCloudStatsDto**](DocsCloudStatsDto.md) | The usage statistics for the current period. | [optional] 

## Methods

### NewDocsCloudTenantInfoDto

`func NewDocsCloudTenantInfoDto() *DocsCloudTenantInfoDto`

NewDocsCloudTenantInfoDto instantiates a new DocsCloudTenantInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudTenantInfoDtoWithDefaults

`func NewDocsCloudTenantInfoDtoWithDefaults() *DocsCloudTenantInfoDto`

NewDocsCloudTenantInfoDtoWithDefaults instantiates a new DocsCloudTenantInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *DocsCloudTenantInfoDto) GetLicense() DocsCloudLicenseInfoDto`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *DocsCloudTenantInfoDto) GetLicenseOk() (*DocsCloudLicenseInfoDto, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *DocsCloudTenantInfoDto) SetLicense(v DocsCloudLicenseInfoDto)`

SetLicense sets License field to given value.

### HasLicense

`func (o *DocsCloudTenantInfoDto) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetServer

`func (o *DocsCloudTenantInfoDto) GetServer() DocsCloudServerInfoDto`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DocsCloudTenantInfoDto) GetServerOk() (*DocsCloudServerInfoDto, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DocsCloudTenantInfoDto) SetServer(v DocsCloudServerInfoDto)`

SetServer sets Server field to given value.

### HasServer

`func (o *DocsCloudTenantInfoDto) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUsersLimit

`func (o *DocsCloudTenantInfoDto) GetUsersLimit() DocsCloudUsersLimitDto`

GetUsersLimit returns the UsersLimit field if non-nil, zero value otherwise.

### GetUsersLimitOk

`func (o *DocsCloudTenantInfoDto) GetUsersLimitOk() (*DocsCloudUsersLimitDto, bool)`

GetUsersLimitOk returns a tuple with the UsersLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersLimit

`func (o *DocsCloudTenantInfoDto) SetUsersLimit(v DocsCloudUsersLimitDto)`

SetUsersLimit sets UsersLimit field to given value.

### HasUsersLimit

`func (o *DocsCloudTenantInfoDto) HasUsersLimit() bool`

HasUsersLimit returns a boolean if a field has been set.

### GetStats

`func (o *DocsCloudTenantInfoDto) GetStats() DocsCloudStatsDto`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *DocsCloudTenantInfoDto) GetStatsOk() (*DocsCloudStatsDto, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *DocsCloudTenantInfoDto) SetStats(v DocsCloudStatsDto)`

SetStats sets Stats field to given value.

### HasStats

`func (o *DocsCloudTenantInfoDto) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


