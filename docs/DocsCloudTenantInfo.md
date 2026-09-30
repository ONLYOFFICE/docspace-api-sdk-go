# DocsCloudTenantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | Pointer to [**DocsCloudLicenseInfo**](DocsCloudLicenseInfo.md) | The license information. | [optional] 
**Server** | Pointer to [**DocsCloudServerInfo**](DocsCloudServerInfo.md) | The Docs Connect server information. | [optional] 
**UsersLimit** | Pointer to [**DocsCloudUsersLimit**](DocsCloudUsersLimit.md) | The user limits of the license. | [optional] 
**Stats** | Pointer to [**DocsCloudStats**](DocsCloudStats.md) | The usage statistics for the current period. | [optional] 

## Methods

### NewDocsCloudTenantInfo

`func NewDocsCloudTenantInfo() *DocsCloudTenantInfo`

NewDocsCloudTenantInfo instantiates a new DocsCloudTenantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudTenantInfoWithDefaults

`func NewDocsCloudTenantInfoWithDefaults() *DocsCloudTenantInfo`

NewDocsCloudTenantInfoWithDefaults instantiates a new DocsCloudTenantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *DocsCloudTenantInfo) GetLicense() DocsCloudLicenseInfo`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *DocsCloudTenantInfo) GetLicenseOk() (*DocsCloudLicenseInfo, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *DocsCloudTenantInfo) SetLicense(v DocsCloudLicenseInfo)`

SetLicense sets License field to given value.

### HasLicense

`func (o *DocsCloudTenantInfo) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetServer

`func (o *DocsCloudTenantInfo) GetServer() DocsCloudServerInfo`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DocsCloudTenantInfo) GetServerOk() (*DocsCloudServerInfo, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DocsCloudTenantInfo) SetServer(v DocsCloudServerInfo)`

SetServer sets Server field to given value.

### HasServer

`func (o *DocsCloudTenantInfo) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUsersLimit

`func (o *DocsCloudTenantInfo) GetUsersLimit() DocsCloudUsersLimit`

GetUsersLimit returns the UsersLimit field if non-nil, zero value otherwise.

### GetUsersLimitOk

`func (o *DocsCloudTenantInfo) GetUsersLimitOk() (*DocsCloudUsersLimit, bool)`

GetUsersLimitOk returns a tuple with the UsersLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersLimit

`func (o *DocsCloudTenantInfo) SetUsersLimit(v DocsCloudUsersLimit)`

SetUsersLimit sets UsersLimit field to given value.

### HasUsersLimit

`func (o *DocsCloudTenantInfo) HasUsersLimit() bool`

HasUsersLimit returns a boolean if a field has been set.

### GetStats

`func (o *DocsCloudTenantInfo) GetStats() DocsCloudStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *DocsCloudTenantInfo) GetStatsOk() (*DocsCloudStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *DocsCloudTenantInfo) SetStats(v DocsCloudStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *DocsCloudTenantInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


