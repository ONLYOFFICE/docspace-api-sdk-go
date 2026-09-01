# DocsCloudConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantName** | Pointer to **NullableString** | The tenant name. | [optional] 
**Security** | Pointer to [**DocsCloudSecurityConfig**](DocsCloudSecurityConfig.md) | The security configuration. | [optional] 
**Server** | Pointer to [**DocsCloudServerConfig**](DocsCloudServerConfig.md) | The server configuration. | [optional] 
**Wopi** | Pointer to [**DocsCloudWopiConfig**](DocsCloudWopiConfig.md) | The WOPI configuration. | [optional] 
**IpFilter** | Pointer to [**DocsCloudIpFilterConfig**](DocsCloudIpFilterConfig.md) | The IP filter configuration. | [optional] 

## Methods

### NewDocsCloudConfig

`func NewDocsCloudConfig() *DocsCloudConfig`

NewDocsCloudConfig instantiates a new DocsCloudConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudConfigWithDefaults

`func NewDocsCloudConfigWithDefaults() *DocsCloudConfig`

NewDocsCloudConfigWithDefaults instantiates a new DocsCloudConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantName

`func (o *DocsCloudConfig) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *DocsCloudConfig) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *DocsCloudConfig) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *DocsCloudConfig) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *DocsCloudConfig) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *DocsCloudConfig) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
### GetSecurity

`func (o *DocsCloudConfig) GetSecurity() DocsCloudSecurityConfig`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DocsCloudConfig) GetSecurityOk() (*DocsCloudSecurityConfig, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DocsCloudConfig) SetSecurity(v DocsCloudSecurityConfig)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DocsCloudConfig) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetServer

`func (o *DocsCloudConfig) GetServer() DocsCloudServerConfig`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DocsCloudConfig) GetServerOk() (*DocsCloudServerConfig, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DocsCloudConfig) SetServer(v DocsCloudServerConfig)`

SetServer sets Server field to given value.

### HasServer

`func (o *DocsCloudConfig) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetWopi

`func (o *DocsCloudConfig) GetWopi() DocsCloudWopiConfig`

GetWopi returns the Wopi field if non-nil, zero value otherwise.

### GetWopiOk

`func (o *DocsCloudConfig) GetWopiOk() (*DocsCloudWopiConfig, bool)`

GetWopiOk returns a tuple with the Wopi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWopi

`func (o *DocsCloudConfig) SetWopi(v DocsCloudWopiConfig)`

SetWopi sets Wopi field to given value.

### HasWopi

`func (o *DocsCloudConfig) HasWopi() bool`

HasWopi returns a boolean if a field has been set.

### GetIpFilter

`func (o *DocsCloudConfig) GetIpFilter() DocsCloudIpFilterConfig`

GetIpFilter returns the IpFilter field if non-nil, zero value otherwise.

### GetIpFilterOk

`func (o *DocsCloudConfig) GetIpFilterOk() (*DocsCloudIpFilterConfig, bool)`

GetIpFilterOk returns a tuple with the IpFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFilter

`func (o *DocsCloudConfig) SetIpFilter(v DocsCloudIpFilterConfig)`

SetIpFilter sets IpFilter field to given value.

### HasIpFilter

`func (o *DocsCloudConfig) HasIpFilter() bool`

HasIpFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


