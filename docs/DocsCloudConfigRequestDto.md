# DocsCloudConfigRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantName** | Pointer to **NullableString** | The tenant name. | [optional] 
**Security** | Pointer to [**DocsCloudSecurityConfigRequest**](DocsCloudSecurityConfigRequest.md) | The security configuration. | [optional] 
**Server** | Pointer to [**DocsCloudServerConfigRequest**](DocsCloudServerConfigRequest.md) | The server configuration. | [optional] 
**Wopi** | Pointer to [**DocsCloudWopiConfigRequest**](DocsCloudWopiConfigRequest.md) | The WOPI configuration. | [optional] 
**IpFilter** | Pointer to [**DocsCloudIpFilterConfigRequest**](DocsCloudIpFilterConfigRequest.md) | The IP filter configuration. | [optional] 

## Methods

### NewDocsCloudConfigRequestDto

`func NewDocsCloudConfigRequestDto() *DocsCloudConfigRequestDto`

NewDocsCloudConfigRequestDto instantiates a new DocsCloudConfigRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudConfigRequestDtoWithDefaults

`func NewDocsCloudConfigRequestDtoWithDefaults() *DocsCloudConfigRequestDto`

NewDocsCloudConfigRequestDtoWithDefaults instantiates a new DocsCloudConfigRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantName

`func (o *DocsCloudConfigRequestDto) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *DocsCloudConfigRequestDto) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *DocsCloudConfigRequestDto) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *DocsCloudConfigRequestDto) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *DocsCloudConfigRequestDto) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *DocsCloudConfigRequestDto) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
### GetSecurity

`func (o *DocsCloudConfigRequestDto) GetSecurity() DocsCloudSecurityConfigRequest`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DocsCloudConfigRequestDto) GetSecurityOk() (*DocsCloudSecurityConfigRequest, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DocsCloudConfigRequestDto) SetSecurity(v DocsCloudSecurityConfigRequest)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DocsCloudConfigRequestDto) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetServer

`func (o *DocsCloudConfigRequestDto) GetServer() DocsCloudServerConfigRequest`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DocsCloudConfigRequestDto) GetServerOk() (*DocsCloudServerConfigRequest, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DocsCloudConfigRequestDto) SetServer(v DocsCloudServerConfigRequest)`

SetServer sets Server field to given value.

### HasServer

`func (o *DocsCloudConfigRequestDto) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetWopi

`func (o *DocsCloudConfigRequestDto) GetWopi() DocsCloudWopiConfigRequest`

GetWopi returns the Wopi field if non-nil, zero value otherwise.

### GetWopiOk

`func (o *DocsCloudConfigRequestDto) GetWopiOk() (*DocsCloudWopiConfigRequest, bool)`

GetWopiOk returns a tuple with the Wopi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWopi

`func (o *DocsCloudConfigRequestDto) SetWopi(v DocsCloudWopiConfigRequest)`

SetWopi sets Wopi field to given value.

### HasWopi

`func (o *DocsCloudConfigRequestDto) HasWopi() bool`

HasWopi returns a boolean if a field has been set.

### GetIpFilter

`func (o *DocsCloudConfigRequestDto) GetIpFilter() DocsCloudIpFilterConfigRequest`

GetIpFilter returns the IpFilter field if non-nil, zero value otherwise.

### GetIpFilterOk

`func (o *DocsCloudConfigRequestDto) GetIpFilterOk() (*DocsCloudIpFilterConfigRequest, bool)`

GetIpFilterOk returns a tuple with the IpFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFilter

`func (o *DocsCloudConfigRequestDto) SetIpFilter(v DocsCloudIpFilterConfigRequest)`

SetIpFilter sets IpFilter field to given value.

### HasIpFilter

`func (o *DocsCloudConfigRequestDto) HasIpFilter() bool`

HasIpFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


