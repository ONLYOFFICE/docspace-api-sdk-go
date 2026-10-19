# DocsCloudConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantName** | Pointer to **NullableString** | The tenant name. | [optional] 
**Security** | Pointer to [**DocsCloudSecurityConfigDto**](DocsCloudSecurityConfigDto.md) | The security configuration. | [optional] 
**Server** | Pointer to [**DocsCloudServerConfigDto**](DocsCloudServerConfigDto.md) | The server configuration. | [optional] 
**Wopi** | Pointer to [**DocsCloudWopiConfigDto**](DocsCloudWopiConfigDto.md) | The WOPI configuration. | [optional] 
**IpFilter** | Pointer to [**DocsCloudIpFilterConfigDto**](DocsCloudIpFilterConfigDto.md) | The IP filter configuration. | [optional] 

## Methods

### NewDocsCloudConfigDto

`func NewDocsCloudConfigDto() *DocsCloudConfigDto`

NewDocsCloudConfigDto instantiates a new DocsCloudConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudConfigDtoWithDefaults

`func NewDocsCloudConfigDtoWithDefaults() *DocsCloudConfigDto`

NewDocsCloudConfigDtoWithDefaults instantiates a new DocsCloudConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantName

`func (o *DocsCloudConfigDto) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *DocsCloudConfigDto) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *DocsCloudConfigDto) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *DocsCloudConfigDto) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *DocsCloudConfigDto) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *DocsCloudConfigDto) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
### GetSecurity

`func (o *DocsCloudConfigDto) GetSecurity() DocsCloudSecurityConfigDto`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DocsCloudConfigDto) GetSecurityOk() (*DocsCloudSecurityConfigDto, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DocsCloudConfigDto) SetSecurity(v DocsCloudSecurityConfigDto)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DocsCloudConfigDto) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetServer

`func (o *DocsCloudConfigDto) GetServer() DocsCloudServerConfigDto`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DocsCloudConfigDto) GetServerOk() (*DocsCloudServerConfigDto, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DocsCloudConfigDto) SetServer(v DocsCloudServerConfigDto)`

SetServer sets Server field to given value.

### HasServer

`func (o *DocsCloudConfigDto) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetWopi

`func (o *DocsCloudConfigDto) GetWopi() DocsCloudWopiConfigDto`

GetWopi returns the Wopi field if non-nil, zero value otherwise.

### GetWopiOk

`func (o *DocsCloudConfigDto) GetWopiOk() (*DocsCloudWopiConfigDto, bool)`

GetWopiOk returns a tuple with the Wopi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWopi

`func (o *DocsCloudConfigDto) SetWopi(v DocsCloudWopiConfigDto)`

SetWopi sets Wopi field to given value.

### HasWopi

`func (o *DocsCloudConfigDto) HasWopi() bool`

HasWopi returns a boolean if a field has been set.

### GetIpFilter

`func (o *DocsCloudConfigDto) GetIpFilter() DocsCloudIpFilterConfigDto`

GetIpFilter returns the IpFilter field if non-nil, zero value otherwise.

### GetIpFilterOk

`func (o *DocsCloudConfigDto) GetIpFilterOk() (*DocsCloudIpFilterConfigDto, bool)`

GetIpFilterOk returns a tuple with the IpFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFilter

`func (o *DocsCloudConfigDto) SetIpFilter(v DocsCloudIpFilterConfigDto)`

SetIpFilter sets IpFilter field to given value.

### HasIpFilter

`func (o *DocsCloudConfigDto) HasIpFilter() bool`

HasIpFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


